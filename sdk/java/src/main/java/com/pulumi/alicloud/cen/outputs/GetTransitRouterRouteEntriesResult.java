// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.outputs;

import com.pulumi.alicloud.cen.outputs.GetTransitRouterRouteEntriesEntry;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetTransitRouterRouteEntriesResult {
    /**
     * @return A list of CEN Route Entries. Each element contains the following attributes:
     * 
     */
    private List<GetTransitRouterRouteEntriesEntry> entries;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of CEN Transit Router Route Entry IDs.
     * 
     */
    private List<String> ids;
    private @Nullable String nameRegex;
    /**
     * @return A list of CEN Transit Router Route Entry Names.
     * 
     */
    private List<String> names;
    private @Nullable String outputFile;
    private @Nullable String status;
    private @Nullable List<String> transitRouterRouteEntryIds;
    private @Nullable List<String> transitRouterRouteEntryNames;
    /**
     * @return The status of the route entry in CEN.
     * 
     */
    private @Nullable String transitRouterRouteEntryStatus;
    private String transitRouterRouteTableId;

    private GetTransitRouterRouteEntriesResult() {}
    /**
     * @return A list of CEN Route Entries. Each element contains the following attributes:
     * 
     */
    public List<GetTransitRouterRouteEntriesEntry> entries() {
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
     * @return A list of CEN Transit Router Route Entry IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of CEN Transit Router Route Entry Names.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    public List<String> transitRouterRouteEntryIds() {
        return this.transitRouterRouteEntryIds == null ? List.of() : this.transitRouterRouteEntryIds;
    }
    public List<String> transitRouterRouteEntryNames() {
        return this.transitRouterRouteEntryNames == null ? List.of() : this.transitRouterRouteEntryNames;
    }
    /**
     * @return The status of the route entry in CEN.
     * 
     */
    public Optional<String> transitRouterRouteEntryStatus() {
        return Optional.ofNullable(this.transitRouterRouteEntryStatus);
    }
    public String transitRouterRouteTableId() {
        return this.transitRouterRouteTableId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTransitRouterRouteEntriesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetTransitRouterRouteEntriesEntry> entries;
        private String id;
        private List<String> ids;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable String status;
        private @Nullable List<String> transitRouterRouteEntryIds;
        private @Nullable List<String> transitRouterRouteEntryNames;
        private @Nullable String transitRouterRouteEntryStatus;
        private String transitRouterRouteTableId;
        public Builder() {}
        public Builder(GetTransitRouterRouteEntriesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.entries = defaults.entries;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.status = defaults.status;
    	      this.transitRouterRouteEntryIds = defaults.transitRouterRouteEntryIds;
    	      this.transitRouterRouteEntryNames = defaults.transitRouterRouteEntryNames;
    	      this.transitRouterRouteEntryStatus = defaults.transitRouterRouteEntryStatus;
    	      this.transitRouterRouteTableId = defaults.transitRouterRouteTableId;
        }

        @CustomType.Setter
        public Builder entries(List<GetTransitRouterRouteEntriesEntry> entries) {
            this.entries = Objects.requireNonNull(entries);
            return this;
        }
        public Builder entries(GetTransitRouterRouteEntriesEntry... entries) {
            return entries(List.of(entries));
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
        public Builder nameRegex(@Nullable String nameRegex) {
            this.nameRegex = nameRegex;
            return this;
        }
        @CustomType.Setter
        public Builder names(List<String> names) {
            this.names = Objects.requireNonNull(names);
            return this;
        }
        public Builder names(String... names) {
            return names(List.of(names));
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder transitRouterRouteEntryIds(@Nullable List<String> transitRouterRouteEntryIds) {
            this.transitRouterRouteEntryIds = transitRouterRouteEntryIds;
            return this;
        }
        public Builder transitRouterRouteEntryIds(String... transitRouterRouteEntryIds) {
            return transitRouterRouteEntryIds(List.of(transitRouterRouteEntryIds));
        }
        @CustomType.Setter
        public Builder transitRouterRouteEntryNames(@Nullable List<String> transitRouterRouteEntryNames) {
            this.transitRouterRouteEntryNames = transitRouterRouteEntryNames;
            return this;
        }
        public Builder transitRouterRouteEntryNames(String... transitRouterRouteEntryNames) {
            return transitRouterRouteEntryNames(List.of(transitRouterRouteEntryNames));
        }
        @CustomType.Setter
        public Builder transitRouterRouteEntryStatus(@Nullable String transitRouterRouteEntryStatus) {
            this.transitRouterRouteEntryStatus = transitRouterRouteEntryStatus;
            return this;
        }
        @CustomType.Setter
        public Builder transitRouterRouteTableId(String transitRouterRouteTableId) {
            this.transitRouterRouteTableId = Objects.requireNonNull(transitRouterRouteTableId);
            return this;
        }
        public GetTransitRouterRouteEntriesResult build() {
            final var o = new GetTransitRouterRouteEntriesResult();
            o.entries = entries;
            o.id = id;
            o.ids = ids;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.status = status;
            o.transitRouterRouteEntryIds = transitRouterRouteEntryIds;
            o.transitRouterRouteEntryNames = transitRouterRouteEntryNames;
            o.transitRouterRouteEntryStatus = transitRouterRouteEntryStatus;
            o.transitRouterRouteTableId = transitRouterRouteTableId;
            return o;
        }
    }
}
