// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr.outputs;

import com.pulumi.alicloud.hbr.outputs.GetBackupJobsJobOtsDetail;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetBackupJobsJob {
    /**
     * @return The actual data volume of the backup task (After deduplication) . Unit byte.
     * 
     */
    private String actualBytes;
    /**
     * @return The actual number of items in the backup task. (Currently only file backup is available).
     * 
     */
    private String actualItems;
    /**
     * @return The name of backup job.
     * 
     */
    private String backJobName;
    /**
     * @return The ID of the backup job.
     * 
     */
    private String backupJobId;
    /**
     * @return Backup type. Valid values: `COMPLETE`(full backup).
     * 
     */
    private String backupType;
    /**
     * @return The name of target OSS bucket.
     * 
     */
    private String bucket;
    /**
     * @return The amount of backup data (Incremental). Unit byte.
     * 
     */
    private String bytesDone;
    /**
     * @return The total amount of data sources. Unit byte.
     * 
     */
    private String bytesTotal;
    /**
     * @return The completion time of backup job. UNIX time seconds.
     * 
     */
    private String completeTime;
    /**
     * @return The creation time of backup job. UNIX time seconds.
     * 
     */
    private String createTime;
    /**
     * @return The role name created in the original account RAM backup by the cross account managed by the current account. It is valid only when `source_type` is `ECS_FILE`, `NAS`, `OSS` or `OTS`.
     * 
     */
    private String crossAccountRoleName;
    /**
     * @return The type of the cross account backup. It is valid only when `source_type` is `ECS_FILE`, `NAS`, `OSS` or `OTS`.
     * 
     */
    private String crossAccountType;
    /**
     * @return The original account ID of the cross account backup managed by the current account. It is valid only when `source_type` is `ECS_FILE`, `NAS`, `OSS` or `OTS`.
     * 
     */
    private Integer crossAccountUserId;
    /**
     * @return Error message.
     * 
     */
    private String errorMessage;
    /**
     * @return Exclude path. String of Json list. Up to 255 characters. e.g. `&#34;[\&#34;/home/work\&#34;]&#34;`
     * 
     */
    private String exclude;
    /**
     * @return The ID of destination file system.
     * 
     */
    private String fileSystemId;
    /**
     * @return The ID of the backup job.
     * 
     */
    private String id;
    /**
     * @return Include path. String of Json list. Up to 255 characters. e.g. `&#34;[\&#34;/var\&#34;]&#34;`
     * 
     */
    private String include;
    /**
     * @return The ID of target ECS instance.
     * 
     */
    private String instanceId;
    /**
     * @return The number of items restore job recovered.
     * 
     */
    private String itemsDone;
    /**
     * @return The total number of items restore job recovered.
     * 
     */
    private String itemsTotal;
    /**
     * @return File system creation time. UNIX time in seconds.
     * 
     */
    private String nasCreateTime;
    private List<GetBackupJobsJobOtsDetail> otsDetails;
    /**
     * @return List of backup path. e.g. `[&#34;/home&#34;, &#34;/var&#34;]`.
     * 
     */
    private List<String> paths;
    /**
     * @return The ID of a backup plan.
     * 
     */
    private String planId;
    /**
     * @return The prefix of Oss bucket files.
     * 
     */
    private String prefix;
    /**
     * @return Backup progress. The value is 100%*100.
     * 
     */
    private String progress;
    /**
     * @return The type of data source. Valid Values: `ECS_FILE`, `OSS`, `NAS`, `UDM_DISK`.
     * 
     */
    private String sourceType;
    /**
     * @return The scheduled backup start time. UNIX time seconds.
     * 
     */
    private String startTime;
    /**
     * @return The status of restore job. Valid values: `COMPLETE` , `PARTIAL_COMPLETE`, `FAILED`.
     * 
     */
    private String status;
    /**
     * @return The update time of backup job. UNIX time seconds.
     * 
     */
    private String updatedTime;
    /**
     * @return The ID of backup vault.
     * 
     */
    private String vaultId;

    private GetBackupJobsJob() {}
    /**
     * @return The actual data volume of the backup task (After deduplication) . Unit byte.
     * 
     */
    public String actualBytes() {
        return this.actualBytes;
    }
    /**
     * @return The actual number of items in the backup task. (Currently only file backup is available).
     * 
     */
    public String actualItems() {
        return this.actualItems;
    }
    /**
     * @return The name of backup job.
     * 
     */
    public String backJobName() {
        return this.backJobName;
    }
    /**
     * @return The ID of the backup job.
     * 
     */
    public String backupJobId() {
        return this.backupJobId;
    }
    /**
     * @return Backup type. Valid values: `COMPLETE`(full backup).
     * 
     */
    public String backupType() {
        return this.backupType;
    }
    /**
     * @return The name of target OSS bucket.
     * 
     */
    public String bucket() {
        return this.bucket;
    }
    /**
     * @return The amount of backup data (Incremental). Unit byte.
     * 
     */
    public String bytesDone() {
        return this.bytesDone;
    }
    /**
     * @return The total amount of data sources. Unit byte.
     * 
     */
    public String bytesTotal() {
        return this.bytesTotal;
    }
    /**
     * @return The completion time of backup job. UNIX time seconds.
     * 
     */
    public String completeTime() {
        return this.completeTime;
    }
    /**
     * @return The creation time of backup job. UNIX time seconds.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The role name created in the original account RAM backup by the cross account managed by the current account. It is valid only when `source_type` is `ECS_FILE`, `NAS`, `OSS` or `OTS`.
     * 
     */
    public String crossAccountRoleName() {
        return this.crossAccountRoleName;
    }
    /**
     * @return The type of the cross account backup. It is valid only when `source_type` is `ECS_FILE`, `NAS`, `OSS` or `OTS`.
     * 
     */
    public String crossAccountType() {
        return this.crossAccountType;
    }
    /**
     * @return The original account ID of the cross account backup managed by the current account. It is valid only when `source_type` is `ECS_FILE`, `NAS`, `OSS` or `OTS`.
     * 
     */
    public Integer crossAccountUserId() {
        return this.crossAccountUserId;
    }
    /**
     * @return Error message.
     * 
     */
    public String errorMessage() {
        return this.errorMessage;
    }
    /**
     * @return Exclude path. String of Json list. Up to 255 characters. e.g. `&#34;[\&#34;/home/work\&#34;]&#34;`
     * 
     */
    public String exclude() {
        return this.exclude;
    }
    /**
     * @return The ID of destination file system.
     * 
     */
    public String fileSystemId() {
        return this.fileSystemId;
    }
    /**
     * @return The ID of the backup job.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Include path. String of Json list. Up to 255 characters. e.g. `&#34;[\&#34;/var\&#34;]&#34;`
     * 
     */
    public String include() {
        return this.include;
    }
    /**
     * @return The ID of target ECS instance.
     * 
     */
    public String instanceId() {
        return this.instanceId;
    }
    /**
     * @return The number of items restore job recovered.
     * 
     */
    public String itemsDone() {
        return this.itemsDone;
    }
    /**
     * @return The total number of items restore job recovered.
     * 
     */
    public String itemsTotal() {
        return this.itemsTotal;
    }
    /**
     * @return File system creation time. UNIX time in seconds.
     * 
     */
    public String nasCreateTime() {
        return this.nasCreateTime;
    }
    public List<GetBackupJobsJobOtsDetail> otsDetails() {
        return this.otsDetails;
    }
    /**
     * @return List of backup path. e.g. `[&#34;/home&#34;, &#34;/var&#34;]`.
     * 
     */
    public List<String> paths() {
        return this.paths;
    }
    /**
     * @return The ID of a backup plan.
     * 
     */
    public String planId() {
        return this.planId;
    }
    /**
     * @return The prefix of Oss bucket files.
     * 
     */
    public String prefix() {
        return this.prefix;
    }
    /**
     * @return Backup progress. The value is 100%*100.
     * 
     */
    public String progress() {
        return this.progress;
    }
    /**
     * @return The type of data source. Valid Values: `ECS_FILE`, `OSS`, `NAS`, `UDM_DISK`.
     * 
     */
    public String sourceType() {
        return this.sourceType;
    }
    /**
     * @return The scheduled backup start time. UNIX time seconds.
     * 
     */
    public String startTime() {
        return this.startTime;
    }
    /**
     * @return The status of restore job. Valid values: `COMPLETE` , `PARTIAL_COMPLETE`, `FAILED`.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The update time of backup job. UNIX time seconds.
     * 
     */
    public String updatedTime() {
        return this.updatedTime;
    }
    /**
     * @return The ID of backup vault.
     * 
     */
    public String vaultId() {
        return this.vaultId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBackupJobsJob defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String actualBytes;
        private String actualItems;
        private String backJobName;
        private String backupJobId;
        private String backupType;
        private String bucket;
        private String bytesDone;
        private String bytesTotal;
        private String completeTime;
        private String createTime;
        private String crossAccountRoleName;
        private String crossAccountType;
        private Integer crossAccountUserId;
        private String errorMessage;
        private String exclude;
        private String fileSystemId;
        private String id;
        private String include;
        private String instanceId;
        private String itemsDone;
        private String itemsTotal;
        private String nasCreateTime;
        private List<GetBackupJobsJobOtsDetail> otsDetails;
        private List<String> paths;
        private String planId;
        private String prefix;
        private String progress;
        private String sourceType;
        private String startTime;
        private String status;
        private String updatedTime;
        private String vaultId;
        public Builder() {}
        public Builder(GetBackupJobsJob defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.actualBytes = defaults.actualBytes;
    	      this.actualItems = defaults.actualItems;
    	      this.backJobName = defaults.backJobName;
    	      this.backupJobId = defaults.backupJobId;
    	      this.backupType = defaults.backupType;
    	      this.bucket = defaults.bucket;
    	      this.bytesDone = defaults.bytesDone;
    	      this.bytesTotal = defaults.bytesTotal;
    	      this.completeTime = defaults.completeTime;
    	      this.createTime = defaults.createTime;
    	      this.crossAccountRoleName = defaults.crossAccountRoleName;
    	      this.crossAccountType = defaults.crossAccountType;
    	      this.crossAccountUserId = defaults.crossAccountUserId;
    	      this.errorMessage = defaults.errorMessage;
    	      this.exclude = defaults.exclude;
    	      this.fileSystemId = defaults.fileSystemId;
    	      this.id = defaults.id;
    	      this.include = defaults.include;
    	      this.instanceId = defaults.instanceId;
    	      this.itemsDone = defaults.itemsDone;
    	      this.itemsTotal = defaults.itemsTotal;
    	      this.nasCreateTime = defaults.nasCreateTime;
    	      this.otsDetails = defaults.otsDetails;
    	      this.paths = defaults.paths;
    	      this.planId = defaults.planId;
    	      this.prefix = defaults.prefix;
    	      this.progress = defaults.progress;
    	      this.sourceType = defaults.sourceType;
    	      this.startTime = defaults.startTime;
    	      this.status = defaults.status;
    	      this.updatedTime = defaults.updatedTime;
    	      this.vaultId = defaults.vaultId;
        }

        @CustomType.Setter
        public Builder actualBytes(String actualBytes) {
            this.actualBytes = Objects.requireNonNull(actualBytes);
            return this;
        }
        @CustomType.Setter
        public Builder actualItems(String actualItems) {
            this.actualItems = Objects.requireNonNull(actualItems);
            return this;
        }
        @CustomType.Setter
        public Builder backJobName(String backJobName) {
            this.backJobName = Objects.requireNonNull(backJobName);
            return this;
        }
        @CustomType.Setter
        public Builder backupJobId(String backupJobId) {
            this.backupJobId = Objects.requireNonNull(backupJobId);
            return this;
        }
        @CustomType.Setter
        public Builder backupType(String backupType) {
            this.backupType = Objects.requireNonNull(backupType);
            return this;
        }
        @CustomType.Setter
        public Builder bucket(String bucket) {
            this.bucket = Objects.requireNonNull(bucket);
            return this;
        }
        @CustomType.Setter
        public Builder bytesDone(String bytesDone) {
            this.bytesDone = Objects.requireNonNull(bytesDone);
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
        public Builder crossAccountRoleName(String crossAccountRoleName) {
            this.crossAccountRoleName = Objects.requireNonNull(crossAccountRoleName);
            return this;
        }
        @CustomType.Setter
        public Builder crossAccountType(String crossAccountType) {
            this.crossAccountType = Objects.requireNonNull(crossAccountType);
            return this;
        }
        @CustomType.Setter
        public Builder crossAccountUserId(Integer crossAccountUserId) {
            this.crossAccountUserId = Objects.requireNonNull(crossAccountUserId);
            return this;
        }
        @CustomType.Setter
        public Builder errorMessage(String errorMessage) {
            this.errorMessage = Objects.requireNonNull(errorMessage);
            return this;
        }
        @CustomType.Setter
        public Builder exclude(String exclude) {
            this.exclude = Objects.requireNonNull(exclude);
            return this;
        }
        @CustomType.Setter
        public Builder fileSystemId(String fileSystemId) {
            this.fileSystemId = Objects.requireNonNull(fileSystemId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder include(String include) {
            this.include = Objects.requireNonNull(include);
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(String instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder itemsDone(String itemsDone) {
            this.itemsDone = Objects.requireNonNull(itemsDone);
            return this;
        }
        @CustomType.Setter
        public Builder itemsTotal(String itemsTotal) {
            this.itemsTotal = Objects.requireNonNull(itemsTotal);
            return this;
        }
        @CustomType.Setter
        public Builder nasCreateTime(String nasCreateTime) {
            this.nasCreateTime = Objects.requireNonNull(nasCreateTime);
            return this;
        }
        @CustomType.Setter
        public Builder otsDetails(List<GetBackupJobsJobOtsDetail> otsDetails) {
            this.otsDetails = Objects.requireNonNull(otsDetails);
            return this;
        }
        public Builder otsDetails(GetBackupJobsJobOtsDetail... otsDetails) {
            return otsDetails(List.of(otsDetails));
        }
        @CustomType.Setter
        public Builder paths(List<String> paths) {
            this.paths = Objects.requireNonNull(paths);
            return this;
        }
        public Builder paths(String... paths) {
            return paths(List.of(paths));
        }
        @CustomType.Setter
        public Builder planId(String planId) {
            this.planId = Objects.requireNonNull(planId);
            return this;
        }
        @CustomType.Setter
        public Builder prefix(String prefix) {
            this.prefix = Objects.requireNonNull(prefix);
            return this;
        }
        @CustomType.Setter
        public Builder progress(String progress) {
            this.progress = Objects.requireNonNull(progress);
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
        public Builder updatedTime(String updatedTime) {
            this.updatedTime = Objects.requireNonNull(updatedTime);
            return this;
        }
        @CustomType.Setter
        public Builder vaultId(String vaultId) {
            this.vaultId = Objects.requireNonNull(vaultId);
            return this;
        }
        public GetBackupJobsJob build() {
            final var o = new GetBackupJobsJob();
            o.actualBytes = actualBytes;
            o.actualItems = actualItems;
            o.backJobName = backJobName;
            o.backupJobId = backupJobId;
            o.backupType = backupType;
            o.bucket = bucket;
            o.bytesDone = bytesDone;
            o.bytesTotal = bytesTotal;
            o.completeTime = completeTime;
            o.createTime = createTime;
            o.crossAccountRoleName = crossAccountRoleName;
            o.crossAccountType = crossAccountType;
            o.crossAccountUserId = crossAccountUserId;
            o.errorMessage = errorMessage;
            o.exclude = exclude;
            o.fileSystemId = fileSystemId;
            o.id = id;
            o.include = include;
            o.instanceId = instanceId;
            o.itemsDone = itemsDone;
            o.itemsTotal = itemsTotal;
            o.nasCreateTime = nasCreateTime;
            o.otsDetails = otsDetails;
            o.paths = paths;
            o.planId = planId;
            o.prefix = prefix;
            o.progress = progress;
            o.sourceType = sourceType;
            o.startTime = startTime;
            o.status = status;
            o.updatedTime = updatedTime;
            o.vaultId = vaultId;
            return o;
        }
    }
}