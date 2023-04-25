// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amqp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a RabbitMQ (AMQP) Queue resource.
//
// For information about RabbitMQ (AMQP) Queue and how to use it, see [What is Queue](https://www.alibabacloud.com/help/doc-detail/101631.htm).
//
// > **NOTE:** Available in v1.127.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/amqp"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleVirtualHost, err := amqp.NewVirtualHost(ctx, "exampleVirtualHost", &amqp.VirtualHostArgs{
//				InstanceId:      pulumi.String("amqp-abc12345"),
//				VirtualHostName: pulumi.String("my-VirtualHost"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = amqp.NewQueue(ctx, "exampleQueue", &amqp.QueueArgs{
//				InstanceId:      exampleVirtualHost.InstanceId,
//				QueueName:       pulumi.String("my-Queue"),
//				VirtualHostName: exampleVirtualHost.VirtualHostName,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// RabbitMQ (AMQP) Queue can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:amqp/queue:Queue example <instance_id>:<virtual_host_name>:<queue_name>
//
// ```
type Queue struct {
	pulumi.CustomResourceState

	// Specifies whether the Auto Delete attribute is configured. Valid values:
	// * true: The Auto Delete attribute is configured. The queue is automatically deleted after the last subscription from consumers to this queue is canceled.
	// * false: The Auto Delete attribute is not configured.
	AutoDeleteState pulumi.BoolPtrOutput `pulumi:"autoDeleteState"`
	// The validity period after which the queue is automatically deleted.
	// If the queue is not accessed within a specified period of time, it is automatically deleted.
	AutoExpireState pulumi.StringPtrOutput `pulumi:"autoExpireState"`
	// The dead-letter exchange. A dead-letter exchange is used to receive rejected messages.
	// If a consumer rejects a message that cannot be retried, this message is routed to a specified dead-letter exchange.
	// Then, the dead-letter exchange routes the message to the queue that is bound to the dead-letter exchange.
	DeadLetterExchange pulumi.StringPtrOutput `pulumi:"deadLetterExchange"`
	// The dead letter routing key.
	DeadLetterRoutingKey pulumi.StringPtrOutput `pulumi:"deadLetterRoutingKey"`
	// Specifies whether the queue is an exclusive queue. Valid values:
	// * true: The queue is an exclusive queue. It can be used only for the connection that declares the exclusive queue. After the connection is closed, the exclusive queue is automatically deleted.
	// * false: The queue is not an exclusive queue.
	ExclusiveState pulumi.BoolPtrOutput `pulumi:"exclusiveState"`
	// The ID of the instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The maximum number of messages that can be stored in the queue.
	// If this threshold is exceeded, the earliest messages that are routed to the queue are discarded.
	MaxLength pulumi.StringPtrOutput `pulumi:"maxLength"`
	// The highest priority supported by the queue. This parameter is set to a positive integer.
	// Valid values: 0 to 255. Recommended values: 1 to 10
	MaximumPriority pulumi.IntPtrOutput `pulumi:"maximumPriority"`
	// The message TTL of the queue.
	// If the retention period of a message in the queue exceeds the message TTL of the queue, the message expires.
	// Message TTL must be set to a non-negative integer, in milliseconds.
	// For example, if the message TTL of the queue is 1000, messages survive for at most 1 second in the queue.
	MessageTtl pulumi.StringPtrOutput `pulumi:"messageTtl"`
	// The name of the queue.
	// The queue name must be 1 to 255 characters in length, and can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
	QueueName pulumi.StringOutput `pulumi:"queueName"`
	// The name of the virtual host.
	VirtualHostName pulumi.StringOutput `pulumi:"virtualHostName"`
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.QueueName == nil {
		return nil, errors.New("invalid value for required argument 'QueueName'")
	}
	if args.VirtualHostName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHostName'")
	}
	var resource Queue
	err := ctx.RegisterResource("alicloud:amqp/queue:Queue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueState, opts ...pulumi.ResourceOption) (*Queue, error) {
	var resource Queue
	err := ctx.ReadResource("alicloud:amqp/queue:Queue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Queue resources.
type queueState struct {
	// Specifies whether the Auto Delete attribute is configured. Valid values:
	// * true: The Auto Delete attribute is configured. The queue is automatically deleted after the last subscription from consumers to this queue is canceled.
	// * false: The Auto Delete attribute is not configured.
	AutoDeleteState *bool `pulumi:"autoDeleteState"`
	// The validity period after which the queue is automatically deleted.
	// If the queue is not accessed within a specified period of time, it is automatically deleted.
	AutoExpireState *string `pulumi:"autoExpireState"`
	// The dead-letter exchange. A dead-letter exchange is used to receive rejected messages.
	// If a consumer rejects a message that cannot be retried, this message is routed to a specified dead-letter exchange.
	// Then, the dead-letter exchange routes the message to the queue that is bound to the dead-letter exchange.
	DeadLetterExchange *string `pulumi:"deadLetterExchange"`
	// The dead letter routing key.
	DeadLetterRoutingKey *string `pulumi:"deadLetterRoutingKey"`
	// Specifies whether the queue is an exclusive queue. Valid values:
	// * true: The queue is an exclusive queue. It can be used only for the connection that declares the exclusive queue. After the connection is closed, the exclusive queue is automatically deleted.
	// * false: The queue is not an exclusive queue.
	ExclusiveState *bool `pulumi:"exclusiveState"`
	// The ID of the instance.
	InstanceId *string `pulumi:"instanceId"`
	// The maximum number of messages that can be stored in the queue.
	// If this threshold is exceeded, the earliest messages that are routed to the queue are discarded.
	MaxLength *string `pulumi:"maxLength"`
	// The highest priority supported by the queue. This parameter is set to a positive integer.
	// Valid values: 0 to 255. Recommended values: 1 to 10
	MaximumPriority *int `pulumi:"maximumPriority"`
	// The message TTL of the queue.
	// If the retention period of a message in the queue exceeds the message TTL of the queue, the message expires.
	// Message TTL must be set to a non-negative integer, in milliseconds.
	// For example, if the message TTL of the queue is 1000, messages survive for at most 1 second in the queue.
	MessageTtl *string `pulumi:"messageTtl"`
	// The name of the queue.
	// The queue name must be 1 to 255 characters in length, and can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
	QueueName *string `pulumi:"queueName"`
	// The name of the virtual host.
	VirtualHostName *string `pulumi:"virtualHostName"`
}

type QueueState struct {
	// Specifies whether the Auto Delete attribute is configured. Valid values:
	// * true: The Auto Delete attribute is configured. The queue is automatically deleted after the last subscription from consumers to this queue is canceled.
	// * false: The Auto Delete attribute is not configured.
	AutoDeleteState pulumi.BoolPtrInput
	// The validity period after which the queue is automatically deleted.
	// If the queue is not accessed within a specified period of time, it is automatically deleted.
	AutoExpireState pulumi.StringPtrInput
	// The dead-letter exchange. A dead-letter exchange is used to receive rejected messages.
	// If a consumer rejects a message that cannot be retried, this message is routed to a specified dead-letter exchange.
	// Then, the dead-letter exchange routes the message to the queue that is bound to the dead-letter exchange.
	DeadLetterExchange pulumi.StringPtrInput
	// The dead letter routing key.
	DeadLetterRoutingKey pulumi.StringPtrInput
	// Specifies whether the queue is an exclusive queue. Valid values:
	// * true: The queue is an exclusive queue. It can be used only for the connection that declares the exclusive queue. After the connection is closed, the exclusive queue is automatically deleted.
	// * false: The queue is not an exclusive queue.
	ExclusiveState pulumi.BoolPtrInput
	// The ID of the instance.
	InstanceId pulumi.StringPtrInput
	// The maximum number of messages that can be stored in the queue.
	// If this threshold is exceeded, the earliest messages that are routed to the queue are discarded.
	MaxLength pulumi.StringPtrInput
	// The highest priority supported by the queue. This parameter is set to a positive integer.
	// Valid values: 0 to 255. Recommended values: 1 to 10
	MaximumPriority pulumi.IntPtrInput
	// The message TTL of the queue.
	// If the retention period of a message in the queue exceeds the message TTL of the queue, the message expires.
	// Message TTL must be set to a non-negative integer, in milliseconds.
	// For example, if the message TTL of the queue is 1000, messages survive for at most 1 second in the queue.
	MessageTtl pulumi.StringPtrInput
	// The name of the queue.
	// The queue name must be 1 to 255 characters in length, and can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
	QueueName pulumi.StringPtrInput
	// The name of the virtual host.
	VirtualHostName pulumi.StringPtrInput
}

func (QueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueState)(nil)).Elem()
}

