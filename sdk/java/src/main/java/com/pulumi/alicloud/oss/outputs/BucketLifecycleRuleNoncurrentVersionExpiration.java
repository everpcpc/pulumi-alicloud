// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class BucketLifecycleRuleNoncurrentVersionExpiration {
    /**
     * @return Specifies the number of days after object creation when the specific rule action takes effect.
     * 
     * `NOTE`: One and only one of &#34;created_before_date&#34; and &#34;days&#34; can be specified in one abort_multipart_upload configuration.
     * 
     */
    private Integer days;

    private BucketLifecycleRuleNoncurrentVersionExpiration() {}
    /**
     * @return Specifies the number of days after object creation when the specific rule action takes effect.
     * 
     * `NOTE`: One and only one of &#34;created_before_date&#34; and &#34;days&#34; can be specified in one abort_multipart_upload configuration.
     * 
     */
    public Integer days() {
        return this.days;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BucketLifecycleRuleNoncurrentVersionExpiration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer days;
        public Builder() {}
        public Builder(BucketLifecycleRuleNoncurrentVersionExpiration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.days = defaults.days;
        }

        @CustomType.Setter
        public Builder days(Integer days) {
            this.days = Objects.requireNonNull(days);
            return this;
        }
        public BucketLifecycleRuleNoncurrentVersionExpiration build() {
            final var o = new BucketLifecycleRuleNoncurrentVersionExpiration();
            o.days = days;
            return o;
        }
    }
}
