// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Eci.Outputs
{

    [OutputType]
    public sealed class ContainerGroupInitContainerVolumeMount
    {
        /// <summary>
        /// The directory of the mounted volume. Data under this directory will be overwritten by the data in the volume.
        /// </summary>
        public readonly string? MountPath;
        /// <summary>
        /// The name of the volume.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Default to `false`.
        /// </summary>
        public readonly bool? ReadOnly;

        [OutputConstructor]
        private ContainerGroupInitContainerVolumeMount(
            string? mountPath,

            string? name,

            bool? readOnly)
        {
            MountPath = mountPath;
            Name = name;
            ReadOnly = readOnly;
        }
    }
}
