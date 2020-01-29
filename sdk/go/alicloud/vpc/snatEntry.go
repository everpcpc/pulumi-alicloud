// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vpc

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a snat resource.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/snat_entry.html.markdown.
type SnatEntry struct {
	pulumi.CustomResourceState

	// The id of the snat entry on the server.
	SnatEntryId pulumi.StringOutput `pulumi:"snatEntryId"`
	// The SNAT ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidthPackages`.
	SnatIp pulumi.StringOutput `pulumi:"snatIp"`
	// The value can get from `vpc.NatGateway` Attributes "snatTableIds".
	SnatTableId pulumi.StringOutput `pulumi:"snatTableId"`
	// The vswitch ID.
	SourceVswitchId pulumi.StringOutput `pulumi:"sourceVswitchId"`
}

// NewSnatEntry registers a new resource with the given unique name, arguments, and options.
func NewSnatEntry(ctx *pulumi.Context,
	name string, args *SnatEntryArgs, opts ...pulumi.ResourceOption) (*SnatEntry, error) {
	if args == nil || args.SnatIp == nil {
		return nil, errors.New("missing required argument 'SnatIp'")
	}
	if args == nil || args.SnatTableId == nil {
		return nil, errors.New("missing required argument 'SnatTableId'")
	}
	if args == nil || args.SourceVswitchId == nil {
		return nil, errors.New("missing required argument 'SourceVswitchId'")
	}
	if args == nil {
		args = &SnatEntryArgs{}
	}
	var resource SnatEntry
	err := ctx.RegisterResource("alicloud:vpc/snatEntry:SnatEntry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnatEntry gets an existing SnatEntry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnatEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnatEntryState, opts ...pulumi.ResourceOption) (*SnatEntry, error) {
	var resource SnatEntry
	err := ctx.ReadResource("alicloud:vpc/snatEntry:SnatEntry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnatEntry resources.
type snatEntryState struct {
	// The id of the snat entry on the server.
	SnatEntryId *string `pulumi:"snatEntryId"`
	// The SNAT ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidthPackages`.
	SnatIp *string `pulumi:"snatIp"`
	// The value can get from `vpc.NatGateway` Attributes "snatTableIds".
	SnatTableId *string `pulumi:"snatTableId"`
	// The vswitch ID.
	SourceVswitchId *string `pulumi:"sourceVswitchId"`
}

type SnatEntryState struct {
	// The id of the snat entry on the server.
	SnatEntryId pulumi.StringPtrInput
	// The SNAT ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidthPackages`.
	SnatIp pulumi.StringPtrInput
	// The value can get from `vpc.NatGateway` Attributes "snatTableIds".
	SnatTableId pulumi.StringPtrInput
	// The vswitch ID.
	SourceVswitchId pulumi.StringPtrInput
}

func (SnatEntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*snatEntryState)(nil)).Elem()
}

type snatEntryArgs struct {
	// The SNAT ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidthPackages`.
	SnatIp string `pulumi:"snatIp"`
	// The value can get from `vpc.NatGateway` Attributes "snatTableIds".
	SnatTableId string `pulumi:"snatTableId"`
	// The vswitch ID.
	SourceVswitchId string `pulumi:"sourceVswitchId"`
}

// The set of arguments for constructing a SnatEntry resource.
type SnatEntryArgs struct {
	// The SNAT ip address, the ip must along bandwidth package public ip which `vpc.NatGateway` argument `bandwidthPackages`.
	SnatIp pulumi.StringInput
	// The value can get from `vpc.NatGateway` Attributes "snatTableIds".
	SnatTableId pulumi.StringInput
	// The vswitch ID.
	SourceVswitchId pulumi.StringInput
}

func (SnatEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snatEntryArgs)(nil)).Elem()
}

