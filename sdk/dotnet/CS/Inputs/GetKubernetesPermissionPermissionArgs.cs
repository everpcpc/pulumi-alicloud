// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Inputs
{

    public sealed class GetKubernetesPermissionPermissionInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ndicates whether the permissions are granted to the cluster owner. Valid values `0`, `1`.
        /// * `is_ram_role` -Indicates whether the permissions are granted to the RAM role. Valid values `0`,`1`.
        /// </summary>
        [Input("isOwner")]
        public Input<bool>? IsOwner { get; set; }

        [Input("isRamRole")]
        public Input<bool>? IsRamRole { get; set; }

        /// <summary>
        /// The permission settings to manage ACK clusters.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// The authorization type. Valid values `cluster`, `namespace` and `console`.
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        /// <summary>
        /// The name of the predefined role. If a custom role is assigned, the value is the name of the assigined custom role.
        /// </summary>
        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

        /// <summary>
        /// The predefined role. Valid values `admin`,`ops`,`dev`,`restricted` and `custom`.
        /// </summary>
        [Input("roleType")]
        public Input<string>? RoleType { get; set; }

        public GetKubernetesPermissionPermissionInputArgs()
        {
        }
    }
}
