// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oss.outputs;

import com.pulumi.alicloud.oss.outputs.GetInstanceAttachmentsAttachment;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetInstanceAttachmentsResult {
    /**
     * @return A list of instance attachments. Each element contains the following attributes:
     * 
     */
    private List<GetInstanceAttachmentsAttachment> attachments;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The instance name.
     * 
     */
    private String instanceName;
    private @Nullable String nameRegex;
    /**
     * @return A list of vpc names.
     * 
     */
    private List<String> names;
    private @Nullable String outputFile;
    /**
     * @return A list of vpc ids.
     * 
     */
    private List<String> vpcIds;

    private GetInstanceAttachmentsResult() {}
    /**
     * @return A list of instance attachments. Each element contains the following attributes:
     * 
     */
    public List<GetInstanceAttachmentsAttachment> attachments() {
        return this.attachments;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The instance name.
     * 
     */
    public String instanceName() {
        return this.instanceName;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of vpc names.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return A list of vpc ids.
     * 
     */
    public List<String> vpcIds() {
        return this.vpcIds;
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
        private String id;
        private String instanceName;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private List<String> vpcIds;
        public Builder() {}
        public Builder(GetInstanceAttachmentsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.attachments = defaults.attachments;
    	      this.id = defaults.id;
    	      this.instanceName = defaults.instanceName;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.vpcIds = defaults.vpcIds;
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
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceName(String instanceName) {
            this.instanceName = Objects.requireNonNull(instanceName);
            return this;
        }
        @CustomType.Setter
        public Builder nameRegex(@Nullable String nameRegex) {
            this.nameRegex = nameRegex;
            return this;
        }
        @CustomType.Setter
        public Builder names(List<String> names) {
            this.names = Objects.requireNonNull(names);
            return this;
        }
        public Builder names(String... names) {
            return names(List.of(names));
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder vpcIds(List<String> vpcIds) {
            this.vpcIds = Objects.requireNonNull(vpcIds);
            return this;
        }
        public Builder vpcIds(String... vpcIds) {
            return vpcIds(List.of(vpcIds));
        }
        public GetInstanceAttachmentsResult build() {
            final var o = new GetInstanceAttachmentsResult();
            o.attachments = attachments;
            o.id = id;
            o.instanceName = instanceName;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.vpcIds = vpcIds;
            return o;
        }
    }
}
