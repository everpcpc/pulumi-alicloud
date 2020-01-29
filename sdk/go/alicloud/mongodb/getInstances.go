// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package mongodb

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `mongodb.getInstances` data source provides a collection of MongoDB instances available in Alicloud account.
// Filters support regular expression for the instance name, engine or instance type.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/mongodb_instances.html.markdown.
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("alicloud:mongodb/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// Instance availability zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The ids list of MongoDB instances
	Ids []string `pulumi:"ids"`
	// Sizing of the instance to be queried.
	InstanceClass *string `pulumi:"instanceClass"`
	// Type of the instance to be queried. If it is set to `sharding`, the sharded cluster instances are listed. If it is set to `replicate`, replica set instances are listed. Default value `replicate`.
	InstanceType *string `pulumi:"instanceType"`
	// A regex string to apply to the instance name.
	NameRegex *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}


// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// Instance availability zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ids list of MongoDB instances
	Ids []string `pulumi:"ids"`
	// Sizing of the MongoDB instance.
	InstanceClass *string `pulumi:"instanceClass"`
	// Instance type. Optional values `sharding` or `replicate`.
	InstanceType *string `pulumi:"instanceType"`
	// A list of MongoDB instances. Its every element contains the following attributes:
	Instances []GetInstancesInstance `pulumi:"instances"`
	NameRegex *string `pulumi:"nameRegex"`
	// The names list of MongoDB instances
	Names []string `pulumi:"names"`
	OutputFile *string `pulumi:"outputFile"`
	Tags map[string]interface{} `pulumi:"tags"`
}

