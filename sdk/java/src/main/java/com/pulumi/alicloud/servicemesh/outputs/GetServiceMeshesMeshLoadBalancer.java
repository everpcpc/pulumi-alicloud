// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.servicemesh.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetServiceMeshesMeshLoadBalancer {
    /**
     * @return The IP address of a public network exposed API Server corresponding to the Load Balance.
     * 
     */
    private String apiServerLoadbalancerId;
    /**
     * @return Whether to use the IP address of a public network exposed the API Server.
     * 
     */
    private Boolean apiServerPublicEip;
    /**
     * @return Whether to use the IP address of a public network exposure the Istio Pilot.
     * 
     */
    private Boolean pilotPublicEip;
    /**
     * @return The IP address of a public network exposure Istio Pilot corresponds to the Load Balance.
     * 
     */
    private String pilotPublicLoadbalancerId;

    private GetServiceMeshesMeshLoadBalancer() {}
    /**
     * @return The IP address of a public network exposed API Server corresponding to the Load Balance.
     * 
     */
    public String apiServerLoadbalancerId() {
        return this.apiServerLoadbalancerId;
    }
    /**
     * @return Whether to use the IP address of a public network exposed the API Server.
     * 
     */
    public Boolean apiServerPublicEip() {
        return this.apiServerPublicEip;
    }
    /**
     * @return Whether to use the IP address of a public network exposure the Istio Pilot.
     * 
     */
    public Boolean pilotPublicEip() {
        return this.pilotPublicEip;
    }
    /**
     * @return The IP address of a public network exposure Istio Pilot corresponds to the Load Balance.
     * 
     */
    public String pilotPublicLoadbalancerId() {
        return this.pilotPublicLoadbalancerId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServiceMeshesMeshLoadBalancer defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String apiServerLoadbalancerId;
        private Boolean apiServerPublicEip;
        private Boolean pilotPublicEip;
        private String pilotPublicLoadbalancerId;
        public Builder() {}
        public Builder(GetServiceMeshesMeshLoadBalancer defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.apiServerLoadbalancerId = defaults.apiServerLoadbalancerId;
    	      this.apiServerPublicEip = defaults.apiServerPublicEip;
    	      this.pilotPublicEip = defaults.pilotPublicEip;
    	      this.pilotPublicLoadbalancerId = defaults.pilotPublicLoadbalancerId;
        }

        @CustomType.Setter
        public Builder apiServerLoadbalancerId(String apiServerLoadbalancerId) {
            this.apiServerLoadbalancerId = Objects.requireNonNull(apiServerLoadbalancerId);
            return this;
        }
        @CustomType.Setter
        public Builder apiServerPublicEip(Boolean apiServerPublicEip) {
            this.apiServerPublicEip = Objects.requireNonNull(apiServerPublicEip);
            return this;
        }
        @CustomType.Setter
        public Builder pilotPublicEip(Boolean pilotPublicEip) {
            this.pilotPublicEip = Objects.requireNonNull(pilotPublicEip);
            return this;
        }
        @CustomType.Setter
        public Builder pilotPublicLoadbalancerId(String pilotPublicLoadbalancerId) {
            this.pilotPublicLoadbalancerId = Objects.requireNonNull(pilotPublicLoadbalancerId);
            return this;
        }
        public GetServiceMeshesMeshLoadBalancer build() {
            final var o = new GetServiceMeshesMeshLoadBalancer();
            o.apiServerLoadbalancerId = apiServerLoadbalancerId;
            o.apiServerPublicEip = apiServerPublicEip;
            o.pilotPublicEip = pilotPublicEip;
            o.pilotPublicLoadbalancerId = pilotPublicLoadbalancerId;
            return o;
        }
    }
}