// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eds.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBundlesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBundlesArgs Empty = new GetBundlesArgs();

    /**
     * The bundle id of the bundle.
     * 
     */
    @Import(name="bundleIds")
    private @Nullable Output<List<String>> bundleIds;

    /**
     * @return The bundle id of the bundle.
     * 
     */
    public Optional<Output<List<String>>> bundleIds() {
        return Optional.ofNullable(this.bundleIds);
    }

    /**
     * The bundle type of  the bundle. Valid values: `SYSTEM`,`CUSTOM`.
     * 
     */
    @Import(name="bundleType")
    private @Nullable Output<String> bundleType;

    /**
     * @return The bundle type of  the bundle. Valid values: `SYSTEM`,`CUSTOM`.
     * 
     */
    public Optional<Output<String>> bundleType() {
        return Optional.ofNullable(this.bundleType);
    }

    /**
     * A list of Bundle IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Bundle IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Bundle name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Bundle name.
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

    private GetBundlesArgs() {}

    private GetBundlesArgs(GetBundlesArgs $) {
        this.bundleIds = $.bundleIds;
        this.bundleType = $.bundleType;
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBundlesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBundlesArgs $;

        public Builder() {
            $ = new GetBundlesArgs();
        }

        public Builder(GetBundlesArgs defaults) {
            $ = new GetBundlesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bundleIds The bundle id of the bundle.
         * 
         * @return builder
         * 
         */
        public Builder bundleIds(@Nullable Output<List<String>> bundleIds) {
            $.bundleIds = bundleIds;
            return this;
        }

        /**
         * @param bundleIds The bundle id of the bundle.
         * 
         * @return builder
         * 
         */
        public Builder bundleIds(List<String> bundleIds) {
            return bundleIds(Output.of(bundleIds));
        }

        /**
         * @param bundleIds The bundle id of the bundle.
         * 
         * @return builder
         * 
         */
        public Builder bundleIds(String... bundleIds) {
            return bundleIds(List.of(bundleIds));
        }

        /**
         * @param bundleType The bundle type of  the bundle. Valid values: `SYSTEM`,`CUSTOM`.
         * 
         * @return builder
         * 
         */
        public Builder bundleType(@Nullable Output<String> bundleType) {
            $.bundleType = bundleType;
            return this;
        }

        /**
         * @param bundleType The bundle type of  the bundle. Valid values: `SYSTEM`,`CUSTOM`.
         * 
         * @return builder
         * 
         */
        public Builder bundleType(String bundleType) {
            return bundleType(Output.of(bundleType));
        }

        /**
         * @param ids A list of Bundle IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Bundle IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Bundle IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Bundle name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Bundle name.
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

        public GetBundlesArgs build() {
            return $;
        }
    }

}
