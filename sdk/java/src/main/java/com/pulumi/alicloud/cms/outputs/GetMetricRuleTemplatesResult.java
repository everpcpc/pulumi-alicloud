// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms.outputs;

import com.pulumi.alicloud.cms.outputs.GetMetricRuleTemplatesTemplate;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetMetricRuleTemplatesResult {
    private @Nullable Boolean enableDetails;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String keyword;
    private @Nullable String metricRuleTemplateName;
    private @Nullable String nameRegex;
    private List<String> names;
    private @Nullable String outputFile;
    private @Nullable String templateId;
    private List<GetMetricRuleTemplatesTemplate> templates;

    private GetMetricRuleTemplatesResult() {}
    public Optional<Boolean> enableDetails() {
        return Optional.ofNullable(this.enableDetails);
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
    public Optional<String> keyword() {
        return Optional.ofNullable(this.keyword);
    }
    public Optional<String> metricRuleTemplateName() {
        return Optional.ofNullable(this.metricRuleTemplateName);
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
    public Optional<String> templateId() {
        return Optional.ofNullable(this.templateId);
    }
    public List<GetMetricRuleTemplatesTemplate> templates() {
        return this.templates;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMetricRuleTemplatesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enableDetails;
        private String id;
        private List<String> ids;
        private @Nullable String keyword;
        private @Nullable String metricRuleTemplateName;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable String templateId;
        private List<GetMetricRuleTemplatesTemplate> templates;
        public Builder() {}
        public Builder(GetMetricRuleTemplatesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enableDetails = defaults.enableDetails;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.keyword = defaults.keyword;
    	      this.metricRuleTemplateName = defaults.metricRuleTemplateName;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.templateId = defaults.templateId;
    	      this.templates = defaults.templates;
        }

        @CustomType.Setter
        public Builder enableDetails(@Nullable Boolean enableDetails) {
            this.enableDetails = enableDetails;
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
        public Builder keyword(@Nullable String keyword) {
            this.keyword = keyword;
            return this;
        }
        @CustomType.Setter
        public Builder metricRuleTemplateName(@Nullable String metricRuleTemplateName) {
            this.metricRuleTemplateName = metricRuleTemplateName;
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
        public Builder templateId(@Nullable String templateId) {
            this.templateId = templateId;
            return this;
        }
        @CustomType.Setter
        public Builder templates(List<GetMetricRuleTemplatesTemplate> templates) {
            this.templates = Objects.requireNonNull(templates);
            return this;
        }
        public Builder templates(GetMetricRuleTemplatesTemplate... templates) {
            return templates(List.of(templates));
        }
        public GetMetricRuleTemplatesResult build() {
            final var o = new GetMetricRuleTemplatesResult();
            o.enableDetails = enableDetails;
            o.id = id;
            o.ids = ids;
            o.keyword = keyword;
            o.metricRuleTemplateName = metricRuleTemplateName;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.templateId = templateId;
            o.templates = templates;
            return o;
        }
    }
}