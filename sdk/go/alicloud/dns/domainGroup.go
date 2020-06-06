// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Alidns Domain Group resource. For information about Alidns Domain Group and how to use it, see [What is Resource Alidns Domain Group](https://www.alibabacloud.com/help/en/doc-detail/29762.htm).
//
// > **NOTE:** Available in v1.84.0+.
type DomainGroup struct {
	pulumi.CustomResourceState

	// Name of the domain group.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// User language.
	Lang pulumi.StringPtrOutput `pulumi:"lang"`
}

// NewDomainGroup registers a new resource with the given unique name, arguments, and options.
func NewDomainGroup(ctx *pulumi.Context,
	name string, args *DomainGroupArgs, opts ...pulumi.ResourceOption) (*DomainGroup, error) {
	if args == nil || args.GroupName == nil {
		return nil, errors.New("missing required argument 'GroupName'")
	}
	if args == nil {
		args = &DomainGroupArgs{}
	}
	var resource DomainGroup
	err := ctx.RegisterResource("alicloud:dns/domainGroup:DomainGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainGroup gets an existing DomainGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainGroupState, opts ...pulumi.ResourceOption) (*DomainGroup, error) {
	var resource DomainGroup
	err := ctx.ReadResource("alicloud:dns/domainGroup:DomainGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainGroup resources.
type domainGroupState struct {
	// Name of the domain group.
	GroupName *string `pulumi:"groupName"`
	// User language.
	Lang *string `pulumi:"lang"`
}

type DomainGroupState struct {
	// Name of the domain group.
	GroupName pulumi.StringPtrInput
	// User language.
	Lang pulumi.StringPtrInput
}

func (DomainGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainGroupState)(nil)).Elem()
}

type domainGroupArgs struct {
	// Name of the domain group.
	GroupName string `pulumi:"groupName"`
	// User language.
	Lang *string `pulumi:"lang"`
}

// The set of arguments for constructing a DomainGroup resource.
type DomainGroupArgs struct {
	// Name of the domain group.
	GroupName pulumi.StringInput
	// User language.
	Lang pulumi.StringPtrInput
}

func (DomainGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainGroupArgs)(nil)).Elem()
}