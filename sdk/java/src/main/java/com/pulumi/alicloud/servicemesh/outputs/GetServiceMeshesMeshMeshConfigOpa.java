// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.servicemesh.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetServiceMeshesMeshMeshConfigOpa {
    /**
     * @return Whether to enable Service grid audit.
     * 
     */
    private Boolean enabled;
    /**
     * @return Sidecar injector Pods on the throttle.
     * 
     */
    private String limitCpu;
    /**
     * @return The memory limit  of the Sidecar injector Pods.
     * 
     */
    private String limitMemory;
    /**
     * @return The log level of the OPA proxy container .
     * 
     */
    private String logLevel;
    /**
     * @return The requested cpu the Sidecar injector Pods.
     * 
     */
    private String requestCpu;
    /**
     * @return The requested memory the Sidecar injector Pods.
     * 
     */
    private String requestMemory;

    private GetServiceMeshesMeshMeshConfigOpa() {}
    /**
     * @return Whether to enable Service grid audit.
     * 
     */
    public Boolean enabled() {
        return this.enabled;
    }
    /**
     * @return Sidecar injector Pods on the throttle.
     * 
     */
    public String limitCpu() {
        return this.limitCpu;
    }
    /**
     * @return The memory limit  of the Sidecar injector Pods.
     * 
     */
    public String limitMemory() {
        return this.limitMemory;
    }
    /**
     * @return The log level of the OPA proxy container .
     * 
     */
    public String logLevel() {
        return this.logLevel;
    }
    /**
     * @return The requested cpu the Sidecar injector Pods.
     * 
     */
    public String requestCpu() {
        return this.requestCpu;
    }
    /**
     * @return The requested memory the Sidecar injector Pods.
     * 
     */
    public String requestMemory() {
        return this.requestMemory;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServiceMeshesMeshMeshConfigOpa defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean enabled;
        private String limitCpu;
        private String limitMemory;
        private String logLevel;
        private String requestCpu;
        private String requestMemory;
        public Builder() {}
        public Builder(GetServiceMeshesMeshMeshConfigOpa defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enabled = defaults.enabled;
    	      this.limitCpu = defaults.limitCpu;
    	      this.limitMemory = defaults.limitMemory;
    	      this.logLevel = defaults.logLevel;
    	      this.requestCpu = defaults.requestCpu;
    	      this.requestMemory = defaults.requestMemory;
        }

        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            this.enabled = Objects.requireNonNull(enabled);
            return this;
        }
        @CustomType.Setter
        public Builder limitCpu(String limitCpu) {
            this.limitCpu = Objects.requireNonNull(limitCpu);
            return this;
        }
        @CustomType.Setter
        public Builder limitMemory(String limitMemory) {
            this.limitMemory = Objects.requireNonNull(limitMemory);
            return this;
        }
        @CustomType.Setter
        public Builder logLevel(String logLevel) {
            this.logLevel = Objects.requireNonNull(logLevel);
            return this;
        }
        @CustomType.Setter
        public Builder requestCpu(String requestCpu) {
            this.requestCpu = Objects.requireNonNull(requestCpu);
            return this;
        }
        @CustomType.Setter
        public Builder requestMemory(String requestMemory) {
            this.requestMemory = Objects.requireNonNull(requestMemory);
            return this;
        }
        public GetServiceMeshesMeshMeshConfigOpa build() {
            final var o = new GetServiceMeshesMeshMeshConfigOpa();
            o.enabled = enabled;
            o.limitCpu = limitCpu;
            o.limitMemory = limitMemory;
            o.logLevel = logLevel;
            o.requestCpu = requestCpu;
            o.requestMemory = requestMemory;
            return o;
        }
    }
}
