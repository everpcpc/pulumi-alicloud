// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudsso.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetScimServerCredentialsCredential {
    /**
     * @return The CreateTime of the resource.
     * 
     */
    private String createTime;
    /**
     * @return The CredentialId of the resource.
     * 
     */
    private String credentialId;
    /**
     * @return The CredentialSecret of the resource.
     * 
     */
    private String credentialSecret;
    /**
     * @return The CredentialType of the resource.
     * 
     */
    private String credentialType;
    /**
     * @return The ID of the Directory.
     * 
     */
    private String directoryId;
    /**
     * @return The ExpireTime of the resource.
     * 
     */
    private String expireTime;
    /**
     * @return The ID of the SCIM Server Credential.
     * 
     */
    private String id;
    /**
     * @return The Status of the resource. Valid values: `Disabled`, `Enabled`.
     * 
     */
    private String status;

    private GetScimServerCredentialsCredential() {}
    /**
     * @return The CreateTime of the resource.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The CredentialId of the resource.
     * 
     */
    public String credentialId() {
        return this.credentialId;
    }
    /**
     * @return The CredentialSecret of the resource.
     * 
     */
    public String credentialSecret() {
        return this.credentialSecret;
    }
    /**
     * @return The CredentialType of the resource.
     * 
     */
    public String credentialType() {
        return this.credentialType;
    }
    /**
     * @return The ID of the Directory.
     * 
     */
    public String directoryId() {
        return this.directoryId;
    }
    /**
     * @return The ExpireTime of the resource.
     * 
     */
    public String expireTime() {
        return this.expireTime;
    }
    /**
     * @return The ID of the SCIM Server Credential.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The Status of the resource. Valid values: `Disabled`, `Enabled`.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetScimServerCredentialsCredential defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createTime;
        private String credentialId;
        private String credentialSecret;
        private String credentialType;
        private String directoryId;
        private String expireTime;
        private String id;
        private String status;
        public Builder() {}
        public Builder(GetScimServerCredentialsCredential defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createTime = defaults.createTime;
    	      this.credentialId = defaults.credentialId;
    	      this.credentialSecret = defaults.credentialSecret;
    	      this.credentialType = defaults.credentialType;
    	      this.directoryId = defaults.directoryId;
    	      this.expireTime = defaults.expireTime;
    	      this.id = defaults.id;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder credentialId(String credentialId) {
            this.credentialId = Objects.requireNonNull(credentialId);
            return this;
        }
        @CustomType.Setter
        public Builder credentialSecret(String credentialSecret) {
            this.credentialSecret = Objects.requireNonNull(credentialSecret);
            return this;
        }
        @CustomType.Setter
        public Builder credentialType(String credentialType) {
            this.credentialType = Objects.requireNonNull(credentialType);
            return this;
        }
        @CustomType.Setter
        public Builder directoryId(String directoryId) {
            this.directoryId = Objects.requireNonNull(directoryId);
            return this;
        }
        @CustomType.Setter
        public Builder expireTime(String expireTime) {
            this.expireTime = Objects.requireNonNull(expireTime);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetScimServerCredentialsCredential build() {
            final var o = new GetScimServerCredentialsCredential();
            o.createTime = createTime;
            o.credentialId = credentialId;
            o.credentialSecret = credentialSecret;
            o.credentialType = credentialType;
            o.directoryId = directoryId;
            o.expireTime = expireTime;
            o.id = id;
            o.status = status;
            return o;
        }
    }
}
