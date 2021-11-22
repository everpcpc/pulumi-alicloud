// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hbr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a HBR Backup vault resource.
//
// For information about HBR Backup vault and how to use it, see [What is Backup vault](https://www.alibabacloud.com/help/doc-detail/62362.htm).
//
// > **NOTE:** Available in v1.129.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/hbr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := hbr.NewVault(ctx, "example", &hbr.VaultArgs{
// 			VaultName: pulumi.String("example_value"),
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
// HBR Vault can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:hbr/vault:Vault example <id>
// ```
type Vault struct {
	pulumi.CustomResourceState

	// The description of Vault. Defaults to an empty string.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The status of Vault. Valid values: `INITIALIZING`, `CREATED`, `ERROR`, `UNKNOWN`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The name of Vault.
	VaultName pulumi.StringOutput `pulumi:"vaultName"`
	// The storage class of Vault. Valid values: `STANDARD`.
	VaultStorageClass pulumi.StringOutput `pulumi:"vaultStorageClass"`
	// The type of Vault. Valid values: `STANDARD`.
	VaultType pulumi.StringOutput `pulumi:"vaultType"`
}

// NewVault registers a new resource with the given unique name, arguments, and options.
func NewVault(ctx *pulumi.Context,
	name string, args *VaultArgs, opts ...pulumi.ResourceOption) (*Vault, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VaultName == nil {
		return nil, errors.New("invalid value for required argument 'VaultName'")
	}
	var resource Vault
	err := ctx.RegisterResource("alicloud:hbr/vault:Vault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVault gets an existing Vault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VaultState, opts ...pulumi.ResourceOption) (*Vault, error) {
	var resource Vault
	err := ctx.ReadResource("alicloud:hbr/vault:Vault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vault resources.
type vaultState struct {
	// The description of Vault. Defaults to an empty string.
	Description *string `pulumi:"description"`
	// The status of Vault. Valid values: `INITIALIZING`, `CREATED`, `ERROR`, `UNKNOWN`.
	Status *string `pulumi:"status"`
	// The name of Vault.
	VaultName *string `pulumi:"vaultName"`
	// The storage class of Vault. Valid values: `STANDARD`.
	VaultStorageClass *string `pulumi:"vaultStorageClass"`
	// The type of Vault. Valid values: `STANDARD`.
	VaultType *string `pulumi:"vaultType"`
}

type VaultState struct {
	// The description of Vault. Defaults to an empty string.
	Description pulumi.StringPtrInput
	// The status of Vault. Valid values: `INITIALIZING`, `CREATED`, `ERROR`, `UNKNOWN`.
	Status pulumi.StringPtrInput
	// The name of Vault.
	VaultName pulumi.StringPtrInput
	// The storage class of Vault. Valid values: `STANDARD`.
	VaultStorageClass pulumi.StringPtrInput
	// The type of Vault. Valid values: `STANDARD`.
	VaultType pulumi.StringPtrInput
}

func (VaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultState)(nil)).Elem()
}

type vaultArgs struct {
	// The description of Vault. Defaults to an empty string.
	Description *string `pulumi:"description"`
	// The name of Vault.
	VaultName string `pulumi:"vaultName"`
	// The storage class of Vault. Valid values: `STANDARD`.
	VaultStorageClass *string `pulumi:"vaultStorageClass"`
	// The type of Vault. Valid values: `STANDARD`.
	VaultType *string `pulumi:"vaultType"`
}

// The set of arguments for constructing a Vault resource.
type VaultArgs struct {
	// The description of Vault. Defaults to an empty string.
	Description pulumi.StringPtrInput
	// The name of Vault.
	VaultName pulumi.StringInput
	// The storage class of Vault. Valid values: `STANDARD`.
	VaultStorageClass pulumi.StringPtrInput
	// The type of Vault. Valid values: `STANDARD`.
	VaultType pulumi.StringPtrInput
}

func (VaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultArgs)(nil)).Elem()
}

type VaultInput interface {
	pulumi.Input

	ToVaultOutput() VaultOutput
	ToVaultOutputWithContext(ctx context.Context) VaultOutput
}

func (*Vault) ElementType() reflect.Type {
	return reflect.TypeOf((*Vault)(nil))
}

func (i *Vault) ToVaultOutput() VaultOutput {
	return i.ToVaultOutputWithContext(context.Background())
}

func (i *Vault) ToVaultOutputWithContext(ctx context.Context) VaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultOutput)
}

func (i *Vault) ToVaultPtrOutput() VaultPtrOutput {
	return i.ToVaultPtrOutputWithContext(context.Background())
}

