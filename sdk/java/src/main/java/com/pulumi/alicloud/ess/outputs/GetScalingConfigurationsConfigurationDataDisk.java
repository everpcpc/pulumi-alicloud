// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetScalingConfigurationsConfigurationDataDisk {
    /**
     * @return Category of data disk.
     * 
     */
    private @Nullable String category;
    /**
     * @return Delete_with_instance attribute of data disk.
     * 
     */
    private @Nullable Boolean deleteWithInstance;
    /**
     * @return Device attribute of data disk.
     * 
     */
    private @Nullable String device;
    /**
     * @return The performance level of the ESSD used as data disk.
     * 
     */
    private @Nullable String performanceLevel;
    /**
     * @return Size of data disk.
     * 
     */
    private @Nullable Integer size;
    /**
     * @return Size of data disk.
     * 
     */
    private @Nullable String snapshotId;

    private GetScalingConfigurationsConfigurationDataDisk() {}
    /**
     * @return Category of data disk.
     * 
     */
    public Optional<String> category() {
        return Optional.ofNullable(this.category);
    }
    /**
     * @return Delete_with_instance attribute of data disk.
     * 
     */
    public Optional<Boolean> deleteWithInstance() {
        return Optional.ofNullable(this.deleteWithInstance);
    }
    /**
     * @return Device attribute of data disk.
     * 
     */
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    /**
     * @return The performance level of the ESSD used as data disk.
     * 
     */
    public Optional<String> performanceLevel() {
        return Optional.ofNullable(this.performanceLevel);
    }
    /**
     * @return Size of data disk.
     * 
     */
    public Optional<Integer> size() {
        return Optional.ofNullable(this.size);
    }
    /**
     * @return Size of data disk.
     * 
     */
    public Optional<String> snapshotId() {
        return Optional.ofNullable(this.snapshotId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetScalingConfigurationsConfigurationDataDisk defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String category;
        private @Nullable Boolean deleteWithInstance;
        private @Nullable String device;
        private @Nullable String performanceLevel;
        private @Nullable Integer size;
        private @Nullable String snapshotId;
        public Builder() {}
        public Builder(GetScalingConfigurationsConfigurationDataDisk defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.category = defaults.category;
    	      this.deleteWithInstance = defaults.deleteWithInstance;
    	      this.device = defaults.device;
    	      this.performanceLevel = defaults.performanceLevel;
    	      this.size = defaults.size;
    	      this.snapshotId = defaults.snapshotId;
        }

        @CustomType.Setter
        public Builder category(@Nullable String category) {
            this.category = category;
            return this;
        }
        @CustomType.Setter
        public Builder deleteWithInstance(@Nullable Boolean deleteWithInstance) {
            this.deleteWithInstance = deleteWithInstance;
            return this;
        }
        @CustomType.Setter
        public Builder device(@Nullable String device) {
            this.device = device;
            return this;
        }
        @CustomType.Setter
        public Builder performanceLevel(@Nullable String performanceLevel) {
            this.performanceLevel = performanceLevel;
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
        public GetScalingConfigurationsConfigurationDataDisk build() {
            final var o = new GetScalingConfigurationsConfigurationDataDisk();
            o.category = category;
            o.deleteWithInstance = deleteWithInstance;
            o.device = device;
            o.performanceLevel = performanceLevel;
            o.size = size;
            o.snapshotId = snapshotId;
            return o;
        }
    }
}
