// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ros.Inputs
{

    public sealed class StackGroupParameterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The parameter key.
        /// </summary>
        [Input("parameterKey")]
        public Input<string>? ParameterKey { get; set; }

        /// <summary>
        /// The parameter value.
        /// </summary>
        [Input("parameterValue")]
        public Input<string>? ParameterValue { get; set; }

        public StackGroupParameterArgs()
        {
        }
    }
}
