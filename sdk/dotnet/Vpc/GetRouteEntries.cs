// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides a list of Route Entries owned by an Alibaba Cloud account.
        /// 
        /// &gt; **NOTE:** Available in 1.37.0+.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/route_entries.html.markdown.
        /// </summary>
        public static Task<GetRouteEntriesResult> GetRouteEntries(GetRouteEntriesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRouteEntriesResult>("alicloud:vpc/getRouteEntries:getRouteEntries", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetRouteEntriesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The destination CIDR block of the route entry.
        /// </summary>
        [Input("cidrBlock")]
        public string? CidrBlock { get; set; }

        /// <summary>
        /// The instance ID of the next hop.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ID of the router table to which the route entry belongs.
        /// </summary>
        [Input("routeTableId", required: true)]
        public string RouteTableId { get; set; } = null!;

        /// <summary>
        /// The type of the route entry.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetRouteEntriesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetRouteEntriesResult
    {
        /// <summary>
        /// The destination CIDR block of the route entry.
        /// </summary>
        public readonly string? CidrBlock;
        /// <summary>
        /// A list of Route Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouteEntriesEntriesResult> Entries;
        /// <summary>
        /// The instance ID of the next hop.
        /// </summary>
        public readonly string? InstanceId;
        public readonly string? OutputFile;
        /// <summary>
        /// The ID of the router table to which the route entry belongs.
        /// </summary>
        public readonly string RouteTableId;
        /// <summary>
        /// The type of the route entry.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetRouteEntriesResult(
            string? cidrBlock,
            ImmutableArray<Outputs.GetRouteEntriesEntriesResult> entries,
            string? instanceId,
            string? outputFile,
            string routeTableId,
            string? type,
            string id)
        {
            CidrBlock = cidrBlock;
            Entries = entries;
            InstanceId = instanceId;
            OutputFile = outputFile;
            RouteTableId = routeTableId;
            Type = type;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetRouteEntriesEntriesResult
    {
        /// <summary>
        /// The destination CIDR block of the route entry.
        /// </summary>
        public readonly string CidrBlock;
        /// <summary>
        /// The instance ID of the next hop.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The type of the next hop.
        /// </summary>
        public readonly string NextHopType;
        /// <summary>
        /// The ID of the router table to which the route entry belongs.
        /// </summary>
        public readonly string RouteTableId;
        /// <summary>
        /// The status of the route entry.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The type of the route entry.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetRouteEntriesEntriesResult(
            string cidrBlock,
            string instanceId,
            string nextHopType,
            string routeTableId,
            string status,
            string type)
        {
            CidrBlock = cidrBlock;
            InstanceId = instanceId;
            NextHopType = nextHopType;
            RouteTableId = routeTableId;
            Status = status;
            Type = type;
        }
    }
    }
}
