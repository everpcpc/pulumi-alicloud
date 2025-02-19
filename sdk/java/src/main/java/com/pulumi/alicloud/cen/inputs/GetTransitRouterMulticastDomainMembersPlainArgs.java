// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTransitRouterMulticastDomainMembersPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTransitRouterMulticastDomainMembersPlainArgs Empty = new GetTransitRouterMulticastDomainMembersPlainArgs();

    /**
     * A list of Transit Router Multicast Domain Member IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Transit Router Multicast Domain Member IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The ID of the ENI.
     * 
     */
    @Import(name="networkInterfaceId")
    private @Nullable String networkInterfaceId;

    /**
     * @return The ID of the ENI.
     * 
     */
    public Optional<String> networkInterfaceId() {
        return Optional.ofNullable(this.networkInterfaceId);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The ID of the multicast domain to which the multicast member belongs.
     * 
     */
    @Import(name="transitRouterMulticastDomainId", required=true)
    private String transitRouterMulticastDomainId;

    /**
     * @return The ID of the multicast domain to which the multicast member belongs.
     * 
     */
    public String transitRouterMulticastDomainId() {
        return this.transitRouterMulticastDomainId;
    }

    private GetTransitRouterMulticastDomainMembersPlainArgs() {}

    private GetTransitRouterMulticastDomainMembersPlainArgs(GetTransitRouterMulticastDomainMembersPlainArgs $) {
        this.ids = $.ids;
        this.networkInterfaceId = $.networkInterfaceId;
        this.outputFile = $.outputFile;
        this.transitRouterMulticastDomainId = $.transitRouterMulticastDomainId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTransitRouterMulticastDomainMembersPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTransitRouterMulticastDomainMembersPlainArgs $;

        public Builder() {
            $ = new GetTransitRouterMulticastDomainMembersPlainArgs();
        }

        public Builder(GetTransitRouterMulticastDomainMembersPlainArgs defaults) {
            $ = new GetTransitRouterMulticastDomainMembersPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Transit Router Multicast Domain Member IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Transit Router Multicast Domain Member IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param networkInterfaceId The ID of the ENI.
         * 
         * @return builder
         * 
         */
        public Builder networkInterfaceId(@Nullable String networkInterfaceId) {
            $.networkInterfaceId = networkInterfaceId;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param transitRouterMulticastDomainId The ID of the multicast domain to which the multicast member belongs.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterMulticastDomainId(String transitRouterMulticastDomainId) {
            $.transitRouterMulticastDomainId = transitRouterMulticastDomainId;
            return this;
        }

        public GetTransitRouterMulticastDomainMembersPlainArgs build() {
            $.transitRouterMulticastDomainId = Objects.requireNonNull($.transitRouterMulticastDomainId, "expected parameter 'transitRouterMulticastDomainId' to be non-null");
            return $;
        }
    }

}
