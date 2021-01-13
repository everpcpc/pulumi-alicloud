// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// RAM policy can be imported using the id or name, e.g.
//
// ```sh
//  $ pulumi import alicloud:ram/policy:Policy example my-policy
// ```
type Policy struct {
	pulumi.CustomResourceState

	// The policy attachment count.
	AttachmentCount pulumi.IntOutput `pulumi:"attachmentCount"`
	// Description of the RAM policy. This name can have a string of 1 to 1024 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Document of the RAM policy. It is required when the `statement` is not specified.
	Document pulumi.StringOutput `pulumi:"document"`
	// This parameter is used for resource destroy. Default value is `false`.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	Name pulumi.StringOutput `pulumi:"name"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Statements of the RAM policy document. It is required when the `document` is not specified.
	//
	// Deprecated: Field 'statement' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Statements PolicyStatementArrayOutput `pulumi:"statements"`
	// The policy type.
	Type pulumi.StringOutput `pulumi:"type"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM policy document. Valid value is `1`. Default value is `1`.
	//
	// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		args = &PolicyArgs{}
	}

	var resource Policy
	err := ctx.RegisterResource("alicloud:ram/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("alicloud:ram/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// The policy attachment count.
	AttachmentCount *int `pulumi:"attachmentCount"`
	// Description of the RAM policy. This name can have a string of 1 to 1024 characters.
	Description *string `pulumi:"description"`
	// Document of the RAM policy. It is required when the `statement` is not specified.
	Document *string `pulumi:"document"`
	// This parameter is used for resource destroy. Default value is `false`.
	Force *bool `pulumi:"force"`
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	Name *string `pulumi:"name"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Statements of the RAM policy document. It is required when the `document` is not specified.
	//
	// Deprecated: Field 'statement' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Statements []PolicyStatement `pulumi:"statements"`
	// The policy type.
	Type *string `pulumi:"type"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM policy document. Valid value is `1`. Default value is `1`.
	//
	// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Version *string `pulumi:"version"`
}

type PolicyState struct {
	// The policy attachment count.
	AttachmentCount pulumi.IntPtrInput
	// Description of the RAM policy. This name can have a string of 1 to 1024 characters.
	Description pulumi.StringPtrInput
	// Document of the RAM policy. It is required when the `statement` is not specified.
	Document pulumi.StringPtrInput
	// This parameter is used for resource destroy. Default value is `false`.
	Force pulumi.BoolPtrInput
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	Name pulumi.StringPtrInput
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Statements of the RAM policy document. It is required when the `document` is not specified.
	//
	// Deprecated: Field 'statement' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Statements PolicyStatementArrayInput
	// The policy type.
	Type pulumi.StringPtrInput
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM policy document. Valid value is `1`. Default value is `1`.
	//
	// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Version pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Description of the RAM policy. This name can have a string of 1 to 1024 characters.
	Description *string `pulumi:"description"`
	// Document of the RAM policy. It is required when the `statement` is not specified.
	Document *string `pulumi:"document"`
	// This parameter is used for resource destroy. Default value is `false`.
	Force *bool `pulumi:"force"`
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	Name *string `pulumi:"name"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Statements of the RAM policy document. It is required when the `document` is not specified.
	//
	// Deprecated: Field 'statement' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Statements []PolicyStatement `pulumi:"statements"`
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM policy document. Valid value is `1`. Default value is `1`.
	//
	// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Description of the RAM policy. This name can have a string of 1 to 1024 characters.
	Description pulumi.StringPtrInput
	// Document of the RAM policy. It is required when the `statement` is not specified.
	Document pulumi.StringPtrInput
	// This parameter is used for resource destroy. Default value is `false`.
	Force pulumi.BoolPtrInput
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	Name pulumi.StringPtrInput
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Statements of the RAM policy document. It is required when the `document` is not specified.
	//
	// Deprecated: Field 'statement' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Statements PolicyStatementArrayInput
	// (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM policy document. Valid value is `1`. Default value is `1`.
	//
	// Deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace.
	Version pulumi.StringPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (Policy) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil)).Elem()
}

func (i Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

type PolicyOutput struct {
	*pulumi.OutputState
}

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyOutput)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PolicyOutput{})
}
