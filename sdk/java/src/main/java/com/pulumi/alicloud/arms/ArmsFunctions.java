// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.arms;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.arms.inputs.GetAlertContactGroupsArgs;
import com.pulumi.alicloud.arms.inputs.GetAlertContactGroupsPlainArgs;
import com.pulumi.alicloud.arms.inputs.GetAlertContactsArgs;
import com.pulumi.alicloud.arms.inputs.GetAlertContactsPlainArgs;
import com.pulumi.alicloud.arms.inputs.GetDispatchRulesArgs;
import com.pulumi.alicloud.arms.inputs.GetDispatchRulesPlainArgs;
import com.pulumi.alicloud.arms.inputs.GetPrometheusAlertRulesArgs;
import com.pulumi.alicloud.arms.inputs.GetPrometheusAlertRulesPlainArgs;
import com.pulumi.alicloud.arms.outputs.GetAlertContactGroupsResult;
import com.pulumi.alicloud.arms.outputs.GetAlertContactsResult;
import com.pulumi.alicloud.arms.outputs.GetDispatchRulesResult;
import com.pulumi.alicloud.arms.outputs.GetPrometheusAlertRulesResult;
import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import java.util.concurrent.CompletableFuture;

public final class ArmsFunctions {
    /**
     * This data source provides the Arms Alert Contact Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.131.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetAlertContactGroupsArgs;
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
     *         final var nameRegex = ArmsFunctions.getAlertContactGroups(GetAlertContactGroupsArgs.builder()
     *             .nameRegex(&#34;^my-AlertContactGroup&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsAlertContactGroupId&#34;, nameRegex.applyValue(getAlertContactGroupsResult -&gt; getAlertContactGroupsResult.groups()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetAlertContactGroupsResult> getAlertContactGroups() {
        return getAlertContactGroups(GetAlertContactGroupsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Arms Alert Contact Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.131.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetAlertContactGroupsArgs;
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
     *         final var nameRegex = ArmsFunctions.getAlertContactGroups(GetAlertContactGroupsArgs.builder()
     *             .nameRegex(&#34;^my-AlertContactGroup&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsAlertContactGroupId&#34;, nameRegex.applyValue(getAlertContactGroupsResult -&gt; getAlertContactGroupsResult.groups()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetAlertContactGroupsResult> getAlertContactGroupsPlain() {
        return getAlertContactGroupsPlain(GetAlertContactGroupsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Arms Alert Contact Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.131.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetAlertContactGroupsArgs;
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
     *         final var nameRegex = ArmsFunctions.getAlertContactGroups(GetAlertContactGroupsArgs.builder()
     *             .nameRegex(&#34;^my-AlertContactGroup&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsAlertContactGroupId&#34;, nameRegex.applyValue(getAlertContactGroupsResult -&gt; getAlertContactGroupsResult.groups()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetAlertContactGroupsResult> getAlertContactGroups(GetAlertContactGroupsArgs args) {
        return getAlertContactGroups(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Arms Alert Contact Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.131.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetAlertContactGroupsArgs;
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
     *         final var nameRegex = ArmsFunctions.getAlertContactGroups(GetAlertContactGroupsArgs.builder()
     *             .nameRegex(&#34;^my-AlertContactGroup&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsAlertContactGroupId&#34;, nameRegex.applyValue(getAlertContactGroupsResult -&gt; getAlertContactGroupsResult.groups()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetAlertContactGroupsResult> getAlertContactGroupsPlain(GetAlertContactGroupsPlainArgs args) {
        return getAlertContactGroupsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Arms Alert Contact Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.131.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetAlertContactGroupsArgs;
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
     *         final var nameRegex = ArmsFunctions.getAlertContactGroups(GetAlertContactGroupsArgs.builder()
     *             .nameRegex(&#34;^my-AlertContactGroup&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsAlertContactGroupId&#34;, nameRegex.applyValue(getAlertContactGroupsResult -&gt; getAlertContactGroupsResult.groups()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetAlertContactGroupsResult> getAlertContactGroups(GetAlertContactGroupsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:arms/getAlertContactGroups:getAlertContactGroups", TypeShape.of(GetAlertContactGroupsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Arms Alert Contact Groups of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.131.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetAlertContactGroupsArgs;
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
     *         final var nameRegex = ArmsFunctions.getAlertContactGroups(GetAlertContactGroupsArgs.builder()
     *             .nameRegex(&#34;^my-AlertContactGroup&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsAlertContactGroupId&#34;, nameRegex.applyValue(getAlertContactGroupsResult -&gt; getAlertContactGroupsResult.groups()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetAlertContactGroupsResult> getAlertContactGroupsPlain(GetAlertContactGroupsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:arms/getAlertContactGroups:getAlertContactGroups", TypeShape.of(GetAlertContactGroupsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Arms Alert Contacts of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.129.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetAlertContactsArgs;
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
     *         final var ids = ArmsFunctions.getAlertContacts();
     * 
     *         ctx.export(&#34;armsAlertContactId1&#34;, ids.applyValue(getAlertContactsResult -&gt; getAlertContactsResult.contacts()[0].id()));
     *         final var nameRegex = ArmsFunctions.getAlertContacts(GetAlertContactsArgs.builder()
     *             .nameRegex(&#34;^my-AlertContact&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsAlertContactId2&#34;, nameRegex.applyValue(getAlertContactsResult -&gt; getAlertContactsResult.contacts()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetAlertContactsResult> getAlertContacts() {
        return getAlertContacts(GetAlertContactsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Arms Alert Contacts of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.129.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetAlertContactsArgs;
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
     *         final var ids = ArmsFunctions.getAlertContacts();
     * 
     *         ctx.export(&#34;armsAlertContactId1&#34;, ids.applyValue(getAlertContactsResult -&gt; getAlertContactsResult.contacts()[0].id()));
     *         final var nameRegex = ArmsFunctions.getAlertContacts(GetAlertContactsArgs.builder()
     *             .nameRegex(&#34;^my-AlertContact&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsAlertContactId2&#34;, nameRegex.applyValue(getAlertContactsResult -&gt; getAlertContactsResult.contacts()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetAlertContactsResult> getAlertContactsPlain() {
        return getAlertContactsPlain(GetAlertContactsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Arms Alert Contacts of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.129.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetAlertContactsArgs;
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
     *         final var ids = ArmsFunctions.getAlertContacts();
     * 
     *         ctx.export(&#34;armsAlertContactId1&#34;, ids.applyValue(getAlertContactsResult -&gt; getAlertContactsResult.contacts()[0].id()));
     *         final var nameRegex = ArmsFunctions.getAlertContacts(GetAlertContactsArgs.builder()
     *             .nameRegex(&#34;^my-AlertContact&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsAlertContactId2&#34;, nameRegex.applyValue(getAlertContactsResult -&gt; getAlertContactsResult.contacts()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetAlertContactsResult> getAlertContacts(GetAlertContactsArgs args) {
        return getAlertContacts(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Arms Alert Contacts of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.129.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetAlertContactsArgs;
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
     *         final var ids = ArmsFunctions.getAlertContacts();
     * 
     *         ctx.export(&#34;armsAlertContactId1&#34;, ids.applyValue(getAlertContactsResult -&gt; getAlertContactsResult.contacts()[0].id()));
     *         final var nameRegex = ArmsFunctions.getAlertContacts(GetAlertContactsArgs.builder()
     *             .nameRegex(&#34;^my-AlertContact&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsAlertContactId2&#34;, nameRegex.applyValue(getAlertContactsResult -&gt; getAlertContactsResult.contacts()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetAlertContactsResult> getAlertContactsPlain(GetAlertContactsPlainArgs args) {
        return getAlertContactsPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Arms Alert Contacts of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.129.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetAlertContactsArgs;
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
     *         final var ids = ArmsFunctions.getAlertContacts();
     * 
     *         ctx.export(&#34;armsAlertContactId1&#34;, ids.applyValue(getAlertContactsResult -&gt; getAlertContactsResult.contacts()[0].id()));
     *         final var nameRegex = ArmsFunctions.getAlertContacts(GetAlertContactsArgs.builder()
     *             .nameRegex(&#34;^my-AlertContact&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsAlertContactId2&#34;, nameRegex.applyValue(getAlertContactsResult -&gt; getAlertContactsResult.contacts()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetAlertContactsResult> getAlertContacts(GetAlertContactsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:arms/getAlertContacts:getAlertContacts", TypeShape.of(GetAlertContactsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Arms Alert Contacts of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.129.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetAlertContactsArgs;
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
     *         final var ids = ArmsFunctions.getAlertContacts();
     * 
     *         ctx.export(&#34;armsAlertContactId1&#34;, ids.applyValue(getAlertContactsResult -&gt; getAlertContactsResult.contacts()[0].id()));
     *         final var nameRegex = ArmsFunctions.getAlertContacts(GetAlertContactsArgs.builder()
     *             .nameRegex(&#34;^my-AlertContact&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsAlertContactId2&#34;, nameRegex.applyValue(getAlertContactsResult -&gt; getAlertContactsResult.contacts()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetAlertContactsResult> getAlertContactsPlain(GetAlertContactsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:arms/getAlertContacts:getAlertContacts", TypeShape.of(GetAlertContactsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Arms Dispatch Rules of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetDispatchRulesArgs;
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
     *         final var ids = ArmsFunctions.getDispatchRules();
     * 
     *         ctx.export(&#34;armsDispatchRuleId1&#34;, ids.applyValue(getDispatchRulesResult -&gt; getDispatchRulesResult.rules()[0].id()));
     *         final var nameRegex = ArmsFunctions.getDispatchRules(GetDispatchRulesArgs.builder()
     *             .nameRegex(&#34;^my-DispatchRule&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsDispatchRuleId2&#34;, nameRegex.applyValue(getDispatchRulesResult -&gt; getDispatchRulesResult.rules()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetDispatchRulesResult> getDispatchRules() {
        return getDispatchRules(GetDispatchRulesArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Arms Dispatch Rules of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetDispatchRulesArgs;
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
     *         final var ids = ArmsFunctions.getDispatchRules();
     * 
     *         ctx.export(&#34;armsDispatchRuleId1&#34;, ids.applyValue(getDispatchRulesResult -&gt; getDispatchRulesResult.rules()[0].id()));
     *         final var nameRegex = ArmsFunctions.getDispatchRules(GetDispatchRulesArgs.builder()
     *             .nameRegex(&#34;^my-DispatchRule&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsDispatchRuleId2&#34;, nameRegex.applyValue(getDispatchRulesResult -&gt; getDispatchRulesResult.rules()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetDispatchRulesResult> getDispatchRulesPlain() {
        return getDispatchRulesPlain(GetDispatchRulesPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Arms Dispatch Rules of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetDispatchRulesArgs;
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
     *         final var ids = ArmsFunctions.getDispatchRules();
     * 
     *         ctx.export(&#34;armsDispatchRuleId1&#34;, ids.applyValue(getDispatchRulesResult -&gt; getDispatchRulesResult.rules()[0].id()));
     *         final var nameRegex = ArmsFunctions.getDispatchRules(GetDispatchRulesArgs.builder()
     *             .nameRegex(&#34;^my-DispatchRule&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsDispatchRuleId2&#34;, nameRegex.applyValue(getDispatchRulesResult -&gt; getDispatchRulesResult.rules()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetDispatchRulesResult> getDispatchRules(GetDispatchRulesArgs args) {
        return getDispatchRules(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Arms Dispatch Rules of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetDispatchRulesArgs;
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
     *         final var ids = ArmsFunctions.getDispatchRules();
     * 
     *         ctx.export(&#34;armsDispatchRuleId1&#34;, ids.applyValue(getDispatchRulesResult -&gt; getDispatchRulesResult.rules()[0].id()));
     *         final var nameRegex = ArmsFunctions.getDispatchRules(GetDispatchRulesArgs.builder()
     *             .nameRegex(&#34;^my-DispatchRule&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsDispatchRuleId2&#34;, nameRegex.applyValue(getDispatchRulesResult -&gt; getDispatchRulesResult.rules()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetDispatchRulesResult> getDispatchRulesPlain(GetDispatchRulesPlainArgs args) {
        return getDispatchRulesPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Arms Dispatch Rules of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetDispatchRulesArgs;
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
     *         final var ids = ArmsFunctions.getDispatchRules();
     * 
     *         ctx.export(&#34;armsDispatchRuleId1&#34;, ids.applyValue(getDispatchRulesResult -&gt; getDispatchRulesResult.rules()[0].id()));
     *         final var nameRegex = ArmsFunctions.getDispatchRules(GetDispatchRulesArgs.builder()
     *             .nameRegex(&#34;^my-DispatchRule&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsDispatchRuleId2&#34;, nameRegex.applyValue(getDispatchRulesResult -&gt; getDispatchRulesResult.rules()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetDispatchRulesResult> getDispatchRules(GetDispatchRulesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:arms/getDispatchRules:getDispatchRules", TypeShape.of(GetDispatchRulesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Arms Dispatch Rules of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetDispatchRulesArgs;
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
     *         final var ids = ArmsFunctions.getDispatchRules();
     * 
     *         ctx.export(&#34;armsDispatchRuleId1&#34;, ids.applyValue(getDispatchRulesResult -&gt; getDispatchRulesResult.rules()[0].id()));
     *         final var nameRegex = ArmsFunctions.getDispatchRules(GetDispatchRulesArgs.builder()
     *             .nameRegex(&#34;^my-DispatchRule&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsDispatchRuleId2&#34;, nameRegex.applyValue(getDispatchRulesResult -&gt; getDispatchRulesResult.rules()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetDispatchRulesResult> getDispatchRulesPlain(GetDispatchRulesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:arms/getDispatchRules:getDispatchRules", TypeShape.of(GetDispatchRulesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Arms Prometheus Alert Rules of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetPrometheusAlertRulesArgs;
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
     *         final var ids = ArmsFunctions.getPrometheusAlertRules(GetPrometheusAlertRulesArgs.builder()
     *             .clusterId(&#34;example_value&#34;)
     *             .ids(            
     *                 &#34;example_value-1&#34;,
     *                 &#34;example_value-2&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheusAlertRuleId1&#34;, ids.applyValue(getPrometheusAlertRulesResult -&gt; getPrometheusAlertRulesResult.rules()[0].id()));
     *         final var nameRegex = ArmsFunctions.getPrometheusAlertRules(GetPrometheusAlertRulesArgs.builder()
     *             .clusterId(&#34;example_value&#34;)
     *             .nameRegex(&#34;^my-PrometheusAlertRule&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheusAlertRuleId2&#34;, nameRegex.applyValue(getPrometheusAlertRulesResult -&gt; getPrometheusAlertRulesResult.rules()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetPrometheusAlertRulesResult> getPrometheusAlertRules(GetPrometheusAlertRulesArgs args) {
        return getPrometheusAlertRules(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Arms Prometheus Alert Rules of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetPrometheusAlertRulesArgs;
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
     *         final var ids = ArmsFunctions.getPrometheusAlertRules(GetPrometheusAlertRulesArgs.builder()
     *             .clusterId(&#34;example_value&#34;)
     *             .ids(            
     *                 &#34;example_value-1&#34;,
     *                 &#34;example_value-2&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheusAlertRuleId1&#34;, ids.applyValue(getPrometheusAlertRulesResult -&gt; getPrometheusAlertRulesResult.rules()[0].id()));
     *         final var nameRegex = ArmsFunctions.getPrometheusAlertRules(GetPrometheusAlertRulesArgs.builder()
     *             .clusterId(&#34;example_value&#34;)
     *             .nameRegex(&#34;^my-PrometheusAlertRule&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheusAlertRuleId2&#34;, nameRegex.applyValue(getPrometheusAlertRulesResult -&gt; getPrometheusAlertRulesResult.rules()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetPrometheusAlertRulesResult> getPrometheusAlertRulesPlain(GetPrometheusAlertRulesPlainArgs args) {
        return getPrometheusAlertRulesPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source provides the Arms Prometheus Alert Rules of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetPrometheusAlertRulesArgs;
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
     *         final var ids = ArmsFunctions.getPrometheusAlertRules(GetPrometheusAlertRulesArgs.builder()
     *             .clusterId(&#34;example_value&#34;)
     *             .ids(            
     *                 &#34;example_value-1&#34;,
     *                 &#34;example_value-2&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheusAlertRuleId1&#34;, ids.applyValue(getPrometheusAlertRulesResult -&gt; getPrometheusAlertRulesResult.rules()[0].id()));
     *         final var nameRegex = ArmsFunctions.getPrometheusAlertRules(GetPrometheusAlertRulesArgs.builder()
     *             .clusterId(&#34;example_value&#34;)
     *             .nameRegex(&#34;^my-PrometheusAlertRule&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheusAlertRuleId2&#34;, nameRegex.applyValue(getPrometheusAlertRulesResult -&gt; getPrometheusAlertRulesResult.rules()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetPrometheusAlertRulesResult> getPrometheusAlertRules(GetPrometheusAlertRulesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("alicloud:arms/getPrometheusAlertRules:getPrometheusAlertRules", TypeShape.of(GetPrometheusAlertRulesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source provides the Arms Prometheus Alert Rules of the current Alibaba Cloud user.
     * 
     * &gt; **NOTE:** Available in v1.136.0+.
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
     * import com.pulumi.alicloud.arms.ArmsFunctions;
     * import com.pulumi.alicloud.arms.inputs.GetPrometheusAlertRulesArgs;
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
     *         final var ids = ArmsFunctions.getPrometheusAlertRules(GetPrometheusAlertRulesArgs.builder()
     *             .clusterId(&#34;example_value&#34;)
     *             .ids(            
     *                 &#34;example_value-1&#34;,
     *                 &#34;example_value-2&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheusAlertRuleId1&#34;, ids.applyValue(getPrometheusAlertRulesResult -&gt; getPrometheusAlertRulesResult.rules()[0].id()));
     *         final var nameRegex = ArmsFunctions.getPrometheusAlertRules(GetPrometheusAlertRulesArgs.builder()
     *             .clusterId(&#34;example_value&#34;)
     *             .nameRegex(&#34;^my-PrometheusAlertRule&#34;)
     *             .build());
     * 
     *         ctx.export(&#34;armsPrometheusAlertRuleId2&#34;, nameRegex.applyValue(getPrometheusAlertRulesResult -&gt; getPrometheusAlertRulesResult.rules()[0].id()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetPrometheusAlertRulesResult> getPrometheusAlertRulesPlain(GetPrometheusAlertRulesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("alicloud:arms/getPrometheusAlertRules:getPrometheusAlertRules", TypeShape.of(GetPrometheusAlertRulesResult.class), args, Utilities.withVersion(options));
    }
}