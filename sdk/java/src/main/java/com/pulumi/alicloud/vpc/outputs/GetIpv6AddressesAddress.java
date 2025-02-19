// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetIpv6AddressesAddress {
    /**
     * @return The ID of the instance that is assigned the IPv6 address.
     * 
     */
    private String associatedInstanceId;
    /**
     * @return The type of the instance that is assigned the IPv6 address.
     * 
     */
    private String associatedInstanceType;
    /**
     * @return The time when the IPv6 address was created.
     * 
     */
    private String createTime;
    /**
     * @return The ID of the Ipv6 Address.
     * 
     */
    private String id;
    /**
     * @return The address of the Ipv6 Address.
     * 
     */
    private String ipv6Address;
    /**
     * @return The ID of the IPv6 address.
     * 
     */
    private String ipv6AddressId;
    /**
     * @return The name of the IPv6 address.
     * 
     */
    private String ipv6AddressName;
    /**
     * @return The ID of the IPv6 gateway to which the IPv6 address belongs.
     * 
     */
    private String ipv6GatewayId;
    /**
     * @return The type of communication supported by the IPv6 address. Valid values:`Private` or `Public`. `Private`: communication within the private network. `Public`: communication over the public network
     * 
     */
    private String networkType;
    /**
     * @return The status of the IPv6 address. Valid values:`Pending` or `Available`.
     * 
     */
    private String status;
    /**
     * @return The ID of the VPC to which the IPv6 address belongs.
     * 
     */
    private String vpcId;
    /**
     * @return The ID of the vSwitch to which the IPv6 address belongs.
     * 
     */
    private String vswitchId;

    private GetIpv6AddressesAddress() {}
    /**
     * @return The ID of the instance that is assigned the IPv6 address.
     * 
     */
    public String associatedInstanceId() {
        return this.associatedInstanceId;
    }
    /**
     * @return The type of the instance that is assigned the IPv6 address.
     * 
     */
    public String associatedInstanceType() {
        return this.associatedInstanceType;
    }
    /**
     * @return The time when the IPv6 address was created.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The ID of the Ipv6 Address.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The address of the Ipv6 Address.
     * 
     */
    public String ipv6Address() {
        return this.ipv6Address;
    }
    /**
     * @return The ID of the IPv6 address.
     * 
     */
    public String ipv6AddressId() {
        return this.ipv6AddressId;
    }
    /**
     * @return The name of the IPv6 address.
     * 
     */
    public String ipv6AddressName() {
        return this.ipv6AddressName;
    }
    /**
     * @return The ID of the IPv6 gateway to which the IPv6 address belongs.
     * 
     */
    public String ipv6GatewayId() {
        return this.ipv6GatewayId;
    }
    /**
     * @return The type of communication supported by the IPv6 address. Valid values:`Private` or `Public`. `Private`: communication within the private network. `Public`: communication over the public network
     * 
     */
    public String networkType() {
        return this.networkType;
    }
    /**
     * @return The status of the IPv6 address. Valid values:`Pending` or `Available`.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The ID of the VPC to which the IPv6 address belongs.
     * 
     */
    public String vpcId() {
        return this.vpcId;
    }
    /**
     * @return The ID of the vSwitch to which the IPv6 address belongs.
     * 
     */
    public String vswitchId() {
        return this.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIpv6AddressesAddress defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String associatedInstanceId;
        private String associatedInstanceType;
        private String createTime;
        private String id;
        private String ipv6Address;
        private String ipv6AddressId;
        private String ipv6AddressName;
        private String ipv6GatewayId;
        private String networkType;
        private String status;
        private String vpcId;
        private String vswitchId;
        public Builder() {}
        public Builder(GetIpv6AddressesAddress defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.associatedInstanceId = defaults.associatedInstanceId;
    	      this.associatedInstanceType = defaults.associatedInstanceType;
    	      this.createTime = defaults.createTime;
    	      this.id = defaults.id;
    	      this.ipv6Address = defaults.ipv6Address;
    	      this.ipv6AddressId = defaults.ipv6AddressId;
    	      this.ipv6AddressName = defaults.ipv6AddressName;
    	      this.ipv6GatewayId = defaults.ipv6GatewayId;
    	      this.networkType = defaults.networkType;
    	      this.status = defaults.status;
    	      this.vpcId = defaults.vpcId;
    	      this.vswitchId = defaults.vswitchId;
        }

        @CustomType.Setter
        public Builder associatedInstanceId(String associatedInstanceId) {
            this.associatedInstanceId = Objects.requireNonNull(associatedInstanceId);
            return this;
        }
        @CustomType.Setter
        public Builder associatedInstanceType(String associatedInstanceType) {
            this.associatedInstanceType = Objects.requireNonNull(associatedInstanceType);
            return this;
        }
        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ipv6Address(String ipv6Address) {
            this.ipv6Address = Objects.requireNonNull(ipv6Address);
            return this;
        }
        @CustomType.Setter
        public Builder ipv6AddressId(String ipv6AddressId) {
            this.ipv6AddressId = Objects.requireNonNull(ipv6AddressId);
            return this;
        }
        @CustomType.Setter
        public Builder ipv6AddressName(String ipv6AddressName) {
            this.ipv6AddressName = Objects.requireNonNull(ipv6AddressName);
            return this;
        }
        @CustomType.Setter
        public Builder ipv6GatewayId(String ipv6GatewayId) {
            this.ipv6GatewayId = Objects.requireNonNull(ipv6GatewayId);
            return this;
        }
        @CustomType.Setter
        public Builder networkType(String networkType) {
            this.networkType = Objects.requireNonNull(networkType);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(String vpcId) {
            this.vpcId = Objects.requireNonNull(vpcId);
            return this;
        }
        @CustomType.Setter
        public Builder vswitchId(String vswitchId) {
            this.vswitchId = Objects.requireNonNull(vswitchId);
            return this;
        }
        public GetIpv6AddressesAddress build() {
            final var o = new GetIpv6AddressesAddress();
            o.associatedInstanceId = associatedInstanceId;
            o.associatedInstanceType = associatedInstanceType;
            o.createTime = createTime;
            o.id = id;
            o.ipv6Address = ipv6Address;
            o.ipv6AddressId = ipv6AddressId;
            o.ipv6AddressName = ipv6AddressName;
            o.ipv6GatewayId = ipv6GatewayId;
            o.networkType = networkType;
            o.status = status;
            o.vpcId = vpcId;
            o.vswitchId = vswitchId;
            return o;
        }
    }
}
