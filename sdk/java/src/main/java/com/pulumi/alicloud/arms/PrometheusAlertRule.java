// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.arms.PrometheusAlertRuleArgs;
import com.pulumi.alicloud.arms.inputs.PrometheusAlertRuleState;
import com.pulumi.alicloud.arms.outputs.PrometheusAlertRuleAnnotation;
import com.pulumi.alicloud.arms.outputs.PrometheusAlertRuleLabel;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Application Real-Time Monitoring Service (ARMS) Prometheus Alert Rule resource.
 * 
 * For information about Application Real-Time Monitoring Service (ARMS) Prometheus Alert Rule and how to use it, see [What is Prometheus Alert Rule](https://www.alibabacloud.com/help/en/doc-detail/212056.htm).
 * 
 * &gt; **NOTE:** Available in v1.136.0+.
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
 * import com.pulumi.alicloud.arms.PrometheusAlertRule;
 * import com.pulumi.alicloud.arms.PrometheusAlertRuleArgs;
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
 *         var example = new PrometheusAlertRule(&#34;example&#34;, PrometheusAlertRuleArgs.builder()        
 *             .clusterId(&#34;example_value&#34;)
 *             .duration(&#34;example_value&#34;)
 *             .expression(&#34;example_value&#34;)
 *             .message(&#34;example_value&#34;)
 *             .prometheusAlertRuleName(&#34;example_value&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Application Real-Time Monitoring Service (ARMS) Prometheus Alert Rule can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:arms/prometheusAlertRule:PrometheusAlertRule example &lt;cluster_id&gt;:&lt;prometheus_alert_rule_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:arms/prometheusAlertRule:PrometheusAlertRule")
public class PrometheusAlertRule extends com.pulumi.resources.CustomResource {
    /**
     * The annotations of the alert rule.. See the following `Block annotations`.
     * 
     */
    @Export(name="annotations", type=List.class, parameters={PrometheusAlertRuleAnnotation.class})
    private Output</* @Nullable */ List<PrometheusAlertRuleAnnotation>> annotations;

    /**
     * @return The annotations of the alert rule.. See the following `Block annotations`.
     * 
     */
    public Output<Optional<List<PrometheusAlertRuleAnnotation>>> annotations() {
        return Codegen.optional(this.annotations);
    }
    /**
     * The ID of the cluster.
     * 
     */
    @Export(name="clusterId", type=String.class, parameters={})
    private Output<String> clusterId;

    /**
     * @return The ID of the cluster.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * The ID of the notification policy. This parameter is required when the `notify_type` parameter is set to `DISPATCH_RULE`.
     * 
     */
    @Export(name="dispatchRuleId", type=String.class, parameters={})
    private Output</* @Nullable */ String> dispatchRuleId;

    /**
     * @return The ID of the notification policy. This parameter is required when the `notify_type` parameter is set to `DISPATCH_RULE`.
     * 
     */
    public Output<Optional<String>> dispatchRuleId() {
        return Codegen.optional(this.dispatchRuleId);
    }
    /**
     * The duration of the alert.
     * 
     */
    @Export(name="duration", type=String.class, parameters={})
    private Output<String> duration;

    /**
     * @return The duration of the alert.
     * 
     */
    public Output<String> duration() {
        return this.duration;
    }
    /**
     * The alert rule expression that follows the PromQL syntax.
     * 
     */
    @Export(name="expression", type=String.class, parameters={})
    private Output<String> expression;

    /**
     * @return The alert rule expression that follows the PromQL syntax.
     * 
     */
    public Output<String> expression() {
        return this.expression;
    }
    /**
     * The labels of the resource. See the following `Block labels`.
     * 
     */
    @Export(name="labels", type=List.class, parameters={PrometheusAlertRuleLabel.class})
    private Output</* @Nullable */ List<PrometheusAlertRuleLabel>> labels;

    /**
     * @return The labels of the resource. See the following `Block labels`.
     * 
     */
    public Output<Optional<List<PrometheusAlertRuleLabel>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The message of the alert notification.
     * 
     */
    @Export(name="message", type=String.class, parameters={})
    private Output<String> message;

    /**
     * @return The message of the alert notification.
     * 
     */
    public Output<String> message() {
        return this.message;
    }
    /**
     * The method of sending the alert notification. Valid values: `ALERT_MANAGER`, `DISPATCH_RULE`.
     * 
     */
    @Export(name="notifyType", type=String.class, parameters={})
    private Output</* @Nullable */ String> notifyType;

    /**
     * @return The method of sending the alert notification. Valid values: `ALERT_MANAGER`, `DISPATCH_RULE`.
     * 
     */
    public Output<Optional<String>> notifyType() {
        return Codegen.optional(this.notifyType);
    }
    /**
     * The first ID of the resource.
     * 
     */
    @Export(name="prometheusAlertRuleId", type=Integer.class, parameters={})
    private Output<Integer> prometheusAlertRuleId;

    /**
     * @return The first ID of the resource.
     * 
     */
    public Output<Integer> prometheusAlertRuleId() {
        return this.prometheusAlertRuleId;
    }
    /**
     * The name of the resource.
     * 
     */
    @Export(name="prometheusAlertRuleName", type=String.class, parameters={})
    private Output<String> prometheusAlertRuleName;

    /**
     * @return The name of the resource.
     * 
     */
    public Output<String> prometheusAlertRuleName() {
        return this.prometheusAlertRuleName;
    }
    /**
     * The status of the resource. Valid values: `0`, `1`.
     * 
     */
    @Export(name="status", type=Integer.class, parameters={})
    private Output<Integer> status;

    /**
     * @return The status of the resource. Valid values: `0`, `1`.
     * 
     */
    public Output<Integer> status() {
        return this.status;
    }
    /**
     * The type of the alert rule.
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output<String> type;

    /**
     * @return The type of the alert rule.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PrometheusAlertRule(String name) {
        this(name, PrometheusAlertRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PrometheusAlertRule(String name, PrometheusAlertRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PrometheusAlertRule(String name, PrometheusAlertRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:arms/prometheusAlertRule:PrometheusAlertRule", name, args == null ? PrometheusAlertRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PrometheusAlertRule(String name, Output<String> id, @Nullable PrometheusAlertRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:arms/prometheusAlertRule:PrometheusAlertRule", name, state, makeResourceOptions(options, id));
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
    public static PrometheusAlertRule get(String name, Output<String> id, @Nullable PrometheusAlertRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PrometheusAlertRule(name, id, state, options);
    }
}
