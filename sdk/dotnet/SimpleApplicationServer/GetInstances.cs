// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.SimpleApplicationServer
{
    public static class GetInstances
    {
        /// <summary>
        /// This data source provides the Simple Application Server Instances of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.135.0+.
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
        ///         var ids = Output.Create(AliCloud.SimpleApplicationServer.GetInstances.InvokeAsync(new AliCloud.SimpleApplicationServer.GetInstancesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_id",
        ///             },
        ///         }));
        ///         this.SimpleApplicationServerInstanceId1 = ids.Apply(ids =&gt; ids.Instances?[0]?.Id);
        ///         var nameRegex = Output.Create(AliCloud.SimpleApplicationServer.GetInstances.InvokeAsync(new AliCloud.SimpleApplicationServer.GetInstancesArgs
        ///         {
        ///             NameRegex = "^my-Instance",
        ///         }));
        ///         this.SimpleApplicationServerInstanceId2 = nameRegex.Apply(nameRegex =&gt; nameRegex.Instances?[0]?.Id);
        ///     }
        /// 
        ///     [Output("simpleApplicationServerInstanceId1")]
        ///     public Output&lt;string&gt; SimpleApplicationServerInstanceId1 { get; set; }
        ///     [Output("simpleApplicationServerInstanceId2")]
        ///     public Output&lt;string&gt; SimpleApplicationServerInstanceId2 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstancesResult> InvokeAsync(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("alicloud:simpleapplicationserver/getInstances:getInstances", args ?? new GetInstancesArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides the Simple Application Server Instances of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.135.0+.
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
        ///         var ids = Output.Create(AliCloud.SimpleApplicationServer.GetInstances.InvokeAsync(new AliCloud.SimpleApplicationServer.GetInstancesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_id",
        ///             },
        ///         }));
        ///         this.SimpleApplicationServerInstanceId1 = ids.Apply(ids =&gt; ids.Instances?[0]?.Id);
        ///         var nameRegex = Output.Create(AliCloud.SimpleApplicationServer.GetInstances.InvokeAsync(new AliCloud.SimpleApplicationServer.GetInstancesArgs
        ///         {
        ///             NameRegex = "^my-Instance",
        ///         }));
        ///         this.SimpleApplicationServerInstanceId2 = nameRegex.Apply(nameRegex =&gt; nameRegex.Instances?[0]?.Id);
        ///     }
        /// 
        ///     [Output("simpleApplicationServerInstanceId1")]
        ///     public Output&lt;string&gt; SimpleApplicationServerInstanceId1 { get; set; }
        ///     [Output("simpleApplicationServerInstanceId2")]
        ///     public Output&lt;string&gt; SimpleApplicationServerInstanceId2 { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstancesResult> Invoke(GetInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstancesResult>("alicloud:simpleapplicationserver/getInstances:getInstances", args ?? new GetInstancesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetInstancesArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Instance IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Instance name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The billing method of the simple application server.
        /// </summary>
        [Input("paymentType")]
        public string? PaymentType { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetInstancesArgs()
        {
        }
    }

    public sealed class GetInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Instance IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Instance name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The billing method of the simple application server.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstancesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<Outputs.GetInstancesInstanceResult> Instances;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? PaymentType;
        public readonly string? Status;

        [OutputConstructor]
        private GetInstancesResult(
            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetInstancesInstanceResult> instances,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? paymentType,

            string? status)
        {
            Id = id;
            Ids = ids;
            Instances = instances;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            PaymentType = paymentType;
            Status = status;
        }
    }
}
