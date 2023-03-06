// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.emr.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetClustersClusterHostPoolInfo {
    /**
     * @return Machine pool ID.
     * 
     */
    private String hpBizId;
    /**
     * @return The name of the machine pool.
     * 
     */
    private String hpName;

    private GetClustersClusterHostPoolInfo() {}
    /**
     * @return Machine pool ID.
     * 
     */
    public String hpBizId() {
        return this.hpBizId;
    }
    /**
     * @return The name of the machine pool.
     * 
     */
    public String hpName() {
        return this.hpName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClustersClusterHostPoolInfo defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String hpBizId;
        private String hpName;
        public Builder() {}
        public Builder(GetClustersClusterHostPoolInfo defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.hpBizId = defaults.hpBizId;
    	      this.hpName = defaults.hpName;
        }

        @CustomType.Setter
        public Builder hpBizId(String hpBizId) {
            this.hpBizId = Objects.requireNonNull(hpBizId);
            return this;
        }
        @CustomType.Setter
        public Builder hpName(String hpName) {
            this.hpName = Objects.requireNonNull(hpName);
            return this;
        }
        public GetClustersClusterHostPoolInfo build() {
            final var o = new GetClustersClusterHostPoolInfo();
            o.hpBizId = hpBizId;
            o.hpName = hpName;
            return o;
        }
    }
}