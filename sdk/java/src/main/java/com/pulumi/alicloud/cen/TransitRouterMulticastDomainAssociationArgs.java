// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class TransitRouterMulticastDomainAssociationArgs extends com.pulumi.resources.ResourceArgs {

    public static final TransitRouterMulticastDomainAssociationArgs Empty = new TransitRouterMulticastDomainAssociationArgs();

    /**
     * The ID of the VPC connection.
     * 
     */
    @Import(name="transitRouterAttachmentId", required=true)
    private Output<String> transitRouterAttachmentId;

    /**
     * @return The ID of the VPC connection.
     * 
     */
    public Output<String> transitRouterAttachmentId() {
        return this.transitRouterAttachmentId;
    }

    /**
     * The ID of the multicast domain.
     * 
     */
    @Import(name="transitRouterMulticastDomainId", required=true)
    private Output<String> transitRouterMulticastDomainId;

    /**
     * @return The ID of the multicast domain.
     * 
     */
    public Output<String> transitRouterMulticastDomainId() {
        return this.transitRouterMulticastDomainId;
    }

    /**
     * The ID of the vSwitch.
     * 
     */
    @Import(name="vswitchId", required=true)
    private Output<String> vswitchId;

    /**
     * @return The ID of the vSwitch.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }

    private TransitRouterMulticastDomainAssociationArgs() {}

    private TransitRouterMulticastDomainAssociationArgs(TransitRouterMulticastDomainAssociationArgs $) {
        this.transitRouterAttachmentId = $.transitRouterAttachmentId;
        this.transitRouterMulticastDomainId = $.transitRouterMulticastDomainId;
        this.vswitchId = $.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TransitRouterMulticastDomainAssociationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TransitRouterMulticastDomainAssociationArgs $;

        public Builder() {
            $ = new TransitRouterMulticastDomainAssociationArgs();
        }

        public Builder(TransitRouterMulticastDomainAssociationArgs defaults) {
            $ = new TransitRouterMulticastDomainAssociationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param transitRouterAttachmentId The ID of the VPC connection.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterAttachmentId(Output<String> transitRouterAttachmentId) {
            $.transitRouterAttachmentId = transitRouterAttachmentId;
            return this;
        }

        /**
         * @param transitRouterAttachmentId The ID of the VPC connection.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterAttachmentId(String transitRouterAttachmentId) {
            return transitRouterAttachmentId(Output.of(transitRouterAttachmentId));
        }

        /**
         * @param transitRouterMulticastDomainId The ID of the multicast domain.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterMulticastDomainId(Output<String> transitRouterMulticastDomainId) {
            $.transitRouterMulticastDomainId = transitRouterMulticastDomainId;
            return this;
        }

        /**
         * @param transitRouterMulticastDomainId The ID of the multicast domain.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterMulticastDomainId(String transitRouterMulticastDomainId) {
            return transitRouterMulticastDomainId(Output.of(transitRouterMulticastDomainId));
        }

        /**
         * @param vswitchId The ID of the vSwitch.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(Output<String> vswitchId) {
            $.vswitchId = vswitchId;
            return this;
        }

        /**
         * @param vswitchId The ID of the vSwitch.
         * 
         * @return builder
         * 
         */
        public Builder vswitchId(String vswitchId) {
            return vswitchId(Output.of(vswitchId));
        }

        public TransitRouterMulticastDomainAssociationArgs build() {
            $.transitRouterAttachmentId = Objects.requireNonNull($.transitRouterAttachmentId, "expected parameter 'transitRouterAttachmentId' to be non-null");
            $.transitRouterMulticastDomainId = Objects.requireNonNull($.transitRouterMulticastDomainId, "expected parameter 'transitRouterMulticastDomainId' to be non-null");
            $.vswitchId = Objects.requireNonNull($.vswitchId, "expected parameter 'vswitchId' to be non-null");
            return $;
        }
    }

}
