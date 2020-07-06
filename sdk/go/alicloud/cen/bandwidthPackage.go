// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CEN bandwidth package resource. The CEN bandwidth package is an abstracted object that includes an interconnection bandwidth and interconnection areas. To buy a bandwidth package, you must specify the areas to connect. An area consists of one or more Alibaba Cloud regions. The areas in CEN include Mainland China, Asia Pacific, North America, and Europe.
//
// For information about CEN and how to use it, see [Manage bandwidth packages](https://www.alibabacloud.com/help/doc-detail/65982.htm).
type BandwidthPackage struct {
	pulumi.CustomResourceState

	// The bandwidth in Mbps of the bandwidth package. Cannot be less than 2Mbps.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// The billing method. Valid value: PostPaid | PrePaid. Default to PostPaid. If set to PrePaid, the bandwidth package can't be deleted before expired time.
	ChargeType pulumi.StringPtrOutput `pulumi:"chargeType"`
	// The description of the bandwidth package. Default to null.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The time of the bandwidth package to expire.
	ExpiredTime pulumi.StringOutput `pulumi:"expiredTime"`
	// List of the two areas to connect. Valid value: China | North-America | Asia-Pacific | Europe | Middle-East | Australia.
	GeographicRegionIds pulumi.StringArrayOutput `pulumi:"geographicRegionIds"`
	// The name of the bandwidth package. Defaults to null.
	Name pulumi.StringOutput `pulumi:"name"`
	// The purchase period in month. Valid value: 1, 2, 3, 6, 12. Default to 1.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// The status of the bandwidth, including "InUse" and "Idle".
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewBandwidthPackage registers a new resource with the given unique name, arguments, and options.
func NewBandwidthPackage(ctx *pulumi.Context,
	name string, args *BandwidthPackageArgs, opts ...pulumi.ResourceOption) (*BandwidthPackage, error) {
	if args == nil || args.Bandwidth == nil {
		return nil, errors.New("missing required argument 'Bandwidth'")
	}
	if args == nil || args.GeographicRegionIds == nil {
		return nil, errors.New("missing required argument 'GeographicRegionIds'")
	}
	if args == nil {
		args = &BandwidthPackageArgs{}
	}
	var resource BandwidthPackage
	err := ctx.RegisterResource("alicloud:cen/bandwidthPackage:BandwidthPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBandwidthPackage gets an existing BandwidthPackage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBandwidthPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BandwidthPackageState, opts ...pulumi.ResourceOption) (*BandwidthPackage, error) {
	var resource BandwidthPackage
	err := ctx.ReadResource("alicloud:cen/bandwidthPackage:BandwidthPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BandwidthPackage resources.
type bandwidthPackageState struct {
	// The bandwidth in Mbps of the bandwidth package. Cannot be less than 2Mbps.
	Bandwidth *int `pulumi:"bandwidth"`
	// The billing method. Valid value: PostPaid | PrePaid. Default to PostPaid. If set to PrePaid, the bandwidth package can't be deleted before expired time.
	ChargeType *string `pulumi:"chargeType"`
	// The description of the bandwidth package. Default to null.
	Description *string `pulumi:"description"`
	// The time of the bandwidth package to expire.
	ExpiredTime *string `pulumi:"expiredTime"`
	// List of the two areas to connect. Valid value: China | North-America | Asia-Pacific | Europe | Middle-East | Australia.
	GeographicRegionIds []string `pulumi:"geographicRegionIds"`
	// The name of the bandwidth package. Defaults to null.
	Name *string `pulumi:"name"`
	// The purchase period in month. Valid value: 1, 2, 3, 6, 12. Default to 1.
	Period *int `pulumi:"period"`
	// The status of the bandwidth, including "InUse" and "Idle".
	Status *string `pulumi:"status"`
}

type BandwidthPackageState struct {
	// The bandwidth in Mbps of the bandwidth package. Cannot be less than 2Mbps.
	Bandwidth pulumi.IntPtrInput
	// The billing method. Valid value: PostPaid | PrePaid. Default to PostPaid. If set to PrePaid, the bandwidth package can't be deleted before expired time.
	ChargeType pulumi.StringPtrInput
	// The description of the bandwidth package. Default to null.
	Description pulumi.StringPtrInput
	// The time of the bandwidth package to expire.
	ExpiredTime pulumi.StringPtrInput
	// List of the two areas to connect. Valid value: China | North-America | Asia-Pacific | Europe | Middle-East | Australia.
	GeographicRegionIds pulumi.StringArrayInput
	// The name of the bandwidth package. Defaults to null.
	Name pulumi.StringPtrInput
	// The purchase period in month. Valid value: 1, 2, 3, 6, 12. Default to 1.
	Period pulumi.IntPtrInput
	// The status of the bandwidth, including "InUse" and "Idle".
	Status pulumi.StringPtrInput
}

func (BandwidthPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthPackageState)(nil)).Elem()
}

type bandwidthPackageArgs struct {
	// The bandwidth in Mbps of the bandwidth package. Cannot be less than 2Mbps.
	Bandwidth int `pulumi:"bandwidth"`
	// The billing method. Valid value: PostPaid | PrePaid. Default to PostPaid. If set to PrePaid, the bandwidth package can't be deleted before expired time.
	ChargeType *string `pulumi:"chargeType"`
	// The description of the bandwidth package. Default to null.
	Description *string `pulumi:"description"`
	// List of the two areas to connect. Valid value: China | North-America | Asia-Pacific | Europe | Middle-East | Australia.
	GeographicRegionIds []string `pulumi:"geographicRegionIds"`
	// The name of the bandwidth package. Defaults to null.
	Name *string `pulumi:"name"`
	// The purchase period in month. Valid value: 1, 2, 3, 6, 12. Default to 1.
	Period *int `pulumi:"period"`
}

// The set of arguments for constructing a BandwidthPackage resource.
type BandwidthPackageArgs struct {
	// The bandwidth in Mbps of the bandwidth package. Cannot be less than 2Mbps.
	Bandwidth pulumi.IntInput
	// The billing method. Valid value: PostPaid | PrePaid. Default to PostPaid. If set to PrePaid, the bandwidth package can't be deleted before expired time.
	ChargeType pulumi.StringPtrInput
	// The description of the bandwidth package. Default to null.
	Description pulumi.StringPtrInput
	// List of the two areas to connect. Valid value: China | North-America | Asia-Pacific | Europe | Middle-East | Australia.
	GeographicRegionIds pulumi.StringArrayInput
	// The name of the bandwidth package. Defaults to null.
	Name pulumi.StringPtrInput
	// The purchase period in month. Valid value: 1, 2, 3, 6, 12. Default to 1.
	Period pulumi.IntPtrInput
}

func (BandwidthPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthPackageArgs)(nil)).Elem()
}
