// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.outputs;

import com.pulumi.alicloud.outputs.GetMscSubSubscriptionsSubscription;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetMscSubSubscriptionsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String outputFile;
    private List<GetMscSubSubscriptionsSubscription> subscriptions;

    private GetMscSubSubscriptionsResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public List<GetMscSubSubscriptionsSubscription> subscriptions() {
        return this.subscriptions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMscSubSubscriptionsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private @Nullable String outputFile;
        private List<GetMscSubSubscriptionsSubscription> subscriptions;
        public Builder() {}
        public Builder(GetMscSubSubscriptionsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.outputFile = defaults.outputFile;
    	      this.subscriptions = defaults.subscriptions;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder subscriptions(List<GetMscSubSubscriptionsSubscription> subscriptions) {
            this.subscriptions = Objects.requireNonNull(subscriptions);
            return this;
        }
        public Builder subscriptions(GetMscSubSubscriptionsSubscription... subscriptions) {
            return subscriptions(List.of(subscriptions));
        }
        public GetMscSubSubscriptionsResult build() {
            final var o = new GetMscSubSubscriptionsResult();
            o.id = id;
            o.outputFile = outputFile;
            o.subscriptions = subscriptions;
            return o;
        }
    }
}