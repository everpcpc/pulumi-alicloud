// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ram;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SamlProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final SamlProviderArgs Empty = new SamlProviderArgs();

    /**
     * The description of SAML Provider.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of SAML Provider.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The metadata file, which is Base64 encoded. The file is provided by an IdP that supports SAML 2.0.
     * 
     */
    @Import(name="encodedsamlMetadataDocument")
    private @Nullable Output<String> encodedsamlMetadataDocument;

    /**
     * @return The metadata file, which is Base64 encoded. The file is provided by an IdP that supports SAML 2.0.
     * 
     */
    public Optional<Output<String>> encodedsamlMetadataDocument() {
        return Optional.ofNullable(this.encodedsamlMetadataDocument);
    }

    /**
     * The name of SAML Provider.
     * 
     */
    @Import(name="samlProviderName", required=true)
    private Output<String> samlProviderName;

    /**
     * @return The name of SAML Provider.
     * 
     */
    public Output<String> samlProviderName() {
        return this.samlProviderName;
    }

    private SamlProviderArgs() {}

    private SamlProviderArgs(SamlProviderArgs $) {
        this.description = $.description;
        this.encodedsamlMetadataDocument = $.encodedsamlMetadataDocument;
        this.samlProviderName = $.samlProviderName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SamlProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SamlProviderArgs $;

        public Builder() {
            $ = new SamlProviderArgs();
        }

        public Builder(SamlProviderArgs defaults) {
            $ = new SamlProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of SAML Provider.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of SAML Provider.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param encodedsamlMetadataDocument The metadata file, which is Base64 encoded. The file is provided by an IdP that supports SAML 2.0.
         * 
         * @return builder
         * 
         */
        public Builder encodedsamlMetadataDocument(@Nullable Output<String> encodedsamlMetadataDocument) {
            $.encodedsamlMetadataDocument = encodedsamlMetadataDocument;
            return this;
        }

        /**
         * @param encodedsamlMetadataDocument The metadata file, which is Base64 encoded. The file is provided by an IdP that supports SAML 2.0.
         * 
         * @return builder
         * 
         */
        public Builder encodedsamlMetadataDocument(String encodedsamlMetadataDocument) {
            return encodedsamlMetadataDocument(Output.of(encodedsamlMetadataDocument));
        }

        /**
         * @param samlProviderName The name of SAML Provider.
         * 
         * @return builder
         * 
         */
        public Builder samlProviderName(Output<String> samlProviderName) {
            $.samlProviderName = samlProviderName;
            return this;
        }

        /**
         * @param samlProviderName The name of SAML Provider.
         * 
         * @return builder
         * 
         */
        public Builder samlProviderName(String samlProviderName) {
            return samlProviderName(Output.of(samlProviderName));
        }

        public SamlProviderArgs build() {
            $.samlProviderName = Objects.requireNonNull($.samlProviderName, "expected parameter 'samlProviderName' to be non-null");
            return $;
        }
    }

}
