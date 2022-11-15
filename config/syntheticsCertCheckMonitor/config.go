package syntheticscertcheckmonitor

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("newrelic_synthetics_cert_check_monitor", func(r *config.Resource) {
		r.ShortGroup = "synthetics_cert_check_monitor"

	})
}
