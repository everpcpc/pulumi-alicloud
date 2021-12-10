// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ros.Outputs
{

    [OutputType]
    public sealed class StackInstanceParameterOverride
    {
        /// <summary>
        /// The key of override parameter. If you do not specify the key and value of the parameter, ROS uses the key and value that you specified when you created the stack group.
        /// </summary>
        public readonly string? ParameterKey;
        /// <summary>
        /// The value of override parameter. If you do not specify the key and value of the parameter, ROS uses the key and value that you specified when you created the stack group.
        /// </summary>
        public readonly string? ParameterValue;

        [OutputConstructor]
        private StackInstanceParameterOverride(
            string? parameterKey,

            string? parameterValue)
        {
            ParameterKey = parameterKey;
            ParameterValue = parameterValue;
        }
    }
}
