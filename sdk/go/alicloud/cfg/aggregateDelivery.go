// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Config Aggregate Delivery resource.
//
// For information about Cloud Config Aggregate Delivery and how to use it, see [What is Aggregate Delivery](https://www.alibabacloud.com/help/en/cloud-config/latest/delivery-destination-services-overview).
//
// > **NOTE:** Available in v1.172.0+.
//
// ## Import
//
// Cloud Config Aggregate Delivery can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:cfg/aggregateDelivery:AggregateDelivery example <aggregator_id>:<delivery_channel_id>
// ```
type AggregateDelivery struct {
	pulumi.CustomResourceState

	// The ID of the Aggregator.
	AggregatorId pulumi.StringOutput `pulumi:"aggregatorId"`
	// Open or close delivery configuration change history.
	ConfigurationItemChangeNotification pulumi.BoolOutput `pulumi:"configurationItemChangeNotification"`
	// Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `deliveryChannelType` is `OSS`.
	ConfigurationSnapshot pulumi.BoolOutput `pulumi:"configurationSnapshot"`
	// The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `deliveryChannelType` is `MNS`.
	DeliveryChannelCondition pulumi.StringPtrOutput `pulumi:"deliveryChannelCondition"`
	// The ID of the delivery method.
	DeliveryChannelId pulumi.StringOutput `pulumi:"deliveryChannelId"`
	// The name of the delivery method.
	DeliveryChannelName pulumi.StringPtrOutput `pulumi:"deliveryChannelName"`
	// The ARN of the delivery destination. The value must be in one of the following formats:
	// * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
	// * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
	// * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
	DeliveryChannelTargetArn pulumi.StringOutput `pulumi:"deliveryChannelTargetArn"`
	// The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
	DeliveryChannelType pulumi.StringOutput `pulumi:"deliveryChannelType"`
	// The description of the delivery method.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `deliveryChannelType` is `SLS` or `MNS`.
	NonCompliantNotification pulumi.BoolOutput `pulumi:"nonCompliantNotification"`
	// The oss ARN of the delivery channel when the value data oversized limit.
	// * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
	// * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
	OversizedDataOssTargetArn pulumi.StringPtrOutput `pulumi:"oversizedDataOssTargetArn"`
	// The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
	Status pulumi.IntOutput `pulumi:"status"`
}

