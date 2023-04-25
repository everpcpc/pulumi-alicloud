// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga
{
    public static class GetForwardingRules
    {
        /// <summary>
        /// This data source provides the Global Accelerator (GA) Forwarding Rules of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.120.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var example = AliCloud.Ga.GetForwardingRules.Invoke(new()
        ///     {
        ///         AcceleratorId = "example_value",
        ///         ListenerId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstGaForwardingRuleId"] = example.Apply(getForwardingRulesResult =&gt; getForwardingRulesResult.ForwardingRules[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetForwardingRulesResult> InvokeAsync(GetForwardingRulesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetForwardingRulesResult>("alicloud:ga/getForwardingRules:getForwardingRules", args ?? new GetForwardingRulesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Global Accelerator (GA) Forwarding Rules of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.120.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var example = AliCloud.Ga.GetForwardingRules.Invoke(new()
        ///     {
        ///         AcceleratorId = "example_value",
        ///         ListenerId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstGaForwardingRuleId"] = example.Apply(getForwardingRulesResult =&gt; getForwardingRulesResult.ForwardingRules[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetForwardingRulesResult> Invoke(GetForwardingRulesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetForwardingRulesResult>("alicloud:ga/getForwardingRules:getForwardingRules", args ?? new GetForwardingRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetForwardingRulesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Global Accelerator instance.
        /// </summary>
        [Input("acceleratorId", required: true)]
        public string AcceleratorId { get; set; } = null!;

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Forwarding Rule IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The ID of the listener.
        /// </summary>
        [Input("listenerId", required: true)]
        public string ListenerId { get; set; } = null!;

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of the acceleration region. Valid values: `active`, `configuring`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetForwardingRulesArgs()
        {
        }
        public static new GetForwardingRulesArgs Empty => new GetForwardingRulesArgs();
    }

    public sealed class GetForwardingRulesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Global Accelerator instance.
        /// </summary>
        [Input("acceleratorId", required: true)]
        public Input<string> AcceleratorId { get; set; } = null!;

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Forwarding Rule IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The ID of the listener.
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The status of the acceleration region. Valid values: `active`, `configuring`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetForwardingRulesInvokeArgs()
        {
        }
        public static new GetForwardingRulesInvokeArgs Empty => new GetForwardingRulesInvokeArgs();
    }


    [OutputType]
    public sealed class GetForwardingRulesResult
    {
        public readonly string AcceleratorId;
        public readonly ImmutableArray<Outputs.GetForwardingRulesForwardingRuleResult> ForwardingRules;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string ListenerId;
        public readonly string? OutputFile;
        public readonly string? Status;

        [OutputConstructor]
        private GetForwardingRulesResult(
            string acceleratorId,

            ImmutableArray<Outputs.GetForwardingRulesForwardingRuleResult> forwardingRules,

            string id,

            ImmutableArray<string> ids,

            string listenerId,

            string? outputFile,

            string? status)
        {
            AcceleratorId = acceleratorId;
            ForwardingRules = forwardingRules;
            Id = id;
            Ids = ids;
            ListenerId = listenerId;
            OutputFile = outputFile;
            Status = status;
        }
    }
}
