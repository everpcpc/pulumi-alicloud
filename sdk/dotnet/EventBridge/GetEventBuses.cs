// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.EventBridge
{
    public static class GetEventBuses
    {
        /// <summary>
        /// This data source provides the Event Bridge Event Buses of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.129.0+.
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
        ///         var ids = Output.Create(AliCloud.EventBridge.GetEventBuses.InvokeAsync());
        ///         this.EventBridgeEventBusId1 = ids.Apply(ids =&gt; ids.Buses[0].Id);
        ///         var nameRegex = Output.Create(AliCloud.EventBridge.GetEventBuses.InvokeAsync(new AliCloud.EventBridge.GetEventBusesArgs
        ///         {
        ///             NameRegex = "^my-EventBus",
        ///         }));
        ///         this.EventBridgeEventBusId2 = nameRegex.Apply(nameRegex =&gt; nameRegex.Buses[0].Id);
        ///     }
        /// 
        ///     [Output("eventBridgeEventBusId1")]
        ///     public Output&lt;string&gt; EventBridgeEventBusId1 { get; set; }
        ///     [Output("eventBridgeEventBusId2")]
        ///     public Output&lt;string&gt; EventBridgeEventBusId2 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEventBusesResult> InvokeAsync(GetEventBusesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEventBusesResult>("alicloud:eventbridge/getEventBuses:getEventBuses", args ?? new GetEventBusesArgs(), options.WithVersion());
    }


    public sealed class GetEventBusesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The event bus type.
        /// </summary>
        [Input("eventBusType")]
        public string? EventBusType { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Event Bus IDs. Its element value is same as Event Bus Name.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name prefix.
        /// </summary>
        [Input("namePrefix")]
        public string? NamePrefix { get; set; }

        /// <summary>
        /// A regex string to filter results by Event Bus name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetEventBusesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEventBusesResult
    {
        public readonly ImmutableArray<Outputs.GetEventBusesBusResult> Buses;
        public readonly string? EventBusType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NamePrefix;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetEventBusesResult(
            ImmutableArray<Outputs.GetEventBusesBusResult> buses,

            string? eventBusType,

            string id,

            ImmutableArray<string> ids,

            string? namePrefix,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile)
        {
            Buses = buses;
            EventBusType = eventBusType;
            Id = id;
            Ids = ids;
            NamePrefix = namePrefix;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
        }
    }
}
