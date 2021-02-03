// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rocketmq

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
	case "alicloud:rocketmq/acl:Acl":
		r, err = NewAcl(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:rocketmq/aclRule:AclRule":
		r, err = NewAclRule(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:rocketmq/clientUser:ClientUser":
		r, err = NewClientUser(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:rocketmq/dnatEntry:DnatEntry":
		r, err = NewDnatEntry(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:rocketmq/group:Group":
		r, err = NewGroup(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:rocketmq/instance:Instance":
		r, err = NewInstance(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:rocketmq/qos:Qos":
		r, err = NewQos(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:rocketmq/qosCar:QosCar":
		r, err = NewQosCar(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:rocketmq/qosPolicy:QosPolicy":
		r, err = NewQosPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:rocketmq/snatEntry:SnatEntry":
		r, err = NewSnatEntry(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:rocketmq/topic:Topic":
		r, err = NewTopic(ctx, name, nil, pulumi.URN_(urn))
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
		"rocketmq/acl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"rocketmq/aclRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"rocketmq/clientUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"rocketmq/dnatEntry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"rocketmq/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"rocketmq/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"rocketmq/qos",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"rocketmq/qosCar",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"rocketmq/qosPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"rocketmq/snatEntry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"rocketmq/topic",
		&module{version},
	)
}