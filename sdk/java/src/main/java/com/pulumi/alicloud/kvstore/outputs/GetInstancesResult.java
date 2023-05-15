// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.kvstore.outputs;

import com.pulumi.alicloud.kvstore.outputs.GetInstancesInstance;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetInstancesResult {
    private @Nullable String architectureType;
    private @Nullable String editionType;
    private @Nullable Boolean enableDetails;
    /**
     * @return The engine version of the instance.
     * 
     */
    private @Nullable String engineVersion;
    private @Nullable String expired;
    private @Nullable Boolean globalInstance;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of KVStore Instance IDs.
     * 
     */
    private List<String> ids;
    /**
     * @return (Optional) Type of the applied ApsaraDB for instance.
     * For more information, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/61135.htm).
     * 
     */
    private @Nullable String instanceClass;
    /**
     * @return (Optional) Database type. Valid Values: `Memcache`, `Redis`. If no value is specified, all types are returned.
     * 
     */
    private @Nullable String instanceType;
    /**
     * @return A list of KVStore Instances. Its every element contains the following attributes:
     * 
     */
    private List<GetInstancesInstance> instances;
    private @Nullable String nameRegex;
    /**
     * @return A list of KVStore Instance names.
     * 
     */
    private List<String> names;
    /**
     * @return The network type of the instance.
     * 
     */
    private @Nullable String networkType;
    private @Nullable String outputFile;
    /**
     * @return Billing method. Valid Values: `PostPaid` for  Pay-As-You-Go and `PrePaid` for subscription.
     * 
     */
    private @Nullable String paymentType;
    private @Nullable String resourceGroupId;
    private @Nullable String searchKey;
    /**
     * @return Status of the instance.
     * 
     */
    private @Nullable String status;
    private @Nullable Map<String,Object> tags;
    /**
     * @return VPC ID the instance belongs to.
     * 
     */
    private @Nullable String vpcId;
    /**
     * @return VSwitch ID the instance belongs to.
     * 
     */
    private @Nullable String vswitchId;
    /**
     * @return The ID of zone.
     * 
     */
    private @Nullable String zoneId;

    private GetInstancesResult() {}
    public Optional<String> architectureType() {
        return Optional.ofNullable(this.architectureType);
    }
    public Optional<String> editionType() {
        return Optional.ofNullable(this.editionType);
    }
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }
    /**
     * @return The engine version of the instance.
     * 
     */
    public Optional<String> engineVersion() {
        return Optional.ofNullable(this.engineVersion);
    }
    public Optional<String> expired() {
        return Optional.ofNullable(this.expired);
    }
    public Optional<Boolean> globalInstance() {
        return Optional.ofNullable(this.globalInstance);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of KVStore Instance IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    /**
     * @return (Optional) Type of the applied ApsaraDB for instance.
     * For more information, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/61135.htm).
     * 
     */
    public Optional<String> instanceClass() {
        return Optional.ofNullable(this.instanceClass);
    }
    /**
     * @return (Optional) Database type. Valid Values: `Memcache`, `Redis`. If no value is specified, all types are returned.
     * 
     */
    public Optional<String> instanceType() {
        return Optional.ofNullable(this.instanceType);
    }
    /**
     * @return A list of KVStore Instances. Its every element contains the following attributes:
     * 
     */
    public List<GetInstancesInstance> instances() {
        return this.instances;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of KVStore Instance names.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    /**
     * @return The network type of the instance.
     * 
     */
    public Optional<String> networkType() {
        return Optional.ofNullable(this.networkType);
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return Billing method. Valid Values: `PostPaid` for  Pay-As-You-Go and `PrePaid` for subscription.
     * 
     */
    public Optional<String> paymentType() {
        return Optional.ofNullable(this.paymentType);
    }
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }
    public Optional<String> searchKey() {
        return Optional.ofNullable(this.searchKey);
    }
    /**
     * @return Status of the instance.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    public Map<String,Object> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }
    /**
     * @return VPC ID the instance belongs to.
     * 
     */
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }
    /**
     * @return VSwitch ID the instance belongs to.
     * 
     */
    public Optional<String> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }
    /**
     * @return The ID of zone.
     * 
     */
    public Optional<String> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstancesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String architectureType;
        private @Nullable String editionType;
        private @Nullable Boolean enableDetails;
        private @Nullable String engineVersion;
        private @Nullable String expired;
        private @Nullable Boolean globalInstance;
        private String id;
        private List<String> ids;
        private @Nullable String instanceClass;
        private @Nullable String instanceType;
        private List<GetInstancesInstance> instances;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String networkType;
        private @Nullable String outputFile;
        private @Nullable String paymentType;
        private @Nullable String resourceGroupId;
        private @Nullable String searchKey;
        private @Nullable String status;
        private @Nullable Map<String,Object> tags;
        private @Nullable String vpcId;
        private @Nullable String vswitchId;
        private @Nullable String zoneId;
        public Builder() {}
        public Builder(GetInstancesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.architectureType = defaults.architectureType;
    	      this.editionType = defaults.editionType;
    	      this.enableDetails = defaults.enableDetails;
    	      this.engineVersion = defaults.engineVersion;
    	      this.expired = defaults.expired;
    	      this.globalInstance = defaults.globalInstance;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.instanceClass = defaults.instanceClass;
    	      this.instanceType = defaults.instanceType;
    	      this.instances = defaults.instances;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.networkType = defaults.networkType;
    	      this.outputFile = defaults.outputFile;
    	      this.paymentType = defaults.paymentType;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.searchKey = defaults.searchKey;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.vpcId = defaults.vpcId;
    	      this.vswitchId = defaults.vswitchId;
    	      this.zoneId = defaults.zoneId;
        }

        @CustomType.Setter
        public Builder architectureType(@Nullable String architectureType) {
            this.architectureType = architectureType;
            return this;
        }
        @CustomType.Setter
        public Builder editionType(@Nullable String editionType) {
            this.editionType = editionType;
            return this;
        }
        @CustomType.Setter
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            this.enableDetails = enableDetails;
            return this;
        }
        @CustomType.Setter
        public Builder engineVersion(@Nullable String engineVersion) {
            this.engineVersion = engineVersion;
            return this;
        }
        @CustomType.Setter
        public Builder expired(@Nullable String expired) {
            this.expired = expired;
            return this;
        }
        @CustomType.Setter
        public Builder globalInstance(@Nullable Boolean globalInstance) {
            this.globalInstance = globalInstance;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ids(List<String> ids) {
            this.ids = Objects.requireNonNull(ids);
            return this;
        }
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }
        @CustomType.Setter
        public Builder instanceClass(@Nullable String instanceClass) {
            this.instanceClass = instanceClass;
            return this;
        }
        @CustomType.Setter
        public Builder instanceType(@Nullable String instanceType) {
            this.instanceType = instanceType;
            return this;
        }
        @CustomType.Setter
        public Builder instances(List<GetInstancesInstance> instances) {
            this.instances = Objects.requireNonNull(instances);
            return this;
        }
        public Builder instances(GetInstancesInstance... instances) {
            return instances(List.of(instances));
        }
        @CustomType.Setter
        public Builder nameRegex(@Nullable String nameRegex) {
            this.nameRegex = nameRegex;
            return this;
        }
        @CustomType.Setter
        public Builder names(List<String> names) {
            this.names = Objects.requireNonNull(names);
            return this;
        }
        public Builder names(String... names) {
            return names(List.of(names));
        }
        @CustomType.Setter
        public Builder networkType(@Nullable String networkType) {
            this.networkType = networkType;
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder paymentType(@Nullable String paymentType) {
            this.paymentType = paymentType;
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            this.resourceGroupId = resourceGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder searchKey(@Nullable String searchKey) {
            this.searchKey = searchKey;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,Object> tags) {
            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(@Nullable String vpcId) {
            this.vpcId = vpcId;
            return this;
        }
        @CustomType.Setter
        public Builder vswitchId(@Nullable String vswitchId) {
            this.vswitchId = vswitchId;
            return this;
        }
        @CustomType.Setter
        public Builder zoneId(@Nullable String zoneId) {
            this.zoneId = zoneId;
            return this;
        }
        public GetInstancesResult build() {
            final var o = new GetInstancesResult();
            o.architectureType = architectureType;
            o.editionType = editionType;
            o.enableDetails = enableDetails;
            o.engineVersion = engineVersion;
            o.expired = expired;
            o.globalInstance = globalInstance;
            o.id = id;
            o.ids = ids;
            o.instanceClass = instanceClass;
            o.instanceType = instanceType;
            o.instances = instances;
            o.nameRegex = nameRegex;
            o.names = names;
            o.networkType = networkType;
            o.outputFile = outputFile;
            o.paymentType = paymentType;
            o.resourceGroupId = resourceGroupId;
            o.searchKey = searchKey;
            o.status = status;
            o.tags = tags;
            o.vpcId = vpcId;
            o.vswitchId = vswitchId;
            o.zoneId = zoneId;
            return o;
        }
    }
}
