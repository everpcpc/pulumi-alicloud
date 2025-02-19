// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cdn.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DomainCacheConfig {
    private String cacheContent;
    private @Nullable String cacheId;
    private String cacheType;
    private Integer ttl;
    private @Nullable Integer weight;

    private DomainCacheConfig() {}
    public String cacheContent() {
        return this.cacheContent;
    }
    public Optional<String> cacheId() {
        return Optional.ofNullable(this.cacheId);
    }
    public String cacheType() {
        return this.cacheType;
    }
    public Integer ttl() {
        return this.ttl;
    }
    public Optional<Integer> weight() {
        return Optional.ofNullable(this.weight);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DomainCacheConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String cacheContent;
        private @Nullable String cacheId;
        private String cacheType;
        private Integer ttl;
        private @Nullable Integer weight;
        public Builder() {}
        public Builder(DomainCacheConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cacheContent = defaults.cacheContent;
    	      this.cacheId = defaults.cacheId;
    	      this.cacheType = defaults.cacheType;
    	      this.ttl = defaults.ttl;
    	      this.weight = defaults.weight;
        }

        @CustomType.Setter
        public Builder cacheContent(String cacheContent) {
            this.cacheContent = Objects.requireNonNull(cacheContent);
            return this;
        }
        @CustomType.Setter
        public Builder cacheId(@Nullable String cacheId) {
            this.cacheId = cacheId;
            return this;
        }
        @CustomType.Setter
        public Builder cacheType(String cacheType) {
            this.cacheType = Objects.requireNonNull(cacheType);
            return this;
        }
        @CustomType.Setter
        public Builder ttl(Integer ttl) {
            this.ttl = Objects.requireNonNull(ttl);
            return this;
        }
        @CustomType.Setter
        public Builder weight(@Nullable Integer weight) {
            this.weight = weight;
            return this;
        }
        public DomainCacheConfig build() {
            final var o = new DomainCacheConfig();
            o.cacheContent = cacheContent;
            o.cacheId = cacheId;
            o.cacheType = cacheType;
            o.ttl = ttl;
            o.weight = weight;
            return o;
        }
    }
}
