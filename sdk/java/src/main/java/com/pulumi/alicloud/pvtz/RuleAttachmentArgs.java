// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.pvtz;

import com.pulumi.alicloud.pvtz.inputs.RuleAttachmentVpcArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class RuleAttachmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final RuleAttachmentArgs Empty = new RuleAttachmentArgs();

    /**
     * The ID of the rule.
     * 
     */
    @Import(name="ruleId", required=true)
    private Output<String> ruleId;

    /**
     * @return The ID of the rule.
     * 
     */
    public Output<String> ruleId() {
        return this.ruleId;
    }

    /**
     * The List of the VPC. See the following `Block vpcs`.
     * 
     */
    @Import(name="vpcs", required=true)
    private Output<List<RuleAttachmentVpcArgs>> vpcs;

    /**
     * @return The List of the VPC. See the following `Block vpcs`.
     * 
     */
    public Output<List<RuleAttachmentVpcArgs>> vpcs() {
        return this.vpcs;
    }

    private RuleAttachmentArgs() {}

    private RuleAttachmentArgs(RuleAttachmentArgs $) {
        this.ruleId = $.ruleId;
        this.vpcs = $.vpcs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RuleAttachmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RuleAttachmentArgs $;

        public Builder() {
            $ = new RuleAttachmentArgs();
        }

        public Builder(RuleAttachmentArgs defaults) {
            $ = new RuleAttachmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ruleId The ID of the rule.
         * 
         * @return builder
         * 
         */
        public Builder ruleId(Output<String> ruleId) {
            $.ruleId = ruleId;
            return this;
        }

        /**
         * @param ruleId The ID of the rule.
         * 
         * @return builder
         * 
         */
        public Builder ruleId(String ruleId) {
            return ruleId(Output.of(ruleId));
        }

        /**
         * @param vpcs The List of the VPC. See the following `Block vpcs`.
         * 
         * @return builder
         * 
         */
        public Builder vpcs(Output<List<RuleAttachmentVpcArgs>> vpcs) {
            $.vpcs = vpcs;
            return this;
        }

        /**
         * @param vpcs The List of the VPC. See the following `Block vpcs`.
         * 
         * @return builder
         * 
         */
        public Builder vpcs(List<RuleAttachmentVpcArgs> vpcs) {
            return vpcs(Output.of(vpcs));
        }

        /**
         * @param vpcs The List of the VPC. See the following `Block vpcs`.
         * 
         * @return builder
         * 
         */
        public Builder vpcs(RuleAttachmentVpcArgs... vpcs) {
            return vpcs(List.of(vpcs));
        }

        public RuleAttachmentArgs build() {
            $.ruleId = Objects.requireNonNull($.ruleId, "expected parameter 'ruleId' to be non-null");
            $.vpcs = Objects.requireNonNull($.vpcs, "expected parameter 'vpcs' to be non-null");
            return $;
        }
    }

}
