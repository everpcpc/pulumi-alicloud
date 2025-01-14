// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cfg.outputs;

import com.pulumi.alicloud.cfg.outputs.GetAggregateCompliancePacksPackConfigRule;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetAggregateCompliancePacksPack {
    /**
     * @return The Aliyun User Id.
     * 
     */
    private String accountId;
    /**
     * @return The Aggregate Compliance Package Name.
     * 
     */
    private String aggregateCompliancePackName;
    /**
     * @return The Aggregate Compliance Package Id.
     * 
     */
    private String aggregatorCompliancePackId;
    /**
     * @return The template ID of the Compliance Package.
     * 
     */
    private String compliancePackTemplateId;
    /**
     * @return A list of The Aggregate Compliance Package Rules.
     * 
     */
    private List<GetAggregateCompliancePacksPackConfigRule> configRules;
    /**
     * @return The description of aggregate compliance pack.
     * 
     */
    private String description;
    /**
     * @return The ID of the Aggregate Compliance Pack.
     * 
     */
    private String id;
    /**
     * @return The Risk Level.
     * 
     */
    private Integer riskLevel;
    /**
     * @return The status of the resource. Valid values `ACTIVE`, `CREATING`, `INACTIVE`.
     * 
     */
    private String status;

    private GetAggregateCompliancePacksPack() {}
    /**
     * @return The Aliyun User Id.
     * 
     */
    public String accountId() {
        return this.accountId;
    }
    /**
     * @return The Aggregate Compliance Package Name.
     * 
     */
    public String aggregateCompliancePackName() {
        return this.aggregateCompliancePackName;
    }
    /**
     * @return The Aggregate Compliance Package Id.
     * 
     */
    public String aggregatorCompliancePackId() {
        return this.aggregatorCompliancePackId;
    }
    /**
     * @return The template ID of the Compliance Package.
     * 
     */
    public String compliancePackTemplateId() {
        return this.compliancePackTemplateId;
    }
    /**
     * @return A list of The Aggregate Compliance Package Rules.
     * 
     */
    public List<GetAggregateCompliancePacksPackConfigRule> configRules() {
        return this.configRules;
    }
    /**
     * @return The description of aggregate compliance pack.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The ID of the Aggregate Compliance Pack.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The Risk Level.
     * 
     */
    public Integer riskLevel() {
        return this.riskLevel;
    }
    /**
     * @return The status of the resource. Valid values `ACTIVE`, `CREATING`, `INACTIVE`.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAggregateCompliancePacksPack defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accountId;
        private String aggregateCompliancePackName;
        private String aggregatorCompliancePackId;
        private String compliancePackTemplateId;
        private List<GetAggregateCompliancePacksPackConfigRule> configRules;
        private String description;
        private String id;
        private Integer riskLevel;
        private String status;
        public Builder() {}
        public Builder(GetAggregateCompliancePacksPack defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accountId = defaults.accountId;
    	      this.aggregateCompliancePackName = defaults.aggregateCompliancePackName;
    	      this.aggregatorCompliancePackId = defaults.aggregatorCompliancePackId;
    	      this.compliancePackTemplateId = defaults.compliancePackTemplateId;
    	      this.configRules = defaults.configRules;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.riskLevel = defaults.riskLevel;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder accountId(String accountId) {
            this.accountId = Objects.requireNonNull(accountId);
            return this;
        }
        @CustomType.Setter
        public Builder aggregateCompliancePackName(String aggregateCompliancePackName) {
            this.aggregateCompliancePackName = Objects.requireNonNull(aggregateCompliancePackName);
            return this;
        }
        @CustomType.Setter
        public Builder aggregatorCompliancePackId(String aggregatorCompliancePackId) {
            this.aggregatorCompliancePackId = Objects.requireNonNull(aggregatorCompliancePackId);
            return this;
        }
        @CustomType.Setter
        public Builder compliancePackTemplateId(String compliancePackTemplateId) {
            this.compliancePackTemplateId = Objects.requireNonNull(compliancePackTemplateId);
            return this;
        }
        @CustomType.Setter
        public Builder configRules(List<GetAggregateCompliancePacksPackConfigRule> configRules) {
            this.configRules = Objects.requireNonNull(configRules);
            return this;
        }
        public Builder configRules(GetAggregateCompliancePacksPackConfigRule... configRules) {
            return configRules(List.of(configRules));
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder riskLevel(Integer riskLevel) {
            this.riskLevel = Objects.requireNonNull(riskLevel);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetAggregateCompliancePacksPack build() {
            final var o = new GetAggregateCompliancePacksPack();
            o.accountId = accountId;
            o.aggregateCompliancePackName = aggregateCompliancePackName;
            o.aggregatorCompliancePackId = aggregatorCompliancePackId;
            o.compliancePackTemplateId = compliancePackTemplateId;
            o.configRules = configRules;
            o.description = description;
            o.id = id;
            o.riskLevel = riskLevel;
            o.status = status;
            return o;
        }
    }
}
