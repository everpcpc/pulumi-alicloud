// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ots

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an OTS search index resource.
//
// For information about OTS search index and how to use it, see [Search index overview](https://www.alibabacloud.com/help/en/tablestore/latest/search-index-overview).
//
// > **NOTE:** Available in v1.187.0+.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ots"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "terraformtest"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			instance1, err := ots.NewInstance(ctx, "instance1", &ots.InstanceArgs{
//				Description: pulumi.String(name),
//				AccessedBy:  pulumi.String("Any"),
//				Tags: pulumi.AnyMap{
//					"Created": pulumi.Any("TF"),
//					"For":     pulumi.Any("acceptance test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			table1, err := ots.NewTable(ctx, "table1", &ots.TableArgs{
//				InstanceName: instance1.Name,
//				TableName:    pulumi.String(name),
//				PrimaryKeys: ots.TablePrimaryKeyArray{
//					&ots.TablePrimaryKeyArgs{
//						Name: pulumi.String("pk1"),
//						Type: pulumi.String("Integer"),
//					},
//					&ots.TablePrimaryKeyArgs{
//						Name: pulumi.String("pk2"),
//						Type: pulumi.String("String"),
//					},
//				},
//				DefinedColumns: ots.TableDefinedColumnArray{
//					&ots.TableDefinedColumnArgs{
//						Name: pulumi.String("col1"),
//						Type: pulumi.String("String"),
//					},
//					&ots.TableDefinedColumnArgs{
//						Name: pulumi.String("col2"),
//						Type: pulumi.String("Integer"),
//					},
//				},
//				TimeToLive:                -1,
//				MaxVersion:                pulumi.Int(1),
//				DeviationCellVersionInSec: pulumi.String("1"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ots.NewSearchIndex(ctx, "default", &ots.SearchIndexArgs{
//				InstanceName: instance1.Name,
//				TableName:    table1.TableName,
//				IndexName:    pulumi.String(name),
//				TimeToLive:   -1,
//				Schemas: ots.SearchIndexSchemaArray{
//					&ots.SearchIndexSchemaArgs{
//						FieldSchemas: ots.SearchIndexSchemaFieldSchemaArray{
//							&ots.SearchIndexSchemaFieldSchemaArgs{
//								FieldName: pulumi.String("col1"),
//								FieldType: pulumi.String("Text"),
//								IsArray:   pulumi.Bool(false),
//								Index:     pulumi.Bool(true),
//								Analyzer:  pulumi.String("Split"),
//								Store:     pulumi.Bool(true),
//							},
//							&ots.SearchIndexSchemaFieldSchemaArgs{
//								FieldName:        pulumi.String("col2"),
//								FieldType:        pulumi.String("Long"),
//								EnableSortAndAgg: pulumi.Bool(true),
//							},
//							&ots.SearchIndexSchemaFieldSchemaArgs{
//								FieldName: pulumi.String("pk1"),
//								FieldType: pulumi.String("Long"),
//							},
//							&ots.SearchIndexSchemaFieldSchemaArgs{
//								FieldName: pulumi.String("pk2"),
//								FieldType: pulumi.String("Text"),
//							},
//						},
//						IndexSettings: ots.SearchIndexSchemaIndexSettingArray{
//							&ots.SearchIndexSchemaIndexSettingArgs{
//								RoutingFields: pulumi.StringArray{
//									pulumi.String("pk1"),
//									pulumi.String("pk2"),
//								},
//							},
//						},
//						IndexSorts: ots.SearchIndexSchemaIndexSortArray{
//							&ots.SearchIndexSchemaIndexSortArgs{
//								Sorters: ots.SearchIndexSchemaIndexSortSorterArray{
//									&ots.SearchIndexSchemaIndexSortSorterArgs{
//										SorterType: pulumi.String("PrimaryKeySort"),
//										Order:      pulumi.String("Asc"),
//									},
//									&ots.SearchIndexSchemaIndexSortSorterArgs{
//										SorterType: pulumi.String("FieldSort"),
//										Order:      pulumi.String("Desc"),
//										FieldName:  pulumi.String("col2"),
//										Mode:       pulumi.String("Max"),
//									},
//								},
//							},
//						},
//					},
//				},
//			})
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
// OTS search index can be imported using id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ots/searchIndex:SearchIndex index1 <instance_name>:<table_name>:<index_name>:<index_type>
//
// ```
type SearchIndex struct {
	pulumi.CustomResourceState

	// The search index create time.
	CreateTime pulumi.IntOutput `pulumi:"createTime"`
	// The timestamp for sync phase.
	CurrentSyncTimestamp pulumi.IntOutput `pulumi:"currentSyncTimestamp"`
	// The index id of the search index which could not be changed.
	IndexId pulumi.StringOutput `pulumi:"indexId"`
	// The index name of the OTS Table. If changed, a new index would be created.
	IndexName pulumi.StringOutput `pulumi:"indexName"`
	// The name of the OTS instance in which table will located.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The schema of the search index. If changed, a new index would be created.
	Schemas SearchIndexSchemaArrayOutput `pulumi:"schemas"`
	// The search index sync phase. possible values: `Full`, `Incr`.
	SyncPhase pulumi.StringOutput `pulumi:"syncPhase"`
	// The name of the OTS table. If changed, a new table would be created.
	TableName pulumi.StringOutput `pulumi:"tableName"`
	// The index type of the OTS Table. Specifies the retention period of data in the search index. Unit: seconds. Default value: -1.
	// If the retention period exceeds the TTL value, OTS automatically deletes expired data.
	TimeToLive pulumi.IntPtrOutput `pulumi:"timeToLive"`
}

