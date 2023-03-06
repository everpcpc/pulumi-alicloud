// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.alicloud.ecs.outputs.GetImagesImage;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetImagesResult {
    private @Nullable String actionType;
    /**
     * @return Platform type of the image system: i386 or x86_64.
     * 
     */
    private @Nullable String architecture;
    private @Nullable Boolean dryRun;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of image IDs.
     * 
     */
    private List<String> ids;
    private @Nullable String imageFamily;
    private @Nullable String imageId;
    private @Nullable String imageName;
    private @Nullable String imageOwnerId;
    /**
     * @return A list of images. Each element contains the following attributes:
     * 
     */
    private List<GetImagesImage> images;
    private @Nullable String instanceType;
    private @Nullable Boolean isSupportCloudInit;
    private @Nullable Boolean isSupportIoOptimized;
    private @Nullable Boolean mostRecent;
    private @Nullable String nameRegex;
    private @Nullable String osType;
    private @Nullable String outputFile;
    private @Nullable String owners;
    private @Nullable String resourceGroupId;
    /**
     * @return Snapshot ID.
     * 
     */
    private @Nullable String snapshotId;
    /**
     * @return Status of the image. Possible values: `UnAvailable`, `Available`, `Creating` and `CreateFailed`.
     * 
     */
    private @Nullable String status;
    private @Nullable Map<String,Object> tags;
    private @Nullable String usage;

    private GetImagesResult() {}
    public Optional<String> actionType() {
        return Optional.ofNullable(this.actionType);
    }
    /**
     * @return Platform type of the image system: i386 or x86_64.
     * 
     */
    public Optional<String> architecture() {
        return Optional.ofNullable(this.architecture);
    }
    public Optional<Boolean> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of image IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> imageFamily() {
        return Optional.ofNullable(this.imageFamily);
    }
    public Optional<String> imageId() {
        return Optional.ofNullable(this.imageId);
    }
    public Optional<String> imageName() {
        return Optional.ofNullable(this.imageName);
    }
    public Optional<String> imageOwnerId() {
        return Optional.ofNullable(this.imageOwnerId);
    }
    /**
     * @return A list of images. Each element contains the following attributes:
     * 
     */
    public List<GetImagesImage> images() {
        return this.images;
    }
    public Optional<String> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }
    public Optional<Boolean> isSupportCloudInit() {
        return Optional.ofNullable(this.isSupportCloudInit);
    }
    public Optional<Boolean> isSupportIoOptimized() {
        return Optional.ofNullable(this.isSupportIoOptimized);
    }
    public Optional<Boolean> mostRecent() {
        return Optional.ofNullable(this.mostRecent);
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    public Optional<String> osType() {
        return Optional.ofNullable(this.osType);
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Optional<String> owners() {
        return Optional.ofNullable(this.owners);
    }
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }
    /**
     * @return Snapshot ID.
     * 
     */
    public Optional<String> snapshotId() {
        return Optional.ofNullable(this.snapshotId);
    }
    /**
     * @return Status of the image. Possible values: `UnAvailable`, `Available`, `Creating` and `CreateFailed`.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    public Map<String,Object> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }
    public Optional<String> usage() {
        return Optional.ofNullable(this.usage);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetImagesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String actionType;
        private @Nullable String architecture;
        private @Nullable Boolean dryRun;
        private String id;
        private List<String> ids;
        private @Nullable String imageFamily;
        private @Nullable String imageId;
        private @Nullable String imageName;
        private @Nullable String imageOwnerId;
        private List<GetImagesImage> images;
        private @Nullable String instanceType;
        private @Nullable Boolean isSupportCloudInit;
        private @Nullable Boolean isSupportIoOptimized;
        private @Nullable Boolean mostRecent;
        private @Nullable String nameRegex;
        private @Nullable String osType;
        private @Nullable String outputFile;
        private @Nullable String owners;
        private @Nullable String resourceGroupId;
        private @Nullable String snapshotId;
        private @Nullable String status;
        private @Nullable Map<String,Object> tags;
        private @Nullable String usage;
        public Builder() {}
        public Builder(GetImagesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.actionType = defaults.actionType;
    	      this.architecture = defaults.architecture;
    	      this.dryRun = defaults.dryRun;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.imageFamily = defaults.imageFamily;
    	      this.imageId = defaults.imageId;
    	      this.imageName = defaults.imageName;
    	      this.imageOwnerId = defaults.imageOwnerId;
    	      this.images = defaults.images;
    	      this.instanceType = defaults.instanceType;
    	      this.isSupportCloudInit = defaults.isSupportCloudInit;
    	      this.isSupportIoOptimized = defaults.isSupportIoOptimized;
    	      this.mostRecent = defaults.mostRecent;
    	      this.nameRegex = defaults.nameRegex;
    	      this.osType = defaults.osType;
    	      this.outputFile = defaults.outputFile;
    	      this.owners = defaults.owners;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.snapshotId = defaults.snapshotId;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.usage = defaults.usage;
        }

        @CustomType.Setter
        public Builder actionType(@Nullable String actionType) {
            this.actionType = actionType;
            return this;
        }
        @CustomType.Setter
        public Builder architecture(@Nullable String architecture) {
            this.architecture = architecture;
            return this;
        }
        @CustomType.Setter
        public Builder dryRun(@Nullable Boolean dryRun) {
            this.dryRun = dryRun;
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
        public Builder imageFamily(@Nullable String imageFamily) {
            this.imageFamily = imageFamily;
            return this;
        }
        @CustomType.Setter
        public Builder imageId(@Nullable String imageId) {
            this.imageId = imageId;
            return this;
        }
        @CustomType.Setter
        public Builder imageName(@Nullable String imageName) {
            this.imageName = imageName;
            return this;
        }
        @CustomType.Setter
        public Builder imageOwnerId(@Nullable String imageOwnerId) {
            this.imageOwnerId = imageOwnerId;
            return this;
        }
        @CustomType.Setter
        public Builder images(List<GetImagesImage> images) {
            this.images = Objects.requireNonNull(images);
            return this;
        }
        public Builder images(GetImagesImage... images) {
            return images(List.of(images));
        }
        @CustomType.Setter
        public Builder instanceType(@Nullable String instanceType) {
            this.instanceType = instanceType;
            return this;
        }
        @CustomType.Setter
        public Builder isSupportCloudInit(@Nullable Boolean isSupportCloudInit) {
            this.isSupportCloudInit = isSupportCloudInit;
            return this;
        }
        @CustomType.Setter
        public Builder isSupportIoOptimized(@Nullable Boolean isSupportIoOptimized) {
            this.isSupportIoOptimized = isSupportIoOptimized;
            return this;
        }
        @CustomType.Setter
        public Builder mostRecent(@Nullable Boolean mostRecent) {
            this.mostRecent = mostRecent;
            return this;
        }
        @CustomType.Setter
        public Builder nameRegex(@Nullable String nameRegex) {
            this.nameRegex = nameRegex;
            return this;
        }
        @CustomType.Setter
        public Builder osType(@Nullable String osType) {
            this.osType = osType;
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder owners(@Nullable String owners) {
            this.owners = owners;
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            this.resourceGroupId = resourceGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder snapshotId(@Nullable String snapshotId) {
            this.snapshotId = snapshotId;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,Object> tags) {
            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder usage(@Nullable String usage) {
            this.usage = usage;
            return this;
        }
        public GetImagesResult build() {
            final var o = new GetImagesResult();
            o.actionType = actionType;
            o.architecture = architecture;
            o.dryRun = dryRun;
            o.id = id;
            o.ids = ids;
            o.imageFamily = imageFamily;
            o.imageId = imageId;
            o.imageName = imageName;
            o.imageOwnerId = imageOwnerId;
            o.images = images;
            o.instanceType = instanceType;
            o.isSupportCloudInit = isSupportCloudInit;
            o.isSupportIoOptimized = isSupportIoOptimized;
            o.mostRecent = mostRecent;
            o.nameRegex = nameRegex;
            o.osType = osType;
            o.outputFile = outputFile;
            o.owners = owners;
            o.resourceGroupId = resourceGroupId;
            o.snapshotId = snapshotId;
            o.status = status;
            o.tags = tags;
            o.usage = usage;
            return o;
        }
    }
}