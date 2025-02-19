// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mongodb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ShardingInstanceConfigServerList {
    /**
     * @return The connection address of the Config Server node.
     * 
     */
    private @Nullable String connectString;
    /**
     * @return The max connections of the Config Server node.
     * 
     */
    private @Nullable Integer maxConnections;
    /**
     * @return The maximum IOPS of the Config Server node.
     * 
     */
    private @Nullable Integer maxIops;
    /**
     * @return Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
     * 
     */
    private @Nullable String nodeClass;
    /**
     * @return The description of the Config Server node.
     * 
     */
    private @Nullable String nodeDescription;
    /**
     * @return The ID of the Config Server node.
     * 
     */
    private @Nullable String nodeId;
    /**
     * @return - Custom storage space; value range: [10, 1,000]
     * - 10-GB increments. Unit: GB.
     * 
     */
    private @Nullable Integer nodeStorage;
    /**
     * @return The connection port of the Config Server node.
     * 
     */
    private @Nullable Integer port;

    private ShardingInstanceConfigServerList() {}
    /**
     * @return The connection address of the Config Server node.
     * 
     */
    public Optional<String> connectString() {
        return Optional.ofNullable(this.connectString);
    }
    /**
     * @return The max connections of the Config Server node.
     * 
     */
    public Optional<Integer> maxConnections() {
        return Optional.ofNullable(this.maxConnections);
    }
    /**
     * @return The maximum IOPS of the Config Server node.
     * 
     */
    public Optional<Integer> maxIops() {
        return Optional.ofNullable(this.maxIops);
    }
    /**
     * @return Node specification. see [Instance specifications](https://www.alibabacloud.com/help/doc-detail/57141.htm).
     * 
     */
    public Optional<String> nodeClass() {
        return Optional.ofNullable(this.nodeClass);
    }
    /**
     * @return The description of the Config Server node.
     * 
     */
    public Optional<String> nodeDescription() {
        return Optional.ofNullable(this.nodeDescription);
    }
    /**
     * @return The ID of the Config Server node.
     * 
     */
    public Optional<String> nodeId() {
        return Optional.ofNullable(this.nodeId);
    }
    /**
     * @return - Custom storage space; value range: [10, 1,000]
     * - 10-GB increments. Unit: GB.
     * 
     */
    public Optional<Integer> nodeStorage() {
        return Optional.ofNullable(this.nodeStorage);
    }
    /**
     * @return The connection port of the Config Server node.
     * 
     */
    public Optional<Integer> port() {
        return Optional.ofNullable(this.port);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ShardingInstanceConfigServerList defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String connectString;
        private @Nullable Integer maxConnections;
        private @Nullable Integer maxIops;
        private @Nullable String nodeClass;
        private @Nullable String nodeDescription;
        private @Nullable String nodeId;
        private @Nullable Integer nodeStorage;
        private @Nullable Integer port;
        public Builder() {}
        public Builder(ShardingInstanceConfigServerList defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.connectString = defaults.connectString;
    	      this.maxConnections = defaults.maxConnections;
    	      this.maxIops = defaults.maxIops;
    	      this.nodeClass = defaults.nodeClass;
    	      this.nodeDescription = defaults.nodeDescription;
    	      this.nodeId = defaults.nodeId;
    	      this.nodeStorage = defaults.nodeStorage;
    	      this.port = defaults.port;
        }

        @CustomType.Setter
        public Builder connectString(@Nullable String connectString) {
            this.connectString = connectString;
            return this;
        }
        @CustomType.Setter
        public Builder maxConnections(@Nullable Integer maxConnections) {
            this.maxConnections = maxConnections;
            return this;
        }
        @CustomType.Setter
        public Builder maxIops(@Nullable Integer maxIops) {
            this.maxIops = maxIops;
            return this;
        }
        @CustomType.Setter
        public Builder nodeClass(@Nullable String nodeClass) {
            this.nodeClass = nodeClass;
            return this;
        }
        @CustomType.Setter
        public Builder nodeDescription(@Nullable String nodeDescription) {
            this.nodeDescription = nodeDescription;
            return this;
        }
        @CustomType.Setter
        public Builder nodeId(@Nullable String nodeId) {
            this.nodeId = nodeId;
            return this;
        }
        @CustomType.Setter
        public Builder nodeStorage(@Nullable Integer nodeStorage) {
            this.nodeStorage = nodeStorage;
            return this;
        }
        @CustomType.Setter
        public Builder port(@Nullable Integer port) {
            this.port = port;
            return this;
        }
        public ShardingInstanceConfigServerList build() {
            final var o = new ShardingInstanceConfigServerList();
            o.connectString = connectString;
            o.maxConnections = maxConnections;
            o.maxIops = maxIops;
            o.nodeClass = nodeClass;
            o.nodeDescription = nodeDescription;
            o.nodeId = nodeId;
            o.nodeStorage = nodeStorage;
            o.port = port;
            return o;
        }
    }
}
