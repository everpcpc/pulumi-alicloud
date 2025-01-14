// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cs.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRegistryEnterpriseReposPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRegistryEnterpriseReposPlainArgs Empty = new GetRegistryEnterpriseReposPlainArgs();

    /**
     * Boolean, false by default, only repository attributes are exported. Set to true if tags belong to this repository are needed. See `tags` in attributes.
     * 
     */
    @Import(name="enableDetails")
    private @Nullable Boolean enableDetails;

    /**
     * @return Boolean, false by default, only repository attributes are exported. Set to true if tags belong to this repository are needed. See `tags` in attributes.
     * 
     */
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * A list of ids to filter results by repository id.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of ids to filter results by repository id.
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
     * A regex string to filter results by repository name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by repository name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * Name of Container Registry Enterprise Edition namespace where the repositories are located in.
     * 
     */
    @Import(name="namespace")
    private @Nullable String namespace;

    /**
     * @return Name of Container Registry Enterprise Edition namespace where the repositories are located in.
     * 
     */
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
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

    private GetRegistryEnterpriseReposPlainArgs() {}

    private GetRegistryEnterpriseReposPlainArgs(GetRegistryEnterpriseReposPlainArgs $) {
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.instanceId = $.instanceId;
        this.nameRegex = $.nameRegex;
        this.namespace = $.namespace;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRegistryEnterpriseReposPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRegistryEnterpriseReposPlainArgs $;

        public Builder() {
            $ = new GetRegistryEnterpriseReposPlainArgs();
        }

        public Builder(GetRegistryEnterpriseReposPlainArgs defaults) {
            $ = new GetRegistryEnterpriseReposPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enableDetails Boolean, false by default, only repository attributes are exported. Set to true if tags belong to this repository are needed. See `tags` in attributes.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param ids A list of ids to filter results by repository id.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of ids to filter results by repository id.
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
         * @param nameRegex A regex string to filter results by repository name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param namespace Name of Container Registry Enterprise Edition namespace where the repositories are located in.
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable String namespace) {
            $.namespace = namespace;
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

        public GetRegistryEnterpriseReposPlainArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
