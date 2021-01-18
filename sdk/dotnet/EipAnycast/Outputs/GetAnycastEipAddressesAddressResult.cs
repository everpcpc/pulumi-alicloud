// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.EipAnycast.Outputs
{

    [OutputType]
    public sealed class GetAnycastEipAddressesAddressResult
    {
        /// <summary>
        /// Anycast EIP instance account ID.
        /// </summary>
        public readonly int AliUid;
        /// <summary>
        /// Anycast EIP instance name.
        /// </summary>
        public readonly string AnycastEipAddressName;
        /// <summary>
        /// AnycastEip binding information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAnycastEipAddressesAddressAnycastEipBindInfoListResult> AnycastEipBindInfoLists;
        /// <summary>
        /// Anycast EIP instance ID.
        /// </summary>
        public readonly string AnycastId;
        /// <summary>
        /// The peak bandwidth of the Anycast EIP instance, in Mbps.
        /// </summary>
        public readonly int Bandwidth;
        /// <summary>
        /// Anycast EIP instance account BID.
        /// </summary>
        public readonly string Bid;
        /// <summary>
        /// The business status of the Anycast EIP instance. -`Normal`: Normal state. -`FinancialLocked`: The status of arrears locked.
        /// </summary>
        public readonly string BusinessStatus;
        /// <summary>
        /// Anycast EIP instance description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The ID of the Anycast Eip Address.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The billing method of Anycast EIP instance. `PayByBandwidth`: refers to the method of billing based on traffic.
        /// </summary>
        public readonly string InternetChargeType;
        /// <summary>
        /// Anycast EIP instance IP address.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// The payment model of Anycast EIP instance. "PostPaid": Refers to the post-paid mode.
        /// </summary>
        public readonly string PaymentType;
        /// <summary>
        /// Anycast EIP instance access area. "international": Refers to areas outside of Mainland China.
        /// </summary>
        public readonly string ServiceLocation;
        /// <summary>
        /// IP status。- `Associating`, `Unassociating`, `Allocated`, `Associated`, `Modifying`, `Releasing`, `Released`.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetAnycastEipAddressesAddressResult(
            int aliUid,

            string anycastEipAddressName,

            ImmutableArray<Outputs.GetAnycastEipAddressesAddressAnycastEipBindInfoListResult> anycastEipBindInfoLists,

            string anycastId,

            int bandwidth,

            string bid,

            string businessStatus,

            string description,

            string id,

            string internetChargeType,

            string ipAddress,

            string paymentType,

            string serviceLocation,

            string status)
        {
            AliUid = aliUid;
            AnycastEipAddressName = anycastEipAddressName;
            AnycastEipBindInfoLists = anycastEipBindInfoLists;
            AnycastId = anycastId;
            Bandwidth = bandwidth;
            Bid = bid;
            BusinessStatus = businessStatus;
            Description = description;
            Id = id;
            InternetChargeType = internetChargeType;
            IpAddress = ipAddress;
            PaymentType = paymentType;
            ServiceLocation = serviceLocation;
            Status = status;
        }
    }
}
