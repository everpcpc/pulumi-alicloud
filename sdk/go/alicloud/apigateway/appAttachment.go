// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/apigateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigateway.NewAppAttachment(ctx, "foo", &apigateway.AppAttachmentArgs{
//				ApiId:     pulumi.String("d29d25b9cfdf4742b1a3f6537299a749"),
//				AppId:     pulumi.String("20898181"),
//				GroupId:   pulumi.String("aaef8cdbb404420f9398a74ed1db7fff"),
//				StageName: pulumi.String("PRE"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type AppAttachment struct {
	pulumi.CustomResourceState

	// The apiId that app apply to access.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// The app that apply to the authorization.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The group that the api belongs to.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// Stage that the app apply to access.
	StageName pulumi.StringOutput `pulumi:"stageName"`
}

// NewAppAttachment registers a new resource with the given unique name, arguments, and options.
func NewAppAttachment(ctx *pulumi.Context,
	name string, args *AppAttachmentArgs, opts ...pulumi.ResourceOption) (*AppAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.StageName == nil {
		return nil, errors.New("invalid value for required argument 'StageName'")
	}
	var resource AppAttachment
	err := ctx.RegisterResource("alicloud:apigateway/appAttachment:AppAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppAttachment gets an existing AppAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppAttachmentState, opts ...pulumi.ResourceOption) (*AppAttachment, error) {
	var resource AppAttachment
	err := ctx.ReadResource("alicloud:apigateway/appAttachment:AppAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppAttachment resources.
type appAttachmentState struct {
	// The apiId that app apply to access.
	ApiId *string `pulumi:"apiId"`
	// The app that apply to the authorization.
	AppId *string `pulumi:"appId"`
	// The group that the api belongs to.
	GroupId *string `pulumi:"groupId"`
	// Stage that the app apply to access.
	StageName *string `pulumi:"stageName"`
}

type AppAttachmentState struct {
	// The apiId that app apply to access.
	ApiId pulumi.StringPtrInput
	// The app that apply to the authorization.
	AppId pulumi.StringPtrInput
	// The group that the api belongs to.
	GroupId pulumi.StringPtrInput
	// Stage that the app apply to access.
	StageName pulumi.StringPtrInput
}

func (AppAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*appAttachmentState)(nil)).Elem()
}

type appAttachmentArgs struct {
	// The apiId that app apply to access.
	ApiId string `pulumi:"apiId"`
	// The app that apply to the authorization.
	AppId string `pulumi:"appId"`
	// The group that the api belongs to.
	GroupId string `pulumi:"groupId"`
	// Stage that the app apply to access.
	StageName string `pulumi:"stageName"`
}

// The set of arguments for constructing a AppAttachment resource.
type AppAttachmentArgs struct {
	// The apiId that app apply to access.
	ApiId pulumi.StringInput
	// The app that apply to the authorization.
	AppId pulumi.StringInput
	// The group that the api belongs to.
	GroupId pulumi.StringInput
	// Stage that the app apply to access.
	StageName pulumi.StringInput
}

func (AppAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appAttachmentArgs)(nil)).Elem()
}

type AppAttachmentInput interface {
	pulumi.Input

	ToAppAttachmentOutput() AppAttachmentOutput
	ToAppAttachmentOutputWithContext(ctx context.Context) AppAttachmentOutput
}

func (*AppAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**AppAttachment)(nil)).Elem()
}

func (i *AppAttachment) ToAppAttachmentOutput() AppAttachmentOutput {
	return i.ToAppAttachmentOutputWithContext(context.Background())
}

func (i *AppAttachment) ToAppAttachmentOutputWithContext(ctx context.Context) AppAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppAttachmentOutput)
}

// AppAttachmentArrayInput is an input type that accepts AppAttachmentArray and AppAttachmentArrayOutput values.
// You can construct a concrete instance of `AppAttachmentArrayInput` via:
//
//	AppAttachmentArray{ AppAttachmentArgs{...} }
type AppAttachmentArrayInput interface {
	pulumi.Input

	ToAppAttachmentArrayOutput() AppAttachmentArrayOutput
	ToAppAttachmentArrayOutputWithContext(context.Context) AppAttachmentArrayOutput
}

