// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpn;

import com.pulumi.alicloud.vpn.inputs.ConnectionBgpConfigArgs;
import com.pulumi.alicloud.vpn.inputs.ConnectionHealthCheckConfigArgs;
import com.pulumi.alicloud.vpn.inputs.ConnectionIkeConfigArgs;
import com.pulumi.alicloud.vpn.inputs.ConnectionIpsecConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionArgs Empty = new ConnectionArgs();

    /**
     * The configurations of the BGP routing protocol. See the following `Block bgp_config`.
     * 
     */
    @Import(name="bgpConfig")
    private @Nullable Output<ConnectionBgpConfigArgs> bgpConfig;

    /**
     * @return The configurations of the BGP routing protocol. See the following `Block bgp_config`.
     * 
     */
    public Optional<Output<ConnectionBgpConfigArgs>> bgpConfig() {
        return Optional.ofNullable(this.bgpConfig);
    }

    /**
     * The ID of the customer gateway.
     * 
     */
    @Import(name="customerGatewayId", required=true)
    private Output<String> customerGatewayId;

    /**
     * @return The ID of the customer gateway.
     * 
     */
    public Output<String> customerGatewayId() {
        return this.customerGatewayId;
    }

    /**
     * Whether to delete a successfully negotiated IPsec tunnel and initiate a negotiation again. Valid value:true,false.
     * 
     */
    @Import(name="effectImmediately")
    private @Nullable Output<Boolean> effectImmediately;

    /**
     * @return Whether to delete a successfully negotiated IPsec tunnel and initiate a negotiation again. Valid value:true,false.
     * 
     */
    public Optional<Output<Boolean>> effectImmediately() {
        return Optional.ofNullable(this.effectImmediately);
    }

    /**
     * Specifies whether to enable the dead peer detection (DPD) feature. Valid values: `true`(default), `false`.
     * 
     */
    @Import(name="enableDpd")
    private @Nullable Output<Boolean> enableDpd;

    /**
     * @return Specifies whether to enable the dead peer detection (DPD) feature. Valid values: `true`(default), `false`.
     * 
     */
    public Optional<Output<Boolean>> enableDpd() {
        return Optional.ofNullable(this.enableDpd);
    }

    /**
     * Specifies whether to enable NAT traversal. Valid values: `true`(default), `false`.
     * 
     */
    @Import(name="enableNatTraversal")
    private @Nullable Output<Boolean> enableNatTraversal;

    /**
     * @return Specifies whether to enable NAT traversal. Valid values: `true`(default), `false`.
     * 
     */
    public Optional<Output<Boolean>> enableNatTraversal() {
        return Optional.ofNullable(this.enableNatTraversal);
    }

    /**
     * The health check configurations. See the following `Block health_check_config`.
     * 
     */
    @Import(name="healthCheckConfig")
    private @Nullable Output<ConnectionHealthCheckConfigArgs> healthCheckConfig;

    /**
     * @return The health check configurations. See the following `Block health_check_config`.
     * 
     */
    public Optional<Output<ConnectionHealthCheckConfigArgs>> healthCheckConfig() {
        return Optional.ofNullable(this.healthCheckConfig);
    }

    /**
     * The configurations of phase-one negotiation. See the following `Block ike_config`.
     * 
     */
    @Import(name="ikeConfig")
    private @Nullable Output<ConnectionIkeConfigArgs> ikeConfig;

    /**
     * @return The configurations of phase-one negotiation. See the following `Block ike_config`.
     * 
     */
    public Optional<Output<ConnectionIkeConfigArgs>> ikeConfig() {
        return Optional.ofNullable(this.ikeConfig);
    }

    /**
     * The configurations of phase-two negotiation. See the following `Block ipsec_config`.
     * 
     */
    @Import(name="ipsecConfig")
    private @Nullable Output<ConnectionIpsecConfigArgs> ipsecConfig;

    /**
     * @return The configurations of phase-two negotiation. See the following `Block ipsec_config`.
     * 
     */
    public Optional<Output<ConnectionIpsecConfigArgs>> ipsecConfig() {
        return Optional.ofNullable(this.ipsecConfig);
    }

    /**
     * The CIDR block of the VPC to be connected with the local data center. This parameter is used for phase-two negotiation.
     * 
     */
    @Import(name="localSubnets", required=true)
    private Output<List<String>> localSubnets;

    /**
     * @return The CIDR block of the VPC to be connected with the local data center. This parameter is used for phase-two negotiation.
     * 
     */
    public Output<List<String>> localSubnets() {
        return this.localSubnets;
    }

    /**
     * The name of the IPsec connection.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the IPsec connection.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The CIDR block of the local data center. This parameter is used for phase-two negotiation.
     * 
     */
    @Import(name="remoteSubnets", required=true)
    private Output<List<String>> remoteSubnets;

    /**
     * @return The CIDR block of the local data center. This parameter is used for phase-two negotiation.
     * 
     */
    public Output<List<String>> remoteSubnets() {
        return this.remoteSubnets;
    }

    /**
     * The ID of the VPN gateway.
     * 
     */
    @Import(name="vpnGatewayId", required=true)
    private Output<String> vpnGatewayId;

    /**
     * @return The ID of the VPN gateway.
     * 
     */
    public Output<String> vpnGatewayId() {
        return this.vpnGatewayId;
    }

    private ConnectionArgs() {}

    private ConnectionArgs(ConnectionArgs $) {
        this.bgpConfig = $.bgpConfig;
        this.customerGatewayId = $.customerGatewayId;
        this.effectImmediately = $.effectImmediately;
        this.enableDpd = $.enableDpd;
        this.enableNatTraversal = $.enableNatTraversal;
        this.healthCheckConfig = $.healthCheckConfig;
        this.ikeConfig = $.ikeConfig;
        this.ipsecConfig = $.ipsecConfig;
        this.localSubnets = $.localSubnets;
        this.name = $.name;
        this.remoteSubnets = $.remoteSubnets;
        this.vpnGatewayId = $.vpnGatewayId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionArgs $;

        public Builder() {
            $ = new ConnectionArgs();
        }

        public Builder(ConnectionArgs defaults) {
            $ = new ConnectionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bgpConfig The configurations of the BGP routing protocol. See the following `Block bgp_config`.
         * 
         * @return builder
         * 
         */
        public Builder bgpConfig(@Nullable Output<ConnectionBgpConfigArgs> bgpConfig) {
            $.bgpConfig = bgpConfig;
            return this;
        }

        /**
         * @param bgpConfig The configurations of the BGP routing protocol. See the following `Block bgp_config`.
         * 
         * @return builder
         * 
         */
        public Builder bgpConfig(ConnectionBgpConfigArgs bgpConfig) {
            return bgpConfig(Output.of(bgpConfig));
        }

        /**
         * @param customerGatewayId The ID of the customer gateway.
         * 
         * @return builder
         * 
         */
        public Builder customerGatewayId(Output<String> customerGatewayId) {
            $.customerGatewayId = customerGatewayId;
            return this;
        }

        /**
         * @param customerGatewayId The ID of the customer gateway.
         * 
         * @return builder
         * 
         */
        public Builder customerGatewayId(String customerGatewayId) {
            return customerGatewayId(Output.of(customerGatewayId));
        }

        /**
         * @param effectImmediately Whether to delete a successfully negotiated IPsec tunnel and initiate a negotiation again. Valid value:true,false.
         * 
         * @return builder
         * 
         */
        public Builder effectImmediately(@Nullable Output<Boolean> effectImmediately) {
            $.effectImmediately = effectImmediately;
            return this;
        }

        /**
         * @param effectImmediately Whether to delete a successfully negotiated IPsec tunnel and initiate a negotiation again. Valid value:true,false.
         * 
         * @return builder
         * 
         */
        public Builder effectImmediately(Boolean effectImmediately) {
            return effectImmediately(Output.of(effectImmediately));
        }

        /**
         * @param enableDpd Specifies whether to enable the dead peer detection (DPD) feature. Valid values: `true`(default), `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableDpd(@Nullable Output<Boolean> enableDpd) {
            $.enableDpd = enableDpd;
            return this;
        }

        /**
         * @param enableDpd Specifies whether to enable the dead peer detection (DPD) feature. Valid values: `true`(default), `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableDpd(Boolean enableDpd) {
            return enableDpd(Output.of(enableDpd));
        }

        /**
         * @param enableNatTraversal Specifies whether to enable NAT traversal. Valid values: `true`(default), `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableNatTraversal(@Nullable Output<Boolean> enableNatTraversal) {
            $.enableNatTraversal = enableNatTraversal;
            return this;
        }

        /**
         * @param enableNatTraversal Specifies whether to enable NAT traversal. Valid values: `true`(default), `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableNatTraversal(Boolean enableNatTraversal) {
            return enableNatTraversal(Output.of(enableNatTraversal));
        }

        /**
         * @param healthCheckConfig The health check configurations. See the following `Block health_check_config`.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckConfig(@Nullable Output<ConnectionHealthCheckConfigArgs> healthCheckConfig) {
            $.healthCheckConfig = healthCheckConfig;
            return this;
        }

        /**
         * @param healthCheckConfig The health check configurations. See the following `Block health_check_config`.
         * 
         * @return builder
         * 
         */
        public Builder healthCheckConfig(ConnectionHealthCheckConfigArgs healthCheckConfig) {
            return healthCheckConfig(Output.of(healthCheckConfig));
        }

        /**
         * @param ikeConfig The configurations of phase-one negotiation. See the following `Block ike_config`.
         * 
         * @return builder
         * 
         */
        public Builder ikeConfig(@Nullable Output<ConnectionIkeConfigArgs> ikeConfig) {
            $.ikeConfig = ikeConfig;
            return this;
        }

        /**
         * @param ikeConfig The configurations of phase-one negotiation. See the following `Block ike_config`.
         * 
         * @return builder
         * 
         */
        public Builder ikeConfig(ConnectionIkeConfigArgs ikeConfig) {
            return ikeConfig(Output.of(ikeConfig));
        }

        /**
         * @param ipsecConfig The configurations of phase-two negotiation. See the following `Block ipsec_config`.
         * 
         * @return builder
         * 
         */
        public Builder ipsecConfig(@Nullable Output<ConnectionIpsecConfigArgs> ipsecConfig) {
            $.ipsecConfig = ipsecConfig;
            return this;
        }

        /**
         * @param ipsecConfig The configurations of phase-two negotiation. See the following `Block ipsec_config`.
         * 
         * @return builder
         * 
         */
        public Builder ipsecConfig(ConnectionIpsecConfigArgs ipsecConfig) {
            return ipsecConfig(Output.of(ipsecConfig));
        }

        /**
         * @param localSubnets The CIDR block of the VPC to be connected with the local data center. This parameter is used for phase-two negotiation.
         * 
         * @return builder
         * 
         */
        public Builder localSubnets(Output<List<String>> localSubnets) {
            $.localSubnets = localSubnets;
            return this;
        }

        /**
         * @param localSubnets The CIDR block of the VPC to be connected with the local data center. This parameter is used for phase-two negotiation.
         * 
         * @return builder
         * 
         */
        public Builder localSubnets(List<String> localSubnets) {
            return localSubnets(Output.of(localSubnets));
        }

        /**
         * @param localSubnets The CIDR block of the VPC to be connected with the local data center. This parameter is used for phase-two negotiation.
         * 
         * @return builder
         * 
         */
        public Builder localSubnets(String... localSubnets) {
            return localSubnets(List.of(localSubnets));
        }

        /**
         * @param name The name of the IPsec connection.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the IPsec connection.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param remoteSubnets The CIDR block of the local data center. This parameter is used for phase-two negotiation.
         * 
         * @return builder
         * 
         */
        public Builder remoteSubnets(Output<List<String>> remoteSubnets) {
            $.remoteSubnets = remoteSubnets;
            return this;
        }

        /**
         * @param remoteSubnets The CIDR block of the local data center. This parameter is used for phase-two negotiation.
         * 
         * @return builder
         * 
         */
        public Builder remoteSubnets(List<String> remoteSubnets) {
            return remoteSubnets(Output.of(remoteSubnets));
        }

        /**
         * @param remoteSubnets The CIDR block of the local data center. This parameter is used for phase-two negotiation.
         * 
         * @return builder
         * 
         */
        public Builder remoteSubnets(String... remoteSubnets) {
            return remoteSubnets(List.of(remoteSubnets));
        }

        /**
         * @param vpnGatewayId The ID of the VPN gateway.
         * 
         * @return builder
         * 
         */
        public Builder vpnGatewayId(Output<String> vpnGatewayId) {
            $.vpnGatewayId = vpnGatewayId;
            return this;
        }

        /**
         * @param vpnGatewayId The ID of the VPN gateway.
         * 
         * @return builder
         * 
         */
        public Builder vpnGatewayId(String vpnGatewayId) {
            return vpnGatewayId(Output.of(vpnGatewayId));
        }

        public ConnectionArgs build() {
            $.customerGatewayId = Objects.requireNonNull($.customerGatewayId, "expected parameter 'customerGatewayId' to be non-null");
            $.localSubnets = Objects.requireNonNull($.localSubnets, "expected parameter 'localSubnets' to be non-null");
            $.remoteSubnets = Objects.requireNonNull($.remoteSubnets, "expected parameter 'remoteSubnets' to be non-null");
            $.vpnGatewayId = Objects.requireNonNull($.vpnGatewayId, "expected parameter 'vpnGatewayId' to be non-null");
            return $;
        }
    }

}
