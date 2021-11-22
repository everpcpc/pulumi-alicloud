// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rocketmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an ONS topic resource.
//
// For more information about how to use it, see [RocketMQ Topic Management API](https://www.alibabacloud.com/help/doc-detail/29591.html).
//
// > **NOTE:** Available in 1.53.0+
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/rocketmq"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		name := "onsInstanceName"
// 		if param := cfg.Get("name"); param != "" {
// 			name = param
// 		}
// 		topic := "onsTopicName"
// 		if param := cfg.Get("topic"); param != "" {
// 			topic = param
// 		}
// 		defaultInstance, err := rocketmq.NewInstance(ctx, "defaultInstance", &rocketmq.InstanceArgs{
// 			Remark: pulumi.String("default_ons_instance_remark"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = rocketmq.NewTopic(ctx, "defaultTopic", &rocketmq.TopicArgs{
// 			TopicName:   pulumi.String(topic),
// 			InstanceId:  defaultInstance.ID(),
// 			MessageType: pulumi.Int(0),
// 			Remark:      pulumi.String("dafault_ons_topic_remark"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ONS TOPIC can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:rocketmq/topic:Topic topic MQ_INST_1234567890_Baso1234567:onsTopicDemo
// ```
type Topic struct {
	pulumi.CustomResourceState

	// ID of the ONS Instance that owns the topics.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The type of the message. Read [Ons Topic Create](https://www.alibabacloud.com/help/doc-detail/29591.html) for further details.
	MessageType pulumi.IntOutput `pulumi:"messageType"`
	// This attribute is used to set the read-write mode for the topic. Read [Request parameters](https://www.alibabacloud.com/help/doc-detail/56880.html) for further details.
	Perm pulumi.IntPtrOutput `pulumi:"perm"`
	// This attribute is a concise description of topic. The length cannot exceed 128.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Replaced by `topicName` after version 1.97.0.
	//
	// Deprecated: Field 'topic' has been deprecated from version 1.97.0. Use 'topic_name' instead.
	Topic pulumi.StringOutput `pulumi:"topic"`
	// Name of the topic. Two topics on a single instance cannot have the same name and the name cannot start with 'GID' or 'CID'. The length cannot exceed 64 characters.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.MessageType == nil {
		return nil, errors.New("invalid value for required argument 'MessageType'")
	}
	var resource Topic
	err := ctx.RegisterResource("alicloud:rocketmq/topic:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("alicloud:rocketmq/topic:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
	// ID of the ONS Instance that owns the topics.
	InstanceId *string `pulumi:"instanceId"`
	// The type of the message. Read [Ons Topic Create](https://www.alibabacloud.com/help/doc-detail/29591.html) for further details.
	MessageType *int `pulumi:"messageType"`
	// This attribute is used to set the read-write mode for the topic. Read [Request parameters](https://www.alibabacloud.com/help/doc-detail/56880.html) for further details.
	Perm *int `pulumi:"perm"`
	// This attribute is a concise description of topic. The length cannot exceed 128.
	Remark *string `pulumi:"remark"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
	// Replaced by `topicName` after version 1.97.0.
	//
	// Deprecated: Field 'topic' has been deprecated from version 1.97.0. Use 'topic_name' instead.
	Topic *string `pulumi:"topic"`
	// Name of the topic. Two topics on a single instance cannot have the same name and the name cannot start with 'GID' or 'CID'. The length cannot exceed 64 characters.
	TopicName *string `pulumi:"topicName"`
}

type TopicState struct {
	// ID of the ONS Instance that owns the topics.
	InstanceId pulumi.StringPtrInput
	// The type of the message. Read [Ons Topic Create](https://www.alibabacloud.com/help/doc-detail/29591.html) for further details.
	MessageType pulumi.IntPtrInput
	// This attribute is used to set the read-write mode for the topic. Read [Request parameters](https://www.alibabacloud.com/help/doc-detail/56880.html) for further details.
	Perm pulumi.IntPtrInput
	// This attribute is a concise description of topic. The length cannot exceed 128.
	Remark pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapInput
	// Replaced by `topicName` after version 1.97.0.
	//
	// Deprecated: Field 'topic' has been deprecated from version 1.97.0. Use 'topic_name' instead.
	Topic pulumi.StringPtrInput
	// Name of the topic. Two topics on a single instance cannot have the same name and the name cannot start with 'GID' or 'CID'. The length cannot exceed 64 characters.
	TopicName pulumi.StringPtrInput
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// ID of the ONS Instance that owns the topics.
	InstanceId string `pulumi:"instanceId"`
	// The type of the message. Read [Ons Topic Create](https://www.alibabacloud.com/help/doc-detail/29591.html) for further details.
	MessageType int `pulumi:"messageType"`
	// This attribute is used to set the read-write mode for the topic. Read [Request parameters](https://www.alibabacloud.com/help/doc-detail/56880.html) for further details.
	Perm *int `pulumi:"perm"`
	// This attribute is a concise description of topic. The length cannot exceed 128.
	Remark *string `pulumi:"remark"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
	// Replaced by `topicName` after version 1.97.0.
	//
	// Deprecated: Field 'topic' has been deprecated from version 1.97.0. Use 'topic_name' instead.
	Topic *string `pulumi:"topic"`
	// Name of the topic. Two topics on a single instance cannot have the same name and the name cannot start with 'GID' or 'CID'. The length cannot exceed 64 characters.
	TopicName *string `pulumi:"topicName"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// ID of the ONS Instance that owns the topics.
	InstanceId pulumi.StringInput
	// The type of the message. Read [Ons Topic Create](https://www.alibabacloud.com/help/doc-detail/29591.html) for further details.
	MessageType pulumi.IntInput
	// This attribute is used to set the read-write mode for the topic. Read [Request parameters](https://www.alibabacloud.com/help/doc-detail/56880.html) for further details.
	Perm pulumi.IntPtrInput
	// This attribute is a concise description of topic. The length cannot exceed 128.
	Remark pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapInput
	// Replaced by `topicName` after version 1.97.0.
	//
	// Deprecated: Field 'topic' has been deprecated from version 1.97.0. Use 'topic_name' instead.
	Topic pulumi.StringPtrInput
	// Name of the topic. Two topics on a single instance cannot have the same name and the name cannot start with 'GID' or 'CID'. The length cannot exceed 64 characters.
	TopicName pulumi.StringPtrInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}

type TopicInput interface {
	pulumi.Input

	ToTopicOutput() TopicOutput
	ToTopicOutputWithContext(ctx context.Context) TopicOutput
}

func (*Topic) ElementType() reflect.Type {
	return reflect.TypeOf((*Topic)(nil))
}

func (i *Topic) ToTopicOutput() TopicOutput {
	return i.ToTopicOutputWithContext(context.Background())
}

func (i *Topic) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput)
}

func (i *Topic) ToTopicPtrOutput() TopicPtrOutput {
	return i.ToTopicPtrOutputWithContext(context.Background())
}

func (i *Topic) ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicPtrOutput)
}

type TopicPtrInput interface {
	pulumi.Input

	ToTopicPtrOutput() TopicPtrOutput
	ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput
}

type topicPtrType TopicArgs

func (*topicPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil))
}

func (i *topicPtrType) ToTopicPtrOutput() TopicPtrOutput {
	return i.ToTopicPtrOutputWithContext(context.Background())
}

func (i *topicPtrType) ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicPtrOutput)
}

// TopicArrayInput is an input type that accepts TopicArray and TopicArrayOutput values.
// You can construct a concrete instance of `TopicArrayInput` via:
//
//          TopicArray{ TopicArgs{...} }
type TopicArrayInput interface {
	pulumi.Input

	ToTopicArrayOutput() TopicArrayOutput
	ToTopicArrayOutputWithContext(context.Context) TopicArrayOutput
}

type TopicArray []TopicInput

func (TopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Topic)(nil)).Elem()
}

func (i TopicArray) ToTopicArrayOutput() TopicArrayOutput {
	return i.ToTopicArrayOutputWithContext(context.Background())
}

func (i TopicArray) ToTopicArrayOutputWithContext(ctx context.Context) TopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicArrayOutput)
}

// TopicMapInput is an input type that accepts TopicMap and TopicMapOutput values.
// You can construct a concrete instance of `TopicMapInput` via:
//
//          TopicMap{ "key": TopicArgs{...} }
type TopicMapInput interface {
	pulumi.Input

	ToTopicMapOutput() TopicMapOutput
	ToTopicMapOutputWithContext(context.Context) TopicMapOutput
}

type TopicMap map[string]TopicInput

func (TopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Topic)(nil)).Elem()
}

func (i TopicMap) ToTopicMapOutput() TopicMapOutput {
	return i.ToTopicMapOutputWithContext(context.Background())
}

func (i TopicMap) ToTopicMapOutputWithContext(ctx context.Context) TopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicMapOutput)
}

type TopicOutput struct{ *pulumi.OutputState }

func (TopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Topic)(nil))
}

func (o TopicOutput) ToTopicOutput() TopicOutput {
	return o
}

func (o TopicOutput) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return o
}

func (o TopicOutput) ToTopicPtrOutput() TopicPtrOutput {
	return o.ToTopicPtrOutputWithContext(context.Background())
}

func (o TopicOutput) ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Topic) *Topic {
		return &v
	}).(TopicPtrOutput)
}

type TopicPtrOutput struct{ *pulumi.OutputState }

func (TopicPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil))
}

func (o TopicPtrOutput) ToTopicPtrOutput() TopicPtrOutput {
	return o
}

func (o TopicPtrOutput) ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput {
	return o
}

func (o TopicPtrOutput) Elem() TopicOutput {
	return o.ApplyT(func(v *Topic) Topic {
		if v != nil {
			return *v
		}
		var ret Topic
		return ret
	}).(TopicOutput)
}

type TopicArrayOutput struct{ *pulumi.OutputState }

func (TopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Topic)(nil))
}

func (o TopicArrayOutput) ToTopicArrayOutput() TopicArrayOutput {
	return o
}

func (o TopicArrayOutput) ToTopicArrayOutputWithContext(ctx context.Context) TopicArrayOutput {
	return o
}

func (o TopicArrayOutput) Index(i pulumi.IntInput) TopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Topic {
		return vs[0].([]Topic)[vs[1].(int)]
	}).(TopicOutput)
}

type TopicMapOutput struct{ *pulumi.OutputState }

func (TopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Topic)(nil))
}

func (o TopicMapOutput) ToTopicMapOutput() TopicMapOutput {
	return o
}

func (o TopicMapOutput) ToTopicMapOutputWithContext(ctx context.Context) TopicMapOutput {
	return o
}

func (o TopicMapOutput) MapIndex(k pulumi.StringInput) TopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Topic {
		return vs[0].(map[string]Topic)[vs[1].(string)]
	}).(TopicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TopicInput)(nil)).Elem(), &Topic{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicPtrInput)(nil)).Elem(), &Topic{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicArrayInput)(nil)).Elem(), TopicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicMapInput)(nil)).Elem(), TopicMap{})
	pulumi.RegisterOutputType(TopicOutput{})
	pulumi.RegisterOutputType(TopicPtrOutput{})
	pulumi.RegisterOutputType(TopicArrayOutput{})
	pulumi.RegisterOutputType(TopicMapOutput{})
}
