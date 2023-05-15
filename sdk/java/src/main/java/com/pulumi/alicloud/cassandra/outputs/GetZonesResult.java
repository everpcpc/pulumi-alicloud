// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cassandra.outputs;

import com.pulumi.alicloud.cassandra.outputs.GetZonesZone;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetZonesResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of zone IDs.
     * 
     */
    private List<String> ids;
    private @Nullable Boolean multi;
    private @Nullable String outputFile;
    /**
     * @return A list of availability zones. Each element contains the following attributes:
     * 
     */
    private List<GetZonesZone> zones;

    private GetZonesResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of zone IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<Boolean> multi() {
        return Optional.ofNullable(this.multi);
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return A list of availability zones. Each element contains the following attributes:
     * 
     */
    public List<GetZonesZone> zones() {
        return this.zones;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetZonesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> ids;
        private @Nullable Boolean multi;
        private @Nullable String outputFile;
        private List<GetZonesZone> zones;
        public Builder() {}
        public Builder(GetZonesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.multi = defaults.multi;
    	      this.outputFile = defaults.outputFile;
    	      this.zones = defaults.zones;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ids(List<String> ids) {
            this.ids = Objects.requireNonNull(ids);
            return this;
        }
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }
        @CustomType.Setter
        public Builder multi(@Nullable Boolean multi) {
            this.multi = multi;
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder zones(List<GetZonesZone> zones) {
            this.zones = Objects.requireNonNull(zones);
            return this;
        }
        public Builder zones(GetZonesZone... zones) {
            return zones(List.of(zones));
        }
        public GetZonesResult build() {
            final var o = new GetZonesResult();
            o.id = id;
            o.ids = ids;
            o.multi = multi;
            o.outputFile = outputFile;
            o.zones = zones;
            return o;
        }
    }
}
