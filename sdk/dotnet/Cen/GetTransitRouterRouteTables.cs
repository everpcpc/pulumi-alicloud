// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.Cen
{
    public static class GetTransitRouterRouteTables
    {
        /// <summary>
        /// This data source provides CEN Transit Router Route Tables available to the user.[What is Cen Transit Router Route Tables](https://help.aliyun.com/document_detail/261237.html)
        /// 
        /// &gt; **NOTE:** Available in 1.126.0+
        /// </summary>
        public static Task<GetTransitRouterRouteTablesResult> InvokeAsync(GetTransitRouterRouteTablesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTransitRouterRouteTablesResult>("alicloud:cen/getTransitRouterRouteTables:getTransitRouterRouteTables", args ?? new GetTransitRouterRouteTablesArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides CEN Transit Router Route Tables available to the user.[What is Cen Transit Router Route Tables](https://help.aliyun.com/document_detail/261237.html)
        /// 
        /// &gt; **NOTE:** Available in 1.126.0+
        /// </summary>
        public static Output<GetTransitRouterRouteTablesResult> Invoke(GetTransitRouterRouteTablesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTransitRouterRouteTablesResult>("alicloud:cen/getTransitRouterRouteTables:getTransitRouterRouteTables", args ?? new GetTransitRouterRouteTablesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetTransitRouterRouteTablesArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of CEN Transit Router Route Table IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// ID of the CEN Transit Router Route Table.
        /// </summary>
        [Input("transitRouterId", required: true)]
        public string TransitRouterId { get; set; } = null!;

        [Input("transitRouterRouteTableIds")]
        private List<string>? _transitRouterRouteTableIds;

        /// <summary>
        /// A list of ID of the CEN Transit Router Route Table.
        /// </summary>
        public List<string> TransitRouterRouteTableIds
        {
            get => _transitRouterRouteTableIds ?? (_transitRouterRouteTableIds = new List<string>());
            set => _transitRouterRouteTableIds = value;
        }

        [Input("transitRouterRouteTableNames")]
        private List<string>? _transitRouterRouteTableNames;

        /// <summary>
        /// A list of name of the CEN Transit Router Route Table.
        /// </summary>
        public List<string> TransitRouterRouteTableNames
        {
            get => _transitRouterRouteTableNames ?? (_transitRouterRouteTableNames = new List<string>());
            set => _transitRouterRouteTableNames = value;
        }

        /// <summary>
        /// The status of the transit router route table to query.
        /// </summary>
        [Input("transitRouterRouteTableStatus")]
        public string? TransitRouterRouteTableStatus { get; set; }

        public GetTransitRouterRouteTablesArgs()
        {
        }
    }

    public sealed class GetTransitRouterRouteTablesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of CEN Transit Router Route Table IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// ID of the CEN Transit Router Route Table.
        /// </summary>
        [Input("transitRouterId", required: true)]
        public Input<string> TransitRouterId { get; set; } = null!;

        [Input("transitRouterRouteTableIds")]
        private InputList<string>? _transitRouterRouteTableIds;

        /// <summary>
        /// A list of ID of the CEN Transit Router Route Table.
        /// </summary>
        public InputList<string> TransitRouterRouteTableIds
        {
            get => _transitRouterRouteTableIds ?? (_transitRouterRouteTableIds = new InputList<string>());
            set => _transitRouterRouteTableIds = value;
        }

        [Input("transitRouterRouteTableNames")]
        private InputList<string>? _transitRouterRouteTableNames;

        /// <summary>
        /// A list of name of the CEN Transit Router Route Table.
        /// </summary>
        public InputList<string> TransitRouterRouteTableNames
        {
            get => _transitRouterRouteTableNames ?? (_transitRouterRouteTableNames = new InputList<string>());
            set => _transitRouterRouteTableNames = value;
        }

        /// <summary>
        /// The status of the transit router route table to query.
        /// </summary>
        [Input("transitRouterRouteTableStatus")]
        public Input<string>? TransitRouterRouteTableStatus { get; set; }

        public GetTransitRouterRouteTablesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTransitRouterRouteTablesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of CEN Transit Router Route Table IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of name of CEN Transit Router Route Tables.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? Status;
        /// <summary>
        /// A list of CEN Route Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTransitRouterRouteTablesTableResult> Tables;
        public readonly string TransitRouterId;
        public readonly ImmutableArray<string> TransitRouterRouteTableIds;
        public readonly ImmutableArray<string> TransitRouterRouteTableNames;
        /// <summary>
        /// The status of the route table.
        /// </summary>
        public readonly string? TransitRouterRouteTableStatus;

        [OutputConstructor]
        private GetTransitRouterRouteTablesResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? status,

            ImmutableArray<Outputs.GetTransitRouterRouteTablesTableResult> tables,

            string transitRouterId,

            ImmutableArray<string> transitRouterRouteTableIds,

            ImmutableArray<string> transitRouterRouteTableNames,

            string? transitRouterRouteTableStatus)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
            Tables = tables;
            TransitRouterId = transitRouterId;
            TransitRouterRouteTableIds = transitRouterRouteTableIds;
            TransitRouterRouteTableNames = transitRouterRouteTableNames;
            TransitRouterRouteTableStatus = transitRouterRouteTableStatus;
        }
    }
}
