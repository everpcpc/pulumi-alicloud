// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dms.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEnterpriseProxyAccessesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEnterpriseProxyAccessesPlainArgs Empty = new GetEnterpriseProxyAccessesPlainArgs();

    @Import(name="enableDetails")
    private @Nullable Boolean enableDetails;

    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
    }

    /**
     * A list of Proxy Access IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Proxy Access IDs.
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
     * The ID of the security agent.
     * 
     */
    @Import(name="proxyId", required=true)
    private String proxyId;

    /**
     * @return The ID of the security agent.
     * 
     */
    public String proxyId() {
        return this.proxyId;
    }

    private GetEnterpriseProxyAccessesPlainArgs() {}

    private GetEnterpriseProxyAccessesPlainArgs(GetEnterpriseProxyAccessesPlainArgs $) {
        this.enableDetails = $.enableDetails;
        this.ids = $.ids;
        this.outputFile = $.outputFile;
        this.proxyId = $.proxyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEnterpriseProxyAccessesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEnterpriseProxyAccessesPlainArgs $;

        public Builder() {
            $ = new GetEnterpriseProxyAccessesPlainArgs();
        }

        public Builder(GetEnterpriseProxyAccessesPlainArgs defaults) {
            $ = new GetEnterpriseProxyAccessesPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder enableDetails(@Nullable Boolean enableDetails) {
            $.enableDetails = enableDetails;
            return this;
        }

        /**
         * @param ids A list of Proxy Access IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Proxy Access IDs.
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
         * @param proxyId The ID of the security agent.
         * 
         * @return builder
         * 
         */
        public Builder proxyId(String proxyId) {
            $.proxyId = proxyId;
            return this;
        }

        public GetEnterpriseProxyAccessesPlainArgs build() {
            $.proxyId = Objects.requireNonNull($.proxyId, "expected parameter 'proxyId' to be non-null");
            return $;
        }
    }

}
