// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.resourcemanager.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetResourceSharesShare {
    /**
     * @return The ID of the Resource Share.
     * 
     */
    private String id;
    /**
     * @return The ID of the resource share.
     * 
     */
    private String resourceShareId;
    /**
     * @return The name of resource share.
     * 
     */
    private String resourceShareName;
    /**
     * @return The owner of resource share.
     * 
     */
    private String resourceShareOwner;
    /**
     * @return The status of resource share.
     * 
     */
    private String status;

    private GetResourceSharesShare() {}
    /**
     * @return The ID of the Resource Share.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The ID of the resource share.
     * 
     */
    public String resourceShareId() {
        return this.resourceShareId;
    }
    /**
     * @return The name of resource share.
     * 
     */
    public String resourceShareName() {
        return this.resourceShareName;
    }
    /**
     * @return The owner of resource share.
     * 
     */
    public String resourceShareOwner() {
        return this.resourceShareOwner;
    }
    /**
     * @return The status of resource share.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetResourceSharesShare defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String resourceShareId;
        private String resourceShareName;
        private String resourceShareOwner;
        private String status;
        public Builder() {}
        public Builder(GetResourceSharesShare defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.resourceShareId = defaults.resourceShareId;
    	      this.resourceShareName = defaults.resourceShareName;
    	      this.resourceShareOwner = defaults.resourceShareOwner;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder resourceShareId(String resourceShareId) {
            this.resourceShareId = Objects.requireNonNull(resourceShareId);
            return this;
        }
        @CustomType.Setter
        public Builder resourceShareName(String resourceShareName) {
            this.resourceShareName = Objects.requireNonNull(resourceShareName);
            return this;
        }
        @CustomType.Setter
        public Builder resourceShareOwner(String resourceShareOwner) {
            this.resourceShareOwner = Objects.requireNonNull(resourceShareOwner);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetResourceSharesShare build() {
            final var o = new GetResourceSharesShare();
            o.id = id;
            o.resourceShareId = resourceShareId;
            o.resourceShareName = resourceShareName;
            o.resourceShareOwner = resourceShareOwner;
            o.status = status;
            return o;
        }
    }
}
