// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dcdn.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDomainsDomainSource {
    /**
     * @return The origin address.
     * 
     */
    private String content;
    /**
     * @return The status of the origin.
     * 
     */
    private String enabled;
    /**
     * @return The port number.
     * 
     */
    private Integer port;
    /**
     * @return The priority of the origin if multiple origins are specified.
     * 
     */
    private String priority;
    /**
     * @return The type of the origin. Valid values:
     * 
     */
    private String type;
    /**
     * @return The weight of the origin if multiple origins are specified.
     * 
     */
    private String weight;

    private GetDomainsDomainSource() {}
    /**
     * @return The origin address.
     * 
     */
    public String content() {
        return this.content;
    }
    /**
     * @return The status of the origin.
     * 
     */
    public String enabled() {
        return this.enabled;
    }
    /**
     * @return The port number.
     * 
     */
    public Integer port() {
        return this.port;
    }
    /**
     * @return The priority of the origin if multiple origins are specified.
     * 
     */
    public String priority() {
        return this.priority;
    }
    /**
     * @return The type of the origin. Valid values:
     * 
     */
    public String type() {
        return this.type;
    }
    /**
     * @return The weight of the origin if multiple origins are specified.
     * 
     */
    public String weight() {
        return this.weight;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDomainsDomainSource defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String content;
        private String enabled;
        private Integer port;
        private String priority;
        private String type;
        private String weight;
        public Builder() {}
        public Builder(GetDomainsDomainSource defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.content = defaults.content;
    	      this.enabled = defaults.enabled;
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
        public Builder enabled(String enabled) {
            this.enabled = Objects.requireNonNull(enabled);
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
        @CustomType.Setter
        public Builder weight(String weight) {
            this.weight = Objects.requireNonNull(weight);
            return this;
        }
        public GetDomainsDomainSource build() {
            final var o = new GetDomainsDomainSource();
            o.content = content;
            o.enabled = enabled;
            o.port = port;
            o.priority = priority;
            o.type = type;
            o.weight = weight;
            return o;
        }
    }
}
