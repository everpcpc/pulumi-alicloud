// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BackendServerBackendServerArgs extends com.pulumi.resources.ResourceArgs {

    public static final BackendServerBackendServerArgs Empty = new BackendServerBackendServerArgs();

    @Import(name="serverId", required=true)
    private Output<String> serverId;

    public Output<String> serverId() {
        return this.serverId;
    }

    @Import(name="serverIp")
    private @Nullable Output<String> serverIp;

    public Optional<Output<String>> serverIp() {
        return Optional.ofNullable(this.serverIp);
    }

    @Import(name="type")
    private @Nullable Output<String> type;

    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    @Import(name="weight", required=true)
    private Output<Integer> weight;

    public Output<Integer> weight() {
        return this.weight;
    }

    private BackendServerBackendServerArgs() {}

    private BackendServerBackendServerArgs(BackendServerBackendServerArgs $) {
        this.serverId = $.serverId;
        this.serverIp = $.serverIp;
        this.type = $.type;
        this.weight = $.weight;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BackendServerBackendServerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BackendServerBackendServerArgs $;

        public Builder() {
            $ = new BackendServerBackendServerArgs();
        }

        public Builder(BackendServerBackendServerArgs defaults) {
            $ = new BackendServerBackendServerArgs(Objects.requireNonNull(defaults));
        }

        public Builder serverId(Output<String> serverId) {
            $.serverId = serverId;
            return this;
        }

        public Builder serverId(String serverId) {
            return serverId(Output.of(serverId));
        }

        public Builder serverIp(@Nullable Output<String> serverIp) {
            $.serverIp = serverIp;
            return this;
        }

        public Builder serverIp(String serverIp) {
            return serverIp(Output.of(serverIp));
        }

        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public Builder weight(Output<Integer> weight) {
            $.weight = weight;
            return this;
        }

        public Builder weight(Integer weight) {
            return weight(Output.of(weight));
        }

        public BackendServerBackendServerArgs build() {
            $.serverId = Objects.requireNonNull($.serverId, "expected parameter 'serverId' to be non-null");
            $.weight = Objects.requireNonNull($.weight, "expected parameter 'weight' to be non-null");
            return $;
        }
    }

}