type queueArgs struct {
	// Specifies whether the Auto Delete attribute is configured. Valid values:
	// * true: The Auto Delete attribute is configured. The queue is automatically deleted after the last subscription from consumers to this queue is canceled.
	// * false: The Auto Delete attribute is not configured.
	AutoDeleteState *bool `pulumi:"autoDeleteState"`
	// The validity period after which the queue is automatically deleted.
	// If the queue is not accessed within a specified period of time, it is automatically deleted.
	AutoExpireState *string `pulumi:"autoExpireState"`
	// The dead-letter exchange. A dead-letter exchange is used to receive rejected messages.
	// If a consumer rejects a message that cannot be retried, this message is routed to a specified dead-letter exchange.
	// Then, the dead-letter exchange routes the message to the queue that is bound to the dead-letter exchange.
	DeadLetterExchange *string `pulumi:"deadLetterExchange"`
	// The dead letter routing key.
	DeadLetterRoutingKey *string `pulumi:"deadLetterRoutingKey"`
	// Specifies whether the queue is an exclusive queue. Valid values:
	// * true: The queue is an exclusive queue. It can be used only for the connection that declares the exclusive queue. After the connection is closed, the exclusive queue is automatically deleted.
	// * false: The queue is not an exclusive queue.
	ExclusiveState *bool `pulumi:"exclusiveState"`
	// The ID of the instance.
	InstanceId string `pulumi:"instanceId"`
	// The maximum number of messages that can be stored in the queue.
	// If this threshold is exceeded, the earliest messages that are routed to the queue are discarded.
	MaxLength *string `pulumi:"maxLength"`
	// The highest priority supported by the queue. This parameter is set to a positive integer.
	// Valid values: 0 to 255. Recommended values: 1 to 10
	MaximumPriority *int `pulumi:"maximumPriority"`
	// The message TTL of the queue.
	// If the retention period of a message in the queue exceeds the message TTL of the queue, the message expires.
	// Message TTL must be set to a non-negative integer, in milliseconds.
	// For example, if the message TTL of the queue is 1000, messages survive for at most 1 second in the queue.
	MessageTtl *string `pulumi:"messageTtl"`
	// The name of the queue.
	// The queue name must be 1 to 255 characters in length, and can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
	QueueName string `pulumi:"queueName"`
	// The name of the virtual host.
	VirtualHostName string `pulumi:"virtualHostName"`
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// Specifies whether the Auto Delete attribute is configured. Valid values:
	// * true: The Auto Delete attribute is configured. The queue is automatically deleted after the last subscription from consumers to this queue is canceled.
	// * false: The Auto Delete attribute is not configured.
	AutoDeleteState pulumi.BoolPtrInput
	// The validity period after which the queue is automatically deleted.
	// If the queue is not accessed within a specified period of time, it is automatically deleted.
	AutoExpireState pulumi.StringPtrInput
	// The dead-letter exchange. A dead-letter exchange is used to receive rejected messages.
	// If a consumer rejects a message that cannot be retried, this message is routed to a specified dead-letter exchange.
	// Then, the dead-letter exchange routes the message to the queue that is bound to the dead-letter exchange.
	DeadLetterExchange pulumi.StringPtrInput
	// The dead letter routing key.
	DeadLetterRoutingKey pulumi.StringPtrInput
	// Specifies whether the queue is an exclusive queue. Valid values:
	// * true: The queue is an exclusive queue. It can be used only for the connection that declares the exclusive queue. After the connection is closed, the exclusive queue is automatically deleted.
	// * false: The queue is not an exclusive queue.
	ExclusiveState pulumi.BoolPtrInput
	// The ID of the instance.
	InstanceId pulumi.StringInput
	// The maximum number of messages that can be stored in the queue.
	// If this threshold is exceeded, the earliest messages that are routed to the queue are discarded.
	MaxLength pulumi.StringPtrInput
	// The highest priority supported by the queue. This parameter is set to a positive integer.
	// Valid values: 0 to 255. Recommended values: 1 to 10
	MaximumPriority pulumi.IntPtrInput
	// The message TTL of the queue.
	// If the retention period of a message in the queue exceeds the message TTL of the queue, the message expires.
	// Message TTL must be set to a non-negative integer, in milliseconds.
	// For example, if the message TTL of the queue is 1000, messages survive for at most 1 second in the queue.
	MessageTtl pulumi.StringPtrInput
	// The name of the queue.
	// The queue name must be 1 to 255 characters in length, and can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
	QueueName pulumi.StringInput
	// The name of the virtual host.
	VirtualHostName pulumi.StringInput
}

