// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hbase

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `hbase.getInstances` data source provides a collection of HBase instances available in Alicloud account.
// Filters support regular expression for the instance name, ids or availability_zone.
//
// > **NOTE:**  Available in 1.67.0+
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("alicloud:hbase/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// Instance availability zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The ids list of HBase instances
	Ids []string `pulumi:"ids"`
	// A regex string to apply to the instance name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ids list of HBase instances
	Ids []string `pulumi:"ids"`
	// A list of HBase instances. Its every element contains the following attributes:
	Instances []GetInstancesInstance `pulumi:"instances"`
	NameRegex *string                `pulumi:"nameRegex"`
	// The names list of HBase instances
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}
