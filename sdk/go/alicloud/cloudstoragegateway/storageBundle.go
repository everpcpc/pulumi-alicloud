// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudstoragegateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Storage Gateway Storage Bundle resource.
//
// For information about Cloud Storage Gateway Storage Bundle and how to use it, see [What is Storage Bundle](https://www.alibabacloud.com/help/en/doc-detail/53972.htm).
//
// > **NOTE:** Available in v1.116.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cloudstoragegateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudstoragegateway.NewStorageBundle(ctx, "example", &cloudstoragegateway.StorageBundleArgs{
// 			StorageBundleName: pulumi.String("example_value"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Cloud Storage Gateway Storage Bundle can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:cloudstoragegateway/storageBundle:StorageBundle example <id>
// ```
type StorageBundle struct {
	pulumi.CustomResourceState

	// The description of storage bundle.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of storage bundle.
	StorageBundleName pulumi.StringOutput `pulumi:"storageBundleName"`
}

// NewStorageBundle registers a new resource with the given unique name, arguments, and options.
func NewStorageBundle(ctx *pulumi.Context,
	name string, args *StorageBundleArgs, opts ...pulumi.ResourceOption) (*StorageBundle, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StorageBundleName == nil {
		return nil, errors.New("invalid value for required argument 'StorageBundleName'")
	}
	var resource StorageBundle
	err := ctx.RegisterResource("alicloud:cloudstoragegateway/storageBundle:StorageBundle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageBundle gets an existing StorageBundle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageBundle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageBundleState, opts ...pulumi.ResourceOption) (*StorageBundle, error) {
	var resource StorageBundle
	err := ctx.ReadResource("alicloud:cloudstoragegateway/storageBundle:StorageBundle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageBundle resources.
type storageBundleState struct {
	// The description of storage bundle.
	Description *string `pulumi:"description"`
	// The name of storage bundle.
	StorageBundleName *string `pulumi:"storageBundleName"`
}

type StorageBundleState struct {
	// The description of storage bundle.
	Description pulumi.StringPtrInput
	// The name of storage bundle.
	StorageBundleName pulumi.StringPtrInput
}

func (StorageBundleState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBundleState)(nil)).Elem()
}

type storageBundleArgs struct {
	// The description of storage bundle.
	Description *string `pulumi:"description"`
	// The name of storage bundle.
	StorageBundleName string `pulumi:"storageBundleName"`
}

// The set of arguments for constructing a StorageBundle resource.
type StorageBundleArgs struct {
	// The description of storage bundle.
	Description pulumi.StringPtrInput
	// The name of storage bundle.
	StorageBundleName pulumi.StringInput
}

func (StorageBundleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBundleArgs)(nil)).Elem()
}

type StorageBundleInput interface {
	pulumi.Input

	ToStorageBundleOutput() StorageBundleOutput
	ToStorageBundleOutputWithContext(ctx context.Context) StorageBundleOutput
}

func (*StorageBundle) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageBundle)(nil))
}

func (i *StorageBundle) ToStorageBundleOutput() StorageBundleOutput {
	return i.ToStorageBundleOutputWithContext(context.Background())
}

func (i *StorageBundle) ToStorageBundleOutputWithContext(ctx context.Context) StorageBundleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBundleOutput)
}

func (i *StorageBundle) ToStorageBundlePtrOutput() StorageBundlePtrOutput {
	return i.ToStorageBundlePtrOutputWithContext(context.Background())
}

func (i *StorageBundle) ToStorageBundlePtrOutputWithContext(ctx context.Context) StorageBundlePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBundlePtrOutput)
}

type StorageBundlePtrInput interface {
	pulumi.Input

	ToStorageBundlePtrOutput() StorageBundlePtrOutput
	ToStorageBundlePtrOutputWithContext(ctx context.Context) StorageBundlePtrOutput
}

type storageBundlePtrType StorageBundleArgs

func (*storageBundlePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBundle)(nil))
}

func (i *storageBundlePtrType) ToStorageBundlePtrOutput() StorageBundlePtrOutput {
	return i.ToStorageBundlePtrOutputWithContext(context.Background())
}

func (i *storageBundlePtrType) ToStorageBundlePtrOutputWithContext(ctx context.Context) StorageBundlePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBundlePtrOutput)
}

// StorageBundleArrayInput is an input type that accepts StorageBundleArray and StorageBundleArrayOutput values.
// You can construct a concrete instance of `StorageBundleArrayInput` via:
//
//          StorageBundleArray{ StorageBundleArgs{...} }
type StorageBundleArrayInput interface {
	pulumi.Input

	ToStorageBundleArrayOutput() StorageBundleArrayOutput
	ToStorageBundleArrayOutputWithContext(context.Context) StorageBundleArrayOutput
}

