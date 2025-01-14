// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cs.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class NodePoolLabel {
    /**
     * @return The label key.
     * 
     */
    private String key;
    /**
     * @return The label value.
     * 
     */
    private @Nullable String value;

    private NodePoolLabel() {}
    /**
     * @return The label key.
     * 
     */
    public String key() {
        return this.key;
    }
    /**
     * @return The label value.
     * 
     */
    public Optional<String> value() {
        return Optional.ofNullable(this.value);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(NodePoolLabel defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String key;
        private @Nullable String value;
        public Builder() {}
        public Builder(NodePoolLabel defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.key = defaults.key;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder key(String key) {
            this.key = Objects.requireNonNull(key);
            return this;
        }
        @CustomType.Setter
        public Builder value(@Nullable String value) {
            this.value = value;
            return this;
        }
        public NodePoolLabel build() {
            final var o = new NodePoolLabel();
            o.key = key;
            o.value = value;
            return o;
        }
    }
}
