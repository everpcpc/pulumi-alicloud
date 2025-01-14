// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetBucketsBucketLogging {
    /**
     * @return Bucket for storing access logs.
     * 
     */
    private String targetBucket;
    /**
     * @return Prefix of the saved access log file paths.
     * 
     */
    private String targetPrefix;

    private GetBucketsBucketLogging() {}
    /**
     * @return Bucket for storing access logs.
     * 
     */
    public String targetBucket() {
        return this.targetBucket;
    }
    /**
     * @return Prefix of the saved access log file paths.
     * 
     */
    public String targetPrefix() {
        return this.targetPrefix;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBucketsBucketLogging defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String targetBucket;
        private String targetPrefix;
        public Builder() {}
        public Builder(GetBucketsBucketLogging defaults) {
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
        public Builder targetPrefix(String targetPrefix) {
            this.targetPrefix = Objects.requireNonNull(targetPrefix);
            return this;
        }
        public GetBucketsBucketLogging build() {
            final var o = new GetBucketsBucketLogging();
            o.targetBucket = targetBucket;
            o.targetPrefix = targetPrefix;
            return o;
        }
    }
}
