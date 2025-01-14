// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudsso;

import com.pulumi.alicloud.cloudsso.inputs.DirectorySamlIdentityProviderConfigurationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DirectoryArgs extends com.pulumi.resources.ResourceArgs {

    public static final DirectoryArgs Empty = new DirectoryArgs();

    /**
     * The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
     * 
     */
    @Import(name="directoryName")
    private @Nullable Output<String> directoryName;

    /**
     * @return The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
     * 
     */
    public Optional<Output<String>> directoryName() {
        return Optional.ofNullable(this.directoryName);
    }

    /**
     * The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Enabled`.
     * 
     */
    @Import(name="mfaAuthenticationStatus")
    private @Nullable Output<String> mfaAuthenticationStatus;

    /**
     * @return The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Enabled`.
     * 
     */
    public Optional<Output<String>> mfaAuthenticationStatus() {
        return Optional.ofNullable(this.mfaAuthenticationStatus);
    }

    /**
     * The saml identity provider configuration.
     * 
     */
    @Import(name="samlIdentityProviderConfiguration")
    private @Nullable Output<DirectorySamlIdentityProviderConfigurationArgs> samlIdentityProviderConfiguration;

    /**
     * @return The saml identity provider configuration.
     * 
     */
    public Optional<Output<DirectorySamlIdentityProviderConfigurationArgs>> samlIdentityProviderConfiguration() {
        return Optional.ofNullable(this.samlIdentityProviderConfiguration);
    }

    /**
     * The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
     * 
     */
    @Import(name="scimSynchronizationStatus")
    private @Nullable Output<String> scimSynchronizationStatus;

    /**
     * @return The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
     * 
     */
    public Optional<Output<String>> scimSynchronizationStatus() {
        return Optional.ofNullable(this.scimSynchronizationStatus);
    }

    private DirectoryArgs() {}

    private DirectoryArgs(DirectoryArgs $) {
        this.directoryName = $.directoryName;
        this.mfaAuthenticationStatus = $.mfaAuthenticationStatus;
        this.samlIdentityProviderConfiguration = $.samlIdentityProviderConfiguration;
        this.scimSynchronizationStatus = $.scimSynchronizationStatus;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DirectoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DirectoryArgs $;

        public Builder() {
            $ = new DirectoryArgs();
        }

        public Builder(DirectoryArgs defaults) {
            $ = new DirectoryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param directoryName The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
         * 
         * @return builder
         * 
         */
        public Builder directoryName(@Nullable Output<String> directoryName) {
            $.directoryName = directoryName;
            return this;
        }

        /**
         * @param directoryName The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
         * 
         * @return builder
         * 
         */
        public Builder directoryName(String directoryName) {
            return directoryName(Output.of(directoryName));
        }

        /**
         * @param mfaAuthenticationStatus The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Enabled`.
         * 
         * @return builder
         * 
         */
        public Builder mfaAuthenticationStatus(@Nullable Output<String> mfaAuthenticationStatus) {
            $.mfaAuthenticationStatus = mfaAuthenticationStatus;
            return this;
        }

        /**
         * @param mfaAuthenticationStatus The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Enabled`.
         * 
         * @return builder
         * 
         */
        public Builder mfaAuthenticationStatus(String mfaAuthenticationStatus) {
            return mfaAuthenticationStatus(Output.of(mfaAuthenticationStatus));
        }

        /**
         * @param samlIdentityProviderConfiguration The saml identity provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder samlIdentityProviderConfiguration(@Nullable Output<DirectorySamlIdentityProviderConfigurationArgs> samlIdentityProviderConfiguration) {
            $.samlIdentityProviderConfiguration = samlIdentityProviderConfiguration;
            return this;
        }

        /**
         * @param samlIdentityProviderConfiguration The saml identity provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder samlIdentityProviderConfiguration(DirectorySamlIdentityProviderConfigurationArgs samlIdentityProviderConfiguration) {
            return samlIdentityProviderConfiguration(Output.of(samlIdentityProviderConfiguration));
        }

        /**
         * @param scimSynchronizationStatus The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
         * 
         * @return builder
         * 
         */
        public Builder scimSynchronizationStatus(@Nullable Output<String> scimSynchronizationStatus) {
            $.scimSynchronizationStatus = scimSynchronizationStatus;
            return this;
        }

        /**
         * @param scimSynchronizationStatus The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
         * 
         * @return builder
         * 
         */
        public Builder scimSynchronizationStatus(String scimSynchronizationStatus) {
            return scimSynchronizationStatus(Output.of(scimSynchronizationStatus));
        }

        public DirectoryArgs build() {
            return $;
        }
    }

}
