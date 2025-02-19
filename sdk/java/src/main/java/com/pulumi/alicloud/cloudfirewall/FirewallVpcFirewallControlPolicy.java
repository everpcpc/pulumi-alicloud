// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudfirewall;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cloudfirewall.FirewallVpcFirewallControlPolicyArgs;
import com.pulumi.alicloud.cloudfirewall.inputs.FirewallVpcFirewallControlPolicyState;
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
 * Provides a Cloud Firewall Vpc Firewall Control Policy resource.
 * 
 * For information about Cloud Firewall Vpc Firewall Control Policy and how to use it, see [What is Vpc Firewall Control Policy](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallcontrolpolicy).
 * 
 * &gt; **NOTE:** Available in v1.194.0+.
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.cen.Instance;
 * import com.pulumi.alicloud.cen.InstanceArgs;
 * import com.pulumi.alicloud.cloudfirewall.FirewallVpcFirewallControlPolicy;
 * import com.pulumi.alicloud.cloudfirewall.FirewallVpcFirewallControlPolicyArgs;
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
 *         final var defaultAccount = AlicloudFunctions.getAccount();
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .cenInstanceName(var_.name())
 *             .description(&#34;example_value&#34;)
 *             .tags(Map.ofEntries(
 *                 Map.entry(&#34;Created&#34;, &#34;TF&#34;),
 *                 Map.entry(&#34;For&#34;, &#34;acceptance test&#34;)
 *             ))
 *             .build());
 * 
 *         var defaultFirewallVpcFirewallControlPolicy = new FirewallVpcFirewallControlPolicy(&#34;defaultFirewallVpcFirewallControlPolicy&#34;, FirewallVpcFirewallControlPolicyArgs.builder()        
 *             .order(&#34;1&#34;)
 *             .destination(&#34;127.0.0.2/32&#34;)
 *             .applicationName(&#34;ANY&#34;)
 *             .description(&#34;example_value&#34;)
 *             .sourceType(&#34;net&#34;)
 *             .destPort(&#34;80/88&#34;)
 *             .aclAction(&#34;accept&#34;)
 *             .lang(&#34;zh&#34;)
 *             .destinationType(&#34;net&#34;)
 *             .source(&#34;127.0.0.1/32&#34;)
 *             .destPortType(&#34;port&#34;)
 *             .proto(&#34;TCP&#34;)
 *             .release(true)
 *             .memberUid(defaultAccount.applyValue(getAccountResult -&gt; getAccountResult.id()))
 *             .vpcFirewallId(defaultInstance.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Cloud Firewall Vpc Firewall Control Policy can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:cloudfirewall/firewallVpcFirewallControlPolicy:FirewallVpcFirewallControlPolicy example &lt;vpc_firewall_id&gt;:&lt;acl_uuid&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cloudfirewall/firewallVpcFirewallControlPolicy:FirewallVpcFirewallControlPolicy")
public class FirewallVpcFirewallControlPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
     * 
     */
    @Export(name="aclAction", type=String.class, parameters={})
    private Output<String> aclAction;

    /**
     * @return The action that Cloud Firewall performs on the traffic. Valid values: `accept`, `drop`, `log`.
     * 
     */
    public Output<String> aclAction() {
        return this.aclAction;
    }
    /**
     * Access control over VPC firewalls strategy unique identifier.
     * 
     */
    @Export(name="aclUuid", type=String.class, parameters={})
    private Output<String> aclUuid;

    /**
     * @return Access control over VPC firewalls strategy unique identifier.
     * 
     */
    public Output<String> aclUuid() {
        return this.aclUuid;
    }
    /**
     * Policy specifies the application ID.
     * 
     */
    @Export(name="applicationId", type=String.class, parameters={})
    private Output<String> applicationId;

    /**
     * @return Policy specifies the application ID.
     * 
     */
    public Output<String> applicationId() {
        return this.applicationId;
    }
    /**
     * The type of the applications that the access control policy supports. Valid values: `FTP`, `HTTP`, `HTTPS`, `MySQL`, `SMTP`, `SMTPS`, `RDP`, `VNC`, `SSH`, `Redis`, `MQTT`, `MongoDB`, `Memcache`, `SSL`, `ANY`.
     * 
     */
    @Export(name="applicationName", type=String.class, parameters={})
    private Output<String> applicationName;

    /**
     * @return The type of the applications that the access control policy supports. Valid values: `FTP`, `HTTP`, `HTTPS`, `MySQL`, `SMTP`, `SMTPS`, `RDP`, `VNC`, `SSH`, `Redis`, `MQTT`, `MongoDB`, `Memcache`, `SSL`, `ANY`.
     * 
     */
    public Output<String> applicationName() {
        return this.applicationName;
    }
    /**
     * Access control over VPC firewalls description of the strategy information.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output<String> description;

    /**
     * @return Access control over VPC firewalls description of the strategy information.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The destination port in the access control policy. **Note:** If `dest_port_type` is set to `port`, you must specify this parameter.
     * 
     */
    @Export(name="destPort", type=String.class, parameters={})
    private Output</* @Nullable */ String> destPort;

    /**
     * @return The destination port in the access control policy. **Note:** If `dest_port_type` is set to `port`, you must specify this parameter.
     * 
     */
    public Output<Optional<String>> destPort() {
        return Codegen.optional(this.destPort);
    }
    /**
     * Access control policy in the access traffic of the destination port address book name. **Note:** If `dest_port_type` is set to `group`, you must specify this parameter.
     * 
     */
    @Export(name="destPortGroup", type=String.class, parameters={})
    private Output</* @Nullable */ String> destPortGroup;

    /**
     * @return Access control policy in the access traffic of the destination port address book name. **Note:** If `dest_port_type` is set to `group`, you must specify this parameter.
     * 
     */
    public Output<Optional<String>> destPortGroup() {
        return Codegen.optional(this.destPortGroup);
    }
    /**
     * Port Address Book port list.
     * 
     */
    @Export(name="destPortGroupPorts", type=List.class, parameters={String.class})
    private Output<List<String>> destPortGroupPorts;

    /**
     * @return Port Address Book port list.
     * 
     */
    public Output<List<String>> destPortGroupPorts() {
        return this.destPortGroupPorts;
    }
    /**
     * The type of the destination port in the access control policy. Valid values: `port`, `group`.
     * 
     */
    @Export(name="destPortType", type=String.class, parameters={})
    private Output</* @Nullable */ String> destPortType;

    /**
     * @return The type of the destination port in the access control policy. Valid values: `port`, `group`.
     * 
     */
    public Output<Optional<String>> destPortType() {
        return Codegen.optional(this.destPortType);
    }
    /**
     * The destination address in the access control policy. Valid values:
     * - If `destination_type` is set to `net`, the value of `destination` must be a CIDR block.
     * - If `destination_type` is set to `group`, the value of `destination` must be an address book.
     * - If `destination_type` is set to `domain`, the value of `destination` must be a domain name.
     * 
     */
    @Export(name="destination", type=String.class, parameters={})
    private Output<String> destination;

    /**
     * @return The destination address in the access control policy. Valid values:
     * - If `destination_type` is set to `net`, the value of `destination` must be a CIDR block.
     * - If `destination_type` is set to `group`, the value of `destination` must be an address book.
     * - If `destination_type` is set to `domain`, the value of `destination` must be a domain name.
     * 
     */
    public Output<String> destination() {
        return this.destination;
    }
    /**
     * Destination address book defined in the address list.
     * 
     */
    @Export(name="destinationGroupCidrs", type=List.class, parameters={String.class})
    private Output<List<String>> destinationGroupCidrs;

    /**
     * @return Destination address book defined in the address list.
     * 
     */
    public Output<List<String>> destinationGroupCidrs() {
        return this.destinationGroupCidrs;
    }
    /**
     * The destination address book type in the access control policy. Value:
     * - **ip**:IP address book, which contains one or more ip address segments.
     * - **domain**: domain address book, which contains one or more domain addresses.
     * 
     */
    @Export(name="destinationGroupType", type=String.class, parameters={})
    private Output<String> destinationGroupType;

    /**
     * @return The destination address book type in the access control policy. Value:
     * - **ip**:IP address book, which contains one or more ip address segments.
     * - **domain**: domain address book, which contains one or more domain addresses.
     * 
     */
    public Output<String> destinationGroupType() {
        return this.destinationGroupType;
    }
    /**
     * The type of the destination address in the access control policy. Valid values: `net`, `group`, `domain`.
     * 
     */
    @Export(name="destinationType", type=String.class, parameters={})
    private Output<String> destinationType;

    /**
     * @return The type of the destination address in the access control policy. Valid values: `net`, `group`, `domain`.
     * 
     */
    public Output<String> destinationType() {
        return this.destinationType;
    }
    /**
     * Control strategy of hits per second.
     * 
     */
    @Export(name="hitTimes", type=Integer.class, parameters={})
    private Output<Integer> hitTimes;

    /**
     * @return Control strategy of hits per second.
     * 
     */
    public Output<Integer> hitTimes() {
        return this.hitTimes;
    }
    /**
     * The language of the content within the request and response. Valid values: `zh`, `en`.
     * 
     */
    @Export(name="lang", type=String.class, parameters={})
    private Output</* @Nullable */ String> lang;

    /**
     * @return The language of the content within the request and response. Valid values: `zh`, `en`.
     * 
     */
    public Output<Optional<String>> lang() {
        return Codegen.optional(this.lang);
    }
    /**
     * The UID of the member account of the current Alibaba cloud account.
     * 
     */
    @Export(name="memberUid", type=String.class, parameters={})
    private Output<String> memberUid;

    /**
     * @return The UID of the member account of the current Alibaba cloud account.
     * 
     */
    public Output<String> memberUid() {
        return this.memberUid;
    }
    /**
     * The priority of the access control policy. The priority value starts from 1. A smaller priority value indicates a higher priority.
     * 
     */
    @Export(name="order", type=Integer.class, parameters={})
    private Output<Integer> order;

    /**
     * @return The priority of the access control policy. The priority value starts from 1. A smaller priority value indicates a higher priority.
     * 
     */
    public Output<Integer> order() {
        return this.order;
    }
    /**
     * The type of the protocol in the access control policy. Valid values: `ANY`, `TCP`, `UDP`, `ICMP`.
     * 
     */
    @Export(name="proto", type=String.class, parameters={})
    private Output<String> proto;

    /**
     * @return The type of the protocol in the access control policy. Valid values: `ANY`, `TCP`, `UDP`, `ICMP`.
     * 
     */
    public Output<String> proto() {
        return this.proto;
    }
    /**
     * The enabled status of the access control policy. The policy is enabled by default after it is created. Value:
     * - **true**: Enable access control policies
     * - **false**: does not enable access control policies.
     * 
     */
    @Export(name="release", type=Boolean.class, parameters={})
    private Output<Boolean> release;

    /**
     * @return The enabled status of the access control policy. The policy is enabled by default after it is created. Value:
     * - **true**: Enable access control policies
     * - **false**: does not enable access control policies.
     * 
     */
    public Output<Boolean> release() {
        return this.release;
    }
    /**
     * Access control over VPC firewalls strategy in the source address.
     * 
     */
    @Export(name="source", type=String.class, parameters={})
    private Output<String> source;

    /**
     * @return Access control over VPC firewalls strategy in the source address.
     * 
     */
    public Output<String> source() {
        return this.source;
    }
    /**
     * SOURCE address of the address list.
     * 
     */
    @Export(name="sourceGroupCidrs", type=List.class, parameters={String.class})
    private Output<List<String>> sourceGroupCidrs;

    /**
     * @return SOURCE address of the address list.
     * 
     */
    public Output<List<String>> sourceGroupCidrs() {
        return this.sourceGroupCidrs;
    }
    /**
     * The source address type in the access control policy. Unique value: **ip**. The IP address book contains one or more IP address segments.
     * 
     */
    @Export(name="sourceGroupType", type=String.class, parameters={})
    private Output<String> sourceGroupType;

    /**
     * @return The source address type in the access control policy. Unique value: **ip**. The IP address book contains one or more IP address segments.
     * 
     */
    public Output<String> sourceGroupType() {
        return this.sourceGroupType;
    }
    /**
     * The type of the source address in the access control policy. Valid values: `net`, `group`.
     * 
     */
    @Export(name="sourceType", type=String.class, parameters={})
    private Output<String> sourceType;

    /**
     * @return The type of the source address in the access control policy. Valid values: `net`, `group`.
     * 
     */
    public Output<String> sourceType() {
        return this.sourceType;
    }
    /**
     * The ID of the VPC firewall instance. Value:
     * - When the VPC firewall protects traffic between two VPCs connected through the cloud enterprise network, the policy group ID uses the cloud enterprise network instance ID.
     * - When the VPC firewall protects traffic between two VPCs connected through the express connection, the policy group ID uses the ID of the VPC firewall instance.
     * 
     */
    @Export(name="vpcFirewallId", type=String.class, parameters={})
    private Output<String> vpcFirewallId;

    /**
     * @return The ID of the VPC firewall instance. Value:
     * - When the VPC firewall protects traffic between two VPCs connected through the cloud enterprise network, the policy group ID uses the cloud enterprise network instance ID.
     * - When the VPC firewall protects traffic between two VPCs connected through the express connection, the policy group ID uses the ID of the VPC firewall instance.
     * 
     */
    public Output<String> vpcFirewallId() {
        return this.vpcFirewallId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FirewallVpcFirewallControlPolicy(String name) {
        this(name, FirewallVpcFirewallControlPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FirewallVpcFirewallControlPolicy(String name, FirewallVpcFirewallControlPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FirewallVpcFirewallControlPolicy(String name, FirewallVpcFirewallControlPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudfirewall/firewallVpcFirewallControlPolicy:FirewallVpcFirewallControlPolicy", name, args == null ? FirewallVpcFirewallControlPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FirewallVpcFirewallControlPolicy(String name, Output<String> id, @Nullable FirewallVpcFirewallControlPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudfirewall/firewallVpcFirewallControlPolicy:FirewallVpcFirewallControlPolicy", name, state, makeResourceOptions(options, id));
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
    public static FirewallVpcFirewallControlPolicy get(String name, Output<String> id, @Nullable FirewallVpcFirewallControlPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FirewallVpcFirewallControlPolicy(name, id, state, options);
    }
}
