// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sae

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Serverless App Engine (SAE) Config Map resource.
//
// For information about Serverless App Engine (SAE) Config Map and how to use it, see [What is Config Map](https://help.aliyun.com/document_detail/97792.html).
//
// > **NOTE:** Available in v1.130.0+.
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
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/sae"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			configMapName := "examplename"
//			if param := cfg.Get("configMapName"); param != "" {
//				configMapName = param
//			}
//			exampleNamespace, err := sae.NewNamespace(ctx, "exampleNamespace", &sae.NamespaceArgs{
//				NamespaceId:          pulumi.String("cn-hangzhou:yourname"),
//				NamespaceName:        pulumi.String("example_value"),
//				NamespaceDescription: pulumi.String("your_description"),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"env.home":  "/root",
//				"env.shell": "/bin/sh",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = sae.NewConfigMap(ctx, "exampleConfigMap", &sae.ConfigMapArgs{
//				Data:        pulumi.String(json0),
//				NamespaceId: exampleNamespace.NamespaceId,
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
// Serverless App Engine (SAE) Config Map can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:sae/configMap:ConfigMap example <id>
//
// ```
type ConfigMap struct {
	pulumi.CustomResourceState

	// ConfigMap instance data.
	Data pulumi.StringOutput `pulumi:"data"`
	// The Description of ConfigMap.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// ConfigMap instance name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The NamespaceId of ConfigMap.It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
}

// NewConfigMap registers a new resource with the given unique name, arguments, and options.
func NewConfigMap(ctx *pulumi.Context,
	name string, args *ConfigMapArgs, opts ...pulumi.ResourceOption) (*ConfigMap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Data == nil {
		return nil, errors.New("invalid value for required argument 'Data'")
	}
	if args.NamespaceId == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceId'")
	}
	var resource ConfigMap
	err := ctx.RegisterResource("alicloud:sae/configMap:ConfigMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigMap gets an existing ConfigMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigMapState, opts ...pulumi.ResourceOption) (*ConfigMap, error) {
	var resource ConfigMap
	err := ctx.ReadResource("alicloud:sae/configMap:ConfigMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigMap resources.
type configMapState struct {
	// ConfigMap instance data.
	Data *string `pulumi:"data"`
	// The Description of ConfigMap.
	Description *string `pulumi:"description"`
	// ConfigMap instance name.
	Name *string `pulumi:"name"`
	// The NamespaceId of ConfigMap.It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`
	NamespaceId *string `pulumi:"namespaceId"`
}

type ConfigMapState struct {
	// ConfigMap instance data.
	Data pulumi.StringPtrInput
	// The Description of ConfigMap.
	Description pulumi.StringPtrInput
	// ConfigMap instance name.
	Name pulumi.StringPtrInput
	// The NamespaceId of ConfigMap.It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`
	NamespaceId pulumi.StringPtrInput
}

func (ConfigMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*configMapState)(nil)).Elem()
}

type configMapArgs struct {
	// ConfigMap instance data.
	Data string `pulumi:"data"`
	// The Description of ConfigMap.
	Description *string `pulumi:"description"`
	// ConfigMap instance name.
	Name *string `pulumi:"name"`
	// The NamespaceId of ConfigMap.It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`
	NamespaceId string `pulumi:"namespaceId"`
}

// The set of arguments for constructing a ConfigMap resource.
type ConfigMapArgs struct {
	// ConfigMap instance data.
	Data pulumi.StringInput
	// The Description of ConfigMap.
	Description pulumi.StringPtrInput
	// ConfigMap instance name.
	Name pulumi.StringPtrInput
	// The NamespaceId of ConfigMap.It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`
	NamespaceId pulumi.StringInput
}

func (ConfigMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configMapArgs)(nil)).Elem()
}

type ConfigMapInput interface {
	pulumi.Input

	ToConfigMapOutput() ConfigMapOutput
	ToConfigMapOutputWithContext(ctx context.Context) ConfigMapOutput
}

func (*ConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigMap)(nil)).Elem()
}

func (i *ConfigMap) ToConfigMapOutput() ConfigMapOutput {
	return i.ToConfigMapOutputWithContext(context.Background())
}

func (i *ConfigMap) ToConfigMapOutputWithContext(ctx context.Context) ConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapOutput)
}

