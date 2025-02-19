// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.clickhouse.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRegionsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRegionsPlainArgs Empty = new GetRegionsPlainArgs();

    /**
     * Set to true to match only the region configured in the provider. Default value: `true`.
     * 
     */
    @Import(name="current")
    private @Nullable Boolean current;

    /**
     * @return Set to true to match only the region configured in the provider. Default value: `true`.
     * 
     */
    public Optional<Boolean> current() {
        return Optional.ofNullable(this.current);
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
     * The Region ID.
     * 
     */
    @Import(name="regionId")
    private @Nullable String regionId;

    /**
     * @return The Region ID.
     * 
     */
    public Optional<String> regionId() {
        return Optional.ofNullable(this.regionId);
    }

    private GetRegionsPlainArgs() {}

    private GetRegionsPlainArgs(GetRegionsPlainArgs $) {
        this.current = $.current;
        this.outputFile = $.outputFile;
        this.regionId = $.regionId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRegionsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRegionsPlainArgs $;

        public Builder() {
            $ = new GetRegionsPlainArgs();
        }

        public Builder(GetRegionsPlainArgs defaults) {
            $ = new GetRegionsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param current Set to true to match only the region configured in the provider. Default value: `true`.
         * 
         * @return builder
         * 
         */
        public Builder current(@Nullable Boolean current) {
            $.current = current;
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
         * @param regionId The Region ID.
         * 
         * @return builder
         * 
         */
        public Builder regionId(@Nullable String regionId) {
            $.regionId = regionId;
            return this;
        }

        public GetRegionsPlainArgs build() {
            return $;
        }
    }

}
