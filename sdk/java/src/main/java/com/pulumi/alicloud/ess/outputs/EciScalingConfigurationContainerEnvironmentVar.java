// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ess.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class EciScalingConfigurationContainerEnvironmentVar {
    /**
     * @return The name of the variable. The name can be 1 to 128 characters in length and can contain letters,
     * digits, and underscores (_). It cannot start with a digit.
     * digits, and underscores (_). It cannot start with a digit.
     * 
     */
    private @Nullable String key;
    /**
     * @return The value of the variable. The value can be 0 to 256 characters in length.
     * 
     */
    private @Nullable String value;

    private EciScalingConfigurationContainerEnvironmentVar() {}
    /**
     * @return The name of the variable. The name can be 1 to 128 characters in length and can contain letters,
     * digits, and underscores (_). It cannot start with a digit.
     * digits, and underscores (_). It cannot start with a digit.
     * 
     */
    public Optional<String> key() {
        return Optional.ofNullable(this.key);
    }
    /**
     * @return The value of the variable. The value can be 0 to 256 characters in length.
     * 
     */
    public Optional<String> value() {
        return Optional.ofNullable(this.value);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(EciScalingConfigurationContainerEnvironmentVar defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String key;
        private @Nullable String value;
        public Builder() {}
        public Builder(EciScalingConfigurationContainerEnvironmentVar defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.key = defaults.key;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder key(@Nullable String key) {
            this.key = key;
            return this;
        }
        @CustomType.Setter
        public Builder value(@Nullable String value) {
            this.value = value;
            return this;
        }
        public EciScalingConfigurationContainerEnvironmentVar build() {
            final var o = new EciScalingConfigurationContainerEnvironmentVar();
            o.key = key;
            o.value = value;
            return o;
        }
    }
}
