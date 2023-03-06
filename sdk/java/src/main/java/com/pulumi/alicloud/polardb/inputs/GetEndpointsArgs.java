// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEndpointsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEndpointsArgs Empty = new GetEndpointsArgs();

    /**
     * PolarDB cluster ID.
     * 
     */
    @Import(name="dbClusterId", required=true)
    private Output<String> dbClusterId;

    /**
     * @return PolarDB cluster ID.
     * 
     */
    public Output<String> dbClusterId() {
        return this.dbClusterId;
    }

    /**
     * endpoint of the cluster.
     * 
     */
    @Import(name="dbEndpointId")
    private @Nullable Output<String> dbEndpointId;

    /**
     * @return endpoint of the cluster.
     * 
     */
    public Optional<Output<String>> dbEndpointId() {
        return Optional.ofNullable(this.dbEndpointId);
    }

    private GetEndpointsArgs() {}

    private GetEndpointsArgs(GetEndpointsArgs $) {
        this.dbClusterId = $.dbClusterId;
        this.dbEndpointId = $.dbEndpointId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEndpointsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEndpointsArgs $;

        public Builder() {
            $ = new GetEndpointsArgs();
        }

        public Builder(GetEndpointsArgs defaults) {
            $ = new GetEndpointsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dbClusterId PolarDB cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterId(Output<String> dbClusterId) {
            $.dbClusterId = dbClusterId;
            return this;
        }

        /**
         * @param dbClusterId PolarDB cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterId(String dbClusterId) {
            return dbClusterId(Output.of(dbClusterId));
        }

        /**
         * @param dbEndpointId endpoint of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder dbEndpointId(@Nullable Output<String> dbEndpointId) {
            $.dbEndpointId = dbEndpointId;
            return this;
        }

        /**
         * @param dbEndpointId endpoint of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder dbEndpointId(String dbEndpointId) {
            return dbEndpointId(Output.of(dbEndpointId));
        }

        public GetEndpointsArgs build() {
            $.dbClusterId = Objects.requireNonNull($.dbClusterId, "expected parameter 'dbClusterId' to be non-null");
            return $;
        }
    }

}