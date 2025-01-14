// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cfg.outputs;

import com.pulumi.alicloud.cfg.outputs.GetAggregateConfigRulesRuleCompliance;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetAggregateConfigRulesRule {
    /**
     * @return The Aliyun User ID.
     * 
     */
    private String accountId;
    /**
     * @return The config rule name.
     * 
     */
    private String aggregateConfigRuleName;
    /**
     * @return The ID of aggregator.
     * 
     */
    private String aggregatorId;
    /**
     * @return The ID of Compliance Package.
     * 
     */
    private String compliancePackId;
    /**
     * @return The Compliance information.
     * 
     */
    private List<GetAggregateConfigRulesRuleCompliance> compliances;
    /**
     * @return The config rule arn.
     * 
     */
    private String configRuleArn;
    /**
     * @return The ID of the rule.
     * 
     */
    private String configRuleId;
    /**
     * @return The trigger types of config rules.
     * 
     */
    private String configRuleTriggerTypes;
    /**
     * @return The description of the rule.
     * 
     */
    private String description;
    /**
     * @return Event source of the Config Rule.
     * 
     */
    private String eventSource;
    /**
     * @return The id of the resources to be evaluated against the rule.
     * 
     */
    private String excludeResourceIdsScope;
    /**
     * @return The ID of the Aggregate Config Rule.
     * 
     */
    private String id;
    /**
     * @return The settings of the input parameters for the rule.
     * 
     */
    private Map<String,Object> inputParameters;
    /**
     * @return The frequency of the compliance evaluations.
     * 
     */
    private String maximumExecutionFrequency;
    /**
     * @return The timestamp when the rule was last modified.
     * 
     */
    private String modifiedTimestamp;
    /**
     * @return The scope of resource region ids.
     * 
     */
    private String regionIdsScope;
    /**
     * @return The scope of resource group ids.
     * 
     */
    private String resourceGroupIdsScope;
    private List<String> resourceTypesScopes;
    /**
     * @return Optional, ForceNew) The Risk Level. Valid values `1`: critical, `2`: warning, `3`: info.
     * 
     */
    private Integer riskLevel;
    /**
     * @return The identifier of the managed rule or the arn of the custom function.
     * 
     */
    private String sourceIdentifier;
    /**
     * @return The source owner of the Config Rule.
     * 
     */
    private String sourceOwner;
    /**
     * @return The state of the config rule, valid values: `ACTIVE`, `DELETING`, `EVALUATING` and `INACTIVE`.
     * 
     */
    private String status;
    /**
     * @return The scope of tay key.
     * 
     */
    private String tagKeyScope;
    /**
     * @return The scope of tay value.
     * 
     */
    private String tagValueScope;

    private GetAggregateConfigRulesRule() {}
    /**
     * @return The Aliyun User ID.
     * 
     */
    public String accountId() {
        return this.accountId;
    }
    /**
     * @return The config rule name.
     * 
     */
    public String aggregateConfigRuleName() {
        return this.aggregateConfigRuleName;
    }
    /**
     * @return The ID of aggregator.
     * 
     */
    public String aggregatorId() {
        return this.aggregatorId;
    }
    /**
     * @return The ID of Compliance Package.
     * 
     */
    public String compliancePackId() {
        return this.compliancePackId;
    }
    /**
     * @return The Compliance information.
     * 
     */
    public List<GetAggregateConfigRulesRuleCompliance> compliances() {
        return this.compliances;
    }
    /**
     * @return The config rule arn.
     * 
     */
    public String configRuleArn() {
        return this.configRuleArn;
    }
    /**
     * @return The ID of the rule.
     * 
     */
    public String configRuleId() {
        return this.configRuleId;
    }
    /**
     * @return The trigger types of config rules.
     * 
     */
    public String configRuleTriggerTypes() {
        return this.configRuleTriggerTypes;
    }
    /**
     * @return The description of the rule.
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
     * @return The id of the resources to be evaluated against the rule.
     * 
     */
    public String excludeResourceIdsScope() {
        return this.excludeResourceIdsScope;
    }
    /**
     * @return The ID of the Aggregate Config Rule.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The settings of the input parameters for the rule.
     * 
     */
    public Map<String,Object> inputParameters() {
        return this.inputParameters;
    }
    /**
     * @return The frequency of the compliance evaluations.
     * 
     */
    public String maximumExecutionFrequency() {
        return this.maximumExecutionFrequency;
    }
    /**
     * @return The timestamp when the rule was last modified.
     * 
     */
    public String modifiedTimestamp() {
        return this.modifiedTimestamp;
    }
    /**
     * @return The scope of resource region ids.
     * 
     */
    public String regionIdsScope() {
        return this.regionIdsScope;
    }
    /**
     * @return The scope of resource group ids.
     * 
     */
    public String resourceGroupIdsScope() {
        return this.resourceGroupIdsScope;
    }
    public List<String> resourceTypesScopes() {
        return this.resourceTypesScopes;
    }
    /**
     * @return Optional, ForceNew) The Risk Level. Valid values `1`: critical, `2`: warning, `3`: info.
     * 
     */
    public Integer riskLevel() {
        return this.riskLevel;
    }
    /**
     * @return The identifier of the managed rule or the arn of the custom function.
     * 
     */
    public String sourceIdentifier() {
        return this.sourceIdentifier;
    }
    /**
     * @return The source owner of the Config Rule.
     * 
     */
    public String sourceOwner() {
        return this.sourceOwner;
    }
    /**
     * @return The state of the config rule, valid values: `ACTIVE`, `DELETING`, `EVALUATING` and `INACTIVE`.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The scope of tay key.
     * 
     */
    public String tagKeyScope() {
        return this.tagKeyScope;
    }
    /**
     * @return The scope of tay value.
     * 
     */
    public String tagValueScope() {
        return this.tagValueScope;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAggregateConfigRulesRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accountId;
        private String aggregateConfigRuleName;
        private String aggregatorId;
        private String compliancePackId;
        private List<GetAggregateConfigRulesRuleCompliance> compliances;
        private String configRuleArn;
        private String configRuleId;
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
        private String sourceIdentifier;
        private String sourceOwner;
        private String status;
        private String tagKeyScope;
        private String tagValueScope;
        public Builder() {}
        public Builder(GetAggregateConfigRulesRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accountId = defaults.accountId;
    	      this.aggregateConfigRuleName = defaults.aggregateConfigRuleName;
    	      this.aggregatorId = defaults.aggregatorId;
    	      this.compliancePackId = defaults.compliancePackId;
    	      this.compliances = defaults.compliances;
    	      this.configRuleArn = defaults.configRuleArn;
    	      this.configRuleId = defaults.configRuleId;
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
    	      this.sourceIdentifier = defaults.sourceIdentifier;
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
        public Builder aggregateConfigRuleName(String aggregateConfigRuleName) {
            this.aggregateConfigRuleName = Objects.requireNonNull(aggregateConfigRuleName);
            return this;
        }
        @CustomType.Setter
        public Builder aggregatorId(String aggregatorId) {
            this.aggregatorId = Objects.requireNonNull(aggregatorId);
            return this;
        }
        @CustomType.Setter
        public Builder compliancePackId(String compliancePackId) {
            this.compliancePackId = Objects.requireNonNull(compliancePackId);
            return this;
        }
        @CustomType.Setter
        public Builder compliances(List<GetAggregateConfigRulesRuleCompliance> compliances) {
            this.compliances = Objects.requireNonNull(compliances);
            return this;
        }
        public Builder compliances(GetAggregateConfigRulesRuleCompliance... compliances) {
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
        public Builder sourceIdentifier(String sourceIdentifier) {
            this.sourceIdentifier = Objects.requireNonNull(sourceIdentifier);
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
        public GetAggregateConfigRulesRule build() {
            final var o = new GetAggregateConfigRulesRule();
            o.accountId = accountId;
            o.aggregateConfigRuleName = aggregateConfigRuleName;
            o.aggregatorId = aggregatorId;
            o.compliancePackId = compliancePackId;
            o.compliances = compliances;
            o.configRuleArn = configRuleArn;
            o.configRuleId = configRuleId;
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
            o.sourceIdentifier = sourceIdentifier;
            o.sourceOwner = sourceOwner;
            o.status = status;
            o.tagKeyScope = tagKeyScope;
            o.tagValueScope = tagValueScope;
            return o;
        }
    }
}
