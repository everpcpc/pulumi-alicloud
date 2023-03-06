// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ListenerXForwardedForConfig {
    /**
     * @return The Custom Header Field Names Only When `x_forwarded_for_client_cert_client_verify_enabled` Has a Value of True, this Value Will Not Take Effect until.The name must be 1 to 40 characters in length, and can contain letters, hyphens (-), underscores (_), and digits.
     * 
     */
    private @Nullable String xForwardedForClientCertClientVerifyAlias;
    /**
     * @return Indicates Whether the `X-Forwarded-Clientcert-clientverify` Header Field Is Used to Obtain Access to the Server Load Balancer Instance of the Client Certificate to Verify the Results.
     * 
     */
    private @Nullable Boolean xForwardedForClientCertClientVerifyEnabled;
    /**
     * @return The Custom Header Field Names Only When `x_forwarded_for_client_certfingerprint_enabled`, Which Evaluates to True When the Entry into Force of.The name must be 1 to 40 characters in length, and can contain letters, hyphens (-), underscores (_), and digits.
     * 
     */
    private @Nullable String xForwardedForClientCertFingerPrintAlias;
    /**
     * @return Indicates Whether the `X-Forwarded-client_cert-fingerprint` Header Field Is Used to Obtain Access to the Server Load Balancer Instance of the Client Certificate Fingerprint Value.
     * 
     */
    private @Nullable Boolean xForwardedForClientCertFingerPrintEnabled;
    /**
     * @return The Custom Header Field Names Only When `x_forwarded_for_client_cert_issuer_dn_enabled`, Which Evaluates to True When the Entry into Force of.
     * 
     */
    private @Nullable String xForwardedForClientCertIssuerDnAlias;
    /**
     * @return Indicates Whether the `X-Forwarded-Clientcert-issuerdn` Header Field Is Used to Obtain Access to the Server Load Balancer Instance of the Client Certificate after the Manifests Are Signed, the Publisher Information.
     * 
     */
    private @Nullable Boolean xForwardedForClientCertIssuerDnEnabled;
    /**
     * @return The name of the custom header. This parameter is valid only if `x_forwarded_for_client_certsubjectdn_enabled` is set to true. The name must be 1 to 40 characters in length, and can contain letters, hyphens (-), underscores (_), and digits.
     * 
     */
    private @Nullable String xForwardedForClientCertSubjectDnAlias;
    /**
     * @return Specifies whether to use the `X-Forwarded-client_cert-subjectdn` header field to obtain information about the owner of the ALB client certificate. Valid values: true and false. Default value: false.
     * 
     */
    private @Nullable Boolean xForwardedForClientCertSubjectDnEnabled;
    /**
     * @return Indicates Whether the X-Forwarded-Client-Port Header Field Is Used to Obtain Access to Server Load Balancer Instances to the Client, and Those of the Ports.
     * 
     */
    private @Nullable Boolean xForwardedForClientSrcPortEnabled;
    /**
     * @return Whether to Enable by X-Forwarded-For Header Field Is Used to Obtain the Client IP Addresses.
     * 
     */
    private @Nullable Boolean xForwardedForEnabled;
    /**
     * @return Indicates Whether the X-Forwarded-Proto Header Field Is Used to Obtain the Server Load Balancer Instance Snooping Protocols.
     * 
     */
    private @Nullable Boolean xForwardedForProtoEnabled;
    /**
     * @return Indicates Whether the SLB-ID Header Field Is Used to Obtain the Load Balancing Instance Id.
     * 
     */
    private @Nullable Boolean xForwardedForSlbIdEnabled;
    /**
     * @return Indicates Whether the X-Forwarded-Port Header Field Is Used to Obtain the Server Load Balancer Instance Listening Port.
     * 
     */
    private @Nullable Boolean xForwardedForSlbPortEnabled;

    private ListenerXForwardedForConfig() {}
    /**
     * @return The Custom Header Field Names Only When `x_forwarded_for_client_cert_client_verify_enabled` Has a Value of True, this Value Will Not Take Effect until.The name must be 1 to 40 characters in length, and can contain letters, hyphens (-), underscores (_), and digits.
     * 
     */
    public Optional<String> xForwardedForClientCertClientVerifyAlias() {
        return Optional.ofNullable(this.xForwardedForClientCertClientVerifyAlias);
    }
    /**
     * @return Indicates Whether the `X-Forwarded-Clientcert-clientverify` Header Field Is Used to Obtain Access to the Server Load Balancer Instance of the Client Certificate to Verify the Results.
     * 
     */
    public Optional<Boolean> xForwardedForClientCertClientVerifyEnabled() {
        return Optional.ofNullable(this.xForwardedForClientCertClientVerifyEnabled);
    }
    /**
     * @return The Custom Header Field Names Only When `x_forwarded_for_client_certfingerprint_enabled`, Which Evaluates to True When the Entry into Force of.The name must be 1 to 40 characters in length, and can contain letters, hyphens (-), underscores (_), and digits.
     * 
     */
    public Optional<String> xForwardedForClientCertFingerPrintAlias() {
        return Optional.ofNullable(this.xForwardedForClientCertFingerPrintAlias);
    }
    /**
     * @return Indicates Whether the `X-Forwarded-client_cert-fingerprint` Header Field Is Used to Obtain Access to the Server Load Balancer Instance of the Client Certificate Fingerprint Value.
     * 
     */
    public Optional<Boolean> xForwardedForClientCertFingerPrintEnabled() {
        return Optional.ofNullable(this.xForwardedForClientCertFingerPrintEnabled);
    }
    /**
     * @return The Custom Header Field Names Only When `x_forwarded_for_client_cert_issuer_dn_enabled`, Which Evaluates to True When the Entry into Force of.
     * 
     */
    public Optional<String> xForwardedForClientCertIssuerDnAlias() {
        return Optional.ofNullable(this.xForwardedForClientCertIssuerDnAlias);
    }
    /**
     * @return Indicates Whether the `X-Forwarded-Clientcert-issuerdn` Header Field Is Used to Obtain Access to the Server Load Balancer Instance of the Client Certificate after the Manifests Are Signed, the Publisher Information.
     * 
     */
    public Optional<Boolean> xForwardedForClientCertIssuerDnEnabled() {
        return Optional.ofNullable(this.xForwardedForClientCertIssuerDnEnabled);
    }
    /**
     * @return The name of the custom header. This parameter is valid only if `x_forwarded_for_client_certsubjectdn_enabled` is set to true. The name must be 1 to 40 characters in length, and can contain letters, hyphens (-), underscores (_), and digits.
     * 
     */
    public Optional<String> xForwardedForClientCertSubjectDnAlias() {
        return Optional.ofNullable(this.xForwardedForClientCertSubjectDnAlias);
    }
    /**
     * @return Specifies whether to use the `X-Forwarded-client_cert-subjectdn` header field to obtain information about the owner of the ALB client certificate. Valid values: true and false. Default value: false.
     * 
     */
    public Optional<Boolean> xForwardedForClientCertSubjectDnEnabled() {
        return Optional.ofNullable(this.xForwardedForClientCertSubjectDnEnabled);
    }
    /**
     * @return Indicates Whether the X-Forwarded-Client-Port Header Field Is Used to Obtain Access to Server Load Balancer Instances to the Client, and Those of the Ports.
     * 
     */
    public Optional<Boolean> xForwardedForClientSrcPortEnabled() {
        return Optional.ofNullable(this.xForwardedForClientSrcPortEnabled);
    }
    /**
     * @return Whether to Enable by X-Forwarded-For Header Field Is Used to Obtain the Client IP Addresses.
     * 
     */
    public Optional<Boolean> xForwardedForEnabled() {
        return Optional.ofNullable(this.xForwardedForEnabled);
    }
    /**
     * @return Indicates Whether the X-Forwarded-Proto Header Field Is Used to Obtain the Server Load Balancer Instance Snooping Protocols.
     * 
     */
    public Optional<Boolean> xForwardedForProtoEnabled() {
        return Optional.ofNullable(this.xForwardedForProtoEnabled);
    }
    /**
     * @return Indicates Whether the SLB-ID Header Field Is Used to Obtain the Load Balancing Instance Id.
     * 
     */
    public Optional<Boolean> xForwardedForSlbIdEnabled() {
        return Optional.ofNullable(this.xForwardedForSlbIdEnabled);
    }
    /**
     * @return Indicates Whether the X-Forwarded-Port Header Field Is Used to Obtain the Server Load Balancer Instance Listening Port.
     * 
     */
    public Optional<Boolean> xForwardedForSlbPortEnabled() {
        return Optional.ofNullable(this.xForwardedForSlbPortEnabled);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ListenerXForwardedForConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String xForwardedForClientCertClientVerifyAlias;
        private @Nullable Boolean xForwardedForClientCertClientVerifyEnabled;
        private @Nullable String xForwardedForClientCertFingerPrintAlias;
        private @Nullable Boolean xForwardedForClientCertFingerPrintEnabled;
        private @Nullable String xForwardedForClientCertIssuerDnAlias;
        private @Nullable Boolean xForwardedForClientCertIssuerDnEnabled;
        private @Nullable String xForwardedForClientCertSubjectDnAlias;
        private @Nullable Boolean xForwardedForClientCertSubjectDnEnabled;
        private @Nullable Boolean xForwardedForClientSrcPortEnabled;
        private @Nullable Boolean xForwardedForEnabled;
        private @Nullable Boolean xForwardedForProtoEnabled;
        private @Nullable Boolean xForwardedForSlbIdEnabled;
        private @Nullable Boolean xForwardedForSlbPortEnabled;
        public Builder() {}
        public Builder(ListenerXForwardedForConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.xForwardedForClientCertClientVerifyAlias = defaults.xForwardedForClientCertClientVerifyAlias;
    	      this.xForwardedForClientCertClientVerifyEnabled = defaults.xForwardedForClientCertClientVerifyEnabled;
    	      this.xForwardedForClientCertFingerPrintAlias = defaults.xForwardedForClientCertFingerPrintAlias;
    	      this.xForwardedForClientCertFingerPrintEnabled = defaults.xForwardedForClientCertFingerPrintEnabled;
    	      this.xForwardedForClientCertIssuerDnAlias = defaults.xForwardedForClientCertIssuerDnAlias;
    	      this.xForwardedForClientCertIssuerDnEnabled = defaults.xForwardedForClientCertIssuerDnEnabled;
    	      this.xForwardedForClientCertSubjectDnAlias = defaults.xForwardedForClientCertSubjectDnAlias;
    	      this.xForwardedForClientCertSubjectDnEnabled = defaults.xForwardedForClientCertSubjectDnEnabled;
    	      this.xForwardedForClientSrcPortEnabled = defaults.xForwardedForClientSrcPortEnabled;
    	      this.xForwardedForEnabled = defaults.xForwardedForEnabled;
    	      this.xForwardedForProtoEnabled = defaults.xForwardedForProtoEnabled;
    	      this.xForwardedForSlbIdEnabled = defaults.xForwardedForSlbIdEnabled;
    	      this.xForwardedForSlbPortEnabled = defaults.xForwardedForSlbPortEnabled;
        }

        @CustomType.Setter
        public Builder xForwardedForClientCertClientVerifyAlias(@Nullable String xForwardedForClientCertClientVerifyAlias) {
            this.xForwardedForClientCertClientVerifyAlias = xForwardedForClientCertClientVerifyAlias;
            return this;
        }
        @CustomType.Setter
        public Builder xForwardedForClientCertClientVerifyEnabled(@Nullable Boolean xForwardedForClientCertClientVerifyEnabled) {
            this.xForwardedForClientCertClientVerifyEnabled = xForwardedForClientCertClientVerifyEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder xForwardedForClientCertFingerPrintAlias(@Nullable String xForwardedForClientCertFingerPrintAlias) {
            this.xForwardedForClientCertFingerPrintAlias = xForwardedForClientCertFingerPrintAlias;
            return this;
        }
        @CustomType.Setter
        public Builder xForwardedForClientCertFingerPrintEnabled(@Nullable Boolean xForwardedForClientCertFingerPrintEnabled) {
            this.xForwardedForClientCertFingerPrintEnabled = xForwardedForClientCertFingerPrintEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder xForwardedForClientCertIssuerDnAlias(@Nullable String xForwardedForClientCertIssuerDnAlias) {
            this.xForwardedForClientCertIssuerDnAlias = xForwardedForClientCertIssuerDnAlias;
            return this;
        }
        @CustomType.Setter
        public Builder xForwardedForClientCertIssuerDnEnabled(@Nullable Boolean xForwardedForClientCertIssuerDnEnabled) {
            this.xForwardedForClientCertIssuerDnEnabled = xForwardedForClientCertIssuerDnEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder xForwardedForClientCertSubjectDnAlias(@Nullable String xForwardedForClientCertSubjectDnAlias) {
            this.xForwardedForClientCertSubjectDnAlias = xForwardedForClientCertSubjectDnAlias;
            return this;
        }
        @CustomType.Setter
        public Builder xForwardedForClientCertSubjectDnEnabled(@Nullable Boolean xForwardedForClientCertSubjectDnEnabled) {
            this.xForwardedForClientCertSubjectDnEnabled = xForwardedForClientCertSubjectDnEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder xForwardedForClientSrcPortEnabled(@Nullable Boolean xForwardedForClientSrcPortEnabled) {
            this.xForwardedForClientSrcPortEnabled = xForwardedForClientSrcPortEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder xForwardedForEnabled(@Nullable Boolean xForwardedForEnabled) {
            this.xForwardedForEnabled = xForwardedForEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder xForwardedForProtoEnabled(@Nullable Boolean xForwardedForProtoEnabled) {
            this.xForwardedForProtoEnabled = xForwardedForProtoEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder xForwardedForSlbIdEnabled(@Nullable Boolean xForwardedForSlbIdEnabled) {
            this.xForwardedForSlbIdEnabled = xForwardedForSlbIdEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder xForwardedForSlbPortEnabled(@Nullable Boolean xForwardedForSlbPortEnabled) {
            this.xForwardedForSlbPortEnabled = xForwardedForSlbPortEnabled;
            return this;
        }
        public ListenerXForwardedForConfig build() {
            final var o = new ListenerXForwardedForConfig();
            o.xForwardedForClientCertClientVerifyAlias = xForwardedForClientCertClientVerifyAlias;
            o.xForwardedForClientCertClientVerifyEnabled = xForwardedForClientCertClientVerifyEnabled;
            o.xForwardedForClientCertFingerPrintAlias = xForwardedForClientCertFingerPrintAlias;
            o.xForwardedForClientCertFingerPrintEnabled = xForwardedForClientCertFingerPrintEnabled;
            o.xForwardedForClientCertIssuerDnAlias = xForwardedForClientCertIssuerDnAlias;
            o.xForwardedForClientCertIssuerDnEnabled = xForwardedForClientCertIssuerDnEnabled;
            o.xForwardedForClientCertSubjectDnAlias = xForwardedForClientCertSubjectDnAlias;
            o.xForwardedForClientCertSubjectDnEnabled = xForwardedForClientCertSubjectDnEnabled;
            o.xForwardedForClientSrcPortEnabled = xForwardedForClientSrcPortEnabled;
            o.xForwardedForEnabled = xForwardedForEnabled;
            o.xForwardedForProtoEnabled = xForwardedForProtoEnabled;
            o.xForwardedForSlbIdEnabled = xForwardedForSlbIdEnabled;
            o.xForwardedForSlbPortEnabled = xForwardedForSlbPortEnabled;
            return o;
        }
    }
}