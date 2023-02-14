// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hbr

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
	case "alicloud:hbr/ecsBackupClient:EcsBackupClient":
		r = &EcsBackupClient{}
	case "alicloud:hbr/ecsBackupPlan:EcsBackupPlan":
		r = &EcsBackupPlan{}
	case "alicloud:hbr/hanaBackupClient:HanaBackupClient":
		r = &HanaBackupClient{}
	case "alicloud:hbr/hanaBackupPlan:HanaBackupPlan":
		r = &HanaBackupPlan{}
	case "alicloud:hbr/hanaInstance:HanaInstance":
		r = &HanaInstance{}
	case "alicloud:hbr/nasBackupPlan:NasBackupPlan":
		r = &NasBackupPlan{}
	case "alicloud:hbr/ossBackupPlan:OssBackupPlan":
		r = &OssBackupPlan{}
	case "alicloud:hbr/otsBackupPlan:OtsBackupPlan":
		r = &OtsBackupPlan{}
	case "alicloud:hbr/replicationVault:ReplicationVault":
		r = &ReplicationVault{}
	case "alicloud:hbr/restoreJob:RestoreJob":
		r = &RestoreJob{}
	case "alicloud:hbr/serverBackupPlan:ServerBackupPlan":
		r = &ServerBackupPlan{}
	case "alicloud:hbr/vault:Vault":
		r = &Vault{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := alicloud.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"alicloud",
		"hbr/ecsBackupClient",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"hbr/ecsBackupPlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"hbr/hanaBackupClient",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"hbr/hanaBackupPlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"hbr/hanaInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"hbr/nasBackupPlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"hbr/ossBackupPlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"hbr/otsBackupPlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"hbr/replicationVault",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"hbr/restoreJob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"hbr/serverBackupPlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"hbr/vault",
		&module{version},
	)
}
