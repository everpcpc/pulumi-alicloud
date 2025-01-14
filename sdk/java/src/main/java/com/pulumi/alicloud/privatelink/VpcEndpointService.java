// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.privatelink;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.privatelink.VpcEndpointServiceArgs;
import com.pulumi.alicloud.privatelink.inputs.VpcEndpointServiceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Private Link Vpc Endpoint Service resource.
 * 
 * For information about Private Link Vpc Endpoint Service and how to use it, see [What is Vpc Endpoint Service](https://help.aliyun.com/document_detail/183540.html).
 * 
 * &gt; **NOTE:** Available in v1.109.0+.
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
 * import com.pulumi.alicloud.privatelink.VpcEndpointService;
 * import com.pulumi.alicloud.privatelink.VpcEndpointServiceArgs;
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
 *         var example = new VpcEndpointService(&#34;example&#34;, VpcEndpointServiceArgs.builder()        
 *             .autoAcceptConnection(false)
 *             .connectBandwidth(103)
 *             .serviceDescription(&#34;tftest&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Private Link Vpc Endpoint Service can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:privatelink/vpcEndpointService:VpcEndpointService example &lt;service_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:privatelink/vpcEndpointService:VpcEndpointService")
public class VpcEndpointService extends com.pulumi.resources.CustomResource {
    /**
     * Whether to automatically accept terminal node connections.
     * 
     */
    @Export(name="autoAcceptConnection", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> autoAcceptConnection;

    /**
     * @return Whether to automatically accept terminal node connections.
     * 
     */
    public Output<Optional<Boolean>> autoAcceptConnection() {
        return Codegen.optional(this.autoAcceptConnection);
    }
    /**
     * The connection bandwidth.
     * 
     */
    @Export(name="connectBandwidth", type=Integer.class, parameters={})
    private Output<Integer> connectBandwidth;

    /**
     * @return The connection bandwidth.
     * 
     */
    public Output<Integer> connectBandwidth() {
        return this.connectBandwidth;
    }
    /**
     * Whether to pre-check this request only. Default to: `false`
     * 
     */
    @Export(name="dryRun", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return Whether to pre-check this request only. Default to: `false`
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * The payer type. Valid Value: `EndpointService`, `Endpoint`. Default to: `Endpoint`.
     * 
     */
    @Export(name="payer", type=String.class, parameters={})
    private Output</* @Nullable */ String> payer;

    /**
     * @return The payer type. Valid Value: `EndpointService`, `Endpoint`. Default to: `Endpoint`.
     * 
     */
    public Output<Optional<String>> payer() {
        return Codegen.optional(this.payer);
    }
    /**
     * The business status of Vpc Endpoint Service.
     * 
     */
    @Export(name="serviceBusinessStatus", type=String.class, parameters={})
    private Output<String> serviceBusinessStatus;

    /**
     * @return The business status of Vpc Endpoint Service.
     * 
     */
    public Output<String> serviceBusinessStatus() {
        return this.serviceBusinessStatus;
    }
    /**
     * The description of the terminal node service.
     * 
     * &gt; **NOTE:** The `resources` only support load balancing instance with private network type and PrivateLink function.
     * 
     */
    @Export(name="serviceDescription", type=String.class, parameters={})
    private Output</* @Nullable */ String> serviceDescription;

    /**
     * @return The description of the terminal node service.
     * 
     * &gt; **NOTE:** The `resources` only support load balancing instance with private network type and PrivateLink function.
     * 
     */
    public Output<Optional<String>> serviceDescription() {
        return Codegen.optional(this.serviceDescription);
    }
    /**
     * Service Domain.
     * 
     */
    @Export(name="serviceDomain", type=String.class, parameters={})
    private Output<String> serviceDomain;

    /**
     * @return Service Domain.
     * 
     */
    public Output<String> serviceDomain() {
        return this.serviceDomain;
    }
    /**
     * The status of Vpc Endpoint Service.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of Vpc Endpoint Service.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcEndpointService(String name) {
        this(name, VpcEndpointServiceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcEndpointService(String name, @Nullable VpcEndpointServiceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcEndpointService(String name, @Nullable VpcEndpointServiceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:privatelink/vpcEndpointService:VpcEndpointService", name, args == null ? VpcEndpointServiceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpcEndpointService(String name, Output<String> id, @Nullable VpcEndpointServiceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:privatelink/vpcEndpointService:VpcEndpointService", name, state, makeResourceOptions(options, id));
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
    public static VpcEndpointService get(String name, Output<String> id, @Nullable VpcEndpointServiceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcEndpointService(name, id, state, options);
    }
}