// NewSearchIndex registers a new resource with the given unique name, arguments, and options.
func NewSearchIndex(ctx *pulumi.Context,
	name string, args *SearchIndexArgs, opts ...pulumi.ResourceOption) (*SearchIndex, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IndexName == nil {
		return nil, errors.New("invalid value for required argument 'IndexName'")
	}
	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.Schemas == nil {
		return nil, errors.New("invalid value for required argument 'Schemas'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	var resource SearchIndex
	err := ctx.RegisterResource("alicloud:ots/searchIndex:SearchIndex", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSearchIndex gets an existing SearchIndex resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSearchIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SearchIndexState, opts ...pulumi.ResourceOption) (*SearchIndex, error) {
	var resource SearchIndex
	err := ctx.ReadResource("alicloud:ots/searchIndex:SearchIndex", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SearchIndex resources.
type searchIndexState struct {
	// The search index create time.
	CreateTime *int `pulumi:"createTime"`
	// The timestamp for sync phase.
	CurrentSyncTimestamp *int `pulumi:"currentSyncTimestamp"`
	// The index id of the search index which could not be changed.
	IndexId *string `pulumi:"indexId"`
	// The index name of the OTS Table. If changed, a new index would be created.
	IndexName *string `pulumi:"indexName"`
	// The name of the OTS instance in which table will located.
	InstanceName *string `pulumi:"instanceName"`
	// The schema of the search index. If changed, a new index would be created.
	Schemas []SearchIndexSchema `pulumi:"schemas"`
	// The search index sync phase. possible values: `Full`, `Incr`.
	SyncPhase *string `pulumi:"syncPhase"`
	// The name of the OTS table. If changed, a new table would be created.
	TableName *string `pulumi:"tableName"`
	// The index type of the OTS Table. Specifies the retention period of data in the search index. Unit: seconds. Default value: -1.
	// If the retention period exceeds the TTL value, OTS automatically deletes expired data.
	TimeToLive *int `pulumi:"timeToLive"`
}

type SearchIndexState struct {
	// The search index create time.
	CreateTime pulumi.IntPtrInput
	// The timestamp for sync phase.
	CurrentSyncTimestamp pulumi.IntPtrInput
	// The index id of the search index which could not be changed.
	IndexId pulumi.StringPtrInput
	// The index name of the OTS Table. If changed, a new index would be created.
	IndexName pulumi.StringPtrInput
	// The name of the OTS instance in which table will located.
	InstanceName pulumi.StringPtrInput
	// The schema of the search index. If changed, a new index would be created.
	Schemas SearchIndexSchemaArrayInput
	// The search index sync phase. possible values: `Full`, `Incr`.
	SyncPhase pulumi.StringPtrInput
	// The name of the OTS table. If changed, a new table would be created.
	TableName pulumi.StringPtrInput
	// The index type of the OTS Table. Specifies the retention period of data in the search index. Unit: seconds. Default value: -1.
	// If the retention period exceeds the TTL value, OTS automatically deletes expired data.
	TimeToLive pulumi.IntPtrInput
}

func (SearchIndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*searchIndexState)(nil)).Elem()
}

type searchIndexArgs struct {
	// The index name of the OTS Table. If changed, a new index would be created.
	IndexName string `pulumi:"indexName"`
	// The name of the OTS instance in which table will located.
	InstanceName string `pulumi:"instanceName"`
	// The schema of the search index. If changed, a new index would be created.
	Schemas []SearchIndexSchema `pulumi:"schemas"`
	// The name of the OTS table. If changed, a new table would be created.
	TableName string `pulumi:"tableName"`
	// The index type of the OTS Table. Specifies the retention period of data in the search index. Unit: seconds. Default value: -1.
	// If the retention period exceeds the TTL value, OTS automatically deletes expired data.
	TimeToLive *int `pulumi:"timeToLive"`
}

// The set of arguments for constructing a SearchIndex resource.
type SearchIndexArgs struct {
	// The index name of the OTS Table. If changed, a new index would be created.
	IndexName pulumi.StringInput
	// The name of the OTS instance in which table will located.
	InstanceName pulumi.StringInput
	// The schema of the search index. If changed, a new index would be created.
	Schemas SearchIndexSchemaArrayInput
	// The name of the OTS table. If changed, a new table would be created.
	TableName pulumi.StringInput
	// The index type of the OTS Table. Specifies the retention period of data in the search index. Unit: seconds. Default value: -1.
	// If the retention period exceeds the TTL value, OTS automatically deletes expired data.
	TimeToLive pulumi.IntPtrInput
}

func (SearchIndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*searchIndexArgs)(nil)).Elem()
}

type SearchIndexInput interface {
	pulumi.Input

	ToSearchIndexOutput() SearchIndexOutput
	ToSearchIndexOutputWithContext(ctx context.Context) SearchIndexOutput
}

func (*SearchIndex) ElementType() reflect.Type {
	return reflect.TypeOf((**SearchIndex)(nil)).Elem()
}

func (i *SearchIndex) ToSearchIndexOutput() SearchIndexOutput {
	return i.ToSearchIndexOutputWithContext(context.Background())
}

func (i *SearchIndex) ToSearchIndexOutputWithContext(ctx context.Context) SearchIndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SearchIndexOutput)
}

