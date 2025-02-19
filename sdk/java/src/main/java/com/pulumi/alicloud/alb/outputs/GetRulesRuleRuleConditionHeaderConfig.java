// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetRulesRuleRuleConditionHeaderConfig {
    /**
     * @return The key of the header field. The key must be 1 to 40 characters in length, and can contain letters, digits, hyphens (-) and underscores (_). The key does not support Cookie or Host.
     * 
     */
    private String key;
    /**
     * @return Add one or more IP addresses or IP address segments.
     * 
     */
    private List<String> values;

    private GetRulesRuleRuleConditionHeaderConfig() {}
    /**
     * @return The key of the header field. The key must be 1 to 40 characters in length, and can contain letters, digits, hyphens (-) and underscores (_). The key does not support Cookie or Host.
     * 
     */
    public String key() {
        return this.key;
    }
    /**
     * @return Add one or more IP addresses or IP address segments.
     * 
     */
    public List<String> values() {
        return this.values;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRulesRuleRuleConditionHeaderConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String key;
        private List<String> values;
        public Builder() {}
        public Builder(GetRulesRuleRuleConditionHeaderConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.key = defaults.key;
    	      this.values = defaults.values;
        }

        @CustomType.Setter
        public Builder key(String key) {
            this.key = Objects.requireNonNull(key);
            return this;
        }
        @CustomType.Setter
        public Builder values(List<String> values) {
            this.values = Objects.requireNonNull(values);
            return this;
        }
        public Builder values(String... values) {
            return values(List.of(values));
        }
        public GetRulesRuleRuleConditionHeaderConfig build() {
            final var o = new GetRulesRuleRuleConditionHeaderConfig();
            o.key = key;
            o.values = values;
            return o;
        }
    }
}
