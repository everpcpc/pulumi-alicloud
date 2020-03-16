// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides the ECS instance type families of Alibaba Cloud.
        /// 
        /// &gt; **NOTE:** Available in 1.54.0+
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/instance_type_families.html.markdown.
        /// </summary>
        public static Task<GetInstanceTypeFamiliesResult> GetInstanceTypeFamilies(GetInstanceTypeFamiliesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceTypeFamiliesResult>("alicloud:ecs/getInstanceTypeFamilies:getInstanceTypeFamilies", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetInstanceTypeFamiliesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The generation of the instance type family, Valid values: `ecs-1`, `ecs-2`, `ecs-3` and `ecs-4`. For more information, see [Instance type families](https://www.alibabacloud.com/help/doc-detail/25378.htm). 
        /// </summary>
        [Input("generation")]
        public string? Generation { get; set; }

        /// <summary>
        /// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
        /// </summary>
        [Input("instanceChargeType")]
        public string? InstanceChargeType { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Filter the results by ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
        /// </summary>
        [Input("spotStrategy")]
        public string? SpotStrategy { get; set; }

        /// <summary>
        /// The Zone to launch the instance.
        /// </summary>
        [Input("zoneId")]
        public string? ZoneId { get; set; }

        public GetInstanceTypeFamiliesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetInstanceTypeFamiliesResult
    {
        public readonly ImmutableArray<Outputs.GetInstanceTypeFamiliesFamiliesResult> Families;
        /// <summary>
        /// The generation of the instance type family.
        /// </summary>
        public readonly string? Generation;
        /// <summary>
        /// A list of instance type family IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? InstanceChargeType;
        public readonly string? OutputFile;
        public readonly string? SpotStrategy;
        public readonly string? ZoneId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetInstanceTypeFamiliesResult(
            ImmutableArray<Outputs.GetInstanceTypeFamiliesFamiliesResult> families,
            string? generation,
            ImmutableArray<string> ids,
            string? instanceChargeType,
            string? outputFile,
            string? spotStrategy,
            string? zoneId,
            string id)
        {
            Families = families;
            Generation = generation;
            Ids = ids;
            InstanceChargeType = instanceChargeType;
            OutputFile = outputFile;
            SpotStrategy = spotStrategy;
            ZoneId = zoneId;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetInstanceTypeFamiliesFamiliesResult
    {
        /// <summary>
        /// The generation of the instance type family, Valid values: `ecs-1`, `ecs-2`, `ecs-3` and `ecs-4`. For more information, see [Instance type families](https://www.alibabacloud.com/help/doc-detail/25378.htm). 
        /// </summary>
        public readonly string Generation;
        /// <summary>
        /// ID of the instance type family.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Zone to launch the instance.
        /// </summary>
        public readonly ImmutableArray<string> ZoneIds;

        [OutputConstructor]
        private GetInstanceTypeFamiliesFamiliesResult(
            string generation,
            string id,
            ImmutableArray<string> zoneIds)
        {
            Generation = generation;
            Id = id;
            ZoneIds = zoneIds;
        }
    }
    }
}