func (QueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueArgs)(nil)).Elem()
}

type QueueInput interface {
	pulumi.Input

	ToQueueOutput() QueueOutput
	ToQueueOutputWithContext(ctx context.Context) QueueOutput
}

func (*Queue) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil)).Elem()
}

func (i *Queue) ToQueueOutput() QueueOutput {
	return i.ToQueueOutputWithContext(context.Background())
}

func (i *Queue) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueOutput)
}

// QueueArrayInput is an input type that accepts QueueArray and QueueArrayOutput values.
// You can construct a concrete instance of `QueueArrayInput` via:
//
//	QueueArray{ QueueArgs{...} }
type QueueArrayInput interface {
	pulumi.Input

	ToQueueArrayOutput() QueueArrayOutput
	ToQueueArrayOutputWithContext(context.Context) QueueArrayOutput
}

type QueueArray []QueueInput

func (QueueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Queue)(nil)).Elem()
}

func (i QueueArray) ToQueueArrayOutput() QueueArrayOutput {
	return i.ToQueueArrayOutputWithContext(context.Background())
}

func (i QueueArray) ToQueueArrayOutputWithContext(ctx context.Context) QueueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueArrayOutput)
}

// QueueMapInput is an input type that accepts QueueMap and QueueMapOutput values.
// You can construct a concrete instance of `QueueMapInput` via:
//
//	QueueMap{ "key": QueueArgs{...} }
type QueueMapInput interface {
	pulumi.Input

	ToQueueMapOutput() QueueMapOutput
	ToQueueMapOutputWithContext(context.Context) QueueMapOutput
}

