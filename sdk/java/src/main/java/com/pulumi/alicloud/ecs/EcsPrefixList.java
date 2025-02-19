// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.EcsPrefixListArgs;
import com.pulumi.alicloud.ecs.inputs.EcsPrefixListState;
import com.pulumi.alicloud.ecs.outputs.EcsPrefixListEntry;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ECS Prefix List resource.
 * 
 * For information about ECS Prefix List and how to use it, see [What is Prefix List.](https://www.alibabacloud.com/help/en/doc-detail/207969.html).
 * 
 * &gt; **NOTE:** Available in v1.152.0+.
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
 * import com.pulumi.alicloud.ecs.EcsPrefixList;
 * import com.pulumi.alicloud.ecs.EcsPrefixListArgs;
 * import com.pulumi.alicloud.ecs.inputs.EcsPrefixListEntryArgs;
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
 *         var default_ = new EcsPrefixList(&#34;default&#34;, EcsPrefixListArgs.builder()        
 *             .addressFamily(&#34;IPv4&#34;)
 *             .description(&#34;description&#34;)
 *             .entries(EcsPrefixListEntryArgs.builder()
 *                 .cidr(&#34;192.168.0.0/24&#34;)
 *                 .description(&#34;description&#34;)
 *                 .build())
 *             .maxEntries(2)
 *             .prefixListName(&#34;tftest&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ECS Prefix List can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ecs/ecsPrefixList:EcsPrefixList example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/ecsPrefixList:EcsPrefixList")
public class EcsPrefixList extends com.pulumi.resources.CustomResource {
    /**
     * The IP address family. Valid values: `IPv4`,`IPv6`.
     * 
     */
    @Export(name="addressFamily", type=String.class, parameters={})
    private Output<String> addressFamily;

    /**
     * @return The IP address family. Valid values: `IPv4`,`IPv6`.
     * 
     */
    public Output<String> addressFamily() {
        return this.addressFamily;
    }
    /**
     * The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The Entry. The details see Block `entry`.
     * 
     */
    @Export(name="entries", type=List.class, parameters={EcsPrefixListEntry.class})
    private Output<List<EcsPrefixListEntry>> entries;

    /**
     * @return The Entry. The details see Block `entry`.
     * 
     */
    public Output<List<EcsPrefixListEntry>> entries() {
        return this.entries;
    }
    /**
     * The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
     * 
     */
    @Export(name="maxEntries", type=Integer.class, parameters={})
    private Output<Integer> maxEntries;

    /**
     * @return The maximum number of entries that the prefix list can contain.  Valid values: 1 to 200.
     * 
     */
    public Output<Integer> maxEntries() {
        return this.maxEntries;
    }
    /**
     * The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
     * 
     */
    @Export(name="prefixListName", type=String.class, parameters={})
    private Output<String> prefixListName;

    /**
     * @return The name of the prefix. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
     * 
     */
    public Output<String> prefixListName() {
        return this.prefixListName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EcsPrefixList(String name) {
        this(name, EcsPrefixListArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EcsPrefixList(String name, EcsPrefixListArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EcsPrefixList(String name, EcsPrefixListArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/ecsPrefixList:EcsPrefixList", name, args == null ? EcsPrefixListArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EcsPrefixList(String name, Output<String> id, @Nullable EcsPrefixListState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/ecsPrefixList:EcsPrefixList", name, state, makeResourceOptions(options, id));
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
    public static EcsPrefixList get(String name, Output<String> id, @Nullable EcsPrefixListState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EcsPrefixList(name, id, state, options);
    }
}
