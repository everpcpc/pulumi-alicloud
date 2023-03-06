// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbase;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.hbase.inputs.GetInstanceTypesArgs;
import com.pulumi.alicloud.hbase.inputs.GetInstanceTypesPlainArgs;
import com.pulumi.alicloud.hbase.inputs.GetInstancesArgs;
import com.pulumi.alicloud.hbase.inputs.GetInstancesPlainArgs;
import com.pulumi.alicloud.hbase.inputs.GetZonesArgs;
import com.pulumi.alicloud.hbase.inputs.GetZonesPlainArgs;
import com.pulumi.alicloud.hbase.outputs.GetInstanceTypesResult;
import com.pulumi.alicloud.hbase.outputs.GetInstancesResult;
import com.pulumi.alicloud.hbase.outputs.GetZonesResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class HbaseFunctions {
    /**
     * This data source provides availability instance_types for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.106.0+.
     * 
     */
    public static Output<GetInstanceTypesResult> getInstanceTypes() {
        return getInstanceTypes(GetInstanceTypesArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides availability instance_types for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.106.0+.
     * 
     */
    public static CompletableFuture<GetInstanceTypesResult> getInstanceTypesPlain() {
        return getInstanceTypesPlain(GetInstanceTypesPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides availability instance_types for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.106.0+.
     * 
     */
    public static Output<GetInstanceTypesResult> getInstanceTypes(GetInstanceTypesArgs args) {
        return getInstanceTypes(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides availability instance_types for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.106.0+.
     * 
     */
    public static CompletableFuture<GetInstanceTypesResult> getInstanceTypesPlain(GetInstanceTypesPlainArgs args) {
        return getInstanceTypesPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides availability instance_types for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.106.0+.
     * 
     */
    public static Output<GetInstanceTypesResult> getInstanceTypes(GetInstanceTypesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:hbase/getInstanceTypes:getInstanceTypes", TypeShape.of(GetInstanceTypesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides availability instance_types for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.106.0+.
     * 
     */
    public static CompletableFuture<GetInstanceTypesResult> getInstanceTypesPlain(GetInstanceTypesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:hbase/getInstanceTypes:getInstanceTypes", TypeShape.of(GetInstanceTypesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * The `alicloud.hbase.getInstances` data source provides a collection of HBase instances available in Alicloud account.
     * Filters support regular expression for the instance name, ids or availability_zone.
     * 
     * &gt; **NOTE:**  Available in 1.67.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.hbase.HbaseFunctions;
     * import com.pulumi.alicloud.hbase.inputs.GetInstancesArgs;
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
     *         final var hbase = HbaseFunctions.getInstances(GetInstancesArgs.builder()
     *             .availabilityZone(&#34;cn-shenzhen-b&#34;)
     *             .nameRegex(&#34;tf_testAccHBase&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetInstancesResult> getInstances() {
        return getInstances(GetInstancesArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * The `alicloud.hbase.getInstances` data source provides a collection of HBase instances available in Alicloud account.
     * Filters support regular expression for the instance name, ids or availability_zone.
     * 
     * &gt; **NOTE:**  Available in 1.67.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.hbase.HbaseFunctions;
     * import com.pulumi.alicloud.hbase.inputs.GetInstancesArgs;
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
     *         final var hbase = HbaseFunctions.getInstances(GetInstancesArgs.builder()
     *             .availabilityZone(&#34;cn-shenzhen-b&#34;)
     *             .nameRegex(&#34;tf_testAccHBase&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetInstancesResult> getInstancesPlain() {
        return getInstancesPlain(GetInstancesPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * The `alicloud.hbase.getInstances` data source provides a collection of HBase instances available in Alicloud account.
     * Filters support regular expression for the instance name, ids or availability_zone.
     * 
     * &gt; **NOTE:**  Available in 1.67.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.hbase.HbaseFunctions;
     * import com.pulumi.alicloud.hbase.inputs.GetInstancesArgs;
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
     *         final var hbase = HbaseFunctions.getInstances(GetInstancesArgs.builder()
     *             .availabilityZone(&#34;cn-shenzhen-b&#34;)
     *             .nameRegex(&#34;tf_testAccHBase&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetInstancesResult> getInstances(GetInstancesArgs args) {
        return getInstances(args, InvokeOptions.Empty);
    }
    /**
     * The `alicloud.hbase.getInstances` data source provides a collection of HBase instances available in Alicloud account.
     * Filters support regular expression for the instance name, ids or availability_zone.
     * 
     * &gt; **NOTE:**  Available in 1.67.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.hbase.HbaseFunctions;
     * import com.pulumi.alicloud.hbase.inputs.GetInstancesArgs;
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
     *         final var hbase = HbaseFunctions.getInstances(GetInstancesArgs.builder()
     *             .availabilityZone(&#34;cn-shenzhen-b&#34;)
     *             .nameRegex(&#34;tf_testAccHBase&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetInstancesResult> getInstancesPlain(GetInstancesPlainArgs args) {
        return getInstancesPlain(args, InvokeOptions.Empty);
    }
    /**
     * The `alicloud.hbase.getInstances` data source provides a collection of HBase instances available in Alicloud account.
     * Filters support regular expression for the instance name, ids or availability_zone.
     * 
     * &gt; **NOTE:**  Available in 1.67.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.hbase.HbaseFunctions;
     * import com.pulumi.alicloud.hbase.inputs.GetInstancesArgs;
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
     *         final var hbase = HbaseFunctions.getInstances(GetInstancesArgs.builder()
     *             .availabilityZone(&#34;cn-shenzhen-b&#34;)
     *             .nameRegex(&#34;tf_testAccHBase&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetInstancesResult> getInstances(GetInstancesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:hbase/getInstances:getInstances", TypeShape.of(GetInstancesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * The `alicloud.hbase.getInstances` data source provides a collection of HBase instances available in Alicloud account.
     * Filters support regular expression for the instance name, ids or availability_zone.
     * 
     * &gt; **NOTE:**  Available in 1.67.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.hbase.HbaseFunctions;
     * import com.pulumi.alicloud.hbase.inputs.GetInstancesArgs;
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
     *         final var hbase = HbaseFunctions.getInstances(GetInstancesArgs.builder()
     *             .availabilityZone(&#34;cn-shenzhen-b&#34;)
     *             .nameRegex(&#34;tf_testAccHBase&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetInstancesResult> getInstancesPlain(GetInstancesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:hbase/getInstances:getInstances", TypeShape.of(GetInstancesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides availability zones for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.73.0+.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.hbase.HbaseFunctions;
     * import com.pulumi.alicloud.hbase.inputs.GetZonesArgs;
     * import com.pulumi.alicloud.hbase.Instance;
     * import com.pulumi.alicloud.hbase.InstanceArgs;
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
     *         final var zonesIds = HbaseFunctions.getZones();
     * 
     *         var hbase = new Instance(&#34;hbase&#34;, InstanceArgs.builder()        
     *             .zoneId(zonesIds.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetZonesResult> getZones() {
        return getZones(GetZonesArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides availability zones for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.73.0+.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.hbase.HbaseFunctions;
     * import com.pulumi.alicloud.hbase.inputs.GetZonesArgs;
     * import com.pulumi.alicloud.hbase.Instance;
     * import com.pulumi.alicloud.hbase.InstanceArgs;
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
     *         final var zonesIds = HbaseFunctions.getZones();
     * 
     *         var hbase = new Instance(&#34;hbase&#34;, InstanceArgs.builder()        
     *             .zoneId(zonesIds.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetZonesResult> getZonesPlain() {
        return getZonesPlain(GetZonesPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides availability zones for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.73.0+.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.hbase.HbaseFunctions;
     * import com.pulumi.alicloud.hbase.inputs.GetZonesArgs;
     * import com.pulumi.alicloud.hbase.Instance;
     * import com.pulumi.alicloud.hbase.InstanceArgs;
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
     *         final var zonesIds = HbaseFunctions.getZones();
     * 
     *         var hbase = new Instance(&#34;hbase&#34;, InstanceArgs.builder()        
     *             .zoneId(zonesIds.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetZonesResult> getZones(GetZonesArgs args) {
        return getZones(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides availability zones for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.73.0+.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.hbase.HbaseFunctions;
     * import com.pulumi.alicloud.hbase.inputs.GetZonesArgs;
     * import com.pulumi.alicloud.hbase.Instance;
     * import com.pulumi.alicloud.hbase.InstanceArgs;
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
     *         final var zonesIds = HbaseFunctions.getZones();
     * 
     *         var hbase = new Instance(&#34;hbase&#34;, InstanceArgs.builder()        
     *             .zoneId(zonesIds.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetZonesResult> getZonesPlain(GetZonesPlainArgs args) {
        return getZonesPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides availability zones for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.73.0+.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.hbase.HbaseFunctions;
     * import com.pulumi.alicloud.hbase.inputs.GetZonesArgs;
     * import com.pulumi.alicloud.hbase.Instance;
     * import com.pulumi.alicloud.hbase.InstanceArgs;
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
     *         final var zonesIds = HbaseFunctions.getZones();
     * 
     *         var hbase = new Instance(&#34;hbase&#34;, InstanceArgs.builder()        
     *             .zoneId(zonesIds.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetZonesResult> getZones(GetZonesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:hbase/getZones:getZones", TypeShape.of(GetZonesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides availability zones for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
     * 
     * &gt; **NOTE:** Available in v1.73.0+.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.hbase.HbaseFunctions;
     * import com.pulumi.alicloud.hbase.inputs.GetZonesArgs;
     * import com.pulumi.alicloud.hbase.Instance;
     * import com.pulumi.alicloud.hbase.InstanceArgs;
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
     *         final var zonesIds = HbaseFunctions.getZones();
     * 
     *         var hbase = new Instance(&#34;hbase&#34;, InstanceArgs.builder()        
     *             .zoneId(zonesIds.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetZonesResult> getZonesPlain(GetZonesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:hbase/getZones:getZones", TypeShape.of(GetZonesResult.class), args, Utilities.withVersion(options));
    }
}