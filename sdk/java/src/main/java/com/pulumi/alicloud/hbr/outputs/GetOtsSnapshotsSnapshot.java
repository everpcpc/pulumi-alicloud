// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetOtsSnapshotsSnapshot {
    /**
     * @return The actual amount of backup snapshots after duplicates are removed. Unit: bytes.
     * 
     */
    private String actualBytes;
    /**
     * @return The backup type. Valid value: `COMPLETE`, which indicates full backup.
     * 
     */
    private String backupType;
    /**
     * @return The total amount of data. Unit: bytes.
     * 
     */
    private String bytesTotal;
    /**
     * @return The time when the backup snapshot was completed. This value is a UNIX timestamp. Unit: seconds.
     * 
     */
    private String completeTime;
    /**
     * @return The time when the Table store instance was created. This value is a UNIX timestamp. Unit: seconds.
     * 
     */
    private String createTime;
    /**
     * @return The time when the backup snapshot was created. This value is a UNIX timestamp. Unit: seconds.
     * 
     */
    private String createdTime;
    /**
     * @return The ID of the backup snapshot.
     * 
     */
    private String id;
    /**
     * @return The name of the Table store instance.
     * 
     */
    private String instanceName;
    /**
     * @return The ID of the backup job.
     * 
     */
    private String jobId;
    /**
     * @return The hash value of the parent backup snapshot.
     * 
     */
    private String parentSnapshotHash;
    /**
     * @return The time when the backup job ended. This value is a UNIX timestamp. Unit: milliseconds.
     * 
     */
    private String rangeEnd;
    /**
     * @return The time when the backup job started. This value is a UNIX timestamp. Unit: milliseconds.
     * 
     */
    private String rangeStart;
    /**
     * @return The retention period of the backup snapshot.
     * 
     */
    private String retention;
    /**
     * @return The hash value of the backup snapshot.
     * 
     */
    private String snapshotHash;
    /**
     * @return The ID of the backup snapshot.
     * 
     */
    private String snapshotId;
    /**
     * @return The type of the data source. Valid values: `ECS_FILE`,`PARTIAL_COMPLETE`,`FAILED`
     * 
     */
    private String sourceType;
    /**
     * @return The start time of the backup snapshot. This value is a UNIX timestamp. Unit: seconds.
     * 
     */
    private String startTime;
    /**
     * @return The status of the backup job. Valid values: `COMPLETE`,`PARTIAL_COMPLETE`,`FAILED`.
     * 
     */
    private String status;
    /**
     * @return The name of the table in the Table store instance.
     * 
     */
    private String tableName;
    /**
     * @return The time when the backup snapshot was updated. This value is a UNIX timestamp. Unit: seconds.
     * 
     */
    private String updatedTime;
    /**
     * @return The ID of the backup vault that stores the backup snapshot.
     * 
     */
    private String vaultId;

    private GetOtsSnapshotsSnapshot() {}
    /**
     * @return The actual amount of backup snapshots after duplicates are removed. Unit: bytes.
     * 
     */
    public String actualBytes() {
        return this.actualBytes;
    }
    /**
     * @return The backup type. Valid value: `COMPLETE`, which indicates full backup.
     * 
     */
    public String backupType() {
        return this.backupType;
    }
    /**
     * @return The total amount of data. Unit: bytes.
     * 
     */
    public String bytesTotal() {
        return this.bytesTotal;
    }
    /**
     * @return The time when the backup snapshot was completed. This value is a UNIX timestamp. Unit: seconds.
     * 
     */
    public String completeTime() {
        return this.completeTime;
    }
    /**
     * @return The time when the Table store instance was created. This value is a UNIX timestamp. Unit: seconds.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The time when the backup snapshot was created. This value is a UNIX timestamp. Unit: seconds.
     * 
     */
    public String createdTime() {
        return this.createdTime;
    }
    /**
     * @return The ID of the backup snapshot.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the Table store instance.
     * 
     */
    public String instanceName() {
        return this.instanceName;
    }
    /**
     * @return The ID of the backup job.
     * 
     */
    public String jobId() {
        return this.jobId;
    }
    /**
     * @return The hash value of the parent backup snapshot.
     * 
     */
    public String parentSnapshotHash() {
        return this.parentSnapshotHash;
    }
    /**
     * @return The time when the backup job ended. This value is a UNIX timestamp. Unit: milliseconds.
     * 
     */
    public String rangeEnd() {
        return this.rangeEnd;
    }
    /**
     * @return The time when the backup job started. This value is a UNIX timestamp. Unit: milliseconds.
     * 
     */
    public String rangeStart() {
        return this.rangeStart;
    }
    /**
     * @return The retention period of the backup snapshot.
     * 
     */
    public String retention() {
        return this.retention;
    }
    /**
     * @return The hash value of the backup snapshot.
     * 
     */
    public String snapshotHash() {
        return this.snapshotHash;
    }
    /**
     * @return The ID of the backup snapshot.
     * 
     */
    public String snapshotId() {
        return this.snapshotId;
    }
    /**
     * @return The type of the data source. Valid values: `ECS_FILE`,`PARTIAL_COMPLETE`,`FAILED`
     * 
     */
    public String sourceType() {
        return this.sourceType;
    }
    /**
     * @return The start time of the backup snapshot. This value is a UNIX timestamp. Unit: seconds.
     * 
     */
    public String startTime() {
        return this.startTime;
    }
    /**
     * @return The status of the backup job. Valid values: `COMPLETE`,`PARTIAL_COMPLETE`,`FAILED`.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The name of the table in the Table store instance.
     * 
     */
    public String tableName() {
        return this.tableName;
    }
    /**
     * @return The time when the backup snapshot was updated. This value is a UNIX timestamp. Unit: seconds.
     * 
     */
    public String updatedTime() {
        return this.updatedTime;
    }
    /**
     * @return The ID of the backup vault that stores the backup snapshot.
     * 
     */
    public String vaultId() {
        return this.vaultId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOtsSnapshotsSnapshot defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String actualBytes;
        private String backupType;
        private String bytesTotal;
        private String completeTime;
        private String createTime;
        private String createdTime;
        private String id;
        private String instanceName;
        private String jobId;
        private String parentSnapshotHash;
        private String rangeEnd;
        private String rangeStart;
        private String retention;
        private String snapshotHash;
        private String snapshotId;
        private String sourceType;
        private String startTime;
        private String status;
        private String tableName;
        private String updatedTime;
        private String vaultId;
        public Builder() {}
        public Builder(GetOtsSnapshotsSnapshot defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.actualBytes = defaults.actualBytes;
    	      this.backupType = defaults.backupType;
    	      this.bytesTotal = defaults.bytesTotal;
    	      this.completeTime = defaults.completeTime;
    	      this.createTime = defaults.createTime;
    	      this.createdTime = defaults.createdTime;
    	      this.id = defaults.id;
    	      this.instanceName = defaults.instanceName;
    	      this.jobId = defaults.jobId;
    	      this.parentSnapshotHash = defaults.parentSnapshotHash;
    	      this.rangeEnd = defaults.rangeEnd;
    	      this.rangeStart = defaults.rangeStart;
    	      this.retention = defaults.retention;
    	      this.snapshotHash = defaults.snapshotHash;
    	      this.snapshotId = defaults.snapshotId;
    	      this.sourceType = defaults.sourceType;
    	      this.startTime = defaults.startTime;
    	      this.status = defaults.status;
    	      this.tableName = defaults.tableName;
    	      this.updatedTime = defaults.updatedTime;
    	      this.vaultId = defaults.vaultId;
        }

        @CustomType.Setter
        public Builder actualBytes(String actualBytes) {
            this.actualBytes = Objects.requireNonNull(actualBytes);
            return this;
        }
        @CustomType.Setter
        public Builder backupType(String backupType) {
            this.backupType = Objects.requireNonNull(backupType);
            return this;
        }
        @CustomType.Setter
        public Builder bytesTotal(String bytesTotal) {
            this.bytesTotal = Objects.requireNonNull(bytesTotal);
            return this;
        }
        @CustomType.Setter
        public Builder completeTime(String completeTime) {
            this.completeTime = Objects.requireNonNull(completeTime);
            return this;
        }
        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder createdTime(String createdTime) {
            this.createdTime = Objects.requireNonNull(createdTime);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceName(String instanceName) {
            this.instanceName = Objects.requireNonNull(instanceName);
            return this;
        }
        @CustomType.Setter
        public Builder jobId(String jobId) {
            this.jobId = Objects.requireNonNull(jobId);
            return this;
        }
        @CustomType.Setter
        public Builder parentSnapshotHash(String parentSnapshotHash) {
            this.parentSnapshotHash = Objects.requireNonNull(parentSnapshotHash);
            return this;
        }
        @CustomType.Setter
        public Builder rangeEnd(String rangeEnd) {
            this.rangeEnd = Objects.requireNonNull(rangeEnd);
            return this;
        }
        @CustomType.Setter
        public Builder rangeStart(String rangeStart) {
            this.rangeStart = Objects.requireNonNull(rangeStart);
            return this;
        }
        @CustomType.Setter
        public Builder retention(String retention) {
            this.retention = Objects.requireNonNull(retention);
            return this;
        }
        @CustomType.Setter
        public Builder snapshotHash(String snapshotHash) {
            this.snapshotHash = Objects.requireNonNull(snapshotHash);
            return this;
        }
        @CustomType.Setter
        public Builder snapshotId(String snapshotId) {
            this.snapshotId = Objects.requireNonNull(snapshotId);
            return this;
        }
        @CustomType.Setter
        public Builder sourceType(String sourceType) {
            this.sourceType = Objects.requireNonNull(sourceType);
            return this;
        }
        @CustomType.Setter
        public Builder startTime(String startTime) {
            this.startTime = Objects.requireNonNull(startTime);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder tableName(String tableName) {
            this.tableName = Objects.requireNonNull(tableName);
            return this;
        }
        @CustomType.Setter
        public Builder updatedTime(String updatedTime) {
            this.updatedTime = Objects.requireNonNull(updatedTime);
            return this;
        }
        @CustomType.Setter
        public Builder vaultId(String vaultId) {
            this.vaultId = Objects.requireNonNull(vaultId);
            return this;
        }
        public GetOtsSnapshotsSnapshot build() {
            final var o = new GetOtsSnapshotsSnapshot();
            o.actualBytes = actualBytes;
            o.backupType = backupType;
            o.bytesTotal = bytesTotal;
            o.completeTime = completeTime;
            o.createTime = createTime;
            o.createdTime = createdTime;
            o.id = id;
            o.instanceName = instanceName;
            o.jobId = jobId;
            o.parentSnapshotHash = parentSnapshotHash;
            o.rangeEnd = rangeEnd;
            o.rangeStart = rangeStart;
            o.retention = retention;
            o.snapshotHash = snapshotHash;
            o.snapshotId = snapshotId;
            o.sourceType = sourceType;
            o.startTime = startTime;
            o.status = status;
            o.tableName = tableName;
            o.updatedTime = updatedTime;
            o.vaultId = vaultId;
            return o;
        }
    }
}
