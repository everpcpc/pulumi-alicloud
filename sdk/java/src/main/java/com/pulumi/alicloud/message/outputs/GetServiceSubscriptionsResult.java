// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.message.outputs;

import com.pulumi.alicloud.message.outputs.GetServiceSubscriptionsSubscription;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetServiceSubscriptionsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String nameRegex;
    /**
     * @return A list of Subscription names.
     * 
     */
    private List<String> names;
    private @Nullable String outputFile;
    private @Nullable Integer pageNumber;
    private @Nullable Integer pageSize;
    /**
     * @return The name of the subscription.
     * 
     */
    private @Nullable String subscriptionName;
    /**
     * @return A list of Subscriptions. Each element contains the following attributes:
     * 
     */
    private List<GetServiceSubscriptionsSubscription> subscriptions;
    /**
     * @return The name of the topic.
     * 
     */
    private String topicName;

    private GetServiceSubscriptionsResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of Subscription names.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Optional<Integer> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }
    public Optional<Integer> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }
    /**
     * @return The name of the subscription.
     * 
     */
    public Optional<String> subscriptionName() {
        return Optional.ofNullable(this.subscriptionName);
    }
    /**
     * @return A list of Subscriptions. Each element contains the following attributes:
     * 
     */
    public List<GetServiceSubscriptionsSubscription> subscriptions() {
        return this.subscriptions;
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

    public static Builder builder(GetServiceSubscriptionsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> ids;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable Integer pageNumber;
        private @Nullable Integer pageSize;
        private @Nullable String subscriptionName;
        private List<GetServiceSubscriptionsSubscription> subscriptions;
        private String topicName;
        public Builder() {}
        public Builder(GetServiceSubscriptionsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.pageNumber = defaults.pageNumber;
    	      this.pageSize = defaults.pageSize;
    	      this.subscriptionName = defaults.subscriptionName;
    	      this.subscriptions = defaults.subscriptions;
    	      this.topicName = defaults.topicName;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ids(List<String> ids) {
            this.ids = Objects.requireNonNull(ids);
            return this;
        }
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }
        @CustomType.Setter
        public Builder nameRegex(@Nullable String nameRegex) {
            this.nameRegex = nameRegex;
            return this;
        }
        @CustomType.Setter
        public Builder names(List<String> names) {
            this.names = Objects.requireNonNull(names);
            return this;
        }
        public Builder names(String... names) {
            return names(List.of(names));
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder pageNumber(@Nullable Integer pageNumber) {
            this.pageNumber = pageNumber;
            return this;
        }
        @CustomType.Setter
        public Builder pageSize(@Nullable Integer pageSize) {
            this.pageSize = pageSize;
            return this;
        }
        @CustomType.Setter
        public Builder subscriptionName(@Nullable String subscriptionName) {
            this.subscriptionName = subscriptionName;
            return this;
        }
        @CustomType.Setter
        public Builder subscriptions(List<GetServiceSubscriptionsSubscription> subscriptions) {
            this.subscriptions = Objects.requireNonNull(subscriptions);
            return this;
        }
        public Builder subscriptions(GetServiceSubscriptionsSubscription... subscriptions) {
            return subscriptions(List.of(subscriptions));
        }
        @CustomType.Setter
        public Builder topicName(String topicName) {
            this.topicName = Objects.requireNonNull(topicName);
            return this;
        }
        public GetServiceSubscriptionsResult build() {
            final var o = new GetServiceSubscriptionsResult();
            o.id = id;
            o.ids = ids;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.pageNumber = pageNumber;
            o.pageSize = pageSize;
            o.subscriptionName = subscriptionName;
            o.subscriptions = subscriptions;
            o.topicName = topicName;
            return o;
        }
    }
}
