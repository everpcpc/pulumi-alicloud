// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetApplicationLoadBalancersBalancerListenerPortsAndProtocol {
    /**
     * @return The description of protocol.
     * 
     */
    private String description;
    /**
     * @return The forward port.
     * 
     */
    private Integer forwardPort;
    /**
     * @return The listener forward.
     * 
     */
    private String listenerForward;
    /**
     * @return The listener port.
     * 
     */
    private Integer listenerPort;
    /**
     * @return The listener protocol.
     * 
     */
    private String listenerProtocol;

    private GetApplicationLoadBalancersBalancerListenerPortsAndProtocol() {}
    /**
     * @return The description of protocol.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The forward port.
     * 
     */
    public Integer forwardPort() {
        return this.forwardPort;
    }
    /**
     * @return The listener forward.
     * 
     */
    public String listenerForward() {
        return this.listenerForward;
    }
    /**
     * @return The listener port.
     * 
     */
    public Integer listenerPort() {
        return this.listenerPort;
    }
    /**
     * @return The listener protocol.
     * 
     */
    public String listenerProtocol() {
        return this.listenerProtocol;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetApplicationLoadBalancersBalancerListenerPortsAndProtocol defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private Integer forwardPort;
        private String listenerForward;
        private Integer listenerPort;
        private String listenerProtocol;
        public Builder() {}
        public Builder(GetApplicationLoadBalancersBalancerListenerPortsAndProtocol defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.forwardPort = defaults.forwardPort;
    	      this.listenerForward = defaults.listenerForward;
    	      this.listenerPort = defaults.listenerPort;
    	      this.listenerProtocol = defaults.listenerProtocol;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder forwardPort(Integer forwardPort) {
            this.forwardPort = Objects.requireNonNull(forwardPort);
            return this;
        }
        @CustomType.Setter
        public Builder listenerForward(String listenerForward) {
            this.listenerForward = Objects.requireNonNull(listenerForward);
            return this;
        }
        @CustomType.Setter
        public Builder listenerPort(Integer listenerPort) {
            this.listenerPort = Objects.requireNonNull(listenerPort);
            return this;
        }
        @CustomType.Setter
        public Builder listenerProtocol(String listenerProtocol) {
            this.listenerProtocol = Objects.requireNonNull(listenerProtocol);
            return this;
        }
        public GetApplicationLoadBalancersBalancerListenerPortsAndProtocol build() {
            final var o = new GetApplicationLoadBalancersBalancerListenerPortsAndProtocol();
            o.description = description;
            o.forwardPort = forwardPort;
            o.listenerForward = listenerForward;
            o.listenerPort = listenerPort;
            o.listenerProtocol = listenerProtocol;
            return o;
        }
    }
}