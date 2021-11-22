// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// MNS Topic subscription can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:mns/topicSubscription:TopicSubscription subscription tf-example-mnstopic:tf-example-mnstopic-sub
// ```
type TopicSubscription struct {
	pulumi.CustomResourceState

	// The endpoint has three format. Available values format:
	// - HTTP Format: http://xxx.com/xxx
	// - Queue Format: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
	// - Email Format: mail:directmail:{MailAddress}
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The length should be shorter than 16.
	FilterTag pulumi.StringPtrOutput `pulumi:"filterTag"`
	// Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: 'SIMPLIFIED', 'XML' and 'JSON'. Default to 'SIMPLIFIED'.
	NotifyContentFormat pulumi.StringPtrOutput `pulumi:"notifyContentFormat"`
	// The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. the attribute has two value EXPONENTIAL_DECAY_RETR or BACKOFF_RETRY. Default value to BACKOFF_RETRY .
	NotifyStrategy pulumi.StringPtrOutput `pulumi:"notifyStrategy"`
	// The topic which The subscription belongs to was named with the name.A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
}

// NewTopicSubscription registers a new resource with the given unique name, arguments, and options.
func NewTopicSubscription(ctx *pulumi.Context,
	name string, args *TopicSubscriptionArgs, opts ...pulumi.ResourceOption) (*TopicSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endpoint == nil {
		return nil, errors.New("invalid value for required argument 'Endpoint'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	var resource TopicSubscription
	err := ctx.RegisterResource("alicloud:mns/topicSubscription:TopicSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopicSubscription gets an existing TopicSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicSubscriptionState, opts ...pulumi.ResourceOption) (*TopicSubscription, error) {
	var resource TopicSubscription
	err := ctx.ReadResource("alicloud:mns/topicSubscription:TopicSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TopicSubscription resources.
type topicSubscriptionState struct {
	// The endpoint has three format. Available values format:
	// - HTTP Format: http://xxx.com/xxx
	// - Queue Format: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
	// - Email Format: mail:directmail:{MailAddress}
	Endpoint *string `pulumi:"endpoint"`
	// The length should be shorter than 16.
	FilterTag *string `pulumi:"filterTag"`
	// Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
	Name *string `pulumi:"name"`
	// The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: 'SIMPLIFIED', 'XML' and 'JSON'. Default to 'SIMPLIFIED'.
	NotifyContentFormat *string `pulumi:"notifyContentFormat"`
	// The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. the attribute has two value EXPONENTIAL_DECAY_RETR or BACKOFF_RETRY. Default value to BACKOFF_RETRY .
	NotifyStrategy *string `pulumi:"notifyStrategy"`
	// The topic which The subscription belongs to was named with the name.A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
	TopicName *string `pulumi:"topicName"`
}

type TopicSubscriptionState struct {
	// The endpoint has three format. Available values format:
	// - HTTP Format: http://xxx.com/xxx
	// - Queue Format: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
	// - Email Format: mail:directmail:{MailAddress}
	Endpoint pulumi.StringPtrInput
	// The length should be shorter than 16.
	FilterTag pulumi.StringPtrInput
	// Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
	Name pulumi.StringPtrInput
	// The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: 'SIMPLIFIED', 'XML' and 'JSON'. Default to 'SIMPLIFIED'.
	NotifyContentFormat pulumi.StringPtrInput
	// The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. the attribute has two value EXPONENTIAL_DECAY_RETR or BACKOFF_RETRY. Default value to BACKOFF_RETRY .
	NotifyStrategy pulumi.StringPtrInput
	// The topic which The subscription belongs to was named with the name.A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
	TopicName pulumi.StringPtrInput
}

func (TopicSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicSubscriptionState)(nil)).Elem()
}

type topicSubscriptionArgs struct {
	// The endpoint has three format. Available values format:
	// - HTTP Format: http://xxx.com/xxx
	// - Queue Format: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
	// - Email Format: mail:directmail:{MailAddress}
	Endpoint string `pulumi:"endpoint"`
	// The length should be shorter than 16.
	FilterTag *string `pulumi:"filterTag"`
	// Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
	Name *string `pulumi:"name"`
	// The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: 'SIMPLIFIED', 'XML' and 'JSON'. Default to 'SIMPLIFIED'.
	NotifyContentFormat *string `pulumi:"notifyContentFormat"`
	// The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. the attribute has two value EXPONENTIAL_DECAY_RETR or BACKOFF_RETRY. Default value to BACKOFF_RETRY .
	NotifyStrategy *string `pulumi:"notifyStrategy"`
	// The topic which The subscription belongs to was named with the name.A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a TopicSubscription resource.
type TopicSubscriptionArgs struct {
	// The endpoint has three format. Available values format:
	// - HTTP Format: http://xxx.com/xxx
	// - Queue Format: acs:mns:{REGION}:{AccountID}:queues/{QueueName}
	// - Email Format: mail:directmail:{MailAddress}
	Endpoint pulumi.StringInput
	// The length should be shorter than 16.
	FilterTag pulumi.StringPtrInput
	// Two topics subscription on a single account in the same topic cannot have the same name. A topic subscription name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
	Name pulumi.StringPtrInput
	// The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users. The valid values: 'SIMPLIFIED', 'XML' and 'JSON'. Default to 'SIMPLIFIED'.
	NotifyContentFormat pulumi.StringPtrInput
	// The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails. the attribute has two value EXPONENTIAL_DECAY_RETR or BACKOFF_RETRY. Default value to BACKOFF_RETRY .
	NotifyStrategy pulumi.StringPtrInput
	// The topic which The subscription belongs to was named with the name.A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.
	TopicName pulumi.StringInput
}

func (TopicSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicSubscriptionArgs)(nil)).Elem()
}

type TopicSubscriptionInput interface {
	pulumi.Input

	ToTopicSubscriptionOutput() TopicSubscriptionOutput
	ToTopicSubscriptionOutputWithContext(ctx context.Context) TopicSubscriptionOutput
}

func (*TopicSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicSubscription)(nil))
}

