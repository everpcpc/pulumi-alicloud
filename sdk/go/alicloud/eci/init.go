// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eci

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
	case "alicloud:eci/containerGroup:ContainerGroup":
		r, err = NewContainerGroup(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:eci/imageCache:ImageCache":
		r, err = NewImageCache(ctx, name, nil, pulumi.URN_(urn))
	case "alicloud:eci/openApiImageCache:OpenApiImageCache":
		r, err = NewOpenApiImageCache(ctx, name, nil, pulumi.URN_(urn))
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
		"eci/containerGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"eci/imageCache",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"alicloud",
		"eci/openApiImageCache",
		&module{version},
	)
}