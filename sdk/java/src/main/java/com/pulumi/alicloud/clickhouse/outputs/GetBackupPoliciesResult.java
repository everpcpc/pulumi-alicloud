// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.clickhouse.outputs;

import com.pulumi.alicloud.clickhouse.outputs.GetBackupPoliciesPolicy;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBackupPoliciesResult {
    private String dbClusterId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String outputFile;
    private List<GetBackupPoliciesPolicy> policies;

    private GetBackupPoliciesResult() {}
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
    public Optional<String> outputFile() {
        return Optional.ofNullable(this.outputFile);
    }
    public List<GetBackupPoliciesPolicy> policies() {
        return this.policies;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBackupPoliciesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String dbClusterId;
        private String id;
        private @Nullable String outputFile;
        private List<GetBackupPoliciesPolicy> policies;
        public Builder() {}
        public Builder(GetBackupPoliciesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dbClusterId = defaults.dbClusterId;
    	      this.id = defaults.id;
    	      this.outputFile = defaults.outputFile;
    	      this.policies = defaults.policies;
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
        public Builder outputFile(@Nullable String outputFile) {
            this.outputFile = outputFile;
            return this;
        }
        @CustomType.Setter
        public Builder policies(List<GetBackupPoliciesPolicy> policies) {
            this.policies = Objects.requireNonNull(policies);
            return this;
        }
        public Builder policies(GetBackupPoliciesPolicy... policies) {
            return policies(List.of(policies));
        }
        public GetBackupPoliciesResult build() {
            final var o = new GetBackupPoliciesResult();
            o.dbClusterId = dbClusterId;
            o.id = id;
            o.outputFile = outputFile;
            o.policies = policies;
            return o;
        }
    }
}
