package alertpolicy

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("newrelic_alert_policy", func(r *config.Resource) {
		r.ShortGroup = "alertpolicy"

	})
}
