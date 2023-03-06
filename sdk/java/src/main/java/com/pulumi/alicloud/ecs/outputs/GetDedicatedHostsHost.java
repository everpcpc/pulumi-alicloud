// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.alicloud.ecs.outputs.GetDedicatedHostsHostCapacity;
import com.pulumi.alicloud.ecs.outputs.GetDedicatedHostsHostNetworkAttribute;
import com.pulumi.alicloud.ecs.outputs.GetDedicatedHostsHostOperationLock;
import com.pulumi.core.annotations.CustomType;
import java.lang.Double;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetDedicatedHostsHost {
    /**
     * @return The policy used to migrate the instances from the dedicated host when the dedicated host fails or needs to be repaired online.
     * 
     */
    private String actionOnMaintenance;
    /**
     * @return Specifies whether to add the dedicated host to the resource pool for automatic deployment.
     * 
     */
    private String autoPlacement;
    /**
     * @return The automatic release time of the dedicated host.
     * 
     */
    private String autoReleaseTime;
    /**
     * @return (Available in 1.123.1+) A collection of proprietary host performance indicators.
     * 
     */
    private List<GetDedicatedHostsHostCapacity> capacities;
    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    private Integer cores;
    /**
     * @return (Available in 1.123.1+) CPU oversold ratio.
     * 
     */
    private Double cpuOverCommitRatio;
    /**
     * @return The ID of ECS Dedicated Host.
     * 
     */
    private String dedicatedHostId;
    /**
     * @return The name of ECS Dedicated Host.
     * 
     */
    private String dedicatedHostName;
    /**
     * @return The type of the dedicated host.
     * 
     */
    private String dedicatedHostType;
    /**
     * @return The description of the dedicated host.
     * 
     */
    private String description;
    /**
     * @return The expiration time of the subscription dedicated host.
     * 
     */
    private String expiredTime;
    /**
     * @return The GPU model.
     * 
     */
    private String gpuSpec;
    /**
     * @return ID of the ECS Dedicated Host.
     * 
     */
    private String id;
    /**
     * @return The machine code of the dedicated host.
     * 
     */
    private String machineId;
    /**
     * @return dedicated host network parameters. contains the following attributes:
     * 
     */
    private List<GetDedicatedHostsHostNetworkAttribute> networkAttributes;
    /**
     * @return The reason why the dedicated host resource is locked.
     * 
     */
    private List<GetDedicatedHostsHostOperationLock> operationLocks;
    /**
     * @return The billing method of the dedicated host.
     * 
     */
    private String paymentType;
    /**
     * @return The number of physical GPUs.
     * 
     */
    private Integer physicalGpus;
    /**
     * @return The ID of the resource group to which the ECS Dedicated Host belongs.
     * 
     */
    private String resourceGroupId;
    /**
     * @return The unit of the subscription billing method.
     * 
     */
    private String saleCycle;
    /**
     * @return The number of physical CPUs.
     * 
     */
    private Integer sockets;
    /**
     * @return The status of the ECS Dedicated Host. validate value: `Available`, `Creating`, `PermanentFailure`, `Released`, `UnderAssessment`.
     * 
     */
    private String status;
    /**
     * @return (Available in 1.123.1+) A custom instance type family supported by a dedicated host.
     * 
     */
    private List<String> supportedCustomInstanceTypeFamilies;
    /**
     * @return (Available in 1.123.1+) ECS instance type family supported by the dedicated host.
     * 
     */
    private List<String> supportedInstanceTypeFamilies;
    /**
     * @return The list of ECS instance
     * 
     */
    private List<String> supportedInstanceTypesLists;
    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    private Map<String,Object> tags;
    /**
     * @return The zone ID of the ECS Dedicated Host.
     * 
     */
    private String zoneId;

    private GetDedicatedHostsHost() {}
    /**
     * @return The policy used to migrate the instances from the dedicated host when the dedicated host fails or needs to be repaired online.
     * 
     */
    public String actionOnMaintenance() {
        return this.actionOnMaintenance;
    }
    /**
     * @return Specifies whether to add the dedicated host to the resource pool for automatic deployment.
     * 
     */
    public String autoPlacement() {
        return this.autoPlacement;
    }
    /**
     * @return The automatic release time of the dedicated host.
     * 
     */
    public String autoReleaseTime() {
        return this.autoReleaseTime;
    }
    /**
     * @return (Available in 1.123.1+) A collection of proprietary host performance indicators.
     * 
     */
    public List<GetDedicatedHostsHostCapacity> capacities() {
        return this.capacities;
    }
    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Integer cores() {
        return this.cores;
    }
    /**
     * @return (Available in 1.123.1+) CPU oversold ratio.
     * 
     */
    public Double cpuOverCommitRatio() {
        return this.cpuOverCommitRatio;
    }
    /**
     * @return The ID of ECS Dedicated Host.
     * 
     */
    public String dedicatedHostId() {
        return this.dedicatedHostId;
    }
    /**
     * @return The name of ECS Dedicated Host.
     * 
     */
    public String dedicatedHostName() {
        return this.dedicatedHostName;
    }
    /**
     * @return The type of the dedicated host.
     * 
     */
    public String dedicatedHostType() {
        return this.dedicatedHostType;
    }
    /**
     * @return The description of the dedicated host.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The expiration time of the subscription dedicated host.
     * 
     */
    public String expiredTime() {
        return this.expiredTime;
    }
    /**
     * @return The GPU model.
     * 
     */
    public String gpuSpec() {
        return this.gpuSpec;
    }
    /**
     * @return ID of the ECS Dedicated Host.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The machine code of the dedicated host.
     * 
     */
    public String machineId() {
        return this.machineId;
    }
    /**
     * @return dedicated host network parameters. contains the following attributes:
     * 
     */
    public List<GetDedicatedHostsHostNetworkAttribute> networkAttributes() {
        return this.networkAttributes;
    }
    /**
     * @return The reason why the dedicated host resource is locked.
     * 
     */
    public List<GetDedicatedHostsHostOperationLock> operationLocks() {
        return this.operationLocks;
    }
    /**
     * @return The billing method of the dedicated host.
     * 
     */
    public String paymentType() {
        return this.paymentType;
    }
    /**
     * @return The number of physical GPUs.
     * 
     */
    public Integer physicalGpus() {
        return this.physicalGpus;
    }
    /**
     * @return The ID of the resource group to which the ECS Dedicated Host belongs.
     * 
     */
    public String resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * @return The unit of the subscription billing method.
     * 
     */
    public String saleCycle() {
        return this.saleCycle;
    }
    /**
     * @return The number of physical CPUs.
     * 
     */
    public Integer sockets() {
        return this.sockets;
    }
    /**
     * @return The status of the ECS Dedicated Host. validate value: `Available`, `Creating`, `PermanentFailure`, `Released`, `UnderAssessment`.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return (Available in 1.123.1+) A custom instance type family supported by a dedicated host.
     * 
     */
    public List<String> supportedCustomInstanceTypeFamilies() {
        return this.supportedCustomInstanceTypeFamilies;
    }
    /**
     * @return (Available in 1.123.1+) ECS instance type family supported by the dedicated host.
     * 
     */
    public List<String> supportedInstanceTypeFamilies() {
        return this.supportedInstanceTypeFamilies;
    }
    /**
     * @return The list of ECS instance
     * 
     */
    public List<String> supportedInstanceTypesLists() {
        return this.supportedInstanceTypesLists;
    }
    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Map<String,Object> tags() {
        return this.tags;
    }
    /**
     * @return The zone ID of the ECS Dedicated Host.
     * 
     */
    public String zoneId() {
        return this.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDedicatedHostsHost defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String actionOnMaintenance;
        private String autoPlacement;
        private String autoReleaseTime;
        private List<GetDedicatedHostsHostCapacity> capacities;
        private Integer cores;
        private Double cpuOverCommitRatio;
        private String dedicatedHostId;
        private String dedicatedHostName;
        private String dedicatedHostType;
        private String description;
        private String expiredTime;
        private String gpuSpec;
        private String id;
        private String machineId;
        private List<GetDedicatedHostsHostNetworkAttribute> networkAttributes;
        private List<GetDedicatedHostsHostOperationLock> operationLocks;
        private String paymentType;
        private Integer physicalGpus;
        private String resourceGroupId;
        private String saleCycle;
        private Integer sockets;
        private String status;
        private List<String> supportedCustomInstanceTypeFamilies;
        private List<String> supportedInstanceTypeFamilies;
        private List<String> supportedInstanceTypesLists;
        private Map<String,Object> tags;
        private String zoneId;
        public Builder() {}
        public Builder(GetDedicatedHostsHost defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.actionOnMaintenance = defaults.actionOnMaintenance;
    	      this.autoPlacement = defaults.autoPlacement;
    	      this.autoReleaseTime = defaults.autoReleaseTime;
    	      this.capacities = defaults.capacities;
    	      this.cores = defaults.cores;
    	      this.cpuOverCommitRatio = defaults.cpuOverCommitRatio;
    	      this.dedicatedHostId = defaults.dedicatedHostId;
    	      this.dedicatedHostName = defaults.dedicatedHostName;
    	      this.dedicatedHostType = defaults.dedicatedHostType;
    	      this.description = defaults.description;
    	      this.expiredTime = defaults.expiredTime;
    	      this.gpuSpec = defaults.gpuSpec;
    	      this.id = defaults.id;
    	      this.machineId = defaults.machineId;
    	      this.networkAttributes = defaults.networkAttributes;
    	      this.operationLocks = defaults.operationLocks;
    	      this.paymentType = defaults.paymentType;
    	      this.physicalGpus = defaults.physicalGpus;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.saleCycle = defaults.saleCycle;
    	      this.sockets = defaults.sockets;
    	      this.status = defaults.status;
    	      this.supportedCustomInstanceTypeFamilies = defaults.supportedCustomInstanceTypeFamilies;
    	      this.supportedInstanceTypeFamilies = defaults.supportedInstanceTypeFamilies;
    	      this.supportedInstanceTypesLists = defaults.supportedInstanceTypesLists;
    	      this.tags = defaults.tags;
    	      this.zoneId = defaults.zoneId;
        }

        @CustomType.Setter
        public Builder actionOnMaintenance(String actionOnMaintenance) {
            this.actionOnMaintenance = Objects.requireNonNull(actionOnMaintenance);
            return this;
        }
        @CustomType.Setter
        public Builder autoPlacement(String autoPlacement) {
            this.autoPlacement = Objects.requireNonNull(autoPlacement);
            return this;
        }
        @CustomType.Setter
        public Builder autoReleaseTime(String autoReleaseTime) {
            this.autoReleaseTime = Objects.requireNonNull(autoReleaseTime);
            return this;
        }
        @CustomType.Setter
        public Builder capacities(List<GetDedicatedHostsHostCapacity> capacities) {
            this.capacities = Objects.requireNonNull(capacities);
            return this;
        }
        public Builder capacities(GetDedicatedHostsHostCapacity... capacities) {
            return capacities(List.of(capacities));
        }
        @CustomType.Setter
        public Builder cores(Integer cores) {
            this.cores = Objects.requireNonNull(cores);
            return this;
        }
        @CustomType.Setter
        public Builder cpuOverCommitRatio(Double cpuOverCommitRatio) {
            this.cpuOverCommitRatio = Objects.requireNonNull(cpuOverCommitRatio);
            return this;
        }
        @CustomType.Setter
        public Builder dedicatedHostId(String dedicatedHostId) {
            this.dedicatedHostId = Objects.requireNonNull(dedicatedHostId);
            return this;
        }
        @CustomType.Setter
        public Builder dedicatedHostName(String dedicatedHostName) {
            this.dedicatedHostName = Objects.requireNonNull(dedicatedHostName);
            return this;
        }
        @CustomType.Setter
        public Builder dedicatedHostType(String dedicatedHostType) {
            this.dedicatedHostType = Objects.requireNonNull(dedicatedHostType);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder expiredTime(String expiredTime) {
            this.expiredTime = Objects.requireNonNull(expiredTime);
            return this;
        }
        @CustomType.Setter
        public Builder gpuSpec(String gpuSpec) {
            this.gpuSpec = Objects.requireNonNull(gpuSpec);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder machineId(String machineId) {
            this.machineId = Objects.requireNonNull(machineId);
            return this;
        }
        @CustomType.Setter
        public Builder networkAttributes(List<GetDedicatedHostsHostNetworkAttribute> networkAttributes) {
            this.networkAttributes = Objects.requireNonNull(networkAttributes);
            return this;
        }
        public Builder networkAttributes(GetDedicatedHostsHostNetworkAttribute... networkAttributes) {
            return networkAttributes(List.of(networkAttributes));
        }
        @CustomType.Setter
        public Builder operationLocks(List<GetDedicatedHostsHostOperationLock> operationLocks) {
            this.operationLocks = Objects.requireNonNull(operationLocks);
            return this;
        }
        public Builder operationLocks(GetDedicatedHostsHostOperationLock... operationLocks) {
            return operationLocks(List.of(operationLocks));
        }
        @CustomType.Setter
        public Builder paymentType(String paymentType) {
            this.paymentType = Objects.requireNonNull(paymentType);
            return this;
        }
        @CustomType.Setter
        public Builder physicalGpus(Integer physicalGpus) {
            this.physicalGpus = Objects.requireNonNull(physicalGpus);
            return this;
        }
        @CustomType.Setter
        public Builder resourceGroupId(String resourceGroupId) {
            this.resourceGroupId = Objects.requireNonNull(resourceGroupId);
            return this;
        }
        @CustomType.Setter
        public Builder saleCycle(String saleCycle) {
            this.saleCycle = Objects.requireNonNull(saleCycle);
            return this;
        }
        @CustomType.Setter
        public Builder sockets(Integer sockets) {
            this.sockets = Objects.requireNonNull(sockets);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder supportedCustomInstanceTypeFamilies(List<String> supportedCustomInstanceTypeFamilies) {
            this.supportedCustomInstanceTypeFamilies = Objects.requireNonNull(supportedCustomInstanceTypeFamilies);
            return this;
        }
        public Builder supportedCustomInstanceTypeFamilies(String... supportedCustomInstanceTypeFamilies) {
            return supportedCustomInstanceTypeFamilies(List.of(supportedCustomInstanceTypeFamilies));
        }
        @CustomType.Setter
        public Builder supportedInstanceTypeFamilies(List<String> supportedInstanceTypeFamilies) {
            this.supportedInstanceTypeFamilies = Objects.requireNonNull(supportedInstanceTypeFamilies);
            return this;
        }
        public Builder supportedInstanceTypeFamilies(String... supportedInstanceTypeFamilies) {
            return supportedInstanceTypeFamilies(List.of(supportedInstanceTypeFamilies));
        }
        @CustomType.Setter
        public Builder supportedInstanceTypesLists(List<String> supportedInstanceTypesLists) {
            this.supportedInstanceTypesLists = Objects.requireNonNull(supportedInstanceTypesLists);
            return this;
        }
        public Builder supportedInstanceTypesLists(String... supportedInstanceTypesLists) {
            return supportedInstanceTypesLists(List.of(supportedInstanceTypesLists));
        }
        @CustomType.Setter
        public Builder tags(Map<String,Object> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        @CustomType.Setter
        public Builder zoneId(String zoneId) {
            this.zoneId = Objects.requireNonNull(zoneId);
            return this;
        }
        public GetDedicatedHostsHost build() {
            final var o = new GetDedicatedHostsHost();
            o.actionOnMaintenance = actionOnMaintenance;
            o.autoPlacement = autoPlacement;
            o.autoReleaseTime = autoReleaseTime;
            o.capacities = capacities;
            o.cores = cores;
            o.cpuOverCommitRatio = cpuOverCommitRatio;
            o.dedicatedHostId = dedicatedHostId;
            o.dedicatedHostName = dedicatedHostName;
            o.dedicatedHostType = dedicatedHostType;
            o.description = description;
            o.expiredTime = expiredTime;
            o.gpuSpec = gpuSpec;
            o.id = id;
            o.machineId = machineId;
            o.networkAttributes = networkAttributes;
            o.operationLocks = operationLocks;
            o.paymentType = paymentType;
            o.physicalGpus = physicalGpus;
            o.resourceGroupId = resourceGroupId;
            o.saleCycle = saleCycle;
            o.sockets = sockets;
            o.status = status;
            o.supportedCustomInstanceTypeFamilies = supportedCustomInstanceTypeFamilies;
            o.supportedInstanceTypeFamilies = supportedInstanceTypeFamilies;
            o.supportedInstanceTypesLists = supportedInstanceTypesLists;
            o.tags = tags;
            o.zoneId = zoneId;
            return o;
        }
    }
}