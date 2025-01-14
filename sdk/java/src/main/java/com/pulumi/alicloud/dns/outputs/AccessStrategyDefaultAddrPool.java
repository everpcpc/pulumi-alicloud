// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AccessStrategyDefaultAddrPool {
    /**
     * @return The ID of the address pool in the secondary address pool group.
     * 
     */
    private String addrPoolId;
    /**
     * @return The weight of the address pool in the secondary address pool group.
     * 
     */
    private @Nullable Integer lbaWeight;

    private AccessStrategyDefaultAddrPool() {}
    /**
     * @return The ID of the address pool in the secondary address pool group.
     * 
     */
    public String addrPoolId() {
        return this.addrPoolId;
    }
    /**
     * @return The weight of the address pool in the secondary address pool group.
     * 
     */
    public Optional<Integer> lbaWeight() {
        return Optional.ofNullable(this.lbaWeight);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AccessStrategyDefaultAddrPool defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String addrPoolId;
        private @Nullable Integer lbaWeight;
        public Builder() {}
        public Builder(AccessStrategyDefaultAddrPool defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.addrPoolId = defaults.addrPoolId;
    	      this.lbaWeight = defaults.lbaWeight;
        }

        @CustomType.Setter
        public Builder addrPoolId(String addrPoolId) {
            this.addrPoolId = Objects.requireNonNull(addrPoolId);
            return this;
        }
        @CustomType.Setter
        public Builder lbaWeight(@Nullable Integer lbaWeight) {
            this.lbaWeight = lbaWeight;
            return this;
        }
        public AccessStrategyDefaultAddrPool build() {
            final var o = new AccessStrategyDefaultAddrPool();
            o.addrPoolId = addrPoolId;
            o.lbaWeight = lbaWeight;
            return o;
        }
    }
}
