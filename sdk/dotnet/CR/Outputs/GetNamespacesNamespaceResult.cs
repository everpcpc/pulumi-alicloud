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
    public sealed class GetNamespacesNamespaceResult
    {
        /// <summary>
        /// Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
        /// </summary>
        public readonly bool AutoCreate;
        /// <summary>
        /// `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
        /// </summary>
        public readonly string DefaultVisibility;
        /// <summary>
        /// Name of Container Registry namespace.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetNamespacesNamespaceResult(
            bool autoCreate,

            string defaultVisibility,

            string name)
        {
            AutoCreate = autoCreate;
            DefaultVisibility = defaultVisibility;
            Name = name;
        }
    }
}