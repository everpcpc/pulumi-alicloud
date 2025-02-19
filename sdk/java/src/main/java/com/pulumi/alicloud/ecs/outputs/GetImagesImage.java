// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.alicloud.ecs.outputs.GetImagesImageDiskDeviceMapping;
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
public final class GetImagesImage {
    /**
     * @return The image architecture. Valid values: `i386` and `x86_64`.
     * 
     */
    private String architecture;
    /**
     * @return Time of creation.
     * 
     */
    private String creationTime;
    /**
     * @return Description of the image.
     * 
     */
    private String description;
    /**
     * @return Description of the system with disks and snapshots under the image.
     * 
     */
    private List<GetImagesImageDiskDeviceMapping> diskDeviceMappings;
    /**
     * @return ID of the image.
     * 
     */
    private String id;
    /**
     * @return The ID of the image.
     * 
     */
    private String imageId;
    /**
     * @return Alias of the image owner.
     * 
     */
    private String imageOwnerAlias;
    /**
     * @return Version of the image.
     * 
     */
    private String imageVersion;
    private Boolean isCopied;
    private String isSelfShared;
    /**
     * @return Whether the user has subscribed to the terms of service for the image product corresponding to the ProductCode.
     * 
     */
    private Boolean isSubscribed;
    /**
     * @return Specifies whether the image can be used on I/O optimized instances.
     * 
     */
    private Boolean isSupportIoOptimized;
    private String name;
    /**
     * @return Display Chinese name of the OS.
     * 
     */
    private String osName;
    /**
     * @return Display English name of the OS.
     * 
     */
    private String osNameEn;
    /**
     * @return The operating system type of the image. Valid values: `windows` and `linux`.
     * 
     */
    private String osType;
    private String platform;
    /**
     * @return Product code of the image on the image market.
     * 
     */
    private String productCode;
    /**
     * @return Progress of image creation, presented in percentages.
     * 
     */
    private String progress;
    /**
     * @return Size of the created disk.
     * 
     */
    private Integer size;
    private String state;
    /**
     * @return The status of the image. The following values are available, Separate multiple parameter values by using commas (,). Default value: `Available`. Valid values:
     * 
     */
    private String status;
    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    private @Nullable Map<String,Object> tags;
    /**
     * @return Specifies whether to check the validity of the request without actually making the request. Valid values:
     * 
     */
    private String usage;

