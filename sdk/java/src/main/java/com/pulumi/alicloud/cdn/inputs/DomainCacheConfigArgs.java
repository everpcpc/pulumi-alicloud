// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cdn.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DomainCacheConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final DomainCacheConfigArgs Empty = new DomainCacheConfigArgs();

    @Import(name="cacheContent", required=true)
    private Output<String> cacheContent;

    public Output<String> cacheContent() {
        return this.cacheContent;
    }

    @Import(name="cacheId")
    private @Nullable Output<String> cacheId;

    public Optional<Output<String>> cacheId() {
        return Optional.ofNullable(this.cacheId);
    }

    @Import(name="cacheType", required=true)
    private Output<String> cacheType;

    public Output<String> cacheType() {
        return this.cacheType;
    }

    @Import(name="ttl", required=true)
    private Output<Integer> ttl;

    public Output<Integer> ttl() {
        return this.ttl;
    }

    @Import(name="weight")
    private @Nullable Output<Integer> weight;

    public Optional<Output<Integer>> weight() {
        return Optional.ofNullable(this.weight);
    }

    private DomainCacheConfigArgs() {}

    private DomainCacheConfigArgs(DomainCacheConfigArgs $) {
        this.cacheContent = $.cacheContent;
        this.cacheId = $.cacheId;
        this.cacheType = $.cacheType;
        this.ttl = $.ttl;
        this.weight = $.weight;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DomainCacheConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DomainCacheConfigArgs $;

        public Builder() {
            $ = new DomainCacheConfigArgs();
        }

        public Builder(DomainCacheConfigArgs defaults) {
            $ = new DomainCacheConfigArgs(Objects.requireNonNull(defaults));
        }

        public Builder cacheContent(Output<String> cacheContent) {
            $.cacheContent = cacheContent;
            return this;
        }

        public Builder cacheContent(String cacheContent) {
            return cacheContent(Output.of(cacheContent));
        }

        public Builder cacheId(@Nullable Output<String> cacheId) {
            $.cacheId = cacheId;
            return this;
        }

        public Builder cacheId(String cacheId) {
            return cacheId(Output.of(cacheId));
        }

        public Builder cacheType(Output<String> cacheType) {
            $.cacheType = cacheType;
            return this;
        }

        public Builder cacheType(String cacheType) {
            return cacheType(Output.of(cacheType));
        }

        public Builder ttl(Output<Integer> ttl) {
            $.ttl = ttl;
            return this;
        }

        public Builder ttl(Integer ttl) {
            return ttl(Output.of(ttl));
        }

        public Builder weight(@Nullable Output<Integer> weight) {
            $.weight = weight;
            return this;
        }

        public Builder weight(Integer weight) {
            return weight(Output.of(weight));
        }

        public DomainCacheConfigArgs build() {
            $.cacheContent = Objects.requireNonNull($.cacheContent, "expected parameter 'cacheContent' to be non-null");
            $.cacheType = Objects.requireNonNull($.cacheType, "expected parameter 'cacheType' to be non-null");
            $.ttl = Objects.requireNonNull($.ttl, "expected parameter 'ttl' to be non-null");
            return $;
        }
    }

}