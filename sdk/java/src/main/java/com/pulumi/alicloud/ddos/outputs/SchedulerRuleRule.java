// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ddos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SchedulerRuleRule {
    /**
     * @return The priority of the rule.
     * 
     */
    private @Nullable Integer priority;
    /**
     * @return The region where the interaction resource that is used in the scheduling rule is deployed. **NOTE:** This parameter is returned only if the RuleType parameter is set to 2.
     * 
     */
    private @Nullable String regionId;
    /**
     * @return The status of the scheduling rule.
     * 
     */
    private @Nullable Integer status;
    /**
     * @return The address type of the interaction resource. Valid values:
     * `A`: IPv4 address.
     * `CNAME`: CNAME record.
     * 
     */
    private @Nullable String type;
    /**
     * @return The address of the interaction resource.
     * 
     */
    private @Nullable String value;
    /**
     * @return Required. The type of the linked resource. It is an Integer. Valid values:
     * `1`: The IP address of Anti-DDoS Pro or Anti-DDoS Premium
     * `2`: the IP address of the interaction resource (in the tiered protection scenario)
     * `3`: the IP address used to accelerate access (in the network acceleration scenario)
     * `6` the IP address of the interaction resource (in the cloud service interaction scenario)
     * 
     */
    private @Nullable Integer valueType;

    private SchedulerRuleRule() {}
    /**
     * @return The priority of the rule.
     * 
     */
    public Optional<Integer> priority() {
        return Optional.ofNullable(this.priority);
    }
    /**
     * @return The region where the interaction resource that is used in the scheduling rule is deployed. **NOTE:** This parameter is returned only if the RuleType parameter is set to 2.
     * 
     */
    public Optional<String> regionId() {
        return Optional.ofNullable(this.regionId);
    }
    /**
     * @return The status of the scheduling rule.
     * 
     */
    public Optional<Integer> status() {
        return Optional.ofNullable(this.status);
    }
    /**
     * @return The address type of the interaction resource. Valid values:
     * `A`: IPv4 address.
     * `CNAME`: CNAME record.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }
    /**
     * @return The address of the interaction resource.
     * 
     */
    public Optional<String> value() {
        return Optional.ofNullable(this.value);
    }
    /**
     * @return Required. The type of the linked resource. It is an Integer. Valid values:
     * `1`: The IP address of Anti-DDoS Pro or Anti-DDoS Premium
     * `2`: the IP address of the interaction resource (in the tiered protection scenario)
     * `3`: the IP address used to accelerate access (in the network acceleration scenario)
     * `6` the IP address of the interaction resource (in the cloud service interaction scenario)
     * 
     */
    public Optional<Integer> valueType() {
        return Optional.ofNullable(this.valueType);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SchedulerRuleRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer priority;
        private @Nullable String regionId;
        private @Nullable Integer status;
        private @Nullable String type;
        private @Nullable String value;
        private @Nullable Integer valueType;
        public Builder() {}
        public Builder(SchedulerRuleRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.priority = defaults.priority;
    	      this.regionId = defaults.regionId;
    	      this.status = defaults.status;
    	      this.type = defaults.type;
    	      this.value = defaults.value;
    	      this.valueType = defaults.valueType;
        }

        @CustomType.Setter
        public Builder priority(@Nullable Integer priority) {
            this.priority = priority;
            return this;
        }
        @CustomType.Setter
        public Builder regionId(@Nullable String regionId) {
            this.regionId = regionId;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable Integer status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        @CustomType.Setter
        public Builder value(@Nullable String value) {
            this.value = value;
            return this;
        }
        @CustomType.Setter
        public Builder valueType(@Nullable Integer valueType) {
            this.valueType = valueType;
            return this;
        }
        public SchedulerRuleRule build() {
            final var o = new SchedulerRuleRule();
            o.priority = priority;
            o.regionId = regionId;
            o.status = status;
            o.type = type;
            o.value = value;
            o.valueType = valueType;
            return o;
        }
    }
}
