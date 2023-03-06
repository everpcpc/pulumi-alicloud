// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dts.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetMigrationJobsJob {
    /**
     * @return Whether or not to execute DTS supports schema migration, full data migration, or full-data initialization.
     * 
     */
    private Boolean dataInitialization;
    /**
     * @return Whether to perform incremental data migration for migration types or synchronization values include:
     * 
     */
    private Boolean dataSynchronization;
    /**
     * @return The Migration object, in the format of JSON strings.
     * 
     */
    private String dbList;
    /**
     * @return The name of migrate the database.
     * 
     */
    private String destinationEndpointDataBaseName;
    /**
     * @return The type of destination database.
     * 
     */
    private String destinationEndpointEngineName;
    /**
     * @return The ID of destination instance.
     * 
     */
    private String destinationEndpointInstanceId;
    /**
     * @return The type of destination instance.
     * 
     */
    private String destinationEndpointInstanceType;
    /**
     * @return The ip of source endpoint.
     * 
     */
    private String destinationEndpointIp;
    /**
     * @return The SID of Oracle database.
     * 
     */
    private String destinationEndpointOracleSid;
    /**
     * @return The port of source endpoint.
     * 
     */
    private String destinationEndpointPort;
    /**
     * @return The region of destination instance.
     * 
     */
    private String destinationEndpointRegion;
    /**
     * @return The username of database account.
     * 
     */
    private String destinationEndpointUserName;
    /**
     * @return The Migration instance ID. The ID of `alicloud.dts.MigrationInstance`.
     * 
     */
    private String dtsInstanceId;
    /**
     * @return The ID of the Migration Job.
     * 
     */
    private String dtsJobId;
    /**
     * @return The name of synchronization job.
     * 
     */
    private String dtsJobName;
    /**
     * @return The ID of the Migration Job. Its value is same as `dts_job_id`.
     * 
     */
    private String id;
    /**
     * @return The payment type of the Migration Instance.
     * 
     */
    private String paymentType;
    /**
     * @return The name of migrate the database.
     * 
     */
    private String sourceEndpointDatabaseName;
    /**
     * @return The type of source database.
     * 
     */
    private String sourceEndpointEngineName;
    /**
     * @return The ID of source instance.
     * 
     */
    private String sourceEndpointInstanceId;
    /**
     * @return The type of source instance.
     * 
     */
    private String sourceEndpointInstanceType;
    /**
     * @return The ip of source endpoint.
     * 
     */
    private String sourceEndpointIp;
    /**
     * @return The SID of Oracle database.
     * 
     */
    private String sourceEndpointOracleSid;
    /**
     * @return The Alibaba Cloud account ID to which the source instance belongs.
     * 
     */
    private String sourceEndpointOwnerId;
    /**
     * @return The port of source endpoint.
     * 
     */
    private String sourceEndpointPort;
    /**
     * @return The region of source instance.
     * 
     */
    private String sourceEndpointRegion;
    /**
     * @return The name of the role configured for the cloud account to which the source instance belongs.
     * 
     */
    private String sourceEndpointRole;
    /**
     * @return The username of database account.
     * 
     */
    private String sourceEndpointUserName;
    /**
     * @return The status of the resource.
     * 
     */
    private String status;
    /**
     * @return Whether to perform a database table structure to migrate or initialization.
     * 
     */
    private Boolean structureInitialization;

    private GetMigrationJobsJob() {}
    /**
     * @return Whether or not to execute DTS supports schema migration, full data migration, or full-data initialization.
     * 
     */
    public Boolean dataInitialization() {
        return this.dataInitialization;
    }
    /**
     * @return Whether to perform incremental data migration for migration types or synchronization values include:
     * 
     */
    public Boolean dataSynchronization() {
        return this.dataSynchronization;
    }
    /**
     * @return The Migration object, in the format of JSON strings.
     * 
     */
    public String dbList() {
        return this.dbList;
    }
    /**
     * @return The name of migrate the database.
     * 
     */
    public String destinationEndpointDataBaseName() {
        return this.destinationEndpointDataBaseName;
    }
    /**
     * @return The type of destination database.
     * 
     */
    public String destinationEndpointEngineName() {
        return this.destinationEndpointEngineName;
    }
    /**
     * @return The ID of destination instance.
     * 
     */
    public String destinationEndpointInstanceId() {
        return this.destinationEndpointInstanceId;
    }
    /**
     * @return The type of destination instance.
     * 
     */
    public String destinationEndpointInstanceType() {
        return this.destinationEndpointInstanceType;
    }
    /**
     * @return The ip of source endpoint.
     * 
     */
    public String destinationEndpointIp() {
        return this.destinationEndpointIp;
    }
    /**
     * @return The SID of Oracle database.
     * 
     */
    public String destinationEndpointOracleSid() {
        return this.destinationEndpointOracleSid;
    }
    /**
     * @return The port of source endpoint.
     * 
     */
    public String destinationEndpointPort() {
        return this.destinationEndpointPort;
    }
    /**
     * @return The region of destination instance.
     * 
     */
    public String destinationEndpointRegion() {
        return this.destinationEndpointRegion;
    }
    /**
     * @return The username of database account.
     * 
     */
    public String destinationEndpointUserName() {
        return this.destinationEndpointUserName;
    }
    /**
     * @return The Migration instance ID. The ID of `alicloud.dts.MigrationInstance`.
     * 
     */
    public String dtsInstanceId() {
        return this.dtsInstanceId;
    }
    /**
     * @return The ID of the Migration Job.
     * 
     */
    public String dtsJobId() {
        return this.dtsJobId;
    }
    /**
     * @return The name of synchronization job.
     * 
     */
    public String dtsJobName() {
        return this.dtsJobName;
    }
    /**
     * @return The ID of the Migration Job. Its value is same as `dts_job_id`.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The payment type of the Migration Instance.
     * 
     */
    public String paymentType() {
        return this.paymentType;
    }
    /**
     * @return The name of migrate the database.
     * 
     */
    public String sourceEndpointDatabaseName() {
        return this.sourceEndpointDatabaseName;
    }
    /**
     * @return The type of source database.
     * 
     */
    public String sourceEndpointEngineName() {
        return this.sourceEndpointEngineName;
    }
    /**
     * @return The ID of source instance.
     * 
     */
    public String sourceEndpointInstanceId() {
        return this.sourceEndpointInstanceId;
    }
    /**
     * @return The type of source instance.
     * 
     */
    public String sourceEndpointInstanceType() {
        return this.sourceEndpointInstanceType;
    }
    /**
     * @return The ip of source endpoint.
     * 
     */
    public String sourceEndpointIp() {
        return this.sourceEndpointIp;
    }
    /**
     * @return The SID of Oracle database.
     * 
     */
    public String sourceEndpointOracleSid() {
        return this.sourceEndpointOracleSid;
    }
    /**
     * @return The Alibaba Cloud account ID to which the source instance belongs.
     * 
     */
    public String sourceEndpointOwnerId() {
        return this.sourceEndpointOwnerId;
    }
    /**
     * @return The port of source endpoint.
     * 
     */
    public String sourceEndpointPort() {
        return this.sourceEndpointPort;
    }
    /**
     * @return The region of source instance.
     * 
     */
    public String sourceEndpointRegion() {
        return this.sourceEndpointRegion;
    }
    /**
     * @return The name of the role configured for the cloud account to which the source instance belongs.
     * 
     */
    public String sourceEndpointRole() {
        return this.sourceEndpointRole;
    }
    /**
     * @return The username of database account.
     * 
     */
    public String sourceEndpointUserName() {
        return this.sourceEndpointUserName;
    }
    /**
     * @return The status of the resource.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return Whether to perform a database table structure to migrate or initialization.
     * 
     */
    public Boolean structureInitialization() {
        return this.structureInitialization;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMigrationJobsJob defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean dataInitialization;
        private Boolean dataSynchronization;
        private String dbList;
        private String destinationEndpointDataBaseName;
        private String destinationEndpointEngineName;
        private String destinationEndpointInstanceId;
        private String destinationEndpointInstanceType;
        private String destinationEndpointIp;
        private String destinationEndpointOracleSid;
        private String destinationEndpointPort;
        private String destinationEndpointRegion;
        private String destinationEndpointUserName;
        private String dtsInstanceId;
        private String dtsJobId;
        private String dtsJobName;
        private String id;
        private String paymentType;
        private String sourceEndpointDatabaseName;
        private String sourceEndpointEngineName;
        private String sourceEndpointInstanceId;
        private String sourceEndpointInstanceType;
        private String sourceEndpointIp;
        private String sourceEndpointOracleSid;
        private String sourceEndpointOwnerId;
        private String sourceEndpointPort;
        private String sourceEndpointRegion;
        private String sourceEndpointRole;
        private String sourceEndpointUserName;
        private String status;
        private Boolean structureInitialization;
        public Builder() {}
        public Builder(GetMigrationJobsJob defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dataInitialization = defaults.dataInitialization;
    	      this.dataSynchronization = defaults.dataSynchronization;
    	      this.dbList = defaults.dbList;
    	      this.destinationEndpointDataBaseName = defaults.destinationEndpointDataBaseName;
    	      this.destinationEndpointEngineName = defaults.destinationEndpointEngineName;
    	      this.destinationEndpointInstanceId = defaults.destinationEndpointInstanceId;
    	      this.destinationEndpointInstanceType = defaults.destinationEndpointInstanceType;
    	      this.destinationEndpointIp = defaults.destinationEndpointIp;
    	      this.destinationEndpointOracleSid = defaults.destinationEndpointOracleSid;
    	      this.destinationEndpointPort = defaults.destinationEndpointPort;
    	      this.destinationEndpointRegion = defaults.destinationEndpointRegion;
    	      this.destinationEndpointUserName = defaults.destinationEndpointUserName;
    	      this.dtsInstanceId = defaults.dtsInstanceId;
    	      this.dtsJobId = defaults.dtsJobId;
    	      this.dtsJobName = defaults.dtsJobName;
    	      this.id = defaults.id;
    	      this.paymentType = defaults.paymentType;
    	      this.sourceEndpointDatabaseName = defaults.sourceEndpointDatabaseName;
    	      this.sourceEndpointEngineName = defaults.sourceEndpointEngineName;
    	      this.sourceEndpointInstanceId = defaults.sourceEndpointInstanceId;
    	      this.sourceEndpointInstanceType = defaults.sourceEndpointInstanceType;
    	      this.sourceEndpointIp = defaults.sourceEndpointIp;
    	      this.sourceEndpointOracleSid = defaults.sourceEndpointOracleSid;
    	      this.sourceEndpointOwnerId = defaults.sourceEndpointOwnerId;
    	      this.sourceEndpointPort = defaults.sourceEndpointPort;
    	      this.sourceEndpointRegion = defaults.sourceEndpointRegion;
    	      this.sourceEndpointRole = defaults.sourceEndpointRole;
    	      this.sourceEndpointUserName = defaults.sourceEndpointUserName;
    	      this.status = defaults.status;
    	      this.structureInitialization = defaults.structureInitialization;
        }

        @CustomType.Setter
        public Builder dataInitialization(Boolean dataInitialization) {
            this.dataInitialization = Objects.requireNonNull(dataInitialization);
            return this;
        }
        @CustomType.Setter
        public Builder dataSynchronization(Boolean dataSynchronization) {
            this.dataSynchronization = Objects.requireNonNull(dataSynchronization);
            return this;
        }
        @CustomType.Setter
        public Builder dbList(String dbList) {
            this.dbList = Objects.requireNonNull(dbList);
            return this;
        }
        @CustomType.Setter
        public Builder destinationEndpointDataBaseName(String destinationEndpointDataBaseName) {
            this.destinationEndpointDataBaseName = Objects.requireNonNull(destinationEndpointDataBaseName);
            return this;
        }
        @CustomType.Setter
        public Builder destinationEndpointEngineName(String destinationEndpointEngineName) {
            this.destinationEndpointEngineName = Objects.requireNonNull(destinationEndpointEngineName);
            return this;
        }
        @CustomType.Setter
        public Builder destinationEndpointInstanceId(String destinationEndpointInstanceId) {
            this.destinationEndpointInstanceId = Objects.requireNonNull(destinationEndpointInstanceId);
            return this;
        }
        @CustomType.Setter
        public Builder destinationEndpointInstanceType(String destinationEndpointInstanceType) {
            this.destinationEndpointInstanceType = Objects.requireNonNull(destinationEndpointInstanceType);
            return this;
        }
        @CustomType.Setter
        public Builder destinationEndpointIp(String destinationEndpointIp) {
            this.destinationEndpointIp = Objects.requireNonNull(destinationEndpointIp);
            return this;
        }
        @CustomType.Setter
        public Builder destinationEndpointOracleSid(String destinationEndpointOracleSid) {
            this.destinationEndpointOracleSid = Objects.requireNonNull(destinationEndpointOracleSid);
            return this;
        }
        @CustomType.Setter
        public Builder destinationEndpointPort(String destinationEndpointPort) {
            this.destinationEndpointPort = Objects.requireNonNull(destinationEndpointPort);
            return this;
        }
        @CustomType.Setter
        public Builder destinationEndpointRegion(String destinationEndpointRegion) {
            this.destinationEndpointRegion = Objects.requireNonNull(destinationEndpointRegion);
            return this;
        }
        @CustomType.Setter
        public Builder destinationEndpointUserName(String destinationEndpointUserName) {
            this.destinationEndpointUserName = Objects.requireNonNull(destinationEndpointUserName);
            return this;
        }
        @CustomType.Setter
        public Builder dtsInstanceId(String dtsInstanceId) {
            this.dtsInstanceId = Objects.requireNonNull(dtsInstanceId);
            return this;
        }
        @CustomType.Setter
        public Builder dtsJobId(String dtsJobId) {
            this.dtsJobId = Objects.requireNonNull(dtsJobId);
            return this;
        }
        @CustomType.Setter
        public Builder dtsJobName(String dtsJobName) {
            this.dtsJobName = Objects.requireNonNull(dtsJobName);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder paymentType(String paymentType) {
            this.paymentType = Objects.requireNonNull(paymentType);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointDatabaseName(String sourceEndpointDatabaseName) {
            this.sourceEndpointDatabaseName = Objects.requireNonNull(sourceEndpointDatabaseName);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointEngineName(String sourceEndpointEngineName) {
            this.sourceEndpointEngineName = Objects.requireNonNull(sourceEndpointEngineName);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointInstanceId(String sourceEndpointInstanceId) {
            this.sourceEndpointInstanceId = Objects.requireNonNull(sourceEndpointInstanceId);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointInstanceType(String sourceEndpointInstanceType) {
            this.sourceEndpointInstanceType = Objects.requireNonNull(sourceEndpointInstanceType);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointIp(String sourceEndpointIp) {
            this.sourceEndpointIp = Objects.requireNonNull(sourceEndpointIp);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointOracleSid(String sourceEndpointOracleSid) {
            this.sourceEndpointOracleSid = Objects.requireNonNull(sourceEndpointOracleSid);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointOwnerId(String sourceEndpointOwnerId) {
            this.sourceEndpointOwnerId = Objects.requireNonNull(sourceEndpointOwnerId);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointPort(String sourceEndpointPort) {
            this.sourceEndpointPort = Objects.requireNonNull(sourceEndpointPort);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointRegion(String sourceEndpointRegion) {
            this.sourceEndpointRegion = Objects.requireNonNull(sourceEndpointRegion);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointRole(String sourceEndpointRole) {
            this.sourceEndpointRole = Objects.requireNonNull(sourceEndpointRole);
            return this;
        }
        @CustomType.Setter
        public Builder sourceEndpointUserName(String sourceEndpointUserName) {
            this.sourceEndpointUserName = Objects.requireNonNull(sourceEndpointUserName);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder structureInitialization(Boolean structureInitialization) {
            this.structureInitialization = Objects.requireNonNull(structureInitialization);
            return this;
        }
        public GetMigrationJobsJob build() {
            final var o = new GetMigrationJobsJob();
            o.dataInitialization = dataInitialization;
            o.dataSynchronization = dataSynchronization;
            o.dbList = dbList;
            o.destinationEndpointDataBaseName = destinationEndpointDataBaseName;
            o.destinationEndpointEngineName = destinationEndpointEngineName;
            o.destinationEndpointInstanceId = destinationEndpointInstanceId;
            o.destinationEndpointInstanceType = destinationEndpointInstanceType;
            o.destinationEndpointIp = destinationEndpointIp;
            o.destinationEndpointOracleSid = destinationEndpointOracleSid;
            o.destinationEndpointPort = destinationEndpointPort;
            o.destinationEndpointRegion = destinationEndpointRegion;
            o.destinationEndpointUserName = destinationEndpointUserName;
            o.dtsInstanceId = dtsInstanceId;
            o.dtsJobId = dtsJobId;
            o.dtsJobName = dtsJobName;
            o.id = id;
            o.paymentType = paymentType;
            o.sourceEndpointDatabaseName = sourceEndpointDatabaseName;
            o.sourceEndpointEngineName = sourceEndpointEngineName;
            o.sourceEndpointInstanceId = sourceEndpointInstanceId;
            o.sourceEndpointInstanceType = sourceEndpointInstanceType;
            o.sourceEndpointIp = sourceEndpointIp;
            o.sourceEndpointOracleSid = sourceEndpointOracleSid;
            o.sourceEndpointOwnerId = sourceEndpointOwnerId;
            o.sourceEndpointPort = sourceEndpointPort;
            o.sourceEndpointRegion = sourceEndpointRegion;
            o.sourceEndpointRole = sourceEndpointRole;
            o.sourceEndpointUserName = sourceEndpointUserName;
            o.status = status;
            o.structureInitialization = structureInitialization;
            return o;
        }
    }
}