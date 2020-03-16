// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package apigateway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type VpcAccess struct {
	pulumi.CustomResourceState

	// ID of the instance in VPC (ECS/Server Load Balance).
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The name of the vpc authorization. 
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the port corresponding to the instance.
	Port pulumi.IntOutput `pulumi:"port"`
	// The vpc id of the vpc authorization. 
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVpcAccess registers a new resource with the given unique name, arguments, and options.
func NewVpcAccess(ctx *pulumi.Context,
	name string, args *VpcAccessArgs, opts ...pulumi.ResourceOption) (*VpcAccess, error) {
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil || args.Port == nil {
		return nil, errors.New("missing required argument 'Port'")
	}
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	if args == nil {
		args = &VpcAccessArgs{}
	}
	var resource VpcAccess
	err := ctx.RegisterResource("alicloud:apigateway/vpcAccess:VpcAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcAccess gets an existing VpcAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcAccessState, opts ...pulumi.ResourceOption) (*VpcAccess, error) {
	var resource VpcAccess
	err := ctx.ReadResource("alicloud:apigateway/vpcAccess:VpcAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcAccess resources.
type vpcAccessState struct {
	// ID of the instance in VPC (ECS/Server Load Balance).
	InstanceId *string `pulumi:"instanceId"`
	// The name of the vpc authorization. 
	Name *string `pulumi:"name"`
	// ID of the port corresponding to the instance.
	Port *int `pulumi:"port"`
	// The vpc id of the vpc authorization. 
	VpcId *string `pulumi:"vpcId"`
}

type VpcAccessState struct {
	// ID of the instance in VPC (ECS/Server Load Balance).
	InstanceId pulumi.StringPtrInput
	// The name of the vpc authorization. 
	Name pulumi.StringPtrInput
	// ID of the port corresponding to the instance.
	Port pulumi.IntPtrInput
	// The vpc id of the vpc authorization. 
	VpcId pulumi.StringPtrInput
}

func (VpcAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcAccessState)(nil)).Elem()
}

type vpcAccessArgs struct {
	// ID of the instance in VPC (ECS/Server Load Balance).
	InstanceId string `pulumi:"instanceId"`
	// The name of the vpc authorization. 
	Name *string `pulumi:"name"`
	// ID of the port corresponding to the instance.
	Port int `pulumi:"port"`
	// The vpc id of the vpc authorization. 
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpcAccess resource.
type VpcAccessArgs struct {
	// ID of the instance in VPC (ECS/Server Load Balance).
	InstanceId pulumi.StringInput
	// The name of the vpc authorization. 
	Name pulumi.StringPtrInput
	// ID of the port corresponding to the instance.
	Port pulumi.IntInput
	// The vpc id of the vpc authorization. 
	VpcId pulumi.StringInput
}

func (VpcAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcAccessArgs)(nil)).Elem()
}

