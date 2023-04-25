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
        /// This data source provides the Global Accelerator (GA) Listeners of the current Alibaba Cloud user.
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
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AliCloud.Ga.GetListeners.Invoke(new()
        ///     {
        ///         AcceleratorId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value",
        ///         },
        ///         NameRegex = "the_resource_name",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstGaListenerId"] = example.Apply(getListenersResult =&gt; getListenersResult.Listeners[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetListenersResult> InvokeAsync(GetListenersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetListenersResult>("alicloud:ga/getListeners:getListeners", args ?? new GetListenersArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Global Accelerator (GA) Listeners of the current Alibaba Cloud user.
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
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AliCloud.Ga.GetListeners.Invoke(new()
        ///     {
        ///         AcceleratorId = "example_value",
        ///         Ids = new[]
        ///         {
        ///             "example_value",
        ///         },
        ///         NameRegex = "the_resource_name",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstGaListenerId"] = example.Apply(getListenersResult =&gt; getListenersResult.Listeners[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetListenersResult> Invoke(GetListenersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetListenersResult>("alicloud:ga/getListeners:getListeners", args ?? new GetListenersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetListenersArgs : global::Pulumi.InvokeArgs
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
        public static new GetListenersArgs Empty => new GetListenersArgs();
    }

    public sealed class GetListenersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The accelerator id.
        /// </summary>
        [Input("acceleratorId", required: true)]
        public Input<string> AcceleratorId { get; set; } = null!;

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Listener IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Listener name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The status of the listener.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetListenersInvokeArgs()
        {
        }
        public static new GetListenersInvokeArgs Empty => new GetListenersInvokeArgs();
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
