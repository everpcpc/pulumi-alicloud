// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oos

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
	case "alicloud:oos/application:Application":
		r = &Application{}
	case "alicloud:oos/applicationGroup:ApplicationGroup":
		r = &ApplicationGroup{}
	case "alicloud:oos/defaultPatchBaseline:DefaultPatchBaseline":
		r = &DefaultPatchBaseline{}
	case "alicloud:oos/execution:Execution":
		r = &Execution{}
	case "alicloud:oos/parameter:Parameter":
		r = &Parameter{}
	case "alicloud:oos/patchBaseline:PatchBaseline":
		r = &PatchBaseline{}
	case "alicloud:oos/secretParameter:SecretParameter":
		r = &SecretParameter{}
	case "alicloud:oos/serviceSetting:ServiceSetting":
		r = &ServiceSetting{}
	case "alicloud:oos/stateConfiguration:StateConfiguration":
		r = &StateConfiguration{}
	case "alicloud:oos/template:Template":
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
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"alicloud",
		"oos/application",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oos/applicationGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oos/defaultPatchBaseline",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oos/execution",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oos/parameter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oos/patchBaseline",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oos/secretParameter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oos/serviceSetting",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oos/stateConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"oos/template",
		&module{version},
	)
}
