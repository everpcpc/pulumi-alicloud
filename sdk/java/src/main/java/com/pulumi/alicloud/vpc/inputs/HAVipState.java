// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HAVipState extends com.pulumi.resources.ResourceArgs {

    public static final HAVipState Empty = new HAVipState();

    /**
     * The description of the HaVip instance.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the HaVip instance.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name of the HaVip instance.
     * 
     */
    @Import(name="havipName")
    private @Nullable Output<String> havipName;

    /**
     * @return The name of the HaVip instance.
     * 
     */
    public Optional<Output<String>> havipName() {
        return Optional.ofNullable(this.havipName);
    }

    /**
     * The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
     * 
     */
    @Import(name="ipAddress")
    private @Nullable Output<String> ipAddress;

    /**
     * @return The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
     * 
     */
    public Optional<Output<String>> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }

    /**
     * (Available in v1.120.0+) The status of the HaVip instance.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return (Available in v1.120.0+) The status of the HaVip instance.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The vswitch_id of the HaVip, the field can&#39;t be changed.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable Output<String> vswitchId;

    /**
     * @return The vswitch_id of the HaVip, the field can&#39;t be changed.
     * 
     */
    public Optional<Output<String>> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    private HAVipState() {}

    private HAVipState(HAVipState $) {
        this.description = $.description;
        this.havipName = $.havipName;
        this.ipAddress = $.ipAddress;
        this.status = $.status;
        this.vswitchId = $.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HAVipState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HAVipState $;

        public Builder() {
            $ = new HAVipState();
        }

        public Builder(HAVipState defaults) {
            $ = new HAVipState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the HaVip instance.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the HaVip instance.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param havipName The name of the HaVip instance.
         * 
         * @return builder
         * 
         */
        public Builder havipName(@Nullable Output<String> havipName) {
            $.havipName = havipName;
            return this;
        }

        /**
         * @param havipName The name of the HaVip instance.
         * 
         * @return builder
         * 
         */
        public Builder havipName(String havipName) {
            return havipName(Output.of(havipName));
        }

        /**
         * @param ipAddress The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(@Nullable Output<String> ipAddress) {
            $.ipAddress = ipAddress;
            return this;
        }

        /**
         * @param ipAddress The ip address of the HaVip. If not filled, the default will be assigned one from the vswitch.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(String ipAddress) {
            return ipAddress(Output.of(ipAddress));
        }

        /**
         * @param status (Available in v1.120.0+) The status of the HaVip instance.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status (Available in v1.120.0+) The status of the HaVip instance.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param vswitchId The vswitch_id of the HaVip, the field can&#39;t be changed.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId The vswitch_id of the HaVip, the field can&#39;t be changed.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        public HAVipState build() {
            return $;
        }
    }

}
