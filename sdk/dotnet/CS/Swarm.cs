// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS
{
    /// <summary>
    /// &gt; **DEPRECATED:** This resource manages swarm cluster, which is being deprecated and will be replaced by Kubernetes cluster.
    /// 
    /// This resource will help you to manager a Swarm Cluster.
    /// 
    /// &gt; **NOTE:** Swarm cluster only supports VPC network and you can specify a VPC network by filed `vswitch_id`.
    /// </summary>
    public partial class Swarm : Pulumi.CustomResource
    {
        /// <summary>
        /// The nodes agent version.
        /// </summary>
        [Output("agentVersion")]
        public Output<string> AgentVersion { get; private set; } = null!;

        /// <summary>
        /// The CIDR block for the Container. It can not be same as the CIDR used by the VPC.
        /// Valid value:
        /// - 192.168.0.0/16
        /// - 172.19-30.0.0/16
        /// - 10.0.0.0/16
        /// </summary>
        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The data disk category of ECS instance node. Its valid value are `cloud`, `cloud_ssd`, `cloud_essd`, `ephemeral_essd` and `cloud_efficiency`. Default to `cloud_efficiency`.
        /// </summary>
        [Output("diskCategory")]
        public Output<string?> DiskCategory { get; private set; } = null!;

        /// <summary>
        /// The data disk size of ECS instance node. Its valid value is 20~32768 GB. Default to 20.
        /// </summary>
        [Output("diskSize")]
        public Output<int?> DiskSize { get; private set; } = null!;

        /// <summary>
        /// The image ID of ECS instance node used. Default to System automate allocated.
        /// </summary>
        [Output("imageId")]
        public Output<string?> ImageId { get; private set; } = null!;

        /// <summary>
        /// The type of ECS instance node.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// Whether to use outdated instance type. Default to false.
        /// </summary>
        [Output("isOutdated")]
        public Output<bool?> IsOutdated { get; private set; } = null!;

        /// <summary>
        /// The container cluster's name. It is the only in one Alicloud account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        /// <summary>
        /// Whether to create the default simple routing Server Load Balancer instance for the cluster. The default value is true.
        /// </summary>
        [Output("needSlb")]
        public Output<bool?> NeedSlb { get; private set; } = null!;

        /// <summary>
        /// The ECS node number of the container cluster. Its value choices are 1~50, and default to 1.
        /// </summary>
        [Output("nodeNumber")]
        public Output<int?> NodeNumber { get; private set; } = null!;

        /// <summary>
        /// List of cluster nodes. It contains several attributes to `Block Nodes`.
        /// </summary>
        [Output("nodes")]
        public Output<ImmutableArray<Outputs.SwarmNode>> Nodes { get; private set; } = null!;

        /// <summary>
        /// The password of ECS instance node.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// Whether to release EIP after creating swarm cluster successfully. Default to false.
        /// </summary>
        [Output("releaseEip")]
        public Output<bool?> ReleaseEip { get; private set; } = null!;

        /// <summary>
        /// The ID of security group where the current cluster worker node is located.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// Field 'size' has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.
        /// </summary>
        [Output("size")]
        public Output<int?> Size { get; private set; } = null!;

        /// <summary>
        /// The ID of load balancer where the current cluster worker node is located.
        /// </summary>
        [Output("slbId")]
        public Output<string> SlbId { get; private set; } = null!;

        /// <summary>
        /// The ID of VPC where the current cluster is located.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The password of ECS instance node. If it is not specified, the container cluster's network mode will be `Classic`.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a Swarm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Swarm(string name, SwarmArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cs/swarm:Swarm", name, args ?? new SwarmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Swarm(string name, Input<string> id, SwarmState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cs/swarm:Swarm", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Swarm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Swarm Get(string name, Input<string> id, SwarmState? state = null, CustomResourceOptions? options = null)
        {
            return new Swarm(name, id, state, options);
        }
    }

    public sealed class SwarmArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CIDR block for the Container. It can not be same as the CIDR used by the VPC.
        /// Valid value:
        /// - 192.168.0.0/16
        /// - 172.19-30.0.0/16
        /// - 10.0.0.0/16
        /// </summary>
        [Input("cidrBlock", required: true)]
        public Input<string> CidrBlock { get; set; } = null!;

        /// <summary>
        /// The data disk category of ECS instance node. Its valid value are `cloud`, `cloud_ssd`, `cloud_essd`, `ephemeral_essd` and `cloud_efficiency`. Default to `cloud_efficiency`.
        /// </summary>
        [Input("diskCategory")]
        public Input<string>? DiskCategory { get; set; }

        /// <summary>
        /// The data disk size of ECS instance node. Its valid value is 20~32768 GB. Default to 20.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// The image ID of ECS instance node used. Default to System automate allocated.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// The type of ECS instance node.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// Whether to use outdated instance type. Default to false.
        /// </summary>
        [Input("isOutdated")]
        public Input<bool>? IsOutdated { get; set; }

        /// <summary>
        /// The container cluster's name. It is the only in one Alicloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// Whether to create the default simple routing Server Load Balancer instance for the cluster. The default value is true.
        /// </summary>
        [Input("needSlb")]
        public Input<bool>? NeedSlb { get; set; }

        /// <summary>
        /// The ECS node number of the container cluster. Its value choices are 1~50, and default to 1.
        /// </summary>
        [Input("nodeNumber")]
        public Input<int>? NodeNumber { get; set; }

        /// <summary>
        /// The password of ECS instance node.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// Whether to release EIP after creating swarm cluster successfully. Default to false.
        /// </summary>
        [Input("releaseEip")]
        public Input<bool>? ReleaseEip { get; set; }

        /// <summary>
        /// Field 'size' has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The password of ECS instance node. If it is not specified, the container cluster's network mode will be `Classic`.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        public SwarmArgs()
        {
        }
    }

    public sealed class SwarmState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The nodes agent version.
        /// </summary>
        [Input("agentVersion")]
        public Input<string>? AgentVersion { get; set; }

        /// <summary>
        /// The CIDR block for the Container. It can not be same as the CIDR used by the VPC.
        /// Valid value:
        /// - 192.168.0.0/16
        /// - 172.19-30.0.0/16
        /// - 10.0.0.0/16
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// The data disk category of ECS instance node. Its valid value are `cloud`, `cloud_ssd`, `cloud_essd`, `ephemeral_essd` and `cloud_efficiency`. Default to `cloud_efficiency`.
        /// </summary>
        [Input("diskCategory")]
        public Input<string>? DiskCategory { get; set; }

        /// <summary>
        /// The data disk size of ECS instance node. Its valid value is 20~32768 GB. Default to 20.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// The image ID of ECS instance node used. Default to System automate allocated.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// The type of ECS instance node.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// Whether to use outdated instance type. Default to false.
        /// </summary>
        [Input("isOutdated")]
        public Input<bool>? IsOutdated { get; set; }

        /// <summary>
        /// The container cluster's name. It is the only in one Alicloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        /// <summary>
        /// Whether to create the default simple routing Server Load Balancer instance for the cluster. The default value is true.
        /// </summary>
        [Input("needSlb")]
        public Input<bool>? NeedSlb { get; set; }

        /// <summary>
        /// The ECS node number of the container cluster. Its value choices are 1~50, and default to 1.
        /// </summary>
        [Input("nodeNumber")]
        public Input<int>? NodeNumber { get; set; }

        [Input("nodes")]
        private InputList<Inputs.SwarmNodeGetArgs>? _nodes;

        /// <summary>
        /// List of cluster nodes. It contains several attributes to `Block Nodes`.
        /// </summary>
        public InputList<Inputs.SwarmNodeGetArgs> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<Inputs.SwarmNodeGetArgs>());
            set => _nodes = value;
        }

        /// <summary>
        /// The password of ECS instance node.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Whether to release EIP after creating swarm cluster successfully. Default to false.
        /// </summary>
        [Input("releaseEip")]
        public Input<bool>? ReleaseEip { get; set; }

        /// <summary>
        /// The ID of security group where the current cluster worker node is located.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// Field 'size' has been deprecated from provider version 1.9.1. New field 'node_number' replaces it.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The ID of load balancer where the current cluster worker node is located.
        /// </summary>
        [Input("slbId")]
        public Input<string>? SlbId { get; set; }

        /// <summary>
        /// The ID of VPC where the current cluster is located.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The password of ECS instance node. If it is not specified, the container cluster's network mode will be `Classic`.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public SwarmState()
        {
        }
    }
}