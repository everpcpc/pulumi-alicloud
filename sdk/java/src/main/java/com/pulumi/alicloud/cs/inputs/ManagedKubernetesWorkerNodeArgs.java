// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ManagedKubernetesWorkerNodeArgs extends com.pulumi.resources.ResourceArgs {

    public static final ManagedKubernetesWorkerNodeArgs Empty = new ManagedKubernetesWorkerNodeArgs();

    /**
     * (Deprecated from version 1.177.0) ID of the node.
     * 
     */
    @Import(name="id")
    private @Nullable Output<String> id;

    /**
     * @return (Deprecated from version 1.177.0) ID of the node.
     * 
     */
    public Optional<Output<String>> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * The kubernetes cluster&#39;s name. It is unique in one Alicloud account.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The kubernetes cluster&#39;s name. It is unique in one Alicloud account.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * (Deprecated from version 1.177.0) The private IP address of node.
     * 
     */
    @Import(name="privateIp")
    private @Nullable Output<String> privateIp;

    /**
     * @return (Deprecated from version 1.177.0) The private IP address of node.
     * 
     */
    public Optional<Output<String>> privateIp() {
        return Optional.ofNullable(this.privateIp);
    }

    private ManagedKubernetesWorkerNodeArgs() {}

    private ManagedKubernetesWorkerNodeArgs(ManagedKubernetesWorkerNodeArgs $) {
        this.id = $.id;
        this.name = $.name;
        this.privateIp = $.privateIp;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ManagedKubernetesWorkerNodeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ManagedKubernetesWorkerNodeArgs $;

        public Builder() {
            $ = new ManagedKubernetesWorkerNodeArgs();
        }

        public Builder(ManagedKubernetesWorkerNodeArgs defaults) {
            $ = new ManagedKubernetesWorkerNodeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id (Deprecated from version 1.177.0) ID of the node.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id (Deprecated from version 1.177.0) ID of the node.
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        /**
         * @param name The kubernetes cluster&#39;s name. It is unique in one Alicloud account.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The kubernetes cluster&#39;s name. It is unique in one Alicloud account.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param privateIp (Deprecated from version 1.177.0) The private IP address of node.
         * 
         * @return builder
         * 
         */
        public Builder privateIp(@Nullable Output<String> privateIp) {
            $.privateIp = privateIp;
            return this;
        }

        /**
         * @param privateIp (Deprecated from version 1.177.0) The private IP address of node.
         * 
         * @return builder
         * 
         */
        public Builder privateIp(String privateIp) {
            return privateIp(Output.of(privateIp));
        }

        public ManagedKubernetesWorkerNodeArgs build() {
            return $;
        }
    }

}
