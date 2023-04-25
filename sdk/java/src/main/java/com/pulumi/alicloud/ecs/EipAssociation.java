// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.EipAssociationArgs;
import com.pulumi.alicloud.ecs.inputs.EipAssociationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * Elastic IP address association can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ecs/eipAssociation:EipAssociation example eip-abc12345678:i-abc12355
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/eipAssociation:EipAssociation")
public class EipAssociation extends com.pulumi.resources.CustomResource {
    /**
     * The allocation EIP ID.
     * 
     */
    @Export(name="allocationId", type=String.class, parameters={})
    private Output<String> allocationId;

    /**
     * @return The allocation EIP ID.
     * 
     */
    public Output<String> allocationId() {
        return this.allocationId;
    }
    /**
     * When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default to `false`.
     * 
     */
    @Export(name="force", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> force;

    /**
     * @return When EIP is bound to a NAT gateway, and the NAT gateway adds a DNAT or SNAT entry, set it for `true` can unassociation any way. Default to `false`.
     * 
     */
    public Output<Optional<Boolean>> force() {
        return Codegen.optional(this.force);
    }
    /**
     * The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
     * 
     */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return The ID of the ECS or SLB instance or Nat Gateway or NetworkInterface or HaVip.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The type of cloud product that the eip instance to bind. Valid values: `EcsInstance`, `SlbInstance`, `Nat`, `NetworkInterface`, `HaVip` and `IpAddress`.
     * 
     */
    @Export(name="instanceType", type=String.class, parameters={})
    private Output<String> instanceType;

    /**
     * @return The type of cloud product that the eip instance to bind. Valid values: `EcsInstance`, `SlbInstance`, `Nat`, `NetworkInterface`, `HaVip` and `IpAddress`.
     * 
     */
    public Output<String> instanceType() {
        return this.instanceType;
    }
    /**
     * The private IP address in the network segment of the vswitch which has been assigned.
     * 
     */
    @Export(name="privateIpAddress", type=String.class, parameters={})
    private Output<String> privateIpAddress;

    /**
     * @return The private IP address in the network segment of the vswitch which has been assigned.
     * 
     */
    public Output<String> privateIpAddress() {
        return this.privateIpAddress;
    }
    /**
     * The ID of the VPC that has IPv4 gateways enabled and that is deployed in the same region as the EIP. When you associate an EIP with an IP address, the system can enable the IP address to access the Internet based on VPC route configurations. **Note:** This parameter is required if `instance_type` is set to IpAddress. In this case, the EIP is associated with an IP address.
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output</* @Nullable */ String> vpcId;

    /**
     * @return The ID of the VPC that has IPv4 gateways enabled and that is deployed in the same region as the EIP. When you associate an EIP with an IP address, the system can enable the IP address to access the Internet based on VPC route configurations. **Note:** This parameter is required if `instance_type` is set to IpAddress. In this case, the EIP is associated with an IP address.
     * 
     */
    public Output<Optional<String>> vpcId() {
        return Codegen.optional(this.vpcId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EipAssociation(String name) {
        this(name, EipAssociationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EipAssociation(String name, EipAssociationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EipAssociation(String name, EipAssociationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/eipAssociation:EipAssociation", name, args == null ? EipAssociationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EipAssociation(String name, Output<String> id, @Nullable EipAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/eipAssociation:EipAssociation", name, state, makeResourceOptions(options, id));
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
    public static EipAssociation get(String name, Output<String> id, @Nullable EipAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EipAssociation(name, id, state, options);
    }
}
