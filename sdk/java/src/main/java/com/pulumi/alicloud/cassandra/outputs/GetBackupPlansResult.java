// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cassandra.outputs;

import com.pulumi.alicloud.cassandra.outputs.GetBackupPlansPlan;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBackupPlansResult {
    private String clusterId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String outputFile;
    private List<GetBackupPlansPlan> plans;

    private GetBackupPlansResult() {}
    public String clusterId() {
        return this.clusterId;
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
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public List<GetBackupPlansPlan> plans() {
        return this.plans;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBackupPlansResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterId;
        private String id;
        private List<String> ids;
        private @Nullable String outputFile;
        private List<GetBackupPlansPlan> plans;
        public Builder() {}
        public Builder(GetBackupPlansResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterId = defaults.clusterId;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.outputFile = defaults.outputFile;
    	      this.plans = defaults.plans;
        }

        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            this.clusterId = Objects.requireNonNull(clusterId);
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
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder plans(List<GetBackupPlansPlan> plans) {
            this.plans = Objects.requireNonNull(plans);
            return this;
        }
        public Builder plans(GetBackupPlansPlan... plans) {
            return plans(List.of(plans));
        }
        public GetBackupPlansResult build() {
            final var o = new GetBackupPlansResult();
            o.clusterId = clusterId;
            o.id = id;
            o.ids = ids;
            o.outputFile = outputFile;
            o.plans = plans;
            return o;
        }
    }
}
