// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.pvtz.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class RuleForwardIp {
    /**
     * @return The ip of the forwarding destination.
     * 
     */
    private String ip;
    /**
     * @return The port of the forwarding destination.
     * 
     */
    private Integer port;

    private RuleForwardIp() {}
    /**
     * @return The ip of the forwarding destination.
     * 
     */
    public String ip() {
        return this.ip;
    }
    /**
     * @return The port of the forwarding destination.
     * 
     */
    public Integer port() {
        return this.port;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RuleForwardIp defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String ip;
        private Integer port;
        public Builder() {}
        public Builder(RuleForwardIp defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ip = defaults.ip;
    	      this.port = defaults.port;
        }

        @CustomType.Setter
        public Builder ip(String ip) {
            this.ip = Objects.requireNonNull(ip);
            return this;
        }
        @CustomType.Setter
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        public RuleForwardIp build() {
            final var o = new RuleForwardIp();
            o.ip = ip;
            o.port = port;
            return o;
        }
    }
}
