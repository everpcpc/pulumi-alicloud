// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource will help you to manager Container Registry Enterprise Edition namespaces.
//
// For information about Container Registry Enterprise Edition namespaces and how to use it, see [Create a Namespace](https://www.alibabacloud.com/help/doc-detail/145483.htm)
//
// > **NOTE:** Available in v1.86.0+.
//
// > **NOTE:** You need to set your registry password in Container Registry Enterprise Edition console before use this resource.
type RegistryEnterpriseNamespace struct {
	pulumi.CustomResourceState

	// Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
	AutoCreate pulumi.BoolOutput `pulumi:"autoCreate"`
	// `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
	DefaultVisibility pulumi.StringOutput `pulumi:"defaultVisibility"`
	// ID of Container Registry Enterprise Edition instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Name of Container Registry Enterprise Edition namespace. It can contain 2 to 30 characters.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewRegistryEnterpriseNamespace registers a new resource with the given unique name, arguments, and options.
func NewRegistryEnterpriseNamespace(ctx *pulumi.Context,
	name string, args *RegistryEnterpriseNamespaceArgs, opts ...pulumi.ResourceOption) (*RegistryEnterpriseNamespace, error) {
	if args == nil || args.AutoCreate == nil {
		return nil, errors.New("missing required argument 'AutoCreate'")
	}
	if args == nil || args.DefaultVisibility == nil {
		return nil, errors.New("missing required argument 'DefaultVisibility'")
	}
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil {
		args = &RegistryEnterpriseNamespaceArgs{}
	}
	var resource RegistryEnterpriseNamespace
	err := ctx.RegisterResource("alicloud:cs/registryEnterpriseNamespace:RegistryEnterpriseNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryEnterpriseNamespace gets an existing RegistryEnterpriseNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryEnterpriseNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryEnterpriseNamespaceState, opts ...pulumi.ResourceOption) (*RegistryEnterpriseNamespace, error) {
	var resource RegistryEnterpriseNamespace
	err := ctx.ReadResource("alicloud:cs/registryEnterpriseNamespace:RegistryEnterpriseNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryEnterpriseNamespace resources.
type registryEnterpriseNamespaceState struct {
	// Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
	AutoCreate *bool `pulumi:"autoCreate"`
	// `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
	DefaultVisibility *string `pulumi:"defaultVisibility"`
	// ID of Container Registry Enterprise Edition instance.
	InstanceId *string `pulumi:"instanceId"`
	// Name of Container Registry Enterprise Edition namespace. It can contain 2 to 30 characters.
	Name *string `pulumi:"name"`
}

type RegistryEnterpriseNamespaceState struct {
	// Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
	AutoCreate pulumi.BoolPtrInput
	// `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
	DefaultVisibility pulumi.StringPtrInput
	// ID of Container Registry Enterprise Edition instance.
	InstanceId pulumi.StringPtrInput
	// Name of Container Registry Enterprise Edition namespace. It can contain 2 to 30 characters.
	Name pulumi.StringPtrInput
}

func (RegistryEnterpriseNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnterpriseNamespaceState)(nil)).Elem()
}

type registryEnterpriseNamespaceArgs struct {
	// Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
	AutoCreate bool `pulumi:"autoCreate"`
	// `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
	DefaultVisibility string `pulumi:"defaultVisibility"`
	// ID of Container Registry Enterprise Edition instance.
	InstanceId string `pulumi:"instanceId"`
	// Name of Container Registry Enterprise Edition namespace. It can contain 2 to 30 characters.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a RegistryEnterpriseNamespace resource.
type RegistryEnterpriseNamespaceArgs struct {
	// Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
	AutoCreate pulumi.BoolInput
	// `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
	DefaultVisibility pulumi.StringInput
	// ID of Container Registry Enterprise Edition instance.
	InstanceId pulumi.StringInput
	// Name of Container Registry Enterprise Edition namespace. It can contain 2 to 30 characters.
	Name pulumi.StringPtrInput
}

func (RegistryEnterpriseNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryEnterpriseNamespaceArgs)(nil)).Elem()
}
