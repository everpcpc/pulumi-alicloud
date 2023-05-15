// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetPbrRouteEntriesEntry {
    /**
     * @return The creation time of the VPN Pbr Route Entry.
     * 
     */
    private String createTime;
    /**
     * @return The id of the vpn pbr route entry. The value formats as `&lt;vpn_gateway_id&gt;:&lt;next_hop&gt;:&lt;route_source&gt;:&lt;route_dest&gt;`.
     * 
     */
    private String id;
    /**
     * @return The next hop of the policy-based route.
     * 
     */
    private String nextHop;
    /**
     * @return The destination CIDR block of the policy-based route.
     * 
     */
    private String routeDest;
    /**
     * @return The source CIDR block of the policy-based route.
     * 
     */
    private String routeSource;
    /**
     * @return The status of the VPN Pbr Route Entry.
     * 
     */
    private String status;
    /**
     * @return The ID of the VPN gateway.
     * 
     */
    private String vpnGatewayId;
    /**
     * @return The weight of the policy-based route. Valid values: 0 and 100.
     * 
     */
    private Integer weight;

    private GetPbrRouteEntriesEntry() {}
    /**
     * @return The creation time of the VPN Pbr Route Entry.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The id of the vpn pbr route entry. The value formats as `&lt;vpn_gateway_id&gt;:&lt;next_hop&gt;:&lt;route_source&gt;:&lt;route_dest&gt;`.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The next hop of the policy-based route.
     * 
     */
    public String nextHop() {
        return this.nextHop;
    }
    /**
     * @return The destination CIDR block of the policy-based route.
     * 
     */
    public String routeDest() {
        return this.routeDest;
    }
    /**
     * @return The source CIDR block of the policy-based route.
     * 
     */
    public String routeSource() {
        return this.routeSource;
    }
    /**
     * @return The status of the VPN Pbr Route Entry.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The ID of the VPN gateway.
     * 
     */
    public String vpnGatewayId() {
        return this.vpnGatewayId;
    }
    /**
     * @return The weight of the policy-based route. Valid values: 0 and 100.
     * 
     */
    public Integer weight() {
        return this.weight;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPbrRouteEntriesEntry defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createTime;
        private String id;
        private String nextHop;
        private String routeDest;
        private String routeSource;
        private String status;
        private String vpnGatewayId;
        private Integer weight;
        public Builder() {}
        public Builder(GetPbrRouteEntriesEntry defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createTime = defaults.createTime;
    	      this.id = defaults.id;
    	      this.nextHop = defaults.nextHop;
    	      this.routeDest = defaults.routeDest;
    	      this.routeSource = defaults.routeSource;
    	      this.status = defaults.status;
    	      this.vpnGatewayId = defaults.vpnGatewayId;
    	      this.weight = defaults.weight;
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
        public Builder nextHop(String nextHop) {
            this.nextHop = Objects.requireNonNull(nextHop);
            return this;
        }
        @CustomType.Setter
        public Builder routeDest(String routeDest) {
            this.routeDest = Objects.requireNonNull(routeDest);
            return this;
        }
        @CustomType.Setter
        public Builder routeSource(String routeSource) {
            this.routeSource = Objects.requireNonNull(routeSource);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder vpnGatewayId(String vpnGatewayId) {
            this.vpnGatewayId = Objects.requireNonNull(vpnGatewayId);
            return this;
        }
        @CustomType.Setter
        public Builder weight(Integer weight) {
            this.weight = Objects.requireNonNull(weight);
            return this;
        }
        public GetPbrRouteEntriesEntry build() {
            final var o = new GetPbrRouteEntriesEntry();
            o.createTime = createTime;
            o.id = id;
            o.nextHop = nextHop;
            o.routeDest = routeDest;
            o.routeSource = routeSource;
            o.status = status;
            o.vpnGatewayId = vpnGatewayId;
            o.weight = weight;
            return o;
        }
    }
}
