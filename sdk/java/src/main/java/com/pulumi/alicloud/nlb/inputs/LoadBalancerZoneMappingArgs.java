// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nlb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LoadBalancerZoneMappingArgs extends com.pulumi.resources.ResourceArgs {

    public static final LoadBalancerZoneMappingArgs Empty = new LoadBalancerZoneMappingArgs();

    /**
     * The ID of the EIP associated with the Internet-facing NLB instance.
     * 
     */
    @Import(name="allocationId")
    private @Nullable Output<String> allocationId;

    /**
     * @return The ID of the EIP associated with the Internet-facing NLB instance.
     * 
     */
    public Optional<Output<String>> allocationId() {
        return Optional.ofNullable(this.allocationId);
    }

    /**
     * The ID of the elastic network interface (ENI).
     * 
     */
    @Import(name="eniId")
    private @Nullable Output<String> eniId;

    /**
     * @return The ID of the elastic network interface (ENI).
     * 
     */
    public Optional<Output<String>> eniId() {
        return Optional.ofNullable(this.eniId);
    }

    /**
     * The IPv6 address of the NLB instance.
     * 
     */
    @Import(name="ipv6Address")
    private @Nullable Output<String> ipv6Address;

    /**
     * @return The IPv6 address of the NLB instance.
     * 
     */
    public Optional<Output<String>> ipv6Address() {
        return Optional.ofNullable(this.ipv6Address);
    }

    /**
     * The private IPv4 address of the NLB instance.
     * 
     */
    @Import(name="privateIpv4Address")
    private @Nullable Output<String> privateIpv4Address;

    /**
     * @return The private IPv4 address of the NLB instance.
     * 
     */
    public Optional<Output<String>> privateIpv4Address() {
        return Optional.ofNullable(this.privateIpv4Address);
    }

    /**
     * The public IPv4 address of the NLB instance.
     * 
     */
    @Import(name="publicIpv4Address")
    private @Nullable Output<String> publicIpv4Address;

    /**
     * @return The public IPv4 address of the NLB instance.
     * 
     */
    public Optional<Output<String>> publicIpv4Address() {
        return Optional.ofNullable(this.publicIpv4Address);
    }

    /**
     * The vSwitch in the zone. You can specify only one vSwitch (subnet) in each zone of an NLB instance.
     * 
     */
    @Import(name="vswitchId", required=true)
    private Output<String> vswitchId;

    /**
     * @return The vSwitch in the zone. You can specify only one vSwitch (subnet) in each zone of an NLB instance.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }

    /**
     * The ID of the zone of the NLB instance.
     * 
     */
    @Import(name="zoneId", required=true)
    private Output<String> zoneId;

    /**
     * @return The ID of the zone of the NLB instance.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    private LoadBalancerZoneMappingArgs() {}

    private LoadBalancerZoneMappingArgs(LoadBalancerZoneMappingArgs $) {
        this.allocationId = $.allocationId;
        this.eniId = $.eniId;
        this.ipv6Address = $.ipv6Address;
        this.privateIpv4Address = $.privateIpv4Address;
        this.publicIpv4Address = $.publicIpv4Address;
        this.vswitchId = $.vswitchId;
        this.zoneId = $.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LoadBalancerZoneMappingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LoadBalancerZoneMappingArgs $;

        public Builder() {
            $ = new LoadBalancerZoneMappingArgs();
        }

        public Builder(LoadBalancerZoneMappingArgs defaults) {
            $ = new LoadBalancerZoneMappingArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allocationId The ID of the EIP associated with the Internet-facing NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder allocationId(@Nullable Output<String> allocationId) {
            $.allocationId = allocationId;
            return this;
        }

        /**
         * @param allocationId The ID of the EIP associated with the Internet-facing NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder allocationId(String allocationId) {
            return allocationId(Output.of(allocationId));
        }

        /**
         * @param eniId The ID of the elastic network interface (ENI).
         * 
         * @return builder
         * 
         */
        public Builder eniId(@Nullable Output<String> eniId) {
            $.eniId = eniId;
            return this;
        }

        /**
         * @param eniId The ID of the elastic network interface (ENI).
         * 
         * @return builder
         * 
         */
        public Builder eniId(String eniId) {
            return eniId(Output.of(eniId));
        }

        /**
         * @param ipv6Address The IPv6 address of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder ipv6Address(@Nullable Output<String> ipv6Address) {
            $.ipv6Address = ipv6Address;
            return this;
        }

        /**
         * @param ipv6Address The IPv6 address of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder ipv6Address(String ipv6Address) {
            return ipv6Address(Output.of(ipv6Address));
        }

        /**
         * @param privateIpv4Address The private IPv4 address of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder privateIpv4Address(@Nullable Output<String> privateIpv4Address) {
            $.privateIpv4Address = privateIpv4Address;
            return this;
        }

        /**
         * @param privateIpv4Address The private IPv4 address of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder privateIpv4Address(String privateIpv4Address) {
            return privateIpv4Address(Output.of(privateIpv4Address));
        }

        /**
         * @param publicIpv4Address The public IPv4 address of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder publicIpv4Address(@Nullable Output<String> publicIpv4Address) {
            $.publicIpv4Address = publicIpv4Address;
            return this;
        }

        /**
         * @param publicIpv4Address The public IPv4 address of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder publicIpv4Address(String publicIpv4Address) {
            return publicIpv4Address(Output.of(publicIpv4Address));
        }

        /**
         * @param vswitchId The vSwitch in the zone. You can specify only one vSwitch (subnet) in each zone of an NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId The vSwitch in the zone. You can specify only one vSwitch (subnet) in each zone of an NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        /**
         * @param zoneId The ID of the zone of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        /**
         * @param zoneId The ID of the zone of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public LoadBalancerZoneMappingArgs build() {
            $.vswitchId = Objects.requireNonNull($.vswitchId, "expected parameter 'vswitchId' to be non-null");
            $.zoneId = Objects.requireNonNull($.zoneId, "expected parameter 'zoneId' to be non-null");
            return $;
        }
    }

}
