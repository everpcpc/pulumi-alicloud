// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae.outputs;

import com.pulumi.alicloud.sae.outputs.GreyTagRouteDubboRuleItem;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GreyTagRouteDubboRule {
    /**
     * @return The Conditional Patterns for Grayscale Rules. Valid values: `AND`, `OR`.
     * 
     */
    private @Nullable String condition;
    /**
     * @return The service group.
     * 
     */
    private @Nullable String group;
    /**
     * @return A list of conditions items. The details see Block `dubbo_rules_items`.
     * 
     */
    private @Nullable List<GreyTagRouteDubboRuleItem> items;
    /**
     * @return The method name
     * 
     */
    private @Nullable String methodName;
    /**
     * @return The service name.
     * 
     */
    private @Nullable String serviceName;
    /**
     * @return The service version.
     * 
     */
    private @Nullable String version;

    private GreyTagRouteDubboRule() {}
    /**
     * @return The Conditional Patterns for Grayscale Rules. Valid values: `AND`, `OR`.
     * 
     */
    public Optional<String> condition() {
        return Optional.ofNullable(this.condition);
    }
    /**
     * @return The service group.
     * 
     */
    public Optional<String> group() {
        return Optional.ofNullable(this.group);
    }
    /**
     * @return A list of conditions items. The details see Block `dubbo_rules_items`.
     * 
     */
    public List<GreyTagRouteDubboRuleItem> items() {
        return this.items == null ? List.of() : this.items;
    }
    /**
     * @return The method name
     * 
     */
    public Optional<String> methodName() {
        return Optional.ofNullable(this.methodName);
    }
    /**
     * @return The service name.
     * 
     */
    public Optional<String> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }
    /**
     * @return The service version.
     * 
     */
    public Optional<String> version() {
        return Optional.ofNullable(this.version);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GreyTagRouteDubboRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String condition;
        private @Nullable String group;
        private @Nullable List<GreyTagRouteDubboRuleItem> items;
        private @Nullable String methodName;
        private @Nullable String serviceName;
        private @Nullable String version;
        public Builder() {}
        public Builder(GreyTagRouteDubboRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.condition = defaults.condition;
    	      this.group = defaults.group;
    	      this.items = defaults.items;
    	      this.methodName = defaults.methodName;
    	      this.serviceName = defaults.serviceName;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder condition(@Nullable String condition) {
            this.condition = condition;
            return this;
        }
        @CustomType.Setter
        public Builder group(@Nullable String group) {
            this.group = group;
            return this;
        }
        @CustomType.Setter
        public Builder items(@Nullable List<GreyTagRouteDubboRuleItem> items) {
            this.items = items;
            return this;
        }
        public Builder items(GreyTagRouteDubboRuleItem... items) {
            return items(List.of(items));
        }
        @CustomType.Setter
        public Builder methodName(@Nullable String methodName) {
            this.methodName = methodName;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(@Nullable String serviceName) {
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder version(@Nullable String version) {
            this.version = version;
            return this;
        }
        public GreyTagRouteDubboRule build() {
            final var o = new GreyTagRouteDubboRule();
            o.condition = condition;
            o.group = group;
            o.items = items;
            o.methodName = methodName;
            o.serviceName = serviceName;
            o.version = version;
            return o;
        }
    }
}
