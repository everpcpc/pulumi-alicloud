// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CR.Outputs
{

    [OutputType]
    public sealed class ChainChainConfig
    {
        /// <summary>
        /// Each node in the delivery chain.
        /// </summary>
        public readonly ImmutableArray<Outputs.ChainChainConfigNode> Nodes;
        /// <summary>
        /// Execution sequence relationship between delivery chain nodes.
        /// </summary>
        public readonly ImmutableArray<Outputs.ChainChainConfigRouter> Routers;

        [OutputConstructor]
        private ChainChainConfig(
            ImmutableArray<Outputs.ChainChainConfigNode> nodes,

            ImmutableArray<Outputs.ChainChainConfigRouter> routers)
        {
            Nodes = nodes;
            Routers = routers;
        }
    }
}