// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GreyTagRouteScRuleItem {
    /**
     * @return The comparison operator. Valid values: `&gt;`, `&lt;`, `&gt;=`, `&lt;=`, `==`, `!=`.
     * 
     */
    private @Nullable String cond;
    /**
     * @return The name of the parameter.
     * 
     */
    private @Nullable String name;
    /**
     * @return The operator. Valid values: `rawvalue`, `list`, `mod`, `deterministic_proportional_steaming_division`
     * 
     */
    private @Nullable String operator;
    /**
     * @return The compare types. Valid values: `param`, `cookie`, `header`.
     * 
     */
    private @Nullable String type;
    /**
     * @return The value of the parameter.
     * 
     */
    private @Nullable String value;

    private GreyTagRouteScRuleItem() {}
    /**
     * @return The comparison operator. Valid values: `&gt;`, `&lt;`, `&gt;=`, `&lt;=`, `==`, `!=`.
     * 
     */
    public Optional<String> cond() {
        return Optional.ofNullable(this.cond);
    }
    /**
     * @return The name of the parameter.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The operator. Valid values: `rawvalue`, `list`, `mod`, `deterministic_proportional_steaming_division`
     * 
     */
    public Optional<String> operator() {
        return Optional.ofNullable(this.operator);
    }
    /**
     * @return The compare types. Valid values: `param`, `cookie`, `header`.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }
    /**
     * @return The value of the parameter.
     * 
     */
    public Optional<String> value() {
        return Optional.ofNullable(this.value);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GreyTagRouteScRuleItem defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String cond;
        private @Nullable String name;
        private @Nullable String operator;
        private @Nullable String type;
        private @Nullable String value;
        public Builder() {}
        public Builder(GreyTagRouteScRuleItem defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cond = defaults.cond;
    	      this.name = defaults.name;
    	      this.operator = defaults.operator;
    	      this.type = defaults.type;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder cond(@Nullable String cond) {
            this.cond = cond;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder operator(@Nullable String operator) {
            this.operator = operator;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        @CustomType.Setter
        public Builder value(@Nullable String value) {
            this.value = value;
            return this;
        }
        public GreyTagRouteScRuleItem build() {
            final var o = new GreyTagRouteScRuleItem();
            o.cond = cond;
            o.name = name;
            o.operator = operator;
            o.type = type;
            o.value = value;
            return o;
        }
    }
}
