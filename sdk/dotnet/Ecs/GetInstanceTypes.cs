// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides the ECS instance types of Alibaba Cloud.
        /// 
        /// &gt; **NOTE:** By default, only the upgraded instance types are returned. If you want to get outdated instance types, you must set `is_outdated` to true.
        /// 
        /// &gt; **NOTE:** If one instance type is sold out, it will not be exported.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/instance_types.html.markdown.
        /// </summary>
        public static Task<GetInstanceTypesResult> GetInstanceTypes(GetInstanceTypesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceTypesResult>("alicloud:ecs/getInstanceTypes:getInstanceTypes", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetInstanceTypesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The zone where instance types are supported.
        /// </summary>
        [Input("availabilityZone")]
        public string? AvailabilityZone { get; set; }

        /// <summary>
        /// Filter the results to a specific number of cpu cores.
        /// </summary>
        [Input("cpuCoreCount")]
        public int? CpuCoreCount { get; set; }

        /// <summary>
        /// Filter the result whose network interface number is no more than `eni_amount`.
        /// </summary>
        [Input("eniAmount")]
        public int? EniAmount { get; set; }

        /// <summary>
        /// Filter the results by charge type. Valid values: `PrePaid` and `PostPaid`. Default to `PostPaid`.
        /// </summary>
        [Input("instanceChargeType")]
        public string? InstanceChargeType { get; set; }

        /// <summary>
        /// Filter the results based on their family name. For example: 'ecs.n4'.
        /// </summary>
        [Input("instanceTypeFamily")]
        public string? InstanceTypeFamily { get; set; }

        /// <summary>
        /// If true, outdated instance types are included in the results. Default to false.
        /// </summary>
        [Input("isOutdated")]
        public bool? IsOutdated { get; set; }

        [Input("kubernetesNodeRole")]
        public string? KubernetesNodeRole { get; set; }

        /// <summary>
        /// Filter the results to a specific memory size in GB.
        /// </summary>
        [Input("memorySize")]
        public double? MemorySize { get; set; }

        /// <summary>
        /// Filter the results by network type. Valid values: `Classic` and `Vpc`.
        /// </summary>
        [Input("networkType")]
        public string? NetworkType { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("sortedBy")]
        public string? SortedBy { get; set; }

        /// <summary>
        /// Filter the results by ECS spot type. Valid values: `NoSpot`, `SpotWithPriceLimit` and `SpotAsPriceGo`. Default to `NoSpot`.
        /// </summary>
        [Input("spotStrategy")]
        public string? SpotStrategy { get; set; }

        public GetInstanceTypesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetInstanceTypesResult
    {
        public readonly string? AvailabilityZone;
        /// <summary>
        /// Number of CPU cores.
        /// </summary>
        public readonly int? CpuCoreCount;
        /// <summary>
        /// The maximum number of network interfaces that an instance type can be attached to.
        /// </summary>
        public readonly int? EniAmount;
        /// <summary>
        /// A list of instance type IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? InstanceChargeType;
        public readonly string? InstanceTypeFamily;
        /// <summary>
        /// A list of image types. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTypesInstanceTypesResult> InstanceTypes;
        public readonly bool? IsOutdated;
        public readonly string? KubernetesNodeRole;
        /// <summary>
        /// Size of memory, measured in GB.
        /// </summary>
        public readonly double? MemorySize;
        public readonly string? NetworkType;
        public readonly string? OutputFile;
        public readonly string? SortedBy;
        public readonly string? SpotStrategy;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetInstanceTypesResult(
            string? availabilityZone,
            int? cpuCoreCount,
            int? eniAmount,
            ImmutableArray<string> ids,
            string? instanceChargeType,
            string? instanceTypeFamily,
            ImmutableArray<Outputs.GetInstanceTypesInstanceTypesResult> instanceTypes,
            bool? isOutdated,
            string? kubernetesNodeRole,
            double? memorySize,
            string? networkType,
            string? outputFile,
            string? sortedBy,
            string? spotStrategy,
            string id)
        {
            AvailabilityZone = availabilityZone;
            CpuCoreCount = cpuCoreCount;
            EniAmount = eniAmount;
            Ids = ids;
            InstanceChargeType = instanceChargeType;
            InstanceTypeFamily = instanceTypeFamily;
            InstanceTypes = instanceTypes;
            IsOutdated = isOutdated;
            KubernetesNodeRole = kubernetesNodeRole;
            MemorySize = memorySize;
            NetworkType = networkType;
            OutputFile = outputFile;
            SortedBy = sortedBy;
            SpotStrategy = spotStrategy;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetInstanceTypesInstanceTypesBurstableInstanceResult
    {
        /// <summary>
        /// The compute performance benchmark CPU credit of a burstable instance.
        /// </summary>
        public readonly string BaselineCredit;
        /// <summary>
        /// The initial CPU credit of a burstable instance.
        /// </summary>
        public readonly string InitialCredit;

        [OutputConstructor]
        private GetInstanceTypesInstanceTypesBurstableInstanceResult(
            string baselineCredit,
            string initialCredit)
        {
            BaselineCredit = baselineCredit;
            InitialCredit = initialCredit;
        }
    }

    [OutputType]
    public sealed class GetInstanceTypesInstanceTypesGpuResult
    {
        /// <summary>
        /// The number of local storage devices that an instance has been attached to.
        /// </summary>
        public readonly string Amount;
        /// <summary>
        /// The category of local storage that an instance has been attached to.
        /// </summary>
        public readonly string Category;

        [OutputConstructor]
        private GetInstanceTypesInstanceTypesGpuResult(
            string amount,
            string category)
        {
            Amount = amount;
            Category = category;
        }
    }

    [OutputType]
    public sealed class GetInstanceTypesInstanceTypesLocalStorageResult
    {
        /// <summary>
        /// The number of local storage devices that an instance has been attached to.
        /// </summary>
        public readonly string Amount;
        /// <summary>
        /// The capacity of a local storage in GB.
        /// </summary>
        public readonly string Capacity;
        /// <summary>
        /// The category of local storage that an instance has been attached to.
        /// </summary>
        public readonly string Category;

        [OutputConstructor]
        private GetInstanceTypesInstanceTypesLocalStorageResult(
            string amount,
            string capacity,
            string category)
        {
            Amount = amount;
            Capacity = capacity;
            Category = category;
        }
    }

    [OutputType]
    public sealed class GetInstanceTypesInstanceTypesResult
    {
        /// <summary>
        /// List of availability zones that support the instance type.
        /// </summary>
        public readonly ImmutableArray<string> AvailabilityZones;
        /// <summary>
        /// The burstable instance attribution:
        /// </summary>
        public readonly GetInstanceTypesInstanceTypesBurstableInstanceResult BurstableInstance;
        /// <summary>
        /// Filter the results to a specific number of cpu cores.
        /// </summary>
        public readonly int CpuCoreCount;
        /// <summary>
        /// Filter the result whose network interface number is no more than `eni_amount`.
        /// </summary>
        public readonly int EniAmount;
        /// <summary>
        /// The instance type family.
        /// </summary>
        public readonly string Family;
        /// <summary>
        /// The GPU attribution of an instance type:
        /// </summary>
        public readonly GetInstanceTypesInstanceTypesGpuResult Gpu;
        /// <summary>
        /// ID of the instance type.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Local storage of an instance type:
        /// </summary>
        public readonly GetInstanceTypesInstanceTypesLocalStorageResult LocalStorage;
        /// <summary>
        /// Filter the results to a specific memory size in GB.
        /// </summary>
        public readonly double MemorySize;
        public readonly string Price;

        [OutputConstructor]
        private GetInstanceTypesInstanceTypesResult(
            ImmutableArray<string> availabilityZones,
            GetInstanceTypesInstanceTypesBurstableInstanceResult burstableInstance,
            int cpuCoreCount,
            int eniAmount,
            string family,
            GetInstanceTypesInstanceTypesGpuResult gpu,
            string id,
            GetInstanceTypesInstanceTypesLocalStorageResult localStorage,
            double memorySize,
            string price)
        {
            AvailabilityZones = availabilityZones;
            BurstableInstance = burstableInstance;
            CpuCoreCount = cpuCoreCount;
            EniAmount = eniAmount;
            Family = family;
            Gpu = gpu;
            Id = id;
            LocalStorage = localStorage;
            MemorySize = memorySize;
            Price = price;
        }
    }
    }
}
