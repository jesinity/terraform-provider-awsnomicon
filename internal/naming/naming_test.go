package naming

import "testing"

func TestDefaultCloudDefaultsAzure(t *testing.T) {
	defaults, err := DefaultCloudDefaults(CloudAzure)
	if err != nil {
		t.Fatalf("unexpected error loading Azure defaults: %v", err)
	}

	if len(defaults.ResourceAcronyms) < 300 {
		t.Fatalf("expected Azure CAF acronyms to be populated, got %d entries", len(defaults.ResourceAcronyms))
	}

	for resource, acronym := range defaults.ResourceAcronyms {
		if len(acronym) != 4 && len(acronym) != 5 {
			t.Fatalf("resource %q has acronym %q, expected 4 or 5 characters", resource, acronym)
		}
	}

	if got := defaults.ResourceAcronyms["azurerm_storage_account"]; got != "stac" {
		t.Fatalf("expected normalized storage account acronym %q, got %q", "stac", got)
	}

	if got := defaults.ResourceAcronyms["azurerm_resource_group"]; got != "regr" {
		t.Fatalf("expected normalized resource group acronym %q, got %q", "regr", got)
	}

	if got := defaults.ResourceAcronyms["azurerm_virtual_machine"]; got != "vimav" {
		t.Fatalf("expected disambiguated virtual machine acronym %q, got %q", "vimav", got)
	}

	seen := map[string]string{}
	for resource, acronym := range defaults.ResourceAcronyms {
		if previous, exists := seen[acronym]; exists {
			t.Fatalf("expected unique normalized acronyms, but %q and %q both use %q", previous, resource, acronym)
		}
		seen[acronym] = resource
	}
}

func TestDefaultCloudDefaultsAzureUseCAFAcronyms(t *testing.T) {
	defaults, err := DefaultCloudDefaultsWithOptions(CloudAzure, CloudDefaultsOptions{UseAzureCAFAcronyms: true})
	if err != nil {
		t.Fatalf("unexpected error loading Azure defaults: %v", err)
	}

	if got := defaults.ResourceAcronyms["azurerm_storage_account"]; got != "st" {
		t.Fatalf("expected CAF storage account acronym %q, got %q", "st", got)
	}

	if got := defaults.ResourceAcronyms["azurerm_resource_group"]; got != "rg" {
		t.Fatalf("expected CAF resource group acronym %q, got %q", "rg", got)
	}

	// Some entries do not provide CAF slugs; keep a non-empty compatibility fallback.
	if got := defaults.ResourceAcronyms["general"]; got != "gene" {
		t.Fatalf("expected fallback acronym %q for resource %q, got %q", "gene", "general", got)
	}
}

func TestBuildNameAzureStorageAccountSelectsStraightStyle(t *testing.T) {
	defaults, err := DefaultCloudDefaults(CloudAzure)
	if err != nil {
		t.Fatalf("unexpected error loading Azure defaults: %v", err)
	}

	result, err := BuildName(Config{
		StylePriority:          DefaultStylePriority(),
		ResourceAcronyms:       defaults.ResourceAcronyms,
		ResourceStyleOverrides: defaults.ResourceStyleOverrides,
		ResourceConstraints:    defaults.ResourceConstraints,
		RegionalResources:      defaults.RegionalResources,
	}, BuildInput{
		Resource:  "azurerm_storage_account",
		Qualifier: "data",
		Recipe:    []string{"resource", "qualifier"},
	})
	if err != nil {
		t.Fatalf("unexpected build error: %v", err)
	}

	if result.Style != StyleStraight {
		t.Fatalf("expected style %q, got %q", StyleStraight, result.Style)
	}

	if result.Name == "" {
		t.Fatal("expected a generated name")
	}
}
