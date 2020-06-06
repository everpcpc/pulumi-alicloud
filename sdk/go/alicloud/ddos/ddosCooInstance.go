// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ddos

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// BGP-Line Anti-DDoS instance resource. "Ddoscoo" is the short term of this product. See [What is Anti-DDoS Pro](https://www.alibabacloud.com/help/doc-detail/69319.htm).
//
// > **NOTE:** The product region only support cn-hangzhou.
//
// > **NOTE:** The endpoint of bssopenapi used only support "business.aliyuncs.com" at present.
//
// > **NOTE:** Available in 1.37.0+ .
type DdosCooInstance struct {
	pulumi.CustomResourceState

	// Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.
	Bandwidth pulumi.StringOutput `pulumi:"bandwidth"`
	// Base defend bandwidth of the instance. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.
	BaseBandwidth pulumi.StringOutput `pulumi:"baseBandwidth"`
	// Domain retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
	DomainCount pulumi.StringOutput `pulumi:"domainCount"`
	// Name of the instance. This name can have a string of 1 to 63 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The duration that you will buy Ddoscoo instance (in month). Valid values: [1~9], 12, 24, 36. Default to 1. At present, the provider does not support modify "period".
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Port retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
	PortCount pulumi.StringOutput `pulumi:"portCount"`
	// Business bandwidth of the instance. At leaset 100. Increased 100 per step, such as 100, 200, 300. The unit is Mbps. Only support upgrade.
	ServiceBandwidth pulumi.StringOutput `pulumi:"serviceBandwidth"`
}

// NewDdosCooInstance registers a new resource with the given unique name, arguments, and options.
func NewDdosCooInstance(ctx *pulumi.Context,
	name string, args *DdosCooInstanceArgs, opts ...pulumi.ResourceOption) (*DdosCooInstance, error) {
	if args == nil || args.Bandwidth == nil {
		return nil, errors.New("missing required argument 'Bandwidth'")
	}
	if args == nil || args.BaseBandwidth == nil {
		return nil, errors.New("missing required argument 'BaseBandwidth'")
	}
	if args == nil || args.DomainCount == nil {
		return nil, errors.New("missing required argument 'DomainCount'")
	}
	if args == nil || args.PortCount == nil {
		return nil, errors.New("missing required argument 'PortCount'")
	}
	if args == nil || args.ServiceBandwidth == nil {
		return nil, errors.New("missing required argument 'ServiceBandwidth'")
	}
	if args == nil {
		args = &DdosCooInstanceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("alicloud:dns/ddosCooInstance:DdosCooInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource DdosCooInstance
	err := ctx.RegisterResource("alicloud:ddos/ddosCooInstance:DdosCooInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDdosCooInstance gets an existing DdosCooInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDdosCooInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DdosCooInstanceState, opts ...pulumi.ResourceOption) (*DdosCooInstance, error) {
	var resource DdosCooInstance
	err := ctx.ReadResource("alicloud:ddos/ddosCooInstance:DdosCooInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DdosCooInstance resources.
type ddosCooInstanceState struct {
	// Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.
	Bandwidth *string `pulumi:"bandwidth"`
	// Base defend bandwidth of the instance. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.
	BaseBandwidth *string `pulumi:"baseBandwidth"`
	// Domain retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
	DomainCount *string `pulumi:"domainCount"`
	// Name of the instance. This name can have a string of 1 to 63 characters.
	Name *string `pulumi:"name"`
	// The duration that you will buy Ddoscoo instance (in month). Valid values: [1~9], 12, 24, 36. Default to 1. At present, the provider does not support modify "period".
	Period *int `pulumi:"period"`
	// Port retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
	PortCount *string `pulumi:"portCount"`
	// Business bandwidth of the instance. At leaset 100. Increased 100 per step, such as 100, 200, 300. The unit is Mbps. Only support upgrade.
	ServiceBandwidth *string `pulumi:"serviceBandwidth"`
}

type DdosCooInstanceState struct {
	// Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.
	Bandwidth pulumi.StringPtrInput
	// Base defend bandwidth of the instance. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.
	BaseBandwidth pulumi.StringPtrInput
	// Domain retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
	DomainCount pulumi.StringPtrInput
	// Name of the instance. This name can have a string of 1 to 63 characters.
	Name pulumi.StringPtrInput
	// The duration that you will buy Ddoscoo instance (in month). Valid values: [1~9], 12, 24, 36. Default to 1. At present, the provider does not support modify "period".
	Period pulumi.IntPtrInput
	// Port retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
	PortCount pulumi.StringPtrInput
	// Business bandwidth of the instance. At leaset 100. Increased 100 per step, such as 100, 200, 300. The unit is Mbps. Only support upgrade.
	ServiceBandwidth pulumi.StringPtrInput
}

func (DdosCooInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosCooInstanceState)(nil)).Elem()
}

type ddosCooInstanceArgs struct {
	// Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.
	Bandwidth string `pulumi:"bandwidth"`
	// Base defend bandwidth of the instance. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.
	BaseBandwidth string `pulumi:"baseBandwidth"`
	// Domain retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
	DomainCount string `pulumi:"domainCount"`
	// Name of the instance. This name can have a string of 1 to 63 characters.
	Name *string `pulumi:"name"`
	// The duration that you will buy Ddoscoo instance (in month). Valid values: [1~9], 12, 24, 36. Default to 1. At present, the provider does not support modify "period".
	Period *int `pulumi:"period"`
	// Port retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
	PortCount string `pulumi:"portCount"`
	// Business bandwidth of the instance. At leaset 100. Increased 100 per step, such as 100, 200, 300. The unit is Mbps. Only support upgrade.
	ServiceBandwidth string `pulumi:"serviceBandwidth"`
}

// The set of arguments for constructing a DdosCooInstance resource.
type DdosCooInstanceArgs struct {
	// Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.
	Bandwidth pulumi.StringInput
	// Base defend bandwidth of the instance. Valid values: 30, 60, 100, 300, 400, 500, 600. The unit is Gbps. Only support upgrade.
	BaseBandwidth pulumi.StringInput
	// Domain retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
	DomainCount pulumi.StringInput
	// Name of the instance. This name can have a string of 1 to 63 characters.
	Name pulumi.StringPtrInput
	// The duration that you will buy Ddoscoo instance (in month). Valid values: [1~9], 12, 24, 36. Default to 1. At present, the provider does not support modify "period".
	Period pulumi.IntPtrInput
	// Port retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
	PortCount pulumi.StringInput
	// Business bandwidth of the instance. At leaset 100. Increased 100 per step, such as 100, 200, 300. The unit is Mbps. Only support upgrade.
	ServiceBandwidth pulumi.StringInput
}

func (DdosCooInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosCooInstanceArgs)(nil)).Elem()
}