// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emrv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a EMR cluster resource. This resource is based on EMR's new version OpenAPI.
//
// For information about EMR New and how to use it, see [Add a domain](https://www.alibabacloud.com/help/doc-detail/28068.htm).
//
// > **NOTE:** Available in v1.199.0+.
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
//	"fmt"
//
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/emrv2"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/ram"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi-alicloud/sdk/v3/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultResourceGroups, err := resourcemanager.GetResourceGroups(ctx, &resourcemanager.GetResourceGroupsArgs{
//				Status: pulumi.StringRef("OK"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultZones, err := alicloud.GetZones(ctx, &alicloud.GetZonesArgs{
//				AvailableInstanceType: pulumi.StringRef("ecs.g7.xlarge"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				VpcName:   pulumi.String("TF-VPC"),
//				CidrBlock: pulumi.String("172.16.0.0/12"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VpcId:       defaultNetwork.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/21"),
//				ZoneId:      *pulumi.String(defaultZones.Zones[0].Id),
//				VswitchName: pulumi.String("TF_VSwitch"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultEcsKeyPair, err := ecs.NewEcsKeyPair(ctx, "defaultEcsKeyPair", &ecs.EcsKeyPairArgs{
//				KeyPairName: pulumi.String("terraform-kp"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSecurityGroup, err := ecs.NewSecurityGroup(ctx, "defaultSecurityGroup", &ecs.SecurityGroupArgs{
//				VpcId: defaultNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			defaultRole, err := ram.NewRole(ctx, "defaultRole", &ram.RoleArgs{
//				Document: pulumi.String(fmt.Sprintf(`    {
//	        "Statement": [
//	        {
//	            "Action": "sts:AssumeRole",
//	            "Effect": "Allow",
//	            "Principal": {
//	            "Service": [
//	                "emr.aliyuncs.com",
//	                "ecs.aliyuncs.com"
//	            ]
//	            }
//	        }
//	        ],
//	        "Version": "1"
//	    }
//
// `)),
//
//				Description: pulumi.String("this is a role test."),
//				Force:       pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = emrv2.NewCluster(ctx, "defaultCluster", &emrv2.ClusterArgs{
//				PaymentType:    pulumi.String("PayAsYouGo"),
//				ClusterType:    pulumi.String("DATALAKE"),
//				ReleaseVersion: pulumi.String("EMR-5.10.0"),
//				ClusterName:    pulumi.String("terraform-emr-cluster-v2"),
//				DeployMode:     pulumi.String("NORMAL"),
//				SecurityMode:   pulumi.String("NORMAL"),
//				Applications: pulumi.StringArray{
//					pulumi.String("HADOOP-COMMON"),
//					pulumi.String("HDFS"),
//					pulumi.String("YARN"),
//					pulumi.String("HIVE"),
//					pulumi.String("SPARK3"),
//					pulumi.String("TEZ"),
//				},
//				ApplicationConfigs: emrv2.ClusterApplicationConfigArray{
//					&emrv2.ClusterApplicationConfigArgs{
//						ApplicationName: pulumi.String("HIVE"),
//						ConfigFileName:  pulumi.String("hivemetastore-site.xml"),
//						ConfigItemKey:   pulumi.String("hive.metastore.type"),
//						ConfigItemValue: pulumi.String("DLF"),
//						ConfigScope:     pulumi.String("CLUSTER"),
//					},
//					&emrv2.ClusterApplicationConfigArgs{
//						ApplicationName: pulumi.String("SPARK3"),
//						ConfigFileName:  pulumi.String("hive-site.xml"),
//						ConfigItemKey:   pulumi.String("hive.metastore.type"),
//						ConfigItemValue: pulumi.String("DLF"),
//						ConfigScope:     pulumi.String("CLUSTER"),
//					},
//				},
//				NodeAttributes: emrv2.ClusterNodeAttributeArray{
//					&emrv2.ClusterNodeAttributeArgs{
//						RamRole:         defaultRole.Name,
//						SecurityGroupId: defaultSecurityGroup.ID(),
//						VpcId:           defaultNetwork.ID(),
//						ZoneId:          *pulumi.String(defaultZones.Zones[0].Id),
//						KeyPairName:     defaultEcsKeyPair.ID(),
//					},
//				},
//				Tags: pulumi.AnyMap{
//					"created": pulumi.Any("tf"),
//				},
//				NodeGroups: emrv2.ClusterNodeGroupArray{
//					&emrv2.ClusterNodeGroupArgs{
//						NodeGroupType: pulumi.String("MASTER"),
//						NodeGroupName: pulumi.String("emr-master"),
//						PaymentType:   pulumi.String("PayAsYouGo"),
//						VswitchIds: pulumi.StringArray{
//							defaultSwitch.ID(),
//						},
//						WithPublicIp: pulumi.Bool(false),
//						InstanceTypes: pulumi.StringArray{
//							pulumi.String("ecs.g7.xlarge"),
//						},
//						NodeCount: pulumi.Int(1),
//						SystemDisk: &emrv2.ClusterNodeGroupSystemDiskArgs{
//							Category: pulumi.String("cloud_essd"),
//							Size:     pulumi.Int(80),
//							Count:    pulumi.Int(1),
//						},
//						DataDisks: emrv2.ClusterNodeGroupDataDiskArray{
//							&emrv2.ClusterNodeGroupDataDiskArgs{
//								Category: pulumi.String("cloud_essd"),
//								Size:     pulumi.Int(80),
//								Count:    pulumi.Int(3),
//							},
//						},
//					},
//					&emrv2.ClusterNodeGroupArgs{
//						NodeGroupType: pulumi.String("CORE"),
//						NodeGroupName: pulumi.String("emr-core"),
//						PaymentType:   pulumi.String("PayAsYouGo"),
//						VswitchIds: pulumi.StringArray{
//							defaultSwitch.ID(),
//						},
//						WithPublicIp: pulumi.Bool(false),
//						InstanceTypes: pulumi.StringArray{
//							pulumi.String("ecs.g7.xlarge"),
//						},
//						NodeCount: pulumi.Int(3),
//						SystemDisk: &emrv2.ClusterNodeGroupSystemDiskArgs{
//							Category: pulumi.String("cloud_essd"),
//							Size:     pulumi.Int(80),
//							Count:    pulumi.Int(1),
//						},
//						DataDisks: emrv2.ClusterNodeGroupDataDiskArray{
//							&emrv2.ClusterNodeGroupDataDiskArgs{
//								Category: pulumi.String("cloud_essd"),
//								Size:     pulumi.Int(80),
//								Count:    pulumi.Int(3),
//							},
//						},
//					},
//				},
//				ResourceGroupId: *pulumi.String(defaultResourceGroups.Ids[0]),
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
// Aliclioud E-MapReduce cluster can be imported using the id e.g.
//
// ```sh
//
//	$ pulumi import alicloud:emrv2/cluster:Cluster default <id>
//
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// The application configurations of EMR cluster.
	ApplicationConfigs ClusterApplicationConfigArrayOutput `pulumi:"applicationConfigs"`
	// The applications of EMR cluster to be installed, e.g. HADOOP-COMMON, HDFS, YARN, HIVE, SPARK2, SPARK3, ZOOKEEPER etc. You can find all valid applications in emr web console.
	Applications pulumi.StringArrayOutput `pulumi:"applications"`
	// The bootstrap scripts to be effected when creating emr-cluster or resize emr-cluster.
	BootstrapScripts ClusterBootstrapScriptArrayOutput `pulumi:"bootstrapScripts"`
	// The name of emr cluster. The name length must be less than 64. Supported characters: chinese character, english character, number, "-", "_".
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// EMR Cluster Type, e.g. DATALAKE, OLAP, DATAFLOW, DATASERVING, CUSTOM etc. You can find all valid EMR cluster type in emr web console.
	ClusterType pulumi.StringOutput `pulumi:"clusterType"`
	// The deploy mode of EMR cluster. Supported value: NORMAL or HA.
	DeployMode pulumi.StringOutput `pulumi:"deployMode"`
	// The node attributes of ecs instances which the emr-cluster belongs.
	NodeAttributes ClusterNodeAttributeArrayOutput `pulumi:"nodeAttributes"`
	// Groups of node, You can specify MASTER as a group, CORE as a group (just like the above example).
	NodeGroups ClusterNodeGroupArrayOutput `pulumi:"nodeGroups"`
	// Payment Type for this cluster. Supported value: PayAsYouGo or Subscription.
	PaymentType pulumi.StringOutput `pulumi:"paymentType"`
	// EMR Version, e.g. EMR-5.10.0. You can find the all valid EMR Version in emr web console.
	ReleaseVersion pulumi.StringOutput `pulumi:"releaseVersion"`
	// The Id of resource group which the emr-cluster belongs.
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// The security mode of EMR cluster. Supported value: NORMAL or KERBEROS.
	SecurityMode pulumi.StringOutput `pulumi:"securityMode"`
	// The detail configuration of subscription payment type.
	SubscriptionConfig ClusterSubscriptionConfigPtrOutput `pulumi:"subscriptionConfig"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Applications == nil {
		return nil, errors.New("invalid value for required argument 'Applications'")
	}
	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ClusterType == nil {
		return nil, errors.New("invalid value for required argument 'ClusterType'")
	}
	if args.NodeAttributes == nil {
		return nil, errors.New("invalid value for required argument 'NodeAttributes'")
	}
	if args.NodeGroups == nil {
		return nil, errors.New("invalid value for required argument 'NodeGroups'")
	}
	if args.ReleaseVersion == nil {
		return nil, errors.New("invalid value for required argument 'ReleaseVersion'")
	}
	var resource Cluster
	err := ctx.RegisterResource("alicloud:emrv2/cluster:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("alicloud:emrv2/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The application configurations of EMR cluster.
	ApplicationConfigs []ClusterApplicationConfig `pulumi:"applicationConfigs"`
	// The applications of EMR cluster to be installed, e.g. HADOOP-COMMON, HDFS, YARN, HIVE, SPARK2, SPARK3, ZOOKEEPER etc. You can find all valid applications in emr web console.
	Applications []string `pulumi:"applications"`
	// The bootstrap scripts to be effected when creating emr-cluster or resize emr-cluster.
	BootstrapScripts []ClusterBootstrapScript `pulumi:"bootstrapScripts"`
	// The name of emr cluster. The name length must be less than 64. Supported characters: chinese character, english character, number, "-", "_".
	ClusterName *string `pulumi:"clusterName"`
	// EMR Cluster Type, e.g. DATALAKE, OLAP, DATAFLOW, DATASERVING, CUSTOM etc. You can find all valid EMR cluster type in emr web console.
	ClusterType *string `pulumi:"clusterType"`
	// The deploy mode of EMR cluster. Supported value: NORMAL or HA.
	DeployMode *string `pulumi:"deployMode"`
	// The node attributes of ecs instances which the emr-cluster belongs.
	NodeAttributes []ClusterNodeAttribute `pulumi:"nodeAttributes"`
	// Groups of node, You can specify MASTER as a group, CORE as a group (just like the above example).
	NodeGroups []ClusterNodeGroup `pulumi:"nodeGroups"`
	// Payment Type for this cluster. Supported value: PayAsYouGo or Subscription.
	PaymentType *string `pulumi:"paymentType"`
	// EMR Version, e.g. EMR-5.10.0. You can find the all valid EMR Version in emr web console.
	ReleaseVersion *string `pulumi:"releaseVersion"`
	// The Id of resource group which the emr-cluster belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The security mode of EMR cluster. Supported value: NORMAL or KERBEROS.
	SecurityMode *string `pulumi:"securityMode"`
	// The detail configuration of subscription payment type.
	SubscriptionConfig *ClusterSubscriptionConfig `pulumi:"subscriptionConfig"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type ClusterState struct {
	// The application configurations of EMR cluster.
	ApplicationConfigs ClusterApplicationConfigArrayInput
	// The applications of EMR cluster to be installed, e.g. HADOOP-COMMON, HDFS, YARN, HIVE, SPARK2, SPARK3, ZOOKEEPER etc. You can find all valid applications in emr web console.
	Applications pulumi.StringArrayInput
	// The bootstrap scripts to be effected when creating emr-cluster or resize emr-cluster.
	BootstrapScripts ClusterBootstrapScriptArrayInput
	// The name of emr cluster. The name length must be less than 64. Supported characters: chinese character, english character, number, "-", "_".
	ClusterName pulumi.StringPtrInput
	// EMR Cluster Type, e.g. DATALAKE, OLAP, DATAFLOW, DATASERVING, CUSTOM etc. You can find all valid EMR cluster type in emr web console.
	ClusterType pulumi.StringPtrInput
	// The deploy mode of EMR cluster. Supported value: NORMAL or HA.
	DeployMode pulumi.StringPtrInput
	// The node attributes of ecs instances which the emr-cluster belongs.
	NodeAttributes ClusterNodeAttributeArrayInput
	// Groups of node, You can specify MASTER as a group, CORE as a group (just like the above example).
	NodeGroups ClusterNodeGroupArrayInput
	// Payment Type for this cluster. Supported value: PayAsYouGo or Subscription.
	PaymentType pulumi.StringPtrInput
	// EMR Version, e.g. EMR-5.10.0. You can find the all valid EMR Version in emr web console.
	ReleaseVersion pulumi.StringPtrInput
	// The Id of resource group which the emr-cluster belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The security mode of EMR cluster. Supported value: NORMAL or KERBEROS.
	SecurityMode pulumi.StringPtrInput
	// The detail configuration of subscription payment type.
	SubscriptionConfig ClusterSubscriptionConfigPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The application configurations of EMR cluster.
	ApplicationConfigs []ClusterApplicationConfig `pulumi:"applicationConfigs"`
	// The applications of EMR cluster to be installed, e.g. HADOOP-COMMON, HDFS, YARN, HIVE, SPARK2, SPARK3, ZOOKEEPER etc. You can find all valid applications in emr web console.
	Applications []string `pulumi:"applications"`
	// The bootstrap scripts to be effected when creating emr-cluster or resize emr-cluster.
	BootstrapScripts []ClusterBootstrapScript `pulumi:"bootstrapScripts"`
	// The name of emr cluster. The name length must be less than 64. Supported characters: chinese character, english character, number, "-", "_".
	ClusterName string `pulumi:"clusterName"`
	// EMR Cluster Type, e.g. DATALAKE, OLAP, DATAFLOW, DATASERVING, CUSTOM etc. You can find all valid EMR cluster type in emr web console.
	ClusterType string `pulumi:"clusterType"`
	// The deploy mode of EMR cluster. Supported value: NORMAL or HA.
	DeployMode *string `pulumi:"deployMode"`
	// The node attributes of ecs instances which the emr-cluster belongs.
	NodeAttributes []ClusterNodeAttribute `pulumi:"nodeAttributes"`
	// Groups of node, You can specify MASTER as a group, CORE as a group (just like the above example).
	NodeGroups []ClusterNodeGroup `pulumi:"nodeGroups"`
	// Payment Type for this cluster. Supported value: PayAsYouGo or Subscription.
	PaymentType *string `pulumi:"paymentType"`
	// EMR Version, e.g. EMR-5.10.0. You can find the all valid EMR Version in emr web console.
	ReleaseVersion string `pulumi:"releaseVersion"`
	// The Id of resource group which the emr-cluster belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// The security mode of EMR cluster. Supported value: NORMAL or KERBEROS.
	SecurityMode *string `pulumi:"securityMode"`
	// The detail configuration of subscription payment type.
	SubscriptionConfig *ClusterSubscriptionConfig `pulumi:"subscriptionConfig"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The application configurations of EMR cluster.
	ApplicationConfigs ClusterApplicationConfigArrayInput
	// The applications of EMR cluster to be installed, e.g. HADOOP-COMMON, HDFS, YARN, HIVE, SPARK2, SPARK3, ZOOKEEPER etc. You can find all valid applications in emr web console.
	Applications pulumi.StringArrayInput
	// The bootstrap scripts to be effected when creating emr-cluster or resize emr-cluster.
	BootstrapScripts ClusterBootstrapScriptArrayInput
	// The name of emr cluster. The name length must be less than 64. Supported characters: chinese character, english character, number, "-", "_".
	ClusterName pulumi.StringInput
	// EMR Cluster Type, e.g. DATALAKE, OLAP, DATAFLOW, DATASERVING, CUSTOM etc. You can find all valid EMR cluster type in emr web console.
	ClusterType pulumi.StringInput
	// The deploy mode of EMR cluster. Supported value: NORMAL or HA.
	DeployMode pulumi.StringPtrInput
	// The node attributes of ecs instances which the emr-cluster belongs.
	NodeAttributes ClusterNodeAttributeArrayInput
	// Groups of node, You can specify MASTER as a group, CORE as a group (just like the above example).
	NodeGroups ClusterNodeGroupArrayInput
	// Payment Type for this cluster. Supported value: PayAsYouGo or Subscription.
	PaymentType pulumi.StringPtrInput
	// EMR Version, e.g. EMR-5.10.0. You can find the all valid EMR Version in emr web console.
	ReleaseVersion pulumi.StringInput
	// The Id of resource group which the emr-cluster belongs.
	ResourceGroupId pulumi.StringPtrInput
	// The security mode of EMR cluster. Supported value: NORMAL or KERBEROS.
	SecurityMode pulumi.StringPtrInput
	// The detail configuration of subscription payment type.
	SubscriptionConfig ClusterSubscriptionConfigPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
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
//	ClusterArray{ ClusterArgs{...} }
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
//	ClusterMap{ "key": ClusterArgs{...} }
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

// The application configurations of EMR cluster.
func (o ClusterOutput) ApplicationConfigs() ClusterApplicationConfigArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterApplicationConfigArrayOutput { return v.ApplicationConfigs }).(ClusterApplicationConfigArrayOutput)
}

// The applications of EMR cluster to be installed, e.g. HADOOP-COMMON, HDFS, YARN, HIVE, SPARK2, SPARK3, ZOOKEEPER etc. You can find all valid applications in emr web console.
func (o ClusterOutput) Applications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.Applications }).(pulumi.StringArrayOutput)
}

// The bootstrap scripts to be effected when creating emr-cluster or resize emr-cluster.
func (o ClusterOutput) BootstrapScripts() ClusterBootstrapScriptArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterBootstrapScriptArrayOutput { return v.BootstrapScripts }).(ClusterBootstrapScriptArrayOutput)
}

// The name of emr cluster. The name length must be less than 64. Supported characters: chinese character, english character, number, "-", "_".
func (o ClusterOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// EMR Cluster Type, e.g. DATALAKE, OLAP, DATAFLOW, DATASERVING, CUSTOM etc. You can find all valid EMR cluster type in emr web console.
func (o ClusterOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterType }).(pulumi.StringOutput)
}

// The deploy mode of EMR cluster. Supported value: NORMAL or HA.
func (o ClusterOutput) DeployMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.DeployMode }).(pulumi.StringOutput)
}

// The node attributes of ecs instances which the emr-cluster belongs.
func (o ClusterOutput) NodeAttributes() ClusterNodeAttributeArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterNodeAttributeArrayOutput { return v.NodeAttributes }).(ClusterNodeAttributeArrayOutput)
}

// Groups of node, You can specify MASTER as a group, CORE as a group (just like the above example).
func (o ClusterOutput) NodeGroups() ClusterNodeGroupArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterNodeGroupArrayOutput { return v.NodeGroups }).(ClusterNodeGroupArrayOutput)
}

// Payment Type for this cluster. Supported value: PayAsYouGo or Subscription.
func (o ClusterOutput) PaymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.PaymentType }).(pulumi.StringOutput)
}

// EMR Version, e.g. EMR-5.10.0. You can find the all valid EMR Version in emr web console.
func (o ClusterOutput) ReleaseVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ReleaseVersion }).(pulumi.StringOutput)
}

// The Id of resource group which the emr-cluster belongs.
func (o ClusterOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// The security mode of EMR cluster. Supported value: NORMAL or KERBEROS.
func (o ClusterOutput) SecurityMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.SecurityMode }).(pulumi.StringOutput)
}

// The detail configuration of subscription payment type.
func (o ClusterOutput) SubscriptionConfig() ClusterSubscriptionConfigPtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterSubscriptionConfigPtrOutput { return v.SubscriptionConfig }).(ClusterSubscriptionConfigPtrOutput)
}

// A mapping of tags to assign to the resource.
func (o ClusterOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
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