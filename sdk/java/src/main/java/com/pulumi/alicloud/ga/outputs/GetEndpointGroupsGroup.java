// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.outputs;

import com.pulumi.alicloud.ga.outputs.GetEndpointGroupsGroupEndpointConfiguration;
import com.pulumi.alicloud.ga.outputs.GetEndpointGroupsGroupPortOverride;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetEndpointGroupsGroup {
    /**
     * @return The description of the endpoint group.
     * 
     */
    private String description;
    /**
     * @return The endpointConfigurations of the endpoint group.
     * 
     */
    private List<GetEndpointGroupsGroupEndpointConfiguration> endpointConfigurations;
    /**
     * @return The endpoint_group_id of the Endpoint Group.
     * 
     */
    private String endpointGroupId;
    /**
     * @return The ID of the region where the endpoint group is deployed.
     * 
     */
    private String endpointGroupRegion;
    /**
     * @return The interval between two consecutive health checks. Unit: seconds.
     * 
     */
    private Integer healthCheckIntervalSeconds;
    /**
     * @return The path specified as the destination of the targets for health checks.
     * 
     */
    private String healthCheckPath;
    /**
     * @return The port that is used for health checks.
     * 
     */
    private Integer healthCheckPort;
    /**
     * @return The protocol that is used to connect to the targets for health checks.
     * 
     */
    private String healthCheckProtocol;
    /**
     * @return The ID of the Endpoint Group.
     * 
     */
    private String id;
    /**
     * @return The ID of the listener that is associated with the endpoint group.
     * 
     */
    private String listenerId;
    /**
     * @return The name of the endpoint group.
     * 
     */
    private String name;
    /**
     * @return Mapping between listening port and forwarding port of boarding point.
     * 
     */
    private List<GetEndpointGroupsGroupPortOverride> portOverrides;
    /**
     * @return The status of the endpoint group.
     * 
     */
    private String status;
    /**
     * @return The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy.
     * 
     */
    private Integer thresholdCount;
    /**
     * @return The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
     * 
     */
    private Integer trafficPercentage;

    private GetEndpointGroupsGroup() {}
    /**
     * @return The description of the endpoint group.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The endpointConfigurations of the endpoint group.
     * 
     */
    public List<GetEndpointGroupsGroupEndpointConfiguration> endpointConfigurations() {
        return this.endpointConfigurations;
    }
    /**
     * @return The endpoint_group_id of the Endpoint Group.
     * 
     */
    public String endpointGroupId() {
        return this.endpointGroupId;
    }
    /**
     * @return The ID of the region where the endpoint group is deployed.
     * 
     */
    public String endpointGroupRegion() {
        return this.endpointGroupRegion;
    }
    /**
     * @return The interval between two consecutive health checks. Unit: seconds.
     * 
     */
    public Integer healthCheckIntervalSeconds() {
        return this.healthCheckIntervalSeconds;
    }
    /**
     * @return The path specified as the destination of the targets for health checks.
     * 
     */
    public String healthCheckPath() {
        return this.healthCheckPath;
    }
    /**
     * @return The port that is used for health checks.
     * 
     */
    public Integer healthCheckPort() {
        return this.healthCheckPort;
    }
    /**
     * @return The protocol that is used to connect to the targets for health checks.
     * 
     */
    public String healthCheckProtocol() {
        return this.healthCheckProtocol;
    }
    /**
     * @return The ID of the Endpoint Group.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The ID of the listener that is associated with the endpoint group.
     * 
     */
    public String listenerId() {
        return this.listenerId;
    }
    /**
     * @return The name of the endpoint group.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Mapping between listening port and forwarding port of boarding point.
     * 
     */
    public List<GetEndpointGroupsGroupPortOverride> portOverrides() {
        return this.portOverrides;
    }
    /**
     * @return The status of the endpoint group.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The number of consecutive failed heath checks that must occur before the endpoint is deemed unhealthy.
     * 
     */
    public Integer thresholdCount() {
        return this.thresholdCount;
    }
    /**
     * @return The weight of the endpoint group when the corresponding listener is associated with multiple endpoint groups.
     * 
     */
    public Integer trafficPercentage() {
        return this.trafficPercentage;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEndpointGroupsGroup defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private List<GetEndpointGroupsGroupEndpointConfiguration> endpointConfigurations;
        private String endpointGroupId;
        private String endpointGroupRegion;
        private Integer healthCheckIntervalSeconds;
        private String healthCheckPath;
        private Integer healthCheckPort;
        private String healthCheckProtocol;
        private String id;
        private String listenerId;
        private String name;
        private List<GetEndpointGroupsGroupPortOverride> portOverrides;
        private String status;
        private Integer thresholdCount;
        private Integer trafficPercentage;
        public Builder() {}
        public Builder(GetEndpointGroupsGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.endpointConfigurations = defaults.endpointConfigurations;
    	      this.endpointGroupId = defaults.endpointGroupId;
    	      this.endpointGroupRegion = defaults.endpointGroupRegion;
    	      this.healthCheckIntervalSeconds = defaults.healthCheckIntervalSeconds;
    	      this.healthCheckPath = defaults.healthCheckPath;
    	      this.healthCheckPort = defaults.healthCheckPort;
    	      this.healthCheckProtocol = defaults.healthCheckProtocol;
    	      this.id = defaults.id;
    	      this.listenerId = defaults.listenerId;
    	      this.name = defaults.name;
    	      this.portOverrides = defaults.portOverrides;
    	      this.status = defaults.status;
    	      this.thresholdCount = defaults.thresholdCount;
    	      this.trafficPercentage = defaults.trafficPercentage;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder endpointConfigurations(List<GetEndpointGroupsGroupEndpointConfiguration> endpointConfigurations) {
            this.endpointConfigurations = Objects.requireNonNull(endpointConfigurations);
            return this;
        }
        public Builder endpointConfigurations(GetEndpointGroupsGroupEndpointConfiguration... endpointConfigurations) {
            return endpointConfigurations(List.of(endpointConfigurations));
        }
        @CustomType.Setter
        public Builder endpointGroupId(String endpointGroupId) {
            this.endpointGroupId = Objects.requireNonNull(endpointGroupId);
            return this;
        }
        @CustomType.Setter
        public Builder endpointGroupRegion(String endpointGroupRegion) {
            this.endpointGroupRegion = Objects.requireNonNull(endpointGroupRegion);
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckIntervalSeconds(Integer healthCheckIntervalSeconds) {
            this.healthCheckIntervalSeconds = Objects.requireNonNull(healthCheckIntervalSeconds);
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckPath(String healthCheckPath) {
            this.healthCheckPath = Objects.requireNonNull(healthCheckPath);
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckPort(Integer healthCheckPort) {
            this.healthCheckPort = Objects.requireNonNull(healthCheckPort);
            return this;
        }
        @CustomType.Setter
        public Builder healthCheckProtocol(String healthCheckProtocol) {
            this.healthCheckProtocol = Objects.requireNonNull(healthCheckProtocol);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder listenerId(String listenerId) {
            this.listenerId = Objects.requireNonNull(listenerId);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder portOverrides(List<GetEndpointGroupsGroupPortOverride> portOverrides) {
            this.portOverrides = Objects.requireNonNull(portOverrides);
            return this;
        }
        public Builder portOverrides(GetEndpointGroupsGroupPortOverride... portOverrides) {
            return portOverrides(List.of(portOverrides));
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder thresholdCount(Integer thresholdCount) {
            this.thresholdCount = Objects.requireNonNull(thresholdCount);
            return this;
        }
        @CustomType.Setter
        public Builder trafficPercentage(Integer trafficPercentage) {
            this.trafficPercentage = Objects.requireNonNull(trafficPercentage);
            return this;
        }
        public GetEndpointGroupsGroup build() {
            final var o = new GetEndpointGroupsGroup();
            o.description = description;
            o.endpointConfigurations = endpointConfigurations;
            o.endpointGroupId = endpointGroupId;
            o.endpointGroupRegion = endpointGroupRegion;
            o.healthCheckIntervalSeconds = healthCheckIntervalSeconds;
            o.healthCheckPath = healthCheckPath;
            o.healthCheckPort = healthCheckPort;
            o.healthCheckProtocol = healthCheckProtocol;
            o.id = id;
            o.listenerId = listenerId;
            o.name = name;
            o.portOverrides = portOverrides;
            o.status = status;
            o.thresholdCount = thresholdCount;
            o.trafficPercentage = trafficPercentage;
            return o;
        }
    }
}
