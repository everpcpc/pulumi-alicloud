// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetKeyPairsKeyPairInstance {
    /**
     * @return The ID of the availability zone where the ECS instance is located.
     * 
     */
    private String availabilityZone;
    private String description;
    private String imageId;
    /**
     * @return The ID of the ECS instance.
     * 
     */
    private String instanceId;
    /**
     * @return The name of the ECS instance.
     * 
     */
    private String instanceName;
    private String instanceType;
    /**
     * @return Name of the key pair.
     * 
     */
    private String keyName;
    /**
     * @return The private IP address of the ECS instance.
     * 
     */
    private String privateIp;
    /**
     * @return The public IP address or EIP of the ECS instance.
     * 
     */
    private String publicIp;
    private String regionId;
    private String status;
    /**
     * @return The ID of the VSwitch attached to the ECS instance.
     * 
     */
    private String vswitchId;

    private GetKeyPairsKeyPairInstance() {}
    /**
     * @return The ID of the availability zone where the ECS instance is located.
     * 
     */
    public String availabilityZone() {
        return this.availabilityZone;
    }
    public String description() {
        return this.description;
    }
    public String imageId() {
        return this.imageId;
    }
    /**
     * @return The ID of the ECS instance.
     * 
     */
    public String instanceId() {
        return this.instanceId;
    }
    /**
     * @return The name of the ECS instance.
     * 
     */
    public String instanceName() {
        return this.instanceName;
    }
    public String instanceType() {
        return this.instanceType;
    }
    /**
     * @return Name of the key pair.
     * 
     */
    public String keyName() {
        return this.keyName;
    }
    /**
     * @return The private IP address of the ECS instance.
     * 
     */
    public String privateIp() {
        return this.privateIp;
    }
    /**
     * @return The public IP address or EIP of the ECS instance.
     * 
     */
    public String publicIp() {
        return this.publicIp;
    }
    public String regionId() {
        return this.regionId;
    }
    public String status() {
        return this.status;
    }
    /**
     * @return The ID of the VSwitch attached to the ECS instance.
     * 
     */
    public String vswitchId() {
        return this.vswitchId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKeyPairsKeyPairInstance defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String availabilityZone;
        private String description;
        private String imageId;
        private String instanceId;
        private String instanceName;
        private String instanceType;
        private String keyName;
        private String privateIp;
        private String publicIp;
        private String regionId;
        private String status;
        private String vswitchId;
        public Builder() {}
        public Builder(GetKeyPairsKeyPairInstance defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.availabilityZone = defaults.availabilityZone;
    	      this.description = defaults.description;
    	      this.imageId = defaults.imageId;
    	      this.instanceId = defaults.instanceId;
    	      this.instanceName = defaults.instanceName;
    	      this.instanceType = defaults.instanceType;
    	      this.keyName = defaults.keyName;
    	      this.privateIp = defaults.privateIp;
    	      this.publicIp = defaults.publicIp;
    	      this.regionId = defaults.regionId;
    	      this.status = defaults.status;
    	      this.vswitchId = defaults.vswitchId;
        }

        @CustomType.Setter
        public Builder availabilityZone(String availabilityZone) {
            this.availabilityZone = Objects.requireNonNull(availabilityZone);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder imageId(String imageId) {
            this.imageId = Objects.requireNonNull(imageId);
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(String instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder instanceName(String instanceName) {
            this.instanceName = Objects.requireNonNull(instanceName);
            return this;
        }
        @CustomType.Setter
        public Builder instanceType(String instanceType) {
            this.instanceType = Objects.requireNonNull(instanceType);
            return this;
        }
        @CustomType.Setter
        public Builder keyName(String keyName) {
            this.keyName = Objects.requireNonNull(keyName);
            return this;
        }
        @CustomType.Setter
        public Builder privateIp(String privateIp) {
            this.privateIp = Objects.requireNonNull(privateIp);
            return this;
        }
        @CustomType.Setter
        public Builder publicIp(String publicIp) {
            this.publicIp = Objects.requireNonNull(publicIp);
            return this;
        }
        @CustomType.Setter
        public Builder regionId(String regionId) {
            this.regionId = Objects.requireNonNull(regionId);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder vswitchId(String vswitchId) {
            this.vswitchId = Objects.requireNonNull(vswitchId);
            return this;
        }
        public GetKeyPairsKeyPairInstance build() {
            final var o = new GetKeyPairsKeyPairInstance();
            o.availabilityZone = availabilityZone;
            o.description = description;
            o.imageId = imageId;
            o.instanceId = instanceId;
            o.instanceName = instanceName;
            o.instanceType = instanceType;
            o.keyName = keyName;
            o.privateIp = privateIp;
            o.publicIp = publicIp;
            o.regionId = regionId;
            o.status = status;
            o.vswitchId = vswitchId;
            return o;
        }
    }
}