// SearchIndexArrayInput is an input type that accepts SearchIndexArray and SearchIndexArrayOutput values.
// You can construct a concrete instance of `SearchIndexArrayInput` via:
//
//	SearchIndexArray{ SearchIndexArgs{...} }
type SearchIndexArrayInput interface {
	pulumi.Input

	ToSearchIndexArrayOutput() SearchIndexArrayOutput
	ToSearchIndexArrayOutputWithContext(context.Context) SearchIndexArrayOutput
}

type SearchIndexArray []SearchIndexInput

func (SearchIndexArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SearchIndex)(nil)).Elem()
}

func (i SearchIndexArray) ToSearchIndexArrayOutput() SearchIndexArrayOutput {
	return i.ToSearchIndexArrayOutputWithContext(context.Background())
}

func (i SearchIndexArray) ToSearchIndexArrayOutputWithContext(ctx context.Context) SearchIndexArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SearchIndexArrayOutput)
}

// SearchIndexMapInput is an input type that accepts SearchIndexMap and SearchIndexMapOutput values.
// You can construct a concrete instance of `SearchIndexMapInput` via:
//
//	SearchIndexMap{ "key": SearchIndexArgs{...} }
type SearchIndexMapInput interface {
	pulumi.Input

	ToSearchIndexMapOutput() SearchIndexMapOutput
	ToSearchIndexMapOutputWithContext(context.Context) SearchIndexMapOutput
}

type SearchIndexMap map[string]SearchIndexInput

func (SearchIndexMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SearchIndex)(nil)).Elem()
}

func (i SearchIndexMap) ToSearchIndexMapOutput() SearchIndexMapOutput {
	return i.ToSearchIndexMapOutputWithContext(context.Background())
}

func (i SearchIndexMap) ToSearchIndexMapOutputWithContext(ctx context.Context) SearchIndexMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SearchIndexMapOutput)
}

type SearchIndexOutput struct{ *pulumi.OutputState }

func (SearchIndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SearchIndex)(nil)).Elem()
}

