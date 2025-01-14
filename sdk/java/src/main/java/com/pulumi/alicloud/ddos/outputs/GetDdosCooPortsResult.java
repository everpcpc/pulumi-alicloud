// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ddos.outputs;

import com.pulumi.alicloud.ddos.outputs.GetDdosCooPortsPort;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetDdosCooPortsResult {
    private @Nullable String frontendPort;
    private @Nullable String frontendProtocol;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private String instanceId;
    private @Nullable String outputFile;
    private List<GetDdosCooPortsPort> ports;

    private GetDdosCooPortsResult() {}
    public Optional<String> frontendPort() {
        return Optional.ofNullable(this.frontendPort);
    }
    public Optional<String> frontendProtocol() {
        return Optional.ofNullable(this.frontendProtocol);
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
    public String instanceId() {
        return this.instanceId;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public List<GetDdosCooPortsPort> ports() {
        return this.ports;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDdosCooPortsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String frontendPort;
        private @Nullable String frontendProtocol;
        private String id;
        private List<String> ids;
        private String instanceId;
        private @Nullable String outputFile;
        private List<GetDdosCooPortsPort> ports;
        public Builder() {}
        public Builder(GetDdosCooPortsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.frontendPort = defaults.frontendPort;
    	      this.frontendProtocol = defaults.frontendProtocol;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.instanceId = defaults.instanceId;
    	      this.outputFile = defaults.outputFile;
    	      this.ports = defaults.ports;
        }

        @CustomType.Setter
        public Builder frontendPort(@Nullable String frontendPort) {
            this.frontendPort = frontendPort;
            return this;
        }
        @CustomType.Setter
        public Builder frontendProtocol(@Nullable String frontendProtocol) {
            this.frontendProtocol = frontendProtocol;
            return this;
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
        public Builder instanceId(String instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder ports(List<GetDdosCooPortsPort> ports) {
            this.ports = Objects.requireNonNull(ports);
            return this;
        }
        public Builder ports(GetDdosCooPortsPort... ports) {
            return ports(List.of(ports));
        }
        public GetDdosCooPortsResult build() {
            final var o = new GetDdosCooPortsResult();
            o.frontendPort = frontendPort;
            o.frontendProtocol = frontendProtocol;
            o.id = id;
            o.ids = ids;
            o.instanceId = instanceId;
            o.outputFile = outputFile;
            o.ports = ports;
            return o;
        }
    }
}
