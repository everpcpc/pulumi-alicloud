// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.message.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetServiceSubscriptionsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetServiceSubscriptionsArgs Empty = new GetServiceSubscriptionsArgs();

    /**
     * A list of Subscription IDs.
     * 
     */
    @Import(name="ids")
    private @Nullable Output<List<String>> ids;

    /**
     * @return A list of Subscription IDs.
     * 
     */
    public Optional<Output<List<String>>> ids() {
        return Optional.ofNullable(this.ids);
    }

    /**
     * A regex string to filter results by Subscription name.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable Output<String> nameRegex;

    /**
     * @return A regex string to filter results by Subscription name.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
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

    /**
     * The name of the subscription.
     * 
     */
    @Import(name="subscriptionName")
    private @Nullable Output<String> subscriptionName;

    /**
     * @return The name of the subscription.
     * 
     */
    public Optional<Output<String>> subscriptionName() {
        return Optional.ofNullable(this.subscriptionName);
    }

    /**
     * The name of the topic.
     * 
     */
    @Import(name="topicName", required=true)
    private Output<String> topicName;

    /**
     * @return The name of the topic.
     * 
     */
    public Output<String> topicName() {
        return this.topicName;
    }

    private GetServiceSubscriptionsArgs() {}

    private GetServiceSubscriptionsArgs(GetServiceSubscriptionsArgs $) {
        this.ids = $.ids;
        this.nameRegex = $.nameRegex;
        this.outputFile = $.outputFile;
        this.pageNumber = $.pageNumber;
        this.pageSize = $.pageSize;
        this.subscriptionName = $.subscriptionName;
        this.topicName = $.topicName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServiceSubscriptionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServiceSubscriptionsArgs $;

        public Builder() {
            $ = new GetServiceSubscriptionsArgs();
        }

        public Builder(GetServiceSubscriptionsArgs defaults) {
            $ = new GetServiceSubscriptionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ids A list of Subscription IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(@Nullable Output<List<String>> ids) {
            $.ids = ids;
            return this;
        }

        /**
         * @param ids A list of Subscription IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(List<String> ids) {
            return ids(Output.of(ids));
        }

        /**
         * @param ids A list of Subscription IDs.
         * 
         * @return builder
         * 
         */
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }

        /**
         * @param nameRegex A regex string to filter results by Subscription name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param nameRegex A regex string to filter results by Subscription name.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
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

        /**
         * @param subscriptionName The name of the subscription.
         * 
         * @return builder
         * 
         */
        public Builder subscriptionName(@Nullable Output<String> subscriptionName) {
            $.subscriptionName = subscriptionName;
            return this;
        }

        /**
         * @param subscriptionName The name of the subscription.
         * 
         * @return builder
         * 
         */
        public Builder subscriptionName(String subscriptionName) {
            return subscriptionName(Output.of(subscriptionName));
        }

        /**
         * @param topicName The name of the topic.
         * 
         * @return builder
         * 
         */
        public Builder topicName(Output<String> topicName) {
            $.topicName = topicName;
            return this;
        }

        /**
         * @param topicName The name of the topic.
         * 
         * @return builder
         * 
         */
        public Builder topicName(String topicName) {
            return topicName(Output.of(topicName));
        }

        public GetServiceSubscriptionsArgs build() {
            $.topicName = Objects.requireNonNull($.topicName, "expected parameter 'topicName' to be non-null");
            return $;
        }
    }

}
