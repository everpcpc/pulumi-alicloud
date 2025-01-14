// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.clickhouse.outputs;

import com.pulumi.alicloud.clickhouse.outputs.GetRegionsRegion;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetRegionsResult {
    private @Nullable Boolean current;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String outputFile;
    private @Nullable String regionId;
    private List<GetRegionsRegion> regions;

    private GetRegionsResult() {}
    public Optional<Boolean> current() {
        return Optional.ofNullable(this.current);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Optional<String> regionId() {
        return Optional.ofNullable(this.regionId);
    }
    public List<GetRegionsRegion> regions() {
        return this.regions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRegionsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean current;
        private String id;
        private @Nullable String outputFile;
        private @Nullable String regionId;
        private List<GetRegionsRegion> regions;
        public Builder() {}
        public Builder(GetRegionsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.current = defaults.current;
    	      this.id = defaults.id;
    	      this.outputFile = defaults.outputFile;
    	      this.regionId = defaults.regionId;
    	      this.regions = defaults.regions;
        }

        @CustomType.Setter
        public Builder current(@Nullable Boolean current) {
            this.current = current;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder regionId(@Nullable String regionId) {
            this.regionId = regionId;
            return this;
        }
        @CustomType.Setter
        public Builder regions(List<GetRegionsRegion> regions) {
            this.regions = Objects.requireNonNull(regions);
            return this;
        }
        public Builder regions(GetRegionsRegion... regions) {
            return regions(List.of(regions));
        }
        public GetRegionsResult build() {
            final var o = new GetRegionsResult();
            o.current = current;
            o.id = id;
            o.outputFile = outputFile;
            o.regionId = regionId;
            o.regions = regions;
            return o;
        }
    }
}
