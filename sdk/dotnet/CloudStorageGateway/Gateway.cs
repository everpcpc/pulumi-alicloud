// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CloudStorageGateway
{
    /// <summary>
    /// Provides a Cloud Storage Gateway: Gateway resource.
    /// 
    /// For information about Cloud Storage Gateway Gateway and how to use it, see [What is Gateway](https://www.alibabacloud.com/help/en/doc-detail/53972.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.132.0+.
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
    ///     var vpc = new AliCloud.Vpc.Network("vpc", new()
    ///     {
    ///         VpcName = "tf_test_foo",
    ///         CidrBlock = "172.16.0.0/12",
    ///     });
    /// 
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VpcId = vpc.Id,
    ///         CidrBlock = "172.16.0.0/21",
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         VswitchName = "tf-testAccCsgName",
    ///     });
    /// 
    ///     var example = new AliCloud.CloudStorageGateway.StorageBundle("example", new()
    ///     {
    ///         StorageBundleName = "example_value",
    ///     });
    /// 
    ///     var defaultGateway = new AliCloud.CloudStorageGateway.Gateway("defaultGateway", new()
    ///     {
    ///         Description = "tf-acctestDesalone",
    ///         GatewayClass = "Standard",
    ///         Type = "File",
    ///         PaymentType = "PayAsYouGo",
    ///         VswitchId = defaultSwitch.Id,
    ///         ReleaseAfterExpiration = false,
    ///         PublicNetworkBandwidth = 40,
    ///         StorageBundleId = example.Id,
    ///         Location = "Cloud",
    ///         GatewayName = "tf-acctestGatewayName",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cloud Storage Gateway Gateway can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cloudstoragegateway/gateway:Gateway example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cloudstoragegateway/gateway:Gateway")]
    public partial class Gateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// the description of gateway.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// the gateway class. the valid values: `Basic`, `Standard`,`Enhanced`,`Advanced`
        /// </summary>
        [Output("gatewayClass")]
        public Output<string?> GatewayClass { get; private set; } = null!;

        /// <summary>
        /// the name of gateway.
        /// </summary>
        [Output("gatewayName")]
        public Output<string> GatewayName { get; private set; } = null!;

        /// <summary>
        /// gateway location. the valid values: `Cloud`, `On_Premise`.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The Payment type of gateway. The valid value: `PayAsYouGo`.
        /// </summary>
        [Output("paymentType")]
        public Output<string?> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The public network bandwidth of gateway. Valid values between `5` and `200`. Defaults to `5`.
        /// </summary>
        [Output("publicNetworkBandwidth")]
        public Output<int> PublicNetworkBandwidth { get; private set; } = null!;

        /// <summary>
        /// The reason detail of gateway.
        /// </summary>
        [Output("reasonDetail")]
        public Output<string?> ReasonDetail { get; private set; } = null!;

        /// <summary>
        /// The reason type when user deletes the gateway.
        /// </summary>
        [Output("reasonType")]
        public Output<string?> ReasonType { get; private set; } = null!;

        /// <summary>
        /// Whether to release the gateway due to expiration.
        /// </summary>
        [Output("releaseAfterExpiration")]
        public Output<bool?> ReleaseAfterExpiration { get; private set; } = null!;

        /// <summary>
        /// gateway status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// storage bundle id.
        /// </summary>
        [Output("storageBundleId")]
        public Output<string> StorageBundleId { get; private set; } = null!;

        /// <summary>
        /// gateway type. the valid values: `Type`, `Iscsi`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The vswitch id of gateway.
        /// </summary>
        [Output("vswitchId")]
        public Output<string?> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a Gateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Gateway(string name, GatewayArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cloudstoragegateway/gateway:Gateway", name, args ?? new GatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Gateway(string name, Input<string> id, GatewayState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cloudstoragegateway/gateway:Gateway", name, state, MakeResourceOptions(options, id))
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
        /// the description of gateway.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// the gateway class. the valid values: `Basic`, `Standard`,`Enhanced`,`Advanced`
        /// </summary>
        [Input("gatewayClass")]
        public Input<string>? GatewayClass { get; set; }

        /// <summary>
        /// the name of gateway.
        /// </summary>
        [Input("gatewayName", required: true)]
        public Input<string> GatewayName { get; set; } = null!;

        /// <summary>
        /// gateway location. the valid values: `Cloud`, `On_Premise`.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The Payment type of gateway. The valid value: `PayAsYouGo`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The public network bandwidth of gateway. Valid values between `5` and `200`. Defaults to `5`.
        /// </summary>
        [Input("publicNetworkBandwidth")]
        public Input<int>? PublicNetworkBandwidth { get; set; }

        /// <summary>
        /// The reason detail of gateway.
        /// </summary>
        [Input("reasonDetail")]
        public Input<string>? ReasonDetail { get; set; }

        /// <summary>
        /// The reason type when user deletes the gateway.
        /// </summary>
        [Input("reasonType")]
        public Input<string>? ReasonType { get; set; }

        /// <summary>
        /// Whether to release the gateway due to expiration.
        /// </summary>
        [Input("releaseAfterExpiration")]
        public Input<bool>? ReleaseAfterExpiration { get; set; }

        /// <summary>
        /// storage bundle id.
        /// </summary>
        [Input("storageBundleId", required: true)]
        public Input<string> StorageBundleId { get; set; } = null!;

        /// <summary>
        /// gateway type. the valid values: `Type`, `Iscsi`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The vswitch id of gateway.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public GatewayArgs()
        {
        }
        public static new GatewayArgs Empty => new GatewayArgs();
    }

    public sealed class GatewayState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the description of gateway.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// the gateway class. the valid values: `Basic`, `Standard`,`Enhanced`,`Advanced`
        /// </summary>
        [Input("gatewayClass")]
        public Input<string>? GatewayClass { get; set; }

        /// <summary>
        /// the name of gateway.
        /// </summary>
        [Input("gatewayName")]
        public Input<string>? GatewayName { get; set; }

        /// <summary>
        /// gateway location. the valid values: `Cloud`, `On_Premise`.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The Payment type of gateway. The valid value: `PayAsYouGo`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The public network bandwidth of gateway. Valid values between `5` and `200`. Defaults to `5`.
        /// </summary>
        [Input("publicNetworkBandwidth")]
        public Input<int>? PublicNetworkBandwidth { get; set; }

        /// <summary>
        /// The reason detail of gateway.
        /// </summary>
        [Input("reasonDetail")]
        public Input<string>? ReasonDetail { get; set; }

        /// <summary>
        /// The reason type when user deletes the gateway.
        /// </summary>
        [Input("reasonType")]
        public Input<string>? ReasonType { get; set; }

        /// <summary>
        /// Whether to release the gateway due to expiration.
        /// </summary>
        [Input("releaseAfterExpiration")]
        public Input<bool>? ReleaseAfterExpiration { get; set; }

        /// <summary>
        /// gateway status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// storage bundle id.
        /// </summary>
        [Input("storageBundleId")]
        public Input<string>? StorageBundleId { get; set; }

        /// <summary>
        /// gateway type. the valid values: `Type`, `Iscsi`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The vswitch id of gateway.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public GatewayState()
        {
        }
        public static new GatewayState Empty => new GatewayState();
    }
}
