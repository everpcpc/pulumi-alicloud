// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCommandsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCommandsPlainArgs Empty = new GetCommandsPlainArgs();

    /**
     * Public order provider.
     * 
     */
    @Import(name="commandProvider")
    private @Nullable String commandProvider;

    /**
     * @return Public order provider.
     * 
     */
    public Optional<String> commandProvider() {
        return Optional.ofNullable(this.commandProvider);
    }

    /**
     * The Base64-encoded content of the command.
     * 
     */
    @Import(name="contentEncoding")
    private @Nullable String contentEncoding;

    /**
     * @return The Base64-encoded content of the command.
     * 
     */
    public Optional<String> contentEncoding() {
        return Optional.ofNullable(this.contentEncoding);
    }

    /**
     * The description of command.
     * 
     */
    @Import(name="description")
    private @Nullable String description;

    /**
     * @return The description of command.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * A list of Command IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable List<String> ids;

    /**
     * @return A list of Command IDs.
     * 
     */
    public Optional<List<String>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * The name of the command
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the command
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A regex string to filter results by Command name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return A regex string to filter results by Command name.
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
     * The command type.
     * 
     */
    @Import(name="type")
    private @Nullable String type;

    /**
     * @return The command type.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    private GetCommandsPlainArgs() {}

    private GetCommandsPlainArgs(GetCommandsPlainArgs $) {
        this.commandProvider = $.commandProvider;
        this.contentEncoding = $.contentEncoding;
        this.description = $.description;
        this.ids = $.ids;
        this.name = $.name;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCommandsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCommandsPlainArgs $;

        public Builder() {
            $ = new GetCommandsPlainArgs();
        }

        public Builder(GetCommandsPlainArgs defaults) {
            $ = new GetCommandsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param commandProvider Public order provider.
         * 
         * @return builder
         * 
         */
        public Builder commandProvider(@Nullable String commandProvider) {
            $.commandProvider = commandProvider;
            return this;
        }

        /**
         * @param contentEncoding The Base64-encoded content of the command.
         * 
         * @return builder
         * 
         */
        public Builder contentEncoding(@Nullable String contentEncoding) {
            $.contentEncoding = contentEncoding;
            return this;
        }

        /**
         * @param description The description of command.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable String description) {
            $.description = description;
            return this;
        }

        /**
         * @param ids A list of Command IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable List<String> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Command IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param name The name of the command
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Command name.
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
         * @param type The command type.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable String type) {
            $.type = type;
            return this;
        }

        public GetCommandsPlainArgs build() {
            return $;
        }
    }

}
