// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.dns;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.dns.AddressPoolArgs;
import com.pulumi.alicloud.dns.inputs.AddressPoolState;
import com.pulumi.alicloud.dns.outputs.AddressPoolAddress;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a Alidns Address Pool resource.
 * 
 * For information about Alidns Address Pool and how to use it, see [What is Address Pool](https://www.alibabacloud.com/help/doc-detail/189621.html).
 * 
 * &gt; **NOTE:** Available in v1.152.0+.
 * 
 * ## Import
 * 
 * Alidns Address Pool can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:dns/addressPool:AddressPool example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:dns/addressPool:AddressPool")
public class AddressPool extends com.pulumi.resources.CustomResource {
    /**
     * The name of the address pool.
     * 
     */
    @Export(name="addressPoolName", type=String.class, parameters={})
    private Output<String> addressPoolName;

    /**
     * @return The name of the address pool.
     * 
     */
    public Output<String> addressPoolName() {
        return this.addressPoolName;
    }
    /**
     * The address lists of the Address Pool. See the following `Block address`.
     * 
     */
    @Export(name="addresses", type=List.class, parameters={AddressPoolAddress.class})
    private Output<List<AddressPoolAddress>> addresses;

    /**
     * @return The address lists of the Address Pool. See the following `Block address`.
     * 
     */
    public Output<List<AddressPoolAddress>> addresses() {
        return this.addresses;
    }
    /**
     * The ID of the instance.
     * 
     */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return The ID of the instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The load balancing policy of the address pool. Valid values:`ALL_RR` or `RATIO`. `ALL_RR`: returns all addresses. `RATIO`: returns addresses by weight.
     * 
     */
    @Export(name="lbaStrategy", type=String.class, parameters={})
    private Output<String> lbaStrategy;

    /**
     * @return The load balancing policy of the address pool. Valid values:`ALL_RR` or `RATIO`. `ALL_RR`: returns all addresses. `RATIO`: returns addresses by weight.
     * 
     */
    public Output<String> lbaStrategy() {
        return this.lbaStrategy;
    }
    /**
     * The type of the address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output<String> type;

    /**
     * @return The type of the address pool. Valid values: `IPV4`, `IPV6`, `DOMAIN`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AddressPool(String name) {
        this(name, AddressPoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AddressPool(String name, AddressPoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AddressPool(String name, AddressPoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/addressPool:AddressPool", name, args == null ? AddressPoolArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AddressPool(String name, Output<String> id, @Nullable AddressPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:dns/addressPool:AddressPool", name, state, makeResourceOptions(options, id));
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
    public static AddressPool get(String name, Output<String> id, @Nullable AddressPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AddressPool(name, id, state, options);
    }
}