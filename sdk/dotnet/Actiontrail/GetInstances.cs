// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ActionTrail
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides a list of ALIKAFKA Instances in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.59.0+
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/alikafka_instances.html.markdown.
        /// </summary>
        public static Task<GetInstancesResult> GetInstances(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("alicloud:actiontrail/getInstances:getInstances", args ?? InvokeArgs.Empty, options.WithVersion());
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
        /// The create time of the instance.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The deploy type of the instance. 0: sharing instance, 1: vpc instance, 2: vpc instance(support ip mapping), 3: eip instance, 4: eip/vpc instance, 5: vpc instance.
        /// </summary>
        public readonly int DeployType;
        /// <summary>
        /// The disk size of the instance.
        /// </summary>
        public readonly int DiskSize;
        /// <summary>
        /// The disk type of the instance. 0: efficient cloud disk , 1: SSD.
        /// </summary>
        public readonly int DiskType;
        /// <summary>
        /// The peak bandwidth of the instance.
        /// </summary>
        public readonly int EipMax;
        /// <summary>
        /// ID of the instance.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The peak value of io of the instance.
        /// </summary>
        public readonly int IoMax;
        /// <summary>
        /// Name of the instance.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The paid type of the instance.
        /// </summary>
        public readonly string PaidType;
        /// <summary>
        /// The current status of the instance. -1: unknown status, 0: wait deploy, 1: initializing, 2: preparing, 3 starting, 5: in service, 7: wait upgrade, 8: upgrading, 10: released, 15: freeze, 101: deploy error, 102: upgrade error. 
        /// </summary>
        public readonly int ServiceStatus;
        /// <summary>
        /// The spec type of the instance.
        /// </summary>
        public readonly string SpecType;
        /// <summary>
        /// The max num of topic can be create of the instance.
        /// </summary>
        public readonly int TopicQuota;
        /// <summary>
        /// The ID of attaching VPC to instance.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// The ID of attaching vswitch to instance.
        /// </summary>
        public readonly string VswitchId;
        /// <summary>
        /// The ID of attaching zone to instance.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private GetInstancesInstancesResult(
            string createTime,
            int deployType,
            int diskSize,
            int diskType,
            int eipMax,
            string id,
            int ioMax,
            string name,
            string paidType,
            int serviceStatus,
            string specType,
            int topicQuota,
            string vpcId,
            string vswitchId,
            string zoneId)
        {
            CreateTime = createTime;
            DeployType = deployType;
            DiskSize = diskSize;
            DiskType = diskType;
            EipMax = eipMax;
            Id = id;
            IoMax = ioMax;
            Name = name;
            PaidType = paidType;
            ServiceStatus = serviceStatus;
            SpecType = specType;
            TopicQuota = topicQuota;
            VpcId = vpcId;
            VswitchId = vswitchId;
            ZoneId = zoneId;
        }
    }
    }
}
