// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.adb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.adb.ResourceGroupArgs;
import com.pulumi.alicloud.adb.inputs.ResourceGroupState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Adb Resource Group resource.
 * 
 * For information about Adb Resource Group and how to use it, see [What is Adb Resource Group](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/latest/create-db-resource-group).
 * 
 * &gt; **NOTE:** Available in v1.195.0+.
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
 * import com.pulumi.alicloud.adb.ResourceGroup;
 * import com.pulumi.alicloud.adb.ResourceGroupArgs;
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
 *         var default_ = new ResourceGroup(&#34;default&#34;, ResourceGroupArgs.builder()        
 *             .dbClusterId(&#34;am-bp1a16357gty69185&#34;)
 *             .groupName(&#34;TESTOPENAPI&#34;)
 *             .groupType(&#34;batch&#34;)
 *             .nodeNum(0)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Adb Resource Group can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:adb/resourceGroup:ResourceGroup example &lt;db_cluster_id&gt;:&lt;group_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:adb/resourceGroup:ResourceGroup")
public class ResourceGroup extends com.pulumi.resources.CustomResource {
    /**
     * Creation time.
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return Creation time.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * DB cluster id.
     * 
     */
    @Export(name="dbClusterId", type=String.class, parameters={})
    private Output<String> dbClusterId;

    /**
     * @return DB cluster id.
     * 
     */
    public Output<String> dbClusterId() {
        return this.dbClusterId;
    }
    /**
     * The name of the resource pool. The group name must be 2 to 30 characters in length, and can contain upper case letters, digits, and underscore(_).
     * 
     */
    @Export(name="groupName", type=String.class, parameters={})
    private Output<String> groupName;

    /**
     * @return The name of the resource pool. The group name must be 2 to 30 characters in length, and can contain upper case letters, digits, and underscore(_).
     * 
     */
    public Output<String> groupName() {
        return this.groupName;
    }
    /**
     * Query type, value description:
     * * **etl**: Batch query mode.
     * * **interactive**: interactive Query mode.
     * * **default_type**: the default query mode.
     * 
     */
    @Export(name="groupType", type=String.class, parameters={})
    private Output<String> groupType;

    /**
     * @return Query type, value description:
     * * **etl**: Batch query mode.
     * * **interactive**: interactive Query mode.
     * * **default_type**: the default query mode.
     * 
     */
    public Output<String> groupType() {
        return this.groupType;
    }
    /**
     * The number of nodes. The default number of nodes is 0. The number of nodes must be less than or equal to the number of nodes whose resource name is USER_DEFAULT.
     * 
     */
    @Export(name="nodeNum", type=Integer.class, parameters={})
    private Output<Integer> nodeNum;

    /**
     * @return The number of nodes. The default number of nodes is 0. The number of nodes must be less than or equal to the number of nodes whose resource name is USER_DEFAULT.
     * 
     */
    public Output<Integer> nodeNum() {
        return this.nodeNum;
    }
    /**
     * Update time.
     * 
     */
    @Export(name="updateTime", type=String.class, parameters={})
    private Output<String> updateTime;

    /**
     * @return Update time.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }
    /**
     * Binding User.
     * 
     */
    @Export(name="user", type=String.class, parameters={})
    private Output<String> user;

    /**
     * @return Binding User.
     * 
     */
    public Output<String> user() {
        return this.user;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ResourceGroup(String name) {
        this(name, ResourceGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ResourceGroup(String name, ResourceGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ResourceGroup(String name, ResourceGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:adb/resourceGroup:ResourceGroup", name, args == null ? ResourceGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ResourceGroup(String name, Output<String> id, @Nullable ResourceGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:adb/resourceGroup:ResourceGroup", name, state, makeResourceOptions(options, id));
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
    public static ResourceGroup get(String name, Output<String> id, @Nullable ResourceGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ResourceGroup(name, id, state, options);
    }
}
