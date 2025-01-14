// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.marketplace.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetProductArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetProductArgs Empty = new GetProductArgs();

    /**
     * A available region id used to filter market place Ecs images.
     * 
     */
    @Import(name="availableRegion")
    private @Nullable Output<String> availableRegion;

    /**
     * @return A available region id used to filter market place Ecs images.
     * 
     */
    public Optional<Output<String>> availableRegion() {
        return Optional.ofNullable(this.availableRegion);
    }

    /**
     * The product code of the market product.
     * 
     */
    @Import(name="productCode", required=true)
    private Output<String> productCode;

    /**
     * @return The product code of the market product.
     * 
     */
    public Output<String> productCode() {
        return this.productCode;
    }

    private GetProductArgs() {}

    private GetProductArgs(GetProductArgs $) {
        this.availableRegion = $.availableRegion;
        this.productCode = $.productCode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProductArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProductArgs $;

        public Builder() {
            $ = new GetProductArgs();
        }

        public Builder(GetProductArgs defaults) {
            $ = new GetProductArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param availableRegion A available region id used to filter market place Ecs images.
         * 
         * @return builder
         * 
         */
        public Builder availableRegion(@Nullable Output<String> availableRegion) {
            $.availableRegion = availableRegion;
            return this;
        }

        /**
         * @param availableRegion A available region id used to filter market place Ecs images.
         * 
         * @return builder
         * 
         */
        public Builder availableRegion(String availableRegion) {
            return availableRegion(Output.of(availableRegion));
        }

        /**
         * @param productCode The product code of the market product.
         * 
         * @return builder
         * 
         */
        public Builder productCode(Output<String> productCode) {
            $.productCode = productCode;
            return this;
        }

        /**
         * @param productCode The product code of the market product.
         * 
         * @return builder
         * 
         */
        public Builder productCode(String productCode) {
            return productCode(Output.of(productCode));
        }

        public GetProductArgs build() {
            $.productCode = Objects.requireNonNull($.productCode, "expected parameter 'productCode' to be non-null");
            return $;
        }
    }

}
