// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RuleRuleActionForwardGroupConfigServerGroupTuple {
    /**
     * @return The ID of the destination server group to which requests are forwarded.
     * 
     */
    private @Nullable String serverGroupId;
    /**
     * @return The Weight of server group. Default value: `100`. **NOTE:** This attribute is required when the number of `server_group_tuples` is greater than 2.
     * 
     */
    private @Nullable Integer weight;

    private RuleRuleActionForwardGroupConfigServerGroupTuple() {}
    /**
     * @return The ID of the destination server group to which requests are forwarded.
     * 
     */
    public Optional<String> serverGroupId() {
        return Optional.ofNullable(this.serverGroupId);
    }
    /**
     * @return The Weight of server group. Default value: `100`. **NOTE:** This attribute is required when the number of `server_group_tuples` is greater than 2.
     * 
     */
    public Optional<Integer> weight() {
        return Optional.ofNullable(this.weight);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RuleRuleActionForwardGroupConfigServerGroupTuple defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String serverGroupId;
        private @Nullable Integer weight;
        public Builder() {}
        public Builder(RuleRuleActionForwardGroupConfigServerGroupTuple defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.serverGroupId = defaults.serverGroupId;
    	      this.weight = defaults.weight;
        }

        @CustomType.Setter
        public Builder serverGroupId(@Nullable String serverGroupId) {
            this.serverGroupId = serverGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder weight(@Nullable Integer weight) {
            this.weight = weight;
            return this;
        }
        public RuleRuleActionForwardGroupConfigServerGroupTuple build() {
            final var o = new RuleRuleActionForwardGroupConfigServerGroupTuple();
            o.serverGroupId = serverGroupId;
            o.weight = weight;
            return o;
        }
    }
}
