// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetBucketsBucketVersioning {
    /**
     * @return A bucket versioning state. Possible values:`Enabled` and `Suspended`.
     * 
     */
    private String status;

    private GetBucketsBucketVersioning() {}
    /**
     * @return A bucket versioning state. Possible values:`Enabled` and `Suspended`.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBucketsBucketVersioning defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String status;
        public Builder() {}
        public Builder(GetBucketsBucketVersioning defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetBucketsBucketVersioning build() {
            final var o = new GetBucketsBucketVersioning();
            o.status = status;
            return o;
        }
    }
}
