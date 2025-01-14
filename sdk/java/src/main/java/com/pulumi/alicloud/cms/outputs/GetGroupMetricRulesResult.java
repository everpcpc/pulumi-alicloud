// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms.outputs;

import com.pulumi.alicloud.cms.outputs.GetGroupMetricRulesRule;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetGroupMetricRulesResult {
    private @Nullable String dimensions;
    private @Nullable Boolean enableState;
    private @Nullable String groupId;
    private @Nullable String groupMetricRuleName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String metricName;
    private @Nullable String nameRegex;
    private List<String> names;
    private @Nullable String namespace;
    private @Nullable String outputFile;
    private List<GetGroupMetricRulesRule> rules;
    private @Nullable String status;

    private GetGroupMetricRulesResult() {}
    public Optional<String> dimensions() {
        return Optional.ofNullable(this.dimensions);
    }
    public Optional<Boolean> enableState() {
        return Optional.ofNullable(this.enableState);
    }
    public Optional<String> groupId() {
        return Optional.ofNullable(this.groupId);
    }
    public Optional<String> groupMetricRuleName() {
        return Optional.ofNullable(this.groupMetricRuleName);
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
    public Optional<String> metricName() {
        return Optional.ofNullable(this.metricName);
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    public List<String> names() {
        return this.names;
    }
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public List<GetGroupMetricRulesRule> rules() {
        return this.rules;
    }
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGroupMetricRulesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String dimensions;
        private @Nullable Boolean enableState;
        private @Nullable String groupId;
        private @Nullable String groupMetricRuleName;
        private String id;
        private List<String> ids;
        private @Nullable String metricName;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String namespace;
        private @Nullable String outputFile;
        private List<GetGroupMetricRulesRule> rules;
        private @Nullable String status;
        public Builder() {}
        public Builder(GetGroupMetricRulesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dimensions = defaults.dimensions;
    	      this.enableState = defaults.enableState;
    	      this.groupId = defaults.groupId;
    	      this.groupMetricRuleName = defaults.groupMetricRuleName;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.metricName = defaults.metricName;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.namespace = defaults.namespace;
    	      this.outputFile = defaults.outputFile;
    	      this.rules = defaults.rules;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder dimensions(@Nullable String dimensions) {
            this.dimensions = dimensions;
            return this;
        }
        @CustomType.Setter
        public Builder enableState(@Nullable Boolean enableState) {
            this.enableState = enableState;
            return this;
        }
        @CustomType.Setter
        public Builder groupId(@Nullable String groupId) {
            this.groupId = groupId;
            return this;
        }
        @CustomType.Setter
        public Builder groupMetricRuleName(@Nullable String groupMetricRuleName) {
            this.groupMetricRuleName = groupMetricRuleName;
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
        public Builder metricName(@Nullable String metricName) {
            this.metricName = metricName;
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
        public Builder namespace(@Nullable String namespace) {
            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder rules(List<GetGroupMetricRulesRule> rules) {
            this.rules = Objects.requireNonNull(rules);
            return this;
        }
        public Builder rules(GetGroupMetricRulesRule... rules) {
            return rules(List.of(rules));
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        public GetGroupMetricRulesResult build() {
            final var o = new GetGroupMetricRulesResult();
            o.dimensions = dimensions;
            o.enableState = enableState;
            o.groupId = groupId;
            o.groupMetricRuleName = groupMetricRuleName;
            o.id = id;
            o.ids = ids;
            o.metricName = metricName;
            o.nameRegex = nameRegex;
            o.names = names;
            o.namespace = namespace;
            o.outputFile = outputFile;
            o.rules = rules;
            o.status = status;
            return o;
        }
    }
}
