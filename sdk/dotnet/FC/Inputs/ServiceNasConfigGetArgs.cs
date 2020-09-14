// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.FC.Inputs
{

    public sealed class ServiceNasConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The group id of your NAS file system.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<int> GroupId { get; set; } = null!;

        [Input("mountPoints", required: true)]
        private InputList<Inputs.ServiceNasConfigMountPointGetArgs>? _mountPoints;

        /// <summary>
        /// Config the NAS mount points, including following attributes:
        /// </summary>
        public InputList<Inputs.ServiceNasConfigMountPointGetArgs> MountPoints
        {
            get => _mountPoints ?? (_mountPoints = new InputList<Inputs.ServiceNasConfigMountPointGetArgs>());
            set => _mountPoints = value;
        }

        /// <summary>
        /// The user id of your NAS file system.
        /// </summary>
        [Input("userId", required: true)]
        public Input<int> UserId { get; set; } = null!;

        public ServiceNasConfigGetArgs()
        {
        }
    }
}
