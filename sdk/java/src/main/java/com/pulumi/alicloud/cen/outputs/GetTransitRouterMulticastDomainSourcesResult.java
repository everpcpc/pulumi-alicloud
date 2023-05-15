// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.outputs;

import com.pulumi.alicloud.cen.outputs.GetTransitRouterMulticastDomainSourcesSource;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetTransitRouterMulticastDomainSourcesResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String outputFile;
    /**
     * @return A list of Transit Router Multicast Domain Source Entries. Each element contains the following attributes:
     * 
     */
    private List<GetTransitRouterMulticastDomainSourcesSource> sources;
    /**
     * @return The ID of the multicast domain to which the multicast source belongs.
     * 
     */
    private String transitRouterMulticastDomainId;

    private GetTransitRouterMulticastDomainSourcesResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return A list of Transit Router Multicast Domain Source Entries. Each element contains the following attributes:
     * 
     */
    public List<GetTransitRouterMulticastDomainSourcesSource> sources() {
        return this.sources;
    }
    /**
     * @return The ID of the multicast domain to which the multicast source belongs.
     * 
     */
    public String transitRouterMulticastDomainId() {
        return this.transitRouterMulticastDomainId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTransitRouterMulticastDomainSourcesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> ids;
        private @Nullable String outputFile;
        private List<GetTransitRouterMulticastDomainSourcesSource> sources;
        private String transitRouterMulticastDomainId;
        public Builder() {}
        public Builder(GetTransitRouterMulticastDomainSourcesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.outputFile = defaults.outputFile;
    	      this.sources = defaults.sources;
    	      this.transitRouterMulticastDomainId = defaults.transitRouterMulticastDomainId;
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
        public Builder sources(List<GetTransitRouterMulticastDomainSourcesSource> sources) {
            this.sources = Objects.requireNonNull(sources);
            return this;
        }
        public Builder sources(GetTransitRouterMulticastDomainSourcesSource... sources) {
            return sources(List.of(sources));
        }
        @CustomType.Setter
        public Builder transitRouterMulticastDomainId(String transitRouterMulticastDomainId) {
            this.transitRouterMulticastDomainId = Objects.requireNonNull(transitRouterMulticastDomainId);
            return this;
        }
        public GetTransitRouterMulticastDomainSourcesResult build() {
            final var o = new GetTransitRouterMulticastDomainSourcesResult();
            o.id = id;
            o.ids = ids;
            o.outputFile = outputFile;
            o.sources = sources;
            o.transitRouterMulticastDomainId = transitRouterMulticastDomainId;
            return o;
        }
    }
}
