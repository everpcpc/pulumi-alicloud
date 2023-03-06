// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.yundun.outputs;

import com.pulumi.alicloud.yundun.outputs.GetBastionHostInstancesInstance;
import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBastionHostInstancesResult {
    private @Nullable String descriptionRegex;
    private List<String> descriptions;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private List<GetBastionHostInstancesInstance> instances;
    private @Nullable String outputFile;
    private @Nullable Map<String,Object> tags;

    private GetBastionHostInstancesResult() {}
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
    public List<GetBastionHostInstancesInstance> instances() {
        return this.instances;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Map<String,Object> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBastionHostInstancesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String descriptionRegex;
        private List<String> descriptions;
        private String id;
        private List<String> ids;
        private List<GetBastionHostInstancesInstance> instances;
        private @Nullable String outputFile;
        private @Nullable Map<String,Object> tags;
        public Builder() {}
        public Builder(GetBastionHostInstancesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.descriptionRegex = defaults.descriptionRegex;
    	      this.descriptions = defaults.descriptions;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.instances = defaults.instances;
    	      this.outputFile = defaults.outputFile;
    	      this.tags = defaults.tags;
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
        public Builder instances(List<GetBastionHostInstancesInstance> instances) {
            this.instances = Objects.requireNonNull(instances);
            return this;
        }
        public Builder instances(GetBastionHostInstancesInstance... instances) {
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
        public GetBastionHostInstancesResult build() {
            final var o = new GetBastionHostInstancesResult();
            o.descriptionRegex = descriptionRegex;
            o.descriptions = descriptions;
            o.id = id;
            o.ids = ids;
            o.instances = instances;
            o.outputFile = outputFile;
            o.tags = tags;
            return o;
        }
    }
}