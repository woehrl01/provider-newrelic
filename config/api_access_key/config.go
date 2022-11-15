package api_access_key

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("newrelic_api_access_key", func(r *config.Resource) {
        r.ShortGroup = "api_access_key"
    })
}
