// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cfg.outputs;

import com.pulumi.alicloud.cfg.outputs.GetConfigurationRecordersRecorder;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetConfigurationRecordersResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String outputFile;
    /**
     * @return A list of Config Configuration Recorders. Each element contains the following attributes:
     * 
     */
    private List<GetConfigurationRecordersRecorder> recorders;

    private GetConfigurationRecordersResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    /**
     * @return A list of Config Configuration Recorders. Each element contains the following attributes:
     * 
     */
    public List<GetConfigurationRecordersRecorder> recorders() {
        return this.recorders;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetConfigurationRecordersResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private @Nullable String outputFile;
        private List<GetConfigurationRecordersRecorder> recorders;
        public Builder() {}
        public Builder(GetConfigurationRecordersResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.outputFile = defaults.outputFile;
    	      this.recorders = defaults.recorders;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder recorders(List<GetConfigurationRecordersRecorder> recorders) {
            this.recorders = Objects.requireNonNull(recorders);
            return this;
        }
        public Builder recorders(GetConfigurationRecordersRecorder... recorders) {
            return recorders(List.of(recorders));
        }
        public GetConfigurationRecordersResult build() {
            final var o = new GetConfigurationRecordersResult();
            o.id = id;
            o.outputFile = outputFile;
            o.recorders = recorders;
            return o;
        }
    }
}
