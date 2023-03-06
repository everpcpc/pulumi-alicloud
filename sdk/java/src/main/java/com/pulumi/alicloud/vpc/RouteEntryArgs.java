// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RouteEntryArgs extends com.pulumi.resources.ResourceArgs {

    public static final RouteEntryArgs Empty = new RouteEntryArgs();

    /**
     * The RouteEntry&#39;s target network segment.
     * 
     */
    @Import(name="destinationCidrblock")
    private @Nullable Output<String> destinationCidrblock;

    /**
     * @return The RouteEntry&#39;s target network segment.
     * 
     */
    public Optional<Output<String>> destinationCidrblock() {
        return Optional.ofNullable(this.destinationCidrblock);
    }

    /**
     * The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The route entry&#39;s next hop. ECS instance ID or VPC router interface ID.
     * 
     */
    @Import(name="nexthopId")
    private @Nullable Output<String> nexthopId;

    /**
     * @return The route entry&#39;s next hop. ECS instance ID or VPC router interface ID.
     * 
     */
    public Optional<Output<String>> nexthopId() {
        return Optional.ofNullable(this.nexthopId);
    }

    /**
     * The next hop type. Available values:
     * 
     */
    @Import(name="nexthopType")
    private @Nullable Output<String> nexthopType;

    /**
     * @return The next hop type. Available values:
     * 
     */
    public Optional<Output<String>> nexthopType() {
        return Optional.ofNullable(this.nexthopType);
    }

    /**
     * The ID of the route table.
     * 
     */
    @Import(name="routeTableId", required=true)
    private Output<String> routeTableId;

    /**
     * @return The ID of the route table.
     * 
     */
    public Output<String> routeTableId() {
        return this.routeTableId;
    }

    /**
     * This argument has been deprecated. Please use other arguments to launch a custom route entry.
     * 
     * @deprecated
     * Attribute router_id has been deprecated and suggest removing it from your template.
     * 
     */
    @Deprecated /* Attribute router_id has been deprecated and suggest removing it from your template. */
    @Import(name="routerId")
    private @Nullable Output<String> routerId;

    /**
     * @return This argument has been deprecated. Please use other arguments to launch a custom route entry.
     * 
     * @deprecated
     * Attribute router_id has been deprecated and suggest removing it from your template.
     * 
     */
    @Deprecated /* Attribute router_id has been deprecated and suggest removing it from your template. */
    public Optional<Output<String>> routerId() {
        return Optional.ofNullable(this.routerId);
    }

    private RouteEntryArgs() {}

    private RouteEntryArgs(RouteEntryArgs $) {
        this.destinationCidrblock = $.destinationCidrblock;
        this.name = $.name;
        this.nexthopId = $.nexthopId;
        this.nexthopType = $.nexthopType;
        this.routeTableId = $.routeTableId;
        this.routerId = $.routerId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RouteEntryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RouteEntryArgs $;

        public Builder() {
            $ = new RouteEntryArgs();
        }

        public Builder(RouteEntryArgs defaults) {
            $ = new RouteEntryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param destinationCidrblock The RouteEntry&#39;s target network segment.
         * 
         * @return builder
         * 
         */
        public Builder destinationCidrblock(@Nullable Output<String> destinationCidrblock) {
            $.destinationCidrblock = destinationCidrblock;
            return this;
        }

        /**
         * @param destinationCidrblock The RouteEntry&#39;s target network segment.
         * 
         * @return builder
         * 
         */
        public Builder destinationCidrblock(String destinationCidrblock) {
            return destinationCidrblock(Output.of(destinationCidrblock));
        }

        /**
         * @param name The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param nexthopId The route entry&#39;s next hop. ECS instance ID or VPC router interface ID.
         * 
         * @return builder
         * 
         */
        public Builder nexthopId(@Nullable Output<String> nexthopId) {
            $.nexthopId = nexthopId;
            return this;
        }

        /**
         * @param nexthopId The route entry&#39;s next hop. ECS instance ID or VPC router interface ID.
         * 
         * @return builder
         * 
         */
        public Builder nexthopId(String nexthopId) {
            return nexthopId(Output.of(nexthopId));
        }

        /**
         * @param nexthopType The next hop type. Available values:
         * 
         * @return builder
         * 
         */
        public Builder nexthopType(@Nullable Output<String> nexthopType) {
            $.nexthopType = nexthopType;
            return this;
        }

        /**
         * @param nexthopType The next hop type. Available values:
         * 
         * @return builder
         * 
         */
        public Builder nexthopType(String nexthopType) {
            return nexthopType(Output.of(nexthopType));
        }

        /**
         * @param routeTableId The ID of the route table.
         * 
         * @return builder
         * 
         */
        public Builder routeTableId(Output<String> routeTableId) {
            $.routeTableId = routeTableId;
            return this;
        }

        /**
         * @param routeTableId The ID of the route table.
         * 
         * @return builder
         * 
         */
        public Builder routeTableId(String routeTableId) {
            return routeTableId(Output.of(routeTableId));
        }

        /**
         * @param routerId This argument has been deprecated. Please use other arguments to launch a custom route entry.
         * 
         * @return builder
         * 
         * @deprecated
         * Attribute router_id has been deprecated and suggest removing it from your template.
         * 
         */
        @Deprecated /* Attribute router_id has been deprecated and suggest removing it from your template. */
        public Builder routerId(@Nullable Output<String> routerId) {
            $.routerId = routerId;
            return this;
        }

        /**
         * @param routerId This argument has been deprecated. Please use other arguments to launch a custom route entry.
         * 
         * @return builder
         * 
         * @deprecated
         * Attribute router_id has been deprecated and suggest removing it from your template.
         * 
         */
        @Deprecated /* Attribute router_id has been deprecated and suggest removing it from your template. */
        public Builder routerId(String routerId) {
            return routerId(Output.of(routerId));
        }

        public RouteEntryArgs build() {
            $.routeTableId = Objects.requireNonNull($.routeTableId, "expected parameter 'routeTableId' to be non-null");
            return $;
        }
    }

}