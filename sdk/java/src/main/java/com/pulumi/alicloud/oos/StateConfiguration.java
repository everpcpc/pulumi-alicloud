// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.oos;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.oos.StateConfigurationArgs;
import com.pulumi.alicloud.oos.inputs.StateConfigurationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a OOS State Configuration resource.
 * 
 * For information about OOS State Configuration and how to use it, see [What is State Configuration](https://www.alibabacloud.com/help/en/doc-detail/208728.html).
 * 
 * &gt; **NOTE:** Available in v1.147.0+.
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
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.oos.StateConfiguration;
 * import com.pulumi.alicloud.oos.StateConfigurationArgs;
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
 *         final var defaultResourceGroups = ResourcemanagerFunctions.getResourceGroups();
 * 
 *         var defaultStateConfiguration = new StateConfiguration(&#34;defaultStateConfiguration&#34;, StateConfigurationArgs.builder()        
 *             .templateName(&#34;ACS-ECS-InventoryDataCollection&#34;)
 *             .configureMode(&#34;ApplyOnly&#34;)
 *             .description(var_.name())
 *             .scheduleType(&#34;rate&#34;)
 *             .scheduleExpression(&#34;1 hour&#34;)
 *             .resourceGroupId(defaultResourceGroups.applyValue(getResourceGroupsResult -&gt; getResourceGroupsResult.ids()[0]))
 *             .targets(&#34;{\&#34;Filters\&#34;: [{\&#34;Type\&#34;: \&#34;All\&#34;, \&#34;Parameters\&#34;: {\&#34;InstanceChargeType\&#34;: \&#34;PrePaid\&#34;}}], \&#34;ResourceType\&#34;: \&#34;ALIYUN::ECS::Instance\&#34;}&#34;)
 *             .parameters(&#34;{\&#34;policy\&#34;: {\&#34;ACS:Application\&#34;: {\&#34;Collection\&#34;: \&#34;Enabled\&#34;}}}&#34;)
 *             .tags(Map.ofEntries(
 *                 Map.entry(&#34;Created&#34;, &#34;TF&#34;),
 *                 Map.entry(&#34;For&#34;, &#34;Test&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * OOS State Configuration can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:oos/stateConfiguration:StateConfiguration example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:oos/stateConfiguration:StateConfiguration")
public class StateConfiguration extends com.pulumi.resources.CustomResource {
    /**
     * Configuration mode. Valid values: `ApplyAndAutoCorrect`, `ApplyAndMonitor`, `ApplyOnly`.
     * 
     */
    @Export(name="configureMode", type=String.class, parameters={})
    private Output<String> configureMode;

    /**
     * @return Configuration mode. Valid values: `ApplyAndAutoCorrect`, `ApplyAndMonitor`, `ApplyOnly`.
     * 
     */
    public Output<String> configureMode() {
        return this.configureMode;
    }
    /**
     * The description of the resource.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The parameter of the Template. This field is in the format of JSON strings. For detailed definition instructions, please refer to [Metadata types that are supported by a configuration list](https://www.alibabacloud.com/help/en/doc-detail/208276.html).
     * 
     */
    @Export(name="parameters", type=String.class, parameters={})
    private Output</* @Nullable */ String> parameters;

    /**
     * @return The parameter of the Template. This field is in the format of JSON strings. For detailed definition instructions, please refer to [Metadata types that are supported by a configuration list](https://www.alibabacloud.com/help/en/doc-detail/208276.html).
     * 
     */
    public Output<Optional<String>> parameters() {
        return Codegen.optional(this.parameters);
    }
    /**
     * The ID of the resource group.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * Timing expression.
     * 
     */
    @Export(name="scheduleExpression", type=String.class, parameters={})
    private Output<String> scheduleExpression;

    /**
     * @return Timing expression.
     * 
     */
    public Output<String> scheduleExpression() {
        return this.scheduleExpression;
    }
    /**
     * Timing type. Valid values: `rate`.
     * 
     */
    @Export(name="scheduleType", type=String.class, parameters={})
    private Output<String> scheduleType;

    /**
     * @return Timing type. Valid values: `rate`.
     * 
     */
    public Output<String> scheduleType() {
        return this.scheduleType;
    }
    /**
     * The tag of the resource.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return The tag of the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The Target resources.  This field is in the format of JSON strings. For detailed definition instructions, please refer to [Parameter](https://www.alibabacloud.com/help/en/doc-detail/120674.html).
     * 
     */
    @Export(name="targets", type=String.class, parameters={})
    private Output<String> targets;

    /**
     * @return The Target resources.  This field is in the format of JSON strings. For detailed definition instructions, please refer to [Parameter](https://www.alibabacloud.com/help/en/doc-detail/120674.html).
     * 
     */
    public Output<String> targets() {
        return this.targets;
    }
    /**
     * The name of the template.
     * 
     */
    @Export(name="templateName", type=String.class, parameters={})
    private Output<String> templateName;

    /**
     * @return The name of the template.
     * 
     */
    public Output<String> templateName() {
        return this.templateName;
    }
    /**
     * The version number. If you do not specify this parameter, the system uses the latest version.
     * 
     */
    @Export(name="templateVersion", type=String.class, parameters={})
    private Output<String> templateVersion;

    /**
     * @return The version number. If you do not specify this parameter, the system uses the latest version.
     * 
     */
    public Output<String> templateVersion() {
        return this.templateVersion;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public StateConfiguration(String name) {
        this(name, StateConfigurationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public StateConfiguration(String name, StateConfigurationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public StateConfiguration(String name, StateConfigurationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:oos/stateConfiguration:StateConfiguration", name, args == null ? StateConfigurationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private StateConfiguration(String name, Output<String> id, @Nullable StateConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:oos/stateConfiguration:StateConfiguration", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static StateConfiguration get(String name, Output<String> id, @Nullable StateConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new StateConfiguration(name, id, state, options);
    }
}