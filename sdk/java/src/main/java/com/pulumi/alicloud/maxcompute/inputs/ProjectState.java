// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.maxcompute.inputs;

import com.pulumi.alicloud.maxcompute.inputs.ProjectIpWhiteListArgs;
import com.pulumi.alicloud.maxcompute.inputs.ProjectPropertiesArgs;
import com.pulumi.alicloud.maxcompute.inputs.ProjectSecurityPropertiesArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectState Empty = new ProjectState();

    /**
     * Comments of project
     * 
     */
    @Import(name="comment")
    private @Nullable Output<String> comment;

    /**
     * @return Comments of project
     * 
     */
    public Optional<Output<String>> comment() {
        return Optional.ofNullable(this.comment);
    }

    /**
     * Default Computing Resource Group
     * 
     */
    @Import(name="defaultQuota")
    private @Nullable Output<String> defaultQuota;

    /**
     * @return Default Computing Resource Group
     * 
     */
    public Optional<Output<String>> defaultQuota() {
        return Optional.ofNullable(this.defaultQuota);
    }

    /**
     * IP whitelistSee the following `Block IpWhiteList`.
     * 
     */
    @Import(name="ipWhiteList")
    private @Nullable Output<ProjectIpWhiteListArgs> ipWhiteList;

    /**
     * @return IP whitelistSee the following `Block IpWhiteList`.
     * 
     */
    public Optional<Output<ProjectIpWhiteListArgs>> ipWhiteList() {
        return Optional.ofNullable(this.ipWhiteList);
    }

    /**
     * Project owner
     * 
     */
    @Import(name="owner")
    private @Nullable Output<String> owner;

    /**
     * @return Project owner
     * 
     */
    public Optional<Output<String>> owner() {
        return Optional.ofNullable(this.owner);
    }

    /**
     * Quota payment type, support `PayAsYouGo`, `Subscription`, `Dev`.
     * 
     */
    @Import(name="productType")
    private @Nullable Output<String> productType;

    /**
     * @return Quota payment type, support `PayAsYouGo`, `Subscription`, `Dev`.
     * 
     */
    public Optional<Output<String>> productType() {
        return Optional.ofNullable(this.productType);
    }

    /**
     * The name of the project
     * 
     */
    @Import(name="projectName")
    private @Nullable Output<String> projectName;

    /**
     * @return The name of the project
     * 
     */
    public Optional<Output<String>> projectName() {
        return Optional.ofNullable(this.projectName);
    }

    /**
     * Project base attributesSee the following `Block Properties`.
     * 
     */
    @Import(name="properties")
    private @Nullable Output<ProjectPropertiesArgs> properties;

    /**
     * @return Project base attributesSee the following `Block Properties`.
     * 
     */
    public Optional<Output<ProjectPropertiesArgs>> properties() {
        return Optional.ofNullable(this.properties);
    }

    /**
     * Security-related attributesSee the following `Block SecurityProperties`.
     * 
     */
    @Import(name="securityProperties")
    private @Nullable Output<ProjectSecurityPropertiesArgs> securityProperties;

    /**
     * @return Security-related attributesSee the following `Block SecurityProperties`.
     * 
     */
    public Optional<Output<ProjectSecurityPropertiesArgs>> securityProperties() {
        return Optional.ofNullable(this.securityProperties);
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

    /**
     * Life cycle type.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Life cycle type.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private ProjectState() {}

    private ProjectState(ProjectState $) {
        this.comment = $.comment;
        this.defaultQuota = $.defaultQuota;
        this.ipWhiteList = $.ipWhiteList;
        this.owner = $.owner;
        this.productType = $.productType;
        this.projectName = $.projectName;
        this.properties = $.properties;
        this.securityProperties = $.securityProperties;
        this.status = $.status;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectState $;

        public Builder() {
            $ = new ProjectState();
        }

        public Builder(ProjectState defaults) {
            $ = new ProjectState(Objects.requireNonNull(defaults));
        }

        /**
         * @param comment Comments of project
         * 
         * @return builder
         * 
         */
        public Builder comment(@Nullable Output<String> comment) {
            $.comment = comment;
            return this;
        }

        /**
         * @param comment Comments of project
         * 
         * @return builder
         * 
         */
        public Builder comment(String comment) {
            return comment(Output.of(comment));
        }

        /**
         * @param defaultQuota Default Computing Resource Group
         * 
         * @return builder
         * 
         */
        public Builder defaultQuota(@Nullable Output<String> defaultQuota) {
            $.defaultQuota = defaultQuota;
            return this;
        }

        /**
         * @param defaultQuota Default Computing Resource Group
         * 
         * @return builder
         * 
         */
        public Builder defaultQuota(String defaultQuota) {
            return defaultQuota(Output.of(defaultQuota));
        }

        /**
         * @param ipWhiteList IP whitelistSee the following `Block IpWhiteList`.
         * 
         * @return builder
         * 
         */
        public Builder ipWhiteList(@Nullable Output<ProjectIpWhiteListArgs> ipWhiteList) {
            $.ipWhiteList = ipWhiteList;
            return this;
        }

        /**
         * @param ipWhiteList IP whitelistSee the following `Block IpWhiteList`.
         * 
         * @return builder
         * 
         */
        public Builder ipWhiteList(ProjectIpWhiteListArgs ipWhiteList) {
            return ipWhiteList(Output.of(ipWhiteList));
        }

        /**
         * @param owner Project owner
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner Project owner
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        /**
         * @param productType Quota payment type, support `PayAsYouGo`, `Subscription`, `Dev`.
         * 
         * @return builder
         * 
         */
        public Builder productType(@Nullable Output<String> productType) {
            $.productType = productType;
            return this;
        }

        /**
         * @param productType Quota payment type, support `PayAsYouGo`, `Subscription`, `Dev`.
         * 
         * @return builder
         * 
         */
        public Builder productType(String productType) {
            return productType(Output.of(productType));
        }

        /**
         * @param projectName The name of the project
         * 
         * @return builder
         * 
         */
        public Builder projectName(@Nullable Output<String> projectName) {
            $.projectName = projectName;
            return this;
        }

        /**
         * @param projectName The name of the project
         * 
         * @return builder
         * 
         */
        public Builder projectName(String projectName) {
            return projectName(Output.of(projectName));
        }

        /**
         * @param properties Project base attributesSee the following `Block Properties`.
         * 
         * @return builder
         * 
         */
        public Builder properties(@Nullable Output<ProjectPropertiesArgs> properties) {
            $.properties = properties;
            return this;
        }

        /**
         * @param properties Project base attributesSee the following `Block Properties`.
         * 
         * @return builder
         * 
         */
        public Builder properties(ProjectPropertiesArgs properties) {
            return properties(Output.of(properties));
        }

        /**
         * @param securityProperties Security-related attributesSee the following `Block SecurityProperties`.
         * 
         * @return builder
         * 
         */
        public Builder securityProperties(@Nullable Output<ProjectSecurityPropertiesArgs> securityProperties) {
            $.securityProperties = securityProperties;
            return this;
        }

        /**
         * @param securityProperties Security-related attributesSee the following `Block SecurityProperties`.
         * 
         * @return builder
         * 
         */
        public Builder securityProperties(ProjectSecurityPropertiesArgs securityProperties) {
            return securityProperties(Output.of(securityProperties));
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

        /**
         * @param type Life cycle type.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Life cycle type.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ProjectState build() {
            return $;
        }
    }

}
