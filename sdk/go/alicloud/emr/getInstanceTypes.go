// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package emr

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `emr.getInstanceTypes` data source provides a collection of ecs
// instance types available in Alibaba Cloud account when create a emr cluster.
// 
// > **NOTE:** Available in 1.59.0+
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/emr_instance_types.html.markdown.
func GetInstanceTypes(ctx *pulumi.Context, args *GetInstanceTypesArgs, opts ...pulumi.InvokeOption) (*GetInstanceTypesResult, error) {
	var rv GetInstanceTypesResult
	err := ctx.Invoke("alicloud:emr/getInstanceTypes:getInstanceTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceTypes.
type GetInstanceTypesArgs struct {
	// The cluster type of the emr cluster instance. Possible values: `HADOOP`, `KAFKA`, `ZOOKEEPER`, `DRUID`.
	ClusterType string `pulumi:"clusterType"`
	// The destination resource of emr cluster instance
	DestinationResource string `pulumi:"destinationResource"`
	// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
	InstanceChargeType string `pulumi:"instanceChargeType"`
	OutputFile *string `pulumi:"outputFile"`
	// Whether the current storage disk is local or not.
	SupportLocalStorage *bool `pulumi:"supportLocalStorage"`
	// The specific supported node type list. 
	// Possible values may be any one or combination of these: ["MASTER", "CORE", "TASK", "GATEWAY"]
	SupportNodeTypes []string `pulumi:"supportNodeTypes"`
	// The supported resources of specific zoneId.
	ZoneId *string `pulumi:"zoneId"`
}


// A collection of values returned by getInstanceTypes.
type GetInstanceTypesResult struct {
	ClusterType string `pulumi:"clusterType"`
	DestinationResource string `pulumi:"destinationResource"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of emr instance types IDs. 
	Ids []string `pulumi:"ids"`
	InstanceChargeType string `pulumi:"instanceChargeType"`
	OutputFile *string `pulumi:"outputFile"`
	SupportLocalStorage *bool `pulumi:"supportLocalStorage"`
	SupportNodeTypes []string `pulumi:"supportNodeTypes"`
	// A list of emr instance types. Each element contains the following attributes:
	Types []GetInstanceTypesType `pulumi:"types"`
	// The available zone id in Alibaba Cloud account
	ZoneId *string `pulumi:"zoneId"`
}

