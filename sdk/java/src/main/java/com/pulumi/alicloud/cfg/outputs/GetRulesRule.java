// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cfg.outputs;

import com.pulumi.alicloud.cfg.outputs.GetRulesRuleCompliance;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetRulesRule {
    /**
     * @return The ID of the Alicloud account.
     * 
     */
    private String accountId;
    private String compliancePackId;
    /**
     * @return The information about the compliance evaluations based on the rule.
     * 
     */
    private List<GetRulesRuleCompliance> compliances;
    /**
     * @return The ARN of the Config Rule.
     * 
     */
    private String configRuleArn;
    /**
     * @return The ID of the Config Rule.
     * 
     */
    private String configRuleId;
    /**
     * @return Field `config_rule_state` has been deprecated from provider version 1.124.1. New field `status` instead.
     * 
     */
    private String configRuleState;
    /**
     * @return (Available in 1.124.1+) A list of trigger types of config rule.
     * 
     */
    private String configRuleTriggerTypes;
    /**
     * @return The description of the Config Rule.
     * 
     */
    private String description;
    /**
     * @return Event source of the Config Rule.
     * 
     */
    private String eventSource;
    /**
     * @return (Available in 1.124.1+) The scope of exclude of resource ids.
     * 
     */
    private String excludeResourceIdsScope;
    /**
     * @return The ID of the Config Rule.
     * 
     */
    private String id;
    /**
     * @return The input parameters of the Config Rule.
     * 
     */
    private Map<String,Object> inputParameters;
    /**
     * @return (Available in 1.124.1+) The frequency of maximum execution.
     * 
     */
    private String maximumExecutionFrequency;
    /**
     * @return the timestamp of the Config Rule modified.
     * 
     */
    private String modifiedTimestamp;
    /**
     * @return (Available in 1.124.1+) The scope of region ids.
     * 
     */
    private String regionIdsScope;
    /**
     * @return (Available in 1.124.1+) The scope of resource group ids.
     * 
     */
    private String resourceGroupIdsScope;
    /**
     * @return (Available in 1.124.1+) The scope of resource types.
     * 
     */
    private List<String> resourceTypesScopes;
    /**
     * @return The risk level of Config Rule. Valid values: `1`: Critical ,`2`: Warning , `3`: Info.
     * 
     */
    private Integer riskLevel;
    /**
     * @return The name of config rule.
     * 
     */
    private String ruleName;
    /**
     * @return The types of the resources to be evaluated against the rule.
     * 
     */
    private List<String> scopeComplianceResourceTypes;
    /**
     * @return Rule trigger mechanism.
     * 
     */
    private String sourceDetailMessageType;
    /**
     * @return The identifier of the managed rule or the arn of the custom function.
     * 
     */
    private String sourceIdentifier;
    /**
     * @return Rule execution cycle.
     * 
     */
    private String sourceMaximumExecutionFrequency;
    /**
     * @return The source owner of the Config Rule.
     * 
     */
    private String sourceOwner;
    /**
     * @return The status of the config rule, valid values: `ACTIVE`, `DELETING`, `EVALUATING` and `INACTIVE`.
     * 
     */
    private String status;
    /**
     * @return (Available in 1.124.1+) The scope of tag key.
     * 
     */
    private String tagKeyScope;
    /**
     * @return (Available in 1.124.1+) The scope of tag value.
     * 
     */
    private String tagValueScope;

    private GetRulesRule() {}
    /**
     * @return The ID of the Alicloud account.
     * 
     */
    public String accountId() {
        return this.accountId;
    }
    public String compliancePackId() {
        return this.compliancePackId;
    }
    /**
     * @return The information about the compliance evaluations based on the rule.
     * 
     */
    public List<GetRulesRuleCompliance> compliances() {
        return this.compliances;
    }
    /**
     * @return The ARN of the Config Rule.
     * 
     */
    public String configRuleArn() {
        return this.configRuleArn;
    }
    /**
     * @return The ID of the Config Rule.
     * 
     */
    public String configRuleId() {
        return this.configRuleId;
    }
    /**
     * @return Field `config_rule_state` has been deprecated from provider version 1.124.1. New field `status` instead.
     * 
     */
    public String configRuleState() {
        return this.configRuleState;
    }
    /**
     * @return (Available in 1.124.1+) A list of trigger types of config rule.
     * 
     */
    public String configRuleTriggerTypes() {
        return this.configRuleTriggerTypes;
    }
    /**
     * @return The description of the Config Rule.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return Event source of the Config Rule.
     * 
     */
    public String eventSource() {
        return this.eventSource;
    }
    /**
     * @return (Available in 1.124.1+) The scope of exclude of resource ids.
     * 
     */
    public String excludeResourceIdsScope() {
        return this.excludeResourceIdsScope;
    }
    /**
     * @return The ID of the Config Rule.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The input parameters of the Config Rule.
     * 
     */
    public Map<String,Object> inputParameters() {
        return this.inputParameters;
    }
    /**
     * @return (Available in 1.124.1+) The frequency of maximum execution.
     * 
     */
    public String maximumExecutionFrequency() {
        return this.maximumExecutionFrequency;
    }
    /**
     * @return the timestamp of the Config Rule modified.
     * 
     */
    public String modifiedTimestamp() {
        return this.modifiedTimestamp;
    }
    /**
     * @return (Available in 1.124.1+) The scope of region ids.
     * 
     */
    public String regionIdsScope() {
        return this.regionIdsScope;
    }
    /**
     * @return (Available in 1.124.1+) The scope of resource group ids.
     * 
     */
    public String resourceGroupIdsScope() {
        return this.resourceGroupIdsScope;
    }
    /**
     * @return (Available in 1.124.1+) The scope of resource types.
     * 
     */
    public List<String> resourceTypesScopes() {
        return this.resourceTypesScopes;
    }
    /**
     * @return The risk level of Config Rule. Valid values: `1`: Critical ,`2`: Warning , `3`: Info.
     * 
     */
    public Integer riskLevel() {
        return this.riskLevel;
    }
    /**
     * @return The name of config rule.
     * 
     */
    public String ruleName() {
        return this.ruleName;
    }
    /**
     * @return The types of the resources to be evaluated against the rule.
     * 
     */
    public List<String> scopeComplianceResourceTypes() {
        return this.scopeComplianceResourceTypes;
    }
    /**
     * @return Rule trigger mechanism.
     * 
     */
    public String sourceDetailMessageType() {
        return this.sourceDetailMessageType;
    }
    /**
     * @return The identifier of the managed rule or the arn of the custom function.
     * 
     */
    public String sourceIdentifier() {
        return this.sourceIdentifier;
    }
    /**
     * @return Rule execution cycle.
     * 
     */
    public String sourceMaximumExecutionFrequency() {
        return this.sourceMaximumExecutionFrequency;
    }
    /**
     * @return The source owner of the Config Rule.
     * 
     */
    public String sourceOwner() {
        return this.sourceOwner;
    }
    /**
     * @return The status of the config rule, valid values: `ACTIVE`, `DELETING`, `EVALUATING` and `INACTIVE`.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return (Available in 1.124.1+) The scope of tag key.
     * 
     */
    public String tagKeyScope() {
        return this.tagKeyScope;
    }
    /**
     * @return (Available in 1.124.1+) The scope of tag value.
     * 
     */
    public String tagValueScope() {
        return this.tagValueScope;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRulesRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accountId;
        private String compliancePackId;
        private List<GetRulesRuleCompliance> compliances;
        private String configRuleArn;
        private String configRuleId;
        private String configRuleState;
        private String configRuleTriggerTypes;
        private String description;
        private String eventSource;
        private String excludeResourceIdsScope;
        private String id;
        private Map<String,Object> inputParameters;
        private String maximumExecutionFrequency;
        private String modifiedTimestamp;
        private String regionIdsScope;
        private String resourceGroupIdsScope;
        private List<String> resourceTypesScopes;
        private Integer riskLevel;
        private String ruleName;
        private List<String> scopeComplianceResourceTypes;
        private String sourceDetailMessageType;
        private String sourceIdentifier;
        private String sourceMaximumExecutionFrequency;
        private String sourceOwner;
        private String status;
        private String tagKeyScope;
        private String tagValueScope;
        public Builder() {}
        public Builder(GetRulesRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accountId = defaults.accountId;
    	      this.compliancePackId = defaults.compliancePackId;
    	      this.compliances = defaults.compliances;
    	      this.configRuleArn = defaults.configRuleArn;
    	      this.configRuleId = defaults.configRuleId;
    	      this.configRuleState = defaults.configRuleState;
    	      this.configRuleTriggerTypes = defaults.configRuleTriggerTypes;
    	      this.description = defaults.description;
    	      this.eventSource = defaults.eventSource;
    	      this.excludeResourceIdsScope = defaults.excludeResourceIdsScope;
    	      this.id = defaults.id;
    	      this.inputParameters = defaults.inputParameters;
    	      this.maximumExecutionFrequency = defaults.maximumExecutionFrequency;
    	      this.modifiedTimestamp = defaults.modifiedTimestamp;
    	      this.regionIdsScope = defaults.regionIdsScope;
    	      this.resourceGroupIdsScope = defaults.resourceGroupIdsScope;
    	      this.resourceTypesScopes = defaults.resourceTypesScopes;
    	      this.riskLevel = defaults.riskLevel;
    	      this.ruleName = defaults.ruleName;
    	      this.scopeComplianceResourceTypes = defaults.scopeComplianceResourceTypes;
    	      this.sourceDetailMessageType = defaults.sourceDetailMessageType;
    	      this.sourceIdentifier = defaults.sourceIdentifier;
    	      this.sourceMaximumExecutionFrequency = defaults.sourceMaximumExecutionFrequency;
    	      this.sourceOwner = defaults.sourceOwner;
    	      this.status = defaults.status;
    	      this.tagKeyScope = defaults.tagKeyScope;
    	      this.tagValueScope = defaults.tagValueScope;
        }

        @CustomType.Setter
        public Builder accountId(String accountId) {
            this.accountId = Objects.requireNonNull(accountId);
            return this;
        }
        @CustomType.Setter
        public Builder compliancePackId(String compliancePackId) {
            this.compliancePackId = Objects.requireNonNull(compliancePackId);
            return this;
        }
        @CustomType.Setter
        public Builder compliances(List<GetRulesRuleCompliance> compliances) {
            this.compliances = Objects.requireNonNull(compliances);
            return this;
        }
        public Builder compliances(GetRulesRuleCompliance... compliances) {
            return compliances(List.of(compliances));
        }
        @CustomType.Setter
        public Builder configRuleArn(String configRuleArn) {
            this.configRuleArn = Objects.requireNonNull(configRuleArn);
            return this;
        }
        @CustomType.Setter
        public Builder configRuleId(String configRuleId) {
            this.configRuleId = Objects.requireNonNull(configRuleId);
            return this;
        }
        @CustomType.Setter
        public Builder configRuleState(String configRuleState) {
            this.configRuleState = Objects.requireNonNull(configRuleState);
            return this;
        }
        @CustomType.Setter
        public Builder configRuleTriggerTypes(String configRuleTriggerTypes) {
            this.configRuleTriggerTypes = Objects.requireNonNull(configRuleTriggerTypes);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder eventSource(String eventSource) {
            this.eventSource = Objects.requireNonNull(eventSource);
            return this;
        }
        @CustomType.Setter
        public Builder excludeResourceIdsScope(String excludeResourceIdsScope) {
            this.excludeResourceIdsScope = Objects.requireNonNull(excludeResourceIdsScope);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder inputParameters(Map<String,Object> inputParameters) {
            this.inputParameters = Objects.requireNonNull(inputParameters);
            return this;
        }
        @CustomType.Setter
        public Builder maximumExecutionFrequency(String maximumExecutionFrequency) {
            this.maximumExecutionFrequency = Objects.requireNonNull(maximumExecutionFrequency);
            return this;
        }
        @CustomType.Setter
        public Builder modifiedTimestamp(String modifiedTimestamp) {
            this.modifiedTimestamp = Objects.requireNonNull(modifiedTimestamp);
            return this;
        }
        @CustomType.Setter
        public Builder regionIdsScope(String regionIdsScope) {
            this.regionIdsScope = Objects.requireNonNull(regionIdsScope);
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroupIdsScope(String resourceGroupIdsScope) {
            this.resourceGroupIdsScope = Objects.requireNonNull(resourceGroupIdsScope);
            return this;
        }
        @CustomType.Setter
        public Builder resourceTypesScopes(List<String> resourceTypesScopes) {
            this.resourceTypesScopes = Objects.requireNonNull(resourceTypesScopes);
            return this;
        }
        public Builder resourceTypesScopes(String... resourceTypesScopes) {
            return resourceTypesScopes(List.of(resourceTypesScopes));
        }
        @CustomType.Setter
        public Builder riskLevel(Integer riskLevel) {
            this.riskLevel = Objects.requireNonNull(riskLevel);
            return this;
        }
        @CustomType.Setter
        public Builder ruleName(String ruleName) {
            this.ruleName = Objects.requireNonNull(ruleName);
            return this;
        }
        @CustomType.Setter
        public Builder scopeComplianceResourceTypes(List<String> scopeComplianceResourceTypes) {
            this.scopeComplianceResourceTypes = Objects.requireNonNull(scopeComplianceResourceTypes);
            return this;
        }
        public Builder scopeComplianceResourceTypes(String... scopeComplianceResourceTypes) {
            return scopeComplianceResourceTypes(List.of(scopeComplianceResourceTypes));
        }
        @CustomType.Setter
        public Builder sourceDetailMessageType(String sourceDetailMessageType) {
            this.sourceDetailMessageType = Objects.requireNonNull(sourceDetailMessageType);
            return this;
        }
        @CustomType.Setter
        public Builder sourceIdentifier(String sourceIdentifier) {
            this.sourceIdentifier = Objects.requireNonNull(sourceIdentifier);
            return this;
        }
        @CustomType.Setter
        public Builder sourceMaximumExecutionFrequency(String sourceMaximumExecutionFrequency) {
            this.sourceMaximumExecutionFrequency = Objects.requireNonNull(sourceMaximumExecutionFrequency);
            return this;
        }
        @CustomType.Setter
        public Builder sourceOwner(String sourceOwner) {
            this.sourceOwner = Objects.requireNonNull(sourceOwner);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder tagKeyScope(String tagKeyScope) {
            this.tagKeyScope = Objects.requireNonNull(tagKeyScope);
            return this;
        }
        @CustomType.Setter
        public Builder tagValueScope(String tagValueScope) {
            this.tagValueScope = Objects.requireNonNull(tagValueScope);
            return this;
        }
        public GetRulesRule build() {
            final var o = new GetRulesRule();
            o.accountId = accountId;
            o.compliancePackId = compliancePackId;
            o.compliances = compliances;
            o.configRuleArn = configRuleArn;
            o.configRuleId = configRuleId;
            o.configRuleState = configRuleState;
            o.configRuleTriggerTypes = configRuleTriggerTypes;
            o.description = description;
            o.eventSource = eventSource;
            o.excludeResourceIdsScope = excludeResourceIdsScope;
            o.id = id;
            o.inputParameters = inputParameters;
            o.maximumExecutionFrequency = maximumExecutionFrequency;
            o.modifiedTimestamp = modifiedTimestamp;
            o.regionIdsScope = regionIdsScope;
            o.resourceGroupIdsScope = resourceGroupIdsScope;
            o.resourceTypesScopes = resourceTypesScopes;
            o.riskLevel = riskLevel;
            o.ruleName = ruleName;
            o.scopeComplianceResourceTypes = scopeComplianceResourceTypes;
            o.sourceDetailMessageType = sourceDetailMessageType;
            o.sourceIdentifier = sourceIdentifier;
            o.sourceMaximumExecutionFrequency = sourceMaximumExecutionFrequency;
            o.sourceOwner = sourceOwner;
            o.status = status;
            o.tagKeyScope = tagKeyScope;
            o.tagValueScope = tagValueScope;
            return o;
        }
    }
}
