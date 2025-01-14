// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mse.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetClustersClusterInstanceModel {
    /**
     * @return The health status of MSE Cluster.
     * 
     */
    private String healthStatus;
    private String instanceType;
    private String internetIp;
    private String ip;
    private String podName;
    private String role;
    private String singleTunnelVip;
    private String vip;

    private GetClustersClusterInstanceModel() {}
    /**
     * @return The health status of MSE Cluster.
     * 
     */
    public String healthStatus() {
        return this.healthStatus;
    }
    public String instanceType() {
        return this.instanceType;
    }
    public String internetIp() {
        return this.internetIp;
    }
    public String ip() {
        return this.ip;
    }
    public String podName() {
        return this.podName;
    }
    public String role() {
        return this.role;
    }
    public String singleTunnelVip() {
        return this.singleTunnelVip;
    }
    public String vip() {
        return this.vip;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClustersClusterInstanceModel defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String healthStatus;
        private String instanceType;
        private String internetIp;
        private String ip;
        private String podName;
        private String role;
        private String singleTunnelVip;
        private String vip;
        public Builder() {}
        public Builder(GetClustersClusterInstanceModel defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.healthStatus = defaults.healthStatus;
    	      this.instanceType = defaults.instanceType;
    	      this.internetIp = defaults.internetIp;
    	      this.ip = defaults.ip;
    	      this.podName = defaults.podName;
    	      this.role = defaults.role;
    	      this.singleTunnelVip = defaults.singleTunnelVip;
    	      this.vip = defaults.vip;
        }

        @CustomType.Setter
        public Builder healthStatus(String healthStatus) {
            this.healthStatus = Objects.requireNonNull(healthStatus);
            return this;
        }
        @CustomType.Setter
        public Builder instanceType(String instanceType) {
            this.instanceType = Objects.requireNonNull(instanceType);
            return this;
        }
        @CustomType.Setter
        public Builder internetIp(String internetIp) {
            this.internetIp = Objects.requireNonNull(internetIp);
            return this;
        }
        @CustomType.Setter
        public Builder ip(String ip) {
            this.ip = Objects.requireNonNull(ip);
            return this;
        }
        @CustomType.Setter
        public Builder podName(String podName) {
            this.podName = Objects.requireNonNull(podName);
            return this;
        }
        @CustomType.Setter
        public Builder role(String role) {
            this.role = Objects.requireNonNull(role);
            return this;
        }
        @CustomType.Setter
        public Builder singleTunnelVip(String singleTunnelVip) {
            this.singleTunnelVip = Objects.requireNonNull(singleTunnelVip);
            return this;
        }
        @CustomType.Setter
        public Builder vip(String vip) {
            this.vip = Objects.requireNonNull(vip);
            return this;
        }
        public GetClustersClusterInstanceModel build() {
            final var o = new GetClustersClusterInstanceModel();
            o.healthStatus = healthStatus;
            o.instanceType = instanceType;
            o.internetIp = internetIp;
            o.ip = ip;
            o.podName = podName;
            o.role = role;
            o.singleTunnelVip = singleTunnelVip;
            o.vip = vip;
            return o;
        }
    }
}
