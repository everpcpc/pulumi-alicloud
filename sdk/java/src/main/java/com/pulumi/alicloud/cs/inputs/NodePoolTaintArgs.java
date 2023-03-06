// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NodePoolTaintArgs extends com.pulumi.resources.ResourceArgs {

    public static final NodePoolTaintArgs Empty = new NodePoolTaintArgs();

    @Import(name="effect")
    private @Nullable Output<String> effect;

    public Optional<Output<String>> effect() {
        return Optional.ofNullable(this.effect);
    }

    /**
     * The label key.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return The label key.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    /**
     * The label value.
     * 
     */
    @Import(name="value")
    private @Nullable Output<String> value;

    /**
     * @return The label value.
     * 
     */
    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    private NodePoolTaintArgs() {}

    private NodePoolTaintArgs(NodePoolTaintArgs $) {
        this.effect = $.effect;
        this.key = $.key;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NodePoolTaintArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NodePoolTaintArgs $;

        public Builder() {
            $ = new NodePoolTaintArgs();
        }

        public Builder(NodePoolTaintArgs defaults) {
            $ = new NodePoolTaintArgs(Objects.requireNonNull(defaults));
        }

        public Builder effect(@Nullable Output<String> effect) {
            $.effect = effect;
            return this;
        }

        public Builder effect(String effect) {
            return effect(Output.of(effect));
        }

        /**
         * @param key The label key.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The label key.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param value The label value.
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value The label value.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public NodePoolTaintArgs build() {
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}