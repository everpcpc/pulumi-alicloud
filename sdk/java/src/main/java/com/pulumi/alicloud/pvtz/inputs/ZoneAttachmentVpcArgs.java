// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.pvtz.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ZoneAttachmentVpcArgs extends com.pulumi.resources.ResourceArgs {

    public static final ZoneAttachmentVpcArgs Empty = new ZoneAttachmentVpcArgs();

    /**
     * The region of the vpc. If not set, the current region will instead of.
     * 
     * Recommend to use `vpcs`.
     * 
     */
    @Import(name="regionId")
    private @Nullable Output<String> regionId;

    /**
     * @return The region of the vpc. If not set, the current region will instead of.
     * 
     * Recommend to use `vpcs`.
     * 
     */
    public Optional<Output<String>> regionId() {
        return Optional.ofNullable(this.regionId);
    }

    /**
     * The Id of the vpc.
     * 
     */
    @Import(name="vpcId", required=true)
    private Output<String> vpcId;

    /**
     * @return The Id of the vpc.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    private ZoneAttachmentVpcArgs() {}

    private ZoneAttachmentVpcArgs(ZoneAttachmentVpcArgs $) {
        this.regionId = $.regionId;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ZoneAttachmentVpcArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ZoneAttachmentVpcArgs $;

        public Builder() {
            $ = new ZoneAttachmentVpcArgs();
        }

        public Builder(ZoneAttachmentVpcArgs defaults) {
            $ = new ZoneAttachmentVpcArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param regionId The region of the vpc. If not set, the current region will instead of.
         * 
         * Recommend to use `vpcs`.
         * 
         * @return builder
         * 
         */
        public Builder regionId(@Nullable Output<String> regionId) {
            $.regionId = regionId;
            return this;
        }

        /**
         * @param regionId The region of the vpc. If not set, the current region will instead of.
         * 
         * Recommend to use `vpcs`.
         * 
         * @return builder
         * 
         */
        public Builder regionId(String regionId) {
            return regionId(Output.of(regionId));
        }

        /**
         * @param vpcId The Id of the vpc.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The Id of the vpc.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        public ZoneAttachmentVpcArgs build() {
            $.vpcId = Objects.requireNonNull($.vpcId, "expected parameter 'vpcId' to be non-null");
            return $;
        }
    }

}
