// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package log

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Logtail access service is a log collection agent provided by Log Service.
// You can use Logtail to collect logs from servers such as Alibaba Cloud Elastic
// Compute Service (ECS) instances in real time in the Log Service console. [Refer to details](https://www.alibabacloud.com/help/doc-detail/29058.htm)
//
// This resource amis to attach one logtail configure to a machine group.
//
// > **NOTE:** One logtail configure can be attached to multiple machine groups and one machine group can attach several logtail configures.
//
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/log"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testProject, err := log.NewProject(ctx, "testProject", &log.ProjectArgs{
// 			Description: pulumi.String("create by terraform"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testStore, err := log.NewStore(ctx, "testStore", &log.StoreArgs{
// 			Project:            testProject.Name,
// 			RetentionPeriod:    pulumi.Int(3650),
// 			ShardCount:         pulumi.Int(3),
// 			AutoSplit:          pulumi.Bool(true),
// 			MaxSplitShardCount: pulumi.Int(60),
// 			AppendMeta:         pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testMachineGroup, err := log.NewMachineGroup(ctx, "testMachineGroup", &log.MachineGroupArgs{
// 			Project: testProject.Name,
// 			Topic:   pulumi.String("terraform"),
// 			IdentifyLists: pulumi.StringArray{
// 				pulumi.String("10.0.0.1"),
// 				pulumi.String("10.0.0.3"),
// 				pulumi.String("10.0.0.2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testLogTailConfig, err := log.NewLogTailConfig(ctx, "testLogTailConfig", &log.LogTailConfigArgs{
// 			Project:    testProject.Name,
// 			Logstore:   testStore.Name,
// 			InputType:  pulumi.String("file"),
// 			LogSample:  pulumi.String("test"),
// 			OutputType: pulumi.String("LogService"),
// 			InputDetail: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v", "  	{\n", "		\"logPath\": \"/logPath\",\n", "		\"filePattern\": \"access.log\",\n", "		\"logType\": \"json_log\",\n", "		\"topicFormat\": \"default\",\n", "		\"discardUnmatch\": false,\n", "		\"enableRawLog\": true,\n", "		\"fileEncoding\": \"gbk\",\n", "		\"maxDepth\": 10\n", "	}\n", "	\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = log.NewLogTailAttachment(ctx, "testLogTailAttachment", &log.LogTailAttachmentArgs{
// 			Project:           testProject.Name,
// 			LogtailConfigName: testLogTailConfig.Name,
// 			MachineGroupName:  testMachineGroup.Name,
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
// Logtial to machine group can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:log/logTailAttachment:LogTailAttachment example tf-log:tf-log-config:tf-log-machine-group
// ```
type LogTailAttachment struct {
	pulumi.CustomResourceState

	// The Logtail configuration name, which is unique in the same project.
	LogtailConfigName pulumi.StringOutput `pulumi:"logtailConfigName"`
	// The machine group name, which is unique in the same project.
	MachineGroupName pulumi.StringOutput `pulumi:"machineGroupName"`
	// The project name to the log store belongs.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewLogTailAttachment registers a new resource with the given unique name, arguments, and options.
func NewLogTailAttachment(ctx *pulumi.Context,
	name string, args *LogTailAttachmentArgs, opts ...pulumi.ResourceOption) (*LogTailAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LogtailConfigName == nil {
		return nil, errors.New("invalid value for required argument 'LogtailConfigName'")
	}
	if args.MachineGroupName == nil {
		return nil, errors.New("invalid value for required argument 'MachineGroupName'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource LogTailAttachment
	err := ctx.RegisterResource("alicloud:log/logTailAttachment:LogTailAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogTailAttachment gets an existing LogTailAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogTailAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogTailAttachmentState, opts ...pulumi.ResourceOption) (*LogTailAttachment, error) {
	var resource LogTailAttachment
	err := ctx.ReadResource("alicloud:log/logTailAttachment:LogTailAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogTailAttachment resources.
type logTailAttachmentState struct {
	// The Logtail configuration name, which is unique in the same project.
	LogtailConfigName *string `pulumi:"logtailConfigName"`
	// The machine group name, which is unique in the same project.
	MachineGroupName *string `pulumi:"machineGroupName"`
	// The project name to the log store belongs.
	Project *string `pulumi:"project"`
}

type LogTailAttachmentState struct {
	// The Logtail configuration name, which is unique in the same project.
	LogtailConfigName pulumi.StringPtrInput
	// The machine group name, which is unique in the same project.
	MachineGroupName pulumi.StringPtrInput
	// The project name to the log store belongs.
	Project pulumi.StringPtrInput
}

func (LogTailAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*logTailAttachmentState)(nil)).Elem()
}

type logTailAttachmentArgs struct {
	// The Logtail configuration name, which is unique in the same project.
	LogtailConfigName string `pulumi:"logtailConfigName"`
	// The machine group name, which is unique in the same project.
	MachineGroupName string `pulumi:"machineGroupName"`
	// The project name to the log store belongs.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a LogTailAttachment resource.
type LogTailAttachmentArgs struct {
	// The Logtail configuration name, which is unique in the same project.
	LogtailConfigName pulumi.StringInput
	// The machine group name, which is unique in the same project.
	MachineGroupName pulumi.StringInput
	// The project name to the log store belongs.
	Project pulumi.StringInput
}

func (LogTailAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logTailAttachmentArgs)(nil)).Elem()
}

type LogTailAttachmentInput interface {
	pulumi.Input

	ToLogTailAttachmentOutput() LogTailAttachmentOutput
	ToLogTailAttachmentOutputWithContext(ctx context.Context) LogTailAttachmentOutput
}

func (LogTailAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*LogTailAttachment)(nil)).Elem()
}

func (i LogTailAttachment) ToLogTailAttachmentOutput() LogTailAttachmentOutput {
	return i.ToLogTailAttachmentOutputWithContext(context.Background())
}

func (i LogTailAttachment) ToLogTailAttachmentOutputWithContext(ctx context.Context) LogTailAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogTailAttachmentOutput)
}

type LogTailAttachmentOutput struct {
	*pulumi.OutputState
}

func (LogTailAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogTailAttachmentOutput)(nil)).Elem()
}

func (o LogTailAttachmentOutput) ToLogTailAttachmentOutput() LogTailAttachmentOutput {
	return o
}

func (o LogTailAttachmentOutput) ToLogTailAttachmentOutputWithContext(ctx context.Context) LogTailAttachmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LogTailAttachmentOutput{})
}
