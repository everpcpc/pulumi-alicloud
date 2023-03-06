// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ServiceCatalog.Outputs
{

    [OutputType]
    public sealed class GetLaunchOptionsLaunchOptionConstraintSummaryResult
    {
        /// <summary>
        /// Constraint type.The value is Launch, which indicates that the constraint is started.
        /// </summary>
        public readonly string ConstraintType;
        /// <summary>
        /// Constraint description.
        /// </summary>
        public readonly string Description;

        [OutputConstructor]
        private GetLaunchOptionsLaunchOptionConstraintSummaryResult(
            string constraintType,

            string description)
        {
            ConstraintType = constraintType;
            Description = description;
        }
    }
}