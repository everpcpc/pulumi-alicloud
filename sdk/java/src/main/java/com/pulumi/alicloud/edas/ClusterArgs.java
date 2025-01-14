// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.edas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterArgs Empty = new ClusterArgs();

    /**
     * The name of the cluster that you want to create.
     * 
     */
    @Import(name="clusterName", required=true)
    private Output<String> clusterName;

    /**
     * @return The name of the cluster that you want to create.
     * 
     */
    public Output<String> clusterName() {
        return this.clusterName;
    }

    /**
     * The type of the cluster that you want to create. Valid values only: 2: ECS cluster.
     * 
     */
    @Import(name="clusterType", required=true)
    private Output<Integer> clusterType;

    /**
     * @return The type of the cluster that you want to create. Valid values only: 2: ECS cluster.
     * 
     */
    public Output<Integer> clusterType() {
        return this.clusterType;
    }

    /**
     * The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
     * 
     */
    @Import(name="logicalRegionId")
    private @Nullable Output<String> logicalRegionId;

    /**
     * @return The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
     * 
     */
    public Optional<Output<String>> logicalRegionId() {
        return Optional.ofNullable(this.logicalRegionId);
    }

    /**
     * The network type of the cluster that you want to create. Valid values: 1: classic network. 2: VPC.
     * 
     */
    @Import(name="networkMode", required=true)
    private Output<Integer> networkMode;

    /**
     * @return The network type of the cluster that you want to create. Valid values: 1: classic network. 2: VPC.
     * 
     */
    public Output<Integer> networkMode() {
        return this.networkMode;
    }

    /**
     * The ID of the Virtual Private Cloud (VPC) for the cluster.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The ID of the Virtual Private Cloud (VPC) for the cluster.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    private ClusterArgs() {}

    private ClusterArgs(ClusterArgs $) {
        this.clusterName = $.clusterName;
        this.clusterType = $.clusterType;
        this.logicalRegionId = $.logicalRegionId;
        this.networkMode = $.networkMode;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterArgs $;

        public Builder() {
            $ = new ClusterArgs();
        }

        public Builder(ClusterArgs defaults) {
            $ = new ClusterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterName The name of the cluster that you want to create.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(Output<String> clusterName) {
            $.clusterName = clusterName;
            return this;
        }

        /**
         * @param clusterName The name of the cluster that you want to create.
         * 
         * @return builder
         * 
         */
        public Builder clusterName(String clusterName) {
            return clusterName(Output.of(clusterName));
        }

        /**
         * @param clusterType The type of the cluster that you want to create. Valid values only: 2: ECS cluster.
         * 
         * @return builder
         * 
         */
        public Builder clusterType(Output<Integer> clusterType) {
            $.clusterType = clusterType;
            return this;
        }

        /**
         * @param clusterType The type of the cluster that you want to create. Valid values only: 2: ECS cluster.
         * 
         * @return builder
         * 
         */
        public Builder clusterType(Integer clusterType) {
            return clusterType(Output.of(clusterType));
        }

        /**
         * @param logicalRegionId The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
         * 
         * @return builder
         * 
         */
        public Builder logicalRegionId(@Nullable Output<String> logicalRegionId) {
            $.logicalRegionId = logicalRegionId;
            return this;
        }

        /**
         * @param logicalRegionId The ID of the namespace where you want to create the application. You can call the ListUserDefineRegion operation to query the namespace ID.
         * 
         * @return builder
         * 
         */
        public Builder logicalRegionId(String logicalRegionId) {
            return logicalRegionId(Output.of(logicalRegionId));
        }

        /**
         * @param networkMode The network type of the cluster that you want to create. Valid values: 1: classic network. 2: VPC.
         * 
         * @return builder
         * 
         */
        public Builder networkMode(Output<Integer> networkMode) {
            $.networkMode = networkMode;
            return this;
        }

        /**
         * @param networkMode The network type of the cluster that you want to create. Valid values: 1: classic network. 2: VPC.
         * 
         * @return builder
         * 
         */
        public Builder networkMode(Integer networkMode) {
            return networkMode(Output.of(networkMode));
        }

        /**
         * @param vpcId The ID of the Virtual Private Cloud (VPC) for the cluster.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The ID of the Virtual Private Cloud (VPC) for the cluster.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        public ClusterArgs build() {
            $.clusterName = Objects.requireNonNull($.clusterName, "expected parameter 'clusterName' to be non-null");
            $.clusterType = Objects.requireNonNull($.clusterType, "expected parameter 'clusterType' to be non-null");
            $.networkMode = Objects.requireNonNull($.networkMode, "expected parameter 'networkMode' to be non-null");
            return $;
        }
    }

}
