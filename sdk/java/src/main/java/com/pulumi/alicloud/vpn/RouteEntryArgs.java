// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpn;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class RouteEntryArgs extends com.pulumi.resources.ResourceArgs {

    public static final RouteEntryArgs Empty = new RouteEntryArgs();

    /**
     * The next hop of the destination route.
     * 
     */
    @Import(name="nextHop", required=true)
    private Output<String> nextHop;

    /**
     * @return The next hop of the destination route.
     * 
     */
    public Output<String> nextHop() {
        return this.nextHop;
    }

    /**
     * Whether to issue the destination route to the VPC.
     * 
     */
    @Import(name="publishVpc", required=true)
    private Output<Boolean> publishVpc;

    /**
     * @return Whether to issue the destination route to the VPC.
     * 
     */
    public Output<Boolean> publishVpc() {
        return this.publishVpc;
    }

    /**
     * The destination network segment of the destination route.
     * 
     */
    @Import(name="routeDest", required=true)
    private Output<String> routeDest;

    /**
     * @return The destination network segment of the destination route.
     * 
     */
    public Output<String> routeDest() {
        return this.routeDest;
    }

    /**
     * The id of the vpn gateway.
     * 
     */
    @Import(name="vpnGatewayId", required=true)
    private Output<String> vpnGatewayId;

    /**
     * @return The id of the vpn gateway.
     * 
     */
    public Output<String> vpnGatewayId() {
        return this.vpnGatewayId;
    }

    /**
     * The value should be 0 or 100.
     * 
     */
    @Import(name="weight", required=true)
    private Output<Integer> weight;

    /**
     * @return The value should be 0 or 100.
     * 
     */
    public Output<Integer> weight() {
        return this.weight;
    }

    private RouteEntryArgs() {}

    private RouteEntryArgs(RouteEntryArgs $) {
        this.nextHop = $.nextHop;
        this.publishVpc = $.publishVpc;
        this.routeDest = $.routeDest;
        this.vpnGatewayId = $.vpnGatewayId;
        this.weight = $.weight;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RouteEntryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RouteEntryArgs $;

        public Builder() {
            $ = new RouteEntryArgs();
        }

        public Builder(RouteEntryArgs defaults) {
            $ = new RouteEntryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param nextHop The next hop of the destination route.
         * 
         * @return builder
         * 
         */
        public Builder nextHop(Output<String> nextHop) {
            $.nextHop = nextHop;
            return this;
        }

        /**
         * @param nextHop The next hop of the destination route.
         * 
         * @return builder
         * 
         */
        public Builder nextHop(String nextHop) {
            return nextHop(Output.of(nextHop));
        }

        /**
         * @param publishVpc Whether to issue the destination route to the VPC.
         * 
         * @return builder
         * 
         */
        public Builder publishVpc(Output<Boolean> publishVpc) {
            $.publishVpc = publishVpc;
            return this;
        }

        /**
         * @param publishVpc Whether to issue the destination route to the VPC.
         * 
         * @return builder
         * 
         */
        public Builder publishVpc(Boolean publishVpc) {
            return publishVpc(Output.of(publishVpc));
        }

        /**
         * @param routeDest The destination network segment of the destination route.
         * 
         * @return builder
         * 
         */
        public Builder routeDest(Output<String> routeDest) {
            $.routeDest = routeDest;
            return this;
        }

        /**
         * @param routeDest The destination network segment of the destination route.
         * 
         * @return builder
         * 
         */
        public Builder routeDest(String routeDest) {
            return routeDest(Output.of(routeDest));
        }

        /**
         * @param vpnGatewayId The id of the vpn gateway.
         * 
         * @return builder
         * 
         */
        public Builder vpnGatewayId(Output<String> vpnGatewayId) {
            $.vpnGatewayId = vpnGatewayId;
            return this;
        }

        /**
         * @param vpnGatewayId The id of the vpn gateway.
         * 
         * @return builder
         * 
         */
        public Builder vpnGatewayId(String vpnGatewayId) {
            return vpnGatewayId(Output.of(vpnGatewayId));
        }

        /**
         * @param weight The value should be 0 or 100.
         * 
         * @return builder
         * 
         */
        public Builder weight(Output<Integer> weight) {
            $.weight = weight;
            return this;
        }

        /**
         * @param weight The value should be 0 or 100.
         * 
         * @return builder
         * 
         */
        public Builder weight(Integer weight) {
            return weight(Output.of(weight));
        }

        public RouteEntryArgs build() {
            $.nextHop = Objects.requireNonNull($.nextHop, "expected parameter 'nextHop' to be non-null");
            $.publishVpc = Objects.requireNonNull($.publishVpc, "expected parameter 'publishVpc' to be non-null");
            $.routeDest = Objects.requireNonNull($.routeDest, "expected parameter 'routeDest' to be non-null");
            $.vpnGatewayId = Objects.requireNonNull($.vpnGatewayId, "expected parameter 'vpnGatewayId' to be non-null");
            $.weight = Objects.requireNonNull($.weight, "expected parameter 'weight' to be non-null");
            return $;
        }
    }

}
