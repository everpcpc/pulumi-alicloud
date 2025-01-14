// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb.outputs;

import com.pulumi.alicloud.slb.outputs.GetBackendServersBackendServer;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBackendServersResult {
    private List<GetBackendServersBackendServer> backendServers;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private String loadBalancerId;
    private @Nullable String outputFile;

    private GetBackendServersResult() {}
    public List<GetBackendServersBackendServer> backendServers() {
        return this.backendServers;
    }
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
    public String loadBalancerId() {
        return this.loadBalancerId;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBackendServersResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetBackendServersBackendServer> backendServers;
        private String id;
        private List<String> ids;
        private String loadBalancerId;
        private @Nullable String outputFile;
        public Builder() {}
        public Builder(GetBackendServersResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backendServers = defaults.backendServers;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.loadBalancerId = defaults.loadBalancerId;
    	      this.outputFile = defaults.outputFile;
        }

        @CustomType.Setter
        public Builder backendServers(List<GetBackendServersBackendServer> backendServers) {
            this.backendServers = Objects.requireNonNull(backendServers);
            return this;
        }
        public Builder backendServers(GetBackendServersBackendServer... backendServers) {
            return backendServers(List.of(backendServers));
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
        public Builder loadBalancerId(String loadBalancerId) {
            this.loadBalancerId = Objects.requireNonNull(loadBalancerId);
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        public GetBackendServersResult build() {
            final var o = new GetBackendServersResult();
            o.backendServers = backendServers;
            o.id = id;
            o.ids = ids;
            o.loadBalancerId = loadBalancerId;
            o.outputFile = outputFile;
            return o;
        }
    }
}
