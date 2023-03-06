// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.alicloud.ecs.outputs.GetEcsDedicatedHostClustersClusterDedicatedHostClusterCapacityLocalStorageCapacity;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetEcsDedicatedHostClustersClusterDedicatedHostClusterCapacity {
    /**
     * @return The available memory size. Unit: `GiB`.
     * 
     */
    private Integer availableMemory;
    /**
     * @return The number of available vCPUs.
     * 
     */
    private Integer availableVcpus;
    /**
     * @return The local storage.
     * 
     */
    private List<GetEcsDedicatedHostClustersClusterDedicatedHostClusterCapacityLocalStorageCapacity> localStorageCapacities;
    /**
     * @return The total memory size. Unit: `GiB`.
     * 
     */
    private Integer totalMemory;
    /**
     * @return The total number of vCPUs.
     * 
     */
    private Integer totalVcpus;

    private GetEcsDedicatedHostClustersClusterDedicatedHostClusterCapacity() {}
    /**
     * @return The available memory size. Unit: `GiB`.
     * 
     */
    public Integer availableMemory() {
        return this.availableMemory;
    }
    /**
     * @return The number of available vCPUs.
     * 
     */
    public Integer availableVcpus() {
        return this.availableVcpus;
    }
    /**
     * @return The local storage.
     * 
     */
    public List<GetEcsDedicatedHostClustersClusterDedicatedHostClusterCapacityLocalStorageCapacity> localStorageCapacities() {
        return this.localStorageCapacities;
    }
    /**
     * @return The total memory size. Unit: `GiB`.
     * 
     */
    public Integer totalMemory() {
        return this.totalMemory;
    }
    /**
     * @return The total number of vCPUs.
     * 
     */
    public Integer totalVcpus() {
        return this.totalVcpus;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEcsDedicatedHostClustersClusterDedicatedHostClusterCapacity defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer availableMemory;
        private Integer availableVcpus;
        private List<GetEcsDedicatedHostClustersClusterDedicatedHostClusterCapacityLocalStorageCapacity> localStorageCapacities;
        private Integer totalMemory;
        private Integer totalVcpus;
        public Builder() {}
        public Builder(GetEcsDedicatedHostClustersClusterDedicatedHostClusterCapacity defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.availableMemory = defaults.availableMemory;
    	      this.availableVcpus = defaults.availableVcpus;
    	      this.localStorageCapacities = defaults.localStorageCapacities;
    	      this.totalMemory = defaults.totalMemory;
    	      this.totalVcpus = defaults.totalVcpus;
        }

        @CustomType.Setter
        public Builder availableMemory(Integer availableMemory) {
            this.availableMemory = Objects.requireNonNull(availableMemory);
            return this;
        }
        @CustomType.Setter
        public Builder availableVcpus(Integer availableVcpus) {
            this.availableVcpus = Objects.requireNonNull(availableVcpus);
            return this;
        }
        @CustomType.Setter
        public Builder localStorageCapacities(List<GetEcsDedicatedHostClustersClusterDedicatedHostClusterCapacityLocalStorageCapacity> localStorageCapacities) {
            this.localStorageCapacities = Objects.requireNonNull(localStorageCapacities);
            return this;
        }
        public Builder localStorageCapacities(GetEcsDedicatedHostClustersClusterDedicatedHostClusterCapacityLocalStorageCapacity... localStorageCapacities) {
            return localStorageCapacities(List.of(localStorageCapacities));
        }
        @CustomType.Setter
        public Builder totalMemory(Integer totalMemory) {
            this.totalMemory = Objects.requireNonNull(totalMemory);
            return this;
        }
        @CustomType.Setter
        public Builder totalVcpus(Integer totalVcpus) {
            this.totalVcpus = Objects.requireNonNull(totalVcpus);
            return this;
        }
        public GetEcsDedicatedHostClustersClusterDedicatedHostClusterCapacity build() {
            final var o = new GetEcsDedicatedHostClustersClusterDedicatedHostClusterCapacity();
            o.availableMemory = availableMemory;
            o.availableVcpus = availableVcpus;
            o.localStorageCapacities = localStorageCapacities;
            o.totalMemory = totalMemory;
            o.totalVcpus = totalVcpus;
            return o;
        }
    }
}