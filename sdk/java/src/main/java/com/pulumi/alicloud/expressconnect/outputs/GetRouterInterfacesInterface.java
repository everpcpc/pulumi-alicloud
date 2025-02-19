// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.expressconnect.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRouterInterfacesInterface {
    /**
     * @return The access point ID to which the VBR belongs.
     * 
     */
    private String accessPointId;
    /**
     * @return The bandwidth of the resource.
     * 
     */
    private Integer bandwidth;
    /**
     * @return The businessStatus of the resource. Valid Values: `Normal`, `FinancialLocked`, `SecurityLocked`.
     * 
     */
    private String businessStatus;
    /**
     * @return The connected time of the resource.
     * 
     */
    private String connectedTime;
    /**
     * @return The creation time of the resource
     * 
     */
    private String createTime;
    /**
     * @return The cross border of the resource.
     * 
     */
    private Boolean crossBorder;
    /**
     * @return The description of the router interface.
     * 
     */
    private String description;
    /**
     * @return The end time of the resource.
     * 
     */
    private String endTime;
    /**
     * @return The has reservation data of the resource.
     * 
     */
    private String hasReservationData;
    /**
     * @return The hc rate of the resource.
     * 
     */
    private Integer hcRate;
    /**
     * @return The hc threshold of the resource.
     * 
     */
    private String hcThreshold;
    /**
     * @return The health check source IP address, must be an unused IP within the local VPC.
     * 
     */
    private String healthCheckSourceIp;
    /**
     * @return The IP address for health screening purposes.
     * 
     */
    private String healthCheckTargetIp;
    private String id;
    /**
     * @return The Access point ID to which the other end belongs.
     * 
     */
    private String oppositeAccessPointId;
    /**
     * @return The opposite bandwidth of the router on the other side.
     * 
     */
    private Integer oppositeBandwidth;
    /**
     * @return The opposite interface business status of the router on the other side. Valid Values: `Normal`, `FinancialLocked`, `SecurityLocked`.
     * 
     */
    private String oppositeInterfaceBusinessStatus;
    /**
     * @return The Interface ID of the router at the other end.
     * 
     */
    private String oppositeInterfaceId;
    /**
     * @return The AliCloud account ID of the owner of the router interface on the other end.
     * 
     */
    private String oppositeInterfaceOwnerId;
    /**
     * @return The opposite interface spec of the router on the other side. Valid Values: `Mini.2`, `Mini.5`, `Mini.5`, `Small.2`, `Small.5`, `Middle.1`, `Middle.2`, `Middle.5`, `Large.1`, `Large.2`, `Large.5`, `XLarge.1`, `Negative`.
     * 
     */
    private String oppositeInterfaceSpec;
    /**
     * @return The opposite interface status of the router on the other side. Valid Values: `Idle`, `AcceptingConnecting`, `Connecting`, `Activating`, `Active`, `Modifying`, `Deactivating`, `Inactive`, `Deleting`.
     * 
     */
    private String oppositeInterfaceStatus;
    /**
     * @return The geographical ID of the location of the receiving end of the connection.
     * 
     */
    private String oppositeRegionId;
    /**
     * @return The id of the router at the other end.
     * 
     */
    private String oppositeRouterId;
    /**
     * @return The opposite router type of the router on the other side. Valid Values: `VRouter`, `VBR`.
     * 
     */
    private String oppositeRouterType;
    /**
     * @return The opposite vpc instance id of the router on the other side.
     * 
     */
    private String oppositeVpcInstanceId;
    /**
     * @return The payment methods for router interfaces. Valid Values: `PrePaid`, `PostPaid`.
     * 
     */
    private String paymentType;
    /**
     * @return The reservation active time of the resource.
     * 
     */
    private String reservationActiveTime;
    /**
     * @return The reservation bandwidth of the resource.
     * 
     */
    private String reservationBandwidth;
    /**
     * @return The reservation internet charge type of the resource.
     * 
     */
    private String reservationInternetChargeType;
    /**
     * @return The reservation order type of the resource.
     * 
     */
    private String reservationOrderType;
    /**
     * @return The role of the router interface. Valid Values: `InitiatingSide`, `AcceptingSide`.
     * 
     */
    private String role;
    /**
     * @return The router id associated with the router interface.
     * 
     */
    private String routerId;
    /**
     * @return The first ID of the resource.
     * 
     */
    private String routerInterfaceId;
    /**
     * @return The name of the resource.
     * 
     */
    private String routerInterfaceName;
    /**
     * @return The type of router associated with the router interface. Valid Values: `VRouter`, `VBR`.
     * 
     */
    private String routerType;
    /**
     * @return The specification of the router interface. Valid Values: `Mini.2`, `Mini.5`, `Mini.5`, `Small.2`, `Small.5`, `Middle.1`, `Middle.2`, `Middle.5`, `Large.1`, `Large.2`, `Large.5`, `XLarge.1`, `Negative`.
     * 
     */
    private String spec;
    /**
     * @return The status of the resource. Valid Values: `Idle`, `AcceptingConnecting`, `Connecting`, `Activating`, `Active`, `Modifying`, `Deactivating`, `Inactive`, `Deleting`.
     * 
     */
    private String status;
    /**
     * @return The vpc instance id of the resource.
     * 
     */
    private String vpcInstanceId;

    private GetRouterInterfacesInterface() {}
    /**
     * @return The access point ID to which the VBR belongs.
     * 
     */
    public String accessPointId() {
        return this.accessPointId;
    }
    /**
     * @return The bandwidth of the resource.
     * 
     */
    public Integer bandwidth() {
        return this.bandwidth;
    }
    /**
     * @return The businessStatus of the resource. Valid Values: `Normal`, `FinancialLocked`, `SecurityLocked`.
     * 
     */
    public String businessStatus() {
        return this.businessStatus;
    }
    /**
     * @return The connected time of the resource.
     * 
     */
    public String connectedTime() {
        return this.connectedTime;
    }
    /**
     * @return The creation time of the resource
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The cross border of the resource.
     * 
     */
    public Boolean crossBorder() {
        return this.crossBorder;
    }
    /**
     * @return The description of the router interface.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The end time of the resource.
     * 
     */
    public String endTime() {
        return this.endTime;
    }
    /**
     * @return The has reservation data of the resource.
     * 
     */
    public String hasReservationData() {
        return this.hasReservationData;
    }
    /**
     * @return The hc rate of the resource.
     * 
     */
    public Integer hcRate() {
        return this.hcRate;
    }
    /**
     * @return The hc threshold of the resource.
     * 
     */
    public String hcThreshold() {
        return this.hcThreshold;
    }
    /**
     * @return The health check source IP address, must be an unused IP within the local VPC.
     * 
     */
    public String healthCheckSourceIp() {
        return this.healthCheckSourceIp;
    }
    /**
     * @return The IP address for health screening purposes.
     * 
     */
    public String healthCheckTargetIp() {
        return this.healthCheckTargetIp;
    }
    public String id() {
        return this.id;
    }
    /**
     * @return The Access point ID to which the other end belongs.
     * 
     */
    public String oppositeAccessPointId() {
        return this.oppositeAccessPointId;
    }
    /**
     * @return The opposite bandwidth of the router on the other side.
     * 
     */
    public Integer oppositeBandwidth() {
        return this.oppositeBandwidth;
    }
    /**
     * @return The opposite interface business status of the router on the other side. Valid Values: `Normal`, `FinancialLocked`, `SecurityLocked`.
     * 
     */
    public String oppositeInterfaceBusinessStatus() {
        return this.oppositeInterfaceBusinessStatus;
    }
    /**
     * @return The Interface ID of the router at the other end.
     * 
     */
    public String oppositeInterfaceId() {
        return this.oppositeInterfaceId;
    }
    /**
     * @return The AliCloud account ID of the owner of the router interface on the other end.
     * 
     */
    public String oppositeInterfaceOwnerId() {
        return this.oppositeInterfaceOwnerId;
    }
    /**
     * @return The opposite interface spec of the router on the other side. Valid Values: `Mini.2`, `Mini.5`, `Mini.5`, `Small.2`, `Small.5`, `Middle.1`, `Middle.2`, `Middle.5`, `Large.1`, `Large.2`, `Large.5`, `XLarge.1`, `Negative`.
     * 
     */
    public String oppositeInterfaceSpec() {
        return this.oppositeInterfaceSpec;
    }
    /**
     * @return The opposite interface status of the router on the other side. Valid Values: `Idle`, `AcceptingConnecting`, `Connecting`, `Activating`, `Active`, `Modifying`, `Deactivating`, `Inactive`, `Deleting`.
     * 
     */
    public String oppositeInterfaceStatus() {
        return this.oppositeInterfaceStatus;
    }
    /**
     * @return The geographical ID of the location of the receiving end of the connection.
     * 
     */
    public String oppositeRegionId() {
        return this.oppositeRegionId;
    }
    /**
     * @return The id of the router at the other end.
     * 
     */
    public String oppositeRouterId() {
        return this.oppositeRouterId;
    }
    /**
     * @return The opposite router type of the router on the other side. Valid Values: `VRouter`, `VBR`.
     * 
     */
    public String oppositeRouterType() {
        return this.oppositeRouterType;
    }
    /**
     * @return The opposite vpc instance id of the router on the other side.
     * 
     */
    public String oppositeVpcInstanceId() {
        return this.oppositeVpcInstanceId;
    }
    /**
     * @return The payment methods for router interfaces. Valid Values: `PrePaid`, `PostPaid`.
     * 
     */
    public String paymentType() {
        return this.paymentType;
    }
    /**
     * @return The reservation active time of the resource.
     * 
     */
    public String reservationActiveTime() {
        return this.reservationActiveTime;
    }
    /**
     * @return The reservation bandwidth of the resource.
     * 
     */
    public String reservationBandwidth() {
        return this.reservationBandwidth;
    }
    /**
     * @return The reservation internet charge type of the resource.
     * 
     */
    public String reservationInternetChargeType() {
        return this.reservationInternetChargeType;
    }
    /**
     * @return The reservation order type of the resource.
     * 
     */
    public String reservationOrderType() {
        return this.reservationOrderType;
    }
    /**
     * @return The role of the router interface. Valid Values: `InitiatingSide`, `AcceptingSide`.
     * 
     */
    public String role() {
        return this.role;
    }
    /**
     * @return The router id associated with the router interface.
     * 
     */
    public String routerId() {
        return this.routerId;
    }
    /**
     * @return The first ID of the resource.
     * 
     */
    public String routerInterfaceId() {
        return this.routerInterfaceId;
    }
    /**
     * @return The name of the resource.
     * 
     */
    public String routerInterfaceName() {
        return this.routerInterfaceName;
    }
    /**
     * @return The type of router associated with the router interface. Valid Values: `VRouter`, `VBR`.
     * 
     */
    public String routerType() {
        return this.routerType;
    }
    /**
     * @return The specification of the router interface. Valid Values: `Mini.2`, `Mini.5`, `Mini.5`, `Small.2`, `Small.5`, `Middle.1`, `Middle.2`, `Middle.5`, `Large.1`, `Large.2`, `Large.5`, `XLarge.1`, `Negative`.
     * 
     */
    public String spec() {
        return this.spec;
    }
    /**
     * @return The status of the resource. Valid Values: `Idle`, `AcceptingConnecting`, `Connecting`, `Activating`, `Active`, `Modifying`, `Deactivating`, `Inactive`, `Deleting`.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The vpc instance id of the resource.
     * 
     */
    public String vpcInstanceId() {
        return this.vpcInstanceId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRouterInterfacesInterface defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accessPointId;
        private Integer bandwidth;
        private String businessStatus;
        private String connectedTime;
        private String createTime;
        private Boolean crossBorder;
        private String description;
        private String endTime;
        private String hasReservationData;
        private Integer hcRate;
        private String hcThreshold;
        private String healthCheckSourceIp;
        private String healthCheckTargetIp;
        private String id;
        private String oppositeAccessPointId;
        private Integer oppositeBandwidth;
        private String oppositeInterfaceBusinessStatus;
        private String oppositeInterfaceId;
        private String oppositeInterfaceOwnerId;
        private String oppositeInterfaceSpec;
        private String oppositeInterfaceStatus;
        private String oppositeRegionId;
        private String oppositeRouterId;
        private String oppositeRouterType;
        private String oppositeVpcInstanceId;
        private String paymentType;
        private String reservationActiveTime;
        private String reservationBandwidth;
        private String reservationInternetChargeType;
        private String reservationOrderType;
        private String role;
        private String routerId;
        private String routerInterfaceId;
        private String routerInterfaceName;
        private String routerType;
        private String spec;
        private String status;
        private String vpcInstanceId;
        public Builder() {}
        public Builder(GetRouterInterfacesInterface defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessPointId = defaults.accessPointId;
    	      this.bandwidth = defaults.bandwidth;
    	      this.businessStatus = defaults.businessStatus;
    	      this.connectedTime = defaults.connectedTime;
    	      this.createTime = defaults.createTime;
    	      this.crossBorder = defaults.crossBorder;
    	      this.description = defaults.description;
    	      this.endTime = defaults.endTime;
    	      this.hasReservationData = defaults.hasReservationData;
    	      this.hcRate = defaults.hcRate;
    	      this.hcThreshold = defaults.hcThreshold;
    	      this.healthCheckSourceIp = defaults.healthCheckSourceIp;
    	      this.healthCheckTargetIp = defaults.healthCheckTargetIp;
    	      this.id = defaults.id;
    	      this.oppositeAccessPointId = defaults.oppositeAccessPointId;
    	      this.oppositeBandwidth = defaults.oppositeBandwidth;
    	      this.oppositeInterfaceBusinessStatus = defaults.oppositeInterfaceBusinessStatus;
    	      this.oppositeInterfaceId = defaults.oppositeInterfaceId;
    	      this.oppositeInterfaceOwnerId = defaults.oppositeInterfaceOwnerId;
    	      this.oppositeInterfaceSpec = defaults.oppositeInterfaceSpec;
    	      this.oppositeInterfaceStatus = defaults.oppositeInterfaceStatus;
    	      this.oppositeRegionId = defaults.oppositeRegionId;
    	      this.oppositeRouterId = defaults.oppositeRouterId;
    	      this.oppositeRouterType = defaults.oppositeRouterType;
    	      this.oppositeVpcInstanceId = defaults.oppositeVpcInstanceId;
    	      this.paymentType = defaults.paymentType;
    	      this.reservationActiveTime = defaults.reservationActiveTime;
    	      this.reservationBandwidth = defaults.reservationBandwidth;
    	      this.reservationInternetChargeType = defaults.reservationInternetChargeType;
    	      this.reservationOrderType = defaults.reservationOrderType;
    	      this.role = defaults.role;
    	      this.routerId = defaults.routerId;
    	      this.routerInterfaceId = defaults.routerInterfaceId;
    	      this.routerInterfaceName = defaults.routerInterfaceName;
    	      this.routerType = defaults.routerType;
    	      this.spec = defaults.spec;
    	      this.status = defaults.status;
    	      this.vpcInstanceId = defaults.vpcInstanceId;
        }

        @CustomType.Setter
        public Builder accessPointId(String accessPointId) {
            this.accessPointId = Objects.requireNonNull(accessPointId);
            return this;
        }
        @CustomType.Setter
        public Builder bandwidth(Integer bandwidth) {
            this.bandwidth = Objects.requireNonNull(bandwidth);
            return this;
        }
        @CustomType.Setter
        public Builder businessStatus(String businessStatus) {
            this.businessStatus = Objects.requireNonNull(businessStatus);
            return this;
        }
        @CustomType.Setter
        public Builder connectedTime(String connectedTime) {
            this.connectedTime = Objects.requireNonNull(connectedTime);
            return this;
        }
        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder crossBorder(Boolean crossBorder) {
            this.crossBorder = Objects.requireNonNull(crossBorder);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder endTime(String endTime) {
            this.endTime = Objects.requireNonNull(endTime);
            return this;
        }
        @CustomType.Setter
        public Builder hasReservationData(String hasReservationData) {
            this.hasReservationData = Objects.requireNonNull(hasReservationData);
            return this;
        }
        @CustomType.Setter
        public Builder hcRate(Integer hcRate) {
            this.hcRate = Objects.requireNonNull(hcRate);
            return this;
        }
        @CustomType.Setter
        public Builder hcThreshold(String hcThreshold) {
            this.hcThreshold = Objects.requireNonNull(hcThreshold);
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckSourceIp(String healthCheckSourceIp) {
            this.healthCheckSourceIp = Objects.requireNonNull(healthCheckSourceIp);
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckTargetIp(String healthCheckTargetIp) {
            this.healthCheckTargetIp = Objects.requireNonNull(healthCheckTargetIp);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder oppositeAccessPointId(String oppositeAccessPointId) {
            this.oppositeAccessPointId = Objects.requireNonNull(oppositeAccessPointId);
            return this;
        }
        @CustomType.Setter
        public Builder oppositeBandwidth(Integer oppositeBandwidth) {
            this.oppositeBandwidth = Objects.requireNonNull(oppositeBandwidth);
            return this;
        }
        @CustomType.Setter
        public Builder oppositeInterfaceBusinessStatus(String oppositeInterfaceBusinessStatus) {
            this.oppositeInterfaceBusinessStatus = Objects.requireNonNull(oppositeInterfaceBusinessStatus);
            return this;
        }
        @CustomType.Setter
        public Builder oppositeInterfaceId(String oppositeInterfaceId) {
            this.oppositeInterfaceId = Objects.requireNonNull(oppositeInterfaceId);
            return this;
        }
        @CustomType.Setter
        public Builder oppositeInterfaceOwnerId(String oppositeInterfaceOwnerId) {
            this.oppositeInterfaceOwnerId = Objects.requireNonNull(oppositeInterfaceOwnerId);
            return this;
        }
        @CustomType.Setter
        public Builder oppositeInterfaceSpec(String oppositeInterfaceSpec) {
            this.oppositeInterfaceSpec = Objects.requireNonNull(oppositeInterfaceSpec);
            return this;
        }
        @CustomType.Setter
        public Builder oppositeInterfaceStatus(String oppositeInterfaceStatus) {
            this.oppositeInterfaceStatus = Objects.requireNonNull(oppositeInterfaceStatus);
            return this;
        }
        @CustomType.Setter
        public Builder oppositeRegionId(String oppositeRegionId) {
            this.oppositeRegionId = Objects.requireNonNull(oppositeRegionId);
            return this;
        }
        @CustomType.Setter
        public Builder oppositeRouterId(String oppositeRouterId) {
            this.oppositeRouterId = Objects.requireNonNull(oppositeRouterId);
            return this;
        }
        @CustomType.Setter
        public Builder oppositeRouterType(String oppositeRouterType) {
            this.oppositeRouterType = Objects.requireNonNull(oppositeRouterType);
            return this;
        }
        @CustomType.Setter
        public Builder oppositeVpcInstanceId(String oppositeVpcInstanceId) {
            this.oppositeVpcInstanceId = Objects.requireNonNull(oppositeVpcInstanceId);
            return this;
        }
        @CustomType.Setter
        public Builder paymentType(String paymentType) {
            this.paymentType = Objects.requireNonNull(paymentType);
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
        public Builder role(String role) {
            this.role = Objects.requireNonNull(role);
            return this;
        }
        @CustomType.Setter
        public Builder routerId(String routerId) {
            this.routerId = Objects.requireNonNull(routerId);
            return this;
        }
        @CustomType.Setter
        public Builder routerInterfaceId(String routerInterfaceId) {
            this.routerInterfaceId = Objects.requireNonNull(routerInterfaceId);
            return this;
        }
        @CustomType.Setter
        public Builder routerInterfaceName(String routerInterfaceName) {
            this.routerInterfaceName = Objects.requireNonNull(routerInterfaceName);
            return this;
        }
        @CustomType.Setter
        public Builder routerType(String routerType) {
            this.routerType = Objects.requireNonNull(routerType);
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
        public Builder vpcInstanceId(String vpcInstanceId) {
            this.vpcInstanceId = Objects.requireNonNull(vpcInstanceId);
            return this;
        }
        public GetRouterInterfacesInterface build() {
            final var o = new GetRouterInterfacesInterface();
            o.accessPointId = accessPointId;
            o.bandwidth = bandwidth;
            o.businessStatus = businessStatus;
            o.connectedTime = connectedTime;
            o.createTime = createTime;
            o.crossBorder = crossBorder;
            o.description = description;
            o.endTime = endTime;
            o.hasReservationData = hasReservationData;
            o.hcRate = hcRate;
            o.hcThreshold = hcThreshold;
            o.healthCheckSourceIp = healthCheckSourceIp;
            o.healthCheckTargetIp = healthCheckTargetIp;
            o.id = id;
            o.oppositeAccessPointId = oppositeAccessPointId;
            o.oppositeBandwidth = oppositeBandwidth;
            o.oppositeInterfaceBusinessStatus = oppositeInterfaceBusinessStatus;
            o.oppositeInterfaceId = oppositeInterfaceId;
            o.oppositeInterfaceOwnerId = oppositeInterfaceOwnerId;
            o.oppositeInterfaceSpec = oppositeInterfaceSpec;
            o.oppositeInterfaceStatus = oppositeInterfaceStatus;
            o.oppositeRegionId = oppositeRegionId;
            o.oppositeRouterId = oppositeRouterId;
            o.oppositeRouterType = oppositeRouterType;
            o.oppositeVpcInstanceId = oppositeVpcInstanceId;
            o.paymentType = paymentType;
            o.reservationActiveTime = reservationActiveTime;
            o.reservationBandwidth = reservationBandwidth;
            o.reservationInternetChargeType = reservationInternetChargeType;
            o.reservationOrderType = reservationOrderType;
            o.role = role;
            o.routerId = routerId;
            o.routerInterfaceId = routerInterfaceId;
            o.routerInterfaceName = routerInterfaceName;
            o.routerType = routerType;
            o.spec = spec;
            o.status = status;
            o.vpcInstanceId = vpcInstanceId;
            return o;
        }
    }
}
