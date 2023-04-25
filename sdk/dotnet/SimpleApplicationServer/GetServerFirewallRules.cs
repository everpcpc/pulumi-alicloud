// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.SimpleApplicationServer
{
    public static class GetServerFirewallRules
    {
        /// <summary>
        /// This data source provides the Simple Application Server Firewall Rules of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.143.0+.
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
        ///     var ids = AliCloud.SimpleApplicationServer.GetServerFirewallRules.Invoke(new()
        ///     {
        ///         InstanceId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["simpleApplicationServerFirewallRuleId1"] = ids.Apply(getServerFirewallRulesResult =&gt; getServerFirewallRulesResult.Rules[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServerFirewallRulesResult> InvokeAsync(GetServerFirewallRulesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServerFirewallRulesResult>("alicloud:simpleapplicationserver/getServerFirewallRules:getServerFirewallRules", args ?? new GetServerFirewallRulesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Simple Application Server Firewall Rules of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.143.0+.
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
        ///     var ids = AliCloud.SimpleApplicationServer.GetServerFirewallRules.Invoke(new()
        ///     {
        ///         InstanceId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value-1",
        ///             "example_value-2",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["simpleApplicationServerFirewallRuleId1"] = ids.Apply(getServerFirewallRulesResult =&gt; getServerFirewallRulesResult.Rules[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetServerFirewallRulesResult> Invoke(GetServerFirewallRulesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerFirewallRulesResult>("alicloud:simpleapplicationserver/getServerFirewallRules:getServerFirewallRules", args ?? new GetServerFirewallRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerFirewallRulesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Firewall Rule IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Alibaba Cloud simple application server instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetServerFirewallRulesArgs()
        {
        }
        public static new GetServerFirewallRulesArgs Empty => new GetServerFirewallRulesArgs();
    }

    public sealed class GetServerFirewallRulesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Firewall Rule IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Alibaba Cloud simple application server instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetServerFirewallRulesInvokeArgs()
        {
        }
        public static new GetServerFirewallRulesInvokeArgs Empty => new GetServerFirewallRulesInvokeArgs();
    }


    [OutputType]
    public sealed class GetServerFirewallRulesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string InstanceId;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetServerFirewallRulesRuleResult> Rules;

        [OutputConstructor]
        private GetServerFirewallRulesResult(
            string id,

            ImmutableArray<string> ids,

            string instanceId,

            string? outputFile,

            ImmutableArray<Outputs.GetServerFirewallRulesRuleResult> rules)
        {
            Id = id;
            Ids = ids;
            InstanceId = instanceId;
            OutputFile = outputFile;
            Rules = rules;
        }
    }
}
