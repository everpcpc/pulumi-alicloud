// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// DNS instance be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:dns/instance:Instance example dns-cn-v0h1ldjhfff
// ```
type Instance struct {
	pulumi.CustomResourceState

	// DNS security level. Valid values: `no`, `basic`, `advanced`.
	DnsSecurity pulumi.StringOutput `pulumi:"dnsSecurity"`
	// Number of domain names bound.
	DomainNumbers pulumi.StringOutput    `pulumi:"domainNumbers"`
	PaymentType   pulumi.StringPtrOutput `pulumi:"paymentType"`
	// Creating a pre-paid instance, it must be set, the unit is month, please enter an integer multiple of 12 for annually paid products.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Automatic renewal period, the unit is month. When setting RenewalStatus to AutoRenewal, it must be set.
	RenewPeriod pulumi.IntPtrOutput `pulumi:"renewPeriod"`
	// Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, default to `ManualRenewal`.
	RenewalStatus pulumi.StringPtrOutput `pulumi:"renewalStatus"`
	// Paid package version. Valid values: `versionPersonal`, `versionEnterpriseBasic`, `versionEnterpriseAdvanced`.
	VersionCode pulumi.StringOutput `pulumi:"versionCode"`
	// Paid package version name.
	VersionName pulumi.StringOutput `pulumi:"versionName"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DnsSecurity == nil {
		return nil, errors.New("invalid value for required argument 'DnsSecurity'")
	}
	if args.DomainNumbers == nil {
		return nil, errors.New("invalid value for required argument 'DomainNumbers'")
	}
	if args.VersionCode == nil {
		return nil, errors.New("invalid value for required argument 'VersionCode'")
	}
	var resource Instance
	err := ctx.RegisterResource("alicloud:dns/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("alicloud:dns/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// DNS security level. Valid values: `no`, `basic`, `advanced`.
	DnsSecurity *string `pulumi:"dnsSecurity"`
	// Number of domain names bound.
	DomainNumbers *string `pulumi:"domainNumbers"`
	PaymentType   *string `pulumi:"paymentType"`
	// Creating a pre-paid instance, it must be set, the unit is month, please enter an integer multiple of 12 for annually paid products.
	Period *int `pulumi:"period"`
	// Automatic renewal period, the unit is month. When setting RenewalStatus to AutoRenewal, it must be set.
	RenewPeriod *int `pulumi:"renewPeriod"`
	// Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, default to `ManualRenewal`.
	RenewalStatus *string `pulumi:"renewalStatus"`
	// Paid package version. Valid values: `versionPersonal`, `versionEnterpriseBasic`, `versionEnterpriseAdvanced`.
	VersionCode *string `pulumi:"versionCode"`
	// Paid package version name.
	VersionName *string `pulumi:"versionName"`
}

type InstanceState struct {
	// DNS security level. Valid values: `no`, `basic`, `advanced`.
	DnsSecurity pulumi.StringPtrInput
	// Number of domain names bound.
	DomainNumbers pulumi.StringPtrInput
	PaymentType   pulumi.StringPtrInput
	// Creating a pre-paid instance, it must be set, the unit is month, please enter an integer multiple of 12 for annually paid products.
	Period pulumi.IntPtrInput
	// Automatic renewal period, the unit is month. When setting RenewalStatus to AutoRenewal, it must be set.
	RenewPeriod pulumi.IntPtrInput
	// Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, default to `ManualRenewal`.
	RenewalStatus pulumi.StringPtrInput
	// Paid package version. Valid values: `versionPersonal`, `versionEnterpriseBasic`, `versionEnterpriseAdvanced`.
	VersionCode pulumi.StringPtrInput
	// Paid package version name.
	VersionName pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// DNS security level. Valid values: `no`, `basic`, `advanced`.
	DnsSecurity string `pulumi:"dnsSecurity"`
	// Number of domain names bound.
	DomainNumbers string  `pulumi:"domainNumbers"`
	PaymentType   *string `pulumi:"paymentType"`
	// Creating a pre-paid instance, it must be set, the unit is month, please enter an integer multiple of 12 for annually paid products.
	Period *int `pulumi:"period"`
	// Automatic renewal period, the unit is month. When setting RenewalStatus to AutoRenewal, it must be set.
	RenewPeriod *int `pulumi:"renewPeriod"`
	// Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, default to `ManualRenewal`.
	RenewalStatus *string `pulumi:"renewalStatus"`
	// Paid package version. Valid values: `versionPersonal`, `versionEnterpriseBasic`, `versionEnterpriseAdvanced`.
	VersionCode string `pulumi:"versionCode"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// DNS security level. Valid values: `no`, `basic`, `advanced`.
	DnsSecurity pulumi.StringInput
	// Number of domain names bound.
	DomainNumbers pulumi.StringInput
	PaymentType   pulumi.StringPtrInput
	// Creating a pre-paid instance, it must be set, the unit is month, please enter an integer multiple of 12 for annually paid products.
	Period pulumi.IntPtrInput
	// Automatic renewal period, the unit is month. When setting RenewalStatus to AutoRenewal, it must be set.
	RenewPeriod pulumi.IntPtrInput
	// Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, default to `ManualRenewal`.
	RenewalStatus pulumi.StringPtrInput
	// Paid package version. Valid values: `versionPersonal`, `versionEnterpriseBasic`, `versionEnterpriseAdvanced`.
	VersionCode pulumi.StringInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (Instance) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil)).Elem()
}

func (i Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

type InstanceOutput struct {
	*pulumi.OutputState
}

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceOutput)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstanceOutput{})
}
