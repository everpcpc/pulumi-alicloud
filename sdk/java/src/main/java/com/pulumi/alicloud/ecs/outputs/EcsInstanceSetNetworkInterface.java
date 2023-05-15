// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class EcsInstanceSetNetworkInterface {
    /**
     * @return The description of ENI.
     * 
     */
    private @Nullable String description;
    /**
     * @return The name of ENI.
     * 
     */
    private @Nullable String networkInterfaceName;
    /**
     * @return The primary private IP address of ENI.
     * 
     */
    private @Nullable String primaryIpAddress;
    /**
     * @return The ID of the security group to which to assign secondary ENI.
     * 
     */
    private String securityGroupId;
    /**
     * @return The ID of the vSwitch to which to connect ENI.
     * 
     */
    private @Nullable String vswitchId;

    private EcsInstanceSetNetworkInterface() {}
    /**
     * @return The description of ENI.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return The name of ENI.
     * 
     */
    public Optional<String> networkInterfaceName() {
        return Optional.ofNullable(this.networkInterfaceName);
    }
    /**
     * @return The primary private IP address of ENI.
     * 
     */
    public Optional<String> primaryIpAddress() {
        return Optional.ofNullable(this.primaryIpAddress);
    }
    /**
     * @return The ID of the security group to which to assign secondary ENI.
     * 
     */
    public String securityGroupId() {
        return this.securityGroupId;
    }
    /**
     * @return The ID of the vSwitch to which to connect ENI.
     * 
     */
    public Optional<String> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(EcsInstanceSetNetworkInterface defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String description;
        private @Nullable String networkInterfaceName;
        private @Nullable String primaryIpAddress;
        private String securityGroupId;
        private @Nullable String vswitchId;
        public Builder() {}
        public Builder(EcsInstanceSetNetworkInterface defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.networkInterfaceName = defaults.networkInterfaceName;
    	      this.primaryIpAddress = defaults.primaryIpAddress;
    	      this.securityGroupId = defaults.securityGroupId;
    	      this.vswitchId = defaults.vswitchId;
        }

        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder networkInterfaceName(@Nullable String networkInterfaceName) {
            this.networkInterfaceName = networkInterfaceName;
            return this;
        }
        @CustomType.Setter
        public Builder primaryIpAddress(@Nullable String primaryIpAddress) {
            this.primaryIpAddress = primaryIpAddress;
            return this;
        }
        @CustomType.Setter
        public Builder securityGroupId(String securityGroupId) {
            this.securityGroupId = Objects.requireNonNull(securityGroupId);
            return this;
        }
        @CustomType.Setter
        public Builder vswitchId(@Nullable String vswitchId) {
            this.vswitchId = vswitchId;
            return this;
        }
        public EcsInstanceSetNetworkInterface build() {
            final var o = new EcsInstanceSetNetworkInterface();
            o.description = description;
            o.networkInterfaceName = networkInterfaceName;
            o.primaryIpAddress = primaryIpAddress;
            o.securityGroupId = securityGroupId;
            o.vswitchId = vswitchId;
            return o;
        }
    }
}
