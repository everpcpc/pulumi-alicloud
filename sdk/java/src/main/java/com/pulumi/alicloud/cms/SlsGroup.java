// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cms.SlsGroupArgs;
import com.pulumi.alicloud.cms.inputs.SlsGroupState;
import com.pulumi.alicloud.cms.outputs.SlsGroupSlsGroupConfig;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Monitor Service Sls Group resource.
 * 
 * For information about Cloud Monitor Service Sls Group and how to use it, see [What is Sls Group](https://www.alibabacloud.com/help/doc-detail/28608.htm).
 * 
 * &gt; **NOTE:** Available in v1.171.0+.
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.cms.SlsGroup;
 * import com.pulumi.alicloud.cms.SlsGroupArgs;
 * import com.pulumi.alicloud.cms.inputs.SlsGroupSlsGroupConfigArgs;
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
 *         final var this = AlicloudFunctions.getAccount();
 * 
 *         var default_ = new SlsGroup(&#34;default&#34;, SlsGroupArgs.builder()        
 *             .slsGroupConfigs(SlsGroupSlsGroupConfigArgs.builder()
 *                 .slsUserId(this_.id())
 *                 .slsLogstore(&#34;Logstore-ECS&#34;)
 *                 .slsProject(&#34;aliyun-project&#34;)
 *                 .slsRegion(&#34;cn-hangzhou&#34;)
 *                 .build())
 *             .slsGroupDescription(var_.name())
 *             .slsGroupName(var_.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Cloud Monitor Service Sls Group can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:cms/slsGroup:SlsGroup example &lt;sls_group_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cms/slsGroup:SlsGroup")
public class SlsGroup extends com.pulumi.resources.CustomResource {
    /**
     * The Config of the Sls Group. You can specify up to 25 Config. See the following `Block sls_group_config`.
     * 
     */
    @Export(name="slsGroupConfigs", type=List.class, parameters={SlsGroupSlsGroupConfig.class})
    private Output<List<SlsGroupSlsGroupConfig>> slsGroupConfigs;

    /**
     * @return The Config of the Sls Group. You can specify up to 25 Config. See the following `Block sls_group_config`.
     * 
     */
    public Output<List<SlsGroupSlsGroupConfig>> slsGroupConfigs() {
        return this.slsGroupConfigs;
    }
    /**
     * The Description of the Sls Group.
     * 
     */
    @Export(name="slsGroupDescription", type=String.class, parameters={})
    private Output</* @Nullable */ String> slsGroupDescription;

    /**
     * @return The Description of the Sls Group.
     * 
     */
    public Output<Optional<String>> slsGroupDescription() {
        return Codegen.optional(this.slsGroupDescription);
    }
    /**
     * The name of the resource. The name must be `2` to `32` characters in length, and can contain letters, digits and underscores (_). It must start with a letter.
     * 
     */
    @Export(name="slsGroupName", type=String.class, parameters={})
    private Output<String> slsGroupName;

    /**
     * @return The name of the resource. The name must be `2` to `32` characters in length, and can contain letters, digits and underscores (_). It must start with a letter.
     * 
     */
    public Output<String> slsGroupName() {
        return this.slsGroupName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SlsGroup(String name) {
        this(name, SlsGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SlsGroup(String name, SlsGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SlsGroup(String name, SlsGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cms/slsGroup:SlsGroup", name, args == null ? SlsGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SlsGroup(String name, Output<String> id, @Nullable SlsGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cms/slsGroup:SlsGroup", name, state, makeResourceOptions(options, id));
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
    public static SlsGroup get(String name, Output<String> id, @Nullable SlsGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SlsGroup(name, id, state, options);
    }
}
