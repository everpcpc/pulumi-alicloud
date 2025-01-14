// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms.outputs;

import com.pulumi.alicloud.arms.outputs.GetDispatchRulesRuleLabelMatchExpressionGridLabelMatchExpressionGroupLabelMatchExpression;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetDispatchRulesRuleLabelMatchExpressionGridLabelMatchExpressionGroup {
    /**
     * @return Sets the dispatch rule.
     * 
     */
    private List<GetDispatchRulesRuleLabelMatchExpressionGridLabelMatchExpressionGroupLabelMatchExpression> labelMatchExpressions;

    private GetDispatchRulesRuleLabelMatchExpressionGridLabelMatchExpressionGroup() {}
    /**
     * @return Sets the dispatch rule.
     * 
     */
    public List<GetDispatchRulesRuleLabelMatchExpressionGridLabelMatchExpressionGroupLabelMatchExpression> labelMatchExpressions() {
        return this.labelMatchExpressions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDispatchRulesRuleLabelMatchExpressionGridLabelMatchExpressionGroup defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetDispatchRulesRuleLabelMatchExpressionGridLabelMatchExpressionGroupLabelMatchExpression> labelMatchExpressions;
        public Builder() {}
        public Builder(GetDispatchRulesRuleLabelMatchExpressionGridLabelMatchExpressionGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.labelMatchExpressions = defaults.labelMatchExpressions;
        }

        @CustomType.Setter
        public Builder labelMatchExpressions(List<GetDispatchRulesRuleLabelMatchExpressionGridLabelMatchExpressionGroupLabelMatchExpression> labelMatchExpressions) {
            this.labelMatchExpressions = Objects.requireNonNull(labelMatchExpressions);
            return this;
        }
        public Builder labelMatchExpressions(GetDispatchRulesRuleLabelMatchExpressionGridLabelMatchExpressionGroupLabelMatchExpression... labelMatchExpressions) {
            return labelMatchExpressions(List.of(labelMatchExpressions));
        }
        public GetDispatchRulesRuleLabelMatchExpressionGridLabelMatchExpressionGroup build() {
            final var o = new GetDispatchRulesRuleLabelMatchExpressionGridLabelMatchExpressionGroup();
            o.labelMatchExpressions = labelMatchExpressions;
            return o;
        }
    }
}
