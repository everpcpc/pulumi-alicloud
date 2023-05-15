// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetLoadBalancersBalancerBackendServer {
    private String description;
    private String serverId;
    private String type;
    private Integer weight;

    private GetLoadBalancersBalancerBackendServer() {}
    public String description() {
        return this.description;
    }
    public String serverId() {
        return this.serverId;
    }
    public String type() {
        return this.type;
    }
    public Integer weight() {
        return this.weight;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLoadBalancersBalancerBackendServer defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String serverId;
        private String type;
        private Integer weight;
        public Builder() {}
        public Builder(GetLoadBalancersBalancerBackendServer defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.serverId = defaults.serverId;
    	      this.type = defaults.type;
    	      this.weight = defaults.weight;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder serverId(String serverId) {
            this.serverId = Objects.requireNonNull(serverId);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        @CustomType.Setter
        public Builder weight(Integer weight) {
            this.weight = Objects.requireNonNull(weight);
            return this;
        }
        public GetLoadBalancersBalancerBackendServer build() {
            final var o = new GetLoadBalancersBalancerBackendServer();
            o.description = description;
            o.serverId = serverId;
            o.type = type;
            o.weight = weight;
            return o;
        }
    }
}
