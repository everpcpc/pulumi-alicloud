// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudstoragegateway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetGatewaysGateway {
    /**
     * @return gateway .
     * 
     */
    private String activatedTime;
    private String buyUrl;
    /**
     * @return gateway category.
     * 
     */
    private String category;
    /**
     * @return gateway created timestamp in second format.
     * 
     */
    private String createTime;
    /**
     * @return gateway description.
     * 
     */
    private String description;
    /**
     * @return gateway ecs instance id.
     * 
     */
    private String ecsInstanceId;
    /**
     * @return gateway expiration status.
     * 
     */
    private Integer expireStatus;
    /**
     * @return gateway expiration timestamp in second format.
     * 
     */
    private String expiredTime;
    /**
     * @return gateway class.
     * 
     */
    private String gatewayClass;
    /**
     * @return gateway id.
     * 
     */
    private String gatewayId;
    /**
     * @return gateway name.
     * 
     */
    private String gatewayName;
    /**
     * @return gateway version.
     * 
     */
    private String gatewayVersion;
    /**
     * @return The ID of the Gateway.
     * 
     */
    private String id;
    /**
     * @return gateway service ip.
     * 
     */
    private String innerIp;
    /**
     * @return gateway public ip.
     * 
     */
    private String ip;
    /**
     * @return whether subscription gateway is released after expiration or not.
     * 
     */
    private Boolean isReleaseAfterExpiration;
    /**
     * @return gateway location.
     * 
     */
    private String location;
    /**
     * @return gateway payment type. The Payment type of gateway. The valid value: `PayAsYouGo`, `Subscription`.
     * 
     */
    private String paymentType;
    /**
     * @return gateway public network bandwidth.
     * 
     */
    private Integer publicNetworkBandwidth;
    private String renewUrl;
    /**
     * @return gateway status.
     * 
     */
    private String status;
    /**
     * @return storage bundle id.
     * 
     */
    private String storageBundleId;
    /**
     * @return gateway task id.
     * 
     */
    private String taskId;
    /**
     * @return gateway type.
     * 
     */
    private String type;
    /**
     * @return gateway vpc id.
     * 
     */
    private String vpcId;
    /**
     * @return The vswitch id.
     * 
     */
    private String vswitchId;

    private GetGatewaysGateway() {}
    /**
     * @return gateway .
     * 
     */
    public String activatedTime() {
        return this.activatedTime;
    }
    public String buyUrl() {
        return this.buyUrl;
    }
    /**
     * @return gateway category.
     * 
     */
    public String category() {
        return this.category;
    }
    /**
     * @return gateway created timestamp in second format.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return gateway description.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return gateway ecs instance id.
     * 
     */
    public String ecsInstanceId() {
        return this.ecsInstanceId;
    }
    /**
     * @return gateway expiration status.
     * 
     */
    public Integer expireStatus() {
        return this.expireStatus;
    }
    /**
     * @return gateway expiration timestamp in second format.
     * 
     */
    public String expiredTime() {
        return this.expiredTime;
    }
    /**
     * @return gateway class.
     * 
     */
    public String gatewayClass() {
        return this.gatewayClass;
    }
    /**
     * @return gateway id.
     * 
     */
    public String gatewayId() {
        return this.gatewayId;
    }
    /**
     * @return gateway name.
     * 
     */
    public String gatewayName() {
        return this.gatewayName;
    }
    /**
     * @return gateway version.
     * 
     */
    public String gatewayVersion() {
        return this.gatewayVersion;
    }
    /**
     * @return The ID of the Gateway.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return gateway service ip.
     * 
     */
    public String innerIp() {
        return this.innerIp;
    }
    /**
     * @return gateway public ip.
     * 
     */
    public String ip() {
        return this.ip;
    }
    /**
     * @return whether subscription gateway is released after expiration or not.
     * 
     */
    public Boolean isReleaseAfterExpiration() {
        return this.isReleaseAfterExpiration;
    }
    /**
     * @return gateway location.
     * 
     */
    public String location() {
        return this.location;
    }
    /**
     * @return gateway payment type. The Payment type of gateway. The valid value: `PayAsYouGo`, `Subscription`.
     * 
     */
    public String paymentType() {
        return this.paymentType;
    }
    /**
     * @return gateway public network bandwidth.
     * 
     */
    public Integer publicNetworkBandwidth() {
        return this.publicNetworkBandwidth;
    }
    public String renewUrl() {
        return this.renewUrl;
    }
    /**
     * @return gateway status.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return storage bundle id.
     * 
     */
    public String storageBundleId() {
        return this.storageBundleId;
    }
    /**
     * @return gateway task id.
     * 
     */
    public String taskId() {
        return this.taskId;
    }
    /**
     * @return gateway type.
     * 
     */
    public String type() {
        return this.type;
    }
    /**
     * @return gateway vpc id.
     * 
     */
    public String vpcId() {
        return this.vpcId;
    }
    /**
     * @return The vswitch id.
     * 
     */
    public String vswitchId() {
        return this.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGatewaysGateway defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String activatedTime;
        private String buyUrl;
        private String category;
        private String createTime;
        private String description;
        private String ecsInstanceId;
        private Integer expireStatus;
        private String expiredTime;
        private String gatewayClass;
        private String gatewayId;
        private String gatewayName;
        private String gatewayVersion;
        private String id;
        private String innerIp;
        private String ip;
        private Boolean isReleaseAfterExpiration;
        private String location;
        private String paymentType;
        private Integer publicNetworkBandwidth;
        private String renewUrl;
        private String status;
        private String storageBundleId;
        private String taskId;
        private String type;
        private String vpcId;
        private String vswitchId;
        public Builder() {}
        public Builder(GetGatewaysGateway defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.activatedTime = defaults.activatedTime;
    	      this.buyUrl = defaults.buyUrl;
    	      this.category = defaults.category;
    	      this.createTime = defaults.createTime;
    	      this.description = defaults.description;
    	      this.ecsInstanceId = defaults.ecsInstanceId;
    	      this.expireStatus = defaults.expireStatus;
    	      this.expiredTime = defaults.expiredTime;
    	      this.gatewayClass = defaults.gatewayClass;
    	      this.gatewayId = defaults.gatewayId;
    	      this.gatewayName = defaults.gatewayName;
    	      this.gatewayVersion = defaults.gatewayVersion;
    	      this.id = defaults.id;
    	      this.innerIp = defaults.innerIp;
    	      this.ip = defaults.ip;
    	      this.isReleaseAfterExpiration = defaults.isReleaseAfterExpiration;
    	      this.location = defaults.location;
    	      this.paymentType = defaults.paymentType;
    	      this.publicNetworkBandwidth = defaults.publicNetworkBandwidth;
    	      this.renewUrl = defaults.renewUrl;
    	      this.status = defaults.status;
    	      this.storageBundleId = defaults.storageBundleId;
    	      this.taskId = defaults.taskId;
    	      this.type = defaults.type;
    	      this.vpcId = defaults.vpcId;
    	      this.vswitchId = defaults.vswitchId;
        }

        @CustomType.Setter
        public Builder activatedTime(String activatedTime) {
            this.activatedTime = Objects.requireNonNull(activatedTime);
            return this;
        }
        @CustomType.Setter
        public Builder buyUrl(String buyUrl) {
            this.buyUrl = Objects.requireNonNull(buyUrl);
            return this;
        }
        @CustomType.Setter
        public Builder category(String category) {
            this.category = Objects.requireNonNull(category);
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
        public Builder ecsInstanceId(String ecsInstanceId) {
            this.ecsInstanceId = Objects.requireNonNull(ecsInstanceId);
            return this;
        }
        @CustomType.Setter
        public Builder expireStatus(Integer expireStatus) {
            this.expireStatus = Objects.requireNonNull(expireStatus);
            return this;
        }
        @CustomType.Setter
        public Builder expiredTime(String expiredTime) {
            this.expiredTime = Objects.requireNonNull(expiredTime);
            return this;
        }
        @CustomType.Setter
        public Builder gatewayClass(String gatewayClass) {
            this.gatewayClass = Objects.requireNonNull(gatewayClass);
            return this;
        }
        @CustomType.Setter
        public Builder gatewayId(String gatewayId) {
            this.gatewayId = Objects.requireNonNull(gatewayId);
            return this;
        }
        @CustomType.Setter
        public Builder gatewayName(String gatewayName) {
            this.gatewayName = Objects.requireNonNull(gatewayName);
            return this;
        }
        @CustomType.Setter
        public Builder gatewayVersion(String gatewayVersion) {
            this.gatewayVersion = Objects.requireNonNull(gatewayVersion);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder innerIp(String innerIp) {
            this.innerIp = Objects.requireNonNull(innerIp);
            return this;
        }
        @CustomType.Setter
        public Builder ip(String ip) {
            this.ip = Objects.requireNonNull(ip);
            return this;
        }
        @CustomType.Setter
        public Builder isReleaseAfterExpiration(Boolean isReleaseAfterExpiration) {
            this.isReleaseAfterExpiration = Objects.requireNonNull(isReleaseAfterExpiration);
            return this;
        }
        @CustomType.Setter
        public Builder location(String location) {
            this.location = Objects.requireNonNull(location);
            return this;
        }
        @CustomType.Setter
        public Builder paymentType(String paymentType) {
            this.paymentType = Objects.requireNonNull(paymentType);
            return this;
        }
        @CustomType.Setter
        public Builder publicNetworkBandwidth(Integer publicNetworkBandwidth) {
            this.publicNetworkBandwidth = Objects.requireNonNull(publicNetworkBandwidth);
            return this;
        }
        @CustomType.Setter
        public Builder renewUrl(String renewUrl) {
            this.renewUrl = Objects.requireNonNull(renewUrl);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder storageBundleId(String storageBundleId) {
            this.storageBundleId = Objects.requireNonNull(storageBundleId);
            return this;
        }
        @CustomType.Setter
        public Builder taskId(String taskId) {
            this.taskId = Objects.requireNonNull(taskId);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
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
        public GetGatewaysGateway build() {
            final var o = new GetGatewaysGateway();
            o.activatedTime = activatedTime;
            o.buyUrl = buyUrl;
            o.category = category;
            o.createTime = createTime;
            o.description = description;
            o.ecsInstanceId = ecsInstanceId;
            o.expireStatus = expireStatus;
            o.expiredTime = expiredTime;
            o.gatewayClass = gatewayClass;
            o.gatewayId = gatewayId;
            o.gatewayName = gatewayName;
            o.gatewayVersion = gatewayVersion;
            o.id = id;
            o.innerIp = innerIp;
            o.ip = ip;
            o.isReleaseAfterExpiration = isReleaseAfterExpiration;
            o.location = location;
            o.paymentType = paymentType;
            o.publicNetworkBandwidth = publicNetworkBandwidth;
            o.renewUrl = renewUrl;
            o.status = status;
            o.storageBundleId = storageBundleId;
            o.taskId = taskId;
            o.type = type;
            o.vpcId = vpcId;
            o.vswitchId = vswitchId;
            return o;
        }
    }
}
