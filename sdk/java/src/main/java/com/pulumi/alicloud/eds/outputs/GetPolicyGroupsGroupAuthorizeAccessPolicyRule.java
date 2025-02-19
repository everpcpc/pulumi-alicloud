// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eds.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetPolicyGroupsGroupAuthorizeAccessPolicyRule {
    /**
     * @return The cidrip of security rules.
     * 
     */
    private String cidrIp;
    /**
     * @return The description of security rules.
     * 
     */
    private String description;

    private GetPolicyGroupsGroupAuthorizeAccessPolicyRule() {}
    /**
     * @return The cidrip of security rules.
     * 
     */
    public String cidrIp() {
        return this.cidrIp;
    }
    /**
     * @return The description of security rules.
     * 
     */
    public String description() {
        return this.description;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPolicyGroupsGroupAuthorizeAccessPolicyRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String cidrIp;
        private String description;
        public Builder() {}
        public Builder(GetPolicyGroupsGroupAuthorizeAccessPolicyRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cidrIp = defaults.cidrIp;
    	      this.description = defaults.description;
        }

        @CustomType.Setter
        public Builder cidrIp(String cidrIp) {
            this.cidrIp = Objects.requireNonNull(cidrIp);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        public GetPolicyGroupsGroupAuthorizeAccessPolicyRule build() {
            final var o = new GetPolicyGroupsGroupAuthorizeAccessPolicyRule();
            o.cidrIp = cidrIp;
            o.description = description;
            return o;
        }
    }
}
