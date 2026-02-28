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
