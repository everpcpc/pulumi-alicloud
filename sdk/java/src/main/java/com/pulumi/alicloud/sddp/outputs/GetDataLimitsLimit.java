// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sddp.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDataLimitsLimit {
    /**
     * @return Whether to enable the log auditing feature.
     * 
     */
    private Integer auditStatus;
    /**
     * @return The status of the connectivity test between the data asset and SDDP.
     * 
     */
    private Integer checkStatus;
    /**
     * @return The first ID of the resource.
     * 
     */
    private String dataLimitId;
    /**
     * @return The type of the database.
     * 
     */
    private String engineType;
    /**
     * @return The ID of the Data Limit.
     * 
     */
    private String id;
    /**
     * @return The name of the service to which the data asset belongs.
     * 
     */
    private String localName;
    /**
     * @return The retention period of raw logs after you enable the log auditing feature.
     * 
     */
    private Integer logStoreDay;
    /**
     * @return The ID of the data asset.
     * 
     */
    private String parentId;
    /**
     * @return The port that is used to connect to the database.
     * 
     */
    private Integer port;
    /**
     * @return The type of the service to which the data asset belongs.
     * 
     */
    private String resourceType;
    /**
     * @return The name of the user who owns the data asset.
     * 
     */
    private String userName;

    private GetDataLimitsLimit() {}
    /**
     * @return Whether to enable the log auditing feature.
     * 
     */
    public Integer auditStatus() {
        return this.auditStatus;
    }
    /**
     * @return The status of the connectivity test between the data asset and SDDP.
     * 
     */
    public Integer checkStatus() {
        return this.checkStatus;
    }
    /**
     * @return The first ID of the resource.
     * 
     */
    public String dataLimitId() {
        return this.dataLimitId;
    }
    /**
     * @return The type of the database.
     * 
     */
    public String engineType() {
        return this.engineType;
    }
    /**
     * @return The ID of the Data Limit.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the service to which the data asset belongs.
     * 
     */
    public String localName() {
        return this.localName;
    }
    /**
     * @return The retention period of raw logs after you enable the log auditing feature.
     * 
     */
    public Integer logStoreDay() {
        return this.logStoreDay;
    }
    /**
     * @return The ID of the data asset.
     * 
     */
    public String parentId() {
        return this.parentId;
    }
    /**
     * @return The port that is used to connect to the database.
     * 
     */
    public Integer port() {
        return this.port;
    }
    /**
     * @return The type of the service to which the data asset belongs.
     * 
     */
    public String resourceType() {
        return this.resourceType;
    }
    /**
     * @return The name of the user who owns the data asset.
     * 
     */
    public String userName() {
        return this.userName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDataLimitsLimit defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer auditStatus;
        private Integer checkStatus;
        private String dataLimitId;
        private String engineType;
        private String id;
        private String localName;
        private Integer logStoreDay;
        private String parentId;
        private Integer port;
        private String resourceType;
        private String userName;
        public Builder() {}
        public Builder(GetDataLimitsLimit defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.auditStatus = defaults.auditStatus;
    	      this.checkStatus = defaults.checkStatus;
    	      this.dataLimitId = defaults.dataLimitId;
    	      this.engineType = defaults.engineType;
    	      this.id = defaults.id;
    	      this.localName = defaults.localName;
    	      this.logStoreDay = defaults.logStoreDay;
    	      this.parentId = defaults.parentId;
    	      this.port = defaults.port;
    	      this.resourceType = defaults.resourceType;
    	      this.userName = defaults.userName;
        }

        @CustomType.Setter
        public Builder auditStatus(Integer auditStatus) {
            this.auditStatus = Objects.requireNonNull(auditStatus);
            return this;
        }
        @CustomType.Setter
        public Builder checkStatus(Integer checkStatus) {
            this.checkStatus = Objects.requireNonNull(checkStatus);
            return this;
        }
        @CustomType.Setter
        public Builder dataLimitId(String dataLimitId) {
            this.dataLimitId = Objects.requireNonNull(dataLimitId);
            return this;
        }
        @CustomType.Setter
        public Builder engineType(String engineType) {
            this.engineType = Objects.requireNonNull(engineType);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder localName(String localName) {
            this.localName = Objects.requireNonNull(localName);
            return this;
        }
        @CustomType.Setter
        public Builder logStoreDay(Integer logStoreDay) {
            this.logStoreDay = Objects.requireNonNull(logStoreDay);
            return this;
        }
        @CustomType.Setter
        public Builder parentId(String parentId) {
            this.parentId = Objects.requireNonNull(parentId);
            return this;
        }
        @CustomType.Setter
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        @CustomType.Setter
        public Builder resourceType(String resourceType) {
            this.resourceType = Objects.requireNonNull(resourceType);
            return this;
        }
        @CustomType.Setter
        public Builder userName(String userName) {
            this.userName = Objects.requireNonNull(userName);
            return this;
        }
        public GetDataLimitsLimit build() {
            final var o = new GetDataLimitsLimit();
            o.auditStatus = auditStatus;
            o.checkStatus = checkStatus;
            o.dataLimitId = dataLimitId;
            o.engineType = engineType;
            o.id = id;
            o.localName = localName;
            o.logStoreDay = logStoreDay;
            o.parentId = parentId;
            o.port = port;
            o.resourceType = resourceType;
            o.userName = userName;
            return o;
        }
    }
}