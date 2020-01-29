// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package log

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/log_project.html.markdown.
type Project struct {
	pulumi.CustomResourceState

	// Description of the log project.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the log project. It is the only in one Alicloud account.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		args = &ProjectArgs{}
	}
	var resource Project
	err := ctx.RegisterResource("alicloud:log/project:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("alicloud:log/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// Description of the log project.
	Description *string `pulumi:"description"`
	// The name of the log project. It is the only in one Alicloud account.
	Name *string `pulumi:"name"`
}

type ProjectState struct {
	// Description of the log project.
	Description pulumi.StringPtrInput
	// The name of the log project. It is the only in one Alicloud account.
	Name pulumi.StringPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// Description of the log project.
	Description *string `pulumi:"description"`
	// The name of the log project. It is the only in one Alicloud account.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// Description of the log project.
	Description pulumi.StringPtrInput
	// The name of the log project. It is the only in one Alicloud account.
	Name pulumi.StringPtrInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

