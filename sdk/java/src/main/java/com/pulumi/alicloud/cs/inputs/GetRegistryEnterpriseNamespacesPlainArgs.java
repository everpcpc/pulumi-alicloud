// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cs.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRegistryEnterpriseNamespacesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRegistryEnterpriseNamespacesPlainArgs Empty = new GetRegistryEnterpriseNamespacesPlainArgs();

    /**
     * A list of ids to filter results by namespace id. Each item formats as `&lt;instance_id&gt;:&lt;namespace_name&gt;`.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of ids to filter results by namespace id. Each item formats as `&lt;instance_id&gt;:&lt;namespace_name&gt;`.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * ID of Container Registry Enterprise Edition instance.
     * 
     */
    @Import(name="instanceId", required=true)
    private String instanceId;

    /**
     * @return ID of Container Registry Enterprise Edition instance.
     * 
     */
    public String instanceId() {
        return this.instanceId;
    }

    /**
     * A regex string to filter results by namespace name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by namespace name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
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

    private GetRegistryEnterpriseNamespacesPlainArgs() {}

    private GetRegistryEnterpriseNamespacesPlainArgs(GetRegistryEnterpriseNamespacesPlainArgs $) {
        this.ids = $.ids;
        this.instanceId = $.instanceId;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRegistryEnterpriseNamespacesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRegistryEnterpriseNamespacesPlainArgs $;

        public Builder() {
            $ = new GetRegistryEnterpriseNamespacesPlainArgs();
        }

        public Builder(GetRegistryEnterpriseNamespacesPlainArgs defaults) {
            $ = new GetRegistryEnterpriseNamespacesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of ids to filter results by namespace id. Each item formats as `&lt;instance_id&gt;:&lt;namespace_name&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of ids to filter results by namespace id. Each item formats as `&lt;instance_id&gt;:&lt;namespace_name&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceId ID of Container Registry Enterprise Edition instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by namespace name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
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

        public GetRegistryEnterpriseNamespacesPlainArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
