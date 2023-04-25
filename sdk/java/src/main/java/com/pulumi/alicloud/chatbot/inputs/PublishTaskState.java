// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.chatbot.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PublishTaskState extends com.pulumi.resources.ResourceArgs {

    public static final PublishTaskState Empty = new PublishTaskState();

    /**
     * The business space key. If you do not set it, the default business space is accessed. The key value is obtained on the business management page of the primary account.
     * 
     */
    @Import(name="agentKey")
    private @Nullable Output<String> agentKey;

    /**
     * @return The business space key. If you do not set it, the default business space is accessed. The key value is obtained on the business management page of the primary account.
     * 
     */
    public Optional<Output<String>> agentKey() {
        return Optional.ofNullable(this.agentKey);
    }

    /**
     * The type of the publishing unit. Please use the CreateInstancePublishTask API to publish the robot.
     * 
     */
    @Import(name="bizType")
    private @Nullable Output<String> bizType;

    /**
     * @return The type of the publishing unit. Please use the CreateInstancePublishTask API to publish the robot.
     * 
     */
    public Optional<Output<String>> bizType() {
        return Optional.ofNullable(this.bizType);
    }

    /**
     * UTC time of task creation
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return UTC time of task creation
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * Additional release information. Currently supported: If the BizType is faq, enter the category Id in this field to indicate that only the knowledge under these categories is published.
     * 
     */
    @Import(name="dataIdLists")
    private @Nullable Output<List<String>> dataIdLists;

    /**
     * @return Additional release information. Currently supported: If the BizType is faq, enter the category Id in this field to indicate that only the knowledge under these categories is published.
     * 
     */
    public Optional<Output<List<String>>> dataIdLists() {
        return Optional.ofNullable(this.dataIdLists);
    }

    /**
     * UTC time for task modification
     * 
     */
    @Import(name="modifyTime")
    private @Nullable Output<String> modifyTime;

    /**
     * @return UTC time for task modification
     * 
     */
    public Optional<Output<String>> modifyTime() {
        return Optional.ofNullable(this.modifyTime);
    }

    /**
     * The status of the task.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the task.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private PublishTaskState() {}

    private PublishTaskState(PublishTaskState $) {
        this.agentKey = $.agentKey;
        this.bizType = $.bizType;
        this.createTime = $.createTime;
        this.dataIdLists = $.dataIdLists;
        this.modifyTime = $.modifyTime;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PublishTaskState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PublishTaskState $;

        public Builder() {
            $ = new PublishTaskState();
        }

        public Builder(PublishTaskState defaults) {
            $ = new PublishTaskState(Objects.requireNonNull(defaults));
        }

        /**
         * @param agentKey The business space key. If you do not set it, the default business space is accessed. The key value is obtained on the business management page of the primary account.
         * 
         * @return builder
         * 
         */
        public Builder agentKey(@Nullable Output<String> agentKey) {
            $.agentKey = agentKey;
            return this;
        }

        /**
         * @param agentKey The business space key. If you do not set it, the default business space is accessed. The key value is obtained on the business management page of the primary account.
         * 
         * @return builder
         * 
         */
        public Builder agentKey(String agentKey) {
            return agentKey(Output.of(agentKey));
        }

        /**
         * @param bizType The type of the publishing unit. Please use the CreateInstancePublishTask API to publish the robot.
         * 
         * @return builder
         * 
         */
        public Builder bizType(@Nullable Output<String> bizType) {
            $.bizType = bizType;
            return this;
        }

        /**
         * @param bizType The type of the publishing unit. Please use the CreateInstancePublishTask API to publish the robot.
         * 
         * @return builder
         * 
         */
        public Builder bizType(String bizType) {
            return bizType(Output.of(bizType));
        }

        /**
         * @param createTime UTC time of task creation
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime UTC time of task creation
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param dataIdLists Additional release information. Currently supported: If the BizType is faq, enter the category Id in this field to indicate that only the knowledge under these categories is published.
         * 
         * @return builder
         * 
         */
        public Builder dataIdLists(@Nullable Output<List<String>> dataIdLists) {
            $.dataIdLists = dataIdLists;
            return this;
        }

        /**
         * @param dataIdLists Additional release information. Currently supported: If the BizType is faq, enter the category Id in this field to indicate that only the knowledge under these categories is published.
         * 
         * @return builder
         * 
         */
        public Builder dataIdLists(List<String> dataIdLists) {
            return dataIdLists(Output.of(dataIdLists));
        }

        /**
         * @param dataIdLists Additional release information. Currently supported: If the BizType is faq, enter the category Id in this field to indicate that only the knowledge under these categories is published.
         * 
         * @return builder
         * 
         */
        public Builder dataIdLists(String... dataIdLists) {
            return dataIdLists(List.of(dataIdLists));
        }

        /**
         * @param modifyTime UTC time for task modification
         * 
         * @return builder
         * 
         */
        public Builder modifyTime(@Nullable Output<String> modifyTime) {
            $.modifyTime = modifyTime;
            return this;
        }

        /**
         * @param modifyTime UTC time for task modification
         * 
         * @return builder
         * 
         */
        public Builder modifyTime(String modifyTime) {
            return modifyTime(Output.of(modifyTime));
        }

        /**
         * @param status The status of the task.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the task.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public PublishTaskState build() {
            return $;
        }
    }

}
