// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.chatbot.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAgentsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAgentsArgs Empty = new GetAgentsArgs();

    /**
     * The name of the agent.
     * 
     */
    @Import(name="agentName")
    private @Nullable Output<String> agentName;

    /**
     * @return The name of the agent.
     * 
     */
    public Optional<Output<String>> agentName() {
        return Optional.ofNullable(this.agentName);
    }

    /**
     * A regex string to filter resulting chatbot agents by name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter resulting chatbot agents by name.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    @Import(name="outputFile")
    private @Nullable Output<String> outputFile;

    public Optional<Output<String>> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }

    @Import(name="pageNumber")
    private @Nullable Output<Integer> pageNumber;

    public Optional<Output<Integer>> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }

    @Import(name="pageSize")
    private @Nullable Output<Integer> pageSize;

    public Optional<Output<Integer>> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }

    private GetAgentsArgs() {}

    private GetAgentsArgs(GetAgentsArgs $) {
        this.agentName = $.agentName;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAgentsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAgentsArgs $;

        public Builder() {
            $ = new GetAgentsArgs();
        }

        public Builder(GetAgentsArgs defaults) {
            $ = new GetAgentsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param agentName The name of the agent.
         * 
         * @return builder
         * 
         */
        public Builder agentName(@Nullable Output<String> agentName) {
            $.agentName = agentName;
            return this;
        }

        /**
         * @param agentName The name of the agent.
         * 
         * @return builder
         * 
         */
        public Builder agentName(String agentName) {
            return agentName(Output.of(agentName));
        }

        /**
         * @param nameRegex A regex string to filter resulting chatbot agents by name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter resulting chatbot agents by name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
        }

        public Builder outputFile(@Nullable Output<String> outputFile) {
            $.outputFile = outputFile;
            return this;
        }

        public Builder outputFile(String outputFile) {
            return outputFile(Output.of(outputFile));
        }

        public Builder pageNumber(@Nullable Output<Integer> pageNumber) {
            $.pageNumber = pageNumber;
            return this;
        }

        public Builder pageNumber(Integer pageNumber) {
            return pageNumber(Output.of(pageNumber));
        }

        public Builder pageSize(@Nullable Output<Integer> pageSize) {
            $.pageSize = pageSize;
            return this;
        }

        public Builder pageSize(Integer pageSize) {
            return pageSize(Output.of(pageSize));
        }

        public GetAgentsArgs build() {
            return $;
        }
    }

}
