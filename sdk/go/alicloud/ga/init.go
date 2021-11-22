// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ga

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
	case "alicloud:ga/accelerator:Accelerator":
		r = &Accelerator{}
	case "alicloud:ga/bandwidthPackage:BandwidthPackage":
		r = &BandwidthPackage{}
	case "alicloud:ga/bandwidthPackageAttachment:BandwidthPackageAttachment":
		r = &BandwidthPackageAttachment{}
	case "alicloud:ga/endpointGroup:EndpointGroup":
		r = &EndpointGroup{}
	case "alicloud:ga/forwardingRule:ForwardingRule":
		r = &ForwardingRule{}
	case "alicloud:ga/ipSet:IpSet":
		r = &IpSet{}
	case "alicloud:ga/listener:Listener":
		r = &Listener{}
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
		"ga/accelerator",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ga/bandwidthPackage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ga/bandwidthPackageAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ga/endpointGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ga/forwardingRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ga/ipSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ga/listener",
		&module{version},
	)
}
