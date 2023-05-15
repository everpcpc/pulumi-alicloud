// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr.outputs;

import com.pulumi.alicloud.hbr.outputs.GetBackupJobsFilter;
import com.pulumi.alicloud.hbr.outputs.GetBackupJobsJob;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBackupJobsResult {
    private @Nullable List<GetBackupJobsFilter> filters;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private List<GetBackupJobsJob> jobs;
    private @Nullable String outputFile;
    private @Nullable String sortDirection;
    private String sourceType;
    private @Nullable String status;

    private GetBackupJobsResult() {}
    public List<GetBackupJobsFilter> filters() {
        return this.filters == null ? List.of() : this.filters;
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
    public List<GetBackupJobsJob> jobs() {
        return this.jobs;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public Optional<String> sortDirection() {
        return Optional.ofNullable(this.sortDirection);
    }
    public String sourceType() {
        return this.sourceType;
    }
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBackupJobsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<GetBackupJobsFilter> filters;
        private String id;
        private List<String> ids;
        private List<GetBackupJobsJob> jobs;
        private @Nullable String outputFile;
        private @Nullable String sortDirection;
        private String sourceType;
        private @Nullable String status;
        public Builder() {}
        public Builder(GetBackupJobsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.filters = defaults.filters;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.jobs = defaults.jobs;
    	      this.outputFile = defaults.outputFile;
    	      this.sortDirection = defaults.sortDirection;
    	      this.sourceType = defaults.sourceType;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder filters(@Nullable List<GetBackupJobsFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetBackupJobsFilter... filters) {
            return filters(List.of(filters));
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
        public Builder jobs(List<GetBackupJobsJob> jobs) {
            this.jobs = Objects.requireNonNull(jobs);
            return this;
        }
        public Builder jobs(GetBackupJobsJob... jobs) {
            return jobs(List.of(jobs));
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder sortDirection(@Nullable String sortDirection) {
            this.sortDirection = sortDirection;
            return this;
        }
        @CustomType.Setter
        public Builder sourceType(String sourceType) {
            this.sourceType = Objects.requireNonNull(sourceType);
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        public GetBackupJobsResult build() {
            final var o = new GetBackupJobsResult();
            o.filters = filters;
            o.id = id;
            o.ids = ids;
            o.jobs = jobs;
            o.outputFile = outputFile;
            o.sortDirection = sortDirection;
            o.sourceType = sourceType;
            o.status = status;
            return o;
        }
    }
}
