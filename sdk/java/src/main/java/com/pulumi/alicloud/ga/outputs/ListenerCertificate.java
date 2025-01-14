// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ListenerCertificate {
    /**
     * @return The id of the certificate.
     * 
     */
    private @Nullable String id;

    private ListenerCertificate() {}
    /**
     * @return The id of the certificate.
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ListenerCertificate defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String id;
        public Builder() {}
        public Builder(ListenerCertificate defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
        }

        @CustomType.Setter
        public Builder id(@Nullable String id) {
            this.id = id;
            return this;
        }
        public ListenerCertificate build() {
            final var o = new ListenerCertificate();
            o.id = id;
            return o;
        }
    }
}
