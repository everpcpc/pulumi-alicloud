// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.apigateway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ApiConstantParameterArgs extends com.pulumi.resources.ResourceArgs {

    public static final ApiConstantParameterArgs Empty = new ApiConstantParameterArgs();

    /**
     * The description of the api. Defaults to null.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the api. Defaults to null.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Request&#39;s parameter location; values: BODY, HEAD, QUERY, and PATH.
     * 
     */
    @Import(name="in", required=true)
    private Output<String> in;

    /**
     * @return Request&#39;s parameter location; values: BODY, HEAD, QUERY, and PATH.
     * 
     */
    public Output<String> in() {
        return this.in;
    }

    /**
     * The name of the api gateway api. Defaults to null.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the api gateway api. Defaults to null.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Constant parameter value.
     * 
     */
    @Import(name="value", required=true)
    private Output<String> value;

    /**
     * @return Constant parameter value.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    private ApiConstantParameterArgs() {}

    private ApiConstantParameterArgs(ApiConstantParameterArgs $) {
        this.description = $.description;
        this.in = $.in;
        this.name = $.name;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApiConstantParameterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApiConstantParameterArgs $;

        public Builder() {
            $ = new ApiConstantParameterArgs();
        }

        public Builder(ApiConstantParameterArgs defaults) {
            $ = new ApiConstantParameterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the api. Defaults to null.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the api. Defaults to null.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param in Request&#39;s parameter location; values: BODY, HEAD, QUERY, and PATH.
         * 
         * @return builder
         * 
         */
        public Builder in(Output<String> in) {
            $.in = in;
            return this;
        }

        /**
         * @param in Request&#39;s parameter location; values: BODY, HEAD, QUERY, and PATH.
         * 
         * @return builder
         * 
         */
        public Builder in(String in) {
            return in(Output.of(in));
        }

        /**
         * @param name The name of the api gateway api. Defaults to null.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the api gateway api. Defaults to null.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param value Constant parameter value.
         * 
         * @return builder
         * 
         */
        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value Constant parameter value.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public ApiConstantParameterArgs build() {
            $.in = Objects.requireNonNull($.in, "expected parameter 'in' to be non-null");
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.value = Objects.requireNonNull($.value, "expected parameter 'value' to be non-null");
            return $;
        }
    }

}
