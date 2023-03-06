// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc.outputs;

import com.pulumi.alicloud.vpc.outputs.GetVpcFlowLogsLog;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetVpcFlowLogsResult {
    private @Nullable String description;
    private @Nullable String flowLogName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String logStoreName;
    private List<GetVpcFlowLogsLog> logs;
    private @Nullable String nameRegex;
    private List<String> names;
    private @Nullable String outputFile;
    private @Nullable String projectName;
    private @Nullable String resourceId;
    private @Nullable String resourceType;
    private @Nullable String status;
    private @Nullable String trafficType;

    private GetVpcFlowLogsResult() {}
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    public Optional<String> flowLogName() {
        return Optional.ofNullable(this.flowLogName);
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
    public Optional<String> logStoreName() {
        return Optional.ofNullable(this.logStoreName);
    }
    public List<GetVpcFlowLogsLog> logs() {
        return this.logs;
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
    public Optional<String> projectName() {
        return Optional.ofNullable(this.projectName);
    }
    public Optional<String> resourceId() {
        return Optional.ofNullable(this.resourceId);
    }
    public Optional<String> resourceType() {
        return Optional.ofNullable(this.resourceType);
    }
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    public Optional<String> trafficType() {
        return Optional.ofNullable(this.trafficType);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVpcFlowLogsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String description;
        private @Nullable String flowLogName;
        private String id;
        private List<String> ids;
        private @Nullable String logStoreName;
        private List<GetVpcFlowLogsLog> logs;
        private @Nullable String nameRegex;
        private List<String> names;
        private @Nullable String outputFile;
        private @Nullable String projectName;
        private @Nullable String resourceId;
        private @Nullable String resourceType;
        private @Nullable String status;
        private @Nullable String trafficType;
        public Builder() {}
        public Builder(GetVpcFlowLogsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.flowLogName = defaults.flowLogName;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.logStoreName = defaults.logStoreName;
    	      this.logs = defaults.logs;
    	      this.nameRegex = defaults.nameRegex;
    	      this.names = defaults.names;
    	      this.outputFile = defaults.outputFile;
    	      this.projectName = defaults.projectName;
    	      this.resourceId = defaults.resourceId;
    	      this.resourceType = defaults.resourceType;
    	      this.status = defaults.status;
    	      this.trafficType = defaults.trafficType;
        }

        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder flowLogName(@Nullable String flowLogName) {
            this.flowLogName = flowLogName;
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
        public Builder logStoreName(@Nullable String logStoreName) {
            this.logStoreName = logStoreName;
            return this;
        }
        @CustomType.Setter
        public Builder logs(List<GetVpcFlowLogsLog> logs) {
            this.logs = Objects.requireNonNull(logs);
            return this;
        }
        public Builder logs(GetVpcFlowLogsLog... logs) {
            return logs(List.of(logs));
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
        public Builder projectName(@Nullable String projectName) {
            this.projectName = projectName;
            return this;
        }
        @CustomType.Setter
        public Builder resourceId(@Nullable String resourceId) {
            this.resourceId = resourceId;
            return this;
        }
        @CustomType.Setter
        public Builder resourceType(@Nullable String resourceType) {
            this.resourceType = resourceType;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder trafficType(@Nullable String trafficType) {
            this.trafficType = trafficType;
            return this;
        }
        public GetVpcFlowLogsResult build() {
            final var o = new GetVpcFlowLogsResult();
            o.description = description;
            o.flowLogName = flowLogName;
            o.id = id;
            o.ids = ids;
            o.logStoreName = logStoreName;
            o.logs = logs;
            o.nameRegex = nameRegex;
            o.names = names;
            o.outputFile = outputFile;
            o.projectName = projectName;
            o.resourceId = resourceId;
            o.resourceType = resourceType;
            o.status = status;
            o.trafficType = trafficType;
            return o;
        }
    }
}