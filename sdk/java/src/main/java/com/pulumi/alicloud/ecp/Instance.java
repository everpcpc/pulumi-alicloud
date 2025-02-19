// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecp;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecp.InstanceArgs;
import com.pulumi.alicloud.ecp.inputs.InstanceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Elastic Cloud Phone (ECP) Instance resource.
 * 
 * For information about Elastic Cloud Phone (ECP) Instance and how to use it,
 * see [What is Instance](https://help.aliyun.com/document_detail/258178.html/).
 * 
 * &gt; **NOTE:** Available in v1.158.0+.
 * 
 * ## Import
 * 
 * Elastic Cloud Phone (ECP) Instance can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:ecp/instance:Instance example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ecp/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * The auto pay.
     * 
     */
    @Export(name="autoPay", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> autoPay;

    /**
     * @return The auto pay.
     * 
     */
    public Output<Optional<Boolean>> autoPay() {
        return Codegen.optional(this.autoPay);
    }
    /**
     * The auto renew.
     * 
     */
    @Export(name="autoRenew", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> autoRenew;

    /**
     * @return The auto renew.
     * 
     */
    public Output<Optional<Boolean>> autoRenew() {
        return Codegen.optional(this.autoRenew);
    }
    /**
     * Description of the instance. 2 to 256 English or Chinese characters in length and cannot
     * start with `http://` and `https`.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the instance. 2 to 256 English or Chinese characters in length and cannot
     * start with `http://` and `https`.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The eip bandwidth.
     * 
     */
    @Export(name="eipBandwidth", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> eipBandwidth;

    /**
     * @return The eip bandwidth.
     * 
     */
    public Output<Optional<Integer>> eipBandwidth() {
        return Codegen.optional(this.eipBandwidth);
    }
    /**
     * The force.
     * 
     */
    @Export(name="force", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> force;

    /**
     * @return The force.
     * 
     */
    public Output<Optional<Boolean>> force() {
        return Codegen.optional(this.force);
    }
    /**
     * The ID Of The Image.
     * 
     */
    @Export(name="imageId", type=String.class, parameters={})
    private Output<String> imageId;

    /**
     * @return The ID Of The Image.
     * 
     */
    public Output<String> imageId() {
        return this.imageId;
    }
    /**
     * The name of the instance. It must be 2 to 128 characters in length and must start with an
     * uppercase letter or Chinese. It cannot start with http:// or https. It can contain Chinese, English, numbers,
     * half-width colons (:), underscores (_), half-width periods (.), or dashes (-). The default value is the InstanceId of
     * the instance.
     * 
     */
    @Export(name="instanceName", type=String.class, parameters={})
    private Output</* @Nullable */ String> instanceName;

    /**
     * @return The name of the instance. It must be 2 to 128 characters in length and must start with an
     * uppercase letter or Chinese. It cannot start with http:// or https. It can contain Chinese, English, numbers,
     * half-width colons (:), underscores (_), half-width periods (.), or dashes (-). The default value is the InstanceId of
     * the instance.
     * 
     */
    public Output<Optional<String>> instanceName() {
        return Codegen.optional(this.instanceName);
    }
    /**
     * Instance Type.
     * 
     */
    @Export(name="instanceType", type=String.class, parameters={})
    private Output<String> instanceType;

    /**
     * @return Instance Type.
     * 
     */
    public Output<String> instanceType() {
        return this.instanceType;
    }
    /**
     * The name of the key pair of the mobile phone instance.
     * 
     */
    @Export(name="keyPairName", type=String.class, parameters={})
    private Output</* @Nullable */ String> keyPairName;

    /**
     * @return The name of the key pair of the mobile phone instance.
     * 
     */
    public Output<Optional<String>> keyPairName() {
        return Codegen.optional(this.keyPairName);
    }
    /**
     * The payment type.Valid values: `PayAsYouGo`,`Subscription`
     * 
     */
    @Export(name="paymentType", type=String.class, parameters={})
    private Output</* @Nullable */ String> paymentType;

    /**
     * @return The payment type.Valid values: `PayAsYouGo`,`Subscription`
     * 
     */
    public Output<Optional<String>> paymentType() {
        return Codegen.optional(this.paymentType);
    }
    /**
     * The period. It is valid when `period_unit` is &#39;Year&#39;. Valid value: `1`, `2`, `3`, `4`, `5`. It
     * is valid when `period_unit` is &#39;Month&#39;. Valid value: `1`, `2`, `3`, `5`
     * 
     */
    @Export(name="period", type=String.class, parameters={})
    private Output</* @Nullable */ String> period;

    /**
     * @return The period. It is valid when `period_unit` is &#39;Year&#39;. Valid value: `1`, `2`, `3`, `4`, `5`. It
     * is valid when `period_unit` is &#39;Month&#39;. Valid value: `1`, `2`, `3`, `5`
     * 
     */
    public Output<Optional<String>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * The duration unit that you will buy the resource. Valid value: `Year`,`Month`. Default
     * to `Month`.
     * 
     */
    @Export(name="periodUnit", type=String.class, parameters={})
    private Output</* @Nullable */ String> periodUnit;

    /**
     * @return The duration unit that you will buy the resource. Valid value: `Year`,`Month`. Default
     * to `Month`.
     * 
     */
    public Output<Optional<String>> periodUnit() {
        return Codegen.optional(this.periodUnit);
    }
    /**
     * The selected resolution for the cloud mobile phone instance.
     * 
     */
    @Export(name="resolution", type=String.class, parameters={})
    private Output<String> resolution;

    /**
     * @return The selected resolution for the cloud mobile phone instance.
     * 
     */
    public Output<String> resolution() {
        return this.resolution;
    }
    /**
     * The ID of the security group. The security group is the same as that of the
     * ECS instance.
     * 
     */
    @Export(name="securityGroupId", type=String.class, parameters={})
    private Output<String> securityGroupId;

    /**
     * @return The ID of the security group. The security group is the same as that of the
     * ECS instance.
     * 
     */
    public Output<String> securityGroupId() {
        return this.securityGroupId;
    }
    /**
     * Instance status. Valid values: `Running`, `Stopped`.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return Instance status. Valid values: `Running`, `Stopped`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Cloud mobile phone VNC password. The password must be six characters in length and must
     * contain only uppercase, lowercase English letters and Arabic numerals.
     * 
     */
    @Export(name="vncPassword", type=String.class, parameters={})
    private Output</* @Nullable */ String> vncPassword;

    /**
     * @return Cloud mobile phone VNC password. The password must be six characters in length and must
     * contain only uppercase, lowercase English letters and Arabic numerals.
     * 
     */
    public Output<Optional<String>> vncPassword() {
        return Codegen.optional(this.vncPassword);
    }
    /**
     * The vswitch id.
     * 
     */
    @Export(name="vswitchId", type=String.class, parameters={})
    private Output<String> vswitchId;

    /**
     * @return The vswitch id.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Instance(String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(String name, InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(String name, InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecp/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecp/instance:Instance", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "vncPassword"
            ))
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
    public static Instance get(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
