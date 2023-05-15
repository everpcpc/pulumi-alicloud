// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.polardb.outputs;

import com.pulumi.alicloud.polardb.outputs.GetDatabasesDatabase;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetDatabasesResult {
    /**
     * @return A list of PolarDB cluster databases. Each element contains the following attributes:
     * 
     */
    private List<GetDatabasesDatabase> databases;
    private String dbClusterId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String nameRegex;
    /**
     * @return database name of the cluster.
     * 
     */
    private List<String> names;

    private GetDatabasesResult() {}
    /**
     * @return A list of PolarDB cluster databases. Each element contains the following attributes:
     * 
     */
    public List<GetDatabasesDatabase> databases() {
        return this.databases;
    }
    public String dbClusterId() {
        return this.dbClusterId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    /**
     * @return database name of the cluster.
     * 
     */
    public List<String> names() {
        return this.names;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDatabasesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetDatabasesDatabase> databases;
        private String dbClusterId;
        private String id;
        private @Nullable String nameRegex;
        private List<String> names;
        public Builder() {}
        public Builder(GetDatabasesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.databases = defaults.databases;
    	      this.dbClusterId = defaults.dbClusterId;
    	      this.id = defaults.id;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
        }

        @CustomType.Setter
        public Builder databases(List<GetDatabasesDatabase> databases) {
            this.databases = Objects.requireNonNull(databases);
            return this;
        }
        public Builder databases(GetDatabasesDatabase... databases) {
            return databases(List.of(databases));
        }
        @CustomType.Setter
        public Builder dbClusterId(String dbClusterId) {
            this.dbClusterId = Objects.requireNonNull(dbClusterId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
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
        public GetDatabasesResult build() {
            final var o = new GetDatabasesResult();
            o.databases = databases;
            o.dbClusterId = dbClusterId;
            o.id = id;
            o.nameRegex = nameRegex;
            o.names = names;
            return o;
        }
    }
}
