// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ListenerCertificates {
    /**
     * @return The ID of the Certificate.
     * 
     */
    private @Nullable String certificateId;

    private ListenerCertificates() {}
    /**
     * @return The ID of the Certificate.
     * 
     */
    public Optional<String> certificateId() {
        return Optional.ofNullable(this.certificateId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ListenerCertificates defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String certificateId;
        public Builder() {}
        public Builder(ListenerCertificates defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.certificateId = defaults.certificateId;
        }

        @CustomType.Setter
        public Builder certificateId(@Nullable String certificateId) {
            this.certificateId = certificateId;
            return this;
        }
        public ListenerCertificates build() {
            final var o = new ListenerCertificates();
            o.certificateId = certificateId;
            return o;
        }
    }
}
