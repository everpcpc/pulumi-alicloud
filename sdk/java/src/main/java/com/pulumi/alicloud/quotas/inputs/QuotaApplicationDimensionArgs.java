// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.quotas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QuotaApplicationDimensionArgs extends com.pulumi.resources.ResourceArgs {

    public static final QuotaApplicationDimensionArgs Empty = new QuotaApplicationDimensionArgs();

    /**
     * The key of dimensions.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return The key of dimensions.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * The value of dimensions.
     * 
     */
    @Import(name="value")
    private @Nullable Output<String> value;

    /**
     * @return The value of dimensions.
     * 
     */
    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    private QuotaApplicationDimensionArgs() {}

    private QuotaApplicationDimensionArgs(QuotaApplicationDimensionArgs $) {
        this.key = $.key;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QuotaApplicationDimensionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QuotaApplicationDimensionArgs $;

        public Builder() {
            $ = new QuotaApplicationDimensionArgs();
        }

        public Builder(QuotaApplicationDimensionArgs defaults) {
            $ = new QuotaApplicationDimensionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param key The key of dimensions.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The key of dimensions.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param value The value of dimensions.
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value The value of dimensions.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public QuotaApplicationDimensionArgs build() {
            return $;
        }
    }

}