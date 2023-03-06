// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.kms.outputs;

import com.pulumi.alicloud.kms.outputs.GetSecretsSecret;
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
public final class GetSecretsResult {
    private @Nullable Boolean enableDetails;
    private @Nullable Boolean fetchTags;
    private @Nullable String filters;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of Kms Secret ids. The value is same as KMS secret_name.
     * 
     */
    private List<String> ids;
    private @Nullable String nameRegex;
    /**
     * @return A list of KMS Secret names.
     * 
     */
    private List<String> names;
    private @Nullable String outputFile;
    /**
     * @return A list of KMS Secrets. Each element contains the following attributes:
     * 
     */
    private List<GetSecretsSecret> secrets;
    /**
     * @return (Optional) A mapping of tags to assign to the resource.
     * 
     */
    private @Nullable Map<String,Object> tags;

    private GetSecretsResult() {}
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }
    public Optional<Boolean> fetchTags() {
        return Optional.ofNullable(this.fetchTags);
    }
    public Optional<String> filters() {
        return Optional.ofNullable(this.filters);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of Kms Secret ids. The value is same as KMS secret_name.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of KMS Secret names.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return A list of KMS Secrets. Each element contains the following attributes:
     * 
     */
    public List<GetSecretsSecret> secrets() {
        return this.secrets;
    }
    /**
     * @return (Optional) A mapping of tags to assign to the resource.
     * 
     */
    public Map<String,Object> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSecretsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enableDetails;
        private @Nullable Boolean fetchTags;
        private @Nullable String filters;
        private String id;
        private List<String> ids;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private List<GetSecretsSecret> secrets;
        private @Nullable Map<String,Object> tags;
        public Builder() {}
        public Builder(GetSecretsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enableDetails = defaults.enableDetails;
    	      this.fetchTags = defaults.fetchTags;
    	      this.filters = defaults.filters;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.secrets = defaults.secrets;
    	      this.tags = defaults.tags;
        }

        @CustomType.Setter
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            this.enableDetails = enableDetails;
            return this;
        }
        @CustomType.Setter
        public Builder fetchTags(@Nullable Boolean fetchTags) {
            this.fetchTags = fetchTags;
            return this;
        }
        @CustomType.Setter
        public Builder filters(@Nullable String filters) {
            this.filters = filters;
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
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder secrets(List<GetSecretsSecret> secrets) {
            this.secrets = Objects.requireNonNull(secrets);
            return this;
        }
        public Builder secrets(GetSecretsSecret... secrets) {
            return secrets(List.of(secrets));
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,Object> tags) {
            this.tags = tags;
            return this;
        }
        public GetSecretsResult build() {
            final var o = new GetSecretsResult();
            o.enableDetails = enableDetails;
            o.fetchTags = fetchTags;
            o.filters = filters;
            o.id = id;
            o.ids = ids;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.secrets = secrets;
            o.tags = tags;
            return o;
        }
    }
}