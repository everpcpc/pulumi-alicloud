// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.apigateway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetLogConfigsConfig {
    /**
     * @return The ID of the Log Config.
     * 
     */
    private String id;
    /**
     * @return The type the of log.
     * 
     */
    private String logType;
    /**
     * @return The region ID of the Log Config.
     * 
     */
    private String regionId;
    /**
     * @return The name of the Log Store.
     * 
     */
    private String slsLogStore;
    /**
     * @return The name of the Project.
     * 
     */
    private String slsProject;

    private GetLogConfigsConfig() {}
    /**
     * @return The ID of the Log Config.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The type the of log.
     * 
     */
    public String logType() {
        return this.logType;
    }
    /**
     * @return The region ID of the Log Config.
     * 
     */
    public String regionId() {
        return this.regionId;
    }
    /**
     * @return The name of the Log Store.
     * 
     */
    public String slsLogStore() {
        return this.slsLogStore;
    }
    /**
     * @return The name of the Project.
     * 
     */
    public String slsProject() {
        return this.slsProject;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLogConfigsConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String logType;
        private String regionId;
        private String slsLogStore;
        private String slsProject;
        public Builder() {}
        public Builder(GetLogConfigsConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.logType = defaults.logType;
    	      this.regionId = defaults.regionId;
    	      this.slsLogStore = defaults.slsLogStore;
    	      this.slsProject = defaults.slsProject;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder logType(String logType) {
            this.logType = Objects.requireNonNull(logType);
            return this;
        }
        @CustomType.Setter
        public Builder regionId(String regionId) {
            this.regionId = Objects.requireNonNull(regionId);
            return this;
        }
        @CustomType.Setter
        public Builder slsLogStore(String slsLogStore) {
            this.slsLogStore = Objects.requireNonNull(slsLogStore);
            return this;
        }
        @CustomType.Setter
        public Builder slsProject(String slsProject) {
            this.slsProject = Objects.requireNonNull(slsProject);
            return this;
        }
        public GetLogConfigsConfig build() {
            final var o = new GetLogConfigsConfig();
            o.id = id;
            o.logType = logType;
            o.regionId = regionId;
            o.slsLogStore = slsLogStore;
            o.slsProject = slsProject;
            return o;
        }
    }
}
