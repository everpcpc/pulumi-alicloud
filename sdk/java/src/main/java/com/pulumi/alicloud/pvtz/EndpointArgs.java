// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.pvtz;

import com.pulumi.alicloud.pvtz.inputs.EndpointIpConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class EndpointArgs extends com.pulumi.resources.ResourceArgs {

    public static final EndpointArgs Empty = new EndpointArgs();

    /**
     * The name of the resource.
     * 
     */
    @Import(name="endpointName", required=true)
    private Output<String> endpointName;

    /**
     * @return The name of the resource.
     * 
     */
    public Output<String> endpointName() {
        return this.endpointName;
    }

    /**
     * The Ip Configs. See the following `Block ip_configs`. **NOTE:** In order to ensure high availability, add at least 2 and up to 6.
     * 
     */
    @Import(name="ipConfigs", required=true)
    private Output<List<EndpointIpConfigArgs>> ipConfigs;

    /**
     * @return The Ip Configs. See the following `Block ip_configs`. **NOTE:** In order to ensure high availability, add at least 2 and up to 6.
     * 
     */
    public Output<List<EndpointIpConfigArgs>> ipConfigs() {
        return this.ipConfigs;
    }

    /**
     * The ID of the Security Group.
     * 
     */
    @Import(name="securityGroupId", required=true)
    private Output<String> securityGroupId;

    /**
     * @return The ID of the Security Group.
     * 
     */
    public Output<String> securityGroupId() {
        return this.securityGroupId;
    }

    /**
     * The VPC ID.
     * 
     */
    @Import(name="vpcId", required=true)
    private Output<String> vpcId;

    /**
     * @return The VPC ID.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     * The Region of the VPC.
     * 
     */
    @Import(name="vpcRegionId", required=true)
    private Output<String> vpcRegionId;

    /**
     * @return The Region of the VPC.
     * 
     */
    public Output<String> vpcRegionId() {
        return this.vpcRegionId;
    }

    private EndpointArgs() {}

    private EndpointArgs(EndpointArgs $) {
        this.endpointName = $.endpointName;
        this.ipConfigs = $.ipConfigs;
        this.securityGroupId = $.securityGroupId;
        this.vpcId = $.vpcId;
        this.vpcRegionId = $.vpcRegionId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EndpointArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EndpointArgs $;

        public Builder() {
            $ = new EndpointArgs();
        }

        public Builder(EndpointArgs defaults) {
            $ = new EndpointArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param endpointName The name of the resource.
         * 
         * @return builder
         * 
         */
        public Builder endpointName(Output<String> endpointName) {
            $.endpointName = endpointName;
            return this;
        }

        /**
         * @param endpointName The name of the resource.
         * 
         * @return builder
         * 
         */
        public Builder endpointName(String endpointName) {
            return endpointName(Output.of(endpointName));
        }

        /**
         * @param ipConfigs The Ip Configs. See the following `Block ip_configs`. **NOTE:** In order to ensure high availability, add at least 2 and up to 6.
         * 
         * @return builder
         * 
         */
        public Builder ipConfigs(Output<List<EndpointIpConfigArgs>> ipConfigs) {
            $.ipConfigs = ipConfigs;
            return this;
        }

        /**
         * @param ipConfigs The Ip Configs. See the following `Block ip_configs`. **NOTE:** In order to ensure high availability, add at least 2 and up to 6.
         * 
         * @return builder
         * 
         */
        public Builder ipConfigs(List<EndpointIpConfigArgs> ipConfigs) {
            return ipConfigs(Output.of(ipConfigs));
        }

        /**
         * @param ipConfigs The Ip Configs. See the following `Block ip_configs`. **NOTE:** In order to ensure high availability, add at least 2 and up to 6.
         * 
         * @return builder
         * 
         */
        public Builder ipConfigs(EndpointIpConfigArgs... ipConfigs) {
            return ipConfigs(List.of(ipConfigs));
        }

        /**
         * @param securityGroupId The ID of the Security Group.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupId(Output<String> securityGroupId) {
            $.securityGroupId = securityGroupId;
            return this;
        }

        /**
         * @param securityGroupId The ID of the Security Group.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupId(String securityGroupId) {
            return securityGroupId(Output.of(securityGroupId));
        }

        /**
         * @param vpcId The VPC ID.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The VPC ID.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param vpcRegionId The Region of the VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcRegionId(Output<String> vpcRegionId) {
            $.vpcRegionId = vpcRegionId;
            return this;
        }

        /**
         * @param vpcRegionId The Region of the VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcRegionId(String vpcRegionId) {
            return vpcRegionId(Output.of(vpcRegionId));
        }

        public EndpointArgs build() {
            $.endpointName = Objects.requireNonNull($.endpointName, "expected parameter 'endpointName' to be non-null");
            $.ipConfigs = Objects.requireNonNull($.ipConfigs, "expected parameter 'ipConfigs' to be non-null");
            $.securityGroupId = Objects.requireNonNull($.securityGroupId, "expected parameter 'securityGroupId' to be non-null");
            $.vpcId = Objects.requireNonNull($.vpcId, "expected parameter 'vpcId' to be non-null");
            $.vpcRegionId = Objects.requireNonNull($.vpcRegionId, "expected parameter 'vpcRegionId' to be non-null");
            return $;
        }
    }

}
