// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A master slave server group contains two ECS instances. The master slave server group can help you to define multiple listening dimension.
//
// > **NOTE:** One ECS instance can be added into multiple master slave server groups.
//
// > **NOTE:** One master slave server group can only add two ECS instances, which are master server and slave server.
//
// > **NOTE:** One master slave server group can be attached with tcp/udp listeners in one load balancer.
//
// > **NOTE:** One Classic and Internet load balancer, its master slave server group can add Classic and VPC ECS instances.
//
// > **NOTE:** One Classic and Intranet load balancer, its master slave server group can only add Classic ECS instances.
//
// > **NOTE:** One VPC load balancer, its master slave server group can only add the same VPC ECS instances.
//
// > **NOTE:** Available in 1.54.0+
//
// ## Block servers
//
// The servers mapping supports the following:
//
// * `serverIds` - (Required) A list backend server ID (ECS instance ID).
// * `port` - (Required) The port used by the backend server. Valid value range: [1-65535].
// * `weight` - (Optional) Weight of the backend server. Valid value range: [0-100]. Default to 100.
// * `type` - (Optional, Available in 1.51.0+) Type of the backend server. Valid value ecs, eni. Default to eni.
// * `serverType` - (Optional) The server type of the backend server. Valid value Master, Slave.
// * `isBackup` - (Removed from v1.63.0) Determine if the server is executing. Valid value 0, 1.
//
// ## Import
//
// Load balancer master slave server group can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup example abc123456
// ```
type MasterSlaveServerGroup struct {
	pulumi.CustomResourceState

	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrOutput `pulumi:"deleteProtectionValidation"`
	// The Load Balancer ID which is used to launch a new master slave server group.
	LoadBalancerId pulumi.StringOutput `pulumi:"loadBalancerId"`
	// Name of the master slave server group.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
	Servers MasterSlaveServerGroupServerArrayOutput `pulumi:"servers"`
}

// NewMasterSlaveServerGroup registers a new resource with the given unique name, arguments, and options.
func NewMasterSlaveServerGroup(ctx *pulumi.Context,
	name string, args *MasterSlaveServerGroupArgs, opts ...pulumi.ResourceOption) (*MasterSlaveServerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadBalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerId'")
	}
	var resource MasterSlaveServerGroup
	err := ctx.RegisterResource("alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMasterSlaveServerGroup gets an existing MasterSlaveServerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMasterSlaveServerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MasterSlaveServerGroupState, opts ...pulumi.ResourceOption) (*MasterSlaveServerGroup, error) {
	var resource MasterSlaveServerGroup
	err := ctx.ReadResource("alicloud:slb/masterSlaveServerGroup:MasterSlaveServerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MasterSlaveServerGroup resources.
type masterSlaveServerGroupState struct {
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation *bool `pulumi:"deleteProtectionValidation"`
	// The Load Balancer ID which is used to launch a new master slave server group.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// Name of the master slave server group.
	Name *string `pulumi:"name"`
	// A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
	Servers []MasterSlaveServerGroupServer `pulumi:"servers"`
}

type MasterSlaveServerGroupState struct {
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrInput
	// The Load Balancer ID which is used to launch a new master slave server group.
	LoadBalancerId pulumi.StringPtrInput
	// Name of the master slave server group.
	Name pulumi.StringPtrInput
	// A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
	Servers MasterSlaveServerGroupServerArrayInput
}

func (MasterSlaveServerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*masterSlaveServerGroupState)(nil)).Elem()
}

type masterSlaveServerGroupArgs struct {
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation *bool `pulumi:"deleteProtectionValidation"`
	// The Load Balancer ID which is used to launch a new master slave server group.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// Name of the master slave server group.
	Name *string `pulumi:"name"`
	// A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
	Servers []MasterSlaveServerGroupServer `pulumi:"servers"`
}

// The set of arguments for constructing a MasterSlaveServerGroup resource.
type MasterSlaveServerGroupArgs struct {
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrInput
	// The Load Balancer ID which is used to launch a new master slave server group.
	LoadBalancerId pulumi.StringInput
	// Name of the master slave server group.
	Name pulumi.StringPtrInput
	// A list of ECS instances to be added. Only two ECS instances can be supported in one resource. It contains six sub-fields as `Block server` follows.
	Servers MasterSlaveServerGroupServerArrayInput
}

func (MasterSlaveServerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*masterSlaveServerGroupArgs)(nil)).Elem()
}

type MasterSlaveServerGroupInput interface {
	pulumi.Input

	ToMasterSlaveServerGroupOutput() MasterSlaveServerGroupOutput
	ToMasterSlaveServerGroupOutputWithContext(ctx context.Context) MasterSlaveServerGroupOutput
}

func (MasterSlaveServerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*MasterSlaveServerGroup)(nil)).Elem()
}

func (i MasterSlaveServerGroup) ToMasterSlaveServerGroupOutput() MasterSlaveServerGroupOutput {
	return i.ToMasterSlaveServerGroupOutputWithContext(context.Background())
}

func (i MasterSlaveServerGroup) ToMasterSlaveServerGroupOutputWithContext(ctx context.Context) MasterSlaveServerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MasterSlaveServerGroupOutput)
}

type MasterSlaveServerGroupOutput struct {
	*pulumi.OutputState
}

func (MasterSlaveServerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MasterSlaveServerGroupOutput)(nil)).Elem()
}

func (o MasterSlaveServerGroupOutput) ToMasterSlaveServerGroupOutput() MasterSlaveServerGroupOutput {
	return o
}

func (o MasterSlaveServerGroupOutput) ToMasterSlaveServerGroupOutputWithContext(ctx context.Context) MasterSlaveServerGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MasterSlaveServerGroupOutput{})
}
