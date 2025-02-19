// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eflo.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VpdState extends com.pulumi.resources.ResourceArgs {

    public static final VpdState Empty = new VpdState();

    /**
     * CIDR network segment
     * 
     */
    @Import(name="cidr")
    private @Nullable Output<String> cidr;

    /**
     * @return CIDR network segment
     * 
     */
    public Optional<Output<String>> cidr() {
        return Optional.ofNullable(this.cidr);
    }

    /**
     * The creation time of the resource
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The creation time of the resource
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * Modification time
     * 
     */
    @Import(name="gmtModified")
    private @Nullable Output<String> gmtModified;

    /**
     * @return Modification time
     * 
     */
    public Optional<Output<String>> gmtModified() {
        return Optional.ofNullable(this.gmtModified);
    }

    /**
     * The Resource group id
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The Resource group id
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The Vpd status.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The Vpd status.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The Name of the VPD.
     * 
     */
    @Import(name="vpdName")
    private @Nullable Output<String> vpdName;

    /**
     * @return The Name of the VPD.
     * 
     */
    public Optional<Output<String>> vpdName() {
        return Optional.ofNullable(this.vpdName);
    }

    private VpdState() {}

    private VpdState(VpdState $) {
        this.cidr = $.cidr;
        this.createTime = $.createTime;
        this.gmtModified = $.gmtModified;
        this.resourceGroupId = $.resourceGroupId;
        this.status = $.status;
        this.vpdName = $.vpdName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpdState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpdState $;

        public Builder() {
            $ = new VpdState();
        }

        public Builder(VpdState defaults) {
            $ = new VpdState(Objects.requireNonNull(defaults));
        }

        /**
         * @param cidr CIDR network segment
         * 
         * @return builder
         * 
         */
        public Builder cidr(@Nullable Output<String> cidr) {
            $.cidr = cidr;
            return this;
        }

        /**
         * @param cidr CIDR network segment
         * 
         * @return builder
         * 
         */
        public Builder cidr(String cidr) {
            return cidr(Output.of(cidr));
        }

        /**
         * @param createTime The creation time of the resource
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The creation time of the resource
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param gmtModified Modification time
         * 
         * @return builder
         * 
         */
        public Builder gmtModified(@Nullable Output<String> gmtModified) {
            $.gmtModified = gmtModified;
            return this;
        }

        /**
         * @param gmtModified Modification time
         * 
         * @return builder
         * 
         */
        public Builder gmtModified(String gmtModified) {
            return gmtModified(Output.of(gmtModified));
        }

        /**
         * @param resourceGroupId The Resource group id
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The Resource group id
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param status The Vpd status.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The Vpd status.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param vpdName The Name of the VPD.
         * 
         * @return builder
         * 
         */
        public Builder vpdName(@Nullable Output<String> vpdName) {
            $.vpdName = vpdName;
            return this;
        }

        /**
         * @param vpdName The Name of the VPD.
         * 
         * @return builder
         * 
         */
        public Builder vpdName(String vpdName) {
            return vpdName(Output.of(vpdName));
        }

        public VpdState build() {
            return $;
        }
    }

}
