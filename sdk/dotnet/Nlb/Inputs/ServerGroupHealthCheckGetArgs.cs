// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Nlb.Inputs
{

    public sealed class ServerGroupHealthCheckGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The backend port that is used for health checks. Valid values: 0 to 65535. Default value: 0. If you set the value to 0, the port of a backend server is used for health checks.
        /// </summary>
        [Input("healthCheckConnectPort")]
        public Input<int>? HealthCheckConnectPort { get; set; }

        /// <summary>
        /// The maximum timeout period of a health check response. Unit: seconds. Valid values: 1 to 300. Default value: 5.
        /// </summary>
        [Input("healthCheckConnectTimeout")]
        public Input<int>? HealthCheckConnectTimeout { get; set; }

        /// <summary>
        /// The domain name that is used for health checks. Valid values:
        /// - `$SERVER_IP`: the private IP address of a backend server.
        /// </summary>
        [Input("healthCheckDomain")]
        public Input<string>? HealthCheckDomain { get; set; }

        /// <summary>
        /// Specifies whether to enable health checks.
        /// </summary>
        [Input("healthCheckEnabled")]
        public Input<bool>? HealthCheckEnabled { get; set; }

        [Input("healthCheckHttpCodes")]
        private InputList<string>? _healthCheckHttpCodes;

        /// <summary>
        /// The HTTP status codes to return to health checks. Separate multiple HTTP status codes with commas (,). Valid values: http_2xx (default), http_3xx, http_4xx, and http_5xx. **Note:** This parameter takes effect only if `health_check_type` is set to `http`.
        /// </summary>
        public InputList<string> HealthCheckHttpCodes
        {
            get => _healthCheckHttpCodes ?? (_healthCheckHttpCodes = new InputList<string>());
            set => _healthCheckHttpCodes = value;
        }

        /// <summary>
        /// The interval between two consecutive health checks. Unit: seconds. Valid values: 5 to 5000. Default value: 10.
        /// </summary>
        [Input("healthCheckInterval")]
        public Input<int>? HealthCheckInterval { get; set; }

        /// <summary>
        /// The protocol that is used for health checks. Valid values: `TCP` (default) and `HTTP`.
        /// </summary>
        [Input("healthCheckType")]
        public Input<string>? HealthCheckType { get; set; }

        /// <summary>
        /// The path to which health check requests are sent. The path must be 1 to 80 characters in length, and can contain only letters, digits, and the following special characters: `- / . % ? # &amp; =`. It can also contain the following extended characters: `_ ; ~ ! ( ) * [ ] @ $ ^ : ' , +`. The path must start with a forward slash (/). **Note:** This parameter takes effect only if `health_check_type` is set to `http`.
        /// </summary>
        [Input("healthCheckUrl")]
        public Input<string>? HealthCheckUrl { get; set; }

        /// <summary>
        /// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from fail to success. Valid values: 2 to 10. Default value: 2.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// The HTTP method that is used for health checks. Valid values: `GET` and `HEAD`. **Note:** This parameter takes effect only if `health_check_type` is set to `http`.
        /// </summary>
        [Input("httpCheckMethod")]
        public Input<string>? HttpCheckMethod { get; set; }

        /// <summary>
        /// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from success to fail. Valid values: 2 to 10. Default value: 2.
        /// </summary>
        [Input("unhealthyThreshold")]
        public Input<int>? UnhealthyThreshold { get; set; }

        public ServerGroupHealthCheckGetArgs()
        {
        }
        public static new ServerGroupHealthCheckGetArgs Empty => new ServerGroupHealthCheckGetArgs();
    }
}
