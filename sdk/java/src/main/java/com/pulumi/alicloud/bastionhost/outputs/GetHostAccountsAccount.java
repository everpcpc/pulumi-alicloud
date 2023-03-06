// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.bastionhost.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetHostAccountsAccount {
    /**
     * @return Whether to set a new password.
     * 
     */
    private Boolean hasPassword;
    /**
     * @return Hosting account ID.
     * 
     */
    private String hostAccountId;
    /**
     * @return Specify the new hosting account&#39;s name, support the longest 128 characters.
     * 
     */
    private String hostAccountName;
    /**
     * @return Specifies the database where you want to create your hosting account&#39;s host ID.
     * 
     */
    private String hostId;
    /**
     * @return The ID of the Host Account.
     * 
     */
    private String id;
    /**
     * @return Specifies the database where you want to create your hosting account&#39;s host bastion host ID of.
     * 
     */
    private String instanceId;
    /**
     * @return The situation where the private keys of the fingerprint information.
     * 
     */
    private String privateKeyFingerprint;
    /**
     * @return Specify the new hosting account of the agreement name. Valid values: USING SSH and RDP.
     * 
     */
    private String protocolName;

    private GetHostAccountsAccount() {}
    /**
     * @return Whether to set a new password.
     * 
     */
    public Boolean hasPassword() {
        return this.hasPassword;
    }
    /**
     * @return Hosting account ID.
     * 
     */
    public String hostAccountId() {
        return this.hostAccountId;
    }
    /**
     * @return Specify the new hosting account&#39;s name, support the longest 128 characters.
     * 
     */
    public String hostAccountName() {
        return this.hostAccountName;
    }
    /**
     * @return Specifies the database where you want to create your hosting account&#39;s host ID.
     * 
     */
    public String hostId() {
        return this.hostId;
    }
    /**
     * @return The ID of the Host Account.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Specifies the database where you want to create your hosting account&#39;s host bastion host ID of.
     * 
     */
    public String instanceId() {
        return this.instanceId;
    }
    /**
     * @return The situation where the private keys of the fingerprint information.
     * 
     */
    public String privateKeyFingerprint() {
        return this.privateKeyFingerprint;
    }
    /**
     * @return Specify the new hosting account of the agreement name. Valid values: USING SSH and RDP.
     * 
     */
    public String protocolName() {
        return this.protocolName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetHostAccountsAccount defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean hasPassword;
        private String hostAccountId;
        private String hostAccountName;
        private String hostId;
        private String id;
        private String instanceId;
        private String privateKeyFingerprint;
        private String protocolName;
        public Builder() {}
        public Builder(GetHostAccountsAccount defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.hasPassword = defaults.hasPassword;
    	      this.hostAccountId = defaults.hostAccountId;
    	      this.hostAccountName = defaults.hostAccountName;
    	      this.hostId = defaults.hostId;
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.privateKeyFingerprint = defaults.privateKeyFingerprint;
    	      this.protocolName = defaults.protocolName;
        }

        @CustomType.Setter
        public Builder hasPassword(Boolean hasPassword) {
            this.hasPassword = Objects.requireNonNull(hasPassword);
            return this;
        }
        @CustomType.Setter
        public Builder hostAccountId(String hostAccountId) {
            this.hostAccountId = Objects.requireNonNull(hostAccountId);
            return this;
        }
        @CustomType.Setter
        public Builder hostAccountName(String hostAccountName) {
            this.hostAccountName = Objects.requireNonNull(hostAccountName);
            return this;
        }
        @CustomType.Setter
        public Builder hostId(String hostId) {
            this.hostId = Objects.requireNonNull(hostId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(String instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder privateKeyFingerprint(String privateKeyFingerprint) {
            this.privateKeyFingerprint = Objects.requireNonNull(privateKeyFingerprint);
            return this;
        }
        @CustomType.Setter
        public Builder protocolName(String protocolName) {
            this.protocolName = Objects.requireNonNull(protocolName);
            return this;
        }
        public GetHostAccountsAccount build() {
            final var o = new GetHostAccountsAccount();
            o.hasPassword = hasPassword;
            o.hostAccountId = hostAccountId;
            o.hostAccountName = hostAccountName;
            o.hostId = hostId;
            o.id = id;
            o.instanceId = instanceId;
            o.privateKeyFingerprint = privateKeyFingerprint;
            o.protocolName = protocolName;
            return o;
        }
    }
}