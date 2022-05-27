// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hbr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a HBR Ots Backup Plan resource.
//
// For information about HBR Ots Backup Plan and how to use it, see [What is Ots Backup Plan](https://www.alibabacloud.com/help/en/hybrid-backup-recovery/latest/overview).
//
// > **NOTE:** Available in v1.163.0+.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/hbr"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ots"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		name := "testAcc"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		_, err := hbr.NewVault(ctx, "default", &hbr.VaultArgs{
// 			VaultName: pulumi.String(name),
// 			VaultType: pulumi.String("OTS_BACKUP"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		foo, err := ots.NewInstance(ctx, "foo", &ots.InstanceArgs{
// 			Description: pulumi.String(name),
// 			AccessedBy:  pulumi.String("Any"),
// 			Tags: pulumi.AnyMap{
// 				"Created": pulumi.Any("TF"),
// 				"For":     pulumi.Any("acceptance test"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		basic, err := ots.NewTable(ctx, "basic", &ots.TableArgs{
// 			InstanceName: foo.Name,
// 			TableName:    pulumi.String(name),
// 			PrimaryKeys: ots.TablePrimaryKeyArray{
// 				&ots.TablePrimaryKeyArgs{
// 					Name: pulumi.String("pk1"),
// 					Type: pulumi.String("Integer"),
// 				},
// 			},
// 			TimeToLive:                -1,
// 			MaxVersion:                pulumi.Int(1),
// 			DeviationCellVersionInSec: pulumi.String("1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = hbr.NewOtsBackupPlan(ctx, "example", &hbr.OtsBackupPlanArgs{
// 			OtsBackupPlanName: pulumi.String(name),
// 			VaultId:           _default.ID(),
// 			BackupType:        pulumi.String("COMPLETE"),
// 			Schedule:          pulumi.String("I|1602673264|PT2H"),
// 			Retention:         pulumi.String("2"),
// 			InstanceName:      foo.Name,
// 			OtsDetails: hbr.OtsBackupPlanOtsDetailArray{
// 				&hbr.OtsBackupPlanOtsDetailArgs{
// 					TableNames: pulumi.StringArray{
// 						basic.TableName,
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// HBR Ots Backup Plan can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:hbr/otsBackupPlan:OtsBackupPlan example <id>
// ```
type OtsBackupPlan struct {
	pulumi.CustomResourceState

	// The name of the tableStore instance. Valid values: `COMPLETE`, `INCREMENTAL`. **Note:** Required while sourceType equals `OTS_TABLE`.
	BackupType pulumi.StringOutput `pulumi:"backupType"`
	// Whether to disable the backup task. Valid values: true, false.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// The name of the Table store instance. **Note:** Required while sourceType equals `OTS_TABLE`.
	InstanceName pulumi.StringPtrOutput `pulumi:"instanceName"`
	// The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
	OtsBackupPlanName pulumi.StringOutput `pulumi:"otsBackupPlanName"`
	// The details about the Table store instance. See the following `Block otsDetail`. **Note:** Required while sourceType equals `OTS_TABLE`.
	OtsDetails OtsBackupPlanOtsDetailArrayOutput `pulumi:"otsDetails"`
	// Backup retention days, the minimum is 1. **Note:** Required while sourceType equals `OTS_TABLE`.
	Retention pulumi.StringOutput `pulumi:"retention"`
	// The backup plan rule. See the following `Block rules`. **Note:** Required while sourceType equals `OTS_TABLE`.
	Rules OtsBackupPlanRuleArrayOutput `pulumi:"rules"`
	// Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered. **Note:** Required while sourceType equals `OTS_TABLE`.
	//
	// Deprecated: Field 'schedule' has been deprecated from version 1.163.0. Use 'rules' instead.
	Schedule pulumi.StringPtrOutput `pulumi:"schedule"`
	// The ID of backup vault.
	VaultId pulumi.StringPtrOutput `pulumi:"vaultId"`
}

