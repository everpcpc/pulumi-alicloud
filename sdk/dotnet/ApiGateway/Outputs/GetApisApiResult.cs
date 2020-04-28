// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ApiGateway.Outputs
{

    [OutputType]
    public sealed class GetApisApiResult
    {
        /// <summary>
        /// API description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// ID of the specified group.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// The group name that the apis belong to.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// API ID, which is generated by the system and globally unique.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// API name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID of the region where the API is located.
        /// </summary>
        public readonly string RegionId;

        [OutputConstructor]
        private GetApisApiResult(
            string description,

            string groupId,

            string groupName,

            string id,

            string name,

            string regionId)
        {
            Description = description;
            GroupId = groupId;
            GroupName = groupName;
            Id = id;
            Name = name;
            RegionId = regionId;
        }
    }
}