// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetEcsImagePipelinePipeline {
    /**
     * @return The IDs of Alibaba Cloud accounts to which the image was shared.
     * 
     */
    private List<String> addAccounts;
    /**
     * @return The source image.
     * 
     */
    private String baseImage;
    /**
     * @return The type of the source image.
     * 
     */
    private String baseImageType;
    /**
     * @return The content of the image template.
     * 
     */
    private String buildContent;
    /**
     * @return The time when the image template was created.
     * 
     */
    private String creationTime;
    /**
     * @return Indicates whether the intermediate instance was released when the image failed to be created.
     * 
     */
    private Boolean deleteInstanceOnFailure;
    /**
     * @return The description of the image template.
     * 
     */
    private String description;
    /**
     * @return The ID of the Image Pipeline.
     * 
     */
    private String id;
    /**
     * @return The name prefix of the created image.
     * 
     */
    private String imageName;
    /**
     * @return The ID of the image template.
     * 
     */
    private String imagePipelineId;
    /**
     * @return The instance type of the intermediate instance.
     * 
     */
    private String instanceType;
    /**
     * @return The size of the outbound public bandwidth for the intermediate instance. Unit: `Mbit/s`.
     * 
     */
    private Integer internetMaxBandwidthOut;
    /**
     * @return The name of the image template.
     * 
     */
    private String name;
    /**
     * @return The ID of the resource group to which the image template belongs.
     * 
     */
    private String resourceGroupId;
    /**
     * @return The system disk size of the intermediate instance. Unit: `GiB`.
     * 
     */
    private Integer systemDiskSize;
    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    private @Nullable Map<String,Object> tags;
    /**
     * @return The IDs of regions to which to distribute the created image.
     * 
     */
    private List<String> toRegionIds;
    /**
     * @return The vswitch id.
     * 
     */
    private String vswitchId;

    private GetEcsImagePipelinePipeline() {}
    /**
     * @return The IDs of Alibaba Cloud accounts to which the image was shared.
     * 
     */
    public List<String> addAccounts() {
        return this.addAccounts;
    }
    /**
     * @return The source image.
     * 
     */
    public String baseImage() {
        return this.baseImage;
    }
    /**
     * @return The type of the source image.
     * 
     */
    public String baseImageType() {
        return this.baseImageType;
    }
    /**
     * @return The content of the image template.
     * 
     */
    public String buildContent() {
        return this.buildContent;
    }
    /**
     * @return The time when the image template was created.
     * 
     */
    public String creationTime() {
        return this.creationTime;
    }
    /**
     * @return Indicates whether the intermediate instance was released when the image failed to be created.
     * 
     */
    public Boolean deleteInstanceOnFailure() {
        return this.deleteInstanceOnFailure;
    }
    /**
     * @return The description of the image template.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The ID of the Image Pipeline.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name prefix of the created image.
     * 
     */
    public String imageName() {
        return this.imageName;
    }
    /**
     * @return The ID of the image template.
     * 
     */
    public String imagePipelineId() {
        return this.imagePipelineId;
    }
    /**
     * @return The instance type of the intermediate instance.
     * 
     */
    public String instanceType() {
        return this.instanceType;
    }
    /**
     * @return The size of the outbound public bandwidth for the intermediate instance. Unit: `Mbit/s`.
     * 
     */
    public Integer internetMaxBandwidthOut() {
        return this.internetMaxBandwidthOut;
    }
    /**
     * @return The name of the image template.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The ID of the resource group to which the image template belongs.
     * 
     */
    public String resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * @return The system disk size of the intermediate instance. Unit: `GiB`.
     * 
     */
    public Integer systemDiskSize() {
        return this.systemDiskSize;
    }
    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Map<String,Object> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }
    /**
     * @return The IDs of regions to which to distribute the created image.
     * 
     */
    public List<String> toRegionIds() {
        return this.toRegionIds;
    }
    /**
     * @return The vswitch id.
     * 
     */
    public String vswitchId() {
        return this.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEcsImagePipelinePipeline defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> addAccounts;
        private String baseImage;
        private String baseImageType;
        private String buildContent;
        private String creationTime;
        private Boolean deleteInstanceOnFailure;
        private String description;
        private String id;
        private String imageName;
        private String imagePipelineId;
        private String instanceType;
        private Integer internetMaxBandwidthOut;
        private String name;
        private String resourceGroupId;
        private Integer systemDiskSize;
        private @Nullable Map<String,Object> tags;
        private List<String> toRegionIds;
        private String vswitchId;
        public Builder() {}
        public Builder(GetEcsImagePipelinePipeline defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.addAccounts = defaults.addAccounts;
    	      this.baseImage = defaults.baseImage;
    	      this.baseImageType = defaults.baseImageType;
    	      this.buildContent = defaults.buildContent;
    	      this.creationTime = defaults.creationTime;
    	      this.deleteInstanceOnFailure = defaults.deleteInstanceOnFailure;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.imageName = defaults.imageName;
    	      this.imagePipelineId = defaults.imagePipelineId;
    	      this.instanceType = defaults.instanceType;
    	      this.internetMaxBandwidthOut = defaults.internetMaxBandwidthOut;
    	      this.name = defaults.name;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.systemDiskSize = defaults.systemDiskSize;
    	      this.tags = defaults.tags;
    	      this.toRegionIds = defaults.toRegionIds;
    	      this.vswitchId = defaults.vswitchId;
        }

        @CustomType.Setter
        public Builder addAccounts(List<String> addAccounts) {
            this.addAccounts = Objects.requireNonNull(addAccounts);
            return this;
        }
        public Builder addAccounts(String... addAccounts) {
            return addAccounts(List.of(addAccounts));
        }
        @CustomType.Setter
        public Builder baseImage(String baseImage) {
            this.baseImage = Objects.requireNonNull(baseImage);
            return this;
        }
        @CustomType.Setter
        public Builder baseImageType(String baseImageType) {
            this.baseImageType = Objects.requireNonNull(baseImageType);
            return this;
        }
        @CustomType.Setter
        public Builder buildContent(String buildContent) {
            this.buildContent = Objects.requireNonNull(buildContent);
            return this;
        }
        @CustomType.Setter
        public Builder creationTime(String creationTime) {
            this.creationTime = Objects.requireNonNull(creationTime);
            return this;
        }
        @CustomType.Setter
        public Builder deleteInstanceOnFailure(Boolean deleteInstanceOnFailure) {
            this.deleteInstanceOnFailure = Objects.requireNonNull(deleteInstanceOnFailure);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder imageName(String imageName) {
            this.imageName = Objects.requireNonNull(imageName);
            return this;
        }
        @CustomType.Setter
        public Builder imagePipelineId(String imagePipelineId) {
            this.imagePipelineId = Objects.requireNonNull(imagePipelineId);
            return this;
        }
        @CustomType.Setter
        public Builder instanceType(String instanceType) {
            this.instanceType = Objects.requireNonNull(instanceType);
            return this;
        }
        @CustomType.Setter
        public Builder internetMaxBandwidthOut(Integer internetMaxBandwidthOut) {
            this.internetMaxBandwidthOut = Objects.requireNonNull(internetMaxBandwidthOut);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroupId(String resourceGroupId) {
            this.resourceGroupId = Objects.requireNonNull(resourceGroupId);
            return this;
        }
        @CustomType.Setter
        public Builder systemDiskSize(Integer systemDiskSize) {
            this.systemDiskSize = Objects.requireNonNull(systemDiskSize);
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,Object> tags) {
            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder toRegionIds(List<String> toRegionIds) {
            this.toRegionIds = Objects.requireNonNull(toRegionIds);
            return this;
        }
        public Builder toRegionIds(String... toRegionIds) {
            return toRegionIds(List.of(toRegionIds));
        }
        @CustomType.Setter
        public Builder vswitchId(String vswitchId) {
            this.vswitchId = Objects.requireNonNull(vswitchId);
            return this;
        }
        public GetEcsImagePipelinePipeline build() {
            final var o = new GetEcsImagePipelinePipeline();
            o.addAccounts = addAccounts;
            o.baseImage = baseImage;
            o.baseImageType = baseImageType;
            o.buildContent = buildContent;
            o.creationTime = creationTime;
            o.deleteInstanceOnFailure = deleteInstanceOnFailure;
            o.description = description;
            o.id = id;
            o.imageName = imageName;
            o.imagePipelineId = imagePipelineId;
            o.instanceType = instanceType;
            o.internetMaxBandwidthOut = internetMaxBandwidthOut;
            o.name = name;
            o.resourceGroupId = resourceGroupId;
            o.systemDiskSize = systemDiskSize;
            o.tags = tags;
            o.toRegionIds = toRegionIds;
            o.vswitchId = vswitchId;
            return o;
        }
    }
}
