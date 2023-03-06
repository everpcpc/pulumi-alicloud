// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cassandra.outputs;

import com.pulumi.alicloud.cassandra.outputs.GetClustersCluster;
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
     * @return A list of Cassandra clusters. Its every element contains the following attributes:
     * 
     */
    private List<GetClustersCluster> clusters;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The list of Cassandra cluster ids.
     * 
     */
    private List<String> ids;
    private @Nullable String nameRegex;
    /**
     * @return The name list of Cassandra clusters.
     * 
     */
    private List<String> names;
    private @Nullable String outputFile;
    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    private @Nullable Map<String,Object> tags;

    private GetClustersResult() {}
    /**
     * @return A list of Cassandra clusters. Its every element contains the following attributes:
     * 
     */
    public List<GetClustersCluster> clusters() {
        return this.clusters;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The list of Cassandra cluster ids.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return The name list of Cassandra clusters.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
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
        private String id;
        private List<String> ids;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable Map<String,Object> tags;
        public Builder() {}
        public Builder(GetClustersResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusters = defaults.clusters;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
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
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
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
            o.id = id;
            o.ids = ids;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.tags = tags;
            return o;
        }
    }
}