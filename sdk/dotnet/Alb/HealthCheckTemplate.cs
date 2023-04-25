// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb
{
    /// <summary>
    /// Provides a Application Load Balancer (ALB) Health Check Template resource.
    /// 
    /// For information about Application Load Balancer (ALB) Health Check Template and how to use it, see [What is Health Check Template](https://www.alibabacloud.com/help/doc-detail/214343.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.134.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AliCloud.Alb.HealthCheckTemplate("example", new()
    ///     {
    ///         HealthCheckTemplateName = "example_name",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Application Load Balancer (ALB) Health Check Template can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:alb/healthCheckTemplate:HealthCheckTemplate example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:alb/healthCheckTemplate:HealthCheckTemplate")]
    public partial class HealthCheckTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to precheck the API request.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The HTTP status code that indicates a successful health check. **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
        /// </summary>
        [Output("healthCheckCodes")]
        public Output<ImmutableArray<string>> HealthCheckCodes { get; private set; } = null!;

        /// <summary>
        /// The number of the port that is used for health checks.  Valid values: `0` to `65535`.  Default value: `0`. This default value indicates that the backend server is used for health checks.
        /// </summary>
        [Output("healthCheckConnectPort")]
        public Output<int> HealthCheckConnectPort { get; private set; } = null!;

        /// <summary>
        /// The domain name that is used for health checks. Default value:  `$SERVER_IP`. The domain name must be 1 to 80 characters in length.  **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
        /// </summary>
        [Output("healthCheckHost")]
        public Output<string> HealthCheckHost { get; private set; } = null!;

        /// <summary>
        /// The version of the HTTP protocol.  Valid values: `HTTP1.0` and `HTTP1.1`.  Default value: `HTTP1.1`. **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
        /// </summary>
        [Output("healthCheckHttpVersion")]
        public Output<string> HealthCheckHttpVersion { get; private set; } = null!;

        /// <summary>
        /// The time interval between two consecutive health checks.  Valid values: `1` to `50`. Unit: seconds.  Default value: `2`.
        /// </summary>
        [Output("healthCheckInterval")]
        public Output<int> HealthCheckInterval { get; private set; } = null!;

        /// <summary>
        /// The health check method.  Valid values: GET and HEAD.  Default value: HEAD. **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
        /// </summary>
        [Output("healthCheckMethod")]
        public Output<string> HealthCheckMethod { get; private set; } = null!;

        /// <summary>
        /// The URL that is used for health checks.  The URL must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), and ampersands (&amp;). The URL can also contain the following extended characters: _ ; ~ ! ( )* [ ] @ $ ^ : ' , +. The URL must start with a forward slash (/). **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
        /// </summary>
        [Output("healthCheckPath")]
        public Output<string> HealthCheckPath { get; private set; } = null!;

        /// <summary>
        /// The protocol that is used for health checks.  Valid values: `HTTP` and `TCP`.  Default value: `HTTP`.
        /// </summary>
        [Output("healthCheckProtocol")]
        public Output<string> HealthCheckProtocol { get; private set; } = null!;

        /// <summary>
        /// The name of the health check template.  The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
        /// </summary>
        [Output("healthCheckTemplateName")]
        public Output<string> HealthCheckTemplateName { get; private set; } = null!;

        /// <summary>
        /// The timeout period of a health check response. If the backend Elastic Compute Service (ECS) instance does not send an expected response within the specified period of time, the health check fails.  Valid values: `1` to `300`. Unit: seconds.  Default value: `5`.
        /// </summary>
        [Output("healthCheckTimeout")]
        public Output<int> HealthCheckTimeout { get; private set; } = null!;

        /// <summary>
        /// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy (from fail to success).  Valid values: `2` to `10`.  Default value: `3`. Unit: seconds.
        /// </summary>
        [Output("healthyThreshold")]
        public Output<int> HealthyThreshold { get; private set; } = null!;

        /// <summary>
        /// The number of times that an healthy backend server must consecutively fail health checks before it is declared unhealthy (from success to fail). Valid values: `2` to `10`.  Default value: `3`. Unit: seconds.
        /// </summary>
        [Output("unhealthyThreshold")]
        public Output<int> UnhealthyThreshold { get; private set; } = null!;


        /// <summary>
        /// Create a HealthCheckTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HealthCheckTemplate(string name, HealthCheckTemplateArgs args, CustomResourceOptions? options = null)
            : base("alicloud:alb/healthCheckTemplate:HealthCheckTemplate", name, args ?? new HealthCheckTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HealthCheckTemplate(string name, Input<string> id, HealthCheckTemplateState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:alb/healthCheckTemplate:HealthCheckTemplate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HealthCheckTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HealthCheckTemplate Get(string name, Input<string> id, HealthCheckTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new HealthCheckTemplate(name, id, state, options);
        }
    }

    public sealed class HealthCheckTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to precheck the API request.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        [Input("healthCheckCodes")]
        private InputList<string>? _healthCheckCodes;

        /// <summary>
        /// The HTTP status code that indicates a successful health check. **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
        /// </summary>
        public InputList<string> HealthCheckCodes
        {
            get => _healthCheckCodes ?? (_healthCheckCodes = new InputList<string>());
            set => _healthCheckCodes = value;
        }

        /// <summary>
        /// The number of the port that is used for health checks.  Valid values: `0` to `65535`.  Default value: `0`. This default value indicates that the backend server is used for health checks.
        /// </summary>
        [Input("healthCheckConnectPort")]
        public Input<int>? HealthCheckConnectPort { get; set; }

        /// <summary>
        /// The domain name that is used for health checks. Default value:  `$SERVER_IP`. The domain name must be 1 to 80 characters in length.  **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
        /// </summary>
        [Input("healthCheckHost")]
        public Input<string>? HealthCheckHost { get; set; }

        /// <summary>
        /// The version of the HTTP protocol.  Valid values: `HTTP1.0` and `HTTP1.1`.  Default value: `HTTP1.1`. **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
        /// </summary>
        [Input("healthCheckHttpVersion")]
        public Input<string>? HealthCheckHttpVersion { get; set; }

        /// <summary>
        /// The time interval between two consecutive health checks.  Valid values: `1` to `50`. Unit: seconds.  Default value: `2`.
        /// </summary>
        [Input("healthCheckInterval")]
        public Input<int>? HealthCheckInterval { get; set; }

        /// <summary>
        /// The health check method.  Valid values: GET and HEAD.  Default value: HEAD. **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
        /// </summary>
        [Input("healthCheckMethod")]
        public Input<string>? HealthCheckMethod { get; set; }

        /// <summary>
        /// The URL that is used for health checks.  The URL must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), and ampersands (&amp;). The URL can also contain the following extended characters: _ ; ~ ! ( )* [ ] @ $ ^ : ' , +. The URL must start with a forward slash (/). **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
        /// </summary>
        [Input("healthCheckPath")]
        public Input<string>? HealthCheckPath { get; set; }

        /// <summary>
        /// The protocol that is used for health checks.  Valid values: `HTTP` and `TCP`.  Default value: `HTTP`.
        /// </summary>
        [Input("healthCheckProtocol")]
        public Input<string>? HealthCheckProtocol { get; set; }

        /// <summary>
        /// The name of the health check template.  The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
        /// </summary>
        [Input("healthCheckTemplateName", required: true)]
        public Input<string> HealthCheckTemplateName { get; set; } = null!;

        /// <summary>
        /// The timeout period of a health check response. If the backend Elastic Compute Service (ECS) instance does not send an expected response within the specified period of time, the health check fails.  Valid values: `1` to `300`. Unit: seconds.  Default value: `5`.
        /// </summary>
        [Input("healthCheckTimeout")]
        public Input<int>? HealthCheckTimeout { get; set; }

        /// <summary>
        /// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy (from fail to success).  Valid values: `2` to `10`.  Default value: `3`. Unit: seconds.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// The number of times that an healthy backend server must consecutively fail health checks before it is declared unhealthy (from success to fail). Valid values: `2` to `10`.  Default value: `3`. Unit: seconds.
        /// </summary>
        [Input("unhealthyThreshold")]
        public Input<int>? UnhealthyThreshold { get; set; }

        public HealthCheckTemplateArgs()
        {
        }
        public static new HealthCheckTemplateArgs Empty => new HealthCheckTemplateArgs();
    }

    public sealed class HealthCheckTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to precheck the API request.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        [Input("healthCheckCodes")]
        private InputList<string>? _healthCheckCodes;

        /// <summary>
        /// The HTTP status code that indicates a successful health check. **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
        /// </summary>
        public InputList<string> HealthCheckCodes
        {
            get => _healthCheckCodes ?? (_healthCheckCodes = new InputList<string>());
            set => _healthCheckCodes = value;
        }

        /// <summary>
        /// The number of the port that is used for health checks.  Valid values: `0` to `65535`.  Default value: `0`. This default value indicates that the backend server is used for health checks.
        /// </summary>
        [Input("healthCheckConnectPort")]
        public Input<int>? HealthCheckConnectPort { get; set; }

        /// <summary>
        /// The domain name that is used for health checks. Default value:  `$SERVER_IP`. The domain name must be 1 to 80 characters in length.  **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
        /// </summary>
        [Input("healthCheckHost")]
        public Input<string>? HealthCheckHost { get; set; }

        /// <summary>
        /// The version of the HTTP protocol.  Valid values: `HTTP1.0` and `HTTP1.1`.  Default value: `HTTP1.1`. **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
        /// </summary>
        [Input("healthCheckHttpVersion")]
        public Input<string>? HealthCheckHttpVersion { get; set; }

        /// <summary>
        /// The time interval between two consecutive health checks.  Valid values: `1` to `50`. Unit: seconds.  Default value: `2`.
        /// </summary>
        [Input("healthCheckInterval")]
        public Input<int>? HealthCheckInterval { get; set; }

        /// <summary>
        /// The health check method.  Valid values: GET and HEAD.  Default value: HEAD. **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
        /// </summary>
        [Input("healthCheckMethod")]
        public Input<string>? HealthCheckMethod { get; set; }

        /// <summary>
        /// The URL that is used for health checks.  The URL must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), and ampersands (&amp;). The URL can also contain the following extended characters: _ ; ~ ! ( )* [ ] @ $ ^ : ' , +. The URL must start with a forward slash (/). **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
        /// </summary>
        [Input("healthCheckPath")]
        public Input<string>? HealthCheckPath { get; set; }

        /// <summary>
        /// The protocol that is used for health checks.  Valid values: `HTTP` and `TCP`.  Default value: `HTTP`.
        /// </summary>
        [Input("healthCheckProtocol")]
        public Input<string>? HealthCheckProtocol { get; set; }

        /// <summary>
        /// The name of the health check template.  The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
        /// </summary>
        [Input("healthCheckTemplateName")]
        public Input<string>? HealthCheckTemplateName { get; set; }

        /// <summary>
        /// The timeout period of a health check response. If the backend Elastic Compute Service (ECS) instance does not send an expected response within the specified period of time, the health check fails.  Valid values: `1` to `300`. Unit: seconds.  Default value: `5`.
        /// </summary>
        [Input("healthCheckTimeout")]
        public Input<int>? HealthCheckTimeout { get; set; }

        /// <summary>
        /// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy (from fail to success).  Valid values: `2` to `10`.  Default value: `3`. Unit: seconds.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// The number of times that an healthy backend server must consecutively fail health checks before it is declared unhealthy (from success to fail). Valid values: `2` to `10`.  Default value: `3`. Unit: seconds.
        /// </summary>
        [Input("unhealthyThreshold")]
        public Input<int>? UnhealthyThreshold { get; set; }

        public HealthCheckTemplateState()
        {
        }
        public static new HealthCheckTemplateState Empty => new HealthCheckTemplateState();
    }
}
