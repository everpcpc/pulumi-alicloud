// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mongodb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ShardingNetworkPrivateAddressNetworkAddress {
    /**
     * @return The remaining duration of the classic network address. Unit: `seconds`.
     * 
     */
    private @Nullable String expiredTime;
    /**
     * @return The IP address of the instance.
     * 
     */
    private @Nullable String ipAddress;
    /**
     * @return The endpoint of the instance.
     * 
     */
    private @Nullable String networkAddress;
    /**
     * @return The network type.
     * 
     */
    private @Nullable String networkType;
    /**
     * @return The ID of the Shard node or the ConfigServer node.
     * 
     */
    private @Nullable String nodeId;
    /**
     * @return The type of the node.
     * 
     */
    private @Nullable String nodeType;
    /**
     * @return The port number.
     * 
     */
    private @Nullable String port;
    /**
     * @return The role of the node.
     * 
     */
    private @Nullable String role;
    /**
     * @return The ID of the VPC.
     * 
     */
    private @Nullable String vpcId;
    /**
     * @return The vSwitch ID of the VPC.
     * 
     */
    private @Nullable String vswitchId;

    private ShardingNetworkPrivateAddressNetworkAddress() {}
    /**
     * @return The remaining duration of the classic network address. Unit: `seconds`.
     * 
     */
    public Optional<String> expiredTime() {
        return Optional.ofNullable(this.expiredTime);
    }
    /**
     * @return The IP address of the instance.
     * 
     */
    public Optional<String> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }
    /**
     * @return The endpoint of the instance.
     * 
     */
    public Optional<String> networkAddress() {
        return Optional.ofNullable(this.networkAddress);
    }
    /**
     * @return The network type.
     * 
     */
    public Optional<String> networkType() {
        return Optional.ofNullable(this.networkType);
    }
    /**
     * @return The ID of the Shard node or the ConfigServer node.
     * 
     */
    public Optional<String> nodeId() {
        return Optional.ofNullable(this.nodeId);
    }
    /**
     * @return The type of the node.
     * 
     */
    public Optional<String> nodeType() {
        return Optional.ofNullable(this.nodeType);
    }
    /**
     * @return The port number.
     * 
     */
    public Optional<String> port() {
        return Optional.ofNullable(this.port);
    }
    /**
     * @return The role of the node.
     * 
     */
    public Optional<String> role() {
        return Optional.ofNullable(this.role);
    }
    /**
     * @return The ID of the VPC.
     * 
     */
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }
    /**
     * @return The vSwitch ID of the VPC.
     * 
     */
    public Optional<String> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ShardingNetworkPrivateAddressNetworkAddress defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String expiredTime;
        private @Nullable String ipAddress;
        private @Nullable String networkAddress;
        private @Nullable String networkType;
        private @Nullable String nodeId;
        private @Nullable String nodeType;
        private @Nullable String port;
        private @Nullable String role;
        private @Nullable String vpcId;
        private @Nullable String vswitchId;
        public Builder() {}
        public Builder(ShardingNetworkPrivateAddressNetworkAddress defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.expiredTime = defaults.expiredTime;
    	      this.ipAddress = defaults.ipAddress;
    	      this.networkAddress = defaults.networkAddress;
    	      this.networkType = defaults.networkType;
    	      this.nodeId = defaults.nodeId;
    	      this.nodeType = defaults.nodeType;
    	      this.port = defaults.port;
    	      this.role = defaults.role;
    	      this.vpcId = defaults.vpcId;
    	      this.vswitchId = defaults.vswitchId;
        }

        @CustomType.Setter
        public Builder expiredTime(@Nullable String expiredTime) {
            this.expiredTime = expiredTime;
            return this;
        }
        @CustomType.Setter
        public Builder ipAddress(@Nullable String ipAddress) {
            this.ipAddress = ipAddress;
            return this;
        }
        @CustomType.Setter
        public Builder networkAddress(@Nullable String networkAddress) {
            this.networkAddress = networkAddress;
            return this;
        }
        @CustomType.Setter
        public Builder networkType(@Nullable String networkType) {
            this.networkType = networkType;
            return this;
        }
        @CustomType.Setter
        public Builder nodeId(@Nullable String nodeId) {
            this.nodeId = nodeId;
            return this;
        }
        @CustomType.Setter
        public Builder nodeType(@Nullable String nodeType) {
            this.nodeType = nodeType;
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable String port) {
            this.port = port;
            return this;
        }
        @CustomType.Setter
        public Builder role(@Nullable String role) {
            this.role = role;
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(@Nullable String vpcId) {
            this.vpcId = vpcId;
            return this;
        }
        @CustomType.Setter
        public Builder vswitchId(@Nullable String vswitchId) {
            this.vswitchId = vswitchId;
            return this;
        }
        public ShardingNetworkPrivateAddressNetworkAddress build() {
            final var o = new ShardingNetworkPrivateAddressNetworkAddress();
            o.expiredTime = expiredTime;
            o.ipAddress = ipAddress;
            o.networkAddress = networkAddress;
            o.networkType = networkType;
            o.nodeId = nodeId;
            o.nodeType = nodeType;
            o.port = port;
            o.role = role;
            o.vpcId = vpcId;
            o.vswitchId = vswitchId;
            return o;
        }
    }
}