type AppAttachmentArray []AppAttachmentInput

func (AppAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppAttachment)(nil)).Elem()
}

func (i AppAttachmentArray) ToAppAttachmentArrayOutput() AppAttachmentArrayOutput {
	return i.ToAppAttachmentArrayOutputWithContext(context.Background())
}

func (i AppAttachmentArray) ToAppAttachmentArrayOutputWithContext(ctx context.Context) AppAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppAttachmentArrayOutput)
}

// AppAttachmentMapInput is an input type that accepts AppAttachmentMap and AppAttachmentMapOutput values.
// You can construct a concrete instance of `AppAttachmentMapInput` via:
//
//	AppAttachmentMap{ "key": AppAttachmentArgs{...} }
type AppAttachmentMapInput interface {
	pulumi.Input

	ToAppAttachmentMapOutput() AppAttachmentMapOutput
	ToAppAttachmentMapOutputWithContext(context.Context) AppAttachmentMapOutput
}

type AppAttachmentMap map[string]AppAttachmentInput

func (AppAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppAttachment)(nil)).Elem()
}

func (i AppAttachmentMap) ToAppAttachmentMapOutput() AppAttachmentMapOutput {
	return i.ToAppAttachmentMapOutputWithContext(context.Background())
}

func (i AppAttachmentMap) ToAppAttachmentMapOutputWithContext(ctx context.Context) AppAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppAttachmentMapOutput)
}

type AppAttachmentOutput struct{ *pulumi.OutputState }

func (AppAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppAttachment)(nil)).Elem()
}

func (o AppAttachmentOutput) ToAppAttachmentOutput() AppAttachmentOutput {
	return o
}

func (o AppAttachmentOutput) ToAppAttachmentOutputWithContext(ctx context.Context) AppAttachmentOutput {
	return o
}

// The apiId that app apply to access.
func (o AppAttachmentOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppAttachment) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// The app that apply to the authorization.
func (o AppAttachmentOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppAttachment) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// The group that the api belongs to.
func (o AppAttachmentOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppAttachment) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// Stage that the app apply to access.
func (o AppAttachmentOutput) StageName() pulumi.StringOutput {
	return o.ApplyT(func(v *AppAttachment) pulumi.StringOutput { return v.StageName }).(pulumi.StringOutput)
}

type AppAttachmentArrayOutput struct{ *pulumi.OutputState }

func (AppAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppAttachment)(nil)).Elem()
}

func (o AppAttachmentArrayOutput) ToAppAttachmentArrayOutput() AppAttachmentArrayOutput {
	return o
}

func (o AppAttachmentArrayOutput) ToAppAttachmentArrayOutputWithContext(ctx context.Context) AppAttachmentArrayOutput {
	return o
}

func (o AppAttachmentArrayOutput) Index(i pulumi.IntInput) AppAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppAttachment {
		return vs[0].([]*AppAttachment)[vs[1].(int)]
	}).(AppAttachmentOutput)
}

type AppAttachmentMapOutput struct{ *pulumi.OutputState }

func (AppAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppAttachment)(nil)).Elem()
}

func (o AppAttachmentMapOutput) ToAppAttachmentMapOutput() AppAttachmentMapOutput {
	return o
}

func (o AppAttachmentMapOutput) ToAppAttachmentMapOutputWithContext(ctx context.Context) AppAttachmentMapOutput {
	return o
}

func (o AppAttachmentMapOutput) MapIndex(k pulumi.StringInput) AppAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppAttachment {
		return vs[0].(map[string]*AppAttachment)[vs[1].(string)]
	}).(AppAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppAttachmentInput)(nil)).Elem(), &AppAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppAttachmentArrayInput)(nil)).Elem(), AppAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppAttachmentMapInput)(nil)).Elem(), AppAttachmentMap{})
	pulumi.RegisterOutputType(AppAttachmentOutput{})
	pulumi.RegisterOutputType(AppAttachmentArrayOutput{})
	pulumi.RegisterOutputType(AppAttachmentMapOutput{})
}
