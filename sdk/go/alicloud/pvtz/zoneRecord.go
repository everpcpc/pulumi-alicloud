// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pvtz

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Private Zone Record can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:pvtz/zoneRecord:ZoneRecord example abc123456
// ```
type ZoneRecord struct {
	pulumi.CustomResourceState

	// User language.
	Lang pulumi.StringPtrOutput `pulumi:"lang"`
	// The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-99]. Default to 1.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The Private Zone Record ID.
	RecordId pulumi.StringOutput `pulumi:"recordId"`
	// The remark of the Private Zone Record.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// The resource record of the Private Zone Record.
	//
	// Deprecated: Field 'resource_record' has been deprecated from version 1.109.0. Use 'rr' instead.
	ResourceRecord pulumi.StringOutput `pulumi:"resourceRecord"`
	// The rr of the Private Zone Record.
	Rr pulumi.StringOutput `pulumi:"rr"`
	// Resolve record status. Value:
	// - ENABLE: enable resolution.
	// - DISABLE: pause parsing.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// The ttl of the Private Zone Record. Default to `60`.
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
	// The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
	Type         pulumi.StringOutput    `pulumi:"type"`
	UserClientIp pulumi.StringPtrOutput `pulumi:"userClientIp"`
	// The value of the Private Zone Record.
	Value pulumi.StringOutput `pulumi:"value"`
	// The name of the Private Zone Record.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewZoneRecord registers a new resource with the given unique name, arguments, and options.
func NewZoneRecord(ctx *pulumi.Context,
	name string, args *ZoneRecordArgs, opts ...pulumi.ResourceOption) (*ZoneRecord, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	var resource ZoneRecord
	err := ctx.RegisterResource("alicloud:pvtz/zoneRecord:ZoneRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZoneRecord gets an existing ZoneRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZoneRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneRecordState, opts ...pulumi.ResourceOption) (*ZoneRecord, error) {
	var resource ZoneRecord
	err := ctx.ReadResource("alicloud:pvtz/zoneRecord:ZoneRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZoneRecord resources.
type zoneRecordState struct {
	// User language.
	Lang *string `pulumi:"lang"`
	// The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-99]. Default to 1.
	Priority *int `pulumi:"priority"`
	// The Private Zone Record ID.
	RecordId *string `pulumi:"recordId"`
	// The remark of the Private Zone Record.
	Remark *string `pulumi:"remark"`
	// The resource record of the Private Zone Record.
	//
	// Deprecated: Field 'resource_record' has been deprecated from version 1.109.0. Use 'rr' instead.
	ResourceRecord *string `pulumi:"resourceRecord"`
	// The rr of the Private Zone Record.
	Rr *string `pulumi:"rr"`
	// Resolve record status. Value:
	// - ENABLE: enable resolution.
	// - DISABLE: pause parsing.
	Status *string `pulumi:"status"`
	// The ttl of the Private Zone Record. Default to `60`.
	Ttl *int `pulumi:"ttl"`
	// The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
	Type         *string `pulumi:"type"`
	UserClientIp *string `pulumi:"userClientIp"`
	// The value of the Private Zone Record.
	Value *string `pulumi:"value"`
	// The name of the Private Zone Record.
	ZoneId *string `pulumi:"zoneId"`
}

type ZoneRecordState struct {
	// User language.
	Lang pulumi.StringPtrInput
	// The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-99]. Default to 1.
	Priority pulumi.IntPtrInput
	// The Private Zone Record ID.
	RecordId pulumi.StringPtrInput
	// The remark of the Private Zone Record.
	Remark pulumi.StringPtrInput
	// The resource record of the Private Zone Record.
	//
	// Deprecated: Field 'resource_record' has been deprecated from version 1.109.0. Use 'rr' instead.
	ResourceRecord pulumi.StringPtrInput
	// The rr of the Private Zone Record.
	Rr pulumi.StringPtrInput
	// Resolve record status. Value:
	// - ENABLE: enable resolution.
	// - DISABLE: pause parsing.
	Status pulumi.StringPtrInput
	// The ttl of the Private Zone Record. Default to `60`.
	Ttl pulumi.IntPtrInput
	// The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
	Type         pulumi.StringPtrInput
	UserClientIp pulumi.StringPtrInput
	// The value of the Private Zone Record.
	Value pulumi.StringPtrInput
	// The name of the Private Zone Record.
	ZoneId pulumi.StringPtrInput
}

func (ZoneRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneRecordState)(nil)).Elem()
}

type zoneRecordArgs struct {
	// User language.
	Lang *string `pulumi:"lang"`
	// The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-99]. Default to 1.
	Priority *int `pulumi:"priority"`
	// The remark of the Private Zone Record.
	Remark *string `pulumi:"remark"`
	// The resource record of the Private Zone Record.
	//
	// Deprecated: Field 'resource_record' has been deprecated from version 1.109.0. Use 'rr' instead.
	ResourceRecord *string `pulumi:"resourceRecord"`
	// The rr of the Private Zone Record.
	Rr *string `pulumi:"rr"`
	// Resolve record status. Value:
	// - ENABLE: enable resolution.
	// - DISABLE: pause parsing.
	Status *string `pulumi:"status"`
	// The ttl of the Private Zone Record. Default to `60`.
	Ttl *int `pulumi:"ttl"`
	// The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
	Type         string  `pulumi:"type"`
	UserClientIp *string `pulumi:"userClientIp"`
	// The value of the Private Zone Record.
	Value string `pulumi:"value"`
	// The name of the Private Zone Record.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a ZoneRecord resource.
type ZoneRecordArgs struct {
	// User language.
	Lang pulumi.StringPtrInput
	// The priority of the Private Zone Record. At present, only can "MX" record support it. Valid values: [1-99]. Default to 1.
	Priority pulumi.IntPtrInput
	// The remark of the Private Zone Record.
	Remark pulumi.StringPtrInput
	// The resource record of the Private Zone Record.
	//
	// Deprecated: Field 'resource_record' has been deprecated from version 1.109.0. Use 'rr' instead.
	ResourceRecord pulumi.StringPtrInput
	// The rr of the Private Zone Record.
	Rr pulumi.StringPtrInput
	// Resolve record status. Value:
	// - ENABLE: enable resolution.
	// - DISABLE: pause parsing.
	Status pulumi.StringPtrInput
	// The ttl of the Private Zone Record. Default to `60`.
	Ttl pulumi.IntPtrInput
	// The type of the Private Zone Record. Valid values: A, CNAME, TXT, MX, PTR, SRV.
	Type         pulumi.StringInput
	UserClientIp pulumi.StringPtrInput
	// The value of the Private Zone Record.
	Value pulumi.StringInput
	// The name of the Private Zone Record.
	ZoneId pulumi.StringInput
}

func (ZoneRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneRecordArgs)(nil)).Elem()
}

type ZoneRecordInput interface {
	pulumi.Input

	ToZoneRecordOutput() ZoneRecordOutput
	ToZoneRecordOutputWithContext(ctx context.Context) ZoneRecordOutput
}

func (*ZoneRecord) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneRecord)(nil))
}

