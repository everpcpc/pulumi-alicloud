// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEcsNetworkInterfacesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEcsNetworkInterfacesPlainArgs Empty = new GetEcsNetworkInterfacesPlainArgs();

    /**
     * A list of Network Interface IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Network Interface IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The instance id.
     * 
     */
    @Import(name="instanceId")
    private @Nullable String instanceId;

    /**
     * @return The instance id.
     * 
     */
    public Optional<String> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * The network interface name.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.123.1. New field &#39;network_interface_name&#39; instead
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The network interface name.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.123.1. New field &#39;network_interface_name&#39; instead
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A regex string to filter results by Network Interface name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Network Interface name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * The network interface name.
     * 
     */
    @Import(name="networkInterfaceName")
    private @Nullable String networkInterfaceName;

    /**
     * @return The network interface name.
     * 
     */
    public Optional<String> networkInterfaceName() {
        return Optional.ofNullable(this.networkInterfaceName);
    }

    @Import(name="outputFile")
    private @Nullable String outputFile;

    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The primary private IP address of the ENI.
     * 
     */
    @Import(name="primaryIpAddress")
    private @Nullable String primaryIpAddress;

    /**
     * @return The primary private IP address of the ENI.
     * 
     */
    public Optional<String> primaryIpAddress() {
        return Optional.ofNullable(this.primaryIpAddress);
    }

    /**
     * The primary private IP address of the ENI.
     * 
     * @deprecated
     * Field &#39;private_ip&#39; has been deprecated from provider version 1.123.1. New field &#39;primary_ip_address&#39; instead
     * 
     */
    @Deprecated /* Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead */
    @Import(name="privateIp")
    private @Nullable String privateIp;

    /**
     * @return The primary private IP address of the ENI.
     * 
     * @deprecated
     * Field &#39;private_ip&#39; has been deprecated from provider version 1.123.1. New field &#39;primary_ip_address&#39; instead
     * 
     */
    @Deprecated /* Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead */
    public Optional<String> privateIp() {
        return Optional.ofNullable(this.privateIp);
    }

    /**
     * The resource group id.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable String resourceGroupId;

    /**
     * @return The resource group id.
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The security group id.
     * 
     */
    @Import(name="securityGroupId")
    private @Nullable String securityGroupId;

    /**
     * @return The security group id.
     * 
     */
    public Optional<String> securityGroupId() {
        return Optional.ofNullable(this.securityGroupId);
    }

    /**
     * Whether the user of the elastic network card is a cloud product or a virtual vendor.
     * 
     */
    @Import(name="serviceManaged")
    private @Nullable Boolean serviceManaged;

    /**
     * @return Whether the user of the elastic network card is a cloud product or a virtual vendor.
     * 
     */
    public Optional<Boolean> serviceManaged() {
        return Optional.ofNullable(this.serviceManaged);
    }

    /**
     * The status of the ENI.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the ENI.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The tags.
     * 
     */
    @Import(name="tags")
    private @Nullable Map<String,Object> tags;

    /**
     * @return The tags.
     * 
     */
    public Optional<Map<String,Object>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The type of the ENI.
     * 
     */
    @Import(name="type")
    private @Nullable String type;

    /**
     * @return The type of the ENI.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * The Vpc Id.
     * 
     */
    @Import(name="vpcId")
    private @Nullable String vpcId;

    /**
     * @return The Vpc Id.
     * 
     */
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * The vswitch id.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable String vswitchId;

    /**
     * @return The vswitch id.
     * 
     */
    public Optional<String> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    private GetEcsNetworkInterfacesPlainArgs() {}

    private GetEcsNetworkInterfacesPlainArgs(GetEcsNetworkInterfacesPlainArgs $) {
        this.ids = $.ids;
        this.instanceId = $.instanceId;
        this.name = $.name;
        this.nameRegex = $.nameRegex;
        this.networkInterfaceName = $.networkInterfaceName;
        this.outputFile = $.outputFile;
        this.primaryIpAddress = $.primaryIpAddress;
        this.privateIp = $.privateIp;
        this.resourceGroupId = $.resourceGroupId;
        this.securityGroupId = $.securityGroupId;
        this.serviceManaged = $.serviceManaged;
        this.status = $.status;
        this.tags = $.tags;
        this.type = $.type;
        this.vpcId = $.vpcId;
        this.vswitchId = $.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEcsNetworkInterfacesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEcsNetworkInterfacesPlainArgs $;

        public Builder() {
            $ = new GetEcsNetworkInterfacesPlainArgs();
        }

        public Builder(GetEcsNetworkInterfacesPlainArgs defaults) {
            $ = new GetEcsNetworkInterfacesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Network Interface IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Network Interface IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceId The instance id.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable String instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param name The network interface name.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from provider version 1.123.1. New field &#39;network_interface_name&#39; instead
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from provider version 1.123.1. New field 'network_interface_name' instead */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Network Interface name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param networkInterfaceName The network interface name.
         * 
         * @return builder
         * 
         */
        public Builder networkInterfaceName(@Nullable String networkInterfaceName) {
            $.networkInterfaceName = networkInterfaceName;
            return this;
        }

        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param primaryIpAddress The primary private IP address of the ENI.
         * 
         * @return builder
         * 
         */
        public Builder primaryIpAddress(@Nullable String primaryIpAddress) {
            $.primaryIpAddress = primaryIpAddress;
            return this;
        }

        /**
         * @param privateIp The primary private IP address of the ENI.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;private_ip&#39; has been deprecated from provider version 1.123.1. New field &#39;primary_ip_address&#39; instead
         * 
         */
        @Deprecated /* Field 'private_ip' has been deprecated from provider version 1.123.1. New field 'primary_ip_address' instead */
        public Builder privateIp(@Nullable String privateIp) {
            $.privateIp = privateIp;
            return this;
        }

        /**
         * @param resourceGroupId The resource group id.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param securityGroupId The security group id.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupId(@Nullable String securityGroupId) {
            $.securityGroupId = securityGroupId;
            return this;
        }

        /**
         * @param serviceManaged Whether the user of the elastic network card is a cloud product or a virtual vendor.
         * 
         * @return builder
         * 
         */
        public Builder serviceManaged(@Nullable Boolean serviceManaged) {
            $.serviceManaged = serviceManaged;
            return this;
        }

        /**
         * @param status The status of the ENI.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param tags The tags.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Map<String,Object> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param type The type of the ENI.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable String type) {
            $.type = type;
            return this;
        }

        /**
         * @param vpcId The Vpc Id.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable String vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vswitchId The vswitch id.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable String vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        public GetEcsNetworkInterfacesPlainArgs build() {
            return $;
        }
    }

}