// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetBasicAccelerateIpsIp {
    /**
     * @return The address of the Basic Accelerate IP.
     * 
     */
    private String accelerateIpAddress;
    /**
     * @return The id of the Basic Accelerate IP.
     * 
     */
    private String accelerateIpId;
    /**
     * @return The id of the Global Accelerator Basic Accelerator instance.
     * 
     */
    private String acceleratorId;
    /**
     * @return The id of the Basic Accelerate IP.
     * 
     */
    private String id;
    /**
     * @return The ID of the Basic Ip Set.
     * 
     */
    private String ipSetId;
    /**
     * @return The status of the Global Accelerator Basic Accelerate IP instance. Valid Value: `active`, `binding`, `bound`, `unbinding`, `deleting`.
     * 
     */
    private String status;

    private GetBasicAccelerateIpsIp() {}
    /**
     * @return The address of the Basic Accelerate IP.
     * 
     */
    public String accelerateIpAddress() {
        return this.accelerateIpAddress;
    }
    /**
     * @return The id of the Basic Accelerate IP.
     * 
     */
    public String accelerateIpId() {
        return this.accelerateIpId;
    }
    /**
     * @return The id of the Global Accelerator Basic Accelerator instance.
     * 
     */
    public String acceleratorId() {
        return this.acceleratorId;
    }
    /**
     * @return The id of the Basic Accelerate IP.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The ID of the Basic Ip Set.
     * 
     */
    public String ipSetId() {
        return this.ipSetId;
    }
    /**
     * @return The status of the Global Accelerator Basic Accelerate IP instance. Valid Value: `active`, `binding`, `bound`, `unbinding`, `deleting`.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBasicAccelerateIpsIp defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accelerateIpAddress;
        private String accelerateIpId;
        private String acceleratorId;
        private String id;
        private String ipSetId;
        private String status;
        public Builder() {}
        public Builder(GetBasicAccelerateIpsIp defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accelerateIpAddress = defaults.accelerateIpAddress;
    	      this.accelerateIpId = defaults.accelerateIpId;
    	      this.acceleratorId = defaults.acceleratorId;
    	      this.id = defaults.id;
    	      this.ipSetId = defaults.ipSetId;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder accelerateIpAddress(String accelerateIpAddress) {
            this.accelerateIpAddress = Objects.requireNonNull(accelerateIpAddress);
            return this;
        }
        @CustomType.Setter
        public Builder accelerateIpId(String accelerateIpId) {
            this.accelerateIpId = Objects.requireNonNull(accelerateIpId);
            return this;
        }
        @CustomType.Setter
        public Builder acceleratorId(String acceleratorId) {
            this.acceleratorId = Objects.requireNonNull(acceleratorId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ipSetId(String ipSetId) {
            this.ipSetId = Objects.requireNonNull(ipSetId);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetBasicAccelerateIpsIp build() {
            final var o = new GetBasicAccelerateIpsIp();
            o.accelerateIpAddress = accelerateIpAddress;
            o.accelerateIpId = accelerateIpId;
            o.acceleratorId = acceleratorId;
            o.id = id;
            o.ipSetId = ipSetId;
            o.status = status;
            return o;
        }
    }
}
