// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hbase

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a HBase instance resource supports replica set instances only. The HBase provides stable, reliable, and automatic scalable database services.
// It offers a full range of database solutions, such as disaster recovery, backup, recovery, monitoring, and alarms.
// You can see detail product introduction [here](https://help.aliyun.com/product/49055.html)
//
// > **NOTE:**  Available in 1.67.0+
//
// > **NOTE:**  The following regions don't support create Classic network HBase instance.
// [`cn-hangzhou`,`cn-shanghai`,`cn-qingdao`,`cn-beijing`,`cn-shenzhen`,`ap-southeast-1a`,.....]
// The official website mark  more regions. or you can call [DescribeRegions](https://help.aliyun.com/document_detail/144489.html)
//
// > **NOTE:**  Create HBase instance or change instance type and storage would cost 15 minutes. Please make full preparation
//
// ## Example Usage
// ### Create a hbase instance
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/hbase"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := hbase.NewInstance(ctx, "_default", &hbase.InstanceArgs{
// 			ColdStorageSize:      pulumi.Int(0),
// 			CoreDiskSize:         pulumi.Int(400),
// 			CoreDiskType:         pulumi.String("cloud_efficiency"),
// 			CoreInstanceQuantity: pulumi.Int(2),
// 			CoreInstanceType:     pulumi.String("hbase.sn1.large"),
// 			EngineVersion:        pulumi.String("2.0"),
// 			MasterInstanceType:   pulumi.String("hbase.sn1.large"),
// 			PayType:              pulumi.String("PostPaid"),
// 			ZoneId:               pulumi.String("cn-shenzhen-b"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// this is a example for class netType instance. you can find more detail with the examples/hbase dir.
//
// ## Import
//
// HBase can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import alicloud:hbase/instance:Instance example hb-wz96815u13k659fvd
// ```
type Instance struct {
	pulumi.CustomResourceState

	// The account of the cluster web ui.
	Account pulumi.StringPtrOutput `pulumi:"account"`
	// Valid values are `true`, `false`, system default to `false`, valid when payType = PrePaid.
	AutoRenew pulumi.BoolOutput `pulumi:"autoRenew"`
	// 0 or 800+. 0 means isColdStorage = false. 800+ means isColdStorage = true.
	ColdStorageSize pulumi.IntPtrOutput `pulumi:"coldStorageSize"`
	// User-defined HBase instance one core node's storage. Valid when engine=hbase/hbaseue. Bds engine no need core_disk_size, space.Unit: GB. Value range:
	// - Custom storage space, value range: [20, 8000].
	// - Cluster min=400GB, 40-GB increments.
	// - Single min=20GB, 1-GB increments.
	CoreDiskSize pulumi.IntPtrOutput `pulumi:"coreDiskSize"`
	// Valid values are `cloudSsd`, `cloudEssdPl1`, `cloudEfficiency`, `localHddPro`, `localSsdPro`，`-`, ``, localDisk size is fixed. When engine=bds, no need to set disk type(or empty string).
	CoreDiskType pulumi.StringPtrOutput `pulumi:"coreDiskType"`
	// Default=2. If coreInstanceQuantity > 1, this is cluster's instance. If coreInstanceQuantity = 1, this is a single instance.
	CoreInstanceQuantity pulumi.IntPtrOutput `pulumi:"coreInstanceQuantity"`
	CoreInstanceType     pulumi.StringOutput `pulumi:"coreInstanceType"`
	// The switch of delete protection. True: delete protect, False: no delete protect. You must set false when you want to delete cluster.
	DeletionProtection pulumi.BoolPtrOutput `pulumi:"deletionProtection"`
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, valid when payType = PrePaid,  unit: month. 12, 24, 36 mean 1, 2, 3 years.
	Duration pulumi.IntOutput `pulumi:"duration"`
	// Valid values are "hbase/hbaseue/bds". The following types are supported after v1.73.0: `hbaseue` and ` bds  `.
	Engine pulumi.StringPtrOutput `pulumi:"engine"`
	// HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://help.aliyun.com/document_detail/144607.html).
	// * `masterInstanceType`, `coreInstanceType` - (Required, ForceNew) Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// The switch of delete immediate. True: delete immediate, False: delete delay. You will not found the cluster no matter set true or false.
	ImmediateDeleteFlag pulumi.BoolPtrOutput `pulumi:"immediateDeleteFlag"`
	// The white ip list of the cluster.
	IpWhite pulumi.StringOutput `pulumi:"ipWhite"`
	// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 04:00Z.
	MaintainEndTime pulumi.StringOutput `pulumi:"maintainEndTime"`
	// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 02:00Z.
	MaintainStartTime pulumi.StringOutput `pulumi:"maintainStartTime"`
	// Count nodes of the master node.
	MasterInstanceQuantity pulumi.IntOutput    `pulumi:"masterInstanceQuantity"`
	MasterInstanceType     pulumi.StringOutput `pulumi:"masterInstanceType"`
	// HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password of the cluster web ui account.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. You can also convert PostPaid to PrePaid. Not support convert PrePaid to PostPaid.
	PayType pulumi.StringPtrOutput `pulumi:"payType"`
	// The security group resource of the cluster.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// (Available in 1.105.0+) The slb service addresses of the cluster.
	SlbConnAddrs InstanceSlbConnAddrArrayOutput `pulumi:"slbConnAddrs"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// (Available in 1.105.0+) The Web UI proxy addresses of the cluster.
	UiProxyConnAddrs InstanceUiProxyConnAddrArrayOutput `pulumi:"uiProxyConnAddrs"`
	// If vswitchId is not empty, that mean netType = vpc and has a same region. If vswitchId is empty, net_type=classic.
	VswitchId pulumi.StringPtrOutput `pulumi:"vswitchId"`
	// (Available in 1.105.0+) The zookeeper addresses of the cluster.
	ZkConnAddrs InstanceZkConnAddrArrayOutput `pulumi:"zkConnAddrs"`
	// The Zone to launch the HBase instance. If vswitchId is not empty, this zoneId can be "" or consistent.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CoreInstanceType == nil {
		return nil, errors.New("invalid value for required argument 'CoreInstanceType'")
	}
	if args.EngineVersion == nil {
		return nil, errors.New("invalid value for required argument 'EngineVersion'")
	}
	if args.MasterInstanceType == nil {
		return nil, errors.New("invalid value for required argument 'MasterInstanceType'")
	}
	var resource Instance
	err := ctx.RegisterResource("alicloud:hbase/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:hbase/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The account of the cluster web ui.
	Account *string `pulumi:"account"`
	// Valid values are `true`, `false`, system default to `false`, valid when payType = PrePaid.
	AutoRenew *bool `pulumi:"autoRenew"`
	// 0 or 800+. 0 means isColdStorage = false. 800+ means isColdStorage = true.
	ColdStorageSize *int `pulumi:"coldStorageSize"`
	// User-defined HBase instance one core node's storage. Valid when engine=hbase/hbaseue. Bds engine no need core_disk_size, space.Unit: GB. Value range:
	// - Custom storage space, value range: [20, 8000].
	// - Cluster min=400GB, 40-GB increments.
	// - Single min=20GB, 1-GB increments.
	CoreDiskSize *int `pulumi:"coreDiskSize"`
	// Valid values are `cloudSsd`, `cloudEssdPl1`, `cloudEfficiency`, `localHddPro`, `localSsdPro`，`-`, ``, localDisk size is fixed. When engine=bds, no need to set disk type(or empty string).
	CoreDiskType *string `pulumi:"coreDiskType"`
	// Default=2. If coreInstanceQuantity > 1, this is cluster's instance. If coreInstanceQuantity = 1, this is a single instance.
	CoreInstanceQuantity *int    `pulumi:"coreInstanceQuantity"`
	CoreInstanceType     *string `pulumi:"coreInstanceType"`
	// The switch of delete protection. True: delete protect, False: no delete protect. You must set false when you want to delete cluster.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, valid when payType = PrePaid,  unit: month. 12, 24, 36 mean 1, 2, 3 years.
	Duration *int `pulumi:"duration"`
	// Valid values are "hbase/hbaseue/bds". The following types are supported after v1.73.0: `hbaseue` and ` bds  `.
	Engine *string `pulumi:"engine"`
	// HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://help.aliyun.com/document_detail/144607.html).
	// * `masterInstanceType`, `coreInstanceType` - (Required, ForceNew) Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
	EngineVersion *string `pulumi:"engineVersion"`
	// The switch of delete immediate. True: delete immediate, False: delete delay. You will not found the cluster no matter set true or false.
	ImmediateDeleteFlag *bool `pulumi:"immediateDeleteFlag"`
	// The white ip list of the cluster.
	IpWhite *string `pulumi:"ipWhite"`
	// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 04:00Z.
	MaintainEndTime *string `pulumi:"maintainEndTime"`
	// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 02:00Z.
	MaintainStartTime *string `pulumi:"maintainStartTime"`
	// Count nodes of the master node.
	MasterInstanceQuantity *int    `pulumi:"masterInstanceQuantity"`
	MasterInstanceType     *string `pulumi:"masterInstanceType"`
	// HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
	Name *string `pulumi:"name"`
	// The password of the cluster web ui account.
	Password *string `pulumi:"password"`
	// Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. You can also convert PostPaid to PrePaid. Not support convert PrePaid to PostPaid.
	PayType *string `pulumi:"payType"`
	// The security group resource of the cluster.
	SecurityGroups []string `pulumi:"securityGroups"`
	// (Available in 1.105.0+) The slb service addresses of the cluster.
	SlbConnAddrs []InstanceSlbConnAddr `pulumi:"slbConnAddrs"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// (Available in 1.105.0+) The Web UI proxy addresses of the cluster.
	UiProxyConnAddrs []InstanceUiProxyConnAddr `pulumi:"uiProxyConnAddrs"`
	// If vswitchId is not empty, that mean netType = vpc and has a same region. If vswitchId is empty, net_type=classic.
	VswitchId *string `pulumi:"vswitchId"`
	// (Available in 1.105.0+) The zookeeper addresses of the cluster.
	ZkConnAddrs []InstanceZkConnAddr `pulumi:"zkConnAddrs"`
	// The Zone to launch the HBase instance. If vswitchId is not empty, this zoneId can be "" or consistent.
	ZoneId *string `pulumi:"zoneId"`
}

