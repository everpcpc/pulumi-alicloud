// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ImageCopyArgs extends com.pulumi.resources.ResourceArgs {

    public static final ImageCopyArgs Empty = new ImageCopyArgs();

    @Import(name="deleteAutoSnapshot")
    private @Nullable Output<Boolean> deleteAutoSnapshot;

    public Optional<Output<Boolean>> deleteAutoSnapshot() {
        return Optional.ofNullable(this.deleteAutoSnapshot);
    }

    /**
     * The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Indicates whether to encrypt the image.
     * 
     */
    @Import(name="encrypted")
    private @Nullable Output<Boolean> encrypted;

    /**
     * @return Indicates whether to encrypt the image.
     * 
     */
    public Optional<Output<Boolean>> encrypted() {
        return Optional.ofNullable(this.encrypted);
    }

    /**
     * Indicates whether to force delete the custom image, Default is `false`.
     * - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
     * - false：Verifies that the image is not currently in use by any other instances before deleting the image.
     * 
     */
    @Import(name="force")
    private @Nullable Output<Boolean> force;

    /**
     * @return Indicates whether to force delete the custom image, Default is `false`.
     * - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
     * - false：Verifies that the image is not currently in use by any other instances before deleting the image.
     * 
     */
    public Optional<Output<Boolean>> force() {
        return Optional.ofNullable(this.force);
    }

    /**
     * The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
     * 
     */
    @Import(name="imageName")
    private @Nullable Output<String> imageName;

    /**
     * @return The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
     * 
     */
    public Optional<Output<String>> imageName() {
        return Optional.ofNullable(this.imageName);
    }

    /**
     * Key ID used to encrypt the image.
     * 
     */
    @Import(name="kmsKeyId")
    private @Nullable Output<String> kmsKeyId;

    /**
     * @return Key ID used to encrypt the image.
     * 
     */
    public Optional<Output<String>> kmsKeyId() {
        return Optional.ofNullable(this.kmsKeyId);
    }

    /**
     * @deprecated
     * Attribute &#39;name&#39; has been deprecated from version 1.69.0. Use `image_name` instead.
     * 
     */
    @Deprecated /* Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead. */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @deprecated
     * Attribute &#39;name&#39; has been deprecated from version 1.69.0. Use `image_name` instead.
     * 
     */
    @Deprecated /* Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead. */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The source image ID.
     * 
     */
    @Import(name="sourceImageId", required=true)
    private Output<String> sourceImageId;

    /**
     * @return The source image ID.
     * 
     */
    public Output<String> sourceImageId() {
        return this.sourceImageId;
    }

    /**
     * The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
     * 
     */
    @Import(name="sourceRegionId", required=true)
    private Output<String> sourceRegionId;

    /**
     * @return The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
     * 
     */
    public Output<String> sourceRegionId() {
        return this.sourceRegionId;
    }

    /**
     * The tag value of an image. The value of N ranges from 1 to 20.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return The tag value of an image. The value of N ranges from 1 to 20.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private ImageCopyArgs() {}

    private ImageCopyArgs(ImageCopyArgs $) {
        this.deleteAutoSnapshot = $.deleteAutoSnapshot;
        this.description = $.description;
        this.encrypted = $.encrypted;
        this.force = $.force;
        this.imageName = $.imageName;
        this.kmsKeyId = $.kmsKeyId;
        this.name = $.name;
        this.sourceImageId = $.sourceImageId;
        this.sourceRegionId = $.sourceRegionId;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ImageCopyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ImageCopyArgs $;

        public Builder() {
            $ = new ImageCopyArgs();
        }

        public Builder(ImageCopyArgs defaults) {
            $ = new ImageCopyArgs(Objects.requireNonNull(defaults));
        }

        public Builder deleteAutoSnapshot(@Nullable Output<Boolean> deleteAutoSnapshot) {
            $.deleteAutoSnapshot = deleteAutoSnapshot;
            return this;
        }

        public Builder deleteAutoSnapshot(Boolean deleteAutoSnapshot) {
            return deleteAutoSnapshot(Output.of(deleteAutoSnapshot));
        }

        /**
         * @param description The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the image. It must be 2 to 256 characters in length and must not start with http:// or https://. Default value: null.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param encrypted Indicates whether to encrypt the image.
         * 
         * @return builder
         * 
         */
        public Builder encrypted(@Nullable Output<Boolean> encrypted) {
            $.encrypted = encrypted;
            return this;
        }

        /**
         * @param encrypted Indicates whether to encrypt the image.
         * 
         * @return builder
         * 
         */
        public Builder encrypted(Boolean encrypted) {
            return encrypted(Output.of(encrypted));
        }

        /**
         * @param force Indicates whether to force delete the custom image, Default is `false`.
         * - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
         * - false：Verifies that the image is not currently in use by any other instances before deleting the image.
         * 
         * @return builder
         * 
         */
        public Builder force(@Nullable Output<Boolean> force) {
            $.force = force;
            return this;
        }

        /**
         * @param force Indicates whether to force delete the custom image, Default is `false`.
         * - true：Force deletes the custom image, regardless of whether the image is currently being used by other instances.
         * - false：Verifies that the image is not currently in use by any other instances before deleting the image.
         * 
         * @return builder
         * 
         */
        public Builder force(Boolean force) {
            return force(Output.of(force));
        }

        /**
         * @param imageName The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
         * 
         * @return builder
         * 
         */
        public Builder imageName(@Nullable Output<String> imageName) {
            $.imageName = imageName;
            return this;
        }

        /**
         * @param imageName The image name. It must be 2 to 128 characters in length, and must begin with a letter or Chinese character (beginning with http:// or https:// is not allowed). It can contain digits, colons (:), underscores (_), or hyphens (-). Default value: null.
         * 
         * @return builder
         * 
         */
        public Builder imageName(String imageName) {
            return imageName(Output.of(imageName));
        }

        /**
         * @param kmsKeyId Key ID used to encrypt the image.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(@Nullable Output<String> kmsKeyId) {
            $.kmsKeyId = kmsKeyId;
            return this;
        }

        /**
         * @param kmsKeyId Key ID used to encrypt the image.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(String kmsKeyId) {
            return kmsKeyId(Output.of(kmsKeyId));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Attribute &#39;name&#39; has been deprecated from version 1.69.0. Use `image_name` instead.
         * 
         */
        @Deprecated /* Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead. */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Attribute &#39;name&#39; has been deprecated from version 1.69.0. Use `image_name` instead.
         * 
         */
        @Deprecated /* Attribute 'name' has been deprecated from version 1.69.0. Use `image_name` instead. */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param sourceImageId The source image ID.
         * 
         * @return builder
         * 
         */
        public Builder sourceImageId(Output<String> sourceImageId) {
            $.sourceImageId = sourceImageId;
            return this;
        }

        /**
         * @param sourceImageId The source image ID.
         * 
         * @return builder
         * 
         */
        public Builder sourceImageId(String sourceImageId) {
            return sourceImageId(Output.of(sourceImageId));
        }

        /**
         * @param sourceRegionId The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
         * 
         * @return builder
         * 
         */
        public Builder sourceRegionId(Output<String> sourceRegionId) {
            $.sourceRegionId = sourceRegionId;
            return this;
        }

        /**
         * @param sourceRegionId The ID of the region to which the source custom image belongs. You can call [DescribeRegions](https://www.alibabacloud.com/help/doc-detail/25609.htm) to view the latest regions of Alibaba Cloud.
         * 
         * @return builder
         * 
         */
        public Builder sourceRegionId(String sourceRegionId) {
            return sourceRegionId(Output.of(sourceRegionId));
        }

        /**
         * @param tags The tag value of an image. The value of N ranges from 1 to 20.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tag value of an image. The value of N ranges from 1 to 20.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        public ImageCopyArgs build() {
            $.sourceImageId = Objects.requireNonNull($.sourceImageId, "expected parameter 'sourceImageId' to be non-null");
            $.sourceRegionId = Objects.requireNonNull($.sourceRegionId, "expected parameter 'sourceRegionId' to be non-null");
            return $;
        }
    }

}