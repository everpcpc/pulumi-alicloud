// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb.outputs;

import com.pulumi.alicloud.slb.outputs.GetTlsCipherPoliciesPolicyRelateListener;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetTlsCipherPoliciesPolicy {
    /**
     * @return The encryption algorithms supported. It depends on the value of `tls_versions`.
     * 
     */
    private List<String> ciphers;
    /**
     * @return The creation time timestamp.
     * 
     */
    private String createTime;
    /**
     * @return The ID of the Tls Cipher Policy.
     * 
     */
    private String id;
    /**
     * @return Array of Relate Listeners.
     * 
     */
    private List<GetTlsCipherPoliciesPolicyRelateListener> relateListeners;
    /**
     * @return TLS policy instance state.
     * 
     */
    private String status;
    /**
     * @return The ID of TLS cipher policy.
     * 
     */
    private String tlsCipherPolicyId;
    /**
     * @return TLS policy name. Length is from 2 to 128, or in both the English and Chinese characters must be with an uppercase/lowercase letter or a Chinese character and the beginning, may contain numbers, in dot `.`, underscore `_` or dash `-`.
     * 
     */
    private String tlsCipherPolicyName;
    /**
     * @return The version of TLS protocol.
     * 
     */
    private List<String> tlsVersions;

    private GetTlsCipherPoliciesPolicy() {}
    /**
     * @return The encryption algorithms supported. It depends on the value of `tls_versions`.
     * 
     */
    public List<String> ciphers() {
        return this.ciphers;
    }
    /**
     * @return The creation time timestamp.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The ID of the Tls Cipher Policy.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Array of Relate Listeners.
     * 
     */
    public List<GetTlsCipherPoliciesPolicyRelateListener> relateListeners() {
        return this.relateListeners;
    }
    /**
     * @return TLS policy instance state.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The ID of TLS cipher policy.
     * 
     */
    public String tlsCipherPolicyId() {
        return this.tlsCipherPolicyId;
    }
    /**
     * @return TLS policy name. Length is from 2 to 128, or in both the English and Chinese characters must be with an uppercase/lowercase letter or a Chinese character and the beginning, may contain numbers, in dot `.`, underscore `_` or dash `-`.
     * 
     */
    public String tlsCipherPolicyName() {
        return this.tlsCipherPolicyName;
    }
    /**
     * @return The version of TLS protocol.
     * 
     */
    public List<String> tlsVersions() {
        return this.tlsVersions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTlsCipherPoliciesPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> ciphers;
        private String createTime;
        private String id;
        private List<GetTlsCipherPoliciesPolicyRelateListener> relateListeners;
        private String status;
        private String tlsCipherPolicyId;
        private String tlsCipherPolicyName;
        private List<String> tlsVersions;
        public Builder() {}
        public Builder(GetTlsCipherPoliciesPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ciphers = defaults.ciphers;
    	      this.createTime = defaults.createTime;
    	      this.id = defaults.id;
    	      this.relateListeners = defaults.relateListeners;
    	      this.status = defaults.status;
    	      this.tlsCipherPolicyId = defaults.tlsCipherPolicyId;
    	      this.tlsCipherPolicyName = defaults.tlsCipherPolicyName;
    	      this.tlsVersions = defaults.tlsVersions;
        }

        @CustomType.Setter
        public Builder ciphers(List<String> ciphers) {
            this.ciphers = Objects.requireNonNull(ciphers);
            return this;
        }
        public Builder ciphers(String... ciphers) {
            return ciphers(List.of(ciphers));
        }
        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder relateListeners(List<GetTlsCipherPoliciesPolicyRelateListener> relateListeners) {
            this.relateListeners = Objects.requireNonNull(relateListeners);
            return this;
        }
        public Builder relateListeners(GetTlsCipherPoliciesPolicyRelateListener... relateListeners) {
            return relateListeners(List.of(relateListeners));
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder tlsCipherPolicyId(String tlsCipherPolicyId) {
            this.tlsCipherPolicyId = Objects.requireNonNull(tlsCipherPolicyId);
            return this;
        }
        @CustomType.Setter
        public Builder tlsCipherPolicyName(String tlsCipherPolicyName) {
            this.tlsCipherPolicyName = Objects.requireNonNull(tlsCipherPolicyName);
            return this;
        }
        @CustomType.Setter
        public Builder tlsVersions(List<String> tlsVersions) {
            this.tlsVersions = Objects.requireNonNull(tlsVersions);
            return this;
        }
        public Builder tlsVersions(String... tlsVersions) {
            return tlsVersions(List.of(tlsVersions));
        }
        public GetTlsCipherPoliciesPolicy build() {
            final var o = new GetTlsCipherPoliciesPolicy();
            o.ciphers = ciphers;
            o.createTime = createTime;
            o.id = id;
            o.relateListeners = relateListeners;
            o.status = status;
            o.tlsCipherPolicyId = tlsCipherPolicyId;
            o.tlsCipherPolicyName = tlsCipherPolicyName;
            o.tlsVersions = tlsVersions;
            return o;
        }
    }
}