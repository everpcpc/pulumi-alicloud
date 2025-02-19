// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.resourcemanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ResourceShareArgs extends com.pulumi.resources.ResourceArgs {

    public static final ResourceShareArgs Empty = new ResourceShareArgs();

    /**
     * The name of resource share.
     * 
     */
    @Import(name="resourceShareName", required=true)
    private Output<String> resourceShareName;

    /**
     * @return The name of resource share.
     * 
     */
    public Output<String> resourceShareName() {
        return this.resourceShareName;
    }

    private ResourceShareArgs() {}

    private ResourceShareArgs(ResourceShareArgs $) {
        this.resourceShareName = $.resourceShareName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ResourceShareArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ResourceShareArgs $;

        public Builder() {
            $ = new ResourceShareArgs();
        }

        public Builder(ResourceShareArgs defaults) {
            $ = new ResourceShareArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param resourceShareName The name of resource share.
         * 
         * @return builder
         * 
         */
        public Builder resourceShareName(Output<String> resourceShareName) {
            $.resourceShareName = resourceShareName;
            return this;
        }

        /**
         * @param resourceShareName The name of resource share.
         * 
         * @return builder
         * 
         */
        public Builder resourceShareName(String resourceShareName) {
            return resourceShareName(Output.of(resourceShareName));
        }

        public ResourceShareArgs build() {
            $.resourceShareName = Objects.requireNonNull($.resourceShareName, "expected parameter 'resourceShareName' to be non-null");
            return $;
        }
    }

}
