// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.DhcpOptionsSetAttachmentArgs;
import com.pulumi.alicloud.vpc.inputs.DhcpOptionsSetAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a VPC Dhcp Options Set Attachment resource.
 * 
 * For information about VPC Dhcp Options Set and how to use it, see [What is Dhcp Options Set](https://www.alibabacloud.com/help/doc-detail/174112.htm).
 * 
 * &gt; **NOTE:** Available in v1.153.0+.
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
 * import com.pulumi.alicloud.vpc.DhcpOptionsSet;
 * import com.pulumi.alicloud.vpc.DhcpOptionsSetArgs;
 * import com.pulumi.alicloud.vpc.DhcpOptionsSetAttachment;
 * import com.pulumi.alicloud.vpc.DhcpOptionsSetAttachmentArgs;
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
 *         var exampleNetwork = new Network(&#34;exampleNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;test&#34;)
 *             .cidrBlock(&#34;172.16.0.0/12&#34;)
 *             .build());
 * 
 *         var exampleDhcpOptionsSet = new DhcpOptionsSet(&#34;exampleDhcpOptionsSet&#34;, DhcpOptionsSetArgs.builder()        
 *             .dhcpOptionsSetName(&#34;example_value&#34;)
 *             .dhcpOptionsSetDescription(&#34;example_value&#34;)
 *             .domainName(&#34;example.com&#34;)
 *             .domainNameServers(&#34;100.100.2.136&#34;)
 *             .build());
 * 
 *         var exampleDhcpOptionsSetAttachment = new DhcpOptionsSetAttachment(&#34;exampleDhcpOptionsSetAttachment&#34;, DhcpOptionsSetAttachmentArgs.builder()        
 *             .vpcId(exampleNetwork.id())
 *             .dhcpOptionsSetId(exampleDhcpOptionsSet.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * VPC Dhcp Options Set Attachment can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:vpc/dhcpOptionsSetAttachment:DhcpOptionsSetAttachment example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/dhcpOptionsSetAttachment:DhcpOptionsSetAttachment")
public class DhcpOptionsSetAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the DHCP options set.
     * 
     */
    @Export(name="dhcpOptionsSetId", type=String.class, parameters={})
    private Output<String> dhcpOptionsSetId;

    /**
     * @return The ID of the DHCP options set.
     * 
     */
    public Output<String> dhcpOptionsSetId() {
        return this.dhcpOptionsSetId;
    }
    /**
     * Specifies whether to precheck this request only. Default values: `false`. Valid values:
     * 
     */
    @Export(name="dryRun", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return Specifies whether to precheck this request only. Default values: `false`. Valid values:
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * The status of the VPC network that is associated with the DHCP options set.  Valid values: `InUse` or `Pending`.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the VPC network that is associated with the DHCP options set.  Valid values: `InUse` or `Pending`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The ID of the VPC network that is to be associated with the DHCP options set..
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output<String> vpcId;

    /**
     * @return The ID of the VPC network that is to be associated with the DHCP options set..
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DhcpOptionsSetAttachment(String name) {
        this(name, DhcpOptionsSetAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DhcpOptionsSetAttachment(String name, DhcpOptionsSetAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DhcpOptionsSetAttachment(String name, DhcpOptionsSetAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/dhcpOptionsSetAttachment:DhcpOptionsSetAttachment", name, args == null ? DhcpOptionsSetAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DhcpOptionsSetAttachment(String name, Output<String> id, @Nullable DhcpOptionsSetAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/dhcpOptionsSetAttachment:DhcpOptionsSetAttachment", name, state, makeResourceOptions(options, id));
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
    public static DhcpOptionsSetAttachment get(String name, Output<String> id, @Nullable DhcpOptionsSetAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DhcpOptionsSetAttachment(name, id, state, options);
    }
}
