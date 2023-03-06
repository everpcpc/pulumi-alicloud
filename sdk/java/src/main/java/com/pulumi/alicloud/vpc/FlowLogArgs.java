// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FlowLogArgs extends com.pulumi.resources.ResourceArgs {

    public static final FlowLogArgs Empty = new FlowLogArgs();

    /**
     * The Description of the VPC Flow Log.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The Description of the VPC Flow Log.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The Name of the VPC Flow Log.
     * 
     */
    @Import(name="flowLogName")
    private @Nullable Output<String> flowLogName;

    /**
     * @return The Name of the VPC Flow Log.
     * 
     */
    public Optional<Output<String>> flowLogName() {
        return Optional.ofNullable(this.flowLogName);
    }

    /**
     * The name of the logstore.
     * 
     */
    @Import(name="logStoreName", required=true)
    private Output<String> logStoreName;

    /**
     * @return The name of the logstore.
     * 
     */
    public Output<String> logStoreName() {
        return this.logStoreName;
    }

    /**
     * The name of the project.
     * 
     */
    @Import(name="projectName", required=true)
    private Output<String> projectName;

    /**
     * @return The name of the project.
     * 
     */
    public Output<String> projectName() {
        return this.projectName;
    }

    /**
     * The ID of the resource.
     * 
     */
    @Import(name="resourceId", required=true)
    private Output<String> resourceId;

    /**
     * @return The ID of the resource.
     * 
     */
    public Output<String> resourceId() {
        return this.resourceId;
    }

    /**
     * The type of the resource to capture traffic. Valid values `NetworkInterface`, `VPC`, and `VSwitch`.
     * 
     */
    @Import(name="resourceType", required=true)
    private Output<String> resourceType;

    /**
     * @return The type of the resource to capture traffic. Valid values `NetworkInterface`, `VPC`, and `VSwitch`.
     * 
     */
    public Output<String> resourceType() {
        return this.resourceType;
    }

    /**
     * The status of the VPC Flow Log. Valid values `Active` and `Inactive`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the VPC Flow Log. Valid values `Active` and `Inactive`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The type of traffic collected. Valid values `All`, `Drop` and `Allow`.
     * 
     */
    @Import(name="trafficType", required=true)
    private Output<String> trafficType;

    /**
     * @return The type of traffic collected. Valid values `All`, `Drop` and `Allow`.
     * 
     */
    public Output<String> trafficType() {
        return this.trafficType;
    }

    private FlowLogArgs() {}

    private FlowLogArgs(FlowLogArgs $) {
        this.description = $.description;
        this.flowLogName = $.flowLogName;
        this.logStoreName = $.logStoreName;
        this.projectName = $.projectName;
        this.resourceId = $.resourceId;
        this.resourceType = $.resourceType;
        this.status = $.status;
        this.trafficType = $.trafficType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FlowLogArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FlowLogArgs $;

        public Builder() {
            $ = new FlowLogArgs();
        }

        public Builder(FlowLogArgs defaults) {
            $ = new FlowLogArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The Description of the VPC Flow Log.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The Description of the VPC Flow Log.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param flowLogName The Name of the VPC Flow Log.
         * 
         * @return builder
         * 
         */
        public Builder flowLogName(@Nullable Output<String> flowLogName) {
            $.flowLogName = flowLogName;
            return this;
        }

        /**
         * @param flowLogName The Name of the VPC Flow Log.
         * 
         * @return builder
         * 
         */
        public Builder flowLogName(String flowLogName) {
            return flowLogName(Output.of(flowLogName));
        }

        /**
         * @param logStoreName The name of the logstore.
         * 
         * @return builder
         * 
         */
        public Builder logStoreName(Output<String> logStoreName) {
            $.logStoreName = logStoreName;
            return this;
        }

        /**
         * @param logStoreName The name of the logstore.
         * 
         * @return builder
         * 
         */
        public Builder logStoreName(String logStoreName) {
            return logStoreName(Output.of(logStoreName));
        }

        /**
         * @param projectName The name of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectName(Output<String> projectName) {
            $.projectName = projectName;
            return this;
        }

        /**
         * @param projectName The name of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectName(String projectName) {
            return projectName(Output.of(projectName));
        }

        /**
         * @param resourceId The ID of the resource.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(Output<String> resourceId) {
            $.resourceId = resourceId;
            return this;
        }

        /**
         * @param resourceId The ID of the resource.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(String resourceId) {
            return resourceId(Output.of(resourceId));
        }

        /**
         * @param resourceType The type of the resource to capture traffic. Valid values `NetworkInterface`, `VPC`, and `VSwitch`.
         * 
         * @return builder
         * 
         */
        public Builder resourceType(Output<String> resourceType) {
            $.resourceType = resourceType;
            return this;
        }

        /**
         * @param resourceType The type of the resource to capture traffic. Valid values `NetworkInterface`, `VPC`, and `VSwitch`.
         * 
         * @return builder
         * 
         */
        public Builder resourceType(String resourceType) {
            return resourceType(Output.of(resourceType));
        }

        /**
         * @param status The status of the VPC Flow Log. Valid values `Active` and `Inactive`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the VPC Flow Log. Valid values `Active` and `Inactive`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param trafficType The type of traffic collected. Valid values `All`, `Drop` and `Allow`.
         * 
         * @return builder
         * 
         */
        public Builder trafficType(Output<String> trafficType) {
            $.trafficType = trafficType;
            return this;
        }

        /**
         * @param trafficType The type of traffic collected. Valid values `All`, `Drop` and `Allow`.
         * 
         * @return builder
         * 
         */
        public Builder trafficType(String trafficType) {
            return trafficType(Output.of(trafficType));
        }

        public FlowLogArgs build() {
            $.logStoreName = Objects.requireNonNull($.logStoreName, "expected parameter 'logStoreName' to be non-null");
            $.projectName = Objects.requireNonNull($.projectName, "expected parameter 'projectName' to be non-null");
            $.resourceId = Objects.requireNonNull($.resourceId, "expected parameter 'resourceId' to be non-null");
            $.resourceType = Objects.requireNonNull($.resourceType, "expected parameter 'resourceType' to be non-null");
            $.trafficType = Objects.requireNonNull($.trafficType, "expected parameter 'trafficType' to be non-null");
            return $;
        }
    }

}