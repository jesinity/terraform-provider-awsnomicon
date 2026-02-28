package naming

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"sync"
	"unicode"
)

//go:embed azure_caf_resource_definition.json
var azureCAFResourceDefinitionJSON []byte

type azureCAFResourceDefinition struct {
	Name            string `json:"name"`
	MinLength       int    `json:"min_length"`
	MaxLength       int    `json:"max_length"`
	ValidationRegex string `json:"validation_regex"`
	Scope           string `json:"scope"`
	Slug            string `json:"slug"`
	Dashes          bool   `json:"dashes"`
	Lowercase       bool   `json:"lowercase"`
}

type azureCloudProfile struct {
	normalizedDefaults azureCloudDefaultsCache
	cafDefaults        azureCloudDefaultsCache
}

type azureCloudDefaultsCache struct {
	once     sync.Once
	defaults CloudDefaults
	err      error
}

func newAzureCloudProfile() CloudProfile {
	return &azureCloudProfile{}
}

func (*azureCloudProfile) Cloud() string {
	return CloudAzure
}

func (p *azureCloudProfile) Defaults(options CloudDefaultsOptions) (CloudDefaults, error) {
	useCAFAcronyms := options.UseAzureCAFAcronyms
	cache := &p.normalizedDefaults
	if useCAFAcronyms {
		cache = &p.cafDefaults
	}

	cache.once.Do(func() {
		cache.defaults, cache.err = loadAzureCloudDefaults(useCAFAcronyms)
	})

	if cache.err != nil {
		return CloudDefaults{}, cache.err
	}
	return copyCloudDefaults(cache.defaults), nil
}

func loadAzureCloudDefaults(useCAFAcronyms bool) (CloudDefaults, error) {
	var definitions []azureCAFResourceDefinition
	if err := json.Unmarshal(azureCAFResourceDefinitionJSON, &definitions); err != nil {
		return CloudDefaults{}, fmt.Errorf("decode Azure CAF resource definitions: %w", err)
	}

	acronyms := make(map[string]string, len(definitions))
	styleOverrides := make(map[string][]string, len(definitions))
	constraints := make(map[string]ResourceConstraint, len(definitions))
	regionalResources := make(map[string]bool, len(definitions))

	for _, definition := range definitions {
		name := strings.ToLower(strings.TrimSpace(definition.Name))
		if name == "" {
			continue
		}

		acronym := azureCAFResourceAcronym(definition.Slug, definition.Name)
		if useCAFAcronyms {
			acronym = azureCAFOfficialResourceAcronym(definition.Slug, definition.Name)
		}

		acronyms[name] = acronym
		styleOverrides[name] = azureCAFStyleOverrides(definition.Lowercase, definition.Dashes)
		regionalResources[name] = azureCAFIsRegionalScope(definition.Scope)
		constraints[name] = azureCAFConstraint(definition)
	}

	if !useCAFAcronyms {
		azureCAFDisambiguateNormalizedAcronyms(definitions, acronyms)
	}

	return CloudDefaults{
		RegionMap:              map[string]string{},
		ResourceAcronyms:       acronyms,
		ResourceStyleOverrides: styleOverrides,
		ResourceConstraints:    constraints,
		RegionalResources:      regionalResources,
	}, nil
}

func azureCAFConstraint(definition azureCAFResourceDefinition) ResourceConstraint {
	constraint := ResourceConstraint{
		MinLen: definition.MinLength,
		MaxLen: definition.MaxLength,
	}

	regexValue := strings.TrimSpace(definition.ValidationRegex)
	regexValue = strings.Trim(regexValue, `"`)
	if regexValue == "" {
		return constraint
	}

	constraint.PatternDescription = fmt.Sprintf("must match Azure CAF regex %q", regexValue)
	pattern, err := regexp.Compile(regexValue)
	if err != nil {
		return constraint
	}
	constraint.Pattern = pattern
	return constraint
}

func azureCAFStyleOverrides(lowercase, dashes bool) []string {
	styles := []string{}
	if lowercase {
		if dashes {
			styles = append(styles, StyleDashed)
		}
		styles = append(styles, StyleStraight)
		return styles
	}

	if dashes {
		styles = append(styles, StyleDashed, StylePascalDashed)
	}
	styles = append(styles, StylePascal, StyleCamel, StyleStraight)
	return styles
}

