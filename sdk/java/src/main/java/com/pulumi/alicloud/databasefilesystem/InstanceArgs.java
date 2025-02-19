// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.databasefilesystem;

import com.pulumi.alicloud.databasefilesystem.inputs.InstanceEcsListArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceArgs Empty = new InstanceArgs();

    /**
     * The type of the Database file system. Valid values: `standard`.
     * 
     */
    @Import(name="category")
    private @Nullable Output<String> category;

    /**
     * @return The type of the Database file system. Valid values: `standard`.
     * 
     */
    public Optional<Output<String>> category() {
        return Optional.ofNullable(this.category);
    }

    /**
     * Whether to delete the original snapshot after the DBFS is created using the snapshot. Valid values : `true` anf `false`.
     * 
     */
    @Import(name="deleteSnapshot")
    private @Nullable Output<Boolean> deleteSnapshot;

    /**
     * @return Whether to delete the original snapshot after the DBFS is created using the snapshot. Valid values : `true` anf `false`.
     * 
     */
    public Optional<Output<Boolean>> deleteSnapshot() {
        return Optional.ofNullable(this.deleteSnapshot);
    }

    /**
     * The collection of ECS instances mounted to the Database file system. See the following `Block ecs_list`. **NOTE:** Field &#39;ecs_list&#39; has been deprecated from provider version 1.156.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_dbfs_instance_attachment&#39; to attach ECS and DBFS.
     * 
     * @deprecated
     * Field &#39;ecs_list&#39; has been deprecated from provider version 1.156.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_dbfs_instance_attachment&#39; to attach ECS and DBFS.
     * 
     */
    @Deprecated /* Field 'ecs_list' has been deprecated from provider version 1.156.0 and it will be removed in the future version. Please use the new resource 'alicloud_dbfs_instance_attachment' to attach ECS and DBFS. */
    @Import(name="ecsLists")
    private @Nullable Output<List<InstanceEcsListArgs>> ecsLists;

    /**
     * @return The collection of ECS instances mounted to the Database file system. See the following `Block ecs_list`. **NOTE:** Field &#39;ecs_list&#39; has been deprecated from provider version 1.156.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_dbfs_instance_attachment&#39; to attach ECS and DBFS.
     * 
     * @deprecated
     * Field &#39;ecs_list&#39; has been deprecated from provider version 1.156.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_dbfs_instance_attachment&#39; to attach ECS and DBFS.
     * 
     */
    @Deprecated /* Field 'ecs_list' has been deprecated from provider version 1.156.0 and it will be removed in the future version. Please use the new resource 'alicloud_dbfs_instance_attachment' to attach ECS and DBFS. */
    public Optional<Output<List<InstanceEcsListArgs>>> ecsLists() {
        return Optional.ofNullable(this.ecsLists);
    }

    /**
     * Whether to create the Database file system in RAID way. Valid values : `true` anf `false`.
     * 
     */
    @Import(name="enableRaid")
    private @Nullable Output<Boolean> enableRaid;

    /**
     * @return Whether to create the Database file system in RAID way. Valid values : `true` anf `false`.
     * 
     */
    public Optional<Output<Boolean>> enableRaid() {
        return Optional.ofNullable(this.enableRaid);
    }

    /**
     * Whether to encrypt the database file system. Valid values: `true` and `false`.
     * 
     */
    @Import(name="encryption")
    private @Nullable Output<Boolean> encryption;

    /**
     * @return Whether to encrypt the database file system. Valid values: `true` and `false`.
     * 
     */
    public Optional<Output<Boolean>> encryption() {
        return Optional.ofNullable(this.encryption);
    }

    /**
     * The name of the Database file system.
     * 
     */
    @Import(name="instanceName", required=true)
    private Output<String> instanceName;

    /**
     * @return The name of the Database file system.
     * 
     */
    public Output<String> instanceName() {
        return this.instanceName;
    }

    /**
     * The KMS key ID of the Database file system used. This parameter is valid When `encryption` parameter is set to `true`.
     * 
     */
    @Import(name="kmsKeyId")
    private @Nullable Output<String> kmsKeyId;

    /**
     * @return The KMS key ID of the Database file system used. This parameter is valid When `encryption` parameter is set to `true`.
     * 
     */
    public Optional<Output<String>> kmsKeyId() {
        return Optional.ofNullable(this.kmsKeyId);
    }

    /**
     * The performance level of the Database file system. Valid values: `PL0`, `PL1`, `PL2`, `PL3`.
     * 
     */
    @Import(name="performanceLevel")
    private @Nullable Output<String> performanceLevel;

    /**
     * @return The performance level of the Database file system. Valid values: `PL0`, `PL1`, `PL2`, `PL3`.
     * 
     */
    public Optional<Output<String>> performanceLevel() {
        return Optional.ofNullable(this.performanceLevel);
    }

    /**
     * The number of strip. This parameter is valid When `enable_raid` parameter is set to `true`.
     * 
     */
    @Import(name="raidStripeUnitNumber")
    private @Nullable Output<String> raidStripeUnitNumber;

    /**
     * @return The number of strip. This parameter is valid When `enable_raid` parameter is set to `true`.
     * 
     */
    public Optional<Output<String>> raidStripeUnitNumber() {
        return Optional.ofNullable(this.raidStripeUnitNumber);
    }

    /**
     * The size Of the Database file system. Unit: GiB.
     * 
     */
    @Import(name="size", required=true)
    private Output<Integer> size;

    /**
     * @return The size Of the Database file system. Unit: GiB.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }

    /**
     * The snapshot id of the Database file system.
     * 
     */
    @Import(name="snapshotId")
    private @Nullable Output<String> snapshotId;

    /**
     * @return The snapshot id of the Database file system.
     * 
     */
    public Optional<Output<String>> snapshotId() {
        return Optional.ofNullable(this.snapshotId);
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

    /**
     * The Zone ID of the Database file system.
     * 
     */
    @Import(name="zoneId", required=true)
    private Output<String> zoneId;

    /**
     * @return The Zone ID of the Database file system.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    private InstanceArgs() {}

    private InstanceArgs(InstanceArgs $) {
        this.category = $.category;
        this.deleteSnapshot = $.deleteSnapshot;
        this.ecsLists = $.ecsLists;
        this.enableRaid = $.enableRaid;
        this.encryption = $.encryption;
        this.instanceName = $.instanceName;
        this.kmsKeyId = $.kmsKeyId;
        this.performanceLevel = $.performanceLevel;
        this.raidStripeUnitNumber = $.raidStripeUnitNumber;
        this.size = $.size;
        this.snapshotId = $.snapshotId;
        this.tags = $.tags;
        this.zoneId = $.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceArgs $;

        public Builder() {
            $ = new InstanceArgs();
        }

        public Builder(InstanceArgs defaults) {
            $ = new InstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param category The type of the Database file system. Valid values: `standard`.
         * 
         * @return builder
         * 
         */
        public Builder category(@Nullable Output<String> category) {
            $.category = category;
            return this;
        }

        /**
         * @param category The type of the Database file system. Valid values: `standard`.
         * 
         * @return builder
         * 
         */
        public Builder category(String category) {
            return category(Output.of(category));
        }

        /**
         * @param deleteSnapshot Whether to delete the original snapshot after the DBFS is created using the snapshot. Valid values : `true` anf `false`.
         * 
         * @return builder
         * 
         */
        public Builder deleteSnapshot(@Nullable Output<Boolean> deleteSnapshot) {
            $.deleteSnapshot = deleteSnapshot;
            return this;
        }

        /**
         * @param deleteSnapshot Whether to delete the original snapshot after the DBFS is created using the snapshot. Valid values : `true` anf `false`.
         * 
         * @return builder
         * 
         */
        public Builder deleteSnapshot(Boolean deleteSnapshot) {
            return deleteSnapshot(Output.of(deleteSnapshot));
        }

        /**
         * @param ecsLists The collection of ECS instances mounted to the Database file system. See the following `Block ecs_list`. **NOTE:** Field &#39;ecs_list&#39; has been deprecated from provider version 1.156.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_dbfs_instance_attachment&#39; to attach ECS and DBFS.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;ecs_list&#39; has been deprecated from provider version 1.156.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_dbfs_instance_attachment&#39; to attach ECS and DBFS.
         * 
         */
        @Deprecated /* Field 'ecs_list' has been deprecated from provider version 1.156.0 and it will be removed in the future version. Please use the new resource 'alicloud_dbfs_instance_attachment' to attach ECS and DBFS. */
        public Builder ecsLists(@Nullable Output<List<InstanceEcsListArgs>> ecsLists) {
            $.ecsLists = ecsLists;
            return this;
        }

        /**
         * @param ecsLists The collection of ECS instances mounted to the Database file system. See the following `Block ecs_list`. **NOTE:** Field &#39;ecs_list&#39; has been deprecated from provider version 1.156.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_dbfs_instance_attachment&#39; to attach ECS and DBFS.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;ecs_list&#39; has been deprecated from provider version 1.156.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_dbfs_instance_attachment&#39; to attach ECS and DBFS.
         * 
         */
        @Deprecated /* Field 'ecs_list' has been deprecated from provider version 1.156.0 and it will be removed in the future version. Please use the new resource 'alicloud_dbfs_instance_attachment' to attach ECS and DBFS. */
        public Builder ecsLists(List<InstanceEcsListArgs> ecsLists) {
            return ecsLists(Output.of(ecsLists));
        }

        /**
         * @param ecsLists The collection of ECS instances mounted to the Database file system. See the following `Block ecs_list`. **NOTE:** Field &#39;ecs_list&#39; has been deprecated from provider version 1.156.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_dbfs_instance_attachment&#39; to attach ECS and DBFS.
         * 
         * @return builder
         * 
         * @deprecated
         * Field &#39;ecs_list&#39; has been deprecated from provider version 1.156.0 and it will be removed in the future version. Please use the new resource &#39;alicloud_dbfs_instance_attachment&#39; to attach ECS and DBFS.
         * 
         */
        @Deprecated /* Field 'ecs_list' has been deprecated from provider version 1.156.0 and it will be removed in the future version. Please use the new resource 'alicloud_dbfs_instance_attachment' to attach ECS and DBFS. */
        public Builder ecsLists(InstanceEcsListArgs... ecsLists) {
            return ecsLists(List.of(ecsLists));
        }

        /**
         * @param enableRaid Whether to create the Database file system in RAID way. Valid values : `true` anf `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableRaid(@Nullable Output<Boolean> enableRaid) {
            $.enableRaid = enableRaid;
            return this;
        }

        /**
         * @param enableRaid Whether to create the Database file system in RAID way. Valid values : `true` anf `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableRaid(Boolean enableRaid) {
            return enableRaid(Output.of(enableRaid));
        }

        /**
         * @param encryption Whether to encrypt the database file system. Valid values: `true` and `false`.
         * 
         * @return builder
         * 
         */
        public Builder encryption(@Nullable Output<Boolean> encryption) {
            $.encryption = encryption;
            return this;
        }

        /**
         * @param encryption Whether to encrypt the database file system. Valid values: `true` and `false`.
         * 
         * @return builder
         * 
         */
        public Builder encryption(Boolean encryption) {
            return encryption(Output.of(encryption));
        }

        /**
         * @param instanceName The name of the Database file system.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(Output<String> instanceName) {
            $.instanceName = instanceName;
            return this;
        }

        /**
         * @param instanceName The name of the Database file system.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(String instanceName) {
            return instanceName(Output.of(instanceName));
        }

        /**
         * @param kmsKeyId The KMS key ID of the Database file system used. This parameter is valid When `encryption` parameter is set to `true`.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(@Nullable Output<String> kmsKeyId) {
            $.kmsKeyId = kmsKeyId;
            return this;
        }

        /**
         * @param kmsKeyId The KMS key ID of the Database file system used. This parameter is valid When `encryption` parameter is set to `true`.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(String kmsKeyId) {
            return kmsKeyId(Output.of(kmsKeyId));
        }

        /**
         * @param performanceLevel The performance level of the Database file system. Valid values: `PL0`, `PL1`, `PL2`, `PL3`.
         * 
         * @return builder
         * 
         */
        public Builder performanceLevel(@Nullable Output<String> performanceLevel) {
            $.performanceLevel = performanceLevel;
            return this;
        }

        /**
         * @param performanceLevel The performance level of the Database file system. Valid values: `PL0`, `PL1`, `PL2`, `PL3`.
         * 
         * @return builder
         * 
         */
        public Builder performanceLevel(String performanceLevel) {
            return performanceLevel(Output.of(performanceLevel));
        }

        /**
         * @param raidStripeUnitNumber The number of strip. This parameter is valid When `enable_raid` parameter is set to `true`.
         * 
         * @return builder
         * 
         */
        public Builder raidStripeUnitNumber(@Nullable Output<String> raidStripeUnitNumber) {
            $.raidStripeUnitNumber = raidStripeUnitNumber;
            return this;
        }

        /**
         * @param raidStripeUnitNumber The number of strip. This parameter is valid When `enable_raid` parameter is set to `true`.
         * 
         * @return builder
         * 
         */
        public Builder raidStripeUnitNumber(String raidStripeUnitNumber) {
            return raidStripeUnitNumber(Output.of(raidStripeUnitNumber));
        }

        /**
         * @param size The size Of the Database file system. Unit: GiB.
         * 
         * @return builder
         * 
         */
        public Builder size(Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size The size Of the Database file system. Unit: GiB.
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param snapshotId The snapshot id of the Database file system.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(@Nullable Output<String> snapshotId) {
            $.snapshotId = snapshotId;
            return this;
        }

        /**
         * @param snapshotId The snapshot id of the Database file system.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(String snapshotId) {
            return snapshotId(Output.of(snapshotId));
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

        /**
         * @param zoneId The Zone ID of the Database file system.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(Output<String> zoneId) {
            $.zoneId = zoneId;
            return this;
        }

        /**
         * @param zoneId The Zone ID of the Database file system.
         * 
         * @return builder
         * 
         */
        public Builder zoneId(String zoneId) {
            return zoneId(Output.of(zoneId));
        }

        public InstanceArgs build() {
            $.instanceName = Objects.requireNonNull($.instanceName, "expected parameter 'instanceName' to be non-null");
            $.size = Objects.requireNonNull($.size, "expected parameter 'size' to be non-null");
            $.zoneId = Objects.requireNonNull($.zoneId, "expected parameter 'zoneId' to be non-null");
            return $;
        }
    }

}
