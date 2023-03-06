// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.resourcemanager;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.resourcemanager.SharedTargetArgs;
import com.pulumi.alicloud.resourcemanager.inputs.SharedTargetState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Resource Manager Shared Target resource.
 * 
 * For information about Resource Manager Shared Target and how to use it, see [What is Shared Target](https://www.alibabacloud.com/help/en/doc-detail/94475.htm).
 * 
 * &gt; **NOTE:** Available in v1.111.0+.
 * 
 * ## Import
 * 
 * Resource Manager Shared Target can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:resourcemanager/sharedTarget:SharedTarget example &lt;resource_share_id&gt;:&lt;target_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:resourcemanager/sharedTarget:SharedTarget")
public class SharedTarget extends com.pulumi.resources.CustomResource {
    /**
     * The resource share ID of resource manager.
     * 
     */
    @Export(name="resourceShareId", type=String.class, parameters={})
    private Output<String> resourceShareId;

    /**
     * @return The resource share ID of resource manager.
     * 
     */
    public Output<String> resourceShareId() {
        return this.resourceShareId;
    }
    /**
     * The status of shared target.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of shared target.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The member account ID in resource directory.
     * 
     */
    @Export(name="targetId", type=String.class, parameters={})
    private Output<String> targetId;

    /**
     * @return The member account ID in resource directory.
     * 
     */
    public Output<String> targetId() {
        return this.targetId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SharedTarget(String name) {
        this(name, SharedTargetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SharedTarget(String name, SharedTargetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SharedTarget(String name, SharedTargetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:resourcemanager/sharedTarget:SharedTarget", name, args == null ? SharedTargetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SharedTarget(String name, Output<String> id, @Nullable SharedTargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:resourcemanager/sharedTarget:SharedTarget", name, state, makeResourceOptions(options, id));
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
    public static SharedTarget get(String name, Output<String> id, @Nullable SharedTargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SharedTarget(name, id, state, options);
    }
}