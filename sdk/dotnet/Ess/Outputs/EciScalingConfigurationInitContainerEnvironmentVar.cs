// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ess.Outputs
{

    [OutputType]
    public sealed class EciScalingConfigurationInitContainerEnvironmentVar
    {
        /// <summary>
        /// The name of the variable. The name can be 1 to 128 characters in length and can contain letters,
        /// digits, and underscores (_). It cannot start with a digit.
        /// digits, and underscores (_). It cannot start with a digit.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// The value of the variable. The value can be 0 to 256 characters in length.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private EciScalingConfigurationInitContainerEnvironmentVar(
            string? key,

            string? value)
        {
            Key = key;
            Value = value;
        }
    }
}
