// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class ForwardingRuleRuleConditionPathConfig {
    /**
     * @return The length of the path is 1-128 characters. It must start with a forward slash (/), and can only contain letters, numbers, dollar sign ($), dash (-), and underscores (_) , half width full stop (.), plus sign (+), forward slash (/), and (&amp;), wavy line (~), at (@), half width colon (:), apostrophe (&#39;). It supports asterisk (*) and half width question mark (?) as wildcards.
     * 
     */
    private @Nullable List<String> values;

    private ForwardingRuleRuleConditionPathConfig() {}
    /**
     * @return The length of the path is 1-128 characters. It must start with a forward slash (/), and can only contain letters, numbers, dollar sign ($), dash (-), and underscores (_) , half width full stop (.), plus sign (+), forward slash (/), and (&amp;), wavy line (~), at (@), half width colon (:), apostrophe (&#39;). It supports asterisk (*) and half width question mark (?) as wildcards.
     * 
     */
    public List<String> values() {
        return this.values == null ? List.of() : this.values;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ForwardingRuleRuleConditionPathConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> values;
        public Builder() {}
        public Builder(ForwardingRuleRuleConditionPathConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.values = defaults.values;
        }

        @CustomType.Setter
        public Builder values(@Nullable List<String> values) {
            this.values = values;
            return this;
        }
        public Builder values(String... values) {
            return values(List.of(values));
        }
        public ForwardingRuleRuleConditionPathConfig build() {
            final var o = new ForwardingRuleRuleConditionPathConfig();
            o.values = values;
            return o;
        }
    }
}
