// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudsso.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetAccessConfigurationsConfigurationPermissionPolicy {
    /**
     * @return The Creation time of policy.
     * 
     */
    private String addTime;
    /**
     * @return The Content of Policy.
     * 
     */
    private String permissionPolicyDocument;
    /**
     * @return The Policy Name of policy.
     * 
     */
    private String permissionPolicyName;
    /**
     * @return The Policy Type of policy. Valid values: `System`, `Inline`.
     * 
     */
    private String permissionPolicyType;

    private GetAccessConfigurationsConfigurationPermissionPolicy() {}
    /**
     * @return The Creation time of policy.
     * 
     */
    public String addTime() {
        return this.addTime;
    }
    /**
     * @return The Content of Policy.
     * 
     */
    public String permissionPolicyDocument() {
        return this.permissionPolicyDocument;
    }
    /**
     * @return The Policy Name of policy.
     * 
     */
    public String permissionPolicyName() {
        return this.permissionPolicyName;
    }
    /**
     * @return The Policy Type of policy. Valid values: `System`, `Inline`.
     * 
     */
    public String permissionPolicyType() {
        return this.permissionPolicyType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAccessConfigurationsConfigurationPermissionPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String addTime;
        private String permissionPolicyDocument;
        private String permissionPolicyName;
        private String permissionPolicyType;
        public Builder() {}
        public Builder(GetAccessConfigurationsConfigurationPermissionPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.addTime = defaults.addTime;
    	      this.permissionPolicyDocument = defaults.permissionPolicyDocument;
    	      this.permissionPolicyName = defaults.permissionPolicyName;
    	      this.permissionPolicyType = defaults.permissionPolicyType;
        }

        @CustomType.Setter
        public Builder addTime(String addTime) {
            this.addTime = Objects.requireNonNull(addTime);
            return this;
        }
        @CustomType.Setter
        public Builder permissionPolicyDocument(String permissionPolicyDocument) {
            this.permissionPolicyDocument = Objects.requireNonNull(permissionPolicyDocument);
            return this;
        }
        @CustomType.Setter
        public Builder permissionPolicyName(String permissionPolicyName) {
            this.permissionPolicyName = Objects.requireNonNull(permissionPolicyName);
            return this;
        }
        @CustomType.Setter
        public Builder permissionPolicyType(String permissionPolicyType) {
            this.permissionPolicyType = Objects.requireNonNull(permissionPolicyType);
            return this;
        }
        public GetAccessConfigurationsConfigurationPermissionPolicy build() {
            final var o = new GetAccessConfigurationsConfigurationPermissionPolicy();
            o.addTime = addTime;
            o.permissionPolicyDocument = permissionPolicyDocument;
            o.permissionPolicyName = permissionPolicyName;
            o.permissionPolicyType = permissionPolicyType;
            return o;
        }
    }
}