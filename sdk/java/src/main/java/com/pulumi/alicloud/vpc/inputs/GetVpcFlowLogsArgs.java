// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVpcFlowLogsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVpcFlowLogsArgs Empty = new GetVpcFlowLogsArgs();

    /**
     * The Description of flow log.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The Description of flow log.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The flow log name.
     * 
     */
    @Import(name="flowLogName")
    private @Nullable Output<String> flowLogName;

    /**
     * @return The flow log name.
     * 
     */
    public Optional<Output<String>> flowLogName() {
        return Optional.ofNullable(this.flowLogName);
    }

    /**
     * A list of Flow Log IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Flow Log IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The log store name.
     * 
     */
    @Import(name="logStoreName")
    private @Nullable Output<String> logStoreName;

    /**
     * @return The log store name.
     * 
     */
    public Optional<Output<String>> logStoreName() {
        return Optional.ofNullable(this.logStoreName);
    }

    /**
     * A regex string to filter results by Flow Log name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Flow Log name.
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

    /**
     * The project name.
     * 
     */
    @Import(name="projectName")
    private @Nullable Output<String> projectName;

    /**
     * @return The project name.
     * 
     */
    public Optional<Output<String>> projectName() {
        return Optional.ofNullable(this.projectName);
    }

    /**
     * The resource id.
     * 
     */
    @Import(name="resourceId")
    private @Nullable Output<String> resourceId;

    /**
     * @return The resource id.
     * 
     */
    public Optional<Output<String>> resourceId() {
        return Optional.ofNullable(this.resourceId);
    }

    /**
     * The resource type.
     * 
     */
    @Import(name="resourceType")
    private @Nullable Output<String> resourceType;

    /**
     * @return The resource type.
     * 
     */
    public Optional<Output<String>> resourceType() {
        return Optional.ofNullable(this.resourceType);
    }

    /**
     * The status of flow log.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of flow log.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The traffic type.
     * 
     */
    @Import(name="trafficType")
    private @Nullable Output<String> trafficType;

    /**
     * @return The traffic type.
     * 
     */
    public Optional<Output<String>> trafficType() {
        return Optional.ofNullable(this.trafficType);
    }

    private GetVpcFlowLogsArgs() {}

    private GetVpcFlowLogsArgs(GetVpcFlowLogsArgs $) {
        this.description = $.description;
        this.flowLogName = $.flowLogName;
        this.ids = $.ids;
        this.logStoreName = $.logStoreName;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.projectName = $.projectName;
        this.resourceId = $.resourceId;
        this.resourceType = $.resourceType;
        this.status = $.status;
        this.trafficType = $.trafficType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVpcFlowLogsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVpcFlowLogsArgs $;

        public Builder() {
            $ = new GetVpcFlowLogsArgs();
        }

        public Builder(GetVpcFlowLogsArgs defaults) {
            $ = new GetVpcFlowLogsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The Description of flow log.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The Description of flow log.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param flowLogName The flow log name.
         * 
         * @return builder
         * 
         */
        public Builder flowLogName(@Nullable Output<String> flowLogName) {
            $.flowLogName = flowLogName;
            return this;
        }

        /**
         * @param flowLogName The flow log name.
         * 
         * @return builder
         * 
         */
        public Builder flowLogName(String flowLogName) {
            return flowLogName(Output.of(flowLogName));
        }

        /**
         * @param ids A list of Flow Log IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Flow Log IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Flow Log IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param logStoreName The log store name.
         * 
         * @return builder
         * 
         */
        public Builder logStoreName(@Nullable Output<String> logStoreName) {
            $.logStoreName = logStoreName;
            return this;
        }

        /**
         * @param logStoreName The log store name.
         * 
         * @return builder
         * 
         */
        public Builder logStoreName(String logStoreName) {
            return logStoreName(Output.of(logStoreName));
        }

        /**
         * @param nameRegex A regex string to filter results by Flow Log name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Flow Log name.
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

        /**
         * @param projectName The project name.
         * 
         * @return builder
         * 
         */
        public Builder projectName(@Nullable Output<String> projectName) {
            $.projectName = projectName;
            return this;
        }

        /**
         * @param projectName The project name.
         * 
         * @return builder
         * 
         */
        public Builder projectName(String projectName) {
            return projectName(Output.of(projectName));
        }

        /**
         * @param resourceId The resource id.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(@Nullable Output<String> resourceId) {
            $.resourceId = resourceId;
            return this;
        }

        /**
         * @param resourceId The resource id.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(String resourceId) {
            return resourceId(Output.of(resourceId));
        }

        /**
         * @param resourceType The resource type.
         * 
         * @return builder
         * 
         */
        public Builder resourceType(@Nullable Output<String> resourceType) {
            $.resourceType = resourceType;
            return this;
        }

        /**
         * @param resourceType The resource type.
         * 
         * @return builder
         * 
         */
        public Builder resourceType(String resourceType) {
            return resourceType(Output.of(resourceType));
        }

        /**
         * @param status The status of flow log.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of flow log.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param trafficType The traffic type.
         * 
         * @return builder
         * 
         */
        public Builder trafficType(@Nullable Output<String> trafficType) {
            $.trafficType = trafficType;
            return this;
        }

        /**
         * @param trafficType The traffic type.
         * 
         * @return builder
         * 
         */
        public Builder trafficType(String trafficType) {
            return trafficType(Output.of(trafficType));
        }

        public GetVpcFlowLogsArgs build() {
            return $;
        }
    }

}
