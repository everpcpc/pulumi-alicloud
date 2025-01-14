// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.outputs;

import com.pulumi.alicloud.ga.outputs.GetCustomRoutingPortMappingsCustomRoutingPortMapping;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetCustomRoutingPortMappingsResult {
    /**
     * @return The ID of the GA instance.
     * 
     */
    private String acceleratorId;
    /**
     * @return A list of Custom Routing Port Mappings. Each element contains the following attributes:
     * 
     */
    private List<GetCustomRoutingPortMappingsCustomRoutingPortMapping> customRoutingPortMappings;
    /**
     * @return The ID of the endpoint group.
     * 
     */
    private @Nullable String endpointGroupId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The ID of the listener.
     * 
     */
    private @Nullable String listenerId;
    private @Nullable String outputFile;
    private @Nullable Integer pageNumber;
    private @Nullable Integer pageSize;
    /**
     * @return The access policy of traffic for the backend instance.
     * 
     */
    private @Nullable String status;

    private GetCustomRoutingPortMappingsResult() {}
    /**
     * @return The ID of the GA instance.
     * 
     */
    public String acceleratorId() {
        return this.acceleratorId;
    }
    /**
     * @return A list of Custom Routing Port Mappings. Each element contains the following attributes:
     * 
     */
    public List<GetCustomRoutingPortMappingsCustomRoutingPortMapping> customRoutingPortMappings() {
        return this.customRoutingPortMappings;
    }
    /**
     * @return The ID of the endpoint group.
     * 
     */
    public Optional<String> endpointGroupId() {
        return Optional.ofNullable(this.endpointGroupId);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The ID of the listener.
     * 
     */
    public Optional<String> listenerId() {
        return Optional.ofNullable(this.listenerId);
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Optional<Integer> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }
    public Optional<Integer> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }
    /**
     * @return The access policy of traffic for the backend instance.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCustomRoutingPortMappingsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String acceleratorId;
        private List<GetCustomRoutingPortMappingsCustomRoutingPortMapping> customRoutingPortMappings;
        private @Nullable String endpointGroupId;
        private String id;
        private @Nullable String listenerId;
        private @Nullable String outputFile;
        private @Nullable Integer pageNumber;
        private @Nullable Integer pageSize;
        private @Nullable String status;
        public Builder() {}
        public Builder(GetCustomRoutingPortMappingsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.acceleratorId = defaults.acceleratorId;
    	      this.customRoutingPortMappings = defaults.customRoutingPortMappings;
    	      this.endpointGroupId = defaults.endpointGroupId;
    	      this.id = defaults.id;
    	      this.listenerId = defaults.listenerId;
    	      this.outputFile = defaults.outputFile;
    	      this.pageNumber = defaults.pageNumber;
    	      this.pageSize = defaults.pageSize;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder acceleratorId(String acceleratorId) {
            this.acceleratorId = Objects.requireNonNull(acceleratorId);
            return this;
        }
        @CustomType.Setter
        public Builder customRoutingPortMappings(List<GetCustomRoutingPortMappingsCustomRoutingPortMapping> customRoutingPortMappings) {
            this.customRoutingPortMappings = Objects.requireNonNull(customRoutingPortMappings);
            return this;
        }
        public Builder customRoutingPortMappings(GetCustomRoutingPortMappingsCustomRoutingPortMapping... customRoutingPortMappings) {
            return customRoutingPortMappings(List.of(customRoutingPortMappings));
        }
        @CustomType.Setter
        public Builder endpointGroupId(@Nullable String endpointGroupId) {
            this.endpointGroupId = endpointGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder listenerId(@Nullable String listenerId) {
            this.listenerId = listenerId;
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder pageNumber(@Nullable Integer pageNumber) {
            this.pageNumber = pageNumber;
            return this;
        }
        @CustomType.Setter
        public Builder pageSize(@Nullable Integer pageSize) {
            this.pageSize = pageSize;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        public GetCustomRoutingPortMappingsResult build() {
            final var o = new GetCustomRoutingPortMappingsResult();
            o.acceleratorId = acceleratorId;
            o.customRoutingPortMappings = customRoutingPortMappings;
            o.endpointGroupId = endpointGroupId;
            o.id = id;
            o.listenerId = listenerId;
            o.outputFile = outputFile;
            o.pageNumber = pageNumber;
            o.pageSize = pageSize;
            o.status = status;
            return o;
        }
    }
}
