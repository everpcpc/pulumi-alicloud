// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hbr

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a HBR Nas Backup Plan resource.
//
// For information about HBR Nas Backup Plan and how to use it, see [What is Nas Backup Plan](https://www.alibabacloud.com/help/doc-detail/132248.htm).
//
// > **NOTE:** Available in v1.132.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/hbr"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/nas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "tf-testAccHBRNas"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultVault, err := hbr.NewVault(ctx, "defaultVault", &hbr.VaultArgs{
//				VaultName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			defaultFileSystem, err := nas.NewFileSystem(ctx, "defaultFileSystem", &nas.FileSystemArgs{
//				ProtocolType: pulumi.String("NFS"),
//				StorageType:  pulumi.String("Performance"),
//				Description:  pulumi.String(name),
//				EncryptType:  pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			defaultFileSystems := defaultFileSystem.Description.ApplyT(func(description *string) (nas.GetFileSystemsResult, error) {
//				return nas.GetFileSystemsOutput(ctx, nas.GetFileSystemsOutputArgs{
//					ProtocolType:     "NFS",
//					DescriptionRegex: description,
//				}, nil), nil
//			}).(nas.GetFileSystemsResultOutput)
//			_, err = hbr.NewNasBackupPlan(ctx, "defaultNasBackupPlan", &hbr.NasBackupPlanArgs{
//				NasBackupPlanName: pulumi.String(name),
//				FileSystemId:      defaultFileSystem.ID(),
//				Schedule:          pulumi.String("I|1602673264|PT2H"),
//				BackupType:        pulumi.String("COMPLETE"),
//				VaultId:           defaultVault.ID(),
//				CreateTime: defaultFileSystems.ApplyT(func(defaultFileSystems nas.GetFileSystemsResult) (*string, error) {
//					return &defaultFileSystems.Systems[0].CreateTime, nil
//				}).(pulumi.StringPtrOutput),
//				Retention: pulumi.String("2"),
//				Paths: pulumi.StringArray{
//					pulumi.String("/"),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				pulumi.Resource("alicloud_nas_file_system.default"),
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// HBR Nas Backup Plan can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:hbr/nasBackupPlan:NasBackupPlan example <id>
//
// ```
type NasBackupPlan struct {
	pulumi.CustomResourceState

	// Backup type. Valid values: `COMPLETE`.
	BackupType pulumi.StringOutput `pulumi:"backupType"`
	// This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
	//
	// Deprecated: Field 'create_time' has been deprecated from provider version 1.153.0.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The role name created in the original account RAM backup by the cross account managed by the current account.
	CrossAccountRoleName pulumi.StringPtrOutput `pulumi:"crossAccountRoleName"`
	// The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
	CrossAccountType pulumi.StringOutput `pulumi:"crossAccountType"`
	// The original account ID of the cross account backup managed by the current account.
	CrossAccountUserId pulumi.IntPtrOutput `pulumi:"crossAccountUserId"`
	// Whether to disable the backup task. Valid values: `true`, `false`.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// The File System ID of Nas.
	FileSystemId pulumi.StringOutput `pulumi:"fileSystemId"`
	// The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
	NasBackupPlanName pulumi.StringOutput `pulumi:"nasBackupPlanName"`
	// This parameter specifies whether to use Windows VSS to define a backup path.
	Options pulumi.StringPtrOutput `pulumi:"options"`
	// List of backup path. Up to 65536 characters. e.g.`["/home", "/var"]`. **Note** You should at least specify a backup path, empty array not allowed here.
	Paths pulumi.StringArrayOutput `pulumi:"paths"`
	// Backup retention days, the minimum is 1.
	Retention pulumi.StringOutput `pulumi:"retention"`
	// Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// The ID of Backup vault.
	VaultId pulumi.StringOutput `pulumi:"vaultId"`
}

