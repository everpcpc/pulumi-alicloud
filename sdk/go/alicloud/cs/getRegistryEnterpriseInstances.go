// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides a list Container Registry Enterprise Edition instances on Alibaba Cloud.
//
// > **NOTE:** Available in v1.86.0+
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/cs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "my-instances"
// 		opt1 := "my-instances-json"
// 		myInstances, err := cs.GetRegistryEnterpriseInstances(ctx, &cs.GetRegistryEnterpriseInstancesArgs{
// 			NameRegex:  &opt0,
// 			OutputFile: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("output", myInstances.Instances)
// 		return nil
// 	})
// }
// ```
func GetRegistryEnterpriseInstances(ctx *pulumi.Context, args *GetRegistryEnterpriseInstancesArgs, opts ...pulumi.InvokeOption) (*GetRegistryEnterpriseInstancesResult, error) {
	var rv GetRegistryEnterpriseInstancesResult
	err := ctx.Invoke("alicloud:cs/getRegistryEnterpriseInstances:getRegistryEnterpriseInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegistryEnterpriseInstances.
type GetRegistryEnterpriseInstancesArgs struct {
	// Default to `true`. Set it to true can output instance authorization token.
	EnableDetails *bool `pulumi:"enableDetails"`
	// A list of ids to filter results by instance id.
	Ids []string `pulumi:"ids"`
	// A regex string to filter results by instance name.
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getRegistryEnterpriseInstances.
type GetRegistryEnterpriseInstancesResult struct {
	EnableDetails *bool `pulumi:"enableDetails"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of matched Container Registry Enterprise Edition instances. Its element is an instance uuid.
	Ids []string `pulumi:"ids"`
	// A list of matched Container Registry Enterprise Editioninstances. Each element contains the following attributes:
	Instances []GetRegistryEnterpriseInstancesInstance `pulumi:"instances"`
	NameRegex *string                                  `pulumi:"nameRegex"`
	// A list of instance names.
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}
