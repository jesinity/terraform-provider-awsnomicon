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
		if len(acronym) != 4 {
			t.Fatalf("resource %q has acronym %q, expected 4 characters", resource, acronym)
		}
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
