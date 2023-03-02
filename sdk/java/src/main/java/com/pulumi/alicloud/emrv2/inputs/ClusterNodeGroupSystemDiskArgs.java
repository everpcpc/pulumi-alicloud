// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.emrv2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterNodeGroupSystemDiskArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterNodeGroupSystemDiskArgs Empty = new ClusterNodeGroupSystemDiskArgs();

    /**
     * The type of the data disk. Valid values: `cloud_efficiency` and `cloud_essd`.
     * 
     */
    @Import(name="category", required=true)
    private Output<String> category;

    /**
     * @return The type of the data disk. Valid values: `cloud_efficiency` and `cloud_essd`.
     * 
     */
    public Output<String> category() {
        return this.category;
    }

    /**
     * The count of a data disk.
     * 
     */
    @Import(name="count")
    private @Nullable Output<Integer> count;

    /**
     * @return The count of a data disk.
     * 
     */
    public Optional<Output<Integer>> count() {
        return Optional.ofNullable(this.count);
    }

    /**
     * Worker node data disk performance level, when `category` values `cloud_essd`, the optional values are `PL0`, `PL1`, `PL2` or `PL3`, but the specific performance level is related to the disk capacity.
     * 
     */
    @Import(name="performanceLevel")
    private @Nullable Output<String> performanceLevel;

    /**
     * @return Worker node data disk performance level, when `category` values `cloud_essd`, the optional values are `PL0`, `PL1`, `PL2` or `PL3`, but the specific performance level is related to the disk capacity.
     * 
     */
    public Optional<Output<String>> performanceLevel() {
        return Optional.ofNullable(this.performanceLevel);
    }

    /**
     * The size of a data disk, at least 40. Unit: GiB.
     * 
     */
    @Import(name="size", required=true)
    private Output<Integer> size;

    /**
     * @return The size of a data disk, at least 40. Unit: GiB.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }

    private ClusterNodeGroupSystemDiskArgs() {}

    private ClusterNodeGroupSystemDiskArgs(ClusterNodeGroupSystemDiskArgs $) {
        this.category = $.category;
        this.count = $.count;
        this.performanceLevel = $.performanceLevel;
        this.size = $.size;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterNodeGroupSystemDiskArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterNodeGroupSystemDiskArgs $;

        public Builder() {
            $ = new ClusterNodeGroupSystemDiskArgs();
        }

        public Builder(ClusterNodeGroupSystemDiskArgs defaults) {
            $ = new ClusterNodeGroupSystemDiskArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param category The type of the data disk. Valid values: `cloud_efficiency` and `cloud_essd`.
         * 
         * @return builder
         * 
         */
        public Builder category(Output<String> category) {
            $.category = category;
            return this;
        }

        /**
         * @param category The type of the data disk. Valid values: `cloud_efficiency` and `cloud_essd`.
         * 
         * @return builder
         * 
         */
        public Builder category(String category) {
            return category(Output.of(category));
        }

        /**
         * @param count The count of a data disk.
         * 
         * @return builder
         * 
         */
        public Builder count(@Nullable Output<Integer> count) {
            $.count = count;
            return this;
        }

        /**
         * @param count The count of a data disk.
         * 
         * @return builder
         * 
         */
        public Builder count(Integer count) {
            return count(Output.of(count));
        }

        /**
         * @param performanceLevel Worker node data disk performance level, when `category` values `cloud_essd`, the optional values are `PL0`, `PL1`, `PL2` or `PL3`, but the specific performance level is related to the disk capacity.
         * 
         * @return builder
         * 
         */
        public Builder performanceLevel(@Nullable Output<String> performanceLevel) {
            $.performanceLevel = performanceLevel;
            return this;
        }

        /**
         * @param performanceLevel Worker node data disk performance level, when `category` values `cloud_essd`, the optional values are `PL0`, `PL1`, `PL2` or `PL3`, but the specific performance level is related to the disk capacity.
         * 
         * @return builder
         * 
         */
        public Builder performanceLevel(String performanceLevel) {
            return performanceLevel(Output.of(performanceLevel));
        }

        /**
         * @param size The size of a data disk, at least 40. Unit: GiB.
         * 
         * @return builder
         * 
         */
        public Builder size(Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size The size of a data disk, at least 40. Unit: GiB.
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        public ClusterNodeGroupSystemDiskArgs build() {
            $.category = Objects.requireNonNull($.category, "expected parameter 'category' to be non-null");
            $.size = Objects.requireNonNull($.size, "expected parameter 'size' to be non-null");
            return $;
        }
    }

}