func (i *TopicSubscription) ToTopicSubscriptionOutput() TopicSubscriptionOutput {
	return i.ToTopicSubscriptionOutputWithContext(context.Background())
}

func (i *TopicSubscription) ToTopicSubscriptionOutputWithContext(ctx context.Context) TopicSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicSubscriptionOutput)
}

func (i *TopicSubscription) ToTopicSubscriptionPtrOutput() TopicSubscriptionPtrOutput {
	return i.ToTopicSubscriptionPtrOutputWithContext(context.Background())
}

func (i *TopicSubscription) ToTopicSubscriptionPtrOutputWithContext(ctx context.Context) TopicSubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicSubscriptionPtrOutput)
}

type TopicSubscriptionPtrInput interface {
	pulumi.Input

	ToTopicSubscriptionPtrOutput() TopicSubscriptionPtrOutput
	ToTopicSubscriptionPtrOutputWithContext(ctx context.Context) TopicSubscriptionPtrOutput
}

type topicSubscriptionPtrType TopicSubscriptionArgs

func (*topicSubscriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicSubscription)(nil))
}

func (i *topicSubscriptionPtrType) ToTopicSubscriptionPtrOutput() TopicSubscriptionPtrOutput {
	return i.ToTopicSubscriptionPtrOutputWithContext(context.Background())
}

func (i *topicSubscriptionPtrType) ToTopicSubscriptionPtrOutputWithContext(ctx context.Context) TopicSubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicSubscriptionPtrOutput)
}

// TopicSubscriptionArrayInput is an input type that accepts TopicSubscriptionArray and TopicSubscriptionArrayOutput values.
// You can construct a concrete instance of `TopicSubscriptionArrayInput` via:
//
//          TopicSubscriptionArray{ TopicSubscriptionArgs{...} }
type TopicSubscriptionArrayInput interface {
	pulumi.Input

	ToTopicSubscriptionArrayOutput() TopicSubscriptionArrayOutput
	ToTopicSubscriptionArrayOutputWithContext(context.Context) TopicSubscriptionArrayOutput
}

type TopicSubscriptionArray []TopicSubscriptionInput

func (TopicSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TopicSubscription)(nil)).Elem()
}

func (i TopicSubscriptionArray) ToTopicSubscriptionArrayOutput() TopicSubscriptionArrayOutput {
	return i.ToTopicSubscriptionArrayOutputWithContext(context.Background())
}

func (i TopicSubscriptionArray) ToTopicSubscriptionArrayOutputWithContext(ctx context.Context) TopicSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicSubscriptionArrayOutput)
}

