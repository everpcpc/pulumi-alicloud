// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RegistryEnterpriseNamespaceArgs extends com.pulumi.resources.ResourceArgs {

    public static final RegistryEnterpriseNamespaceArgs Empty = new RegistryEnterpriseNamespaceArgs();

    /**
     * Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
     * 
     */
    @Import(name="autoCreate", required=true)
    private Output<Boolean> autoCreate;

    /**
     * @return Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
     * 
     */
    public Output<Boolean> autoCreate() {
        return this.autoCreate;
    }

    /**
     * `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
     * 
     */
    @Import(name="defaultVisibility", required=true)
    private Output<String> defaultVisibility;

    /**
     * @return `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
     * 
     */
    public Output<String> defaultVisibility() {
        return this.defaultVisibility;
    }

    /**
     * ID of Container Registry Enterprise Edition instance.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return ID of Container Registry Enterprise Edition instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     * Name of Container Registry Enterprise Edition namespace. It can contain 2 to 30 characters.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of Container Registry Enterprise Edition namespace. It can contain 2 to 30 characters.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private RegistryEnterpriseNamespaceArgs() {}

    private RegistryEnterpriseNamespaceArgs(RegistryEnterpriseNamespaceArgs $) {
        this.autoCreate = $.autoCreate;
        this.defaultVisibility = $.defaultVisibility;
        this.instanceId = $.instanceId;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RegistryEnterpriseNamespaceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RegistryEnterpriseNamespaceArgs $;

        public Builder() {
            $ = new RegistryEnterpriseNamespaceArgs();
        }

        public Builder(RegistryEnterpriseNamespaceArgs defaults) {
            $ = new RegistryEnterpriseNamespaceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param autoCreate Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
         * 
         * @return builder
         * 
         */
        public Builder autoCreate(Output<Boolean> autoCreate) {
            $.autoCreate = autoCreate;
            return this;
        }

        /**
         * @param autoCreate Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.
         * 
         * @return builder
         * 
         */
        public Builder autoCreate(Boolean autoCreate) {
            return autoCreate(Output.of(autoCreate));
        }

        /**
         * @param defaultVisibility `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
         * 
         * @return builder
         * 
         */
        public Builder defaultVisibility(Output<String> defaultVisibility) {
            $.defaultVisibility = defaultVisibility;
            return this;
        }

        /**
         * @param defaultVisibility `PUBLIC` or `PRIVATE`, default repository visibility in this namespace.
         * 
         * @return builder
         * 
         */
        public Builder defaultVisibility(String defaultVisibility) {
            return defaultVisibility(Output.of(defaultVisibility));
        }

        /**
         * @param instanceId ID of Container Registry Enterprise Edition instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId ID of Container Registry Enterprise Edition instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param name Name of Container Registry Enterprise Edition namespace. It can contain 2 to 30 characters.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of Container Registry Enterprise Edition namespace. It can contain 2 to 30 characters.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public RegistryEnterpriseNamespaceArgs build() {
            $.autoCreate = Objects.requireNonNull($.autoCreate, "expected parameter 'autoCreate' to be non-null");
            $.defaultVisibility = Objects.requireNonNull($.defaultVisibility, "expected parameter 'defaultVisibility' to be non-null");
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
