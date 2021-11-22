// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cassandra

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
	case "alicloud:cassandra/backupPlan:BackupPlan":
		r = &BackupPlan{}
	case "alicloud:cassandra/cluster:Cluster":
		r = &Cluster{}
	case "alicloud:cassandra/dataCenter:DataCenter":
		r = &DataCenter{}
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
		"cassandra/backupPlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"cassandra/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"cassandra/dataCenter",
		&module{version},
	)
}
