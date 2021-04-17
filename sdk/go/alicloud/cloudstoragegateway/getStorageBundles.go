// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudstoragegateway

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetStorageBundles(ctx *pulumi.Context, args *GetStorageBundlesArgs, opts ...pulumi.InvokeOption) (*GetStorageBundlesResult, error) {
	var rv GetStorageBundlesResult
	err := ctx.Invoke("alicloud:cloudstoragegateway/getStorageBundles:getStorageBundles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStorageBundles.
type GetStorageBundlesArgs struct {
	BackendBucketRegionId string   `pulumi:"backendBucketRegionId"`
	Ids                   []string `pulumi:"ids"`
	NameRegex             *string  `pulumi:"nameRegex"`
	OutputFile            *string  `pulumi:"outputFile"`
}

// A collection of values returned by getStorageBundles.
type GetStorageBundlesResult struct {
	BackendBucketRegionId string                    `pulumi:"backendBucketRegionId"`
	Bundles               []GetStorageBundlesBundle `pulumi:"bundles"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	Names      []string `pulumi:"names"`
	OutputFile *string  `pulumi:"outputFile"`
}
