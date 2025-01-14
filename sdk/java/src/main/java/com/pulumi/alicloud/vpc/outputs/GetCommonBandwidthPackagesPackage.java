// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.outputs;

import com.pulumi.alicloud.vpc.outputs.GetCommonBandwidthPackagesPackagePublicIpAddress;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetCommonBandwidthPackagesPackage {
    /**
     * @return The peak bandwidth of the Internet Shared Bandwidth instance.
     * 
     */
    private String bandwidth;
    /**
     * @return The resource ID of bandwidth package.
     * 
     */
    private String bandwidthPackageId;
    /**
     * @return The name of bandwidth package.
     * 
     */
    private String bandwidthPackageName;
    /**
     * @return The business status of the Common Bandwidth Package instance.
     * 
     */
    private String businessStatus;
    /**
     * @return The deletion protection of bandwidth package.
     * 
     */
    private Boolean deletionProtection;
    /**
     * @return The description of the Common Bandwidth Package instance.
     * 
     */
    private String description;
    /**
     * @return The expired time of bandwidth package.
     * 
     */
    private String expiredTime;
    /**
     * @return Is has reservation data.
     * 
     */
    private Boolean hasReservationData;
    /**
     * @return ID of the Common Bandwidth Package.
     * 
     */
    private String id;
    /**
     * @return The internet charge type of bandwidth package.
     * 
     */
    private String internetChargeType;
    /**
     * @return ISP of the Common Bandwidth Package.
     * 
     */
    private String isp;
    /**
     * @return Name of the Common Bandwidth Package.
     * 
     */
    private String name;
    /**
     * @return The payment type of bandwidth package.
     * 
     */
    private String paymentType;
    /**
     * @return Public ip addresses that in the Common Bandwidth Pakcage.
     * 
     */
    private List<GetCommonBandwidthPackagesPackagePublicIpAddress> publicIpAddresses;
    /**
     * @return The ratio of bandwidth package.
     * 
     */
    private Integer ratio;
    /**
     * @return The active time of reservation.
     * 
     */
    private String reservationActiveTime;
    /**
     * @return The bandwidth of reservation.
     * 
     */
    private String reservationBandwidth;
    /**
     * @return The charge type of reservation internet.
     * 
     */
    private String reservationInternetChargeType;
    /**
     * @return The type of reservation order.
     * 
     */
    private String reservationOrderType;
    /**
     * @return The Id of resource group which the common bandwidth package belongs.
     * 
     */
    private String resourceGroupId;
    /**
     * @return The service managed.
     * 
     */
    private Integer serviceManaged;
    /**
     * @return The status of bandwidth package. Valid values: `Available` and `Pending`.
     * 
     */
    private String status;

    private GetCommonBandwidthPackagesPackage() {}
    /**
     * @return The peak bandwidth of the Internet Shared Bandwidth instance.
     * 
     */
    public String bandwidth() {
        return this.bandwidth;
    }
    /**
     * @return The resource ID of bandwidth package.
     * 
     */
    public String bandwidthPackageId() {
        return this.bandwidthPackageId;
    }
    /**
     * @return The name of bandwidth package.
     * 
     */
    public String bandwidthPackageName() {
        return this.bandwidthPackageName;
    }
    /**
     * @return The business status of the Common Bandwidth Package instance.
     * 
     */
    public String businessStatus() {
        return this.businessStatus;
    }
    /**
     * @return The deletion protection of bandwidth package.
     * 
     */
    public Boolean deletionProtection() {
        return this.deletionProtection;
    }
    /**
     * @return The description of the Common Bandwidth Package instance.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The expired time of bandwidth package.
     * 
     */
    public String expiredTime() {
        return this.expiredTime;
    }
    /**
     * @return Is has reservation data.
     * 
     */
    public Boolean hasReservationData() {
        return this.hasReservationData;
    }
    /**
     * @return ID of the Common Bandwidth Package.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The internet charge type of bandwidth package.
     * 
     */
    public String internetChargeType() {
        return this.internetChargeType;
    }
    /**
     * @return ISP of the Common Bandwidth Package.
     * 
     */
    public String isp() {
        return this.isp;
    }
    /**
     * @return Name of the Common Bandwidth Package.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The payment type of bandwidth package.
     * 
     */
    public String paymentType() {
        return this.paymentType;
    }
    /**
     * @return Public ip addresses that in the Common Bandwidth Pakcage.
     * 
     */
    public List<GetCommonBandwidthPackagesPackagePublicIpAddress> publicIpAddresses() {
        return this.publicIpAddresses;
    }
    /**
     * @return The ratio of bandwidth package.
     * 
     */
    public Integer ratio() {
        return this.ratio;
    }
    /**
     * @return The active time of reservation.
     * 
     */
    public String reservationActiveTime() {
        return this.reservationActiveTime;
    }
    /**
     * @return The bandwidth of reservation.
     * 
     */
    public String reservationBandwidth() {
        return this.reservationBandwidth;
    }
    /**
     * @return The charge type of reservation internet.
     * 
     */
    public String reservationInternetChargeType() {
        return this.reservationInternetChargeType;
    }
    /**
     * @return The type of reservation order.
     * 
     */
    public String reservationOrderType() {
        return this.reservationOrderType;
    }
    /**
     * @return The Id of resource group which the common bandwidth package belongs.
     * 
     */
    public String resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * @return The service managed.
     * 
     */
    public Integer serviceManaged() {
        return this.serviceManaged;
    }
    /**
     * @return The status of bandwidth package. Valid values: `Available` and `Pending`.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCommonBandwidthPackagesPackage defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String bandwidth;
        private String bandwidthPackageId;
        private String bandwidthPackageName;
        private String businessStatus;
        private Boolean deletionProtection;
        private String description;
        private String expiredTime;
        private Boolean hasReservationData;
        private String id;
        private String internetChargeType;
        private String isp;
        private String name;
        private String paymentType;
        private List<GetCommonBandwidthPackagesPackagePublicIpAddress> publicIpAddresses;
        private Integer ratio;
        private String reservationActiveTime;
        private String reservationBandwidth;
        private String reservationInternetChargeType;
        private String reservationOrderType;
        private String resourceGroupId;
        private Integer serviceManaged;
        private String status;
        public Builder() {}
        public Builder(GetCommonBandwidthPackagesPackage defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bandwidth = defaults.bandwidth;
    	      this.bandwidthPackageId = defaults.bandwidthPackageId;
    	      this.bandwidthPackageName = defaults.bandwidthPackageName;
    	      this.businessStatus = defaults.businessStatus;
    	      this.deletionProtection = defaults.deletionProtection;
    	      this.description = defaults.description;
    	      this.expiredTime = defaults.expiredTime;
    	      this.hasReservationData = defaults.hasReservationData;
    	      this.id = defaults.id;
    	      this.internetChargeType = defaults.internetChargeType;
    	      this.isp = defaults.isp;
    	      this.name = defaults.name;
    	      this.paymentType = defaults.paymentType;
    	      this.publicIpAddresses = defaults.publicIpAddresses;
    	      this.ratio = defaults.ratio;
    	      this.reservationActiveTime = defaults.reservationActiveTime;
    	      this.reservationBandwidth = defaults.reservationBandwidth;
    	      this.reservationInternetChargeType = defaults.reservationInternetChargeType;
    	      this.reservationOrderType = defaults.reservationOrderType;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.serviceManaged = defaults.serviceManaged;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder bandwidth(String bandwidth) {
            this.bandwidth = Objects.requireNonNull(bandwidth);
            return this;
        }
        @CustomType.Setter
        public Builder bandwidthPackageId(String bandwidthPackageId) {
            this.bandwidthPackageId = Objects.requireNonNull(bandwidthPackageId);
            return this;
        }
        @CustomType.Setter
        public Builder bandwidthPackageName(String bandwidthPackageName) {
            this.bandwidthPackageName = Objects.requireNonNull(bandwidthPackageName);
            return this;
        }
        @CustomType.Setter
        public Builder businessStatus(String businessStatus) {
            this.businessStatus = Objects.requireNonNull(businessStatus);
            return this;
        }
        @CustomType.Setter
        public Builder deletionProtection(Boolean deletionProtection) {
            this.deletionProtection = Objects.requireNonNull(deletionProtection);
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
        public Builder hasReservationData(Boolean hasReservationData) {
            this.hasReservationData = Objects.requireNonNull(hasReservationData);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder internetChargeType(String internetChargeType) {
            this.internetChargeType = Objects.requireNonNull(internetChargeType);
            return this;
        }
        @CustomType.Setter
        public Builder isp(String isp) {
            this.isp = Objects.requireNonNull(isp);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder paymentType(String paymentType) {
            this.paymentType = Objects.requireNonNull(paymentType);
            return this;
        }
        @CustomType.Setter
        public Builder publicIpAddresses(List<GetCommonBandwidthPackagesPackagePublicIpAddress> publicIpAddresses) {
            this.publicIpAddresses = Objects.requireNonNull(publicIpAddresses);
            return this;
        }
        public Builder publicIpAddresses(GetCommonBandwidthPackagesPackagePublicIpAddress... publicIpAddresses) {
            return publicIpAddresses(List.of(publicIpAddresses));
        }
        @CustomType.Setter
        public Builder ratio(Integer ratio) {
            this.ratio = Objects.requireNonNull(ratio);
            return this;
        }
        @CustomType.Setter
        public Builder reservationActiveTime(String reservationActiveTime) {
            this.reservationActiveTime = Objects.requireNonNull(reservationActiveTime);
            return this;
        }
        @CustomType.Setter
        public Builder reservationBandwidth(String reservationBandwidth) {
            this.reservationBandwidth = Objects.requireNonNull(reservationBandwidth);
            return this;
        }
        @CustomType.Setter
        public Builder reservationInternetChargeType(String reservationInternetChargeType) {
            this.reservationInternetChargeType = Objects.requireNonNull(reservationInternetChargeType);
            return this;
        }
        @CustomType.Setter
        public Builder reservationOrderType(String reservationOrderType) {
            this.reservationOrderType = Objects.requireNonNull(reservationOrderType);
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroupId(String resourceGroupId) {
            this.resourceGroupId = Objects.requireNonNull(resourceGroupId);
            return this;
        }
        @CustomType.Setter
        public Builder serviceManaged(Integer serviceManaged) {
            this.serviceManaged = Objects.requireNonNull(serviceManaged);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetCommonBandwidthPackagesPackage build() {
            final var o = new GetCommonBandwidthPackagesPackage();
            o.bandwidth = bandwidth;
            o.bandwidthPackageId = bandwidthPackageId;
            o.bandwidthPackageName = bandwidthPackageName;
            o.businessStatus = businessStatus;
            o.deletionProtection = deletionProtection;
            o.description = description;
            o.expiredTime = expiredTime;
            o.hasReservationData = hasReservationData;
            o.id = id;
            o.internetChargeType = internetChargeType;
            o.isp = isp;
            o.name = name;
            o.paymentType = paymentType;
            o.publicIpAddresses = publicIpAddresses;
            o.ratio = ratio;
            o.reservationActiveTime = reservationActiveTime;
            o.reservationBandwidth = reservationBandwidth;
            o.reservationInternetChargeType = reservationInternetChargeType;
            o.reservationOrderType = reservationOrderType;
            o.resourceGroupId = resourceGroupId;
            o.serviceManaged = serviceManaged;
            o.status = status;
            return o;
        }
    }
}
