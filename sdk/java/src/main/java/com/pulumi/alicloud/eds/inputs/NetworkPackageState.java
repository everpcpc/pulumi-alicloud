// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eds.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NetworkPackageState extends com.pulumi.resources.ResourceArgs {

    public static final NetworkPackageState Empty = new NetworkPackageState();

    /**
     * The bandwidth of package public network bandwidth peak. Valid values: 1~200. Unit:Mbps.
     * 
     */
    @Import(name="bandwidth")
    private @Nullable Output<Integer> bandwidth;

    /**
     * @return The bandwidth of package public network bandwidth peak. Valid values: 1~200. Unit:Mbps.
     * 
     */
    public Optional<Output<Integer>> bandwidth() {
        return Optional.ofNullable(this.bandwidth);
    }

    /**
     * The internet charge type  of  package.
     * 
     */
    @Import(name="internetChargeType")
    private @Nullable Output<String> internetChargeType;

    /**
     * @return The internet charge type  of  package.
     * 
     */
    public Optional<Output<String>> internetChargeType() {
        return Optional.ofNullable(this.internetChargeType);
    }

    /**
     * The ID of office site.
     * 
     */
    @Import(name="officeSiteId")
    private @Nullable Output<String> officeSiteId;

    /**
     * @return The ID of office site.
     * 
     */
    public Optional<Output<String>> officeSiteId() {
        return Optional.ofNullable(this.officeSiteId);
    }

    /**
     * The status of network package. Valid values: `Creating`, `InUse`, `Releasing`,`Released`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of network package. Valid values: `Creating`, `InUse`, `Releasing`,`Released`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private NetworkPackageState() {}

    private NetworkPackageState(NetworkPackageState $) {
        this.bandwidth = $.bandwidth;
        this.internetChargeType = $.internetChargeType;
        this.officeSiteId = $.officeSiteId;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NetworkPackageState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NetworkPackageState $;

        public Builder() {
            $ = new NetworkPackageState();
        }

        public Builder(NetworkPackageState defaults) {
            $ = new NetworkPackageState(Objects.requireNonNull(defaults));
        }

        /**
         * @param bandwidth The bandwidth of package public network bandwidth peak. Valid values: 1~200. Unit:Mbps.
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(@Nullable Output<Integer> bandwidth) {
            $.bandwidth = bandwidth;
            return this;
        }

        /**
         * @param bandwidth The bandwidth of package public network bandwidth peak. Valid values: 1~200. Unit:Mbps.
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(Integer bandwidth) {
            return bandwidth(Output.of(bandwidth));
        }

        /**
         * @param internetChargeType The internet charge type  of  package.
         * 
         * @return builder
         * 
         */
        public Builder internetChargeType(@Nullable Output<String> internetChargeType) {
            $.internetChargeType = internetChargeType;
            return this;
        }

        /**
         * @param internetChargeType The internet charge type  of  package.
         * 
         * @return builder
         * 
         */
        public Builder internetChargeType(String internetChargeType) {
            return internetChargeType(Output.of(internetChargeType));
        }

        /**
         * @param officeSiteId The ID of office site.
         * 
         * @return builder
         * 
         */
        public Builder officeSiteId(@Nullable Output<String> officeSiteId) {
            $.officeSiteId = officeSiteId;
            return this;
        }

        /**
         * @param officeSiteId The ID of office site.
         * 
         * @return builder
         * 
         */
        public Builder officeSiteId(String officeSiteId) {
            return officeSiteId(Output.of(officeSiteId));
        }

        /**
         * @param status The status of network package. Valid values: `Creating`, `InUse`, `Releasing`,`Released`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of network package. Valid values: `Creating`, `InUse`, `Releasing`,`Released`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public NetworkPackageState build() {
            return $;
        }
    }

}