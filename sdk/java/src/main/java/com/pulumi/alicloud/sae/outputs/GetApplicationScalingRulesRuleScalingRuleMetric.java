// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae.outputs;

import com.pulumi.alicloud.sae.outputs.GetApplicationScalingRulesRuleScalingRuleMetricMetric;
import com.pulumi.alicloud.sae.outputs.GetApplicationScalingRulesRuleScalingRuleMetricMetricsStatus;
import com.pulumi.alicloud.sae.outputs.GetApplicationScalingRulesRuleScalingRuleMetricScaleDownRule;
import com.pulumi.alicloud.sae.outputs.GetApplicationScalingRulesRuleScalingRuleMetricScaleUpRule;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetApplicationScalingRulesRuleScalingRuleMetric {
    /**
     * @return The maximum number of instances.
     * 
     */
    private Integer maxReplicas;
    /**
     * @return The auto scaling list of monitoring indicators.
     * 
     */
    private List<GetApplicationScalingRulesRuleScalingRuleMetricMetric> metrics;
    /**
     * @return Monitor indicator elasticity status.
     * 
     */
    private List<GetApplicationScalingRulesRuleScalingRuleMetricMetricsStatus> metricsStatuses;
    /**
     * @return The minimum number of instances.
     * 
     */
    private Integer minReplicas;
    /**
     * @return The shrink rule.
     * 
     */
    private List<GetApplicationScalingRulesRuleScalingRuleMetricScaleDownRule> scaleDownRules;
    /**
     * @return The expansion rules.
     * 
     */
    private List<GetApplicationScalingRulesRuleScalingRuleMetricScaleUpRule> scaleUpRules;

    private GetApplicationScalingRulesRuleScalingRuleMetric() {}
    /**
     * @return The maximum number of instances.
     * 
     */
    public Integer maxReplicas() {
        return this.maxReplicas;
    }
    /**
     * @return The auto scaling list of monitoring indicators.
     * 
     */
    public List<GetApplicationScalingRulesRuleScalingRuleMetricMetric> metrics() {
        return this.metrics;
    }
    /**
     * @return Monitor indicator elasticity status.
     * 
     */
    public List<GetApplicationScalingRulesRuleScalingRuleMetricMetricsStatus> metricsStatuses() {
        return this.metricsStatuses;
    }
    /**
     * @return The minimum number of instances.
     * 
     */
    public Integer minReplicas() {
        return this.minReplicas;
    }
    /**
     * @return The shrink rule.
     * 
     */
    public List<GetApplicationScalingRulesRuleScalingRuleMetricScaleDownRule> scaleDownRules() {
        return this.scaleDownRules;
    }
    /**
     * @return The expansion rules.
     * 
     */
    public List<GetApplicationScalingRulesRuleScalingRuleMetricScaleUpRule> scaleUpRules() {
        return this.scaleUpRules;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetApplicationScalingRulesRuleScalingRuleMetric defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer maxReplicas;
        private List<GetApplicationScalingRulesRuleScalingRuleMetricMetric> metrics;
        private List<GetApplicationScalingRulesRuleScalingRuleMetricMetricsStatus> metricsStatuses;
        private Integer minReplicas;
        private List<GetApplicationScalingRulesRuleScalingRuleMetricScaleDownRule> scaleDownRules;
        private List<GetApplicationScalingRulesRuleScalingRuleMetricScaleUpRule> scaleUpRules;
        public Builder() {}
        public Builder(GetApplicationScalingRulesRuleScalingRuleMetric defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.maxReplicas = defaults.maxReplicas;
    	      this.metrics = defaults.metrics;
    	      this.metricsStatuses = defaults.metricsStatuses;
    	      this.minReplicas = defaults.minReplicas;
    	      this.scaleDownRules = defaults.scaleDownRules;
    	      this.scaleUpRules = defaults.scaleUpRules;
        }

        @CustomType.Setter
        public Builder maxReplicas(Integer maxReplicas) {
            this.maxReplicas = Objects.requireNonNull(maxReplicas);
            return this;
        }
        @CustomType.Setter
        public Builder metrics(List<GetApplicationScalingRulesRuleScalingRuleMetricMetric> metrics) {
            this.metrics = Objects.requireNonNull(metrics);
            return this;
        }
        public Builder metrics(GetApplicationScalingRulesRuleScalingRuleMetricMetric... metrics) {
            return metrics(List.of(metrics));
        }
        @CustomType.Setter
        public Builder metricsStatuses(List<GetApplicationScalingRulesRuleScalingRuleMetricMetricsStatus> metricsStatuses) {
            this.metricsStatuses = Objects.requireNonNull(metricsStatuses);
            return this;
        }
        public Builder metricsStatuses(GetApplicationScalingRulesRuleScalingRuleMetricMetricsStatus... metricsStatuses) {
            return metricsStatuses(List.of(metricsStatuses));
        }
        @CustomType.Setter
        public Builder minReplicas(Integer minReplicas) {
            this.minReplicas = Objects.requireNonNull(minReplicas);
            return this;
        }
        @CustomType.Setter
        public Builder scaleDownRules(List<GetApplicationScalingRulesRuleScalingRuleMetricScaleDownRule> scaleDownRules) {
            this.scaleDownRules = Objects.requireNonNull(scaleDownRules);
            return this;
        }
        public Builder scaleDownRules(GetApplicationScalingRulesRuleScalingRuleMetricScaleDownRule... scaleDownRules) {
            return scaleDownRules(List.of(scaleDownRules));
        }
        @CustomType.Setter
        public Builder scaleUpRules(List<GetApplicationScalingRulesRuleScalingRuleMetricScaleUpRule> scaleUpRules) {
            this.scaleUpRules = Objects.requireNonNull(scaleUpRules);
            return this;
        }
        public Builder scaleUpRules(GetApplicationScalingRulesRuleScalingRuleMetricScaleUpRule... scaleUpRules) {
            return scaleUpRules(List.of(scaleUpRules));
        }
        public GetApplicationScalingRulesRuleScalingRuleMetric build() {
            final var o = new GetApplicationScalingRulesRuleScalingRuleMetric();
            o.maxReplicas = maxReplicas;
            o.metrics = metrics;
            o.metricsStatuses = metricsStatuses;
            o.minReplicas = minReplicas;
            o.scaleDownRules = scaleDownRules;
            o.scaleUpRules = scaleUpRules;
            return o;
        }
    }
}
