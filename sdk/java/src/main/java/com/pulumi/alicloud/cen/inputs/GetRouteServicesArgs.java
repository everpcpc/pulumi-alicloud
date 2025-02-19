// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRouteServicesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRouteServicesArgs Empty = new GetRouteServicesArgs();

    /**
     * The region of the network instances that access the cloud services.
     * 
     */
    @Import(name="accessRegionId")
    private @Nullable Output<String> accessRegionId;

    /**
     * @return The region of the network instances that access the cloud services.
     * 
     */
    public Optional<Output<String>> accessRegionId() {
        return Optional.ofNullable(this.accessRegionId);
    }

    /**
     * The ID of the CEN instance.
     * 
     */
    @Import(name="cenId", required=true)
    private Output<String> cenId;

    /**
     * @return The ID of the CEN instance.
     * 
     */
    public Output<String> cenId() {
        return this.cenId;
    }

    /**
     * The domain name or IP address of the cloud service.
     * 
     */
    @Import(name="host")
    private @Nullable Output<String> host;

    /**
     * @return The domain name or IP address of the cloud service.
     * 
     */
    public Optional<Output<String>> host() {
        return Optional.ofNullable(this.host);
    }

    /**
     * The region of the cloud service.
     * 
     */
    @Import(name="hostRegionId")
    private @Nullable Output<String> hostRegionId;

    /**
     * @return The region of the cloud service.
     * 
     */
    public Optional<Output<String>> hostRegionId() {
        return Optional.ofNullable(this.hostRegionId);
    }

    /**
     * The VPC associated with the cloud service.
     * 
     */
    @Import(name="hostVpcId")
    private @Nullable Output<String> hostVpcId;

    /**
     * @return The VPC associated with the cloud service.
     * 
     */
    public Optional<Output<String>> hostVpcId() {
        return Optional.ofNullable(this.hostVpcId);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The status of the cloud service. Valid values: `Active`, `Creating` and `Deleting`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the cloud service. Valid values: `Active`, `Creating` and `Deleting`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private GetRouteServicesArgs() {}

    private GetRouteServicesArgs(GetRouteServicesArgs $) {
        this.accessRegionId = $.accessRegionId;
        this.cenId = $.cenId;
        this.host = $.host;
        this.hostRegionId = $.hostRegionId;
        this.hostVpcId = $.hostVpcId;
        this.outputFile = $.outputFile;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRouteServicesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRouteServicesArgs $;

        public Builder() {
            $ = new GetRouteServicesArgs();
        }

        public Builder(GetRouteServicesArgs defaults) {
            $ = new GetRouteServicesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessRegionId The region of the network instances that access the cloud services.
         * 
         * @return builder
         * 
         */
        public Builder accessRegionId(@Nullable Output<String> accessRegionId) {
            $.accessRegionId = accessRegionId;
            return this;
        }

        /**
         * @param accessRegionId The region of the network instances that access the cloud services.
         * 
         * @return builder
         * 
         */
        public Builder accessRegionId(String accessRegionId) {
            return accessRegionId(Output.of(accessRegionId));
        }

        /**
         * @param cenId The ID of the CEN instance.
         * 
         * @return builder
         * 
         */
        public Builder cenId(Output<String> cenId) {
            $.cenId = cenId;
            return this;
        }

        /**
         * @param cenId The ID of the CEN instance.
         * 
         * @return builder
         * 
         */
        public Builder cenId(String cenId) {
            return cenId(Output.of(cenId));
        }

        /**
         * @param host The domain name or IP address of the cloud service.
         * 
         * @return builder
         * 
         */
        public Builder host(@Nullable Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host The domain name or IP address of the cloud service.
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param hostRegionId The region of the cloud service.
         * 
         * @return builder
         * 
         */
        public Builder hostRegionId(@Nullable Output<String> hostRegionId) {
            $.hostRegionId = hostRegionId;
            return this;
        }

        /**
         * @param hostRegionId The region of the cloud service.
         * 
         * @return builder
         * 
         */
        public Builder hostRegionId(String hostRegionId) {
            return hostRegionId(Output.of(hostRegionId));
        }

        /**
         * @param hostVpcId The VPC associated with the cloud service.
         * 
         * @return builder
         * 
         */
        public Builder hostVpcId(@Nullable Output<String> hostVpcId) {
            $.hostVpcId = hostVpcId;
            return this;
        }

        /**
         * @param hostVpcId The VPC associated with the cloud service.
         * 
         * @return builder
         * 
         */
        public Builder hostVpcId(String hostVpcId) {
            return hostVpcId(Output.of(hostVpcId));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        /**
         * @param status The status of the cloud service. Valid values: `Active`, `Creating` and `Deleting`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the cloud service. Valid values: `Active`, `Creating` and `Deleting`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public GetRouteServicesArgs build() {
            $.cenId = Objects.requireNonNull($.cenId, "expected parameter 'cenId' to be non-null");
            return $;
        }
    }

}
