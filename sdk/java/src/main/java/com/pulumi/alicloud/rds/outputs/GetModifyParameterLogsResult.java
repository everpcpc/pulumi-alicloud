// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds.outputs;

import com.pulumi.alicloud.rds.outputs.GetModifyParameterLogsLog;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetModifyParameterLogsResult {
    private String dbInstanceId;
    private String endTime;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<GetModifyParameterLogsLog> logs;
    private @Nullable String outputFile;
    private String startTime;

    private GetModifyParameterLogsResult() {}
    public String dbInstanceId() {
        return this.dbInstanceId;
    }
    public String endTime() {
        return this.endTime;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<GetModifyParameterLogsLog> logs() {
        return this.logs;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public String startTime() {
        return this.startTime;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetModifyParameterLogsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String dbInstanceId;
        private String endTime;
        private String id;
        private List<GetModifyParameterLogsLog> logs;
        private @Nullable String outputFile;
        private String startTime;
        public Builder() {}
        public Builder(GetModifyParameterLogsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dbInstanceId = defaults.dbInstanceId;
    	      this.endTime = defaults.endTime;
    	      this.id = defaults.id;
    	      this.logs = defaults.logs;
    	      this.outputFile = defaults.outputFile;
    	      this.startTime = defaults.startTime;
        }

        @CustomType.Setter
        public Builder dbInstanceId(String dbInstanceId) {
            this.dbInstanceId = Objects.requireNonNull(dbInstanceId);
            return this;
        }
        @CustomType.Setter
        public Builder endTime(String endTime) {
            this.endTime = Objects.requireNonNull(endTime);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder logs(List<GetModifyParameterLogsLog> logs) {
            this.logs = Objects.requireNonNull(logs);
            return this;
        }
        public Builder logs(GetModifyParameterLogsLog... logs) {
            return logs(List.of(logs));
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder startTime(String startTime) {
            this.startTime = Objects.requireNonNull(startTime);
            return this;
        }
        public GetModifyParameterLogsResult build() {
            final var o = new GetModifyParameterLogsResult();
            o.dbInstanceId = dbInstanceId;
            o.endTime = endTime;
            o.id = id;
            o.logs = logs;
            o.outputFile = outputFile;
            o.startTime = startTime;
            return o;
        }
    }
}