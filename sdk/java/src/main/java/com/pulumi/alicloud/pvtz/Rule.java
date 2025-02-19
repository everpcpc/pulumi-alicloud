// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.pvtz;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.pvtz.RuleArgs;
import com.pulumi.alicloud.pvtz.inputs.RuleState;
import com.pulumi.alicloud.pvtz.outputs.RuleForwardIp;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Private Zone Rule resource.
 * 
 * For information about Private Zone Rule and how to use it, see [What is Rule](https://www.alibabacloud.com/help/en/doc-detail/177601.htm).
 * 
 * &gt; **NOTE:** Available in v1.143.0+.
 * 
 * ## Import
 * 
 * Private Zone Rule can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:pvtz/rule:Rule example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:pvtz/rule:Rule")
public class Rule extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the Endpoint.
     * 
     */
    @Export(name="endpointId", type=String.class, parameters={})
    private Output<String> endpointId;

    /**
     * @return The ID of the Endpoint.
     * 
     */
    public Output<String> endpointId() {
        return this.endpointId;
    }
    /**
     * Forwarding target. See the following `Block forward_ip`.
     * 
     */
    @Export(name="forwardIps", type=List.class, parameters={RuleForwardIp.class})
    private Output<List<RuleForwardIp>> forwardIps;

    /**
     * @return Forwarding target. See the following `Block forward_ip`.
     * 
     */
    public Output<List<RuleForwardIp>> forwardIps() {
        return this.forwardIps;
    }
    /**
     * The name of the resource.
     * 
     */
    @Export(name="ruleName", type=String.class, parameters={})
    private Output<String> ruleName;

    /**
     * @return The name of the resource.
     * 
     */
    public Output<String> ruleName() {
        return this.ruleName;
    }
    /**
     * The type of the rule. Valid values: `OUTBOUND`.
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output</* @Nullable */ String> type;

    /**
     * @return The type of the rule. Valid values: `OUTBOUND`.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }
    /**
     * The name of the forwarding zone.
     * 
     */
    @Export(name="zoneName", type=String.class, parameters={})
    private Output<String> zoneName;

    /**
     * @return The name of the forwarding zone.
     * 
     */
    public Output<String> zoneName() {
        return this.zoneName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Rule(String name) {
        this(name, RuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Rule(String name, RuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Rule(String name, RuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:pvtz/rule:Rule", name, args == null ? RuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Rule(String name, Output<String> id, @Nullable RuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:pvtz/rule:Rule", name, state, makeResourceOptions(options, id));
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
    public static Rule get(String name, Output<String> id, @Nullable RuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Rule(name, id, state, options);
    }
}
