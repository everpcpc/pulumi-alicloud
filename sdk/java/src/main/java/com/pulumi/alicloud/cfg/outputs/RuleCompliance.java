// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cfg.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RuleCompliance {
    /**
     * @return The type of compliance. Valid values: `COMPLIANT`, `NON_COMPLIANT`, `NOT_APPLICABLE`, `INSUFFICIENT_DATA`.
     * 
     */
    private @Nullable String complianceType;
    /**
     * @return The count of compliance.
     * 
     */
    private @Nullable Integer count;

    private RuleCompliance() {}
    /**
     * @return The type of compliance. Valid values: `COMPLIANT`, `NON_COMPLIANT`, `NOT_APPLICABLE`, `INSUFFICIENT_DATA`.
     * 
     */
    public Optional<String> complianceType() {
        return Optional.ofNullable(this.complianceType);
    }
    /**
     * @return The count of compliance.
     * 
     */
    public Optional<Integer> count() {
        return Optional.ofNullable(this.count);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RuleCompliance defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String complianceType;
        private @Nullable Integer count;
        public Builder() {}
        public Builder(RuleCompliance defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.complianceType = defaults.complianceType;
    	      this.count = defaults.count;
        }

        @CustomType.Setter
        public Builder complianceType(@Nullable String complianceType) {
            this.complianceType = complianceType;
            return this;
        }
        @CustomType.Setter
        public Builder count(@Nullable Integer count) {
            this.count = count;
            return this;
        }
        public RuleCompliance build() {
            final var o = new RuleCompliance();
            o.complianceType = complianceType;
            o.count = count;
            return o;
        }
    }
}
