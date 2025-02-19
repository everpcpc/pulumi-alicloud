// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetTlsCipherPoliciesPolicyRelateListener {
    /**
     * @return The ID of SLB instance.
     * 
     */
    private String loadBalancerId;
    /**
     * @return Listening port. Valid value: 1 to 65535.
     * 
     */
    private Integer port;
    /**
     * @return Snooping protocols. Valid values: `TCP`, `UDP`, `HTTP`, or `HTTPS`.
     * 
     */
    private String protocol;

    private GetTlsCipherPoliciesPolicyRelateListener() {}
    /**
     * @return The ID of SLB instance.
     * 
     */
    public String loadBalancerId() {
        return this.loadBalancerId;
    }
    /**
     * @return Listening port. Valid value: 1 to 65535.
     * 
     */
    public Integer port() {
        return this.port;
    }
    /**
     * @return Snooping protocols. Valid values: `TCP`, `UDP`, `HTTP`, or `HTTPS`.
     * 
     */
    public String protocol() {
        return this.protocol;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTlsCipherPoliciesPolicyRelateListener defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String loadBalancerId;
        private Integer port;
        private String protocol;
        public Builder() {}
        public Builder(GetTlsCipherPoliciesPolicyRelateListener defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.loadBalancerId = defaults.loadBalancerId;
    	      this.port = defaults.port;
    	      this.protocol = defaults.protocol;
        }

        @CustomType.Setter
        public Builder loadBalancerId(String loadBalancerId) {
            this.loadBalancerId = Objects.requireNonNull(loadBalancerId);
            return this;
        }
        @CustomType.Setter
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        @CustomType.Setter
        public Builder protocol(String protocol) {
            this.protocol = Objects.requireNonNull(protocol);
            return this;
        }
        public GetTlsCipherPoliciesPolicyRelateListener build() {
            final var o = new GetTlsCipherPoliciesPolicyRelateListener();
            o.loadBalancerId = loadBalancerId;
            o.port = port;
            o.protocol = protocol;
            return o;
        }
    }
}
