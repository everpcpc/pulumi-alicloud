// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mongodb;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ShardingNetworkPublicAddressArgs extends com.pulumi.resources.ResourceArgs {

    public static final ShardingNetworkPublicAddressArgs Empty = new ShardingNetworkPublicAddressArgs();

    /**
     * The ID of the instance.
     * 
     */
    @Import(name="dbInstanceId", required=true)
    private Output<String> dbInstanceId;

    /**
     * @return The ID of the instance.
     * 
     */
    public Output<String> dbInstanceId() {
        return this.dbInstanceId;
    }

    /**
     * The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
     * 
     */
    @Import(name="nodeId", required=true)
    private Output<String> nodeId;

    /**
     * @return The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
     * 
     */
    public Output<String> nodeId() {
        return this.nodeId;
    }

    private ShardingNetworkPublicAddressArgs() {}

    private ShardingNetworkPublicAddressArgs(ShardingNetworkPublicAddressArgs $) {
        this.dbInstanceId = $.dbInstanceId;
        this.nodeId = $.nodeId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ShardingNetworkPublicAddressArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ShardingNetworkPublicAddressArgs $;

        public Builder() {
            $ = new ShardingNetworkPublicAddressArgs();
        }

        public Builder(ShardingNetworkPublicAddressArgs defaults) {
            $ = new ShardingNetworkPublicAddressArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dbInstanceId The ID of the instance.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceId(Output<String> dbInstanceId) {
            $.dbInstanceId = dbInstanceId;
            return this;
        }

        /**
         * @param dbInstanceId The ID of the instance.
         * 
         * @return builder
         * 
         */
        public Builder dbInstanceId(String dbInstanceId) {
            return dbInstanceId(Output.of(dbInstanceId));
        }

        /**
         * @param nodeId The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
         * 
         * @return builder
         * 
         */
        public Builder nodeId(Output<String> nodeId) {
            $.nodeId = nodeId;
            return this;
        }

        /**
         * @param nodeId The ID of the `mongos`, `shard`, or `Configserver` node in the sharded cluster instance.
         * 
         * @return builder
         * 
         */
        public Builder nodeId(String nodeId) {
            return nodeId(Output.of(nodeId));
        }

        public ShardingNetworkPublicAddressArgs build() {
            $.dbInstanceId = Objects.requireNonNull($.dbInstanceId, "expected parameter 'dbInstanceId' to be non-null");
            $.nodeId = Objects.requireNonNull($.nodeId, "expected parameter 'nodeId' to be non-null");
            return $;
        }
    }

}
