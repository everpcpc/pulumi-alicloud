// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eci.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ContainerGroupContainerLivenessProbeTcpSocket {
    /**
     * @return The port number. Valid values: 1 to 65535.
     * 
     */
    private @Nullable Integer port;

    private ContainerGroupContainerLivenessProbeTcpSocket() {}
    /**
     * @return The port number. Valid values: 1 to 65535.
     * 
     */
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContainerGroupContainerLivenessProbeTcpSocket defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer port;
        public Builder() {}
        public Builder(ContainerGroupContainerLivenessProbeTcpSocket defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.port = defaults.port;
        }

        @CustomType.Setter
        public Builder port(@Nullable Integer port) {
            this.port = port;
            return this;
        }
        public ContainerGroupContainerLivenessProbeTcpSocket build() {
            final var o = new ContainerGroupContainerLivenessProbeTcpSocket();
            o.port = port;
            return o;
        }
    }
}