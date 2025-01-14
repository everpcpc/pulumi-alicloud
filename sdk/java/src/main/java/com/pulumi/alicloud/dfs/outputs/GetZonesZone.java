// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dfs.outputs;

import com.pulumi.alicloud.dfs.outputs.GetZonesZoneOption;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetZonesZone {
    /**
     * @return A list of available configurations of the Zone.
     * 
     */
    private List<GetZonesZoneOption> options;
    /**
     * @return The zone ID.
     * 
     */
    private String zoneId;

    private GetZonesZone() {}
    /**
     * @return A list of available configurations of the Zone.
     * 
     */
    public List<GetZonesZoneOption> options() {
        return this.options;
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

    public static Builder builder(GetZonesZone defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetZonesZoneOption> options;
        private String zoneId;
        public Builder() {}
        public Builder(GetZonesZone defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.options = defaults.options;
    	      this.zoneId = defaults.zoneId;
        }

        @CustomType.Setter
        public Builder options(List<GetZonesZoneOption> options) {
            this.options = Objects.requireNonNull(options);
            return this;
        }
        public Builder options(GetZonesZoneOption... options) {
            return options(List.of(options));
        }
        @CustomType.Setter
        public Builder zoneId(String zoneId) {
            this.zoneId = Objects.requireNonNull(zoneId);
            return this;
        }
        public GetZonesZone build() {
            final var o = new GetZonesZone();
            o.options = options;
            o.zoneId = zoneId;
            return o;
        }
    }
}
