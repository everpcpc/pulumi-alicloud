// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a MSE Cluster resource. It is a one-stop microservice platform for the industry's mainstream open source microservice frameworks Spring Cloud and Dubbo, providing governance center, managed registry and managed configuration center.
//
// > **NOTE:** Available in 1.94.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/mse"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := mse.NewCluster(ctx, "example", &mse.ClusterArgs{
// 			AclEntryLists: pulumi.StringArray{
// 				pulumi.String("127.0.0.1/32"),
// 			},
// 			ClusterAliasName:     pulumi.String("tf-testAccMseCluster"),
// 			ClusterSpecification: pulumi.String("MSE_SC_1_2_200_c"),
// 			ClusterType:          pulumi.String("Eureka"),
// 			ClusterVersion:       pulumi.String("EUREKA_1_9_3"),
// 			InstanceCount:        pulumi.Int(1),
// 			NetType:              pulumi.String("privatenet"),
// 			PubNetworkFlow:       pulumi.String("1"),
// 			VswitchId:            pulumi.String("vsw-123456"),
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
// MSE Cluster can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:mse/cluster:Cluster example mse-cn-0d9xxxx
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// The whitelist.
	AclEntryLists pulumi.StringArrayOutput `pulumi:"aclEntryLists"`
	// The alias of MSE Cluster.
	ClusterAliasName pulumi.StringPtrOutput `pulumi:"clusterAliasName"`
	// The engine specification of MSE Cluster. Valid values:
	// `MSE_SC_1_2_200_c`：1C2G
	// `MSE_SC_2_4_200_c`：2C4G
	// `MSE_SC_4_8_200_c`：4C8G
	// `MSE_SC_8_16_200_c`：8C16G
	ClusterSpecification pulumi.StringOutput `pulumi:"clusterSpecification"`
	// The type of MSE Cluster.
	ClusterType pulumi.StringOutput `pulumi:"clusterType"`
	// The version of MSE Cluster.
	ClusterVersion pulumi.StringOutput `pulumi:"clusterVersion"`
	// The type of Disk.
	DiskType pulumi.StringPtrOutput `pulumi:"diskType"`
	// The count of instance.
	InstanceCount pulumi.IntOutput `pulumi:"instanceCount"`
	// The type of network. Valid values: "privatenet" and "pubnet".
	NetType pulumi.StringOutput `pulumi:"netType"`
	// The specification of private network SLB.
	PrivateSlbSpecification pulumi.StringPtrOutput `pulumi:"privateSlbSpecification"`
	// The public network bandwidth. `0` means no access to the public network.
	PubNetworkFlow pulumi.StringPtrOutput `pulumi:"pubNetworkFlow"`
	// The specification of public network SLB.
	PubSlbSpecification pulumi.StringPtrOutput `pulumi:"pubSlbSpecification"`
	// The status of MSE Cluster.
	Status pulumi.StringOutput `pulumi:"status"`
	// The id of VSwitch.
	VswitchId pulumi.StringPtrOutput `pulumi:"vswitchId"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterSpecification == nil {
		return nil, errors.New("invalid value for required argument 'ClusterSpecification'")
	}
	if args.ClusterType == nil {
		return nil, errors.New("invalid value for required argument 'ClusterType'")
	}
	if args.ClusterVersion == nil {
		return nil, errors.New("invalid value for required argument 'ClusterVersion'")
	}
	if args.InstanceCount == nil {
		return nil, errors.New("invalid value for required argument 'InstanceCount'")
	}
	if args.NetType == nil {
		return nil, errors.New("invalid value for required argument 'NetType'")
	}
	var resource Cluster
	err := ctx.RegisterResource("alicloud:mse/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("alicloud:mse/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The whitelist.
	AclEntryLists []string `pulumi:"aclEntryLists"`
	// The alias of MSE Cluster.
	ClusterAliasName *string `pulumi:"clusterAliasName"`
	// The engine specification of MSE Cluster. Valid values:
	// `MSE_SC_1_2_200_c`：1C2G
	// `MSE_SC_2_4_200_c`：2C4G
	// `MSE_SC_4_8_200_c`：4C8G
	// `MSE_SC_8_16_200_c`：8C16G
	ClusterSpecification *string `pulumi:"clusterSpecification"`
	// The type of MSE Cluster.
	ClusterType *string `pulumi:"clusterType"`
	// The version of MSE Cluster.
	ClusterVersion *string `pulumi:"clusterVersion"`
	// The type of Disk.
	DiskType *string `pulumi:"diskType"`
	// The count of instance.
	InstanceCount *int `pulumi:"instanceCount"`
	// The type of network. Valid values: "privatenet" and "pubnet".
	NetType *string `pulumi:"netType"`
	// The specification of private network SLB.
	PrivateSlbSpecification *string `pulumi:"privateSlbSpecification"`
	// The public network bandwidth. `0` means no access to the public network.
	PubNetworkFlow *string `pulumi:"pubNetworkFlow"`
	// The specification of public network SLB.
	PubSlbSpecification *string `pulumi:"pubSlbSpecification"`
	// The status of MSE Cluster.
	Status *string `pulumi:"status"`
	// The id of VSwitch.
	VswitchId *string `pulumi:"vswitchId"`
}

type ClusterState struct {
	// The whitelist.
	AclEntryLists pulumi.StringArrayInput
	// The alias of MSE Cluster.
	ClusterAliasName pulumi.StringPtrInput
	// The engine specification of MSE Cluster. Valid values:
	// `MSE_SC_1_2_200_c`：1C2G
	// `MSE_SC_2_4_200_c`：2C4G
	// `MSE_SC_4_8_200_c`：4C8G
	// `MSE_SC_8_16_200_c`：8C16G
	ClusterSpecification pulumi.StringPtrInput
	// The type of MSE Cluster.
	ClusterType pulumi.StringPtrInput
	// The version of MSE Cluster.
	ClusterVersion pulumi.StringPtrInput
	// The type of Disk.
	DiskType pulumi.StringPtrInput
	// The count of instance.
	InstanceCount pulumi.IntPtrInput
	// The type of network. Valid values: "privatenet" and "pubnet".
	NetType pulumi.StringPtrInput
	// The specification of private network SLB.
	PrivateSlbSpecification pulumi.StringPtrInput
	// The public network bandwidth. `0` means no access to the public network.
	PubNetworkFlow pulumi.StringPtrInput
	// The specification of public network SLB.
	PubSlbSpecification pulumi.StringPtrInput
	// The status of MSE Cluster.
	Status pulumi.StringPtrInput
	// The id of VSwitch.
	VswitchId pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The whitelist.
	AclEntryLists []string `pulumi:"aclEntryLists"`
	// The alias of MSE Cluster.
	ClusterAliasName *string `pulumi:"clusterAliasName"`
	// The engine specification of MSE Cluster. Valid values:
	// `MSE_SC_1_2_200_c`：1C2G
	// `MSE_SC_2_4_200_c`：2C4G
	// `MSE_SC_4_8_200_c`：4C8G
	// `MSE_SC_8_16_200_c`：8C16G
	ClusterSpecification string `pulumi:"clusterSpecification"`
	// The type of MSE Cluster.
	ClusterType string `pulumi:"clusterType"`
	// The version of MSE Cluster.
	ClusterVersion string `pulumi:"clusterVersion"`
	// The type of Disk.
	DiskType *string `pulumi:"diskType"`
	// The count of instance.
	InstanceCount int `pulumi:"instanceCount"`
	// The type of network. Valid values: "privatenet" and "pubnet".
	NetType string `pulumi:"netType"`
	// The specification of private network SLB.
	PrivateSlbSpecification *string `pulumi:"privateSlbSpecification"`
	// The public network bandwidth. `0` means no access to the public network.
	PubNetworkFlow *string `pulumi:"pubNetworkFlow"`
	// The specification of public network SLB.
	PubSlbSpecification *string `pulumi:"pubSlbSpecification"`
	// The id of VSwitch.
	VswitchId *string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The whitelist.
	AclEntryLists pulumi.StringArrayInput
	// The alias of MSE Cluster.
	ClusterAliasName pulumi.StringPtrInput
	// The engine specification of MSE Cluster. Valid values:
	// `MSE_SC_1_2_200_c`：1C2G
	// `MSE_SC_2_4_200_c`：2C4G
	// `MSE_SC_4_8_200_c`：4C8G
	// `MSE_SC_8_16_200_c`：8C16G
	ClusterSpecification pulumi.StringInput
	// The type of MSE Cluster.
	ClusterType pulumi.StringInput
	// The version of MSE Cluster.
	ClusterVersion pulumi.StringInput
	// The type of Disk.
	DiskType pulumi.StringPtrInput
	// The count of instance.
	InstanceCount pulumi.IntInput
	// The type of network. Valid values: "privatenet" and "pubnet".
	NetType pulumi.StringInput
	// The specification of private network SLB.
	PrivateSlbSpecification pulumi.StringPtrInput
	// The public network bandwidth. `0` means no access to the public network.
	PubNetworkFlow pulumi.StringPtrInput
	// The specification of public network SLB.
	PubSlbSpecification pulumi.StringPtrInput
	// The id of VSwitch.
	VswitchId pulumi.StringPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil))
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

func (i *Cluster) ToClusterPtrOutput() ClusterPtrOutput {
	return i.ToClusterPtrOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPtrOutput)
}

type ClusterPtrInput interface {
	pulumi.Input

	ToClusterPtrOutput() ClusterPtrOutput
	ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput
}

type clusterPtrType ClusterArgs

func (*clusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil))
}

func (i *clusterPtrType) ToClusterPtrOutput() ClusterPtrOutput {
	return i.ToClusterPtrOutputWithContext(context.Background())
}

func (i *clusterPtrType) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPtrOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//          ClusterArray{ ClusterArgs{...} }
type ClusterArrayInput interface {
	pulumi.Input

	ToClusterArrayOutput() ClusterArrayOutput
	ToClusterArrayOutputWithContext(context.Context) ClusterArrayOutput
}

type ClusterArray []ClusterInput

func (ClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (i ClusterArray) ToClusterArrayOutput() ClusterArrayOutput {
	return i.ToClusterArrayOutputWithContext(context.Background())
}

func (i ClusterArray) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterArrayOutput)
}

// ClusterMapInput is an input type that accepts ClusterMap and ClusterMapOutput values.
// You can construct a concrete instance of `ClusterMapInput` via:
//
//          ClusterMap{ "key": ClusterArgs{...} }
type ClusterMapInput interface {
	pulumi.Input

	ToClusterMapOutput() ClusterMapOutput
	ToClusterMapOutputWithContext(context.Context) ClusterMapOutput
}

type ClusterMap map[string]ClusterInput

func (ClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (i ClusterMap) ToClusterMapOutput() ClusterMapOutput {
	return i.ToClusterMapOutputWithContext(context.Background())
}

func (i ClusterMap) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMapOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil))
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterPtrOutput() ClusterPtrOutput {
	return o.ToClusterPtrOutputWithContext(context.Background())
}

func (o ClusterOutput) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Cluster) *Cluster {
		return &v
	}).(ClusterPtrOutput)
}

type ClusterPtrOutput struct{ *pulumi.OutputState }

func (ClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil))
}

func (o ClusterPtrOutput) ToClusterPtrOutput() ClusterPtrOutput {
	return o
}

func (o ClusterPtrOutput) ToClusterPtrOutputWithContext(ctx context.Context) ClusterPtrOutput {
	return o
}

func (o ClusterPtrOutput) Elem() ClusterOutput {
	return o.ApplyT(func(v *Cluster) Cluster {
		if v != nil {
			return *v
		}
		var ret Cluster
		return ret
	}).(ClusterOutput)
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Cluster)(nil))
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Cluster {
		return vs[0].([]Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Cluster)(nil))
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Cluster {
		return vs[0].(map[string]Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterPtrInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterArrayInput)(nil)).Elem(), ClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMapInput)(nil)).Elem(), ClusterMap{})
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterPtrOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}
