// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.pvtz.outputs;

import com.pulumi.alicloud.pvtz.outputs.GetRulesRuleBindVpc;
import com.pulumi.alicloud.pvtz.outputs.GetRulesRuleForwardIp;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetRulesRule {
    /**
     * @return The List of the VPC. See the following `Block bind_vpcs`. **NOTE:** Available in v1.158.0+.
     * 
     */
    private List<GetRulesRuleBindVpc> bindVpcs;
    /**
     * @return The creation time of the resource.
     * 
     */
    private String createTime;
    /**
     * @return The ID of the Endpoint.
     * 
     */
    private String endpointId;
    /**
     * @return The Name of the Endpoint.
     * 
     */
    private String endpointName;
    private List<GetRulesRuleForwardIp> forwardIps;
    /**
     * @return The ID of the Rule.
     * 
     */
    private String id;
    /**
     * @return The first ID of the resource.
     * 
     */
    private String ruleId;
    /**
     * @return The name of the resource.
     * 
     */
    private String ruleName;
    /**
     * @return The type of the rule.
     * 
     */
    private String type;
    /**
     * @return The name of the forwarding zone.
     * 
     */
    private String zoneName;

    private GetRulesRule() {}
    /**
     * @return The List of the VPC. See the following `Block bind_vpcs`. **NOTE:** Available in v1.158.0+.
     * 
     */
    public List<GetRulesRuleBindVpc> bindVpcs() {
        return this.bindVpcs;
    }
    /**
     * @return The creation time of the resource.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The ID of the Endpoint.
     * 
     */
    public String endpointId() {
        return this.endpointId;
    }
    /**
     * @return The Name of the Endpoint.
     * 
     */
    public String endpointName() {
        return this.endpointName;
    }
    public List<GetRulesRuleForwardIp> forwardIps() {
        return this.forwardIps;
    }
    /**
     * @return The ID of the Rule.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The first ID of the resource.
     * 
     */
    public String ruleId() {
        return this.ruleId;
    }
    /**
     * @return The name of the resource.
     * 
     */
    public String ruleName() {
        return this.ruleName;
    }
    /**
     * @return The type of the rule.
     * 
     */
    public String type() {
        return this.type;
    }
    /**
     * @return The name of the forwarding zone.
     * 
     */
    public String zoneName() {
        return this.zoneName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRulesRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetRulesRuleBindVpc> bindVpcs;
        private String createTime;
        private String endpointId;
        private String endpointName;
        private List<GetRulesRuleForwardIp> forwardIps;
        private String id;
        private String ruleId;
        private String ruleName;
        private String type;
        private String zoneName;
        public Builder() {}
        public Builder(GetRulesRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bindVpcs = defaults.bindVpcs;
    	      this.createTime = defaults.createTime;
    	      this.endpointId = defaults.endpointId;
    	      this.endpointName = defaults.endpointName;
    	      this.forwardIps = defaults.forwardIps;
    	      this.id = defaults.id;
    	      this.ruleId = defaults.ruleId;
    	      this.ruleName = defaults.ruleName;
    	      this.type = defaults.type;
    	      this.zoneName = defaults.zoneName;
        }

        @CustomType.Setter
        public Builder bindVpcs(List<GetRulesRuleBindVpc> bindVpcs) {
            this.bindVpcs = Objects.requireNonNull(bindVpcs);
            return this;
        }
        public Builder bindVpcs(GetRulesRuleBindVpc... bindVpcs) {
            return bindVpcs(List.of(bindVpcs));
        }
        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder endpointId(String endpointId) {
            this.endpointId = Objects.requireNonNull(endpointId);
            return this;
        }
        @CustomType.Setter
        public Builder endpointName(String endpointName) {
            this.endpointName = Objects.requireNonNull(endpointName);
            return this;
        }
        @CustomType.Setter
        public Builder forwardIps(List<GetRulesRuleForwardIp> forwardIps) {
            this.forwardIps = Objects.requireNonNull(forwardIps);
            return this;
        }
        public Builder forwardIps(GetRulesRuleForwardIp... forwardIps) {
            return forwardIps(List.of(forwardIps));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ruleId(String ruleId) {
            this.ruleId = Objects.requireNonNull(ruleId);
            return this;
        }
        @CustomType.Setter
        public Builder ruleName(String ruleName) {
            this.ruleName = Objects.requireNonNull(ruleName);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        @CustomType.Setter
        public Builder zoneName(String zoneName) {
            this.zoneName = Objects.requireNonNull(zoneName);
            return this;
        }
        public GetRulesRule build() {
            final var o = new GetRulesRule();
            o.bindVpcs = bindVpcs;
            o.createTime = createTime;
            o.endpointId = endpointId;
            o.endpointName = endpointName;
            o.forwardIps = forwardIps;
            o.id = id;
            o.ruleId = ruleId;
            o.ruleName = ruleName;
            o.type = type;
            o.zoneName = zoneName;
            return o;
        }
    }
}
