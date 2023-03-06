// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eventbridge.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SlrState extends com.pulumi.resources.ResourceArgs {

    public static final SlrState Empty = new SlrState();

    @Import(name="productName")
    private @Nullable Output<String> productName;

    public Optional<Output<String>> productName() {
        return Optional.ofNullable(this.productName);
    }

    private SlrState() {}

    private SlrState(SlrState $) {
        this.productName = $.productName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SlrState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SlrState $;

        public Builder() {
            $ = new SlrState();
        }

        public Builder(SlrState defaults) {
            $ = new SlrState(Objects.requireNonNull(defaults));
        }

        public Builder productName(@Nullable Output<String> productName) {
            $.productName = productName;
            return this;
        }

        public Builder productName(String productName) {
            return productName(Output.of(productName));
        }

        public SlrState build() {
            return $;
        }
    }

}