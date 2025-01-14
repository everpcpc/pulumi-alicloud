// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetInstanceAttachmentsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInstanceAttachmentsPlainArgs Empty = new GetInstanceAttachmentsPlainArgs();

    /**
     * The region to which the network to be queried belongs.
     * 
     */
    @Import(name="childInstanceRegionId")
    private @Nullable String childInstanceRegionId;

    /**
     * @return The region to which the network to be queried belongs.
     * 
     */
    public Optional<String> childInstanceRegionId() {
        return Optional.ofNullable(this.childInstanceRegionId);
    }

    /**
     * The type of the associated network. Valid values: `VPC`, `VBR` and `CCN`.
     * 
     */
    @Import(name="childInstanceType")
    private @Nullable String childInstanceType;

    /**
     * @return The type of the associated network. Valid values: `VPC`, `VBR` and `CCN`.
     * 
     */
    public Optional<String> childInstanceType() {
        return Optional.ofNullable(this.childInstanceType);
    }

    /**
     * The ID of the CEN instance.
     * 
     */
    @Import(name="instanceId", required=true)
    private String instanceId;

    /**
     * @return The ID of the CEN instance.
     * 
     */
    public String instanceId() {
        return this.instanceId;
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

    /**
     * The status of the Cen Child Instance Attachment. Valid value: `Attaching`, `Attached` and `Aetaching`.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the Cen Child Instance Attachment. Valid value: `Attaching`, `Attached` and `Aetaching`.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    private GetInstanceAttachmentsPlainArgs() {}

    private GetInstanceAttachmentsPlainArgs(GetInstanceAttachmentsPlainArgs $) {
        this.childInstanceRegionId = $.childInstanceRegionId;
        this.childInstanceType = $.childInstanceType;
        this.instanceId = $.instanceId;
        this.outputFile = $.outputFile;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstanceAttachmentsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstanceAttachmentsPlainArgs $;

        public Builder() {
            $ = new GetInstanceAttachmentsPlainArgs();
        }

        public Builder(GetInstanceAttachmentsPlainArgs defaults) {
            $ = new GetInstanceAttachmentsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param childInstanceRegionId The region to which the network to be queried belongs.
         * 
         * @return builder
         * 
         */
        public Builder childInstanceRegionId(@Nullable String childInstanceRegionId) {
            $.childInstanceRegionId = childInstanceRegionId;
            return this;
        }

        /**
         * @param childInstanceType The type of the associated network. Valid values: `VPC`, `VBR` and `CCN`.
         * 
         * @return builder
         * 
         */
        public Builder childInstanceType(@Nullable String childInstanceType) {
            $.childInstanceType = childInstanceType;
            return this;
        }

        /**
         * @param instanceId The ID of the CEN instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            $.instanceId = instanceId;
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

        /**
         * @param status The status of the Cen Child Instance Attachment. Valid value: `Attaching`, `Attached` and `Aetaching`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        public GetInstanceAttachmentsPlainArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
