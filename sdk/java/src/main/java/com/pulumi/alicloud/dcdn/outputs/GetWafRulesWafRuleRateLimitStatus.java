// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dcdn.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetWafRulesWafRuleRateLimitStatus {
    private String code;
    /**
     * @return The number of times that the HTTP status code that was returned.
     * 
     */
    private Integer count;
    /**
     * @return The percentage of HTTP status codes.
     * 
     */
    private Integer ratio;

    private GetWafRulesWafRuleRateLimitStatus() {}
    public String code() {
        return this.code;
    }
    /**
     * @return The number of times that the HTTP status code that was returned.
     * 
     */
    public Integer count() {
        return this.count;
    }
    /**
     * @return The percentage of HTTP status codes.
     * 
     */
    public Integer ratio() {
        return this.ratio;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetWafRulesWafRuleRateLimitStatus defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String code;
        private Integer count;
        private Integer ratio;
        public Builder() {}
        public Builder(GetWafRulesWafRuleRateLimitStatus defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.code = defaults.code;
    	      this.count = defaults.count;
    	      this.ratio = defaults.ratio;
        }

        @CustomType.Setter
        public Builder code(String code) {
            this.code = Objects.requireNonNull(code);
            return this;
        }
        @CustomType.Setter
        public Builder count(Integer count) {
            this.count = Objects.requireNonNull(count);
            return this;
        }
        @CustomType.Setter
        public Builder ratio(Integer ratio) {
            this.ratio = Objects.requireNonNull(ratio);
            return this;
        }
        public GetWafRulesWafRuleRateLimitStatus build() {
            final var o = new GetWafRulesWafRuleRateLimitStatus();
            o.code = code;
            o.count = count;
            o.ratio = ratio;
            return o;
        }
    }
}
