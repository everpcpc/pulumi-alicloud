// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ots.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetInstancesInstance {
    /**
     * @return The cluster type of the instance. Possible values: `SSD`, `HYBRID`.
     * 
     */
    private String clusterType;
    /**
     * @return The create time of the instance.
     * 
     */
    private String createTime;
    /**
     * @return The description of the instance.
     * 
     */
    private String description;
    /**
     * @return The instance quota which indicating the maximum number of tables.
     * 
     */
    private Integer entityQuota;
    /**
     * @return ID of the instance.
     * 
     */
    private String id;
    /**
     * @return Instance name.
     * 
     */
    private String name;
    /**
     * @return The network type of the instance. Possible values: `NORMAL`, `VPC`, `VPC_CONSOLE`.
     * 
     */
    private String network;
    /**
     * @return The maximum adjustable read capacity unit of the instance.
     * 
     */
    private Integer readCapacity;
    /**
     * @return Instance status. Possible values: `Running`, `Disabled`, `Deleting`.
     * 
     */
    private String status;
    /**
     * @return A map of tags assigned to the instance. It must be in the format:
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.ots.OtsFunctions;
     * import com.pulumi.alicloud.ots.inputs.GetInstancesArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var instancesDs = OtsFunctions.getInstances(GetInstancesArgs.builder()
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;tagKey1&#34;, &#34;tagValue1&#34;),
     *                 Map.entry(&#34;tagKey2&#34;, &#34;tagValue2&#34;)
     *             ))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    private Map<String,Object> tags;
    /**
     * @return The user id of the instance.
     * 
     */
    private String userId;
    /**
     * @return The maximum adjustable write capacity unit of the instance.
     * 
     */
    private Integer writeCapacity;

    private GetInstancesInstance() {}
    /**
     * @return The cluster type of the instance. Possible values: `SSD`, `HYBRID`.
     * 
     */
    public String clusterType() {
        return this.clusterType;
    }
    /**
     * @return The create time of the instance.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The description of the instance.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The instance quota which indicating the maximum number of tables.
     * 
     */
    public Integer entityQuota() {
        return this.entityQuota;
    }
    /**
     * @return ID of the instance.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Instance name.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The network type of the instance. Possible values: `NORMAL`, `VPC`, `VPC_CONSOLE`.
     * 
     */
    public String network() {
        return this.network;
    }
    /**
     * @return The maximum adjustable read capacity unit of the instance.
     * 
     */
    public Integer readCapacity() {
        return this.readCapacity;
    }
    /**
     * @return Instance status. Possible values: `Running`, `Disabled`, `Deleting`.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return A map of tags assigned to the instance. It must be in the format:
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.ots.OtsFunctions;
     * import com.pulumi.alicloud.ots.inputs.GetInstancesArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var instancesDs = OtsFunctions.getInstances(GetInstancesArgs.builder()
     *             .tags(Map.ofEntries(
     *                 Map.entry(&#34;tagKey1&#34;, &#34;tagValue1&#34;),
     *                 Map.entry(&#34;tagKey2&#34;, &#34;tagValue2&#34;)
     *             ))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public Map<String,Object> tags() {
        return this.tags;
    }
    /**
     * @return The user id of the instance.
     * 
     */
    public String userId() {
        return this.userId;
    }
    /**
     * @return The maximum adjustable write capacity unit of the instance.
     * 
     */
    public Integer writeCapacity() {
        return this.writeCapacity;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstancesInstance defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterType;
        private String createTime;
        private String description;
        private Integer entityQuota;
        private String id;
        private String name;
        private String network;
        private Integer readCapacity;
        private String status;
        private Map<String,Object> tags;
        private String userId;
        private Integer writeCapacity;
        public Builder() {}
        public Builder(GetInstancesInstance defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterType = defaults.clusterType;
    	      this.createTime = defaults.createTime;
    	      this.description = defaults.description;
    	      this.entityQuota = defaults.entityQuota;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.network = defaults.network;
    	      this.readCapacity = defaults.readCapacity;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.userId = defaults.userId;
    	      this.writeCapacity = defaults.writeCapacity;
        }

        @CustomType.Setter
        public Builder clusterType(String clusterType) {
            this.clusterType = Objects.requireNonNull(clusterType);
            return this;
        }
        @CustomType.Setter
        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder entityQuota(Integer entityQuota) {
            this.entityQuota = Objects.requireNonNull(entityQuota);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder network(String network) {
            this.network = Objects.requireNonNull(network);
            return this;
        }
        @CustomType.Setter
        public Builder readCapacity(Integer readCapacity) {
            this.readCapacity = Objects.requireNonNull(readCapacity);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,Object> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        @CustomType.Setter
        public Builder userId(String userId) {
            this.userId = Objects.requireNonNull(userId);
            return this;
        }
        @CustomType.Setter
        public Builder writeCapacity(Integer writeCapacity) {
            this.writeCapacity = Objects.requireNonNull(writeCapacity);
            return this;
        }
        public GetInstancesInstance build() {
            final var o = new GetInstancesInstance();
            o.clusterType = clusterType;
            o.createTime = createTime;
            o.description = description;
            o.entityQuota = entityQuota;
            o.id = id;
            o.name = name;
            o.network = network;
            o.readCapacity = readCapacity;
            o.status = status;
            o.tags = tags;
            o.userId = userId;
            o.writeCapacity = writeCapacity;
            return o;
        }
    }
}