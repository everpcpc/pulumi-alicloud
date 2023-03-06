// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PublicIpAddressPoolState extends com.pulumi.resources.ResourceArgs {

    public static final PublicIpAddressPoolState Empty = new PublicIpAddressPoolState();

    /**
     * The description of the VPC Public IP address pool.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the VPC Public IP address pool.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The Internet service provider. Valid values: `BGP`, `BGP_PRO`, `ChinaTelecom`, `ChinaUnicom`, `ChinaMobile`, `ChinaTelecom_L2`, `ChinaUnicom_L2`, `ChinaMobile_L2`, `BGP_FinanceCloud`. Default Value: `BGP`.
     * 
     */
    @Import(name="isp")
    private @Nullable Output<String> isp;

    /**
     * @return The Internet service provider. Valid values: `BGP`, `BGP_PRO`, `ChinaTelecom`, `ChinaUnicom`, `ChinaMobile`, `ChinaTelecom_L2`, `ChinaUnicom_L2`, `ChinaMobile_L2`, `BGP_FinanceCloud`. Default Value: `BGP`.
     * 
     */
    public Optional<Output<String>> isp() {
        return Optional.ofNullable(this.isp);
    }

    /**
     * The name of the VPC Public IP address pool.
     * 
     */
    @Import(name="publicIpAddressPoolName")
    private @Nullable Output<String> publicIpAddressPoolName;

    /**
     * @return The name of the VPC Public IP address pool.
     * 
     */
    public Optional<Output<String>> publicIpAddressPoolName() {
        return Optional.ofNullable(this.publicIpAddressPoolName);
    }

    /**
     * The status of the VPC Public IP address pool.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the VPC Public IP address pool.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private PublicIpAddressPoolState() {}

    private PublicIpAddressPoolState(PublicIpAddressPoolState $) {
        this.description = $.description;
        this.isp = $.isp;
        this.publicIpAddressPoolName = $.publicIpAddressPoolName;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PublicIpAddressPoolState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PublicIpAddressPoolState $;

        public Builder() {
            $ = new PublicIpAddressPoolState();
        }

        public Builder(PublicIpAddressPoolState defaults) {
            $ = new PublicIpAddressPoolState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the VPC Public IP address pool.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the VPC Public IP address pool.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param isp The Internet service provider. Valid values: `BGP`, `BGP_PRO`, `ChinaTelecom`, `ChinaUnicom`, `ChinaMobile`, `ChinaTelecom_L2`, `ChinaUnicom_L2`, `ChinaMobile_L2`, `BGP_FinanceCloud`. Default Value: `BGP`.
         * 
         * @return builder
         * 
         */
        public Builder isp(@Nullable Output<String> isp) {
            $.isp = isp;
            return this;
        }

        /**
         * @param isp The Internet service provider. Valid values: `BGP`, `BGP_PRO`, `ChinaTelecom`, `ChinaUnicom`, `ChinaMobile`, `ChinaTelecom_L2`, `ChinaUnicom_L2`, `ChinaMobile_L2`, `BGP_FinanceCloud`. Default Value: `BGP`.
         * 
         * @return builder
         * 
         */
        public Builder isp(String isp) {
            return isp(Output.of(isp));
        }

        /**
         * @param publicIpAddressPoolName The name of the VPC Public IP address pool.
         * 
         * @return builder
         * 
         */
        public Builder publicIpAddressPoolName(@Nullable Output<String> publicIpAddressPoolName) {
            $.publicIpAddressPoolName = publicIpAddressPoolName;
            return this;
        }

        /**
         * @param publicIpAddressPoolName The name of the VPC Public IP address pool.
         * 
         * @return builder
         * 
         */
        public Builder publicIpAddressPoolName(String publicIpAddressPoolName) {
            return publicIpAddressPoolName(Output.of(publicIpAddressPoolName));
        }

        /**
         * @param status The status of the VPC Public IP address pool.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the VPC Public IP address pool.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public PublicIpAddressPoolState build() {
            return $;
        }
    }

}