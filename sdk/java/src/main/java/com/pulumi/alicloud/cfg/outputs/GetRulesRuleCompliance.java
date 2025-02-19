// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cfg.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRulesRuleCompliance {
    /**
     * @return The compliance evaluation result of the target resources.
     * 
     */
    private String complianceType;
    /**
     * @return The number of resources with the specified compliance evaluation result.
     * 
     */
    private Integer count;

    private GetRulesRuleCompliance() {}
    /**
     * @return The compliance evaluation result of the target resources.
     * 
     */
    public String complianceType() {
        return this.complianceType;
    }
    /**
     * @return The number of resources with the specified compliance evaluation result.
     * 
     */
    public Integer count() {
        return this.count;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRulesRuleCompliance defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String complianceType;
        private Integer count;
        public Builder() {}
        public Builder(GetRulesRuleCompliance defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.complianceType = defaults.complianceType;
    	      this.count = defaults.count;
        }

        @CustomType.Setter
        public Builder complianceType(String complianceType) {
            this.complianceType = Objects.requireNonNull(complianceType);
            return this;
        }
        @CustomType.Setter
        public Builder count(Integer count) {
            this.count = Objects.requireNonNull(count);
            return this;
        }
        public GetRulesRuleCompliance build() {
            final var o = new GetRulesRuleCompliance();
            o.complianceType = complianceType;
            o.count = count;
            return o;
        }
    }
}
