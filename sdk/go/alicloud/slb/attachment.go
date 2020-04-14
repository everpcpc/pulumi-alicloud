// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Attachment struct {
	pulumi.CustomResourceState

	// The backend servers of the load balancer.
	BackendServers pulumi.StringOutput `pulumi:"backendServers"`
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrOutput `pulumi:"deleteProtectionValidation"`
	// A list of instance ids to added backend server in the SLB.
	InstanceIds pulumi.StringArrayOutput `pulumi:"instanceIds"`
	// ID of the load balancer.
	LoadBalancerId pulumi.StringOutput `pulumi:"loadBalancerId"`
	// Type of the instances. Valid value ecs, eni. Default to ecs.
	ServerType pulumi.StringPtrOutput `pulumi:"serverType"`
	// Weight of the instances. Valid value range: [0-100]. Default to 100.
	Weight pulumi.IntPtrOutput `pulumi:"weight"`
}

// NewAttachment registers a new resource with the given unique name, arguments, and options.
func NewAttachment(ctx *pulumi.Context,
	name string, args *AttachmentArgs, opts ...pulumi.ResourceOption) (*Attachment, error) {
	if args == nil || args.InstanceIds == nil {
		return nil, errors.New("missing required argument 'InstanceIds'")
	}
	if args == nil || args.LoadBalancerId == nil {
		return nil, errors.New("missing required argument 'LoadBalancerId'")
	}
	if args == nil {
		args = &AttachmentArgs{}
	}
	var resource Attachment
	err := ctx.RegisterResource("alicloud:slb/attachment:Attachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachment gets an existing Attachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachmentState, opts ...pulumi.ResourceOption) (*Attachment, error) {
	var resource Attachment
	err := ctx.ReadResource("alicloud:slb/attachment:Attachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Attachment resources.
type attachmentState struct {
	// The backend servers of the load balancer.
	BackendServers *string `pulumi:"backendServers"`
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation *bool `pulumi:"deleteProtectionValidation"`
	// A list of instance ids to added backend server in the SLB.
	InstanceIds []string `pulumi:"instanceIds"`
	// ID of the load balancer.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// Type of the instances. Valid value ecs, eni. Default to ecs.
	ServerType *string `pulumi:"serverType"`
	// Weight of the instances. Valid value range: [0-100]. Default to 100.
	Weight *int `pulumi:"weight"`
}

type AttachmentState struct {
	// The backend servers of the load balancer.
	BackendServers pulumi.StringPtrInput
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrInput
	// A list of instance ids to added backend server in the SLB.
	InstanceIds pulumi.StringArrayInput
	// ID of the load balancer.
	LoadBalancerId pulumi.StringPtrInput
	// Type of the instances. Valid value ecs, eni. Default to ecs.
	ServerType pulumi.StringPtrInput
	// Weight of the instances. Valid value range: [0-100]. Default to 100.
	Weight pulumi.IntPtrInput
}

func (AttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachmentState)(nil)).Elem()
}

type attachmentArgs struct {
	// The backend servers of the load balancer.
	BackendServers *string `pulumi:"backendServers"`
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation *bool `pulumi:"deleteProtectionValidation"`
	// A list of instance ids to added backend server in the SLB.
	InstanceIds []string `pulumi:"instanceIds"`
	// ID of the load balancer.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// Type of the instances. Valid value ecs, eni. Default to ecs.
	ServerType *string `pulumi:"serverType"`
	// Weight of the instances. Valid value range: [0-100]. Default to 100.
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a Attachment resource.
type AttachmentArgs struct {
	// The backend servers of the load balancer.
	BackendServers pulumi.StringPtrInput
	// Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
	DeleteProtectionValidation pulumi.BoolPtrInput
	// A list of instance ids to added backend server in the SLB.
	InstanceIds pulumi.StringArrayInput
	// ID of the load balancer.
	LoadBalancerId pulumi.StringInput
	// Type of the instances. Valid value ecs, eni. Default to ecs.
	ServerType pulumi.StringPtrInput
	// Weight of the instances. Valid value range: [0-100]. Default to 100.
	Weight pulumi.IntPtrInput
}

func (AttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachmentArgs)(nil)).Elem()
}
