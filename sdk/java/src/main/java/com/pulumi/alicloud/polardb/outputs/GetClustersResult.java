// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb.outputs;

import com.pulumi.alicloud.polardb.outputs.GetClustersCluster;
import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetClustersResult {
    /**
     * @return A list of PolarDB clusters. Each element contains the following attributes:
     * 
     */
    private List<GetClustersCluster> clusters;
    /**
     * @return `Primary` for primary cluster, `ReadOnly` for read-only cluster, `Guard` for disaster recovery cluster, and `Temp` for temporary cluster.
     * 
     */
    private @Nullable String dbType;
    private @Nullable String descriptionRegex;
    /**
     * @return A list of RDS cluster descriptions.
     * 
     */
    private List<String> descriptions;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of RDS cluster IDs.
     * 
     */
    private List<String> ids;
    private @Nullable String outputFile;
    /**
     * @return Status of the cluster.
     * 
     */
    private @Nullable String status;
    private @Nullable Map<String,Object> tags;

    private GetClustersResult() {}
    /**
     * @return A list of PolarDB clusters. Each element contains the following attributes:
     * 
     */
    public List<GetClustersCluster> clusters() {
        return this.clusters;
    }
    /**
     * @return `Primary` for primary cluster, `ReadOnly` for read-only cluster, `Guard` for disaster recovery cluster, and `Temp` for temporary cluster.
     * 
     */
    public Optional<String> dbType() {
        return Optional.ofNullable(this.dbType);
    }
    public Optional<String> descriptionRegex() {
        return Optional.ofNullable(this.descriptionRegex);
    }
    /**
     * @return A list of RDS cluster descriptions.
     * 
     */
    public List<String> descriptions() {
        return this.descriptions;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of RDS cluster IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return Status of the cluster.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    public Map<String,Object> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClustersResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetClustersCluster> clusters;
        private @Nullable String dbType;
        private @Nullable String descriptionRegex;
        private List<String> descriptions;
        private String id;
        private List<String> ids;
        private @Nullable String outputFile;
        private @Nullable String status;
        private @Nullable Map<String,Object> tags;
        public Builder() {}
        public Builder(GetClustersResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusters = defaults.clusters;
    	      this.dbType = defaults.dbType;
    	      this.descriptionRegex = defaults.descriptionRegex;
    	      this.descriptions = defaults.descriptions;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.outputFile = defaults.outputFile;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
        }

        @CustomType.Setter
        public Builder clusters(List<GetClustersCluster> clusters) {
            this.clusters = Objects.requireNonNull(clusters);
            return this;
        }
        public Builder clusters(GetClustersCluster... clusters) {
            return clusters(List.of(clusters));
        }
        @CustomType.Setter
        public Builder dbType(@Nullable String dbType) {
            this.dbType = dbType;
            return this;
        }
        @CustomType.Setter
        public Builder descriptionRegex(@Nullable String descriptionRegex) {
            this.descriptionRegex = descriptionRegex;
            return this;
        }
        @CustomType.Setter
        public Builder descriptions(List<String> descriptions) {
            this.descriptions = Objects.requireNonNull(descriptions);
            return this;
        }
        public Builder descriptions(String... descriptions) {
            return descriptions(List.of(descriptions));
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
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
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
        public GetClustersResult build() {
            final var o = new GetClustersResult();
            o.clusters = clusters;
            o.dbType = dbType;
            o.descriptionRegex = descriptionRegex;
            o.descriptions = descriptions;
            o.id = id;
            o.ids = ids;
            o.outputFile = outputFile;
            o.status = status;
            o.tags = tags;
            return o;
        }
    }
}
