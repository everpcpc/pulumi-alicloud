// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oss.Outputs
{

    [OutputType]
    public sealed class GetInstancesInstanceResult
    {
        /// <summary>
        /// The cluster type of the instance. Possible values: `SSD`, `HYBRID`.
        /// </summary>
        public readonly string ClusterType;
        /// <summary>
        /// The create time of the instance.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of the instance.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The instance quota which indicating the maximum number of tables.
        /// </summary>
        public readonly int EntityQuota;
        /// <summary>
        /// ID of the instance.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Instance name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The network type of the instance. Possible values: `NORMAL`, `VPC`, `VPC_CONSOLE`.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// The maximum adjustable read capacity unit of the instance.
        /// </summary>
        public readonly int ReadCapacity;
        /// <summary>
        /// Instance status. Possible values: `Running`, `Disabled`, `Deleting`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A map of tags assigned to the instance. It must be in the format:
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var instancesDs = AliCloud.Ots.GetInstances.Invoke(new()
        ///     {
        ///         Tags = 
        ///         {
        ///             { "tagKey1", "tagValue1" },
        ///             { "tagKey2", "tagValue2" },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// The user id of the instance.
        /// </summary>
        public readonly string UserId;
        /// <summary>
        /// The maximum adjustable write capacity unit of the instance.
        /// </summary>
        public readonly int WriteCapacity;

        [OutputConstructor]
        private GetInstancesInstanceResult(
            string clusterType,

            string createTime,

            string description,

            int entityQuota,

            string id,

            string name,

            string network,

            int readCapacity,

            string status,

            ImmutableDictionary<string, object> tags,

            string userId,

            int writeCapacity)
        {
            ClusterType = clusterType;
            CreateTime = createTime;
            Description = description;
            EntityQuota = entityQuota;
            Id = id;
            Name = name;
            Network = network;
            ReadCapacity = readCapacity;
            Status = status;
            Tags = tags;
            UserId = userId;
            WriteCapacity = writeCapacity;
        }
    }
}
