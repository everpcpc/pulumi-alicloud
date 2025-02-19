// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class GetCustomRoutingEndpointTrafficPoliciesCustomRoutingEndpointTrafficPolicyPortRange {
    /**
     * @return The first port of the port range.
     * 
     */
    private Integer fromPort;
    /**
     * @return The last port of the port range.
     * 
     */
    private Integer toPort;

    private GetCustomRoutingEndpointTrafficPoliciesCustomRoutingEndpointTrafficPolicyPortRange() {}
    /**
     * @return The first port of the port range.
     * 
     */
    public Integer fromPort() {
        return this.fromPort;
    }
    /**
     * @return The last port of the port range.
     * 
     */
    public Integer toPort() {
        return this.toPort;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCustomRoutingEndpointTrafficPoliciesCustomRoutingEndpointTrafficPolicyPortRange defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer fromPort;
        private Integer toPort;
        public Builder() {}
        public Builder(GetCustomRoutingEndpointTrafficPoliciesCustomRoutingEndpointTrafficPolicyPortRange defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fromPort = defaults.fromPort;
    	      this.toPort = defaults.toPort;
        }

        @CustomType.Setter
        public Builder fromPort(Integer fromPort) {
            this.fromPort = Objects.requireNonNull(fromPort);
            return this;
        }
        @CustomType.Setter
        public Builder toPort(Integer toPort) {
            this.toPort = Objects.requireNonNull(toPort);
            return this;
        }
        public GetCustomRoutingEndpointTrafficPoliciesCustomRoutingEndpointTrafficPolicyPortRange build() {
            final var o = new GetCustomRoutingEndpointTrafficPoliciesCustomRoutingEndpointTrafficPolicyPortRange();
            o.fromPort = fromPort;
            o.toPort = toPort;
            return o;
        }
    }
}
