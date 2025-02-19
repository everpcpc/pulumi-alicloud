// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.Ipv6GatewayArgs;
import com.pulumi.alicloud.vpc.inputs.Ipv6GatewayState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a VPC Ipv6 Gateway resource.
 * 
 * For information about VPC Ipv6 Gateway and how to use it, see [What is Ipv6 Gateway](https://www.alibabacloud.com/help/doc-detail/102214.htm).
 * 
 * &gt; **NOTE:** Available in v1.142.0+.
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
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Ipv6Gateway;
 * import com.pulumi.alicloud.vpc.Ipv6GatewayArgs;
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
 *         var default_ = new Network(&#34;default&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;example_value&#34;)
 *             .enableIpv6(&#34;true&#34;)
 *             .build());
 * 
 *         var example = new Ipv6Gateway(&#34;example&#34;, Ipv6GatewayArgs.builder()        
 *             .ipv6GatewayName(&#34;example_value&#34;)
 *             .vpcId(default_.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * VPC Ipv6 Gateway can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:vpc/ipv6Gateway:Ipv6Gateway example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/ipv6Gateway:Ipv6Gateway")
public class Ipv6Gateway extends com.pulumi.resources.CustomResource {
    /**
     * The description of the IPv6 gateway. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the IPv6 gateway. The description must be `2` to `256` characters in length. It cannot start with `http://` or `https://`.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The name of the IPv6 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
     * 
     */
    @Export(name="ipv6GatewayName", type=String.class, parameters={})
    private Output</* @Nullable */ String> ipv6GatewayName;

    /**
     * @return The name of the IPv6 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
     * 
     */
    public Output<Optional<String>> ipv6GatewayName() {
        return Codegen.optional(this.ipv6GatewayName);
    }
    /**
     * The edition of the IPv6 gateway. Valid values: `Large`, `Medium` and `Small`. `Small` (default): Free Edition. `Medium`: Enterprise Edition . `Large`: Enhanced Enterprise Edition. The throughput capacity of an IPv6 gateway varies based on the edition. For more information, see [Editions of IPv6 gateways](https://www.alibabacloud.com/help/doc-detail/98926.htm).
     * 
     */
    @Export(name="spec", type=String.class, parameters={})
    private Output<String> spec;

    /**
     * @return The edition of the IPv6 gateway. Valid values: `Large`, `Medium` and `Small`. `Small` (default): Free Edition. `Medium`: Enterprise Edition . `Large`: Enhanced Enterprise Edition. The throughput capacity of an IPv6 gateway varies based on the edition. For more information, see [Editions of IPv6 gateways](https://www.alibabacloud.com/help/doc-detail/98926.htm).
     * 
     */
    public Output<String> spec() {
        return this.spec;
    }
    /**
     * The status of the resource. Valid values: `Available`, `Pending` and `Deleting`.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the resource. Valid values: `Available`, `Pending` and `Deleting`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The ID of the virtual private cloud (VPC) for which you want to create the IPv6 gateway.
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output<String> vpcId;

    /**
     * @return The ID of the virtual private cloud (VPC) for which you want to create the IPv6 gateway.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Ipv6Gateway(String name) {
        this(name, Ipv6GatewayArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Ipv6Gateway(String name, Ipv6GatewayArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Ipv6Gateway(String name, Ipv6GatewayArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/ipv6Gateway:Ipv6Gateway", name, args == null ? Ipv6GatewayArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Ipv6Gateway(String name, Output<String> id, @Nullable Ipv6GatewayState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/ipv6Gateway:Ipv6Gateway", name, state, makeResourceOptions(options, id));
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
    public static Ipv6Gateway get(String name, Output<String> id, @Nullable Ipv6GatewayState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Ipv6Gateway(name, id, state, options);
    }
}
