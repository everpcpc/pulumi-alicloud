// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cen.InterRegionTrafficQosQueueArgs;
import com.pulumi.alicloud.cen.inputs.InterRegionTrafficQosQueueState;
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
 * Provides a Cen Inter Region Traffic Qos Queue resource.
 * 
 * For information about Cen Inter Region Traffic Qos Queue and how to use it, see [What is Inter Region Traffic Qos Queue](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/api-doc-cbn-2017-09-12-api-doc-createceninterregiontrafficqosqueue).
 * 
 * &gt; **NOTE:** Available in v1.195.0+.
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
 * import com.pulumi.alicloud.cen.InterRegionTrafficQosQueue;
 * import com.pulumi.alicloud.cen.InterRegionTrafficQosQueueArgs;
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
 *         var default_ = new InterRegionTrafficQosQueue(&#34;default&#34;, InterRegionTrafficQosQueueArgs.builder()        
 *             .dscps(            
 *                 1,
 *                 2)
 *             .interRegionTrafficQosQueueDescription(&#34;test&#34;)
 *             .remainBandwidthPercent(20)
 *             .trafficQosPolicyId(&#34;qos-xxxxxx&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Cen Inter Region Traffic Qos Queue can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:cen/interRegionTrafficQosQueue:InterRegionTrafficQosQueue example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cen/interRegionTrafficQosQueue:InterRegionTrafficQosQueue")
public class InterRegionTrafficQosQueue extends com.pulumi.resources.CustomResource {
    /**
     * The DSCP value of the traffic packet to be matched in the current queue, ranging from 0 to 63.
     * 
     */
    @Export(name="dscps", type=List.class, parameters={String.class})
    private Output<List<String>> dscps;

    /**
     * @return The DSCP value of the traffic packet to be matched in the current queue, ranging from 0 to 63.
     * 
     */
    public Output<List<String>> dscps() {
        return this.dscps;
    }
    /**
     * The description information of the traffic scheduling policy.
     * 
     */
    @Export(name="interRegionTrafficQosQueueDescription", type=String.class, parameters={})
    private Output</* @Nullable */ String> interRegionTrafficQosQueueDescription;

    /**
     * @return The description information of the traffic scheduling policy.
     * 
     */
    public Output<Optional<String>> interRegionTrafficQosQueueDescription() {
        return Codegen.optional(this.interRegionTrafficQosQueueDescription);
    }
    /**
     * The name of the traffic scheduling policy.
     * 
     */
    @Export(name="interRegionTrafficQosQueueName", type=String.class, parameters={})
    private Output</* @Nullable */ String> interRegionTrafficQosQueueName;

    /**
     * @return The name of the traffic scheduling policy.
     * 
     */
    public Output<Optional<String>> interRegionTrafficQosQueueName() {
        return Codegen.optional(this.interRegionTrafficQosQueueName);
    }
    /**
     * The percentage of cross-region bandwidth that the current queue can use.
     * 
     */
    @Export(name="remainBandwidthPercent", type=Integer.class, parameters={})
    private Output<Integer> remainBandwidthPercent;

    /**
     * @return The percentage of cross-region bandwidth that the current queue can use.
     * 
     */
    public Output<Integer> remainBandwidthPercent() {
        return this.remainBandwidthPercent;
    }
    /**
     * The status of the traffic scheduling policy. -**Creating**: The function is being created.-**Active**: available.-**Modifying**: is being modified.-**Deleting**: Deleted.-**Deleted**: Deleted.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the traffic scheduling policy. -**Creating**: The function is being created.-**Active**: available.-**Modifying**: is being modified.-**Deleting**: Deleted.-**Deleted**: Deleted.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The ID of the traffic scheduling policy.
     * 
     */
    @Export(name="trafficQosPolicyId", type=String.class, parameters={})
    private Output<String> trafficQosPolicyId;

    /**
     * @return The ID of the traffic scheduling policy.
     * 
     */
    public Output<String> trafficQosPolicyId() {
        return this.trafficQosPolicyId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InterRegionTrafficQosQueue(String name) {
        this(name, InterRegionTrafficQosQueueArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InterRegionTrafficQosQueue(String name, InterRegionTrafficQosQueueArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InterRegionTrafficQosQueue(String name, InterRegionTrafficQosQueueArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cen/interRegionTrafficQosQueue:InterRegionTrafficQosQueue", name, args == null ? InterRegionTrafficQosQueueArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InterRegionTrafficQosQueue(String name, Output<String> id, @Nullable InterRegionTrafficQosQueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cen/interRegionTrafficQosQueue:InterRegionTrafficQosQueue", name, state, makeResourceOptions(options, id));
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
    public static InterRegionTrafficQosQueue get(String name, Output<String> id, @Nullable InterRegionTrafficQosQueueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InterRegionTrafficQosQueue(name, id, state, options);
    }
}
