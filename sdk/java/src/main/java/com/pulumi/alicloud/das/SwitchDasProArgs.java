// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.das;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SwitchDasProArgs extends com.pulumi.resources.ResourceArgs {

    public static final SwitchDasProArgs Empty = new SwitchDasProArgs();

    /**
     * The ID of the database instance.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return The ID of the database instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     * The storage duration of SQL Explorer data. Valid values: `30`, `180`, `365`, `1095`, `1825`. Unit: days. Default value: `30`.
     * 
     */
    @Import(name="sqlRetention")
    private @Nullable Output<Integer> sqlRetention;

    /**
     * @return The storage duration of SQL Explorer data. Valid values: `30`, `180`, `365`, `1095`, `1825`. Unit: days. Default value: `30`.
     * 
     */
    public Optional<Output<Integer>> sqlRetention() {
        return Optional.ofNullable(this.sqlRetention);
    }

    /**
     * The ID of the Alibaba Cloud account that is used to create the database instance.
     * 
     */
    @Import(name="userId")
    private @Nullable Output<String> userId;

    /**
     * @return The ID of the Alibaba Cloud account that is used to create the database instance.
     * 
     */
    public Optional<Output<String>> userId() {
        return Optional.ofNullable(this.userId);
    }

    private SwitchDasProArgs() {}

    private SwitchDasProArgs(SwitchDasProArgs $) {
        this.instanceId = $.instanceId;
        this.sqlRetention = $.sqlRetention;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SwitchDasProArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SwitchDasProArgs $;

        public Builder() {
            $ = new SwitchDasProArgs();
        }

        public Builder(SwitchDasProArgs defaults) {
            $ = new SwitchDasProArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param instanceId The ID of the database instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The ID of the database instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param sqlRetention The storage duration of SQL Explorer data. Valid values: `30`, `180`, `365`, `1095`, `1825`. Unit: days. Default value: `30`.
         * 
         * @return builder
         * 
         */
        public Builder sqlRetention(@Nullable Output<Integer> sqlRetention) {
            $.sqlRetention = sqlRetention;
            return this;
        }

        /**
         * @param sqlRetention The storage duration of SQL Explorer data. Valid values: `30`, `180`, `365`, `1095`, `1825`. Unit: days. Default value: `30`.
         * 
         * @return builder
         * 
         */
        public Builder sqlRetention(Integer sqlRetention) {
            return sqlRetention(Output.of(sqlRetention));
        }

        /**
         * @param userId The ID of the Alibaba Cloud account that is used to create the database instance.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Output<String> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId The ID of the Alibaba Cloud account that is used to create the database instance.
         * 
         * @return builder
         * 
         */
        public Builder userId(String userId) {
            return userId(Output.of(userId));
        }

        public SwitchDasProArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}