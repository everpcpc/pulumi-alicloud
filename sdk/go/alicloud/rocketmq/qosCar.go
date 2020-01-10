// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rocketmq

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Sag qos car resource. 
// You need to create a QoS car to set priorities, rate limits, and quintuple rules for different messages.
// 
// For information about Sag Qos Car and how to use it, see [What is Qos Car](https://www.alibabacloud.com/help/doc-detail/140065.htm).
// 
// > **NOTE:** Available in 1.60.0+
// 
// > **NOTE:** Only the following regions support. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/sag_qos_car.html.markdown.
type QosCar struct {
	s *pulumi.ResourceState
}

// NewQosCar registers a new resource with the given unique name, arguments, and options.
func NewQosCar(ctx *pulumi.Context,
	name string, args *QosCarArgs, opts ...pulumi.ResourceOpt) (*QosCar, error) {
	if args == nil || args.LimitType == nil {
		return nil, errors.New("missing required argument 'LimitType'")
	}
	if args == nil || args.Priority == nil {
		return nil, errors.New("missing required argument 'Priority'")
	}
	if args == nil || args.QosId == nil {
		return nil, errors.New("missing required argument 'QosId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["limitType"] = nil
		inputs["maxBandwidthAbs"] = nil
		inputs["maxBandwidthPercent"] = nil
		inputs["minBandwidthAbs"] = nil
		inputs["minBandwidthPercent"] = nil
		inputs["name"] = nil
		inputs["percentSourceType"] = nil
		inputs["priority"] = nil
		inputs["qosId"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["limitType"] = args.LimitType
		inputs["maxBandwidthAbs"] = args.MaxBandwidthAbs
		inputs["maxBandwidthPercent"] = args.MaxBandwidthPercent
		inputs["minBandwidthAbs"] = args.MinBandwidthAbs
		inputs["minBandwidthPercent"] = args.MinBandwidthPercent
		inputs["name"] = args.Name
		inputs["percentSourceType"] = args.PercentSourceType
		inputs["priority"] = args.Priority
		inputs["qosId"] = args.QosId
	}
	s, err := ctx.RegisterResource("alicloud:rocketmq/qosCar:QosCar", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &QosCar{s: s}, nil
}

// GetQosCar gets an existing QosCar resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQosCar(ctx *pulumi.Context,
	name string, id pulumi.ID, state *QosCarState, opts ...pulumi.ResourceOpt) (*QosCar, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["limitType"] = state.LimitType
		inputs["maxBandwidthAbs"] = state.MaxBandwidthAbs
		inputs["maxBandwidthPercent"] = state.MaxBandwidthPercent
		inputs["minBandwidthAbs"] = state.MinBandwidthAbs
		inputs["minBandwidthPercent"] = state.MinBandwidthPercent
		inputs["name"] = state.Name
		inputs["percentSourceType"] = state.PercentSourceType
		inputs["priority"] = state.Priority
		inputs["qosId"] = state.QosId
	}
	s, err := ctx.ReadResource("alicloud:rocketmq/qosCar:QosCar", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &QosCar{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *QosCar) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *QosCar) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The description of the QoS speed limiting rule.
func (r *QosCar) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The speed limiting method. Valid values: Absolute, Percent.
func (r *QosCar) LimitType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["limitType"])
}

// The maximum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType is Absolute.
func (r *QosCar) MaxBandwidthAbs() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["maxBandwidthAbs"])
}

// The maximum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated Smart Access Gateway (SAG) instance.This parameter is required when the value of the LimitType parameter is Percent.
func (r *QosCar) MaxBandwidthPercent() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["maxBandwidthPercent"])
}

// The minimum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType parameter is Absolute.
func (r *QosCar) MinBandwidthAbs() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["minBandwidthAbs"])
}

// The minimum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated SAG instance.This parameter is required when the value of the LimitType parameter is Percent.
func (r *QosCar) MinBandwidthPercent() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["minBandwidthPercent"])
}

// The name of the QoS speed limiting rule..
func (r *QosCar) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The bandwidth type when the speed is limited based on percentage. Valid values: CcnBandwidth, InternetUpBandwidth.The default value is InternetUpBandwidth.
func (r *QosCar) PercentSourceType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["percentSourceType"])
}

// The priority of the specified stream.
func (r *QosCar) Priority() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["priority"])
}

// The instance ID of the QoS.
func (r *QosCar) QosId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["qosId"])
}

// Input properties used for looking up and filtering QosCar resources.
type QosCarState struct {
	// The description of the QoS speed limiting rule.
	Description interface{}
	// The speed limiting method. Valid values: Absolute, Percent.
	LimitType interface{}
	// The maximum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType is Absolute.
	MaxBandwidthAbs interface{}
	// The maximum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated Smart Access Gateway (SAG) instance.This parameter is required when the value of the LimitType parameter is Percent.
	MaxBandwidthPercent interface{}
	// The minimum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType parameter is Absolute.
	MinBandwidthAbs interface{}
	// The minimum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated SAG instance.This parameter is required when the value of the LimitType parameter is Percent.
	MinBandwidthPercent interface{}
	// The name of the QoS speed limiting rule..
	Name interface{}
	// The bandwidth type when the speed is limited based on percentage. Valid values: CcnBandwidth, InternetUpBandwidth.The default value is InternetUpBandwidth.
	PercentSourceType interface{}
	// The priority of the specified stream.
	Priority interface{}
	// The instance ID of the QoS.
	QosId interface{}
}

// The set of arguments for constructing a QosCar resource.
type QosCarArgs struct {
	// The description of the QoS speed limiting rule.
	Description interface{}
	// The speed limiting method. Valid values: Absolute, Percent.
	LimitType interface{}
	// The maximum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType is Absolute.
	MaxBandwidthAbs interface{}
	// The maximum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated Smart Access Gateway (SAG) instance.This parameter is required when the value of the LimitType parameter is Percent.
	MaxBandwidthPercent interface{}
	// The minimum bandwidth allowed for the stream specified in the quintuple rule. This parameter is required when the value of the LimitType parameter is Absolute.
	MinBandwidthAbs interface{}
	// The minimum bandwidth percentage allowed for the stream specified in the quintuple rule. It is based on the maximum upstream bandwidth you set for the associated SAG instance.This parameter is required when the value of the LimitType parameter is Percent.
	MinBandwidthPercent interface{}
	// The name of the QoS speed limiting rule..
	Name interface{}
	// The bandwidth type when the speed is limited based on percentage. Valid values: CcnBandwidth, InternetUpBandwidth.The default value is InternetUpBandwidth.
	PercentSourceType interface{}
	// The priority of the specified stream.
	Priority interface{}
	// The instance ID of the QoS.
	QosId interface{}
}
