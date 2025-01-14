// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.edas;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.edas.InstanceClusterAttachmentArgs;
import com.pulumi.alicloud.edas.inputs.InstanceClusterAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import javax.annotation.Nullable;

/**
 * Provides an EDAS instance cluster attachment resource.
 * 
 * &gt; **NOTE:** Available in 1.82.0+
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
 * import com.pulumi.alicloud.edas.InstanceClusterAttachment;
 * import com.pulumi.alicloud.edas.InstanceClusterAttachmentArgs;
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
 *         var default_ = new InstanceClusterAttachment(&#34;default&#34;, InstanceClusterAttachmentArgs.builder()        
 *             .clusterId(var_.cluster_id())
 *             .instanceIds(var_.instance_ids())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="alicloud:edas/instanceClusterAttachment:InstanceClusterAttachment")
public class InstanceClusterAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the cluster that you want to create the application.
     * 
     */
    @Export(name="clusterId", type=String.class, parameters={})
    private Output<String> clusterId;

    /**
     * @return The ID of the cluster that you want to create the application.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * The cluster members map of the resource supplied above. The key is instance_id and the value is cluster_member_id.
     * 
     */
    @Export(name="clusterMemberIds", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> clusterMemberIds;

    /**
     * @return The cluster members map of the resource supplied above. The key is instance_id and the value is cluster_member_id.
     * 
     */
    public Output<Map<String,String>> clusterMemberIds() {
        return this.clusterMemberIds;
    }
    /**
     * The ecu map of the resource supplied above. The key is instance_id and the value is ecu_id.
     * 
     */
    @Export(name="ecuMap", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> ecuMap;

    /**
     * @return The ecu map of the resource supplied above. The key is instance_id and the value is ecu_id.
     * 
     */
    public Output<Map<String,String>> ecuMap() {
        return this.ecuMap;
    }
    /**
     * The ID of instance. Type: list.
     * 
     */
    @Export(name="instanceIds", type=List.class, parameters={String.class})
    private Output<List<String>> instanceIds;

    /**
     * @return The ID of instance. Type: list.
     * 
     */
    public Output<List<String>> instanceIds() {
        return this.instanceIds;
    }
    /**
     * The status map of the resource supplied above. The key is instance_id and the values are 1(running) 0(converting) -1(failed) and -2(offline).
     * 
     */
    @Export(name="statusMap", type=Map.class, parameters={String.class, Integer.class})
    private Output<Map<String,Integer>> statusMap;

    /**
     * @return The status map of the resource supplied above. The key is instance_id and the values are 1(running) 0(converting) -1(failed) and -2(offline).
     * 
     */
    public Output<Map<String,Integer>> statusMap() {
        return this.statusMap;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstanceClusterAttachment(String name) {
        this(name, InstanceClusterAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstanceClusterAttachment(String name, InstanceClusterAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstanceClusterAttachment(String name, InstanceClusterAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:edas/instanceClusterAttachment:InstanceClusterAttachment", name, args == null ? InstanceClusterAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InstanceClusterAttachment(String name, Output<String> id, @Nullable InstanceClusterAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:edas/instanceClusterAttachment:InstanceClusterAttachment", name, state, makeResourceOptions(options, id));
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
    public static InstanceClusterAttachment get(String name, Output<String> id, @Nullable InstanceClusterAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstanceClusterAttachment(name, id, state, options);
    }
}
