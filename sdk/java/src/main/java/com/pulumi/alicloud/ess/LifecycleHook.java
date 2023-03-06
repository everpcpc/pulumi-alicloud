// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ess.LifecycleHookArgs;
import com.pulumi.alicloud.ess.inputs.LifecycleHookState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * Ess lifecycle hook can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ess/lifecycleHook:LifecycleHook example ash-l12345
 * ```
 * 
 */
@ResourceType(type="alicloud:ess/lifecycleHook:LifecycleHook")
public class LifecycleHook extends com.pulumi.resources.CustomResource {
    /**
     * Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses. Applicable value: CONTINUE, ABANDON, ROLLBACK, default value: CONTINUE.
     * 
     */
    @Export(name="defaultResult", type=String.class, parameters={})
    private Output</* @Nullable */ String> defaultResult;

    /**
     * @return Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses. Applicable value: CONTINUE, ABANDON, ROLLBACK, default value: CONTINUE.
     * 
     */
    public Output<Optional<String>> defaultResult() {
        return Codegen.optional(this.defaultResult);
    }
    /**
     * Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter. Default value: 600.
     * 
     */
    @Export(name="heartbeatTimeout", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> heartbeatTimeout;

    /**
     * @return Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter. Default value: 600.
     * 
     */
    public Output<Optional<Integer>> heartbeatTimeout() {
        return Codegen.optional(this.heartbeatTimeout);
    }
    /**
     * Type of Scaling activity attached to lifecycle hook. Supported value: SCALE_OUT, SCALE_IN.
     * 
     */
    @Export(name="lifecycleTransition", type=String.class, parameters={})
    private Output<String> lifecycleTransition;

    /**
     * @return Type of Scaling activity attached to lifecycle hook. Supported value: SCALE_OUT, SCALE_IN.
     * 
     */
    public Output<String> lifecycleTransition() {
        return this.lifecycleTransition;
    }
    /**
     * The name of the lifecycle hook, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is lifecycle hook id.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the lifecycle hook, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is lifecycle hook id.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The Arn of notification target.
     * 
     */
    @Export(name="notificationArn", type=String.class, parameters={})
    private Output<String> notificationArn;

    /**
     * @return The Arn of notification target.
     * 
     */
    public Output<String> notificationArn() {
        return this.notificationArn;
    }
    /**
     * Additional information that you want to include when Auto Scaling sends a message to the notification target.
     * 
     */
    @Export(name="notificationMetadata", type=String.class, parameters={})
    private Output<String> notificationMetadata;

    /**
     * @return Additional information that you want to include when Auto Scaling sends a message to the notification target.
     * 
     */
    public Output<String> notificationMetadata() {
        return this.notificationMetadata;
    }
    /**
     * The ID of the Auto Scaling group to which you want to assign the lifecycle hook.
     * 
     */
    @Export(name="scalingGroupId", type=String.class, parameters={})
    private Output<String> scalingGroupId;

    /**
     * @return The ID of the Auto Scaling group to which you want to assign the lifecycle hook.
     * 
     */
    public Output<String> scalingGroupId() {
        return this.scalingGroupId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LifecycleHook(String name) {
        this(name, LifecycleHookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LifecycleHook(String name, LifecycleHookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LifecycleHook(String name, LifecycleHookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ess/lifecycleHook:LifecycleHook", name, args == null ? LifecycleHookArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LifecycleHook(String name, Output<String> id, @Nullable LifecycleHookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ess/lifecycleHook:LifecycleHook", name, state, makeResourceOptions(options, id));
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
    public static LifecycleHook get(String name, Output<String> id, @Nullable LifecycleHookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LifecycleHook(name, id, state, options);
    }
}