// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudsso.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DirectorySamlIdentityProviderConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final DirectorySamlIdentityProviderConfigurationArgs Empty = new DirectorySamlIdentityProviderConfigurationArgs();

    /**
     * Base64 encoded IdP metadata document. **NOTE:** If the IdP Metadata has been uploaded, no update will be made if this parameter is not specified, otherwise the update will be made according to the parameter content. If IdP Metadata has not been uploaded, and the parameter `sso_status` is `Enabled`, this parameter must be provided. If the IdP Metadata has not been uploaded, and the parameter `sso_status` is `Disabled`, this parameter can be omitted, and the IdP Metadata will remain empty.
     * 
     */
    @Import(name="encodedMetadataDocument")
    private @Nullable Output<String> encodedMetadataDocument;

    /**
     * @return Base64 encoded IdP metadata document. **NOTE:** If the IdP Metadata has been uploaded, no update will be made if this parameter is not specified, otherwise the update will be made according to the parameter content. If IdP Metadata has not been uploaded, and the parameter `sso_status` is `Enabled`, this parameter must be provided. If the IdP Metadata has not been uploaded, and the parameter `sso_status` is `Disabled`, this parameter can be omitted, and the IdP Metadata will remain empty.
     * 
     */
    public Optional<Output<String>> encodedMetadataDocument() {
        return Optional.ofNullable(this.encodedMetadataDocument);
    }

    /**
     * SAML SSO login enabled status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
     * 
     * &gt; **NOTE:** The `saml_identity_provider_configuration` will be removed automatically when the resource is deleted, please operate with caution. If there are left more configuration in the directory, please remove them before deleting the directory.
     * 
     */
    @Import(name="ssoStatus")
    private @Nullable Output<String> ssoStatus;

    /**
     * @return SAML SSO login enabled status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
     * 
     * &gt; **NOTE:** The `saml_identity_provider_configuration` will be removed automatically when the resource is deleted, please operate with caution. If there are left more configuration in the directory, please remove them before deleting the directory.
     * 
     */
    public Optional<Output<String>> ssoStatus() {
        return Optional.ofNullable(this.ssoStatus);
    }

    private DirectorySamlIdentityProviderConfigurationArgs() {}

    private DirectorySamlIdentityProviderConfigurationArgs(DirectorySamlIdentityProviderConfigurationArgs $) {
        this.encodedMetadataDocument = $.encodedMetadataDocument;
        this.ssoStatus = $.ssoStatus;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DirectorySamlIdentityProviderConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DirectorySamlIdentityProviderConfigurationArgs $;

        public Builder() {
            $ = new DirectorySamlIdentityProviderConfigurationArgs();
        }

        public Builder(DirectorySamlIdentityProviderConfigurationArgs defaults) {
            $ = new DirectorySamlIdentityProviderConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param encodedMetadataDocument Base64 encoded IdP metadata document. **NOTE:** If the IdP Metadata has been uploaded, no update will be made if this parameter is not specified, otherwise the update will be made according to the parameter content. If IdP Metadata has not been uploaded, and the parameter `sso_status` is `Enabled`, this parameter must be provided. If the IdP Metadata has not been uploaded, and the parameter `sso_status` is `Disabled`, this parameter can be omitted, and the IdP Metadata will remain empty.
         * 
         * @return builder
         * 
         */
        public Builder encodedMetadataDocument(@Nullable Output<String> encodedMetadataDocument) {
            $.encodedMetadataDocument = encodedMetadataDocument;
            return this;
        }

        /**
         * @param encodedMetadataDocument Base64 encoded IdP metadata document. **NOTE:** If the IdP Metadata has been uploaded, no update will be made if this parameter is not specified, otherwise the update will be made according to the parameter content. If IdP Metadata has not been uploaded, and the parameter `sso_status` is `Enabled`, this parameter must be provided. If the IdP Metadata has not been uploaded, and the parameter `sso_status` is `Disabled`, this parameter can be omitted, and the IdP Metadata will remain empty.
         * 
         * @return builder
         * 
         */
        public Builder encodedMetadataDocument(String encodedMetadataDocument) {
            return encodedMetadataDocument(Output.of(encodedMetadataDocument));
        }

        /**
         * @param ssoStatus SAML SSO login enabled status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
         * 
         * &gt; **NOTE:** The `saml_identity_provider_configuration` will be removed automatically when the resource is deleted, please operate with caution. If there are left more configuration in the directory, please remove them before deleting the directory.
         * 
         * @return builder
         * 
         */
        public Builder ssoStatus(@Nullable Output<String> ssoStatus) {
            $.ssoStatus = ssoStatus;
            return this;
        }

        /**
         * @param ssoStatus SAML SSO login enabled status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
         * 
         * &gt; **NOTE:** The `saml_identity_provider_configuration` will be removed automatically when the resource is deleted, please operate with caution. If there are left more configuration in the directory, please remove them before deleting the directory.
         * 
         * @return builder
         * 
         */
        public Builder ssoStatus(String ssoStatus) {
            return ssoStatus(Output.of(ssoStatus));
        }

        public DirectorySamlIdentityProviderConfigurationArgs build() {
            return $;
        }
    }

}
