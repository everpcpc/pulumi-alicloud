// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package arms

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "alicloud:arms/alertContact:AlertContact":
		r = &AlertContact{}
	case "alicloud:arms/alertContactGroup:AlertContactGroup":
		r = &AlertContactGroup{}
	case "alicloud:arms/dispatchRule:DispatchRule":
		r = &DispatchRule{}
	case "alicloud:arms/prometheusAlertRule:PrometheusAlertRule":
		r = &PrometheusAlertRule{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := alicloud.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"alicloud",
		"arms/alertContact",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"arms/alertContactGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"arms/dispatchRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"arms/prometheusAlertRule",
		&module{version},
	)
}
