// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetHybridMonitorFcTasksTask {
    /**
     * @return Create the timestamp of the monitoring task. Unit: milliseconds.
     * 
     */
    private String createTime;
    /**
     * @return The ID of the monitoring task.
     * 
     */
    private String hybridMonitorFcTaskId;
    /**
     * @return The ID of the Hybrid Monitor Fc Task. The value formats as `&lt;hybrid_monitor_fc_task_id&gt;:&lt;namespace&gt;`.
     * 
     */
    private String id;
    /**
     * @return The index warehouse where the host belongs.
     * 
     */
    private String namespace;
    /**
     * @return The ID of the member account.
     * 
     */
    private String targetUserId;
    /**
     * @return The configuration file of the Alibaba Cloud service that you want to monitor by using Hybrid Cloud Monitoring.
     * 
     */
    private String yarmConfig;

    private GetHybridMonitorFcTasksTask() {}
    /**
     * @return Create the timestamp of the monitoring task. Unit: milliseconds.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The ID of the monitoring task.
     * 
     */
    public String hybridMonitorFcTaskId() {
        return this.hybridMonitorFcTaskId;
    }
    /**
     * @return The ID of the Hybrid Monitor Fc Task. The value formats as `&lt;hybrid_monitor_fc_task_id&gt;:&lt;namespace&gt;`.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The index warehouse where the host belongs.
     * 
     */
    public String namespace() {
        return this.namespace;
    }
    /**
     * @return The ID of the member account.
     * 
     */
    public String targetUserId() {
        return this.targetUserId;
    }
    /**
     * @return The configuration file of the Alibaba Cloud service that you want to monitor by using Hybrid Cloud Monitoring.
     * 
     */
    public String yarmConfig() {
        return this.yarmConfig;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetHybridMonitorFcTasksTask defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createTime;
        private String hybridMonitorFcTaskId;
        private String id;
        private String namespace;
        private String targetUserId;
        private String yarmConfig;
        public Builder() {}
        public Builder(GetHybridMonitorFcTasksTask defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createTime = defaults.createTime;
    	      this.hybridMonitorFcTaskId = defaults.hybridMonitorFcTaskId;
    	      this.id = defaults.id;
    	      this.namespace = defaults.namespace;
    	      this.targetUserId = defaults.targetUserId;
    	      this.yarmConfig = defaults.yarmConfig;
        }

        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder hybridMonitorFcTaskId(String hybridMonitorFcTaskId) {
            this.hybridMonitorFcTaskId = Objects.requireNonNull(hybridMonitorFcTaskId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder namespace(String namespace) {
            this.namespace = Objects.requireNonNull(namespace);
            return this;
        }
        @CustomType.Setter
        public Builder targetUserId(String targetUserId) {
            this.targetUserId = Objects.requireNonNull(targetUserId);
            return this;
        }
        @CustomType.Setter
        public Builder yarmConfig(String yarmConfig) {
            this.yarmConfig = Objects.requireNonNull(yarmConfig);
            return this;
        }
        public GetHybridMonitorFcTasksTask build() {
            final var o = new GetHybridMonitorFcTasksTask();
            o.createTime = createTime;
            o.hybridMonitorFcTaskId = hybridMonitorFcTaskId;
            o.id = id;
            o.namespace = namespace;
            o.targetUserId = targetUserId;
            o.yarmConfig = yarmConfig;
            return o;
        }
    }
}
