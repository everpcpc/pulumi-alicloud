// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eci;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.eci.inputs.GetContainerGroupsArgs;
import com.pulumi.alicloud.eci.inputs.GetContainerGroupsPlainArgs;
import com.pulumi.alicloud.eci.inputs.GetImageCachesArgs;
import com.pulumi.alicloud.eci.inputs.GetImageCachesPlainArgs;
import com.pulumi.alicloud.eci.inputs.GetVirtualNodesArgs;
import com.pulumi.alicloud.eci.inputs.GetVirtualNodesPlainArgs;
import com.pulumi.alicloud.eci.inputs.GetZonesArgs;
import com.pulumi.alicloud.eci.inputs.GetZonesPlainArgs;
import com.pulumi.alicloud.eci.outputs.GetContainerGroupsResult;
import com.pulumi.alicloud.eci.outputs.GetImageCachesResult;
import com.pulumi.alicloud.eci.outputs.GetVirtualNodesResult;
import com.pulumi.alicloud.eci.outputs.GetZonesResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class EciFunctions {
    /**
     * This data source provides the Eci Container Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.111.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetContainerGroupsArgs;
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
     *         final var example = EciFunctions.getContainerGroups(GetContainerGroupsArgs.builder()
     *             .ids(&#34;example_value&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstEciContainerGroupId&#34;, example.applyValue(getContainerGroupsResult -&gt; getContainerGroupsResult.groups()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetContainerGroupsResult> getContainerGroups() {
        return getContainerGroups(GetContainerGroupsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Eci Container Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.111.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetContainerGroupsArgs;
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
     *         final var example = EciFunctions.getContainerGroups(GetContainerGroupsArgs.builder()
     *             .ids(&#34;example_value&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstEciContainerGroupId&#34;, example.applyValue(getContainerGroupsResult -&gt; getContainerGroupsResult.groups()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetContainerGroupsResult> getContainerGroupsPlain() {
        return getContainerGroupsPlain(GetContainerGroupsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Eci Container Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.111.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetContainerGroupsArgs;
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
     *         final var example = EciFunctions.getContainerGroups(GetContainerGroupsArgs.builder()
     *             .ids(&#34;example_value&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstEciContainerGroupId&#34;, example.applyValue(getContainerGroupsResult -&gt; getContainerGroupsResult.groups()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetContainerGroupsResult> getContainerGroups(GetContainerGroupsArgs args) {
        return getContainerGroups(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Eci Container Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.111.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetContainerGroupsArgs;
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
     *         final var example = EciFunctions.getContainerGroups(GetContainerGroupsArgs.builder()
     *             .ids(&#34;example_value&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstEciContainerGroupId&#34;, example.applyValue(getContainerGroupsResult -&gt; getContainerGroupsResult.groups()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetContainerGroupsResult> getContainerGroupsPlain(GetContainerGroupsPlainArgs args) {
        return getContainerGroupsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Eci Container Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.111.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetContainerGroupsArgs;
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
     *         final var example = EciFunctions.getContainerGroups(GetContainerGroupsArgs.builder()
     *             .ids(&#34;example_value&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstEciContainerGroupId&#34;, example.applyValue(getContainerGroupsResult -&gt; getContainerGroupsResult.groups()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetContainerGroupsResult> getContainerGroups(GetContainerGroupsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:eci/getContainerGroups:getContainerGroups", TypeShape.of(GetContainerGroupsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Eci Container Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.111.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetContainerGroupsArgs;
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
     *         final var example = EciFunctions.getContainerGroups(GetContainerGroupsArgs.builder()
     *             .ids(&#34;example_value&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstEciContainerGroupId&#34;, example.applyValue(getContainerGroupsResult -&gt; getContainerGroupsResult.groups()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetContainerGroupsResult> getContainerGroupsPlain(GetContainerGroupsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:eci/getContainerGroups:getContainerGroups", TypeShape.of(GetContainerGroupsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Provides a collection of ECI Image Cache to the specified filters.
     * 
     * &gt; **NOTE:** Available in 1.90.0+.
     * 
     */
    public static Output<GetImageCachesResult> getImageCaches() {
        return getImageCaches(GetImageCachesArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Provides a collection of ECI Image Cache to the specified filters.
     * 
     * &gt; **NOTE:** Available in 1.90.0+.
     * 
     */
    public static CompletableFuture<GetImageCachesResult> getImageCachesPlain() {
        return getImageCachesPlain(GetImageCachesPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Provides a collection of ECI Image Cache to the specified filters.
     * 
     * &gt; **NOTE:** Available in 1.90.0+.
     * 
     */
    public static Output<GetImageCachesResult> getImageCaches(GetImageCachesArgs args) {
        return getImageCaches(args, InvokeOptions.Empty);
    }
    /**
     * Provides a collection of ECI Image Cache to the specified filters.
     * 
     * &gt; **NOTE:** Available in 1.90.0+.
     * 
     */
    public static CompletableFuture<GetImageCachesResult> getImageCachesPlain(GetImageCachesPlainArgs args) {
        return getImageCachesPlain(args, InvokeOptions.Empty);
    }
    /**
     * Provides a collection of ECI Image Cache to the specified filters.
     * 
     * &gt; **NOTE:** Available in 1.90.0+.
     * 
     */
    public static Output<GetImageCachesResult> getImageCaches(GetImageCachesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:eci/getImageCaches:getImageCaches", TypeShape.of(GetImageCachesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Provides a collection of ECI Image Cache to the specified filters.
     * 
     * &gt; **NOTE:** Available in 1.90.0+.
     * 
     */
    public static CompletableFuture<GetImageCachesResult> getImageCachesPlain(GetImageCachesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:eci/getImageCaches:getImageCaches", TypeShape.of(GetImageCachesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Eci Virtual Nodes of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.145.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetVirtualNodesArgs;
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
     *         final var ids = EciFunctions.getVirtualNodes(GetVirtualNodesArgs.builder()
     *             .ids(            
     *                 &#34;example_value-1&#34;,
     *                 &#34;example_value-2&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eciVirtualNodeId1&#34;, ids.applyValue(getVirtualNodesResult -&gt; getVirtualNodesResult.nodes()[0].id()));
     *         final var nameRegex = EciFunctions.getVirtualNodes(GetVirtualNodesArgs.builder()
     *             .nameRegex(&#34;^my-VirtualNode&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eciVirtualNodeId2&#34;, nameRegex.applyValue(getVirtualNodesResult -&gt; getVirtualNodesResult.nodes()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetVirtualNodesResult> getVirtualNodes() {
        return getVirtualNodes(GetVirtualNodesArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Eci Virtual Nodes of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.145.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetVirtualNodesArgs;
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
     *         final var ids = EciFunctions.getVirtualNodes(GetVirtualNodesArgs.builder()
     *             .ids(            
     *                 &#34;example_value-1&#34;,
     *                 &#34;example_value-2&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eciVirtualNodeId1&#34;, ids.applyValue(getVirtualNodesResult -&gt; getVirtualNodesResult.nodes()[0].id()));
     *         final var nameRegex = EciFunctions.getVirtualNodes(GetVirtualNodesArgs.builder()
     *             .nameRegex(&#34;^my-VirtualNode&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eciVirtualNodeId2&#34;, nameRegex.applyValue(getVirtualNodesResult -&gt; getVirtualNodesResult.nodes()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetVirtualNodesResult> getVirtualNodesPlain() {
        return getVirtualNodesPlain(GetVirtualNodesPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Eci Virtual Nodes of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.145.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetVirtualNodesArgs;
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
     *         final var ids = EciFunctions.getVirtualNodes(GetVirtualNodesArgs.builder()
     *             .ids(            
     *                 &#34;example_value-1&#34;,
     *                 &#34;example_value-2&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eciVirtualNodeId1&#34;, ids.applyValue(getVirtualNodesResult -&gt; getVirtualNodesResult.nodes()[0].id()));
     *         final var nameRegex = EciFunctions.getVirtualNodes(GetVirtualNodesArgs.builder()
     *             .nameRegex(&#34;^my-VirtualNode&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eciVirtualNodeId2&#34;, nameRegex.applyValue(getVirtualNodesResult -&gt; getVirtualNodesResult.nodes()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetVirtualNodesResult> getVirtualNodes(GetVirtualNodesArgs args) {
        return getVirtualNodes(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Eci Virtual Nodes of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.145.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetVirtualNodesArgs;
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
     *         final var ids = EciFunctions.getVirtualNodes(GetVirtualNodesArgs.builder()
     *             .ids(            
     *                 &#34;example_value-1&#34;,
     *                 &#34;example_value-2&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eciVirtualNodeId1&#34;, ids.applyValue(getVirtualNodesResult -&gt; getVirtualNodesResult.nodes()[0].id()));
     *         final var nameRegex = EciFunctions.getVirtualNodes(GetVirtualNodesArgs.builder()
     *             .nameRegex(&#34;^my-VirtualNode&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eciVirtualNodeId2&#34;, nameRegex.applyValue(getVirtualNodesResult -&gt; getVirtualNodesResult.nodes()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetVirtualNodesResult> getVirtualNodesPlain(GetVirtualNodesPlainArgs args) {
        return getVirtualNodesPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Eci Virtual Nodes of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.145.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetVirtualNodesArgs;
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
     *         final var ids = EciFunctions.getVirtualNodes(GetVirtualNodesArgs.builder()
     *             .ids(            
     *                 &#34;example_value-1&#34;,
     *                 &#34;example_value-2&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eciVirtualNodeId1&#34;, ids.applyValue(getVirtualNodesResult -&gt; getVirtualNodesResult.nodes()[0].id()));
     *         final var nameRegex = EciFunctions.getVirtualNodes(GetVirtualNodesArgs.builder()
     *             .nameRegex(&#34;^my-VirtualNode&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eciVirtualNodeId2&#34;, nameRegex.applyValue(getVirtualNodesResult -&gt; getVirtualNodesResult.nodes()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetVirtualNodesResult> getVirtualNodes(GetVirtualNodesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:eci/getVirtualNodes:getVirtualNodes", TypeShape.of(GetVirtualNodesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Eci Virtual Nodes of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.145.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetVirtualNodesArgs;
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
     *         final var ids = EciFunctions.getVirtualNodes(GetVirtualNodesArgs.builder()
     *             .ids(            
     *                 &#34;example_value-1&#34;,
     *                 &#34;example_value-2&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eciVirtualNodeId1&#34;, ids.applyValue(getVirtualNodesResult -&gt; getVirtualNodesResult.nodes()[0].id()));
     *         final var nameRegex = EciFunctions.getVirtualNodes(GetVirtualNodesArgs.builder()
     *             .nameRegex(&#34;^my-VirtualNode&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;eciVirtualNodeId2&#34;, nameRegex.applyValue(getVirtualNodesResult -&gt; getVirtualNodesResult.nodes()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetVirtualNodesResult> getVirtualNodesPlain(GetVirtualNodesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:eci/getVirtualNodes:getVirtualNodes", TypeShape.of(GetVirtualNodesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the available zones with the Application Load Balancer (ALB) Instance of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.145.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetZonesArgs;
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
     *         final var default = EciFunctions.getZones();
     * 
     *         ctx.export(&#34;firstEciZonesId&#34;, default_.zones()[0].zoneIds()[0]);
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetZonesResult> getZones() {
        return getZones(GetZonesArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the available zones with the Application Load Balancer (ALB) Instance of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.145.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetZonesArgs;
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
     *         final var default = EciFunctions.getZones();
     * 
     *         ctx.export(&#34;firstEciZonesId&#34;, default_.zones()[0].zoneIds()[0]);
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetZonesResult> getZonesPlain() {
        return getZonesPlain(GetZonesPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the available zones with the Application Load Balancer (ALB) Instance of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.145.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetZonesArgs;
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
     *         final var default = EciFunctions.getZones();
     * 
     *         ctx.export(&#34;firstEciZonesId&#34;, default_.zones()[0].zoneIds()[0]);
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetZonesResult> getZones(GetZonesArgs args) {
        return getZones(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the available zones with the Application Load Balancer (ALB) Instance of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.145.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetZonesArgs;
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
     *         final var default = EciFunctions.getZones();
     * 
     *         ctx.export(&#34;firstEciZonesId&#34;, default_.zones()[0].zoneIds()[0]);
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetZonesResult> getZonesPlain(GetZonesPlainArgs args) {
        return getZonesPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the available zones with the Application Load Balancer (ALB) Instance of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.145.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetZonesArgs;
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
     *         final var default = EciFunctions.getZones();
     * 
     *         ctx.export(&#34;firstEciZonesId&#34;, default_.zones()[0].zoneIds()[0]);
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetZonesResult> getZones(GetZonesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:eci/getZones:getZones", TypeShape.of(GetZonesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the available zones with the Application Load Balancer (ALB) Instance of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.145.0+.
     * 
     * ## Example Usage
     * 
     * Basic Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.eci.EciFunctions;
     * import com.pulumi.alicloud.eci.inputs.GetZonesArgs;
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
     *         final var default = EciFunctions.getZones();
     * 
     *         ctx.export(&#34;firstEciZonesId&#34;, default_.zones()[0].zoneIds()[0]);
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetZonesResult> getZonesPlain(GetZonesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:eci/getZones:getZones", TypeShape.of(GetZonesResult.class), args, Utilities.withVersion(options));
    }
}
