// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mongodb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceReplicaSetArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceReplicaSetArgs Empty = new InstanceReplicaSetArgs();

    /**
     * The connection address of the node.
     * 
     */
    @Import(name="connectionDomain")
    private @Nullable Output<String> connectionDomain;

    /**
     * @return The connection address of the node.
     * 
     */
    public Optional<Output<String>> connectionDomain() {
        return Optional.ofNullable(this.connectionDomain);
    }

    /**
     * The connection port of the node.
     * 
     */
    @Import(name="connectionPort")
    private @Nullable Output<String> connectionPort;

    /**
     * @return The connection port of the node.
     * 
     */
    public Optional<Output<String>> connectionPort() {
        return Optional.ofNullable(this.connectionPort);
    }

    /**
     * The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
     * 
     */
    @Import(name="networkType")
    private @Nullable Output<String> networkType;

    /**
     * @return The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
     * 
     */
    public Optional<Output<String>> networkType() {
        return Optional.ofNullable(this.networkType);
    }

    /**
     * The role of the node. Valid values: `Primary`,`Secondary`.
     * 
     */
    @Import(name="replicaSetRole")
    private @Nullable Output<String> replicaSetRole;

    /**
     * @return The role of the node. Valid values: `Primary`,`Secondary`.
     * 
     */
    public Optional<Output<String>> replicaSetRole() {
        return Optional.ofNullable(this.replicaSetRole);
    }

    /**
     * VPC instance ID.
     * 
     */
    @Import(name="vpcCloudInstanceId")
    private @Nullable Output<String> vpcCloudInstanceId;

    /**
     * @return VPC instance ID.
     * 
     */
    public Optional<Output<String>> vpcCloudInstanceId() {
        return Optional.ofNullable(this.vpcCloudInstanceId);
    }

    /**
     * The ID of the VPC. &gt; **NOTE:** This parameter is valid only when NetworkType is set to VPC.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return The ID of the VPC. &gt; **NOTE:** This parameter is valid only when NetworkType is set to VPC.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    /**
     * The virtual switch ID to launch DB instances in one VPC.
     * 
     */
    @Import(name="vswitchId")
    private @Nullable Output<String> vswitchId;

    /**
     * @return The virtual switch ID to launch DB instances in one VPC.
     * 
     */
    public Optional<Output<String>> vswitchId() {
        return Optional.ofNullable(this.vswitchId);
    }

    private InstanceReplicaSetArgs() {}

    private InstanceReplicaSetArgs(InstanceReplicaSetArgs $) {
        this.connectionDomain = $.connectionDomain;
        this.connectionPort = $.connectionPort;
        this.networkType = $.networkType;
        this.replicaSetRole = $.replicaSetRole;
        this.vpcCloudInstanceId = $.vpcCloudInstanceId;
        this.vpcId = $.vpcId;
        this.vswitchId = $.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceReplicaSetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceReplicaSetArgs $;

        public Builder() {
            $ = new InstanceReplicaSetArgs();
        }

        public Builder(InstanceReplicaSetArgs defaults) {
            $ = new InstanceReplicaSetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param connectionDomain The connection address of the node.
         * 
         * @return builder
         * 
         */
        public Builder connectionDomain(@Nullable Output<String> connectionDomain) {
            $.connectionDomain = connectionDomain;
            return this;
        }

        /**
         * @param connectionDomain The connection address of the node.
         * 
         * @return builder
         * 
         */
        public Builder connectionDomain(String connectionDomain) {
            return connectionDomain(Output.of(connectionDomain));
        }

        /**
         * @param connectionPort The connection port of the node.
         * 
         * @return builder
         * 
         */
        public Builder connectionPort(@Nullable Output<String> connectionPort) {
            $.connectionPort = connectionPort;
            return this;
        }

        /**
         * @param connectionPort The connection port of the node.
         * 
         * @return builder
         * 
         */
        public Builder connectionPort(String connectionPort) {
            return connectionPort(Output.of(connectionPort));
        }

        /**
         * @param networkType The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
         * 
         * @return builder
         * 
         */
        public Builder networkType(@Nullable Output<String> networkType) {
            $.networkType = networkType;
            return this;
        }

        /**
         * @param networkType The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
         * 
         * @return builder
         * 
         */
        public Builder networkType(String networkType) {
            return networkType(Output.of(networkType));
        }

        /**
         * @param replicaSetRole The role of the node. Valid values: `Primary`,`Secondary`.
         * 
         * @return builder
         * 
         */
        public Builder replicaSetRole(@Nullable Output<String> replicaSetRole) {
            $.replicaSetRole = replicaSetRole;
            return this;
        }

        /**
         * @param replicaSetRole The role of the node. Valid values: `Primary`,`Secondary`.
         * 
         * @return builder
         * 
         */
        public Builder replicaSetRole(String replicaSetRole) {
            return replicaSetRole(Output.of(replicaSetRole));
        }

        /**
         * @param vpcCloudInstanceId VPC instance ID.
         * 
         * @return builder
         * 
         */
        public Builder vpcCloudInstanceId(@Nullable Output<String> vpcCloudInstanceId) {
            $.vpcCloudInstanceId = vpcCloudInstanceId;
            return this;
        }

        /**
         * @param vpcCloudInstanceId VPC instance ID.
         * 
         * @return builder
         * 
         */
        public Builder vpcCloudInstanceId(String vpcCloudInstanceId) {
            return vpcCloudInstanceId(Output.of(vpcCloudInstanceId));
        }

        /**
         * @param vpcId The ID of the VPC. &gt; **NOTE:** This parameter is valid only when NetworkType is set to VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId The ID of the VPC. &gt; **NOTE:** This parameter is valid only when NetworkType is set to VPC.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        /**
         * @param vswitchId The virtual switch ID to launch DB instances in one VPC.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(@Nullable Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId The virtual switch ID to launch DB instances in one VPC.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        public InstanceReplicaSetArgs build() {
            return $;
        }
    }

}