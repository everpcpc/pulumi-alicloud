// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source provides CEN Route Service available to the user.
//
// > **NOTE:** Available in v1.102.0+
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/cen"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := cen.GetRouteServices(ctx, &cen.GetRouteServicesArgs{
// 			CenId: "cen-7qthudw0ll6jmc****",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstCenRouteServiceId", example.Services[0].Id)
// 		return nil
// 	})
// }
// ```
func GetRouteServices(ctx *pulumi.Context, args *GetRouteServicesArgs, opts ...pulumi.InvokeOption) (*GetRouteServicesResult, error) {
	var rv GetRouteServicesResult
	err := ctx.Invoke("alicloud:cen/getRouteServices:getRouteServices", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteServices.
type GetRouteServicesArgs struct {
	// The region of the network instances that access the cloud services.
	AccessRegionId *string `pulumi:"accessRegionId"`
	// -(Required, ForceNew) The ID of the CEN instance.
	CenId string `pulumi:"cenId"`
	// -(Optional, ForceNew) The domain name or IP address of the cloud service.
	Host *string `pulumi:"host"`
	// The region of the cloud service.
	HostRegionId *string `pulumi:"hostRegionId"`
	// The VPC associated with the cloud service.
	HostVpcId  *string `pulumi:"hostVpcId"`
	OutputFile *string `pulumi:"outputFile"`
	// The status of the cloud service. Valid values: `Active`, `Creating` and `Deleting`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getRouteServices.
type GetRouteServicesResult struct {
	// The region of the network instances that access the cloud services.
	AccessRegionId *string `pulumi:"accessRegionId"`
	// The ID of the CEN instance.
	CenId string `pulumi:"cenId"`
	// The domain name or IP address of the cloud service.
	Host *string `pulumi:"host"`
	// The region of the cloud service.
	HostRegionId *string `pulumi:"hostRegionId"`
	// The VPC associated with the cloud service.
	HostVpcId *string `pulumi:"hostVpcId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of CEN Route Service IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// A list of CEN Route Services. Each element contains the following attributes:
	Services []GetRouteServicesService `pulumi:"services"`
	// The status of the cloud service.
	Status *string `pulumi:"status"`
}
