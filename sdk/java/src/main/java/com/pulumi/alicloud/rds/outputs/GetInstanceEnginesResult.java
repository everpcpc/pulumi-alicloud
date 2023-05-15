// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds.outputs;

import com.pulumi.alicloud.rds.outputs.GetInstanceEnginesInstanceEngine;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetInstanceEnginesResult {
    /**
     * @return DB Instance category.
     * 
     */
    private @Nullable String category;
    private @Nullable String dbInstanceStorageType;
    /**
     * @return Database type.
     * 
     */
    private @Nullable String engine;
    /**
     * @return DB Instance version.
     * 
     */
    private @Nullable String engineVersion;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of engines.
     * 
     */
    private List<String> ids;
    private @Nullable String instanceChargeType;
    /**
     * @return A list of Rds available resource. Each element contains the following attributes:
     * 
     */
    private List<GetInstanceEnginesInstanceEngine> instanceEngines;
    private @Nullable Boolean multiZone;
    private @Nullable String outputFile;
    private @Nullable String zoneId;

    private GetInstanceEnginesResult() {}
    /**
     * @return DB Instance category.
     * 
     */
    public Optional<String> category() {
        return Optional.ofNullable(this.category);
    }
    public Optional<String> dbInstanceStorageType() {
        return Optional.ofNullable(this.dbInstanceStorageType);
    }
    /**
     * @return Database type.
     * 
     */
    public Optional<String> engine() {
        return Optional.ofNullable(this.engine);
    }
    /**
     * @return DB Instance version.
     * 
     */
    public Optional<String> engineVersion() {
        return Optional.ofNullable(this.engineVersion);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of engines.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> instanceChargeType() {
        return Optional.ofNullable(this.instanceChargeType);
    }
    /**
     * @return A list of Rds available resource. Each element contains the following attributes:
     * 
     */
    public List<GetInstanceEnginesInstanceEngine> instanceEngines() {
        return this.instanceEngines;
    }
    public Optional<Boolean> multiZone() {
        return Optional.ofNullable(this.multiZone);
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Optional<String> zoneId() {
        return Optional.ofNullable(this.zoneId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstanceEnginesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String category;
        private @Nullable String dbInstanceStorageType;
        private @Nullable String engine;
        private @Nullable String engineVersion;
        private String id;
        private List<String> ids;
        private @Nullable String instanceChargeType;
        private List<GetInstanceEnginesInstanceEngine> instanceEngines;
        private @Nullable Boolean multiZone;
        private @Nullable String outputFile;
        private @Nullable String zoneId;
        public Builder() {}
        public Builder(GetInstanceEnginesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.category = defaults.category;
    	      this.dbInstanceStorageType = defaults.dbInstanceStorageType;
    	      this.engine = defaults.engine;
    	      this.engineVersion = defaults.engineVersion;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.instanceChargeType = defaults.instanceChargeType;
    	      this.instanceEngines = defaults.instanceEngines;
    	      this.multiZone = defaults.multiZone;
    	      this.outputFile = defaults.outputFile;
    	      this.zoneId = defaults.zoneId;
        }

        @CustomType.Setter
        public Builder category(@Nullable String category) {
            this.category = category;
            return this;
        }
        @CustomType.Setter
        public Builder dbInstanceStorageType(@Nullable String dbInstanceStorageType) {
            this.dbInstanceStorageType = dbInstanceStorageType;
            return this;
        }
        @CustomType.Setter
        public Builder engine(@Nullable String engine) {
            this.engine = engine;
            return this;
        }
        @CustomType.Setter
        public Builder engineVersion(@Nullable String engineVersion) {
            this.engineVersion = engineVersion;
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
        public Builder instanceChargeType(@Nullable String instanceChargeType) {
            this.instanceChargeType = instanceChargeType;
            return this;
        }
        @CustomType.Setter
        public Builder instanceEngines(List<GetInstanceEnginesInstanceEngine> instanceEngines) {
            this.instanceEngines = Objects.requireNonNull(instanceEngines);
            return this;
        }
        public Builder instanceEngines(GetInstanceEnginesInstanceEngine... instanceEngines) {
            return instanceEngines(List.of(instanceEngines));
        }
        @CustomType.Setter
        public Builder multiZone(@Nullable Boolean multiZone) {
            this.multiZone = multiZone;
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder zoneId(@Nullable String zoneId) {
            this.zoneId = zoneId;
            return this;
        }
        public GetInstanceEnginesResult build() {
            final var o = new GetInstanceEnginesResult();
            o.category = category;
            o.dbInstanceStorageType = dbInstanceStorageType;
            o.engine = engine;
            o.engineVersion = engineVersion;
            o.id = id;
            o.ids = ids;
            o.instanceChargeType = instanceChargeType;
            o.instanceEngines = instanceEngines;
            o.multiZone = multiZone;
            o.outputFile = outputFile;
            o.zoneId = zoneId;
            return o;
        }
    }
}
