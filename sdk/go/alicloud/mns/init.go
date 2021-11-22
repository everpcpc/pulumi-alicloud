// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mns

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
	case "alicloud:mns/queue:Queue":
		r = &Queue{}
	case "alicloud:mns/topic:Topic":
		r = &Topic{}
	case "alicloud:mns/topicSubscription:TopicSubscription":
		r = &TopicSubscription{}
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
		"mns/queue",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"mns/topic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"mns/topicSubscription",
		&module{version},
	)
}