func azureCAFResourceAcronym(slug, name string) string {
	base := toLowerAlnum(slug)
	fallback := toLowerAlnum(name)
	for len(base) < 4 && fallback != "" {
		base += fallback[:1]
		fallback = fallback[1:]
	}
	for len(base) < 4 {
		base += "x"
	}
	acronym := base[:4]
	if strings.HasSuffix(acronym, "az") {
		return azureCAFResourceSemanticAcronym(name)
	}
	return acronym
}

func azureCAFOfficialResourceAcronym(slug, name string) string {
	base := toLowerAlnum(slug)
	if base != "" {
		return base
	}
	// Keep compatibility for definitions that do not publish a CAF slug.
	return azureCAFResourceAcronym(slug, name)
}

func azureCAFResourceSemanticAcronym(name string) string {
	normalized := strings.ToLower(strings.TrimSpace(name))
	for _, prefix := range []string{"azurerm_", "azure_", "azapi_"} {
		if strings.HasPrefix(normalized, prefix) {
			normalized = strings.TrimPrefix(normalized, prefix)
			break
		}
	}

	tokens := []string{}
	for _, token := range strings.Split(normalized, "_") {
		token = toLowerAlnum(token)
		if token == "" {
			continue
		}
		tokens = append(tokens, token)
	}

	var base string
	switch {
	case len(tokens) >= 2:
		last := tokens[len(tokens)-1]
		prev := tokens[len(tokens)-2]
		base = firstN(prev, 2) + firstN(last, 2)
	case len(tokens) == 1:
		base = firstN(tokens[0], 4)
	default:
		base = firstN(toLowerAlnum(name), 4)
	}

	for len(base) < 4 {
		base += "x"
	}
	return base[:4]
}

func firstN(value string, n int) string {
	if n <= 0 || value == "" {
		return ""
	}
	if len(value) <= n {
		return value
	}
	return value[:n]
}

func azureCAFDisambiguateNormalizedAcronyms(definitions []azureCAFResourceDefinition, acronyms map[string]string) {
	groups := map[string][]string{}
	for _, definition := range definitions {
		name := strings.ToLower(strings.TrimSpace(definition.Name))
		if name == "" {
			continue
		}
		acronym := strings.TrimSpace(acronyms[name])
		if acronym == "" {
			continue
		}
		groups[acronym] = append(groups[acronym], name)
	}

	for base, names := range groups {
		if len(names) <= 1 {
			continue
		}

		// Stable ordering keeps outputs deterministic across runs.
		sort.Strings(names)
		resourceTokens := map[string][]string{}
		tokenFrequency := map[string]int{}
		for _, name := range names {
			tokens := azureCAFResourceTokens(name)
			resourceTokens[name] = tokens
			seen := map[string]bool{}
			for _, token := range tokens {
				if token == "" || seen[token] {
					continue
				}
				tokenFrequency[token]++
				seen[token] = true
			}
		}

		primary := azureCAFPrimaryResourceForBase(names, resourceTokens)
		used := map[string]bool{}
		acronyms[primary] = base
		used[base] = true

		type disambiguationItem struct {
			name           string
			tokens         []string
			preferredToken string
		}

		items := []disambiguationItem{}
		for _, name := range names {
			if name == primary {
				continue
			}
			tokens := resourceTokens[name]
			items = append(items, disambiguationItem{
				name:           name,
				tokens:         tokens,
				preferredToken: azureCAFPreferredDisambiguationToken(tokens, tokenFrequency),
			})
		}

		sort.Slice(items, func(i, j int) bool {
			left := items[i]
			right := items[j]
			if len(left.preferredToken) != len(right.preferredToken) {
				return len(left.preferredToken) < len(right.preferredToken)
			}
			if left.preferredToken != right.preferredToken {
				return left.preferredToken < right.preferredToken
			}
			return left.name < right.name
		})

		for _, item := range items {
			acronyms[item.name] = azureCAFDisambiguatedAcronym(base, item.tokens, tokenFrequency, used)
			used[acronyms[item.name]] = true
		}
	}
}

func azureCAFPrimaryResourceForBase(names []string, resourceTokens map[string][]string) string {
	if len(names) == 0 {
		return ""
	}

	primary := names[0]
	for _, candidate := range names[1:] {
		primaryTokens := resourceTokens[primary]
		candidateTokens := resourceTokens[candidate]
		if len(candidateTokens) < len(primaryTokens) {
			primary = candidate
			continue
		}
		if len(candidateTokens) > len(primaryTokens) {
			continue
		}
		if len(candidate) < len(primary) {
			primary = candidate
			continue
		}
		if len(candidate) > len(primary) {
			continue
		}
		if candidate < primary {
			primary = candidate
		}
	}

	return primary
}

