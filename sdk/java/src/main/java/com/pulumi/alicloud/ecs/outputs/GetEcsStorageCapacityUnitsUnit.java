// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetEcsStorageCapacityUnitsUnit {
    /**
     * @return When the AllocationType value is Shared, this parameter indicates the allocation status of Storage Capacity Unit. Valid values: `allocated`, `Ignored`.
     * 
     */
    private String allocationStatus;
    /**
     * @return The capacity of the Storage Capacity Unit.
     * 
     */
    private Integer capacity;
    /**
     * @return The time when the Storage Capacity Unit was created.
     * 
     */
    private String createTime;
    /**
     * @return The description of the Storage Capacity Unit.
     * 
     */
    private String description;
    /**
     * @return The time when the Storage Capacity Unit expires.
     * 
     */
    private String expiredTime;
    /**
     * @return The ID of the Storage Capacity Unit.
     * 
     */
    private String id;
    /**
     * @return The effective time of the Storage Capacity Unit.
     * 
     */
    private String startTime;
    /**
     * @return The status of Storage Capacity Unit.
     * 
     */
    private String status;
    /**
     * @return The ID of Storage Capacity Unit.
     * 
     */
    private String storageCapacityUnitId;
    /**
     * @return The name of the Storage Capacity Unit.
     * 
     */
    private String storageCapacityUnitName;

    private GetEcsStorageCapacityUnitsUnit() {}
    /**
     * @return When the AllocationType value is Shared, this parameter indicates the allocation status of Storage Capacity Unit. Valid values: `allocated`, `Ignored`.
     * 
     */
    public String allocationStatus() {
        return this.allocationStatus;
    }
    /**
     * @return The capacity of the Storage Capacity Unit.
     * 
     */
    public Integer capacity() {
        return this.capacity;
    }
    /**
     * @return The time when the Storage Capacity Unit was created.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The description of the Storage Capacity Unit.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The time when the Storage Capacity Unit expires.
     * 
     */
    public String expiredTime() {
        return this.expiredTime;
    }
    /**
     * @return The ID of the Storage Capacity Unit.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The effective time of the Storage Capacity Unit.
     * 
     */
    public String startTime() {
        return this.startTime;
    }
    /**
     * @return The status of Storage Capacity Unit.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The ID of Storage Capacity Unit.
     * 
     */
    public String storageCapacityUnitId() {
        return this.storageCapacityUnitId;
    }
    /**
     * @return The name of the Storage Capacity Unit.
     * 
     */
    public String storageCapacityUnitName() {
        return this.storageCapacityUnitName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEcsStorageCapacityUnitsUnit defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String allocationStatus;
        private Integer capacity;
        private String createTime;
        private String description;
        private String expiredTime;
        private String id;
        private String startTime;
        private String status;
        private String storageCapacityUnitId;
        private String storageCapacityUnitName;
        public Builder() {}
        public Builder(GetEcsStorageCapacityUnitsUnit defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allocationStatus = defaults.allocationStatus;
    	      this.capacity = defaults.capacity;
    	      this.createTime = defaults.createTime;
    	      this.description = defaults.description;
    	      this.expiredTime = defaults.expiredTime;
    	      this.id = defaults.id;
    	      this.startTime = defaults.startTime;
    	      this.status = defaults.status;
    	      this.storageCapacityUnitId = defaults.storageCapacityUnitId;
    	      this.storageCapacityUnitName = defaults.storageCapacityUnitName;
        }

        @CustomType.Setter
        public Builder allocationStatus(String allocationStatus) {
            this.allocationStatus = Objects.requireNonNull(allocationStatus);
            return this;
        }
        @CustomType.Setter
        public Builder capacity(Integer capacity) {
            this.capacity = Objects.requireNonNull(capacity);
            return this;
        }
        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder expiredTime(String expiredTime) {
            this.expiredTime = Objects.requireNonNull(expiredTime);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
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
        public Builder storageCapacityUnitId(String storageCapacityUnitId) {
            this.storageCapacityUnitId = Objects.requireNonNull(storageCapacityUnitId);
            return this;
        }
        @CustomType.Setter
        public Builder storageCapacityUnitName(String storageCapacityUnitName) {
            this.storageCapacityUnitName = Objects.requireNonNull(storageCapacityUnitName);
            return this;
        }
        public GetEcsStorageCapacityUnitsUnit build() {
            final var o = new GetEcsStorageCapacityUnitsUnit();
            o.allocationStatus = allocationStatus;
            o.capacity = capacity;
            o.createTime = createTime;
            o.description = description;
            o.expiredTime = expiredTime;
            o.id = id;
            o.startTime = startTime;
            o.status = status;
            o.storageCapacityUnitId = storageCapacityUnitId;
            o.storageCapacityUnitName = storageCapacityUnitName;
            return o;
        }
    }
}
