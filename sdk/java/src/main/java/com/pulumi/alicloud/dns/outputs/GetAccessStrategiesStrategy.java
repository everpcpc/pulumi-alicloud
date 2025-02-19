// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns.outputs;

import com.pulumi.alicloud.dns.outputs.GetAccessStrategiesStrategyDefaultAddrPool;
import com.pulumi.alicloud.dns.outputs.GetAccessStrategiesStrategyFailoverAddrPool;
import com.pulumi.alicloud.dns.outputs.GetAccessStrategiesStrategyLine;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetAccessStrategiesStrategy {
    /**
     * @return The primary/secondary switchover policy for address pool groups.
     * 
     */
    private String accessMode;
    /**
     * @return The first ID of the resource.
     * 
     */
    private String accessStrategyId;
    /**
     * @return The time when the access policy was created.
     * 
     */
    private String createTime;
    /**
     * @return The timestamp that indicates when the access policy was created.
     * 
     */
    private String createTimestamp;
    /**
     * @return The type of the primary address pool.
     * 
     */
    private String defaultAddrPoolType;
    /**
     * @return The address pools in the primary address pool group.
     * 
     */
    private List<GetAccessStrategiesStrategyDefaultAddrPool> defaultAddrPools;
    /**
     * @return The number of addresses currently available in the primary address pool.
     * 
     */
    private Integer defaultAvailableAddrNum;
    /**
     * @return Indicates whether scheduling optimization for latency resolution was enabled for the primary address pool group.
     * 
     */
    private String defaultLatencyOptimization;
    /**
     * @return The load balancing policy of the primary address pool group.
     * 
     */
    private String defaultLbaStrategy;
    /**
     * @return The maximum number of addresses returned by the primary address pool set.
     * 
     */
    private Integer defaultMaxReturnAddrNum;
    /**
     * @return The minimum number of available addresses for the primary address pool set.
     * 
     */
    private Integer defaultMinAvailableAddrNum;
    /**
     * @return The type of the active address pool group.
     * 
     */
    private String effectiveAddrPoolGroupType;
    /**
     * @return The type of the secondary address pool.
     * 
     */
    private String failoverAddrPoolType;
    /**
     * @return The address pools in the secondary address pool group.
     * 
     */
    private List<GetAccessStrategiesStrategyFailoverAddrPool> failoverAddrPools;
    /**
     * @return The number of available addresses in the standby address pool.
     * 
     */
    private Integer failoverAvailableAddrNum;
    /**
     * @return Indicates whether scheduling optimization for latency resolution was enabled for the secondary address pool group.
     * 
     */
    private String failoverLatencyOptimization;
    /**
     * @return The load balancing policy of the secondary address pool group.
     * 
     */
    private String failoverLbaStrategy;
    /**
     * @return The maximum number of returned addresses in the standby address pool.
     * 
     */
    private Integer failoverMaxReturnAddrNum;
    /**
     * @return The minimum number of available addresses in the standby address pool.
     * 
     */
    private Integer failoverMinAvailableAddrNum;
    /**
     * @return The ID of the Access Strategy.
     * 
     */
    private String id;
    /**
     * @return The Id of the associated instance.
     * 
     */
    private String instanceId;
    /**
     * @return List of source regions.
     * 
     */
    private List<GetAccessStrategiesStrategyLine> lines;
    /**
     * @return The type of the access policy.
     * 
     */
    private String strategyMode;
    /**
     * @return The name of the access policy.
     * 
     */
    private String strategyName;

    private GetAccessStrategiesStrategy() {}
    /**
     * @return The primary/secondary switchover policy for address pool groups.
     * 
     */
    public String accessMode() {
        return this.accessMode;
    }
    /**
     * @return The first ID of the resource.
     * 
     */
    public String accessStrategyId() {
        return this.accessStrategyId;
    }
    /**
     * @return The time when the access policy was created.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The timestamp that indicates when the access policy was created.
     * 
     */
    public String createTimestamp() {
        return this.createTimestamp;
    }
    /**
     * @return The type of the primary address pool.
     * 
     */
    public String defaultAddrPoolType() {
        return this.defaultAddrPoolType;
    }
    /**
     * @return The address pools in the primary address pool group.
     * 
     */
    public List<GetAccessStrategiesStrategyDefaultAddrPool> defaultAddrPools() {
        return this.defaultAddrPools;
    }
    /**
     * @return The number of addresses currently available in the primary address pool.
     * 
     */
    public Integer defaultAvailableAddrNum() {
        return this.defaultAvailableAddrNum;
    }
    /**
     * @return Indicates whether scheduling optimization for latency resolution was enabled for the primary address pool group.
     * 
     */
    public String defaultLatencyOptimization() {
        return this.defaultLatencyOptimization;
    }
    /**
     * @return The load balancing policy of the primary address pool group.
     * 
     */
    public String defaultLbaStrategy() {
        return this.defaultLbaStrategy;
    }
    /**
     * @return The maximum number of addresses returned by the primary address pool set.
     * 
     */
    public Integer defaultMaxReturnAddrNum() {
        return this.defaultMaxReturnAddrNum;
    }
    /**
     * @return The minimum number of available addresses for the primary address pool set.
     * 
     */
    public Integer defaultMinAvailableAddrNum() {
        return this.defaultMinAvailableAddrNum;
    }
    /**
     * @return The type of the active address pool group.
     * 
     */
    public String effectiveAddrPoolGroupType() {
        return this.effectiveAddrPoolGroupType;
    }
    /**
     * @return The type of the secondary address pool.
     * 
     */
    public String failoverAddrPoolType() {
        return this.failoverAddrPoolType;
    }
    /**
     * @return The address pools in the secondary address pool group.
     * 
     */
    public List<GetAccessStrategiesStrategyFailoverAddrPool> failoverAddrPools() {
        return this.failoverAddrPools;
    }
    /**
     * @return The number of available addresses in the standby address pool.
     * 
     */
    public Integer failoverAvailableAddrNum() {
        return this.failoverAvailableAddrNum;
    }
    /**
     * @return Indicates whether scheduling optimization for latency resolution was enabled for the secondary address pool group.
     * 
     */
    public String failoverLatencyOptimization() {
        return this.failoverLatencyOptimization;
    }
    /**
     * @return The load balancing policy of the secondary address pool group.
     * 
     */
    public String failoverLbaStrategy() {
        return this.failoverLbaStrategy;
    }
    /**
     * @return The maximum number of returned addresses in the standby address pool.
     * 
     */
    public Integer failoverMaxReturnAddrNum() {
        return this.failoverMaxReturnAddrNum;
    }
    /**
     * @return The minimum number of available addresses in the standby address pool.
     * 
     */
    public Integer failoverMinAvailableAddrNum() {
        return this.failoverMinAvailableAddrNum;
    }
    /**
     * @return The ID of the Access Strategy.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The Id of the associated instance.
     * 
     */
    public String instanceId() {
        return this.instanceId;
    }
    /**
     * @return List of source regions.
     * 
     */
    public List<GetAccessStrategiesStrategyLine> lines() {
        return this.lines;
    }
    /**
     * @return The type of the access policy.
     * 
     */
    public String strategyMode() {
        return this.strategyMode;
    }
    /**
     * @return The name of the access policy.
     * 
     */
    public String strategyName() {
        return this.strategyName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAccessStrategiesStrategy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accessMode;
        private String accessStrategyId;
        private String createTime;
        private String createTimestamp;
        private String defaultAddrPoolType;
        private List<GetAccessStrategiesStrategyDefaultAddrPool> defaultAddrPools;
        private Integer defaultAvailableAddrNum;
        private String defaultLatencyOptimization;
        private String defaultLbaStrategy;
        private Integer defaultMaxReturnAddrNum;
        private Integer defaultMinAvailableAddrNum;
        private String effectiveAddrPoolGroupType;
        private String failoverAddrPoolType;
        private List<GetAccessStrategiesStrategyFailoverAddrPool> failoverAddrPools;
        private Integer failoverAvailableAddrNum;
        private String failoverLatencyOptimization;
        private String failoverLbaStrategy;
        private Integer failoverMaxReturnAddrNum;
        private Integer failoverMinAvailableAddrNum;
        private String id;
        private String instanceId;
        private List<GetAccessStrategiesStrategyLine> lines;
        private String strategyMode;
        private String strategyName;
        public Builder() {}
        public Builder(GetAccessStrategiesStrategy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessMode = defaults.accessMode;
    	      this.accessStrategyId = defaults.accessStrategyId;
    	      this.createTime = defaults.createTime;
    	      this.createTimestamp = defaults.createTimestamp;
    	      this.defaultAddrPoolType = defaults.defaultAddrPoolType;
    	      this.defaultAddrPools = defaults.defaultAddrPools;
    	      this.defaultAvailableAddrNum = defaults.defaultAvailableAddrNum;
    	      this.defaultLatencyOptimization = defaults.defaultLatencyOptimization;
    	      this.defaultLbaStrategy = defaults.defaultLbaStrategy;
    	      this.defaultMaxReturnAddrNum = defaults.defaultMaxReturnAddrNum;
    	      this.defaultMinAvailableAddrNum = defaults.defaultMinAvailableAddrNum;
    	      this.effectiveAddrPoolGroupType = defaults.effectiveAddrPoolGroupType;
    	      this.failoverAddrPoolType = defaults.failoverAddrPoolType;
    	      this.failoverAddrPools = defaults.failoverAddrPools;
    	      this.failoverAvailableAddrNum = defaults.failoverAvailableAddrNum;
    	      this.failoverLatencyOptimization = defaults.failoverLatencyOptimization;
    	      this.failoverLbaStrategy = defaults.failoverLbaStrategy;
    	      this.failoverMaxReturnAddrNum = defaults.failoverMaxReturnAddrNum;
    	      this.failoverMinAvailableAddrNum = defaults.failoverMinAvailableAddrNum;
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.lines = defaults.lines;
    	      this.strategyMode = defaults.strategyMode;
    	      this.strategyName = defaults.strategyName;
        }

        @CustomType.Setter
        public Builder accessMode(String accessMode) {
            this.accessMode = Objects.requireNonNull(accessMode);
            return this;
        }
        @CustomType.Setter
        public Builder accessStrategyId(String accessStrategyId) {
            this.accessStrategyId = Objects.requireNonNull(accessStrategyId);
            return this;
        }
        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder createTimestamp(String createTimestamp) {
            this.createTimestamp = Objects.requireNonNull(createTimestamp);
            return this;
        }
        @CustomType.Setter
        public Builder defaultAddrPoolType(String defaultAddrPoolType) {
            this.defaultAddrPoolType = Objects.requireNonNull(defaultAddrPoolType);
            return this;
        }
        @CustomType.Setter
        public Builder defaultAddrPools(List<GetAccessStrategiesStrategyDefaultAddrPool> defaultAddrPools) {
            this.defaultAddrPools = Objects.requireNonNull(defaultAddrPools);
            return this;
        }
        public Builder defaultAddrPools(GetAccessStrategiesStrategyDefaultAddrPool... defaultAddrPools) {
            return defaultAddrPools(List.of(defaultAddrPools));
        }
        @CustomType.Setter
        public Builder defaultAvailableAddrNum(Integer defaultAvailableAddrNum) {
            this.defaultAvailableAddrNum = Objects.requireNonNull(defaultAvailableAddrNum);
            return this;
        }
        @CustomType.Setter
        public Builder defaultLatencyOptimization(String defaultLatencyOptimization) {
            this.defaultLatencyOptimization = Objects.requireNonNull(defaultLatencyOptimization);
            return this;
        }
        @CustomType.Setter
        public Builder defaultLbaStrategy(String defaultLbaStrategy) {
            this.defaultLbaStrategy = Objects.requireNonNull(defaultLbaStrategy);
            return this;
        }
        @CustomType.Setter
        public Builder defaultMaxReturnAddrNum(Integer defaultMaxReturnAddrNum) {
            this.defaultMaxReturnAddrNum = Objects.requireNonNull(defaultMaxReturnAddrNum);
            return this;
        }
        @CustomType.Setter
        public Builder defaultMinAvailableAddrNum(Integer defaultMinAvailableAddrNum) {
            this.defaultMinAvailableAddrNum = Objects.requireNonNull(defaultMinAvailableAddrNum);
            return this;
        }
        @CustomType.Setter
        public Builder effectiveAddrPoolGroupType(String effectiveAddrPoolGroupType) {
            this.effectiveAddrPoolGroupType = Objects.requireNonNull(effectiveAddrPoolGroupType);
            return this;
        }
        @CustomType.Setter
        public Builder failoverAddrPoolType(String failoverAddrPoolType) {
            this.failoverAddrPoolType = Objects.requireNonNull(failoverAddrPoolType);
            return this;
        }
        @CustomType.Setter
        public Builder failoverAddrPools(List<GetAccessStrategiesStrategyFailoverAddrPool> failoverAddrPools) {
            this.failoverAddrPools = Objects.requireNonNull(failoverAddrPools);
            return this;
        }
        public Builder failoverAddrPools(GetAccessStrategiesStrategyFailoverAddrPool... failoverAddrPools) {
            return failoverAddrPools(List.of(failoverAddrPools));
        }
        @CustomType.Setter
        public Builder failoverAvailableAddrNum(Integer failoverAvailableAddrNum) {
            this.failoverAvailableAddrNum = Objects.requireNonNull(failoverAvailableAddrNum);
            return this;
        }
        @CustomType.Setter
        public Builder failoverLatencyOptimization(String failoverLatencyOptimization) {
            this.failoverLatencyOptimization = Objects.requireNonNull(failoverLatencyOptimization);
            return this;
        }
        @CustomType.Setter
        public Builder failoverLbaStrategy(String failoverLbaStrategy) {
            this.failoverLbaStrategy = Objects.requireNonNull(failoverLbaStrategy);
            return this;
        }
        @CustomType.Setter
        public Builder failoverMaxReturnAddrNum(Integer failoverMaxReturnAddrNum) {
            this.failoverMaxReturnAddrNum = Objects.requireNonNull(failoverMaxReturnAddrNum);
            return this;
        }
        @CustomType.Setter
        public Builder failoverMinAvailableAddrNum(Integer failoverMinAvailableAddrNum) {
            this.failoverMinAvailableAddrNum = Objects.requireNonNull(failoverMinAvailableAddrNum);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(String instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder lines(List<GetAccessStrategiesStrategyLine> lines) {
            this.lines = Objects.requireNonNull(lines);
            return this;
        }
        public Builder lines(GetAccessStrategiesStrategyLine... lines) {
            return lines(List.of(lines));
        }
        @CustomType.Setter
        public Builder strategyMode(String strategyMode) {
            this.strategyMode = Objects.requireNonNull(strategyMode);
            return this;
        }
        @CustomType.Setter
        public Builder strategyName(String strategyName) {
            this.strategyName = Objects.requireNonNull(strategyName);
            return this;
        }
        public GetAccessStrategiesStrategy build() {
            final var o = new GetAccessStrategiesStrategy();
            o.accessMode = accessMode;
            o.accessStrategyId = accessStrategyId;
            o.createTime = createTime;
            o.createTimestamp = createTimestamp;
            o.defaultAddrPoolType = defaultAddrPoolType;
            o.defaultAddrPools = defaultAddrPools;
            o.defaultAvailableAddrNum = defaultAvailableAddrNum;
            o.defaultLatencyOptimization = defaultLatencyOptimization;
            o.defaultLbaStrategy = defaultLbaStrategy;
            o.defaultMaxReturnAddrNum = defaultMaxReturnAddrNum;
            o.defaultMinAvailableAddrNum = defaultMinAvailableAddrNum;
            o.effectiveAddrPoolGroupType = effectiveAddrPoolGroupType;
            o.failoverAddrPoolType = failoverAddrPoolType;
            o.failoverAddrPools = failoverAddrPools;
            o.failoverAvailableAddrNum = failoverAvailableAddrNum;
            o.failoverLatencyOptimization = failoverLatencyOptimization;
            o.failoverLbaStrategy = failoverLbaStrategy;
            o.failoverMaxReturnAddrNum = failoverMaxReturnAddrNum;
            o.failoverMinAvailableAddrNum = failoverMinAvailableAddrNum;
            o.id = id;
            o.instanceId = instanceId;
            o.lines = lines;
            o.strategyMode = strategyMode;
            o.strategyName = strategyName;
            return o;
        }
    }
}
