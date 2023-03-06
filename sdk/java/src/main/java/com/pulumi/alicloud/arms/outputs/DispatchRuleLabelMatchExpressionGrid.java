// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms.outputs;

import com.pulumi.alicloud.arms.outputs.DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroup;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;

@CustomType
public final class DispatchRuleLabelMatchExpressionGrid {
    /**
     * @return Sets the dispatch rule. See the following `Block label_match_expression_groups`.
     * 
     */
    private List<DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroup> labelMatchExpressionGroups;

    private DispatchRuleLabelMatchExpressionGrid() {}
    /**
     * @return Sets the dispatch rule. See the following `Block label_match_expression_groups`.
     * 
     */
    public List<DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroup> labelMatchExpressionGroups() {
        return this.labelMatchExpressionGroups;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DispatchRuleLabelMatchExpressionGrid defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroup> labelMatchExpressionGroups;
        public Builder() {}
        public Builder(DispatchRuleLabelMatchExpressionGrid defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.labelMatchExpressionGroups = defaults.labelMatchExpressionGroups;
        }

        @CustomType.Setter
        public Builder labelMatchExpressionGroups(List<DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroup> labelMatchExpressionGroups) {
            this.labelMatchExpressionGroups = Objects.requireNonNull(labelMatchExpressionGroups);
            return this;
        }
        public Builder labelMatchExpressionGroups(DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroup... labelMatchExpressionGroups) {
            return labelMatchExpressionGroups(List.of(labelMatchExpressionGroups));
        }
        public DispatchRuleLabelMatchExpressionGrid build() {
            final var o = new DispatchRuleLabelMatchExpressionGrid();
            o.labelMatchExpressionGroups = labelMatchExpressionGroups;
            return o;
        }
    }
}