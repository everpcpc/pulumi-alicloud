// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.alb.HealthCheckTemplateArgs;
import com.pulumi.alicloud.alb.inputs.HealthCheckTemplateState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Application Load Balancer (ALB) Health Check Template resource.
 * 
 * For information about Application Load Balancer (ALB) Health Check Template and how to use it, see [What is Health Check Template](https://www.alibabacloud.com/help/doc-detail/214343.htm).
 * 
 * &gt; **NOTE:** Available in v1.134.0+.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.alb.HealthCheckTemplate;
 * import com.pulumi.alicloud.alb.HealthCheckTemplateArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new HealthCheckTemplate(&#34;example&#34;, HealthCheckTemplateArgs.builder()        
 *             .healthCheckTemplateName(&#34;example_name&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Application Load Balancer (ALB) Health Check Template can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:alb/healthCheckTemplate:HealthCheckTemplate example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:alb/healthCheckTemplate:HealthCheckTemplate")
public class HealthCheckTemplate extends com.pulumi.resources.CustomResource {
    /**
     * Whether to precheck the API request.
     * 
     */
    @Export(name="dryRun", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return Whether to precheck the API request.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * The HTTP status code that indicates a successful health check. **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
     * 
     */
    @Export(name="healthCheckCodes", type=List.class, parameters={String.class})
    private Output<List<String>> healthCheckCodes;

    /**
     * @return The HTTP status code that indicates a successful health check. **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
     * 
     */
    public Output<List<String>> healthCheckCodes() {
        return this.healthCheckCodes;
    }
    /**
     * The number of the port that is used for health checks.  Valid values: `0` to `65535`.  Default value: `0`. This default value indicates that the backend server is used for health checks.
     * 
     */
    @Export(name="healthCheckConnectPort", type=Integer.class, parameters={})
    private Output<Integer> healthCheckConnectPort;

    /**
     * @return The number of the port that is used for health checks.  Valid values: `0` to `65535`.  Default value: `0`. This default value indicates that the backend server is used for health checks.
     * 
     */
    public Output<Integer> healthCheckConnectPort() {
        return this.healthCheckConnectPort;
    }
    /**
     * The domain name that is used for health checks. Default value:  `$SERVER_IP`. The domain name must be 1 to 80 characters in length.  **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
     * 
     */
    @Export(name="healthCheckHost", type=String.class, parameters={})
    private Output<String> healthCheckHost;

    /**
     * @return The domain name that is used for health checks. Default value:  `$SERVER_IP`. The domain name must be 1 to 80 characters in length.  **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
     * 
     */
    public Output<String> healthCheckHost() {
        return this.healthCheckHost;
    }
    /**
     * The version of the HTTP protocol.  Valid values: `HTTP1.0` and `HTTP1.1`.  Default value: `HTTP1.1`. **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
     * 
     */
    @Export(name="healthCheckHttpVersion", type=String.class, parameters={})
    private Output<String> healthCheckHttpVersion;

    /**
     * @return The version of the HTTP protocol.  Valid values: `HTTP1.0` and `HTTP1.1`.  Default value: `HTTP1.1`. **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
     * 
     */
    public Output<String> healthCheckHttpVersion() {
        return this.healthCheckHttpVersion;
    }
    /**
     * The time interval between two consecutive health checks.  Valid values: `1` to `50`. Unit: seconds.  Default value: `2`.
     * 
     */
    @Export(name="healthCheckInterval", type=Integer.class, parameters={})
    private Output<Integer> healthCheckInterval;

    /**
     * @return The time interval between two consecutive health checks.  Valid values: `1` to `50`. Unit: seconds.  Default value: `2`.
     * 
     */
    public Output<Integer> healthCheckInterval() {
        return this.healthCheckInterval;
    }
    /**
     * The health check method.  Valid values: GET and HEAD.  Default value: HEAD. **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
     * 
     */
    @Export(name="healthCheckMethod", type=String.class, parameters={})
    private Output<String> healthCheckMethod;

    /**
     * @return The health check method.  Valid values: GET and HEAD.  Default value: HEAD. **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
     * 
     */
    public Output<String> healthCheckMethod() {
        return this.healthCheckMethod;
    }
    /**
     * The URL that is used for health checks.  The URL must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), and ampersands (&amp;). The URL can also contain the following extended characters: _ ; ~ ! ( )* [ ] @ $ ^ : &#39; , +. The URL must start with a forward slash (/). **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
     * 
     */
    @Export(name="healthCheckPath", type=String.class, parameters={})
    private Output<String> healthCheckPath;

    /**
     * @return The URL that is used for health checks.  The URL must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), and ampersands (&amp;). The URL can also contain the following extended characters: _ ; ~ ! ( )* [ ] @ $ ^ : &#39; , +. The URL must start with a forward slash (/). **NOTE:** The attribute `HealthCheckProtocol` is valid when the attribute is  `HTTP` .
     * 
     */
    public Output<String> healthCheckPath() {
        return this.healthCheckPath;
    }
    /**
     * The protocol that is used for health checks.  Valid values: `HTTP` and `TCP`.  Default value: `HTTP`.
     * 
     */
    @Export(name="healthCheckProtocol", type=String.class, parameters={})
    private Output<String> healthCheckProtocol;

    /**
     * @return The protocol that is used for health checks.  Valid values: `HTTP` and `TCP`.  Default value: `HTTP`.
     * 
     */
    public Output<String> healthCheckProtocol() {
        return this.healthCheckProtocol;
    }
    /**
     * The name of the health check template.  The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    @Export(name="healthCheckTemplateName", type=String.class, parameters={})
    private Output<String> healthCheckTemplateName;

    /**
     * @return The name of the health check template.  The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    public Output<String> healthCheckTemplateName() {
        return this.healthCheckTemplateName;
    }
    /**
     * The timeout period of a health check response. If the backend Elastic Compute Service (ECS) instance does not send an expected response within the specified period of time, the health check fails.  Valid values: `1` to `300`. Unit: seconds.  Default value: `5`.
     * 
     */
    @Export(name="healthCheckTimeout", type=Integer.class, parameters={})
    private Output<Integer> healthCheckTimeout;

    /**
     * @return The timeout period of a health check response. If the backend Elastic Compute Service (ECS) instance does not send an expected response within the specified period of time, the health check fails.  Valid values: `1` to `300`. Unit: seconds.  Default value: `5`.
     * 
     */
    public Output<Integer> healthCheckTimeout() {
        return this.healthCheckTimeout;
    }
    /**
     * The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy (from fail to success).  Valid values: `2` to `10`.  Default value: `3`. Unit: seconds.
     * 
     */
    @Export(name="healthyThreshold", type=Integer.class, parameters={})
    private Output<Integer> healthyThreshold;

    /**
     * @return The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy (from fail to success).  Valid values: `2` to `10`.  Default value: `3`. Unit: seconds.
     * 
     */
    public Output<Integer> healthyThreshold() {
        return this.healthyThreshold;
    }
    /**
     * The number of times that an healthy backend server must consecutively fail health checks before it is declared unhealthy (from success to fail). Valid values: `2` to `10`.  Default value: `3`. Unit: seconds.
     * 
     */
    @Export(name="unhealthyThreshold", type=Integer.class, parameters={})
    private Output<Integer> unhealthyThreshold;

    /**
     * @return The number of times that an healthy backend server must consecutively fail health checks before it is declared unhealthy (from success to fail). Valid values: `2` to `10`.  Default value: `3`. Unit: seconds.
     * 
     */
    public Output<Integer> unhealthyThreshold() {
        return this.unhealthyThreshold;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HealthCheckTemplate(String name) {
        this(name, HealthCheckTemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HealthCheckTemplate(String name, HealthCheckTemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HealthCheckTemplate(String name, HealthCheckTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alb/healthCheckTemplate:HealthCheckTemplate", name, args == null ? HealthCheckTemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HealthCheckTemplate(String name, Output<String> id, @Nullable HealthCheckTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:alb/healthCheckTemplate:HealthCheckTemplate", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static HealthCheckTemplate get(String name, Output<String> id, @Nullable HealthCheckTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HealthCheckTemplate(name, id, state, options);
    }
}