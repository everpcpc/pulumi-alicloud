// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.Ga
{
    public static class GetBandwidthPackages
    {
        /// <summary>
        /// This data source provides the Global Accelerator (GA) Bandwidth Packages of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.112.0+.
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
        ///         var example = Output.Create(AliCloud.Ga.GetBandwidthPackages.InvokeAsync(new AliCloud.Ga.GetBandwidthPackagesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_value",
        ///             },
        ///             NameRegex = "the_resource_name",
        ///         }));
        ///         this.FirstGaBandwidthPackageId = example.Apply(example =&gt; example.Packages?[0]?.Id);
        ///     }
        /// 
        ///     [Output("firstGaBandwidthPackageId")]
        ///     public Output&lt;string&gt; FirstGaBandwidthPackageId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBandwidthPackagesResult> InvokeAsync(GetBandwidthPackagesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBandwidthPackagesResult>("alicloud:ga/getBandwidthPackages:getBandwidthPackages", args ?? new GetBandwidthPackagesArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides the Global Accelerator (GA) Bandwidth Packages of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.112.0+.
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
        ///         var example = Output.Create(AliCloud.Ga.GetBandwidthPackages.InvokeAsync(new AliCloud.Ga.GetBandwidthPackagesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_value",
        ///             },
        ///             NameRegex = "the_resource_name",
        ///         }));
        ///         this.FirstGaBandwidthPackageId = example.Apply(example =&gt; example.Packages?[0]?.Id);
        ///     }
        /// 
        ///     [Output("firstGaBandwidthPackageId")]
        ///     public Output&lt;string&gt; FirstGaBandwidthPackageId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBandwidthPackagesResult> Invoke(GetBandwidthPackagesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBandwidthPackagesResult>("alicloud:ga/getBandwidthPackages:getBandwidthPackages", args ?? new GetBandwidthPackagesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetBandwidthPackagesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Bandwidth Package IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Bandwidth Package name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of the bandwidth plan.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The type of the bandwidth packet. China station only supports return to basic.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetBandwidthPackagesArgs()
        {
        }
    }

    public sealed class GetBandwidthPackagesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Bandwidth Package IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Bandwidth Package name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The status of the bandwidth plan.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The type of the bandwidth packet. China station only supports return to basic.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetBandwidthPackagesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBandwidthPackagesResult
    {
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetBandwidthPackagesPackageResult> Packages;
        public readonly string? Status;
        public readonly string? Type;

        [OutputConstructor]
        private GetBandwidthPackagesResult(
            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetBandwidthPackagesPackageResult> packages,

            string? status,

            string? type)
        {
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Packages = packages;
            Status = status;
            Type = type;
        }
    }
}
