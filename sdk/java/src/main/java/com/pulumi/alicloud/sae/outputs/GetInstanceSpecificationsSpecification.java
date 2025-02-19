// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.sae.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetInstanceSpecificationsSpecification {
    /**
     * @return CPU Size, Specifications for Micronucleus.
     * 
     */
    private Integer cpu;
    /**
     * @return Whether the instance is available. The value description is as follows:
     * 
     */
    private Boolean enable;
    /**
     * @return The ID of the Instance Specification.
     * 
     */
    private String id;
    /**
     * @return The first ID of the resource.
     * 
     */
    private String instanceSpecificationId;
    /**
     * @return The Memory specifications for the MB.
     * 
     */
    private Integer memory;
    /**
     * @return The specification configuration name.
     * 
     */
    private String specInfo;
    /**
     * @return The specification configuration version.
     * 
     */
    private Integer version;

    private GetInstanceSpecificationsSpecification() {}
    /**
     * @return CPU Size, Specifications for Micronucleus.
     * 
     */
    public Integer cpu() {
        return this.cpu;
    }
    /**
     * @return Whether the instance is available. The value description is as follows:
     * 
     */
    public Boolean enable() {
        return this.enable;
    }
    /**
     * @return The ID of the Instance Specification.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The first ID of the resource.
     * 
     */
    public String instanceSpecificationId() {
        return this.instanceSpecificationId;
    }
    /**
     * @return The Memory specifications for the MB.
     * 
     */
    public Integer memory() {
        return this.memory;
    }
    /**
     * @return The specification configuration name.
     * 
     */
    public String specInfo() {
        return this.specInfo;
    }
    /**
     * @return The specification configuration version.
     * 
     */
    public Integer version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstanceSpecificationsSpecification defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer cpu;
        private Boolean enable;
        private String id;
        private String instanceSpecificationId;
        private Integer memory;
        private String specInfo;
        private Integer version;
        public Builder() {}
        public Builder(GetInstanceSpecificationsSpecification defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cpu = defaults.cpu;
    	      this.enable = defaults.enable;
    	      this.id = defaults.id;
    	      this.instanceSpecificationId = defaults.instanceSpecificationId;
    	      this.memory = defaults.memory;
    	      this.specInfo = defaults.specInfo;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder cpu(Integer cpu) {
            this.cpu = Objects.requireNonNull(cpu);
            return this;
        }
        @CustomType.Setter
        public Builder enable(Boolean enable) {
            this.enable = Objects.requireNonNull(enable);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceSpecificationId(String instanceSpecificationId) {
            this.instanceSpecificationId = Objects.requireNonNull(instanceSpecificationId);
            return this;
        }
        @CustomType.Setter
        public Builder memory(Integer memory) {
            this.memory = Objects.requireNonNull(memory);
            return this;
        }
        @CustomType.Setter
        public Builder specInfo(String specInfo) {
            this.specInfo = Objects.requireNonNull(specInfo);
            return this;
        }
        @CustomType.Setter
        public Builder version(Integer version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }
        public GetInstanceSpecificationsSpecification build() {
            final var o = new GetInstanceSpecificationsSpecification();
            o.cpu = cpu;
            o.enable = enable;
            o.id = id;
            o.instanceSpecificationId = instanceSpecificationId;
            o.memory = memory;
            o.specInfo = specInfo;
            o.version = version;
            return o;
        }
    }
}
