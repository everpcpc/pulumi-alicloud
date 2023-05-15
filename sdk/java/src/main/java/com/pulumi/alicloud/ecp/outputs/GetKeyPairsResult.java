// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecp.outputs;

import com.pulumi.alicloud.ecp.outputs.GetKeyPairsPair;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetKeyPairsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String keyPairFingerPrint;
    private @Nullable String nameRegex;
    private List<String> names;
    private @Nullable String outputFile;
    private List<GetKeyPairsPair> pairs;

    private GetKeyPairsResult() {}
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
    public Optional<String> keyPairFingerPrint() {
        return Optional.ofNullable(this.keyPairFingerPrint);
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public List<GetKeyPairsPair> pairs() {
        return this.pairs;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKeyPairsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> ids;
        private @Nullable String keyPairFingerPrint;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private List<GetKeyPairsPair> pairs;
        public Builder() {}
        public Builder(GetKeyPairsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.keyPairFingerPrint = defaults.keyPairFingerPrint;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.pairs = defaults.pairs;
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
        public Builder keyPairFingerPrint(@Nullable String keyPairFingerPrint) {
            this.keyPairFingerPrint = keyPairFingerPrint;
            return this;
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
        public GetKeyPairsResult build() {
            final var o = new GetKeyPairsResult();
            o.id = id;
            o.ids = ids;
            o.keyPairFingerPrint = keyPairFingerPrint;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.pairs = pairs;
            return o;
        }
    }
}
