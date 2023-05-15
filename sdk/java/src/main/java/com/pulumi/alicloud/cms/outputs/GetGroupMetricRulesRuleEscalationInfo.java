// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetGroupMetricRulesRuleEscalationInfo {
    /**
     * @return The comparison operator of the threshold for warn-level alerts.
     * 
     */
    private String comparisonOperator;
    /**
     * @return The statistical aggregation method for warn-level alerts.
     * 
     */
    private String statistics;
    /**
     * @return The threshold for warn-level alerts.
     * 
     */
    private String threshold;
    /**
     * @return The consecutive number of times for which the metric value is measured before a warn-level alert is triggered.
     * 
     */
    private Integer times;

    private GetGroupMetricRulesRuleEscalationInfo() {}
    /**
     * @return The comparison operator of the threshold for warn-level alerts.
     * 
     */
    public String comparisonOperator() {
        return this.comparisonOperator;
    }
    /**
     * @return The statistical aggregation method for warn-level alerts.
     * 
     */
    public String statistics() {
        return this.statistics;
    }
    /**
     * @return The threshold for warn-level alerts.
     * 
     */
    public String threshold() {
        return this.threshold;
    }
    /**
     * @return The consecutive number of times for which the metric value is measured before a warn-level alert is triggered.
     * 
     */
    public Integer times() {
        return this.times;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGroupMetricRulesRuleEscalationInfo defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String comparisonOperator;
        private String statistics;
        private String threshold;
        private Integer times;
        public Builder() {}
        public Builder(GetGroupMetricRulesRuleEscalationInfo defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.comparisonOperator = defaults.comparisonOperator;
    	      this.statistics = defaults.statistics;
    	      this.threshold = defaults.threshold;
    	      this.times = defaults.times;
        }

        @CustomType.Setter
        public Builder comparisonOperator(String comparisonOperator) {
            this.comparisonOperator = Objects.requireNonNull(comparisonOperator);
            return this;
        }
        @CustomType.Setter
        public Builder statistics(String statistics) {
            this.statistics = Objects.requireNonNull(statistics);
            return this;
        }
        @CustomType.Setter
        public Builder threshold(String threshold) {
            this.threshold = Objects.requireNonNull(threshold);
            return this;
        }
        @CustomType.Setter
        public Builder times(Integer times) {
            this.times = Objects.requireNonNull(times);
            return this;
        }
        public GetGroupMetricRulesRuleEscalationInfo build() {
            final var o = new GetGroupMetricRulesRuleEscalationInfo();
            o.comparisonOperator = comparisonOperator;
            o.statistics = statistics;
            o.threshold = threshold;
            o.times = times;
            return o;
        }
    }
}
