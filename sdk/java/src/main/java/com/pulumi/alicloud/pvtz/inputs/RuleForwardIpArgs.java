// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.pvtz.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class RuleForwardIpArgs extends com.pulumi.resources.ResourceArgs {

    public static final RuleForwardIpArgs Empty = new RuleForwardIpArgs();

    /**
     * The ip of the forwarding destination.
     * 
     */
    @Import(name="ip", required=true)
    private Output<String> ip;

    /**
     * @return The ip of the forwarding destination.
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }

    /**
     * The port of the forwarding destination.
     * 
     */
    @Import(name="port", required=true)
    private Output<Integer> port;

    /**
     * @return The port of the forwarding destination.
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }

    private RuleForwardIpArgs() {}

    private RuleForwardIpArgs(RuleForwardIpArgs $) {
        this.ip = $.ip;
        this.port = $.port;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RuleForwardIpArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RuleForwardIpArgs $;

        public Builder() {
            $ = new RuleForwardIpArgs();
        }

        public Builder(RuleForwardIpArgs defaults) {
            $ = new RuleForwardIpArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ip The ip of the forwarding destination.
         * 
         * @return builder
         * 
         */
        public Builder ip(Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip The ip of the forwarding destination.
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        /**
         * @param port The port of the forwarding destination.
         * 
         * @return builder
         * 
         */
        public Builder port(Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port of the forwarding destination.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        public RuleForwardIpArgs build() {
            $.ip = Objects.requireNonNull($.ip, "expected parameter 'ip' to be non-null");
            $.port = Objects.requireNonNull($.port, "expected parameter 'port' to be non-null");
            return $;
        }
    }

}
