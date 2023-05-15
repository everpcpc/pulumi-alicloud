// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cs.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class EdgeKubernetesCertificateAuthority {
    /**
     * @return The path of client certificate, like `~/.kube/client-cert.pem`.
     * 
     */
    private @Nullable String clientCert;
    /**
     * @return The path of client key, like `~/.kube/client-key.pem`.
     * 
     */
    private @Nullable String clientKey;
    /**
     * @return The base64 encoded cluster certificate data required to communicate with your cluster. Add this to the certificate-authority-data section of the kubeconfig file for your cluster.
     * 
     */
    private @Nullable String clusterCert;

    private EdgeKubernetesCertificateAuthority() {}
    /**
     * @return The path of client certificate, like `~/.kube/client-cert.pem`.
     * 
     */
    public Optional<String> clientCert() {
        return Optional.ofNullable(this.clientCert);
    }
    /**
     * @return The path of client key, like `~/.kube/client-key.pem`.
     * 
     */
    public Optional<String> clientKey() {
        return Optional.ofNullable(this.clientKey);
    }
    /**
     * @return The base64 encoded cluster certificate data required to communicate with your cluster. Add this to the certificate-authority-data section of the kubeconfig file for your cluster.
     * 
     */
    public Optional<String> clusterCert() {
        return Optional.ofNullable(this.clusterCert);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(EdgeKubernetesCertificateAuthority defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String clientCert;
        private @Nullable String clientKey;
        private @Nullable String clusterCert;
        public Builder() {}
        public Builder(EdgeKubernetesCertificateAuthority defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clientCert = defaults.clientCert;
    	      this.clientKey = defaults.clientKey;
    	      this.clusterCert = defaults.clusterCert;
        }

        @CustomType.Setter
        public Builder clientCert(@Nullable String clientCert) {
            this.clientCert = clientCert;
            return this;
        }
        @CustomType.Setter
        public Builder clientKey(@Nullable String clientKey) {
            this.clientKey = clientKey;
            return this;
        }
        @CustomType.Setter
        public Builder clusterCert(@Nullable String clusterCert) {
            this.clusterCert = clusterCert;
            return this;
        }
        public EdgeKubernetesCertificateAuthority build() {
            final var o = new EdgeKubernetesCertificateAuthority();
            o.clientCert = clientCert;
            o.clientKey = clientKey;
            o.clusterCert = clusterCert;
            return o;
        }
    }
}
