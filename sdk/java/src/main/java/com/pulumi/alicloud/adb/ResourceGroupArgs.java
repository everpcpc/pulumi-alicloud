// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.adb;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ResourceGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final ResourceGroupArgs Empty = new ResourceGroupArgs();

    /**
     * DB cluster id.
     * 
     */
    @Import(name="dbClusterId", required=true)
    private Output<String> dbClusterId;

    /**
     * @return DB cluster id.
     * 
     */
    public Output<String> dbClusterId() {
        return this.dbClusterId;
    }

    /**
     * The name of the resource pool. The group name must be 2 to 30 characters in length, and can contain upper case letters, digits, and underscore(_).
     * 
     */
    @Import(name="groupName", required=true)
    private Output<String> groupName;

    /**
     * @return The name of the resource pool. The group name must be 2 to 30 characters in length, and can contain upper case letters, digits, and underscore(_).
     * 
     */
    public Output<String> groupName() {
        return this.groupName;
    }

    /**
     * Query type, value description:
     * * **etl**: Batch query mode.
     * * **interactive**: interactive Query mode.
     * * **default_type**: the default query mode.
     * 
     */
    @Import(name="groupType")
    private @Nullable Output<String> groupType;

    /**
     * @return Query type, value description:
     * * **etl**: Batch query mode.
     * * **interactive**: interactive Query mode.
     * * **default_type**: the default query mode.
     * 
     */
    public Optional<Output<String>> groupType() {
        return Optional.ofNullable(this.groupType);
    }

    /**
     * The number of nodes. The default number of nodes is 0. The number of nodes must be less than or equal to the number of nodes whose resource name is USER_DEFAULT.
     * 
     */
    @Import(name="nodeNum")
    private @Nullable Output<Integer> nodeNum;

    /**
     * @return The number of nodes. The default number of nodes is 0. The number of nodes must be less than or equal to the number of nodes whose resource name is USER_DEFAULT.
     * 
     */
    public Optional<Output<Integer>> nodeNum() {
        return Optional.ofNullable(this.nodeNum);
    }

    private ResourceGroupArgs() {}

    private ResourceGroupArgs(ResourceGroupArgs $) {
        this.dbClusterId = $.dbClusterId;
        this.groupName = $.groupName;
        this.groupType = $.groupType;
        this.nodeNum = $.nodeNum;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ResourceGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ResourceGroupArgs $;

        public Builder() {
            $ = new ResourceGroupArgs();
        }

        public Builder(ResourceGroupArgs defaults) {
            $ = new ResourceGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dbClusterId DB cluster id.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterId(Output<String> dbClusterId) {
            $.dbClusterId = dbClusterId;
            return this;
        }

        /**
         * @param dbClusterId DB cluster id.
         * 
         * @return builder
         * 
         */
        public Builder dbClusterId(String dbClusterId) {
            return dbClusterId(Output.of(dbClusterId));
        }

        /**
         * @param groupName The name of the resource pool. The group name must be 2 to 30 characters in length, and can contain upper case letters, digits, and underscore(_).
         * 
         * @return builder
         * 
         */
        public Builder groupName(Output<String> groupName) {
            $.groupName = groupName;
            return this;
        }

        /**
         * @param groupName The name of the resource pool. The group name must be 2 to 30 characters in length, and can contain upper case letters, digits, and underscore(_).
         * 
         * @return builder
         * 
         */
        public Builder groupName(String groupName) {
            return groupName(Output.of(groupName));
        }

        /**
         * @param groupType Query type, value description:
         * * **etl**: Batch query mode.
         * * **interactive**: interactive Query mode.
         * * **default_type**: the default query mode.
         * 
         * @return builder
         * 
         */
        public Builder groupType(@Nullable Output<String> groupType) {
            $.groupType = groupType;
            return this;
        }

        /**
         * @param groupType Query type, value description:
         * * **etl**: Batch query mode.
         * * **interactive**: interactive Query mode.
         * * **default_type**: the default query mode.
         * 
         * @return builder
         * 
         */
        public Builder groupType(String groupType) {
            return groupType(Output.of(groupType));
        }

        /**
         * @param nodeNum The number of nodes. The default number of nodes is 0. The number of nodes must be less than or equal to the number of nodes whose resource name is USER_DEFAULT.
         * 
         * @return builder
         * 
         */
        public Builder nodeNum(@Nullable Output<Integer> nodeNum) {
            $.nodeNum = nodeNum;
            return this;
        }

        /**
         * @param nodeNum The number of nodes. The default number of nodes is 0. The number of nodes must be less than or equal to the number of nodes whose resource name is USER_DEFAULT.
         * 
         * @return builder
         * 
         */
        public Builder nodeNum(Integer nodeNum) {
            return nodeNum(Output.of(nodeNum));
        }

        public ResourceGroupArgs build() {
            $.dbClusterId = Objects.requireNonNull($.dbClusterId, "expected parameter 'dbClusterId' to be non-null");
            $.groupName = Objects.requireNonNull($.groupName, "expected parameter 'groupName' to be non-null");
            return $;
        }
    }

}