// NewNasBackupPlan registers a new resource with the given unique name, arguments, and options.
func NewNasBackupPlan(ctx *pulumi.Context,
	name string, args *NasBackupPlanArgs, opts ...pulumi.ResourceOption) (*NasBackupPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupType == nil {
		return nil, errors.New("invalid value for required argument 'BackupType'")
	}
	if args.FileSystemId == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemId'")
	}
	if args.NasBackupPlanName == nil {
		return nil, errors.New("invalid value for required argument 'NasBackupPlanName'")
	}
	if args.Paths == nil {
		return nil, errors.New("invalid value for required argument 'Paths'")
	}
	if args.Retention == nil {
		return nil, errors.New("invalid value for required argument 'Retention'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	if args.VaultId == nil {
		return nil, errors.New("invalid value for required argument 'VaultId'")
	}
	var resource NasBackupPlan
	err := ctx.RegisterResource("alicloud:hbr/nasBackupPlan:NasBackupPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNasBackupPlan gets an existing NasBackupPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNasBackupPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NasBackupPlanState, opts ...pulumi.ResourceOption) (*NasBackupPlan, error) {
	var resource NasBackupPlan
	err := ctx.ReadResource("alicloud:hbr/nasBackupPlan:NasBackupPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NasBackupPlan resources.
type nasBackupPlanState struct {
	// Backup type. Valid values: `COMPLETE`.
	BackupType *string `pulumi:"backupType"`
	// This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
	//
	// Deprecated: Field 'create_time' has been deprecated from provider version 1.153.0.
	CreateTime *string `pulumi:"createTime"`
	// The role name created in the original account RAM backup by the cross account managed by the current account.
	CrossAccountRoleName *string `pulumi:"crossAccountRoleName"`
	// The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
	CrossAccountType *string `pulumi:"crossAccountType"`
	// The original account ID of the cross account backup managed by the current account.
	CrossAccountUserId *int `pulumi:"crossAccountUserId"`
	// Whether to disable the backup task. Valid values: `true`, `false`.
	Disabled *bool `pulumi:"disabled"`
	// The File System ID of Nas.
	FileSystemId *string `pulumi:"fileSystemId"`
	// The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
	NasBackupPlanName *string `pulumi:"nasBackupPlanName"`
	// This parameter specifies whether to use Windows VSS to define a backup path.
	Options *string `pulumi:"options"`
	// List of backup path. Up to 65536 characters. e.g.`["/home", "/var"]`. **Note** You should at least specify a backup path, empty array not allowed here.
	Paths []string `pulumi:"paths"`
	// Backup retention days, the minimum is 1.
	Retention *string `pulumi:"retention"`
	// Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
	Schedule *string `pulumi:"schedule"`
	// The ID of Backup vault.
	VaultId *string `pulumi:"vaultId"`
}

type NasBackupPlanState struct {
	// Backup type. Valid values: `COMPLETE`.
	BackupType pulumi.StringPtrInput
	// This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
	//
	// Deprecated: Field 'create_time' has been deprecated from provider version 1.153.0.
	CreateTime pulumi.StringPtrInput
	// The role name created in the original account RAM backup by the cross account managed by the current account.
	CrossAccountRoleName pulumi.StringPtrInput
	// The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
	CrossAccountType pulumi.StringPtrInput
	// The original account ID of the cross account backup managed by the current account.
	CrossAccountUserId pulumi.IntPtrInput
	// Whether to disable the backup task. Valid values: `true`, `false`.
	Disabled pulumi.BoolPtrInput
	// The File System ID of Nas.
	FileSystemId pulumi.StringPtrInput
	// The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
	NasBackupPlanName pulumi.StringPtrInput
	// This parameter specifies whether to use Windows VSS to define a backup path.
	Options pulumi.StringPtrInput
	// List of backup path. Up to 65536 characters. e.g.`["/home", "/var"]`. **Note** You should at least specify a backup path, empty array not allowed here.
	Paths pulumi.StringArrayInput
	// Backup retention days, the minimum is 1.
	Retention pulumi.StringPtrInput
	// Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
	Schedule pulumi.StringPtrInput
	// The ID of Backup vault.
	VaultId pulumi.StringPtrInput
}

func (NasBackupPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*nasBackupPlanState)(nil)).Elem()
}

type nasBackupPlanArgs struct {
	// Backup type. Valid values: `COMPLETE`.
	BackupType string `pulumi:"backupType"`
	// This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
	//
	// Deprecated: Field 'create_time' has been deprecated from provider version 1.153.0.
	CreateTime *string `pulumi:"createTime"`
	// The role name created in the original account RAM backup by the cross account managed by the current account.
	CrossAccountRoleName *string `pulumi:"crossAccountRoleName"`
	// The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
	CrossAccountType *string `pulumi:"crossAccountType"`
	// The original account ID of the cross account backup managed by the current account.
	CrossAccountUserId *int `pulumi:"crossAccountUserId"`
	// Whether to disable the backup task. Valid values: `true`, `false`.
	Disabled *bool `pulumi:"disabled"`
	// The File System ID of Nas.
	FileSystemId string `pulumi:"fileSystemId"`
	// The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
	NasBackupPlanName string `pulumi:"nasBackupPlanName"`
	// This parameter specifies whether to use Windows VSS to define a backup path.
	Options *string `pulumi:"options"`
	// List of backup path. Up to 65536 characters. e.g.`["/home", "/var"]`. **Note** You should at least specify a backup path, empty array not allowed here.
	Paths []string `pulumi:"paths"`
	// Backup retention days, the minimum is 1.
	Retention string `pulumi:"retention"`
	// Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
	Schedule string `pulumi:"schedule"`
	// The ID of Backup vault.
	VaultId string `pulumi:"vaultId"`
}

// The set of arguments for constructing a NasBackupPlan resource.
type NasBackupPlanArgs struct {
	// Backup type. Valid values: `COMPLETE`.
	BackupType pulumi.StringInput
	// This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
	//
	// Deprecated: Field 'create_time' has been deprecated from provider version 1.153.0.
	CreateTime pulumi.StringPtrInput
	// The role name created in the original account RAM backup by the cross account managed by the current account.
	CrossAccountRoleName pulumi.StringPtrInput
	// The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
	CrossAccountType pulumi.StringPtrInput
	// The original account ID of the cross account backup managed by the current account.
	CrossAccountUserId pulumi.IntPtrInput
	// Whether to disable the backup task. Valid values: `true`, `false`.
	Disabled pulumi.BoolPtrInput
	// The File System ID of Nas.
	FileSystemId pulumi.StringInput
	// The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
	NasBackupPlanName pulumi.StringInput
	// This parameter specifies whether to use Windows VSS to define a backup path.
	Options pulumi.StringPtrInput
	// List of backup path. Up to 65536 characters. e.g.`["/home", "/var"]`. **Note** You should at least specify a backup path, empty array not allowed here.
	Paths pulumi.StringArrayInput
	// Backup retention days, the minimum is 1.
	Retention pulumi.StringInput
	// Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
	Schedule pulumi.StringInput
	// The ID of Backup vault.
	VaultId pulumi.StringInput
}

func (NasBackupPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nasBackupPlanArgs)(nil)).Elem()
}

