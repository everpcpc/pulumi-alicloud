// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mse.outputs;

import com.pulumi.alicloud.mse.outputs.GetGatewaysGateway;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetGatewaysResult {
    private @Nullable Boolean enableDetails;
    private @Nullable String gatewayName;
    private List<GetGatewaysGateway> gateways;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String nameRegex;
    private List<String> names;
    private @Nullable String outputFile;
    private @Nullable String status;
    private @Nullable String vpcId;

    private GetGatewaysResult() {}
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }
    public Optional<String> gatewayName() {
        return Optional.ofNullable(this.gatewayName);
    }
    public List<GetGatewaysGateway> gateways() {
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

    public static Builder builder(GetGatewaysResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enableDetails;
        private @Nullable String gatewayName;
        private List<GetGatewaysGateway> gateways;
        private String id;
        private List<String> ids;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable String status;
        private @Nullable String vpcId;
        public Builder() {}
        public Builder(GetGatewaysResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enableDetails = defaults.enableDetails;
    	      this.gatewayName = defaults.gatewayName;
    	      this.gateways = defaults.gateways;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.status = defaults.status;
    	      this.vpcId = defaults.vpcId;
        }

        @CustomType.Setter
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            this.enableDetails = enableDetails;
            return this;
        }
        @CustomType.Setter
        public Builder gatewayName(@Nullable String gatewayName) {
            this.gatewayName = gatewayName;
            return this;
        }
        @CustomType.Setter
        public Builder gateways(List<GetGatewaysGateway> gateways) {
            this.gateways = Objects.requireNonNull(gateways);
            return this;
        }
        public Builder gateways(GetGatewaysGateway... gateways) {
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
        public GetGatewaysResult build() {
            final var o = new GetGatewaysResult();
            o.enableDetails = enableDetails;
            o.gatewayName = gatewayName;
            o.gateways = gateways;
            o.id = id;
            o.ids = ids;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.status = status;
            o.vpcId = vpcId;
            return o;
        }
    }
}