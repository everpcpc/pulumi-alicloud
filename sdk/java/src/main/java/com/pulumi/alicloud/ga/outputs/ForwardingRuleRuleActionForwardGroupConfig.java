// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.outputs;

import com.pulumi.alicloud.ga.outputs.ForwardingRuleRuleActionForwardGroupConfigServerGroupTuple;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;

@CustomType
public final class ForwardingRuleRuleActionForwardGroupConfig {
    /**
     * @return Terminal node group configuration.
     * 
     */
    private List<ForwardingRuleRuleActionForwardGroupConfigServerGroupTuple> serverGroupTuples;

    private ForwardingRuleRuleActionForwardGroupConfig() {}
    /**
     * @return Terminal node group configuration.
     * 
     */
    public List<ForwardingRuleRuleActionForwardGroupConfigServerGroupTuple> serverGroupTuples() {
        return this.serverGroupTuples;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ForwardingRuleRuleActionForwardGroupConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<ForwardingRuleRuleActionForwardGroupConfigServerGroupTuple> serverGroupTuples;
        public Builder() {}
        public Builder(ForwardingRuleRuleActionForwardGroupConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.serverGroupTuples = defaults.serverGroupTuples;
        }

        @CustomType.Setter
        public Builder serverGroupTuples(List<ForwardingRuleRuleActionForwardGroupConfigServerGroupTuple> serverGroupTuples) {
            this.serverGroupTuples = Objects.requireNonNull(serverGroupTuples);
            return this;
        }
        public Builder serverGroupTuples(ForwardingRuleRuleActionForwardGroupConfigServerGroupTuple... serverGroupTuples) {
            return serverGroupTuples(List.of(serverGroupTuples));
        }
        public ForwardingRuleRuleActionForwardGroupConfig build() {
            final var o = new ForwardingRuleRuleActionForwardGroupConfig();
            o.serverGroupTuples = serverGroupTuples;
            return o;
        }
    }
}
