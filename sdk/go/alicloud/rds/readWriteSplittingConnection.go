// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an RDS read write splitting connection resource to allocate an Intranet connection string for RDS instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/rds"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		creation := "Rds"
// 		if param := cfg.Get("creation"); param != "" {
// 			creation = param
// 		}
// 		name := "dbInstancevpc"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		opt0 := creation
// 		defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
// 			AvailableResourceCreation: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
// 			VpcId:            defaultNetwork.ID(),
// 			CidrBlock:        pulumi.String("172.16.0.0/24"),
// 			AvailabilityZone: pulumi.String(defaultZones.Zones[0].Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultInstance, err := rds.NewInstance(ctx, "defaultInstance", &rds.InstanceArgs{
// 			Engine:             pulumi.String("MySQL"),
// 			EngineVersion:      pulumi.String("5.6"),
// 			InstanceType:       pulumi.String("rds.mysql.t1.small"),
// 			InstanceStorage:    pulumi.Int(20),
// 			InstanceChargeType: pulumi.String("Postpaid"),
// 			InstanceName:       pulumi.String(name),
// 			VswitchId:          defaultSwitch.ID(),
// 			SecurityIps: pulumi.StringArray{
// 				pulumi.String("10.168.1.12"),
// 				pulumi.String("100.69.7.112"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultReadOnlyInstance, err := rds.NewReadOnlyInstance(ctx, "defaultReadOnlyInstance", &rds.ReadOnlyInstanceArgs{
// 			MasterDbInstanceId: defaultInstance.ID(),
// 			ZoneId:             defaultInstance.ZoneId,
// 			EngineVersion:      defaultInstance.EngineVersion,
// 			InstanceType:       defaultInstance.InstanceType,
// 			InstanceStorage:    pulumi.Int(30),
// 			InstanceName:       pulumi.String(fmt.Sprintf("%v%v", name, "ro")),
// 			VswitchId:          defaultSwitch.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = rds.NewReadWriteSplittingConnection(ctx, "defaultReadWriteSplittingConnection", &rds.ReadWriteSplittingConnectionArgs{
// 			InstanceId:       defaultInstance.ID(),
// 			ConnectionPrefix: pulumi.String("t-con-123"),
// 			DistributionType: pulumi.String("Standard"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			defaultReadOnlyInstance,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// > **NOTE:** Resource `rds.ReadWriteSplittingConnection` should be created after `rds.ReadOnlyInstance`, so the `dependsOn` statement is necessary.
//
// ## Import
//
// RDS read write splitting connection can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:rds/readWriteSplittingConnection:ReadWriteSplittingConnection example abc12345678
// ```
type ReadWriteSplittingConnection struct {
	pulumi.CustomResourceState

	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
	ConnectionPrefix pulumi.StringPtrOutput `pulumi:"connectionPrefix"`
	// Connection instance string.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution.
	DistributionType pulumi.StringOutput `pulumi:"distributionType"`
	// The Id of instance that can run database.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.
	MaxDelayTime pulumi.IntOutput `pulumi:"maxDelayTime"`
	// Intranet connection port. Valid value: [3001-3999]. Default to 3306.
	Port pulumi.IntOutput `pulumi:"port"`
	// Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distributionType is set to Custom.
	Weight pulumi.MapOutput `pulumi:"weight"`
}

