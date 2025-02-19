// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae.outputs;

import com.pulumi.alicloud.sae.outputs.GetGreyTagRoutesRouteDubboRuleItem;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetGreyTagRoutesRouteDubboRule {
    /**
     * @return The Conditional Patterns for Grayscale Rules.
     * 
     */
    private String condition;
    /**
     * @return The service group.
     * 
     */
    private String group;
    /**
     * @return A list of conditions items.
     * 
     */
    private List<GetGreyTagRoutesRouteDubboRuleItem> items;
    /**
     * @return The method name
     * 
     */
    private String methodName;
    /**
     * @return The service name.
     * 
     */
    private String serviceName;
    /**
     * @return The service version.
     * 
     */
    private String version;

    private GetGreyTagRoutesRouteDubboRule() {}
    /**
     * @return The Conditional Patterns for Grayscale Rules.
     * 
     */
    public String condition() {
        return this.condition;
    }
    /**
     * @return The service group.
     * 
     */
    public String group() {
        return this.group;
    }
    /**
     * @return A list of conditions items.
     * 
     */
    public List<GetGreyTagRoutesRouteDubboRuleItem> items() {
        return this.items;
    }
    /**
     * @return The method name
     * 
     */
    public String methodName() {
        return this.methodName;
    }
    /**
     * @return The service name.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return The service version.
     * 
     */
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGreyTagRoutesRouteDubboRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String condition;
        private String group;
        private List<GetGreyTagRoutesRouteDubboRuleItem> items;
        private String methodName;
        private String serviceName;
        private String version;
        public Builder() {}
        public Builder(GetGreyTagRoutesRouteDubboRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.condition = defaults.condition;
    	      this.group = defaults.group;
    	      this.items = defaults.items;
    	      this.methodName = defaults.methodName;
    	      this.serviceName = defaults.serviceName;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder condition(String condition) {
            this.condition = Objects.requireNonNull(condition);
            return this;
        }
        @CustomType.Setter
        public Builder group(String group) {
            this.group = Objects.requireNonNull(group);
            return this;
        }
        @CustomType.Setter
        public Builder items(List<GetGreyTagRoutesRouteDubboRuleItem> items) {
            this.items = Objects.requireNonNull(items);
            return this;
        }
        public Builder items(GetGreyTagRoutesRouteDubboRuleItem... items) {
            return items(List.of(items));
        }
        @CustomType.Setter
        public Builder methodName(String methodName) {
            this.methodName = Objects.requireNonNull(methodName);
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            this.serviceName = Objects.requireNonNull(serviceName);
            return this;
        }
        @CustomType.Setter
        public Builder version(String version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }
        public GetGreyTagRoutesRouteDubboRule build() {
            final var o = new GetGreyTagRoutesRouteDubboRule();
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
