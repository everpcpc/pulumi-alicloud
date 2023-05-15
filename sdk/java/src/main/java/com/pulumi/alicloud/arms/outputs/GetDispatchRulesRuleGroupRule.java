// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetDispatchRulesRuleGroupRule {
    private Integer groupId;
    /**
     * @return The duration for which the system waits after the first alert is sent. After the duration, all alerts are sent in a single notification to the handler.
     * 
     */
    private Integer groupInterval;
    /**
     * @return The duration for which the system waits after the first alert is sent. After the duration, all alerts are sent in a single notification to the handler.
     * 
     */
    private Integer groupWaitTime;
    /**
     * @return The fields that are used to group events. Events with the same field content are assigned to a group. Alerts with the same specified grouping field are sent to the handler in separate notifications.
     * 
     */
    private List<String> groupingFields;
    /**
     * @return The silence period of repeated alerts. All alerts are repeatedly sent at specified intervals until the alerts are cleared. The minimum value is 61. Default to 600.
     * 
     */
    private Integer repeatInterval;

    private GetDispatchRulesRuleGroupRule() {}
    public Integer groupId() {
        return this.groupId;
    }
    /**
     * @return The duration for which the system waits after the first alert is sent. After the duration, all alerts are sent in a single notification to the handler.
     * 
     */
    public Integer groupInterval() {
        return this.groupInterval;
    }
    /**
     * @return The duration for which the system waits after the first alert is sent. After the duration, all alerts are sent in a single notification to the handler.
     * 
     */
    public Integer groupWaitTime() {
        return this.groupWaitTime;
    }
    /**
     * @return The fields that are used to group events. Events with the same field content are assigned to a group. Alerts with the same specified grouping field are sent to the handler in separate notifications.
     * 
     */
    public List<String> groupingFields() {
        return this.groupingFields;
    }
    /**
     * @return The silence period of repeated alerts. All alerts are repeatedly sent at specified intervals until the alerts are cleared. The minimum value is 61. Default to 600.
     * 
     */
    public Integer repeatInterval() {
        return this.repeatInterval;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDispatchRulesRuleGroupRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer groupId;
        private Integer groupInterval;
        private Integer groupWaitTime;
        private List<String> groupingFields;
        private Integer repeatInterval;
        public Builder() {}
        public Builder(GetDispatchRulesRuleGroupRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.groupId = defaults.groupId;
    	      this.groupInterval = defaults.groupInterval;
    	      this.groupWaitTime = defaults.groupWaitTime;
    	      this.groupingFields = defaults.groupingFields;
    	      this.repeatInterval = defaults.repeatInterval;
        }

        @CustomType.Setter
        public Builder groupId(Integer groupId) {
            this.groupId = Objects.requireNonNull(groupId);
            return this;
        }
        @CustomType.Setter
        public Builder groupInterval(Integer groupInterval) {
            this.groupInterval = Objects.requireNonNull(groupInterval);
            return this;
        }
        @CustomType.Setter
        public Builder groupWaitTime(Integer groupWaitTime) {
            this.groupWaitTime = Objects.requireNonNull(groupWaitTime);
            return this;
        }
        @CustomType.Setter
        public Builder groupingFields(List<String> groupingFields) {
            this.groupingFields = Objects.requireNonNull(groupingFields);
            return this;
        }
        public Builder groupingFields(String... groupingFields) {
            return groupingFields(List.of(groupingFields));
        }
        @CustomType.Setter
        public Builder repeatInterval(Integer repeatInterval) {
            this.repeatInterval = Objects.requireNonNull(repeatInterval);
            return this;
        }
        public GetDispatchRulesRuleGroupRule build() {
            final var o = new GetDispatchRulesRuleGroupRule();
            o.groupId = groupId;
            o.groupInterval = groupInterval;
            o.groupWaitTime = groupWaitTime;
            o.groupingFields = groupingFields;
            o.repeatInterval = repeatInterval;
            return o;
        }
    }
}
