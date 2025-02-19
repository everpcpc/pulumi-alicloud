// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cassandra.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetBackupPlansPlan {
    /**
     * @return Specifies whether to activate the backup plan.
     * 
     */
    private Boolean active;
    /**
     * @return The backup cycle. Valid values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, and Sunday.
     * 
     */
    private String backupPeriod;
    /**
     * @return The start time of the backup task each day. The time is displayed in UTC and denoted by Z.
     * 
     */
    private String backupTime;
    /**
     * @return The ID of the cluster for the backup.
     * 
     */
    private String clusterId;
    /**
     * @return The time when the backup plan was created.
     * 
     */
    private String createTime;
    /**
     * @return The ID of the data center for the backup in the cluster.
     * 
     */
    private String dataCenterId;
    /**
     * @return The ID of the Backup Plan.
     * 
     */
    private String id;
    /**
     * @return The duration for which you want to retain the backup. Valid values: 1 to 30. Unit: days.
     * 
     */
    private Integer retentionPeriod;

    private GetBackupPlansPlan() {}
    /**
     * @return Specifies whether to activate the backup plan.
     * 
     */
    public Boolean active() {
        return this.active;
    }
    /**
     * @return The backup cycle. Valid values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, and Sunday.
     * 
     */
    public String backupPeriod() {
        return this.backupPeriod;
    }
    /**
     * @return The start time of the backup task each day. The time is displayed in UTC and denoted by Z.
     * 
     */
    public String backupTime() {
        return this.backupTime;
    }
    /**
     * @return The ID of the cluster for the backup.
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }
    /**
     * @return The time when the backup plan was created.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The ID of the data center for the backup in the cluster.
     * 
     */
    public String dataCenterId() {
        return this.dataCenterId;
    }
    /**
     * @return The ID of the Backup Plan.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The duration for which you want to retain the backup. Valid values: 1 to 30. Unit: days.
     * 
     */
    public Integer retentionPeriod() {
        return this.retentionPeriod;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBackupPlansPlan defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean active;
        private String backupPeriod;
        private String backupTime;
        private String clusterId;
        private String createTime;
        private String dataCenterId;
        private String id;
        private Integer retentionPeriod;
        public Builder() {}
        public Builder(GetBackupPlansPlan defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.active = defaults.active;
    	      this.backupPeriod = defaults.backupPeriod;
    	      this.backupTime = defaults.backupTime;
    	      this.clusterId = defaults.clusterId;
    	      this.createTime = defaults.createTime;
    	      this.dataCenterId = defaults.dataCenterId;
    	      this.id = defaults.id;
    	      this.retentionPeriod = defaults.retentionPeriod;
        }

        @CustomType.Setter
        public Builder active(Boolean active) {
            this.active = Objects.requireNonNull(active);
            return this;
        }
        @CustomType.Setter
        public Builder backupPeriod(String backupPeriod) {
            this.backupPeriod = Objects.requireNonNull(backupPeriod);
            return this;
        }
        @CustomType.Setter
        public Builder backupTime(String backupTime) {
            this.backupTime = Objects.requireNonNull(backupTime);
            return this;
        }
        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            this.clusterId = Objects.requireNonNull(clusterId);
            return this;
        }
        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder dataCenterId(String dataCenterId) {
            this.dataCenterId = Objects.requireNonNull(dataCenterId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder retentionPeriod(Integer retentionPeriod) {
            this.retentionPeriod = Objects.requireNonNull(retentionPeriod);
            return this;
        }
        public GetBackupPlansPlan build() {
            final var o = new GetBackupPlansPlan();
            o.active = active;
            o.backupPeriod = backupPeriod;
            o.backupTime = backupTime;
            o.clusterId = clusterId;
            o.createTime = createTime;
            o.dataCenterId = dataCenterId;
            o.id = id;
            o.retentionPeriod = retentionPeriod;
            return o;
        }
    }
}
