// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rocketmq;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.rocketmq.QosPolicyArgs;
import com.pulumi.alicloud.rocketmq.inputs.QosPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Sag qos policy resource.
 * You need to create a QoS policy to set priorities, rate limits, and quintuple rules for different messages.
 * 
 * For information about Sag Qos Policy and how to use it, see [What is Qos Policy](https://www.alibabacloud.com/help/doc-detail/140065.htm).
 * 
 * &gt; **NOTE:** Available in 1.60.0+
 * 
 * &gt; **NOTE:** Only the following regions support. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
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
 * import com.pulumi.alicloud.rocketmq.Qos;
 * import com.pulumi.alicloud.rocketmq.QosPolicy;
 * import com.pulumi.alicloud.rocketmq.QosPolicyArgs;
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
 *         var defaultQos = new Qos(&#34;defaultQos&#34;);
 * 
 *         var defaultQosPolicy = new QosPolicy(&#34;defaultQosPolicy&#34;, QosPolicyArgs.builder()        
 *             .qosId(defaultQos.id())
 *             .description(&#34;tf-testSagQosPolicyDescription&#34;)
 *             .priority(&#34;1&#34;)
 *             .ipProtocol(&#34;ALL&#34;)
 *             .sourceCidr(&#34;192.168.0.0/24&#34;)
 *             .sourcePortRange(&#34;-1/-1&#34;)
 *             .destCidr(&#34;10.10.0.0/24&#34;)
 *             .destPortRange(&#34;-1/-1&#34;)
 *             .startTime(&#34;2019-10-25T16:41:33+0800&#34;)
 *             .endTime(&#34;2019-10-26T16:41:33+0800&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * The Sag Qos Policy can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:rocketmq/qosPolicy:QosPolicy example qos-abc123456:qospy-abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:rocketmq/qosPolicy:QosPolicy")
public class QosPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The description of the QoS policy.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the QoS policy.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The destination CIDR block.
     * 
     */
    @Export(name="destCidr", type=String.class, parameters={})
    private Output<String> destCidr;

    /**
     * @return The destination CIDR block.
     * 
     */
    public Output<String> destCidr() {
        return this.destCidr;
    }
    /**
     * The destination port range.
     * 
     */
    @Export(name="destPortRange", type=String.class, parameters={})
    private Output<String> destPortRange;

    /**
     * @return The destination port range.
     * 
     */
    public Output<String> destPortRange() {
        return this.destPortRange;
    }
    /**
     * The expiration time of the quintuple rule.
     * 
     */
    @Export(name="endTime", type=String.class, parameters={})
    private Output</* @Nullable */ String> endTime;

    /**
     * @return The expiration time of the quintuple rule.
     * 
     */
    public Output<Optional<String>> endTime() {
        return Codegen.optional(this.endTime);
    }
    /**
     * The transport layer protocol.
     * 
     */
    @Export(name="ipProtocol", type=String.class, parameters={})
    private Output<String> ipProtocol;

    /**
     * @return The transport layer protocol.
     * 
     */
    public Output<String> ipProtocol() {
        return this.ipProtocol;
    }
    /**
     * The name of the QoS policy.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the QoS policy.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The priority of the quintuple rule. A smaller value indicates a higher priority. If the priorities of two quintuple rules are the same, the rule created earlier is applied first.Value range: 1 to 7.
     * 
     */
    @Export(name="priority", type=Integer.class, parameters={})
    private Output<Integer> priority;

    /**
     * @return The priority of the quintuple rule. A smaller value indicates a higher priority. If the priorities of two quintuple rules are the same, the rule created earlier is applied first.Value range: 1 to 7.
     * 
     */
    public Output<Integer> priority() {
        return this.priority;
    }
    /**
     * The instance ID of the QoS policy to which the quintuple rule is created.
     * 
     */
    @Export(name="qosId", type=String.class, parameters={})
    private Output<String> qosId;

    /**
     * @return The instance ID of the QoS policy to which the quintuple rule is created.
     * 
     */
    public Output<String> qosId() {
        return this.qosId;
    }
    /**
     * The source CIDR block.
     * 
     */
    @Export(name="sourceCidr", type=String.class, parameters={})
    private Output<String> sourceCidr;

    /**
     * @return The source CIDR block.
     * 
     */
    public Output<String> sourceCidr() {
        return this.sourceCidr;
    }
    /**
     * The source port range of the transport layer.
     * 
     */
    @Export(name="sourcePortRange", type=String.class, parameters={})
    private Output<String> sourcePortRange;

    /**
     * @return The source port range of the transport layer.
     * 
     */
    public Output<String> sourcePortRange() {
        return this.sourcePortRange;
    }
    /**
     * The time when the quintuple rule takes effect.
     * 
     */
    @Export(name="startTime", type=String.class, parameters={})
    private Output</* @Nullable */ String> startTime;

    /**
     * @return The time when the quintuple rule takes effect.
     * 
     */
    public Output<Optional<String>> startTime() {
        return Codegen.optional(this.startTime);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public QosPolicy(String name) {
        this(name, QosPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public QosPolicy(String name, QosPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public QosPolicy(String name, QosPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rocketmq/qosPolicy:QosPolicy", name, args == null ? QosPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private QosPolicy(String name, Output<String> id, @Nullable QosPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rocketmq/qosPolicy:QosPolicy", name, state, makeResourceOptions(options, id));
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
    public static QosPolicy get(String name, Output<String> id, @Nullable QosPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new QosPolicy(name, id, state, options);
    }
}