type StorageBundleArray []StorageBundleInput

func (StorageBundleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageBundle)(nil)).Elem()
}

func (i StorageBundleArray) ToStorageBundleArrayOutput() StorageBundleArrayOutput {
	return i.ToStorageBundleArrayOutputWithContext(context.Background())
}

func (i StorageBundleArray) ToStorageBundleArrayOutputWithContext(ctx context.Context) StorageBundleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBundleArrayOutput)
}

// StorageBundleMapInput is an input type that accepts StorageBundleMap and StorageBundleMapOutput values.
// You can construct a concrete instance of `StorageBundleMapInput` via:
//
//          StorageBundleMap{ "key": StorageBundleArgs{...} }
type StorageBundleMapInput interface {
	pulumi.Input

	ToStorageBundleMapOutput() StorageBundleMapOutput
	ToStorageBundleMapOutputWithContext(context.Context) StorageBundleMapOutput
}

type StorageBundleMap map[string]StorageBundleInput

func (StorageBundleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageBundle)(nil)).Elem()
}

func (i StorageBundleMap) ToStorageBundleMapOutput() StorageBundleMapOutput {
	return i.ToStorageBundleMapOutputWithContext(context.Background())
}

func (i StorageBundleMap) ToStorageBundleMapOutputWithContext(ctx context.Context) StorageBundleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBundleMapOutput)
}

type StorageBundleOutput struct{ *pulumi.OutputState }

func (StorageBundleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageBundle)(nil))
}

func (o StorageBundleOutput) ToStorageBundleOutput() StorageBundleOutput {
	return o
}

func (o StorageBundleOutput) ToStorageBundleOutputWithContext(ctx context.Context) StorageBundleOutput {
	return o
}

func (o StorageBundleOutput) ToStorageBundlePtrOutput() StorageBundlePtrOutput {
	return o.ToStorageBundlePtrOutputWithContext(context.Background())
}

func (o StorageBundleOutput) ToStorageBundlePtrOutputWithContext(ctx context.Context) StorageBundlePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageBundle) *StorageBundle {
		return &v
	}).(StorageBundlePtrOutput)
}

type StorageBundlePtrOutput struct{ *pulumi.OutputState }

func (StorageBundlePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBundle)(nil))
}

func (o StorageBundlePtrOutput) ToStorageBundlePtrOutput() StorageBundlePtrOutput {
	return o
}

func (o StorageBundlePtrOutput) ToStorageBundlePtrOutputWithContext(ctx context.Context) StorageBundlePtrOutput {
	return o
}

func (o StorageBundlePtrOutput) Elem() StorageBundleOutput {
	return o.ApplyT(func(v *StorageBundle) StorageBundle {
		if v != nil {
			return *v
		}
		var ret StorageBundle
		return ret
	}).(StorageBundleOutput)
}

type StorageBundleArrayOutput struct{ *pulumi.OutputState }

func (StorageBundleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageBundle)(nil))
}

func (o StorageBundleArrayOutput) ToStorageBundleArrayOutput() StorageBundleArrayOutput {
	return o
}

func (o StorageBundleArrayOutput) ToStorageBundleArrayOutputWithContext(ctx context.Context) StorageBundleArrayOutput {
	return o
}

func (o StorageBundleArrayOutput) Index(i pulumi.IntInput) StorageBundleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageBundle {
		return vs[0].([]StorageBundle)[vs[1].(int)]
	}).(StorageBundleOutput)
}

type StorageBundleMapOutput struct{ *pulumi.OutputState }

func (StorageBundleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StorageBundle)(nil))
}

func (o StorageBundleMapOutput) ToStorageBundleMapOutput() StorageBundleMapOutput {
	return o
}

func (o StorageBundleMapOutput) ToStorageBundleMapOutputWithContext(ctx context.Context) StorageBundleMapOutput {
	return o
}

func (o StorageBundleMapOutput) MapIndex(k pulumi.StringInput) StorageBundleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StorageBundle {
		return vs[0].(map[string]StorageBundle)[vs[1].(string)]
	}).(StorageBundleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBundleInput)(nil)).Elem(), &StorageBundle{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBundlePtrInput)(nil)).Elem(), &StorageBundle{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBundleArrayInput)(nil)).Elem(), StorageBundleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBundleMapInput)(nil)).Elem(), StorageBundleMap{})
	pulumi.RegisterOutputType(StorageBundleOutput{})
	pulumi.RegisterOutputType(StorageBundlePtrOutput{})
	pulumi.RegisterOutputType(StorageBundleArrayOutput{})
	pulumi.RegisterOutputType(StorageBundleMapOutput{})
}
