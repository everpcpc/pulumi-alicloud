// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.elasticsearch.outputs;

import com.pulumi.alicloud.elasticsearch.outputs.GetInstancesInstance;
import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetInstancesResult {
    private @Nullable String descriptionRegex;
    private List<String> descriptions;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private List<GetInstancesInstance> instances;
    private @Nullable String outputFile;
    private @Nullable Map<String,Object> tags;
    private @Nullable String version;

    private GetInstancesResult() {}
    public Optional<String> descriptionRegex() {
        return Optional.ofNullable(this.descriptionRegex);
    }
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
    public List<String> ids() {
        return this.ids;
    }
    public List<GetInstancesInstance> instances() {
        return this.instances;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Map<String,Object> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }
    public Optional<String> version() {
        return Optional.ofNullable(this.version);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstancesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String descriptionRegex;
        private List<String> descriptions;
        private String id;
        private List<String> ids;
        private List<GetInstancesInstance> instances;
        private @Nullable String outputFile;
        private @Nullable Map<String,Object> tags;
        private @Nullable String version;
        public Builder() {}
        public Builder(GetInstancesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.descriptionRegex = defaults.descriptionRegex;
    	      this.descriptions = defaults.descriptions;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.instances = defaults.instances;
    	      this.outputFile = defaults.outputFile;
    	      this.tags = defaults.tags;
    	      this.version = defaults.version;
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
        public Builder instances(List<GetInstancesInstance> instances) {
            this.instances = Objects.requireNonNull(instances);
            return this;
        }
        public Builder instances(GetInstancesInstance... instances) {
            return instances(List.of(instances));
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
        @CustomType.Setter
        public Builder version(@Nullable String version) {
            this.version = version;
            return this;
        }
        public GetInstancesResult build() {
            final var o = new GetInstancesResult();
            o.descriptionRegex = descriptionRegex;
            o.descriptions = descriptions;
            o.id = id;
            o.ids = ids;
            o.instances = instances;
            o.outputFile = outputFile;
            o.tags = tags;
            o.version = version;
            return o;
        }
    }
}
