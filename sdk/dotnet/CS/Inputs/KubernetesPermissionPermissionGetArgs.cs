// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Inputs
{

    public sealed class KubernetesPermissionPermissionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the cluster that you want to manage.
        /// </summary>
        [Input("cluster", required: true)]
        public Input<string> Cluster { get; set; } = null!;

        /// <summary>
        /// Specifies whether to perform a custom authorization. To perform a custom authorization, set `role_name` to a custom cluster role.
        /// </summary>
        [Input("isCustom")]
        public Input<bool>? IsCustom { get; set; }

        /// <summary>
        /// Specifies whether the permissions are granted to a RAM role. When `uid` is ram role id, the value of `is_ram_role` must be `true`.
        /// </summary>
        [Input("isRamRole")]
        public Input<bool>? IsRamRole { get; set; }

        /// <summary>
        /// The namespace to which the permissions are scoped. This parameter is required only if you set role_type to namespace.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Specifies the predefined role that you want to assign. Valid values `admin`, `ops`, `dev`, `restricted` and the custom cluster roles.
        /// </summary>
        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

        /// <summary>
        /// The authorization type. Valid values `cluster`, `namespace`.
        /// </summary>
        [Input("roleType", required: true)]
        public Input<string> RoleType { get; set; } = null!;

        public KubernetesPermissionPermissionGetArgs()
        {
        }
    }
}
