// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.inputs;

import com.pulumi.alicloud.alb.inputs.ListenerDefaultActionForwardGroupConfigServerGroupTupleArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.List;
import java.util.Objects;


public final class ListenerDefaultActionForwardGroupConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final ListenerDefaultActionForwardGroupConfigArgs Empty = new ListenerDefaultActionForwardGroupConfigArgs();

    /**
     * The destination server group to which requests are forwarded.
     * 
     */
    @Import(name="serverGroupTuples", required=true)
    private Output<List<ListenerDefaultActionForwardGroupConfigServerGroupTupleArgs>> serverGroupTuples;

    /**
     * @return The destination server group to which requests are forwarded.
     * 
     */
    public Output<List<ListenerDefaultActionForwardGroupConfigServerGroupTupleArgs>> serverGroupTuples() {
        return this.serverGroupTuples;
    }

    private ListenerDefaultActionForwardGroupConfigArgs() {}

    private ListenerDefaultActionForwardGroupConfigArgs(ListenerDefaultActionForwardGroupConfigArgs $) {
        this.serverGroupTuples = $.serverGroupTuples;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ListenerDefaultActionForwardGroupConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ListenerDefaultActionForwardGroupConfigArgs $;

        public Builder() {
            $ = new ListenerDefaultActionForwardGroupConfigArgs();
        }

        public Builder(ListenerDefaultActionForwardGroupConfigArgs defaults) {
            $ = new ListenerDefaultActionForwardGroupConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serverGroupTuples The destination server group to which requests are forwarded.
         * 
         * @return builder
         * 
         */
        public Builder serverGroupTuples(Output<List<ListenerDefaultActionForwardGroupConfigServerGroupTupleArgs>> serverGroupTuples) {
            $.serverGroupTuples = serverGroupTuples;
            return this;
        }

        /**
         * @param serverGroupTuples The destination server group to which requests are forwarded.
         * 
         * @return builder
         * 
         */
        public Builder serverGroupTuples(List<ListenerDefaultActionForwardGroupConfigServerGroupTupleArgs> serverGroupTuples) {
            return serverGroupTuples(Output.of(serverGroupTuples));
        }

        /**
         * @param serverGroupTuples The destination server group to which requests are forwarded.
         * 
         * @return builder
         * 
         */
        public Builder serverGroupTuples(ListenerDefaultActionForwardGroupConfigServerGroupTupleArgs... serverGroupTuples) {
            return serverGroupTuples(List.of(serverGroupTuples));
        }

        public ListenerDefaultActionForwardGroupConfigArgs build() {
            $.serverGroupTuples = Objects.requireNonNull($.serverGroupTuples, "expected parameter 'serverGroupTuples' to be non-null");
            return $;
        }
    }

}
