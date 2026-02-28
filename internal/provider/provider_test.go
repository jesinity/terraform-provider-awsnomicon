package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jesinity/terraform-provider-sigil/internal/naming"
)

func TestResolveCloudPrecedence(t *testing.T) {
	base := providerConfigModel{Cloud: types.StringValue(naming.CloudAzure)}
	override := providerConfigModel{Cloud: types.StringValue(naming.CloudAWS)}

	resolved := resolveCloud(types.StringValue(naming.CloudAzure), base, true, override, true)
	if resolved != naming.CloudAWS {
		t.Fatalf("expected override cloud to win (%q), got %q", naming.CloudAWS, resolved)
	}

	resolved = resolveCloud(types.StringNull(), base, true, providerConfigModel{}, false)
	if resolved != naming.CloudAzure {
		t.Fatalf("expected base cloud (%q), got %q", naming.CloudAzure, resolved)
	}

	resolved = resolveCloud(types.StringNull(), providerConfigModel{}, false, providerConfigModel{}, false)
	if resolved != naming.CloudAWS {
		t.Fatalf("expected default cloud (%q), got %q", naming.CloudAWS, resolved)
	}
}

func TestResolveUseAzureCAFAcronymsPrecedence(t *testing.T) {
	base := providerConfigModel{UseAzureCAFAcronyms: types.BoolValue(true)}
	override := providerConfigModel{UseAzureCAFAcronyms: types.BoolValue(false)}

	resolved := resolveUseAzureCAFAcronyms(types.BoolNull(), base, true, override, true)
	if resolved {
		t.Fatalf("expected override value to win (false), got %t", resolved)
	}

	resolved = resolveUseAzureCAFAcronyms(types.BoolValue(true), base, true, providerConfigModel{}, false)
	if !resolved {
		t.Fatalf("expected top-level value to win (true), got %t", resolved)
	}

	resolved = resolveUseAzureCAFAcronyms(types.BoolNull(), providerConfigModel{}, false, providerConfigModel{}, false)
	if resolved {
		t.Fatalf("expected default value (false), got %t", resolved)
	}
}
