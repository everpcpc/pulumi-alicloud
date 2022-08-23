// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.emr.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ClusterConfig {
    /**
     * @return Custom configuration service config key, e.g. ’dfs.replication’.
     * 
     */
    private final String configKey;
    /**
     * @return Custom configuration service config value, e.g. ’3’.
     * 
     */
    private final String configValue;
    /**
     * @return Custom configuration service file name, e.g. ’hdfs-site’.
     * 
     */
    private final String fileName;
    /**
     * @return Cluster service configuration modification name, e.g. ’HDFS’.
     * 
     */
    private final String serviceName;

    @CustomType.Constructor
    private ClusterConfig(
        @CustomType.Parameter("configKey") String configKey,
        @CustomType.Parameter("configValue") String configValue,
        @CustomType.Parameter("fileName") String fileName,
        @CustomType.Parameter("serviceName") String serviceName) {
        this.configKey = configKey;
        this.configValue = configValue;
        this.fileName = fileName;
        this.serviceName = serviceName;
    }

    /**
     * @return Custom configuration service config key, e.g. ’dfs.replication’.
     * 
     */
    public String configKey() {
        return this.configKey;
    }
    /**
     * @return Custom configuration service config value, e.g. ’3’.
     * 
     */
    public String configValue() {
        return this.configValue;
    }
    /**
     * @return Custom configuration service file name, e.g. ’hdfs-site’.
     * 
     */
    public String fileName() {
        return this.fileName;
    }
    /**
     * @return Cluster service configuration modification name, e.g. ’HDFS’.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClusterConfig defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String configKey;
        private String configValue;
        private String fileName;
        private String serviceName;

        public Builder() {
    	      // Empty
        }

        public Builder(ClusterConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.configKey = defaults.configKey;
    	      this.configValue = defaults.configValue;
    	      this.fileName = defaults.fileName;
    	      this.serviceName = defaults.serviceName;
        }

        public Builder configKey(String configKey) {
            this.configKey = Objects.requireNonNull(configKey);
            return this;
        }
        public Builder configValue(String configValue) {
            this.configValue = Objects.requireNonNull(configValue);
            return this;
        }
        public Builder fileName(String fileName) {
            this.fileName = Objects.requireNonNull(fileName);
            return this;
        }
        public Builder serviceName(String serviceName) {
            this.serviceName = Objects.requireNonNull(serviceName);
            return this;
        }        public ClusterConfig build() {
            return new ClusterConfig(configKey, configValue, fileName, serviceName);
        }
    }
}
