// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cdn.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DomainNewCertificateConfig {
    /**
     * @return The SSL certificate name.
     * 
     */
    private @Nullable String certName;
    /**
     * @return The SSL certificate type, can be &#34;upload&#34;, &#34;cas&#34; and &#34;free&#34;.
     * 
     */
    private @Nullable String certType;
    /**
     * @return Set `1` to ignore the repeated verification for certificate name, and cover the information of the origin certificate (with the same name). Set `0` to work the verification.
     * 
     */
    private @Nullable String forceSet;
    /**
     * @return The SSL private key. This is required if `server_certificate_status` is `on`
     * 
     */
    private @Nullable String privateKey;
    /**
     * @return The SSL server certificate string. This is required if `server_certificate_status` is `on`
     * 
     */
    private @Nullable String serverCertificate;
    /**
     * @return This parameter indicates whether or not enable https. Valid values are `on` and `off`. Default value is `on`.
     * 
     */
    private @Nullable String serverCertificateStatus;

    private DomainNewCertificateConfig() {}
    /**
     * @return The SSL certificate name.
     * 
     */
    public Optional<String> certName() {
        return Optional.ofNullable(this.certName);
    }
    /**
     * @return The SSL certificate type, can be &#34;upload&#34;, &#34;cas&#34; and &#34;free&#34;.
     * 
     */
    public Optional<String> certType() {
        return Optional.ofNullable(this.certType);
    }
    /**
     * @return Set `1` to ignore the repeated verification for certificate name, and cover the information of the origin certificate (with the same name). Set `0` to work the verification.
     * 
     */
    public Optional<String> forceSet() {
        return Optional.ofNullable(this.forceSet);
    }
    /**
     * @return The SSL private key. This is required if `server_certificate_status` is `on`
     * 
     */
    public Optional<String> privateKey() {
        return Optional.ofNullable(this.privateKey);
    }
    /**
     * @return The SSL server certificate string. This is required if `server_certificate_status` is `on`
     * 
     */
    public Optional<String> serverCertificate() {
        return Optional.ofNullable(this.serverCertificate);
    }
    /**
     * @return This parameter indicates whether or not enable https. Valid values are `on` and `off`. Default value is `on`.
     * 
     */
    public Optional<String> serverCertificateStatus() {
        return Optional.ofNullable(this.serverCertificateStatus);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DomainNewCertificateConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String certName;
        private @Nullable String certType;
        private @Nullable String forceSet;
        private @Nullable String privateKey;
        private @Nullable String serverCertificate;
        private @Nullable String serverCertificateStatus;
        public Builder() {}
        public Builder(DomainNewCertificateConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.certName = defaults.certName;
    	      this.certType = defaults.certType;
    	      this.forceSet = defaults.forceSet;
    	      this.privateKey = defaults.privateKey;
    	      this.serverCertificate = defaults.serverCertificate;
    	      this.serverCertificateStatus = defaults.serverCertificateStatus;
        }

        @CustomType.Setter
        public Builder certName(@Nullable String certName) {
            this.certName = certName;
            return this;
        }
        @CustomType.Setter
        public Builder certType(@Nullable String certType) {
            this.certType = certType;
            return this;
        }
        @CustomType.Setter
        public Builder forceSet(@Nullable String forceSet) {
            this.forceSet = forceSet;
            return this;
        }
        @CustomType.Setter
        public Builder privateKey(@Nullable String privateKey) {
            this.privateKey = privateKey;
            return this;
        }
        @CustomType.Setter
        public Builder serverCertificate(@Nullable String serverCertificate) {
            this.serverCertificate = serverCertificate;
            return this;
        }
        @CustomType.Setter
        public Builder serverCertificateStatus(@Nullable String serverCertificateStatus) {
            this.serverCertificateStatus = serverCertificateStatus;
            return this;
        }
        public DomainNewCertificateConfig build() {
            final var o = new DomainNewCertificateConfig();
            o.certName = certName;
            o.certType = certType;
            o.forceSet = forceSet;
            o.privateKey = privateKey;
            o.serverCertificate = serverCertificate;
            o.serverCertificateStatus = serverCertificateStatus;
            return o;
        }
    }
}
