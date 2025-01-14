// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dfs.outputs;

import com.pulumi.alicloud.dfs.outputs.GetAccessRulesRule;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAccessRulesResult {
    private String accessGroupId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String outputFile;
    private List<GetAccessRulesRule> rules;

    private GetAccessRulesResult() {}
    public String accessGroupId() {
        return this.accessGroupId;
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
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public List<GetAccessRulesRule> rules() {
        return this.rules;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAccessRulesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accessGroupId;
        private String id;
        private List<String> ids;
        private @Nullable String outputFile;
        private List<GetAccessRulesRule> rules;
        public Builder() {}
        public Builder(GetAccessRulesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessGroupId = defaults.accessGroupId;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.outputFile = defaults.outputFile;
    	      this.rules = defaults.rules;
        }

        @CustomType.Setter
        public Builder accessGroupId(String accessGroupId) {
            this.accessGroupId = Objects.requireNonNull(accessGroupId);
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
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder rules(List<GetAccessRulesRule> rules) {
            this.rules = Objects.requireNonNull(rules);
            return this;
        }
        public Builder rules(GetAccessRulesRule... rules) {
            return rules(List.of(rules));
        }
        public GetAccessRulesResult build() {
            final var o = new GetAccessRulesResult();
            o.accessGroupId = accessGroupId;
            o.id = id;
            o.ids = ids;
            o.outputFile = outputFile;
            o.rules = rules;
            return o;
        }
    }
}
