// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ros.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class StackGroupParameter {
    /**
     * @return The parameter key.
     * 
     */
    private @Nullable String parameterKey;
    /**
     * @return The parameter value.
     * 
     */
    private @Nullable String parameterValue;

    private StackGroupParameter() {}
    /**
     * @return The parameter key.
     * 
     */
    public Optional<String> parameterKey() {
        return Optional.ofNullable(this.parameterKey);
    }
    /**
     * @return The parameter value.
     * 
     */
    public Optional<String> parameterValue() {
        return Optional.ofNullable(this.parameterValue);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(StackGroupParameter defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String parameterKey;
        private @Nullable String parameterValue;
        public Builder() {}
        public Builder(StackGroupParameter defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.parameterKey = defaults.parameterKey;
    	      this.parameterValue = defaults.parameterValue;
        }

        @CustomType.Setter
        public Builder parameterKey(@Nullable String parameterKey) {
            this.parameterKey = parameterKey;
            return this;
        }
        @CustomType.Setter
        public Builder parameterValue(@Nullable String parameterValue) {
            this.parameterValue = parameterValue;
            return this;
        }
        public StackGroupParameter build() {
            final var o = new StackGroupParameter();
            o.parameterKey = parameterKey;
            o.parameterValue = parameterValue;
            return o;
        }
    }
}