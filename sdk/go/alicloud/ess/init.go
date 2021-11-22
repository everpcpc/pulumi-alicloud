// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ess

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
	case "alicloud:ess/alarm:Alarm":
		r = &Alarm{}
	case "alicloud:ess/attachment:Attachment":
		r = &Attachment{}
	case "alicloud:ess/lifecycleHook:LifecycleHook":
		r = &LifecycleHook{}
	case "alicloud:ess/notification:Notification":
		r = &Notification{}
	case "alicloud:ess/scalingConfiguration:ScalingConfiguration":
		r = &ScalingConfiguration{}
	case "alicloud:ess/scalingGroup:ScalingGroup":
		r = &ScalingGroup{}
	case "alicloud:ess/scalingGroupVServerGroups:ScalingGroupVServerGroups":
		r = &ScalingGroupVServerGroups{}
	case "alicloud:ess/scalingRule:ScalingRule":
		r = &ScalingRule{}
	case "alicloud:ess/schedule:Schedule":
		r = &Schedule{}
	case "alicloud:ess/scheduledTask:ScheduledTask":
		r = &ScheduledTask{}
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
		"ess/alarm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ess/attachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ess/lifecycleHook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ess/notification",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ess/scalingConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ess/scalingGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ess/scalingGroupVServerGroups",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ess/scalingRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ess/schedule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"ess/scheduledTask",
		&module{version},
	)
}