type NasBackupPlanInput interface {
	pulumi.Input

	ToNasBackupPlanOutput() NasBackupPlanOutput
	ToNasBackupPlanOutputWithContext(ctx context.Context) NasBackupPlanOutput
}

func (*NasBackupPlan) ElementType() reflect.Type {
	return reflect.TypeOf((**NasBackupPlan)(nil)).Elem()
}

func (i *NasBackupPlan) ToNasBackupPlanOutput() NasBackupPlanOutput {
	return i.ToNasBackupPlanOutputWithContext(context.Background())
}

func (i *NasBackupPlan) ToNasBackupPlanOutputWithContext(ctx context.Context) NasBackupPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NasBackupPlanOutput)
}

// NasBackupPlanArrayInput is an input type that accepts NasBackupPlanArray and NasBackupPlanArrayOutput values.
// You can construct a concrete instance of `NasBackupPlanArrayInput` via:
//
//	NasBackupPlanArray{ NasBackupPlanArgs{...} }
type NasBackupPlanArrayInput interface {
	pulumi.Input

	ToNasBackupPlanArrayOutput() NasBackupPlanArrayOutput
	ToNasBackupPlanArrayOutputWithContext(context.Context) NasBackupPlanArrayOutput
}

type NasBackupPlanArray []NasBackupPlanInput

