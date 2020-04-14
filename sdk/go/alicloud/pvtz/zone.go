// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pvtz

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Zone struct {
	pulumi.CustomResourceState

	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	IsPtr        pulumi.BoolOutput   `pulumi:"isPtr"`
	// The language. Valid values: "zh", "en", "jp".
	Lang pulumi.StringPtrOutput `pulumi:"lang"`
	// The name of the Private Zone.
	Name pulumi.StringOutput `pulumi:"name"`
	// The recursive DNS proxy. Valid values:
	// - ZONE: indicates that the recursive DNS proxy is disabled.
	// - RECORD: indicates that the recursive DNS proxy is enabled.
	ProxyPattern pulumi.StringPtrOutput `pulumi:"proxyPattern"`
	// The count of the Private Zone Record.
	RecordCount pulumi.IntOutput `pulumi:"recordCount"`
	// The remark of the Private Zone.
	Remark     pulumi.StringPtrOutput `pulumi:"remark"`
	UpdateTime pulumi.StringOutput    `pulumi:"updateTime"`
	// The IP address of the client.
	UserClientIp pulumi.StringPtrOutput `pulumi:"userClientIp"`
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
	CreationTime *string `pulumi:"creationTime"`
	IsPtr        *bool   `pulumi:"isPtr"`
	// The language. Valid values: "zh", "en", "jp".
	Lang *string `pulumi:"lang"`
	// The name of the Private Zone.
	Name *string `pulumi:"name"`
	// The recursive DNS proxy. Valid values:
	// - ZONE: indicates that the recursive DNS proxy is disabled.
	// - RECORD: indicates that the recursive DNS proxy is enabled.
	ProxyPattern *string `pulumi:"proxyPattern"`
	// The count of the Private Zone Record.
	RecordCount *int `pulumi:"recordCount"`
	// The remark of the Private Zone.
	Remark     *string `pulumi:"remark"`
	UpdateTime *string `pulumi:"updateTime"`
	// The IP address of the client.
	UserClientIp *string `pulumi:"userClientIp"`
}

type ZoneState struct {
	CreationTime pulumi.StringPtrInput
	IsPtr        pulumi.BoolPtrInput
	// The language. Valid values: "zh", "en", "jp".
	Lang pulumi.StringPtrInput
	// The name of the Private Zone.
	Name pulumi.StringPtrInput
	// The recursive DNS proxy. Valid values:
	// - ZONE: indicates that the recursive DNS proxy is disabled.
	// - RECORD: indicates that the recursive DNS proxy is enabled.
	ProxyPattern pulumi.StringPtrInput
	// The count of the Private Zone Record.
	RecordCount pulumi.IntPtrInput
	// The remark of the Private Zone.
	Remark     pulumi.StringPtrInput
	UpdateTime pulumi.StringPtrInput
	// The IP address of the client.
	UserClientIp pulumi.StringPtrInput
}

func (ZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneState)(nil)).Elem()
}

type zoneArgs struct {
	// The language. Valid values: "zh", "en", "jp".
	Lang *string `pulumi:"lang"`
	// The name of the Private Zone.
	Name *string `pulumi:"name"`
	// The recursive DNS proxy. Valid values:
	// - ZONE: indicates that the recursive DNS proxy is disabled.
	// - RECORD: indicates that the recursive DNS proxy is enabled.
	ProxyPattern *string `pulumi:"proxyPattern"`
	// The remark of the Private Zone.
	Remark *string `pulumi:"remark"`
	// The IP address of the client.
	UserClientIp *string `pulumi:"userClientIp"`
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	// The language. Valid values: "zh", "en", "jp".
	Lang pulumi.StringPtrInput
	// The name of the Private Zone.
	Name pulumi.StringPtrInput
	// The recursive DNS proxy. Valid values:
	// - ZONE: indicates that the recursive DNS proxy is disabled.
	// - RECORD: indicates that the recursive DNS proxy is enabled.
	ProxyPattern pulumi.StringPtrInput
	// The remark of the Private Zone.
	Remark pulumi.StringPtrInput
	// The IP address of the client.
	UserClientIp pulumi.StringPtrInput
}

func (ZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneArgs)(nil)).Elem()
}
