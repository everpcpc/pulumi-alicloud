// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This topic describes how to configure the health check feature for a Cloud Enterprise Network (CEN) instance.
// After you attach a Virtual Border Router (VBR) to the CEN instance and configure the health check feature, you can monitor the network conditions of the on-premises data center connected to the VBR.
//
// For information about CEN VBR HealthCheck and how to use it, see [Manage CEN VBR HealthCheck](https://www.alibabacloud.com/help/en/doc-detail/71141.htm).
//
// > **NOTE:** Available in 1.88.0+
type VbrHealthCheck struct {
	pulumi.CustomResourceState

	// The ID of the CEN instance.
	CenId pulumi.StringOutput `pulumi:"cenId"`
	// Specifies the interval at which the health check sends continuous detection packets. Default value: 2. Value range: 2 to 3.
	HealthCheckInterval pulumi.IntPtrOutput `pulumi:"healthCheckInterval"`
	// The source IP address of health checks.
	HealthCheckSourceIp pulumi.StringPtrOutput `pulumi:"healthCheckSourceIp"`
	// The destination IP address of health checks.
	HealthCheckTargetIp pulumi.StringOutput `pulumi:"healthCheckTargetIp"`
	// Specifies the number of probe messages sent by the health check. Default value: 8. Value range: 3 to 8.
	HealthyThreshold pulumi.IntPtrOutput `pulumi:"healthyThreshold"`
	// The ID of the VBR.
	VbrInstanceId pulumi.StringOutput `pulumi:"vbrInstanceId"`
	// The ID of the account to which the VBR belongs.
	VbrInstanceOwnerId pulumi.IntPtrOutput `pulumi:"vbrInstanceOwnerId"`
	// The ID of the region to which the VBR belongs.
	VbrInstanceRegionId pulumi.StringOutput `pulumi:"vbrInstanceRegionId"`
}

// NewVbrHealthCheck registers a new resource with the given unique name, arguments, and options.
func NewVbrHealthCheck(ctx *pulumi.Context,
	name string, args *VbrHealthCheckArgs, opts ...pulumi.ResourceOption) (*VbrHealthCheck, error) {
	if args == nil || args.CenId == nil {
		return nil, errors.New("missing required argument 'CenId'")
	}
	if args == nil || args.HealthCheckTargetIp == nil {
		return nil, errors.New("missing required argument 'HealthCheckTargetIp'")
	}
	if args == nil || args.VbrInstanceId == nil {
		return nil, errors.New("missing required argument 'VbrInstanceId'")
	}
	if args == nil || args.VbrInstanceRegionId == nil {
		return nil, errors.New("missing required argument 'VbrInstanceRegionId'")
	}
	if args == nil {
		args = &VbrHealthCheckArgs{}
	}
	var resource VbrHealthCheck
	err := ctx.RegisterResource("alicloud:cen/vbrHealthCheck:VbrHealthCheck", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVbrHealthCheck gets an existing VbrHealthCheck resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVbrHealthCheck(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VbrHealthCheckState, opts ...pulumi.ResourceOption) (*VbrHealthCheck, error) {
	var resource VbrHealthCheck
	err := ctx.ReadResource("alicloud:cen/vbrHealthCheck:VbrHealthCheck", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VbrHealthCheck resources.
type vbrHealthCheckState struct {
	// The ID of the CEN instance.
	CenId *string `pulumi:"cenId"`
	// Specifies the interval at which the health check sends continuous detection packets. Default value: 2. Value range: 2 to 3.
	HealthCheckInterval *int `pulumi:"healthCheckInterval"`
	// The source IP address of health checks.
	HealthCheckSourceIp *string `pulumi:"healthCheckSourceIp"`
	// The destination IP address of health checks.
	HealthCheckTargetIp *string `pulumi:"healthCheckTargetIp"`
	// Specifies the number of probe messages sent by the health check. Default value: 8. Value range: 3 to 8.
	HealthyThreshold *int `pulumi:"healthyThreshold"`
	// The ID of the VBR.
	VbrInstanceId *string `pulumi:"vbrInstanceId"`
	// The ID of the account to which the VBR belongs.
	VbrInstanceOwnerId *int `pulumi:"vbrInstanceOwnerId"`
	// The ID of the region to which the VBR belongs.
	VbrInstanceRegionId *string `pulumi:"vbrInstanceRegionId"`
}

type VbrHealthCheckState struct {
	// The ID of the CEN instance.
	CenId pulumi.StringPtrInput
	// Specifies the interval at which the health check sends continuous detection packets. Default value: 2. Value range: 2 to 3.
	HealthCheckInterval pulumi.IntPtrInput
	// The source IP address of health checks.
	HealthCheckSourceIp pulumi.StringPtrInput
	// The destination IP address of health checks.
	HealthCheckTargetIp pulumi.StringPtrInput
	// Specifies the number of probe messages sent by the health check. Default value: 8. Value range: 3 to 8.
	HealthyThreshold pulumi.IntPtrInput
	// The ID of the VBR.
	VbrInstanceId pulumi.StringPtrInput
	// The ID of the account to which the VBR belongs.
	VbrInstanceOwnerId pulumi.IntPtrInput
	// The ID of the region to which the VBR belongs.
	VbrInstanceRegionId pulumi.StringPtrInput
}

func (VbrHealthCheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*vbrHealthCheckState)(nil)).Elem()
}

type vbrHealthCheckArgs struct {
	// The ID of the CEN instance.
	CenId string `pulumi:"cenId"`
	// Specifies the interval at which the health check sends continuous detection packets. Default value: 2. Value range: 2 to 3.
	HealthCheckInterval *int `pulumi:"healthCheckInterval"`
	// The source IP address of health checks.
	HealthCheckSourceIp *string `pulumi:"healthCheckSourceIp"`
	// The destination IP address of health checks.
	HealthCheckTargetIp string `pulumi:"healthCheckTargetIp"`
	// Specifies the number of probe messages sent by the health check. Default value: 8. Value range: 3 to 8.
	HealthyThreshold *int `pulumi:"healthyThreshold"`
	// The ID of the VBR.
	VbrInstanceId string `pulumi:"vbrInstanceId"`
	// The ID of the account to which the VBR belongs.
	VbrInstanceOwnerId *int `pulumi:"vbrInstanceOwnerId"`
	// The ID of the region to which the VBR belongs.
	VbrInstanceRegionId string `pulumi:"vbrInstanceRegionId"`
}

// The set of arguments for constructing a VbrHealthCheck resource.
type VbrHealthCheckArgs struct {
	// The ID of the CEN instance.
	CenId pulumi.StringInput
	// Specifies the interval at which the health check sends continuous detection packets. Default value: 2. Value range: 2 to 3.
	HealthCheckInterval pulumi.IntPtrInput
	// The source IP address of health checks.
	HealthCheckSourceIp pulumi.StringPtrInput
	// The destination IP address of health checks.
	HealthCheckTargetIp pulumi.StringInput
	// Specifies the number of probe messages sent by the health check. Default value: 8. Value range: 3 to 8.
	HealthyThreshold pulumi.IntPtrInput
	// The ID of the VBR.
	VbrInstanceId pulumi.StringInput
	// The ID of the account to which the VBR belongs.
	VbrInstanceOwnerId pulumi.IntPtrInput
	// The ID of the region to which the VBR belongs.
	VbrInstanceRegionId pulumi.StringInput
}

func (VbrHealthCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vbrHealthCheckArgs)(nil)).Elem()
}
