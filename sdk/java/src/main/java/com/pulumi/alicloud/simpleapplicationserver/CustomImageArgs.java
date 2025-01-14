// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.simpleapplicationserver;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CustomImageArgs extends com.pulumi.resources.ResourceArgs {

    public static final CustomImageArgs Empty = new CustomImageArgs();

    /**
     * The name of the resource. The name must be `2` to `128` characters in length. It must start with a letter or a number. It can contain letters, digits, colons (:), underscores (_) and hyphens (-).
     * 
     */
    @Import(name="customImageName", required=true)
    private Output<String> customImageName;

    /**
     * @return The name of the resource. The name must be `2` to `128` characters in length. It must start with a letter or a number. It can contain letters, digits, colons (:), underscores (_) and hyphens (-).
     * 
     */
    public Output<String> customImageName() {
        return this.customImageName;
    }

    /**
     * The description of the Custom Image.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the Custom Image.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The ID of the instance.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return The ID of the instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     * The Shared status of the Custom Image. Valid values: `Share`, `UnShare`.
     * 
     * **NOTE:** The `status` will be automatically change to `UnShare` when the resource is deleted, please operate with caution.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The Shared status of the Custom Image. Valid values: `Share`, `UnShare`.
     * 
     * **NOTE:** The `status` will be automatically change to `UnShare` when the resource is deleted, please operate with caution.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The ID of the system snapshot.
     * 
     */
    @Import(name="systemSnapshotId", required=true)
    private Output<String> systemSnapshotId;

    /**
     * @return The ID of the system snapshot.
     * 
     */
    public Output<String> systemSnapshotId() {
        return this.systemSnapshotId;
    }

    private CustomImageArgs() {}

    private CustomImageArgs(CustomImageArgs $) {
        this.customImageName = $.customImageName;
        this.description = $.description;
        this.instanceId = $.instanceId;
        this.status = $.status;
        this.systemSnapshotId = $.systemSnapshotId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CustomImageArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CustomImageArgs $;

        public Builder() {
            $ = new CustomImageArgs();
        }

        public Builder(CustomImageArgs defaults) {
            $ = new CustomImageArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param customImageName The name of the resource. The name must be `2` to `128` characters in length. It must start with a letter or a number. It can contain letters, digits, colons (:), underscores (_) and hyphens (-).
         * 
         * @return builder
         * 
         */
        public Builder customImageName(Output<String> customImageName) {
            $.customImageName = customImageName;
            return this;
        }

        /**
         * @param customImageName The name of the resource. The name must be `2` to `128` characters in length. It must start with a letter or a number. It can contain letters, digits, colons (:), underscores (_) and hyphens (-).
         * 
         * @return builder
         * 
         */
        public Builder customImageName(String customImageName) {
            return customImageName(Output.of(customImageName));
        }

        /**
         * @param description The description of the Custom Image.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the Custom Image.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param instanceId The ID of the instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The ID of the instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param status The Shared status of the Custom Image. Valid values: `Share`, `UnShare`.
         * 
         * **NOTE:** The `status` will be automatically change to `UnShare` when the resource is deleted, please operate with caution.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The Shared status of the Custom Image. Valid values: `Share`, `UnShare`.
         * 
         * **NOTE:** The `status` will be automatically change to `UnShare` when the resource is deleted, please operate with caution.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param systemSnapshotId The ID of the system snapshot.
         * 
         * @return builder
         * 
         */
        public Builder systemSnapshotId(Output<String> systemSnapshotId) {
            $.systemSnapshotId = systemSnapshotId;
            return this;
        }

        /**
         * @param systemSnapshotId The ID of the system snapshot.
         * 
         * @return builder
         * 
         */
        public Builder systemSnapshotId(String systemSnapshotId) {
            return systemSnapshotId(Output.of(systemSnapshotId));
        }

        public CustomImageArgs build() {
            $.customImageName = Objects.requireNonNull($.customImageName, "expected parameter 'customImageName' to be non-null");
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            $.systemSnapshotId = Objects.requireNonNull($.systemSnapshotId, "expected parameter 'systemSnapshotId' to be non-null");
            return $;
        }
    }

}
