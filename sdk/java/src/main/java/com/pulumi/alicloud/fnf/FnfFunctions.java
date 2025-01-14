// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.fnf;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.fnf.inputs.GetExecutionsArgs;
import com.pulumi.alicloud.fnf.inputs.GetExecutionsPlainArgs;
import com.pulumi.alicloud.fnf.inputs.GetFlowsArgs;
import com.pulumi.alicloud.fnf.inputs.GetFlowsPlainArgs;
import com.pulumi.alicloud.fnf.inputs.GetSchedulesArgs;
import com.pulumi.alicloud.fnf.inputs.GetSchedulesPlainArgs;
import com.pulumi.alicloud.fnf.inputs.GetServiceArgs;
import com.pulumi.alicloud.fnf.inputs.GetServicePlainArgs;
import com.pulumi.alicloud.fnf.outputs.GetExecutionsResult;
import com.pulumi.alicloud.fnf.outputs.GetFlowsResult;
import com.pulumi.alicloud.fnf.outputs.GetSchedulesResult;
import com.pulumi.alicloud.fnf.outputs.GetServiceResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class FnfFunctions {
    /**
     * This data source provides the FnF Executions of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.149.0+.
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
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetExecutionsArgs;
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
     *         final var ids = FnfFunctions.getExecutions(GetExecutionsArgs.builder()
     *             .flowName(&#34;example_value&#34;)
     *             .ids(            
     *                 &#34;my-Execution-1&#34;,
     *                 &#34;my-Execution-2&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;fnfExecutionId1&#34;, data.alicloud_fn_f_executions().ids().executions()[0].id());
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetExecutionsResult> getExecutions(GetExecutionsArgs args) {
        return getExecutions(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the FnF Executions of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.149.0+.
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
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetExecutionsArgs;
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
     *         final var ids = FnfFunctions.getExecutions(GetExecutionsArgs.builder()
     *             .flowName(&#34;example_value&#34;)
     *             .ids(            
     *                 &#34;my-Execution-1&#34;,
     *                 &#34;my-Execution-2&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;fnfExecutionId1&#34;, data.alicloud_fn_f_executions().ids().executions()[0].id());
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetExecutionsResult> getExecutionsPlain(GetExecutionsPlainArgs args) {
        return getExecutionsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the FnF Executions of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.149.0+.
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
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetExecutionsArgs;
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
     *         final var ids = FnfFunctions.getExecutions(GetExecutionsArgs.builder()
     *             .flowName(&#34;example_value&#34;)
     *             .ids(            
     *                 &#34;my-Execution-1&#34;,
     *                 &#34;my-Execution-2&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;fnfExecutionId1&#34;, data.alicloud_fn_f_executions().ids().executions()[0].id());
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetExecutionsResult> getExecutions(GetExecutionsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:fnf/getExecutions:getExecutions", TypeShape.of(GetExecutionsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the FnF Executions of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.149.0+.
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
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetExecutionsArgs;
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
     *         final var ids = FnfFunctions.getExecutions(GetExecutionsArgs.builder()
     *             .flowName(&#34;example_value&#34;)
     *             .ids(            
     *                 &#34;my-Execution-1&#34;,
     *                 &#34;my-Execution-2&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;fnfExecutionId1&#34;, data.alicloud_fn_f_executions().ids().executions()[0].id());
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetExecutionsResult> getExecutionsPlain(GetExecutionsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:fnf/getExecutions:getExecutions", TypeShape.of(GetExecutionsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Fnf Flows of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.105.0+.
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
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetFlowsArgs;
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
     *         final var example = FnfFunctions.getFlows(GetFlowsArgs.builder()
     *             .ids(&#34;example_value&#34;)
     *             .nameRegex(&#34;the_resource_name&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstFnfFlowId&#34;, example.applyValue(getFlowsResult -&gt; getFlowsResult.flows()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetFlowsResult> getFlows() {
        return getFlows(GetFlowsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Fnf Flows of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.105.0+.
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
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetFlowsArgs;
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
     *         final var example = FnfFunctions.getFlows(GetFlowsArgs.builder()
     *             .ids(&#34;example_value&#34;)
     *             .nameRegex(&#34;the_resource_name&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstFnfFlowId&#34;, example.applyValue(getFlowsResult -&gt; getFlowsResult.flows()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetFlowsResult> getFlowsPlain() {
        return getFlowsPlain(GetFlowsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Fnf Flows of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.105.0+.
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
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetFlowsArgs;
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
     *         final var example = FnfFunctions.getFlows(GetFlowsArgs.builder()
     *             .ids(&#34;example_value&#34;)
     *             .nameRegex(&#34;the_resource_name&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstFnfFlowId&#34;, example.applyValue(getFlowsResult -&gt; getFlowsResult.flows()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetFlowsResult> getFlows(GetFlowsArgs args) {
        return getFlows(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Fnf Flows of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.105.0+.
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
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetFlowsArgs;
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
     *         final var example = FnfFunctions.getFlows(GetFlowsArgs.builder()
     *             .ids(&#34;example_value&#34;)
     *             .nameRegex(&#34;the_resource_name&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstFnfFlowId&#34;, example.applyValue(getFlowsResult -&gt; getFlowsResult.flows()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetFlowsResult> getFlowsPlain(GetFlowsPlainArgs args) {
        return getFlowsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Fnf Flows of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.105.0+.
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
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetFlowsArgs;
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
     *         final var example = FnfFunctions.getFlows(GetFlowsArgs.builder()
     *             .ids(&#34;example_value&#34;)
     *             .nameRegex(&#34;the_resource_name&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstFnfFlowId&#34;, example.applyValue(getFlowsResult -&gt; getFlowsResult.flows()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetFlowsResult> getFlows(GetFlowsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:fnf/getFlows:getFlows", TypeShape.of(GetFlowsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Fnf Flows of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.105.0+.
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
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetFlowsArgs;
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
     *         final var example = FnfFunctions.getFlows(GetFlowsArgs.builder()
     *             .ids(&#34;example_value&#34;)
     *             .nameRegex(&#34;the_resource_name&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstFnfFlowId&#34;, example.applyValue(getFlowsResult -&gt; getFlowsResult.flows()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetFlowsResult> getFlowsPlain(GetFlowsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:fnf/getFlows:getFlows", TypeShape.of(GetFlowsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Fnf Schedules of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.105.0+.
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
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetSchedulesArgs;
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
     *         final var example = FnfFunctions.getSchedules(GetSchedulesArgs.builder()
     *             .flowName(&#34;example_value&#34;)
     *             .ids(&#34;example_value&#34;)
     *             .nameRegex(&#34;the_resource_name&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstFnfScheduleId&#34;, example.applyValue(getSchedulesResult -&gt; getSchedulesResult.schedules()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetSchedulesResult> getSchedules(GetSchedulesArgs args) {
        return getSchedules(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Fnf Schedules of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.105.0+.
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
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetSchedulesArgs;
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
     *         final var example = FnfFunctions.getSchedules(GetSchedulesArgs.builder()
     *             .flowName(&#34;example_value&#34;)
     *             .ids(&#34;example_value&#34;)
     *             .nameRegex(&#34;the_resource_name&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstFnfScheduleId&#34;, example.applyValue(getSchedulesResult -&gt; getSchedulesResult.schedules()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetSchedulesResult> getSchedulesPlain(GetSchedulesPlainArgs args) {
        return getSchedulesPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Fnf Schedules of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.105.0+.
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
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetSchedulesArgs;
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
     *         final var example = FnfFunctions.getSchedules(GetSchedulesArgs.builder()
     *             .flowName(&#34;example_value&#34;)
     *             .ids(&#34;example_value&#34;)
     *             .nameRegex(&#34;the_resource_name&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstFnfScheduleId&#34;, example.applyValue(getSchedulesResult -&gt; getSchedulesResult.schedules()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetSchedulesResult> getSchedules(GetSchedulesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:fnf/getSchedules:getSchedules", TypeShape.of(GetSchedulesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Fnf Schedules of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.105.0+.
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
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetSchedulesArgs;
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
     *         final var example = FnfFunctions.getSchedules(GetSchedulesArgs.builder()
     *             .flowName(&#34;example_value&#34;)
     *             .ids(&#34;example_value&#34;)
     *             .nameRegex(&#34;the_resource_name&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;firstFnfScheduleId&#34;, example.applyValue(getSchedulesResult -&gt; getSchedulesResult.schedules()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetSchedulesResult> getSchedulesPlain(GetSchedulesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:fnf/getSchedules:getSchedules", TypeShape.of(GetSchedulesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Using this data source can open Fnf service automatically. If the service has been opened, it will return opened.
     * 
     * For information about Fnf and how to use it, see [What is Fnf](https://www.alibabacloud.com/help/en/product/113549.htm).
     * 
     * &gt; **NOTE:** Available in v1.114.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetServiceArgs;
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
     *         final var open = FnfFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetServiceResult> getService() {
        return getService(GetServiceArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Using this data source can open Fnf service automatically. If the service has been opened, it will return opened.
     * 
     * For information about Fnf and how to use it, see [What is Fnf](https://www.alibabacloud.com/help/en/product/113549.htm).
     * 
     * &gt; **NOTE:** Available in v1.114.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetServiceArgs;
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
     *         final var open = FnfFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetServiceResult> getServicePlain() {
        return getServicePlain(GetServicePlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Using this data source can open Fnf service automatically. If the service has been opened, it will return opened.
     * 
     * For information about Fnf and how to use it, see [What is Fnf](https://www.alibabacloud.com/help/en/product/113549.htm).
     * 
     * &gt; **NOTE:** Available in v1.114.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetServiceArgs;
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
     *         final var open = FnfFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetServiceResult> getService(GetServiceArgs args) {
        return getService(args, InvokeOptions.Empty);
    }
    /**
     * Using this data source can open Fnf service automatically. If the service has been opened, it will return opened.
     * 
     * For information about Fnf and how to use it, see [What is Fnf](https://www.alibabacloud.com/help/en/product/113549.htm).
     * 
     * &gt; **NOTE:** Available in v1.114.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetServiceArgs;
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
     *         final var open = FnfFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetServiceResult> getServicePlain(GetServicePlainArgs args) {
        return getServicePlain(args, InvokeOptions.Empty);
    }
    /**
     * Using this data source can open Fnf service automatically. If the service has been opened, it will return opened.
     * 
     * For information about Fnf and how to use it, see [What is Fnf](https://www.alibabacloud.com/help/en/product/113549.htm).
     * 
     * &gt; **NOTE:** Available in v1.114.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetServiceArgs;
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
     *         final var open = FnfFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetServiceResult> getService(GetServiceArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:fnf/getService:getService", TypeShape.of(GetServiceResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Using this data source can open Fnf service automatically. If the service has been opened, it will return opened.
     * 
     * For information about Fnf and how to use it, see [What is Fnf](https://www.alibabacloud.com/help/en/product/113549.htm).
     * 
     * &gt; **NOTE:** Available in v1.114.0+
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.alicloud.fnf.FnfFunctions;
     * import com.pulumi.alicloud.fnf.inputs.GetServiceArgs;
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
     *         final var open = FnfFunctions.getService(GetServiceArgs.builder()
     *             .enable(&#34;On&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetServiceResult> getServicePlain(GetServicePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:fnf/getService:getService", TypeShape.of(GetServiceResult.class), args, Utilities.withVersion(options));
    }
}
