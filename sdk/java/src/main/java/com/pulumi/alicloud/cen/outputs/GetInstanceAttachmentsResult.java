// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.outputs;

import com.pulumi.alicloud.cen.outputs.GetInstanceAttachmentsAttachment;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetInstanceAttachmentsResult {
    /**
     * @return A list of CEN Instance Attachments. Each element contains the following attributes:
     * 
     */
    private List<GetInstanceAttachmentsAttachment> attachments;
    /**
     * @return The ID of the region to which the network belongs.
     * 
     */
    private @Nullable String childInstanceRegionId;
    /**
     * @return The type of the associated network.
     * 
     */
    private @Nullable String childInstanceType;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of CEN Instance Attachment IDs.
     * 
     */
    private List<String> ids;
    /**
     * @return The ID of the CEN instance.
     * 
     */
    private String instanceId;
    private @Nullable String outputFile;
    /**
     * @return The status of the network.
     * 
     */
    private @Nullable String status;

    private GetInstanceAttachmentsResult() {}
    /**
     * @return A list of CEN Instance Attachments. Each element contains the following attributes:
     * 
     */
    public List<GetInstanceAttachmentsAttachment> attachments() {
        return this.attachments;
    }
    /**
     * @return The ID of the region to which the network belongs.
     * 
     */
    public Optional<String> childInstanceRegionId() {
        return Optional.ofNullable(this.childInstanceRegionId);
    }
    /**
     * @return The type of the associated network.
     * 
     */
    public Optional<String> childInstanceType() {
        return Optional.ofNullable(this.childInstanceType);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of CEN Instance Attachment IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    /**
     * @return The ID of the CEN instance.
     * 
     */
    public String instanceId() {
        return this.instanceId;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return The status of the network.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstanceAttachmentsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetInstanceAttachmentsAttachment> attachments;
        private @Nullable String childInstanceRegionId;
        private @Nullable String childInstanceType;
        private String id;
        private List<String> ids;
        private String instanceId;
        private @Nullable String outputFile;
        private @Nullable String status;
        public Builder() {}
        public Builder(GetInstanceAttachmentsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.attachments = defaults.attachments;
    	      this.childInstanceRegionId = defaults.childInstanceRegionId;
    	      this.childInstanceType = defaults.childInstanceType;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.instanceId = defaults.instanceId;
    	      this.outputFile = defaults.outputFile;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder attachments(List<GetInstanceAttachmentsAttachment> attachments) {
            this.attachments = Objects.requireNonNull(attachments);
            return this;
        }
        public Builder attachments(GetInstanceAttachmentsAttachment... attachments) {
            return attachments(List.of(attachments));
        }
        @CustomType.Setter
        public Builder childInstanceRegionId(@Nullable String childInstanceRegionId) {
            this.childInstanceRegionId = childInstanceRegionId;
            return this;
        }
        @CustomType.Setter
        public Builder childInstanceType(@Nullable String childInstanceType) {
            this.childInstanceType = childInstanceType;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ids(List<String> ids) {
            this.ids = Objects.requireNonNull(ids);
            return this;
        }
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }
        @CustomType.Setter
        public Builder instanceId(String instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        public GetInstanceAttachmentsResult build() {
            final var o = new GetInstanceAttachmentsResult();
            o.attachments = attachments;
            o.childInstanceRegionId = childInstanceRegionId;
            o.childInstanceType = childInstanceType;
            o.id = id;
            o.ids = ids;
            o.instanceId = instanceId;
            o.outputFile = outputFile;
            o.status = status;
            return o;
        }
    }
}
