// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cdn
{
    /// <summary>
    /// Provides a CDN Fc Trigger resource.
    /// 
    /// For information about CDN Fc Trigger and how to use it, see [What is Fc Trigger](https://www.alibabacloud.com/help/zh/alibaba-cloud-cdn/latest/add-function-calculation-trigger).
    /// 
    /// &gt; **NOTE:** Available in v1.165.0+.
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
    ///     var defaultAccount = AliCloud.GetAccount.Invoke();
    /// 
    ///     var defaultRegions = AliCloud.GetRegions.Invoke(new()
    ///     {
    ///         Current = true,
    ///     });
    /// 
    ///     var example = new AliCloud.Cdn.FcTrigger("example", new()
    ///     {
    ///         EventMetaName = "LogFileCreated",
    ///         EventMetaVersion = "1.0.0",
    ///         Notes = "example_value",
    ///         RoleArn = $"acs:ram::{defaultAccount.Apply(getAccountResult =&gt; getAccountResult.Id)}:role/aliyuncdneventnotificationrole",
    ///         SourceArn = $"acs:cdn:*:{defaultAccount.Apply(getAccountResult =&gt; getAccountResult.Id)}:domain/example.com",
    ///         TriggerArn = Output.Tuple(defaultRegions, defaultAccount).Apply(values =&gt;
    ///         {
    ///             var defaultRegions = values.Item1;
    ///             var defaultAccount = values.Item2;
    ///             return $"acs:fc:{defaultRegions.Apply(getRegionsResult =&gt; getRegionsResult.Regions[0]?.Id)}:{defaultAccount.Apply(getAccountResult =&gt; getAccountResult.Id)}:services/FCTestService/functions/printEvent/triggers/testtrigger";
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CDN Fc Trigger can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cdn/fcTrigger:FcTrigger example &lt;trigger_arn&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cdn/fcTrigger:FcTrigger")]
    public partial class FcTrigger : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the Event.
        /// </summary>
        [Output("eventMetaName")]
        public Output<string> EventMetaName { get; private set; } = null!;

        /// <summary>
        /// The version of the Event.
        /// </summary>
        [Output("eventMetaVersion")]
        public Output<string> EventMetaVersion { get; private set; } = null!;

        /// <summary>
        /// The function arn. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`.
        /// </summary>
        [Output("functionArn")]
        public Output<string?> FunctionArn { get; private set; } = null!;

        /// <summary>
        /// The Note information.
        /// </summary>
        [Output("notes")]
        public Output<string> Notes { get; private set; } = null!;

        /// <summary>
        /// The role authorized by RAM. The value formats as `acs:ram::{AccountID}:role/{RoleName}`.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// Resources and filters for event listening. The value formats as `acs:cdn:{RegionID}:{AccountID}:{Filter}`.
        /// </summary>
        [Output("sourceArn")]
        public Output<string> SourceArn { get; private set; } = null!;

        /// <summary>
        /// The trigger corresponding to the function Compute Service. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`. See [Create a CDN Fc Trigger](https://www.alibabacloud.com/help/zh/alibaba-cloud-cdn/latest/add-function-calculation-trigger) for more details.
        /// </summary>
        [Output("triggerArn")]
        public Output<string> TriggerArn { get; private set; } = null!;


        /// <summary>
        /// Create a FcTrigger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FcTrigger(string name, FcTriggerArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cdn/fcTrigger:FcTrigger", name, args ?? new FcTriggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FcTrigger(string name, Input<string> id, FcTriggerState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cdn/fcTrigger:FcTrigger", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FcTrigger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FcTrigger Get(string name, Input<string> id, FcTriggerState? state = null, CustomResourceOptions? options = null)
        {
            return new FcTrigger(name, id, state, options);
        }
    }

    public sealed class FcTriggerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Event.
        /// </summary>
        [Input("eventMetaName", required: true)]
        public Input<string> EventMetaName { get; set; } = null!;

        /// <summary>
        /// The version of the Event.
        /// </summary>
        [Input("eventMetaVersion", required: true)]
        public Input<string> EventMetaVersion { get; set; } = null!;

        /// <summary>
        /// The function arn. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`.
        /// </summary>
        [Input("functionArn")]
        public Input<string>? FunctionArn { get; set; }

        /// <summary>
        /// The Note information.
        /// </summary>
        [Input("notes", required: true)]
        public Input<string> Notes { get; set; } = null!;

        /// <summary>
        /// The role authorized by RAM. The value formats as `acs:ram::{AccountID}:role/{RoleName}`.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// Resources and filters for event listening. The value formats as `acs:cdn:{RegionID}:{AccountID}:{Filter}`.
        /// </summary>
        [Input("sourceArn", required: true)]
        public Input<string> SourceArn { get; set; } = null!;

        /// <summary>
        /// The trigger corresponding to the function Compute Service. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`. See [Create a CDN Fc Trigger](https://www.alibabacloud.com/help/zh/alibaba-cloud-cdn/latest/add-function-calculation-trigger) for more details.
        /// </summary>
        [Input("triggerArn", required: true)]
        public Input<string> TriggerArn { get; set; } = null!;

        public FcTriggerArgs()
        {
        }
        public static new FcTriggerArgs Empty => new FcTriggerArgs();
    }

    public sealed class FcTriggerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Event.
        /// </summary>
        [Input("eventMetaName")]
        public Input<string>? EventMetaName { get; set; }

        /// <summary>
        /// The version of the Event.
        /// </summary>
        [Input("eventMetaVersion")]
        public Input<string>? EventMetaVersion { get; set; }

        /// <summary>
        /// The function arn. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`.
        /// </summary>
        [Input("functionArn")]
        public Input<string>? FunctionArn { get; set; }

        /// <summary>
        /// The Note information.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// The role authorized by RAM. The value formats as `acs:ram::{AccountID}:role/{RoleName}`.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// Resources and filters for event listening. The value formats as `acs:cdn:{RegionID}:{AccountID}:{Filter}`.
        /// </summary>
        [Input("sourceArn")]
        public Input<string>? SourceArn { get; set; }

        /// <summary>
        /// The trigger corresponding to the function Compute Service. The value formats as `acs:fc:{RegionID}:{AccountID}:{Filter}`. See [Create a CDN Fc Trigger](https://www.alibabacloud.com/help/zh/alibaba-cloud-cdn/latest/add-function-calculation-trigger) for more details.
        /// </summary>
        [Input("triggerArn")]
        public Input<string>? TriggerArn { get; set; }

        public FcTriggerState()
        {
        }
        public static new FcTriggerState Empty => new FcTriggerState();
    }
}
