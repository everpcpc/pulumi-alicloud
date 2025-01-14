// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEcsPrefixListsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEcsPrefixListsArgs Empty = new GetEcsPrefixListsArgs();

    /**
     * The address family of the prefix list. Valid values:`IPv4`,`IPv6`.
     * 
     */
    @Import(name="addressFamily")
    private @Nullable Output<String> addressFamily;

    /**
     * @return The address family of the prefix list. Valid values:`IPv4`,`IPv6`.
     * 
     */
    public Optional<Output<String>> addressFamily() {
        return Optional.ofNullable(this.addressFamily);
    }

    @Import(name="enableDetails")
    private @Nullable Output<Boolean> enableDetails;

    public Optional<Output<Boolean>> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * A list of Prefix List IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Prefix List IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by `prefix_list_name`.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by `prefix_list_name`.
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

    private GetEcsPrefixListsArgs() {}

    private GetEcsPrefixListsArgs(GetEcsPrefixListsArgs $) {
        this.addressFamily = $.addressFamily;
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEcsPrefixListsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEcsPrefixListsArgs $;

        public Builder() {
            $ = new GetEcsPrefixListsArgs();
        }

        public Builder(GetEcsPrefixListsArgs defaults) {
            $ = new GetEcsPrefixListsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressFamily The address family of the prefix list. Valid values:`IPv4`,`IPv6`.
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(@Nullable Output<String> addressFamily) {
            $.addressFamily = addressFamily;
            return this;
        }

        /**
         * @param addressFamily The address family of the prefix list. Valid values:`IPv4`,`IPv6`.
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(String addressFamily) {
            return addressFamily(Output.of(addressFamily));
        }

        public Builder enableDetails(@Nullable Output<Boolean> enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        public Builder enableDetails(Boolean enableDetails) {
            return enableDetails(Output.of(enableDetails));
        }

        /**
         * @param ids A list of Prefix List IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Prefix List IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Prefix List IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by `prefix_list_name`.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by `prefix_list_name`.
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

        public GetEcsPrefixListsArgs build() {
            return $;
        }
    }

}
