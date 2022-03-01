// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dns
{
    public static class GetInstances
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(AliCloud.Dns.GetInstances.InvokeAsync(new AliCloud.Dns.GetInstancesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "dns-cn-oew1npk****",
        ///             },
        ///         }));
        ///         this.FirstInstanceId = example.Apply(example =&gt; example.Instances?[0]?.Id);
        ///     }
        /// 
        ///     [Output("firstInstanceId")]
        ///     public Output&lt;string&gt; FirstInstanceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstancesResult> InvokeAsync(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("alicloud:dns/getInstances:getInstances", args ?? new GetInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(AliCloud.Dns.GetInstances.InvokeAsync(new AliCloud.Dns.GetInstancesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "dns-cn-oew1npk****",
        ///             },
        ///         }));
        ///         this.FirstInstanceId = example.Apply(example =&gt; example.Instances?[0]?.Id);
        ///     }
        /// 
        ///     [Output("firstInstanceId")]
        ///     public Output&lt;string&gt; FirstInstanceId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstancesResult> Invoke(GetInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstancesResult>("alicloud:dns/getInstances:getInstances", args ?? new GetInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstancesArgs : Pulumi.InvokeArgs
    {
        [Input("domainType")]
        public string? DomainType { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of instance IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("lang")]
        public string? Lang { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("userClientIp")]
        public string? UserClientIp { get; set; }

        public GetInstancesArgs()
        {
        }
    }

    public sealed class GetInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("domainType")]
        public Input<string>? DomainType { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of instance IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("lang")]
        public Input<string>? Lang { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("userClientIp")]
        public Input<string>? UserClientIp { get; set; }

        public GetInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstancesResult
    {
        public readonly string? DomainType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of instance IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of instances. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstancesInstanceResult> Instances;
        public readonly string? Lang;
        public readonly string? OutputFile;
        public readonly string? UserClientIp;

        [OutputConstructor]
        private GetInstancesResult(
            string? domainType,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetInstancesInstanceResult> instances,

            string? lang,

            string? outputFile,

            string? userClientIp)
        {
            DomainType = domainType;
            Id = id;
            Ids = ids;
            Instances = instances;
            Lang = lang;
            OutputFile = outputFile;
            UserClientIp = userClientIp;
        }
    }
}
