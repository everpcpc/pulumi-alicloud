// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb.outputs;

import com.pulumi.alicloud.slb.outputs.GetCaCertificatesCertificate;
import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetCaCertificatesResult {
    /**
     * @return A list of SLB ca certificates. Each element contains the following attributes:
     * 
     */
    private List<GetCaCertificatesCertificate> certificates;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of SLB ca certificates IDs.
     * 
     */
    private List<String> ids;
    private @Nullable String nameRegex;
    /**
     * @return A list of SLB ca certificates names.
     * 
     */
    private List<String> names;
    private @Nullable String outputFile;
    /**
     * @return The resource group Id of CA certificate.
     * 
     */
    private @Nullable String resourceGroupId;
    /**
     * @return (Available in v1.66.0+) A mapping of tags to assign to the resource.
     * 
     */
    private @Nullable Map<String,Object> tags;

    private GetCaCertificatesResult() {}
    /**
     * @return A list of SLB ca certificates. Each element contains the following attributes:
     * 
     */
    public List<GetCaCertificatesCertificate> certificates() {
        return this.certificates;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of SLB ca certificates IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of SLB ca certificates names.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return The resource group Id of CA certificate.
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }
    /**
     * @return (Available in v1.66.0+) A mapping of tags to assign to the resource.
     * 
     */
    public Map<String,Object> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCaCertificatesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetCaCertificatesCertificate> certificates;
        private String id;
        private List<String> ids;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable String resourceGroupId;
        private @Nullable Map<String,Object> tags;
        public Builder() {}
        public Builder(GetCaCertificatesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.certificates = defaults.certificates;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.tags = defaults.tags;
        }

        @CustomType.Setter
        public Builder certificates(List<GetCaCertificatesCertificate> certificates) {
            this.certificates = Objects.requireNonNull(certificates);
            return this;
        }
        public Builder certificates(GetCaCertificatesCertificate... certificates) {
            return certificates(List.of(certificates));
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
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            this.resourceGroupId = resourceGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,Object> tags) {
            this.tags = tags;
            return this;
        }
        public GetCaCertificatesResult build() {
            final var o = new GetCaCertificatesResult();
            o.certificates = certificates;
            o.id = id;
            o.ids = ids;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.resourceGroupId = resourceGroupId;
            o.tags = tags;
            return o;
        }
    }
}
