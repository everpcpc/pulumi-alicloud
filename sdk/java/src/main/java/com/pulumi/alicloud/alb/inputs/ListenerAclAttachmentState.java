// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ListenerAclAttachmentState extends com.pulumi.resources.ResourceArgs {

    public static final ListenerAclAttachmentState Empty = new ListenerAclAttachmentState();

    /**
     * The ID of the Acl.
     * 
     */
    @Import(name="aclId")
    private @Nullable Output<String> aclId;

    /**
     * @return The ID of the Acl.
     * 
     */
    public Optional<Output<String>> aclId() {
        return Optional.ofNullable(this.aclId);
    }

    /**
     * The type of the ACL. Valid values:
     * - White: a whitelist. Only requests from the IP addresses or CIDR blocks in the ACL are forwarded. The whitelist applies to scenarios in which you want to allow only specific IP addresses to access an application. Risks may arise if you specify an ACL as a whitelist. After a whitelist is configured, only IP addresses in the whitelist can access the Application Load Balancer (ALB) listener. If you enable a whitelist but the whitelist does not contain an IP address, the listener forwards all requests.
     * - Black: a blacklist. All requests from the IP addresses or CIDR blocks in the ACL are blocked. The blacklist applies to scenarios in which you want to block access from specific IP addresses to an application. If you enable a blacklist but the blacklist does not contain an IP address, the listener forwards all requests.
     * 
     */
    @Import(name="aclType")
    private @Nullable Output<String> aclType;

    /**
     * @return The type of the ACL. Valid values:
     * - White: a whitelist. Only requests from the IP addresses or CIDR blocks in the ACL are forwarded. The whitelist applies to scenarios in which you want to allow only specific IP addresses to access an application. Risks may arise if you specify an ACL as a whitelist. After a whitelist is configured, only IP addresses in the whitelist can access the Application Load Balancer (ALB) listener. If you enable a whitelist but the whitelist does not contain an IP address, the listener forwards all requests.
     * - Black: a blacklist. All requests from the IP addresses or CIDR blocks in the ACL are blocked. The blacklist applies to scenarios in which you want to block access from specific IP addresses to an application. If you enable a blacklist but the blacklist does not contain an IP address, the listener forwards all requests.
     * 
     */
    public Optional<Output<String>> aclType() {
        return Optional.ofNullable(this.aclType);
    }

    /**
     * The ID of the ALB listener.
     * 
     */
    @Import(name="listenerId")
    private @Nullable Output<String> listenerId;

    /**
     * @return The ID of the ALB listener.
     * 
     */
    public Optional<Output<String>> listenerId() {
        return Optional.ofNullable(this.listenerId);
    }

    /**
     * The status of the Listener Acl Attachment.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the Listener Acl Attachment.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private ListenerAclAttachmentState() {}

    private ListenerAclAttachmentState(ListenerAclAttachmentState $) {
        this.aclId = $.aclId;
        this.aclType = $.aclType;
        this.listenerId = $.listenerId;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ListenerAclAttachmentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ListenerAclAttachmentState $;

        public Builder() {
            $ = new ListenerAclAttachmentState();
        }

        public Builder(ListenerAclAttachmentState defaults) {
            $ = new ListenerAclAttachmentState(Objects.requireNonNull(defaults));
        }

        /**
         * @param aclId The ID of the Acl.
         * 
         * @return builder
         * 
         */
        public Builder aclId(@Nullable Output<String> aclId) {
            $.aclId = aclId;
            return this;
        }

        /**
         * @param aclId The ID of the Acl.
         * 
         * @return builder
         * 
         */
        public Builder aclId(String aclId) {
            return aclId(Output.of(aclId));
        }

        /**
         * @param aclType The type of the ACL. Valid values:
         * - White: a whitelist. Only requests from the IP addresses or CIDR blocks in the ACL are forwarded. The whitelist applies to scenarios in which you want to allow only specific IP addresses to access an application. Risks may arise if you specify an ACL as a whitelist. After a whitelist is configured, only IP addresses in the whitelist can access the Application Load Balancer (ALB) listener. If you enable a whitelist but the whitelist does not contain an IP address, the listener forwards all requests.
         * - Black: a blacklist. All requests from the IP addresses or CIDR blocks in the ACL are blocked. The blacklist applies to scenarios in which you want to block access from specific IP addresses to an application. If you enable a blacklist but the blacklist does not contain an IP address, the listener forwards all requests.
         * 
         * @return builder
         * 
         */
        public Builder aclType(@Nullable Output<String> aclType) {
            $.aclType = aclType;
            return this;
        }

        /**
         * @param aclType The type of the ACL. Valid values:
         * - White: a whitelist. Only requests from the IP addresses or CIDR blocks in the ACL are forwarded. The whitelist applies to scenarios in which you want to allow only specific IP addresses to access an application. Risks may arise if you specify an ACL as a whitelist. After a whitelist is configured, only IP addresses in the whitelist can access the Application Load Balancer (ALB) listener. If you enable a whitelist but the whitelist does not contain an IP address, the listener forwards all requests.
         * - Black: a blacklist. All requests from the IP addresses or CIDR blocks in the ACL are blocked. The blacklist applies to scenarios in which you want to block access from specific IP addresses to an application. If you enable a blacklist but the blacklist does not contain an IP address, the listener forwards all requests.
         * 
         * @return builder
         * 
         */
        public Builder aclType(String aclType) {
            return aclType(Output.of(aclType));
        }

        /**
         * @param listenerId The ID of the ALB listener.
         * 
         * @return builder
         * 
         */
        public Builder listenerId(@Nullable Output<String> listenerId) {
            $.listenerId = listenerId;
            return this;
        }

        /**
         * @param listenerId The ID of the ALB listener.
         * 
         * @return builder
         * 
         */
        public Builder listenerId(String listenerId) {
            return listenerId(Output.of(listenerId));
        }

        /**
         * @param status The status of the Listener Acl Attachment.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the Listener Acl Attachment.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public ListenerAclAttachmentState build() {
            return $;
        }
    }

}
