// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.MaxCompute.Inputs
{

    public sealed class ProjectSecurityPropertiesProjectProtectionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Exclusion policy.
        /// </summary>
        [Input("exceptionPolicy")]
        public Input<string>? ExceptionPolicy { get; set; }

        /// <summary>
        /// Is it turned on.
        /// </summary>
        [Input("protected")]
        public Input<bool>? Protected { get; set; }

        public ProjectSecurityPropertiesProjectProtectionGetArgs()
        {
        }
        public static new ProjectSecurityPropertiesProjectProtectionGetArgs Empty => new ProjectSecurityPropertiesProjectProtectionGetArgs();
    }
}
