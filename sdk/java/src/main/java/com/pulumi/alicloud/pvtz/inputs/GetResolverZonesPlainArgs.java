// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.pvtz.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetResolverZonesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetResolverZonesPlainArgs Empty = new GetResolverZonesPlainArgs();

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
     * The status of the Zone.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the Zone.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    private GetResolverZonesPlainArgs() {}

    private GetResolverZonesPlainArgs(GetResolverZonesPlainArgs $) {
        this.outputFile = $.outputFile;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetResolverZonesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetResolverZonesPlainArgs $;

        public Builder() {
            $ = new GetResolverZonesPlainArgs();
        }

        public Builder(GetResolverZonesPlainArgs defaults) {
            $ = new GetResolverZonesPlainArgs(Objects.requireNonNull(defaults));
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
         * @param status The status of the Zone.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        public GetResolverZonesPlainArgs build() {
            return $;
        }
    }

}
