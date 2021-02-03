// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "alicloud:vpc/commonBandwithPackage:CommonBandwithPackage":
		r, err = NewCommonBandwithPackage(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/commonBandwithPackageAttachment:CommonBandwithPackageAttachment":
		r, err = NewCommonBandwithPackageAttachment(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/forwardEntry:ForwardEntry":
		r, err = NewForwardEntry(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/hAVip:HAVip":
		r, err = NewHAVip(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/hAVipAttachment:HAVipAttachment":
		r, err = NewHAVipAttachment(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/natGateway:NatGateway":
		r, err = NewNatGateway(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/network:Network":
		r, err = NewNetwork(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/networkAcl:NetworkAcl":
		r, err = NewNetworkAcl(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/networkAclAttachment:NetworkAclAttachment":
		r, err = NewNetworkAclAttachment(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/networkAclEntries:NetworkAclEntries":
		r, err = NewNetworkAclEntries(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/networkInterface:NetworkInterface":
		r, err = NewNetworkInterface(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/networkInterfaceAttachment:NetworkInterfaceAttachment":
		r, err = NewNetworkInterfaceAttachment(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/routeEntry:RouteEntry":
		r, err = NewRouteEntry(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/routeTable:RouteTable":
		r, err = NewRouteTable(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/routeTableAttachment:RouteTableAttachment":
		r, err = NewRouteTableAttachment(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/routerInterface:RouterInterface":
		r, err = NewRouterInterface(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/routerInterfaceConnection:RouterInterfaceConnection":
		r, err = NewRouterInterfaceConnection(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/snatEntry:SnatEntry":
		r, err = NewSnatEntry(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/subnet:Subnet":
		r, err = NewSubnet(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:vpc/switch:Switch":
		r, err = NewSwitch(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := alicloud.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/commonBandwithPackage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/commonBandwithPackageAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/forwardEntry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/hAVip",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/hAVipAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/natGateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/network",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/networkAcl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/networkAclAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/networkAclEntries",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/networkInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/networkInterfaceAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/routeEntry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/routeTable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/routeTableAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/routerInterface",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/routerInterfaceConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/snatEntry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/subnet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"vpc/switch",
		&module{version},
	)
}