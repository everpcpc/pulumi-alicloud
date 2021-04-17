// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package brain

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
	case "alicloud:brain/industrialPidLoop:IndustrialPidLoop":
		r = &IndustrialPidLoop{}
	case "alicloud:brain/industrialPidOrganization:IndustrialPidOrganization":
		r = &IndustrialPidOrganization{}
	case "alicloud:brain/industrialPidProject:IndustrialPidProject":
		r = &IndustrialPidProject{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := alicloud.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"alicloud",
		"brain/industrialPidLoop",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"brain/industrialPidOrganization",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"brain/industrialPidProject",
		&module{version},
	)
}
