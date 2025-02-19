// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTransitRouterRouteEntriesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTransitRouterRouteEntriesPlainArgs Empty = new GetTransitRouterRouteEntriesPlainArgs();

    /**
     * A list of CEN Transit Router Route Entry IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of CEN Transit Router Route Entry IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
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

    @Import(name="status")
    private @Nullable String status;

    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * A list of ID of the cen transit router route entry.
     * 
     */
    @Import(name="transitRouterRouteEntryIds")
    private @Nullable List<String> transitRouterRouteEntryIds;

    /**
     * @return A list of ID of the cen transit router route entry.
     * 
     */
    public Optional<List<String>> transitRouterRouteEntryIds() {
        return Optional.ofNullable(this.transitRouterRouteEntryIds);
    }

    /**
     * A list of name of the cen transit router route entry.
     * 
     */
    @Import(name="transitRouterRouteEntryNames")
    private @Nullable List<String> transitRouterRouteEntryNames;

    /**
     * @return A list of name of the cen transit router route entry.
     * 
     */
    public Optional<List<String>> transitRouterRouteEntryNames() {
        return Optional.ofNullable(this.transitRouterRouteEntryNames);
    }

    /**
     * The status of the resource.Valid values `Creating`, `Active` and `Deleting`.
     * 
     */
    @Import(name="transitRouterRouteEntryStatus")
    private @Nullable String transitRouterRouteEntryStatus;

    /**
     * @return The status of the resource.Valid values `Creating`, `Active` and `Deleting`.
     * 
     */
    public Optional<String> transitRouterRouteEntryStatus() {
        return Optional.ofNullable(this.transitRouterRouteEntryStatus);
    }

    /**
     * ID of the CEN Transit Router Route Table.
     * 
     */
    @Import(name="transitRouterRouteTableId", required=true)
    private String transitRouterRouteTableId;

    /**
     * @return ID of the CEN Transit Router Route Table.
     * 
     */
    public String transitRouterRouteTableId() {
        return this.transitRouterRouteTableId;
    }

    private GetTransitRouterRouteEntriesPlainArgs() {}

    private GetTransitRouterRouteEntriesPlainArgs(GetTransitRouterRouteEntriesPlainArgs $) {
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.status = $.status;
        this.transitRouterRouteEntryIds = $.transitRouterRouteEntryIds;
        this.transitRouterRouteEntryNames = $.transitRouterRouteEntryNames;
        this.transitRouterRouteEntryStatus = $.transitRouterRouteEntryStatus;
        this.transitRouterRouteTableId = $.transitRouterRouteTableId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTransitRouterRouteEntriesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTransitRouterRouteEntriesPlainArgs $;

        public Builder() {
            $ = new GetTransitRouterRouteEntriesPlainArgs();
        }

        public Builder(GetTransitRouterRouteEntriesPlainArgs defaults) {
            $ = new GetTransitRouterRouteEntriesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of CEN Transit Router Route Entry IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of CEN Transit Router Route Entry IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
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

        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param transitRouterRouteEntryIds A list of ID of the cen transit router route entry.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryIds(@Nullable List<String> transitRouterRouteEntryIds) {
            $.transitRouterRouteEntryIds = transitRouterRouteEntryIds;
            return this;
        }

        /**
         * @param transitRouterRouteEntryIds A list of ID of the cen transit router route entry.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryIds(String... transitRouterRouteEntryIds) {
            return transitRouterRouteEntryIds(List.of(transitRouterRouteEntryIds));
        }

        /**
         * @param transitRouterRouteEntryNames A list of name of the cen transit router route entry.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryNames(@Nullable List<String> transitRouterRouteEntryNames) {
            $.transitRouterRouteEntryNames = transitRouterRouteEntryNames;
            return this;
        }

        /**
         * @param transitRouterRouteEntryNames A list of name of the cen transit router route entry.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryNames(String... transitRouterRouteEntryNames) {
            return transitRouterRouteEntryNames(List.of(transitRouterRouteEntryNames));
        }

        /**
         * @param transitRouterRouteEntryStatus The status of the resource.Valid values `Creating`, `Active` and `Deleting`.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteEntryStatus(@Nullable String transitRouterRouteEntryStatus) {
            $.transitRouterRouteEntryStatus = transitRouterRouteEntryStatus;
            return this;
        }

        /**
         * @param transitRouterRouteTableId ID of the CEN Transit Router Route Table.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterRouteTableId(String transitRouterRouteTableId) {
            $.transitRouterRouteTableId = transitRouterRouteTableId;
            return this;
        }

        public GetTransitRouterRouteEntriesPlainArgs build() {
            $.transitRouterRouteTableId = Objects.requireNonNull($.transitRouterRouteTableId, "expected parameter 'transitRouterRouteTableId' to be non-null");
            return $;
        }
    }

}
