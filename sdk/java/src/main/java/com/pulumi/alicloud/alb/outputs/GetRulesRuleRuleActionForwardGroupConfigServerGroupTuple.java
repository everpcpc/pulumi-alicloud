// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRulesRuleRuleActionForwardGroupConfigServerGroupTuple {
    /**
     * @return The ID of the destination server group to which requests are forwarded.
     * 
     */
    private String serverGroupId;
    /**
     * @return The Weight of server group.
     * 
     */
    private Integer weight;

    private GetRulesRuleRuleActionForwardGroupConfigServerGroupTuple() {}
    /**
     * @return The ID of the destination server group to which requests are forwarded.
     * 
     */
    public String serverGroupId() {
        return this.serverGroupId;
    }
    /**
     * @return The Weight of server group.
     * 
     */
    public Integer weight() {
        return this.weight;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRulesRuleRuleActionForwardGroupConfigServerGroupTuple defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String serverGroupId;
        private Integer weight;
        public Builder() {}
        public Builder(GetRulesRuleRuleActionForwardGroupConfigServerGroupTuple defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.serverGroupId = defaults.serverGroupId;
    	      this.weight = defaults.weight;
        }

        @CustomType.Setter
        public Builder serverGroupId(String serverGroupId) {
            this.serverGroupId = Objects.requireNonNull(serverGroupId);
            return this;
        }
        @CustomType.Setter
        public Builder weight(Integer weight) {
            this.weight = Objects.requireNonNull(weight);
            return this;
        }
        public GetRulesRuleRuleActionForwardGroupConfigServerGroupTuple build() {
            final var o = new GetRulesRuleRuleActionForwardGroupConfigServerGroupTuple();
            o.serverGroupId = serverGroupId;
            o.weight = weight;
            return o;
        }
    }
}