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
 * Provides an Alicloud EIP Association resource for associating Elastic IP to ECS Instance, SLB Instance or Nat Gateway.
 * 
 * &gt; **NOTE:** `alicloud.ecs.EipAssociation` is useful in scenarios where EIPs are either
 *  pre-existing or distributed to customers or users and therefore cannot be changed.
 * 
 * &gt; **NOTE:** From version 1.7.1, the resource support to associate EIP to SLB Instance or Nat Gateway.
 * 
 * &gt; **NOTE:** One EIP can only be associated with ECS or SLB instance which in the VPC.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.EcsFunctions;
 * import com.pulumi.alicloud.ecs.inputs.GetInstanceTypesArgs;
 * import com.pulumi.alicloud.ecs.inputs.GetImagesArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.ecs.Instance;
 * import com.pulumi.alicloud.ecs.InstanceArgs;
 * import com.pulumi.alicloud.ecs.EipAddress;
 * import com.pulumi.alicloud.ecs.EipAssociation;
 * import com.pulumi.alicloud.ecs.EipAssociationArgs;
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
 *         final var defaultZones = AlicloudFunctions.getZones();
 * 
 *         var vpc = new Network(&#34;vpc&#34;, NetworkArgs.builder()        
 *             .cidrBlock(&#34;10.1.0.0/21&#34;)
 *             .build());
 * 
 *         var vsw = new Switch(&#34;vsw&#34;, SwitchArgs.builder()        
 *             .vpcId(vpc.id())
 *             .cidrBlock(&#34;10.1.1.0/24&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(vpc)
 *                 .build());
 * 
 *         final var defaultInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .availabilityZone(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         final var defaultImages = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .nameRegex(&#34;^ubuntu_18.*64&#34;)
 *             .mostRecent(true)
 *             .owners(&#34;system&#34;)
 *             .build());
 * 
 *         var group = new SecurityGroup(&#34;group&#34;, SecurityGroupArgs.builder()        
 *             .description(&#34;New security group&#34;)
 *             .vpcId(vpc.id())
 *             .build());
 * 
 *         var ecsInstance = new Instance(&#34;ecsInstance&#34;, InstanceArgs.builder()        
 *             .imageId(defaultImages.applyValue(getImagesResult -&gt; getImagesResult.images()[0].id()))
 *             .instanceType(defaultInstanceTypes.applyValue(getInstanceTypesResult -&gt; getInstanceTypesResult.instanceTypes()[0].id()))
 *             .availabilityZone(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .securityGroups(group.id())
 *             .vswitchId(vsw.id())
 *             .instanceName(&#34;hello&#34;)
 *             .tags(Map.of(&#34;Name&#34;, &#34;TerraformTest-instance&#34;))
 *             .build());
 * 
 *         var eip = new EipAddress(&#34;eip&#34;);
 * 
 *         var eipAsso = new EipAssociation(&#34;eipAsso&#34;, EipAssociationArgs.builder()        
 *             .allocationId(eip.id())
 *             .instanceId(ecsInstance.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Module Support
 * 
 * You can use the existing eip module
 * to create several EIP instances and associate them with other resources one-click, like ECS instances, SLB, Nat Gateway and so on.
 * 
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
