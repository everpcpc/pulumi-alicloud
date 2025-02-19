// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.threatdetection;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.threatdetection.HoneypotNodeArgs;
import com.pulumi.alicloud.threatdetection.inputs.HoneypotNodeState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Threat Detection Honeypot Node resource.
 * 
 * For information about Threat Detection Honeypot Node and how to use it, see [What is Honeypot Node](https://www.alibabacloud.com/help/en/security-center/latest/api-doc-sas-2018-12-03-api-doc-createhoneypotnode).
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
 * import com.pulumi.alicloud.threatdetection.HoneypotNode;
 * import com.pulumi.alicloud.threatdetection.HoneypotNodeArgs;
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
 *         var default_ = new HoneypotNode(&#34;default&#34;, HoneypotNodeArgs.builder()        
 *             .availableProbeNum(20)
 *             .nodeName(&#34;apispec_test&#34;)
 *             .securityGroupProbeIpLists(&#34;0.0.0.0/0&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Threat Detection Honeypot Node can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:threatdetection/honeypotNode:HoneypotNode example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:threatdetection/honeypotNode:HoneypotNode")
public class HoneypotNode extends com.pulumi.resources.CustomResource {
    /**
     * Whether to allow honeypot access to the external network. Value:-**true**: Allow-**false**: Disabled
     * 
     */
    @Export(name="allowHoneypotAccessInternet", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> allowHoneypotAccessInternet;

    /**
     * @return Whether to allow honeypot access to the external network. Value:-**true**: Allow-**false**: Disabled
     * 
     */
    public Output<Optional<Boolean>> allowHoneypotAccessInternet() {
        return Codegen.optional(this.allowHoneypotAccessInternet);
    }
    /**
     * Number of probes available.
     * 
     */
    @Export(name="availableProbeNum", type=Integer.class, parameters={})
    private Output<Integer> availableProbeNum;

    /**
     * @return Number of probes available.
     * 
     */
    public Output<Integer> availableProbeNum() {
        return this.availableProbeNum;
    }
    /**
     * The creation time of the resource
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return The creation time of the resource
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Management node name.
     * 
     */
    @Export(name="nodeName", type=String.class, parameters={})
    private Output<String> nodeName;

    /**
     * @return Management node name.
     * 
     */
    public Output<String> nodeName() {
        return this.nodeName;
    }
    /**
     * Release the collection of network segments.
     * 
     */
    @Export(name="securityGroupProbeIpLists", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> securityGroupProbeIpLists;

    /**
     * @return Release the collection of network segments.
     * 
     */
    public Output<Optional<List<String>>> securityGroupProbeIpLists() {
        return Codegen.optional(this.securityGroupProbeIpLists);
    }
    /**
     * The status of the resource
     * 
     */
    @Export(name="status", type=Integer.class, parameters={})
    private Output<Integer> status;

    /**
     * @return The status of the resource
     * 
     */
    public Output<Integer> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HoneypotNode(String name) {
        this(name, HoneypotNodeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HoneypotNode(String name, HoneypotNodeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HoneypotNode(String name, HoneypotNodeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/honeypotNode:HoneypotNode", name, args == null ? HoneypotNodeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HoneypotNode(String name, Output<String> id, @Nullable HoneypotNodeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:threatdetection/honeypotNode:HoneypotNode", name, state, makeResourceOptions(options, id));
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
    public static HoneypotNode get(String name, Output<String> id, @Nullable HoneypotNodeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HoneypotNode(name, id, state, options);
    }
}
