package naming

import "testing"

func TestDefaultCloudDefaultsAzure(t *testing.T) {
	defaults, err := DefaultCloudDefaults(CloudAzure)
	if err != nil {
		t.Fatalf("unexpected error loading Azure defaults: %v", err)
	}

	if len(defaults.RegionMap) == 0 {
		t.Fatal("expected Azure region map defaults to be populated")
	}

	if len(defaults.ResourceAcronyms) < 300 {
		t.Fatalf("expected Azure CAF acronyms to be populated, got %d entries", len(defaults.ResourceAcronyms))
	}

	for resource, acronym := range defaults.ResourceAcronyms {
		if acronym == "" {
			t.Fatalf("resource %q has empty acronym", resource)
		}
	}

	if got := defaults.ResourceAcronyms["azurerm_storage_account"]; got != "st" {
		t.Fatalf("expected CAF storage account acronym %q, got %q", "st", got)
	}

	if got := defaults.ResourceAcronyms["azurerm_resource_group"]; got != "rg" {
		t.Fatalf("expected CAF resource group acronym %q, got %q", "rg", got)
	}

	if got := defaults.ResourceAcronyms["azurerm_virtual_machine"]; got != "vm" {
		t.Fatalf("expected CAF virtual machine acronym %q, got %q", "vm", got)
	}

	if got := defaults.ResourceAcronyms["azurerm_linux_virtual_machine"]; got != "vm" {
		t.Fatalf("expected CAF linux virtual machine acronym %q, got %q", "vm", got)
	}

	if got := defaults.ResourceAcronyms["azurerm_windows_virtual_machine"]; got != "vm" {
		t.Fatalf("expected CAF windows virtual machine acronym %q, got %q", "vm", got)
	}

	if got := defaults.ResourceAcronyms["azurerm_api_management"]; got != "apim" {
		t.Fatalf("expected CAF API management acronym %q, got %q", "apim", got)
	}

	if got := defaults.ResourceAcronyms["azurerm_api_management_group"]; got != "apimgr" {
		t.Fatalf("expected CAF API management group acronym %q, got %q", "apimgr", got)
	}

	if got := defaults.ResourceAcronyms["azurerm_api_management_logger"]; got != "apimlg" {
		t.Fatalf("expected CAF API management logger acronym %q, got %q", "apimlg", got)
	}

	if got := defaults.ResourceAcronyms["azurerm_api_management_service"]; got != "apim" {
		t.Fatalf("expected CAF API management service acronym %q, got %q", "apim", got)
	}

	if _, ok := defaults.ResourceAcronyms["general"]; ok {
		t.Fatalf("expected no acronym entry for %q because CAF slug is empty", "general")
	}

	if got := defaults.RegionMap["westeurope"]; got != "weu" {
		t.Fatalf("expected Azure region short code %q for westeurope, got %q", "weu", got)
	}
	if got := defaults.RegionMap["eastus2"]; got != "eus2" {
		t.Fatalf("expected Azure region short code %q for eastus2, got %q", "eus2", got)
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

func TestBuildNameAzureRegionLookupNormalizesSeparators(t *testing.T) {
	defaults, err := DefaultCloudDefaults(CloudAzure)
	if err != nil {
		t.Fatalf("unexpected error loading Azure defaults: %v", err)
	}

	result, err := BuildName(Config{
		Cloud:                            CloudAzure,
		Region:                           "west-europe",
		RegionMap:                        defaults.RegionMap,
		IgnoreRegionForRegionalResources: false,
		ResourceAcronyms:                 defaults.ResourceAcronyms,
		ResourceStyleOverrides:           defaults.ResourceStyleOverrides,
		ResourceConstraints:              defaults.ResourceConstraints,
		RegionalResources:                defaults.RegionalResources,
	}, BuildInput{
		Resource: "azurerm_storage_account",
		Recipe:   []string{"region"},
	})
	if err != nil {
		t.Fatalf("unexpected build error: %v", err)
	}

	if result.RegionCode != "weu" {
		t.Fatalf("expected region code %q, got %q", "weu", result.RegionCode)
	}
	if result.Name != "weu" {
		t.Fatalf("expected generated name %q, got %q", "weu", result.Name)
	}
}

func TestBuildNameAzureFallsBackToAllowedStyleWhenPriorityDoesNotMatch(t *testing.T) {
	defaults, err := DefaultCloudDefaults(CloudAzure)
	if err != nil {
		t.Fatalf("unexpected error loading Azure defaults: %v", err)
	}

	result, err := BuildName(Config{
		Cloud:                            CloudAzure,
		OrgPrefix:                        "acme",
		Project:                          "payments",
		Env:                              "prod",
		Region:                           "westeurope",
		RegionMap:                        defaults.RegionMap,
		IgnoreRegionForRegionalResources: false,
		StylePriority:                    []string{StylePascal}, // not allowed for storage accounts
		ResourceAcronyms:                 defaults.ResourceAcronyms,
		ResourceStyleOverrides:           defaults.ResourceStyleOverrides,
		ResourceConstraints:              defaults.ResourceConstraints,
		RegionalResources:                defaults.RegionalResources,
	}, BuildInput{
		Resource:  "azurerm_storage_account",
		Qualifier: "raw",
		Recipe:    []string{"org", "proj", "env", "resource", "qualifier"},
	})
	if err != nil {
		t.Fatalf("unexpected build error: %v", err)
	}

	if result.Style != StyleStraight {
		t.Fatalf("expected fallback style %q, got %q", StyleStraight, result.Style)
	}
}
