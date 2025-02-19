// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.inputs;

import com.pulumi.alicloud.alb.inputs.ListenerDefaultActionForwardGroupConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ListenerDefaultActionArgs extends com.pulumi.resources.ResourceArgs {

    public static final ListenerDefaultActionArgs Empty = new ListenerDefaultActionArgs();

    /**
     * The configurations of the actions. This parameter is required if Type is set to FowardGroup.
     * 
     */
    @Import(name="forwardGroupConfig", required=true)
    private Output<ListenerDefaultActionForwardGroupConfigArgs> forwardGroupConfig;

    /**
     * @return The configurations of the actions. This parameter is required if Type is set to FowardGroup.
     * 
     */
    public Output<ListenerDefaultActionForwardGroupConfigArgs> forwardGroupConfig() {
        return this.forwardGroupConfig;
    }

    /**
     * Action Type.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Action Type.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private ListenerDefaultActionArgs() {}

    private ListenerDefaultActionArgs(ListenerDefaultActionArgs $) {
        this.forwardGroupConfig = $.forwardGroupConfig;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ListenerDefaultActionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ListenerDefaultActionArgs $;

        public Builder() {
            $ = new ListenerDefaultActionArgs();
        }

        public Builder(ListenerDefaultActionArgs defaults) {
            $ = new ListenerDefaultActionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param forwardGroupConfig The configurations of the actions. This parameter is required if Type is set to FowardGroup.
         * 
         * @return builder
         * 
         */
        public Builder forwardGroupConfig(Output<ListenerDefaultActionForwardGroupConfigArgs> forwardGroupConfig) {
            $.forwardGroupConfig = forwardGroupConfig;
            return this;
        }

        /**
         * @param forwardGroupConfig The configurations of the actions. This parameter is required if Type is set to FowardGroup.
         * 
         * @return builder
         * 
         */
        public Builder forwardGroupConfig(ListenerDefaultActionForwardGroupConfigArgs forwardGroupConfig) {
            return forwardGroupConfig(Output.of(forwardGroupConfig));
        }

        /**
         * @param type Action Type.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Action Type.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ListenerDefaultActionArgs build() {
            $.forwardGroupConfig = Objects.requireNonNull($.forwardGroupConfig, "expected parameter 'forwardGroupConfig' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
