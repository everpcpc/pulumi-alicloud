// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the Rds Parameter Groups of the current Alibaba Cloud user.
//
// > **NOTE:** Available in v1.196.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			backups, err := rds.GetCrossRegionBackups(ctx, &rds.GetCrossRegionBackupsArgs{
//				DbInstanceId: "example_value",
//				StartTime:    pulumi.StringRef("2022-12-01T00:00:00Z"),
//				EndTime:      pulumi.StringRef("2022-12-16T00:00:00Z"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstRdsCrossRegionBackups", backups.Backups[0].Id)
//			return nil
//		})
//	}
//
// ```
func GetCrossRegionBackups(ctx *pulumi.Context, args *GetCrossRegionBackupsArgs, opts ...pulumi.InvokeOption) (*GetCrossRegionBackupsResult, error) {
	var rv GetCrossRegionBackupsResult
	err := ctx.Invoke("alicloud:rds/getCrossRegionBackups:getCrossRegionBackups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCrossRegionBackups.
type GetCrossRegionBackupsArgs struct {
	// The ID of the cross-region data backup file.
	BackupId *string `pulumi:"backupId"`
	// The ID of the cross-region data backup file.
	CrossBackupId *string `pulumi:"crossBackupId"`
	// The ID of the destination region where the cross-region data backup file of the instance is stored.
	CrossBackupRegion *string `pulumi:"crossBackupRegion"`
	// The db instance id.
	DbInstanceId string `pulumi:"dbInstanceId"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	EndTime *string `pulumi:"endTime"`
	// A list of Cross Region Backup IDs.
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The ID of the resource group.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	StartTime *string `pulumi:"startTime"`
}

// A collection of values returned by getCrossRegionBackups.
type GetCrossRegionBackupsResult struct {
	BackupId          *string                       `pulumi:"backupId"`
	Backups           []GetCrossRegionBackupsBackup `pulumi:"backups"`
	CrossBackupId     *string                       `pulumi:"crossBackupId"`
	CrossBackupRegion *string                       `pulumi:"crossBackupRegion"`
	DbInstanceId      string                        `pulumi:"dbInstanceId"`
	EndTime           *string                       `pulumi:"endTime"`
	// The provider-assigned unique ID for this managed resource.
	Id              string   `pulumi:"id"`
	Ids             []string `pulumi:"ids"`
	OutputFile      *string  `pulumi:"outputFile"`
	ResourceGroupId *string  `pulumi:"resourceGroupId"`
	StartTime       *string  `pulumi:"startTime"`
}

func GetCrossRegionBackupsOutput(ctx *pulumi.Context, args GetCrossRegionBackupsOutputArgs, opts ...pulumi.InvokeOption) GetCrossRegionBackupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCrossRegionBackupsResult, error) {
			args := v.(GetCrossRegionBackupsArgs)
			r, err := GetCrossRegionBackups(ctx, &args, opts...)
			var s GetCrossRegionBackupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCrossRegionBackupsResultOutput)
}

// A collection of arguments for invoking getCrossRegionBackups.
type GetCrossRegionBackupsOutputArgs struct {
	// The ID of the cross-region data backup file.
	BackupId pulumi.StringPtrInput `pulumi:"backupId"`
	// The ID of the cross-region data backup file.
	CrossBackupId pulumi.StringPtrInput `pulumi:"crossBackupId"`
	// The ID of the destination region where the cross-region data backup file of the instance is stored.
	CrossBackupRegion pulumi.StringPtrInput `pulumi:"crossBackupRegion"`
	// The db instance id.
	DbInstanceId pulumi.StringInput `pulumi:"dbInstanceId"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	EndTime pulumi.StringPtrInput `pulumi:"endTime"`
	// A list of Cross Region Backup IDs.
	Ids        pulumi.StringArrayInput `pulumi:"ids"`
	OutputFile pulumi.StringPtrInput   `pulumi:"outputFile"`
	// The ID of the resource group.
	ResourceGroupId pulumi.StringPtrInput `pulumi:"resourceGroupId"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	StartTime pulumi.StringPtrInput `pulumi:"startTime"`
}

func (GetCrossRegionBackupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCrossRegionBackupsArgs)(nil)).Elem()
}

// A collection of values returned by getCrossRegionBackups.
type GetCrossRegionBackupsResultOutput struct{ *pulumi.OutputState }

func (GetCrossRegionBackupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCrossRegionBackupsResult)(nil)).Elem()
}

func (o GetCrossRegionBackupsResultOutput) ToGetCrossRegionBackupsResultOutput() GetCrossRegionBackupsResultOutput {
	return o
}

func (o GetCrossRegionBackupsResultOutput) ToGetCrossRegionBackupsResultOutputWithContext(ctx context.Context) GetCrossRegionBackupsResultOutput {
	return o
}

func (o GetCrossRegionBackupsResultOutput) BackupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCrossRegionBackupsResult) *string { return v.BackupId }).(pulumi.StringPtrOutput)
}

func (o GetCrossRegionBackupsResultOutput) Backups() GetCrossRegionBackupsBackupArrayOutput {
	return o.ApplyT(func(v GetCrossRegionBackupsResult) []GetCrossRegionBackupsBackup { return v.Backups }).(GetCrossRegionBackupsBackupArrayOutput)
}

func (o GetCrossRegionBackupsResultOutput) CrossBackupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCrossRegionBackupsResult) *string { return v.CrossBackupId }).(pulumi.StringPtrOutput)
}

func (o GetCrossRegionBackupsResultOutput) CrossBackupRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCrossRegionBackupsResult) *string { return v.CrossBackupRegion }).(pulumi.StringPtrOutput)
}

func (o GetCrossRegionBackupsResultOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCrossRegionBackupsResult) string { return v.DbInstanceId }).(pulumi.StringOutput)
}

func (o GetCrossRegionBackupsResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCrossRegionBackupsResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCrossRegionBackupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCrossRegionBackupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCrossRegionBackupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCrossRegionBackupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetCrossRegionBackupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCrossRegionBackupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetCrossRegionBackupsResultOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCrossRegionBackupsResult) *string { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o GetCrossRegionBackupsResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCrossRegionBackupsResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCrossRegionBackupsResultOutput{})
}