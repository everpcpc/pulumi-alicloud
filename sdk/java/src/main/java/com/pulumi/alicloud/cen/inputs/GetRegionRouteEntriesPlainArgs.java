// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRegionRouteEntriesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRegionRouteEntriesPlainArgs Empty = new GetRegionRouteEntriesPlainArgs();

    /**
     * ID of the CEN instance.
     * 
     */
    @Import(name="instanceId", required=true)
    private String instanceId;

    /**
     * @return ID of the CEN instance.
     * 
     */
    public String instanceId() {
        return this.instanceId;
    }

    @Import(name="outputFile")
    private @Nullable String outputFile;

    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * ID of the region.
     * 
     */
    @Import(name="regionId", required=true)
    private String regionId;

    /**
     * @return ID of the region.
     * 
     */
    public String regionId() {
        return this.regionId;
    }

    private GetRegionRouteEntriesPlainArgs() {}

    private GetRegionRouteEntriesPlainArgs(GetRegionRouteEntriesPlainArgs $) {
        this.instanceId = $.instanceId;
        this.outputFile = $.outputFile;
        this.regionId = $.regionId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRegionRouteEntriesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRegionRouteEntriesPlainArgs $;

        public Builder() {
            $ = new GetRegionRouteEntriesPlainArgs();
        }

        public Builder(GetRegionRouteEntriesPlainArgs defaults) {
            $ = new GetRegionRouteEntriesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param instanceId ID of the CEN instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param regionId ID of the region.
         * 
         * @return builder
         * 
         */
        public Builder regionId(String regionId) {
            $.regionId = regionId;
            return this;
        }

        public GetRegionRouteEntriesPlainArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            $.regionId = Objects.requireNonNull($.regionId, "expected parameter 'regionId' to be non-null");
            return $;
        }
    }

}
