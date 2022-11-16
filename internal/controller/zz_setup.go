/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	policy "github.com/woehrl01/provider-newrelic/internal/controller/alertpolicy/policy"
	accesskey "github.com/woehrl01/provider-newrelic/internal/controller/apiaccesskey/accesskey"
	alertcondition "github.com/woehrl01/provider-newrelic/internal/controller/nrqlalertcondition/alertcondition"
	providerconfig "github.com/woehrl01/provider-newrelic/internal/controller/providerconfig"
	certcheckmonitor "github.com/woehrl01/provider-newrelic/internal/controller/syntheticscertcheckmonitor/certcheckmonitor"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		policy.Setup,
		accesskey.Setup,
		alertcondition.Setup,
		providerconfig.Setup,
		certcheckmonitor.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
