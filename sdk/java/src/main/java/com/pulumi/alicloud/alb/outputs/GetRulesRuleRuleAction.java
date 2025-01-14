// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.outputs;

import com.pulumi.alicloud.alb.outputs.GetRulesRuleRuleActionFixedResponseConfig;
import com.pulumi.alicloud.alb.outputs.GetRulesRuleRuleActionForwardGroupConfig;
import com.pulumi.alicloud.alb.outputs.GetRulesRuleRuleActionInsertHeaderConfig;
import com.pulumi.alicloud.alb.outputs.GetRulesRuleRuleActionRedirectConfig;
import com.pulumi.alicloud.alb.outputs.GetRulesRuleRuleActionRewriteConfig;
import com.pulumi.alicloud.alb.outputs.GetRulesRuleRuleActionTrafficLimitConfig;
import com.pulumi.alicloud.alb.outputs.GetRulesRuleRuleActionTrafficMirrorConfig;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetRulesRuleRuleAction {
    /**
     * @return The configuration of the fixed response.
     * 
     */
    private List<GetRulesRuleRuleActionFixedResponseConfig> fixedResponseConfigs;
    /**
     * @return The configurations of the destination server groups.
     * 
     */
    private List<GetRulesRuleRuleActionForwardGroupConfig> forwardGroupConfigs;
    /**
     * @return The configuration of the inserted header field.
     * 
     */
    private List<GetRulesRuleRuleActionInsertHeaderConfig> insertHeaderConfigs;
    /**
     * @return The order of the forwarding rule actions. Valid values:1 to 50000. The actions are performed in ascending order. You cannot leave this parameter empty. Each value must be unique.
     * 
     */
    private Integer order;
    /**
     * @return The configuration of the external redirect action.
     * 
     */
    private List<GetRulesRuleRuleActionRedirectConfig> redirectConfigs;
    /**
     * @return The redirect action within ALB.
     * 
     */
    private List<GetRulesRuleRuleActionRewriteConfig> rewriteConfigs;
    /**
     * @return The Flow speed limit.
     * 
     */
    private List<GetRulesRuleRuleActionTrafficLimitConfig> trafficLimitConfigs;
    /**
     * @return The Traffic mirroring.
     * 
     */
    private List<GetRulesRuleRuleActionTrafficMirrorConfig> trafficMirrorConfigs;
    /**
     * @return The type of the forwarding rule.
     * 
     */
    private String type;

    private GetRulesRuleRuleAction() {}
    /**
     * @return The configuration of the fixed response.
     * 
     */
    public List<GetRulesRuleRuleActionFixedResponseConfig> fixedResponseConfigs() {
        return this.fixedResponseConfigs;
    }
    /**
     * @return The configurations of the destination server groups.
     * 
     */
    public List<GetRulesRuleRuleActionForwardGroupConfig> forwardGroupConfigs() {
        return this.forwardGroupConfigs;
    }
    /**
     * @return The configuration of the inserted header field.
     * 
     */
    public List<GetRulesRuleRuleActionInsertHeaderConfig> insertHeaderConfigs() {
        return this.insertHeaderConfigs;
    }
    /**
     * @return The order of the forwarding rule actions. Valid values:1 to 50000. The actions are performed in ascending order. You cannot leave this parameter empty. Each value must be unique.
     * 
     */
    public Integer order() {
        return this.order;
    }
    /**
     * @return The configuration of the external redirect action.
     * 
     */
    public List<GetRulesRuleRuleActionRedirectConfig> redirectConfigs() {
        return this.redirectConfigs;
    }
    /**
     * @return The redirect action within ALB.
     * 
     */
    public List<GetRulesRuleRuleActionRewriteConfig> rewriteConfigs() {
        return this.rewriteConfigs;
    }
    /**
     * @return The Flow speed limit.
     * 
     */
    public List<GetRulesRuleRuleActionTrafficLimitConfig> trafficLimitConfigs() {
        return this.trafficLimitConfigs;
    }
    /**
     * @return The Traffic mirroring.
     * 
     */
    public List<GetRulesRuleRuleActionTrafficMirrorConfig> trafficMirrorConfigs() {
        return this.trafficMirrorConfigs;
    }
    /**
     * @return The type of the forwarding rule.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRulesRuleRuleAction defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetRulesRuleRuleActionFixedResponseConfig> fixedResponseConfigs;
        private List<GetRulesRuleRuleActionForwardGroupConfig> forwardGroupConfigs;
        private List<GetRulesRuleRuleActionInsertHeaderConfig> insertHeaderConfigs;
        private Integer order;
        private List<GetRulesRuleRuleActionRedirectConfig> redirectConfigs;
        private List<GetRulesRuleRuleActionRewriteConfig> rewriteConfigs;
        private List<GetRulesRuleRuleActionTrafficLimitConfig> trafficLimitConfigs;
        private List<GetRulesRuleRuleActionTrafficMirrorConfig> trafficMirrorConfigs;
        private String type;
        public Builder() {}
        public Builder(GetRulesRuleRuleAction defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fixedResponseConfigs = defaults.fixedResponseConfigs;
    	      this.forwardGroupConfigs = defaults.forwardGroupConfigs;
    	      this.insertHeaderConfigs = defaults.insertHeaderConfigs;
    	      this.order = defaults.order;
    	      this.redirectConfigs = defaults.redirectConfigs;
    	      this.rewriteConfigs = defaults.rewriteConfigs;
    	      this.trafficLimitConfigs = defaults.trafficLimitConfigs;
    	      this.trafficMirrorConfigs = defaults.trafficMirrorConfigs;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder fixedResponseConfigs(List<GetRulesRuleRuleActionFixedResponseConfig> fixedResponseConfigs) {
            this.fixedResponseConfigs = Objects.requireNonNull(fixedResponseConfigs);
            return this;
        }
        public Builder fixedResponseConfigs(GetRulesRuleRuleActionFixedResponseConfig... fixedResponseConfigs) {
            return fixedResponseConfigs(List.of(fixedResponseConfigs));
        }
        @CustomType.Setter
        public Builder forwardGroupConfigs(List<GetRulesRuleRuleActionForwardGroupConfig> forwardGroupConfigs) {
            this.forwardGroupConfigs = Objects.requireNonNull(forwardGroupConfigs);
            return this;
        }
        public Builder forwardGroupConfigs(GetRulesRuleRuleActionForwardGroupConfig... forwardGroupConfigs) {
            return forwardGroupConfigs(List.of(forwardGroupConfigs));
        }
        @CustomType.Setter
        public Builder insertHeaderConfigs(List<GetRulesRuleRuleActionInsertHeaderConfig> insertHeaderConfigs) {
            this.insertHeaderConfigs = Objects.requireNonNull(insertHeaderConfigs);
            return this;
        }
        public Builder insertHeaderConfigs(GetRulesRuleRuleActionInsertHeaderConfig... insertHeaderConfigs) {
            return insertHeaderConfigs(List.of(insertHeaderConfigs));
        }
        @CustomType.Setter
        public Builder order(Integer order) {
            this.order = Objects.requireNonNull(order);
            return this;
        }
        @CustomType.Setter
        public Builder redirectConfigs(List<GetRulesRuleRuleActionRedirectConfig> redirectConfigs) {
            this.redirectConfigs = Objects.requireNonNull(redirectConfigs);
            return this;
        }
        public Builder redirectConfigs(GetRulesRuleRuleActionRedirectConfig... redirectConfigs) {
            return redirectConfigs(List.of(redirectConfigs));
        }
        @CustomType.Setter
        public Builder rewriteConfigs(List<GetRulesRuleRuleActionRewriteConfig> rewriteConfigs) {
            this.rewriteConfigs = Objects.requireNonNull(rewriteConfigs);
            return this;
        }
        public Builder rewriteConfigs(GetRulesRuleRuleActionRewriteConfig... rewriteConfigs) {
            return rewriteConfigs(List.of(rewriteConfigs));
        }
        @CustomType.Setter
        public Builder trafficLimitConfigs(List<GetRulesRuleRuleActionTrafficLimitConfig> trafficLimitConfigs) {
            this.trafficLimitConfigs = Objects.requireNonNull(trafficLimitConfigs);
            return this;
        }
        public Builder trafficLimitConfigs(GetRulesRuleRuleActionTrafficLimitConfig... trafficLimitConfigs) {
            return trafficLimitConfigs(List.of(trafficLimitConfigs));
        }
        @CustomType.Setter
        public Builder trafficMirrorConfigs(List<GetRulesRuleRuleActionTrafficMirrorConfig> trafficMirrorConfigs) {
            this.trafficMirrorConfigs = Objects.requireNonNull(trafficMirrorConfigs);
            return this;
        }
        public Builder trafficMirrorConfigs(GetRulesRuleRuleActionTrafficMirrorConfig... trafficMirrorConfigs) {
            return trafficMirrorConfigs(List.of(trafficMirrorConfigs));
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetRulesRuleRuleAction build() {
            final var o = new GetRulesRuleRuleAction();
            o.fixedResponseConfigs = fixedResponseConfigs;
            o.forwardGroupConfigs = forwardGroupConfigs;
            o.insertHeaderConfigs = insertHeaderConfigs;
            o.order = order;
            o.redirectConfigs = redirectConfigs;
            o.rewriteConfigs = rewriteConfigs;
            o.trafficLimitConfigs = trafficLimitConfigs;
            o.trafficMirrorConfigs = trafficMirrorConfigs;
            o.type = type;
            return o;
        }
    }
}
