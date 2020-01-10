// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an RDS read write splitting connection resource to allocate an Intranet connection string for RDS instance.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/db_read_write_splitting_connection.html.markdown.
type ReadWriteSplittingConnection struct {
	s *pulumi.ResourceState
}

// NewReadWriteSplittingConnection registers a new resource with the given unique name, arguments, and options.
func NewReadWriteSplittingConnection(ctx *pulumi.Context,
	name string, args *ReadWriteSplittingConnectionArgs, opts ...pulumi.ResourceOpt) (*ReadWriteSplittingConnection, error) {
	if args == nil || args.DistributionType == nil {
		return nil, errors.New("missing required argument 'DistributionType'")
	}
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["connectionPrefix"] = nil
		inputs["distributionType"] = nil
		inputs["instanceId"] = nil
		inputs["maxDelayTime"] = nil
		inputs["port"] = nil
		inputs["weight"] = nil
	} else {
		inputs["connectionPrefix"] = args.ConnectionPrefix
		inputs["distributionType"] = args.DistributionType
		inputs["instanceId"] = args.InstanceId
		inputs["maxDelayTime"] = args.MaxDelayTime
		inputs["port"] = args.Port
		inputs["weight"] = args.Weight
	}
	inputs["connectionString"] = nil
	s, err := ctx.RegisterResource("alicloud:rds/readWriteSplittingConnection:ReadWriteSplittingConnection", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ReadWriteSplittingConnection{s: s}, nil
}

// GetReadWriteSplittingConnection gets an existing ReadWriteSplittingConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReadWriteSplittingConnection(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ReadWriteSplittingConnectionState, opts ...pulumi.ResourceOpt) (*ReadWriteSplittingConnection, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["connectionPrefix"] = state.ConnectionPrefix
		inputs["connectionString"] = state.ConnectionString
		inputs["distributionType"] = state.DistributionType
		inputs["instanceId"] = state.InstanceId
		inputs["maxDelayTime"] = state.MaxDelayTime
		inputs["port"] = state.Port
		inputs["weight"] = state.Weight
	}
	s, err := ctx.ReadResource("alicloud:rds/readWriteSplittingConnection:ReadWriteSplittingConnection", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ReadWriteSplittingConnection{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ReadWriteSplittingConnection) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ReadWriteSplittingConnection) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
func (r *ReadWriteSplittingConnection) ConnectionPrefix() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["connectionPrefix"])
}

// Connection instance string.
func (r *ReadWriteSplittingConnection) ConnectionString() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["connectionString"])
}

// Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution. 
func (r *ReadWriteSplittingConnection) DistributionType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["distributionType"])
}

// The Id of instance that can run database.
func (r *ReadWriteSplittingConnection) InstanceId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["instanceId"])
}

// Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.  
func (r *ReadWriteSplittingConnection) MaxDelayTime() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["maxDelayTime"])
}

// Intranet connection port. Valid value: [3001-3999]. Default to 3306.
func (r *ReadWriteSplittingConnection) Port() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["port"])
}

// Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distributionType is set to Custom. 
func (r *ReadWriteSplittingConnection) Weight() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["weight"])
}

// Input properties used for looking up and filtering ReadWriteSplittingConnection resources.
type ReadWriteSplittingConnectionState struct {
	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
	ConnectionPrefix interface{}
	// Connection instance string.
	ConnectionString interface{}
	// Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution. 
	DistributionType interface{}
	// The Id of instance that can run database.
	InstanceId interface{}
	// Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.  
	MaxDelayTime interface{}
	// Intranet connection port. Valid value: [3001-3999]. Default to 3306.
	Port interface{}
	// Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distributionType is set to Custom. 
	Weight interface{}
}

// The set of arguments for constructing a ReadWriteSplittingConnection resource.
type ReadWriteSplittingConnectionArgs struct {
	// Prefix of an Internet connection string. It must be checked for uniqueness. It may consist of lowercase letters, numbers, and underlines, and must start with a letter and have no more than 30 characters. Default to <instance_id> + 'rw'.
	ConnectionPrefix interface{}
	// Read weight distribution mode. Values are as follows: `Standard` indicates automatic weight distribution based on types, `Custom` indicates custom weight distribution. 
	DistributionType interface{}
	// The Id of instance that can run database.
	InstanceId interface{}
	// Delay threshold, in seconds. The value range is 0 to 7200. Default to 30. Read requests are not routed to the read-only instances with a delay greater than the threshold.  
	MaxDelayTime interface{}
	// Intranet connection port. Valid value: [3001-3999]. Default to 3306.
	Port interface{}
	// Read weight distribution. Read weights increase at a step of 100 up to 10,000. Enter weights in the following format: {"Instanceid":"Weight","Instanceid":"Weight"}. This parameter must be set when distributionType is set to Custom. 
	Weight interface{}
}
