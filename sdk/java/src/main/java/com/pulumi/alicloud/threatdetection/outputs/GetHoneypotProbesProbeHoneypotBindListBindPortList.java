// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class GetHoneypotProbesProbeHoneypotBindListBindPortList {
    /**
     * @return Whether to bind the port.
     * 
     */
    private Boolean bindPort;
    /**
     * @return End port.
     * 
     */
    private Integer endPort;
    /**
     * @return Whether the port is fixed.
     * 
     */
    private Boolean fixed;
    /**
     * @return Start port.
     * 
     */
    private Integer startPort;
    /**
     * @return Destination port.
     * 
     */
    private Integer targetPort;

    private GetHoneypotProbesProbeHoneypotBindListBindPortList() {}
    /**
     * @return Whether to bind the port.
     * 
     */
    public Boolean bindPort() {
        return this.bindPort;
    }
    /**
     * @return End port.
     * 
     */
    public Integer endPort() {
        return this.endPort;
    }
    /**
     * @return Whether the port is fixed.
     * 
     */
    public Boolean fixed() {
        return this.fixed;
    }
    /**
     * @return Start port.
     * 
     */
    public Integer startPort() {
        return this.startPort;
    }
    /**
     * @return Destination port.
     * 
     */
    public Integer targetPort() {
        return this.targetPort;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetHoneypotProbesProbeHoneypotBindListBindPortList defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean bindPort;
        private Integer endPort;
        private Boolean fixed;
        private Integer startPort;
        private Integer targetPort;
        public Builder() {}
        public Builder(GetHoneypotProbesProbeHoneypotBindListBindPortList defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bindPort = defaults.bindPort;
    	      this.endPort = defaults.endPort;
    	      this.fixed = defaults.fixed;
    	      this.startPort = defaults.startPort;
    	      this.targetPort = defaults.targetPort;
        }

        @CustomType.Setter
        public Builder bindPort(Boolean bindPort) {
            this.bindPort = Objects.requireNonNull(bindPort);
            return this;
        }
        @CustomType.Setter
        public Builder endPort(Integer endPort) {
            this.endPort = Objects.requireNonNull(endPort);
            return this;
        }
        @CustomType.Setter
        public Builder fixed(Boolean fixed) {
            this.fixed = Objects.requireNonNull(fixed);
            return this;
        }
        @CustomType.Setter
        public Builder startPort(Integer startPort) {
            this.startPort = Objects.requireNonNull(startPort);
            return this;
        }
        @CustomType.Setter
        public Builder targetPort(Integer targetPort) {
            this.targetPort = Objects.requireNonNull(targetPort);
            return this;
        }
        public GetHoneypotProbesProbeHoneypotBindListBindPortList build() {
            final var o = new GetHoneypotProbesProbeHoneypotBindListBindPortList();
            o.bindPort = bindPort;
            o.endPort = endPort;
            o.fixed = fixed;
            o.startPort = startPort;
            o.targetPort = targetPort;
            return o;
        }
    }
}
