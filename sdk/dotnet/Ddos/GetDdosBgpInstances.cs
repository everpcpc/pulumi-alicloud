// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ddos
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides a list of Anti-DDoS Advanced instances in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.57.0+ .
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ddosbgp_instances.html.markdown.
        /// </summary>
        public static Task<GetDdosBgpInstancesResult> GetDdosBgpInstances(GetDdosBgpInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDdosBgpInstancesResult>("alicloud:ddos/getDdosBgpInstances:getDdosBgpInstances", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetDdosBgpInstancesArgs : Pulumi.InvokeArgs
    {
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

        /// <summary>
        /// A regex string to filter results by the instance name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetDdosBgpInstancesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetDdosBgpInstancesResult
    {
        /// <summary>
        /// A list of instance IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of apis. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDdosBgpInstancesInstancesResult> Instances;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of instance names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDdosBgpInstancesResult(
            ImmutableArray<string> ids,
            ImmutableArray<Outputs.GetDdosBgpInstancesInstancesResult> instances,
            string? nameRegex,
            ImmutableArray<string> names,
            string? outputFile,
            string id)
        {
            Ids = ids;
            Instances = instances;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetDdosBgpInstancesInstancesResult
    {
        /// <summary>
        /// The instance's elastic defend bandwidth.
        /// </summary>
        public readonly int Bandwidth;
        /// <summary>
        /// The instance's base defend bandwidth.
        /// </summary>
        public readonly int BaseBandwidth;
        /// <summary>
        /// The instance's id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The instance's count of ip config.
        /// </summary>
        public readonly int IpCount;
        /// <summary>
        /// The instance's IP version.
        /// </summary>
        public readonly string IpType;
        /// <summary>
        /// The instance's remark.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A region of instance.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The instance's type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDdosBgpInstancesInstancesResult(
            int bandwidth,
            int baseBandwidth,
            string id,
            int ipCount,
            string ipType,
            string name,
            string region,
            string type)
        {
            Bandwidth = bandwidth;
            BaseBandwidth = baseBandwidth;
            Id = id;
            IpCount = ipCount;
            IpType = ipType;
            Name = name;
            Region = region;
            Type = type;
        }
    }
    }
}
