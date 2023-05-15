// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class EciScalingConfigurationInitContainerPort {
    /**
     * @return The port number. Valid values: 1 to 65535.
     * 
     */
    private @Nullable Integer port;
    /**
     * @return Valid values: TCP and UDP.
     * 
     */
    private @Nullable String protocol;

    private EciScalingConfigurationInitContainerPort() {}
    /**
     * @return The port number. Valid values: 1 to 65535.
     * 
     */
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }
    /**
     * @return Valid values: TCP and UDP.
     * 
     */
    public Optional<String> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(EciScalingConfigurationInitContainerPort defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer port;
        private @Nullable String protocol;
        public Builder() {}
        public Builder(EciScalingConfigurationInitContainerPort defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.port = defaults.port;
    	      this.protocol = defaults.protocol;
        }

        @CustomType.Setter
        public Builder port(@Nullable Integer port) {
            this.port = port;
            return this;
        }
        @CustomType.Setter
        public Builder protocol(@Nullable String protocol) {
            this.protocol = protocol;
            return this;
        }
        public EciScalingConfigurationInitContainerPort build() {
            final var o = new EciScalingConfigurationInitContainerPort();
            o.port = port;
            o.protocol = protocol;
            return o;
        }
    }
}