func (NasBackupPlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NasBackupPlan)(nil)).Elem()
}

func (i NasBackupPlanArray) ToNasBackupPlanArrayOutput() NasBackupPlanArrayOutput {
	return i.ToNasBackupPlanArrayOutputWithContext(context.Background())
}

func (i NasBackupPlanArray) ToNasBackupPlanArrayOutputWithContext(ctx context.Context) NasBackupPlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NasBackupPlanArrayOutput)
}

// NasBackupPlanMapInput is an input type that accepts NasBackupPlanMap and NasBackupPlanMapOutput values.
// You can construct a concrete instance of `NasBackupPlanMapInput` via:
//
//	NasBackupPlanMap{ "key": NasBackupPlanArgs{...} }
type NasBackupPlanMapInput interface {
	pulumi.Input

	ToNasBackupPlanMapOutput() NasBackupPlanMapOutput
	ToNasBackupPlanMapOutputWithContext(context.Context) NasBackupPlanMapOutput
}

type NasBackupPlanMap map[string]NasBackupPlanInput

func (NasBackupPlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NasBackupPlan)(nil)).Elem()
}

func (i NasBackupPlanMap) ToNasBackupPlanMapOutput() NasBackupPlanMapOutput {
	return i.ToNasBackupPlanMapOutputWithContext(context.Background())
}

func (i NasBackupPlanMap) ToNasBackupPlanMapOutputWithContext(ctx context.Context) NasBackupPlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NasBackupPlanMapOutput)
}

type NasBackupPlanOutput struct{ *pulumi.OutputState }

func (NasBackupPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NasBackupPlan)(nil)).Elem()
}

func (o NasBackupPlanOutput) ToNasBackupPlanOutput() NasBackupPlanOutput {
	return o
}

func (o NasBackupPlanOutput) ToNasBackupPlanOutputWithContext(ctx context.Context) NasBackupPlanOutput {
	return o
}

// Backup type. Valid values: `COMPLETE`.
func (o NasBackupPlanOutput) BackupType() pulumi.StringOutput {
	return o.ApplyT(func(v *NasBackupPlan) pulumi.StringOutput { return v.BackupType }).(pulumi.StringOutput)
}

// This field has been deprecated from provider version 1.153.0+. The creation time of NAS file system. **Note** The time format of the API adopts the ISO 8601, such as `2021-07-09T15:45:30CST` or `2021-07-09T07:45:30Z`.
//
// Deprecated: Field 'create_time' has been deprecated from provider version 1.153.0.
func (o NasBackupPlanOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NasBackupPlan) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The role name created in the original account RAM backup by the cross account managed by the current account.
func (o NasBackupPlanOutput) CrossAccountRoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NasBackupPlan) pulumi.StringPtrOutput { return v.CrossAccountRoleName }).(pulumi.StringPtrOutput)
}

// The type of the cross account backup. Valid values: `SELF_ACCOUNT`, `CROSS_ACCOUNT`.
func (o NasBackupPlanOutput) CrossAccountType() pulumi.StringOutput {
	return o.ApplyT(func(v *NasBackupPlan) pulumi.StringOutput { return v.CrossAccountType }).(pulumi.StringOutput)
}

// The original account ID of the cross account backup managed by the current account.
func (o NasBackupPlanOutput) CrossAccountUserId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NasBackupPlan) pulumi.IntPtrOutput { return v.CrossAccountUserId }).(pulumi.IntPtrOutput)
}

