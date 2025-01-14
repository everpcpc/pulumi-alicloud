// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.outputs;

import com.pulumi.alicloud.vpc.outputs.GetRouteEntriesEntry;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetRouteEntriesResult {
    /**
     * @return The destination CIDR block of the route entry.
     * 
     */
    private @Nullable String cidrBlock;
    /**
     * @return A list of Route Entries. Each element contains the following attributes:
     * 
     */
    private List<GetRouteEntriesEntry> entries;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The instance ID of the next hop.
     * 
     */
    private @Nullable String instanceId;
    private @Nullable String outputFile;
    /**
     * @return The ID of the router table to which the route entry belongs.
     * 
     */
    private String routeTableId;
    /**
     * @return The type of the route entry.
     * 
     */
    private @Nullable String type;

    private GetRouteEntriesResult() {}
    /**
     * @return The destination CIDR block of the route entry.
     * 
     */
    public Optional<String> cidrBlock() {
        return Optional.ofNullable(this.cidrBlock);
    }
    /**
     * @return A list of Route Entries. Each element contains the following attributes:
     * 
     */
    public List<GetRouteEntriesEntry> entries() {
        return this.entries;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The instance ID of the next hop.
     * 
     */
    public Optional<String> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return The ID of the router table to which the route entry belongs.
     * 
     */
    public String routeTableId() {
        return this.routeTableId;
    }
    /**
     * @return The type of the route entry.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRouteEntriesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String cidrBlock;
        private List<GetRouteEntriesEntry> entries;
        private String id;
        private @Nullable String instanceId;
        private @Nullable String outputFile;
        private String routeTableId;
        private @Nullable String type;
        public Builder() {}
        public Builder(GetRouteEntriesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cidrBlock = defaults.cidrBlock;
    	      this.entries = defaults.entries;
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.outputFile = defaults.outputFile;
    	      this.routeTableId = defaults.routeTableId;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder cidrBlock(@Nullable String cidrBlock) {
            this.cidrBlock = cidrBlock;
            return this;
        }
        @CustomType.Setter
        public Builder entries(List<GetRouteEntriesEntry> entries) {
            this.entries = Objects.requireNonNull(entries);
            return this;
        }
        public Builder entries(GetRouteEntriesEntry... entries) {
            return entries(List.of(entries));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(@Nullable String instanceId) {
            this.instanceId = instanceId;
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder routeTableId(String routeTableId) {
            this.routeTableId = Objects.requireNonNull(routeTableId);
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        public GetRouteEntriesResult build() {
            final var o = new GetRouteEntriesResult();
            o.cidrBlock = cidrBlock;
            o.entries = entries;
            o.id = id;
            o.instanceId = instanceId;
            o.outputFile = outputFile;
            o.routeTableId = routeTableId;
            o.type = type;
            return o;
        }
    }
}
