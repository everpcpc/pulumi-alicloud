// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.slb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetListenersArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetListenersArgs Empty = new GetListenersArgs();

    /**
     * A regex string to filter results by SLB listener description.
     * 
     */
    @Import(name="descriptionRegex")
    private @Nullable Output<String> descriptionRegex;

    /**
     * @return A regex string to filter results by SLB listener description.
     * 
     */
    public Optional<Output<String>> descriptionRegex() {
        return Optional.ofNullable(this.descriptionRegex);
    }

    /**
     * Filter listeners by the specified frontend port.
     * 
     */
    @Import(name="frontendPort")
    private @Nullable Output<Integer> frontendPort;

    /**
     * @return Filter listeners by the specified frontend port.
     * 
     */
    public Optional<Output<Integer>> frontendPort() {
        return Optional.ofNullable(this.frontendPort);
    }

    /**
     * ID of the SLB with listeners.
     * 
     */
    @Import(name="loadBalancerId", required=true)
    private Output<String> loadBalancerId;

    /**
     * @return ID of the SLB with listeners.
     * 
     */
    public Output<String> loadBalancerId() {
        return this.loadBalancerId;
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
     * Filter listeners by the specified protocol. Valid values: `http`, `https`, `tcp` and `udp`.
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return Filter listeners by the specified protocol. Valid values: `http`, `https`, `tcp` and `udp`.
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    private GetListenersArgs() {}

    private GetListenersArgs(GetListenersArgs $) {
        this.descriptionRegex = $.descriptionRegex;
        this.frontendPort = $.frontendPort;
        this.loadBalancerId = $.loadBalancerId;
        this.outputFile = $.outputFile;
        this.protocol = $.protocol;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetListenersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetListenersArgs $;

        public Builder() {
            $ = new GetListenersArgs();
        }

        public Builder(GetListenersArgs defaults) {
            $ = new GetListenersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param descriptionRegex A regex string to filter results by SLB listener description.
         * 
         * @return builder
         * 
         */
        public Builder descriptionRegex(@Nullable Output<String> descriptionRegex) {
            $.descriptionRegex = descriptionRegex;
            return this;
        }

        /**
         * @param descriptionRegex A regex string to filter results by SLB listener description.
         * 
         * @return builder
         * 
         */
        public Builder descriptionRegex(String descriptionRegex) {
            return descriptionRegex(Output.of(descriptionRegex));
        }

        /**
         * @param frontendPort Filter listeners by the specified frontend port.
         * 
         * @return builder
         * 
         */
        public Builder frontendPort(@Nullable Output<Integer> frontendPort) {
            $.frontendPort = frontendPort;
            return this;
        }

        /**
         * @param frontendPort Filter listeners by the specified frontend port.
         * 
         * @return builder
         * 
         */
        public Builder frontendPort(Integer frontendPort) {
            return frontendPort(Output.of(frontendPort));
        }

        /**
         * @param loadBalancerId ID of the SLB with listeners.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerId(Output<String> loadBalancerId) {
            $.loadBalancerId = loadBalancerId;
            return this;
        }

        /**
         * @param loadBalancerId ID of the SLB with listeners.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerId(String loadBalancerId) {
            return loadBalancerId(Output.of(loadBalancerId));
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
         * @param protocol Filter listeners by the specified protocol. Valid values: `http`, `https`, `tcp` and `udp`.
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol Filter listeners by the specified protocol. Valid values: `http`, `https`, `tcp` and `udp`.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        public GetListenersArgs build() {
            $.loadBalancerId = Objects.requireNonNull($.loadBalancerId, "expected parameter 'loadBalancerId' to be non-null");
            return $;
        }
    }

}
