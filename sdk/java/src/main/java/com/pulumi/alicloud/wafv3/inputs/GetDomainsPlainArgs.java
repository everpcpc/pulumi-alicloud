// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.wafv3.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDomainsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDomainsPlainArgs Empty = new GetDomainsPlainArgs();

    /**
     * The address type of the origin server. The address can be an IP address or a domain name. You can specify only one type of address.
     * 
     */
    @Import(name="backend")
    private @Nullable String backend;

    /**
     * @return The address type of the origin server. The address can be an IP address or a domain name. You can specify only one type of address.
     * 
     */
    public Optional<String> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * The name of the domain name to query.
     * 
     */
    @Import(name="domain")
    private @Nullable String domain;

    /**
     * @return The name of the domain name to query.
     * 
     */
    public Optional<String> domain() {
        return Optional.ofNullable(this.domain);
    }

    /**
     * Default to `false`. Set it to `true` can output more details about resource attributes.
     * 
     */
    @Import(name="enableDetails")
    private @Nullable Boolean enableDetails;

    /**
     * @return Default to `false`. Set it to `true` can output more details about resource attributes.
     * 
     */
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * A list of domain IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of domain IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The WAF instance ID.
     * 
     */
    @Import(name="instanceId", required=true)
    private String instanceId;

    /**
     * @return The WAF instance ID.
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

    private GetDomainsPlainArgs() {}

    private GetDomainsPlainArgs(GetDomainsPlainArgs $) {
        this.backend = $.backend;
        this.domain = $.domain;
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.instanceId = $.instanceId;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDomainsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDomainsPlainArgs $;

        public Builder() {
            $ = new GetDomainsPlainArgs();
        }

        public Builder(GetDomainsPlainArgs defaults) {
            $ = new GetDomainsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The address type of the origin server. The address can be an IP address or a domain name. You can specify only one type of address.
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable String backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param domain The name of the domain name to query.
         * 
         * @return builder
         * 
         */
        public Builder domain(@Nullable String domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param enableDetails Default to `false`. Set it to `true` can output more details about resource attributes.
         * 
         * @return builder
         * 
         */
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param ids A list of domain IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of domain IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceId The WAF instance ID.
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

        public GetDomainsPlainArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
