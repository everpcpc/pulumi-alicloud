// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.log.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectsProject {
    /**
     * @return The description of the project.
     * 
     */
    private String description;
    /**
     * @return The ID of the project.
     * 
     */
    private String id;
    /**
     * @return The last modify time of project.
     * 
     */
    private String lastModifyTime;
    /**
     * @return The owner of project.
     * 
     */
    private String owner;
    /**
     * @return The policy of project.
     * 
     */
    private String policy;
    /**
     * @return The name of the project.
     * 
     */
    private String projectName;
    /**
     * @return The region of project.
     * 
     */
    private String region;
    /**
     * @return The status of project.
     * 
     */
    private String status;

    private GetProjectsProject() {}
    /**
     * @return The description of the project.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The ID of the project.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The last modify time of project.
     * 
     */
    public String lastModifyTime() {
        return this.lastModifyTime;
    }
    /**
     * @return The owner of project.
     * 
     */
    public String owner() {
        return this.owner;
    }
    /**
     * @return The policy of project.
     * 
     */
    public String policy() {
        return this.policy;
    }
    /**
     * @return The name of the project.
     * 
     */
    public String projectName() {
        return this.projectName;
    }
    /**
     * @return The region of project.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return The status of project.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectsProject defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String id;
        private String lastModifyTime;
        private String owner;
        private String policy;
        private String projectName;
        private String region;
        private String status;
        public Builder() {}
        public Builder(GetProjectsProject defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.lastModifyTime = defaults.lastModifyTime;
    	      this.owner = defaults.owner;
    	      this.policy = defaults.policy;
    	      this.projectName = defaults.projectName;
    	      this.region = defaults.region;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder lastModifyTime(String lastModifyTime) {
            this.lastModifyTime = Objects.requireNonNull(lastModifyTime);
            return this;
        }
        @CustomType.Setter
        public Builder owner(String owner) {
            this.owner = Objects.requireNonNull(owner);
            return this;
        }
        @CustomType.Setter
        public Builder policy(String policy) {
            this.policy = Objects.requireNonNull(policy);
            return this;
        }
        @CustomType.Setter
        public Builder projectName(String projectName) {
            this.projectName = Objects.requireNonNull(projectName);
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetProjectsProject build() {
            final var o = new GetProjectsProject();
            o.description = description;
            o.id = id;
            o.lastModifyTime = lastModifyTime;
            o.owner = owner;
            o.policy = policy;
            o.projectName = projectName;
            o.region = region;
            o.status = status;
            return o;
        }
    }
}
