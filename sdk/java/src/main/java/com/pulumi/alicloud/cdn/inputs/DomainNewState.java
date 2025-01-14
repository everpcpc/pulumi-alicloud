// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cdn.inputs;

import com.pulumi.alicloud.cdn.inputs.DomainNewCertificateConfigArgs;
import com.pulumi.alicloud.cdn.inputs.DomainNewSourceArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DomainNewState extends com.pulumi.resources.ResourceArgs {

    public static final DomainNewState Empty = new DomainNewState();

    /**
     * Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
     * 
     */
    @Import(name="cdnType")
    private @Nullable Output<String> cdnType;

    /**
     * @return Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
     * 
     */
    public Optional<Output<String>> cdnType() {
        return Optional.ofNullable(this.cdnType);
    }

    /**
     * Certificate config of the accelerated domain. It&#39;s a list and consist of at most 1 item.
     * 
     */
    @Import(name="certificateConfig")
    private @Nullable Output<DomainNewCertificateConfigArgs> certificateConfig;

    /**
     * @return Certificate config of the accelerated domain. It&#39;s a list and consist of at most 1 item.
     * 
     */
    public Optional<Output<DomainNewCertificateConfigArgs>> certificateConfig() {
        return Optional.ofNullable(this.certificateConfig);
    }

    /**
     * (Available in v1.90.0+) The CNAME of the CDN domain.
     * 
     */
    @Import(name="cname")
    private @Nullable Output<String> cname;

    /**
     * @return (Available in v1.90.0+) The CNAME of the CDN domain.
     * 
     */
    public Optional<Output<String>> cname() {
        return Optional.ofNullable(this.cname);
    }

    /**
     * Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or &#34;-&#34;, and must not begin or end with &#34;-&#34;, and &#34;-&#34; must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     * 
     */
    @Import(name="domainName")
    private @Nullable Output<String> domainName;

    /**
     * @return Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or &#34;-&#34;, and must not begin or end with &#34;-&#34;, and &#34;-&#34; must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     * 
     */
    public Optional<Output<String>> domainName() {
        return Optional.ofNullable(this.domainName);
    }

    /**
     * Resource group ID.
     * 
     */
    @Import(name="resourceGroupId")
    private @Nullable Output<String> resourceGroupId;

    /**
     * @return Resource group ID.
     * 
     */
    public Optional<Output<String>> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }

    /**
     * Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter&#39;s setting is valid Only for the international users and domestic L3 and above users .
     * 
     */
    @Import(name="scope")
    private @Nullable Output<String> scope;

    /**
     * @return Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter&#39;s setting is valid Only for the international users and domestic L3 and above users .
     * 
     */
    public Optional<Output<String>> scope() {
        return Optional.ofNullable(this.scope);
    }

    /**
     * The source address list of the accelerated domain. Defaults to null. See Block Sources.
     * 
     */
    @Import(name="sources")
    private @Nullable Output<List<DomainNewSourceArgs>> sources;

    /**
     * @return The source address list of the accelerated domain. Defaults to null. See Block Sources.
     * 
     */
    public Optional<Output<List<DomainNewSourceArgs>>> sources() {
        return Optional.ofNullable(this.sources);
    }

    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Optional<Output<Map<String,Object>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private DomainNewState() {}

    private DomainNewState(DomainNewState $) {
        this.cdnType = $.cdnType;
        this.certificateConfig = $.certificateConfig;
        this.cname = $.cname;
        this.domainName = $.domainName;
        this.resourceGroupId = $.resourceGroupId;
        this.scope = $.scope;
        this.sources = $.sources;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DomainNewState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DomainNewState $;

        public Builder() {
            $ = new DomainNewState();
        }

        public Builder(DomainNewState defaults) {
            $ = new DomainNewState(Objects.requireNonNull(defaults));
        }

        /**
         * @param cdnType Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
         * 
         * @return builder
         * 
         */
        public Builder cdnType(@Nullable Output<String> cdnType) {
            $.cdnType = cdnType;
            return this;
        }

        /**
         * @param cdnType Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
         * 
         * @return builder
         * 
         */
        public Builder cdnType(String cdnType) {
            return cdnType(Output.of(cdnType));
        }

        /**
         * @param certificateConfig Certificate config of the accelerated domain. It&#39;s a list and consist of at most 1 item.
         * 
         * @return builder
         * 
         */
        public Builder certificateConfig(@Nullable Output<DomainNewCertificateConfigArgs> certificateConfig) {
            $.certificateConfig = certificateConfig;
            return this;
        }

        /**
         * @param certificateConfig Certificate config of the accelerated domain. It&#39;s a list and consist of at most 1 item.
         * 
         * @return builder
         * 
         */
        public Builder certificateConfig(DomainNewCertificateConfigArgs certificateConfig) {
            return certificateConfig(Output.of(certificateConfig));
        }

        /**
         * @param cname (Available in v1.90.0+) The CNAME of the CDN domain.
         * 
         * @return builder
         * 
         */
        public Builder cname(@Nullable Output<String> cname) {
            $.cname = cname;
            return this;
        }

        /**
         * @param cname (Available in v1.90.0+) The CNAME of the CDN domain.
         * 
         * @return builder
         * 
         */
        public Builder cname(String cname) {
            return cname(Output.of(cname));
        }

        /**
         * @param domainName Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or &#34;-&#34;, and must not begin or end with &#34;-&#34;, and &#34;-&#34; must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
         * 
         * @return builder
         * 
         */
        public Builder domainName(@Nullable Output<String> domainName) {
            $.domainName = domainName;
            return this;
        }

        /**
         * @param domainName Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or &#34;-&#34;, and must not begin or end with &#34;-&#34;, and &#34;-&#34; must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
         * 
         * @return builder
         * 
         */
        public Builder domainName(String domainName) {
            return domainName(Output.of(domainName));
        }

        /**
         * @param resourceGroupId Resource group ID.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(@Nullable Output<String> resourceGroupId) {
            $.resourceGroupId = resourceGroupId;
            return this;
        }

        /**
         * @param resourceGroupId Resource group ID.
         * 
         * @return builder
         * 
         */
        public Builder resourceGroupId(String resourceGroupId) {
            return resourceGroupId(Output.of(resourceGroupId));
        }

        /**
         * @param scope Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter&#39;s setting is valid Only for the international users and domestic L3 and above users .
         * 
         * @return builder
         * 
         */
        public Builder scope(@Nullable Output<String> scope) {
            $.scope = scope;
            return this;
        }

        /**
         * @param scope Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter&#39;s setting is valid Only for the international users and domestic L3 and above users .
         * 
         * @return builder
         * 
         */
        public Builder scope(String scope) {
            return scope(Output.of(scope));
        }

        /**
         * @param sources The source address list of the accelerated domain. Defaults to null. See Block Sources.
         * 
         * @return builder
         * 
         */
        public Builder sources(@Nullable Output<List<DomainNewSourceArgs>> sources) {
            $.sources = sources;
            return this;
        }

        /**
         * @param sources The source address list of the accelerated domain. Defaults to null. See Block Sources.
         * 
         * @return builder
         * 
         */
        public Builder sources(List<DomainNewSourceArgs> sources) {
            return sources(Output.of(sources));
        }

        /**
         * @param sources The source address list of the accelerated domain. Defaults to null. See Block Sources.
         * 
         * @return builder
         * 
         */
        public Builder sources(DomainNewSourceArgs... sources) {
            return sources(List.of(sources));
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,Object>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags to assign to the resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,Object> tags) {
            return tags(Output.of(tags));
        }

        public DomainNewState build() {
            return $;
        }
    }

}