// NewOtsBackupPlan registers a new resource with the given unique name, arguments, and options.
func NewOtsBackupPlan(ctx *pulumi.Context,
	name string, args *OtsBackupPlanArgs, opts ...pulumi.ResourceOption) (*OtsBackupPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupType == nil {
		return nil, errors.New("invalid value for required argument 'BackupType'")
	}
	if args.OtsBackupPlanName == nil {
		return nil, errors.New("invalid value for required argument 'OtsBackupPlanName'")
	}
	if args.Retention == nil {
		return nil, errors.New("invalid value for required argument 'Retention'")
	}
	var resource OtsBackupPlan
	err := ctx.RegisterResource("alicloud:hbr/otsBackupPlan:OtsBackupPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOtsBackupPlan gets an existing OtsBackupPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOtsBackupPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OtsBackupPlanState, opts ...pulumi.ResourceOption) (*OtsBackupPlan, error) {
	var resource OtsBackupPlan
	err := ctx.ReadResource("alicloud:hbr/otsBackupPlan:OtsBackupPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OtsBackupPlan resources.
type otsBackupPlanState struct {
	// The name of the tableStore instance. Valid values: `COMPLETE`, `INCREMENTAL`. **Note:** Required while sourceType equals `OTS_TABLE`.
	BackupType *string `pulumi:"backupType"`
	// Whether to disable the backup task. Valid values: true, false.
	Disabled *bool `pulumi:"disabled"`
	// The name of the Table store instance. **Note:** Required while sourceType equals `OTS_TABLE`.
	InstanceName *string `pulumi:"instanceName"`
	// The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
	OtsBackupPlanName *string `pulumi:"otsBackupPlanName"`
	// The details about the Table store instance. See the following `Block otsDetail`. **Note:** Required while sourceType equals `OTS_TABLE`.
	OtsDetails []OtsBackupPlanOtsDetail `pulumi:"otsDetails"`
	// Backup retention days, the minimum is 1. **Note:** Required while sourceType equals `OTS_TABLE`.
	Retention *string `pulumi:"retention"`
	// The backup plan rule. See the following `Block rules`. **Note:** Required while sourceType equals `OTS_TABLE`.
	Rules []OtsBackupPlanRule `pulumi:"rules"`
	// Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered. **Note:** Required while sourceType equals `OTS_TABLE`.
	//
	// Deprecated: Field 'schedule' has been deprecated from version 1.163.0. Use 'rules' instead.
	Schedule *string `pulumi:"schedule"`
	// The ID of backup vault.
	VaultId *string `pulumi:"vaultId"`
}

type OtsBackupPlanState struct {
	// The name of the tableStore instance. Valid values: `COMPLETE`, `INCREMENTAL`. **Note:** Required while sourceType equals `OTS_TABLE`.
	BackupType pulumi.StringPtrInput
	// Whether to disable the backup task. Valid values: true, false.
	Disabled pulumi.BoolPtrInput
	// The name of the Table store instance. **Note:** Required while sourceType equals `OTS_TABLE`.
	InstanceName pulumi.StringPtrInput
	// The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
	OtsBackupPlanName pulumi.StringPtrInput
	// The details about the Table store instance. See the following `Block otsDetail`. **Note:** Required while sourceType equals `OTS_TABLE`.
	OtsDetails OtsBackupPlanOtsDetailArrayInput
	// Backup retention days, the minimum is 1. **Note:** Required while sourceType equals `OTS_TABLE`.
	Retention pulumi.StringPtrInput
	// The backup plan rule. See the following `Block rules`. **Note:** Required while sourceType equals `OTS_TABLE`.
	Rules OtsBackupPlanRuleArrayInput
	// Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered. **Note:** Required while sourceType equals `OTS_TABLE`.
	//
	// Deprecated: Field 'schedule' has been deprecated from version 1.163.0. Use 'rules' instead.
	Schedule pulumi.StringPtrInput
	// The ID of backup vault.
	VaultId pulumi.StringPtrInput
}

func (OtsBackupPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*otsBackupPlanState)(nil)).Elem()
}

type otsBackupPlanArgs struct {
	// The name of the tableStore instance. Valid values: `COMPLETE`, `INCREMENTAL`. **Note:** Required while sourceType equals `OTS_TABLE`.
	BackupType string `pulumi:"backupType"`
	// Whether to disable the backup task. Valid values: true, false.
	Disabled *bool `pulumi:"disabled"`
	// The name of the Table store instance. **Note:** Required while sourceType equals `OTS_TABLE`.
	InstanceName *string `pulumi:"instanceName"`
	// The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
	OtsBackupPlanName string `pulumi:"otsBackupPlanName"`
	// The details about the Table store instance. See the following `Block otsDetail`. **Note:** Required while sourceType equals `OTS_TABLE`.
	OtsDetails []OtsBackupPlanOtsDetail `pulumi:"otsDetails"`
	// Backup retention days, the minimum is 1. **Note:** Required while sourceType equals `OTS_TABLE`.
	Retention string `pulumi:"retention"`
	// The backup plan rule. See the following `Block rules`. **Note:** Required while sourceType equals `OTS_TABLE`.
	Rules []OtsBackupPlanRule `pulumi:"rules"`
	// Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered. **Note:** Required while sourceType equals `OTS_TABLE`.
	//
	// Deprecated: Field 'schedule' has been deprecated from version 1.163.0. Use 'rules' instead.
	Schedule *string `pulumi:"schedule"`
	// The ID of backup vault.
	VaultId *string `pulumi:"vaultId"`
}

// The set of arguments for constructing a OtsBackupPlan resource.
type OtsBackupPlanArgs struct {
	// The name of the tableStore instance. Valid values: `COMPLETE`, `INCREMENTAL`. **Note:** Required while sourceType equals `OTS_TABLE`.
	BackupType pulumi.StringInput
	// Whether to disable the backup task. Valid values: true, false.
	Disabled pulumi.BoolPtrInput
	// The name of the Table store instance. **Note:** Required while sourceType equals `OTS_TABLE`.
	InstanceName pulumi.StringPtrInput
	// The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
	OtsBackupPlanName pulumi.StringInput
	// The details about the Table store instance. See the following `Block otsDetail`. **Note:** Required while sourceType equals `OTS_TABLE`.
	OtsDetails OtsBackupPlanOtsDetailArrayInput
	// Backup retention days, the minimum is 1. **Note:** Required while sourceType equals `OTS_TABLE`.
	Retention pulumi.StringInput
	// The backup plan rule. See the following `Block rules`. **Note:** Required while sourceType equals `OTS_TABLE`.
	Rules OtsBackupPlanRuleArrayInput
	// Backup strategy. Optional format: `I|{startTime}|{interval}`. It means to execute a backup task every `{interval}` starting from `{startTime}`. The backup task for the elapsed time will not be compensated. If the last backup task has not completed yet, the next backup task will not be triggered. **Note:** Required while sourceType equals `OTS_TABLE`.
	//
	// Deprecated: Field 'schedule' has been deprecated from version 1.163.0. Use 'rules' instead.
	Schedule pulumi.StringPtrInput
	// The ID of backup vault.
	VaultId pulumi.StringPtrInput
}

func (OtsBackupPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*otsBackupPlanArgs)(nil)).Elem()
}

