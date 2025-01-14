// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eci.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ContainerGroupContainerPortArgs extends com.pulumi.resources.ResourceArgs {

    public static final ContainerGroupContainerPortArgs Empty = new ContainerGroupContainerPortArgs();

    /**
     * The port number. Valid values: 1 to 65535.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return The port number. Valid values: 1 to 65535.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * Valid values: TCP and UDP.
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return Valid values: TCP and UDP.
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    private ContainerGroupContainerPortArgs() {}

    private ContainerGroupContainerPortArgs(ContainerGroupContainerPortArgs $) {
        this.port = $.port;
        this.protocol = $.protocol;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerGroupContainerPortArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerGroupContainerPortArgs $;

        public Builder() {
            $ = new ContainerGroupContainerPortArgs();
        }

        public Builder(ContainerGroupContainerPortArgs defaults) {
            $ = new ContainerGroupContainerPortArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param port The port number. Valid values: 1 to 65535.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port number. Valid values: 1 to 65535.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param protocol Valid values: TCP and UDP.
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol Valid values: TCP and UDP.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        public ContainerGroupContainerPortArgs build() {
            return $;
        }
    }

}
