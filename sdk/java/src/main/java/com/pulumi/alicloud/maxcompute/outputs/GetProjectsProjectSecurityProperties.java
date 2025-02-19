// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.maxcompute.outputs;

import com.pulumi.alicloud.maxcompute.outputs.GetProjectsProjectSecurityPropertiesProjectProtection;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;

@CustomType
public final class GetProjectsProjectSecurityProperties {
    /**
     * @return Whether to enable download permission check.
     * 
     */
    private Boolean enableDownloadPrivilege;
    /**
     * @return Label authorization.
     * 
     */
    private Boolean labelSecurity;
    /**
     * @return Project creator permissions.
     * 
     */
    private Boolean objectCreatorHasAccessPermission;
    /**
     * @return Does the project creator have authorization rights.
     * 
     */
    private Boolean objectCreatorHasGrantPermission;
    /**
     * @return Project protection.
     * 
     */
    private GetProjectsProjectSecurityPropertiesProjectProtection projectProtection;
    /**
     * @return Whether to turn on ACL.
     * 
     */
    private Boolean usingAcl;
    /**
     * @return Whether to enable Policy.
     * 
     */
    private Boolean usingPolicy;

    private GetProjectsProjectSecurityProperties() {}
    /**
     * @return Whether to enable download permission check.
     * 
     */
    public Boolean enableDownloadPrivilege() {
        return this.enableDownloadPrivilege;
    }
    /**
     * @return Label authorization.
     * 
     */
    public Boolean labelSecurity() {
        return this.labelSecurity;
    }
    /**
     * @return Project creator permissions.
     * 
     */
    public Boolean objectCreatorHasAccessPermission() {
        return this.objectCreatorHasAccessPermission;
    }
    /**
     * @return Does the project creator have authorization rights.
     * 
     */
    public Boolean objectCreatorHasGrantPermission() {
        return this.objectCreatorHasGrantPermission;
    }
    /**
     * @return Project protection.
     * 
     */
    public GetProjectsProjectSecurityPropertiesProjectProtection projectProtection() {
        return this.projectProtection;
    }
    /**
     * @return Whether to turn on ACL.
     * 
     */
    public Boolean usingAcl() {
        return this.usingAcl;
    }
    /**
     * @return Whether to enable Policy.
     * 
     */
    public Boolean usingPolicy() {
        return this.usingPolicy;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectsProjectSecurityProperties defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean enableDownloadPrivilege;
        private Boolean labelSecurity;
        private Boolean objectCreatorHasAccessPermission;
        private Boolean objectCreatorHasGrantPermission;
        private GetProjectsProjectSecurityPropertiesProjectProtection projectProtection;
        private Boolean usingAcl;
        private Boolean usingPolicy;
        public Builder() {}
        public Builder(GetProjectsProjectSecurityProperties defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enableDownloadPrivilege = defaults.enableDownloadPrivilege;
    	      this.labelSecurity = defaults.labelSecurity;
    	      this.objectCreatorHasAccessPermission = defaults.objectCreatorHasAccessPermission;
    	      this.objectCreatorHasGrantPermission = defaults.objectCreatorHasGrantPermission;
    	      this.projectProtection = defaults.projectProtection;
    	      this.usingAcl = defaults.usingAcl;
    	      this.usingPolicy = defaults.usingPolicy;
        }

        @CustomType.Setter
        public Builder enableDownloadPrivilege(Boolean enableDownloadPrivilege) {
            this.enableDownloadPrivilege = Objects.requireNonNull(enableDownloadPrivilege);
            return this;
        }
        @CustomType.Setter
        public Builder labelSecurity(Boolean labelSecurity) {
            this.labelSecurity = Objects.requireNonNull(labelSecurity);
            return this;
        }
        @CustomType.Setter
        public Builder objectCreatorHasAccessPermission(Boolean objectCreatorHasAccessPermission) {
            this.objectCreatorHasAccessPermission = Objects.requireNonNull(objectCreatorHasAccessPermission);
            return this;
        }
        @CustomType.Setter
        public Builder objectCreatorHasGrantPermission(Boolean objectCreatorHasGrantPermission) {
            this.objectCreatorHasGrantPermission = Objects.requireNonNull(objectCreatorHasGrantPermission);
            return this;
        }
        @CustomType.Setter
        public Builder projectProtection(GetProjectsProjectSecurityPropertiesProjectProtection projectProtection) {
            this.projectProtection = Objects.requireNonNull(projectProtection);
            return this;
        }
        @CustomType.Setter
        public Builder usingAcl(Boolean usingAcl) {
            this.usingAcl = Objects.requireNonNull(usingAcl);
            return this;
        }
        @CustomType.Setter
        public Builder usingPolicy(Boolean usingPolicy) {
            this.usingPolicy = Objects.requireNonNull(usingPolicy);
            return this;
        }
        public GetProjectsProjectSecurityProperties build() {
            final var o = new GetProjectsProjectSecurityProperties();
            o.enableDownloadPrivilege = enableDownloadPrivilege;
            o.labelSecurity = labelSecurity;
            o.objectCreatorHasAccessPermission = objectCreatorHasAccessPermission;
            o.objectCreatorHasGrantPermission = objectCreatorHasGrantPermission;
            o.projectProtection = projectProtection;
            o.usingAcl = usingAcl;
            o.usingPolicy = usingPolicy;
            return o;
        }
    }
}
