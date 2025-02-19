// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.servicecatalog.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetLaunchOptionsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLaunchOptionsPlainArgs Empty = new GetLaunchOptionsPlainArgs();

    @Import(name="ids")
    private @Nullable List<String> ids;

    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by portfolio name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by portfolio name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * Product ID.
     * 
     */
    @Import(name="productId", required=true)
    private String productId;

    /**
     * @return Product ID.
     * 
     */
    public String productId() {
        return this.productId;
    }

    private GetLaunchOptionsPlainArgs() {}

    private GetLaunchOptionsPlainArgs(GetLaunchOptionsPlainArgs $) {
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.productId = $.productId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLaunchOptionsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLaunchOptionsPlainArgs $;

        public Builder() {
            $ = new GetLaunchOptionsPlainArgs();
        }

        public Builder(GetLaunchOptionsPlainArgs defaults) {
            $ = new GetLaunchOptionsPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by portfolio name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param productId Product ID.
         * 
         * @return builder
         * 
         */
        public Builder productId(String productId) {
            $.productId = productId;
            return this;
        }

        public GetLaunchOptionsPlainArgs build() {
            $.productId = Objects.requireNonNull($.productId, "expected parameter 'productId' to be non-null");
            return $;
        }
    }

}
