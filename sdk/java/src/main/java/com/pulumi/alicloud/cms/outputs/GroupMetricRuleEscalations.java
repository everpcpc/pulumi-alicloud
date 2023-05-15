// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms.outputs;

import com.pulumi.alicloud.cms.outputs.GroupMetricRuleEscalationsCritical;
import com.pulumi.alicloud.cms.outputs.GroupMetricRuleEscalationsInfo;
import com.pulumi.alicloud.cms.outputs.GroupMetricRuleEscalationsWarn;
import com.pulumi.core.annotations.CustomType;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GroupMetricRuleEscalations {
    /**
     * @return The critical level.
     * 
     */
    private @Nullable GroupMetricRuleEscalationsCritical critical;
    /**
     * @return The info level.
     * 
     */
    private @Nullable GroupMetricRuleEscalationsInfo info;
    /**
     * @return The warn level.
     * 
     */
    private @Nullable GroupMetricRuleEscalationsWarn warn;

    private GroupMetricRuleEscalations() {}
    /**
     * @return The critical level.
     * 
     */
    public Optional<GroupMetricRuleEscalationsCritical> critical() {
        return Optional.ofNullable(this.critical);
    }
    /**
     * @return The info level.
     * 
     */
    public Optional<GroupMetricRuleEscalationsInfo> info() {
        return Optional.ofNullable(this.info);
    }
    /**
     * @return The warn level.
     * 
     */
    public Optional<GroupMetricRuleEscalationsWarn> warn() {
        return Optional.ofNullable(this.warn);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GroupMetricRuleEscalations defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable GroupMetricRuleEscalationsCritical critical;
        private @Nullable GroupMetricRuleEscalationsInfo info;
        private @Nullable GroupMetricRuleEscalationsWarn warn;
        public Builder() {}
        public Builder(GroupMetricRuleEscalations defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.critical = defaults.critical;
    	      this.info = defaults.info;
    	      this.warn = defaults.warn;
        }

        @CustomType.Setter
        public Builder critical(@Nullable GroupMetricRuleEscalationsCritical critical) {
            this.critical = critical;
            return this;
        }
        @CustomType.Setter
        public Builder info(@Nullable GroupMetricRuleEscalationsInfo info) {
            this.info = info;
            return this;
        }
        @CustomType.Setter
        public Builder warn(@Nullable GroupMetricRuleEscalationsWarn warn) {
            this.warn = warn;
            return this;
        }
        public GroupMetricRuleEscalations build() {
            final var o = new GroupMetricRuleEscalations();
            o.critical = critical;
            o.info = info;
            o.warn = warn;
            return o;
        }
    }
}
