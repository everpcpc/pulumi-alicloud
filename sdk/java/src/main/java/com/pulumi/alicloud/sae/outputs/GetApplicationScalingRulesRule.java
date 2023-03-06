// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae.outputs;

import com.pulumi.alicloud.sae.outputs.GetApplicationScalingRulesRuleScalingRuleMetric;
import com.pulumi.alicloud.sae.outputs.GetApplicationScalingRulesRuleScalingRuleTimer;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetApplicationScalingRulesRule {
    /**
     * @return The ID of the Application.
     * 
     */
    private String appId;
    /**
     * @return The CreateTime of the Application Scaling Rule.
     * 
     */
    private String createTime;
    /**
     * @return The ID of the Application Scaling Rule.
     * 
     */
    private String id;
    /**
     * @return Whether to enable the auto scaling policy.
     * 
     */
    private Boolean scalingRuleEnable;
    /**
     * @return Monitoring indicators for elastic scaling.
     * 
     */
    private List<GetApplicationScalingRulesRuleScalingRuleMetric> scalingRuleMetrics;
    /**
     * @return The name of the scaling rule.
     * 
     */
    private String scalingRuleName;
    /**
     * @return Timing elastic expansion.
     * 
     */
    private List<GetApplicationScalingRulesRuleScalingRuleTimer> scalingRuleTimers;
    /**
     * @return Flexible strategy type.
     * 
     */
    private String scalingRuleType;

    private GetApplicationScalingRulesRule() {}
    /**
     * @return The ID of the Application.
     * 
     */
    public String appId() {
        return this.appId;
    }
    /**
     * @return The CreateTime of the Application Scaling Rule.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The ID of the Application Scaling Rule.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Whether to enable the auto scaling policy.
     * 
     */
    public Boolean scalingRuleEnable() {
        return this.scalingRuleEnable;
    }
    /**
     * @return Monitoring indicators for elastic scaling.
     * 
     */
    public List<GetApplicationScalingRulesRuleScalingRuleMetric> scalingRuleMetrics() {
        return this.scalingRuleMetrics;
    }
    /**
     * @return The name of the scaling rule.
     * 
     */
    public String scalingRuleName() {
        return this.scalingRuleName;
    }
    /**
     * @return Timing elastic expansion.
     * 
     */
    public List<GetApplicationScalingRulesRuleScalingRuleTimer> scalingRuleTimers() {
        return this.scalingRuleTimers;
    }
    /**
     * @return Flexible strategy type.
     * 
     */
    public String scalingRuleType() {
        return this.scalingRuleType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetApplicationScalingRulesRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String appId;
        private String createTime;
        private String id;
        private Boolean scalingRuleEnable;
        private List<GetApplicationScalingRulesRuleScalingRuleMetric> scalingRuleMetrics;
        private String scalingRuleName;
        private List<GetApplicationScalingRulesRuleScalingRuleTimer> scalingRuleTimers;
        private String scalingRuleType;
        public Builder() {}
        public Builder(GetApplicationScalingRulesRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.appId = defaults.appId;
    	      this.createTime = defaults.createTime;
    	      this.id = defaults.id;
    	      this.scalingRuleEnable = defaults.scalingRuleEnable;
    	      this.scalingRuleMetrics = defaults.scalingRuleMetrics;
    	      this.scalingRuleName = defaults.scalingRuleName;
    	      this.scalingRuleTimers = defaults.scalingRuleTimers;
    	      this.scalingRuleType = defaults.scalingRuleType;
        }

        @CustomType.Setter
        public Builder appId(String appId) {
            this.appId = Objects.requireNonNull(appId);
            return this;
        }
        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder scalingRuleEnable(Boolean scalingRuleEnable) {
            this.scalingRuleEnable = Objects.requireNonNull(scalingRuleEnable);
            return this;
        }
        @CustomType.Setter
        public Builder scalingRuleMetrics(List<GetApplicationScalingRulesRuleScalingRuleMetric> scalingRuleMetrics) {
            this.scalingRuleMetrics = Objects.requireNonNull(scalingRuleMetrics);
            return this;
        }
        public Builder scalingRuleMetrics(GetApplicationScalingRulesRuleScalingRuleMetric... scalingRuleMetrics) {
            return scalingRuleMetrics(List.of(scalingRuleMetrics));
        }
        @CustomType.Setter
        public Builder scalingRuleName(String scalingRuleName) {
            this.scalingRuleName = Objects.requireNonNull(scalingRuleName);
            return this;
        }
        @CustomType.Setter
        public Builder scalingRuleTimers(List<GetApplicationScalingRulesRuleScalingRuleTimer> scalingRuleTimers) {
            this.scalingRuleTimers = Objects.requireNonNull(scalingRuleTimers);
            return this;
        }
        public Builder scalingRuleTimers(GetApplicationScalingRulesRuleScalingRuleTimer... scalingRuleTimers) {
            return scalingRuleTimers(List.of(scalingRuleTimers));
        }
        @CustomType.Setter
        public Builder scalingRuleType(String scalingRuleType) {
            this.scalingRuleType = Objects.requireNonNull(scalingRuleType);
            return this;
        }
        public GetApplicationScalingRulesRule build() {
            final var o = new GetApplicationScalingRulesRule();
            o.appId = appId;
            o.createTime = createTime;
            o.id = id;
            o.scalingRuleEnable = scalingRuleEnable;
            o.scalingRuleMetrics = scalingRuleMetrics;
            o.scalingRuleName = scalingRuleName;
            o.scalingRuleTimers = scalingRuleTimers;
            o.scalingRuleType = scalingRuleType;
            return o;
        }
    }
}