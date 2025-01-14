// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetInstanceAttachmentsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInstanceAttachmentsArgs Empty = new GetInstanceAttachmentsArgs();

    /**
     * The region to which the network to be queried belongs.
     * 
     */
    @Import(name="childInstanceRegionId")
    private @Nullable Output<String> childInstanceRegionId;

    /**
     * @return The region to which the network to be queried belongs.
     * 
     */
    public Optional<Output<String>> childInstanceRegionId() {
        return Optional.ofNullable(this.childInstanceRegionId);
    }

    /**
     * The type of the associated network. Valid values: `VPC`, `VBR` and `CCN`.
     * 
     */
    @Import(name="childInstanceType")
    private @Nullable Output<String> childInstanceType;

    /**
     * @return The type of the associated network. Valid values: `VPC`, `VBR` and `CCN`.
     * 
     */
    public Optional<Output<String>> childInstanceType() {
        return Optional.ofNullable(this.childInstanceType);
    }

    /**
     * The ID of the CEN instance.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return The ID of the CEN instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
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
     * The status of the Cen Child Instance Attachment. Valid value: `Attaching`, `Attached` and `Aetaching`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the Cen Child Instance Attachment. Valid value: `Attaching`, `Attached` and `Aetaching`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private GetInstanceAttachmentsArgs() {}

    private GetInstanceAttachmentsArgs(GetInstanceAttachmentsArgs $) {
        this.childInstanceRegionId = $.childInstanceRegionId;
        this.childInstanceType = $.childInstanceType;
        this.instanceId = $.instanceId;
        this.outputFile = $.outputFile;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstanceAttachmentsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstanceAttachmentsArgs $;

        public Builder() {
            $ = new GetInstanceAttachmentsArgs();
        }

        public Builder(GetInstanceAttachmentsArgs defaults) {
            $ = new GetInstanceAttachmentsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param childInstanceRegionId The region to which the network to be queried belongs.
         * 
         * @return builder
         * 
         */
        public Builder childInstanceRegionId(@Nullable Output<String> childInstanceRegionId) {
            $.childInstanceRegionId = childInstanceRegionId;
            return this;
        }

        /**
         * @param childInstanceRegionId The region to which the network to be queried belongs.
         * 
         * @return builder
         * 
         */
        public Builder childInstanceRegionId(String childInstanceRegionId) {
            return childInstanceRegionId(Output.of(childInstanceRegionId));
        }

        /**
         * @param childInstanceType The type of the associated network. Valid values: `VPC`, `VBR` and `CCN`.
         * 
         * @return builder
         * 
         */
        public Builder childInstanceType(@Nullable Output<String> childInstanceType) {
            $.childInstanceType = childInstanceType;
            return this;
        }

        /**
         * @param childInstanceType The type of the associated network. Valid values: `VPC`, `VBR` and `CCN`.
         * 
         * @return builder
         * 
         */
        public Builder childInstanceType(String childInstanceType) {
            return childInstanceType(Output.of(childInstanceType));
        }

        /**
         * @param instanceId The ID of the CEN instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The ID of the CEN instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
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
         * @param status The status of the Cen Child Instance Attachment. Valid value: `Attaching`, `Attached` and `Aetaching`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the Cen Child Instance Attachment. Valid value: `Attaching`, `Attached` and `Aetaching`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public GetInstanceAttachmentsArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
