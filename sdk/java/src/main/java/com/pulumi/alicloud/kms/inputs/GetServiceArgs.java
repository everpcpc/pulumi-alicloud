// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.kms.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetServiceArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetServiceArgs Empty = new GetServiceArgs();

    /**
     * Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: &#34;On&#34; or &#34;Off&#34;. Default to &#34;Off&#34;.
     * 
     */
    @Import(name="enable")
    private @Nullable Output<String> enable;

    /**
     * @return Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: &#34;On&#34; or &#34;Off&#34;. Default to &#34;Off&#34;.
     * 
     */
    public Optional<Output<String>> enable() {
        return Optional.ofNullable(this.enable);
    }

    private GetServiceArgs() {}

    private GetServiceArgs(GetServiceArgs $) {
        this.enable = $.enable;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServiceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServiceArgs $;

        public Builder() {
            $ = new GetServiceArgs();
        }

        public Builder(GetServiceArgs defaults) {
            $ = new GetServiceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enable Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: &#34;On&#34; or &#34;Off&#34;. Default to &#34;Off&#34;.
         * 
         * @return builder
         * 
         */
        public Builder enable(@Nullable Output<String> enable) {
            $.enable = enable;
            return this;
        }

        /**
         * @param enable Setting the value to `On` to enable the service. If has been enabled, return the result. Valid values: &#34;On&#34; or &#34;Off&#34;. Default to &#34;Off&#34;.
         * 
         * @return builder
         * 
         */
        public Builder enable(String enable) {
            return enable(Output.of(enable));
        }

        public GetServiceArgs build() {
            return $;
        }
    }

}