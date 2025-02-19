// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.edas.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class K8sSlbAttachmentSlbConfigPortMappingServicePort {
    /**
     * @return The port of k8s service, values should be within range [1, 65535].
     * 
     */
    private Integer port;
    /**
     * @return The protocol of k8s service, values can be &#39;TCP&#39; or &#39;UDP&#39;.
     * 
     */
    private String protocol;
    /**
     * @return The port of k8s pod, values should be within range [1, 65535].
     * 
     */
    private Integer targetPort;

    private K8sSlbAttachmentSlbConfigPortMappingServicePort() {}
    /**
     * @return The port of k8s service, values should be within range [1, 65535].
     * 
     */
    public Integer port() {
        return this.port;
    }
    /**
     * @return The protocol of k8s service, values can be &#39;TCP&#39; or &#39;UDP&#39;.
     * 
     */
    public String protocol() {
        return this.protocol;
    }
    /**
     * @return The port of k8s pod, values should be within range [1, 65535].
     * 
     */
    public Integer targetPort() {
        return this.targetPort;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(K8sSlbAttachmentSlbConfigPortMappingServicePort defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer port;
        private String protocol;
        private Integer targetPort;
        public Builder() {}
        public Builder(K8sSlbAttachmentSlbConfigPortMappingServicePort defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.port = defaults.port;
    	      this.protocol = defaults.protocol;
    	      this.targetPort = defaults.targetPort;
        }

        @CustomType.Setter
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        @CustomType.Setter
        public Builder protocol(String protocol) {
            this.protocol = Objects.requireNonNull(protocol);
            return this;
        }
        @CustomType.Setter
        public Builder targetPort(Integer targetPort) {
            this.targetPort = Objects.requireNonNull(targetPort);
            return this;
        }
        public K8sSlbAttachmentSlbConfigPortMappingServicePort build() {
            final var o = new K8sSlbAttachmentSlbConfigPortMappingServicePort();
            o.port = port;
            o.protocol = protocol;
            o.targetPort = targetPort;
            return o;
        }
    }
}
