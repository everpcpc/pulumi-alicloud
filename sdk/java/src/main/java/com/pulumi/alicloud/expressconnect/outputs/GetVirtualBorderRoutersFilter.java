// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.expressconnect.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetVirtualBorderRoutersFilter {
    /**
     * @return The key of the field to filter by, as defined by
     * [Alibaba Cloud API](https://www.alibabacloud.com/help/en/doc-detail/124791.htm).
     * 
     */
    private @Nullable String key;
    /**
     * @return Set of values that are accepted for the given field.
     * 
     */
    private @Nullable List<String> values;

    private GetVirtualBorderRoutersFilter() {}
    /**
     * @return The key of the field to filter by, as defined by
     * [Alibaba Cloud API](https://www.alibabacloud.com/help/en/doc-detail/124791.htm).
     * 
     */
    public Optional<String> key() {
        return Optional.ofNullable(this.key);
    }
    /**
     * @return Set of values that are accepted for the given field.
     * 
     */
    public List<String> values() {
        return this.values == null ? List.of() : this.values;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVirtualBorderRoutersFilter defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String key;
        private @Nullable List<String> values;
        public Builder() {}
        public Builder(GetVirtualBorderRoutersFilter defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.key = defaults.key;
    	      this.values = defaults.values;
        }

        @CustomType.Setter
        public Builder key(@Nullable String key) {
            this.key = key;
            return this;
        }
        @CustomType.Setter
        public Builder values(@Nullable List<String> values) {
            this.values = values;
            return this;
        }
        public Builder values(String... values) {
            return values(List.of(values));
        }
        public GetVirtualBorderRoutersFilter build() {
            final var o = new GetVirtualBorderRoutersFilter();
            o.key = key;
            o.values = values;
            return o;
        }
    }
}