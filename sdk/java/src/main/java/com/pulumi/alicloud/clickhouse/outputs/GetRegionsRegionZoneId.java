// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.clickhouse.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRegionsRegionZoneId {
    /**
     * @return Whether to support vpc network.
     * 
     */
    private Boolean vpcEnabled;
    /**
     * @return The zone ID.
     * 
     */
    private String zoneId;

    private GetRegionsRegionZoneId() {}
    /**
     * @return Whether to support vpc network.
     * 
     */
    public Boolean vpcEnabled() {
        return this.vpcEnabled;
    }
    /**
     * @return The zone ID.
     * 
     */
    public String zoneId() {
        return this.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRegionsRegionZoneId defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean vpcEnabled;
        private String zoneId;
        public Builder() {}
        public Builder(GetRegionsRegionZoneId defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.vpcEnabled = defaults.vpcEnabled;
    	      this.zoneId = defaults.zoneId;
        }

        @CustomType.Setter
        public Builder vpcEnabled(Boolean vpcEnabled) {
            this.vpcEnabled = Objects.requireNonNull(vpcEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder zoneId(String zoneId) {
            this.zoneId = Objects.requireNonNull(zoneId);
            return this;
        }
        public GetRegionsRegionZoneId build() {
            final var o = new GetRegionsRegionZoneId();
            o.vpcEnabled = vpcEnabled;
            o.zoneId = zoneId;
            return o;
        }
    }
}
