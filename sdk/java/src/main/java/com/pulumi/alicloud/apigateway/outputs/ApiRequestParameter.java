// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.apigateway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ApiRequestParameter {
    /**
     * @return The default value of the parameter.
     * 
     */
    private @Nullable String defaultValue;
    /**
     * @return The description of the api. Defaults to null.
     * 
     */
    private @Nullable String description;
    /**
     * @return Request&#39;s parameter location; values: BODY, HEAD, QUERY, and PATH.
     * 
     */
    private String in;
    /**
     * @return Backend service&#39;s parameter location; values: BODY, HEAD, QUERY, and PATH.
     * 
     */
    private String inService;
    /**
     * @return The name of the api gateway api. Defaults to null.
     * 
     */
    private String name;
    /**
     * @return Backend service&#39;s parameter name.
     * 
     */
    private String nameService;
    /**
     * @return Parameter required or not; values: REQUIRED and OPTIONAL.
     * 
     */
    private String required;
    /**
     * @return Parameter type which supports values of &#39;STRING&#39;,&#39;INT&#39;,&#39;BOOLEAN&#39;,&#39;LONG&#39;,&#34;FLOAT&#34; and &#34;DOUBLE&#34;.
     * 
     */
    private String type;

    private ApiRequestParameter() {}
    /**
     * @return The default value of the parameter.
     * 
     */
    public Optional<String> defaultValue() {
        return Optional.ofNullable(this.defaultValue);
    }
    /**
     * @return The description of the api. Defaults to null.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return Request&#39;s parameter location; values: BODY, HEAD, QUERY, and PATH.
     * 
     */
    public String in() {
        return this.in;
    }
    /**
     * @return Backend service&#39;s parameter location; values: BODY, HEAD, QUERY, and PATH.
     * 
     */
    public String inService() {
        return this.inService;
    }
    /**
     * @return The name of the api gateway api. Defaults to null.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Backend service&#39;s parameter name.
     * 
     */
    public String nameService() {
        return this.nameService;
    }
    /**
     * @return Parameter required or not; values: REQUIRED and OPTIONAL.
     * 
     */
    public String required() {
        return this.required;
    }
    /**
     * @return Parameter type which supports values of &#39;STRING&#39;,&#39;INT&#39;,&#39;BOOLEAN&#39;,&#39;LONG&#39;,&#34;FLOAT&#34; and &#34;DOUBLE&#34;.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ApiRequestParameter defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String defaultValue;
        private @Nullable String description;
        private String in;
        private String inService;
        private String name;
        private String nameService;
        private String required;
        private String type;
        public Builder() {}
        public Builder(ApiRequestParameter defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.defaultValue = defaults.defaultValue;
    	      this.description = defaults.description;
    	      this.in = defaults.in;
    	      this.inService = defaults.inService;
    	      this.name = defaults.name;
    	      this.nameService = defaults.nameService;
    	      this.required = defaults.required;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder defaultValue(@Nullable String defaultValue) {
            this.defaultValue = defaultValue;
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder in(String in) {
            this.in = Objects.requireNonNull(in);
            return this;
        }
        @CustomType.Setter
        public Builder inService(String inService) {
            this.inService = Objects.requireNonNull(inService);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder nameService(String nameService) {
            this.nameService = Objects.requireNonNull(nameService);
            return this;
        }
        @CustomType.Setter
        public Builder required(String required) {
            this.required = Objects.requireNonNull(required);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public ApiRequestParameter build() {
            final var o = new ApiRequestParameter();
            o.defaultValue = defaultValue;
            o.description = description;
            o.in = in;
            o.inService = inService;
            o.name = name;
            o.nameService = nameService;
            o.required = required;
            o.type = type;
            return o;
        }
    }
}
