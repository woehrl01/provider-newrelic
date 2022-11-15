/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/woehrl01/provider-newrelic/config/apiaccesskey"
	"github.com/woehrl01/provider-newrelic/config/syntheticscertcheckmonitor"
)

const (
	resourcePrefix = "newrelic"
	modulePath     = "github.com/woehrl01/provider-newrelic"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		apiaccesskey.Configure,
		syntheticscertcheckmonitor.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
