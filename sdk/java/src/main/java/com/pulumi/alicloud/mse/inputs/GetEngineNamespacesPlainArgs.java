// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mse.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEngineNamespacesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEngineNamespacesPlainArgs Empty = new GetEngineNamespacesPlainArgs();

    /**
     * The language type of the returned information. Valid values: `zh`, `en`.
     * 
     */
    @Import(name="acceptLanguage")
    private @Nullable String acceptLanguage;

    /**
     * @return The language type of the returned information. Valid values: `zh`, `en`.
     * 
     */
    public Optional<String> acceptLanguage() {
        return Optional.ofNullable(this.acceptLanguage);
    }

    /**
     * The id of the cluster.
     * 
     */
    @Import(name="clusterId", required=true)
    private String clusterId;

    /**
     * @return The id of the cluster.
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }

    /**
     * A list of Engine Namespace IDs. It is formatted to `&lt;cluster_id&gt;:&lt;namespace_id&gt;`.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Engine Namespace IDs. It is formatted to `&lt;cluster_id&gt;:&lt;namespace_id&gt;`.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    private GetEngineNamespacesPlainArgs() {}

    private GetEngineNamespacesPlainArgs(GetEngineNamespacesPlainArgs $) {
        this.acceptLanguage = $.acceptLanguage;
        this.clusterId = $.clusterId;
        this.ids = $.ids;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEngineNamespacesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEngineNamespacesPlainArgs $;

        public Builder() {
            $ = new GetEngineNamespacesPlainArgs();
        }

        public Builder(GetEngineNamespacesPlainArgs defaults) {
            $ = new GetEngineNamespacesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param acceptLanguage The language type of the returned information. Valid values: `zh`, `en`.
         * 
         * @return builder
         * 
         */
        public Builder acceptLanguage(@Nullable String acceptLanguage) {
            $.acceptLanguage = acceptLanguage;
            return this;
        }

        /**
         * @param clusterId The id of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param ids A list of Engine Namespace IDs. It is formatted to `&lt;cluster_id&gt;:&lt;namespace_id&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Engine Namespace IDs. It is formatted to `&lt;cluster_id&gt;:&lt;namespace_id&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public GetEngineNamespacesPlainArgs build() {
            $.clusterId = Objects.requireNonNull($.clusterId, "expected parameter 'clusterId' to be non-null");
            return $;
        }
    }

}
