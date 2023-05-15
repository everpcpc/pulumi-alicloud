// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRouterInterfacesInterface {
    /**
     * @return ID of the access point used by the VBR.
     * 
     */
    private String accessPointId;
    /**
     * @return Router interface creation time.
     * 
     */
    private String creationTime;
    /**
     * @return Router interface description.
     * 
     */
    private String description;
    /**
     * @return Source IP address used to perform health check on the physical connection.
     * 
     */
    private String healthCheckSourceIp;
    /**
     * @return Destination IP address used to perform health check on the physical connection.
     * 
     */
    private String healthCheckTargetIp;
    /**
     * @return Router interface ID.
     * 
     */
    private String id;
    /**
     * @return Router interface name.
     * 
     */
    private String name;
    /**
     * @return ID of the peer router interface.
     * 
     */
    private String oppositeInterfaceId;
    /**
     * @return Account ID of the owner of the peer router interface.
     * 
     */
    private String oppositeInterfaceOwnerId;
    /**
     * @return Peer router region ID.
     * 
     */
    private String oppositeRegionId;
    /**
     * @return Peer router ID.
     * 
     */
    private String oppositeRouterId;
    /**
     * @return Router type in the peer region. Possible values: `VRouter` and `VBR`.
     * 
     */
    private String oppositeRouterType;
    /**
     * @return Role of the router interface. Valid values are `InitiatingSide` (connection initiator) and
     * `AcceptingSide` (connection receiver). The value of this parameter must be `InitiatingSide` if the `router_type` is set to `VBR`.
     * 
     */
    private String role;
    /**
     * @return ID of the VRouter located in the local region.
     * 
     */
    private String routerId;
    /**
     * @return Router type in the local region. Valid values are `VRouter` and `VBR` (physical connection).
     * 
     */
    private String routerType;
    /**
     * @return Specification of the link, such as `Small.1` (10Mb), `Middle.1` (100Mb), `Large.2` (2Gb), ...etc.
     * 
     */
    private String specification;
    /**
     * @return Expected status. Valid values are `Active`, `Inactive` and `Idle`.
     * 
     */
    private String status;
    /**
     * @return ID of the VPC that owns the router in the local region.
     * 
     */
    private String vpcId;

    private GetRouterInterfacesInterface() {}
    /**
     * @return ID of the access point used by the VBR.
     * 
     */
    public String accessPointId() {
        return this.accessPointId;
    }
    /**
     * @return Router interface creation time.
     * 
     */
    public String creationTime() {
        return this.creationTime;
    }
    /**
     * @return Router interface description.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return Source IP address used to perform health check on the physical connection.
     * 
     */
    public String healthCheckSourceIp() {
        return this.healthCheckSourceIp;
    }
    /**
     * @return Destination IP address used to perform health check on the physical connection.
     * 
     */
    public String healthCheckTargetIp() {
        return this.healthCheckTargetIp;
    }
    /**
     * @return Router interface ID.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Router interface name.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return ID of the peer router interface.
     * 
     */
    public String oppositeInterfaceId() {
        return this.oppositeInterfaceId;
    }
    /**
     * @return Account ID of the owner of the peer router interface.
     * 
     */
    public String oppositeInterfaceOwnerId() {
        return this.oppositeInterfaceOwnerId;
    }
    /**
     * @return Peer router region ID.
     * 
     */
    public String oppositeRegionId() {
        return this.oppositeRegionId;
    }
    /**
     * @return Peer router ID.
     * 
     */
    public String oppositeRouterId() {
        return this.oppositeRouterId;
    }
    /**
     * @return Router type in the peer region. Possible values: `VRouter` and `VBR`.
     * 
     */
    public String oppositeRouterType() {
        return this.oppositeRouterType;
    }
    /**
     * @return Role of the router interface. Valid values are `InitiatingSide` (connection initiator) and
     * `AcceptingSide` (connection receiver). The value of this parameter must be `InitiatingSide` if the `router_type` is set to `VBR`.
     * 
     */
    public String role() {
        return this.role;
    }
    /**
     * @return ID of the VRouter located in the local region.
     * 
     */
    public String routerId() {
        return this.routerId;
    }
    /**
     * @return Router type in the local region. Valid values are `VRouter` and `VBR` (physical connection).
     * 
     */
    public String routerType() {
        return this.routerType;
    }
    /**
     * @return Specification of the link, such as `Small.1` (10Mb), `Middle.1` (100Mb), `Large.2` (2Gb), ...etc.
     * 
     */
    public String specification() {
        return this.specification;
    }
    /**
     * @return Expected status. Valid values are `Active`, `Inactive` and `Idle`.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return ID of the VPC that owns the router in the local region.
     * 
     */
    public String vpcId() {
        return this.vpcId;
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
        private String creationTime;
        private String description;
        private String healthCheckSourceIp;
        private String healthCheckTargetIp;
        private String id;
        private String name;
        private String oppositeInterfaceId;
        private String oppositeInterfaceOwnerId;
        private String oppositeRegionId;
        private String oppositeRouterId;
        private String oppositeRouterType;
        private String role;
        private String routerId;
        private String routerType;
        private String specification;
        private String status;
        private String vpcId;
        public Builder() {}
        public Builder(GetRouterInterfacesInterface defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessPointId = defaults.accessPointId;
    	      this.creationTime = defaults.creationTime;
    	      this.description = defaults.description;
    	      this.healthCheckSourceIp = defaults.healthCheckSourceIp;
    	      this.healthCheckTargetIp = defaults.healthCheckTargetIp;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.oppositeInterfaceId = defaults.oppositeInterfaceId;
    	      this.oppositeInterfaceOwnerId = defaults.oppositeInterfaceOwnerId;
    	      this.oppositeRegionId = defaults.oppositeRegionId;
    	      this.oppositeRouterId = defaults.oppositeRouterId;
    	      this.oppositeRouterType = defaults.oppositeRouterType;
    	      this.role = defaults.role;
    	      this.routerId = defaults.routerId;
    	      this.routerType = defaults.routerType;
    	      this.specification = defaults.specification;
    	      this.status = defaults.status;
    	      this.vpcId = defaults.vpcId;
        }

        @CustomType.Setter
        public Builder accessPointId(String accessPointId) {
            this.accessPointId = Objects.requireNonNull(accessPointId);
            return this;
        }
        @CustomType.Setter
        public Builder creationTime(String creationTime) {
            this.creationTime = Objects.requireNonNull(creationTime);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
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
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
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
        public Builder routerType(String routerType) {
            this.routerType = Objects.requireNonNull(routerType);
            return this;
        }
        @CustomType.Setter
        public Builder specification(String specification) {
            this.specification = Objects.requireNonNull(specification);
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
        public GetRouterInterfacesInterface build() {
            final var o = new GetRouterInterfacesInterface();
            o.accessPointId = accessPointId;
            o.creationTime = creationTime;
            o.description = description;
            o.healthCheckSourceIp = healthCheckSourceIp;
            o.healthCheckTargetIp = healthCheckTargetIp;
            o.id = id;
            o.name = name;
            o.oppositeInterfaceId = oppositeInterfaceId;
            o.oppositeInterfaceOwnerId = oppositeInterfaceOwnerId;
            o.oppositeRegionId = oppositeRegionId;
            o.oppositeRouterId = oppositeRouterId;
            o.oppositeRouterType = oppositeRouterType;
            o.role = role;
            o.routerId = routerId;
            o.routerType = routerType;
            o.specification = specification;
            o.status = status;
            o.vpcId = vpcId;
            return o;
        }
    }
}
