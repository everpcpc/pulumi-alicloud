// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.privatelink;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VpcEndpointServiceResourceArgs extends com.pulumi.resources.ResourceArgs {

    public static final VpcEndpointServiceResourceArgs Empty = new VpcEndpointServiceResourceArgs();

    /**
     * The dry run.
     * 
     */
    @Import(name="dryRun")
    private @Nullable Output<Boolean> dryRun;

    /**
     * @return The dry run.
     * 
     */
    public Optional<Output<Boolean>> dryRun() {
        return Optional.ofNullable(this.dryRun);
    }

    /**
     * The ID of Resource.
     * 
     */
    @Import(name="resourceId", required=true)
    private Output<String> resourceId;

    /**
     * @return The ID of Resource.
     * 
     */
    public Output<String> resourceId() {
        return this.resourceId;
    }

    /**
     * The Type of Resource.
     * 
     */
    @Import(name="resourceType", required=true)
    private Output<String> resourceType;

    /**
     * @return The Type of Resource.
     * 
     */
    public Output<String> resourceType() {
        return this.resourceType;
    }

    /**
     * The ID of Vpc Endpoint Service.
     * 
     */
    @Import(name="serviceId", required=true)
    private Output<String> serviceId;

    /**
     * @return The ID of Vpc Endpoint Service.
     * 
     */
    public Output<String> serviceId() {
        return this.serviceId;
    }

    private VpcEndpointServiceResourceArgs() {}

    private VpcEndpointServiceResourceArgs(VpcEndpointServiceResourceArgs $) {
        this.dryRun = $.dryRun;
        this.resourceId = $.resourceId;
        this.resourceType = $.resourceType;
        this.serviceId = $.serviceId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpcEndpointServiceResourceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpcEndpointServiceResourceArgs $;

        public Builder() {
            $ = new VpcEndpointServiceResourceArgs();
        }

        public Builder(VpcEndpointServiceResourceArgs defaults) {
            $ = new VpcEndpointServiceResourceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dryRun The dry run.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(@Nullable Output<Boolean> dryRun) {
            $.dryRun = dryRun;
            return this;
        }

        /**
         * @param dryRun The dry run.
         * 
         * @return builder
         * 
         */
        public Builder dryRun(Boolean dryRun) {
            return dryRun(Output.of(dryRun));
        }

        /**
         * @param resourceId The ID of Resource.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(Output<String> resourceId) {
            $.resourceId = resourceId;
            return this;
        }

        /**
         * @param resourceId The ID of Resource.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(String resourceId) {
            return resourceId(Output.of(resourceId));
        }

        /**
         * @param resourceType The Type of Resource.
         * 
         * @return builder
         * 
         */
        public Builder resourceType(Output<String> resourceType) {
            $.resourceType = resourceType;
            return this;
        }

        /**
         * @param resourceType The Type of Resource.
         * 
         * @return builder
         * 
         */
        public Builder resourceType(String resourceType) {
            return resourceType(Output.of(resourceType));
        }

        /**
         * @param serviceId The ID of Vpc Endpoint Service.
         * 
         * @return builder
         * 
         */
        public Builder serviceId(Output<String> serviceId) {
            $.serviceId = serviceId;
            return this;
        }

        /**
         * @param serviceId The ID of Vpc Endpoint Service.
         * 
         * @return builder
         * 
         */
        public Builder serviceId(String serviceId) {
            return serviceId(Output.of(serviceId));
        }

        public VpcEndpointServiceResourceArgs build() {
            $.resourceId = Objects.requireNonNull($.resourceId, "expected parameter 'resourceId' to be non-null");
            $.resourceType = Objects.requireNonNull($.resourceType, "expected parameter 'resourceType' to be non-null");
            $.serviceId = Objects.requireNonNull($.serviceId, "expected parameter 'serviceId' to be non-null");
            return $;
        }
    }

}
