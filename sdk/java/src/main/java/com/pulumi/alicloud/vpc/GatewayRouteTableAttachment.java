// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.GatewayRouteTableAttachmentArgs;
import com.pulumi.alicloud.vpc.inputs.GatewayRouteTableAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a VPC Gateway Route Table Attachment resource.
 * 
 * For information about VPC Gateway Route Table Attachment and how to use it, see [What is Gateway Route Table Attachment](https://www.alibabacloud.com/help/doc-detail/174112.htm).
 * 
 * &gt; **NOTE:** Available in v1.194.0+.
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
 * import com.pulumi.alicloud.vpc.GatewayRouteTableAttachment;
 * import com.pulumi.alicloud.vpc.GatewayRouteTableAttachmentArgs;
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
 *         var example = new GatewayRouteTableAttachment(&#34;example&#34;, GatewayRouteTableAttachmentArgs.builder()        
 *             .ipv4GatewayId(&#34;example_value&#34;)
 *             .routeTableId(&#34;example_value&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * VPC Gateway Route Table Attachment can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:vpc/gatewayRouteTableAttachment:GatewayRouteTableAttachment example &lt;route_table_id&gt;:&lt;ipv4_gateway_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/gatewayRouteTableAttachment:GatewayRouteTableAttachment")
public class GatewayRouteTableAttachment extends com.pulumi.resources.CustomResource {
    /**
     * Specifies whether to only precheck this request. Default value: `false`.
     * 
     */
    @Export(name="dryRun", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return Specifies whether to only precheck this request. Default value: `false`.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * The ID of the IPv4 Gateway instance.
     * 
     */
    @Export(name="ipv4GatewayId", type=String.class, parameters={})
    private Output<String> ipv4GatewayId;

    /**
     * @return The ID of the IPv4 Gateway instance.
     * 
     */
    public Output<String> ipv4GatewayId() {
        return this.ipv4GatewayId;
    }
    /**
     * The ID of the Gateway route table to be bound.
     * 
     */
    @Export(name="routeTableId", type=String.class, parameters={})
    private Output<String> routeTableId;

    /**
     * @return The ID of the Gateway route table to be bound.
     * 
     */
    public Output<String> routeTableId() {
        return this.routeTableId;
    }
    /**
     * The status of the IPv4 Gateway instance. Value:
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the IPv4 Gateway instance. Value:
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GatewayRouteTableAttachment(String name) {
        this(name, GatewayRouteTableAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GatewayRouteTableAttachment(String name, GatewayRouteTableAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GatewayRouteTableAttachment(String name, GatewayRouteTableAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/gatewayRouteTableAttachment:GatewayRouteTableAttachment", name, args == null ? GatewayRouteTableAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GatewayRouteTableAttachment(String name, Output<String> id, @Nullable GatewayRouteTableAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/gatewayRouteTableAttachment:GatewayRouteTableAttachment", name, state, makeResourceOptions(options, id));
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
    public static GatewayRouteTableAttachment get(String name, Output<String> id, @Nullable GatewayRouteTableAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GatewayRouteTableAttachment(name, id, state, options);
    }
}
