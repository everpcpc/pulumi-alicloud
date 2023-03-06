// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.brain.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIndustrialPidLoopsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIndustrialPidLoopsArgs Empty = new GetIndustrialPidLoopsArgs();

    @Import(name="enableDetails")
    private @Nullable Output<Boolean> enableDetails;

    public Optional<Output<Boolean>> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * A list of Pid Loop IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Pid Loop IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Pid Loop name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Pid Loop name.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The name of Pid Loop.
     * 
     */
    @Import(name="pidLoopName")
    private @Nullable Output<String> pidLoopName;

    /**
     * @return The name of Pid Loop.
     * 
     */
    public Optional<Output<String>> pidLoopName() {
        return Optional.ofNullable(this.pidLoopName);
    }

    /**
     * The pid project id.
     * 
     */
    @Import(name="pidProjectId", required=true)
    private Output<String> pidProjectId;

    /**
     * @return The pid project id.
     * 
     */
    public Output<String> pidProjectId() {
        return this.pidProjectId;
    }

    /**
     * The status of Pid Loop.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of Pid Loop.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private GetIndustrialPidLoopsArgs() {}

    private GetIndustrialPidLoopsArgs(GetIndustrialPidLoopsArgs $) {
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.pidLoopName = $.pidLoopName;
        this.pidProjectId = $.pidProjectId;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIndustrialPidLoopsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIndustrialPidLoopsArgs $;

        public Builder() {
            $ = new GetIndustrialPidLoopsArgs();
        }

        public Builder(GetIndustrialPidLoopsArgs defaults) {
            $ = new GetIndustrialPidLoopsArgs(Objects.requireNonNull(defaults));
        }

        public Builder enableDetails(@Nullable Output<Boolean> enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        public Builder enableDetails(Boolean enableDetails) {
            return enableDetails(Output.of(enableDetails));
        }

        /**
         * @param ids A list of Pid Loop IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Pid Loop IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Pid Loop IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Pid Loop name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Pid Loop name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
        }

        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        /**
         * @param pidLoopName The name of Pid Loop.
         * 
         * @return builder
         * 
         */
        public Builder pidLoopName(@Nullable Output<String> pidLoopName) {
            $.pidLoopName = pidLoopName;
            return this;
        }

        /**
         * @param pidLoopName The name of Pid Loop.
         * 
         * @return builder
         * 
         */
        public Builder pidLoopName(String pidLoopName) {
            return pidLoopName(Output.of(pidLoopName));
        }

        /**
         * @param pidProjectId The pid project id.
         * 
         * @return builder
         * 
         */
        public Builder pidProjectId(Output<String> pidProjectId) {
            $.pidProjectId = pidProjectId;
            return this;
        }

        /**
         * @param pidProjectId The pid project id.
         * 
         * @return builder
         * 
         */
        public Builder pidProjectId(String pidProjectId) {
            return pidProjectId(Output.of(pidProjectId));
        }

        /**
         * @param status The status of Pid Loop.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of Pid Loop.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public GetIndustrialPidLoopsArgs build() {
            $.pidProjectId = Objects.requireNonNull($.pidProjectId, "expected parameter 'pidProjectId' to be non-null");
            return $;
        }
    }

}