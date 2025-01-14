// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTransitRouterMulticastDomainPeerMembersArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTransitRouterMulticastDomainPeerMembersArgs Empty = new GetTransitRouterMulticastDomainPeerMembersArgs();

    /**
     * A list of Cen Transit Router Multicast Domain Peer Member IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Cen Transit Router Multicast Domain Peer Member IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The IDs of the inter-region multicast domains.
     * 
     */
    @Import(name="peerTransitRouterMulticastDomains")
    private @Nullable Output<List<String>> peerTransitRouterMulticastDomains;

    /**
     * @return The IDs of the inter-region multicast domains.
     * 
     */
    public Optional<Output<List<String>>> peerTransitRouterMulticastDomains() {
        return Optional.ofNullable(this.peerTransitRouterMulticastDomains);
    }

    /**
     * The ID of the resource associated with the multicast resource.
     * 
     */
    @Import(name="resourceId")
    private @Nullable Output<String> resourceId;

    /**
     * @return The ID of the resource associated with the multicast resource.
     * 
     */
    public Optional<Output<String>> resourceId() {
        return Optional.ofNullable(this.resourceId);
    }

    /**
     * The type of the multicast resource. Valid values:
     * * VPC: queries multicast resources by VPC.
     * * TR: queries multicast resources that are also deployed in a different region.
     * 
     */
    @Import(name="resourceType")
    private @Nullable Output<String> resourceType;

    /**
     * @return The type of the multicast resource. Valid values:
     * * VPC: queries multicast resources by VPC.
     * * TR: queries multicast resources that are also deployed in a different region.
     * 
     */
    public Optional<Output<String>> resourceType() {
        return Optional.ofNullable(this.resourceType);
    }

    /**
     * The ID of the network instance connection.
     * 
     */
    @Import(name="transitRouterAttachmentId")
    private @Nullable Output<String> transitRouterAttachmentId;

    /**
     * @return The ID of the network instance connection.
     * 
     */
    public Optional<Output<String>> transitRouterAttachmentId() {
        return Optional.ofNullable(this.transitRouterAttachmentId);
    }

    /**
     * The ID of the multicast domain to which the multicast member belongs.
     * 
     */
    @Import(name="transitRouterMulticastDomainId", required=true)
    private Output<String> transitRouterMulticastDomainId;

    /**
     * @return The ID of the multicast domain to which the multicast member belongs.
     * 
     */
    public Output<String> transitRouterMulticastDomainId() {
        return this.transitRouterMulticastDomainId;
    }

    private GetTransitRouterMulticastDomainPeerMembersArgs() {}

    private GetTransitRouterMulticastDomainPeerMembersArgs(GetTransitRouterMulticastDomainPeerMembersArgs $) {
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.peerTransitRouterMulticastDomains = $.peerTransitRouterMulticastDomains;
        this.resourceId = $.resourceId;
        this.resourceType = $.resourceType;
        this.transitRouterAttachmentId = $.transitRouterAttachmentId;
        this.transitRouterMulticastDomainId = $.transitRouterMulticastDomainId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTransitRouterMulticastDomainPeerMembersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTransitRouterMulticastDomainPeerMembersArgs $;

        public Builder() {
            $ = new GetTransitRouterMulticastDomainPeerMembersArgs();
        }

        public Builder(GetTransitRouterMulticastDomainPeerMembersArgs defaults) {
            $ = new GetTransitRouterMulticastDomainPeerMembersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Cen Transit Router Multicast Domain Peer Member IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Cen Transit Router Multicast Domain Peer Member IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Cen Transit Router Multicast Domain Peer Member IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        /**
         * @param peerTransitRouterMulticastDomains The IDs of the inter-region multicast domains.
         * 
         * @return builder
         * 
         */
        public Builder peerTransitRouterMulticastDomains(@Nullable Output<List<String>> peerTransitRouterMulticastDomains) {
            $.peerTransitRouterMulticastDomains = peerTransitRouterMulticastDomains;
            return this;
        }

        /**
         * @param peerTransitRouterMulticastDomains The IDs of the inter-region multicast domains.
         * 
         * @return builder
         * 
         */
        public Builder peerTransitRouterMulticastDomains(List<String> peerTransitRouterMulticastDomains) {
            return peerTransitRouterMulticastDomains(Output.of(peerTransitRouterMulticastDomains));
        }

        /**
         * @param peerTransitRouterMulticastDomains The IDs of the inter-region multicast domains.
         * 
         * @return builder
         * 
         */
        public Builder peerTransitRouterMulticastDomains(String... peerTransitRouterMulticastDomains) {
            return peerTransitRouterMulticastDomains(List.of(peerTransitRouterMulticastDomains));
        }

        /**
         * @param resourceId The ID of the resource associated with the multicast resource.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(@Nullable Output<String> resourceId) {
            $.resourceId = resourceId;
            return this;
        }

        /**
         * @param resourceId The ID of the resource associated with the multicast resource.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(String resourceId) {
            return resourceId(Output.of(resourceId));
        }

        /**
         * @param resourceType The type of the multicast resource. Valid values:
         * * VPC: queries multicast resources by VPC.
         * * TR: queries multicast resources that are also deployed in a different region.
         * 
         * @return builder
         * 
         */
        public Builder resourceType(@Nullable Output<String> resourceType) {
            $.resourceType = resourceType;
            return this;
        }

        /**
         * @param resourceType The type of the multicast resource. Valid values:
         * * VPC: queries multicast resources by VPC.
         * * TR: queries multicast resources that are also deployed in a different region.
         * 
         * @return builder
         * 
         */
        public Builder resourceType(String resourceType) {
            return resourceType(Output.of(resourceType));
        }

        /**
         * @param transitRouterAttachmentId The ID of the network instance connection.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterAttachmentId(@Nullable Output<String> transitRouterAttachmentId) {
            $.transitRouterAttachmentId = transitRouterAttachmentId;
            return this;
        }

        /**
         * @param transitRouterAttachmentId The ID of the network instance connection.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterAttachmentId(String transitRouterAttachmentId) {
            return transitRouterAttachmentId(Output.of(transitRouterAttachmentId));
        }

        /**
         * @param transitRouterMulticastDomainId The ID of the multicast domain to which the multicast member belongs.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterMulticastDomainId(Output<String> transitRouterMulticastDomainId) {
            $.transitRouterMulticastDomainId = transitRouterMulticastDomainId;
            return this;
        }

        /**
         * @param transitRouterMulticastDomainId The ID of the multicast domain to which the multicast member belongs.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterMulticastDomainId(String transitRouterMulticastDomainId) {
            return transitRouterMulticastDomainId(Output.of(transitRouterMulticastDomainId));
        }

        public GetTransitRouterMulticastDomainPeerMembersArgs build() {
            $.transitRouterMulticastDomainId = Objects.requireNonNull($.transitRouterMulticastDomainId, "expected parameter 'transitRouterMulticastDomainId' to be non-null");
            return $;
        }
    }

}
