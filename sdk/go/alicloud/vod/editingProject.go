// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vod

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VOD Editing Project resource.
//
// For information about VOD Editing Project and how to use it, see [What is Editing Project](https://www.alibabacloud.com/help/en/apsaravideo-for-vod/latest/addeditingproject#doc-api-vod-AddEditingProject).
//
// > **NOTE:** Available in v1.187.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vod"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vod.NewEditingProject(ctx, "example", &vod.EditingProjectArgs{
//				EditingProjectName: pulumi.String("example_value"),
//				Timeline:           pulumi.String("example_value"),
//				Title:              pulumi.String("example_value"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// VOD Editing Project can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:vod/editingProject:EditingProject example <id>
//
// ```
type EditingProject struct {
	pulumi.CustomResourceState

	// The thumbnail URL of the online editing project. If you do not specify this parameter and the video track in the timeline has mezzanine files, the thumbnail of the first mezzanine file in the timeline is used.
	CoverUrl pulumi.StringPtrOutput `pulumi:"coverUrl"`
	// The region where you want to create the online editing project.
	Division pulumi.StringPtrOutput `pulumi:"division"`
	// The description of the online editing project.
	EditingProjectName pulumi.StringOutput `pulumi:"editingProjectName"`
	// The Status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// The timeline of the online editing project, in JSON format. For more information about the structure, see [Timeline](https://www.alibabacloud.com/help/en/apsaravideo-for-vod/latest/basic-structures). If you do not specify this parameter, an empty timeline is created and the duration of the online editing project is zero.
	Timeline pulumi.StringOutput `pulumi:"timeline"`
	// The title of the online editing project.
	Title pulumi.StringOutput `pulumi:"title"`
}

// NewEditingProject registers a new resource with the given unique name, arguments, and options.
func NewEditingProject(ctx *pulumi.Context,
	name string, args *EditingProjectArgs, opts ...pulumi.ResourceOption) (*EditingProject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	var resource EditingProject
	err := ctx.RegisterResource("alicloud:vod/editingProject:EditingProject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEditingProject gets an existing EditingProject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEditingProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EditingProjectState, opts ...pulumi.ResourceOption) (*EditingProject, error) {
	var resource EditingProject
	err := ctx.ReadResource("alicloud:vod/editingProject:EditingProject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EditingProject resources.
type editingProjectState struct {
	// The thumbnail URL of the online editing project. If you do not specify this parameter and the video track in the timeline has mezzanine files, the thumbnail of the first mezzanine file in the timeline is used.
	CoverUrl *string `pulumi:"coverUrl"`
	// The region where you want to create the online editing project.
	Division *string `pulumi:"division"`
	// The description of the online editing project.
	EditingProjectName *string `pulumi:"editingProjectName"`
	// The Status of the resource.
	Status *string `pulumi:"status"`
	// The timeline of the online editing project, in JSON format. For more information about the structure, see [Timeline](https://www.alibabacloud.com/help/en/apsaravideo-for-vod/latest/basic-structures). If you do not specify this parameter, an empty timeline is created and the duration of the online editing project is zero.
	Timeline *string `pulumi:"timeline"`
	// The title of the online editing project.
	Title *string `pulumi:"title"`
}

type EditingProjectState struct {
	// The thumbnail URL of the online editing project. If you do not specify this parameter and the video track in the timeline has mezzanine files, the thumbnail of the first mezzanine file in the timeline is used.
	CoverUrl pulumi.StringPtrInput
	// The region where you want to create the online editing project.
	Division pulumi.StringPtrInput
	// The description of the online editing project.
	EditingProjectName pulumi.StringPtrInput
	// The Status of the resource.
	Status pulumi.StringPtrInput
	// The timeline of the online editing project, in JSON format. For more information about the structure, see [Timeline](https://www.alibabacloud.com/help/en/apsaravideo-for-vod/latest/basic-structures). If you do not specify this parameter, an empty timeline is created and the duration of the online editing project is zero.
	Timeline pulumi.StringPtrInput
	// The title of the online editing project.
	Title pulumi.StringPtrInput
}

func (EditingProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*editingProjectState)(nil)).Elem()
}

type editingProjectArgs struct {
	// The thumbnail URL of the online editing project. If you do not specify this parameter and the video track in the timeline has mezzanine files, the thumbnail of the first mezzanine file in the timeline is used.
	CoverUrl *string `pulumi:"coverUrl"`
	// The region where you want to create the online editing project.
	Division *string `pulumi:"division"`
	// The description of the online editing project.
	EditingProjectName *string `pulumi:"editingProjectName"`
	// The timeline of the online editing project, in JSON format. For more information about the structure, see [Timeline](https://www.alibabacloud.com/help/en/apsaravideo-for-vod/latest/basic-structures). If you do not specify this parameter, an empty timeline is created and the duration of the online editing project is zero.
	Timeline *string `pulumi:"timeline"`
	// The title of the online editing project.
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a EditingProject resource.
type EditingProjectArgs struct {
	// The thumbnail URL of the online editing project. If you do not specify this parameter and the video track in the timeline has mezzanine files, the thumbnail of the first mezzanine file in the timeline is used.
	CoverUrl pulumi.StringPtrInput
	// The region where you want to create the online editing project.
	Division pulumi.StringPtrInput
	// The description of the online editing project.
	EditingProjectName pulumi.StringPtrInput
	// The timeline of the online editing project, in JSON format. For more information about the structure, see [Timeline](https://www.alibabacloud.com/help/en/apsaravideo-for-vod/latest/basic-structures). If you do not specify this parameter, an empty timeline is created and the duration of the online editing project is zero.
	Timeline pulumi.StringPtrInput
	// The title of the online editing project.
	Title pulumi.StringInput
}

func (EditingProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*editingProjectArgs)(nil)).Elem()
}

