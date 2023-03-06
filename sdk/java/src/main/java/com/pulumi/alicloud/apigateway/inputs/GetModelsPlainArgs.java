// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.apigateway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetModelsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetModelsPlainArgs Empty = new GetModelsPlainArgs();

    /**
     * The ID of the api group.
     * 
     */
    @Import(name="groupId", required=true)
    private String groupId;

    /**
     * @return The ID of the api group.
     * 
     */
    public String groupId() {
        return this.groupId;
    }

    /**
     * A list of Model IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Model IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The name of the Model.
     * 
     */
    @Import(name="modelName")
    private @Nullable String modelName;

    /**
     * @return The name of the Model.
     * 
     */
    public Optional<String> modelName() {
        return Optional.ofNullable(this.modelName);
    }

    /**
     * A regex string to filter results by Model name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Model name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    @Import(name="outputFile")
    private @Nullable String outputFile;

    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    @Import(name="pageNumber")
    private @Nullable Integer pageNumber;

    public Optional<Integer> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }

    @Import(name="pageSize")
    private @Nullable Integer pageSize;

    public Optional<Integer> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }

    private GetModelsPlainArgs() {}

    private GetModelsPlainArgs(GetModelsPlainArgs $) {
        this.groupId = $.groupId;
        this.ids = $.ids;
        this.modelName = $.modelName;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetModelsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetModelsPlainArgs $;

        public Builder() {
            $ = new GetModelsPlainArgs();
        }

        public Builder(GetModelsPlainArgs defaults) {
            $ = new GetModelsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param groupId The ID of the api group.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param ids A list of Model IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Model IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param modelName The name of the Model.
         * 
         * @return builder
         * 
         */
        public Builder modelName(@Nullable String modelName) {
            $.modelName = modelName;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Model name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder pageNumber(@Nullable Integer pageNumber) {
            $.pageNumber = pageNumber;
            return this;
        }

        public Builder pageSize(@Nullable Integer pageSize) {
            $.pageSize = pageSize;
            return this;
        }

        public GetModelsPlainArgs build() {
            $.groupId = Objects.requireNonNull($.groupId, "expected parameter 'groupId' to be non-null");
            return $;
        }
    }

}