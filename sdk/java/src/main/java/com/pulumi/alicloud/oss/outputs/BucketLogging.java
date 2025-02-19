// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BucketLogging {
    /**
     * @return The name of the bucket that will receive the log objects.
     * 
     */
    private String targetBucket;
    /**
     * @return To specify a key prefix for log objects.
     * 
     */
    private @Nullable String targetPrefix;

    private BucketLogging() {}
    /**
     * @return The name of the bucket that will receive the log objects.
     * 
     */
    public String targetBucket() {
        return this.targetBucket;
    }
    /**
     * @return To specify a key prefix for log objects.
     * 
     */
    public Optional<String> targetPrefix() {
        return Optional.ofNullable(this.targetPrefix);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BucketLogging defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String targetBucket;
        private @Nullable String targetPrefix;
        public Builder() {}
        public Builder(BucketLogging defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.targetBucket = defaults.targetBucket;
    	      this.targetPrefix = defaults.targetPrefix;
        }

        @CustomType.Setter
        public Builder targetBucket(String targetBucket) {
            this.targetBucket = Objects.requireNonNull(targetBucket);
            return this;
        }
        @CustomType.Setter
        public Builder targetPrefix(@Nullable String targetPrefix) {
            this.targetPrefix = targetPrefix;
            return this;
        }
        public BucketLogging build() {
            final var o = new BucketLogging();
            o.targetBucket = targetBucket;
            o.targetPrefix = targetPrefix;
            return o;
        }
    }
}