type EditingProjectInput interface {
	pulumi.Input

	ToEditingProjectOutput() EditingProjectOutput
	ToEditingProjectOutputWithContext(ctx context.Context) EditingProjectOutput
}

func (*EditingProject) ElementType() reflect.Type {
	return reflect.TypeOf((**EditingProject)(nil)).Elem()
}

func (i *EditingProject) ToEditingProjectOutput() EditingProjectOutput {
	return i.ToEditingProjectOutputWithContext(context.Background())
}

func (i *EditingProject) ToEditingProjectOutputWithContext(ctx context.Context) EditingProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EditingProjectOutput)
}

// EditingProjectArrayInput is an input type that accepts EditingProjectArray and EditingProjectArrayOutput values.
// You can construct a concrete instance of `EditingProjectArrayInput` via:
//
//	EditingProjectArray{ EditingProjectArgs{...} }
type EditingProjectArrayInput interface {
	pulumi.Input

	ToEditingProjectArrayOutput() EditingProjectArrayOutput
	ToEditingProjectArrayOutputWithContext(context.Context) EditingProjectArrayOutput
}

type EditingProjectArray []EditingProjectInput

func (EditingProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EditingProject)(nil)).Elem()
}

func (i EditingProjectArray) ToEditingProjectArrayOutput() EditingProjectArrayOutput {
	return i.ToEditingProjectArrayOutputWithContext(context.Background())
}

func (i EditingProjectArray) ToEditingProjectArrayOutputWithContext(ctx context.Context) EditingProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EditingProjectArrayOutput)
}

// EditingProjectMapInput is an input type that accepts EditingProjectMap and EditingProjectMapOutput values.
// You can construct a concrete instance of `EditingProjectMapInput` via:
//
//	EditingProjectMap{ "key": EditingProjectArgs{...} }
type EditingProjectMapInput interface {
	pulumi.Input

	ToEditingProjectMapOutput() EditingProjectMapOutput
	ToEditingProjectMapOutputWithContext(context.Context) EditingProjectMapOutput
}

