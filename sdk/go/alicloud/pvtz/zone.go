// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pvtz

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// Private Zone can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:pvtz/zone:Zone example abc123456
// ```
type Zone struct {
	pulumi.CustomResourceState

	// Whether the Private Zone is ptr.
	IsPtr pulumi.BoolOutput `pulumi:"isPtr"`
	// The language. Valid values: "zh", "en", "jp".
	Lang pulumi.StringPtrOutput `pulumi:"lang"`
	// The name of the Private Zone.
	//
	// Deprecated: Field 'name' has been deprecated from version 1.107.0. Use 'zone_name' instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// The recursive DNS proxy. Valid values:
	// - ZONE: indicates that the recursive DNS proxy is disabled.
	// - RECORD: indicates that the recursive DNS proxy is enabled.
	//   Default to "ZONE".
	ProxyPattern pulumi.StringPtrOutput `pulumi:"proxyPattern"`
	// The count of the Private Zone Record.
	RecordCount pulumi.IntOutput `pulumi:"recordCount"`
	// The remark of the Private Zone.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// The Id of resource group which the Private Zone belongs.
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
	// The IP address of the client.
	UserClientIp pulumi.StringPtrOutput `pulumi:"userClientIp"`
	// The zoneName of the Private Zone.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewZone registers a new resource with the given unique name, arguments, and options.
func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOption) (*Zone, error) {
	if args == nil {
		args = &ZoneArgs{}
	}

	var resource Zone
	err := ctx.RegisterResource("alicloud:pvtz/zone:Zone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZone gets an existing Zone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneState, opts ...pulumi.ResourceOption) (*Zone, error) {
	var resource Zone
	err := ctx.ReadResource("alicloud:pvtz/zone:Zone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Zone resources.
type zoneState struct {
	// Whether the Private Zone is ptr.
	IsPtr *bool `pulumi:"isPtr"`
	// The language. Valid values: "zh", "en", "jp".
	Lang *string `pulumi:"lang"`
	// The name of the Private Zone.
	//
	// Deprecated: Field 'name' has been deprecated from version 1.107.0. Use 'zone_name' instead.
	Name *string `pulumi:"name"`
	// The recursive DNS proxy. Valid values:
	// - ZONE: indicates that the recursive DNS proxy is disabled.
	// - RECORD: indicates that the recursive DNS proxy is enabled.
	//   Default to "ZONE".
	ProxyPattern *string `pulumi:"proxyPattern"`
	// The count of the Private Zone Record.
	RecordCount *int `pulumi:"recordCount"`
	// The remark of the Private Zone.
	Remark *string `pulumi:"remark"`
	// The Id of resource group which the Private Zone belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The IP address of the client.
	UserClientIp *string `pulumi:"userClientIp"`
	// The zoneName of the Private Zone.
	ZoneName *string `pulumi:"zoneName"`
}

type ZoneState struct {
	// Whether the Private Zone is ptr.
	IsPtr pulumi.BoolPtrInput
	// The language. Valid values: "zh", "en", "jp".
	Lang pulumi.StringPtrInput
	// The name of the Private Zone.
	//
	// Deprecated: Field 'name' has been deprecated from version 1.107.0. Use 'zone_name' instead.
	Name pulumi.StringPtrInput
	// The recursive DNS proxy. Valid values:
	// - ZONE: indicates that the recursive DNS proxy is disabled.
	// - RECORD: indicates that the recursive DNS proxy is enabled.
	//   Default to "ZONE".
	ProxyPattern pulumi.StringPtrInput
	// The count of the Private Zone Record.
	RecordCount pulumi.IntPtrInput
	// The remark of the Private Zone.
	Remark pulumi.StringPtrInput
	// The Id of resource group which the Private Zone belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The IP address of the client.
	UserClientIp pulumi.StringPtrInput
	// The zoneName of the Private Zone.
	ZoneName pulumi.StringPtrInput
}

func (ZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneState)(nil)).Elem()
}

type zoneArgs struct {
	// The language. Valid values: "zh", "en", "jp".
	Lang *string `pulumi:"lang"`
	// The name of the Private Zone.
	//
	// Deprecated: Field 'name' has been deprecated from version 1.107.0. Use 'zone_name' instead.
	Name *string `pulumi:"name"`
	// The recursive DNS proxy. Valid values:
	// - ZONE: indicates that the recursive DNS proxy is disabled.
	// - RECORD: indicates that the recursive DNS proxy is enabled.
	//   Default to "ZONE".
	ProxyPattern *string `pulumi:"proxyPattern"`
	// The remark of the Private Zone.
	Remark *string `pulumi:"remark"`
	// The Id of resource group which the Private Zone belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The IP address of the client.
	UserClientIp *string `pulumi:"userClientIp"`
	// The zoneName of the Private Zone.
	ZoneName *string `pulumi:"zoneName"`
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	// The language. Valid values: "zh", "en", "jp".
	Lang pulumi.StringPtrInput
	// The name of the Private Zone.
	//
	// Deprecated: Field 'name' has been deprecated from version 1.107.0. Use 'zone_name' instead.
	Name pulumi.StringPtrInput
	// The recursive DNS proxy. Valid values:
	// - ZONE: indicates that the recursive DNS proxy is disabled.
	// - RECORD: indicates that the recursive DNS proxy is enabled.
	//   Default to "ZONE".
	ProxyPattern pulumi.StringPtrInput
	// The remark of the Private Zone.
	Remark pulumi.StringPtrInput
	// The Id of resource group which the Private Zone belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The IP address of the client.
	UserClientIp pulumi.StringPtrInput
	// The zoneName of the Private Zone.
	ZoneName pulumi.StringPtrInput
}

func (ZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneArgs)(nil)).Elem()
}

type ZoneInput interface {
	pulumi.Input

	ToZoneOutput() ZoneOutput
	ToZoneOutputWithContext(ctx context.Context) ZoneOutput
}

func (Zone) ElementType() reflect.Type {
	return reflect.TypeOf((*Zone)(nil)).Elem()
}

func (i Zone) ToZoneOutput() ZoneOutput {
	return i.ToZoneOutputWithContext(context.Background())
}

func (i Zone) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneOutput)
}

type ZoneOutput struct {
	*pulumi.OutputState
}

func (ZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneOutput)(nil)).Elem()
}

func (o ZoneOutput) ToZoneOutput() ZoneOutput {
	return o
}

func (o ZoneOutput) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ZoneOutput{})
}
