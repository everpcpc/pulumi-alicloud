// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class GetEndpointGroupsGroupPortOverride {
    /**
     * @return Forwarding port.
     * 
     */
    private Integer endpointPort;
    /**
     * @return Listener port.
     * 
     */
    private Integer listenerPort;

    private GetEndpointGroupsGroupPortOverride() {}
    /**
     * @return Forwarding port.
     * 
     */
    public Integer endpointPort() {
        return this.endpointPort;
    }
    /**
     * @return Listener port.
     * 
     */
    public Integer listenerPort() {
        return this.listenerPort;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEndpointGroupsGroupPortOverride defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer endpointPort;
        private Integer listenerPort;
        public Builder() {}
        public Builder(GetEndpointGroupsGroupPortOverride defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.endpointPort = defaults.endpointPort;
    	      this.listenerPort = defaults.listenerPort;
        }

        @CustomType.Setter
        public Builder endpointPort(Integer endpointPort) {
            this.endpointPort = Objects.requireNonNull(endpointPort);
            return this;
        }
        @CustomType.Setter
        public Builder listenerPort(Integer listenerPort) {
            this.listenerPort = Objects.requireNonNull(listenerPort);
            return this;
        }
        public GetEndpointGroupsGroupPortOverride build() {
            final var o = new GetEndpointGroupsGroupPortOverride();
            o.endpointPort = endpointPort;
            o.listenerPort = listenerPort;
            return o;
        }
    }
}
