// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.servicemesh.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceMeshMeshConfigSidecarInjectorArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceMeshMeshConfigSidecarInjectorArgs Empty = new ServiceMeshMeshConfigSidecarInjectorArgs();

    /**
     * Whether to enable by Pod Annotations automatic injection Sidecar.
     * 
     */
    @Import(name="autoInjectionPolicyEnabled")
    private @Nullable Output<Boolean> autoInjectionPolicyEnabled;

    /**
     * @return Whether to enable by Pod Annotations automatic injection Sidecar.
     * 
     */
    public Optional<Output<Boolean>> autoInjectionPolicyEnabled() {
        return Optional.ofNullable(this.autoInjectionPolicyEnabled);
    }

    /**
     * Whether it is the all namespaces you turn on the auto injection capabilities.
     * 
     */
    @Import(name="enableNamespacesByDefault")
    private @Nullable Output<Boolean> enableNamespacesByDefault;

    /**
     * @return Whether it is the all namespaces you turn on the auto injection capabilities.
     * 
     */
    public Optional<Output<Boolean>> enableNamespacesByDefault() {
        return Optional.ofNullable(this.enableNamespacesByDefault);
    }

    /**
     * The limit cpu of the Sidecar injector Pods.
     * 
     */
    @Import(name="limitCpu")
    private @Nullable Output<String> limitCpu;

    /**
     * @return The limit cpu of the Sidecar injector Pods.
     * 
     */
    public Optional<Output<String>> limitCpu() {
        return Optional.ofNullable(this.limitCpu);
    }

    /**
     * Sidecar injector Pods on the throttle.
     * 
     */
    @Import(name="limitMemory")
    private @Nullable Output<String> limitMemory;

    /**
     * @return Sidecar injector Pods on the throttle.
     * 
     */
    public Optional<Output<String>> limitMemory() {
        return Optional.ofNullable(this.limitMemory);
    }

    /**
     * The requested cpu the Sidecar injector Pods.
     * 
     */
    @Import(name="requestCpu")
    private @Nullable Output<String> requestCpu;

    /**
     * @return The requested cpu the Sidecar injector Pods.
     * 
     */
    public Optional<Output<String>> requestCpu() {
        return Optional.ofNullable(this.requestCpu);
    }

    /**
     * The requested memory the Sidecar injector Pods.
     * 
     */
    @Import(name="requestMemory")
    private @Nullable Output<String> requestMemory;

    /**
     * @return The requested memory the Sidecar injector Pods.
     * 
     */
    public Optional<Output<String>> requestMemory() {
        return Optional.ofNullable(this.requestMemory);
    }

    private ServiceMeshMeshConfigSidecarInjectorArgs() {}

    private ServiceMeshMeshConfigSidecarInjectorArgs(ServiceMeshMeshConfigSidecarInjectorArgs $) {
        this.autoInjectionPolicyEnabled = $.autoInjectionPolicyEnabled;
        this.enableNamespacesByDefault = $.enableNamespacesByDefault;
        this.limitCpu = $.limitCpu;
        this.limitMemory = $.limitMemory;
        this.requestCpu = $.requestCpu;
        this.requestMemory = $.requestMemory;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceMeshMeshConfigSidecarInjectorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceMeshMeshConfigSidecarInjectorArgs $;

        public Builder() {
            $ = new ServiceMeshMeshConfigSidecarInjectorArgs();
        }

        public Builder(ServiceMeshMeshConfigSidecarInjectorArgs defaults) {
            $ = new ServiceMeshMeshConfigSidecarInjectorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param autoInjectionPolicyEnabled Whether to enable by Pod Annotations automatic injection Sidecar.
         * 
         * @return builder
         * 
         */
        public Builder autoInjectionPolicyEnabled(@Nullable Output<Boolean> autoInjectionPolicyEnabled) {
            $.autoInjectionPolicyEnabled = autoInjectionPolicyEnabled;
            return this;
        }

        /**
         * @param autoInjectionPolicyEnabled Whether to enable by Pod Annotations automatic injection Sidecar.
         * 
         * @return builder
         * 
         */
        public Builder autoInjectionPolicyEnabled(Boolean autoInjectionPolicyEnabled) {
            return autoInjectionPolicyEnabled(Output.of(autoInjectionPolicyEnabled));
        }

        /**
         * @param enableNamespacesByDefault Whether it is the all namespaces you turn on the auto injection capabilities.
         * 
         * @return builder
         * 
         */
        public Builder enableNamespacesByDefault(@Nullable Output<Boolean> enableNamespacesByDefault) {
            $.enableNamespacesByDefault = enableNamespacesByDefault;
            return this;
        }

        /**
         * @param enableNamespacesByDefault Whether it is the all namespaces you turn on the auto injection capabilities.
         * 
         * @return builder
         * 
         */
        public Builder enableNamespacesByDefault(Boolean enableNamespacesByDefault) {
            return enableNamespacesByDefault(Output.of(enableNamespacesByDefault));
        }

        /**
         * @param limitCpu The limit cpu of the Sidecar injector Pods.
         * 
         * @return builder
         * 
         */
        public Builder limitCpu(@Nullable Output<String> limitCpu) {
            $.limitCpu = limitCpu;
            return this;
        }

        /**
         * @param limitCpu The limit cpu of the Sidecar injector Pods.
         * 
         * @return builder
         * 
         */
        public Builder limitCpu(String limitCpu) {
            return limitCpu(Output.of(limitCpu));
        }

        /**
         * @param limitMemory Sidecar injector Pods on the throttle.
         * 
         * @return builder
         * 
         */
        public Builder limitMemory(@Nullable Output<String> limitMemory) {
            $.limitMemory = limitMemory;
            return this;
        }

        /**
         * @param limitMemory Sidecar injector Pods on the throttle.
         * 
         * @return builder
         * 
         */
        public Builder limitMemory(String limitMemory) {
            return limitMemory(Output.of(limitMemory));
        }

        /**
         * @param requestCpu The requested cpu the Sidecar injector Pods.
         * 
         * @return builder
         * 
         */
        public Builder requestCpu(@Nullable Output<String> requestCpu) {
            $.requestCpu = requestCpu;
            return this;
        }

        /**
         * @param requestCpu The requested cpu the Sidecar injector Pods.
         * 
         * @return builder
         * 
         */
        public Builder requestCpu(String requestCpu) {
            return requestCpu(Output.of(requestCpu));
        }

        /**
         * @param requestMemory The requested memory the Sidecar injector Pods.
         * 
         * @return builder
         * 
         */
        public Builder requestMemory(@Nullable Output<String> requestMemory) {
            $.requestMemory = requestMemory;
            return this;
        }

        /**
         * @param requestMemory The requested memory the Sidecar injector Pods.
         * 
         * @return builder
         * 
         */
        public Builder requestMemory(String requestMemory) {
            return requestMemory(Output.of(requestMemory));
        }

        public ServiceMeshMeshConfigSidecarInjectorArgs build() {
            return $;
        }
    }

}
