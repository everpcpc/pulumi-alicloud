// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.MongoDB.Outputs
{

    [OutputType]
    public sealed class InstanceParameter
    {
        /// <summary>
        /// The name of DB instance. It a string of 2 to 256 characters.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of the parameter.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private InstanceParameter(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
