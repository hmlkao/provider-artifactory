package localterraformproviderrepository

import (
	"errors"

	"github.com/crossplane/upjet/pkg/config"
)

const shortGroup string = "artifactory"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_local_terraform_provider_repository", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		// Specify Kubernetes kind
		r.Kind = "LocalTerraformProviderRepository"
		// Fix an issue with ExternalName configuration
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if id, ok := tfstate["key"].(string); ok && id != "" {
				return id, nil
			}
			return "", errors.New("cannot find id in tfstate")
		}
	})
}
