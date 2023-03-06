// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTransitRouterCidrsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTransitRouterCidrsPlainArgs Empty = new GetTransitRouterCidrsPlainArgs();

    /**
     * A list of Cen Transit Router Cidr IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Cen Transit Router Cidr IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Transit Router Cidr name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Transit Router Cidr name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    @Import(name="outputFile")
    private @Nullable String outputFile;

    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The ID of the transit router cidr.
     * 
     */
    @Import(name="transitRouterCidrId")
    private @Nullable String transitRouterCidrId;

    /**
     * @return The ID of the transit router cidr.
     * 
     */
    public Optional<String> transitRouterCidrId() {
        return Optional.ofNullable(this.transitRouterCidrId);
    }

    /**
     * The ID of the transit router.
     * 
     */
    @Import(name="transitRouterId", required=true)
    private String transitRouterId;

    /**
     * @return The ID of the transit router.
     * 
     */
    public String transitRouterId() {
        return this.transitRouterId;
    }

    private GetTransitRouterCidrsPlainArgs() {}

    private GetTransitRouterCidrsPlainArgs(GetTransitRouterCidrsPlainArgs $) {
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.transitRouterCidrId = $.transitRouterCidrId;
        this.transitRouterId = $.transitRouterId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTransitRouterCidrsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTransitRouterCidrsPlainArgs $;

        public Builder() {
            $ = new GetTransitRouterCidrsPlainArgs();
        }

        public Builder(GetTransitRouterCidrsPlainArgs defaults) {
            $ = new GetTransitRouterCidrsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Cen Transit Router Cidr IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Cen Transit Router Cidr IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Transit Router Cidr name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param transitRouterCidrId The ID of the transit router cidr.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterCidrId(@Nullable String transitRouterCidrId) {
            $.transitRouterCidrId = transitRouterCidrId;
            return this;
        }

        /**
         * @param transitRouterId The ID of the transit router.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterId(String transitRouterId) {
            $.transitRouterId = transitRouterId;
            return this;
        }

        public GetTransitRouterCidrsPlainArgs build() {
            $.transitRouterId = Objects.requireNonNull($.transitRouterId, "expected parameter 'transitRouterId' to be non-null");
            return $;
        }
    }

}