// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mse.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetGatewaysArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGatewaysArgs Empty = new GetGatewaysArgs();

    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     * 
     */
    @Import(name="enableDetails")
    private @Nullable Output<Boolean> enableDetails;

    /**
     * @return Default to `false`. Set it to `true` can output more details about resource attributes.
     * 
     */
    public Optional<Output<Boolean>> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * The name of the Gateway.
     * 
     */
    @Import(name="gatewayName")
    private @Nullable Output<String> gatewayName;

    /**
     * @return The name of the Gateway.
     * 
     */
    public Optional<Output<String>> gatewayName() {
        return Optional.ofNullable(this.gatewayName);
    }

    /**
     * A list of Gateway IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Gateway IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Gateway name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Gateway name.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
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
     * The status of the gateway.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the gateway.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The ID of the vpc.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The ID of the vpc.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    private GetGatewaysArgs() {}

    private GetGatewaysArgs(GetGatewaysArgs $) {
        this.enableDetails = $.enableDetails;
        this.gatewayName = $.gatewayName;
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.status = $.status;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGatewaysArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGatewaysArgs $;

        public Builder() {
            $ = new GetGatewaysArgs();
        }

        public Builder(GetGatewaysArgs defaults) {
            $ = new GetGatewaysArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enableDetails Default to `false`. Set it to `true` can output more details about resource attributes.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(@Nullable Output<Boolean> enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param enableDetails Default to `false`. Set it to `true` can output more details about resource attributes.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(Boolean enableDetails) {
            return enableDetails(Output.of(enableDetails));
        }

        /**
         * @param gatewayName The name of the Gateway.
         * 
         * @return builder
         * 
         */
        public Builder gatewayName(@Nullable Output<String> gatewayName) {
            $.gatewayName = gatewayName;
            return this;
        }

        /**
         * @param gatewayName The name of the Gateway.
         * 
         * @return builder
         * 
         */
        public Builder gatewayName(String gatewayName) {
            return gatewayName(Output.of(gatewayName));
        }

        /**
         * @param ids A list of Gateway IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Gateway IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Gateway IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Gateway name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Gateway name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
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
         * @param status The status of the gateway.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the gateway.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param vpcId The ID of the vpc.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The ID of the vpc.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        public GetGatewaysArgs build() {
            return $;
        }
    }

}
