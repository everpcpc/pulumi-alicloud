// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ServiceMesh.Inputs
{

    public sealed class ServiceMeshMeshConfigKialiGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable of the access logging. Valid values: `true` and `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public ServiceMeshMeshConfigKialiGetArgs()
        {
        }
        public static new ServiceMeshMeshConfigKialiGetArgs Empty => new ServiceMeshMeshConfigKialiGetArgs();
    }
}