func (o SearchIndexOutput) ToSearchIndexOutput() SearchIndexOutput {
	return o
}

func (o SearchIndexOutput) ToSearchIndexOutputWithContext(ctx context.Context) SearchIndexOutput {
	return o
}

// The search index create time.
func (o SearchIndexOutput) CreateTime() pulumi.IntOutput {
	return o.ApplyT(func(v *SearchIndex) pulumi.IntOutput { return v.CreateTime }).(pulumi.IntOutput)
}

// The timestamp for sync phase.
func (o SearchIndexOutput) CurrentSyncTimestamp() pulumi.IntOutput {
	return o.ApplyT(func(v *SearchIndex) pulumi.IntOutput { return v.CurrentSyncTimestamp }).(pulumi.IntOutput)
}

// The index id of the search index which could not be changed.
func (o SearchIndexOutput) IndexId() pulumi.StringOutput {
	return o.ApplyT(func(v *SearchIndex) pulumi.StringOutput { return v.IndexId }).(pulumi.StringOutput)
}

// The index name of the OTS Table. If changed, a new index would be created.
func (o SearchIndexOutput) IndexName() pulumi.StringOutput {
	return o.ApplyT(func(v *SearchIndex) pulumi.StringOutput { return v.IndexName }).(pulumi.StringOutput)
}

// The name of the OTS instance in which table will located.
func (o SearchIndexOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *SearchIndex) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// The schema of the search index. If changed, a new index would be created.
func (o SearchIndexOutput) Schemas() SearchIndexSchemaArrayOutput {
	return o.ApplyT(func(v *SearchIndex) SearchIndexSchemaArrayOutput { return v.Schemas }).(SearchIndexSchemaArrayOutput)
}

// The search index sync phase. possible values: `Full`, `Incr`.
func (o SearchIndexOutput) SyncPhase() pulumi.StringOutput {
	return o.ApplyT(func(v *SearchIndex) pulumi.StringOutput { return v.SyncPhase }).(pulumi.StringOutput)
}

// The name of the OTS table. If changed, a new table would be created.
func (o SearchIndexOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *SearchIndex) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

// The index type of the OTS Table. Specifies the retention period of data in the search index. Unit: seconds. Default value: -1.
// If the retention period exceeds the TTL value, OTS automatically deletes expired data.
func (o SearchIndexOutput) TimeToLive() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SearchIndex) pulumi.IntPtrOutput { return v.TimeToLive }).(pulumi.IntPtrOutput)
}

type SearchIndexArrayOutput struct{ *pulumi.OutputState }

func (SearchIndexArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SearchIndex)(nil)).Elem()
}

func (o SearchIndexArrayOutput) ToSearchIndexArrayOutput() SearchIndexArrayOutput {
	return o
}

func (o SearchIndexArrayOutput) ToSearchIndexArrayOutputWithContext(ctx context.Context) SearchIndexArrayOutput {
	return o
}

func (o SearchIndexArrayOutput) Index(i pulumi.IntInput) SearchIndexOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SearchIndex {
		return vs[0].([]*SearchIndex)[vs[1].(int)]
	}).(SearchIndexOutput)
}

type SearchIndexMapOutput struct{ *pulumi.OutputState }

func (SearchIndexMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SearchIndex)(nil)).Elem()
}

func (o SearchIndexMapOutput) ToSearchIndexMapOutput() SearchIndexMapOutput {
	return o
}

func (o SearchIndexMapOutput) ToSearchIndexMapOutputWithContext(ctx context.Context) SearchIndexMapOutput {
	return o
}

func (o SearchIndexMapOutput) MapIndex(k pulumi.StringInput) SearchIndexOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SearchIndex {
		return vs[0].(map[string]*SearchIndex)[vs[1].(string)]
	}).(SearchIndexOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SearchIndexInput)(nil)).Elem(), &SearchIndex{})
	pulumi.RegisterInputType(reflect.TypeOf((*SearchIndexArrayInput)(nil)).Elem(), SearchIndexArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SearchIndexMapInput)(nil)).Elem(), SearchIndexMap{})
	pulumi.RegisterOutputType(SearchIndexOutput{})
	pulumi.RegisterOutputType(SearchIndexArrayOutput{})
	pulumi.RegisterOutputType(SearchIndexMapOutput{})
}