type EditingProjectMap map[string]EditingProjectInput

func (EditingProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EditingProject)(nil)).Elem()
}

func (i EditingProjectMap) ToEditingProjectMapOutput() EditingProjectMapOutput {
	return i.ToEditingProjectMapOutputWithContext(context.Background())
}

func (i EditingProjectMap) ToEditingProjectMapOutputWithContext(ctx context.Context) EditingProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EditingProjectMapOutput)
}

type EditingProjectOutput struct{ *pulumi.OutputState }

func (EditingProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EditingProject)(nil)).Elem()
}

func (o EditingProjectOutput) ToEditingProjectOutput() EditingProjectOutput {
	return o
}

func (o EditingProjectOutput) ToEditingProjectOutputWithContext(ctx context.Context) EditingProjectOutput {
	return o
}

// The thumbnail URL of the online editing project. If you do not specify this parameter and the video track in the timeline has mezzanine files, the thumbnail of the first mezzanine file in the timeline is used.
func (o EditingProjectOutput) CoverUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EditingProject) pulumi.StringPtrOutput { return v.CoverUrl }).(pulumi.StringPtrOutput)
}

// The region where you want to create the online editing project.
func (o EditingProjectOutput) Division() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EditingProject) pulumi.StringPtrOutput { return v.Division }).(pulumi.StringPtrOutput)
}

// The description of the online editing project.
func (o EditingProjectOutput) EditingProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *EditingProject) pulumi.StringOutput { return v.EditingProjectName }).(pulumi.StringOutput)
}

// The Status of the resource.
func (o EditingProjectOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *EditingProject) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The timeline of the online editing project, in JSON format. For more information about the structure, see [Timeline](https://www.alibabacloud.com/help/en/apsaravideo-for-vod/latest/basic-structures). If you do not specify this parameter, an empty timeline is created and the duration of the online editing project is zero.
func (o EditingProjectOutput) Timeline() pulumi.StringOutput {
	return o.ApplyT(func(v *EditingProject) pulumi.StringOutput { return v.Timeline }).(pulumi.StringOutput)
}

// The title of the online editing project.
func (o EditingProjectOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *EditingProject) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

type EditingProjectArrayOutput struct{ *pulumi.OutputState }

func (EditingProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EditingProject)(nil)).Elem()
}

func (o EditingProjectArrayOutput) ToEditingProjectArrayOutput() EditingProjectArrayOutput {
	return o
}

func (o EditingProjectArrayOutput) ToEditingProjectArrayOutputWithContext(ctx context.Context) EditingProjectArrayOutput {
	return o
}

func (o EditingProjectArrayOutput) Index(i pulumi.IntInput) EditingProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EditingProject {
		return vs[0].([]*EditingProject)[vs[1].(int)]
	}).(EditingProjectOutput)
}

type EditingProjectMapOutput struct{ *pulumi.OutputState }

func (EditingProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EditingProject)(nil)).Elem()
}

func (o EditingProjectMapOutput) ToEditingProjectMapOutput() EditingProjectMapOutput {
	return o
}

func (o EditingProjectMapOutput) ToEditingProjectMapOutputWithContext(ctx context.Context) EditingProjectMapOutput {
	return o
}

func (o EditingProjectMapOutput) MapIndex(k pulumi.StringInput) EditingProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EditingProject {
		return vs[0].(map[string]*EditingProject)[vs[1].(string)]
	}).(EditingProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EditingProjectInput)(nil)).Elem(), &EditingProject{})
	pulumi.RegisterInputType(reflect.TypeOf((*EditingProjectArrayInput)(nil)).Elem(), EditingProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EditingProjectMapInput)(nil)).Elem(), EditingProjectMap{})
	pulumi.RegisterOutputType(EditingProjectOutput{})
	pulumi.RegisterOutputType(EditingProjectArrayOutput{})
	pulumi.RegisterOutputType(EditingProjectMapOutput{})
}
