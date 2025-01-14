// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.pvtz.inputs;

import com.pulumi.alicloud.pvtz.inputs.RuleAttachmentVpcArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RuleAttachmentState extends com.pulumi.resources.ResourceArgs {

    public static final RuleAttachmentState Empty = new RuleAttachmentState();

    /**
     * The ID of the rule.
     * 
     */
    @Import(name="ruleId")
    private @Nullable Output<String> ruleId;

    /**
     * @return The ID of the rule.
     * 
     */
    public Optional<Output<String>> ruleId() {
        return Optional.ofNullable(this.ruleId);
    }

    /**
     * The List of the VPC. See the following `Block vpcs`.
     * 
     */
    @Import(name="vpcs")
    private @Nullable Output<List<RuleAttachmentVpcArgs>> vpcs;

    /**
     * @return The List of the VPC. See the following `Block vpcs`.
     * 
     */
    public Optional<Output<List<RuleAttachmentVpcArgs>>> vpcs() {
        return Optional.ofNullable(this.vpcs);
    }

    private RuleAttachmentState() {}

    private RuleAttachmentState(RuleAttachmentState $) {
        this.ruleId = $.ruleId;
        this.vpcs = $.vpcs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RuleAttachmentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RuleAttachmentState $;

        public Builder() {
            $ = new RuleAttachmentState();
        }

        public Builder(RuleAttachmentState defaults) {
            $ = new RuleAttachmentState(Objects.requireNonNull(defaults));
        }

        /**
         * @param ruleId The ID of the rule.
         * 
         * @return builder
         * 
         */
        public Builder ruleId(@Nullable Output<String> ruleId) {
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
        public Builder vpcs(@Nullable Output<List<RuleAttachmentVpcArgs>> vpcs) {
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

        public RuleAttachmentState build() {
            return $;
        }
    }

}
