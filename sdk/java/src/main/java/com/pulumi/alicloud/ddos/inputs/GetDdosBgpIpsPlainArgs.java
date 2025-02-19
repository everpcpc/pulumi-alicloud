// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ddos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDdosBgpIpsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDdosBgpIpsPlainArgs Empty = new GetDdosBgpIpsPlainArgs();

    /**
     * A list of Ip IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Ip IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The ID of the native protection enterprise instance to be operated.
     * 
     */
    @Import(name="instanceId", required=true)
    private String instanceId;

    /**
     * @return The ID of the native protection enterprise instance to be operated.
     * 
     */
    public String instanceId() {
        return this.instanceId;
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

    @Import(name="pageNumber")
    private @Nullable Integer pageNumber;

    public Optional<Integer> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }

    @Import(name="pageSize")
    private @Nullable Integer pageSize;

    public Optional<Integer> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }

    /**
     * The product name. Valid Value:`ECS`, `SLB`, `EIP`, `WAF`.
     * 
     */
    @Import(name="productName")
    private @Nullable String productName;

    /**
     * @return The product name. Valid Value:`ECS`, `SLB`, `EIP`, `WAF`.
     * 
     */
    public Optional<String> productName() {
        return Optional.ofNullable(this.productName);
    }

    /**
     * The current state of the IP address.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The current state of the IP address.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    private GetDdosBgpIpsPlainArgs() {}

    private GetDdosBgpIpsPlainArgs(GetDdosBgpIpsPlainArgs $) {
        this.ids = $.ids;
        this.instanceId = $.instanceId;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
        this.productName = $.productName;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDdosBgpIpsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDdosBgpIpsPlainArgs $;

        public Builder() {
            $ = new GetDdosBgpIpsPlainArgs();
        }

        public Builder(GetDdosBgpIpsPlainArgs defaults) {
            $ = new GetDdosBgpIpsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Ip IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Ip IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceId The ID of the native protection enterprise instance to be operated.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            $.instanceId = instanceId;
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

        public Builder pageNumber(@Nullable Integer pageNumber) {
            $.pageNumber = pageNumber;
            return this;
        }

        public Builder pageSize(@Nullable Integer pageSize) {
            $.pageSize = pageSize;
            return this;
        }

        /**
         * @param productName The product name. Valid Value:`ECS`, `SLB`, `EIP`, `WAF`.
         * 
         * @return builder
         * 
         */
        public Builder productName(@Nullable String productName) {
            $.productName = productName;
            return this;
        }

        /**
         * @param status The current state of the IP address.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        public GetDdosBgpIpsPlainArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
