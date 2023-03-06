// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dcdn.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KvNamespaceState extends com.pulumi.resources.ResourceArgs {

    public static final KvNamespaceState Empty = new KvNamespaceState();

    /**
     * Namespace description information
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Namespace description information
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The status of the resource
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the resource
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private KvNamespaceState() {}

    private KvNamespaceState(KvNamespaceState $) {
        this.description = $.description;
        this.namespace = $.namespace;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KvNamespaceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KvNamespaceState $;

        public Builder() {
            $ = new KvNamespaceState();
        }

        public Builder(KvNamespaceState defaults) {
            $ = new KvNamespaceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Namespace description information
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Namespace description information
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param namespace Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace Namespace name. The name can contain letters, digits, hyphens (-), and underscores (_).
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param status The status of the resource
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the resource
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public KvNamespaceState build() {
            return $;
        }
    }

}