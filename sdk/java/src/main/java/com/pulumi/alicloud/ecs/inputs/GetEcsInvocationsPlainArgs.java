// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEcsInvocationsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEcsInvocationsPlainArgs Empty = new GetEcsInvocationsPlainArgs();

    /**
     * The ID of the command.
     * 
     */
    @Import(name="commandId")
    private @Nullable String commandId;

    /**
     * @return The ID of the command.
     * 
     */
    public Optional<String> commandId() {
        return Optional.ofNullable(this.commandId);
    }

    /**
     * The encoding mode of the CommandContent and Output response parameters. Valid values: `PlainText`, `Base64`.
     * 
     */
    @Import(name="contentEncoding")
    private @Nullable String contentEncoding;

    /**
     * @return The encoding mode of the CommandContent and Output response parameters. Valid values: `PlainText`, `Base64`.
     * 
     */
    public Optional<String> contentEncoding() {
        return Optional.ofNullable(this.contentEncoding);
    }

    /**
     * A list of Invocation IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Invocation IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The overall execution state of the command. **Note:** We recommend that you ignore this parameter and check the value of the `invocation_status` response parameter for the overall execution state.
     * 
     */
    @Import(name="invokeStatus")
    private @Nullable String invokeStatus;

    /**
     * @return The overall execution state of the command. **Note:** We recommend that you ignore this parameter and check the value of the `invocation_status` response parameter for the overall execution state.
     * 
     */
    public Optional<String> invokeStatus() {
        return Optional.ofNullable(this.invokeStatus);
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

    private GetEcsInvocationsPlainArgs() {}

    private GetEcsInvocationsPlainArgs(GetEcsInvocationsPlainArgs $) {
        this.commandId = $.commandId;
        this.contentEncoding = $.contentEncoding;
        this.ids = $.ids;
        this.invokeStatus = $.invokeStatus;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEcsInvocationsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEcsInvocationsPlainArgs $;

        public Builder() {
            $ = new GetEcsInvocationsPlainArgs();
        }

        public Builder(GetEcsInvocationsPlainArgs defaults) {
            $ = new GetEcsInvocationsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param commandId The ID of the command.
         * 
         * @return builder
         * 
         */
        public Builder commandId(@Nullable String commandId) {
            $.commandId = commandId;
            return this;
        }

        /**
         * @param contentEncoding The encoding mode of the CommandContent and Output response parameters. Valid values: `PlainText`, `Base64`.
         * 
         * @return builder
         * 
         */
        public Builder contentEncoding(@Nullable String contentEncoding) {
            $.contentEncoding = contentEncoding;
            return this;
        }

        /**
         * @param ids A list of Invocation IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Invocation IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param invokeStatus The overall execution state of the command. **Note:** We recommend that you ignore this parameter and check the value of the `invocation_status` response parameter for the overall execution state.
         * 
         * @return builder
         * 
         */
        public Builder invokeStatus(@Nullable String invokeStatus) {
            $.invokeStatus = invokeStatus;
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

        public GetEcsInvocationsPlainArgs build() {
            return $;
        }
    }

}
