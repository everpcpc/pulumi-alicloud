// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae.outputs;

import com.pulumi.alicloud.sae.outputs.ApplicationScalingRuleScalingRuleTimerSchedule;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ApplicationScalingRuleScalingRuleTimer {
    /**
     * @return The Start date. When the `begin_date` and `end_date` values are empty. it indicates long-term execution and is the default value.
     * 
     */
    private @Nullable String beginDate;
    /**
     * @return The End Date. When the `begin_date` and `end_date` values are empty. it indicates long-term execution and is the default value.
     * 
     */
    private @Nullable String endDate;
    /**
     * @return The period in which a timed elastic scaling strategy is executed.
     * 
     */
    private @Nullable String period;
    /**
     * @return Resilient Scaling Strategy Trigger Timing. See the following `Block schedules`.
     * 
     */
    private @Nullable List<ApplicationScalingRuleScalingRuleTimerSchedule> schedules;

    private ApplicationScalingRuleScalingRuleTimer() {}
    /**
     * @return The Start date. When the `begin_date` and `end_date` values are empty. it indicates long-term execution and is the default value.
     * 
     */
    public Optional<String> beginDate() {
        return Optional.ofNullable(this.beginDate);
    }
    /**
     * @return The End Date. When the `begin_date` and `end_date` values are empty. it indicates long-term execution and is the default value.
     * 
     */
    public Optional<String> endDate() {
        return Optional.ofNullable(this.endDate);
    }
    /**
     * @return The period in which a timed elastic scaling strategy is executed.
     * 
     */
    public Optional<String> period() {
        return Optional.ofNullable(this.period);
    }
    /**
     * @return Resilient Scaling Strategy Trigger Timing. See the following `Block schedules`.
     * 
     */
    public List<ApplicationScalingRuleScalingRuleTimerSchedule> schedules() {
        return this.schedules == null ? List.of() : this.schedules;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ApplicationScalingRuleScalingRuleTimer defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String beginDate;
        private @Nullable String endDate;
        private @Nullable String period;
        private @Nullable List<ApplicationScalingRuleScalingRuleTimerSchedule> schedules;
        public Builder() {}
        public Builder(ApplicationScalingRuleScalingRuleTimer defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.beginDate = defaults.beginDate;
    	      this.endDate = defaults.endDate;
    	      this.period = defaults.period;
    	      this.schedules = defaults.schedules;
        }

        @CustomType.Setter
        public Builder beginDate(@Nullable String beginDate) {
            this.beginDate = beginDate;
            return this;
        }
        @CustomType.Setter
        public Builder endDate(@Nullable String endDate) {
            this.endDate = endDate;
            return this;
        }
        @CustomType.Setter
        public Builder period(@Nullable String period) {
            this.period = period;
            return this;
        }
        @CustomType.Setter
        public Builder schedules(@Nullable List<ApplicationScalingRuleScalingRuleTimerSchedule> schedules) {
            this.schedules = schedules;
            return this;
        }
        public Builder schedules(ApplicationScalingRuleScalingRuleTimerSchedule... schedules) {
            return schedules(List.of(schedules));
        }
        public ApplicationScalingRuleScalingRuleTimer build() {
            final var o = new ApplicationScalingRuleScalingRuleTimer();
            o.beginDate = beginDate;
            o.endDate = endDate;
            o.period = period;
            o.schedules = schedules;
            return o;
        }
    }
}