func (i *ZoneRecord) ToZoneRecordOutput() ZoneRecordOutput {
	return i.ToZoneRecordOutputWithContext(context.Background())
}

func (i *ZoneRecord) ToZoneRecordOutputWithContext(ctx context.Context) ZoneRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneRecordOutput)
}

func (i *ZoneRecord) ToZoneRecordPtrOutput() ZoneRecordPtrOutput {
	return i.ToZoneRecordPtrOutputWithContext(context.Background())
}

func (i *ZoneRecord) ToZoneRecordPtrOutputWithContext(ctx context.Context) ZoneRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneRecordPtrOutput)
}

type ZoneRecordPtrInput interface {
	pulumi.Input

	ToZoneRecordPtrOutput() ZoneRecordPtrOutput
	ToZoneRecordPtrOutputWithContext(ctx context.Context) ZoneRecordPtrOutput
}

type zoneRecordPtrType ZoneRecordArgs

func (*zoneRecordPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneRecord)(nil))
}

func (i *zoneRecordPtrType) ToZoneRecordPtrOutput() ZoneRecordPtrOutput {
	return i.ToZoneRecordPtrOutputWithContext(context.Background())
}

func (i *zoneRecordPtrType) ToZoneRecordPtrOutputWithContext(ctx context.Context) ZoneRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneRecordPtrOutput)
}

// ZoneRecordArrayInput is an input type that accepts ZoneRecordArray and ZoneRecordArrayOutput values.
// You can construct a concrete instance of `ZoneRecordArrayInput` via:
//
//          ZoneRecordArray{ ZoneRecordArgs{...} }
type ZoneRecordArrayInput interface {
	pulumi.Input

	ToZoneRecordArrayOutput() ZoneRecordArrayOutput
	ToZoneRecordArrayOutputWithContext(context.Context) ZoneRecordArrayOutput
}

type ZoneRecordArray []ZoneRecordInput

func (ZoneRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZoneRecord)(nil)).Elem()
}

