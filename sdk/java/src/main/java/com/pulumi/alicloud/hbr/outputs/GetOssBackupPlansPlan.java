// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetOssBackupPlansPlan {
    /**
     * @return Backup type. Valid values: `COMPLETE`.
     * 
     */
    private String backupType;
    /**
     * @return The name of OSS bucket.
     * 
     */
    private String bucket;
    /**
     * @return The creation time of the backup plan. UNIX time in seconds.
     * 
     */
    private String createdTime;
    /**
     * @return Whether to be suspended. Valid values: `true`, `false`.
     * 
     */
    private Boolean disabled;
    /**
     * @return The ID of Oss backup plan.
     * 
     */
    private String id;
    /**
     * @return The ID of Oss backup plan.
     * 
     */
    private String ossBackupPlanId;
    /**
     * @return The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
     * 
     */
    private String ossBackupPlanName;
    /**
     * @return Backup prefix.
     * 
     */
    private String prefix;
    /**
     * @return Backup retention days, the minimum is 1.
     * 
     */
    private String retention;
    /**
     * @return Backup strategy. Optional format: I|{startTime}|{interval}. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed yet, the next backup task will not be triggered.
     * 
     */
    private String schedule;
    /**
     * @return The update time of the backup plan. UNIX time in seconds.
     * 
     */
    private String updatedTime;
    /**
     * @return The ID of backup vault.
     * 
     */
    private String vaultId;

    private GetOssBackupPlansPlan() {}
    /**
     * @return Backup type. Valid values: `COMPLETE`.
     * 
     */
    public String backupType() {
        return this.backupType;
    }
    /**
     * @return The name of OSS bucket.
     * 
     */
    public String bucket() {
        return this.bucket;
    }
    /**
     * @return The creation time of the backup plan. UNIX time in seconds.
     * 
     */
    public String createdTime() {
        return this.createdTime;
    }
    /**
     * @return Whether to be suspended. Valid values: `true`, `false`.
     * 
     */
    public Boolean disabled() {
        return this.disabled;
    }
    /**
     * @return The ID of Oss backup plan.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The ID of Oss backup plan.
     * 
     */
    public String ossBackupPlanId() {
        return this.ossBackupPlanId;
    }
    /**
     * @return The name of the backup plan. 1~64 characters, the backup plan name of each data source type in a single warehouse required to be unique.
     * 
     */
    public String ossBackupPlanName() {
        return this.ossBackupPlanName;
    }
    /**
     * @return Backup prefix.
     * 
     */
    public String prefix() {
        return this.prefix;
    }
    /**
     * @return Backup retention days, the minimum is 1.
     * 
     */
    public String retention() {
        return this.retention;
    }
    /**
     * @return Backup strategy. Optional format: I|{startTime}|{interval}. It means to execute a backup task every {interval} starting from {startTime}. The backup task for the elapsed time will not be compensated. If the last backup task is not completed yet, the next backup task will not be triggered.
     * 
     */
    public String schedule() {
        return this.schedule;
    }
    /**
     * @return The update time of the backup plan. UNIX time in seconds.
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

    public static Builder builder(GetOssBackupPlansPlan defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String backupType;
        private String bucket;
        private String createdTime;
        private Boolean disabled;
        private String id;
        private String ossBackupPlanId;
        private String ossBackupPlanName;
        private String prefix;
        private String retention;
        private String schedule;
        private String updatedTime;
        private String vaultId;
        public Builder() {}
        public Builder(GetOssBackupPlansPlan defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backupType = defaults.backupType;
    	      this.bucket = defaults.bucket;
    	      this.createdTime = defaults.createdTime;
    	      this.disabled = defaults.disabled;
    	      this.id = defaults.id;
    	      this.ossBackupPlanId = defaults.ossBackupPlanId;
    	      this.ossBackupPlanName = defaults.ossBackupPlanName;
    	      this.prefix = defaults.prefix;
    	      this.retention = defaults.retention;
    	      this.schedule = defaults.schedule;
    	      this.updatedTime = defaults.updatedTime;
    	      this.vaultId = defaults.vaultId;
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
        public Builder createdTime(String createdTime) {
            this.createdTime = Objects.requireNonNull(createdTime);
            return this;
        }
        @CustomType.Setter
        public Builder disabled(Boolean disabled) {
            this.disabled = Objects.requireNonNull(disabled);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ossBackupPlanId(String ossBackupPlanId) {
            this.ossBackupPlanId = Objects.requireNonNull(ossBackupPlanId);
            return this;
        }
        @CustomType.Setter
        public Builder ossBackupPlanName(String ossBackupPlanName) {
            this.ossBackupPlanName = Objects.requireNonNull(ossBackupPlanName);
            return this;
        }
        @CustomType.Setter
        public Builder prefix(String prefix) {
            this.prefix = Objects.requireNonNull(prefix);
            return this;
        }
        @CustomType.Setter
        public Builder retention(String retention) {
            this.retention = Objects.requireNonNull(retention);
            return this;
        }
        @CustomType.Setter
        public Builder schedule(String schedule) {
            this.schedule = Objects.requireNonNull(schedule);
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
        public GetOssBackupPlansPlan build() {
            final var o = new GetOssBackupPlansPlan();
            o.backupType = backupType;
            o.bucket = bucket;
            o.createdTime = createdTime;
            o.disabled = disabled;
            o.id = id;
            o.ossBackupPlanId = ossBackupPlanId;
            o.ossBackupPlanName = ossBackupPlanName;
            o.prefix = prefix;
            o.retention = retention;
            o.schedule = schedule;
            o.updatedTime = updatedTime;
            o.vaultId = vaultId;
            return o;
        }
    }
}