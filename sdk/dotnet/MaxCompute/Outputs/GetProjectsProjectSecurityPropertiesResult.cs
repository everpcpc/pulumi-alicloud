// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.MaxCompute.Outputs
{

    [OutputType]
    public sealed class GetProjectsProjectSecurityPropertiesResult
    {
        /// <summary>
        /// Whether to enable download permission check.
        /// </summary>
        public readonly bool EnableDownloadPrivilege;
        /// <summary>
        /// Label authorization.
        /// </summary>
        public readonly bool LabelSecurity;
        /// <summary>
        /// Project creator permissions.
        /// </summary>
        public readonly bool ObjectCreatorHasAccessPermission;
        /// <summary>
        /// Does the project creator have authorization rights.
        /// </summary>
        public readonly bool ObjectCreatorHasGrantPermission;
        /// <summary>
        /// Project protection.
        /// </summary>
        public readonly Outputs.GetProjectsProjectSecurityPropertiesProjectProtectionResult ProjectProtection;
        /// <summary>
        /// Whether to turn on ACL.
        /// </summary>
        public readonly bool UsingAcl;
        /// <summary>
        /// Whether to enable Policy.
        /// </summary>
        public readonly bool UsingPolicy;

        [OutputConstructor]
        private GetProjectsProjectSecurityPropertiesResult(
            bool enableDownloadPrivilege,

            bool labelSecurity,

            bool objectCreatorHasAccessPermission,

            bool objectCreatorHasGrantPermission,

            Outputs.GetProjectsProjectSecurityPropertiesProjectProtectionResult projectProtection,

            bool usingAcl,

            bool usingPolicy)
        {
            EnableDownloadPrivilege = enableDownloadPrivilege;
            LabelSecurity = labelSecurity;
            ObjectCreatorHasAccessPermission = objectCreatorHasAccessPermission;
            ObjectCreatorHasGrantPermission = objectCreatorHasGrantPermission;
            ProjectProtection = projectProtection;
            UsingAcl = usingAcl;
            UsingPolicy = usingPolicy;
        }
    }
}
