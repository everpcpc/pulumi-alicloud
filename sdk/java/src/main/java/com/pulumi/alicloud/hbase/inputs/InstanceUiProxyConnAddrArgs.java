// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbase.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceUiProxyConnAddrArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceUiProxyConnAddrArgs Empty = new InstanceUiProxyConnAddrArgs();

    @Import(name="connAddr")
    private @Nullable Output<String> connAddr;

    public Optional<Output<String>> connAddr() {
        return Optional.ofNullable(this.connAddr);
    }

    @Import(name="connAddrPort")
    private @Nullable Output<String> connAddrPort;

    public Optional<Output<String>> connAddrPort() {
        return Optional.ofNullable(this.connAddrPort);
    }

    @Import(name="netType")
    private @Nullable Output<String> netType;

    public Optional<Output<String>> netType() {
        return Optional.ofNullable(this.netType);
    }

    private InstanceUiProxyConnAddrArgs() {}

    private InstanceUiProxyConnAddrArgs(InstanceUiProxyConnAddrArgs $) {
        this.connAddr = $.connAddr;
        this.connAddrPort = $.connAddrPort;
        this.netType = $.netType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceUiProxyConnAddrArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceUiProxyConnAddrArgs $;

        public Builder() {
            $ = new InstanceUiProxyConnAddrArgs();
        }

        public Builder(InstanceUiProxyConnAddrArgs defaults) {
            $ = new InstanceUiProxyConnAddrArgs(Objects.requireNonNull(defaults));
        }

        public Builder connAddr(@Nullable Output<String> connAddr) {
            $.connAddr = connAddr;
            return this;
        }

        public Builder connAddr(String connAddr) {
            return connAddr(Output.of(connAddr));
        }

        public Builder connAddrPort(@Nullable Output<String> connAddrPort) {
            $.connAddrPort = connAddrPort;
            return this;
        }

        public Builder connAddrPort(String connAddrPort) {
            return connAddrPort(Output.of(connAddrPort));
        }

        public Builder netType(@Nullable Output<String> netType) {
            $.netType = netType;
            return this;
        }

        public Builder netType(String netType) {
            return netType(Output.of(netType));
        }

        public InstanceUiProxyConnAddrArgs build() {
            return $;
        }
    }

}