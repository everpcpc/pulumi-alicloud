// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Outputs
{

    [OutputType]
    public sealed class GetEdgeKubernetesClustersClusterConnectionsResult
    {
        /// <summary>
        /// API Server Internet endpoint.
        /// </summary>
        public readonly string ApiServerInternet;
        /// <summary>
        /// API Server Intranet endpoint.
        /// </summary>
        public readonly string ApiServerIntranet;

        [OutputConstructor]
        private GetEdgeKubernetesClustersClusterConnectionsResult(
            string apiServerInternet,

            string apiServerIntranet)
        {
            ApiServerInternet = apiServerInternet;
            ApiServerIntranet = apiServerIntranet;
        }
    }
}
