// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dds.outputs;

import com.pulumi.alicloud.dds.outputs.GetMongoInstancesInstanceMongo;
import com.pulumi.alicloud.dds.outputs.GetMongoInstancesInstanceShard;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetMongoInstancesInstance {
    private String availabilityZone;
    private String chargeType;
    private String creationTime;
    private String engine;
    private String engineVersion;
    private String expirationTime;
    private String id;
    private String instanceClass;
    private String instanceType;
    private String lockMode;
    private List<GetMongoInstancesInstanceMongo> mongos;
    private String name;
    private String networkType;
    private String regionId;
    private String replication;
    private List<GetMongoInstancesInstanceShard> shards;
    private String status;
    private Integer storage;
    private Map<String,Object> tags;

    private GetMongoInstancesInstance() {}
    public String availabilityZone() {
        return this.availabilityZone;
    }
    public String chargeType() {
        return this.chargeType;
    }
    public String creationTime() {
        return this.creationTime;
    }
    public String engine() {
        return this.engine;
    }
    public String engineVersion() {
        return this.engineVersion;
    }
    public String expirationTime() {
        return this.expirationTime;
    }
    public String id() {
        return this.id;
    }
    public String instanceClass() {
        return this.instanceClass;
    }
    public String instanceType() {
        return this.instanceType;
    }
    public String lockMode() {
        return this.lockMode;
    }
    public List<GetMongoInstancesInstanceMongo> mongos() {
        return this.mongos;
    }
    public String name() {
        return this.name;
    }
    public String networkType() {
        return this.networkType;
    }
    public String regionId() {
        return this.regionId;
    }
    public String replication() {
        return this.replication;
    }
    public List<GetMongoInstancesInstanceShard> shards() {
        return this.shards;
    }
    public String status() {
        return this.status;
    }
    public Integer storage() {
        return this.storage;
    }
    public Map<String,Object> tags() {
        return this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMongoInstancesInstance defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String availabilityZone;
        private String chargeType;
        private String creationTime;
        private String engine;
        private String engineVersion;
        private String expirationTime;
        private String id;
        private String instanceClass;
        private String instanceType;
        private String lockMode;
        private List<GetMongoInstancesInstanceMongo> mongos;
        private String name;
        private String networkType;
        private String regionId;
        private String replication;
        private List<GetMongoInstancesInstanceShard> shards;
        private String status;
        private Integer storage;
        private Map<String,Object> tags;
        public Builder() {}
        public Builder(GetMongoInstancesInstance defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.availabilityZone = defaults.availabilityZone;
    	      this.chargeType = defaults.chargeType;
    	      this.creationTime = defaults.creationTime;
    	      this.engine = defaults.engine;
    	      this.engineVersion = defaults.engineVersion;
    	      this.expirationTime = defaults.expirationTime;
    	      this.id = defaults.id;
    	      this.instanceClass = defaults.instanceClass;
    	      this.instanceType = defaults.instanceType;
    	      this.lockMode = defaults.lockMode;
    	      this.mongos = defaults.mongos;
    	      this.name = defaults.name;
    	      this.networkType = defaults.networkType;
    	      this.regionId = defaults.regionId;
    	      this.replication = defaults.replication;
    	      this.shards = defaults.shards;
    	      this.status = defaults.status;
    	      this.storage = defaults.storage;
    	      this.tags = defaults.tags;
        }

        @CustomType.Setter
        public Builder availabilityZone(String availabilityZone) {
            this.availabilityZone = Objects.requireNonNull(availabilityZone);
            return this;
        }
        @CustomType.Setter
        public Builder chargeType(String chargeType) {
            this.chargeType = Objects.requireNonNull(chargeType);
            return this;
        }
        @CustomType.Setter
        public Builder creationTime(String creationTime) {
            this.creationTime = Objects.requireNonNull(creationTime);
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
        public Builder expirationTime(String expirationTime) {
            this.expirationTime = Objects.requireNonNull(expirationTime);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceClass(String instanceClass) {
            this.instanceClass = Objects.requireNonNull(instanceClass);
            return this;
        }
        @CustomType.Setter
        public Builder instanceType(String instanceType) {
            this.instanceType = Objects.requireNonNull(instanceType);
            return this;
        }
        @CustomType.Setter
        public Builder lockMode(String lockMode) {
            this.lockMode = Objects.requireNonNull(lockMode);
            return this;
        }
        @CustomType.Setter
        public Builder mongos(List<GetMongoInstancesInstanceMongo> mongos) {
            this.mongos = Objects.requireNonNull(mongos);
            return this;
        }
        public Builder mongos(GetMongoInstancesInstanceMongo... mongos) {
            return mongos(List.of(mongos));
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
        public Builder regionId(String regionId) {
            this.regionId = Objects.requireNonNull(regionId);
            return this;
        }
        @CustomType.Setter
        public Builder replication(String replication) {
            this.replication = Objects.requireNonNull(replication);
            return this;
        }
        @CustomType.Setter
        public Builder shards(List<GetMongoInstancesInstanceShard> shards) {
            this.shards = Objects.requireNonNull(shards);
            return this;
        }
        public Builder shards(GetMongoInstancesInstanceShard... shards) {
            return shards(List.of(shards));
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder storage(Integer storage) {
            this.storage = Objects.requireNonNull(storage);
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,Object> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public GetMongoInstancesInstance build() {
            final var o = new GetMongoInstancesInstance();
            o.availabilityZone = availabilityZone;
            o.chargeType = chargeType;
            o.creationTime = creationTime;
            o.engine = engine;
            o.engineVersion = engineVersion;
            o.expirationTime = expirationTime;
            o.id = id;
            o.instanceClass = instanceClass;
            o.instanceType = instanceType;
            o.lockMode = lockMode;
            o.mongos = mongos;
            o.name = name;
            o.networkType = networkType;
            o.regionId = regionId;
            o.replication = replication;
            o.shards = shards;
            o.status = status;
            o.storage = storage;
            o.tags = tags;
            return o;
        }
    }
}