// TopicSubscriptionMapInput is an input type that accepts TopicSubscriptionMap and TopicSubscriptionMapOutput values.
// You can construct a concrete instance of `TopicSubscriptionMapInput` via:
//
//          TopicSubscriptionMap{ "key": TopicSubscriptionArgs{...} }
type TopicSubscriptionMapInput interface {
	pulumi.Input

	ToTopicSubscriptionMapOutput() TopicSubscriptionMapOutput
	ToTopicSubscriptionMapOutputWithContext(context.Context) TopicSubscriptionMapOutput
}

type TopicSubscriptionMap map[string]TopicSubscriptionInput

func (TopicSubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TopicSubscription)(nil)).Elem()
}

func (i TopicSubscriptionMap) ToTopicSubscriptionMapOutput() TopicSubscriptionMapOutput {
	return i.ToTopicSubscriptionMapOutputWithContext(context.Background())
}

func (i TopicSubscriptionMap) ToTopicSubscriptionMapOutputWithContext(ctx context.Context) TopicSubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicSubscriptionMapOutput)
}

type TopicSubscriptionOutput struct{ *pulumi.OutputState }

func (TopicSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicSubscription)(nil))
}

func (o TopicSubscriptionOutput) ToTopicSubscriptionOutput() TopicSubscriptionOutput {
	return o
}

func (o TopicSubscriptionOutput) ToTopicSubscriptionOutputWithContext(ctx context.Context) TopicSubscriptionOutput {
	return o
}

func (o TopicSubscriptionOutput) ToTopicSubscriptionPtrOutput() TopicSubscriptionPtrOutput {
	return o.ToTopicSubscriptionPtrOutputWithContext(context.Background())
}

func (o TopicSubscriptionOutput) ToTopicSubscriptionPtrOutputWithContext(ctx context.Context) TopicSubscriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TopicSubscription) *TopicSubscription {
		return &v
	}).(TopicSubscriptionPtrOutput)
}

type TopicSubscriptionPtrOutput struct{ *pulumi.OutputState }

func (TopicSubscriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicSubscription)(nil))
}

func (o TopicSubscriptionPtrOutput) ToTopicSubscriptionPtrOutput() TopicSubscriptionPtrOutput {
	return o
}

func (o TopicSubscriptionPtrOutput) ToTopicSubscriptionPtrOutputWithContext(ctx context.Context) TopicSubscriptionPtrOutput {
	return o
}

func (o TopicSubscriptionPtrOutput) Elem() TopicSubscriptionOutput {
	return o.ApplyT(func(v *TopicSubscription) TopicSubscription {
		if v != nil {
			return *v
		}
		var ret TopicSubscription
		return ret
	}).(TopicSubscriptionOutput)
}

type TopicSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (TopicSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TopicSubscription)(nil))
}

func (o TopicSubscriptionArrayOutput) ToTopicSubscriptionArrayOutput() TopicSubscriptionArrayOutput {
	return o
}

func (o TopicSubscriptionArrayOutput) ToTopicSubscriptionArrayOutputWithContext(ctx context.Context) TopicSubscriptionArrayOutput {
	return o
}

func (o TopicSubscriptionArrayOutput) Index(i pulumi.IntInput) TopicSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TopicSubscription {
		return vs[0].([]TopicSubscription)[vs[1].(int)]
	}).(TopicSubscriptionOutput)
}

type TopicSubscriptionMapOutput struct{ *pulumi.OutputState }

func (TopicSubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TopicSubscription)(nil))
}

func (o TopicSubscriptionMapOutput) ToTopicSubscriptionMapOutput() TopicSubscriptionMapOutput {
	return o
}

func (o TopicSubscriptionMapOutput) ToTopicSubscriptionMapOutputWithContext(ctx context.Context) TopicSubscriptionMapOutput {
	return o
}

func (o TopicSubscriptionMapOutput) MapIndex(k pulumi.StringInput) TopicSubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TopicSubscription {
		return vs[0].(map[string]TopicSubscription)[vs[1].(string)]
	}).(TopicSubscriptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TopicSubscriptionInput)(nil)).Elem(), &TopicSubscription{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicSubscriptionPtrInput)(nil)).Elem(), &TopicSubscription{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicSubscriptionArrayInput)(nil)).Elem(), TopicSubscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicSubscriptionMapInput)(nil)).Elem(), TopicSubscriptionMap{})
	pulumi.RegisterOutputType(TopicSubscriptionOutput{})
	pulumi.RegisterOutputType(TopicSubscriptionPtrOutput{})
	pulumi.RegisterOutputType(TopicSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(TopicSubscriptionMapOutput{})
}
