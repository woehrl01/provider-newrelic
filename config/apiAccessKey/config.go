package apiAccessKey

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("newrelic_api_access_key", func(r *config.Resource) {
        r.ShortGroup = "api_access_key"

        r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
            conn := map[string][]byte{}
            if a, ok := attr["key"]; ok {
                conn["key"] = []byte(a.(string))
            }
            return conn, nil
        }
    })
}
