// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.privatelink.outputs;

import com.pulumi.alicloud.privatelink.outputs.GetVpcEndpointsEndpoint;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetVpcEndpointsResult {
    private @Nullable String connectionStatus;
    private @Nullable Boolean enableDetails;
    private List<GetVpcEndpointsEndpoint> endpoints;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String nameRegex;
    private List<String> names;
    private @Nullable String outputFile;
    private @Nullable String serviceName;
    private @Nullable String status;
    private @Nullable String vpcEndpointName;
    private @Nullable String vpcId;

    private GetVpcEndpointsResult() {}
    public Optional<String> connectionStatus() {
        return Optional.ofNullable(this.connectionStatus);
    }
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }
    public List<GetVpcEndpointsEndpoint> endpoints() {
        return this.endpoints;
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
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Optional<String> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    public Optional<String> vpcEndpointName() {
        return Optional.ofNullable(this.vpcEndpointName);
    }
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVpcEndpointsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String connectionStatus;
        private @Nullable Boolean enableDetails;
        private List<GetVpcEndpointsEndpoint> endpoints;
        private String id;
        private List<String> ids;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable String serviceName;
        private @Nullable String status;
        private @Nullable String vpcEndpointName;
        private @Nullable String vpcId;
        public Builder() {}
        public Builder(GetVpcEndpointsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.connectionStatus = defaults.connectionStatus;
    	      this.enableDetails = defaults.enableDetails;
    	      this.endpoints = defaults.endpoints;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.serviceName = defaults.serviceName;
    	      this.status = defaults.status;
    	      this.vpcEndpointName = defaults.vpcEndpointName;
    	      this.vpcId = defaults.vpcId;
        }

        @CustomType.Setter
        public Builder connectionStatus(@Nullable String connectionStatus) {
            this.connectionStatus = connectionStatus;
            return this;
        }
        @CustomType.Setter
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            this.enableDetails = enableDetails;
            return this;
        }
        @CustomType.Setter
        public Builder endpoints(List<GetVpcEndpointsEndpoint> endpoints) {
            this.endpoints = Objects.requireNonNull(endpoints);
            return this;
        }
        public Builder endpoints(GetVpcEndpointsEndpoint... endpoints) {
            return endpoints(List.of(endpoints));
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
        public Builder serviceName(@Nullable String serviceName) {
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder vpcEndpointName(@Nullable String vpcEndpointName) {
            this.vpcEndpointName = vpcEndpointName;
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(@Nullable String vpcId) {
            this.vpcId = vpcId;
            return this;
        }
        public GetVpcEndpointsResult build() {
            final var o = new GetVpcEndpointsResult();
            o.connectionStatus = connectionStatus;
            o.enableDetails = enableDetails;
            o.endpoints = endpoints;
            o.id = id;
            o.ids = ids;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.serviceName = serviceName;
            o.status = status;
            o.vpcEndpointName = vpcEndpointName;
            o.vpcId = vpcId;
            return o;
        }
    }
}
