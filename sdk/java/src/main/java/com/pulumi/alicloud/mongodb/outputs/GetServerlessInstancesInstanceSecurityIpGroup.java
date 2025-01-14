// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mongodb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetServerlessInstancesInstanceSecurityIpGroup {
    /**
     * @return The attribute of the IP whitelist. This parameter is empty by default.
     * 
     */
    private String securityIpGroupAttribute;
    /**
     * @return The name of the IP whitelist.
     * 
     */
    private String securityIpGroupName;
    /**
     * @return The IP addresses in the whitelist.
     * 
     */
    private String securityIpList;

    private GetServerlessInstancesInstanceSecurityIpGroup() {}
    /**
     * @return The attribute of the IP whitelist. This parameter is empty by default.
     * 
     */
    public String securityIpGroupAttribute() {
        return this.securityIpGroupAttribute;
    }
    /**
     * @return The name of the IP whitelist.
     * 
     */
    public String securityIpGroupName() {
        return this.securityIpGroupName;
    }
    /**
     * @return The IP addresses in the whitelist.
     * 
     */
    public String securityIpList() {
        return this.securityIpList;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServerlessInstancesInstanceSecurityIpGroup defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String securityIpGroupAttribute;
        private String securityIpGroupName;
        private String securityIpList;
        public Builder() {}
        public Builder(GetServerlessInstancesInstanceSecurityIpGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.securityIpGroupAttribute = defaults.securityIpGroupAttribute;
    	      this.securityIpGroupName = defaults.securityIpGroupName;
    	      this.securityIpList = defaults.securityIpList;
        }

        @CustomType.Setter
        public Builder securityIpGroupAttribute(String securityIpGroupAttribute) {
            this.securityIpGroupAttribute = Objects.requireNonNull(securityIpGroupAttribute);
            return this;
        }
        @CustomType.Setter
        public Builder securityIpGroupName(String securityIpGroupName) {
            this.securityIpGroupName = Objects.requireNonNull(securityIpGroupName);
            return this;
        }
        @CustomType.Setter
        public Builder securityIpList(String securityIpList) {
            this.securityIpList = Objects.requireNonNull(securityIpList);
            return this;
        }
        public GetServerlessInstancesInstanceSecurityIpGroup build() {
            final var o = new GetServerlessInstancesInstanceSecurityIpGroup();
            o.securityIpGroupAttribute = securityIpGroupAttribute;
            o.securityIpGroupName = securityIpGroupName;
            o.securityIpList = securityIpList;
            return o;
        }
    }
}
