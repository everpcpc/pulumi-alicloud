// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.CloudStorageGateway
{
    public static class GetStorageBundles
    {
        public static Task<GetStorageBundlesResult> InvokeAsync(GetStorageBundlesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStorageBundlesResult>("alicloud:cloudstoragegateway/getStorageBundles:getStorageBundles", args ?? new GetStorageBundlesArgs(), options.WithVersion());

        public static Output<GetStorageBundlesResult> Invoke(GetStorageBundlesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetStorageBundlesResult>("alicloud:cloudstoragegateway/getStorageBundles:getStorageBundles", args ?? new GetStorageBundlesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetStorageBundlesArgs : Pulumi.InvokeArgs
    {
        [Input("backendBucketRegionId", required: true)]
        public string BackendBucketRegionId { get; set; } = null!;

        [Input("ids")]
        private List<string>? _ids;
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetStorageBundlesArgs()
        {
        }
    }

    public sealed class GetStorageBundlesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("backendBucketRegionId", required: true)]
        public Input<string> BackendBucketRegionId { get; set; } = null!;

        [Input("ids")]
        private InputList<string>? _ids;
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetStorageBundlesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStorageBundlesResult
    {
        public readonly string BackendBucketRegionId;
        public readonly ImmutableArray<Outputs.GetStorageBundlesBundleResult> Bundles;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetStorageBundlesResult(
            string backendBucketRegionId,

            ImmutableArray<Outputs.GetStorageBundlesBundleResult> bundles,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile)
        {
            BackendBucketRegionId = backendBucketRegionId;
            Bundles = bundles;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
        }
    }
}
