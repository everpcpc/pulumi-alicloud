// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.servicemesh.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ServiceMeshNetwork {
    /**
     * @return The ID of the VPC.
     * 
     */
    private String vpcId;
    /**
     * @return The list of Virtual Switch.
     * 
     */
    private String vswitcheList;

    private ServiceMeshNetwork() {}
    /**
     * @return The ID of the VPC.
     * 
     */
    public String vpcId() {
        return this.vpcId;
    }
    /**
     * @return The list of Virtual Switch.
     * 
     */
    public String vswitcheList() {
        return this.vswitcheList;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceMeshNetwork defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String vpcId;
        private String vswitcheList;
        public Builder() {}
        public Builder(ServiceMeshNetwork defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.vpcId = defaults.vpcId;
    	      this.vswitcheList = defaults.vswitcheList;
        }

        @CustomType.Setter
        public Builder vpcId(String vpcId) {
            this.vpcId = Objects.requireNonNull(vpcId);
            return this;
        }
        @CustomType.Setter
        public Builder vswitcheList(String vswitcheList) {
            this.vswitcheList = Objects.requireNonNull(vswitcheList);
            return this;
        }
        public ServiceMeshNetwork build() {
            final var o = new ServiceMeshNetwork();
            o.vpcId = vpcId;
            o.vswitcheList = vswitcheList;
            return o;
        }
    }
}