func (i *Vault) ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPtrOutput)
}

type VaultPtrInput interface {
	pulumi.Input

	ToVaultPtrOutput() VaultPtrOutput
	ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput
}

type vaultPtrType VaultArgs

func (*vaultPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Vault)(nil))
}

func (i *vaultPtrType) ToVaultPtrOutput() VaultPtrOutput {
	return i.ToVaultPtrOutputWithContext(context.Background())
}

func (i *vaultPtrType) ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPtrOutput)
}

// VaultArrayInput is an input type that accepts VaultArray and VaultArrayOutput values.
// You can construct a concrete instance of `VaultArrayInput` via:
//
//          VaultArray{ VaultArgs{...} }
type VaultArrayInput interface {
	pulumi.Input

	ToVaultArrayOutput() VaultArrayOutput
	ToVaultArrayOutputWithContext(context.Context) VaultArrayOutput
}

type VaultArray []VaultInput

func (VaultArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vault)(nil)).Elem()
}

func (i VaultArray) ToVaultArrayOutput() VaultArrayOutput {
	return i.ToVaultArrayOutputWithContext(context.Background())
}

func (i VaultArray) ToVaultArrayOutputWithContext(ctx context.Context) VaultArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultArrayOutput)
}

// VaultMapInput is an input type that accepts VaultMap and VaultMapOutput values.
// You can construct a concrete instance of `VaultMapInput` via:
//
//          VaultMap{ "key": VaultArgs{...} }
type VaultMapInput interface {
	pulumi.Input

	ToVaultMapOutput() VaultMapOutput
	ToVaultMapOutputWithContext(context.Context) VaultMapOutput
}

type VaultMap map[string]VaultInput

func (VaultMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vault)(nil)).Elem()
}

func (i VaultMap) ToVaultMapOutput() VaultMapOutput {
	return i.ToVaultMapOutputWithContext(context.Background())
}

func (i VaultMap) ToVaultMapOutputWithContext(ctx context.Context) VaultMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultMapOutput)
}

type VaultOutput struct{ *pulumi.OutputState }

func (VaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Vault)(nil))
}

func (o VaultOutput) ToVaultOutput() VaultOutput {
	return o
}

func (o VaultOutput) ToVaultOutputWithContext(ctx context.Context) VaultOutput {
	return o
}

func (o VaultOutput) ToVaultPtrOutput() VaultPtrOutput {
	return o.ToVaultPtrOutputWithContext(context.Background())
}

func (o VaultOutput) ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Vault) *Vault {
		return &v
	}).(VaultPtrOutput)
}

type VaultPtrOutput struct{ *pulumi.OutputState }

func (VaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vault)(nil))
}

func (o VaultPtrOutput) ToVaultPtrOutput() VaultPtrOutput {
	return o
}

func (o VaultPtrOutput) ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput {
	return o
}

func (o VaultPtrOutput) Elem() VaultOutput {
	return o.ApplyT(func(v *Vault) Vault {
		if v != nil {
			return *v
		}
		var ret Vault
		return ret
	}).(VaultOutput)
}

type VaultArrayOutput struct{ *pulumi.OutputState }

func (VaultArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Vault)(nil))
}

func (o VaultArrayOutput) ToVaultArrayOutput() VaultArrayOutput {
	return o
}

func (o VaultArrayOutput) ToVaultArrayOutputWithContext(ctx context.Context) VaultArrayOutput {
	return o
}

func (o VaultArrayOutput) Index(i pulumi.IntInput) VaultOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Vault {
		return vs[0].([]Vault)[vs[1].(int)]
	}).(VaultOutput)
}

type VaultMapOutput struct{ *pulumi.OutputState }

func (VaultMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Vault)(nil))
}

func (o VaultMapOutput) ToVaultMapOutput() VaultMapOutput {
	return o
}

func (o VaultMapOutput) ToVaultMapOutputWithContext(ctx context.Context) VaultMapOutput {
	return o
}

func (o VaultMapOutput) MapIndex(k pulumi.StringInput) VaultOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Vault {
		return vs[0].(map[string]Vault)[vs[1].(string)]
	}).(VaultOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VaultInput)(nil)).Elem(), &Vault{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultPtrInput)(nil)).Elem(), &Vault{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultArrayInput)(nil)).Elem(), VaultArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultMapInput)(nil)).Elem(), VaultMap{})
	pulumi.RegisterOutputType(VaultOutput{})
	pulumi.RegisterOutputType(VaultPtrOutput{})
	pulumi.RegisterOutputType(VaultArrayOutput{})
	pulumi.RegisterOutputType(VaultMapOutput{})
}
