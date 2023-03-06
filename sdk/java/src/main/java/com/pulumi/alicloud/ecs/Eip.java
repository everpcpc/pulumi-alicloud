// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.EipArgs;
import com.pulumi.alicloud.ecs.inputs.EipState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * Elastic IP address can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ecs/eip:Eip example eip-abc12345678
 * ```
 * 
 * @deprecated
 * This resource has been deprecated in favour of the EipAddress resource
 * 
 */
@Deprecated /* This resource has been deprecated in favour of the EipAddress resource */
@ResourceType(type="alicloud:ecs/eip:Eip")
public class Eip extends com.pulumi.resources.CustomResource {
    @Export(name="activityId", type=String.class, parameters={})
    private Output</* @Nullable */ String> activityId;

    public Output<Optional<String>> activityId() {
        return Codegen.optional(this.activityId);
    }
    /**
     * The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://.
     * 
     */
    @Export(name="addressName", type=String.class, parameters={})
    private Output<String> addressName;

    /**
     * @return The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin or end with a hyphen, and must not begin with http:// or https://.
     * 
     */
    public Output<String> addressName() {
        return this.addressName;
    }
    @Export(name="autoPay", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> autoPay;

    public Output<Optional<Boolean>> autoPay() {
        return Codegen.optional(this.autoPay);
    }
    /**
     * Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
     * 
     */
    @Export(name="bandwidth", type=String.class, parameters={})
    private Output<String> bandwidth;

    /**
     * @return Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 5 Mbps.
     * 
     */
    public Output<String> bandwidth() {
        return this.bandwidth;
    }
    /**
     * Whether enable the deletion protection or not. Default value: `false`.
     * - true: Enable deletion protection.
     * - false: Disable deletion protection.
     * 
     */
    @Export(name="deletionProtection", type=Boolean.class, parameters={})
    private Output<Boolean> deletionProtection;

    /**
     * @return Whether enable the deletion protection or not. Default value: `false`.
     * - true: Enable deletion protection.
     * - false: Disable deletion protection.
     * 
     */
    public Output<Boolean> deletionProtection() {
        return this.deletionProtection;
    }
    /**
     * Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the EIP instance, This description can have a string of 2 to 256 characters, It cannot begin with http:// or https://. Default value is null.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    @Export(name="highDefinitionMonitorLogStatus", type=String.class, parameters={})
    private Output<String> highDefinitionMonitorLogStatus;

    public Output<String> highDefinitionMonitorLogStatus() {
        return this.highDefinitionMonitorLogStatus;
    }
    /**
     * (It has been deprecated from version 1.126.0 and using new attribute `payment_type` instead) Elastic IP instance charge type. Valid values are &#34;PrePaid&#34; and &#34;PostPaid&#34;. Default to &#34;PostPaid&#34;.
     * 
     * @deprecated
     * Field &#39;instance_charge_type&#39; has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute &#39;payment_type&#39; instead.
     * 
     */
    @Deprecated /* Field 'instance_charge_type' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'payment_type' instead. */
    @Export(name="instanceChargeType", type=String.class, parameters={})
    private Output<String> instanceChargeType;

    /**
     * @return (It has been deprecated from version 1.126.0 and using new attribute `payment_type` instead) Elastic IP instance charge type. Valid values are &#34;PrePaid&#34; and &#34;PostPaid&#34;. Default to &#34;PostPaid&#34;.
     * 
     */
    public Output<String> instanceChargeType() {
        return this.instanceChargeType;
    }
    /**
     * Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** From version `1.7.1` to `1.125.0`, it defaults to `PayByTraffic`. It is only &#34;PayByBandwidth&#34; when `instance_charge_type` is PrePaid.
     * 
     */
    @Export(name="internetChargeType", type=String.class, parameters={})
    private Output<String> internetChargeType;

    /**
     * @return Internet charge type of the EIP, Valid values are `PayByBandwidth`, `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** From version `1.7.1` to `1.125.0`, it defaults to `PayByTraffic`. It is only &#34;PayByBandwidth&#34; when `instance_charge_type` is PrePaid.
     * 
     */
    public Output<String> internetChargeType() {
        return this.internetChargeType;
    }
    /**
     * The elastic ip address
     * 
     */
    @Export(name="ipAddress", type=String.class, parameters={})
    private Output<String> ipAddress;

    /**
     * @return The elastic ip address
     * 
     */
    public Output<String> ipAddress() {
        return this.ipAddress;
    }
    /**
     * The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
     * 
     */
    @Export(name="isp", type=String.class, parameters={})
    private Output<String> isp;

    /**
     * @return The line type of the Elastic IP instance. Default to `BGP`. Other type of the isp need to open a whitelist.
     * 
     */
    public Output<String> isp() {
        return this.isp;
    }
    @Export(name="logProject", type=String.class, parameters={})
    private Output</* @Nullable */ String> logProject;

    public Output<Optional<String>> logProject() {
        return Codegen.optional(this.logProject);
    }
    @Export(name="logStore", type=String.class, parameters={})
    private Output</* @Nullable */ String> logStore;

    public Output<Optional<String>> logStore() {
        return Codegen.optional(this.logStore);
    }
    /**
     * It has been deprecated from version 1.126.0 and using new attribute `address_name` instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute &#39;address_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.126.0 and it will be remove in the future version. Please use the new attribute 'address_name' instead. */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return It has been deprecated from version 1.126.0 and using new attribute `address_name` instead.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="netmode", type=String.class, parameters={})
    private Output</* @Nullable */ String> netmode;

    public Output<Optional<String>> netmode() {
        return Codegen.optional(this.netmode);
    }
    /**
     * The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
     * 
     */
    @Export(name="paymentType", type=String.class, parameters={})
    private Output<String> paymentType;

    /**
     * @return The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    @Export(name="period", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> period;

    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    @Export(name="publicIpAddressPoolId", type=String.class, parameters={})
    private Output</* @Nullable */ String> publicIpAddressPoolId;

    public Output<Optional<String>> publicIpAddressPoolId() {
        return Codegen.optional(this.publicIpAddressPoolId);
    }
    /**
     * The Id of resource group which the eip belongs.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output<String> resourceGroupId;

    /**
     * @return The Id of resource group which the eip belongs.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    @Export(name="securityProtectionTypes", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> securityProtectionTypes;

    public Output<Optional<List<String>>> securityProtectionTypes() {
        return Codegen.optional(this.securityProtectionTypes);
    }
    /**
     * The EIP current status.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The EIP current status.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Eip(String name) {
        this(name, EipArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Eip(String name, @Nullable EipArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Eip(String name, @Nullable EipArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/eip:Eip", name, args == null ? EipArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Eip(String name, Output<String> id, @Nullable EipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/eip:Eip", name, state, makeResourceOptions(options, id));
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
    public static Eip get(String name, Output<String> id, @Nullable EipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Eip(name, id, state, options);
    }
}
