// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ddos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDdosCooDomainResourcesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDdosCooDomainResourcesArgs Empty = new GetDdosCooDomainResourcesArgs();

    /**
     * A list of Domain Resource IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Domain Resource IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A list ID of instance that you want to associate.
     * 
     */
    @Import(name="instanceIds")
    private @Nullable Output<List<String>> instanceIds;

    /**
     * @return A list ID of instance that you want to associate.
     * 
     */
    public Optional<Output<List<String>>> instanceIds() {
        return Optional.ofNullable(this.instanceIds);
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
     * Match the pattern.
     * 
     */
    @Import(name="queryDomainPattern")
    private @Nullable Output<String> queryDomainPattern;

    /**
     * @return Match the pattern.
     * 
     */
    public Optional<Output<String>> queryDomainPattern() {
        return Optional.ofNullable(this.queryDomainPattern);
    }

    private GetDdosCooDomainResourcesArgs() {}

    private GetDdosCooDomainResourcesArgs(GetDdosCooDomainResourcesArgs $) {
        this.ids = $.ids;
        this.instanceIds = $.instanceIds;
        this.outputFile = $.outputFile;
        this.queryDomainPattern = $.queryDomainPattern;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDdosCooDomainResourcesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDdosCooDomainResourcesArgs $;

        public Builder() {
            $ = new GetDdosCooDomainResourcesArgs();
        }

        public Builder(GetDdosCooDomainResourcesArgs defaults) {
            $ = new GetDdosCooDomainResourcesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Domain Resource IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Domain Resource IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Domain Resource IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceIds A list ID of instance that you want to associate.
         * 
         * @return builder
         * 
         */
        public Builder instanceIds(@Nullable Output<List<String>> instanceIds) {
            $.instanceIds = instanceIds;
            return this;
        }

        /**
         * @param instanceIds A list ID of instance that you want to associate.
         * 
         * @return builder
         * 
         */
        public Builder instanceIds(List<String> instanceIds) {
            return instanceIds(Output.of(instanceIds));
        }

        /**
         * @param instanceIds A list ID of instance that you want to associate.
         * 
         * @return builder
         * 
         */
        public Builder instanceIds(String... instanceIds) {
            return instanceIds(List.of(instanceIds));
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
         * @param queryDomainPattern Match the pattern.
         * 
         * @return builder
         * 
         */
        public Builder queryDomainPattern(@Nullable Output<String> queryDomainPattern) {
            $.queryDomainPattern = queryDomainPattern;
            return this;
        }

        /**
         * @param queryDomainPattern Match the pattern.
         * 
         * @return builder
         * 
         */
        public Builder queryDomainPattern(String queryDomainPattern) {
            return queryDomainPattern(Output.of(queryDomainPattern));
        }

        public GetDdosCooDomainResourcesArgs build() {
            return $;
        }
    }

}
