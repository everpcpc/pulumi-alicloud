// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.graphdatabase.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DbInstanceDbInstanceIpArray {
    /**
     * @return The default is empty. To distinguish between the different property console does not display a `hidden` label grouping.
     * 
     */
    private @Nullable String dbInstanceIpArrayAttribute;
    /**
     * @return IP ADDRESS whitelist group name.
     * 
     */
    private @Nullable String dbInstanceIpArrayName;
    /**
     * @return IP ADDRESS whitelist addresses in the IP ADDRESS list, and a maximum of 1000 comma-separated format is as follows: `0.0.0.0/0` and `10.23.12.24`(IP) or `10.23.12.24/24`(CIDR mode, CIDR (Classless Inter-Domain Routing)/24 represents the address prefixes in the length of the range [1,32]).
     * 
     */
    private @Nullable String securityIps;

    private DbInstanceDbInstanceIpArray() {}
    /**
     * @return The default is empty. To distinguish between the different property console does not display a `hidden` label grouping.
     * 
     */
    public Optional<String> dbInstanceIpArrayAttribute() {
        return Optional.ofNullable(this.dbInstanceIpArrayAttribute);
    }
    /**
     * @return IP ADDRESS whitelist group name.
     * 
     */
    public Optional<String> dbInstanceIpArrayName() {
        return Optional.ofNullable(this.dbInstanceIpArrayName);
    }
    /**
     * @return IP ADDRESS whitelist addresses in the IP ADDRESS list, and a maximum of 1000 comma-separated format is as follows: `0.0.0.0/0` and `10.23.12.24`(IP) or `10.23.12.24/24`(CIDR mode, CIDR (Classless Inter-Domain Routing)/24 represents the address prefixes in the length of the range [1,32]).
     * 
     */
    public Optional<String> securityIps() {
        return Optional.ofNullable(this.securityIps);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DbInstanceDbInstanceIpArray defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String dbInstanceIpArrayAttribute;
        private @Nullable String dbInstanceIpArrayName;
        private @Nullable String securityIps;
        public Builder() {}
        public Builder(DbInstanceDbInstanceIpArray defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dbInstanceIpArrayAttribute = defaults.dbInstanceIpArrayAttribute;
    	      this.dbInstanceIpArrayName = defaults.dbInstanceIpArrayName;
    	      this.securityIps = defaults.securityIps;
        }

        @CustomType.Setter
        public Builder dbInstanceIpArrayAttribute(@Nullable String dbInstanceIpArrayAttribute) {
            this.dbInstanceIpArrayAttribute = dbInstanceIpArrayAttribute;
            return this;
        }
        @CustomType.Setter
        public Builder dbInstanceIpArrayName(@Nullable String dbInstanceIpArrayName) {
            this.dbInstanceIpArrayName = dbInstanceIpArrayName;
            return this;
        }
        @CustomType.Setter
        public Builder securityIps(@Nullable String securityIps) {
            this.securityIps = securityIps;
            return this;
        }
        public DbInstanceDbInstanceIpArray build() {
            final var o = new DbInstanceDbInstanceIpArray();
            o.dbInstanceIpArrayAttribute = dbInstanceIpArrayAttribute;
            o.dbInstanceIpArrayName = dbInstanceIpArrayName;
            o.securityIps = securityIps;
            return o;
        }
    }
}
