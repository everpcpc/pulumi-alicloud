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
    public sealed class GetNetworkInterfacesInterfaceAssociatedPublicIpResult
    {
        public readonly string PublicIpAddress;

        [OutputConstructor]
        private GetNetworkInterfacesInterfaceAssociatedPublicIpResult(string publicIpAddress)
        {
            PublicIpAddress = publicIpAddress;
        }
    }
}
