// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs.Outputs
{

    [OutputType]
    public sealed class AutoProvisioningGroupLaunchTemplateConfig
    {
        public readonly string? InstanceType;
        public readonly string MaxPrice;
        public readonly string? Priority;
        public readonly string VswitchId;
        public readonly string WeightedCapacity;

        [OutputConstructor]
        private AutoProvisioningGroupLaunchTemplateConfig(
            string? instanceType,

            string maxPrice,

            string? priority,

            string vswitchId,

            string weightedCapacity)
        {
            InstanceType = instanceType;
            MaxPrice = maxPrice;
            Priority = priority;
            VswitchId = vswitchId;
            WeightedCapacity = weightedCapacity;
        }
    }
}