type InstanceState struct {
	// The account of the cluster web ui.
	Account pulumi.StringPtrInput
	// Valid values are `true`, `false`, system default to `false`, valid when payType = PrePaid.
	AutoRenew pulumi.BoolPtrInput
	// 0 or 800+. 0 means isColdStorage = false. 800+ means isColdStorage = true.
	ColdStorageSize pulumi.IntPtrInput
	// User-defined HBase instance one core node's storage. Valid when engine=hbase/hbaseue. Bds engine no need core_disk_size, space.Unit: GB. Value range:
	// - Custom storage space, value range: [20, 8000].
	// - Cluster min=400GB, 40-GB increments.
	// - Single min=20GB, 1-GB increments.
	CoreDiskSize pulumi.IntPtrInput
	// Valid values are `cloudSsd`, `cloudEssdPl1`, `cloudEfficiency`, `localHddPro`, `localSsdPro`，`-`, ``, localDisk size is fixed. When engine=bds, no need to set disk type(or empty string).
	CoreDiskType pulumi.StringPtrInput
	// Default=2. If coreInstanceQuantity > 1, this is cluster's instance. If coreInstanceQuantity = 1, this is a single instance.
	CoreInstanceQuantity pulumi.IntPtrInput
	CoreInstanceType     pulumi.StringPtrInput
	// The switch of delete protection. True: delete protect, False: no delete protect. You must set false when you want to delete cluster.
	DeletionProtection pulumi.BoolPtrInput
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, valid when payType = PrePaid,  unit: month. 12, 24, 36 mean 1, 2, 3 years.
	Duration pulumi.IntPtrInput
	// Valid values are "hbase/hbaseue/bds". The following types are supported after v1.73.0: `hbaseue` and ` bds  `.
	Engine pulumi.StringPtrInput
	// HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://help.aliyun.com/document_detail/144607.html).
	// * `masterInstanceType`, `coreInstanceType` - (Required, ForceNew) Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
	EngineVersion pulumi.StringPtrInput
	// The switch of delete immediate. True: delete immediate, False: delete delay. You will not found the cluster no matter set true or false.
	ImmediateDeleteFlag pulumi.BoolPtrInput
	// The white ip list of the cluster.
	IpWhite pulumi.StringPtrInput
	// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 04:00Z.
	MaintainEndTime pulumi.StringPtrInput
	// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 02:00Z.
	MaintainStartTime pulumi.StringPtrInput
	// Count nodes of the master node.
	MasterInstanceQuantity pulumi.IntPtrInput
	MasterInstanceType     pulumi.StringPtrInput
	// HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
	Name pulumi.StringPtrInput
	// The password of the cluster web ui account.
	Password pulumi.StringPtrInput
	// Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. You can also convert PostPaid to PrePaid. Not support convert PrePaid to PostPaid.
	PayType pulumi.StringPtrInput
	// The security group resource of the cluster.
	SecurityGroups pulumi.StringArrayInput
	// (Available in 1.105.0+) The slb service addresses of the cluster.
	SlbConnAddrs InstanceSlbConnAddrArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// (Available in 1.105.0+) The Web UI proxy addresses of the cluster.
	UiProxyConnAddrs InstanceUiProxyConnAddrArrayInput
	// If vswitchId is not empty, that mean netType = vpc and has a same region. If vswitchId is empty, net_type=classic.
	VswitchId pulumi.StringPtrInput
	// (Available in 1.105.0+) The zookeeper addresses of the cluster.
	ZkConnAddrs InstanceZkConnAddrArrayInput
	// The Zone to launch the HBase instance. If vswitchId is not empty, this zoneId can be "" or consistent.
	ZoneId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The account of the cluster web ui.
	Account *string `pulumi:"account"`
	// Valid values are `true`, `false`, system default to `false`, valid when payType = PrePaid.
	AutoRenew *bool `pulumi:"autoRenew"`
	// 0 or 800+. 0 means isColdStorage = false. 800+ means isColdStorage = true.
	ColdStorageSize *int `pulumi:"coldStorageSize"`
	// User-defined HBase instance one core node's storage. Valid when engine=hbase/hbaseue. Bds engine no need core_disk_size, space.Unit: GB. Value range:
	// - Custom storage space, value range: [20, 8000].
	// - Cluster min=400GB, 40-GB increments.
	// - Single min=20GB, 1-GB increments.
	CoreDiskSize *int `pulumi:"coreDiskSize"`
	// Valid values are `cloudSsd`, `cloudEssdPl1`, `cloudEfficiency`, `localHddPro`, `localSsdPro`，`-`, ``, localDisk size is fixed. When engine=bds, no need to set disk type(or empty string).
	CoreDiskType *string `pulumi:"coreDiskType"`
	// Default=2. If coreInstanceQuantity > 1, this is cluster's instance. If coreInstanceQuantity = 1, this is a single instance.
	CoreInstanceQuantity *int   `pulumi:"coreInstanceQuantity"`
	CoreInstanceType     string `pulumi:"coreInstanceType"`
	// The switch of delete protection. True: delete protect, False: no delete protect. You must set false when you want to delete cluster.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, valid when payType = PrePaid,  unit: month. 12, 24, 36 mean 1, 2, 3 years.
	Duration *int `pulumi:"duration"`
	// Valid values are "hbase/hbaseue/bds". The following types are supported after v1.73.0: `hbaseue` and ` bds  `.
	Engine *string `pulumi:"engine"`
	// HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://help.aliyun.com/document_detail/144607.html).
	// * `masterInstanceType`, `coreInstanceType` - (Required, ForceNew) Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
	EngineVersion string `pulumi:"engineVersion"`
	// The switch of delete immediate. True: delete immediate, False: delete delay. You will not found the cluster no matter set true or false.
	ImmediateDeleteFlag *bool `pulumi:"immediateDeleteFlag"`
	// The white ip list of the cluster.
	IpWhite *string `pulumi:"ipWhite"`
	// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 04:00Z.
	MaintainEndTime *string `pulumi:"maintainEndTime"`
	// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 02:00Z.
	MaintainStartTime  *string `pulumi:"maintainStartTime"`
	MasterInstanceType string  `pulumi:"masterInstanceType"`
	// HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
	Name *string `pulumi:"name"`
	// The password of the cluster web ui account.
	Password *string `pulumi:"password"`
	// Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. You can also convert PostPaid to PrePaid. Not support convert PrePaid to PostPaid.
	PayType *string `pulumi:"payType"`
	// The security group resource of the cluster.
	SecurityGroups []string `pulumi:"securityGroups"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// If vswitchId is not empty, that mean netType = vpc and has a same region. If vswitchId is empty, net_type=classic.
	VswitchId *string `pulumi:"vswitchId"`
	// The Zone to launch the HBase instance. If vswitchId is not empty, this zoneId can be "" or consistent.
	ZoneId *string `pulumi:"zoneId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The account of the cluster web ui.
	Account pulumi.StringPtrInput
	// Valid values are `true`, `false`, system default to `false`, valid when payType = PrePaid.
	AutoRenew pulumi.BoolPtrInput
	// 0 or 800+. 0 means isColdStorage = false. 800+ means isColdStorage = true.
	ColdStorageSize pulumi.IntPtrInput
	// User-defined HBase instance one core node's storage. Valid when engine=hbase/hbaseue. Bds engine no need core_disk_size, space.Unit: GB. Value range:
	// - Custom storage space, value range: [20, 8000].
	// - Cluster min=400GB, 40-GB increments.
	// - Single min=20GB, 1-GB increments.
	CoreDiskSize pulumi.IntPtrInput
	// Valid values are `cloudSsd`, `cloudEssdPl1`, `cloudEfficiency`, `localHddPro`, `localSsdPro`，`-`, ``, localDisk size is fixed. When engine=bds, no need to set disk type(or empty string).
	CoreDiskType pulumi.StringPtrInput
	// Default=2. If coreInstanceQuantity > 1, this is cluster's instance. If coreInstanceQuantity = 1, this is a single instance.
	CoreInstanceQuantity pulumi.IntPtrInput
	CoreInstanceType     pulumi.StringInput
	// The switch of delete protection. True: delete protect, False: no delete protect. You must set false when you want to delete cluster.
	DeletionProtection pulumi.BoolPtrInput
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, valid when payType = PrePaid,  unit: month. 12, 24, 36 mean 1, 2, 3 years.
	Duration pulumi.IntPtrInput
	// Valid values are "hbase/hbaseue/bds". The following types are supported after v1.73.0: `hbaseue` and ` bds  `.
	Engine pulumi.StringPtrInput
	// HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://help.aliyun.com/document_detail/144607.html).
	// * `masterInstanceType`, `coreInstanceType` - (Required, ForceNew) Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
	EngineVersion pulumi.StringInput
	// The switch of delete immediate. True: delete immediate, False: delete delay. You will not found the cluster no matter set true or false.
	ImmediateDeleteFlag pulumi.BoolPtrInput
	// The white ip list of the cluster.
	IpWhite pulumi.StringPtrInput
	// The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 04:00Z.
	MaintainEndTime pulumi.StringPtrInput
	// The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 02:00Z.
	MaintainStartTime  pulumi.StringPtrInput
	MasterInstanceType pulumi.StringInput
	// HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
	Name pulumi.StringPtrInput
	// The password of the cluster web ui account.
	Password pulumi.StringPtrInput
	// Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. You can also convert PostPaid to PrePaid. Not support convert PrePaid to PostPaid.
	PayType pulumi.StringPtrInput
	// The security group resource of the cluster.
	SecurityGroups pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// If vswitchId is not empty, that mean netType = vpc and has a same region. If vswitchId is empty, net_type=classic.
	VswitchId pulumi.StringPtrInput
	// The Zone to launch the HBase instance. If vswitchId is not empty, this zoneId can be "" or consistent.
	ZoneId pulumi.StringPtrInput
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
