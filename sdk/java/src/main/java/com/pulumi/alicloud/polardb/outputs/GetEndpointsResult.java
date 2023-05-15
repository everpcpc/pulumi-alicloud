// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb.outputs;

import com.pulumi.alicloud.polardb.outputs.GetEndpointsEndpoint;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetEndpointsResult {
    private String dbClusterId;
    /**
     * @return The endpoint ID.
     * 
     */
    private @Nullable String dbEndpointId;
    /**
     * @return A list of PolarDB cluster endpoints. Each element contains the following attributes:
     * 
     */
    private List<GetEndpointsEndpoint> endpoints;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;

    private GetEndpointsResult() {}
    public String dbClusterId() {
        return this.dbClusterId;
    }
    /**
     * @return The endpoint ID.
     * 
     */
    public Optional<String> dbEndpointId() {
        return Optional.ofNullable(this.dbEndpointId);
    }
    /**
     * @return A list of PolarDB cluster endpoints. Each element contains the following attributes:
     * 
     */
    public List<GetEndpointsEndpoint> endpoints() {
        return this.endpoints;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEndpointsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String dbClusterId;
        private @Nullable String dbEndpointId;
        private List<GetEndpointsEndpoint> endpoints;
        private String id;
        public Builder() {}
        public Builder(GetEndpointsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dbClusterId = defaults.dbClusterId;
    	      this.dbEndpointId = defaults.dbEndpointId;
    	      this.endpoints = defaults.endpoints;
    	      this.id = defaults.id;
        }

        @CustomType.Setter
        public Builder dbClusterId(String dbClusterId) {
            this.dbClusterId = Objects.requireNonNull(dbClusterId);
            return this;
        }
        @CustomType.Setter
        public Builder dbEndpointId(@Nullable String dbEndpointId) {
            this.dbEndpointId = dbEndpointId;
            return this;
        }
        @CustomType.Setter
        public Builder endpoints(List<GetEndpointsEndpoint> endpoints) {
            this.endpoints = Objects.requireNonNull(endpoints);
            return this;
        }
        public Builder endpoints(GetEndpointsEndpoint... endpoints) {
            return endpoints(List.of(endpoints));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public GetEndpointsResult build() {
            final var o = new GetEndpointsResult();
            o.dbClusterId = dbClusterId;
            o.dbEndpointId = dbEndpointId;
            o.endpoints = endpoints;
            o.id = id;
            return o;
        }
    }
}
