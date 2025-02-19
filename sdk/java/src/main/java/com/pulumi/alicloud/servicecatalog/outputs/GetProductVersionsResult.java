// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.servicecatalog.outputs;

import com.pulumi.alicloud.servicecatalog.outputs.GetProductVersionsProductVersion;
import com.pulumi.alicloud.servicecatalog.outputs.GetProductVersionsVersion;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetProductVersionsResult {
    private @Nullable Boolean enableDetails;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of Product Version IDs.
     * 
     */
    private List<String> ids;
    private @Nullable String nameRegex;
    /**
     * @return A list of name of Product Versions.
     * 
     */
    private List<String> names;
    private @Nullable String outputFile;
    private String productId;
    /**
     * @return A list of Product Version Entries. Each element contains the following attributes:
     * 
     */
    private List<GetProductVersionsProductVersion> productVersions;
    /**
     * @deprecated
     * Field &#39;versions&#39; has been deprecated from provider version 1.197.0.
     * 
     */
    @Deprecated /* Field 'versions' has been deprecated from provider version 1.197.0. */
    private List<GetProductVersionsVersion> versions;

    private GetProductVersionsResult() {}
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
     * @return A list of Product Version IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of name of Product Versions.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public String productId() {
        return this.productId;
    }
    /**
     * @return A list of Product Version Entries. Each element contains the following attributes:
     * 
     */
    public List<GetProductVersionsProductVersion> productVersions() {
        return this.productVersions;
    }
    /**
     * @deprecated
     * Field &#39;versions&#39; has been deprecated from provider version 1.197.0.
     * 
     */
    @Deprecated /* Field 'versions' has been deprecated from provider version 1.197.0. */
    public List<GetProductVersionsVersion> versions() {
        return this.versions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProductVersionsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enableDetails;
        private String id;
        private List<String> ids;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private String productId;
        private List<GetProductVersionsProductVersion> productVersions;
        private List<GetProductVersionsVersion> versions;
        public Builder() {}
        public Builder(GetProductVersionsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enableDetails = defaults.enableDetails;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.productId = defaults.productId;
    	      this.productVersions = defaults.productVersions;
    	      this.versions = defaults.versions;
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
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder productId(String productId) {
            this.productId = Objects.requireNonNull(productId);
            return this;
        }
        @CustomType.Setter
        public Builder productVersions(List<GetProductVersionsProductVersion> productVersions) {
            this.productVersions = Objects.requireNonNull(productVersions);
            return this;
        }
        public Builder productVersions(GetProductVersionsProductVersion... productVersions) {
            return productVersions(List.of(productVersions));
        }
        @CustomType.Setter
        public Builder versions(List<GetProductVersionsVersion> versions) {
            this.versions = Objects.requireNonNull(versions);
            return this;
        }
        public Builder versions(GetProductVersionsVersion... versions) {
            return versions(List.of(versions));
        }
        public GetProductVersionsResult build() {
            final var o = new GetProductVersionsResult();
            o.enableDetails = enableDetails;
            o.id = id;
            o.ids = ids;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.productId = productId;
            o.productVersions = productVersions;
            o.versions = versions;
            return o;
        }
    }
}
