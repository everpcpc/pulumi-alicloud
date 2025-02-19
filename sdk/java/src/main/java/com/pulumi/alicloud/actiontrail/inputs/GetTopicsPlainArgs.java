// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.actiontrail.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTopicsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTopicsPlainArgs Empty = new GetTopicsPlainArgs();

    /**
     * A list of ALIKAFKA Topics IDs, It is formatted to `&lt;instance_id&gt;:&lt;topic&gt;`.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of ALIKAFKA Topics IDs, It is formatted to `&lt;instance_id&gt;:&lt;topic&gt;`.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * ID of the instance.
     * 
     */
    @Import(name="instanceId", required=true)
    private String instanceId;

    /**
     * @return ID of the instance.
     * 
     */
    public String instanceId() {
        return this.instanceId;
    }

    /**
     * A regex string to filter results by the topic name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by the topic name.
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
     * A topic to filter results by the topic name.
     * 
     */
    @Import(name="topic")
    private @Nullable String topic;

    /**
     * @return A topic to filter results by the topic name.
     * 
     */
    public Optional<String> topic() {
        return Optional.ofNullable(this.topic);
    }

    private GetTopicsPlainArgs() {}

    private GetTopicsPlainArgs(GetTopicsPlainArgs $) {
        this.ids = $.ids;
        this.instanceId = $.instanceId;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
        this.topic = $.topic;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTopicsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTopicsPlainArgs $;

        public Builder() {
            $ = new GetTopicsPlainArgs();
        }

        public Builder(GetTopicsPlainArgs defaults) {
            $ = new GetTopicsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of ALIKAFKA Topics IDs, It is formatted to `&lt;instance_id&gt;:&lt;topic&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of ALIKAFKA Topics IDs, It is formatted to `&lt;instance_id&gt;:&lt;topic&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param instanceId ID of the instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by the topic name.
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

        public Builder pageNumber(@Nullable Integer pageNumber) {
            $.pageNumber = pageNumber;
            return this;
        }

        public Builder pageSize(@Nullable Integer pageSize) {
            $.pageSize = pageSize;
            return this;
        }

        /**
         * @param topic A topic to filter results by the topic name.
         * 
         * @return builder
         * 
         */
        public Builder topic(@Nullable String topic) {
            $.topic = topic;
            return this;
        }

        public GetTopicsPlainArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
