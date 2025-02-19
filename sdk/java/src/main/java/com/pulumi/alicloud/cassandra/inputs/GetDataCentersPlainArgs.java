// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cassandra.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDataCentersPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDataCentersPlainArgs Empty = new GetDataCentersPlainArgs();

    /**
     * The cluster id of dataCenters belongs to.
     * 
     */
    @Import(name="clusterId", required=true)
    private String clusterId;

    /**
     * @return The cluster id of dataCenters belongs to.
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }

    /**
     * The list of Cassandra data center ids.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return The list of Cassandra data center ids.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to apply to the cluster name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to apply to the cluster name.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * The name of file that can save the collection of data centers after running `pulumi preview`.
     * 
     */
    @Import(name="outputFile")
    private @Nullable String outputFile;

    /**
     * @return The name of file that can save the collection of data centers after running `pulumi preview`.
     * 
     */
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    private GetDataCentersPlainArgs() {}

    private GetDataCentersPlainArgs(GetDataCentersPlainArgs $) {
        this.clusterId = $.clusterId;
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDataCentersPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDataCentersPlainArgs $;

        public Builder() {
            $ = new GetDataCentersPlainArgs();
        }

        public Builder(GetDataCentersPlainArgs defaults) {
            $ = new GetDataCentersPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId The cluster id of dataCenters belongs to.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param ids The list of Cassandra data center ids.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids The list of Cassandra data center ids.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to apply to the cluster name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param outputFile The name of file that can save the collection of data centers after running `pulumi preview`.
         * 
         * @return builder
         * 
         */
        public Builder outputFile(@Nullable String outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public GetDataCentersPlainArgs build() {
            $.clusterId = Objects.requireNonNull($.clusterId, "expected parameter 'clusterId' to be non-null");
            return $;
        }
    }

}
