// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetChildInstanceRouteEntryToAttachmentsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetChildInstanceRouteEntryToAttachmentsArgs Empty = new GetChildInstanceRouteEntryToAttachmentsArgs();

    /**
     * The ID of the CEN instance.
     * 
     */
    @Import(name="cenId")
    private @Nullable Output<String> cenId;

    /**
     * @return The ID of the CEN instance.
     * 
     */
    public Optional<Output<String>> cenId() {
        return Optional.ofNullable(this.cenId);
    }

    /**
     * The first ID of the resource
     * 
     */
    @Import(name="childInstanceRouteTableId", required=true)
    private Output<String> childInstanceRouteTableId;

    /**
     * @return The first ID of the resource
     * 
     */
    public Output<String> childInstanceRouteTableId() {
        return this.childInstanceRouteTableId;
    }

    /**
     * Limit search to a list of specific IDs.The value is formulated as `&lt;cen_id&gt;:&lt;child_instance_route_table_id&gt;:&lt;transit_router_attachment_id&gt;:&lt;destination_cidr_block&gt;`.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return Limit search to a list of specific IDs.The value is formulated as `&lt;cen_id&gt;:&lt;child_instance_route_table_id&gt;:&lt;transit_router_attachment_id&gt;:&lt;destination_cidr_block&gt;`.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
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
     * ServiceType
     * 
     */
    @Import(name="serviceType")
    private @Nullable Output<String> serviceType;

    /**
     * @return ServiceType
     * 
     */
    public Optional<Output<String>> serviceType() {
        return Optional.ofNullable(this.serviceType);
    }

    /**
     * TransitRouterAttachmentId
     * 
     */
    @Import(name="transitRouterAttachmentId", required=true)
    private Output<String> transitRouterAttachmentId;

    /**
     * @return TransitRouterAttachmentId
     * 
     */
    public Output<String> transitRouterAttachmentId() {
        return this.transitRouterAttachmentId;
    }

    private GetChildInstanceRouteEntryToAttachmentsArgs() {}

    private GetChildInstanceRouteEntryToAttachmentsArgs(GetChildInstanceRouteEntryToAttachmentsArgs $) {
        this.cenId = $.cenId;
        this.childInstanceRouteTableId = $.childInstanceRouteTableId;
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.serviceType = $.serviceType;
        this.transitRouterAttachmentId = $.transitRouterAttachmentId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetChildInstanceRouteEntryToAttachmentsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetChildInstanceRouteEntryToAttachmentsArgs $;

        public Builder() {
            $ = new GetChildInstanceRouteEntryToAttachmentsArgs();
        }

        public Builder(GetChildInstanceRouteEntryToAttachmentsArgs defaults) {
            $ = new GetChildInstanceRouteEntryToAttachmentsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cenId The ID of the CEN instance.
         * 
         * @return builder
         * 
         */
        public Builder cenId(@Nullable Output<String> cenId) {
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
         * @param childInstanceRouteTableId The first ID of the resource
         * 
         * @return builder
         * 
         */
        public Builder childInstanceRouteTableId(Output<String> childInstanceRouteTableId) {
            $.childInstanceRouteTableId = childInstanceRouteTableId;
            return this;
        }

        /**
         * @param childInstanceRouteTableId The first ID of the resource
         * 
         * @return builder
         * 
         */
        public Builder childInstanceRouteTableId(String childInstanceRouteTableId) {
            return childInstanceRouteTableId(Output.of(childInstanceRouteTableId));
        }

        /**
         * @param ids Limit search to a list of specific IDs.The value is formulated as `&lt;cen_id&gt;:&lt;child_instance_route_table_id&gt;:&lt;transit_router_attachment_id&gt;:&lt;destination_cidr_block&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids Limit search to a list of specific IDs.The value is formulated as `&lt;cen_id&gt;:&lt;child_instance_route_table_id&gt;:&lt;transit_router_attachment_id&gt;:&lt;destination_cidr_block&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids Limit search to a list of specific IDs.The value is formulated as `&lt;cen_id&gt;:&lt;child_instance_route_table_id&gt;:&lt;transit_router_attachment_id&gt;:&lt;destination_cidr_block&gt;`.
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
         * @param serviceType ServiceType
         * 
         * @return builder
         * 
         */
        public Builder serviceType(@Nullable Output<String> serviceType) {
            $.serviceType = serviceType;
            return this;
        }

        /**
         * @param serviceType ServiceType
         * 
         * @return builder
         * 
         */
        public Builder serviceType(String serviceType) {
            return serviceType(Output.of(serviceType));
        }

        /**
         * @param transitRouterAttachmentId TransitRouterAttachmentId
         * 
         * @return builder
         * 
         */
        public Builder transitRouterAttachmentId(Output<String> transitRouterAttachmentId) {
            $.transitRouterAttachmentId = transitRouterAttachmentId;
            return this;
        }

        /**
         * @param transitRouterAttachmentId TransitRouterAttachmentId
         * 
         * @return builder
         * 
         */
        public Builder transitRouterAttachmentId(String transitRouterAttachmentId) {
            return transitRouterAttachmentId(Output.of(transitRouterAttachmentId));
        }

        public GetChildInstanceRouteEntryToAttachmentsArgs build() {
            $.childInstanceRouteTableId = Objects.requireNonNull($.childInstanceRouteTableId, "expected parameter 'childInstanceRouteTableId' to be non-null");
            $.transitRouterAttachmentId = Objects.requireNonNull($.transitRouterAttachmentId, "expected parameter 'transitRouterAttachmentId' to be non-null");
            return $;
        }
    }

}
