// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// Function Compute trigger can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:fc/trigger:Trigger foo my-fc-service:hello-world:hello-trigger
// ```
type Trigger struct {
	pulumi.CustomResourceState

	// The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
	Config pulumi.StringPtrOutput `pulumi:"config"`
	// The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
	ConfigMns pulumi.StringPtrOutput `pulumi:"configMns"`
	// The Function Compute function name.
	Function pulumi.StringOutput `pulumi:"function"`
	// The date this resource was last modified.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
	Name pulumi.StringOutput `pulumi:"name"`
	// Setting a prefix to get a only trigger name. It is conflict with "name".
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	Role pulumi.StringPtrOutput `pulumi:"role"`
	// The Function Compute service name.
	Service pulumi.StringOutput `pulumi:"service"`
	// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	SourceArn pulumi.StringPtrOutput `pulumi:"sourceArn"`
	// The Function Compute trigger ID.
	TriggerId pulumi.StringOutput `pulumi:"triggerId"`
	// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents"].
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTrigger registers a new resource with the given unique name, arguments, and options.
func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Function == nil {
		return nil, errors.New("invalid value for required argument 'Function'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Trigger
	err := ctx.RegisterResource("alicloud:fc/trigger:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrigger gets an existing Trigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("alicloud:fc/trigger:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trigger resources.
type triggerState struct {
	// The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
	Config *string `pulumi:"config"`
	// The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
	ConfigMns *string `pulumi:"configMns"`
	// The Function Compute function name.
	Function *string `pulumi:"function"`
	// The date this resource was last modified.
	LastModified *string `pulumi:"lastModified"`
	// The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
	Name *string `pulumi:"name"`
	// Setting a prefix to get a only trigger name. It is conflict with "name".
	NamePrefix *string `pulumi:"namePrefix"`
	// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	Role *string `pulumi:"role"`
	// The Function Compute service name.
	Service *string `pulumi:"service"`
	// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	SourceArn *string `pulumi:"sourceArn"`
	// The Function Compute trigger ID.
	TriggerId *string `pulumi:"triggerId"`
	// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents"].
	Type *string `pulumi:"type"`
}

type TriggerState struct {
	// The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
	Config pulumi.StringPtrInput
	// The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
	ConfigMns pulumi.StringPtrInput
	// The Function Compute function name.
	Function pulumi.StringPtrInput
	// The date this resource was last modified.
	LastModified pulumi.StringPtrInput
	// The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
	Name pulumi.StringPtrInput
	// Setting a prefix to get a only trigger name. It is conflict with "name".
	NamePrefix pulumi.StringPtrInput
	// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	Role pulumi.StringPtrInput
	// The Function Compute service name.
	Service pulumi.StringPtrInput
	// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	SourceArn pulumi.StringPtrInput
	// The Function Compute trigger ID.
	TriggerId pulumi.StringPtrInput
	// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents"].
	Type pulumi.StringPtrInput
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	// The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
	Config *string `pulumi:"config"`
	// The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
	ConfigMns *string `pulumi:"configMns"`
	// The Function Compute function name.
	Function string `pulumi:"function"`
	// The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
	Name *string `pulumi:"name"`
	// Setting a prefix to get a only trigger name. It is conflict with "name".
	NamePrefix *string `pulumi:"namePrefix"`
	// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	Role *string `pulumi:"role"`
	// The Function Compute service name.
	Service string `pulumi:"service"`
	// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	SourceArn *string `pulumi:"sourceArn"`
	// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents"].
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	// The config of Function Compute trigger.It is valid when `type` is not "mnsTopic".See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.
	Config pulumi.StringPtrInput
	// The config of Function Compute trigger when the type is "mnsTopic".It is conflict with `config`.
	ConfigMns pulumi.StringPtrInput
	// The Function Compute function name.
	Function pulumi.StringInput
	// The Function Compute trigger name. It is the only in one service and is conflict with "namePrefix".
	Name pulumi.StringPtrInput
	// Setting a prefix to get a only trigger name. It is conflict with "name".
	NamePrefix pulumi.StringPtrInput
	// RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	Role pulumi.StringPtrInput
	// The Function Compute service name.
	Service pulumi.StringInput
	// Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.
	SourceArn pulumi.StringPtrInput
	// The Type of the trigger. Valid values: ["oss", "log", "timer", "http", "mnsTopic", "cdnEvents"].
	Type pulumi.StringInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}

type TriggerInput interface {
	pulumi.Input

	ToTriggerOutput() TriggerOutput
	ToTriggerOutputWithContext(ctx context.Context) TriggerOutput
}

func (Trigger) ElementType() reflect.Type {
	return reflect.TypeOf((*Trigger)(nil)).Elem()
}

func (i Trigger) ToTriggerOutput() TriggerOutput {
	return i.ToTriggerOutputWithContext(context.Background())
}

func (i Trigger) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerOutput)
}

type TriggerOutput struct {
	*pulumi.OutputState
}

func (TriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerOutput)(nil)).Elem()
}

func (o TriggerOutput) ToTriggerOutput() TriggerOutput {
	return o
}

func (o TriggerOutput) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TriggerOutput{})
}
