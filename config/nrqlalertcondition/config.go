package nrqlalertcondition

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("newrelic_nrql_alert_condition", func(r *config.Resource) {
		r.ShortGroup = "nrqlalertcondition"

		r.References["PolicyID"] = config.Reference{
			Type: "github.com/woehrl01/provider-newrelic/apis/alertpolicy/v1alpha1.Policy",
			RefFieldName: "ID",
		}

	})
}
