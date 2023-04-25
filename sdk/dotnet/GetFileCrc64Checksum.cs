// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud
{
    public static class GetFileCrc64Checksum
    {
        /// <summary>
        /// This data source compute file crc64 checksum.
        /// 
        /// &gt; **NOTE:** Available in 1.59.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = AliCloud.GetFileCrc64Checksum.Invoke(new()
        ///     {
        ///         Filename = "exampleFileName",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["fileCrc64Checksum"] = data.Alicloud_file_crc64_checksum.Defualt.Checksum,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFileCrc64ChecksumResult> InvokeAsync(GetFileCrc64ChecksumArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFileCrc64ChecksumResult>("alicloud:index/getFileCrc64Checksum:getFileCrc64Checksum", args ?? new GetFileCrc64ChecksumArgs(), options.WithDefaults());

        /// <summary>
        /// This data source compute file crc64 checksum.
        /// 
        /// &gt; **NOTE:** Available in 1.59.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = AliCloud.GetFileCrc64Checksum.Invoke(new()
        ///     {
        ///         Filename = "exampleFileName",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["fileCrc64Checksum"] = data.Alicloud_file_crc64_checksum.Defualt.Checksum,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFileCrc64ChecksumResult> Invoke(GetFileCrc64ChecksumInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFileCrc64ChecksumResult>("alicloud:index/getFileCrc64Checksum:getFileCrc64Checksum", args ?? new GetFileCrc64ChecksumInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFileCrc64ChecksumArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the file to be computed crc64 checksum.
        /// </summary>
        [Input("filename", required: true)]
        public string Filename { get; set; } = null!;

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetFileCrc64ChecksumArgs()
        {
        }
        public static new GetFileCrc64ChecksumArgs Empty => new GetFileCrc64ChecksumArgs();
    }

    public sealed class GetFileCrc64ChecksumInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the file to be computed crc64 checksum.
        /// </summary>
        [Input("filename", required: true)]
        public Input<string> Filename { get; set; } = null!;

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetFileCrc64ChecksumInvokeArgs()
        {
        }
        public static new GetFileCrc64ChecksumInvokeArgs Empty => new GetFileCrc64ChecksumInvokeArgs();
    }


    [OutputType]
    public sealed class GetFileCrc64ChecksumResult
    {
        /// <summary>
        /// the file checksum of crc64.
        /// </summary>
        public readonly string Checksum;
        public readonly string Filename;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetFileCrc64ChecksumResult(
            string checksum,

            string filename,

            string id,

            string? outputFile)
        {
            Checksum = checksum;
            Filename = filename;
            Id = id;
            OutputFile = outputFile;
        }
    }
}
