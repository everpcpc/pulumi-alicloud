// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs.outputs;

import com.pulumi.alicloud.ecs.outputs.GetKeyPairsKeyPair;
import com.pulumi.alicloud.ecs.outputs.GetKeyPairsPair;
import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetKeyPairsResult {
    /**
     * @return Finger print of the key pair.
     * 
     */
    private @Nullable String fingerPrint;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    /**
     * @return A list of key pairs. Each element contains the following attributes:
     * 
     * @deprecated
     * Field &#39;key_pairs&#39; has been deprecated from provider version 1.121.0. New field &#39;pairs&#39; instead.
     * 
     */
    @Deprecated /* Field 'key_pairs' has been deprecated from provider version 1.121.0. New field 'pairs' instead. */
    private List<GetKeyPairsKeyPair> keyPairs;
    private @Nullable String nameRegex;
    /**
     * @return A list of key pair names.
     * 
     */
    private List<String> names;
    private @Nullable String outputFile;
    private List<GetKeyPairsPair> pairs;
    /**
     * @return The Id of resource group.
     * 
     */
    private @Nullable String resourceGroupId;
    /**
     * @return (Optional, Available in v1.66.0+) A mapping of tags to assign to the resource.
     * 
     */
    private @Nullable Map<String,Object> tags;

    private GetKeyPairsResult() {}
    /**
     * @return Finger print of the key pair.
     * 
     */
    public Optional<String> fingerPrint() {
        return Optional.ofNullable(this.fingerPrint);
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
    /**
     * @return A list of key pairs. Each element contains the following attributes:
     * 
     * @deprecated
     * Field &#39;key_pairs&#39; has been deprecated from provider version 1.121.0. New field &#39;pairs&#39; instead.
     * 
     */
    @Deprecated /* Field 'key_pairs' has been deprecated from provider version 1.121.0. New field 'pairs' instead. */
    public List<GetKeyPairsKeyPair> keyPairs() {
        return this.keyPairs;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of key pair names.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public List<GetKeyPairsPair> pairs() {
        return this.pairs;
    }
    /**
     * @return The Id of resource group.
     * 
     */
    public Optional<String> resourceGroupId() {
        return Optional.ofNullable(this.resourceGroupId);
    }
    /**
     * @return (Optional, Available in v1.66.0+) A mapping of tags to assign to the resource.
     * 
     */
    public Map<String,Object> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKeyPairsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String fingerPrint;
        private String id;
        private List<String> ids;
        private List<GetKeyPairsKeyPair> keyPairs;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private List<GetKeyPairsPair> pairs;
        private @Nullable String resourceGroupId;
        private @Nullable Map<String,Object> tags;
        public Builder() {}
        public Builder(GetKeyPairsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fingerPrint = defaults.fingerPrint;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.keyPairs = defaults.keyPairs;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.pairs = defaults.pairs;
    	      this.resourceGroupId = defaults.resourceGroupId;
    	      this.tags = defaults.tags;
        }

        @CustomType.Setter
        public Builder fingerPrint(@Nullable String fingerPrint) {
            this.fingerPrint = fingerPrint;
            return this;
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
        public Builder keyPairs(List<GetKeyPairsKeyPair> keyPairs) {
            this.keyPairs = Objects.requireNonNull(keyPairs);
            return this;
        }
        public Builder keyPairs(GetKeyPairsKeyPair... keyPairs) {
            return keyPairs(List.of(keyPairs));
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
        public Builder pairs(List<GetKeyPairsPair> pairs) {
            this.pairs = Objects.requireNonNull(pairs);
            return this;
        }
        public Builder pairs(GetKeyPairsPair... pairs) {
            return pairs(List.of(pairs));
        }
        @CustomType.Setter
        public Builder resourceGroupId(@Nullable String resourceGroupId) {
            this.resourceGroupId = resourceGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,Object> tags) {
            this.tags = tags;
            return this;
        }
        public GetKeyPairsResult build() {
            final var o = new GetKeyPairsResult();
            o.fingerPrint = fingerPrint;
            o.id = id;
            o.ids = ids;
            o.keyPairs = keyPairs;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.pairs = pairs;
            o.resourceGroupId = resourceGroupId;
            o.tags = tags;
            return o;
        }
    }
}
