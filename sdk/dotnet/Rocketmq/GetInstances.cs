// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.RocketMQ
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides a list of ONS Instances in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.52.0+
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/ons_instances.html.markdown.
        /// </summary>
        public static Task<GetInstancesResult> GetInstances(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("alicloud:rocketmq/getInstances:getInstances", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetInstancesArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of instance IDs to filter results.
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

        public GetInstancesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetInstancesResult
    {
        /// <summary>
        /// A list of instance IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A list of instances. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstancesInstancesResult> Instances;
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
        private GetInstancesResult(
            ImmutableArray<string> ids,
            ImmutableArray<Outputs.GetInstancesInstancesResult> instances,
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
    public sealed class GetInstancesInstancesResult
    {
        /// <summary>
        /// ID of the instance.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// ID of the instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// Name of the instance.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// The status of the instance. Read [Fields in InstanceVO](https://www.alibabacloud.com/help/doc-detail/106351.html) for further details.
        /// </summary>
        public readonly int InstanceStatus;
        /// <summary>
        /// The type of the instance. Read [Fields in InstanceVO](https://www.alibabacloud.com/help/doc-detail/106351.html) for further details.
        /// </summary>
        public readonly int InstanceType;
        /// <summary>
        /// The automatic release time of an Enterprise Platinum Edition instance.
        /// </summary>
        public readonly int ReleaseTime;

        [OutputConstructor]
        private GetInstancesInstancesResult(
            string id,
            string instanceId,
            string instanceName,
            int instanceStatus,
            int instanceType,
            int releaseTime)
        {
            Id = id;
            InstanceId = instanceId;
            InstanceName = instanceName;
            InstanceStatus = instanceStatus;
            InstanceType = instanceType;
            ReleaseTime = releaseTime;
        }
    }
    }
}
