// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetBucketsBucketWebsite {
    /**
     * @return Key of the HTML document containing the error page.
     * 
     */
    private String errorDocument;
    /**
     * @return Key of the HTML document containing the home page.
     * 
     */
    private String indexDocument;

    private GetBucketsBucketWebsite() {}
    /**
     * @return Key of the HTML document containing the error page.
     * 
     */
    public String errorDocument() {
        return this.errorDocument;
    }
    /**
     * @return Key of the HTML document containing the home page.
     * 
     */
    public String indexDocument() {
        return this.indexDocument;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBucketsBucketWebsite defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String errorDocument;
        private String indexDocument;
        public Builder() {}
        public Builder(GetBucketsBucketWebsite defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.errorDocument = defaults.errorDocument;
    	      this.indexDocument = defaults.indexDocument;
        }

        @CustomType.Setter
        public Builder errorDocument(String errorDocument) {
            this.errorDocument = Objects.requireNonNull(errorDocument);
            return this;
        }
        @CustomType.Setter
        public Builder indexDocument(String indexDocument) {
            this.indexDocument = Objects.requireNonNull(indexDocument);
            return this;
        }
        public GetBucketsBucketWebsite build() {
            final var o = new GetBucketsBucketWebsite();
            o.errorDocument = errorDocument;
            o.indexDocument = indexDocument;
            return o;
        }
    }
}
