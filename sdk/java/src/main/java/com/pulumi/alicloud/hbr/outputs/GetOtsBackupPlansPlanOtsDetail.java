// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbr.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetOtsBackupPlansPlanOtsDetail {
    private List<String> tableNames;

    private GetOtsBackupPlansPlanOtsDetail() {}
    public List<String> tableNames() {
        return this.tableNames;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOtsBackupPlansPlanOtsDetail defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> tableNames;
        public Builder() {}
        public Builder(GetOtsBackupPlansPlanOtsDetail defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.tableNames = defaults.tableNames;
        }

        @CustomType.Setter
        public Builder tableNames(List<String> tableNames) {
            this.tableNames = Objects.requireNonNull(tableNames);
            return this;
        }
        public Builder tableNames(String... tableNames) {
            return tableNames(List.of(tableNames));
        }
        public GetOtsBackupPlansPlanOtsDetail build() {
            final var o = new GetOtsBackupPlansPlanOtsDetail();
            o.tableNames = tableNames;
            return o;
        }
    }
}
