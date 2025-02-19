// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetInterRegionTrafficQosQueuesQueue {
    /**
     * @return The DSCP value of the traffic packet to be matched in the current queue, ranging from 0 to 63.
     * 
     */
    private List<String> dscps;
    private String id;
    /**
     * @return The description information of the traffic scheduling policy.
     * 
     */
    private String interRegionTrafficQosQueueDescription;
    /**
     * @return The ID of the resource.
     * 
     */
    private String interRegionTrafficQosQueueId;
    /**
     * @return The name of the traffic scheduling policy.
     * 
     */
    private String interRegionTrafficQosQueueName;
    /**
     * @return The percentage of cross-region bandwidth that the current queue can use.
     * 
     */
    private Integer remainBandwidthPercent;
    /**
     * @return The status of the traffic scheduling policy. -**Creating**: The function is being created.-**Active**: available.-**Modifying**: is being modified.-**Deleting**: Deleted.-**Deleted**: Deleted.
     * 
     */
    private String status;
    /**
     * @return The ID of the traffic scheduling policy.
     * 
     */
    private String trafficQosPolicyId;

    private GetInterRegionTrafficQosQueuesQueue() {}
    /**
     * @return The DSCP value of the traffic packet to be matched in the current queue, ranging from 0 to 63.
     * 
     */
    public List<String> dscps() {
        return this.dscps;
    }
    public String id() {
        return this.id;
    }
    /**
     * @return The description information of the traffic scheduling policy.
     * 
     */
    public String interRegionTrafficQosQueueDescription() {
        return this.interRegionTrafficQosQueueDescription;
    }
    /**
     * @return The ID of the resource.
     * 
     */
    public String interRegionTrafficQosQueueId() {
        return this.interRegionTrafficQosQueueId;
    }
    /**
     * @return The name of the traffic scheduling policy.
     * 
     */
    public String interRegionTrafficQosQueueName() {
        return this.interRegionTrafficQosQueueName;
    }
    /**
     * @return The percentage of cross-region bandwidth that the current queue can use.
     * 
     */
    public Integer remainBandwidthPercent() {
        return this.remainBandwidthPercent;
    }
    /**
     * @return The status of the traffic scheduling policy. -**Creating**: The function is being created.-**Active**: available.-**Modifying**: is being modified.-**Deleting**: Deleted.-**Deleted**: Deleted.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The ID of the traffic scheduling policy.
     * 
     */
    public String trafficQosPolicyId() {
        return this.trafficQosPolicyId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInterRegionTrafficQosQueuesQueue defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> dscps;
        private String id;
        private String interRegionTrafficQosQueueDescription;
        private String interRegionTrafficQosQueueId;
        private String interRegionTrafficQosQueueName;
        private Integer remainBandwidthPercent;
        private String status;
        private String trafficQosPolicyId;
        public Builder() {}
        public Builder(GetInterRegionTrafficQosQueuesQueue defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dscps = defaults.dscps;
    	      this.id = defaults.id;
    	      this.interRegionTrafficQosQueueDescription = defaults.interRegionTrafficQosQueueDescription;
    	      this.interRegionTrafficQosQueueId = defaults.interRegionTrafficQosQueueId;
    	      this.interRegionTrafficQosQueueName = defaults.interRegionTrafficQosQueueName;
    	      this.remainBandwidthPercent = defaults.remainBandwidthPercent;
    	      this.status = defaults.status;
    	      this.trafficQosPolicyId = defaults.trafficQosPolicyId;
        }

        @CustomType.Setter
        public Builder dscps(List<String> dscps) {
            this.dscps = Objects.requireNonNull(dscps);
            return this;
        }
        public Builder dscps(String... dscps) {
            return dscps(List.of(dscps));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder interRegionTrafficQosQueueDescription(String interRegionTrafficQosQueueDescription) {
            this.interRegionTrafficQosQueueDescription = Objects.requireNonNull(interRegionTrafficQosQueueDescription);
            return this;
        }
        @CustomType.Setter
        public Builder interRegionTrafficQosQueueId(String interRegionTrafficQosQueueId) {
            this.interRegionTrafficQosQueueId = Objects.requireNonNull(interRegionTrafficQosQueueId);
            return this;
        }
        @CustomType.Setter
        public Builder interRegionTrafficQosQueueName(String interRegionTrafficQosQueueName) {
            this.interRegionTrafficQosQueueName = Objects.requireNonNull(interRegionTrafficQosQueueName);
            return this;
        }
        @CustomType.Setter
        public Builder remainBandwidthPercent(Integer remainBandwidthPercent) {
            this.remainBandwidthPercent = Objects.requireNonNull(remainBandwidthPercent);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder trafficQosPolicyId(String trafficQosPolicyId) {
            this.trafficQosPolicyId = Objects.requireNonNull(trafficQosPolicyId);
            return this;
        }
        public GetInterRegionTrafficQosQueuesQueue build() {
            final var o = new GetInterRegionTrafficQosQueuesQueue();
            o.dscps = dscps;
            o.id = id;
            o.interRegionTrafficQosQueueDescription = interRegionTrafficQosQueueDescription;
            o.interRegionTrafficQosQueueId = interRegionTrafficQosQueueId;
            o.interRegionTrafficQosQueueName = interRegionTrafficQosQueueName;
            o.remainBandwidthPercent = remainBandwidthPercent;
            o.status = status;
            o.trafficQosPolicyId = trafficQosPolicyId;
            return o;
        }
    }
}