// ConfigMapArrayInput is an input type that accepts ConfigMapArray and ConfigMapArrayOutput values.
// You can construct a concrete instance of `ConfigMapArrayInput` via:
//
//	ConfigMapArray{ ConfigMapArgs{...} }
type ConfigMapArrayInput interface {
	pulumi.Input

	ToConfigMapArrayOutput() ConfigMapArrayOutput
	ToConfigMapArrayOutputWithContext(context.Context) ConfigMapArrayOutput
}

type ConfigMapArray []ConfigMapInput

func (ConfigMapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigMap)(nil)).Elem()
}

func (i ConfigMapArray) ToConfigMapArrayOutput() ConfigMapArrayOutput {
	return i.ToConfigMapArrayOutputWithContext(context.Background())
}

func (i ConfigMapArray) ToConfigMapArrayOutputWithContext(ctx context.Context) ConfigMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapArrayOutput)
}

// ConfigMapMapInput is an input type that accepts ConfigMapMap and ConfigMapMapOutput values.
// You can construct a concrete instance of `ConfigMapMapInput` via:
//
//	ConfigMapMap{ "key": ConfigMapArgs{...} }
type ConfigMapMapInput interface {
	pulumi.Input

	ToConfigMapMapOutput() ConfigMapMapOutput
	ToConfigMapMapOutputWithContext(context.Context) ConfigMapMapOutput
}

type ConfigMapMap map[string]ConfigMapInput

func (ConfigMapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigMap)(nil)).Elem()
}

func (i ConfigMapMap) ToConfigMapMapOutput() ConfigMapMapOutput {
	return i.ToConfigMapMapOutputWithContext(context.Background())
}

func (i ConfigMapMap) ToConfigMapMapOutputWithContext(ctx context.Context) ConfigMapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapMapOutput)
}

type ConfigMapOutput struct{ *pulumi.OutputState }

func (ConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigMap)(nil)).Elem()
}

func (o ConfigMapOutput) ToConfigMapOutput() ConfigMapOutput {
	return o
}

func (o ConfigMapOutput) ToConfigMapOutputWithContext(ctx context.Context) ConfigMapOutput {
	return o
}

// ConfigMap instance data.
func (o ConfigMapOutput) Data() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigMap) pulumi.StringOutput { return v.Data }).(pulumi.StringOutput)
}

// The Description of ConfigMap.
func (o ConfigMapOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigMap) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// ConfigMap instance name.
func (o ConfigMapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigMap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The NamespaceId of ConfigMap.It can contain 2 to 32 lowercase characters.The value is in format `{RegionId}:{namespace}`
func (o ConfigMapOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigMap) pulumi.StringOutput { return v.NamespaceId }).(pulumi.StringOutput)
}

type ConfigMapArrayOutput struct{ *pulumi.OutputState }

func (ConfigMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigMap)(nil)).Elem()
}

func (o ConfigMapArrayOutput) ToConfigMapArrayOutput() ConfigMapArrayOutput {
	return o
}

func (o ConfigMapArrayOutput) ToConfigMapArrayOutputWithContext(ctx context.Context) ConfigMapArrayOutput {
	return o
}

func (o ConfigMapArrayOutput) Index(i pulumi.IntInput) ConfigMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConfigMap {
		return vs[0].([]*ConfigMap)[vs[1].(int)]
	}).(ConfigMapOutput)
}

type ConfigMapMapOutput struct{ *pulumi.OutputState }

func (ConfigMapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigMap)(nil)).Elem()
}

func (o ConfigMapMapOutput) ToConfigMapMapOutput() ConfigMapMapOutput {
	return o
}

func (o ConfigMapMapOutput) ToConfigMapMapOutputWithContext(ctx context.Context) ConfigMapMapOutput {
	return o
}

func (o ConfigMapMapOutput) MapIndex(k pulumi.StringInput) ConfigMapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConfigMap {
		return vs[0].(map[string]*ConfigMap)[vs[1].(string)]
	}).(ConfigMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapInput)(nil)).Elem(), &ConfigMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapArrayInput)(nil)).Elem(), ConfigMapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapMapInput)(nil)).Elem(), ConfigMapMap{})
	pulumi.RegisterOutputType(ConfigMapOutput{})
	pulumi.RegisterOutputType(ConfigMapArrayOutput{})
	pulumi.RegisterOutputType(ConfigMapMapOutput{})
}
