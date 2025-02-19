// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudfirewall.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAddressBooksArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAddressBooksArgs Empty = new GetAddressBooksArgs();

    /**
     * The type of the Address Book.
     * 
     */
    @Import(name="groupType")
    private @Nullable Output<String> groupType;

    /**
     * @return The type of the Address Book.
     * 
     */
    public Optional<Output<String>> groupType() {
        return Optional.ofNullable(this.groupType);
    }

    /**
     * A list of Address Book IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Address Book IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results Address Book name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results Address Book name.
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

    private GetAddressBooksArgs() {}

    private GetAddressBooksArgs(GetAddressBooksArgs $) {
        this.groupType = $.groupType;
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAddressBooksArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAddressBooksArgs $;

        public Builder() {
            $ = new GetAddressBooksArgs();
        }

        public Builder(GetAddressBooksArgs defaults) {
            $ = new GetAddressBooksArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param groupType The type of the Address Book.
         * 
         * @return builder
         * 
         */
        public Builder groupType(@Nullable Output<String> groupType) {
            $.groupType = groupType;
            return this;
        }

        /**
         * @param groupType The type of the Address Book.
         * 
         * @return builder
         * 
         */
        public Builder groupType(String groupType) {
            return groupType(Output.of(groupType));
        }

        /**
         * @param ids A list of Address Book IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Address Book IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Address Book IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results Address Book name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results Address Book name.
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

        public GetAddressBooksArgs build() {
            return $;
        }
    }

}
