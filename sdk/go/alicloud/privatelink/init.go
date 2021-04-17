// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

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
	case "alicloud:privatelink/vpcEndpoint:VpcEndpoint":
		r = &VpcEndpoint{}
	case "alicloud:privatelink/vpcEndpointService:VpcEndpointService":
		r = &VpcEndpointService{}
	case "alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection":
		r = &VpcEndpointServiceConnection{}
	case "alicloud:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource":
		r = &VpcEndpointServiceResource{}
	case "alicloud:privatelink/vpcEndpointServiceUser:VpcEndpointServiceUser":
		r = &VpcEndpointServiceUser{}
	case "alicloud:privatelink/vpcEndpointZone:VpcEndpointZone":
		r = &VpcEndpointZone{}
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
		"privatelink/vpcEndpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"privatelink/vpcEndpointService",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"privatelink/vpcEndpointServiceConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"privatelink/vpcEndpointServiceResource",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"privatelink/vpcEndpointServiceUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"privatelink/vpcEndpointZone",
		&module{version},
	)
}
