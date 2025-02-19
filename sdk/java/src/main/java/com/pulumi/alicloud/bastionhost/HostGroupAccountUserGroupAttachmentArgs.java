// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.bastionhost;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class HostGroupAccountUserGroupAttachmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final HostGroupAccountUserGroupAttachmentArgs Empty = new HostGroupAccountUserGroupAttachmentArgs();

    /**
     * A list names of the host account.
     * 
     */
    @Import(name="hostAccountNames", required=true)
    private Output<List<String>> hostAccountNames;

    /**
     * @return A list names of the host account.
     * 
     */
    public Output<List<String>> hostAccountNames() {
        return this.hostAccountNames;
    }

    /**
     * The ID of the host group.
     * 
     */
    @Import(name="hostGroupId", required=true)
    private Output<String> hostGroupId;

    /**
     * @return The ID of the host group.
     * 
     */
    public Output<String> hostGroupId() {
        return this.hostGroupId;
    }

    /**
     * The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     * The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
     * 
     */
    @Import(name="userGroupId", required=true)
    private Output<String> userGroupId;

    /**
     * @return The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
     * 
     */
    public Output<String> userGroupId() {
        return this.userGroupId;
    }

    private HostGroupAccountUserGroupAttachmentArgs() {}

    private HostGroupAccountUserGroupAttachmentArgs(HostGroupAccountUserGroupAttachmentArgs $) {
        this.hostAccountNames = $.hostAccountNames;
        this.hostGroupId = $.hostGroupId;
        this.instanceId = $.instanceId;
        this.userGroupId = $.userGroupId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HostGroupAccountUserGroupAttachmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HostGroupAccountUserGroupAttachmentArgs $;

        public Builder() {
            $ = new HostGroupAccountUserGroupAttachmentArgs();
        }

        public Builder(HostGroupAccountUserGroupAttachmentArgs defaults) {
            $ = new HostGroupAccountUserGroupAttachmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param hostAccountNames A list names of the host account.
         * 
         * @return builder
         * 
         */
        public Builder hostAccountNames(Output<List<String>> hostAccountNames) {
            $.hostAccountNames = hostAccountNames;
            return this;
        }

        /**
         * @param hostAccountNames A list names of the host account.
         * 
         * @return builder
         * 
         */
        public Builder hostAccountNames(List<String> hostAccountNames) {
            return hostAccountNames(Output.of(hostAccountNames));
        }

        /**
         * @param hostAccountNames A list names of the host account.
         * 
         * @return builder
         * 
         */
        public Builder hostAccountNames(String... hostAccountNames) {
            return hostAccountNames(List.of(hostAccountNames));
        }

        /**
         * @param hostGroupId The ID of the host group.
         * 
         * @return builder
         * 
         */
        public Builder hostGroupId(Output<String> hostGroupId) {
            $.hostGroupId = hostGroupId;
            return this;
        }

        /**
         * @param hostGroupId The ID of the host group.
         * 
         * @return builder
         * 
         */
        public Builder hostGroupId(String hostGroupId) {
            return hostGroupId(Output.of(hostGroupId));
        }

        /**
         * @param instanceId The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param userGroupId The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
         * 
         * @return builder
         * 
         */
        public Builder userGroupId(Output<String> userGroupId) {
            $.userGroupId = userGroupId;
            return this;
        }

        /**
         * @param userGroupId The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
         * 
         * @return builder
         * 
         */
        public Builder userGroupId(String userGroupId) {
            return userGroupId(Output.of(userGroupId));
        }

        public HostGroupAccountUserGroupAttachmentArgs build() {
            $.hostAccountNames = Objects.requireNonNull($.hostAccountNames, "expected parameter 'hostAccountNames' to be non-null");
            $.hostGroupId = Objects.requireNonNull($.hostGroupId, "expected parameter 'hostGroupId' to be non-null");
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            $.userGroupId = Objects.requireNonNull($.userGroupId, "expected parameter 'userGroupId' to be non-null");
            return $;
        }
    }

}
