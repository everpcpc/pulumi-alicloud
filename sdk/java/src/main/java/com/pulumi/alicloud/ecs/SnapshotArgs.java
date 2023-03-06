// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SnapshotArgs extends com.pulumi.resources.ResourceArgs {

    public static final SnapshotArgs Empty = new SnapshotArgs();

    @Import(name="category")
    private @Nullable Output<String> category;

    public Optional<Output<String>> category() {
        return Optional.ofNullable(this.category);
    }

    /**
     * Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The source disk ID.
     * 
     */
    @Import(name="diskId", required=true)
    private Output<String> diskId;

    /**
     * @return The source disk ID.
     * 
     */
    public Output<String> diskId() {
        return this.diskId;
    }

    @Import(name="force")
    private @Nullable Output<Boolean> force;

    public Optional<Output<Boolean>> force() {
        return Optional.ofNullable(this.force);
    }

    @Import(name="instantAccess")
    private @Nullable Output<Boolean> instantAccess;

    public Optional<Output<Boolean>> instantAccess() {
        return Optional.ofNullable(this.instantAccess);
    }

    @Import(name="instantAccessRetentionDays")
    private @Nullable Output<Integer> instantAccessRetentionDays;

    public Optional<Output<Integer>> instantAccessRetentionDays() {
        return Optional.ofNullable(this.instantAccessRetentionDays);
    }

    /**
     * The name of the snapshot to be created. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
     * It cannot start with auto, because snapshot names starting with auto are recognized as automatic snapshots.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.120.0. New field &#39;snapshot_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead. */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the snapshot to be created. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
     * It cannot start with auto, because snapshot names starting with auto are recognized as automatic snapshots.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.120.0. New field &#39;snapshot_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead. */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the resource group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    @Import(name="retentionDays")
    private @Nullable Output<Integer> retentionDays;

    public Optional<Output<Integer>> retentionDays() {
        return Optional.ofNullable(this.retentionDays);
    }

    @Import(name="snapshotName")
    private @Nullable Output<String> snapshotName;

    public Optional<Output<String>> snapshotName() {
        return Optional.ofNullable(this.snapshotName);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private SnapshotArgs() {}

    private SnapshotArgs(SnapshotArgs $) {
        this.category = $.category;
        this.description = $.description;
        this.diskId = $.diskId;
        this.force = $.force;
        this.instantAccess = $.instantAccess;
        this.instantAccessRetentionDays = $.instantAccessRetentionDays;
        this.name = $.name;
        this.resourceGroupId = $.resourceGroupId;
        this.retentionDays = $.retentionDays;
        this.snapshotName = $.snapshotName;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SnapshotArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SnapshotArgs $;

        public Builder() {
            $ = new SnapshotArgs();
        }

        public Builder(SnapshotArgs defaults) {
            $ = new SnapshotArgs(Objects.requireNonNull(defaults));
        }

        public Builder category(@Nullable Output<String> category) {
            $.category = category;
            return this;
        }

        public Builder category(String category) {
            return category(Output.of(category));
        }

        /**
         * @param description Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the snapshot. This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param diskId The source disk ID.
         * 
         * @return builder
         * 
         */
        public Builder diskId(Output<String> diskId) {
            $.diskId = diskId;
            return this;
        }

        /**
         * @param diskId The source disk ID.
         * 
         * @return builder
         * 
         */
        public Builder diskId(String diskId) {
            return diskId(Output.of(diskId));
        }

        public Builder force(@Nullable Output<Boolean> force) {
            $.force = force;
            return this;
        }

        public Builder force(Boolean force) {
            return force(Output.of(force));
        }

        public Builder instantAccess(@Nullable Output<Boolean> instantAccess) {
            $.instantAccess = instantAccess;
            return this;
        }

        public Builder instantAccess(Boolean instantAccess) {
            return instantAccess(Output.of(instantAccess));
        }

        public Builder instantAccessRetentionDays(@Nullable Output<Integer> instantAccessRetentionDays) {
            $.instantAccessRetentionDays = instantAccessRetentionDays;
            return this;
        }

        public Builder instantAccessRetentionDays(Integer instantAccessRetentionDays) {
            return instantAccessRetentionDays(Output.of(instantAccessRetentionDays));
        }

        /**
         * @param name The name of the snapshot to be created. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
         * It cannot start with auto, because snapshot names starting with auto are recognized as automatic snapshots.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from provider version 1.120.0. New field &#39;snapshot_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead. */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the snapshot to be created. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
         * It cannot start with auto, because snapshot names starting with auto are recognized as automatic snapshots.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;name&#39; has been deprecated from provider version 1.120.0. New field &#39;snapshot_name&#39; instead.
         * 
         */
        @Deprecated /* Field 'name' has been deprecated from provider version 1.120.0. New field 'snapshot_name' instead. */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        public Builder retentionDays(@Nullable Output<Integer> retentionDays) {
            $.retentionDays = retentionDays;
            return this;
        }

        public Builder retentionDays(Integer retentionDays) {
            return retentionDays(Output.of(retentionDays));
        }

        public Builder snapshotName(@Nullable Output<String> snapshotName) {
            $.snapshotName = snapshotName;
            return this;
        }

        public Builder snapshotName(String snapshotName) {
            return snapshotName(Output.of(snapshotName));
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        public SnapshotArgs build() {
            $.diskId = Objects.requireNonNull($.diskId, "expected parameter 'diskId' to be non-null");
            return $;
        }
    }

}