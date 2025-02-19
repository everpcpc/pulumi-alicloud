// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.marketplace;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OrderArgs extends com.pulumi.resources.ResourceArgs {

    public static final OrderArgs Empty = new OrderArgs();

    /**
     * Service providers customize additional components.
     * 
     */
    @Import(name="components")
    private @Nullable Output<Map<String,Object>> components;

    /**
     * @return Service providers customize additional components.
     * 
     */
    public Optional<Output<Map<String,Object>>> components() {
        return Optional.ofNullable(this.components);
    }

    /**
     * The coupon id of the market product.
     * 
     */
    @Import(name="couponId")
    private @Nullable Output<String> couponId;

    /**
     * @return The coupon id of the market product.
     * 
     */
    public Optional<Output<String>> couponId() {
        return Optional.ofNullable(this.couponId);
    }

    /**
     * The number of purchase cycles.
     * 
     */
    @Import(name="duration")
    private @Nullable Output<Integer> duration;

    /**
     * @return The number of purchase cycles.
     * 
     */
    public Optional<Output<Integer>> duration() {
        return Optional.ofNullable(this.duration);
    }

    /**
     * The package version of the market product.
     * 
     */
    @Import(name="packageVersion", required=true)
    private Output<String> packageVersion;

    /**
     * @return The package version of the market product.
     * 
     */
    public Output<String> packageVersion() {
        return this.packageVersion;
    }

    /**
     * Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
     * 
     */
    @Import(name="payType")
    private @Nullable Output<String> payType;

    /**
     * @return Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
     * 
     */
    public Optional<Output<String>> payType() {
        return Optional.ofNullable(this.payType);
    }

    /**
     * The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
     * 
     */
    @Import(name="pricingCycle", required=true)
    private Output<String> pricingCycle;

    /**
     * @return The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
     * 
     */
    public Output<String> pricingCycle() {
        return this.pricingCycle;
    }

    /**
     * The product_code of market place product.
     * 
     */
    @Import(name="productCode", required=true)
    private Output<String> productCode;

    /**
     * @return The product_code of market place product.
     * 
     */
    public Output<String> productCode() {
        return this.productCode;
    }

    /**
     * The quantity of the market product will be purchased.
     * 
     */
    @Import(name="quantity")
    private @Nullable Output<Integer> quantity;

    /**
     * @return The quantity of the market product will be purchased.
     * 
     */
    public Optional<Output<Integer>> quantity() {
        return Optional.ofNullable(this.quantity);
    }

    private OrderArgs() {}

    private OrderArgs(OrderArgs $) {
        this.components = $.components;
        this.couponId = $.couponId;
        this.duration = $.duration;
        this.packageVersion = $.packageVersion;
        this.payType = $.payType;
        this.pricingCycle = $.pricingCycle;
        this.productCode = $.productCode;
        this.quantity = $.quantity;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OrderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OrderArgs $;

        public Builder() {
            $ = new OrderArgs();
        }

        public Builder(OrderArgs defaults) {
            $ = new OrderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param components Service providers customize additional components.
         * 
         * @return builder
         * 
         */
        public Builder components(@Nullable Output<Map<String,Object>> components) {
            $.components = components;
            return this;
        }

        /**
         * @param components Service providers customize additional components.
         * 
         * @return builder
         * 
         */
        public Builder components(Map<String,Object> components) {
            return components(Output.of(components));
        }

        /**
         * @param couponId The coupon id of the market product.
         * 
         * @return builder
         * 
         */
        public Builder couponId(@Nullable Output<String> couponId) {
            $.couponId = couponId;
            return this;
        }

        /**
         * @param couponId The coupon id of the market product.
         * 
         * @return builder
         * 
         */
        public Builder couponId(String couponId) {
            return couponId(Output.of(couponId));
        }

        /**
         * @param duration The number of purchase cycles.
         * 
         * @return builder
         * 
         */
        public Builder duration(@Nullable Output<Integer> duration) {
            $.duration = duration;
            return this;
        }

        /**
         * @param duration The number of purchase cycles.
         * 
         * @return builder
         * 
         */
        public Builder duration(Integer duration) {
            return duration(Output.of(duration));
        }

        /**
         * @param packageVersion The package version of the market product.
         * 
         * @return builder
         * 
         */
        public Builder packageVersion(Output<String> packageVersion) {
            $.packageVersion = packageVersion;
            return this;
        }

        /**
         * @param packageVersion The package version of the market product.
         * 
         * @return builder
         * 
         */
        public Builder packageVersion(String packageVersion) {
            return packageVersion(Output.of(packageVersion));
        }

        /**
         * @param payType Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
         * 
         * @return builder
         * 
         */
        public Builder payType(@Nullable Output<String> payType) {
            $.payType = payType;
            return this;
        }

        /**
         * @param payType Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`.
         * 
         * @return builder
         * 
         */
        public Builder payType(String payType) {
            return payType(Output.of(payType));
        }

        /**
         * @param pricingCycle The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
         * 
         * @return builder
         * 
         */
        public Builder pricingCycle(Output<String> pricingCycle) {
            $.pricingCycle = pricingCycle;
            return this;
        }

        /**
         * @param pricingCycle The purchase cycle of the product, valid values are `Day`, `Month` and `Year`.
         * 
         * @return builder
         * 
         */
        public Builder pricingCycle(String pricingCycle) {
            return pricingCycle(Output.of(pricingCycle));
        }

        /**
         * @param productCode The product_code of market place product.
         * 
         * @return builder
         * 
         */
        public Builder productCode(Output<String> productCode) {
            $.productCode = productCode;
            return this;
        }

        /**
         * @param productCode The product_code of market place product.
         * 
         * @return builder
         * 
         */
        public Builder productCode(String productCode) {
            return productCode(Output.of(productCode));
        }

        /**
         * @param quantity The quantity of the market product will be purchased.
         * 
         * @return builder
         * 
         */
        public Builder quantity(@Nullable Output<Integer> quantity) {
            $.quantity = quantity;
            return this;
        }

        /**
         * @param quantity The quantity of the market product will be purchased.
         * 
         * @return builder
         * 
         */
        public Builder quantity(Integer quantity) {
            return quantity(Output.of(quantity));
        }

        public OrderArgs build() {
            $.packageVersion = Objects.requireNonNull($.packageVersion, "expected parameter 'packageVersion' to be non-null");
            $.pricingCycle = Objects.requireNonNull($.pricingCycle, "expected parameter 'pricingCycle' to be non-null");
            $.productCode = Objects.requireNonNull($.productCode, "expected parameter 'productCode' to be non-null");
            return $;
        }
    }

}
