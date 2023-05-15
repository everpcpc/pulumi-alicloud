// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetMscSubWebhooksWebhook {
    /**
     * @return The ID of the Webhook.
     * 
     */
    private String id;
    /**
     * @return The serverUrl of the Subscription.
     * 
     */
    private String serverUrl;
    /**
     * @return The first ID of the resource.
     * 
     */
    private String webhookId;
    /**
     * @return The name of the Webhook. **Note:** The name must be `2` to `12` characters in length, and can contain uppercase and lowercase letters.
     * 
     */
    private String webhookName;

    private GetMscSubWebhooksWebhook() {}
    /**
     * @return The ID of the Webhook.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The serverUrl of the Subscription.
     * 
     */
    public String serverUrl() {
        return this.serverUrl;
    }
    /**
     * @return The first ID of the resource.
     * 
     */
    public String webhookId() {
        return this.webhookId;
    }
    /**
     * @return The name of the Webhook. **Note:** The name must be `2` to `12` characters in length, and can contain uppercase and lowercase letters.
     * 
     */
    public String webhookName() {
        return this.webhookName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMscSubWebhooksWebhook defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String serverUrl;
        private String webhookId;
        private String webhookName;
        public Builder() {}
        public Builder(GetMscSubWebhooksWebhook defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.serverUrl = defaults.serverUrl;
    	      this.webhookId = defaults.webhookId;
    	      this.webhookName = defaults.webhookName;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder serverUrl(String serverUrl) {
            this.serverUrl = Objects.requireNonNull(serverUrl);
            return this;
        }
        @CustomType.Setter
        public Builder webhookId(String webhookId) {
            this.webhookId = Objects.requireNonNull(webhookId);
            return this;
        }
        @CustomType.Setter
        public Builder webhookName(String webhookName) {
            this.webhookName = Objects.requireNonNull(webhookName);
            return this;
        }
        public GetMscSubWebhooksWebhook build() {
            final var o = new GetMscSubWebhooksWebhook();
            o.id = id;
            o.serverUrl = serverUrl;
            o.webhookId = webhookId;
            o.webhookName = webhookName;
            return o;
        }
    }
}
