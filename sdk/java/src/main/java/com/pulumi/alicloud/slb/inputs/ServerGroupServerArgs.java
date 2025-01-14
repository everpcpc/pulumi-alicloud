// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServerGroupServerArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServerGroupServerArgs Empty = new ServerGroupServerArgs();

    @Import(name="port", required=true)
    private Output<Integer> port;

    public Output<Integer> port() {
        return this.port;
    }

    @Import(name="serverIds", required=true)
    private Output<List<String>> serverIds;

    public Output<List<String>> serverIds() {
        return this.serverIds;
    }

    @Import(name="type")
    private @Nullable Output<String> type;

    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    @Import(name="weight")
    private @Nullable Output<Integer> weight;

    public Optional<Output<Integer>> weight() {
        return Optional.ofNullable(this.weight);
    }

    private ServerGroupServerArgs() {}

    private ServerGroupServerArgs(ServerGroupServerArgs $) {
        this.port = $.port;
        this.serverIds = $.serverIds;
        this.type = $.type;
        this.weight = $.weight;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerGroupServerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerGroupServerArgs $;

        public Builder() {
            $ = new ServerGroupServerArgs();
        }

        public Builder(ServerGroupServerArgs defaults) {
            $ = new ServerGroupServerArgs(Objects.requireNonNull(defaults));
        }

        public Builder port(Output<Integer> port) {
            $.port = port;
            return this;
        }

        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        public Builder serverIds(Output<List<String>> serverIds) {
            $.serverIds = serverIds;
            return this;
        }

        public Builder serverIds(List<String> serverIds) {
            return serverIds(Output.of(serverIds));
        }

        public Builder serverIds(String... serverIds) {
            return serverIds(List.of(serverIds));
        }

        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public Builder weight(@Nullable Output<Integer> weight) {
            $.weight = weight;
            return this;
        }

        public Builder weight(Integer weight) {
            return weight(Output.of(weight));
        }

        public ServerGroupServerArgs build() {
            $.port = Objects.requireNonNull($.port, "expected parameter 'port' to be non-null");
            $.serverIds = Objects.requireNonNull($.serverIds, "expected parameter 'serverIds' to be non-null");
            return $;
        }
    }

}
