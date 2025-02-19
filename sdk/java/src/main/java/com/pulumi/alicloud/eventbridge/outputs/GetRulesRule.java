// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eventbridge.outputs;

import com.pulumi.alicloud.eventbridge.outputs.GetRulesRuleTarget;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetRulesRule {
    /**
     * @return The description of rule.
     * 
     */
    private String description;
    /**
     * @return The name of event bus.
     * 
     */
    private String eventBusName;
    /**
     * @return The pattern to match interested events.
     * 
     */
    private String filterPattern;
    /**
     * @return The ID of the Rule.
     * 
     */
    private String id;
    /**
     * @return The name of rule.
     * 
     */
    private String ruleName;
    /**
     * @return Rule status, either Enable or Disable.
     * 
     */
    private String status;
    /**
     * @return The target for rule.
     * 
     */
    private List<GetRulesRuleTarget> targets;

    private GetRulesRule() {}
    /**
     * @return The description of rule.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The name of event bus.
     * 
     */
    public String eventBusName() {
        return this.eventBusName;
    }
    /**
     * @return The pattern to match interested events.
     * 
     */
    public String filterPattern() {
        return this.filterPattern;
    }
    /**
     * @return The ID of the Rule.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of rule.
     * 
     */
    public String ruleName() {
        return this.ruleName;
    }
    /**
     * @return Rule status, either Enable or Disable.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The target for rule.
     * 
     */
    public List<GetRulesRuleTarget> targets() {
        return this.targets;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRulesRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String eventBusName;
        private String filterPattern;
        private String id;
        private String ruleName;
        private String status;
        private List<GetRulesRuleTarget> targets;
        public Builder() {}
        public Builder(GetRulesRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.eventBusName = defaults.eventBusName;
    	      this.filterPattern = defaults.filterPattern;
    	      this.id = defaults.id;
    	      this.ruleName = defaults.ruleName;
    	      this.status = defaults.status;
    	      this.targets = defaults.targets;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder eventBusName(String eventBusName) {
            this.eventBusName = Objects.requireNonNull(eventBusName);
            return this;
        }
        @CustomType.Setter
        public Builder filterPattern(String filterPattern) {
            this.filterPattern = Objects.requireNonNull(filterPattern);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ruleName(String ruleName) {
            this.ruleName = Objects.requireNonNull(ruleName);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder targets(List<GetRulesRuleTarget> targets) {
            this.targets = Objects.requireNonNull(targets);
            return this;
        }
        public Builder targets(GetRulesRuleTarget... targets) {
            return targets(List.of(targets));
        }
        public GetRulesRule build() {
            final var o = new GetRulesRule();
            o.description = description;
            o.eventBusName = eventBusName;
            o.filterPattern = filterPattern;
            o.id = id;
            o.ruleName = ruleName;
            o.status = status;
            o.targets = targets;
            return o;
        }
    }
}