type QueueMap map[string]QueueInput

func (QueueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Queue)(nil)).Elem()
}

func (i QueueMap) ToQueueMapOutput() QueueMapOutput {
	return i.ToQueueMapOutputWithContext(context.Background())
}

func (i QueueMap) ToQueueMapOutputWithContext(ctx context.Context) QueueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueMapOutput)
}

type QueueOutput struct{ *pulumi.OutputState }

func (QueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil)).Elem()
}

func (o QueueOutput) ToQueueOutput() QueueOutput {
	return o
}

func (o QueueOutput) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return o
}

// Specifies whether the Auto Delete attribute is configured. Valid values:
// * true: The Auto Delete attribute is configured. The queue is automatically deleted after the last subscription from consumers to this queue is canceled.
// * false: The Auto Delete attribute is not configured.
func (o QueueOutput) AutoDeleteState() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.BoolPtrOutput { return v.AutoDeleteState }).(pulumi.BoolPtrOutput)
}

// The validity period after which the queue is automatically deleted.
// If the queue is not accessed within a specified period of time, it is automatically deleted.
func (o QueueOutput) AutoExpireState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.AutoExpireState }).(pulumi.StringPtrOutput)
}

// The dead-letter exchange. A dead-letter exchange is used to receive rejected messages.
// If a consumer rejects a message that cannot be retried, this message is routed to a specified dead-letter exchange.
// Then, the dead-letter exchange routes the message to the queue that is bound to the dead-letter exchange.
func (o QueueOutput) DeadLetterExchange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.DeadLetterExchange }).(pulumi.StringPtrOutput)
}

