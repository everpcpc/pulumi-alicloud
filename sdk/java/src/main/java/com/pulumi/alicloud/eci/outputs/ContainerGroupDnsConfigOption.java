// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eci.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ContainerGroupDnsConfigOption {
    /**
     * @return The name of the volume.
     * 
     */
    private @Nullable String name;
    /**
     * @return The value of the variable. The value can be 0 to 256 characters in length.
     * 
     */
    private @Nullable String value;

    private ContainerGroupDnsConfigOption() {}
    /**
     * @return The name of the volume.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
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

    public static Builder builder(ContainerGroupDnsConfigOption defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String name;
        private @Nullable String value;
        public Builder() {}
        public Builder(ContainerGroupDnsConfigOption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder value(@Nullable String value) {
            this.value = value;
            return this;
        }
        public ContainerGroupDnsConfigOption build() {
            final var o = new ContainerGroupDnsConfigOption();
            o.name = name;
            o.value = value;
            return o;
        }
    }
}
