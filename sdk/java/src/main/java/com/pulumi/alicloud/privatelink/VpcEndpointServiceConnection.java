// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.privatelink;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.privatelink.VpcEndpointServiceConnectionArgs;
import com.pulumi.alicloud.privatelink.inputs.VpcEndpointServiceConnectionState;
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
 * Provides a Private Link Vpc Endpoint Connection resource.
 * 
 * For information about Private Link Vpc Endpoint Connection and how to use it, see [What is Vpc Endpoint Connection](https://help.aliyun.com/document_detail/183551.html).
 * 
 * &gt; **NOTE:** Available in v1.110.0+.
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
 * import com.pulumi.alicloud.privatelink.VpcEndpointServiceConnection;
 * import com.pulumi.alicloud.privatelink.VpcEndpointServiceConnectionArgs;
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
 *         var example = new VpcEndpointServiceConnection(&#34;example&#34;, VpcEndpointServiceConnectionArgs.builder()        
 *             .bandwidth(&#34;1024&#34;)
 *             .endpointId(&#34;example_value&#34;)
 *             .serviceId(&#34;example_value&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Private Link Vpc Endpoint Connection can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection example &lt;service_id&gt;:&lt;endpoint_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection")
public class VpcEndpointServiceConnection extends com.pulumi.resources.CustomResource {
    /**
     * The Bandwidth.
     * 
     */
    @Export(name="bandwidth", type=Integer.class, parameters={})
    private Output<Integer> bandwidth;

    /**
     * @return The Bandwidth.
     * 
     */
    public Output<Integer> bandwidth() {
        return this.bandwidth;
    }
    /**
     * The dry run.
     * 
     */
    @Export(name="dryRun", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return The dry run.
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * The ID of the Vpc Endpoint.
     * 
     */
    @Export(name="endpointId", type=String.class, parameters={})
    private Output<String> endpointId;

    /**
     * @return The ID of the Vpc Endpoint.
     * 
     */
    public Output<String> endpointId() {
        return this.endpointId;
    }
    /**
     * The ID of the Vpc Endpoint Service.
     * 
     */
    @Export(name="serviceId", type=String.class, parameters={})
    private Output<String> serviceId;

    /**
     * @return The ID of the Vpc Endpoint Service.
     * 
     */
    public Output<String> serviceId() {
        return this.serviceId;
    }
    /**
     * The status of Vpc Endpoint Connection.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of Vpc Endpoint Connection.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcEndpointServiceConnection(String name) {
        this(name, VpcEndpointServiceConnectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcEndpointServiceConnection(String name, VpcEndpointServiceConnectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcEndpointServiceConnection(String name, VpcEndpointServiceConnectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection", name, args == null ? VpcEndpointServiceConnectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpcEndpointServiceConnection(String name, Output<String> id, @Nullable VpcEndpointServiceConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:privatelink/vpcEndpointServiceConnection:VpcEndpointServiceConnection", name, state, makeResourceOptions(options, id));
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
    public static VpcEndpointServiceConnection get(String name, Output<String> id, @Nullable VpcEndpointServiceConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcEndpointServiceConnection(name, id, state, options);
    }
}
