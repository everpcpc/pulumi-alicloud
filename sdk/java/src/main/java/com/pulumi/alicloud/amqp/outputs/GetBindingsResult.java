// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.amqp.outputs;

import com.pulumi.alicloud.amqp.outputs.GetBindingsBinding;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBindingsResult {
    private List<GetBindingsBinding> bindings;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private String instanceId;
    private @Nullable String outputFile;
    private String virtualHostName;

    private GetBindingsResult() {}
    public List<GetBindingsBinding> bindings() {
        return this.bindings;
    }
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
    public String instanceId() {
        return this.instanceId;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public String virtualHostName() {
        return this.virtualHostName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBindingsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetBindingsBinding> bindings;
        private String id;
        private List<String> ids;
        private String instanceId;
        private @Nullable String outputFile;
        private String virtualHostName;
        public Builder() {}
        public Builder(GetBindingsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bindings = defaults.bindings;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.instanceId = defaults.instanceId;
    	      this.outputFile = defaults.outputFile;
    	      this.virtualHostName = defaults.virtualHostName;
        }

        @CustomType.Setter
        public Builder bindings(List<GetBindingsBinding> bindings) {
            this.bindings = Objects.requireNonNull(bindings);
            return this;
        }
        public Builder bindings(GetBindingsBinding... bindings) {
            return bindings(List.of(bindings));
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
        public Builder instanceId(String instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder virtualHostName(String virtualHostName) {
            this.virtualHostName = Objects.requireNonNull(virtualHostName);
            return this;
        }
        public GetBindingsResult build() {
            final var o = new GetBindingsResult();
            o.bindings = bindings;
            o.id = id;
            o.ids = ids;
            o.instanceId = instanceId;
            o.outputFile = outputFile;
            o.virtualHostName = virtualHostName;
            return o;
        }
    }
}
