// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Inputs
{

    public sealed class KubernetesConnectionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// API Server Internet endpoint.
        /// </summary>
        [Input("apiServerInternet")]
        public Input<string>? ApiServerInternet { get; set; }

        /// <summary>
        /// API Server Intranet endpoint.
        /// </summary>
        [Input("apiServerIntranet")]
        public Input<string>? ApiServerIntranet { get; set; }

        /// <summary>
        /// Master node SSH IP address.
        /// </summary>
        [Input("masterPublicIp")]
        public Input<string>? MasterPublicIp { get; set; }

        /// <summary>
        /// Service Access Domain.
        /// </summary>
        [Input("serviceDomain")]
        public Input<string>? ServiceDomain { get; set; }

        public KubernetesConnectionsArgs()
        {
        }
    }
}
