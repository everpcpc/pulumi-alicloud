// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVaultsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVaultsArgs Empty = new GetVaultsArgs();

    /**
     * A list of Vault IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Vault IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Vault name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Vault name.
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
     * The status of Vault. Valid values: `CREATED`, `ERROR`, `UNKNOWN`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of Vault. Valid values: `CREATED`, `ERROR`, `UNKNOWN`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The type of Vault. Valid values: `STANDARD`,`OTS_BACKUP`.
     * 
     */
    @Import(name="vaultType")
    private @Nullable Output<String> vaultType;

    /**
     * @return The type of Vault. Valid values: `STANDARD`,`OTS_BACKUP`.
     * 
     */
    public Optional<Output<String>> vaultType() {
        return Optional.ofNullable(this.vaultType);
    }

    private GetVaultsArgs() {}

    private GetVaultsArgs(GetVaultsArgs $) {
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.status = $.status;
        this.vaultType = $.vaultType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVaultsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVaultsArgs $;

        public Builder() {
            $ = new GetVaultsArgs();
        }

        public Builder(GetVaultsArgs defaults) {
            $ = new GetVaultsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Vault IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Vault IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Vault IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Vault name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Vault name.
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
         * @param status The status of Vault. Valid values: `CREATED`, `ERROR`, `UNKNOWN`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of Vault. Valid values: `CREATED`, `ERROR`, `UNKNOWN`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param vaultType The type of Vault. Valid values: `STANDARD`,`OTS_BACKUP`.
         * 
         * @return builder
         * 
         */
        public Builder vaultType(@Nullable Output<String> vaultType) {
            $.vaultType = vaultType;
            return this;
        }

        /**
         * @param vaultType The type of Vault. Valid values: `STANDARD`,`OTS_BACKUP`.
         * 
         * @return builder
         * 
         */
        public Builder vaultType(String vaultType) {
            return vaultType(Output.of(vaultType));
        }

        public GetVaultsArgs build() {
            return $;
        }
    }

}
