// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cddc.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetZonesZone {
    /**
     * @return The ID of the zone.
     * 
     */
    private String id;
    /**
     * @return The ID of the region.
     * 
     */
    private String regionId;
    /**
     * @return The ID of the zone.
     * 
     */
    private String zoneId;

    private GetZonesZone() {}
    /**
     * @return The ID of the zone.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The ID of the region.
     * 
     */
    public String regionId() {
        return this.regionId;
    }
    /**
     * @return The ID of the zone.
     * 
     */
    public String zoneId() {
        return this.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetZonesZone defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String regionId;
        private String zoneId;
        public Builder() {}
        public Builder(GetZonesZone defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.regionId = defaults.regionId;
    	      this.zoneId = defaults.zoneId;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder regionId(String regionId) {
            this.regionId = Objects.requireNonNull(regionId);
            return this;
        }
        @CustomType.Setter
        public Builder zoneId(String zoneId) {
            this.zoneId = Objects.requireNonNull(zoneId);
            return this;
        }
        public GetZonesZone build() {
            final var o = new GetZonesZone();
            o.id = id;
            o.regionId = regionId;
            o.zoneId = zoneId;
            return o;
        }
    }
}