// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetEcsSnapshotsSnapshot {
    /**
     * @return The category of the snapshot.
     * 
     */
    private String category;
    private String creationTime;
    /**
     * @return The description of the snapshot.
     * 
     */
    private String description;
    /**
     * @return The source disk id.
     * 
     */
    private String diskId;
    /**
     * @return Whether the snapshot is encrypted.
     * 
     */
    private Boolean encrypted;
    /**
     * @return The ID of the Snapshot.
     * 
     */
    private String id;
    /**
     * @return Whether snapshot speed availability is enabled.
     * 
     */
    private Boolean instantAccess;
    /**
     * @return Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
     * 
     */
    private Integer instantAccessRetentionDays;
    private String name;
    /**
     * @return The product number inherited from the mirror market.
     * 
     */
    private String productCode;
    /**
     * @return Snapshot creation progress, in percentage.
     * 
     */
    private String progress;
    /**
     * @return Remaining completion time for the snapshot being created.
     * 
     */
    private Integer remainTime;
    /**
     * @return The resource group id.
     * 
     */
    private String resourceGroupId;
    /**
     * @return Automatic snapshot retention days.
     * 
     */
    private Integer retentionDays;
    /**
     * @return The snapshot id.
     * 
     */
    private String snapshotId;
    /**
     * @return Snapshot Display Name.
     * 
     */
    private String snapshotName;
    /**
     * @return The serial number of the snapshot.
     * 
     */
    private String snapshotSn;
    /**
     * @return Snapshot creation type.
     * 
     */
    private String snapshotType;
    private String sourceDiskId;
    /**
     * @return Source disk capacity.
     * 
     */
    private String sourceDiskSize;
    /**
     * @return Source disk attributes.
     * 
     */
    private String sourceDiskType;
    /**
     * @return Original disk type.
     * 
     */
    private String sourceStorageType;
    /**
     * @return The status of the snapshot.
     * 
     */
    private String status;
    /**
     * @return The tags.
     * 
     */
    private Map<String,Object> tags;
    private String type;
    /**
     * @return A resource type that has a reference relationship.
     * 
     */
    private String usage;

    private GetEcsSnapshotsSnapshot() {}
    /**
     * @return The category of the snapshot.
     * 
     */
    public String category() {
        return this.category;
    }
    public String creationTime() {
        return this.creationTime;
    }
    /**
     * @return The description of the snapshot.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The source disk id.
     * 
     */
    public String diskId() {
        return this.diskId;
    }
    /**
     * @return Whether the snapshot is encrypted.
     * 
     */
    public Boolean encrypted() {
        return this.encrypted;
    }
    /**
     * @return The ID of the Snapshot.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Whether snapshot speed availability is enabled.
     * 
     */
    public Boolean instantAccess() {
        return this.instantAccess;
    }
    /**
     * @return Specifies the retention period of the instant access feature. After the retention period ends, the snapshot is automatically released.
     * 
     */
    public Integer instantAccessRetentionDays() {
        return this.instantAccessRetentionDays;
    }
    public String name() {
        return this.name;
    }
    /**
     * @return The product number inherited from the mirror market.
     * 
     */
    public String productCode() {
        return this.productCode;
    }
    /**
     * @return Snapshot creation progress, in percentage.
     * 
     */
    public String progress() {
        return this.progress;
    }
    /**
     * @return Remaining completion time for the snapshot being created.
     * 
     */
    public Integer remainTime() {
        return this.remainTime;
    }
    /**
     * @return The resource group id.
     * 
     */
    public String resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * @return Automatic snapshot retention days.
     * 
     */
    public Integer retentionDays() {
        return this.retentionDays;
    }
    /**
     * @return The snapshot id.
     * 
     */
    public String snapshotId() {
        return this.snapshotId;
    }
    /**
     * @return Snapshot Display Name.
     * 
     */
    public String snapshotName() {
        return this.snapshotName;
    }
    /**
     * @return The serial number of the snapshot.
     * 
     */
    public String snapshotSn() {
        return this.snapshotSn;
    }
    /**
     * @return Snapshot creation type.
     * 
     */
    public String snapshotType() {
        return this.snapshotType;
    }
    public String sourceDiskId() {
        return this.sourceDiskId;
    }
    /**
     * @return Source disk capacity.
     * 
     */
    public String sourceDiskSize() {
        return this.sourceDiskSize;
    }
    /**
     * @return Source disk attributes.
     * 
     */
    public String sourceDiskType() {
        return this.sourceDiskType;
    }
    /**
     * @return Original disk type.
     * 
     */
    public String sourceStorageType() {
        return this.sourceStorageType;
    }
    /**
     * @return The status of the snapshot.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The tags.
     * 
     */
    public Map<String,Object> tags() {
        return this.tags;
    }
    public String type() {
        return this.type;
    }
    /**
     * @return A resource type that has a reference relationship.
     * 
     */
    public String usage() {
        return this.usage;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEcsSnapshotsSnapshot defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String category;
        private String creationTime;
        private String description;
        private String diskId;
        private Boolean encrypted;
        private String id;
        private Boolean instantAccess;
        private Integer instantAccessRetentionDays;
        private String name;
        private String productCode;
        private String progress;
        private Integer remainTime;
        private String resourceGroupId;
        private Integer retentionDays;
        private String snapshotId;
        private String snapshotName;
        private String snapshotSn;
        private String snapshotType;
        private String sourceDiskId;
        private String sourceDiskSize;
        private String sourceDiskType;
        private String sourceStorageType;
        private String status;
        private Map<String,Object> tags;
        private String type;
        private String usage;
        public Builder() {}
        public Builder(GetEcsSnapshotsSnapshot defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.category = defaults.category;
    	      this.creationTime = defaults.creationTime;
    	      this.description = defaults.description;
    	      this.diskId = defaults.diskId;
    	      this.encrypted = defaults.encrypted;
    	      this.id = defaults.id;
    	      this.instantAccess = defaults.instantAccess;
    	      this.instantAccessRetentionDays = defaults.instantAccessRetentionDays;
    	      this.name = defaults.name;
    	      this.productCode = defaults.productCode;
    	      this.progress = defaults.progress;
    	      this.remainTime = defaults.remainTime;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.retentionDays = defaults.retentionDays;
    	      this.snapshotId = defaults.snapshotId;
    	      this.snapshotName = defaults.snapshotName;
    	      this.snapshotSn = defaults.snapshotSn;
    	      this.snapshotType = defaults.snapshotType;
    	      this.sourceDiskId = defaults.sourceDiskId;
    	      this.sourceDiskSize = defaults.sourceDiskSize;
    	      this.sourceDiskType = defaults.sourceDiskType;
    	      this.sourceStorageType = defaults.sourceStorageType;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.type = defaults.type;
    	      this.usage = defaults.usage;
        }

        @CustomType.Setter
        public Builder category(String category) {
            this.category = Objects.requireNonNull(category);
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
        public Builder diskId(String diskId) {
            this.diskId = Objects.requireNonNull(diskId);
            return this;
        }
        @CustomType.Setter
        public Builder encrypted(Boolean encrypted) {
            this.encrypted = Objects.requireNonNull(encrypted);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instantAccess(Boolean instantAccess) {
            this.instantAccess = Objects.requireNonNull(instantAccess);
            return this;
        }
        @CustomType.Setter
        public Builder instantAccessRetentionDays(Integer instantAccessRetentionDays) {
            this.instantAccessRetentionDays = Objects.requireNonNull(instantAccessRetentionDays);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
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
        public Builder remainTime(Integer remainTime) {
            this.remainTime = Objects.requireNonNull(remainTime);
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroupId(String resourceGroupId) {
            this.resourceGroupId = Objects.requireNonNull(resourceGroupId);
            return this;
        }
        @CustomType.Setter
        public Builder retentionDays(Integer retentionDays) {
            this.retentionDays = Objects.requireNonNull(retentionDays);
            return this;
        }
        @CustomType.Setter
        public Builder snapshotId(String snapshotId) {
            this.snapshotId = Objects.requireNonNull(snapshotId);
            return this;
        }
        @CustomType.Setter
        public Builder snapshotName(String snapshotName) {
            this.snapshotName = Objects.requireNonNull(snapshotName);
            return this;
        }
        @CustomType.Setter
        public Builder snapshotSn(String snapshotSn) {
            this.snapshotSn = Objects.requireNonNull(snapshotSn);
            return this;
        }
        @CustomType.Setter
        public Builder snapshotType(String snapshotType) {
            this.snapshotType = Objects.requireNonNull(snapshotType);
            return this;
        }
        @CustomType.Setter
        public Builder sourceDiskId(String sourceDiskId) {
            this.sourceDiskId = Objects.requireNonNull(sourceDiskId);
            return this;
        }
        @CustomType.Setter
        public Builder sourceDiskSize(String sourceDiskSize) {
            this.sourceDiskSize = Objects.requireNonNull(sourceDiskSize);
            return this;
        }
        @CustomType.Setter
        public Builder sourceDiskType(String sourceDiskType) {
            this.sourceDiskType = Objects.requireNonNull(sourceDiskType);
            return this;
        }
        @CustomType.Setter
        public Builder sourceStorageType(String sourceStorageType) {
            this.sourceStorageType = Objects.requireNonNull(sourceStorageType);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,Object> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        @CustomType.Setter
        public Builder usage(String usage) {
            this.usage = Objects.requireNonNull(usage);
            return this;
        }
        public GetEcsSnapshotsSnapshot build() {
            final var o = new GetEcsSnapshotsSnapshot();
            o.category = category;
            o.creationTime = creationTime;
            o.description = description;
            o.diskId = diskId;
            o.encrypted = encrypted;
            o.id = id;
            o.instantAccess = instantAccess;
            o.instantAccessRetentionDays = instantAccessRetentionDays;
            o.name = name;
            o.productCode = productCode;
            o.progress = progress;
            o.remainTime = remainTime;
            o.resourceGroupId = resourceGroupId;
            o.retentionDays = retentionDays;
            o.snapshotId = snapshotId;
            o.snapshotName = snapshotName;
            o.snapshotSn = snapshotSn;
            o.snapshotType = snapshotType;
            o.sourceDiskId = sourceDiskId;
            o.sourceDiskSize = sourceDiskSize;
            o.sourceDiskType = sourceDiskType;
            o.sourceStorageType = sourceStorageType;
            o.status = status;
            o.tags = tags;
            o.type = type;
            o.usage = usage;
            return o;
        }
    }
}