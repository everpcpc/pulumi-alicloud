// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ots.outputs;

import com.pulumi.alicloud.ots.outputs.GetSearchIndexesIndex;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetSearchIndexesResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of search index IDs.
     * 
     */
    private List<String> ids;
    /**
     * @return A list of indexes. Each element contains the following attributes:
     * 
     */
    private List<GetSearchIndexesIndex> indexes;
    /**
     * @return The OTS instance name.
     * 
     */
    private String instanceName;
    private @Nullable String nameRegex;
    /**
     * @return A list of search index  names.
     * 
     */
    private List<String> names;
    private @Nullable String outputFile;
    /**
     * @return The table name of the OTS which could not be changed.
     * 
     */
    private String tableName;

    private GetSearchIndexesResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of search index IDs.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    /**
     * @return A list of indexes. Each element contains the following attributes:
     * 
     */
    public List<GetSearchIndexesIndex> indexes() {
        return this.indexes;
    }
    /**
     * @return The OTS instance name.
     * 
     */
    public String instanceName() {
        return this.instanceName;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return A list of search index  names.
     * 
     */
    public List<String> names() {
        return this.names;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return The table name of the OTS which could not be changed.
     * 
     */
    public String tableName() {
        return this.tableName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSearchIndexesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> ids;
        private List<GetSearchIndexesIndex> indexes;
        private String instanceName;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private String tableName;
        public Builder() {}
        public Builder(GetSearchIndexesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.indexes = defaults.indexes;
    	      this.instanceName = defaults.instanceName;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.tableName = defaults.tableName;
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
        public Builder indexes(List<GetSearchIndexesIndex> indexes) {
            this.indexes = Objects.requireNonNull(indexes);
            return this;
        }
        public Builder indexes(GetSearchIndexesIndex... indexes) {
            return indexes(List.of(indexes));
        }
        @CustomType.Setter
        public Builder instanceName(String instanceName) {
            this.instanceName = Objects.requireNonNull(instanceName);
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
        public Builder tableName(String tableName) {
            this.tableName = Objects.requireNonNull(tableName);
            return this;
        }
        public GetSearchIndexesResult build() {
            final var o = new GetSearchIndexesResult();
            o.id = id;
            o.ids = ids;
            o.indexes = indexes;
            o.instanceName = instanceName;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.tableName = tableName;
            return o;
        }
    }
}
