// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.pvtz.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class EndpointIpConfig {
    /**
     * @return The Subnet mask.
     * 
     */
    private String cidrBlock;
    /**
     * @return The IP address within the parameter range of the subnet mask.  It is recommended to use the IP address assigned by the system.
     * 
     */
    private @Nullable String ip;
    /**
     * @return The Vswitch id.
     * 
     */
    private String vswitchId;
    /**
     * @return The Zone ID.
     * 
     */
    private String zoneId;

    private EndpointIpConfig() {}
    /**
     * @return The Subnet mask.
     * 
     */
    public String cidrBlock() {
        return this.cidrBlock;
    }
    /**
     * @return The IP address within the parameter range of the subnet mask.  It is recommended to use the IP address assigned by the system.
     * 
     */
    public Optional<String> ip() {
        return Optional.ofNullable(this.ip);
    }
    /**
     * @return The Vswitch id.
     * 
     */
    public String vswitchId() {
        return this.vswitchId;
    }
    /**
     * @return The Zone ID.
     * 
     */
    public String zoneId() {
        return this.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(EndpointIpConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String cidrBlock;
        private @Nullable String ip;
        private String vswitchId;
        private String zoneId;
        public Builder() {}
        public Builder(EndpointIpConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cidrBlock = defaults.cidrBlock;
    	      this.ip = defaults.ip;
    	      this.vswitchId = defaults.vswitchId;
    	      this.zoneId = defaults.zoneId;
        }

        @CustomType.Setter
        public Builder cidrBlock(String cidrBlock) {
            this.cidrBlock = Objects.requireNonNull(cidrBlock);
            return this;
        }
        @CustomType.Setter
        public Builder ip(@Nullable String ip) {
            this.ip = ip;
            return this;
        }
        @CustomType.Setter
        public Builder vswitchId(String vswitchId) {
            this.vswitchId = Objects.requireNonNull(vswitchId);
            return this;
        }
        @CustomType.Setter
        public Builder zoneId(String zoneId) {
            this.zoneId = Objects.requireNonNull(zoneId);
            return this;
        }
        public EndpointIpConfig build() {
            final var o = new EndpointIpConfig();
            o.cidrBlock = cidrBlock;
            o.ip = ip;
            o.vswitchId = vswitchId;
            o.zoneId = zoneId;
            return o;
        }
    }
}
