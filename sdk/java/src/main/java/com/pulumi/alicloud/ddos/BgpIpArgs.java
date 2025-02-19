// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ddos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BgpIpArgs extends com.pulumi.resources.ResourceArgs {

    public static final BgpIpArgs Empty = new BgpIpArgs();

    /**
     * The ID of the native protection enterprise instance to be operated.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return The ID of the native protection enterprise instance to be operated.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     * The IP address.
     * 
     */
    @Import(name="ip", required=true)
    private Output<String> ip;

    /**
     * @return The IP address.
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }

    /**
     * The ID of the resource group.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    private BgpIpArgs() {}

    private BgpIpArgs(BgpIpArgs $) {
        this.instanceId = $.instanceId;
        this.ip = $.ip;
        this.resourceGroupId = $.resourceGroupId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BgpIpArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BgpIpArgs $;

        public Builder() {
            $ = new BgpIpArgs();
        }

        public Builder(BgpIpArgs defaults) {
            $ = new BgpIpArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param instanceId The ID of the native protection enterprise instance to be operated.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The ID of the native protection enterprise instance to be operated.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param ip The IP address.
         * 
         * @return builder
         * 
         */
        public Builder ip(Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip The IP address.
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId The ID of the resource group.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        public BgpIpArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            $.ip = Objects.requireNonNull($.ip, "expected parameter 'ip' to be non-null");
            return $;
        }
    }

}
