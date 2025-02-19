// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dcdn.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DomainSource {
    /**
     * @return The origin address.
     * 
     */
    private String content;
    /**
     * @return The port number. Valid values: `443` and `80`. Default to `80`.
     * 
     */
    private @Nullable Integer port;
    /**
     * @return The priority of the origin if multiple origins are specified. Default to `20`.
     * 
     */
    private @Nullable String priority;
    /**
     * @return The type of the origin. Valid values:
     * `ipaddr`: The origin is configured using an IP address.
     * `domain`: The origin is configured using a domain name.
     * `oss`: The origin is configured using the Internet domain name of an Alibaba Cloud Object Storage Service (OSS) bucket.
     * 
     */
    private String type;
    /**
     * @return The weight of the origin if multiple origins are specified. Default to `10`.
     * 
     */
    private @Nullable String weight;

    private DomainSource() {}
    /**
     * @return The origin address.
     * 
     */
    public String content() {
        return this.content;
    }
    /**
     * @return The port number. Valid values: `443` and `80`. Default to `80`.
     * 
     */
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }
    /**
     * @return The priority of the origin if multiple origins are specified. Default to `20`.
     * 
     */
    public Optional<String> priority() {
        return Optional.ofNullable(this.priority);
    }
    /**
     * @return The type of the origin. Valid values:
     * `ipaddr`: The origin is configured using an IP address.
     * `domain`: The origin is configured using a domain name.
     * `oss`: The origin is configured using the Internet domain name of an Alibaba Cloud Object Storage Service (OSS) bucket.
     * 
     */
    public String type() {
        return this.type;
    }
    /**
     * @return The weight of the origin if multiple origins are specified. Default to `10`.
     * 
     */
    public Optional<String> weight() {
        return Optional.ofNullable(this.weight);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DomainSource defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String content;
        private @Nullable Integer port;
        private @Nullable String priority;
        private String type;
        private @Nullable String weight;
        public Builder() {}
        public Builder(DomainSource defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.content = defaults.content;
    	      this.port = defaults.port;
    	      this.priority = defaults.priority;
    	      this.type = defaults.type;
    	      this.weight = defaults.weight;
        }

        @CustomType.Setter
        public Builder content(String content) {
            this.content = Objects.requireNonNull(content);
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable Integer port) {
            this.port = port;
            return this;
        }
        @CustomType.Setter
        public Builder priority(@Nullable String priority) {
            this.priority = priority;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        @CustomType.Setter
        public Builder weight(@Nullable String weight) {
            this.weight = weight;
            return this;
        }
        public DomainSource build() {
            final var o = new DomainSource();
            o.content = content;
            o.port = port;
            o.priority = priority;
            o.type = type;
            o.weight = weight;
            return o;
        }
    }
}
