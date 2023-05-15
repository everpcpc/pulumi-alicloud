// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetIpv6GatewaysGateway {
    /**
     * @return The status of the IPv6 gateway. Valid values:`Normal`, `FinancialLocked` and `SecurityLocked`. `Normal`: working as expected. `FinancialLocked`: locked due to overdue payments. `SecurityLocked`: locked due to security reasons.
     * 
     */
    private String businessStatus;
    /**
     * @return The creation time of the resource.
     * 
     */
    private String createTime;
    /**
     * @return The description of the IPv6 gateway.
     * 
     */
    private String description;
    /**
     * @return The time when the IPv6 gateway expires.
     * 
     */
    private String expiredTime;
    /**
     * @return The ID of the Ipv6 Gateway.
     * 
     */
    private String id;
    /**
     * @return The metering method of the IPv6 gateway. Valid values: `PayAsYouGo`.
     * 
     */
    private String instanceChargeType;
    /**
     * @return The first ID of the resource.
     * 
     */
    private String ipv6GatewayId;
    /**
     * @return The name of the IPv6 gateway.
     * 
     */
    private String ipv6GatewayName;
    /**
     * @return The specification of the IPv6 gateway. Valid values: `Large`, `Medium` and `Small`. `Small` (default): Free Edition. `Medium`: Enterprise Edition . `Large`: Enhanced Enterprise Edition. The throughput capacity of an IPv6 gateway varies based on the edition. For more information, see [Editions of IPv6 gateways](https://www.alibabacloud.com/help/doc-detail/98926.htm).
     * 
     */
    private String spec;
    /**
     * @return The status of the IPv6 gateway. Valid values: `Available`, `Deleting`, `Pending`.
     * 
     */
    private String status;
    /**
     * @return The ID of the virtual private cloud (VPC) to which the IPv6 gateway belongs.
     * 
     */
    private String vpcId;

    private GetIpv6GatewaysGateway() {}
    /**
     * @return The status of the IPv6 gateway. Valid values:`Normal`, `FinancialLocked` and `SecurityLocked`. `Normal`: working as expected. `FinancialLocked`: locked due to overdue payments. `SecurityLocked`: locked due to security reasons.
     * 
     */
    public String businessStatus() {
        return this.businessStatus;
    }
    /**
     * @return The creation time of the resource.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The description of the IPv6 gateway.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The time when the IPv6 gateway expires.
     * 
     */
    public String expiredTime() {
        return this.expiredTime;
    }
    /**
     * @return The ID of the Ipv6 Gateway.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The metering method of the IPv6 gateway. Valid values: `PayAsYouGo`.
     * 
     */
    public String instanceChargeType() {
        return this.instanceChargeType;
    }
    /**
     * @return The first ID of the resource.
     * 
     */
    public String ipv6GatewayId() {
        return this.ipv6GatewayId;
    }
    /**
     * @return The name of the IPv6 gateway.
     * 
     */
    public String ipv6GatewayName() {
        return this.ipv6GatewayName;
    }
    /**
     * @return The specification of the IPv6 gateway. Valid values: `Large`, `Medium` and `Small`. `Small` (default): Free Edition. `Medium`: Enterprise Edition . `Large`: Enhanced Enterprise Edition. The throughput capacity of an IPv6 gateway varies based on the edition. For more information, see [Editions of IPv6 gateways](https://www.alibabacloud.com/help/doc-detail/98926.htm).
     * 
     */
    public String spec() {
        return this.spec;
    }
    /**
     * @return The status of the IPv6 gateway. Valid values: `Available`, `Deleting`, `Pending`.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The ID of the virtual private cloud (VPC) to which the IPv6 gateway belongs.
     * 
     */
    public String vpcId() {
        return this.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIpv6GatewaysGateway defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String businessStatus;
        private String createTime;
        private String description;
        private String expiredTime;
        private String id;
        private String instanceChargeType;
        private String ipv6GatewayId;
        private String ipv6GatewayName;
        private String spec;
        private String status;
        private String vpcId;
        public Builder() {}
        public Builder(GetIpv6GatewaysGateway defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.businessStatus = defaults.businessStatus;
    	      this.createTime = defaults.createTime;
    	      this.description = defaults.description;
    	      this.expiredTime = defaults.expiredTime;
    	      this.id = defaults.id;
    	      this.instanceChargeType = defaults.instanceChargeType;
    	      this.ipv6GatewayId = defaults.ipv6GatewayId;
    	      this.ipv6GatewayName = defaults.ipv6GatewayName;
    	      this.spec = defaults.spec;
    	      this.status = defaults.status;
    	      this.vpcId = defaults.vpcId;
        }

        @CustomType.Setter
        public Builder businessStatus(String businessStatus) {
            this.businessStatus = Objects.requireNonNull(businessStatus);
            return this;
        }
        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
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
        public Builder instanceChargeType(String instanceChargeType) {
            this.instanceChargeType = Objects.requireNonNull(instanceChargeType);
            return this;
        }
        @CustomType.Setter
        public Builder ipv6GatewayId(String ipv6GatewayId) {
            this.ipv6GatewayId = Objects.requireNonNull(ipv6GatewayId);
            return this;
        }
        @CustomType.Setter
        public Builder ipv6GatewayName(String ipv6GatewayName) {
            this.ipv6GatewayName = Objects.requireNonNull(ipv6GatewayName);
            return this;
        }
        @CustomType.Setter
        public Builder spec(String spec) {
            this.spec = Objects.requireNonNull(spec);
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
        public GetIpv6GatewaysGateway build() {
            final var o = new GetIpv6GatewaysGateway();
            o.businessStatus = businessStatus;
            o.createTime = createTime;
            o.description = description;
            o.expiredTime = expiredTime;
            o.id = id;
            o.instanceChargeType = instanceChargeType;
            o.ipv6GatewayId = ipv6GatewayId;
            o.ipv6GatewayName = ipv6GatewayName;
            o.spec = spec;
            o.status = status;
            o.vpcId = vpcId;
            return o;
        }
    }
}