// NewReadWriteSplittingConnection registers a new resource with the given unique name, arguments, and options.
func NewReadWriteSplittingConnection(ctx *pulumi.Context,
	name string, args *ReadWriteSplittingConnectionArgs, opts ...pulumi.ResourceOption) (*ReadWriteSplittingConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DistributionType == nil {
		return nil, errors.New("invalid value for required argument 'DistributionType'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource ReadWriteSplittingConnection
	err := ctx.RegisterResource("alicloud:rds/readWriteSplittingConnection:ReadWriteSplittingConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReadWriteSplittingConnection gets an existing ReadWriteSplittingConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReadWriteSplittingConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReadWriteSplittingConnectionState, opts ...pulumi.ResourceOption) (*ReadWriteSplittingConnection, error) {
	var resource ReadWriteSplittingConnection
	err := ctx.ReadResource("alicloud:rds/readWriteSplittingConnection:ReadWriteSplittingConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReadWriteSplittingConnection resources.
type readWriteSplittingConnectionState struct {
	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
	ConnectionPrefix *string `pulumi:"connectionPrefix"`
	// Connection instance string.
	ConnectionString *string `pulumi:"connectionString"`
	// Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution.
	DistributionType *string `pulumi:"distributionType"`
	// The Id of instance that can run database.
	InstanceId *string `pulumi:"instanceId"`
	// Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.
	MaxDelayTime *int `pulumi:"maxDelayTime"`
	// Intranet connection port. Valid value: [3001-3999]. Default to 3306.
	Port *int `pulumi:"port"`
	// Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distributionType is set to Custom.
	Weight map[string]interface{} `pulumi:"weight"`
}

type ReadWriteSplittingConnectionState struct {
	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
	ConnectionPrefix pulumi.StringPtrInput
	// Connection instance string.
	ConnectionString pulumi.StringPtrInput
	// Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution.
	DistributionType pulumi.StringPtrInput
	// The Id of instance that can run database.
	InstanceId pulumi.StringPtrInput
	// Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.
	MaxDelayTime pulumi.IntPtrInput
	// Intranet connection port. Valid value: [3001-3999]. Default to 3306.
	Port pulumi.IntPtrInput
	// Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distributionType is set to Custom.
	Weight pulumi.MapInput
}

func (ReadWriteSplittingConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*readWriteSplittingConnectionState)(nil)).Elem()
}

type readWriteSplittingConnectionArgs struct {
	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
	ConnectionPrefix *string `pulumi:"connectionPrefix"`
	// Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution.
	DistributionType string `pulumi:"distributionType"`
	// The Id of instance that can run database.
	InstanceId string `pulumi:"instanceId"`
	// Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.
	MaxDelayTime *int `pulumi:"maxDelayTime"`
	// Intranet connection port. Valid value: [3001-3999]. Default to 3306.
	Port *int `pulumi:"port"`
	// Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distributionType is set to Custom.
	Weight map[string]interface{} `pulumi:"weight"`
}

// The set of arguments for constructing a ReadWriteSplittingConnection resource.
type ReadWriteSplittingConnectionArgs struct {
	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
	ConnectionPrefix pulumi.StringPtrInput
	// Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution.
	DistributionType pulumi.StringInput
	// The Id of instance that can run database.
	InstanceId pulumi.StringInput
	// Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.
	MaxDelayTime pulumi.IntPtrInput
	// Intranet connection port. Valid value: [3001-3999]. Default to 3306.
	Port pulumi.IntPtrInput
	// Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distributionType is set to Custom.
	Weight pulumi.MapInput
}

func (ReadWriteSplittingConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*readWriteSplittingConnectionArgs)(nil)).Elem()
}

type ReadWriteSplittingConnectionInput interface {
	pulumi.Input

	ToReadWriteSplittingConnectionOutput() ReadWriteSplittingConnectionOutput
	ToReadWriteSplittingConnectionOutputWithContext(ctx context.Context) ReadWriteSplittingConnectionOutput
}

func (*ReadWriteSplittingConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*ReadWriteSplittingConnection)(nil))
}

func (i *ReadWriteSplittingConnection) ToReadWriteSplittingConnectionOutput() ReadWriteSplittingConnectionOutput {
	return i.ToReadWriteSplittingConnectionOutputWithContext(context.Background())
}

func (i *ReadWriteSplittingConnection) ToReadWriteSplittingConnectionOutputWithContext(ctx context.Context) ReadWriteSplittingConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadWriteSplittingConnectionOutput)
}

func (i *ReadWriteSplittingConnection) ToReadWriteSplittingConnectionPtrOutput() ReadWriteSplittingConnectionPtrOutput {
	return i.ToReadWriteSplittingConnectionPtrOutputWithContext(context.Background())
}

