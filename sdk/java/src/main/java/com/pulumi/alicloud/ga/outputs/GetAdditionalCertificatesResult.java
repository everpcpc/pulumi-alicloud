// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.outputs;

import com.pulumi.alicloud.ga.outputs.GetAdditionalCertificatesCertificate;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAdditionalCertificatesResult {
    private String acceleratorId;
    private List<GetAdditionalCertificatesCertificate> certificates;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private String listenerId;
    private @Nullable String outputFile;

    private GetAdditionalCertificatesResult() {}
    public String acceleratorId() {
        return this.acceleratorId;
    }
    public List<GetAdditionalCertificatesCertificate> certificates() {
        return this.certificates;
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
    public String listenerId() {
        return this.listenerId;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAdditionalCertificatesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String acceleratorId;
        private List<GetAdditionalCertificatesCertificate> certificates;
        private String id;
        private List<String> ids;
        private String listenerId;
        private @Nullable String outputFile;
        public Builder() {}
        public Builder(GetAdditionalCertificatesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.acceleratorId = defaults.acceleratorId;
    	      this.certificates = defaults.certificates;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.listenerId = defaults.listenerId;
    	      this.outputFile = defaults.outputFile;
        }

        @CustomType.Setter
        public Builder acceleratorId(String acceleratorId) {
            this.acceleratorId = Objects.requireNonNull(acceleratorId);
            return this;
        }
        @CustomType.Setter
        public Builder certificates(List<GetAdditionalCertificatesCertificate> certificates) {
            this.certificates = Objects.requireNonNull(certificates);
            return this;
        }
        public Builder certificates(GetAdditionalCertificatesCertificate... certificates) {
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
        public Builder listenerId(String listenerId) {
            this.listenerId = Objects.requireNonNull(listenerId);
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        public GetAdditionalCertificatesResult build() {
            final var o = new GetAdditionalCertificatesResult();
            o.acceleratorId = acceleratorId;
            o.certificates = certificates;
            o.id = id;
            o.ids = ids;
            o.listenerId = listenerId;
            o.outputFile = outputFile;
            return o;
        }
    }
}
