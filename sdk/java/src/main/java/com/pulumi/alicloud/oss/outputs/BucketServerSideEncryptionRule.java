// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BucketServerSideEncryptionRule {
    /**
     * @return The alibaba cloud KMS master key ID used for the SSE-KMS encryption.
     * 
     */
    private @Nullable String kmsMasterKeyId;
    /**
     * @return The server-side encryption algorithm to use. Possible values: `AES256` and `KMS`.
     * 
     */
    private String sseAlgorithm;

    private BucketServerSideEncryptionRule() {}
    /**
     * @return The alibaba cloud KMS master key ID used for the SSE-KMS encryption.
     * 
     */
    public Optional<String> kmsMasterKeyId() {
        return Optional.ofNullable(this.kmsMasterKeyId);
    }
    /**
     * @return The server-side encryption algorithm to use. Possible values: `AES256` and `KMS`.
     * 
     */
    public String sseAlgorithm() {
        return this.sseAlgorithm;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BucketServerSideEncryptionRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String kmsMasterKeyId;
        private String sseAlgorithm;
        public Builder() {}
        public Builder(BucketServerSideEncryptionRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.kmsMasterKeyId = defaults.kmsMasterKeyId;
    	      this.sseAlgorithm = defaults.sseAlgorithm;
        }

        @CustomType.Setter
        public Builder kmsMasterKeyId(@Nullable String kmsMasterKeyId) {
            this.kmsMasterKeyId = kmsMasterKeyId;
            return this;
        }
        @CustomType.Setter
        public Builder sseAlgorithm(String sseAlgorithm) {
            this.sseAlgorithm = Objects.requireNonNull(sseAlgorithm);
            return this;
        }
        public BucketServerSideEncryptionRule build() {
            final var o = new BucketServerSideEncryptionRule();
            o.kmsMasterKeyId = kmsMasterKeyId;
            o.sseAlgorithm = sseAlgorithm;
            return o;
        }
    }
}
