// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class PublicIpAddressPoolCidrBlockArgs extends com.pulumi.resources.ResourceArgs {

    public static final PublicIpAddressPoolCidrBlockArgs Empty = new PublicIpAddressPoolCidrBlockArgs();

    /**
     * The CIDR block.
     * 
     */
    @Import(name="cidrBlock", required=true)
    private Output<String> cidrBlock;

    /**
     * @return The CIDR block.
     * 
     */
    public Output<String> cidrBlock() {
        return this.cidrBlock;
    }

    /**
     * The ID of the VPC Public IP address pool.
     * 
     */
    @Import(name="publicIpAddressPoolId", required=true)
    private Output<String> publicIpAddressPoolId;

    /**
     * @return The ID of the VPC Public IP address pool.
     * 
     */
    public Output<String> publicIpAddressPoolId() {
        return this.publicIpAddressPoolId;
    }

    private PublicIpAddressPoolCidrBlockArgs() {}

    private PublicIpAddressPoolCidrBlockArgs(PublicIpAddressPoolCidrBlockArgs $) {
        this.cidrBlock = $.cidrBlock;
        this.publicIpAddressPoolId = $.publicIpAddressPoolId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PublicIpAddressPoolCidrBlockArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PublicIpAddressPoolCidrBlockArgs $;

        public Builder() {
            $ = new PublicIpAddressPoolCidrBlockArgs();
        }

        public Builder(PublicIpAddressPoolCidrBlockArgs defaults) {
            $ = new PublicIpAddressPoolCidrBlockArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cidrBlock The CIDR block.
         * 
         * @return builder
         * 
         */
        public Builder cidrBlock(Output<String> cidrBlock) {
            $.cidrBlock = cidrBlock;
            return this;
        }

        /**
         * @param cidrBlock The CIDR block.
         * 
         * @return builder
         * 
         */
        public Builder cidrBlock(String cidrBlock) {
            return cidrBlock(Output.of(cidrBlock));
        }

        /**
         * @param publicIpAddressPoolId The ID of the VPC Public IP address pool.
         * 
         * @return builder
         * 
         */
        public Builder publicIpAddressPoolId(Output<String> publicIpAddressPoolId) {
            $.publicIpAddressPoolId = publicIpAddressPoolId;
            return this;
        }

        /**
         * @param publicIpAddressPoolId The ID of the VPC Public IP address pool.
         * 
         * @return builder
         * 
         */
        public Builder publicIpAddressPoolId(String publicIpAddressPoolId) {
            return publicIpAddressPoolId(Output.of(publicIpAddressPoolId));
        }

        public PublicIpAddressPoolCidrBlockArgs build() {
            $.cidrBlock = Objects.requireNonNull($.cidrBlock, "expected parameter 'cidrBlock' to be non-null");
            $.publicIpAddressPoolId = Objects.requireNonNull($.publicIpAddressPoolId, "expected parameter 'publicIpAddressPoolId' to be non-null");
            return $;
        }
    }

}
