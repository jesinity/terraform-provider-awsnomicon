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
		used := map[string]bool{}
		for _, name := range names {
			acronyms[name] = azureCAFDisambiguatedAcronym(base, name, used)
			used[acronyms[name]] = true
		}
	}
}

func azureCAFDisambiguatedAcronym(base, name string, used map[string]bool) string {
	clean := azureCAFResourceDisambiguationSource(name)
	for _, r := range clean {
		candidate := base + string(r)
		if !used[candidate] {
			return candidate
		}
	}

	for _, r := range "abcdefghijklmnopqrstuvwxyz0123456789" {
		candidate := base + string(r)
		if !used[candidate] {
			return candidate
		}
	}

	// Extremely unlikely: keep deterministic output if all 36 suffixes are exhausted.
	return base + "x"
}

func azureCAFResourceDisambiguationSource(name string) string {
	normalized := strings.ToLower(strings.TrimSpace(name))
	for _, prefix := range []string{"azurerm_", "azure_", "azapi_"} {
		if strings.HasPrefix(normalized, prefix) {
			normalized = strings.TrimPrefix(normalized, prefix)
			break
		}
	}
	return toLowerAlnum(normalized)
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
