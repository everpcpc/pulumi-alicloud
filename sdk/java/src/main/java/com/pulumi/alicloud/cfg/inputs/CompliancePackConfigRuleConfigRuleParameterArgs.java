// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cfg.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CompliancePackConfigRuleConfigRuleParameterArgs extends com.pulumi.resources.ResourceArgs {

    public static final CompliancePackConfigRuleConfigRuleParameterArgs Empty = new CompliancePackConfigRuleConfigRuleParameterArgs();

    /**
     * The parameter name.
     * 
     */
    @Import(name="parameterName")
    private @Nullable Output<String> parameterName;

    /**
     * @return The parameter name.
     * 
     */
    public Optional<Output<String>> parameterName() {
        return Optional.ofNullable(this.parameterName);
    }

    /**
     * The parameter value.
     * 
     */
    @Import(name="parameterValue")
    private @Nullable Output<String> parameterValue;

    /**
     * @return The parameter value.
     * 
     */
    public Optional<Output<String>> parameterValue() {
        return Optional.ofNullable(this.parameterValue);
    }

    private CompliancePackConfigRuleConfigRuleParameterArgs() {}

    private CompliancePackConfigRuleConfigRuleParameterArgs(CompliancePackConfigRuleConfigRuleParameterArgs $) {
        this.parameterName = $.parameterName;
        this.parameterValue = $.parameterValue;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CompliancePackConfigRuleConfigRuleParameterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CompliancePackConfigRuleConfigRuleParameterArgs $;

        public Builder() {
            $ = new CompliancePackConfigRuleConfigRuleParameterArgs();
        }

        public Builder(CompliancePackConfigRuleConfigRuleParameterArgs defaults) {
            $ = new CompliancePackConfigRuleConfigRuleParameterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param parameterName The parameter name.
         * 
         * @return builder
         * 
         */
        public Builder parameterName(@Nullable Output<String> parameterName) {
            $.parameterName = parameterName;
            return this;
        }

        /**
         * @param parameterName The parameter name.
         * 
         * @return builder
         * 
         */
        public Builder parameterName(String parameterName) {
            return parameterName(Output.of(parameterName));
        }

        /**
         * @param parameterValue The parameter value.
         * 
         * @return builder
         * 
         */
        public Builder parameterValue(@Nullable Output<String> parameterValue) {
            $.parameterValue = parameterValue;
            return this;
        }

        /**
         * @param parameterValue The parameter value.
         * 
         * @return builder
         * 
         */
        public Builder parameterValue(String parameterValue) {
            return parameterValue(Output.of(parameterValue));
        }

        public CompliancePackConfigRuleConfigRuleParameterArgs build() {
            return $;
        }
    }

}