// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetMscSubSubscriptionsSubscription {
    /**
     * @return The channel the Subscription.
     * 
     */
    private String channel;
    /**
     * @return The ids of subscribed contacts.
     * 
     */
    private List<Integer> contactIds;
    /**
     * @return The description of the Subscription.
     * 
     */
    private String description;
    /**
     * @return The status of email subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
     * 
     */
    private Integer emailStatus;
    /**
     * @return The ID of the Subscription.
     * 
     */
    private String id;
    /**
     * @return The ID of the Subscription.
     * 
     */
    private String itemId;
    /**
     * @return The name of the Subscription.
     * 
     */
    private String itemName;
    /**
     * @return The status of pmsg subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
     * 
     */
    private Integer pmsgStatus;
    /**
     * @return The status of sms subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
     * 
     */
    private Integer smsStatus;
    /**
     * @return The status of tts subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
     * 
     */
    private Integer ttsStatus;
    /**
     * @return The ids of subscribed webhooks.
     * 
     */
    private List<Integer> webhookIds;
    /**
     * @return The status of webhook subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
     * 
     */
    private Integer webhookStatus;

    private GetMscSubSubscriptionsSubscription() {}
    /**
     * @return The channel the Subscription.
     * 
     */
    public String channel() {
        return this.channel;
    }
    /**
     * @return The ids of subscribed contacts.
     * 
     */
    public List<Integer> contactIds() {
        return this.contactIds;
    }
    /**
     * @return The description of the Subscription.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The status of email subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
     * 
     */
    public Integer emailStatus() {
        return this.emailStatus;
    }
    /**
     * @return The ID of the Subscription.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The ID of the Subscription.
     * 
     */
    public String itemId() {
        return this.itemId;
    }
    /**
     * @return The name of the Subscription.
     * 
     */
    public String itemName() {
        return this.itemName;
    }
    /**
     * @return The status of pmsg subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
     * 
     */
    public Integer pmsgStatus() {
        return this.pmsgStatus;
    }
    /**
     * @return The status of sms subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
     * 
     */
    public Integer smsStatus() {
        return this.smsStatus;
    }
    /**
     * @return The status of tts subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
     * 
     */
    public Integer ttsStatus() {
        return this.ttsStatus;
    }
    /**
     * @return The ids of subscribed webhooks.
     * 
     */
    public List<Integer> webhookIds() {
        return this.webhookIds;
    }
    /**
     * @return The status of webhook subscription. Valid values: `-1`, `-2`, `0`, `1`. `-1` means required, `-2` means banned; `1` means subscribed; `0` means not subscribed.
     * 
     */
    public Integer webhookStatus() {
        return this.webhookStatus;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMscSubSubscriptionsSubscription defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String channel;
        private List<Integer> contactIds;
        private String description;
        private Integer emailStatus;
        private String id;
        private String itemId;
        private String itemName;
        private Integer pmsgStatus;
        private Integer smsStatus;
        private Integer ttsStatus;
        private List<Integer> webhookIds;
        private Integer webhookStatus;
        public Builder() {}
        public Builder(GetMscSubSubscriptionsSubscription defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.channel = defaults.channel;
    	      this.contactIds = defaults.contactIds;
    	      this.description = defaults.description;
    	      this.emailStatus = defaults.emailStatus;
    	      this.id = defaults.id;
    	      this.itemId = defaults.itemId;
    	      this.itemName = defaults.itemName;
    	      this.pmsgStatus = defaults.pmsgStatus;
    	      this.smsStatus = defaults.smsStatus;
    	      this.ttsStatus = defaults.ttsStatus;
    	      this.webhookIds = defaults.webhookIds;
    	      this.webhookStatus = defaults.webhookStatus;
        }

        @CustomType.Setter
        public Builder channel(String channel) {
            this.channel = Objects.requireNonNull(channel);
            return this;
        }
        @CustomType.Setter
        public Builder contactIds(List<Integer> contactIds) {
            this.contactIds = Objects.requireNonNull(contactIds);
            return this;
        }
        public Builder contactIds(Integer... contactIds) {
            return contactIds(List.of(contactIds));
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder emailStatus(Integer emailStatus) {
            this.emailStatus = Objects.requireNonNull(emailStatus);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder itemId(String itemId) {
            this.itemId = Objects.requireNonNull(itemId);
            return this;
        }
        @CustomType.Setter
        public Builder itemName(String itemName) {
            this.itemName = Objects.requireNonNull(itemName);
            return this;
        }
        @CustomType.Setter
        public Builder pmsgStatus(Integer pmsgStatus) {
            this.pmsgStatus = Objects.requireNonNull(pmsgStatus);
            return this;
        }
        @CustomType.Setter
        public Builder smsStatus(Integer smsStatus) {
            this.smsStatus = Objects.requireNonNull(smsStatus);
            return this;
        }
        @CustomType.Setter
        public Builder ttsStatus(Integer ttsStatus) {
            this.ttsStatus = Objects.requireNonNull(ttsStatus);
            return this;
        }
        @CustomType.Setter
        public Builder webhookIds(List<Integer> webhookIds) {
            this.webhookIds = Objects.requireNonNull(webhookIds);
            return this;
        }
        public Builder webhookIds(Integer... webhookIds) {
            return webhookIds(List.of(webhookIds));
        }
        @CustomType.Setter
        public Builder webhookStatus(Integer webhookStatus) {
            this.webhookStatus = Objects.requireNonNull(webhookStatus);
            return this;
        }
        public GetMscSubSubscriptionsSubscription build() {
            final var o = new GetMscSubSubscriptionsSubscription();
            o.channel = channel;
            o.contactIds = contactIds;
            o.description = description;
            o.emailStatus = emailStatus;
            o.id = id;
            o.itemId = itemId;
            o.itemName = itemName;
            o.pmsgStatus = pmsgStatus;
            o.smsStatus = smsStatus;
            o.ttsStatus = ttsStatus;
            o.webhookIds = webhookIds;
            o.webhookStatus = webhookStatus;
            return o;
        }
    }
}
