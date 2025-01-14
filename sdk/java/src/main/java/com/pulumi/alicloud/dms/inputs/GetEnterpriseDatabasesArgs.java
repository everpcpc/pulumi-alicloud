// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dms.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEnterpriseDatabasesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEnterpriseDatabasesArgs Empty = new GetEnterpriseDatabasesArgs();

    /**
     * A list of Database IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Database IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The instance ID of the target database.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return The instance ID of the target database.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     * A regex string to filter the results by the database Schema Name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter the results by the database Schema Name.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    private GetEnterpriseDatabasesArgs() {}

    private GetEnterpriseDatabasesArgs(GetEnterpriseDatabasesArgs $) {
        this.ids = $.ids;
        this.instanceId = $.instanceId;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEnterpriseDatabasesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEnterpriseDatabasesArgs $;

        public Builder() {
            $ = new GetEnterpriseDatabasesArgs();
        }

        public Builder(GetEnterpriseDatabasesArgs defaults) {
            $ = new GetEnterpriseDatabasesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Database IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Database IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Database IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceId The instance ID of the target database.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The instance ID of the target database.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param nameRegex A regex string to filter the results by the database Schema Name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter the results by the database Schema Name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        public GetEnterpriseDatabasesArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
