// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nlb.inputs;

import com.pulumi.alicloud.nlb.inputs.LoadBalancerZoneMappingArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LoadBalancerState extends com.pulumi.resources.ResourceArgs {

    public static final LoadBalancerState Empty = new LoadBalancerState();

    /**
     * The protocol version. Valid values:
     * - ipv4 (default): IPv4
     * - DualStack: dual stack
     * 
     */
    @Import(name="addressIpVersion")
    private @Nullable Output<String> addressIpVersion;

    /**
     * @return The protocol version. Valid values:
     * - ipv4 (default): IPv4
     * - DualStack: dual stack
     * 
     */
    public Optional<Output<String>> addressIpVersion() {
        return Optional.ofNullable(this.addressIpVersion);
    }

    /**
     * The type of IPv4 address used by the NLB instance. Valid values:
     * - Internet: The NLB instance uses a public IP address. The domain name of the NLB instance is resolved to the public IP address. Therefore, the NLB instance can be accessed over the Internet.
     * - Intranet: The NLB instance uses a private IP address. The domain name of the NLB instance is resolved to the private IP address. Therefore, the NLB instance can be accessed over the virtual private cloud (VPC) where the NLB instance is deployed.
     * 
     */
    @Import(name="addressType")
    private @Nullable Output<String> addressType;

    /**
     * @return The type of IPv4 address used by the NLB instance. Valid values:
     * - Internet: The NLB instance uses a public IP address. The domain name of the NLB instance is resolved to the public IP address. Therefore, the NLB instance can be accessed over the Internet.
     * - Intranet: The NLB instance uses a private IP address. The domain name of the NLB instance is resolved to the private IP address. Therefore, the NLB instance can be accessed over the virtual private cloud (VPC) where the NLB instance is deployed.
     * 
     */
    public Optional<Output<String>> addressType() {
        return Optional.ofNullable(this.addressType);
    }

    /**
     * The ID of the EIP bandwidth plan that is associated with the NLB instance if the NLB instance uses a public IP address.
     * 
     */
    @Import(name="bandwidthPackageId")
    private @Nullable Output<String> bandwidthPackageId;

    /**
     * @return The ID of the EIP bandwidth plan that is associated with the NLB instance if the NLB instance uses a public IP address.
     * 
     */
    public Optional<Output<String>> bandwidthPackageId() {
        return Optional.ofNullable(this.bandwidthPackageId);
    }

    /**
     * The time when the resource was created. The time is displayed in UTC in `yyyy-MM-ddTHH:mm:ssZ` format.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The time when the resource was created. The time is displayed in UTC in `yyyy-MM-ddTHH:mm:ssZ` format.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * Specifies whether to enable cross-zone load balancing for the NLB instance.
     * 
     */
    @Import(name="crossZoneEnabled")
    private @Nullable Output<Boolean> crossZoneEnabled;

    /**
     * @return Specifies whether to enable cross-zone load balancing for the NLB instance.
     * 
     */
    public Optional<Output<Boolean>> crossZoneEnabled() {
        return Optional.ofNullable(this.crossZoneEnabled);
    }

    /**
     * The domain name of the NLB instance.
     * 
     */
    @Import(name="dnsName")
    private @Nullable Output<String> dnsName;

    /**
     * @return The domain name of the NLB instance.
     * 
     */
    public Optional<Output<String>> dnsName() {
        return Optional.ofNullable(this.dnsName);
    }

    /**
     * The type of IPv6 address used by the NLB instance.
     * 
     */
    @Import(name="ipv6AddressType")
    private @Nullable Output<String> ipv6AddressType;

    /**
     * @return The type of IPv6 address used by the NLB instance.
     * 
     */
    public Optional<Output<String>> ipv6AddressType() {
        return Optional.ofNullable(this.ipv6AddressType);
    }

    /**
     * The business status of the NLB instance.
     * 
     */
    @Import(name="loadBalancerBusinessStatus")
    private @Nullable Output<String> loadBalancerBusinessStatus;

    /**
     * @return The business status of the NLB instance.
     * 
     */
    public Optional<Output<String>> loadBalancerBusinessStatus() {
        return Optional.ofNullable(this.loadBalancerBusinessStatus);
    }

    /**
     * The name of the NLB instance. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    @Import(name="loadBalancerName")
    private @Nullable Output<String> loadBalancerName;

    /**
     * @return The name of the NLB instance. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    public Optional<Output<String>> loadBalancerName() {
        return Optional.ofNullable(this.loadBalancerName);
    }

    /**
     * The type of the instance. Set the value to `Network`, which specifies an NLB instance.
     * 
     */
    @Import(name="loadBalancerType")
    private @Nullable Output<String> loadBalancerType;

    /**
     * @return The type of the instance. Set the value to `Network`, which specifies an NLB instance.
     * 
     */
    public Optional<Output<String>> loadBalancerType() {
        return Optional.ofNullable(this.loadBalancerType);
    }

    /**
     * The ID of the resource group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The status of the NLB instance.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the NLB instance.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The ID of the VPC where the NLB instance is deployed.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The ID of the VPC where the NLB instance is deployed.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * Available Area Configuration List. You must add at least two zones. You can add a maximum of 10 zones. See the following `Block zone_mappings`.
     * 
     */
    @Import(name="zoneMappings")
    private @Nullable Output<List<LoadBalancerZoneMappingArgs>> zoneMappings;

    /**
     * @return Available Area Configuration List. You must add at least two zones. You can add a maximum of 10 zones. See the following `Block zone_mappings`.
     * 
     */
    public Optional<Output<List<LoadBalancerZoneMappingArgs>>> zoneMappings() {
        return Optional.ofNullable(this.zoneMappings);
    }

    private LoadBalancerState() {}

    private LoadBalancerState(LoadBalancerState $) {
        this.addressIpVersion = $.addressIpVersion;
        this.addressType = $.addressType;
        this.bandwidthPackageId = $.bandwidthPackageId;
        this.createTime = $.createTime;
        this.crossZoneEnabled = $.crossZoneEnabled;
        this.dnsName = $.dnsName;
        this.ipv6AddressType = $.ipv6AddressType;
        this.loadBalancerBusinessStatus = $.loadBalancerBusinessStatus;
        this.loadBalancerName = $.loadBalancerName;
        this.loadBalancerType = $.loadBalancerType;
        this.resourceGroupId = $.resourceGroupId;
        this.status = $.status;
        this.tags = $.tags;
        this.vpcId = $.vpcId;
        this.zoneMappings = $.zoneMappings;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LoadBalancerState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LoadBalancerState $;

        public Builder() {
            $ = new LoadBalancerState();
        }

        public Builder(LoadBalancerState defaults) {
            $ = new LoadBalancerState(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressIpVersion The protocol version. Valid values:
         * - ipv4 (default): IPv4
         * - DualStack: dual stack
         * 
         * @return builder
         * 
         */
        public Builder addressIpVersion(@Nullable Output<String> addressIpVersion) {
            $.addressIpVersion = addressIpVersion;
            return this;
        }

        /**
         * @param addressIpVersion The protocol version. Valid values:
         * - ipv4 (default): IPv4
         * - DualStack: dual stack
         * 
         * @return builder
         * 
         */
        public Builder addressIpVersion(String addressIpVersion) {
            return addressIpVersion(Output.of(addressIpVersion));
        }

        /**
         * @param addressType The type of IPv4 address used by the NLB instance. Valid values:
         * - Internet: The NLB instance uses a public IP address. The domain name of the NLB instance is resolved to the public IP address. Therefore, the NLB instance can be accessed over the Internet.
         * - Intranet: The NLB instance uses a private IP address. The domain name of the NLB instance is resolved to the private IP address. Therefore, the NLB instance can be accessed over the virtual private cloud (VPC) where the NLB instance is deployed.
         * 
         * @return builder
         * 
         */
        public Builder addressType(@Nullable Output<String> addressType) {
            $.addressType = addressType;
            return this;
        }

        /**
         * @param addressType The type of IPv4 address used by the NLB instance. Valid values:
         * - Internet: The NLB instance uses a public IP address. The domain name of the NLB instance is resolved to the public IP address. Therefore, the NLB instance can be accessed over the Internet.
         * - Intranet: The NLB instance uses a private IP address. The domain name of the NLB instance is resolved to the private IP address. Therefore, the NLB instance can be accessed over the virtual private cloud (VPC) where the NLB instance is deployed.
         * 
         * @return builder
         * 
         */
        public Builder addressType(String addressType) {
            return addressType(Output.of(addressType));
        }

        /**
         * @param bandwidthPackageId The ID of the EIP bandwidth plan that is associated with the NLB instance if the NLB instance uses a public IP address.
         * 
         * @return builder
         * 
         */
        public Builder bandwidthPackageId(@Nullable Output<String> bandwidthPackageId) {
            $.bandwidthPackageId = bandwidthPackageId;
            return this;
        }

        /**
         * @param bandwidthPackageId The ID of the EIP bandwidth plan that is associated with the NLB instance if the NLB instance uses a public IP address.
         * 
         * @return builder
         * 
         */
        public Builder bandwidthPackageId(String bandwidthPackageId) {
            return bandwidthPackageId(Output.of(bandwidthPackageId));
        }

        /**
         * @param createTime The time when the resource was created. The time is displayed in UTC in `yyyy-MM-ddTHH:mm:ssZ` format.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The time when the resource was created. The time is displayed in UTC in `yyyy-MM-ddTHH:mm:ssZ` format.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param crossZoneEnabled Specifies whether to enable cross-zone load balancing for the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder crossZoneEnabled(@Nullable Output<Boolean> crossZoneEnabled) {
            $.crossZoneEnabled = crossZoneEnabled;
            return this;
        }

        /**
         * @param crossZoneEnabled Specifies whether to enable cross-zone load balancing for the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder crossZoneEnabled(Boolean crossZoneEnabled) {
            return crossZoneEnabled(Output.of(crossZoneEnabled));
        }

        /**
         * @param dnsName The domain name of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder dnsName(@Nullable Output<String> dnsName) {
            $.dnsName = dnsName;
            return this;
        }

        /**
         * @param dnsName The domain name of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder dnsName(String dnsName) {
            return dnsName(Output.of(dnsName));
        }

        /**
         * @param ipv6AddressType The type of IPv6 address used by the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder ipv6AddressType(@Nullable Output<String> ipv6AddressType) {
            $.ipv6AddressType = ipv6AddressType;
            return this;
        }

        /**
         * @param ipv6AddressType The type of IPv6 address used by the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder ipv6AddressType(String ipv6AddressType) {
            return ipv6AddressType(Output.of(ipv6AddressType));
        }

        /**
         * @param loadBalancerBusinessStatus The business status of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerBusinessStatus(@Nullable Output<String> loadBalancerBusinessStatus) {
            $.loadBalancerBusinessStatus = loadBalancerBusinessStatus;
            return this;
        }

        /**
         * @param loadBalancerBusinessStatus The business status of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerBusinessStatus(String loadBalancerBusinessStatus) {
            return loadBalancerBusinessStatus(Output.of(loadBalancerBusinessStatus));
        }

        /**
         * @param loadBalancerName The name of the NLB instance. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerName(@Nullable Output<String> loadBalancerName) {
            $.loadBalancerName = loadBalancerName;
            return this;
        }

        /**
         * @param loadBalancerName The name of the NLB instance. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerName(String loadBalancerName) {
            return loadBalancerName(Output.of(loadBalancerName));
        }

        /**
         * @param loadBalancerType The type of the instance. Set the value to `Network`, which specifies an NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerType(@Nullable Output<String> loadBalancerType) {
            $.loadBalancerType = loadBalancerType;
            return this;
        }

        /**
         * @param loadBalancerType The type of the instance. Set the value to `Network`, which specifies an NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerType(String loadBalancerType) {
            return loadBalancerType(Output.of(loadBalancerType));
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param status The status of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the NLB instance.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param vpcId The ID of the VPC where the NLB instance is deployed.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The ID of the VPC where the NLB instance is deployed.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param zoneMappings Available Area Configuration List. You must add at least two zones. You can add a maximum of 10 zones. See the following `Block zone_mappings`.
         * 
         * @return builder
         * 
         */
        public Builder zoneMappings(@Nullable Output<List<LoadBalancerZoneMappingArgs>> zoneMappings) {
            $.zoneMappings = zoneMappings;
            return this;
        }

        /**
         * @param zoneMappings Available Area Configuration List. You must add at least two zones. You can add a maximum of 10 zones. See the following `Block zone_mappings`.
         * 
         * @return builder
         * 
         */
        public Builder zoneMappings(List<LoadBalancerZoneMappingArgs> zoneMappings) {
            return zoneMappings(Output.of(zoneMappings));
        }

        /**
         * @param zoneMappings Available Area Configuration List. You must add at least two zones. You can add a maximum of 10 zones. See the following `Block zone_mappings`.
         * 
         * @return builder
         * 
         */
        public Builder zoneMappings(LoadBalancerZoneMappingArgs... zoneMappings) {
            return zoneMappings(List.of(zoneMappings));
        }

        public LoadBalancerState build() {
            return $;
        }
    }

}
