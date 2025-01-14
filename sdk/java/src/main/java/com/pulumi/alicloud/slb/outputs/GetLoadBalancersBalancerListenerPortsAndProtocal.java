// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetLoadBalancersBalancerListenerPortsAndProtocal {
    private Integer listenerPort;
    private String listenerProtocal;

    private GetLoadBalancersBalancerListenerPortsAndProtocal() {}
    public Integer listenerPort() {
        return this.listenerPort;
    }
    public String listenerProtocal() {
        return this.listenerProtocal;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLoadBalancersBalancerListenerPortsAndProtocal defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer listenerPort;
        private String listenerProtocal;
        public Builder() {}
        public Builder(GetLoadBalancersBalancerListenerPortsAndProtocal defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.listenerPort = defaults.listenerPort;
    	      this.listenerProtocal = defaults.listenerProtocal;
        }

        @CustomType.Setter
        public Builder listenerPort(Integer listenerPort) {
            this.listenerPort = Objects.requireNonNull(listenerPort);
            return this;
        }
        @CustomType.Setter
        public Builder listenerProtocal(String listenerProtocal) {
            this.listenerProtocal = Objects.requireNonNull(listenerProtocal);
            return this;
        }
        public GetLoadBalancersBalancerListenerPortsAndProtocal build() {
            final var o = new GetLoadBalancersBalancerListenerPortsAndProtocal();
            o.listenerPort = listenerPort;
            o.listenerProtocal = listenerProtocal;
            return o;
        }
    }
}
