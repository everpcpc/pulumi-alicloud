// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.scdn.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DomainSource {
    /**
     * @return The Back-to-Source Address.
     * 
     */
    private String content;
    /**
     * @return The source status. Valid values: online, offline.
     * 
     */
    private @Nullable String enabled;
    /**
     * @return Port.
     * 
     */
    private Integer port;
    /**
     * @return Priority.
     * 
     */
    private String priority;
    /**
     * @return The Origin Server Type. Valid Values:
     * * ipaddr: IP Source Station
     * * domain: the Domain Name
     * * oss: OSS Bucket as a Source Station.
     * 
     */
    private String type;

    private DomainSource() {}
    /**
     * @return The Back-to-Source Address.
     * 
     */
    public String content() {
        return this.content;
    }
    /**
     * @return The source status. Valid values: online, offline.
     * 
     */
    public Optional<String> enabled() {
        return Optional.ofNullable(this.enabled);
    }
    /**
     * @return Port.
     * 
     */
    public Integer port() {
        return this.port;
    }
    /**
     * @return Priority.
     * 
     */
    public String priority() {
        return this.priority;
    }
    /**
     * @return The Origin Server Type. Valid Values:
     * * ipaddr: IP Source Station
     * * domain: the Domain Name
     * * oss: OSS Bucket as a Source Station.
     * 
     */
    public String type() {
        return this.type;
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
        private @Nullable String enabled;
        private Integer port;
        private String priority;
        private String type;
        public Builder() {}
        public Builder(DomainSource defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.content = defaults.content;
    	      this.enabled = defaults.enabled;
    	      this.port = defaults.port;
    	      this.priority = defaults.priority;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder content(String content) {
            this.content = Objects.requireNonNull(content);
            return this;
        }
        @CustomType.Setter
        public Builder enabled(@Nullable String enabled) {
            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        @CustomType.Setter
        public Builder priority(String priority) {
            this.priority = Objects.requireNonNull(priority);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public DomainSource build() {
            final var o = new DomainSource();
            o.content = content;
            o.enabled = enabled;
            o.port = port;
            o.priority = priority;
            o.type = type;
            return o;
        }
    }
}
