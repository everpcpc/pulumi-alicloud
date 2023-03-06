// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.clickhouse.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetBackupPoliciesPolicy {
    /**
     * @return Data backup days. Valid values: `7` to `730`.
     * 
     */
    private Integer backupRetentionPeriod;
    /**
     * @return The db cluster id.
     * 
     */
    private String dbClusterId;
    /**
     * @return The ID of the Backup Policy.
     * 
     */
    private String id;
    /**
     * @return DBCluster Backup period.
     * 
     */
    private List<String> preferredBackupPeriods;
    /**
     * @return Backup Time, UTC time.
     * 
     */
    private String preferredBackupTime;
    /**
     * @return The status of the resource.
     * 
     */
    private String status;

    private GetBackupPoliciesPolicy() {}
    /**
     * @return Data backup days. Valid values: `7` to `730`.
     * 
     */
    public Integer backupRetentionPeriod() {
        return this.backupRetentionPeriod;
    }
    /**
     * @return The db cluster id.
     * 
     */
    public String dbClusterId() {
        return this.dbClusterId;
    }
    /**
     * @return The ID of the Backup Policy.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return DBCluster Backup period.
     * 
     */
    public List<String> preferredBackupPeriods() {
        return this.preferredBackupPeriods;
    }
    /**
     * @return Backup Time, UTC time.
     * 
     */
    public String preferredBackupTime() {
        return this.preferredBackupTime;
    }
    /**
     * @return The status of the resource.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBackupPoliciesPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer backupRetentionPeriod;
        private String dbClusterId;
        private String id;
        private List<String> preferredBackupPeriods;
        private String preferredBackupTime;
        private String status;
        public Builder() {}
        public Builder(GetBackupPoliciesPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backupRetentionPeriod = defaults.backupRetentionPeriod;
    	      this.dbClusterId = defaults.dbClusterId;
    	      this.id = defaults.id;
    	      this.preferredBackupPeriods = defaults.preferredBackupPeriods;
    	      this.preferredBackupTime = defaults.preferredBackupTime;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder backupRetentionPeriod(Integer backupRetentionPeriod) {
            this.backupRetentionPeriod = Objects.requireNonNull(backupRetentionPeriod);
            return this;
        }
        @CustomType.Setter
        public Builder dbClusterId(String dbClusterId) {
            this.dbClusterId = Objects.requireNonNull(dbClusterId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder preferredBackupPeriods(List<String> preferredBackupPeriods) {
            this.preferredBackupPeriods = Objects.requireNonNull(preferredBackupPeriods);
            return this;
        }
        public Builder preferredBackupPeriods(String... preferredBackupPeriods) {
            return preferredBackupPeriods(List.of(preferredBackupPeriods));
        }
        @CustomType.Setter
        public Builder preferredBackupTime(String preferredBackupTime) {
            this.preferredBackupTime = Objects.requireNonNull(preferredBackupTime);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetBackupPoliciesPolicy build() {
            final var o = new GetBackupPoliciesPolicy();
            o.backupRetentionPeriod = backupRetentionPeriod;
            o.dbClusterId = dbClusterId;
            o.id = id;
            o.preferredBackupPeriods = preferredBackupPeriods;
            o.preferredBackupTime = preferredBackupTime;
            o.status = status;
            return o;
        }
    }
}