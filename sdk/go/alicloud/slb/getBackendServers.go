// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package slb

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides the server load balancer backend servers related to a server load balancer..
//
// > **NOTE:** Available in 1.53.0+
func GetBackendServers(ctx *pulumi.Context, args *GetBackendServersArgs, opts ...pulumi.InvokeOption) (*GetBackendServersResult, error) {
	var rv GetBackendServersResult
	err := ctx.Invoke("alicloud:slb/getBackendServers:getBackendServers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackendServers.
type GetBackendServersArgs struct {
	// List of attached ECS instance IDs.
	Ids []string `pulumi:"ids"`
	// ID of the SLB with attachments.
	LoadBalancerId string  `pulumi:"loadBalancerId"`
	OutputFile     *string `pulumi:"outputFile"`
}

// A collection of values returned by getBackendServers.
type GetBackendServersResult struct {
	BackendServers []GetBackendServersBackendServer `pulumi:"backendServers"`
	// id is the provider-assigned unique ID for this managed resource.
	Id             string   `pulumi:"id"`
	Ids            []string `pulumi:"ids"`
	LoadBalancerId string   `pulumi:"loadBalancerId"`
	OutputFile     *string  `pulumi:"outputFile"`
}
