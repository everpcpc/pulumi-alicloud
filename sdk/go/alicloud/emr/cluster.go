// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a EMR Cluster resource. With this you can create, read, and release  EMR Cluster.
//
// > **NOTE:** Available in 1.57.0+.
//
// ## Example Usage
//
// ## Import
//
// Aliclioud E-MapReduce cluster can be imported using the id e.g.
//
// ```sh
//  $ pulumi import alicloud:emr/cluster:Cluster default C-B47FB8FE96C67XXXX
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// Boot action parameters.
	BootstrapActions ClusterBootstrapActionArrayOutput `pulumi:"bootstrapActions"`
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType pulumi.StringPtrOutput `pulumi:"chargeType"`
	// EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
	ClusterType pulumi.StringOutput `pulumi:"clusterType"`
	// Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
	DepositType pulumi.StringPtrOutput `pulumi:"depositType"`
	// High security cluster (true) or not. Default value is false.
	EasEnable pulumi.BoolPtrOutput `pulumi:"easEnable"`
	// EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
	EmrVer pulumi.StringOutput `pulumi:"emrVer"`
	// High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
	HighAvailabilityEnable pulumi.BoolPtrOutput `pulumi:"highAvailabilityEnable"`
	// Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
	HostGroups ClusterHostGroupArrayOutput `pulumi:"hostGroups"`
	// Whether the MASTER node has a public IP address enabled. Default value is false.
	IsOpenPublicIp pulumi.BoolPtrOutput `pulumi:"isOpenPublicIp"`
	// Ssh key pair.
	KeyPairName pulumi.StringPtrOutput `pulumi:"keyPairName"`
	// Master ssh password.
	MasterPwd pulumi.StringPtrOutput `pulumi:"masterPwd"`
	// bootstrap action name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional software list.
	OptionSoftwareLists pulumi.StringArrayOutput `pulumi:"optionSoftwareLists"`
	// If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// This specify the related cluster id, if this cluster is a Gateway.
	RelatedClusterId pulumi.StringPtrOutput `pulumi:"relatedClusterId"`
	// Security Group ID for Cluster, you can also specify this key for each host group.
	SecurityGroupId pulumi.StringPtrOutput `pulumi:"securityGroupId"`
	// If this is set true, we can ssh into cluster. Default value is false.
	SshEnable pulumi.BoolPtrOutput `pulumi:"sshEnable"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Use local metadb. Default is false.
	UseLocalMetadb pulumi.BoolPtrOutput `pulumi:"useLocalMetadb"`
	// Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
	UserDefinedEmrEcsRole pulumi.StringPtrOutput `pulumi:"userDefinedEmrEcsRole"`
	// Global vswitch id, you can also specify it in host group.
	VswitchId pulumi.StringPtrOutput `pulumi:"vswitchId"`
	// Zone ID, e.g. cn-huhehaote-a
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterType == nil {
		return nil, errors.New("invalid value for required argument 'ClusterType'")
	}
	if args.EmrVer == nil {
		return nil, errors.New("invalid value for required argument 'EmrVer'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	var resource Cluster
	err := ctx.RegisterResource("alicloud:emr/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("alicloud:emr/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// Boot action parameters.
	BootstrapActions []ClusterBootstrapAction `pulumi:"bootstrapActions"`
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType *string `pulumi:"chargeType"`
	// EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
	ClusterType *string `pulumi:"clusterType"`
	// Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
	DepositType *string `pulumi:"depositType"`
	// High security cluster (true) or not. Default value is false.
	EasEnable *bool `pulumi:"easEnable"`
	// EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
	EmrVer *string `pulumi:"emrVer"`
	// High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
	HighAvailabilityEnable *bool `pulumi:"highAvailabilityEnable"`
	// Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
	HostGroups []ClusterHostGroup `pulumi:"hostGroups"`
	// Whether the MASTER node has a public IP address enabled. Default value is false.
	IsOpenPublicIp *bool `pulumi:"isOpenPublicIp"`
	// Ssh key pair.
	KeyPairName *string `pulumi:"keyPairName"`
	// Master ssh password.
	MasterPwd *string `pulumi:"masterPwd"`
	// bootstrap action name.
	Name *string `pulumi:"name"`
	// Optional software list.
	OptionSoftwareLists []string `pulumi:"optionSoftwareLists"`
	// If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
	Period *int `pulumi:"period"`
	// This specify the related cluster id, if this cluster is a Gateway.
	RelatedClusterId *string `pulumi:"relatedClusterId"`
	// Security Group ID for Cluster, you can also specify this key for each host group.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// If this is set true, we can ssh into cluster. Default value is false.
	SshEnable *bool `pulumi:"sshEnable"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// Use local metadb. Default is false.
	UseLocalMetadb *bool `pulumi:"useLocalMetadb"`
	// Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
	UserDefinedEmrEcsRole *string `pulumi:"userDefinedEmrEcsRole"`
	// Global vswitch id, you can also specify it in host group.
	VswitchId *string `pulumi:"vswitchId"`
	// Zone ID, e.g. cn-huhehaote-a
	ZoneId *string `pulumi:"zoneId"`
}