// Whether to disable the backup task. Valid values: `true`, `false`.
func (o NasBackupPlanOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *NasBackupPlan) pulumi.BoolOutput { return v.Disabled }).(pulumi.BoolOutput)
}

// The File System ID of Nas.
func (o NasBackupPlanOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *NasBackupPlan) pulumi.StringOutput { return v.FileSystemId }).(pulumi.StringOutput)
}

// The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
func (o NasBackupPlanOutput) NasBackupPlanName() pulumi.StringOutput {
	return o.ApplyT(func(v *NasBackupPlan) pulumi.StringOutput { return v.NasBackupPlanName }).(pulumi.StringOutput)
}

// This parameter specifies whether to use Windows VSS to define a backup path.
func (o NasBackupPlanOutput) Options() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NasBackupPlan) pulumi.StringPtrOutput { return v.Options }).(pulumi.StringPtrOutput)
}

// List of backup path. Up to 65536 characters. e.g.`["/home", "/var"]`. **Note** You should at least specify a backup path, empty array not allowed here.
func (o NasBackupPlanOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NasBackupPlan) pulumi.StringArrayOutput { return v.Paths }).(pulumi.StringArrayOutput)
}

// Backup retention days, the minimum is 1.
func (o NasBackupPlanOutput) Retention() pulumi.StringOutput {
	return o.ApplyT(func(v *NasBackupPlan) pulumi.StringOutput { return v.Retention }).(pulumi.StringOutput)
}

// Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered.
func (o NasBackupPlanOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v *NasBackupPlan) pulumi.StringOutput { return v.Schedule }).(pulumi.StringOutput)
}

// The ID of Backup vault.
func (o NasBackupPlanOutput) VaultId() pulumi.StringOutput {
	return o.ApplyT(func(v *NasBackupPlan) pulumi.StringOutput { return v.VaultId }).(pulumi.StringOutput)
}

type NasBackupPlanArrayOutput struct{ *pulumi.OutputState }

func (NasBackupPlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NasBackupPlan)(nil)).Elem()
}

func (o NasBackupPlanArrayOutput) ToNasBackupPlanArrayOutput() NasBackupPlanArrayOutput {
	return o
}

func (o NasBackupPlanArrayOutput) ToNasBackupPlanArrayOutputWithContext(ctx context.Context) NasBackupPlanArrayOutput {
	return o
}

func (o NasBackupPlanArrayOutput) Index(i pulumi.IntInput) NasBackupPlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NasBackupPlan {
		return vs[0].([]*NasBackupPlan)[vs[1].(int)]
	}).(NasBackupPlanOutput)
}

type NasBackupPlanMapOutput struct{ *pulumi.OutputState }

func (NasBackupPlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NasBackupPlan)(nil)).Elem()
}

func (o NasBackupPlanMapOutput) ToNasBackupPlanMapOutput() NasBackupPlanMapOutput {
	return o
}

func (o NasBackupPlanMapOutput) ToNasBackupPlanMapOutputWithContext(ctx context.Context) NasBackupPlanMapOutput {
	return o
}

func (o NasBackupPlanMapOutput) MapIndex(k pulumi.StringInput) NasBackupPlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NasBackupPlan {
		return vs[0].(map[string]*NasBackupPlan)[vs[1].(string)]
	}).(NasBackupPlanOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NasBackupPlanInput)(nil)).Elem(), &NasBackupPlan{})
	pulumi.RegisterInputType(reflect.TypeOf((*NasBackupPlanArrayInput)(nil)).Elem(), NasBackupPlanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NasBackupPlanMapInput)(nil)).Elem(), NasBackupPlanMap{})
	pulumi.RegisterOutputType(NasBackupPlanOutput{})
	pulumi.RegisterOutputType(NasBackupPlanArrayOutput{})
	pulumi.RegisterOutputType(NasBackupPlanMapOutput{})
}