type OtsBackupPlanInput interface {
	pulumi.Input

	ToOtsBackupPlanOutput() OtsBackupPlanOutput
	ToOtsBackupPlanOutputWithContext(ctx context.Context) OtsBackupPlanOutput
}

func (*OtsBackupPlan) ElementType() reflect.Type {
	return reflect.TypeOf((**OtsBackupPlan)(nil)).Elem()
}

func (i *OtsBackupPlan) ToOtsBackupPlanOutput() OtsBackupPlanOutput {
	return i.ToOtsBackupPlanOutputWithContext(context.Background())
}

func (i *OtsBackupPlan) ToOtsBackupPlanOutputWithContext(ctx context.Context) OtsBackupPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OtsBackupPlanOutput)
}

// OtsBackupPlanArrayInput is an input type that accepts OtsBackupPlanArray and OtsBackupPlanArrayOutput values.
// You can construct a concrete instance of `OtsBackupPlanArrayInput` via:
//
//          OtsBackupPlanArray{ OtsBackupPlanArgs{...} }
type OtsBackupPlanArrayInput interface {
	pulumi.Input

	ToOtsBackupPlanArrayOutput() OtsBackupPlanArrayOutput
	ToOtsBackupPlanArrayOutputWithContext(context.Context) OtsBackupPlanArrayOutput
}

