// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oos.outputs;

import com.pulumi.alicloud.oos.outputs.GetExecutionsExecution;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetExecutionsResult {
    private @Nullable String category;
    private @Nullable String endDate;
    private @Nullable String endDateAfter;
    private @Nullable String executedBy;
    /**
     * @return A list of OOS Executions. Each element contains the following attributes:
     * 
     */
    private List<GetExecutionsExecution> executions;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of OOS Execution ids.
     * 
     */
    private List<String> ids;
    private @Nullable Boolean includeChildExecution;
    private @Nullable String mode;
    private @Nullable String outputFile;
    private @Nullable String parentExecutionId;
    private @Nullable String ramRole;
    private @Nullable String sortField;
    private @Nullable String sortOrder;
    private @Nullable String startDateAfter;
    private @Nullable String startDateBefore;
    private @Nullable String status;
    private @Nullable Map<String,Object> tags;
    private @Nullable String templateName;

    private GetExecutionsResult() {}
    public Optional<String> category() {
        return Optional.ofNullable(this.category);
    }
    public Optional<String> endDate() {
        return Optional.ofNullable(this.endDate);
    }
    public Optional<String> endDateAfter() {
        return Optional.ofNullable(this.endDateAfter);
    }
    public Optional<String> executedBy() {
        return Optional.ofNullable(this.executedBy);
    }
    /**
     * @return A list of OOS Executions. Each element contains the following attributes:
     * 
     */
    public List<GetExecutionsExecution> executions() {
        return this.executions;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of OOS Execution ids.
     * 
     */
    public List<String> ids() {
        return this.ids;
    }
    public Optional<Boolean> includeChildExecution() {
        return Optional.ofNullable(this.includeChildExecution);
    }
    public Optional<String> mode() {
        return Optional.ofNullable(this.mode);
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Optional<String> parentExecutionId() {
        return Optional.ofNullable(this.parentExecutionId);
    }
    public Optional<String> ramRole() {
        return Optional.ofNullable(this.ramRole);
    }
    public Optional<String> sortField() {
        return Optional.ofNullable(this.sortField);
    }
    public Optional<String> sortOrder() {
        return Optional.ofNullable(this.sortOrder);
    }
    public Optional<String> startDateAfter() {
        return Optional.ofNullable(this.startDateAfter);
    }
    public Optional<String> startDateBefore() {
        return Optional.ofNullable(this.startDateBefore);
    }
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    public Map<String,Object> tags() {
        return this.tags == null ? Map.of() : this.tags;
    }
    public Optional<String> templateName() {
        return Optional.ofNullable(this.templateName);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetExecutionsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String category;
        private @Nullable String endDate;
        private @Nullable String endDateAfter;
        private @Nullable String executedBy;
        private List<GetExecutionsExecution> executions;
        private String id;
        private List<String> ids;
        private @Nullable Boolean includeChildExecution;
        private @Nullable String mode;
        private @Nullable String outputFile;
        private @Nullable String parentExecutionId;
        private @Nullable String ramRole;
        private @Nullable String sortField;
        private @Nullable String sortOrder;
        private @Nullable String startDateAfter;
        private @Nullable String startDateBefore;
        private @Nullable String status;
        private @Nullable Map<String,Object> tags;
        private @Nullable String templateName;
        public Builder() {}
        public Builder(GetExecutionsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.category = defaults.category;
    	      this.endDate = defaults.endDate;
    	      this.endDateAfter = defaults.endDateAfter;
    	      this.executedBy = defaults.executedBy;
    	      this.executions = defaults.executions;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.includeChildExecution = defaults.includeChildExecution;
    	      this.mode = defaults.mode;
    	      this.outputFile = defaults.outputFile;
    	      this.parentExecutionId = defaults.parentExecutionId;
    	      this.ramRole = defaults.ramRole;
    	      this.sortField = defaults.sortField;
    	      this.sortOrder = defaults.sortOrder;
    	      this.startDateAfter = defaults.startDateAfter;
    	      this.startDateBefore = defaults.startDateBefore;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.templateName = defaults.templateName;
        }

        @CustomType.Setter
        public Builder category(@Nullable String category) {
            this.category = category;
            return this;
        }
        @CustomType.Setter
        public Builder endDate(@Nullable String endDate) {
            this.endDate = endDate;
            return this;
        }
        @CustomType.Setter
        public Builder endDateAfter(@Nullable String endDateAfter) {
            this.endDateAfter = endDateAfter;
            return this;
        }
        @CustomType.Setter
        public Builder executedBy(@Nullable String executedBy) {
            this.executedBy = executedBy;
            return this;
        }
        @CustomType.Setter
        public Builder executions(List<GetExecutionsExecution> executions) {
            this.executions = Objects.requireNonNull(executions);
            return this;
        }
        public Builder executions(GetExecutionsExecution... executions) {
            return executions(List.of(executions));
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
        public Builder includeChildExecution(@Nullable Boolean includeChildExecution) {
            this.includeChildExecution = includeChildExecution;
            return this;
        }
        @CustomType.Setter
        public Builder mode(@Nullable String mode) {
            this.mode = mode;
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder parentExecutionId(@Nullable String parentExecutionId) {
            this.parentExecutionId = parentExecutionId;
            return this;
        }
        @CustomType.Setter
        public Builder ramRole(@Nullable String ramRole) {
            this.ramRole = ramRole;
            return this;
        }
        @CustomType.Setter
        public Builder sortField(@Nullable String sortField) {
            this.sortField = sortField;
            return this;
        }
        @CustomType.Setter
        public Builder sortOrder(@Nullable String sortOrder) {
            this.sortOrder = sortOrder;
            return this;
        }
        @CustomType.Setter
        public Builder startDateAfter(@Nullable String startDateAfter) {
            this.startDateAfter = startDateAfter;
            return this;
        }
        @CustomType.Setter
        public Builder startDateBefore(@Nullable String startDateBefore) {
            this.startDateBefore = startDateBefore;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable Map<String,Object> tags) {
            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder templateName(@Nullable String templateName) {
            this.templateName = templateName;
            return this;
        }
        public GetExecutionsResult build() {
            final var o = new GetExecutionsResult();
            o.category = category;
            o.endDate = endDate;
            o.endDateAfter = endDateAfter;
            o.executedBy = executedBy;
            o.executions = executions;
            o.id = id;
            o.ids = ids;
            o.includeChildExecution = includeChildExecution;
            o.mode = mode;
            o.outputFile = outputFile;
            o.parentExecutionId = parentExecutionId;
            o.ramRole = ramRole;
            o.sortField = sortField;
            o.sortOrder = sortOrder;
            o.startDateAfter = startDateAfter;
            o.startDateBefore = startDateBefore;
            o.status = status;
            o.tags = tags;
            o.templateName = templateName;
            return o;
        }
    }
}
