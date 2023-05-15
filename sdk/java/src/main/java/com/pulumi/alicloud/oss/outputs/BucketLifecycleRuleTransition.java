// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BucketLifecycleRuleTransition {
    /**
     * @return Specifies the time before which the rules take effect. The date must conform to the ISO8601 format and always be UTC 00:00. For example: 2002-10-11T00:00:00.000Z indicates that objects updated before 2002-10-11T00:00:00.000Z are deleted or converted to another storage class, and objects updated after this time (including this time) are not deleted or converted.
     * 
     */
    private @Nullable String createdBeforeDate;
    /**
     * @return Specifies the number of days after object creation when the specific rule action takes effect.
     * 
     * `NOTE`: One and only one of &#34;created_before_date&#34; and &#34;days&#34; can be specified in one abort_multipart_upload configuration.
     * 
     */
    private @Nullable Integer days;
    /**
     * @return The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be &#34;Standard&#34;, &#34;IA&#34;, &#34;Archive&#34; and &#34;ColdArchive&#34;. Defaults to &#34;Standard&#34;. &#34;ColdArchive&#34; is available in 1.203.0+.
     * 
     */
    private @Nullable String storageClass;

    private BucketLifecycleRuleTransition() {}
    /**
     * @return Specifies the time before which the rules take effect. The date must conform to the ISO8601 format and always be UTC 00:00. For example: 2002-10-11T00:00:00.000Z indicates that objects updated before 2002-10-11T00:00:00.000Z are deleted or converted to another storage class, and objects updated after this time (including this time) are not deleted or converted.
     * 
     */
    public Optional<String> createdBeforeDate() {
        return Optional.ofNullable(this.createdBeforeDate);
    }
    /**
     * @return Specifies the number of days after object creation when the specific rule action takes effect.
     * 
     * `NOTE`: One and only one of &#34;created_before_date&#34; and &#34;days&#34; can be specified in one abort_multipart_upload configuration.
     * 
     */
    public Optional<Integer> days() {
        return Optional.ofNullable(this.days);
    }
    /**
     * @return The [storage class](https://www.alibabacloud.com/help/doc-detail/51374.htm) to apply. Can be &#34;Standard&#34;, &#34;IA&#34;, &#34;Archive&#34; and &#34;ColdArchive&#34;. Defaults to &#34;Standard&#34;. &#34;ColdArchive&#34; is available in 1.203.0+.
     * 
     */
    public Optional<String> storageClass() {
        return Optional.ofNullable(this.storageClass);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BucketLifecycleRuleTransition defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String createdBeforeDate;
        private @Nullable Integer days;
        private @Nullable String storageClass;
        public Builder() {}
        public Builder(BucketLifecycleRuleTransition defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdBeforeDate = defaults.createdBeforeDate;
    	      this.days = defaults.days;
    	      this.storageClass = defaults.storageClass;
        }

        @CustomType.Setter
        public Builder createdBeforeDate(@Nullable String createdBeforeDate) {
            this.createdBeforeDate = createdBeforeDate;
            return this;
        }
        @CustomType.Setter
        public Builder days(@Nullable Integer days) {
            this.days = days;
            return this;
        }
        @CustomType.Setter
        public Builder storageClass(@Nullable String storageClass) {
            this.storageClass = storageClass;
            return this;
        }
        public BucketLifecycleRuleTransition build() {
            final var o = new BucketLifecycleRuleTransition();
            o.createdBeforeDate = createdBeforeDate;
            o.days = days;
            o.storageClass = storageClass;
            return o;
        }
    }
}
