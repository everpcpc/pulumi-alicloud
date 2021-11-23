// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a ECD Network Package resource.
//
// For information about ECD Network Package and how to use it, see [What is Network Package](https://help.aliyun.com/document_detail/188382.html).
//
// > **NOTE:** Available in v1.142.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/eds"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := eds.NewSimpleOfficeSite(ctx, "_default", &eds.SimpleOfficeSiteArgs{
// 			CidrBlock:         pulumi.String("172.16.0.0/12"),
// 			DesktopAccessType: pulumi.String("Internet"),
// 			OfficeSiteName:    pulumi.String("your_office_site_name"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = eds.NewNetworkPackage(ctx, "example", &eds.NetworkPackageArgs{
// 			Bandwidth:    pulumi.Int(10),
// 			OfficeSiteId: _default.ID(),
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
// ECD Network Package can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:eds/networkPackage:NetworkPackage example <id>
// ```
type NetworkPackage struct {
	pulumi.CustomResourceState

	// The bandwidth of package public network bandwidth peak. Valid values: 1~200. Unit:Mbps.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// The internet charge type  of  package.
	InternetChargeType pulumi.StringOutput `pulumi:"internetChargeType"`
	// The ID of office site.
	OfficeSiteId pulumi.StringOutput `pulumi:"officeSiteId"`
	// The status of network package. Valid values: `Creating`, `InUse`, `Releasing`,`Released`.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewNetworkPackage registers a new resource with the given unique name, arguments, and options.
func NewNetworkPackage(ctx *pulumi.Context,
	name string, args *NetworkPackageArgs, opts ...pulumi.ResourceOption) (*NetworkPackage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bandwidth == nil {
		return nil, errors.New("invalid value for required argument 'Bandwidth'")
	}
	if args.OfficeSiteId == nil {
		return nil, errors.New("invalid value for required argument 'OfficeSiteId'")
	}
	var resource NetworkPackage
	err := ctx.RegisterResource("alicloud:eds/networkPackage:NetworkPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkPackage gets an existing NetworkPackage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkPackageState, opts ...pulumi.ResourceOption) (*NetworkPackage, error) {
	var resource NetworkPackage
	err := ctx.ReadResource("alicloud:eds/networkPackage:NetworkPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkPackage resources.
type networkPackageState struct {
	// The bandwidth of package public network bandwidth peak. Valid values: 1~200. Unit:Mbps.
	Bandwidth *int `pulumi:"bandwidth"`
	// The internet charge type  of  package.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The ID of office site.
	OfficeSiteId *string `pulumi:"officeSiteId"`
	// The status of network package. Valid values: `Creating`, `InUse`, `Releasing`,`Released`.
	Status *string `pulumi:"status"`
}

type NetworkPackageState struct {
	// The bandwidth of package public network bandwidth peak. Valid values: 1~200. Unit:Mbps.
	Bandwidth pulumi.IntPtrInput
	// The internet charge type  of  package.
	InternetChargeType pulumi.StringPtrInput
	// The ID of office site.
	OfficeSiteId pulumi.StringPtrInput
	// The status of network package. Valid values: `Creating`, `InUse`, `Releasing`,`Released`.
	Status pulumi.StringPtrInput
}

func (NetworkPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPackageState)(nil)).Elem()
}

type networkPackageArgs struct {
	// The bandwidth of package public network bandwidth peak. Valid values: 1~200. Unit:Mbps.
	Bandwidth int `pulumi:"bandwidth"`
	// The ID of office site.
	OfficeSiteId string `pulumi:"officeSiteId"`
}

// The set of arguments for constructing a NetworkPackage resource.
type NetworkPackageArgs struct {
	// The bandwidth of package public network bandwidth peak. Valid values: 1~200. Unit:Mbps.
	Bandwidth pulumi.IntInput
	// The ID of office site.
	OfficeSiteId pulumi.StringInput
}

func (NetworkPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPackageArgs)(nil)).Elem()
}

type NetworkPackageInput interface {
	pulumi.Input

	ToNetworkPackageOutput() NetworkPackageOutput
	ToNetworkPackageOutputWithContext(ctx context.Context) NetworkPackageOutput
}

func (*NetworkPackage) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPackage)(nil))
}

func (i *NetworkPackage) ToNetworkPackageOutput() NetworkPackageOutput {
	return i.ToNetworkPackageOutputWithContext(context.Background())
}

func (i *NetworkPackage) ToNetworkPackageOutputWithContext(ctx context.Context) NetworkPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPackageOutput)
}

func (i *NetworkPackage) ToNetworkPackagePtrOutput() NetworkPackagePtrOutput {
	return i.ToNetworkPackagePtrOutputWithContext(context.Background())
}

func (i *NetworkPackage) ToNetworkPackagePtrOutputWithContext(ctx context.Context) NetworkPackagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPackagePtrOutput)
}

type NetworkPackagePtrInput interface {
	pulumi.Input

	ToNetworkPackagePtrOutput() NetworkPackagePtrOutput
	ToNetworkPackagePtrOutputWithContext(ctx context.Context) NetworkPackagePtrOutput
}

type networkPackagePtrType NetworkPackageArgs

func (*networkPackagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPackage)(nil))
}

func (i *networkPackagePtrType) ToNetworkPackagePtrOutput() NetworkPackagePtrOutput {
	return i.ToNetworkPackagePtrOutputWithContext(context.Background())
}

