// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.adb.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDBClustersPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDBClustersPlainArgs Empty = new GetDBClustersPlainArgs();

    /**
     * The description of DBCluster.
     * 
     */
    @Import(name="description")
    private @Nullable String description;

    /**
     * @return The description of DBCluster.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * A regex string to filter results by DBCluster description.
     * 
     */
    @Import(name="descriptionRegex")
    private @Nullable String descriptionRegex;

    /**
     * @return A regex string to filter results by DBCluster description.
     * 
     */
    public Optional<String> descriptionRegex() {
        return Optional.ofNullable(this.descriptionRegex);
    }

    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     * 
     */
    @Import(name="enableDetails")
    private @Nullable Boolean enableDetails;

    /**
     * @return Default to `false`. Set it to `true` can output more details about resource attributes.
     * 
     */
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * A list of DBCluster IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of DBCluster IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    @Import(name="pageNumber")
    private @Nullable Integer pageNumber;

    public Optional<Integer> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }

    @Import(name="pageSize")
    private @Nullable Integer pageSize;

    public Optional<Integer> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }

    /**
     * The ID of the resource group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable String resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * The status of the resource.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the resource.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The tag of the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Map<String,Object> tags;

    /**
     * @return The tag of the resource.
     * 
     */
    public Optional<Map<String,Object>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private GetDBClustersPlainArgs() {}

    private GetDBClustersPlainArgs(GetDBClustersPlainArgs $) {
        this.description = $.description;
        this.descriptionRegex = $.descriptionRegex;
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
        this.resourceGroupId = $.resourceGroupId;
        this.status = $.status;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDBClustersPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDBClustersPlainArgs $;

        public Builder() {
            $ = new GetDBClustersPlainArgs();
        }

        public Builder(GetDBClustersPlainArgs defaults) {
            $ = new GetDBClustersPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of DBCluster.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable String description) {
            $.description = description;
            return this;
        }

        /**
         * @param descriptionRegex A regex string to filter results by DBCluster description.
         * 
         * @return builder
         * 
         */
        public Builder descriptionRegex(@Nullable String descriptionRegex) {
            $.descriptionRegex = descriptionRegex;
            return this;
        }

        /**
         * @param enableDetails Default to `false`. Set it to `true` can output more details about resource attributes.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param ids A list of DBCluster IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of DBCluster IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder pageNumber(@Nullable Integer pageNumber) {
            $.pageNumber = pageNumber;
            return this;
        }

        public Builder pageSize(@Nullable Integer pageSize) {
            $.pageSize = pageSize;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param status The status of the resource.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param tags The tag of the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Map<String,Object> tags) {
            $.tags = tags;
            return this;
        }

        public GetDBClustersPlainArgs build() {
            return $;
        }
    }

}
