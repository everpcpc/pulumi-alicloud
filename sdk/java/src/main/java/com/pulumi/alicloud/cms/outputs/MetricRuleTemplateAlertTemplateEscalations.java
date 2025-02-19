// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms.outputs;

import com.pulumi.alicloud.cms.outputs.MetricRuleTemplateAlertTemplateEscalationsCritical;
import com.pulumi.alicloud.cms.outputs.MetricRuleTemplateAlertTemplateEscalationsInfo;
import com.pulumi.alicloud.cms.outputs.MetricRuleTemplateAlertTemplateEscalationsWarn;
import com.pulumi.core.annotations.CustomType;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class MetricRuleTemplateAlertTemplateEscalations {
    /**
     * @return The condition for triggering critical-level alerts. See the following `Block critical`.
     * 
     */
    private @Nullable MetricRuleTemplateAlertTemplateEscalationsCritical critical;
    /**
     * @return The condition for triggering info-level alerts. See the following `Block info`.
     * 
     */
    private @Nullable MetricRuleTemplateAlertTemplateEscalationsInfo info;
    /**
     * @return The condition for triggering warn-level alerts. See the following `Block warn`.
     * 
     */
    private @Nullable MetricRuleTemplateAlertTemplateEscalationsWarn warn;

    private MetricRuleTemplateAlertTemplateEscalations() {}
    /**
     * @return The condition for triggering critical-level alerts. See the following `Block critical`.
     * 
     */
    public Optional<MetricRuleTemplateAlertTemplateEscalationsCritical> critical() {
        return Optional.ofNullable(this.critical);
    }
    /**
     * @return The condition for triggering info-level alerts. See the following `Block info`.
     * 
     */
    public Optional<MetricRuleTemplateAlertTemplateEscalationsInfo> info() {
        return Optional.ofNullable(this.info);
    }
    /**
     * @return The condition for triggering warn-level alerts. See the following `Block warn`.
     * 
     */
    public Optional<MetricRuleTemplateAlertTemplateEscalationsWarn> warn() {
        return Optional.ofNullable(this.warn);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(MetricRuleTemplateAlertTemplateEscalations defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable MetricRuleTemplateAlertTemplateEscalationsCritical critical;
        private @Nullable MetricRuleTemplateAlertTemplateEscalationsInfo info;
        private @Nullable MetricRuleTemplateAlertTemplateEscalationsWarn warn;
        public Builder() {}
        public Builder(MetricRuleTemplateAlertTemplateEscalations defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.critical = defaults.critical;
    	      this.info = defaults.info;
    	      this.warn = defaults.warn;
        }

        @CustomType.Setter
        public Builder critical(@Nullable MetricRuleTemplateAlertTemplateEscalationsCritical critical) {
            this.critical = critical;
            return this;
        }
        @CustomType.Setter
        public Builder info(@Nullable MetricRuleTemplateAlertTemplateEscalationsInfo info) {
            this.info = info;
            return this;
        }
        @CustomType.Setter
        public Builder warn(@Nullable MetricRuleTemplateAlertTemplateEscalationsWarn warn) {
            this.warn = warn;
            return this;
        }
        public MetricRuleTemplateAlertTemplateEscalations build() {
            final var o = new MetricRuleTemplateAlertTemplateEscalations();
            o.critical = critical;
            o.info = info;
            o.warn = warn;
            return o;
        }
    }
}
