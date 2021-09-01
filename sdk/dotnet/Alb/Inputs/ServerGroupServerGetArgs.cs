// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Inputs
{

    public sealed class ServerGroupServerGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the server.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The port that is used by the server. Valid values: `1` to `65535`.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The ID of the ECS instance, ENI instance or ECI instance.
        /// </summary>
        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

        /// <summary>
        /// The IP address of the ENI instance when it is in the inclusive ENI mode.
        /// </summary>
        [Input("serverIp")]
        public Input<string>? ServerIp { get; set; }

        /// <summary>
        /// The type of the server. The type of the server. Valid values: `Ecs`, `Eni` and `Eci`.
        /// </summary>
        [Input("serverType")]
        public Input<string>? ServerType { get; set; }

        /// <summary>
        /// The status of the resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The weight of the server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no
        /// requests are forwarded to the server.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ServerGroupServerGetArgs()
        {
        }
    }
}