func azureCAFDisambiguatedAcronym(base string, tokens []string, tokenFrequency map[string]int, used map[string]bool) string {
	for _, r := range azureCAFDisambiguationSuffixRunes(tokens, tokenFrequency) {
		candidate := base + string(r)
		if !used[candidate] {
			return candidate
		}
	}

	// Extremely unlikely: keep deterministic output if all 36 suffixes are exhausted.
	return base + "x"
}

func azureCAFDisambiguationSuffixRunes(tokens []string, tokenFrequency map[string]int) []rune {
	candidates := []rune{}
	seen := map[rune]bool{}
	addRune := func(r rune) {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return
		}
		r = unicode.ToLower(r)
		if seen[r] {
			return
		}
		seen[r] = true
		candidates = append(candidates, r)
	}
	addRunes := func(value string) {
		for _, r := range toLowerAlnum(value) {
			addRune(r)
		}
	}

	preferred := azureCAFPreferredDisambiguationToken(tokens, tokenFrequency)
	if preferred != "" {
		preferredRunes := []rune(toLowerAlnum(preferred))
		addRune(preferredRunes[0])
		for _, r := range preferredRunes[1:] {
			if !azureCAFIsVowelRune(r) {
				addRune(r)
			}
		}
		for _, r := range preferredRunes[1:] {
			if azureCAFIsVowelRune(r) {
				addRune(r)
			}
		}
	}

	for i := len(tokens) - 1; i >= 0; i-- {
		if tokens[i] == "" {
			continue
		}
		addRunes(tokens[i][:1])
	}

	addRunes(strings.Join(tokens, ""))
	addRunes("abcdefghijklmnopqrstuvwxyz0123456789")
	return candidates
}

func azureCAFPreferredDisambiguationToken(tokens []string, tokenFrequency map[string]int) string {
	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		if tokenFrequency[token] == 1 {
			return token
		}
	}
	if len(tokens) > 0 {
		return tokens[len(tokens)-1]
	}
	return ""
}

func azureCAFIsVowelRune(r rune) bool {
	switch unicode.ToLower(r) {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

func azureCAFResourceTokens(name string) []string {
	normalized := strings.ToLower(strings.TrimSpace(name))
	for _, prefix := range []string{"azurerm_", "azure_", "azapi_"} {
		if strings.HasPrefix(normalized, prefix) {
			normalized = strings.TrimPrefix(normalized, prefix)
			break
		}
	}

	tokens := []string{}
	for _, token := range strings.Split(normalized, "_") {
		token = toLowerAlnum(token)
		if token == "" {
			continue
		}
		tokens = append(tokens, token)
	}
	return tokens
}

func azureCAFIsRegionalScope(scope string) bool {
	switch strings.ToLower(strings.TrimSpace(scope)) {
	case "resourcegroup", "region", "location", "parent":
		return true
	default:
		return false
	}
}

func toLowerAlnum(value string) string {
	var b strings.Builder
	b.Grow(len(value))
	for _, r := range value {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}

func copyCloudDefaults(in CloudDefaults) CloudDefaults {
	return CloudDefaults{
		RegionMap:              copyStringMap(in.RegionMap),
		ResourceAcronyms:       copyStringMap(in.ResourceAcronyms),
		ResourceStyleOverrides: copyStringSliceMap(in.ResourceStyleOverrides),
		ResourceConstraints:    copyConstraintMap(in.ResourceConstraints),
		RegionalResources:      copyBoolMap(in.RegionalResources),
	}
}

func copyStringMap(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))
	for key, value := range in {
		out[key] = value
	}
	return out
}

func copyBoolMap(in map[string]bool) map[string]bool {
	out := make(map[string]bool, len(in))
	for key, value := range in {
		out[key] = value
	}
	return out
}

func copyStringSliceMap(in map[string][]string) map[string][]string {
	out := make(map[string][]string, len(in))
	for key, value := range in {
		cloned := make([]string, len(value))
		copy(cloned, value)
		out[key] = cloned
	}
	return out
}

func copyConstraintMap(in map[string]ResourceConstraint) map[string]ResourceConstraint {
	out := make(map[string]ResourceConstraint, len(in))
	for key, value := range in {
		out[key] = value
	}
	return out
}
