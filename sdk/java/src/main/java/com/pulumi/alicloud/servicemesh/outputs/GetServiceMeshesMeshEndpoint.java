// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.servicemesh.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetServiceMeshesMeshEndpoint {
    /**
     * @return The internal address of the API Server.
     * 
     */
    private String intranetApiServerEndpoint;
    /**
     * @return The internal address of the Istio Pilot.
     * 
     */
    private String intranetPilotEndpoint;
    /**
     * @return The public address of the API Server.
     * 
     */
    private String publicApiServerEndpoint;
    /**
     * @return The public address of the Istio Pilot.
     * 
     */
    private String publicPilotEndpoint;

    private GetServiceMeshesMeshEndpoint() {}
    /**
     * @return The internal address of the API Server.
     * 
     */
    public String intranetApiServerEndpoint() {
        return this.intranetApiServerEndpoint;
    }
    /**
     * @return The internal address of the Istio Pilot.
     * 
     */
    public String intranetPilotEndpoint() {
        return this.intranetPilotEndpoint;
    }
    /**
     * @return The public address of the API Server.
     * 
     */
    public String publicApiServerEndpoint() {
        return this.publicApiServerEndpoint;
    }
    /**
     * @return The public address of the Istio Pilot.
     * 
     */
    public String publicPilotEndpoint() {
        return this.publicPilotEndpoint;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServiceMeshesMeshEndpoint defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String intranetApiServerEndpoint;
        private String intranetPilotEndpoint;
        private String publicApiServerEndpoint;
        private String publicPilotEndpoint;
        public Builder() {}
        public Builder(GetServiceMeshesMeshEndpoint defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.intranetApiServerEndpoint = defaults.intranetApiServerEndpoint;
    	      this.intranetPilotEndpoint = defaults.intranetPilotEndpoint;
    	      this.publicApiServerEndpoint = defaults.publicApiServerEndpoint;
    	      this.publicPilotEndpoint = defaults.publicPilotEndpoint;
        }

        @CustomType.Setter
        public Builder intranetApiServerEndpoint(String intranetApiServerEndpoint) {
            this.intranetApiServerEndpoint = Objects.requireNonNull(intranetApiServerEndpoint);
            return this;
        }
        @CustomType.Setter
        public Builder intranetPilotEndpoint(String intranetPilotEndpoint) {
            this.intranetPilotEndpoint = Objects.requireNonNull(intranetPilotEndpoint);
            return this;
        }
        @CustomType.Setter
        public Builder publicApiServerEndpoint(String publicApiServerEndpoint) {
            this.publicApiServerEndpoint = Objects.requireNonNull(publicApiServerEndpoint);
            return this;
        }
        @CustomType.Setter
        public Builder publicPilotEndpoint(String publicPilotEndpoint) {
            this.publicPilotEndpoint = Objects.requireNonNull(publicPilotEndpoint);
            return this;
        }
        public GetServiceMeshesMeshEndpoint build() {
            final var o = new GetServiceMeshesMeshEndpoint();
            o.intranetApiServerEndpoint = intranetApiServerEndpoint;
            o.intranetPilotEndpoint = intranetPilotEndpoint;
            o.publicApiServerEndpoint = publicApiServerEndpoint;
            o.publicPilotEndpoint = publicPilotEndpoint;
            return o;
        }
    }
}
