// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ga
{
    public static class GetListeners
    {
        /// <summary>
        /// This data source provides the Ga Listeners of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.111.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic Usage
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(AliCloud.Ga.GetListeners.InvokeAsync(new AliCloud.Ga.GetListenersArgs
        ///         {
        ///             AcceleratorId = "example_value",
        ///             Ids = 
        ///             {
        ///                 "example_value",
        ///             },
        ///             NameRegex = "the_resource_name",
        ///         }));
        ///         this.FirstGaListenerId = example.Apply(example =&gt; example.Listeners[0].Id);
        ///     }
        /// 
        ///     [Output("firstGaListenerId")]
        ///     public Output&lt;string&gt; FirstGaListenerId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetListenersResult> InvokeAsync(GetListenersArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetListenersResult>("alicloud:ga/getListeners:getListeners", args ?? new GetListenersArgs(), options.WithVersion());
    }


    public sealed class GetListenersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The accelerator id.
        /// </summary>
        [Input("acceleratorId", required: true)]
        public string AcceleratorId { get; set; } = null!;

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Listener IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Listener name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of the listener.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetListenersArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetListenersResult
    {
        public readonly string AcceleratorId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<Outputs.GetListenersListenerResult> Listeners;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? Status;

        [OutputConstructor]
        private GetListenersResult(
            string acceleratorId,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetListenersListenerResult> listeners,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? status)
        {
            AcceleratorId = acceleratorId;
            Id = id;
            Ids = ids;
            Listeners = listeners;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
        }
    }
}