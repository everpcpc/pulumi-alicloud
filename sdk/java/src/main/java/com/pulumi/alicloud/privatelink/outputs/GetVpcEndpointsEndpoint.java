// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.privatelink.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetVpcEndpointsEndpoint {
    /**
     * @return The Bandwidth.
     * 
     */
    private Integer bandwidth;
    /**
     * @return The status of Connection.
     * 
     */
    private String connectionStatus;
    /**
     * @return The status of Endpoint Business.
     * 
     */
    private String endpointBusinessStatus;
    /**
     * @return The description of Vpc Endpoint.
     * 
     */
    private String endpointDescription;
    /**
     * @return The Endpoint Domain.
     * 
     */
    private String endpointDomain;
    /**
     * @return The ID of the Vpc Endpoint.
     * 
     */
    private String endpointId;
    /**
     * @return The ID of the Vpc Endpoint.
     * 
     */
    private String id;
    /**
     * @return The security group associated with the terminal node network card.
     * 
     */
    private List<String> securityGroupIds;
    /**
     * @return The terminal node service associated with the terminal node.
     * 
     */
    private String serviceId;
    /**
     * @return The name of the terminal node service associated with the terminal node.
     * 
     */
    private String serviceName;
    /**
     * @return The status of Vpc Endpoint.
     * 
     */
    private String status;
    /**
     * @return The name of Vpc Endpoint.
     * 
     */
    private String vpcEndpointName;
    /**
     * @return The private network to which the terminal node belongs.
     * 
     */
    private String vpcId;

    private GetVpcEndpointsEndpoint() {}
    /**
     * @return The Bandwidth.
     * 
     */
    public Integer bandwidth() {
        return this.bandwidth;
    }
    /**
     * @return The status of Connection.
     * 
     */
    public String connectionStatus() {
        return this.connectionStatus;
    }
    /**
     * @return The status of Endpoint Business.
     * 
     */
    public String endpointBusinessStatus() {
        return this.endpointBusinessStatus;
    }
    /**
     * @return The description of Vpc Endpoint.
     * 
     */
    public String endpointDescription() {
        return this.endpointDescription;
    }
    /**
     * @return The Endpoint Domain.
     * 
     */
    public String endpointDomain() {
        return this.endpointDomain;
    }
    /**
     * @return The ID of the Vpc Endpoint.
     * 
     */
    public String endpointId() {
        return this.endpointId;
    }
    /**
     * @return The ID of the Vpc Endpoint.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The security group associated with the terminal node network card.
     * 
     */
    public List<String> securityGroupIds() {
        return this.securityGroupIds;
    }
    /**
     * @return The terminal node service associated with the terminal node.
     * 
     */
    public String serviceId() {
        return this.serviceId;
    }
    /**
     * @return The name of the terminal node service associated with the terminal node.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return The status of Vpc Endpoint.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The name of Vpc Endpoint.
     * 
     */
    public String vpcEndpointName() {
        return this.vpcEndpointName;
    }
    /**
     * @return The private network to which the terminal node belongs.
     * 
     */
    public String vpcId() {
        return this.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVpcEndpointsEndpoint defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer bandwidth;
        private String connectionStatus;
        private String endpointBusinessStatus;
        private String endpointDescription;
        private String endpointDomain;
        private String endpointId;
        private String id;
        private List<String> securityGroupIds;
        private String serviceId;
        private String serviceName;
        private String status;
        private String vpcEndpointName;
        private String vpcId;
        public Builder() {}
        public Builder(GetVpcEndpointsEndpoint defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bandwidth = defaults.bandwidth;
    	      this.connectionStatus = defaults.connectionStatus;
    	      this.endpointBusinessStatus = defaults.endpointBusinessStatus;
    	      this.endpointDescription = defaults.endpointDescription;
    	      this.endpointDomain = defaults.endpointDomain;
    	      this.endpointId = defaults.endpointId;
    	      this.id = defaults.id;
    	      this.securityGroupIds = defaults.securityGroupIds;
    	      this.serviceId = defaults.serviceId;
    	      this.serviceName = defaults.serviceName;
    	      this.status = defaults.status;
    	      this.vpcEndpointName = defaults.vpcEndpointName;
    	      this.vpcId = defaults.vpcId;
        }

        @CustomType.Setter
        public Builder bandwidth(Integer bandwidth) {
            this.bandwidth = Objects.requireNonNull(bandwidth);
            return this;
        }
        @CustomType.Setter
        public Builder connectionStatus(String connectionStatus) {
            this.connectionStatus = Objects.requireNonNull(connectionStatus);
            return this;
        }
        @CustomType.Setter
        public Builder endpointBusinessStatus(String endpointBusinessStatus) {
            this.endpointBusinessStatus = Objects.requireNonNull(endpointBusinessStatus);
            return this;
        }
        @CustomType.Setter
        public Builder endpointDescription(String endpointDescription) {
            this.endpointDescription = Objects.requireNonNull(endpointDescription);
            return this;
        }
        @CustomType.Setter
        public Builder endpointDomain(String endpointDomain) {
            this.endpointDomain = Objects.requireNonNull(endpointDomain);
            return this;
        }
        @CustomType.Setter
        public Builder endpointId(String endpointId) {
            this.endpointId = Objects.requireNonNull(endpointId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder securityGroupIds(List<String> securityGroupIds) {
            this.securityGroupIds = Objects.requireNonNull(securityGroupIds);
            return this;
        }
        public Builder securityGroupIds(String... securityGroupIds) {
            return securityGroupIds(List.of(securityGroupIds));
        }
        @CustomType.Setter
        public Builder serviceId(String serviceId) {
            this.serviceId = Objects.requireNonNull(serviceId);
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            this.serviceName = Objects.requireNonNull(serviceName);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder vpcEndpointName(String vpcEndpointName) {
            this.vpcEndpointName = Objects.requireNonNull(vpcEndpointName);
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(String vpcId) {
            this.vpcId = Objects.requireNonNull(vpcId);
            return this;
        }
        public GetVpcEndpointsEndpoint build() {
            final var o = new GetVpcEndpointsEndpoint();
            o.bandwidth = bandwidth;
            o.connectionStatus = connectionStatus;
            o.endpointBusinessStatus = endpointBusinessStatus;
            o.endpointDescription = endpointDescription;
            o.endpointDomain = endpointDomain;
            o.endpointId = endpointId;
            o.id = id;
            o.securityGroupIds = securityGroupIds;
            o.serviceId = serviceId;
            o.serviceName = serviceName;
            o.status = status;
            o.vpcEndpointName = vpcEndpointName;
            o.vpcId = vpcId;
            return o;
        }
    }
}
