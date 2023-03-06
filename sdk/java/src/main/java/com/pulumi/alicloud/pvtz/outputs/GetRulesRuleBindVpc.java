// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.pvtz.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRulesRuleBindVpc {
    /**
     * @return The region ID of the vpc.
     * 
     */
    private String regionId;
    /**
     * @return The Region Name of the vpc.
     * 
     */
    private String regionName;
    /**
     * @return The ID of the VPC.
     * 
     */
    private String vpcId;
    /**
     * @return The Name of the VPC.
     * 
     */
    private String vpcName;

    private GetRulesRuleBindVpc() {}
    /**
     * @return The region ID of the vpc.
     * 
     */
    public String regionId() {
        return this.regionId;
    }
    /**
     * @return The Region Name of the vpc.
     * 
     */
    public String regionName() {
        return this.regionName;
    }
    /**
     * @return The ID of the VPC.
     * 
     */
    public String vpcId() {
        return this.vpcId;
    }
    /**
     * @return The Name of the VPC.
     * 
     */
    public String vpcName() {
        return this.vpcName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRulesRuleBindVpc defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String regionId;
        private String regionName;
        private String vpcId;
        private String vpcName;
        public Builder() {}
        public Builder(GetRulesRuleBindVpc defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.regionId = defaults.regionId;
    	      this.regionName = defaults.regionName;
    	      this.vpcId = defaults.vpcId;
    	      this.vpcName = defaults.vpcName;
        }

        @CustomType.Setter
        public Builder regionId(String regionId) {
            this.regionId = Objects.requireNonNull(regionId);
            return this;
        }
        @CustomType.Setter
        public Builder regionName(String regionName) {
            this.regionName = Objects.requireNonNull(regionName);
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(String vpcId) {
            this.vpcId = Objects.requireNonNull(vpcId);
            return this;
        }
        @CustomType.Setter
        public Builder vpcName(String vpcName) {
            this.vpcName = Objects.requireNonNull(vpcName);
            return this;
        }
        public GetRulesRuleBindVpc build() {
            final var o = new GetRulesRuleBindVpc();
            o.regionId = regionId;
            o.regionName = regionName;
            o.vpcId = vpcId;
            o.vpcName = vpcName;
            return o;
        }
    }
}