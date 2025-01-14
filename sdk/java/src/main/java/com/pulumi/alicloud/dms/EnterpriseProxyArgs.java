// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dms;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EnterpriseProxyArgs extends com.pulumi.resources.ResourceArgs {

    public static final EnterpriseProxyArgs Empty = new EnterpriseProxyArgs();

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
     * The password of the database account.
     * 
     */
    @Import(name="password", required=true)
    private Output<String> password;

    /**
     * @return The password of the database account.
     * 
     */
    public Output<String> password() {
        return this.password;
    }

    /**
     * The ID of the tenant.
     * 
     */
    @Import(name="tid")
    private @Nullable Output<String> tid;

    /**
     * @return The ID of the tenant.
     * 
     */
    public Optional<Output<String>> tid() {
        return Optional.ofNullable(this.tid);
    }

    /**
     * The username of the database account.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The username of the database account.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private EnterpriseProxyArgs() {}

    private EnterpriseProxyArgs(EnterpriseProxyArgs $) {
        this.instanceId = $.instanceId;
        this.password = $.password;
        this.tid = $.tid;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EnterpriseProxyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EnterpriseProxyArgs $;

        public Builder() {
            $ = new EnterpriseProxyArgs();
        }

        public Builder(EnterpriseProxyArgs defaults) {
            $ = new EnterpriseProxyArgs(Objects.requireNonNull(defaults));
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
         * @param password The password of the database account.
         * 
         * @return builder
         * 
         */
        public Builder password(Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The password of the database account.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param tid The ID of the tenant.
         * 
         * @return builder
         * 
         */
        public Builder tid(@Nullable Output<String> tid) {
            $.tid = tid;
            return this;
        }

        /**
         * @param tid The ID of the tenant.
         * 
         * @return builder
         * 
         */
        public Builder tid(String tid) {
            return tid(Output.of(tid));
        }

        /**
         * @param username The username of the database account.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username of the database account.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public EnterpriseProxyArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            $.password = Objects.requireNonNull($.password, "expected parameter 'password' to be non-null");
            $.username = Objects.requireNonNull($.username, "expected parameter 'username' to be non-null");
            return $;
        }
    }

}
