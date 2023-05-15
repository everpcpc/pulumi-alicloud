// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rocketmq.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetTopicsTopic {
    /**
     * @return The id of the topic.
     * 
     */
    private String id;
    /**
     * @return Indicates whether namespaces are available. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
     * 
     */
    private Boolean independentNaming;
    /**
     * @return ID of the ONS Instance that owns the topics.
     * 
     */
    private String instanceId;
    /**
     * @return The type of the message. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
     * 
     */
    private Integer messageType;
    /**
     * @return The ID of the topic owner, which is the Alibaba Cloud UID.
     * 
     */
    private String owner;
    /**
     * @return This attribute is used to set the read-write mode for the topic.
     * 
     */
    private Integer perm;
    /**
     * @return The relation ID. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
     * 
     */
    private Integer relation;
    /**
     * @return The name of the relation, for example, owner, publishable, subscribable, and publishable and subscribable.
     * 
     */
    private String relationName;
    /**
     * @return Remark of the topic.
     * 
     */
    private String remark;
    /**
     * @return A map of tags assigned to the Ons instance.
     * 
     */
    private Map<String,Object> tags;
    /**
     * @return The name of the topic.
     * 
     */
    private String topic;
    /**
     * @return The name of the topic.
     * 
     */
    private String topicName;

    private GetTopicsTopic() {}
    /**
     * @return The id of the topic.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Indicates whether namespaces are available. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
     * 
     */
    public Boolean independentNaming() {
        return this.independentNaming;
    }
    /**
     * @return ID of the ONS Instance that owns the topics.
     * 
     */
    public String instanceId() {
        return this.instanceId;
    }
    /**
     * @return The type of the message. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
     * 
     */
    public Integer messageType() {
        return this.messageType;
    }
    /**
     * @return The ID of the topic owner, which is the Alibaba Cloud UID.
     * 
     */
    public String owner() {
        return this.owner;
    }
    /**
     * @return This attribute is used to set the read-write mode for the topic.
     * 
     */
    public Integer perm() {
        return this.perm;
    }
    /**
     * @return The relation ID. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.
     * 
     */
    public Integer relation() {
        return this.relation;
    }
    /**
     * @return The name of the relation, for example, owner, publishable, subscribable, and publishable and subscribable.
     * 
     */
    public String relationName() {
        return this.relationName;
    }
    /**
     * @return Remark of the topic.
     * 
     */
    public String remark() {
        return this.remark;
    }
    /**
     * @return A map of tags assigned to the Ons instance.
     * 
     */
    public Map<String,Object> tags() {
        return this.tags;
    }
    /**
     * @return The name of the topic.
     * 
     */
    public String topic() {
        return this.topic;
    }
    /**
     * @return The name of the topic.
     * 
     */
    public String topicName() {
        return this.topicName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTopicsTopic defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private Boolean independentNaming;
        private String instanceId;
        private Integer messageType;
        private String owner;
        private Integer perm;
        private Integer relation;
        private String relationName;
        private String remark;
        private Map<String,Object> tags;
        private String topic;
        private String topicName;
        public Builder() {}
        public Builder(GetTopicsTopic defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.independentNaming = defaults.independentNaming;
    	      this.instanceId = defaults.instanceId;
    	      this.messageType = defaults.messageType;
    	      this.owner = defaults.owner;
    	      this.perm = defaults.perm;
    	      this.relation = defaults.relation;
    	      this.relationName = defaults.relationName;
    	      this.remark = defaults.remark;
    	      this.tags = defaults.tags;
    	      this.topic = defaults.topic;
    	      this.topicName = defaults.topicName;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder independentNaming(Boolean independentNaming) {
            this.independentNaming = Objects.requireNonNull(independentNaming);
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(String instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder messageType(Integer messageType) {
            this.messageType = Objects.requireNonNull(messageType);
            return this;
        }
        @CustomType.Setter
        public Builder owner(String owner) {
            this.owner = Objects.requireNonNull(owner);
            return this;
        }
        @CustomType.Setter
        public Builder perm(Integer perm) {
            this.perm = Objects.requireNonNull(perm);
            return this;
        }
        @CustomType.Setter
        public Builder relation(Integer relation) {
            this.relation = Objects.requireNonNull(relation);
            return this;
        }
        @CustomType.Setter
        public Builder relationName(String relationName) {
            this.relationName = Objects.requireNonNull(relationName);
            return this;
        }
        @CustomType.Setter
        public Builder remark(String remark) {
            this.remark = Objects.requireNonNull(remark);
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,Object> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        @CustomType.Setter
        public Builder topic(String topic) {
            this.topic = Objects.requireNonNull(topic);
            return this;
        }
        @CustomType.Setter
        public Builder topicName(String topicName) {
            this.topicName = Objects.requireNonNull(topicName);
            return this;
        }
        public GetTopicsTopic build() {
            final var o = new GetTopicsTopic();
            o.id = id;
            o.independentNaming = independentNaming;
            o.instanceId = instanceId;
            o.messageType = messageType;
            o.owner = owner;
            o.perm = perm;
            o.relation = relation;
            o.relationName = relationName;
            o.remark = remark;
            o.tags = tags;
            o.topic = topic;
            o.topicName = topicName;
            return o;
        }
    }
}
