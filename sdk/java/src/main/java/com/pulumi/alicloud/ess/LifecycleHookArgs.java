// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LifecycleHookArgs extends com.pulumi.resources.ResourceArgs {

    public static final LifecycleHookArgs Empty = new LifecycleHookArgs();

    /**
     * Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses. Applicable value: CONTINUE, ABANDON, ROLLBACK, default value: CONTINUE.
     * 
     */
    @Import(name="defaultResult")
    private @Nullable Output<String> defaultResult;

    /**
     * @return Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses. Applicable value: CONTINUE, ABANDON, ROLLBACK, default value: CONTINUE.
     * 
     */
    public Optional<Output<String>> defaultResult() {
        return Optional.ofNullable(this.defaultResult);
    }

    /**
     * Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter. Default value: 600.
     * 
     */
    @Import(name="heartbeatTimeout")
    private @Nullable Output<Integer> heartbeatTimeout;

    /**
     * @return Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter. Default value: 600.
     * 
     */
    public Optional<Output<Integer>> heartbeatTimeout() {
        return Optional.ofNullable(this.heartbeatTimeout);
    }

    /**
     * Type of Scaling activity attached to lifecycle hook. Supported value: SCALE_OUT, SCALE_IN.
     * 
     */
    @Import(name="lifecycleTransition", required=true)
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
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the lifecycle hook, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is lifecycle hook id.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The Arn of notification target.
     * 
     */
    @Import(name="notificationArn")
    private @Nullable Output<String> notificationArn;

    /**
     * @return The Arn of notification target.
     * 
     */
    public Optional<Output<String>> notificationArn() {
        return Optional.ofNullable(this.notificationArn);
    }

    /**
     * Additional information that you want to include when Auto Scaling sends a message to the notification target.
     * 
     */
    @Import(name="notificationMetadata")
    private @Nullable Output<String> notificationMetadata;

    /**
     * @return Additional information that you want to include when Auto Scaling sends a message to the notification target.
     * 
     */
    public Optional<Output<String>> notificationMetadata() {
        return Optional.ofNullable(this.notificationMetadata);
    }

    /**
     * The ID of the Auto Scaling group to which you want to assign the lifecycle hook.
     * 
     */
    @Import(name="scalingGroupId", required=true)
    private Output<String> scalingGroupId;

    /**
     * @return The ID of the Auto Scaling group to which you want to assign the lifecycle hook.
     * 
     */
    public Output<String> scalingGroupId() {
        return this.scalingGroupId;
    }

    private LifecycleHookArgs() {}

    private LifecycleHookArgs(LifecycleHookArgs $) {
        this.defaultResult = $.defaultResult;
        this.heartbeatTimeout = $.heartbeatTimeout;
        this.lifecycleTransition = $.lifecycleTransition;
        this.name = $.name;
        this.notificationArn = $.notificationArn;
        this.notificationMetadata = $.notificationMetadata;
        this.scalingGroupId = $.scalingGroupId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LifecycleHookArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LifecycleHookArgs $;

        public Builder() {
            $ = new LifecycleHookArgs();
        }

        public Builder(LifecycleHookArgs defaults) {
            $ = new LifecycleHookArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param defaultResult Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses. Applicable value: CONTINUE, ABANDON, ROLLBACK, default value: CONTINUE.
         * 
         * @return builder
         * 
         */
        public Builder defaultResult(@Nullable Output<String> defaultResult) {
            $.defaultResult = defaultResult;
            return this;
        }

        /**
         * @param defaultResult Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses. Applicable value: CONTINUE, ABANDON, ROLLBACK, default value: CONTINUE.
         * 
         * @return builder
         * 
         */
        public Builder defaultResult(String defaultResult) {
            return defaultResult(Output.of(defaultResult));
        }

        /**
         * @param heartbeatTimeout Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter. Default value: 600.
         * 
         * @return builder
         * 
         */
        public Builder heartbeatTimeout(@Nullable Output<Integer> heartbeatTimeout) {
            $.heartbeatTimeout = heartbeatTimeout;
            return this;
        }

        /**
         * @param heartbeatTimeout Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter. Default value: 600.
         * 
         * @return builder
         * 
         */
        public Builder heartbeatTimeout(Integer heartbeatTimeout) {
            return heartbeatTimeout(Output.of(heartbeatTimeout));
        }

        /**
         * @param lifecycleTransition Type of Scaling activity attached to lifecycle hook. Supported value: SCALE_OUT, SCALE_IN.
         * 
         * @return builder
         * 
         */
        public Builder lifecycleTransition(Output<String> lifecycleTransition) {
            $.lifecycleTransition = lifecycleTransition;
            return this;
        }

        /**
         * @param lifecycleTransition Type of Scaling activity attached to lifecycle hook. Supported value: SCALE_OUT, SCALE_IN.
         * 
         * @return builder
         * 
         */
        public Builder lifecycleTransition(String lifecycleTransition) {
            return lifecycleTransition(Output.of(lifecycleTransition));
        }

        /**
         * @param name The name of the lifecycle hook, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is lifecycle hook id.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the lifecycle hook, which must contain 2-64 characters (English or Chinese), starting with numbers, English letters or Chinese characters, and can contain number, underscores `_`, hypens `-`, and decimal point `.`. If this parameter value is not specified, the default value is lifecycle hook id.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param notificationArn The Arn of notification target.
         * 
         * @return builder
         * 
         */
        public Builder notificationArn(@Nullable Output<String> notificationArn) {
            $.notificationArn = notificationArn;
            return this;
        }

        /**
         * @param notificationArn The Arn of notification target.
         * 
         * @return builder
         * 
         */
        public Builder notificationArn(String notificationArn) {
            return notificationArn(Output.of(notificationArn));
        }

        /**
         * @param notificationMetadata Additional information that you want to include when Auto Scaling sends a message to the notification target.
         * 
         * @return builder
         * 
         */
        public Builder notificationMetadata(@Nullable Output<String> notificationMetadata) {
            $.notificationMetadata = notificationMetadata;
            return this;
        }

        /**
         * @param notificationMetadata Additional information that you want to include when Auto Scaling sends a message to the notification target.
         * 
         * @return builder
         * 
         */
        public Builder notificationMetadata(String notificationMetadata) {
            return notificationMetadata(Output.of(notificationMetadata));
        }

        /**
         * @param scalingGroupId The ID of the Auto Scaling group to which you want to assign the lifecycle hook.
         * 
         * @return builder
         * 
         */
        public Builder scalingGroupId(Output<String> scalingGroupId) {
            $.scalingGroupId = scalingGroupId;
            return this;
        }

        /**
         * @param scalingGroupId The ID of the Auto Scaling group to which you want to assign the lifecycle hook.
         * 
         * @return builder
         * 
         */
        public Builder scalingGroupId(String scalingGroupId) {
            return scalingGroupId(Output.of(scalingGroupId));
        }

        public LifecycleHookArgs build() {
            $.lifecycleTransition = Objects.requireNonNull($.lifecycleTransition, "expected parameter 'lifecycleTransition' to be non-null");
            $.scalingGroupId = Objects.requireNonNull($.scalingGroupId, "expected parameter 'scalingGroupId' to be non-null");
            return $;
        }
    }

}
