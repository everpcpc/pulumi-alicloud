// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dcdn.outputs;

import com.pulumi.alicloud.dcdn.outputs.GetWafRulesWafRule;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetWafRulesResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String outputFile;
    private @Nullable Integer pageNumber;
    private @Nullable Integer pageSize;
    private @Nullable String queryArgs;
    /**
     * @return A list of Waf Rule Entries. Each element contains the following attributes:
     * 
     */
    private List<GetWafRulesWafRule> wafRules;

    private GetWafRulesResult() {}
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
    public Optional<Integer> pageNumber() {
        return Optional.ofNullable(this.pageNumber);
    }
    public Optional<Integer> pageSize() {
        return Optional.ofNullable(this.pageSize);
    }
    public Optional<String> queryArgs() {
        return Optional.ofNullable(this.queryArgs);
    }
    /**
     * @return A list of Waf Rule Entries. Each element contains the following attributes:
     * 
     */
    public List<GetWafRulesWafRule> wafRules() {
        return this.wafRules;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetWafRulesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> ids;
        private @Nullable String outputFile;
        private @Nullable Integer pageNumber;
        private @Nullable Integer pageSize;
        private @Nullable String queryArgs;
        private List<GetWafRulesWafRule> wafRules;
        public Builder() {}
        public Builder(GetWafRulesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.outputFile = defaults.outputFile;
    	      this.pageNumber = defaults.pageNumber;
    	      this.pageSize = defaults.pageSize;
    	      this.queryArgs = defaults.queryArgs;
    	      this.wafRules = defaults.wafRules;
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
        public Builder queryArgs(@Nullable String queryArgs) {
            this.queryArgs = queryArgs;
            return this;
        }
        @CustomType.Setter
        public Builder wafRules(List<GetWafRulesWafRule> wafRules) {
            this.wafRules = Objects.requireNonNull(wafRules);
            return this;
        }
        public Builder wafRules(GetWafRulesWafRule... wafRules) {
            return wafRules(List.of(wafRules));
        }
        public GetWafRulesResult build() {
            final var o = new GetWafRulesResult();
            o.id = id;
            o.ids = ids;
            o.outputFile = outputFile;
            o.pageNumber = pageNumber;
            o.pageSize = pageSize;
            o.queryArgs = queryArgs;
            o.wafRules = wafRules;
            return o;
        }
    }
}
