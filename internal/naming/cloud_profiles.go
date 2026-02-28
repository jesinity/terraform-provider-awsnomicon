package naming

import "fmt"

const (
	CloudAWS   = "aws"
	CloudAzure = "azure"
)

type CloudDefaults struct {
	RegionMap              map[string]string
	ResourceAcronyms       map[string]string
	ResourceStyleOverrides map[string][]string
	ResourceConstraints    map[string]ResourceConstraint
	RegionalResources      map[string]bool
}

type CloudDefaultsOptions struct {
	UseAzureCAFAcronyms bool
}

type CloudProfile interface {
	Cloud() string
	Defaults(options CloudDefaultsOptions) (CloudDefaults, error)
}

var cloudProfiles = map[string]CloudProfile{
	CloudAWS:   awsCloudProfile{},
	CloudAzure: newAzureCloudProfile(),
}

func DefaultCloud() string {
	return CloudAWS
}

func NormalizeCloud(cloud string) string {
	normalized := normalizeStyle(cloud)
	if normalized == "" {
		return DefaultCloud()
	}
	return normalized
}

func IsSupportedCloud(cloud string) bool {
	_, ok := cloudProfiles[NormalizeCloud(cloud)]
	return ok
}

func DefaultCloudDefaults(cloud string) (CloudDefaults, error) {
	return DefaultCloudDefaultsWithOptions(cloud, CloudDefaultsOptions{})
}

func DefaultCloudDefaultsWithOptions(cloud string, options CloudDefaultsOptions) (CloudDefaults, error) {
	normalized := NormalizeCloud(cloud)
	profile, ok := cloudProfiles[normalized]
	if !ok {
		return CloudDefaults{}, fmt.Errorf("unsupported cloud %q", cloud)
	}
	return profile.Defaults(options)
}
