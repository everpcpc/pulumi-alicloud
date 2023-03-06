// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.actiontrail;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.actiontrail.HistoryDeliveryJobArgs;
import com.pulumi.alicloud.actiontrail.inputs.HistoryDeliveryJobState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Actiontrail History Delivery Job resource.
 * 
 * For information about Actiontrail History Delivery Job and how to use it, see [What is History Delivery Job](https://www.alibabacloud.com/help/doc-detail/199999.htm).
 * 
 * &gt; **NOTE:** Available in v1.139.0+.
 * 
 * &gt; **NOTE:** You are authorized to use the historical event delivery task feature. To use this feature, [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=a2c63.p38356.0.0.e29f552bb6odNZ#/ticket/createIndex) or ask the sales manager to add you to the whitelist.
 * 
 * &gt; **NOTE:** Make sure that you have called the `alicloud.actiontrail.Trail` to create a single-account or multi-account trace that delivered to Log Service SLS.
 * 
 * &gt; **NOTE:** An Alibaba cloud account can only have one running delivery history job at the same time.
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetRegionsArgs;
 * import com.pulumi.alicloud.log.Project;
 * import com.pulumi.alicloud.log.ProjectArgs;
 * import com.pulumi.alicloud.actiontrail.Trail;
 * import com.pulumi.alicloud.actiontrail.TrailArgs;
 * import com.pulumi.alicloud.actiontrail.HistoryDeliveryJob;
 * import com.pulumi.alicloud.actiontrail.HistoryDeliveryJobArgs;
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
 *         final var exampleRegions = AlicloudFunctions.getRegions(GetRegionsArgs.builder()
 *             .current(true)
 *             .build());
 * 
 *         final var exampleAccount = AlicloudFunctions.getAccount();
 * 
 *         var exampleProject = new Project(&#34;exampleProject&#34;, ProjectArgs.builder()        
 *             .description(&#34;tf actiontrail test&#34;)
 *             .build());
 * 
 *         var exampleTrail = new Trail(&#34;exampleTrail&#34;, TrailArgs.builder()        
 *             .trailName(&#34;example_value&#34;)
 *             .slsProjectArn(exampleProject.name().applyValue(name -&gt; String.format(&#34;acs:log:%s:%s:project/%s&#34;, exampleRegions.applyValue(getRegionsResult -&gt; getRegionsResult.regions()[0].id()),exampleAccount.applyValue(getAccountResult -&gt; getAccountResult.id()),name)))
 *             .build());
 * 
 *         var exampleHistoryDeliveryJob = new HistoryDeliveryJob(&#34;exampleHistoryDeliveryJob&#34;, HistoryDeliveryJobArgs.builder()        
 *             .trailName(exampleTrail.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Actiontrail History Delivery Job can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:actiontrail/historyDeliveryJob:HistoryDeliveryJob example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:actiontrail/historyDeliveryJob:HistoryDeliveryJob")
public class HistoryDeliveryJob extends com.pulumi.resources.CustomResource {
    /**
     * The status of the task. Valid values: `0`, `1`, `2`, `3`. `0`: The task is initializing. `1`: The task is delivering historical events. `2`: The delivery of historical events is complete. `3`: The task fails.
     * 
     */
    @Export(name="status", type=Integer.class, parameters={})
    private Output<Integer> status;

    /**
     * @return The status of the task. Valid values: `0`, `1`, `2`, `3`. `0`: The task is initializing. `1`: The task is delivering historical events. `2`: The delivery of historical events is complete. `3`: The task fails.
     * 
     */
    public Output<Integer> status() {
        return this.status;
    }
    /**
     * The name of the trail for which you want to create a historical event delivery task.
     * 
     */
    @Export(name="trailName", type=String.class, parameters={})
    private Output<String> trailName;

    /**
     * @return The name of the trail for which you want to create a historical event delivery task.
     * 
     */
    public Output<String> trailName() {
        return this.trailName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HistoryDeliveryJob(String name) {
        this(name, HistoryDeliveryJobArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HistoryDeliveryJob(String name, HistoryDeliveryJobArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HistoryDeliveryJob(String name, HistoryDeliveryJobArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:actiontrail/historyDeliveryJob:HistoryDeliveryJob", name, args == null ? HistoryDeliveryJobArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HistoryDeliveryJob(String name, Output<String> id, @Nullable HistoryDeliveryJobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:actiontrail/historyDeliveryJob:HistoryDeliveryJob", name, state, makeResourceOptions(options, id));
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
    public static HistoryDeliveryJob get(String name, Output<String> id, @Nullable HistoryDeliveryJobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HistoryDeliveryJob(name, id, state, options);
    }
}