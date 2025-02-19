// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.alb.outputs;

import com.pulumi.alicloud.alb.outputs.GetRulesRule;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetRulesResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable List<String> listenerIds;
    private @Nullable List<String> loadBalancerIds;
    private @Nullable String nameRegex;
    private List<String> names;
    private @Nullable String outputFile;
    private @Nullable List<String> ruleIds;
    private List<GetRulesRule> rules;
    private @Nullable String status;

    private GetRulesResult() {}
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
    public List<String> listenerIds() {
        return this.listenerIds == null ? List.of() : this.listenerIds;
    }
    public List<String> loadBalancerIds() {
        return this.loadBalancerIds == null ? List.of() : this.loadBalancerIds;
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
    public List<String> ruleIds() {
        return this.ruleIds == null ? List.of() : this.ruleIds;
    }
    public List<GetRulesRule> rules() {
        return this.rules;
    }
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRulesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> ids;
        private @Nullable List<String> listenerIds;
        private @Nullable List<String> loadBalancerIds;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable List<String> ruleIds;
        private List<GetRulesRule> rules;
        private @Nullable String status;
        public Builder() {}
        public Builder(GetRulesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.listenerIds = defaults.listenerIds;
    	      this.loadBalancerIds = defaults.loadBalancerIds;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.ruleIds = defaults.ruleIds;
    	      this.rules = defaults.rules;
    	      this.status = defaults.status;
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
        public Builder listenerIds(@Nullable List<String> listenerIds) {
            this.listenerIds = listenerIds;
            return this;
        }
        public Builder listenerIds(String... listenerIds) {
            return listenerIds(List.of(listenerIds));
        }
        @CustomType.Setter
        public Builder loadBalancerIds(@Nullable List<String> loadBalancerIds) {
            this.loadBalancerIds = loadBalancerIds;
            return this;
        }
        public Builder loadBalancerIds(String... loadBalancerIds) {
            return loadBalancerIds(List.of(loadBalancerIds));
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
        public Builder ruleIds(@Nullable List<String> ruleIds) {
            this.ruleIds = ruleIds;
            return this;
        }
        public Builder ruleIds(String... ruleIds) {
            return ruleIds(List.of(ruleIds));
        }
        @CustomType.Setter
        public Builder rules(List<GetRulesRule> rules) {
            this.rules = Objects.requireNonNull(rules);
            return this;
        }
        public Builder rules(GetRulesRule... rules) {
            return rules(List.of(rules));
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        public GetRulesResult build() {
            final var o = new GetRulesResult();
            o.id = id;
            o.ids = ids;
            o.listenerIds = listenerIds;
            o.loadBalancerIds = loadBalancerIds;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.ruleIds = ruleIds;
            o.rules = rules;
            o.status = status;
            return o;
        }
    }
}
