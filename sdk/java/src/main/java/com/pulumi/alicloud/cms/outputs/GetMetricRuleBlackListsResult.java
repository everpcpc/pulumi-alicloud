// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms.outputs;

import com.pulumi.alicloud.cms.outputs.GetMetricRuleBlackListsList;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetMetricRuleBlackListsResult {
    /**
     * @return Cloud service classification. For example, Redis includes kvstore_standard, kvstore_sharding, and kvstore_splitrw.
     * 
     */
    private @Nullable String category;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of Metric Rule Black List IDs.
     * 
     */
    private List<String> ids;
    /**
     * @return A list of Metric Rule Black List Entries. Each element contains the following attributes:
     * 
     */
    private List<GetMetricRuleBlackListsList> lists;
    /**
     * @return The first ID of the resource
     * 
     */
    private @Nullable String metricRuleBlackListId;
    private @Nullable String nameRegex;
    /**
     * @return A list of name of Metric Rule Black Lists.
     * 
     */
    private List<String> names;
    /**
     * @return The data namespace of the cloud service.
     * 
     */
    private @Nullable String namespace;
    private @Nullable Integer order;
    private @Nullable String outputFile;
    private @Nullable Integer pageNumber;
    private @Nullable Integer pageSize;

    private GetMetricRuleBlackListsResult() {}
    /**
     * @return Cloud service classification. For example, Redis includes kvstore_standard, kvstore_sharding, and kvstore_splitrw.
     * 
     */
    public Optional<String> category() {
        return Optional.ofNullable(this.category);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of Metric Rule Black List IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    /**
     * @return A list of Metric Rule Black List Entries. Each element contains the following attributes:
     * 
     */
    public List<GetMetricRuleBlackListsList> lists() {
        return this.lists;
    }
    /**
     * @return The first ID of the resource
     * 
     */
    public Optional<String> metricRuleBlackListId() {
        return Optional.ofNullable(this.metricRuleBlackListId);
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of name of Metric Rule Black Lists.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    /**
     * @return The data namespace of the cloud service.
     * 
     */
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    public Optional<Integer> order() {
        return Optional.ofNullable(this.order);
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

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMetricRuleBlackListsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String category;
        private String id;
        private List<String> ids;
        private List<GetMetricRuleBlackListsList> lists;
        private @Nullable String metricRuleBlackListId;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String namespace;
        private @Nullable Integer order;
        private @Nullable String outputFile;
        private @Nullable Integer pageNumber;
        private @Nullable Integer pageSize;
        public Builder() {}
        public Builder(GetMetricRuleBlackListsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.category = defaults.category;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.lists = defaults.lists;
    	      this.metricRuleBlackListId = defaults.metricRuleBlackListId;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.namespace = defaults.namespace;
    	      this.order = defaults.order;
    	      this.outputFile = defaults.outputFile;
    	      this.pageNumber = defaults.pageNumber;
    	      this.pageSize = defaults.pageSize;
        }

        @CustomType.Setter
        public Builder category(@Nullable String category) {
            this.category = category;
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
        public Builder lists(List<GetMetricRuleBlackListsList> lists) {
            this.lists = Objects.requireNonNull(lists);
            return this;
        }
        public Builder lists(GetMetricRuleBlackListsList... lists) {
            return lists(List.of(lists));
        }
        @CustomType.Setter
        public Builder metricRuleBlackListId(@Nullable String metricRuleBlackListId) {
            this.metricRuleBlackListId = metricRuleBlackListId;
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
        public Builder order(@Nullable Integer order) {
            this.order = order;
            return this;
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
        public GetMetricRuleBlackListsResult build() {
            final var o = new GetMetricRuleBlackListsResult();
            o.category = category;
            o.id = id;
            o.ids = ids;
            o.lists = lists;
            o.metricRuleBlackListId = metricRuleBlackListId;
            o.nameRegex = nameRegex;
            o.names = names;
            o.namespace = namespace;
            o.order = order;
            o.outputFile = outputFile;
            o.pageNumber = pageNumber;
            o.pageSize = pageSize;
            return o;
        }
    }
}
