// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpn
{
    public static partial class Invokes
    {
        /// <summary>
        /// The VPNs data source lists a number of VPNs resource information owned by an Alicloud account.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/vpn_gateways.html.markdown.
        /// </summary>
        public static Task<GetGatewaysResult> GetGateways(GetGatewaysArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGatewaysResult>("alicloud:vpn/getGateways:getGateways", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetGatewaysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Limit search to specific business status - valid value is "Normal", "FinancialLocked".
        /// </summary>
        [Input("businessStatus")]
        public string? BusinessStatus { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// IDs of the VPN.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string of VPN name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// Save the result to the file.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Limit search to specific status - valid value is "Init", "Provisioning", "Active", "Updating", "Deleting".
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// Use the VPC ID as the search key.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetGatewaysArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetGatewaysResult
    {
        /// <summary>
        /// The business status of the VPN gateway.
        /// </summary>
        public readonly string? BusinessStatus;
        /// <summary>
        /// A list of VPN gateways. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGatewaysGatewaysResult> Gateways;
        /// <summary>
        /// IDs of the VPN.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// names of the VPN.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// The status of the VPN
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// ID of the VPC that the VPN belongs.
        /// </summary>
        public readonly string? VpcId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetGatewaysResult(
            string? businessStatus,
            ImmutableArray<Outputs.GetGatewaysGatewaysResult> gateways,
            ImmutableArray<string> ids,
            string? nameRegex,
            ImmutableArray<string> names,
            string? outputFile,
            string? status,
            string? vpcId,
            string id)
        {
            BusinessStatus = businessStatus;
            Gateways = gateways;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
            VpcId = vpcId;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetGatewaysGatewaysResult
    {
        /// <summary>
        /// Limit search to specific business status - valid value is "Normal", "FinancialLocked".
        /// </summary>
        public readonly string BusinessStatus;
        /// <summary>
        /// The creation time of the VPN gateway.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of the VPN
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Whether the ipsec function is enabled.
        /// </summary>
        public readonly string EnableIpsec;
        /// <summary>
        /// Whether the ssl function is enabled.
        /// </summary>
        public readonly string EnableSsl;
        /// <summary>
        /// The expiration time of the VPN gateway.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// ID of the VPN.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The charge type of the VPN gateway.
        /// </summary>
        public readonly string InstanceChargeType;
        /// <summary>
        /// The internet ip of the VPN.
        /// </summary>
        public readonly string InternetIp;
        /// <summary>
        /// The name of the VPN.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Specification of the VPN
        /// </summary>
        public readonly string Specification;
        /// <summary>
        /// Total count of ssl vpn connections.
        /// </summary>
        public readonly int SslConnections;
        /// <summary>
        /// Limit search to specific status - valid value is "Init", "Provisioning", "Active", "Updating", "Deleting".
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Use the VPC ID as the search key.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private GetGatewaysGatewaysResult(
            string businessStatus,
            string createTime,
            string description,
            string enableIpsec,
            string enableSsl,
            string endTime,
            string id,
            string instanceChargeType,
            string internetIp,
            string name,
            string specification,
            int sslConnections,
            string status,
            string vpcId)
        {
            BusinessStatus = businessStatus;
            CreateTime = createTime;
            Description = description;
            EnableIpsec = enableIpsec;
            EnableSsl = enableSsl;
            EndTime = endTime;
            Id = id;
            InstanceChargeType = instanceChargeType;
            InternetIp = internetIp;
            Name = name;
            Specification = specification;
            SslConnections = sslConnections;
            Status = status;
            VpcId = vpcId;
        }
    }
    }
}
