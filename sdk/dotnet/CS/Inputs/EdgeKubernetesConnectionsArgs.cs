// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Inputs
{

    public sealed class EdgeKubernetesConnectionsArgs : Pulumi.ResourceArgs
    {
        [Input("apiServerInternet")]
        public Input<string>? ApiServerInternet { get; set; }

        [Input("apiServerIntranet")]
        public Input<string>? ApiServerIntranet { get; set; }

        [Input("masterPublicIp")]
        public Input<string>? MasterPublicIp { get; set; }

        [Input("serviceDomain")]
        public Input<string>? ServiceDomain { get; set; }

        public EdgeKubernetesConnectionsArgs()
        {
        }
    }
}
