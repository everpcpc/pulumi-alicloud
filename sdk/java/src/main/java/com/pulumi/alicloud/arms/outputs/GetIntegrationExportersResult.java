// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms.outputs;

import com.pulumi.alicloud.arms.outputs.GetIntegrationExportersIntegrationExporter;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetIntegrationExportersResult {
    /**
     * @return The ID of the Prometheus instance.
     * 
     */
    private String clusterId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    /**
     * @return A list of Integration Exporters. Each element contains the following attributes:
     * 
     */
    private List<GetIntegrationExportersIntegrationExporter> integrationExporters;
    /**
     * @return The type of prometheus integration.
     * 
     */
    private String integrationType;
    private @Nullable String outputFile;

    private GetIntegrationExportersResult() {}
    /**
     * @return The ID of the Prometheus instance.
     * 
     */
    public String clusterId() {
        return this.clusterId;
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
    /**
     * @return A list of Integration Exporters. Each element contains the following attributes:
     * 
     */
    public List<GetIntegrationExportersIntegrationExporter> integrationExporters() {
        return this.integrationExporters;
    }
    /**
     * @return The type of prometheus integration.
     * 
     */
    public String integrationType() {
        return this.integrationType;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIntegrationExportersResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterId;
        private String id;
        private List<String> ids;
        private List<GetIntegrationExportersIntegrationExporter> integrationExporters;
        private String integrationType;
        private @Nullable String outputFile;
        public Builder() {}
        public Builder(GetIntegrationExportersResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterId = defaults.clusterId;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.integrationExporters = defaults.integrationExporters;
    	      this.integrationType = defaults.integrationType;
    	      this.outputFile = defaults.outputFile;
        }

        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            this.clusterId = Objects.requireNonNull(clusterId);
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
        public Builder integrationExporters(List<GetIntegrationExportersIntegrationExporter> integrationExporters) {
            this.integrationExporters = Objects.requireNonNull(integrationExporters);
            return this;
        }
        public Builder integrationExporters(GetIntegrationExportersIntegrationExporter... integrationExporters) {
            return integrationExporters(List.of(integrationExporters));
        }
        @CustomType.Setter
        public Builder integrationType(String integrationType) {
            this.integrationType = Objects.requireNonNull(integrationType);
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        public GetIntegrationExportersResult build() {
            final var o = new GetIntegrationExportersResult();
            o.clusterId = clusterId;
            o.id = id;
            o.ids = ids;
            o.integrationExporters = integrationExporters;
            o.integrationType = integrationType;
            o.outputFile = outputFile;
            return o;
        }
    }
}
