// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EndpointArgs extends com.pulumi.resources.ResourceArgs {

    public static final EndpointArgs Empty = new EndpointArgs();

    @Import(name="autoAddNewNodes")
    private @Nullable Output<String> autoAddNewNodes;

    public Optional<Output<String>> autoAddNewNodes() {
        return Optional.ofNullable(this.autoAddNewNodes);
    }

    @Import(name="dbClusterId", required=true)
    private Output<String> dbClusterId;

    public Output<String> dbClusterId() {
        return this.dbClusterId;
    }

    @Import(name="endpointConfig")
    private @Nullable Output<Map<String,Object>> endpointConfig;

    public Optional<Output<Map<String,Object>>> endpointConfig() {
        return Optional.ofNullable(this.endpointConfig);
    }

    /**
     * Type of endpoint.
     * 
     */
    @Import(name="endpointType")
    private @Nullable Output<String> endpointType;

    /**
     * @return Type of endpoint.
     * 
     */
    public Optional<Output<String>> endpointType() {
        return Optional.ofNullable(this.endpointType);
    }

    @Import(name="netType")
    private @Nullable Output<String> netType;

    public Optional<Output<String>> netType() {
        return Optional.ofNullable(this.netType);
    }

    @Import(name="nodes")
    private @Nullable Output<List<String>> nodes;

    public Optional<Output<List<String>>> nodes() {
        return Optional.ofNullable(this.nodes);
    }

    @Import(name="readWriteMode")
    private @Nullable Output<String> readWriteMode;

    public Optional<Output<String>> readWriteMode() {
        return Optional.ofNullable(this.readWriteMode);
    }

    @Import(name="sslAutoRotate")
    private @Nullable Output<String> sslAutoRotate;

    public Optional<Output<String>> sslAutoRotate() {
        return Optional.ofNullable(this.sslAutoRotate);
    }

    @Import(name="sslEnabled")
    private @Nullable Output<String> sslEnabled;

    public Optional<Output<String>> sslEnabled() {
        return Optional.ofNullable(this.sslEnabled);
    }

    private EndpointArgs() {}

    private EndpointArgs(EndpointArgs $) {
        this.autoAddNewNodes = $.autoAddNewNodes;
        this.dbClusterId = $.dbClusterId;
        this.endpointConfig = $.endpointConfig;
        this.endpointType = $.endpointType;
        this.netType = $.netType;
        this.nodes = $.nodes;
        this.readWriteMode = $.readWriteMode;
        this.sslAutoRotate = $.sslAutoRotate;
        this.sslEnabled = $.sslEnabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EndpointArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EndpointArgs $;

        public Builder() {
            $ = new EndpointArgs();
        }

        public Builder(EndpointArgs defaults) {
            $ = new EndpointArgs(Objects.requireNonNull(defaults));
        }

        public Builder autoAddNewNodes(@Nullable Output<String> autoAddNewNodes) {
            $.autoAddNewNodes = autoAddNewNodes;
            return this;
        }

        public Builder autoAddNewNodes(String autoAddNewNodes) {
            return autoAddNewNodes(Output.of(autoAddNewNodes));
        }

        public Builder dbClusterId(Output<String> dbClusterId) {
            $.dbClusterId = dbClusterId;
            return this;
        }

        public Builder dbClusterId(String dbClusterId) {
            return dbClusterId(Output.of(dbClusterId));
        }

        public Builder endpointConfig(@Nullable Output<Map<String,Object>> endpointConfig) {
            $.endpointConfig = endpointConfig;
            return this;
        }

        public Builder endpointConfig(Map<String,Object> endpointConfig) {
            return endpointConfig(Output.of(endpointConfig));
        }

        /**
         * @param endpointType Type of endpoint.
         * 
         * @return builder
         * 
         */
        public Builder endpointType(@Nullable Output<String> endpointType) {
            $.endpointType = endpointType;
            return this;
        }

        /**
         * @param endpointType Type of endpoint.
         * 
         * @return builder
         * 
         */
        public Builder endpointType(String endpointType) {
            return endpointType(Output.of(endpointType));
        }

        public Builder netType(@Nullable Output<String> netType) {
            $.netType = netType;
            return this;
        }

        public Builder netType(String netType) {
            return netType(Output.of(netType));
        }

        public Builder nodes(@Nullable Output<List<String>> nodes) {
            $.nodes = nodes;
            return this;
        }

        public Builder nodes(List<String> nodes) {
            return nodes(Output.of(nodes));
        }

        public Builder nodes(String... nodes) {
            return nodes(List.of(nodes));
        }

        public Builder readWriteMode(@Nullable Output<String> readWriteMode) {
            $.readWriteMode = readWriteMode;
            return this;
        }

        public Builder readWriteMode(String readWriteMode) {
            return readWriteMode(Output.of(readWriteMode));
        }

        public Builder sslAutoRotate(@Nullable Output<String> sslAutoRotate) {
            $.sslAutoRotate = sslAutoRotate;
            return this;
        }

        public Builder sslAutoRotate(String sslAutoRotate) {
            return sslAutoRotate(Output.of(sslAutoRotate));
        }

        public Builder sslEnabled(@Nullable Output<String> sslEnabled) {
            $.sslEnabled = sslEnabled;
            return this;
        }

        public Builder sslEnabled(String sslEnabled) {
            return sslEnabled(Output.of(sslEnabled));
        }

        public EndpointArgs build() {
            $.dbClusterId = Objects.requireNonNull($.dbClusterId, "expected parameter 'dbClusterId' to be non-null");
            return $;
        }
    }

}