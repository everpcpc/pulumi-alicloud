// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.outputs;

import com.pulumi.alicloud.alb.outputs.RuleRuleActionTrafficMirrorConfigMirrorGroupConfigServerGroupTuple;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class RuleRuleActionTrafficMirrorConfigMirrorGroupConfig {
    /**
     * @return The destination server group to which requests are forwarded.
     * 
     */
    private @Nullable List<RuleRuleActionTrafficMirrorConfigMirrorGroupConfigServerGroupTuple> serverGroupTuples;

    private RuleRuleActionTrafficMirrorConfigMirrorGroupConfig() {}
    /**
     * @return The destination server group to which requests are forwarded.
     * 
     */
    public List<RuleRuleActionTrafficMirrorConfigMirrorGroupConfigServerGroupTuple> serverGroupTuples() {
        return this.serverGroupTuples == null ? List.of() : this.serverGroupTuples;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RuleRuleActionTrafficMirrorConfigMirrorGroupConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<RuleRuleActionTrafficMirrorConfigMirrorGroupConfigServerGroupTuple> serverGroupTuples;
        public Builder() {}
        public Builder(RuleRuleActionTrafficMirrorConfigMirrorGroupConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.serverGroupTuples = defaults.serverGroupTuples;
        }

        @CustomType.Setter
        public Builder serverGroupTuples(@Nullable List<RuleRuleActionTrafficMirrorConfigMirrorGroupConfigServerGroupTuple> serverGroupTuples) {
            this.serverGroupTuples = serverGroupTuples;
            return this;
        }
        public Builder serverGroupTuples(RuleRuleActionTrafficMirrorConfigMirrorGroupConfigServerGroupTuple... serverGroupTuples) {
            return serverGroupTuples(List.of(serverGroupTuples));
        }
        public RuleRuleActionTrafficMirrorConfigMirrorGroupConfig build() {
            final var o = new RuleRuleActionTrafficMirrorConfigMirrorGroupConfig();
            o.serverGroupTuples = serverGroupTuples;
            return o;
        }
    }
}