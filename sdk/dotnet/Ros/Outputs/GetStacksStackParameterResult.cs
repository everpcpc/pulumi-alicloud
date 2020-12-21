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
    public sealed class GetStacksStackParameterResult
    {
        /// <summary>
        /// The key of parameters.
        /// </summary>
        public readonly string ParameterKey;
        /// <summary>
        /// The value of parameters.
        /// </summary>
        public readonly string ParameterValue;

        [OutputConstructor]
        private GetStacksStackParameterResult(
            string parameterKey,

            string parameterValue)
        {
            ParameterKey = parameterKey;
            ParameterValue = parameterValue;
        }
    }
}
