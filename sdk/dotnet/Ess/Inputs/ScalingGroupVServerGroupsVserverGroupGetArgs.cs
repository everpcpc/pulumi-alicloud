// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ess.Inputs
{

    public sealed class ScalingGroupVServerGroupsVserverGroupGetArgs : Pulumi.ResourceArgs
    {
        [Input("loadbalancerId", required: true)]
        public Input<string> LoadbalancerId { get; set; } = null!;

        [Input("vserverAttributes", required: true)]
        private InputList<Inputs.ScalingGroupVServerGroupsVserverGroupVserverAttributeGetArgs>? _vserverAttributes;
        public InputList<Inputs.ScalingGroupVServerGroupsVserverGroupVserverAttributeGetArgs> VserverAttributes
        {
            get => _vserverAttributes ?? (_vserverAttributes = new InputList<Inputs.ScalingGroupVServerGroupsVserverGroupVserverAttributeGetArgs>());
            set => _vserverAttributes = value;
        }

        public ScalingGroupVServerGroupsVserverGroupGetArgs()
        {
        }
    }
}