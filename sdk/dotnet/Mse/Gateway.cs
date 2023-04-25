// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Mse
{
    /// <summary>
    /// Provides a Microservice Engine (MSE) Gateway resource.
    /// 
    /// For information about Microservice Engine (MSE) Gateway and how to use it, see [What is Gateway](https://help.aliyun.com/document_detail/347638.html).
    /// 
    /// &gt; **NOTE:** Available in v1.157.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var defaultNetworks = AliCloud.Vpc.GetNetworks.Invoke(new()
    ///     {
    ///         NameRegex = "default-NODELETING",
    ///     });
    /// 
    ///     var defaultSwitches = AliCloud.Vpc.GetSwitches.Invoke(new()
    ///     {
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var example = new AliCloud.Mse.Gateway("example", new()
    ///     {
    ///         GatewayName = "example_value",
    ///         Replica = 2,
    ///         Spec = "MSE_GTW_2_4_200_c",
    ///         VswitchId = defaultSwitches.Apply(getSwitchesResult =&gt; getSwitchesResult.Ids[0]),
    ///         BackupVswitchId = defaultSwitches.Apply(getSwitchesResult =&gt; getSwitchesResult.Ids[1]),
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Microservice Engine (MSE) Gateway can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:mse/gateway:Gateway example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:mse/gateway:Gateway")]
    public partial class Gateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The backup vswitch id.
        /// </summary>
        [Output("backupVswitchId")]
        public Output<string?> BackupVswitchId { get; private set; } = null!;

        /// <summary>
        /// Whether to delete the SLB purchased on behalf of the gateway at the same time.
        /// </summary>
        [Output("deleteSlb")]
        public Output<bool?> DeleteSlb { get; private set; } = null!;

        /// <summary>
        /// Whether the enterprise security group type.
        /// </summary>
        [Output("enterpriseSecurityGroup")]
        public Output<bool?> EnterpriseSecurityGroup { get; private set; } = null!;

        /// <summary>
        /// The name of the Gateway .
        /// </summary>
        [Output("gatewayName")]
        public Output<string?> GatewayName { get; private set; } = null!;

        /// <summary>
        /// Public network SLB specifications.
        /// </summary>
        [Output("internetSlbSpec")]
        public Output<string?> InternetSlbSpec { get; private set; } = null!;

        /// <summary>
        /// Number of Gateway Nodes.
        /// </summary>
        [Output("replica")]
        public Output<int> Replica { get; private set; } = null!;

        /// <summary>
        /// A list of gateway Slb.
        /// </summary>
        [Output("slbLists")]
        public Output<ImmutableArray<Outputs.GatewaySlbList>> SlbLists { get; private set; } = null!;

        /// <summary>
        /// Private network SLB specifications.
        /// </summary>
        [Output("slbSpec")]
        public Output<string?> SlbSpec { get; private set; } = null!;

        /// <summary>
        /// Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
        /// </summary>
        [Output("spec")]
        public Output<string> Spec { get; private set; } = null!;

        /// <summary>
        /// The status of the gateway.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of the vpc.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The ID of the vswitch.
        /// </summary>
        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a Gateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Gateway(string name, GatewayArgs args, CustomResourceOptions? options = null)
            : base("alicloud:mse/gateway:Gateway", name, args ?? new GatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Gateway(string name, Input<string> id, GatewayState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:mse/gateway:Gateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Gateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Gateway Get(string name, Input<string> id, GatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new Gateway(name, id, state, options);
        }
    }

    public sealed class GatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The backup vswitch id.
        /// </summary>
        [Input("backupVswitchId")]
        public Input<string>? BackupVswitchId { get; set; }

        /// <summary>
        /// Whether to delete the SLB purchased on behalf of the gateway at the same time.
        /// </summary>
        [Input("deleteSlb")]
        public Input<bool>? DeleteSlb { get; set; }

        /// <summary>
        /// Whether the enterprise security group type.
        /// </summary>
        [Input("enterpriseSecurityGroup")]
        public Input<bool>? EnterpriseSecurityGroup { get; set; }

        /// <summary>
        /// The name of the Gateway .
        /// </summary>
        [Input("gatewayName")]
        public Input<string>? GatewayName { get; set; }

        /// <summary>
        /// Public network SLB specifications.
        /// </summary>
        [Input("internetSlbSpec")]
        public Input<string>? InternetSlbSpec { get; set; }

        /// <summary>
        /// Number of Gateway Nodes.
        /// </summary>
        [Input("replica", required: true)]
        public Input<int> Replica { get; set; } = null!;

        /// <summary>
        /// Private network SLB specifications.
        /// </summary>
        [Input("slbSpec")]
        public Input<string>? SlbSpec { get; set; }

        /// <summary>
        /// Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
        /// </summary>
        [Input("spec", required: true)]
        public Input<string> Spec { get; set; } = null!;

        /// <summary>
        /// The ID of the vpc.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// The ID of the vswitch.
        /// </summary>
        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        public GatewayArgs()
        {
        }
        public static new GatewayArgs Empty => new GatewayArgs();
    }

    public sealed class GatewayState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The backup vswitch id.
        /// </summary>
        [Input("backupVswitchId")]
        public Input<string>? BackupVswitchId { get; set; }

        /// <summary>
        /// Whether to delete the SLB purchased on behalf of the gateway at the same time.
        /// </summary>
        [Input("deleteSlb")]
        public Input<bool>? DeleteSlb { get; set; }

        /// <summary>
        /// Whether the enterprise security group type.
        /// </summary>
        [Input("enterpriseSecurityGroup")]
        public Input<bool>? EnterpriseSecurityGroup { get; set; }

        /// <summary>
        /// The name of the Gateway .
        /// </summary>
        [Input("gatewayName")]
        public Input<string>? GatewayName { get; set; }

        /// <summary>
        /// Public network SLB specifications.
        /// </summary>
        [Input("internetSlbSpec")]
        public Input<string>? InternetSlbSpec { get; set; }

        /// <summary>
        /// Number of Gateway Nodes.
        /// </summary>
        [Input("replica")]
        public Input<int>? Replica { get; set; }

        [Input("slbLists")]
        private InputList<Inputs.GatewaySlbListGetArgs>? _slbLists;

        /// <summary>
        /// A list of gateway Slb.
        /// </summary>
        public InputList<Inputs.GatewaySlbListGetArgs> SlbLists
        {
            get => _slbLists ?? (_slbLists = new InputList<Inputs.GatewaySlbListGetArgs>());
            set => _slbLists = value;
        }

        /// <summary>
        /// Private network SLB specifications.
        /// </summary>
        [Input("slbSpec")]
        public Input<string>? SlbSpec { get; set; }

        /// <summary>
        /// Gateway Node Specifications. Valid values: `MSE_GTW_2_4_200_c`, `MSE_GTW_4_8_200_c`, `MSE_GTW_8_16_200_c`, `MSE_GTW_16_32_200_c`.
        /// </summary>
        [Input("spec")]
        public Input<string>? Spec { get; set; }

        /// <summary>
        /// The status of the gateway.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the vpc.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The ID of the vswitch.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public GatewayState()
        {
        }
        public static new GatewayState Empty => new GatewayState();
    }
}
