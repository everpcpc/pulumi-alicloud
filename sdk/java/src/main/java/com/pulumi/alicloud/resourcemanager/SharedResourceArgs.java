// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.resourcemanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class SharedResourceArgs extends com.pulumi.resources.ResourceArgs {

    public static final SharedResourceArgs Empty = new SharedResourceArgs();

    /**
     * The resource ID need shared.
     * 
     */
    @Import(name="resourceId", required=true)
    private Output<String> resourceId;

    /**
     * @return The resource ID need shared.
     * 
     */
    public Output<String> resourceId() {
        return this.resourceId;
    }

    /**
     * The resource share ID of resource manager.
     * 
     */
    @Import(name="resourceShareId", required=true)
    private Output<String> resourceShareId;

    /**
     * @return The resource share ID of resource manager.
     * 
     */
    public Output<String> resourceShareId() {
        return this.resourceShareId;
    }

    /**
     * The resource type of should shared, valid value
     * 
     */
    @Import(name="resourceType", required=true)
    private Output<String> resourceType;

    /**
     * @return The resource type of should shared, valid value
     * 
     */
    public Output<String> resourceType() {
        return this.resourceType;
    }

    private SharedResourceArgs() {}

    private SharedResourceArgs(SharedResourceArgs $) {
        this.resourceId = $.resourceId;
        this.resourceShareId = $.resourceShareId;
        this.resourceType = $.resourceType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SharedResourceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SharedResourceArgs $;

        public Builder() {
            $ = new SharedResourceArgs();
        }

        public Builder(SharedResourceArgs defaults) {
            $ = new SharedResourceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param resourceId The resource ID need shared.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(Output<String> resourceId) {
            $.resourceId = resourceId;
            return this;
        }

        /**
         * @param resourceId The resource ID need shared.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(String resourceId) {
            return resourceId(Output.of(resourceId));
        }

        /**
         * @param resourceShareId The resource share ID of resource manager.
         * 
         * @return builder
         * 
         */
        public Builder resourceShareId(Output<String> resourceShareId) {
            $.resourceShareId = resourceShareId;
            return this;
        }

        /**
         * @param resourceShareId The resource share ID of resource manager.
         * 
         * @return builder
         * 
         */
        public Builder resourceShareId(String resourceShareId) {
            return resourceShareId(Output.of(resourceShareId));
        }

        /**
         * @param resourceType The resource type of should shared, valid value
         * 
         * @return builder
         * 
         */
        public Builder resourceType(Output<String> resourceType) {
            $.resourceType = resourceType;
            return this;
        }

        /**
         * @param resourceType The resource type of should shared, valid value
         * 
         * @return builder
         * 
         */
        public Builder resourceType(String resourceType) {
            return resourceType(Output.of(resourceType));
        }

        public SharedResourceArgs build() {
            $.resourceId = Objects.requireNonNull($.resourceId, "expected parameter 'resourceId' to be non-null");
            $.resourceShareId = Objects.requireNonNull($.resourceShareId, "expected parameter 'resourceShareId' to be non-null");
            $.resourceType = Objects.requireNonNull($.resourceType, "expected parameter 'resourceType' to be non-null");
            return $;
        }
    }

}