type OtsBackupPlanArray []OtsBackupPlanInput

func (OtsBackupPlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OtsBackupPlan)(nil)).Elem()
}

func (i OtsBackupPlanArray) ToOtsBackupPlanArrayOutput() OtsBackupPlanArrayOutput {
	return i.ToOtsBackupPlanArrayOutputWithContext(context.Background())
}

func (i OtsBackupPlanArray) ToOtsBackupPlanArrayOutputWithContext(ctx context.Context) OtsBackupPlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OtsBackupPlanArrayOutput)
}

// OtsBackupPlanMapInput is an input type that accepts OtsBackupPlanMap and OtsBackupPlanMapOutput values.
// You can construct a concrete instance of `OtsBackupPlanMapInput` via:
//
//          OtsBackupPlanMap{ "key": OtsBackupPlanArgs{...} }
type OtsBackupPlanMapInput interface {
	pulumi.Input

	ToOtsBackupPlanMapOutput() OtsBackupPlanMapOutput
	ToOtsBackupPlanMapOutputWithContext(context.Context) OtsBackupPlanMapOutput
}

type OtsBackupPlanMap map[string]OtsBackupPlanInput

func (OtsBackupPlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OtsBackupPlan)(nil)).Elem()
}

func (i OtsBackupPlanMap) ToOtsBackupPlanMapOutput() OtsBackupPlanMapOutput {
	return i.ToOtsBackupPlanMapOutputWithContext(context.Background())
}

func (i OtsBackupPlanMap) ToOtsBackupPlanMapOutputWithContext(ctx context.Context) OtsBackupPlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OtsBackupPlanMapOutput)
}

type OtsBackupPlanOutput struct{ *pulumi.OutputState }

func (OtsBackupPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OtsBackupPlan)(nil)).Elem()
}

func (o OtsBackupPlanOutput) ToOtsBackupPlanOutput() OtsBackupPlanOutput {
	return o
}

func (o OtsBackupPlanOutput) ToOtsBackupPlanOutputWithContext(ctx context.Context) OtsBackupPlanOutput {
	return o
}

type OtsBackupPlanArrayOutput struct{ *pulumi.OutputState }

func (OtsBackupPlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OtsBackupPlan)(nil)).Elem()
}

func (o OtsBackupPlanArrayOutput) ToOtsBackupPlanArrayOutput() OtsBackupPlanArrayOutput {
	return o
}

func (o OtsBackupPlanArrayOutput) ToOtsBackupPlanArrayOutputWithContext(ctx context.Context) OtsBackupPlanArrayOutput {
	return o
}

func (o OtsBackupPlanArrayOutput) Index(i pulumi.IntInput) OtsBackupPlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OtsBackupPlan {
		return vs[0].([]*OtsBackupPlan)[vs[1].(int)]
	}).(OtsBackupPlanOutput)
}

type OtsBackupPlanMapOutput struct{ *pulumi.OutputState }

func (OtsBackupPlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OtsBackupPlan)(nil)).Elem()
}

func (o OtsBackupPlanMapOutput) ToOtsBackupPlanMapOutput() OtsBackupPlanMapOutput {
	return o
}

func (o OtsBackupPlanMapOutput) ToOtsBackupPlanMapOutputWithContext(ctx context.Context) OtsBackupPlanMapOutput {
	return o
}

func (o OtsBackupPlanMapOutput) MapIndex(k pulumi.StringInput) OtsBackupPlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OtsBackupPlan {
		return vs[0].(map[string]*OtsBackupPlan)[vs[1].(string)]
	}).(OtsBackupPlanOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OtsBackupPlanInput)(nil)).Elem(), &OtsBackupPlan{})
	pulumi.RegisterInputType(reflect.TypeOf((*OtsBackupPlanArrayInput)(nil)).Elem(), OtsBackupPlanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OtsBackupPlanMapInput)(nil)).Elem(), OtsBackupPlanMap{})
	pulumi.RegisterOutputType(OtsBackupPlanOutput{})
	pulumi.RegisterOutputType(OtsBackupPlanArrayOutput{})
	pulumi.RegisterOutputType(OtsBackupPlanMapOutput{})
}
