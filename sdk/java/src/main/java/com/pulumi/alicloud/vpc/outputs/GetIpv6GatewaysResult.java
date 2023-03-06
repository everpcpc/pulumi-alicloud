// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.outputs;

import com.pulumi.alicloud.vpc.outputs.GetIpv6GatewaysGateway;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetIpv6GatewaysResult {
    private List<GetIpv6GatewaysGateway> gateways;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String ipv6GatewayName;
    private @Nullable String nameRegex;
    private List<String> names;
    private @Nullable String outputFile;
    private @Nullable String status;
    private @Nullable String vpcId;

    private GetIpv6GatewaysResult() {}
    public List<GetIpv6GatewaysGateway> gateways() {
        return this.gateways;
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
    public Optional<String> ipv6GatewayName() {
        return Optional.ofNullable(this.ipv6GatewayName);
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
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    public Optional<String> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIpv6GatewaysResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetIpv6GatewaysGateway> gateways;
        private String id;
        private List<String> ids;
        private @Nullable String ipv6GatewayName;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable String status;
        private @Nullable String vpcId;
        public Builder() {}
        public Builder(GetIpv6GatewaysResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.gateways = defaults.gateways;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.ipv6GatewayName = defaults.ipv6GatewayName;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.status = defaults.status;
    	      this.vpcId = defaults.vpcId;
        }

        @CustomType.Setter
        public Builder gateways(List<GetIpv6GatewaysGateway> gateways) {
            this.gateways = Objects.requireNonNull(gateways);
            return this;
        }
        public Builder gateways(GetIpv6GatewaysGateway... gateways) {
            return gateways(List.of(gateways));
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
        public Builder ipv6GatewayName(@Nullable String ipv6GatewayName) {
            this.ipv6GatewayName = ipv6GatewayName;
            return this;
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
        public Builder vpcId(@Nullable String vpcId) {
            this.vpcId = vpcId;
            return this;
        }
        public GetIpv6GatewaysResult build() {
            final var o = new GetIpv6GatewaysResult();
            o.gateways = gateways;
            o.id = id;
            o.ids = ids;
            o.ipv6GatewayName = ipv6GatewayName;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.status = status;
            o.vpcId = vpcId;
            return o;
        }
    }
}