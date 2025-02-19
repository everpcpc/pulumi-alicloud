// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRouteEntriesEntry {
    /**
     * @return The destination CIDR block of the route entry.
     * 
     */
    private String cidrBlock;
    /**
     * @return The instance ID of the next hop.
     * 
     */
    private String instanceId;
    /**
     * @return The type of the next hop.
     * 
     */
    private String nextHopType;
    /**
     * @return The ID of the router table to which the route entry belongs.
     * 
     */
    private String routeTableId;
    /**
     * @return The status of the route entry.
     * 
     */
    private String status;
    /**
     * @return The type of the route entry.
     * 
     */
    private String type;

    private GetRouteEntriesEntry() {}
    /**
     * @return The destination CIDR block of the route entry.
     * 
     */
    public String cidrBlock() {
        return this.cidrBlock;
    }
    /**
     * @return The instance ID of the next hop.
     * 
     */
    public String instanceId() {
        return this.instanceId;
    }
    /**
     * @return The type of the next hop.
     * 
     */
    public String nextHopType() {
        return this.nextHopType;
    }
    /**
     * @return The ID of the router table to which the route entry belongs.
     * 
     */
    public String routeTableId() {
        return this.routeTableId;
    }
    /**
     * @return The status of the route entry.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The type of the route entry.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRouteEntriesEntry defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String cidrBlock;
        private String instanceId;
        private String nextHopType;
        private String routeTableId;
        private String status;
        private String type;
        public Builder() {}
        public Builder(GetRouteEntriesEntry defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cidrBlock = defaults.cidrBlock;
    	      this.instanceId = defaults.instanceId;
    	      this.nextHopType = defaults.nextHopType;
    	      this.routeTableId = defaults.routeTableId;
    	      this.status = defaults.status;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder cidrBlock(String cidrBlock) {
            this.cidrBlock = Objects.requireNonNull(cidrBlock);
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(String instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder nextHopType(String nextHopType) {
            this.nextHopType = Objects.requireNonNull(nextHopType);
            return this;
        }
        @CustomType.Setter
        public Builder routeTableId(String routeTableId) {
            this.routeTableId = Objects.requireNonNull(routeTableId);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetRouteEntriesEntry build() {
            final var o = new GetRouteEntriesEntry();
            o.cidrBlock = cidrBlock;
            o.instanceId = instanceId;
            o.nextHopType = nextHopType;
            o.routeTableId = routeTableId;
            o.status = status;
            o.type = type;
            return o;
        }
    }
}
