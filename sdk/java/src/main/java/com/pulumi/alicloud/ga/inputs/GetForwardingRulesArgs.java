// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetForwardingRulesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetForwardingRulesArgs Empty = new GetForwardingRulesArgs();

    /**
     * The ID of the Global Accelerator instance.
     * 
     */
    @Import(name="acceleratorId", required=true)
    private Output<String> acceleratorId;

    /**
     * @return The ID of the Global Accelerator instance.
     * 
     */
    public Output<String> acceleratorId() {
        return this.acceleratorId;
    }

    /**
     * A list of Forwarding Rule IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Forwarding Rule IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The ID of the listener.
     * 
     */
    @Import(name="listenerId", required=true)
    private Output<String> listenerId;

    /**
     * @return The ID of the listener.
     * 
     */
    public Output<String> listenerId() {
        return this.listenerId;
    }

    /**
     * File name where to save data source results (after running `pulumi preview`).
     * 
     */
    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    /**
     * @return File name where to save data source results (after running `pulumi preview`).
     * 
     */
    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    /**
     * The status of the acceleration region. Valid values: `active`, `configuring`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the acceleration region. Valid values: `active`, `configuring`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private GetForwardingRulesArgs() {}

    private GetForwardingRulesArgs(GetForwardingRulesArgs $) {
        this.acceleratorId = $.acceleratorId;
        this.ids = $.ids;
        this.listenerId = $.listenerId;
        this.outputFile = $.outputFile;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetForwardingRulesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetForwardingRulesArgs $;

        public Builder() {
            $ = new GetForwardingRulesArgs();
        }

        public Builder(GetForwardingRulesArgs defaults) {
            $ = new GetForwardingRulesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param acceleratorId The ID of the Global Accelerator instance.
         * 
         * @return builder
         * 
         */
        public Builder acceleratorId(Output<String> acceleratorId) {
            $.acceleratorId = acceleratorId;
            return this;
        }

        /**
         * @param acceleratorId The ID of the Global Accelerator instance.
         * 
         * @return builder
         * 
         */
        public Builder acceleratorId(String acceleratorId) {
            return acceleratorId(Output.of(acceleratorId));
        }

        /**
         * @param ids A list of Forwarding Rule IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Forwarding Rule IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Forwarding Rule IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param listenerId The ID of the listener.
         * 
         * @return builder
         * 
         */
        public Builder listenerId(Output<String> listenerId) {
            $.listenerId = listenerId;
            return this;
        }

        /**
         * @param listenerId The ID of the listener.
         * 
         * @return builder
         * 
         */
        public Builder listenerId(String listenerId) {
            return listenerId(Output.of(listenerId));
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        /**
         * @param outputFile File name where to save data source results (after running `pulumi preview`).
         * 
         * @return builder
         * 
         */
        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        /**
         * @param status The status of the acceleration region. Valid values: `active`, `configuring`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the acceleration region. Valid values: `active`, `configuring`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public GetForwardingRulesArgs build() {
            $.acceleratorId = Objects.requireNonNull($.acceleratorId, "expected parameter 'acceleratorId' to be non-null");
            $.listenerId = Objects.requireNonNull($.listenerId, "expected parameter 'listenerId' to be non-null");
            return $;
        }
    }

}
