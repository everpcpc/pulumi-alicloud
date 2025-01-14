// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Function Compute service alias. Creates an alias that points to the specified Function Compute service version.
//
//	For the detailed information, please refer to the [developer guide](https://www.alibabacloud.com/help/doc-detail/171635.htm).
//
// > **NOTE:** Available in 1.104.0+
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/fc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fc.NewAlias(ctx, "example", &fc.AliasArgs{
//				AliasName:   pulumi.String("my_alias"),
//				Description: pulumi.String("a sample description"),
//				RoutingConfig: &fc.AliasRoutingConfigArgs{
//					AdditionalVersionWeights: pulumi.Float64Map{
//						"2": pulumi.Float64(0.5),
//					},
//				},
//				ServiceName:    pulumi.String("my_service_name"),
//				ServiceVersion: pulumi.String("1"),
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
// Function Compute alias can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:fc/alias:Alias example my_alias_id
//
// ```
type Alias struct {
	pulumi.CustomResourceState

	// Name for the alias you are creating.
	AliasName pulumi.StringOutput `pulumi:"aliasName"`
	// Description of the alias.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Function Compute alias' route configuration settings. Fields documented below.
	//
	// **routing_config** includes the following arguments:
	RoutingConfig AliasRoutingConfigPtrOutput `pulumi:"routingConfig"`
	// The Function Compute service name.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// The Function Compute service version for which you are creating the alias. Pattern: (LATEST|[0-9]+).
	ServiceVersion pulumi.StringOutput `pulumi:"serviceVersion"`
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOption) (*Alias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AliasName == nil {
		return nil, errors.New("invalid value for required argument 'AliasName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.ServiceVersion == nil {
		return nil, errors.New("invalid value for required argument 'ServiceVersion'")
	}
	var resource Alias
	err := ctx.RegisterResource("alicloud:fc/alias:Alias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AliasState, opts ...pulumi.ResourceOption) (*Alias, error) {
	var resource Alias
	err := ctx.ReadResource("alicloud:fc/alias:Alias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alias resources.
type aliasState struct {
	// Name for the alias you are creating.
	AliasName *string `pulumi:"aliasName"`
	// Description of the alias.
	Description *string `pulumi:"description"`
	// The Function Compute alias' route configuration settings. Fields documented below.
	//
	// **routing_config** includes the following arguments:
	RoutingConfig *AliasRoutingConfig `pulumi:"routingConfig"`
	// The Function Compute service name.
	ServiceName *string `pulumi:"serviceName"`
	// The Function Compute service version for which you are creating the alias. Pattern: (LATEST|[0-9]+).
	ServiceVersion *string `pulumi:"serviceVersion"`
}

type AliasState struct {
	// Name for the alias you are creating.
	AliasName pulumi.StringPtrInput
	// Description of the alias.
	Description pulumi.StringPtrInput
	// The Function Compute alias' route configuration settings. Fields documented below.
	//
	// **routing_config** includes the following arguments:
	RoutingConfig AliasRoutingConfigPtrInput
	// The Function Compute service name.
	ServiceName pulumi.StringPtrInput
	// The Function Compute service version for which you are creating the alias. Pattern: (LATEST|[0-9]+).
	ServiceVersion pulumi.StringPtrInput
}

func (AliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasState)(nil)).Elem()
}

type aliasArgs struct {
	// Name for the alias you are creating.
	AliasName string `pulumi:"aliasName"`
	// Description of the alias.
	Description *string `pulumi:"description"`
	// The Function Compute alias' route configuration settings. Fields documented below.
	//
	// **routing_config** includes the following arguments:
	RoutingConfig *AliasRoutingConfig `pulumi:"routingConfig"`
	// The Function Compute service name.
	ServiceName string `pulumi:"serviceName"`
	// The Function Compute service version for which you are creating the alias. Pattern: (LATEST|[0-9]+).
	ServiceVersion string `pulumi:"serviceVersion"`
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	// Name for the alias you are creating.
	AliasName pulumi.StringInput
	// Description of the alias.
	Description pulumi.StringPtrInput
	// The Function Compute alias' route configuration settings. Fields documented below.
	//
	// **routing_config** includes the following arguments:
	RoutingConfig AliasRoutingConfigPtrInput
	// The Function Compute service name.
	ServiceName pulumi.StringInput
	// The Function Compute service version for which you are creating the alias. Pattern: (LATEST|[0-9]+).
	ServiceVersion pulumi.StringInput
}

func (AliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasArgs)(nil)).Elem()
}

type AliasInput interface {
	pulumi.Input

	ToAliasOutput() AliasOutput
	ToAliasOutputWithContext(ctx context.Context) AliasOutput
}

func (*Alias) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (i *Alias) ToAliasOutput() AliasOutput {
	return i.ToAliasOutputWithContext(context.Background())
}

func (i *Alias) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasOutput)
}

