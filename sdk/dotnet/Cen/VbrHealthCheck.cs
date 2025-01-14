// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    /// <summary>
    /// This topic describes how to configure the health check feature for a Cloud Enterprise Network (CEN) instance.
    /// After you attach a Virtual Border Router (VBR) to the CEN instance and configure the health check feature, you can monitor the network conditions of the on-premises data center connected to the VBR.
    /// 
    /// For information about CEN VBR HealthCheck and how to use it, see [Manage CEN VBR HealthCheck](https://www.alibabacloud.com/help/en/doc-detail/71141.htm).
    /// 
    /// &gt; **NOTE:** Available in 1.88.0+
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
    ///     // Create a cen vbr HealrhCheck resource and use it.
    ///     var defaultInstance = new AliCloud.Cen.Instance("defaultInstance", new()
    ///     {
    ///         CenInstanceName = "test_name",
    ///     });
    /// 
    ///     var defaultInstanceAttachment = new AliCloud.Cen.InstanceAttachment("defaultInstanceAttachment", new()
    ///     {
    ///         InstanceId = defaultInstance.Id,
    ///         ChildInstanceId = "vbr-xxxxx",
    ///         ChildInstanceType = "VBR",
    ///         ChildInstanceRegionId = "cn-hangzhou",
    ///     });
    /// 
    ///     var defaultVbrHealthCheck = new AliCloud.Cen.VbrHealthCheck("defaultVbrHealthCheck", new()
    ///     {
    ///         CenId = defaultInstance.Id,
    ///         HealthCheckSourceIp = "192.168.1.2",
    ///         HealthCheckTargetIp = "10.0.0.2",
    ///         VbrInstanceId = "vbr-xxxxx",
    ///         VbrInstanceRegionId = "cn-hangzhou",
    ///         HealthCheckInterval = 2,
    ///         HealthyThreshold = 8,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             defaultInstanceAttachment,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CEN VBR HealthCheck can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cen/vbrHealthCheck:VbrHealthCheck example vbr-xxxxx:cn-hangzhou
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cen/vbrHealthCheck:VbrHealthCheck")]
    public partial class VbrHealthCheck : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        [Output("cenId")]
        public Output<string> CenId { get; private set; } = null!;

        /// <summary>
        /// Specifies the interval at which the health check sends continuous detection packets. Default value: 2. Value range: 2 to 3.
        /// </summary>
        [Output("healthCheckInterval")]
        public Output<int?> HealthCheckInterval { get; private set; } = null!;

        /// <summary>
        /// The source IP address of health checks.
        /// </summary>
        [Output("healthCheckSourceIp")]
        public Output<string?> HealthCheckSourceIp { get; private set; } = null!;

        /// <summary>
        /// The destination IP address of health checks.
        /// </summary>
        [Output("healthCheckTargetIp")]
        public Output<string> HealthCheckTargetIp { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of probe messages sent by the health check. Default value: 8. Value range: 3 to 8.
        /// </summary>
        [Output("healthyThreshold")]
        public Output<int?> HealthyThreshold { get; private set; } = null!;

        /// <summary>
        /// The ID of the VBR.
        /// </summary>
        [Output("vbrInstanceId")]
        public Output<string> VbrInstanceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the account to which the VBR belongs.
        /// </summary>
        [Output("vbrInstanceOwnerId")]
        public Output<int?> VbrInstanceOwnerId { get; private set; } = null!;

        /// <summary>
        /// The ID of the region to which the VBR belongs.
        /// 
        /// -&gt;**NOTE:** The `alicloud.cen.VbrHealthCheck` resource depends on the related `alicloud.cen.InstanceAttachment` resource.
        /// </summary>
        [Output("vbrInstanceRegionId")]
        public Output<string> VbrInstanceRegionId { get; private set; } = null!;


        /// <summary>
        /// Create a VbrHealthCheck resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VbrHealthCheck(string name, VbrHealthCheckArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/vbrHealthCheck:VbrHealthCheck", name, args ?? new VbrHealthCheckArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VbrHealthCheck(string name, Input<string> id, VbrHealthCheckState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/vbrHealthCheck:VbrHealthCheck", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VbrHealthCheck resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VbrHealthCheck Get(string name, Input<string> id, VbrHealthCheckState? state = null, CustomResourceOptions? options = null)
        {
            return new VbrHealthCheck(name, id, state, options);
        }
    }

    public sealed class VbrHealthCheckArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        [Input("cenId", required: true)]
        public Input<string> CenId { get; set; } = null!;

        /// <summary>
        /// Specifies the interval at which the health check sends continuous detection packets. Default value: 2. Value range: 2 to 3.
        /// </summary>
        [Input("healthCheckInterval")]
        public Input<int>? HealthCheckInterval { get; set; }

        /// <summary>
        /// The source IP address of health checks.
        /// </summary>
        [Input("healthCheckSourceIp")]
        public Input<string>? HealthCheckSourceIp { get; set; }

        /// <summary>
        /// The destination IP address of health checks.
        /// </summary>
        [Input("healthCheckTargetIp", required: true)]
        public Input<string> HealthCheckTargetIp { get; set; } = null!;

        /// <summary>
        /// Specifies the number of probe messages sent by the health check. Default value: 8. Value range: 3 to 8.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// The ID of the VBR.
        /// </summary>
        [Input("vbrInstanceId", required: true)]
        public Input<string> VbrInstanceId { get; set; } = null!;

        /// <summary>
        /// The ID of the account to which the VBR belongs.
        /// </summary>
        [Input("vbrInstanceOwnerId")]
        public Input<int>? VbrInstanceOwnerId { get; set; }

        /// <summary>
        /// The ID of the region to which the VBR belongs.
        /// 
        /// -&gt;**NOTE:** The `alicloud.cen.VbrHealthCheck` resource depends on the related `alicloud.cen.InstanceAttachment` resource.
        /// </summary>
        [Input("vbrInstanceRegionId", required: true)]
        public Input<string> VbrInstanceRegionId { get; set; } = null!;

        public VbrHealthCheckArgs()
        {
        }
        public static new VbrHealthCheckArgs Empty => new VbrHealthCheckArgs();
    }

    public sealed class VbrHealthCheckState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the CEN instance.
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        /// <summary>
        /// Specifies the interval at which the health check sends continuous detection packets. Default value: 2. Value range: 2 to 3.
        /// </summary>
        [Input("healthCheckInterval")]
        public Input<int>? HealthCheckInterval { get; set; }

        /// <summary>
        /// The source IP address of health checks.
        /// </summary>
        [Input("healthCheckSourceIp")]
        public Input<string>? HealthCheckSourceIp { get; set; }

        /// <summary>
        /// The destination IP address of health checks.
        /// </summary>
        [Input("healthCheckTargetIp")]
        public Input<string>? HealthCheckTargetIp { get; set; }

        /// <summary>
        /// Specifies the number of probe messages sent by the health check. Default value: 8. Value range: 3 to 8.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// The ID of the VBR.
        /// </summary>
        [Input("vbrInstanceId")]
        public Input<string>? VbrInstanceId { get; set; }

        /// <summary>
        /// The ID of the account to which the VBR belongs.
        /// </summary>
        [Input("vbrInstanceOwnerId")]
        public Input<int>? VbrInstanceOwnerId { get; set; }

        /// <summary>
        /// The ID of the region to which the VBR belongs.
        /// 
        /// -&gt;**NOTE:** The `alicloud.cen.VbrHealthCheck` resource depends on the related `alicloud.cen.InstanceAttachment` resource.
        /// </summary>
        [Input("vbrInstanceRegionId")]
        public Input<string>? VbrInstanceRegionId { get; set; }

        public VbrHealthCheckState()
        {
        }
        public static new VbrHealthCheckState Empty => new VbrHealthCheckState();
    }
}
