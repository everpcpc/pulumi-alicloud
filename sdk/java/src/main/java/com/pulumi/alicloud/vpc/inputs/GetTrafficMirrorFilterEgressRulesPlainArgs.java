// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTrafficMirrorFilterEgressRulesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTrafficMirrorFilterEgressRulesPlainArgs Empty = new GetTrafficMirrorFilterEgressRulesPlainArgs();

    /**
     * A list of Traffic Mirror Filter Egress Rule IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Traffic Mirror Filter Egress Rule IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
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
     * The status of the resource. Valid values:`Creating`, `Created`, `Modifying` and `Deleting`.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the resource. Valid values:`Creating`, `Created`, `Modifying` and `Deleting`.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The ID of the filter associated with the outbound rule.
     * 
     */
    @Import(name="trafficMirrorFilterId", required=true)
    private String trafficMirrorFilterId;

    /**
     * @return The ID of the filter associated with the outbound rule.
     * 
     */
    public String trafficMirrorFilterId() {
        return this.trafficMirrorFilterId;
    }

    private GetTrafficMirrorFilterEgressRulesPlainArgs() {}

    private GetTrafficMirrorFilterEgressRulesPlainArgs(GetTrafficMirrorFilterEgressRulesPlainArgs $) {
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.status = $.status;
        this.trafficMirrorFilterId = $.trafficMirrorFilterId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTrafficMirrorFilterEgressRulesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTrafficMirrorFilterEgressRulesPlainArgs $;

        public Builder() {
            $ = new GetTrafficMirrorFilterEgressRulesPlainArgs();
        }

        public Builder(GetTrafficMirrorFilterEgressRulesPlainArgs defaults) {
            $ = new GetTrafficMirrorFilterEgressRulesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Traffic Mirror Filter Egress Rule IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Traffic Mirror Filter Egress Rule IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
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
         * @param status The status of the resource. Valid values:`Creating`, `Created`, `Modifying` and `Deleting`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param trafficMirrorFilterId The ID of the filter associated with the outbound rule.
         * 
         * @return builder
         * 
         */
        public Builder trafficMirrorFilterId(String trafficMirrorFilterId) {
            $.trafficMirrorFilterId = trafficMirrorFilterId;
            return this;
        }

        public GetTrafficMirrorFilterEgressRulesPlainArgs build() {
            $.trafficMirrorFilterId = Objects.requireNonNull($.trafficMirrorFilterId, "expected parameter 'trafficMirrorFilterId' to be non-null");
            return $;
        }
    }

}