func (i *ReadWriteSplittingConnection) ToReadWriteSplittingConnectionPtrOutputWithContext(ctx context.Context) ReadWriteSplittingConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadWriteSplittingConnectionPtrOutput)
}

type ReadWriteSplittingConnectionPtrInput interface {
	pulumi.Input

	ToReadWriteSplittingConnectionPtrOutput() ReadWriteSplittingConnectionPtrOutput
	ToReadWriteSplittingConnectionPtrOutputWithContext(ctx context.Context) ReadWriteSplittingConnectionPtrOutput
}

type readWriteSplittingConnectionPtrType ReadWriteSplittingConnectionArgs

func (*readWriteSplittingConnectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadWriteSplittingConnection)(nil))
}

func (i *readWriteSplittingConnectionPtrType) ToReadWriteSplittingConnectionPtrOutput() ReadWriteSplittingConnectionPtrOutput {
	return i.ToReadWriteSplittingConnectionPtrOutputWithContext(context.Background())
}

func (i *readWriteSplittingConnectionPtrType) ToReadWriteSplittingConnectionPtrOutputWithContext(ctx context.Context) ReadWriteSplittingConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadWriteSplittingConnectionPtrOutput)
}

// ReadWriteSplittingConnectionArrayInput is an input type that accepts ReadWriteSplittingConnectionArray and ReadWriteSplittingConnectionArrayOutput values.
// You can construct a concrete instance of `ReadWriteSplittingConnectionArrayInput` via:
//
//          ReadWriteSplittingConnectionArray{ ReadWriteSplittingConnectionArgs{...} }
type ReadWriteSplittingConnectionArrayInput interface {
	pulumi.Input

	ToReadWriteSplittingConnectionArrayOutput() ReadWriteSplittingConnectionArrayOutput
	ToReadWriteSplittingConnectionArrayOutputWithContext(context.Context) ReadWriteSplittingConnectionArrayOutput
}

type ReadWriteSplittingConnectionArray []ReadWriteSplittingConnectionInput

func (ReadWriteSplittingConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ReadWriteSplittingConnection)(nil))
}

func (i ReadWriteSplittingConnectionArray) ToReadWriteSplittingConnectionArrayOutput() ReadWriteSplittingConnectionArrayOutput {
	return i.ToReadWriteSplittingConnectionArrayOutputWithContext(context.Background())
}

func (i ReadWriteSplittingConnectionArray) ToReadWriteSplittingConnectionArrayOutputWithContext(ctx context.Context) ReadWriteSplittingConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadWriteSplittingConnectionArrayOutput)
}

// ReadWriteSplittingConnectionMapInput is an input type that accepts ReadWriteSplittingConnectionMap and ReadWriteSplittingConnectionMapOutput values.
// You can construct a concrete instance of `ReadWriteSplittingConnectionMapInput` via:
//
//          ReadWriteSplittingConnectionMap{ "key": ReadWriteSplittingConnectionArgs{...} }
type ReadWriteSplittingConnectionMapInput interface {
	pulumi.Input

	ToReadWriteSplittingConnectionMapOutput() ReadWriteSplittingConnectionMapOutput
	ToReadWriteSplittingConnectionMapOutputWithContext(context.Context) ReadWriteSplittingConnectionMapOutput
}

type ReadWriteSplittingConnectionMap map[string]ReadWriteSplittingConnectionInput

func (ReadWriteSplittingConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ReadWriteSplittingConnection)(nil))
}

func (i ReadWriteSplittingConnectionMap) ToReadWriteSplittingConnectionMapOutput() ReadWriteSplittingConnectionMapOutput {
	return i.ToReadWriteSplittingConnectionMapOutputWithContext(context.Background())
}

func (i ReadWriteSplittingConnectionMap) ToReadWriteSplittingConnectionMapOutputWithContext(ctx context.Context) ReadWriteSplittingConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadWriteSplittingConnectionMapOutput)
}

