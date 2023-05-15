// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbase.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetInstancesInstance {
    /**
     * @return The Backup Status of the instance.
     * 
     */
    private String backupStatus;
    /**
     * @return Core node disk size, unit:GB.
     * 
     */
    private Integer coreDiskSize;
    /**
     * @return Cloud_ssd or cloud_efficiency
     * 
     */
    private String coreDiskType;
    /**
     * @return Like hbase.sn2.2xlarge, hbase.sn2.4xlarge, hbase.sn2.8xlarge and so on.
     * 
     */
    private String coreInstanceType;
    /**
     * @return Same with &#34;core_instance_quantity&#34;
     * 
     */
    private Integer coreNodeCount;
    /**
     * @return The created time of the instance.
     * 
     */
    private String createdTime;
    /**
     * @return The switch of delete protection.
     * 
     */
    private Boolean deletionProtection;
    /**
     * @return The engine of the instance.
     * 
     */
    private String engine;
    /**
     * @return The engine_version of the instance.
     * 
     */
    private String engineVersion;
    /**
     * @return The expire time of the instance.
     * 
     */
    private String expireTime;
    /**
     * @return The ID of the HBase instance.
     * 
     */
    private String id;
    /**
     * @return Like hbase.sn2.2xlarge, hbase.sn2.4xlarge, hbase.sn2.8xlarge and so on.
     * 
     */
    private String masterInstanceType;
    /**
     * @return The node count of master
     * 
     */
    private Integer masterNodeCount;
    /**
     * @return The name of the HBase instance.
     * 
     */
    private String name;
    /**
     * @return Classic network or VPC.
     * 
     */
    private String networkType;
    /**
     * @return Billing method. Value options are `PostPaid` for  Pay-As-You-Go and `PrePaid` for yearly or monthly subscription.
     * 
     */
    private String payType;
    /**
     * @return Region ID the instance belongs to.
     * 
     */
    private String regionId;
    /**
     * @return Status of the instance.
     * 
     */
    private String status;
    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    private @Nullable Map<String,Object> tags;
    /**
     * @return VPC ID the instance belongs to.
     * 
     */
    private String vpcId;
    /**
     * @return VSwitch ID the instance belongs to.
     * 
     */
    private String vswitchId;
    /**
     * @return Zone ID the instance belongs to.
     * 
     */
    private String zoneId;

    private GetInstancesInstance() {}
    /**
     * @return The Backup Status of the instance.
     * 
     */
    public String backupStatus() {
        return this.backupStatus;
    }
    /**
     * @return Core node disk size, unit:GB.
     * 
     */
    public Integer coreDiskSize() {
        return this.coreDiskSize;
    }
    /**
     * @return Cloud_ssd or cloud_efficiency
     * 
     */
    public String coreDiskType() {
        return this.coreDiskType;
    }
    /**
     * @return Like hbase.sn2.2xlarge, hbase.sn2.4xlarge, hbase.sn2.8xlarge and so on.
     * 
     */
    public String coreInstanceType() {
        return this.coreInstanceType;
    }
    /**
     * @return Same with &#34;core_instance_quantity&#34;
     * 
     */
    public Integer coreNodeCount() {
        return this.coreNodeCount;
    }
    /**
     * @return The created time of the instance.
     * 
     */
    public String createdTime() {
        return this.createdTime;
    }
    /**
     * @return The switch of delete protection.
     * 
     */
    public Boolean deletionProtection() {
        return this.deletionProtection;
    }
    /**
     * @return The engine of the instance.
     * 
     */
    public String engine() {
        return this.engine;
    }
    /**
     * @return The engine_version of the instance.
     * 
     */
    public String engineVersion() {
        return this.engineVersion;
    }
    /**
     * @return The expire time of the instance.
     * 
     */
    public String expireTime() {
        return this.expireTime;
    }
    /**
     * @return The ID of the HBase instance.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Like hbase.sn2.2xlarge, hbase.sn2.4xlarge, hbase.sn2.8xlarge and so on.
     * 
     */
    public String masterInstanceType() {
        return this.masterInstanceType;
    }
    /**
     * @return The node count of master
     * 
     */
    public Integer masterNodeCount() {
        return this.masterNodeCount;
    }
    /**
     * @return The name of the HBase instance.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Classic network or VPC.
     * 
     */
    public String networkType() {
        return this.networkType;
    }
    /**
     * @return Billing method. Value options are `PostPaid` for  Pay-As-You-Go and `PrePaid` for yearly or monthly subscription.
     * 
     */
    public String payType() {
        return this.payType;
    }
    /**
     * @return Region ID the instance belongs to.
     * 
     */
    public String regionId() {
        return this.regionId;
    }
    /**
     * @return Status of the instance.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Map<String,Object> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }
    /**
     * @return VPC ID the instance belongs to.
     * 
     */
    public String vpcId() {
        return this.vpcId;
    }
    /**
     * @return VSwitch ID the instance belongs to.
     * 
     */
    public String vswitchId() {
        return this.vswitchId;
    }
    /**
     * @return Zone ID the instance belongs to.
     * 
     */
    public String zoneId() {
        return this.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstancesInstance defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String backupStatus;
        private Integer coreDiskSize;
        private String coreDiskType;
        private String coreInstanceType;
        private Integer coreNodeCount;
        private String createdTime;
        private Boolean deletionProtection;
        private String engine;
        private String engineVersion;
        private String expireTime;
        private String id;
        private String masterInstanceType;
        private Integer masterNodeCount;
        private String name;
        private String networkType;
        private String payType;
        private String regionId;
        private String status;
        private @Nullable Map<String,Object> tags;
        private String vpcId;
        private String vswitchId;
        private String zoneId;
        public Builder() {}
        public Builder(GetInstancesInstance defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backupStatus = defaults.backupStatus;
    	      this.coreDiskSize = defaults.coreDiskSize;
    	      this.coreDiskType = defaults.coreDiskType;
    	      this.coreInstanceType = defaults.coreInstanceType;
    	      this.coreNodeCount = defaults.coreNodeCount;
    	      this.createdTime = defaults.createdTime;
    	      this.deletionProtection = defaults.deletionProtection;
    	      this.engine = defaults.engine;
    	      this.engineVersion = defaults.engineVersion;
    	      this.expireTime = defaults.expireTime;
    	      this.id = defaults.id;
    	      this.masterInstanceType = defaults.masterInstanceType;
    	      this.masterNodeCount = defaults.masterNodeCount;
    	      this.name = defaults.name;
    	      this.networkType = defaults.networkType;
    	      this.payType = defaults.payType;
    	      this.regionId = defaults.regionId;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.vpcId = defaults.vpcId;
    	      this.vswitchId = defaults.vswitchId;
    	      this.zoneId = defaults.zoneId;
        }

        @CustomType.Setter
        public Builder backupStatus(String backupStatus) {
            this.backupStatus = Objects.requireNonNull(backupStatus);
            return this;
        }
        @CustomType.Setter
        public Builder coreDiskSize(Integer coreDiskSize) {
            this.coreDiskSize = Objects.requireNonNull(coreDiskSize);
            return this;
        }
        @CustomType.Setter
        public Builder coreDiskType(String coreDiskType) {
            this.coreDiskType = Objects.requireNonNull(coreDiskType);
            return this;
        }
        @CustomType.Setter
        public Builder coreInstanceType(String coreInstanceType) {
            this.coreInstanceType = Objects.requireNonNull(coreInstanceType);
            return this;
        }
        @CustomType.Setter
        public Builder coreNodeCount(Integer coreNodeCount) {
            this.coreNodeCount = Objects.requireNonNull(coreNodeCount);
            return this;
        }
        @CustomType.Setter
        public Builder createdTime(String createdTime) {
            this.createdTime = Objects.requireNonNull(createdTime);
            return this;
        }
        @CustomType.Setter
        public Builder deletionProtection(Boolean deletionProtection) {
            this.deletionProtection = Objects.requireNonNull(deletionProtection);
            return this;
        }
        @CustomType.Setter
        public Builder engine(String engine) {
            this.engine = Objects.requireNonNull(engine);
            return this;
        }
        @CustomType.Setter
        public Builder engineVersion(String engineVersion) {
            this.engineVersion = Objects.requireNonNull(engineVersion);
            return this;
        }
        @CustomType.Setter
        public Builder expireTime(String expireTime) {
            this.expireTime = Objects.requireNonNull(expireTime);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder masterInstanceType(String masterInstanceType) {
            this.masterInstanceType = Objects.requireNonNull(masterInstanceType);
            return this;
        }
        @CustomType.Setter
        public Builder masterNodeCount(Integer masterNodeCount) {
            this.masterNodeCount = Objects.requireNonNull(masterNodeCount);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder networkType(String networkType) {
            this.networkType = Objects.requireNonNull(networkType);
            return this;
        }
        @CustomType.Setter
        public Builder payType(String payType) {
            this.payType = Objects.requireNonNull(payType);
            return this;
        }
        @CustomType.Setter
        public Builder regionId(String regionId) {
            this.regionId = Objects.requireNonNull(regionId);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,Object> tags) {
            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(String vpcId) {
            this.vpcId = Objects.requireNonNull(vpcId);
            return this;
        }
        @CustomType.Setter
        public Builder vswitchId(String vswitchId) {
            this.vswitchId = Objects.requireNonNull(vswitchId);
            return this;
        }
        @CustomType.Setter
        public Builder zoneId(String zoneId) {
            this.zoneId = Objects.requireNonNull(zoneId);
            return this;
        }
        public GetInstancesInstance build() {
            final var o = new GetInstancesInstance();
            o.backupStatus = backupStatus;
            o.coreDiskSize = coreDiskSize;
            o.coreDiskType = coreDiskType;
            o.coreInstanceType = coreInstanceType;
            o.coreNodeCount = coreNodeCount;
            o.createdTime = createdTime;
            o.deletionProtection = deletionProtection;
            o.engine = engine;
            o.engineVersion = engineVersion;
            o.expireTime = expireTime;
            o.id = id;
            o.masterInstanceType = masterInstanceType;
            o.masterNodeCount = masterNodeCount;
            o.name = name;
            o.networkType = networkType;
            o.payType = payType;
            o.regionId = regionId;
            o.status = status;
            o.tags = tags;
            o.vpcId = vpcId;
            o.vswitchId = vswitchId;
            o.zoneId = zoneId;
            return o;
        }
    }
}
