// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cr.outputs;

import com.pulumi.alicloud.cr.outputs.GetReposRepo;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetReposResult {
    private @Nullable Boolean enableDetails;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of matched Container Registry Repositories. Its element is set to `names`.
     * 
     */
    private List<String> ids;
    private @Nullable String nameRegex;
    /**
     * @return A list of repository names.
     * 
     */
    private List<String> names;
    /**
     * @return Name of container registry namespace where repo is located.
     * 
     */
    private @Nullable String namespace;
    private @Nullable String outputFile;
    /**
     * @return A list of matched Container Registry Namespaces. Each element contains the following attributes:
     * 
     */
    private List<GetReposRepo> repos;

    private GetReposResult() {}
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of matched Container Registry Repositories. Its element is set to `names`.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of repository names.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    /**
     * @return Name of container registry namespace where repo is located.
     * 
     */
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return A list of matched Container Registry Namespaces. Each element contains the following attributes:
     * 
     */
    public List<GetReposRepo> repos() {
        return this.repos;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetReposResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enableDetails;
        private String id;
        private List<String> ids;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String namespace;
        private @Nullable String outputFile;
        private List<GetReposRepo> repos;
        public Builder() {}
        public Builder(GetReposResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enableDetails = defaults.enableDetails;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.namespace = defaults.namespace;
    	      this.outputFile = defaults.outputFile;
    	      this.repos = defaults.repos;
        }

        @CustomType.Setter
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            this.enableDetails = enableDetails;
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
        public Builder namespace(@Nullable String namespace) {
            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder repos(List<GetReposRepo> repos) {
            this.repos = Objects.requireNonNull(repos);
            return this;
        }
        public Builder repos(GetReposRepo... repos) {
            return repos(List.of(repos));
        }
        public GetReposResult build() {
            final var o = new GetReposResult();
            o.enableDetails = enableDetails;
            o.id = id;
            o.ids = ids;
            o.nameRegex = nameRegex;
            o.names = names;
            o.namespace = namespace;
            o.outputFile = outputFile;
            o.repos = repos;
            return o;
        }
    }
}
