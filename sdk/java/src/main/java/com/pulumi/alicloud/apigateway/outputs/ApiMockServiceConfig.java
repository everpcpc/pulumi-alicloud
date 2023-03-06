// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.apigateway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ApiMockServiceConfig {
    private @Nullable String aoneName;
    /**
     * @return The result of the mock service.
     * 
     */
    private String result;

    private ApiMockServiceConfig() {}
    public Optional<String> aoneName() {
        return Optional.ofNullable(this.aoneName);
    }
    /**
     * @return The result of the mock service.
     * 
     */
    public String result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ApiMockServiceConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String aoneName;
        private String result;
        public Builder() {}
        public Builder(ApiMockServiceConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.aoneName = defaults.aoneName;
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder aoneName(@Nullable String aoneName) {
            this.aoneName = aoneName;
            return this;
        }
        @CustomType.Setter
        public Builder result(String result) {
            this.result = Objects.requireNonNull(result);
            return this;
        }
        public ApiMockServiceConfig build() {
            final var o = new ApiMockServiceConfig();
            o.aoneName = aoneName;
            o.result = result;
            return o;
        }
    }
}