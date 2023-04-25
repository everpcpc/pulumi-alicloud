// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIntegrationExportersPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIntegrationExportersPlainArgs Empty = new GetIntegrationExportersPlainArgs();

    /**
     * The ID of the Prometheus instance.
     * 
     */
    @Import(name="clusterId", required=true)
    private String clusterId;

    /**
     * @return The ID of the Prometheus instance.
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }

    /**
     * A list of Integration Exporter IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Integration Exporter IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The type of prometheus integration.
     * 
     */
    @Import(name="integrationType", required=true)
    private String integrationType;

    /**
     * @return The type of prometheus integration.
     * 
     */
    public String integrationType() {
        return this.integrationType;
    }

    @Import(name="outputFile")
    private @Nullable String outputFile;

    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    private GetIntegrationExportersPlainArgs() {}

    private GetIntegrationExportersPlainArgs(GetIntegrationExportersPlainArgs $) {
        this.clusterId = $.clusterId;
        this.ids = $.ids;
        this.integrationType = $.integrationType;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIntegrationExportersPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIntegrationExportersPlainArgs $;

        public Builder() {
            $ = new GetIntegrationExportersPlainArgs();
        }

        public Builder(GetIntegrationExportersPlainArgs defaults) {
            $ = new GetIntegrationExportersPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId The ID of the Prometheus instance.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param ids A list of Integration Exporter IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Integration Exporter IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param integrationType The type of prometheus integration.
         * 
         * @return builder
         * 
         */
        public Builder integrationType(String integrationType) {
            $.integrationType = integrationType;
            return this;
        }

        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public GetIntegrationExportersPlainArgs build() {
            $.clusterId = Objects.requireNonNull($.clusterId, "expected parameter 'clusterId' to be non-null");
            $.integrationType = Objects.requireNonNull($.integrationType, "expected parameter 'integrationType' to be non-null");
            return $;
        }
    }

}