func (i *networkPackagePtrType) ToNetworkPackagePtrOutputWithContext(ctx context.Context) NetworkPackagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPackagePtrOutput)
}

// NetworkPackageArrayInput is an input type that accepts NetworkPackageArray and NetworkPackageArrayOutput values.
// You can construct a concrete instance of `NetworkPackageArrayInput` via:
//
//          NetworkPackageArray{ NetworkPackageArgs{...} }
type NetworkPackageArrayInput interface {
	pulumi.Input

	ToNetworkPackageArrayOutput() NetworkPackageArrayOutput
	ToNetworkPackageArrayOutputWithContext(context.Context) NetworkPackageArrayOutput
}

type NetworkPackageArray []NetworkPackageInput

func (NetworkPackageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkPackage)(nil)).Elem()
}

func (i NetworkPackageArray) ToNetworkPackageArrayOutput() NetworkPackageArrayOutput {
	return i.ToNetworkPackageArrayOutputWithContext(context.Background())
}

func (i NetworkPackageArray) ToNetworkPackageArrayOutputWithContext(ctx context.Context) NetworkPackageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPackageArrayOutput)
}

// NetworkPackageMapInput is an input type that accepts NetworkPackageMap and NetworkPackageMapOutput values.
// You can construct a concrete instance of `NetworkPackageMapInput` via:
//
//          NetworkPackageMap{ "key": NetworkPackageArgs{...} }
type NetworkPackageMapInput interface {
	pulumi.Input

	ToNetworkPackageMapOutput() NetworkPackageMapOutput
	ToNetworkPackageMapOutputWithContext(context.Context) NetworkPackageMapOutput
}

type NetworkPackageMap map[string]NetworkPackageInput

func (NetworkPackageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkPackage)(nil)).Elem()
}

func (i NetworkPackageMap) ToNetworkPackageMapOutput() NetworkPackageMapOutput {
	return i.ToNetworkPackageMapOutputWithContext(context.Background())
}

func (i NetworkPackageMap) ToNetworkPackageMapOutputWithContext(ctx context.Context) NetworkPackageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPackageMapOutput)
}

type NetworkPackageOutput struct{ *pulumi.OutputState }

func (NetworkPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPackage)(nil))
}

func (o NetworkPackageOutput) ToNetworkPackageOutput() NetworkPackageOutput {
	return o
}

func (o NetworkPackageOutput) ToNetworkPackageOutputWithContext(ctx context.Context) NetworkPackageOutput {
	return o
}

func (o NetworkPackageOutput) ToNetworkPackagePtrOutput() NetworkPackagePtrOutput {
	return o.ToNetworkPackagePtrOutputWithContext(context.Background())
}

func (o NetworkPackageOutput) ToNetworkPackagePtrOutputWithContext(ctx context.Context) NetworkPackagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkPackage) *NetworkPackage {
		return &v
	}).(NetworkPackagePtrOutput)
}

type NetworkPackagePtrOutput struct{ *pulumi.OutputState }

func (NetworkPackagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPackage)(nil))
}

func (o NetworkPackagePtrOutput) ToNetworkPackagePtrOutput() NetworkPackagePtrOutput {
	return o
}

func (o NetworkPackagePtrOutput) ToNetworkPackagePtrOutputWithContext(ctx context.Context) NetworkPackagePtrOutput {
	return o
}

func (o NetworkPackagePtrOutput) Elem() NetworkPackageOutput {
	return o.ApplyT(func(v *NetworkPackage) NetworkPackage {
		if v != nil {
			return *v
		}
		var ret NetworkPackage
		return ret
	}).(NetworkPackageOutput)
}

type NetworkPackageArrayOutput struct{ *pulumi.OutputState }

func (NetworkPackageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkPackage)(nil))
}

func (o NetworkPackageArrayOutput) ToNetworkPackageArrayOutput() NetworkPackageArrayOutput {
	return o
}

func (o NetworkPackageArrayOutput) ToNetworkPackageArrayOutputWithContext(ctx context.Context) NetworkPackageArrayOutput {
	return o
}

func (o NetworkPackageArrayOutput) Index(i pulumi.IntInput) NetworkPackageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkPackage {
		return vs[0].([]NetworkPackage)[vs[1].(int)]
	}).(NetworkPackageOutput)
}

type NetworkPackageMapOutput struct{ *pulumi.OutputState }

func (NetworkPackageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkPackage)(nil))
}

func (o NetworkPackageMapOutput) ToNetworkPackageMapOutput() NetworkPackageMapOutput {
	return o
}

func (o NetworkPackageMapOutput) ToNetworkPackageMapOutputWithContext(ctx context.Context) NetworkPackageMapOutput {
	return o
}

func (o NetworkPackageMapOutput) MapIndex(k pulumi.StringInput) NetworkPackageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkPackage {
		return vs[0].(map[string]NetworkPackage)[vs[1].(string)]
	}).(NetworkPackageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPackageInput)(nil)).Elem(), &NetworkPackage{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPackagePtrInput)(nil)).Elem(), &NetworkPackage{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPackageArrayInput)(nil)).Elem(), NetworkPackageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPackageMapInput)(nil)).Elem(), NetworkPackageMap{})
	pulumi.RegisterOutputType(NetworkPackageOutput{})
	pulumi.RegisterOutputType(NetworkPackagePtrOutput{})
	pulumi.RegisterOutputType(NetworkPackageArrayOutput{})
	pulumi.RegisterOutputType(NetworkPackageMapOutput{})
}