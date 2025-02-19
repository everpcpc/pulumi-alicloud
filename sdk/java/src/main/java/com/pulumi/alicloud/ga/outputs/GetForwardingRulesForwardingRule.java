// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.outputs;

import com.pulumi.alicloud.ga.outputs.GetForwardingRulesForwardingRuleRuleAction;
import com.pulumi.alicloud.ga.outputs.GetForwardingRulesForwardingRuleRuleCondition;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetForwardingRulesForwardingRule {
    /**
     * @return Forwarding Policy ID.
     * 
     */
    private String forwardingRuleId;
    /**
     * @return Forwarding policy name. The length of the name is 2-128 English or Chinese characters.
     * 
     */
    private String forwardingRuleName;
    /**
     * @return Forwarding Policy Status.
     * 
     */
    private String forwardingRuleStatus;
    private String id;
    /**
     * @return The ID of the listener.
     * 
     */
    private String listenerId;
    /**
     * @return Forwarding policy priority.
     * 
     */
    private Integer priority;
    /**
     * @return The IP protocol used by the GA instance.
     * `order` - Forwarding priority.
     * `rule_action_type` - Forward action type.
     * `forward_group_config` - Forwarding configuration.
     * `server_group_tuples` - Terminal node group configuration.
     * `endpoint_group_id` - Terminal node group ID.
     * 
     */
    private List<GetForwardingRulesForwardingRuleRuleAction> ruleActions;
    /**
     * @return Forward action.
     * `rule_condition_type` - Forwarding condition type.
     * `path_config` - Path configuration information.
     * `values` - The length of the path is 1-128 characters.
     * `host_config` - Domain name configuration information.
     * `values` - The domain name is 3-128 characters long.
     * 
     */
    private List<GetForwardingRulesForwardingRuleRuleCondition> ruleConditions;

    private GetForwardingRulesForwardingRule() {}
    /**
     * @return Forwarding Policy ID.
     * 
     */
    public String forwardingRuleId() {
        return this.forwardingRuleId;
    }
    /**
     * @return Forwarding policy name. The length of the name is 2-128 English or Chinese characters.
     * 
     */
    public String forwardingRuleName() {
        return this.forwardingRuleName;
    }
    /**
     * @return Forwarding Policy Status.
     * 
     */
    public String forwardingRuleStatus() {
        return this.forwardingRuleStatus;
    }
    public String id() {
        return this.id;
    }
    /**
     * @return The ID of the listener.
     * 
     */
    public String listenerId() {
        return this.listenerId;
    }
    /**
     * @return Forwarding policy priority.
     * 
     */
    public Integer priority() {
        return this.priority;
    }
    /**
     * @return The IP protocol used by the GA instance.
     * `order` - Forwarding priority.
     * `rule_action_type` - Forward action type.
     * `forward_group_config` - Forwarding configuration.
     * `server_group_tuples` - Terminal node group configuration.
     * `endpoint_group_id` - Terminal node group ID.
     * 
     */
    public List<GetForwardingRulesForwardingRuleRuleAction> ruleActions() {
        return this.ruleActions;
    }
    /**
     * @return Forward action.
     * `rule_condition_type` - Forwarding condition type.
     * `path_config` - Path configuration information.
     * `values` - The length of the path is 1-128 characters.
     * `host_config` - Domain name configuration information.
     * `values` - The domain name is 3-128 characters long.
     * 
     */
    public List<GetForwardingRulesForwardingRuleRuleCondition> ruleConditions() {
        return this.ruleConditions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetForwardingRulesForwardingRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String forwardingRuleId;
        private String forwardingRuleName;
        private String forwardingRuleStatus;
        private String id;
        private String listenerId;
        private Integer priority;
        private List<GetForwardingRulesForwardingRuleRuleAction> ruleActions;
        private List<GetForwardingRulesForwardingRuleRuleCondition> ruleConditions;
        public Builder() {}
        public Builder(GetForwardingRulesForwardingRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.forwardingRuleId = defaults.forwardingRuleId;
    	      this.forwardingRuleName = defaults.forwardingRuleName;
    	      this.forwardingRuleStatus = defaults.forwardingRuleStatus;
    	      this.id = defaults.id;
    	      this.listenerId = defaults.listenerId;
    	      this.priority = defaults.priority;
    	      this.ruleActions = defaults.ruleActions;
    	      this.ruleConditions = defaults.ruleConditions;
        }

        @CustomType.Setter
        public Builder forwardingRuleId(String forwardingRuleId) {
            this.forwardingRuleId = Objects.requireNonNull(forwardingRuleId);
            return this;
        }
        @CustomType.Setter
        public Builder forwardingRuleName(String forwardingRuleName) {
            this.forwardingRuleName = Objects.requireNonNull(forwardingRuleName);
            return this;
        }
        @CustomType.Setter
        public Builder forwardingRuleStatus(String forwardingRuleStatus) {
            this.forwardingRuleStatus = Objects.requireNonNull(forwardingRuleStatus);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder listenerId(String listenerId) {
            this.listenerId = Objects.requireNonNull(listenerId);
            return this;
        }
        @CustomType.Setter
        public Builder priority(Integer priority) {
            this.priority = Objects.requireNonNull(priority);
            return this;
        }
        @CustomType.Setter
        public Builder ruleActions(List<GetForwardingRulesForwardingRuleRuleAction> ruleActions) {
            this.ruleActions = Objects.requireNonNull(ruleActions);
            return this;
        }
        public Builder ruleActions(GetForwardingRulesForwardingRuleRuleAction... ruleActions) {
            return ruleActions(List.of(ruleActions));
        }
        @CustomType.Setter
        public Builder ruleConditions(List<GetForwardingRulesForwardingRuleRuleCondition> ruleConditions) {
            this.ruleConditions = Objects.requireNonNull(ruleConditions);
            return this;
        }
        public Builder ruleConditions(GetForwardingRulesForwardingRuleRuleCondition... ruleConditions) {
            return ruleConditions(List.of(ruleConditions));
        }
        public GetForwardingRulesForwardingRule build() {
            final var o = new GetForwardingRulesForwardingRule();
            o.forwardingRuleId = forwardingRuleId;
            o.forwardingRuleName = forwardingRuleName;
            o.forwardingRuleStatus = forwardingRuleStatus;
            o.id = id;
            o.listenerId = listenerId;
            o.priority = priority;
            o.ruleActions = ruleActions;
            o.ruleConditions = ruleConditions;
            return o;
        }
    }
}
