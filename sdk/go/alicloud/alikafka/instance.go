// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alikafka

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an ALIKAFKA instance resource.
//
// > **NOTE:** Available in 1.59.0+
//
// > **NOTE:** Creation or modification may took about 10-40 minutes.
//
// > **NOTE:** Only the following regions support create alikafka pre paid instance.
// [`cn-hangzhou`,`cn-beijing`,`cn-shenzhen`,`cn-shanghai`,`cn-qingdao`,`cn-hongkong`,`cn-huhehaote`,`cn-zhangjiakou`,`cn-chengdu`,`cn-heyuan`,`ap-southeast-1`,`ap-southeast-3`,`ap-southeast-5`,`ap-south-1`,`ap-northeast-1`,`eu-central-1`,`eu-west-1`,`us-west-1`,`us-east-1`]
//
// > **NOTE:** Only the following regions support create alikafka post paid instance.
// [`cn-hangzhou`,`cn-beijing`,`cn-shenzhen`,`cn-shanghai`,`cn-qingdao`,`cn-hongkong`,`cn-huhehaote`,`cn-zhangjiakou`,`cn-chengdu`,`cn-heyuan`,`ap-southeast-1`,`ap-southeast-3`,`ap-southeast-5`,`ap-south-1`,`ap-northeast-1`,`eu-central-1`,`eu-west-1`,`us-west-1`,`us-east-1`]
// ## Example Usage
//
// Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/alikafka"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
// 	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		instanceName := "alikafkaInstanceName"
// 		if param := cfg.Get("instanceName"); param != "" {
// 			instanceName = param
// 		}
// 		defaultZones, err := alicloud.GetZones(ctx, &GetZonesArgs{
// 			AvailableResourceCreation: pulumi.StringRef("VSwitch"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
// 			CidrBlock: pulumi.String("172.16.0.0/12"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
// 			VpcId:     defaultNetwork.ID(),
// 			CidrBlock: pulumi.String("172.16.0.0/24"),
// 			ZoneId:    pulumi.String(defaultZones.Zones[0].Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultSecurityGroup, err := ecs.NewSecurityGroup(ctx, "defaultSecurityGroup", &ecs.SecurityGroupArgs{
// 			VpcId: defaultNetwork.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = alikafka.NewInstance(ctx, "defaultInstance", &alikafka.InstanceArgs{
// 			TopicQuota:    pulumi.Int(50),
// 			DiskType:      pulumi.Int(1),
// 			DiskSize:      pulumi.Int(500),
// 			DeployType:    pulumi.Int(4),
// 			IoMax:         pulumi.Int(20),
// 			VswitchId:     defaultSwitch.ID(),
// 			SecurityGroup: defaultSecurityGroup.ID(),
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
// ALIKAFKA TOPIC can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:alikafka/instance:Instance instance alikafka_post-cn-123455abc
// ```
type Instance struct {
	pulumi.CustomResourceState

	// （Optional, Available in v1.112.0+） The basic config for this instance. The input should be json type, only the following key allowed: enable.acl, enable.vpc_sasl_ssl, kafka.log.retention.hours, kafka.message.max.bytes.
	Config pulumi.StringOutput `pulumi:"config"`
	// The deploy type of the instance. Currently only support two deploy type, 4: eip/vpc instance, 5: vpc instance.
	DeployType pulumi.IntOutput `pulumi:"deployType"`
	// The disk size of the instance. When modify this value, it only support adjust to a greater value.
	DiskSize pulumi.IntOutput `pulumi:"diskSize"`
	// The disk type of the instance. 0: efficient cloud disk , 1: SSD.
	DiskType pulumi.IntOutput `pulumi:"diskType"`
	// The max bandwidth of the instance. When modify this value, it only support adjust to a greater value.
	EipMax pulumi.IntOutput `pulumi:"eipMax"`
	// The EndPoint to access the kafka instance.
	EndPoint pulumi.StringOutput `pulumi:"endPoint"`
	// The max value of io of the instance. When modify this value, it only support adjust to a greater value.
	IoMax pulumi.IntOutput `pulumi:"ioMax"`
	// Name of your Kafka instance. The length should between 3 and 64 characters. If not set, will use instance id as instance name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The paid type of the instance. Support two type, "PrePaid": pre paid type instance, "PostPaid": post paid type instance. Default is PostPaid. When modify this value, it only support adjust from post pay to pre pay.
	PaidType pulumi.StringPtrOutput `pulumi:"paidType"`
	// （Optional, ForceNew, Available in v1.93.0+） The ID of security group for this instance. If the security group is empty, system will create a default one.
	SecurityGroup pulumi.StringOutput `pulumi:"securityGroup"`
	// （Optional, Available in v1.112.0+） The kafka openSource version for this instance. Only 0.10.2 or 2.2.0 is allowed, default is 0.10.2.
	ServiceVersion pulumi.StringOutput `pulumi:"serviceVersion"`
	// The spec type of the instance. Support two type, "normal": normal version instance, "professional": professional version instance. Default is normal. When modify this value, it only support adjust from normal to professional. Note only pre paid type instance support professional specific type.
	SpecType pulumi.StringPtrOutput `pulumi:"specType"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The max num of topic can be create of the instance. When modify this value, it only adjust to a greater value.
	TopicQuota pulumi.IntOutput `pulumi:"topicQuota"`
	// The ID of attaching VPC to instance.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The ID of attaching vswitch to instance.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
	// The Zone to launch the kafka instance.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeployType == nil {
		return nil, errors.New("invalid value for required argument 'DeployType'")
	}
	if args.DiskSize == nil {
		return nil, errors.New("invalid value for required argument 'DiskSize'")
	}
	if args.DiskType == nil {
		return nil, errors.New("invalid value for required argument 'DiskType'")
	}
	if args.IoMax == nil {
		return nil, errors.New("invalid value for required argument 'IoMax'")
	}
	if args.TopicQuota == nil {
		return nil, errors.New("invalid value for required argument 'TopicQuota'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	var resource Instance
	err := ctx.RegisterResource("alicloud:alikafka/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:alikafka/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// （Optional, Available in v1.112.0+） The basic config for this instance. The input should be json type, only the following key allowed: enable.acl, enable.vpc_sasl_ssl, kafka.log.retention.hours, kafka.message.max.bytes.
	Config *string `pulumi:"config"`
	// The deploy type of the instance. Currently only support two deploy type, 4: eip/vpc instance, 5: vpc instance.
	DeployType *int `pulumi:"deployType"`
	// The disk size of the instance. When modify this value, it only support adjust to a greater value.
	DiskSize *int `pulumi:"diskSize"`
	// The disk type of the instance. 0: efficient cloud disk , 1: SSD.
	DiskType *int `pulumi:"diskType"`
	// The max bandwidth of the instance. When modify this value, it only support adjust to a greater value.
	EipMax *int `pulumi:"eipMax"`
	// The EndPoint to access the kafka instance.
	EndPoint *string `pulumi:"endPoint"`
	// The max value of io of the instance. When modify this value, it only support adjust to a greater value.
	IoMax *int `pulumi:"ioMax"`
	// Name of your Kafka instance. The length should between 3 and 64 characters. If not set, will use instance id as instance name.
	Name *string `pulumi:"name"`
	// The paid type of the instance. Support two type, "PrePaid": pre paid type instance, "PostPaid": post paid type instance. Default is PostPaid. When modify this value, it only support adjust from post pay to pre pay.
	PaidType *string `pulumi:"paidType"`
	// （Optional, ForceNew, Available in v1.93.0+） The ID of security group for this instance. If the security group is empty, system will create a default one.
	SecurityGroup *string `pulumi:"securityGroup"`
	// （Optional, Available in v1.112.0+） The kafka openSource version for this instance. Only 0.10.2 or 2.2.0 is allowed, default is 0.10.2.
	ServiceVersion *string `pulumi:"serviceVersion"`
	// The spec type of the instance. Support two type, "normal": normal version instance, "professional": professional version instance. Default is normal. When modify this value, it only support adjust from normal to professional. Note only pre paid type instance support professional specific type.
	SpecType *string `pulumi:"specType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The max num of topic can be create of the instance. When modify this value, it only adjust to a greater value.
	TopicQuota *int `pulumi:"topicQuota"`
	// The ID of attaching VPC to instance.
	VpcId *string `pulumi:"vpcId"`
	// The ID of attaching vswitch to instance.
	VswitchId *string `pulumi:"vswitchId"`
	// The Zone to launch the kafka instance.
	ZoneId *string `pulumi:"zoneId"`
}

