// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.maxcompute.outputs;

import com.pulumi.alicloud.maxcompute.outputs.GetProjectsProjectIpWhiteList;
import com.pulumi.alicloud.maxcompute.outputs.GetProjectsProjectProperties;
import com.pulumi.alicloud.maxcompute.outputs.GetProjectsProjectSecurityProperties;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectsProject {
    private String comment;
    /**
     * @return Default Computing Resource Group
     * 
     */
    private String defaultQuota;
    /**
     * @return Project ID. The value is the same as `project_name`.
     * 
     */
    private String id;
    /**
     * @return IP whitelist
     * 
     */
    private GetProjectsProjectIpWhiteList ipWhiteList;
    /**
     * @return Project owner
     * 
     */
    private String owner;
    /**
     * @return The name of the resource
     * 
     */
    private String projectName;
    /**
     * @return Project base attributes
     * 
     */
    private GetProjectsProjectProperties properties;
    /**
     * @return Security-related attributes
     * 
     */
    private GetProjectsProjectSecurityProperties securityProperties;
    /**
     * @return The status of the resource
     * 
     */
    private String status;
    /**
     * @return Project type
     * 
     */
    private String type;

    private GetProjectsProject() {}
    public String comment() {
        return this.comment;
    }
    /**
     * @return Default Computing Resource Group
     * 
     */
    public String defaultQuota() {
        return this.defaultQuota;
    }
    /**
     * @return Project ID. The value is the same as `project_name`.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return IP whitelist
     * 
     */
    public GetProjectsProjectIpWhiteList ipWhiteList() {
        return this.ipWhiteList;
    }
    /**
     * @return Project owner
     * 
     */
    public String owner() {
        return this.owner;
    }
    /**
     * @return The name of the resource
     * 
     */
    public String projectName() {
        return this.projectName;
    }
    /**
     * @return Project base attributes
     * 
     */
    public GetProjectsProjectProperties properties() {
        return this.properties;
    }
    /**
     * @return Security-related attributes
     * 
     */
    public GetProjectsProjectSecurityProperties securityProperties() {
        return this.securityProperties;
    }
    /**
     * @return The status of the resource
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return Project type
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectsProject defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String comment;
        private String defaultQuota;
        private String id;
        private GetProjectsProjectIpWhiteList ipWhiteList;
        private String owner;
        private String projectName;
        private GetProjectsProjectProperties properties;
        private GetProjectsProjectSecurityProperties securityProperties;
        private String status;
        private String type;
        public Builder() {}
        public Builder(GetProjectsProject defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.comment = defaults.comment;
    	      this.defaultQuota = defaults.defaultQuota;
    	      this.id = defaults.id;
    	      this.ipWhiteList = defaults.ipWhiteList;
    	      this.owner = defaults.owner;
    	      this.projectName = defaults.projectName;
    	      this.properties = defaults.properties;
    	      this.securityProperties = defaults.securityProperties;
    	      this.status = defaults.status;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder comment(String comment) {
            this.comment = Objects.requireNonNull(comment);
            return this;
        }
        @CustomType.Setter
        public Builder defaultQuota(String defaultQuota) {
            this.defaultQuota = Objects.requireNonNull(defaultQuota);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ipWhiteList(GetProjectsProjectIpWhiteList ipWhiteList) {
            this.ipWhiteList = Objects.requireNonNull(ipWhiteList);
            return this;
        }
        @CustomType.Setter
        public Builder owner(String owner) {
            this.owner = Objects.requireNonNull(owner);
            return this;
        }
        @CustomType.Setter
        public Builder projectName(String projectName) {
            this.projectName = Objects.requireNonNull(projectName);
            return this;
        }
        @CustomType.Setter
        public Builder properties(GetProjectsProjectProperties properties) {
            this.properties = Objects.requireNonNull(properties);
            return this;
        }
        @CustomType.Setter
        public Builder securityProperties(GetProjectsProjectSecurityProperties securityProperties) {
            this.securityProperties = Objects.requireNonNull(securityProperties);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetProjectsProject build() {
            final var o = new GetProjectsProject();
            o.comment = comment;
            o.defaultQuota = defaultQuota;
            o.id = id;
            o.ipWhiteList = ipWhiteList;
            o.owner = owner;
            o.projectName = projectName;
            o.properties = properties;
            o.securityProperties = securityProperties;
            o.status = status;
            o.type = type;
            return o;
        }
    }
}