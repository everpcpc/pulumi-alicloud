// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ros

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
	case "alicloud:ros/changeSet:ChangeSet":
		r = &ChangeSet{}
	case "alicloud:ros/stack:Stack":
		r = &Stack{}
	case "alicloud:ros/stackGroup:StackGroup":
		r = &StackGroup{}
	case "alicloud:ros/template:Template":
		r = &Template{}
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
		"ros/changeSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ros/stack",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ros/stackGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ros/template",
		&module{version},
	)
}
