// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.actiontrail.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetGlobalEventsStorageRegionResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String storageRegion;

    private GetGlobalEventsStorageRegionResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String storageRegion() {
        return this.storageRegion;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGlobalEventsStorageRegionResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String storageRegion;
        public Builder() {}
        public Builder(GetGlobalEventsStorageRegionResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.storageRegion = defaults.storageRegion;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder storageRegion(String storageRegion) {
            this.storageRegion = Objects.requireNonNull(storageRegion);
            return this;
        }
        public GetGlobalEventsStorageRegionResult build() {
            final var o = new GetGlobalEventsStorageRegionResult();
            o.id = id;
            o.storageRegion = storageRegion;
            return o;
        }
    }
}