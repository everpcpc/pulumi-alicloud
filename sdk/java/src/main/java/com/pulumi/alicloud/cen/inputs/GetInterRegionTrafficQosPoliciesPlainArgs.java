// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetInterRegionTrafficQosPoliciesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInterRegionTrafficQosPoliciesPlainArgs Empty = new GetInterRegionTrafficQosPoliciesPlainArgs();

    /**
     * A list of Inter Region Traffic Qos Policy IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Inter Region Traffic Qos Policy IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Inter Region Traffic Qos Policy name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Inter Region Traffic Qos Policy name.
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
     * The status of the traffic scheduling policy. Valid Value: `Creating`, `Active`, `Modifying`, `Deleting`, `Deleted`.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the traffic scheduling policy. Valid Value: `Creating`, `Active`, `Modifying`, `Deleting`, `Deleted`.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The description of the QoS policy.
     * 
     */
    @Import(name="trafficQosPolicyDescription")
    private @Nullable String trafficQosPolicyDescription;

    /**
     * @return The description of the QoS policy.
     * 
     */
    public Optional<String> trafficQosPolicyDescription() {
        return Optional.ofNullable(this.trafficQosPolicyDescription);
    }

    /**
     * The ID of the QoS policy.
     * 
     */
    @Import(name="trafficQosPolicyId")
    private @Nullable String trafficQosPolicyId;

    /**
     * @return The ID of the QoS policy.
     * 
     */
    public Optional<String> trafficQosPolicyId() {
        return Optional.ofNullable(this.trafficQosPolicyId);
    }

    /**
     * The name of the QoS policy.
     * 
     */
    @Import(name="trafficQosPolicyName")
    private @Nullable String trafficQosPolicyName;

    /**
     * @return The name of the QoS policy.
     * 
     */
    public Optional<String> trafficQosPolicyName() {
        return Optional.ofNullable(this.trafficQosPolicyName);
    }

    /**
     * The ID of the inter-region connection.
     * 
     */
    @Import(name="transitRouterAttachmentId", required=true)
    private String transitRouterAttachmentId;

    /**
     * @return The ID of the inter-region connection.
     * 
     */
    public String transitRouterAttachmentId() {
        return this.transitRouterAttachmentId;
    }

    /**
     * The ID of the transit router.
     * 
     */
    @Import(name="transitRouterId", required=true)
    private String transitRouterId;

    /**
     * @return The ID of the transit router.
     * 
     */
    public String transitRouterId() {
        return this.transitRouterId;
    }

    private GetInterRegionTrafficQosPoliciesPlainArgs() {}

    private GetInterRegionTrafficQosPoliciesPlainArgs(GetInterRegionTrafficQosPoliciesPlainArgs $) {
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.status = $.status;
        this.trafficQosPolicyDescription = $.trafficQosPolicyDescription;
        this.trafficQosPolicyId = $.trafficQosPolicyId;
        this.trafficQosPolicyName = $.trafficQosPolicyName;
        this.transitRouterAttachmentId = $.transitRouterAttachmentId;
        this.transitRouterId = $.transitRouterId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInterRegionTrafficQosPoliciesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInterRegionTrafficQosPoliciesPlainArgs $;

        public Builder() {
            $ = new GetInterRegionTrafficQosPoliciesPlainArgs();
        }

        public Builder(GetInterRegionTrafficQosPoliciesPlainArgs defaults) {
            $ = new GetInterRegionTrafficQosPoliciesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Inter Region Traffic Qos Policy IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Inter Region Traffic Qos Policy IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Inter Region Traffic Qos Policy name.
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
         * @param status The status of the traffic scheduling policy. Valid Value: `Creating`, `Active`, `Modifying`, `Deleting`, `Deleted`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param trafficQosPolicyDescription The description of the QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder trafficQosPolicyDescription(@Nullable String trafficQosPolicyDescription) {
            $.trafficQosPolicyDescription = trafficQosPolicyDescription;
            return this;
        }

        /**
         * @param trafficQosPolicyId The ID of the QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder trafficQosPolicyId(@Nullable String trafficQosPolicyId) {
            $.trafficQosPolicyId = trafficQosPolicyId;
            return this;
        }

        /**
         * @param trafficQosPolicyName The name of the QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder trafficQosPolicyName(@Nullable String trafficQosPolicyName) {
            $.trafficQosPolicyName = trafficQosPolicyName;
            return this;
        }

        /**
         * @param transitRouterAttachmentId The ID of the inter-region connection.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterAttachmentId(String transitRouterAttachmentId) {
            $.transitRouterAttachmentId = transitRouterAttachmentId;
            return this;
        }

        /**
         * @param transitRouterId The ID of the transit router.
         * 
         * @return builder
         * 
         */
        public Builder transitRouterId(String transitRouterId) {
            $.transitRouterId = transitRouterId;
            return this;
        }

        public GetInterRegionTrafficQosPoliciesPlainArgs build() {
            $.transitRouterAttachmentId = Objects.requireNonNull($.transitRouterAttachmentId, "expected parameter 'transitRouterAttachmentId' to be non-null");
            $.transitRouterId = Objects.requireNonNull($.transitRouterId, "expected parameter 'transitRouterId' to be non-null");
            return $;
        }
    }

}