type ReadWriteSplittingConnectionOutput struct {
	*pulumi.OutputState
}

func (ReadWriteSplittingConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReadWriteSplittingConnection)(nil))
}

func (o ReadWriteSplittingConnectionOutput) ToReadWriteSplittingConnectionOutput() ReadWriteSplittingConnectionOutput {
	return o
}

func (o ReadWriteSplittingConnectionOutput) ToReadWriteSplittingConnectionOutputWithContext(ctx context.Context) ReadWriteSplittingConnectionOutput {
	return o
}

func (o ReadWriteSplittingConnectionOutput) ToReadWriteSplittingConnectionPtrOutput() ReadWriteSplittingConnectionPtrOutput {
	return o.ToReadWriteSplittingConnectionPtrOutputWithContext(context.Background())
}

func (o ReadWriteSplittingConnectionOutput) ToReadWriteSplittingConnectionPtrOutputWithContext(ctx context.Context) ReadWriteSplittingConnectionPtrOutput {
	return o.ApplyT(func(v ReadWriteSplittingConnection) *ReadWriteSplittingConnection {
		return &v
	}).(ReadWriteSplittingConnectionPtrOutput)
}

type ReadWriteSplittingConnectionPtrOutput struct {
	*pulumi.OutputState
}

func (ReadWriteSplittingConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadWriteSplittingConnection)(nil))
}

func (o ReadWriteSplittingConnectionPtrOutput) ToReadWriteSplittingConnectionPtrOutput() ReadWriteSplittingConnectionPtrOutput {
	return o
}

func (o ReadWriteSplittingConnectionPtrOutput) ToReadWriteSplittingConnectionPtrOutputWithContext(ctx context.Context) ReadWriteSplittingConnectionPtrOutput {
	return o
}

type ReadWriteSplittingConnectionArrayOutput struct{ *pulumi.OutputState }

func (ReadWriteSplittingConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReadWriteSplittingConnection)(nil))
}

func (o ReadWriteSplittingConnectionArrayOutput) ToReadWriteSplittingConnectionArrayOutput() ReadWriteSplittingConnectionArrayOutput {
	return o
}

func (o ReadWriteSplittingConnectionArrayOutput) ToReadWriteSplittingConnectionArrayOutputWithContext(ctx context.Context) ReadWriteSplittingConnectionArrayOutput {
	return o
}

func (o ReadWriteSplittingConnectionArrayOutput) Index(i pulumi.IntInput) ReadWriteSplittingConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReadWriteSplittingConnection {
		return vs[0].([]ReadWriteSplittingConnection)[vs[1].(int)]
	}).(ReadWriteSplittingConnectionOutput)
}

type ReadWriteSplittingConnectionMapOutput struct{ *pulumi.OutputState }

func (ReadWriteSplittingConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReadWriteSplittingConnection)(nil))
}

func (o ReadWriteSplittingConnectionMapOutput) ToReadWriteSplittingConnectionMapOutput() ReadWriteSplittingConnectionMapOutput {
	return o
}

func (o ReadWriteSplittingConnectionMapOutput) ToReadWriteSplittingConnectionMapOutputWithContext(ctx context.Context) ReadWriteSplittingConnectionMapOutput {
	return o
}

func (o ReadWriteSplittingConnectionMapOutput) MapIndex(k pulumi.StringInput) ReadWriteSplittingConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReadWriteSplittingConnection {
		return vs[0].(map[string]ReadWriteSplittingConnection)[vs[1].(string)]
	}).(ReadWriteSplittingConnectionOutput)
}

func init() {
	pulumi.RegisterOutputType(ReadWriteSplittingConnectionOutput{})
	pulumi.RegisterOutputType(ReadWriteSplittingConnectionPtrOutput{})
	pulumi.RegisterOutputType(ReadWriteSplittingConnectionArrayOutput{})
	pulumi.RegisterOutputType(ReadWriteSplittingConnectionMapOutput{})
}