// NewAggregateDelivery registers a new resource with the given unique name, arguments, and options.
func NewAggregateDelivery(ctx *pulumi.Context,
	name string, args *AggregateDeliveryArgs, opts ...pulumi.ResourceOption) (*AggregateDelivery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AggregatorId == nil {
		return nil, errors.New("invalid value for required argument 'AggregatorId'")
	}
	if args.DeliveryChannelTargetArn == nil {
		return nil, errors.New("invalid value for required argument 'DeliveryChannelTargetArn'")
	}
	if args.DeliveryChannelType == nil {
		return nil, errors.New("invalid value for required argument 'DeliveryChannelType'")
	}
	var resource AggregateDelivery
	err := ctx.RegisterResource("alicloud:cfg/aggregateDelivery:AggregateDelivery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAggregateDelivery gets an existing AggregateDelivery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAggregateDelivery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AggregateDeliveryState, opts ...pulumi.ResourceOption) (*AggregateDelivery, error) {
	var resource AggregateDelivery
	err := ctx.ReadResource("alicloud:cfg/aggregateDelivery:AggregateDelivery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AggregateDelivery resources.
type aggregateDeliveryState struct {
	// The ID of the Aggregator.
	AggregatorId *string `pulumi:"aggregatorId"`
	// Open or close delivery configuration change history.
	ConfigurationItemChangeNotification *bool `pulumi:"configurationItemChangeNotification"`
	// Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `deliveryChannelType` is `OSS`.
	ConfigurationSnapshot *bool `pulumi:"configurationSnapshot"`
	// The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `deliveryChannelType` is `MNS`.
	DeliveryChannelCondition *string `pulumi:"deliveryChannelCondition"`
	// The ID of the delivery method.
	DeliveryChannelId *string `pulumi:"deliveryChannelId"`
	// The name of the delivery method.
	DeliveryChannelName *string `pulumi:"deliveryChannelName"`
	// The ARN of the delivery destination. The value must be in one of the following formats:
	// * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
	// * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
	// * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
	DeliveryChannelTargetArn *string `pulumi:"deliveryChannelTargetArn"`
	// The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
	DeliveryChannelType *string `pulumi:"deliveryChannelType"`
	// The description of the delivery method.
	Description *string `pulumi:"description"`
	// Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `deliveryChannelType` is `SLS` or `MNS`.
	NonCompliantNotification *bool `pulumi:"nonCompliantNotification"`
	// The oss ARN of the delivery channel when the value data oversized limit.
	// * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
	// * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
	OversizedDataOssTargetArn *string `pulumi:"oversizedDataOssTargetArn"`
	// The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
	Status *int `pulumi:"status"`
}

type AggregateDeliveryState struct {
	// The ID of the Aggregator.
	AggregatorId pulumi.StringPtrInput
	// Open or close delivery configuration change history.
	ConfigurationItemChangeNotification pulumi.BoolPtrInput
	// Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `deliveryChannelType` is `OSS`.
	ConfigurationSnapshot pulumi.BoolPtrInput
	// The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `deliveryChannelType` is `MNS`.
	DeliveryChannelCondition pulumi.StringPtrInput
	// The ID of the delivery method.
	DeliveryChannelId pulumi.StringPtrInput
	// The name of the delivery method.
	DeliveryChannelName pulumi.StringPtrInput
	// The ARN of the delivery destination. The value must be in one of the following formats:
	// * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
	// * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
	// * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
	DeliveryChannelTargetArn pulumi.StringPtrInput
	// The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
	DeliveryChannelType pulumi.StringPtrInput
	// The description of the delivery method.
	Description pulumi.StringPtrInput
	// Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `deliveryChannelType` is `SLS` or `MNS`.
	NonCompliantNotification pulumi.BoolPtrInput
	// The oss ARN of the delivery channel when the value data oversized limit.
	// * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
	// * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
	OversizedDataOssTargetArn pulumi.StringPtrInput
	// The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
	Status pulumi.IntPtrInput
}

func (AggregateDeliveryState) ElementType() reflect.Type {
	return reflect.TypeOf((*aggregateDeliveryState)(nil)).Elem()
}

type aggregateDeliveryArgs struct {
	// The ID of the Aggregator.
	AggregatorId string `pulumi:"aggregatorId"`
	// Open or close delivery configuration change history.
	ConfigurationItemChangeNotification *bool `pulumi:"configurationItemChangeNotification"`
	// Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `deliveryChannelType` is `OSS`.
	ConfigurationSnapshot *bool `pulumi:"configurationSnapshot"`
	// The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `deliveryChannelType` is `MNS`.
	DeliveryChannelCondition *string `pulumi:"deliveryChannelCondition"`
	// The name of the delivery method.
	DeliveryChannelName *string `pulumi:"deliveryChannelName"`
	// The ARN of the delivery destination. The value must be in one of the following formats:
	// * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
	// * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
	// * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
	DeliveryChannelTargetArn string `pulumi:"deliveryChannelTargetArn"`
	// The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
	DeliveryChannelType string `pulumi:"deliveryChannelType"`
	// The description of the delivery method.
	Description *string `pulumi:"description"`
	// Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `deliveryChannelType` is `SLS` or `MNS`.
	NonCompliantNotification *bool `pulumi:"nonCompliantNotification"`
	// The oss ARN of the delivery channel when the value data oversized limit.
	// * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
	// * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
	OversizedDataOssTargetArn *string `pulumi:"oversizedDataOssTargetArn"`
	// The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
	Status *int `pulumi:"status"`
}

// The set of arguments for constructing a AggregateDelivery resource.
type AggregateDeliveryArgs struct {
	// The ID of the Aggregator.
	AggregatorId pulumi.StringInput
	// Open or close delivery configuration change history.
	ConfigurationItemChangeNotification pulumi.BoolPtrInput
	// Open or close timed snapshot of shipping resources. **NOTE:** The attribute is valid when the attribute `deliveryChannelType` is `OSS`.
	ConfigurationSnapshot pulumi.BoolPtrInput
	// The rule attached to the delivery method. Please refer to api [CreateConfigDeliveryChannel](https://help.aliyun.com/document_detail/429798.html) for example format. **NOTE:** The attribute is valid when the attribute `deliveryChannelType` is `MNS`.
	DeliveryChannelCondition pulumi.StringPtrInput
	// The name of the delivery method.
	DeliveryChannelName pulumi.StringPtrInput
	// The ARN of the delivery destination. The value must be in one of the following formats:
	// * `acs:oss:{RegionId}:{Aliuid}:{bucketName}`: if your delivery destination is an Object Storage Service (OSS) bucket.
	// * `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}`: if your delivery destination is a Message Service (MNS) topic.
	// * `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}`: if your delivery destination is a Log Service Logstore.
	DeliveryChannelTargetArn pulumi.StringInput
	// The type of the delivery method. Valid values: `OSS`: Object Storage, `MNS`: Message Service, `SLS`: Log Service.
	DeliveryChannelType pulumi.StringInput
	// The description of the delivery method.
	Description pulumi.StringPtrInput
	// Open or close non-compliance events of delivery resources. **NOTE:** The attribute is valid when the attribute `deliveryChannelType` is `SLS` or `MNS`.
	NonCompliantNotification pulumi.BoolPtrInput
	// The oss ARN of the delivery channel when the value data oversized limit.
	// * The value must be in one of the following formats: `acs:oss:{RegionId}:{accountId}:{bucketName}`, if your delivery destination is an Object Storage Service (OSS) bucket.
	// * Only delivery channels `SLS` and `MNS` are supported. The delivery channel limit for Log Service SLS is 1 MB, and the delivery channel limit for Message Service MNS is 64 KB.
	OversizedDataOssTargetArn pulumi.StringPtrInput
	// The status of the delivery method. Valid values: `0`: The delivery method is disabled. `1`: The delivery destination is enabled. This is the default value.
	Status pulumi.IntPtrInput
}

func (AggregateDeliveryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aggregateDeliveryArgs)(nil)).Elem()
}

type AggregateDeliveryInput interface {
	pulumi.Input

	ToAggregateDeliveryOutput() AggregateDeliveryOutput
	ToAggregateDeliveryOutputWithContext(ctx context.Context) AggregateDeliveryOutput
}

func (*AggregateDelivery) ElementType() reflect.Type {
	return reflect.TypeOf((**AggregateDelivery)(nil)).Elem()
}

func (i *AggregateDelivery) ToAggregateDeliveryOutput() AggregateDeliveryOutput {
	return i.ToAggregateDeliveryOutputWithContext(context.Background())
}

func (i *AggregateDelivery) ToAggregateDeliveryOutputWithContext(ctx context.Context) AggregateDeliveryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregateDeliveryOutput)
}