type InstanceState struct {
	// （Optional, Available in v1.112.0+） The basic config for this instance. The input should be json type, only the following key allowed: enable.acl, enable.vpc_sasl_ssl, kafka.log.retention.hours, kafka.message.max.bytes.
	Config pulumi.StringPtrInput
	// The deploy type of the instance. Currently only support two deploy type, 4: eip/vpc instance, 5: vpc instance.
	DeployType pulumi.IntPtrInput
	// The disk size of the instance. When modify this value, it only support adjust to a greater value.
	DiskSize pulumi.IntPtrInput
	// The disk type of the instance. 0: efficient cloud disk , 1: SSD.
	DiskType pulumi.IntPtrInput
	// The max bandwidth of the instance. When modify this value, it only support adjust to a greater value.
	EipMax pulumi.IntPtrInput
	// The EndPoint to access the kafka instance.
	EndPoint pulumi.StringPtrInput
	// The max value of io of the instance. When modify this value, it only support adjust to a greater value.
	IoMax pulumi.IntPtrInput
	// Name of your Kafka instance. The length should between 3 and 64 characters. If not set, will use instance id as instance name.
	Name pulumi.StringPtrInput
	// The paid type of the instance. Support two type, "PrePaid": pre paid type instance, "PostPaid": post paid type instance. Default is PostPaid. When modify this value, it only support adjust from post pay to pre pay.
	PaidType pulumi.StringPtrInput
	// （Optional, ForceNew, Available in v1.93.0+） The ID of security group for this instance. If the security group is empty, system will create a default one.
	SecurityGroup pulumi.StringPtrInput
	// （Optional, Available in v1.112.0+） The kafka openSource version for this instance. Only 0.10.2 or 2.2.0 is allowed, default is 0.10.2.
	ServiceVersion pulumi.StringPtrInput
	// The spec type of the instance. Support two type, "normal": normal version instance, "professional": professional version instance. Default is normal. When modify this value, it only support adjust from normal to professional. Note only pre paid type instance support professional specific type.
	SpecType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The max num of topic can be create of the instance. When modify this value, it only adjust to a greater value.
	TopicQuota pulumi.IntPtrInput
	// The ID of attaching VPC to instance.
	VpcId pulumi.StringPtrInput
	// The ID of attaching vswitch to instance.
	VswitchId pulumi.StringPtrInput
	// The Zone to launch the kafka instance.
	ZoneId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// （Optional, Available in v1.112.0+） The basic config for this instance. The input should be json type, only the following key allowed: enable.acl, enable.vpc_sasl_ssl, kafka.log.retention.hours, kafka.message.max.bytes.
	Config *string `pulumi:"config"`
	// The deploy type of the instance. Currently only support two deploy type, 4: eip/vpc instance, 5: vpc instance.
	DeployType int `pulumi:"deployType"`
	// The disk size of the instance. When modify this value, it only support adjust to a greater value.
	DiskSize int `pulumi:"diskSize"`
	// The disk type of the instance. 0: efficient cloud disk , 1: SSD.
	DiskType int `pulumi:"diskType"`
	// The max bandwidth of the instance. When modify this value, it only support adjust to a greater value.
	EipMax *int `pulumi:"eipMax"`
	// The max value of io of the instance. When modify this value, it only support adjust to a greater value.
	IoMax int `pulumi:"ioMax"`
	// Name of your Kafka instance. The length should between 3 and 64 characters. If not set, will use instance id as instance name.
	Name *string `pulumi:"name"`
	// The paid type of the instance. Support two type, "PrePaid": pre paid type instance, "PostPaid": post paid type instance. Default is PostPaid. When modify this value, it only support adjust from post pay to pre pay.
	PaidType *string `pulumi:"paidType"`
	// （Optional, ForceNew, Available in v1.93.0+） The ID of security group for this instance. If the security group is empty, system will create a default one.
	SecurityGroup *string `pulumi:"securityGroup"`
	// （Optional, Available in v1.112.0+） The kafka openSource version for this instance. Only 0.10.2 or 2.2.0 is allowed, default is 0.10.2.
	ServiceVersion *string `pulumi:"serviceVersion"`
	// The spec type of the instance. Support two type, "normal": normal version instance, "professional": professional version instance. Default is normal. When modify this value, it only support adjust from normal to professional. Note only pre paid type instance support professional specific type.
	SpecType *string `pulumi:"specType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The max num of topic can be create of the instance. When modify this value, it only adjust to a greater value.
	TopicQuota int `pulumi:"topicQuota"`
	// The ID of attaching vswitch to instance.
	VswitchId string `pulumi:"vswitchId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// （Optional, Available in v1.112.0+） The basic config for this instance. The input should be json type, only the following key allowed: enable.acl, enable.vpc_sasl_ssl, kafka.log.retention.hours, kafka.message.max.bytes.
	Config pulumi.StringPtrInput
	// The deploy type of the instance. Currently only support two deploy type, 4: eip/vpc instance, 5: vpc instance.
	DeployType pulumi.IntInput
	// The disk size of the instance. When modify this value, it only support adjust to a greater value.
	DiskSize pulumi.IntInput
	// The disk type of the instance. 0: efficient cloud disk , 1: SSD.
	DiskType pulumi.IntInput
	// The max bandwidth of the instance. When modify this value, it only support adjust to a greater value.
	EipMax pulumi.IntPtrInput
	// The max value of io of the instance. When modify this value, it only support adjust to a greater value.
	IoMax pulumi.IntInput
	// Name of your Kafka instance. The length should between 3 and 64 characters. If not set, will use instance id as instance name.
	Name pulumi.StringPtrInput
	// The paid type of the instance. Support two type, "PrePaid": pre paid type instance, "PostPaid": post paid type instance. Default is PostPaid. When modify this value, it only support adjust from post pay to pre pay.
	PaidType pulumi.StringPtrInput
	// （Optional, ForceNew, Available in v1.93.0+） The ID of security group for this instance. If the security group is empty, system will create a default one.
	SecurityGroup pulumi.StringPtrInput
	// （Optional, Available in v1.112.0+） The kafka openSource version for this instance. Only 0.10.2 or 2.2.0 is allowed, default is 0.10.2.
	ServiceVersion pulumi.StringPtrInput
	// The spec type of the instance. Support two type, "normal": normal version instance, "professional": professional version instance. Default is normal. When modify this value, it only support adjust from normal to professional. Note only pre paid type instance support professional specific type.
	SpecType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The max num of topic can be create of the instance. When modify this value, it only adjust to a greater value.
	TopicQuota pulumi.IntInput
	// The ID of attaching vswitch to instance.
	VswitchId pulumi.StringInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//          InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//          InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