    private GetImagesImage() {}
    /**
     * @return The image architecture. Valid values: `i386` and `x86_64`.
     * 
     */
    public String architecture() {
        return this.architecture;
    }
    /**
     * @return Time of creation.
     * 
     */
    public String creationTime() {
        return this.creationTime;
    }
    /**
     * @return Description of the image.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return Description of the system with disks and snapshots under the image.
     * 
     */
    public List<GetImagesImageDiskDeviceMapping> diskDeviceMappings() {
        return this.diskDeviceMappings;
    }
    /**
     * @return ID of the image.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The ID of the image.
     * 
     */
    public String imageId() {
        return this.imageId;
    }
    /**
     * @return Alias of the image owner.
     * 
     */
    public String imageOwnerAlias() {
        return this.imageOwnerAlias;
    }
    /**
     * @return Version of the image.
     * 
     */
    public String imageVersion() {
        return this.imageVersion;
    }
    public Boolean isCopied() {
        return this.isCopied;
    }
    public String isSelfShared() {
        return this.isSelfShared;
    }
    /**
     * @return Whether the user has subscribed to the terms of service for the image product corresponding to the ProductCode.
     * 
     */
    public Boolean isSubscribed() {
        return this.isSubscribed;
    }
    /**
     * @return Specifies whether the image can be used on I/O optimized instances.
     * 
     */
    public Boolean isSupportIoOptimized() {
        return this.isSupportIoOptimized;
    }
    public String name() {
        return this.name;
    }
    /**
     * @return Display Chinese name of the OS.
     * 
     */
    public String osName() {
        return this.osName;
    }
    /**
     * @return Display English name of the OS.
     * 
     */
    public String osNameEn() {
        return this.osNameEn;
    }
    /**
     * @return The operating system type of the image. Valid values: `windows` and `linux`.
     * 
     */
    public String osType() {
        return this.osType;
    }
    public String platform() {
        return this.platform;
    }
    /**
     * @return Product code of the image on the image market.
     * 
     */
    public String productCode() {
        return this.productCode;
    }
    /**
     * @return Progress of image creation, presented in percentages.
     * 
     */
    public String progress() {
        return this.progress;
    }
    /**
     * @return Size of the created disk.
     * 
     */
    public Integer size() {
        return this.size;
    }
    public String state() {
        return this.state;
    }
    /**
     * @return The status of the image. The following values are available, Separate multiple parameter values by using commas (,). Default value: `Available`. Valid values:
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Map<String,Object> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }
    /**
     * @return Specifies whether to check the validity of the request without actually making the request. Valid values:
     * 
     */
    public String usage() {
        return this.usage;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetImagesImage defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String architecture;
        private String creationTime;
        private String description;
        private List<GetImagesImageDiskDeviceMapping> diskDeviceMappings;
        private String id;
        private String imageId;
        private String imageOwnerAlias;
        private String imageVersion;
        private Boolean isCopied;
        private String isSelfShared;
        private Boolean isSubscribed;
        private Boolean isSupportIoOptimized;
        private String name;
        private String osName;
        private String osNameEn;
        private String osType;
        private String platform;
        private String productCode;
        private String progress;
        private Integer size;
        private String state;
        private String status;
        private @Nullable Map<String,Object> tags;
        private String usage;
        public Builder() {}
        public Builder(GetImagesImage defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.architecture = defaults.architecture;
    	      this.creationTime = defaults.creationTime;
    	      this.description = defaults.description;
    	      this.diskDeviceMappings = defaults.diskDeviceMappings;
    	      this.id = defaults.id;
    	      this.imageId = defaults.imageId;
    	      this.imageOwnerAlias = defaults.imageOwnerAlias;
    	      this.imageVersion = defaults.imageVersion;
    	      this.isCopied = defaults.isCopied;
    	      this.isSelfShared = defaults.isSelfShared;
    	      this.isSubscribed = defaults.isSubscribed;
    	      this.isSupportIoOptimized = defaults.isSupportIoOptimized;
    	      this.name = defaults.name;
    	      this.osName = defaults.osName;
    	      this.osNameEn = defaults.osNameEn;
    	      this.osType = defaults.osType;
    	      this.platform = defaults.platform;
    	      this.productCode = defaults.productCode;
    	      this.progress = defaults.progress;
    	      this.size = defaults.size;
    	      this.state = defaults.state;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.usage = defaults.usage;
        }

        @CustomType.Setter
        public Builder architecture(String architecture) {
            this.architecture = Objects.requireNonNull(architecture);
            return this;
        }
        @CustomType.Setter
        public Builder creationTime(String creationTime) {
            this.creationTime = Objects.requireNonNull(creationTime);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder diskDeviceMappings(List<GetImagesImageDiskDeviceMapping> diskDeviceMappings) {
            this.diskDeviceMappings = Objects.requireNonNull(diskDeviceMappings);
            return this;
        }
        public Builder diskDeviceMappings(GetImagesImageDiskDeviceMapping... diskDeviceMappings) {
            return diskDeviceMappings(List.of(diskDeviceMappings));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder imageId(String imageId) {
            this.imageId = Objects.requireNonNull(imageId);
            return this;
        }
        @CustomType.Setter
        public Builder imageOwnerAlias(String imageOwnerAlias) {
            this.imageOwnerAlias = Objects.requireNonNull(imageOwnerAlias);
            return this;
        }
        @CustomType.Setter
        public Builder imageVersion(String imageVersion) {
            this.imageVersion = Objects.requireNonNull(imageVersion);
            return this;
        }
        @CustomType.Setter
        public Builder isCopied(Boolean isCopied) {
            this.isCopied = Objects.requireNonNull(isCopied);
            return this;
        }
        @CustomType.Setter
        public Builder isSelfShared(String isSelfShared) {
            this.isSelfShared = Objects.requireNonNull(isSelfShared);
            return this;
        }
        @CustomType.Setter
        public Builder isSubscribed(Boolean isSubscribed) {
            this.isSubscribed = Objects.requireNonNull(isSubscribed);
            return this;
        }
        @CustomType.Setter
        public Builder isSupportIoOptimized(Boolean isSupportIoOptimized) {
            this.isSupportIoOptimized = Objects.requireNonNull(isSupportIoOptimized);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder osName(String osName) {
            this.osName = Objects.requireNonNull(osName);
            return this;
        }
        @CustomType.Setter
        public Builder osNameEn(String osNameEn) {
            this.osNameEn = Objects.requireNonNull(osNameEn);
            return this;
        }
        @CustomType.Setter
        public Builder osType(String osType) {
            this.osType = Objects.requireNonNull(osType);
            return this;
        }
        @CustomType.Setter
        public Builder platform(String platform) {
            this.platform = Objects.requireNonNull(platform);
            return this;
        }
        @CustomType.Setter
        public Builder productCode(String productCode) {
            this.productCode = Objects.requireNonNull(productCode);
            return this;
        }
        @CustomType.Setter
        public Builder progress(String progress) {
            this.progress = Objects.requireNonNull(progress);
            return this;
        }
        @CustomType.Setter
        public Builder size(Integer size) {
            this.size = Objects.requireNonNull(size);
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            this.state = Objects.requireNonNull(state);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,Object> tags) {
            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder usage(String usage) {
            this.usage = Objects.requireNonNull(usage);
            return this;
        }
        public GetImagesImage build() {
            final var o = new GetImagesImage();
            o.architecture = architecture;
            o.creationTime = creationTime;
            o.description = description;
            o.diskDeviceMappings = diskDeviceMappings;
            o.id = id;
            o.imageId = imageId;
            o.imageOwnerAlias = imageOwnerAlias;
            o.imageVersion = imageVersion;
            o.isCopied = isCopied;
            o.isSelfShared = isSelfShared;
            o.isSubscribed = isSubscribed;
            o.isSupportIoOptimized = isSupportIoOptimized;
            o.name = name;
            o.osName = osName;
            o.osNameEn = osNameEn;
            o.osType = osType;
            o.platform = platform;
            o.productCode = productCode;
            o.progress = progress;
            o.size = size;
            o.state = state;
            o.status = status;
            o.tags = tags;
            o.usage = usage;
            return o;
        }
    }
}
