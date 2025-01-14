// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dts;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dts.JobMonitorRuleArgs;
import com.pulumi.alicloud.dts.inputs.JobMonitorRuleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a DTS Job Monitor Rule resource.
 * 
 * For information about DTS Job Monitor Rule and how to use it, see [What is Job Monitor Rule](https://www.aliyun.com/product/dts).
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
 * import com.pulumi.alicloud.dts.JobMonitorRule;
 * import com.pulumi.alicloud.dts.JobMonitorRuleArgs;
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
 *         var example = new JobMonitorRule(&#34;example&#34;, JobMonitorRuleArgs.builder()        
 *             .dtsJobId(&#34;example_value&#34;)
 *             .type(&#34;delay&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * DTS Job Monitor Rule can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:dts/jobMonitorRule:JobMonitorRule example &lt;dts_job_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dts/jobMonitorRule:JobMonitorRule")
public class JobMonitorRule extends com.pulumi.resources.CustomResource {
    /**
     * Trigger delay alarm threshold, which is measured in seconds.
     * 
     */
    @Export(name="delayRuleTime", type=String.class, parameters={})
    private Output<String> delayRuleTime;

    /**
     * @return Trigger delay alarm threshold, which is measured in seconds.
     * 
     */
    public Output<String> delayRuleTime() {
        return this.delayRuleTime;
    }
    /**
     * Migration, synchronization or subscription task ID can be by calling the [DescribeDtsJobs] get.
     * 
     */
    @Export(name="dtsJobId", type=String.class, parameters={})
    private Output<String> dtsJobId;

    /**
     * @return Migration, synchronization or subscription task ID can be by calling the [DescribeDtsJobs] get.
     * 
     */
    public Output<String> dtsJobId() {
        return this.dtsJobId;
    }
    /**
     * The alarm is triggered after notification of the contact phone number, A plurality of phone numbers between them with a comma (,) to separate.
     * 
     */
    @Export(name="phone", type=String.class, parameters={})
    private Output</* @Nullable */ String> phone;

    /**
     * @return The alarm is triggered after notification of the contact phone number, A plurality of phone numbers between them with a comma (,) to separate.
     * 
     */
    public Output<Optional<String>> phone() {
        return Codegen.optional(this.phone);
    }
    /**
     * Whether to enable monitoring rules, valid values: `Y`, `N`.
     * 
     */
    @Export(name="state", type=String.class, parameters={})
    private Output<String> state;

    /**
     * @return Whether to enable monitoring rules, valid values: `Y`, `N`.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Monitoring rules of type, valid values: `delay`, `error`. **delay**: delay alarm. **error**: abnormal alarm.
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output<String> type;

    /**
     * @return Monitoring rules of type, valid values: `delay`, `error`. **delay**: delay alarm. **error**: abnormal alarm.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public JobMonitorRule(String name) {
        this(name, JobMonitorRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public JobMonitorRule(String name, JobMonitorRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public JobMonitorRule(String name, JobMonitorRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dts/jobMonitorRule:JobMonitorRule", name, args == null ? JobMonitorRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private JobMonitorRule(String name, Output<String> id, @Nullable JobMonitorRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dts/jobMonitorRule:JobMonitorRule", name, state, makeResourceOptions(options, id));
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
    public static JobMonitorRule get(String name, Output<String> id, @Nullable JobMonitorRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new JobMonitorRule(name, id, state, options);
    }
}
