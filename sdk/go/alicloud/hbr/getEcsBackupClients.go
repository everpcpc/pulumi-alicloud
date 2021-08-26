// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hbr

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Hbr Ecs Backup Clients of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.132.0+.
func GetEcsBackupClients(ctx *pulumi.Context, args *GetEcsBackupClientsArgs, opts ...pulumi.InvokeOption) (*GetEcsBackupClientsResult, error) {
	var rv GetEcsBackupClientsResult
	err := ctx.Invoke("alicloud:hbr/getEcsBackupClients:getEcsBackupClients", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEcsBackupClients.
type GetEcsBackupClientsArgs struct {
	// A list of Ecs Backup Client IDs.
	Ids []string `pulumi:"ids"`
	// A list of ECS Instance IDs.
	InstanceIds []string `pulumi:"instanceIds"`
	OutputFile  *string  `pulumi:"outputFile"`
	// The status of the resource.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getEcsBackupClients.
type GetEcsBackupClientsResult struct {
	Clients []GetEcsBackupClientsClient `pulumi:"clients"`
	// The provider-assigned unique ID for this managed resource.
	Id          string   `pulumi:"id"`
	Ids         []string `pulumi:"ids"`
	InstanceIds []string `pulumi:"instanceIds"`
	OutputFile  *string  `pulumi:"outputFile"`
	Status      *string  `pulumi:"status"`
}
