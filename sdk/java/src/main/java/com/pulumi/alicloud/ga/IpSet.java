// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ga;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ga.IpSetArgs;
import com.pulumi.alicloud.ga.inputs.IpSetState;
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
 * Provides a Global Accelerator (GA) Ip Set resource.
 * 
 * For information about Global Accelerator (GA) Ip Set and how to use it, see [What is Ip Set](https://www.alibabacloud.com/help/en/doc-detail/153246.htm).
 * 
 * &gt; **NOTE:** Available in v1.113.0+.
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
 * import com.pulumi.alicloud.ga.Accelerator;
 * import com.pulumi.alicloud.ga.AcceleratorArgs;
 * import com.pulumi.alicloud.ga.BandwidthPackage;
 * import com.pulumi.alicloud.ga.BandwidthPackageArgs;
 * import com.pulumi.alicloud.ga.BandwidthPackageAttachment;
 * import com.pulumi.alicloud.ga.BandwidthPackageAttachmentArgs;
 * import com.pulumi.alicloud.ga.IpSet;
 * import com.pulumi.alicloud.ga.IpSetArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var exampleAccelerator = new Accelerator(&#34;exampleAccelerator&#34;, AcceleratorArgs.builder()        
 *             .duration(1)
 *             .autoUseCoupon(true)
 *             .spec(&#34;1&#34;)
 *             .build());
 * 
 *         var exampleBandwidthPackage = new BandwidthPackage(&#34;exampleBandwidthPackage&#34;, BandwidthPackageArgs.builder()        
 *             .bandwidth(20)
 *             .type(&#34;Basic&#34;)
 *             .bandwidthType(&#34;Basic&#34;)
 *             .duration(1)
 *             .autoPay(true)
 *             .ratio(30)
 *             .build());
 * 
 *         var exampleBandwidthPackageAttachment = new BandwidthPackageAttachment(&#34;exampleBandwidthPackageAttachment&#34;, BandwidthPackageAttachmentArgs.builder()        
 *             .acceleratorId(exampleAccelerator.id())
 *             .bandwidthPackageId(exampleBandwidthPackage.id())
 *             .build());
 * 
 *         var exampleIpSet = new IpSet(&#34;exampleIpSet&#34;, IpSetArgs.builder()        
 *             .accelerateRegionId(&#34;cn-hangzhou&#34;)
 *             .bandwidth(&#34;5&#34;)
 *             .acceleratorId(exampleAccelerator.id())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(exampleBandwidthPackageAttachment)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Ga Ip Set can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ga/ipSet:IpSet example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ga/ipSet:IpSet")
public class IpSet extends com.pulumi.resources.CustomResource {
    /**
     * The ID of an acceleration region.
     * 
     */
    @Export(name="accelerateRegionId", type=String.class, parameters={})
    private Output<String> accelerateRegionId;

    /**
     * @return The ID of an acceleration region.
     * 
     */
    public Output<String> accelerateRegionId() {
        return this.accelerateRegionId;
    }
    /**
     * The ID of the Global Accelerator (GA) instance.
     * 
     */
    @Export(name="acceleratorId", type=String.class, parameters={})
    private Output<String> acceleratorId;

    /**
     * @return The ID of the Global Accelerator (GA) instance.
     * 
     */
    public Output<String> acceleratorId() {
        return this.acceleratorId;
    }
    /**
     * The bandwidth allocated to the acceleration region.
     * 
     * &gt; **NOTE:** The minimum bandwidth of each accelerated region is 2Mbps. The total bandwidth of the acceleration region should be less than or equal to the bandwidth of the basic bandwidth package you purchased.
     * 
     */
    @Export(name="bandwidth", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> bandwidth;

    /**
     * @return The bandwidth allocated to the acceleration region.
     * 
     * &gt; **NOTE:** The minimum bandwidth of each accelerated region is 2Mbps. The total bandwidth of the acceleration region should be less than or equal to the bandwidth of the basic bandwidth package you purchased.
     * 
     */
    public Output<Optional<Integer>> bandwidth() {
        return Codegen.optional(this.bandwidth);
    }
    /**
     * The list of accelerated IP addresses in the acceleration region.
     * 
     */
    @Export(name="ipAddressLists", type=List.class, parameters={String.class})
    private Output<List<String>> ipAddressLists;

    /**
     * @return The list of accelerated IP addresses in the acceleration region.
     * 
     */
    public Output<List<String>> ipAddressLists() {
        return this.ipAddressLists;
    }
    /**
     * The IP protocol used by the GA instance. Valid values: `IPv4`, `IPv6`. Default value is `IPv4`.
     * 
     */
    @Export(name="ipVersion", type=String.class, parameters={})
    private Output</* @Nullable */ String> ipVersion;

    /**
     * @return The IP protocol used by the GA instance. Valid values: `IPv4`, `IPv6`. Default value is `IPv4`.
     * 
     */
    public Output<Optional<String>> ipVersion() {
        return Codegen.optional(this.ipVersion);
    }
    /**
     * The status of the acceleration region.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the acceleration region.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IpSet(String name) {
        this(name, IpSetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IpSet(String name, IpSetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IpSet(String name, IpSetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/ipSet:IpSet", name, args == null ? IpSetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IpSet(String name, Output<String> id, @Nullable IpSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ga/ipSet:IpSet", name, state, makeResourceOptions(options, id));
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
    public static IpSet get(String name, Output<String> id, @Nullable IpSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IpSet(name, id, state, options);
    }
}
