// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudstoragegateway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetStocksStock {
    /**
     * @return A list of available gateway class in this Zone ID.
     * 
     */
    private List<String> availableGatewayClasses;
    /**
     * @return The Zone ID.
     * 
     */
    private String zoneId;

    private GetStocksStock() {}
    /**
     * @return A list of available gateway class in this Zone ID.
     * 
     */
    public List<String> availableGatewayClasses() {
        return this.availableGatewayClasses;
    }
    /**
     * @return The Zone ID.
     * 
     */
    public String zoneId() {
        return this.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetStocksStock defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> availableGatewayClasses;
        private String zoneId;
        public Builder() {}
        public Builder(GetStocksStock defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.availableGatewayClasses = defaults.availableGatewayClasses;
    	      this.zoneId = defaults.zoneId;
        }

        @CustomType.Setter
        public Builder availableGatewayClasses(List<String> availableGatewayClasses) {
            this.availableGatewayClasses = Objects.requireNonNull(availableGatewayClasses);
            return this;
        }
        public Builder availableGatewayClasses(String... availableGatewayClasses) {
            return availableGatewayClasses(List.of(availableGatewayClasses));
        }
        @CustomType.Setter
        public Builder zoneId(String zoneId) {
            this.zoneId = Objects.requireNonNull(zoneId);
            return this;
        }
        public GetStocksStock build() {
            final var o = new GetStocksStock();
            o.availableGatewayClasses = availableGatewayClasses;
            o.zoneId = zoneId;
            return o;
        }
    }
}
