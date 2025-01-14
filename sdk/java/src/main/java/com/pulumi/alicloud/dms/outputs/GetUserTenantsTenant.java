// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dms.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetUserTenantsTenant {
    /**
     * @return The user tenant id.
     * 
     */
    private String id;
    /**
     * @return The status of the user tenant.
     * 
     */
    private String status;
    /**
     * @return The name of the user tenant.
     * 
     */
    private String tenantName;
    /**
     * @return The user tenant id. Same as id.
     * 
     */
    private String tid;

    private GetUserTenantsTenant() {}
    /**
     * @return The user tenant id.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The status of the user tenant.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The name of the user tenant.
     * 
     */
    public String tenantName() {
        return this.tenantName;
    }
    /**
     * @return The user tenant id. Same as id.
     * 
     */
    public String tid() {
        return this.tid;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUserTenantsTenant defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String status;
        private String tenantName;
        private String tid;
        public Builder() {}
        public Builder(GetUserTenantsTenant defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.status = defaults.status;
    	      this.tenantName = defaults.tenantName;
    	      this.tid = defaults.tid;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder tenantName(String tenantName) {
            this.tenantName = Objects.requireNonNull(tenantName);
            return this;
        }
        @CustomType.Setter
        public Builder tid(String tid) {
            this.tid = Objects.requireNonNull(tid);
            return this;
        }
        public GetUserTenantsTenant build() {
            final var o = new GetUserTenantsTenant();
            o.id = id;
            o.status = status;
            o.tenantName = tenantName;
            o.tid = tid;
            return o;
        }
    }
}
