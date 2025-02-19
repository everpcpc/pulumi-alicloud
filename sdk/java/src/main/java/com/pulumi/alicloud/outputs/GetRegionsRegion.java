// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRegionsRegion {
    /**
     * @return ID of the region.
     * 
     */
    private String id;
    /**
     * @return Name of the region in the local language.
     * 
     */
    private String localName;
    private String regionId;

    private GetRegionsRegion() {}
    /**
     * @return ID of the region.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Name of the region in the local language.
     * 
     */
    public String localName() {
        return this.localName;
    }
    public String regionId() {
        return this.regionId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRegionsRegion defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String localName;
        private String regionId;
        public Builder() {}
        public Builder(GetRegionsRegion defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.localName = defaults.localName;
    	      this.regionId = defaults.regionId;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder localName(String localName) {
            this.localName = Objects.requireNonNull(localName);
            return this;
        }
        @CustomType.Setter
        public Builder regionId(String regionId) {
            this.regionId = Objects.requireNonNull(regionId);
            return this;
        }
        public GetRegionsRegion build() {
            final var o = new GetRegionsRegion();
            o.id = id;
            o.localName = localName;
            o.regionId = regionId;
            return o;
        }
    }
}
