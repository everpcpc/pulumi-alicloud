// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cdn.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetBlockedRegionsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBlockedRegionsArgs Empty = new GetBlockedRegionsArgs();

    /**
     * The language. Valid values: `zh`, `en`, `jp`.
     * 
     */
    @Import(name="language", required=true)
    private Output<String> language;

    /**
     * @return The language. Valid values: `zh`, `en`, `jp`.
     * 
     */
    public Output<String> language() {
        return this.language;
    }

    private GetBlockedRegionsArgs() {}

    private GetBlockedRegionsArgs(GetBlockedRegionsArgs $) {
        this.language = $.language;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBlockedRegionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBlockedRegionsArgs $;

        public Builder() {
            $ = new GetBlockedRegionsArgs();
        }

        public Builder(GetBlockedRegionsArgs defaults) {
            $ = new GetBlockedRegionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param language The language. Valid values: `zh`, `en`, `jp`.
         * 
         * @return builder
         * 
         */
        public Builder language(Output<String> language) {
            $.language = language;
            return this;
        }

        /**
         * @param language The language. Valid values: `zh`, `en`, `jp`.
         * 
         * @return builder
         * 
         */
        public Builder language(String language) {
            return language(Output.of(language));
        }

        public GetBlockedRegionsArgs build() {
            $.language = Objects.requireNonNull($.language, "expected parameter 'language' to be non-null");
            return $;
        }
    }

}