// AggregateDeliveryArrayInput is an input type that accepts AggregateDeliveryArray and AggregateDeliveryArrayOutput values.
// You can construct a concrete instance of `AggregateDeliveryArrayInput` via:
//
//          AggregateDeliveryArray{ AggregateDeliveryArgs{...} }
type AggregateDeliveryArrayInput interface {
	pulumi.Input

	ToAggregateDeliveryArrayOutput() AggregateDeliveryArrayOutput
	ToAggregateDeliveryArrayOutputWithContext(context.Context) AggregateDeliveryArrayOutput
}

type AggregateDeliveryArray []AggregateDeliveryInput

func (AggregateDeliveryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AggregateDelivery)(nil)).Elem()
}

func (i AggregateDeliveryArray) ToAggregateDeliveryArrayOutput() AggregateDeliveryArrayOutput {
	return i.ToAggregateDeliveryArrayOutputWithContext(context.Background())
}

func (i AggregateDeliveryArray) ToAggregateDeliveryArrayOutputWithContext(ctx context.Context) AggregateDeliveryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregateDeliveryArrayOutput)
}

// AggregateDeliveryMapInput is an input type that accepts AggregateDeliveryMap and AggregateDeliveryMapOutput values.
// You can construct a concrete instance of `AggregateDeliveryMapInput` via:
//
//          AggregateDeliveryMap{ "key": AggregateDeliveryArgs{...} }
type AggregateDeliveryMapInput interface {
	pulumi.Input

	ToAggregateDeliveryMapOutput() AggregateDeliveryMapOutput
	ToAggregateDeliveryMapOutputWithContext(context.Context) AggregateDeliveryMapOutput
}