type ClusterState struct {
	// Boot action parameters.
	BootstrapActions ClusterBootstrapActionArrayInput
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType pulumi.StringPtrInput
	// EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
	ClusterType pulumi.StringPtrInput
	// Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
	DepositType pulumi.StringPtrInput
	// High security cluster (true) or not. Default value is false.
	EasEnable pulumi.BoolPtrInput
	// EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
	EmrVer pulumi.StringPtrInput
	// High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
	HighAvailabilityEnable pulumi.BoolPtrInput
	// Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
	HostGroups ClusterHostGroupArrayInput
	// Whether the MASTER node has a public IP address enabled. Default value is false.
	IsOpenPublicIp pulumi.BoolPtrInput
	// Ssh key pair.
	KeyPairName pulumi.StringPtrInput
	// Master ssh password.
	MasterPwd pulumi.StringPtrInput
	// bootstrap action name.
	Name pulumi.StringPtrInput
	// Optional software list.
	OptionSoftwareLists pulumi.StringArrayInput
	// If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
	Period pulumi.IntPtrInput
	// This specify the related cluster id, if this cluster is a Gateway.
	RelatedClusterId pulumi.StringPtrInput
	// Security Group ID for Cluster, you can also specify this key for each host group.
	SecurityGroupId pulumi.StringPtrInput
	// If this is set true, we can ssh into cluster. Default value is false.
	SshEnable pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// Use local metadb. Default is false.
	UseLocalMetadb pulumi.BoolPtrInput
	// Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
	UserDefinedEmrEcsRole pulumi.StringPtrInput
	// Global vswitch id, you can also specify it in host group.
	VswitchId pulumi.StringPtrInput
	// Zone ID, e.g. cn-huhehaote-a
	ZoneId pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Boot action parameters.
	BootstrapActions []ClusterBootstrapAction `pulumi:"bootstrapActions"`
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType *string `pulumi:"chargeType"`
	// EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
	ClusterType string `pulumi:"clusterType"`
	// Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
	DepositType *string `pulumi:"depositType"`
	// High security cluster (true) or not. Default value is false.
	EasEnable *bool `pulumi:"easEnable"`
	// EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
	EmrVer string `pulumi:"emrVer"`
	// High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
	HighAvailabilityEnable *bool `pulumi:"highAvailabilityEnable"`
	// Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
	HostGroups []ClusterHostGroup `pulumi:"hostGroups"`
	// Whether the MASTER node has a public IP address enabled. Default value is false.
	IsOpenPublicIp *bool `pulumi:"isOpenPublicIp"`
	// Ssh key pair.
	KeyPairName *string `pulumi:"keyPairName"`
	// Master ssh password.
	MasterPwd *string `pulumi:"masterPwd"`
	// bootstrap action name.
	Name *string `pulumi:"name"`
	// Optional software list.
	OptionSoftwareLists []string `pulumi:"optionSoftwareLists"`
	// If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
	Period *int `pulumi:"period"`
	// This specify the related cluster id, if this cluster is a Gateway.
	RelatedClusterId *string `pulumi:"relatedClusterId"`
	// Security Group ID for Cluster, you can also specify this key for each host group.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// If this is set true, we can ssh into cluster. Default value is false.
	SshEnable *bool `pulumi:"sshEnable"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// Use local metadb. Default is false.
	UseLocalMetadb *bool `pulumi:"useLocalMetadb"`
	// Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
	UserDefinedEmrEcsRole *string `pulumi:"userDefinedEmrEcsRole"`
	// Global vswitch id, you can also specify it in host group.
	VswitchId *string `pulumi:"vswitchId"`
	// Zone ID, e.g. cn-huhehaote-a
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Boot action parameters.
	BootstrapActions ClusterBootstrapActionArrayInput
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType pulumi.StringPtrInput
	// EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
	ClusterType pulumi.StringInput
	// Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
	DepositType pulumi.StringPtrInput
	// High security cluster (true) or not. Default value is false.
	EasEnable pulumi.BoolPtrInput
	// EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
	EmrVer pulumi.StringInput
	// High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
	HighAvailabilityEnable pulumi.BoolPtrInput
	// Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
	HostGroups ClusterHostGroupArrayInput
	// Whether the MASTER node has a public IP address enabled. Default value is false.
	IsOpenPublicIp pulumi.BoolPtrInput
	// Ssh key pair.
	KeyPairName pulumi.StringPtrInput
	// Master ssh password.
	MasterPwd pulumi.StringPtrInput
	// bootstrap action name.
	Name pulumi.StringPtrInput
	// Optional software list.
	OptionSoftwareLists pulumi.StringArrayInput
	// If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
	Period pulumi.IntPtrInput
	// This specify the related cluster id, if this cluster is a Gateway.
	RelatedClusterId pulumi.StringPtrInput
	// Security Group ID for Cluster, you can also specify this key for each host group.
	SecurityGroupId pulumi.StringPtrInput
	// If this is set true, we can ssh into cluster. Default value is false.
	SshEnable pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// Use local metadb. Default is false.
	UseLocalMetadb pulumi.BoolPtrInput
	// Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
	UserDefinedEmrEcsRole pulumi.StringPtrInput
	// Global vswitch id, you can also specify it in host group.
	VswitchId pulumi.StringPtrInput
	// Zone ID, e.g. cn-huhehaote-a
	ZoneId pulumi.StringInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//          ClusterArray{ ClusterArgs{...} }
type ClusterArrayInput interface {
	pulumi.Input

	ToClusterArrayOutput() ClusterArrayOutput
	ToClusterArrayOutputWithContext(context.Context) ClusterArrayOutput
}

type ClusterArray []ClusterInput

func (ClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (i ClusterArray) ToClusterArrayOutput() ClusterArrayOutput {
	return i.ToClusterArrayOutputWithContext(context.Background())
}

func (i ClusterArray) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterArrayOutput)
}

// ClusterMapInput is an input type that accepts ClusterMap and ClusterMapOutput values.
// You can construct a concrete instance of `ClusterMapInput` via:
//
//          ClusterMap{ "key": ClusterArgs{...} }
type ClusterMapInput interface {
	pulumi.Input

	ToClusterMapOutput() ClusterMapOutput
	ToClusterMapOutputWithContext(context.Context) ClusterMapOutput
}

type ClusterMap map[string]ClusterInput

func (ClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (i ClusterMap) ToClusterMapOutput() ClusterMapOutput {
	return i.ToClusterMapOutputWithContext(context.Background())
}

func (i ClusterMap) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMapOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].([]*Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].(map[string]*Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterArrayInput)(nil)).Elem(), ClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMapInput)(nil)).Elem(), ClusterMap{})
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}
