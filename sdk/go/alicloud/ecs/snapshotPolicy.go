// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ecs

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an ECS snapshot policy resource.
//
// For information about snapshot policy and how to use it, see [Snapshot](https://www.alibabacloud.com/help/doc-detail/25460.html).
//
// > **NOTE:** Available in 1.42.0+.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/snapshot_policy.html.markdown.
type SnapshotPolicy struct {
	pulumi.CustomResourceState

	// The snapshot policy name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1  indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array.
	// - A maximum of seven time points can be selected.
	// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
	RepeatWeekdays pulumi.StringArrayOutput `pulumi:"repeatWeekdays"`
	// The snapshot retention time, and the unit of measurement is day. Optional values:
	// - -1: The automatic snapshots are retained permanently.
	// - [1, 65536]: The number of days retained.
	RetentionDays pulumi.IntOutput `pulumi:"retentionDays"`
	// The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00,  for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array.
	// - A maximum of 24 time points can be selected.
	// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
	TimePoints pulumi.StringArrayOutput `pulumi:"timePoints"`
}

// NewSnapshotPolicy registers a new resource with the given unique name, arguments, and options.
func NewSnapshotPolicy(ctx *pulumi.Context,
	name string, args *SnapshotPolicyArgs, opts ...pulumi.ResourceOption) (*SnapshotPolicy, error) {
	if args == nil || args.RepeatWeekdays == nil {
		return nil, errors.New("missing required argument 'RepeatWeekdays'")
	}
	if args == nil || args.RetentionDays == nil {
		return nil, errors.New("missing required argument 'RetentionDays'")
	}
	if args == nil || args.TimePoints == nil {
		return nil, errors.New("missing required argument 'TimePoints'")
	}
	if args == nil {
		args = &SnapshotPolicyArgs{}
	}
	var resource SnapshotPolicy
	err := ctx.RegisterResource("alicloud:ecs/snapshotPolicy:SnapshotPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotPolicy gets an existing SnapshotPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotPolicyState, opts ...pulumi.ResourceOption) (*SnapshotPolicy, error) {
	var resource SnapshotPolicy
	err := ctx.ReadResource("alicloud:ecs/snapshotPolicy:SnapshotPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotPolicy resources.
type snapshotPolicyState struct {
	// The snapshot policy name.
	Name *string `pulumi:"name"`
	// The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1  indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array.
	// - A maximum of seven time points can be selected.
	// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
	RepeatWeekdays []string `pulumi:"repeatWeekdays"`
	// The snapshot retention time, and the unit of measurement is day. Optional values:
	// - -1: The automatic snapshots are retained permanently.
	// - [1, 65536]: The number of days retained.
	RetentionDays *int `pulumi:"retentionDays"`
	// The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00,  for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array.
	// - A maximum of 24 time points can be selected.
	// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
	TimePoints []string `pulumi:"timePoints"`
}

type SnapshotPolicyState struct {
	// The snapshot policy name.
	Name pulumi.StringPtrInput
	// The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1  indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array.
	// - A maximum of seven time points can be selected.
	// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
	RepeatWeekdays pulumi.StringArrayInput
	// The snapshot retention time, and the unit of measurement is day. Optional values:
	// - -1: The automatic snapshots are retained permanently.
	// - [1, 65536]: The number of days retained.
	RetentionDays pulumi.IntPtrInput
	// The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00,  for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array.
	// - A maximum of 24 time points can be selected.
	// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
	TimePoints pulumi.StringArrayInput
}

func (SnapshotPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotPolicyState)(nil)).Elem()
}

type snapshotPolicyArgs struct {
	// The snapshot policy name.
	Name *string `pulumi:"name"`
	// The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1  indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array.
	// - A maximum of seven time points can be selected.
	// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
	RepeatWeekdays []string `pulumi:"repeatWeekdays"`
	// The snapshot retention time, and the unit of measurement is day. Optional values:
	// - -1: The automatic snapshots are retained permanently.
	// - [1, 65536]: The number of days retained.
	RetentionDays int `pulumi:"retentionDays"`
	// The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00,  for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array.
	// - A maximum of 24 time points can be selected.
	// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
	TimePoints []string `pulumi:"timePoints"`
}

// The set of arguments for constructing a SnapshotPolicy resource.
type SnapshotPolicyArgs struct {
	// The snapshot policy name.
	Name pulumi.StringPtrInput
	// The automatic snapshot repetition dates. The unit of measurement is day and the repeating cycle is a week. Value range: [1, 7], which represents days starting from Monday to Sunday, for example 1  indicates Monday. When you want to schedule multiple automatic snapshot tasks for a disk in a week, you can set the RepeatWeekdays to an array.
	// - A maximum of seven time points can be selected.
	// - The format is  an JSON array of ["1", "2", … "7"]  and the time points are separated by commas (,).
	RepeatWeekdays pulumi.StringArrayInput
	// The snapshot retention time, and the unit of measurement is day. Optional values:
	// - -1: The automatic snapshots are retained permanently.
	// - [1, 65536]: The number of days retained.
	RetentionDays pulumi.IntInput
	// The automatic snapshot creation schedule, and the unit of measurement is hour. Value range: [0, 23], which represents from 00:00 to 24:00,  for example 1 indicates 01:00. When you want to schedule multiple automatic snapshot tasks for a disk in a day, you can set the TimePoints to an array.
	// - A maximum of 24 time points can be selected.
	// - The format is  an JSON array of ["0", "1", … "23"] and the time points are separated by commas (,).
	TimePoints pulumi.StringArrayInput
}

func (SnapshotPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotPolicyArgs)(nil)).Elem()
}

