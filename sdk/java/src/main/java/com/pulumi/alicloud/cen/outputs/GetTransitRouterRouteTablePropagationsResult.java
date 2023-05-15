// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.outputs;

import com.pulumi.alicloud.cen.outputs.GetTransitRouterRouteTablePropagationsPropagation;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetTransitRouterRouteTablePropagationsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of CEN Transit Router Route Table Association IDs.
     * 
     */
    private List<String> ids;
    private @Nullable String outputFile;
    /**
     * @return A list of CEN Transit Router Route Table Propagations. Each element contains the following attributes:
     * 
     */
    private List<GetTransitRouterRouteTablePropagationsPropagation> propagations;
    /**
     * @return The status of the route table.
     * 
     */
    private @Nullable String status;
    /**
     * @return ID of the transit router route table.
     * 
     */
    private String transitRouterRouteTableId;

    private GetTransitRouterRouteTablePropagationsResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of CEN Transit Router Route Table Association IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return A list of CEN Transit Router Route Table Propagations. Each element contains the following attributes:
     * 
     */
    public List<GetTransitRouterRouteTablePropagationsPropagation> propagations() {
        return this.propagations;
    }
    /**
     * @return The status of the route table.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    /**
     * @return ID of the transit router route table.
     * 
     */
    public String transitRouterRouteTableId() {
        return this.transitRouterRouteTableId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTransitRouterRouteTablePropagationsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> ids;
        private @Nullable String outputFile;
        private List<GetTransitRouterRouteTablePropagationsPropagation> propagations;
        private @Nullable String status;
        private String transitRouterRouteTableId;
        public Builder() {}
        public Builder(GetTransitRouterRouteTablePropagationsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.outputFile = defaults.outputFile;
    	      this.propagations = defaults.propagations;
    	      this.status = defaults.status;
    	      this.transitRouterRouteTableId = defaults.transitRouterRouteTableId;
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
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder propagations(List<GetTransitRouterRouteTablePropagationsPropagation> propagations) {
            this.propagations = Objects.requireNonNull(propagations);
            return this;
        }
        public Builder propagations(GetTransitRouterRouteTablePropagationsPropagation... propagations) {
            return propagations(List.of(propagations));
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder transitRouterRouteTableId(String transitRouterRouteTableId) {
            this.transitRouterRouteTableId = Objects.requireNonNull(transitRouterRouteTableId);
            return this;
        }
        public GetTransitRouterRouteTablePropagationsResult build() {
            final var o = new GetTransitRouterRouteTablePropagationsResult();
            o.id = id;
            o.ids = ids;
            o.outputFile = outputFile;
            o.propagations = propagations;
            o.status = status;
            o.transitRouterRouteTableId = transitRouterRouteTableId;
            return o;
        }
    }
}
