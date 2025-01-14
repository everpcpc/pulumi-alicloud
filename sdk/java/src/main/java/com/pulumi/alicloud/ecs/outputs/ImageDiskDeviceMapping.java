// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ImageDiskDeviceMapping {
    /**
     * @return Specifies the name of a disk in the combined custom image. Value range: /dev/xvda to /dev/xvdz.
     * 
     */
    private @Nullable String device;
    /**
     * @return Specifies the type of a disk in the combined custom image. If you specify this parameter, you can use a data disk snapshot as the data source of a system disk for creating an image. If it is not specified, the disk type is determined by the corresponding snapshot. Valid values: `system`, `data`,
     * 
     */
    private @Nullable String diskType;
    /**
     * @return Specifies the size of a disk in the combined custom image, in GiB. Value range: 5 to 2000.
     * 
     */
    private @Nullable Integer size;
    /**
     * @return Specifies a snapshot that is used to create a combined custom image.
     * 
     */
    private @Nullable String snapshotId;

    private ImageDiskDeviceMapping() {}
    /**
     * @return Specifies the name of a disk in the combined custom image. Value range: /dev/xvda to /dev/xvdz.
     * 
     */
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    /**
     * @return Specifies the type of a disk in the combined custom image. If you specify this parameter, you can use a data disk snapshot as the data source of a system disk for creating an image. If it is not specified, the disk type is determined by the corresponding snapshot. Valid values: `system`, `data`,
     * 
     */
    public Optional<String> diskType() {
        return Optional.ofNullable(this.diskType);
    }
    /**
     * @return Specifies the size of a disk in the combined custom image, in GiB. Value range: 5 to 2000.
     * 
     */
    public Optional<Integer> size() {
        return Optional.ofNullable(this.size);
    }
    /**
     * @return Specifies a snapshot that is used to create a combined custom image.
     * 
     */
    public Optional<String> snapshotId() {
        return Optional.ofNullable(this.snapshotId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ImageDiskDeviceMapping defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String device;
        private @Nullable String diskType;
        private @Nullable Integer size;
        private @Nullable String snapshotId;
        public Builder() {}
        public Builder(ImageDiskDeviceMapping defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.device = defaults.device;
    	      this.diskType = defaults.diskType;
    	      this.size = defaults.size;
    	      this.snapshotId = defaults.snapshotId;
        }

        @CustomType.Setter
        public Builder device(@Nullable String device) {
            this.device = device;
            return this;
        }
        @CustomType.Setter
        public Builder diskType(@Nullable String diskType) {
            this.diskType = diskType;
            return this;
        }
        @CustomType.Setter
        public Builder size(@Nullable Integer size) {
            this.size = size;
            return this;
        }
        @CustomType.Setter
        public Builder snapshotId(@Nullable String snapshotId) {
            this.snapshotId = snapshotId;
            return this;
        }
        public ImageDiskDeviceMapping build() {
            final var o = new ImageDiskDeviceMapping();
            o.device = device;
            o.diskType = diskType;
            o.size = size;
            o.snapshotId = snapshotId;
            return o;
        }
    }
}
