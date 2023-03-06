// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms.inputs;

import com.pulumi.alicloud.arms.inputs.DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.List;
import java.util.Objects;


public final class DispatchRuleLabelMatchExpressionGridArgs extends com.pulumi.resources.ResourceArgs {

    public static final DispatchRuleLabelMatchExpressionGridArgs Empty = new DispatchRuleLabelMatchExpressionGridArgs();

    /**
     * Sets the dispatch rule. See the following `Block label_match_expression_groups`.
     * 
     */
    @Import(name="labelMatchExpressionGroups", required=true)
    private Output<List<DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupArgs>> labelMatchExpressionGroups;

    /**
     * @return Sets the dispatch rule. See the following `Block label_match_expression_groups`.
     * 
     */
    public Output<List<DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupArgs>> labelMatchExpressionGroups() {
        return this.labelMatchExpressionGroups;
    }

    private DispatchRuleLabelMatchExpressionGridArgs() {}

    private DispatchRuleLabelMatchExpressionGridArgs(DispatchRuleLabelMatchExpressionGridArgs $) {
        this.labelMatchExpressionGroups = $.labelMatchExpressionGroups;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DispatchRuleLabelMatchExpressionGridArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DispatchRuleLabelMatchExpressionGridArgs $;

        public Builder() {
            $ = new DispatchRuleLabelMatchExpressionGridArgs();
        }

        public Builder(DispatchRuleLabelMatchExpressionGridArgs defaults) {
            $ = new DispatchRuleLabelMatchExpressionGridArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param labelMatchExpressionGroups Sets the dispatch rule. See the following `Block label_match_expression_groups`.
         * 
         * @return builder
         * 
         */
        public Builder labelMatchExpressionGroups(Output<List<DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupArgs>> labelMatchExpressionGroups) {
            $.labelMatchExpressionGroups = labelMatchExpressionGroups;
            return this;
        }

        /**
         * @param labelMatchExpressionGroups Sets the dispatch rule. See the following `Block label_match_expression_groups`.
         * 
         * @return builder
         * 
         */
        public Builder labelMatchExpressionGroups(List<DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupArgs> labelMatchExpressionGroups) {
            return labelMatchExpressionGroups(Output.of(labelMatchExpressionGroups));
        }

        /**
         * @param labelMatchExpressionGroups Sets the dispatch rule. See the following `Block label_match_expression_groups`.
         * 
         * @return builder
         * 
         */
        public Builder labelMatchExpressionGroups(DispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupArgs... labelMatchExpressionGroups) {
            return labelMatchExpressionGroups(List.of(labelMatchExpressionGroups));
        }

        public DispatchRuleLabelMatchExpressionGridArgs build() {
            $.labelMatchExpressionGroups = Objects.requireNonNull($.labelMatchExpressionGroups, "expected parameter 'labelMatchExpressionGroups' to be non-null");
            return $;
        }
    }

}