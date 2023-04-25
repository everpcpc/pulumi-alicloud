// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    [Obsolete(@"This function has been deprecated in favour of the getEipAddresses function")]
    public static class GetEips
    {
        /// <summary>
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
        ///     var eipsDs = AliCloud.Ecs.GetEips.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstEipId"] = eipsDs.Apply(getEipsResult =&gt; getEipsResult.Eips[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEipsResult> InvokeAsync(GetEipsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEipsResult>("alicloud:ecs/getEips:getEips", args ?? new GetEipsArgs(), options.WithDefaults());

        /// <summary>
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
        ///     var eipsDs = AliCloud.Ecs.GetEips.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstEipId"] = eipsDs.Apply(getEipsResult =&gt; getEipsResult.Eips[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetEipsResult> Invoke(GetEipsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEipsResult>("alicloud:ecs/getEips:getEips", args ?? new GetEipsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEipsArgs : global::Pulumi.InvokeArgs
    {
        [Input("addressName")]
        public string? AddressName { get; set; }

        [Input("associatedInstanceId")]
        public string? AssociatedInstanceId { get; set; }

        [Input("associatedInstanceType")]
        public string? AssociatedInstanceType { get; set; }

        [Input("dryRun")]
        public bool? DryRun { get; set; }

        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of EIP IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("includeReservationData")]
        public bool? IncludeReservationData { get; set; }

        /// <summary>
        /// Public IP Address of the the EIP.
        /// </summary>
        [Input("ipAddress")]
        public string? IpAddress { get; set; }

        [Input("ipAddresses")]
        private List<string>? _ipAddresses;

        /// <summary>
        /// A list of EIP public IP addresses.
        /// </summary>
        [Obsolete(@"Field 'ip_addresses' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'ip_address' instead.")]
        public List<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new List<string>());
            set => _ipAddresses = value;
        }

        [Input("isp")]
        public string? Isp { get; set; }

        [Input("lockReason")]
        public string? LockReason { get; set; }

        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("paymentType")]
        public string? PaymentType { get; set; }

        /// <summary>
        /// The Id of resource group which the eips belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        [Input("segmentInstanceId")]
        public string? SegmentInstanceId { get; set; }

        /// <summary>
        /// EIP status. Possible values are: `Associating`, `Unassociating`, `InUse` and `Available`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetEipsArgs()
        {
        }
        public static new GetEipsArgs Empty => new GetEipsArgs();
    }

    public sealed class GetEipsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("addressName")]
        public Input<string>? AddressName { get; set; }

        [Input("associatedInstanceId")]
        public Input<string>? AssociatedInstanceId { get; set; }

        [Input("associatedInstanceType")]
        public Input<string>? AssociatedInstanceType { get; set; }

        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        [Input("enableDetails")]
        public Input<bool>? EnableDetails { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of EIP IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("includeReservationData")]
        public Input<bool>? IncludeReservationData { get; set; }

        /// <summary>
        /// Public IP Address of the the EIP.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        [Input("ipAddresses")]
        private InputList<string>? _ipAddresses;

        /// <summary>
        /// A list of EIP public IP addresses.
        /// </summary>
        [Obsolete(@"Field 'ip_addresses' has been deprecated from provider version 1.126.0 and it will be removed in the future version. Please use the new attribute 'ip_address' instead.")]
        public InputList<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<string>());
            set => _ipAddresses = value;
        }

        [Input("isp")]
        public Input<string>? Isp { get; set; }

        [Input("lockReason")]
        public Input<string>? LockReason { get; set; }

        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The Id of resource group which the eips belongs.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("segmentInstanceId")]
        public Input<string>? SegmentInstanceId { get; set; }

        /// <summary>
        /// EIP status. Possible values are: `Associating`, `Unassociating`, `InUse` and `Available`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetEipsInvokeArgs()
        {
        }
        public static new GetEipsInvokeArgs Empty => new GetEipsInvokeArgs();
    }


    [OutputType]
    public sealed class GetEipsResult
    {
        public readonly string? AddressName;
        public readonly ImmutableArray<Outputs.GetEipsAddressResult> Addresses;
        public readonly string? AssociatedInstanceId;
        public readonly string? AssociatedInstanceType;
        public readonly bool? DryRun;
        /// <summary>
        /// A list of EIPs. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetEipsEipResult> Eips;
        public readonly bool? EnableDetails;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Optional) A list of EIP IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly bool? IncludeReservationData;
        /// <summary>
        /// Public IP Address of the the EIP.
        /// </summary>
        public readonly string? IpAddress;
        public readonly ImmutableArray<string> IpAddresses;
        public readonly string? Isp;
        public readonly string? LockReason;
        public readonly string? NameRegex;
        /// <summary>
        /// (Optional) A list of EIP names.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? PaymentType;
        /// <summary>
        /// The Id of resource group which the eips belongs.
        /// </summary>
        public readonly string? ResourceGroupId;
        public readonly string? SegmentInstanceId;
        /// <summary>
        /// EIP status. Possible values are: `Associating`, `Unassociating`, `InUse` and `Available`.
        /// </summary>
        public readonly string? Status;
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetEipsResult(
            string? addressName,

            ImmutableArray<Outputs.GetEipsAddressResult> addresses,

            string? associatedInstanceId,

            string? associatedInstanceType,

            bool? dryRun,

            ImmutableArray<Outputs.GetEipsEipResult> eips,

            bool? enableDetails,

            string id,

            ImmutableArray<string> ids,

            bool? includeReservationData,

            string? ipAddress,

            ImmutableArray<string> ipAddresses,

            string? isp,

            string? lockReason,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? paymentType,

            string? resourceGroupId,

            string? segmentInstanceId,

            string? status,

            ImmutableDictionary<string, object>? tags)
        {
            AddressName = addressName;
            Addresses = addresses;
            AssociatedInstanceId = associatedInstanceId;
            AssociatedInstanceType = associatedInstanceType;
            DryRun = dryRun;
            Eips = eips;
            EnableDetails = enableDetails;
            Id = id;
            Ids = ids;
            IncludeReservationData = includeReservationData;
            IpAddress = ipAddress;
            IpAddresses = ipAddresses;
            Isp = isp;
            LockReason = lockReason;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            PaymentType = paymentType;
            ResourceGroupId = resourceGroupId;
            SegmentInstanceId = segmentInstanceId;
            Status = status;
            Tags = tags;
        }
    }
}