func (i ZoneRecordArray) ToZoneRecordArrayOutput() ZoneRecordArrayOutput {
	return i.ToZoneRecordArrayOutputWithContext(context.Background())
}

func (i ZoneRecordArray) ToZoneRecordArrayOutputWithContext(ctx context.Context) ZoneRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneRecordArrayOutput)
}

// ZoneRecordMapInput is an input type that accepts ZoneRecordMap and ZoneRecordMapOutput values.
// You can construct a concrete instance of `ZoneRecordMapInput` via:
//
//          ZoneRecordMap{ "key": ZoneRecordArgs{...} }
type ZoneRecordMapInput interface {
	pulumi.Input

	ToZoneRecordMapOutput() ZoneRecordMapOutput
	ToZoneRecordMapOutputWithContext(context.Context) ZoneRecordMapOutput
}

type ZoneRecordMap map[string]ZoneRecordInput

func (ZoneRecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZoneRecord)(nil)).Elem()
}

func (i ZoneRecordMap) ToZoneRecordMapOutput() ZoneRecordMapOutput {
	return i.ToZoneRecordMapOutputWithContext(context.Background())
}

func (i ZoneRecordMap) ToZoneRecordMapOutputWithContext(ctx context.Context) ZoneRecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneRecordMapOutput)
}

type ZoneRecordOutput struct{ *pulumi.OutputState }

func (ZoneRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneRecord)(nil))
}

func (o ZoneRecordOutput) ToZoneRecordOutput() ZoneRecordOutput {
	return o
}

func (o ZoneRecordOutput) ToZoneRecordOutputWithContext(ctx context.Context) ZoneRecordOutput {
	return o
}

func (o ZoneRecordOutput) ToZoneRecordPtrOutput() ZoneRecordPtrOutput {
	return o.ToZoneRecordPtrOutputWithContext(context.Background())
}

func (o ZoneRecordOutput) ToZoneRecordPtrOutputWithContext(ctx context.Context) ZoneRecordPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ZoneRecord) *ZoneRecord {
		return &v
	}).(ZoneRecordPtrOutput)
}

type ZoneRecordPtrOutput struct{ *pulumi.OutputState }

func (ZoneRecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneRecord)(nil))
}

func (o ZoneRecordPtrOutput) ToZoneRecordPtrOutput() ZoneRecordPtrOutput {
	return o
}

func (o ZoneRecordPtrOutput) ToZoneRecordPtrOutputWithContext(ctx context.Context) ZoneRecordPtrOutput {
	return o
}

func (o ZoneRecordPtrOutput) Elem() ZoneRecordOutput {
	return o.ApplyT(func(v *ZoneRecord) ZoneRecord {
		if v != nil {
			return *v
		}
		var ret ZoneRecord
		return ret
	}).(ZoneRecordOutput)
}

type ZoneRecordArrayOutput struct{ *pulumi.OutputState }

func (ZoneRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ZoneRecord)(nil))
}

func (o ZoneRecordArrayOutput) ToZoneRecordArrayOutput() ZoneRecordArrayOutput {
	return o
}

func (o ZoneRecordArrayOutput) ToZoneRecordArrayOutputWithContext(ctx context.Context) ZoneRecordArrayOutput {
	return o
}

func (o ZoneRecordArrayOutput) Index(i pulumi.IntInput) ZoneRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ZoneRecord {
		return vs[0].([]ZoneRecord)[vs[1].(int)]
	}).(ZoneRecordOutput)
}

type ZoneRecordMapOutput struct{ *pulumi.OutputState }

func (ZoneRecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ZoneRecord)(nil))
}

func (o ZoneRecordMapOutput) ToZoneRecordMapOutput() ZoneRecordMapOutput {
	return o
}

func (o ZoneRecordMapOutput) ToZoneRecordMapOutputWithContext(ctx context.Context) ZoneRecordMapOutput {
	return o
}

func (o ZoneRecordMapOutput) MapIndex(k pulumi.StringInput) ZoneRecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ZoneRecord {
		return vs[0].(map[string]ZoneRecord)[vs[1].(string)]
	}).(ZoneRecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneRecordInput)(nil)).Elem(), &ZoneRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneRecordPtrInput)(nil)).Elem(), &ZoneRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneRecordArrayInput)(nil)).Elem(), ZoneRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneRecordMapInput)(nil)).Elem(), ZoneRecordMap{})
	pulumi.RegisterOutputType(ZoneRecordOutput{})
	pulumi.RegisterOutputType(ZoneRecordPtrOutput{})
	pulumi.RegisterOutputType(ZoneRecordArrayOutput{})
	pulumi.RegisterOutputType(ZoneRecordMapOutput{})
}
