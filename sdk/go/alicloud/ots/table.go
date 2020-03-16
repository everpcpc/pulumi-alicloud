// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ots

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an OTS table resource.
//
// > **NOTE:** From Provider version 1.10.0, the provider field 'ots_instance_name' has been deprecated and
// you should use resource alicloud_ots_table's new field 'instance_name' and 'table_name' to re-import this resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/ots_table.html.markdown.
type Table struct {
	pulumi.CustomResourceState

	// The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
	DeviationCellVersionInSec pulumi.StringPtrOutput `pulumi:"deviationCellVersionInSec"`
	// The name of the OTS instance in which table will located.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The maximum number of versions stored in this table. The valid value is 1-2147483647.
	MaxVersion pulumi.IntOutput `pulumi:"maxVersion"`
	// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primaryKey` should not be less than one and not be more than four.
	PrimaryKeys TablePrimaryKeyArrayOutput `pulumi:"primaryKeys"`
	// The table name of the OTS instance. If changed, a new table would be created.
	TableName pulumi.StringOutput `pulumi:"tableName"`
	// The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
	TimeToLive pulumi.IntOutput `pulumi:"timeToLive"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil || args.InstanceName == nil {
		return nil, errors.New("missing required argument 'InstanceName'")
	}
	if args == nil || args.MaxVersion == nil {
		return nil, errors.New("missing required argument 'MaxVersion'")
	}
	if args == nil || args.PrimaryKeys == nil {
		return nil, errors.New("missing required argument 'PrimaryKeys'")
	}
	if args == nil || args.TableName == nil {
		return nil, errors.New("missing required argument 'TableName'")
	}
	if args == nil || args.TimeToLive == nil {
		return nil, errors.New("missing required argument 'TimeToLive'")
	}
	if args == nil {
		args = &TableArgs{}
	}
	var resource Table
	err := ctx.RegisterResource("alicloud:ots/table:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("alicloud:ots/table:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
	// The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
	DeviationCellVersionInSec *string `pulumi:"deviationCellVersionInSec"`
	// The name of the OTS instance in which table will located.
	InstanceName *string `pulumi:"instanceName"`
	// The maximum number of versions stored in this table. The valid value is 1-2147483647.
	MaxVersion *int `pulumi:"maxVersion"`
	// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primaryKey` should not be less than one and not be more than four.
	PrimaryKeys []TablePrimaryKey `pulumi:"primaryKeys"`
	// The table name of the OTS instance. If changed, a new table would be created.
	TableName *string `pulumi:"tableName"`
	// The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
	TimeToLive *int `pulumi:"timeToLive"`
}

type TableState struct {
	// The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
	DeviationCellVersionInSec pulumi.StringPtrInput
	// The name of the OTS instance in which table will located.
	InstanceName pulumi.StringPtrInput
	// The maximum number of versions stored in this table. The valid value is 1-2147483647.
	MaxVersion pulumi.IntPtrInput
	// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primaryKey` should not be less than one and not be more than four.
	PrimaryKeys TablePrimaryKeyArrayInput
	// The table name of the OTS instance. If changed, a new table would be created.
	TableName pulumi.StringPtrInput
	// The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
	TimeToLive pulumi.IntPtrInput
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	// The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
	DeviationCellVersionInSec *string `pulumi:"deviationCellVersionInSec"`
	// The name of the OTS instance in which table will located.
	InstanceName string `pulumi:"instanceName"`
	// The maximum number of versions stored in this table. The valid value is 1-2147483647.
	MaxVersion int `pulumi:"maxVersion"`
	// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primaryKey` should not be less than one and not be more than four.
	PrimaryKeys []TablePrimaryKey `pulumi:"primaryKeys"`
	// The table name of the OTS instance. If changed, a new table would be created.
	TableName string `pulumi:"tableName"`
	// The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
	TimeToLive int `pulumi:"timeToLive"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// The max version offset of the table. The valid value is 1-9223372036854775807. Defaults to 86400.
	DeviationCellVersionInSec pulumi.StringPtrInput
	// The name of the OTS instance in which table will located.
	InstanceName pulumi.StringInput
	// The maximum number of versions stored in this table. The valid value is 1-2147483647.
	MaxVersion pulumi.IntInput
	// The property of `TableMeta` which indicates the structure information of a table. It describes the attribute value of primary key. The number of `primaryKey` should not be less than one and not be more than four.
	PrimaryKeys TablePrimaryKeyArrayInput
	// The table name of the OTS instance. If changed, a new table would be created.
	TableName pulumi.StringInput
	// The retention time of data stored in this table (unit: second). The value maximum is 2147483647 and -1 means never expired.
	TimeToLive pulumi.IntInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

