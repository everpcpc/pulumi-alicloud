// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupLabelMatchExpression {
    /**
     * @return The key of the tag of the dispatch rule. Valud values:
     * * _aliyun_arms_userid: user ID
     * * _aliyun_arms_involvedObject_kind: type of the associated object
     * * _aliyun_arms_involvedObject_id: ID of the associated object
     * * _aliyun_arms_involvedObject_name: name of the associated object
     * * _aliyun_arms_alert_name: alert name
     * * _aliyun_arms_alert_rule_id: alert rule ID
     * * _aliyun_arms_alert_type: alert type
     * * _aliyun_arms_alert_level: alert severity
     * 
     */
    private String key;
    /**
     * @return The operator used in the dispatch rule. Valid values:
     * * eq: equals to.
     * * re: matches a regular expression.
     * 
     */
    private String operator;
    /**
     * @return The value of the tag.
     * 
     */
    private String value;

    private DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupLabelMatchExpression() {}
    /**
     * @return The key of the tag of the dispatch rule. Valud values:
     * * _aliyun_arms_userid: user ID
     * * _aliyun_arms_involvedObject_kind: type of the associated object
     * * _aliyun_arms_involvedObject_id: ID of the associated object
     * * _aliyun_arms_involvedObject_name: name of the associated object
     * * _aliyun_arms_alert_name: alert name
     * * _aliyun_arms_alert_rule_id: alert rule ID
     * * _aliyun_arms_alert_type: alert type
     * * _aliyun_arms_alert_level: alert severity
     * 
     */
    public String key() {
        return this.key;
    }
    /**
     * @return The operator used in the dispatch rule. Valid values:
     * * eq: equals to.
     * * re: matches a regular expression.
     * 
     */
    public String operator() {
        return this.operator;
    }
    /**
     * @return The value of the tag.
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupLabelMatchExpression defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String key;
        private String operator;
        private String value;
        public Builder() {}
        public Builder(DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupLabelMatchExpression defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.key = defaults.key;
    	      this.operator = defaults.operator;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder key(String key) {
            this.key = Objects.requireNonNull(key);
            return this;
        }
        @CustomType.Setter
        public Builder operator(String operator) {
            this.operator = Objects.requireNonNull(operator);
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }
        public DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupLabelMatchExpression build() {
            final var o = new DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupLabelMatchExpression();
            o.key = key;
            o.operator = operator;
            o.value = value;
            return o;
        }
    }
}
