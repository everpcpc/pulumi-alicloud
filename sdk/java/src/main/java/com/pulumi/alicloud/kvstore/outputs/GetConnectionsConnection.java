// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.kvstore.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetConnectionsConnection {
    /**
     * @return The connection string of the instance.
     * 
     */
    private String connectionString;
    /**
     * @return The network type of the instance.
     * 
     */
    private String dbInstanceNetType;
    /**
     * @return The expiration time of the classic network address.
     * 
     */
    private String expiredTime;
    private String id;
    private String instanceId;
    /**
     * @return The IP address of the instance.
     * 
     */
    private String ipAddress;
    /**
     * @return The port number of the instance.
     * 
     */
    private String port;
    /**
     * @return The remaining validity period of the endpoint of the classic network.
     * 
     */
    private String upgradeable;
    /**
     * @return The ID of the VPC where the instance is deployed.
     * 
     */
    private String vpcId;
    /**
     * @return The ID of the instance. It is returned only when the value of the DBInstanceNetType parameter is 2 (indicating VPC).
     * 
     */
    private String vpcInstanceId;
    /**
     * @return The ID of the VSwitch.
     * 
     */
    private String vswitchId;

    private GetConnectionsConnection() {}
    /**
     * @return The connection string of the instance.
     * 
     */
    public String connectionString() {
        return this.connectionString;
    }
    /**
     * @return The network type of the instance.
     * 
     */
    public String dbInstanceNetType() {
        return this.dbInstanceNetType;
    }
    /**
     * @return The expiration time of the classic network address.
     * 
     */
    public String expiredTime() {
        return this.expiredTime;
    }
    public String id() {
        return this.id;
    }
    public String instanceId() {
        return this.instanceId;
    }
    /**
     * @return The IP address of the instance.
     * 
     */
    public String ipAddress() {
        return this.ipAddress;
    }
    /**
     * @return The port number of the instance.
     * 
     */
    public String port() {
        return this.port;
    }
    /**
     * @return The remaining validity period of the endpoint of the classic network.
     * 
     */
    public String upgradeable() {
        return this.upgradeable;
    }
    /**
     * @return The ID of the VPC where the instance is deployed.
     * 
     */
    public String vpcId() {
        return this.vpcId;
    }
    /**
     * @return The ID of the instance. It is returned only when the value of the DBInstanceNetType parameter is 2 (indicating VPC).
     * 
     */
    public String vpcInstanceId() {
        return this.vpcInstanceId;
    }
    /**
     * @return The ID of the VSwitch.
     * 
     */
    public String vswitchId() {
        return this.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetConnectionsConnection defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String connectionString;
        private String dbInstanceNetType;
        private String expiredTime;
        private String id;
        private String instanceId;
        private String ipAddress;
        private String port;
        private String upgradeable;
        private String vpcId;
        private String vpcInstanceId;
        private String vswitchId;
        public Builder() {}
        public Builder(GetConnectionsConnection defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.connectionString = defaults.connectionString;
    	      this.dbInstanceNetType = defaults.dbInstanceNetType;
    	      this.expiredTime = defaults.expiredTime;
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.ipAddress = defaults.ipAddress;
    	      this.port = defaults.port;
    	      this.upgradeable = defaults.upgradeable;
    	      this.vpcId = defaults.vpcId;
    	      this.vpcInstanceId = defaults.vpcInstanceId;
    	      this.vswitchId = defaults.vswitchId;
        }

        @CustomType.Setter
        public Builder connectionString(String connectionString) {
            this.connectionString = Objects.requireNonNull(connectionString);
            return this;
        }
        @CustomType.Setter
        public Builder dbInstanceNetType(String dbInstanceNetType) {
            this.dbInstanceNetType = Objects.requireNonNull(dbInstanceNetType);
            return this;
        }
        @CustomType.Setter
        public Builder expiredTime(String expiredTime) {
            this.expiredTime = Objects.requireNonNull(expiredTime);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(String instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder ipAddress(String ipAddress) {
            this.ipAddress = Objects.requireNonNull(ipAddress);
            return this;
        }
        @CustomType.Setter
        public Builder port(String port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        @CustomType.Setter
        public Builder upgradeable(String upgradeable) {
            this.upgradeable = Objects.requireNonNull(upgradeable);
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(String vpcId) {
            this.vpcId = Objects.requireNonNull(vpcId);
            return this;
        }
        @CustomType.Setter
        public Builder vpcInstanceId(String vpcInstanceId) {
            this.vpcInstanceId = Objects.requireNonNull(vpcInstanceId);
            return this;
        }
        @CustomType.Setter
        public Builder vswitchId(String vswitchId) {
            this.vswitchId = Objects.requireNonNull(vswitchId);
            return this;
        }
        public GetConnectionsConnection build() {
            final var o = new GetConnectionsConnection();
            o.connectionString = connectionString;
            o.dbInstanceNetType = dbInstanceNetType;
            o.expiredTime = expiredTime;
            o.id = id;
            o.instanceId = instanceId;
            o.ipAddress = ipAddress;
            o.port = port;
            o.upgradeable = upgradeable;
            o.vpcId = vpcId;
            o.vpcInstanceId = vpcInstanceId;
            o.vswitchId = vswitchId;
            return o;
        }
    }
}