type AggregateDeliveryMap map[string]AggregateDeliveryInput

func (AggregateDeliveryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AggregateDelivery)(nil)).Elem()
}

func (i AggregateDeliveryMap) ToAggregateDeliveryMapOutput() AggregateDeliveryMapOutput {
	return i.ToAggregateDeliveryMapOutputWithContext(context.Background())
}

func (i AggregateDeliveryMap) ToAggregateDeliveryMapOutputWithContext(ctx context.Context) AggregateDeliveryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregateDeliveryMapOutput)
}

type AggregateDeliveryOutput struct{ *pulumi.OutputState }

func (AggregateDeliveryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AggregateDelivery)(nil)).Elem()
}

func (o AggregateDeliveryOutput) ToAggregateDeliveryOutput() AggregateDeliveryOutput {
	return o
}

func (o AggregateDeliveryOutput) ToAggregateDeliveryOutputWithContext(ctx context.Context) AggregateDeliveryOutput {
	return o
}

type AggregateDeliveryArrayOutput struct{ *pulumi.OutputState }

func (AggregateDeliveryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AggregateDelivery)(nil)).Elem()
}

func (o AggregateDeliveryArrayOutput) ToAggregateDeliveryArrayOutput() AggregateDeliveryArrayOutput {
	return o
}

func (o AggregateDeliveryArrayOutput) ToAggregateDeliveryArrayOutputWithContext(ctx context.Context) AggregateDeliveryArrayOutput {
	return o
}

func (o AggregateDeliveryArrayOutput) Index(i pulumi.IntInput) AggregateDeliveryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AggregateDelivery {
		return vs[0].([]*AggregateDelivery)[vs[1].(int)]
	}).(AggregateDeliveryOutput)
}

type AggregateDeliveryMapOutput struct{ *pulumi.OutputState }

func (AggregateDeliveryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AggregateDelivery)(nil)).Elem()
}

func (o AggregateDeliveryMapOutput) ToAggregateDeliveryMapOutput() AggregateDeliveryMapOutput {
	return o
}

func (o AggregateDeliveryMapOutput) ToAggregateDeliveryMapOutputWithContext(ctx context.Context) AggregateDeliveryMapOutput {
	return o
}

func (o AggregateDeliveryMapOutput) MapIndex(k pulumi.StringInput) AggregateDeliveryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AggregateDelivery {
		return vs[0].(map[string]*AggregateDelivery)[vs[1].(string)]
	}).(AggregateDeliveryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AggregateDeliveryInput)(nil)).Elem(), &AggregateDelivery{})
	pulumi.RegisterInputType(reflect.TypeOf((*AggregateDeliveryArrayInput)(nil)).Elem(), AggregateDeliveryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AggregateDeliveryMapInput)(nil)).Elem(), AggregateDeliveryMap{})
	pulumi.RegisterOutputType(AggregateDeliveryOutput{})
	pulumi.RegisterOutputType(AggregateDeliveryArrayOutput{})
	pulumi.RegisterOutputType(AggregateDeliveryMapOutput{})
}