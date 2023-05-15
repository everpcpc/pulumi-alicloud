// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.message.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetServiceTopicsTopic {
    /**
     * @return The time when the topic was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
     * 
     */
    private Integer createTime;
    /**
     * @return The id of the Topic. Its value is same as Topic Name.
     * 
     */
    private String id;
    /**
     * @return The time when the topic was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
     * 
     */
    private Integer lastModifyTime;
    /**
     * @return Indicates whether the log management feature is enabled.
     * 
     */
    private Boolean loggingEnabled;
    /**
     * @return The maximum size of a message body that can be sent to the topic. Unit: bytes.
     * 
     */
    private Integer maxMessageSize;
    /**
     * @return The number of messages in the topic.
     * 
     */
    private Integer messageCount;
    /**
     * @return The maximum period for which a message can be retained in the topic. A message that is sent to the topic can be retained for a specified period. After the specified period ends, the message is deleted no matter whether it is pushed to the specified endpoints. Unit: seconds.
     * 
     */
    private Integer messageRetentionPeriod;
    /**
     * @return The inner url of the topic.
     * 
     */
    private String topicInnerUrl;
    /**
     * @return The name of the topic.
     * 
     */
    private String topicName;
    /**
     * @return The url of the topic.
     * 
     */
    private String topicUrl;

    private GetServiceTopicsTopic() {}
    /**
     * @return The time when the topic was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
     * 
     */
    public Integer createTime() {
        return this.createTime;
    }
    /**
     * @return The id of the Topic. Its value is same as Topic Name.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The time when the topic was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
     * 
     */
    public Integer lastModifyTime() {
        return this.lastModifyTime;
    }
    /**
     * @return Indicates whether the log management feature is enabled.
     * 
     */
    public Boolean loggingEnabled() {
        return this.loggingEnabled;
    }
    /**
     * @return The maximum size of a message body that can be sent to the topic. Unit: bytes.
     * 
     */
    public Integer maxMessageSize() {
        return this.maxMessageSize;
    }
    /**
     * @return The number of messages in the topic.
     * 
     */
    public Integer messageCount() {
        return this.messageCount;
    }
    /**
     * @return The maximum period for which a message can be retained in the topic. A message that is sent to the topic can be retained for a specified period. After the specified period ends, the message is deleted no matter whether it is pushed to the specified endpoints. Unit: seconds.
     * 
     */
    public Integer messageRetentionPeriod() {
        return this.messageRetentionPeriod;
    }
    /**
     * @return The inner url of the topic.
     * 
     */
    public String topicInnerUrl() {
        return this.topicInnerUrl;
    }
    /**
     * @return The name of the topic.
     * 
     */
    public String topicName() {
        return this.topicName;
    }
    /**
     * @return The url of the topic.
     * 
     */
    public String topicUrl() {
        return this.topicUrl;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServiceTopicsTopic defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer createTime;
        private String id;
        private Integer lastModifyTime;
        private Boolean loggingEnabled;
        private Integer maxMessageSize;
        private Integer messageCount;
        private Integer messageRetentionPeriod;
        private String topicInnerUrl;
        private String topicName;
        private String topicUrl;
        public Builder() {}
        public Builder(GetServiceTopicsTopic defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createTime = defaults.createTime;
    	      this.id = defaults.id;
    	      this.lastModifyTime = defaults.lastModifyTime;
    	      this.loggingEnabled = defaults.loggingEnabled;
    	      this.maxMessageSize = defaults.maxMessageSize;
    	      this.messageCount = defaults.messageCount;
    	      this.messageRetentionPeriod = defaults.messageRetentionPeriod;
    	      this.topicInnerUrl = defaults.topicInnerUrl;
    	      this.topicName = defaults.topicName;
    	      this.topicUrl = defaults.topicUrl;
        }

        @CustomType.Setter
        public Builder createTime(Integer createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder lastModifyTime(Integer lastModifyTime) {
            this.lastModifyTime = Objects.requireNonNull(lastModifyTime);
            return this;
        }
        @CustomType.Setter
        public Builder loggingEnabled(Boolean loggingEnabled) {
            this.loggingEnabled = Objects.requireNonNull(loggingEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder maxMessageSize(Integer maxMessageSize) {
            this.maxMessageSize = Objects.requireNonNull(maxMessageSize);
            return this;
        }
        @CustomType.Setter
        public Builder messageCount(Integer messageCount) {
            this.messageCount = Objects.requireNonNull(messageCount);
            return this;
        }
        @CustomType.Setter
        public Builder messageRetentionPeriod(Integer messageRetentionPeriod) {
            this.messageRetentionPeriod = Objects.requireNonNull(messageRetentionPeriod);
            return this;
        }
        @CustomType.Setter
        public Builder topicInnerUrl(String topicInnerUrl) {
            this.topicInnerUrl = Objects.requireNonNull(topicInnerUrl);
            return this;
        }
        @CustomType.Setter
        public Builder topicName(String topicName) {
            this.topicName = Objects.requireNonNull(topicName);
            return this;
        }
        @CustomType.Setter
        public Builder topicUrl(String topicUrl) {
            this.topicUrl = Objects.requireNonNull(topicUrl);
            return this;
        }
        public GetServiceTopicsTopic build() {
            final var o = new GetServiceTopicsTopic();
            o.createTime = createTime;
            o.id = id;
            o.lastModifyTime = lastModifyTime;
            o.loggingEnabled = loggingEnabled;
            o.maxMessageSize = maxMessageSize;
            o.messageCount = messageCount;
            o.messageRetentionPeriod = messageRetentionPeriod;
            o.topicInnerUrl = topicInnerUrl;
            o.topicName = topicName;
            o.topicUrl = topicUrl;
            return o;
        }
    }
}
