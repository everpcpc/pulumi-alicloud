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
    public sealed class ContainerGroupContainer
    {
        /// <summary>
        /// The arguments passed to the commands.
        /// </summary>
        public readonly ImmutableArray<string> Args;
        /// <summary>
        /// The commands run by the init container.
        /// </summary>
        public readonly ImmutableArray<string> Commands;
        /// <summary>
        /// The amount of CPU resources allocated to the container group.
        /// </summary>
        public readonly double? Cpu;
        /// <summary>
        /// The structure of environmentVars.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerGroupContainerEnvironmentVar> EnvironmentVars;
        /// <summary>
        /// The number GPUs.
        /// </summary>
        public readonly int? Gpu;
        /// <summary>
        /// The image of the container.
        /// </summary>
        public readonly string Image;
        /// <summary>
        /// The restart policy of the image.
        /// </summary>
        public readonly string? ImagePullPolicy;
        /// <summary>
        /// The health check of the container.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerGroupContainerLivenessProbe> LivenessProbes;
        /// <summary>
        /// The amount of memory resources allocated to the container group.
        /// </summary>
        public readonly double? Memory;
        /// <summary>
        /// The name of the volume.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The structure of port.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerGroupContainerPort> Ports;
        /// <summary>
        /// The health check of the container.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerGroupContainerReadinessProbe> ReadinessProbes;
        public readonly bool? Ready;
        public readonly int? RestartCount;
        /// <summary>
        /// The structure of volumeMounts.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerGroupContainerVolumeMount> VolumeMounts;
        /// <summary>
        /// The working directory of the container.
        /// </summary>
        public readonly string? WorkingDir;

        [OutputConstructor]
        private ContainerGroupContainer(
            ImmutableArray<string> args,

            ImmutableArray<string> commands,

            double? cpu,

            ImmutableArray<Outputs.ContainerGroupContainerEnvironmentVar> environmentVars,

            int? gpu,

            string image,

            string? imagePullPolicy,

            ImmutableArray<Outputs.ContainerGroupContainerLivenessProbe> livenessProbes,

            double? memory,

            string name,

            ImmutableArray<Outputs.ContainerGroupContainerPort> ports,

            ImmutableArray<Outputs.ContainerGroupContainerReadinessProbe> readinessProbes,

            bool? ready,

            int? restartCount,

            ImmutableArray<Outputs.ContainerGroupContainerVolumeMount> volumeMounts,

            string? workingDir)
        {
            Args = args;
            Commands = commands;
            Cpu = cpu;
            EnvironmentVars = environmentVars;
            Gpu = gpu;
            Image = image;
            ImagePullPolicy = imagePullPolicy;
            LivenessProbes = livenessProbes;
            Memory = memory;
            Name = name;
            Ports = ports;
            ReadinessProbes = readinessProbes;
            Ready = ready;
            RestartCount = restartCount;
            VolumeMounts = volumeMounts;
            WorkingDir = workingDir;
        }
    }
}