// AliasArrayInput is an input type that accepts AliasArray and AliasArrayOutput values.
// You can construct a concrete instance of `AliasArrayInput` via:
//
//	AliasArray{ AliasArgs{...} }
type AliasArrayInput interface {
	pulumi.Input

	ToAliasArrayOutput() AliasArrayOutput
	ToAliasArrayOutputWithContext(context.Context) AliasArrayOutput
}

type AliasArray []AliasInput

func (AliasArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alias)(nil)).Elem()
}

func (i AliasArray) ToAliasArrayOutput() AliasArrayOutput {
	return i.ToAliasArrayOutputWithContext(context.Background())
}

func (i AliasArray) ToAliasArrayOutputWithContext(ctx context.Context) AliasArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasArrayOutput)
}

// AliasMapInput is an input type that accepts AliasMap and AliasMapOutput values.
// You can construct a concrete instance of `AliasMapInput` via:
//
//	AliasMap{ "key": AliasArgs{...} }
type AliasMapInput interface {
	pulumi.Input

	ToAliasMapOutput() AliasMapOutput
	ToAliasMapOutputWithContext(context.Context) AliasMapOutput
}

type AliasMap map[string]AliasInput

func (AliasMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alias)(nil)).Elem()
}

func (i AliasMap) ToAliasMapOutput() AliasMapOutput {
	return i.ToAliasMapOutputWithContext(context.Background())
}

func (i AliasMap) ToAliasMapOutputWithContext(ctx context.Context) AliasMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasMapOutput)
}

type AliasOutput struct{ *pulumi.OutputState }

func (AliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (o AliasOutput) ToAliasOutput() AliasOutput {
	return o
}

func (o AliasOutput) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return o
}

// Name for the alias you are creating.
func (o AliasOutput) AliasName() pulumi.StringOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringOutput { return v.AliasName }).(pulumi.StringOutput)
}

// Description of the alias.
func (o AliasOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Function Compute alias' route configuration settings. Fields documented below.
//
// **routing_config** includes the following arguments:
func (o AliasOutput) RoutingConfig() AliasRoutingConfigPtrOutput {
	return o.ApplyT(func(v *Alias) AliasRoutingConfigPtrOutput { return v.RoutingConfig }).(AliasRoutingConfigPtrOutput)
}

// The Function Compute service name.
func (o AliasOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// The Function Compute service version for which you are creating the alias. Pattern: (LATEST|[0-9]+).
func (o AliasOutput) ServiceVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringOutput { return v.ServiceVersion }).(pulumi.StringOutput)
}

type AliasArrayOutput struct{ *pulumi.OutputState }

func (AliasArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alias)(nil)).Elem()
}

func (o AliasArrayOutput) ToAliasArrayOutput() AliasArrayOutput {
	return o
}

func (o AliasArrayOutput) ToAliasArrayOutputWithContext(ctx context.Context) AliasArrayOutput {
	return o
}

func (o AliasArrayOutput) Index(i pulumi.IntInput) AliasOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Alias {
		return vs[0].([]*Alias)[vs[1].(int)]
	}).(AliasOutput)
}

type AliasMapOutput struct{ *pulumi.OutputState }

func (AliasMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alias)(nil)).Elem()
}

func (o AliasMapOutput) ToAliasMapOutput() AliasMapOutput {
	return o
}

func (o AliasMapOutput) ToAliasMapOutputWithContext(ctx context.Context) AliasMapOutput {
	return o
}

func (o AliasMapOutput) MapIndex(k pulumi.StringInput) AliasOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Alias {
		return vs[0].(map[string]*Alias)[vs[1].(string)]
	}).(AliasOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AliasInput)(nil)).Elem(), &Alias{})
	pulumi.RegisterInputType(reflect.TypeOf((*AliasArrayInput)(nil)).Elem(), AliasArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AliasMapInput)(nil)).Elem(), AliasMap{})
	pulumi.RegisterOutputType(AliasOutput{})
	pulumi.RegisterOutputType(AliasArrayOutput{})
	pulumi.RegisterOutputType(AliasMapOutput{})
}
