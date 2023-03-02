// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.bss.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetOpenApiPricingModulesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOpenApiPricingModulesPlainArgs Empty = new GetOpenApiPricingModulesPlainArgs();

    @Import(name="ids")
    private @Nullable List<String> ids;

    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Property name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Property name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    @Import(name="outputFile")
    private @Nullable String outputFile;

    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The product code.
     * 
     */
    @Import(name="productCode", required=true)
    private String productCode;

    /**
     * @return The product code.
     * 
     */
    public String productCode() {
        return this.productCode;
    }

    /**
     * The product type.
     * 
     */
    @Import(name="productType")
    private @Nullable String productType;

    /**
     * @return The product type.
     * 
     */
    public Optional<String> productType() {
        return Optional.ofNullable(this.productType);
    }

    /**
     * Subscription type. Value:
     * * Subscription: Prepaid.
     * * PayAsYouGo: postpaid.
     * 
     */
    @Import(name="subscriptionType", required=true)
    private String subscriptionType;

    /**
     * @return Subscription type. Value:
     * * Subscription: Prepaid.
     * * PayAsYouGo: postpaid.
     * 
     */
    public String subscriptionType() {
        return this.subscriptionType;
    }

    private GetOpenApiPricingModulesPlainArgs() {}

    private GetOpenApiPricingModulesPlainArgs(GetOpenApiPricingModulesPlainArgs $) {
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.productCode = $.productCode;
        this.productType = $.productType;
        this.subscriptionType = $.subscriptionType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOpenApiPricingModulesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOpenApiPricingModulesPlainArgs $;

        public Builder() {
            $ = new GetOpenApiPricingModulesPlainArgs();
        }

        public Builder(GetOpenApiPricingModulesPlainArgs defaults) {
            $ = new GetOpenApiPricingModulesPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Property name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param productCode The product code.
         * 
         * @return builder
         * 
         */
        public Builder productCode(String productCode) {
            $.productCode = productCode;
            return this;
        }

        /**
         * @param productType The product type.
         * 
         * @return builder
         * 
         */
        public Builder productType(@Nullable String productType) {
            $.productType = productType;
            return this;
        }

        /**
         * @param subscriptionType Subscription type. Value:
         * * Subscription: Prepaid.
         * * PayAsYouGo: postpaid.
         * 
         * @return builder
         * 
         */
        public Builder subscriptionType(String subscriptionType) {
            $.subscriptionType = subscriptionType;
            return this;
        }

        public GetOpenApiPricingModulesPlainArgs build() {
            $.productCode = Objects.requireNonNull($.productCode, "expected parameter 'productCode' to be non-null");
            $.subscriptionType = Objects.requireNonNull($.subscriptionType, "expected parameter 'subscriptionType' to be non-null");
            return $;
        }
    }

}
