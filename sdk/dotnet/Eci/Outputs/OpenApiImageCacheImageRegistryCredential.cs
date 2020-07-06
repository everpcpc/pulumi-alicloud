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
    public sealed class OpenApiImageCacheImageRegistryCredential
    {
        /// <summary>
        /// The password of the Image Registry.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// The address of Image Registry without `http://` or `https://`.
        /// </summary>
        public readonly string? Server;
        /// <summary>
        /// The user name of Image Registry.
        /// </summary>
        public readonly string? UserName;

        [OutputConstructor]
        private OpenApiImageCacheImageRegistryCredential(
            string? password,

            string? server,

            string? userName)
        {
            Password = password;
            Server = server;
            UserName = userName;
        }
    }
}