// The dead letter routing key.
func (o QueueOutput) DeadLetterRoutingKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.DeadLetterRoutingKey }).(pulumi.StringPtrOutput)
}

// Specifies whether the queue is an exclusive queue. Valid values:
// * true: The queue is an exclusive queue. It can be used only for the connection that declares the exclusive queue. After the connection is closed, the exclusive queue is automatically deleted.
// * false: The queue is not an exclusive queue.
func (o QueueOutput) ExclusiveState() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.BoolPtrOutput { return v.ExclusiveState }).(pulumi.BoolPtrOutput)
}

// The ID of the instance.
func (o QueueOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The maximum number of messages that can be stored in the queue.
// If this threshold is exceeded, the earliest messages that are routed to the queue are discarded.
func (o QueueOutput) MaxLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.MaxLength }).(pulumi.StringPtrOutput)
}

// The highest priority supported by the queue. This parameter is set to a positive integer.
// Valid values: 0 to 255. Recommended values: 1 to 10
func (o QueueOutput) MaximumPriority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.IntPtrOutput { return v.MaximumPriority }).(pulumi.IntPtrOutput)
}

// The message TTL of the queue.
// If the retention period of a message in the queue exceeds the message TTL of the queue, the message expires.
// Message TTL must be set to a non-negative integer, in milliseconds.
// For example, if the message TTL of the queue is 1000, messages survive for at most 1 second in the queue.
func (o QueueOutput) MessageTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.MessageTtl }).(pulumi.StringPtrOutput)
}

// The name of the queue.
// The queue name must be 1 to 255 characters in length, and can contain only letters, digits, hyphens (-), underscores (_), periods (.), and at signs (@).
func (o QueueOutput) QueueName() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.QueueName }).(pulumi.StringOutput)
}

// The name of the virtual host.
func (o QueueOutput) VirtualHostName() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.VirtualHostName }).(pulumi.StringOutput)
}

type QueueArrayOutput struct{ *pulumi.OutputState }

func (QueueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Queue)(nil)).Elem()
}

func (o QueueArrayOutput) ToQueueArrayOutput() QueueArrayOutput {
	return o
}

func (o QueueArrayOutput) ToQueueArrayOutputWithContext(ctx context.Context) QueueArrayOutput {
	return o
}

func (o QueueArrayOutput) Index(i pulumi.IntInput) QueueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Queue {
		return vs[0].([]*Queue)[vs[1].(int)]
	}).(QueueOutput)
}

type QueueMapOutput struct{ *pulumi.OutputState }

func (QueueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Queue)(nil)).Elem()
}

func (o QueueMapOutput) ToQueueMapOutput() QueueMapOutput {
	return o
}

func (o QueueMapOutput) ToQueueMapOutputWithContext(ctx context.Context) QueueMapOutput {
	return o
}

func (o QueueMapOutput) MapIndex(k pulumi.StringInput) QueueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Queue {
		return vs[0].(map[string]*Queue)[vs[1].(string)]
	}).(QueueOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueueInput)(nil)).Elem(), &Queue{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueueArrayInput)(nil)).Elem(), QueueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueueMapInput)(nil)).Elem(), QueueMap{})
	pulumi.RegisterOutputType(QueueOutput{})
	pulumi.RegisterOutputType(QueueArrayOutput{})
	pulumi.RegisterOutputType(QueueMapOutput{})
}
