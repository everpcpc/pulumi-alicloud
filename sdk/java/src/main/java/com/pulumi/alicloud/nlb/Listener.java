// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.nlb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.nlb.ListenerArgs;
import com.pulumi.alicloud.nlb.inputs.ListenerState;
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
 * Provides a NLB Listener resource.
 * 
 * For information about NLB Listener and how to use it, see [What is Listener](https://www.alibabacloud.com/help/en/server-load-balancer/latest/createlistener-nl).
 * 
 * &gt; **NOTE:** Available in v1.191.0+.
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
 * import com.pulumi.alicloud.vpc.VpcFunctions;
 * import com.pulumi.alicloud.vpc.inputs.GetNetworksArgs;
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.nlb.ServerGroup;
 * import com.pulumi.alicloud.nlb.ServerGroupArgs;
 * import com.pulumi.alicloud.nlb.inputs.ServerGroupHealthCheckArgs;
 * import com.pulumi.alicloud.nlb.NlbFunctions;
 * import com.pulumi.alicloud.nlb.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.inputs.GetSwitchesArgs;
 * import com.pulumi.alicloud.nlb.LoadBalancer;
 * import com.pulumi.alicloud.nlb.LoadBalancerArgs;
 * import com.pulumi.alicloud.nlb.inputs.LoadBalancerZoneMappingArgs;
 * import com.pulumi.alicloud.nlb.Listener;
 * import com.pulumi.alicloud.nlb.ListenerArgs;
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
 *         final var defaultNetworks = VpcFunctions.getNetworks(GetNetworksArgs.builder()
 *             .nameRegex(&#34;default-NODELETING&#34;)
 *             .build());
 * 
 *         final var defaultResourceGroups = ResourcemanagerFunctions.getResourceGroups();
 * 
 *         var defaultServerGroup = new ServerGroup(&#34;defaultServerGroup&#34;, ServerGroupArgs.builder()        
 *             .resourceGroupId(defaultResourceGroups.applyValue(getResourceGroupsResult -&gt; getResourceGroupsResult.ids()[0]))
 *             .serverGroupName(var_.name())
 *             .serverGroupType(&#34;Instance&#34;)
 *             .vpcId(defaultNetworks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .scheduler(&#34;Wrr&#34;)
 *             .protocol(&#34;TCP&#34;)
 *             .healthCheck(ServerGroupHealthCheckArgs.builder()
 *                 .healthCheckUrl(&#34;/test/index.html&#34;)
 *                 .healthCheckDomain(&#34;tf-testAcc.com&#34;)
 *                 .healthCheckEnabled(true)
 *                 .healthCheckType(&#34;TCP&#34;)
 *                 .healthCheckConnectPort(0)
 *                 .healthyThreshold(2)
 *                 .unhealthyThreshold(2)
 *                 .healthCheckConnectTimeout(5)
 *                 .healthCheckInterval(10)
 *                 .httpCheckMethod(&#34;GET&#34;)
 *                 .healthCheckHttpCodes(                
 *                     &#34;http_2xx&#34;,
 *                     &#34;http_3xx&#34;,
 *                     &#34;http_4xx&#34;)
 *                 .build())
 *             .connectionDrain(true)
 *             .connectionDrainTimeout(60)
 *             .preserveClientIpEnabled(true)
 *             .tags(Map.of(&#34;Created&#34;, &#34;TF&#34;))
 *             .addressIpVersion(&#34;Ipv4&#34;)
 *             .build());
 * 
 *         final var defaultZones = NlbFunctions.getZones();
 * 
 *         final var default1 = VpcFunctions.getSwitches(GetSwitchesArgs.builder()
 *             .vpcId(defaultNetworks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         final var default2 = VpcFunctions.getSwitches(GetSwitchesArgs.builder()
 *             .vpcId(defaultNetworks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[1].id()))
 *             .build());
 * 
 *         final var zoneId1 = defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id());
 * 
 *         final var vswitchId1 = default1.applyValue(getSwitchesResult -&gt; getSwitchesResult.ids()[0]);
 * 
 *         final var zoneId2 = defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[1].id());
 * 
 *         final var vswitchId2 = default2.applyValue(getSwitchesResult -&gt; getSwitchesResult.ids()[0]);
 * 
 *         var defaultLoadBalancer = new LoadBalancer(&#34;defaultLoadBalancer&#34;, LoadBalancerArgs.builder()        
 *             .loadBalancerName(var_.name())
 *             .resourceGroupId(defaultResourceGroups.applyValue(getResourceGroupsResult -&gt; getResourceGroupsResult.ids()[0]))
 *             .loadBalancerType(&#34;Network&#34;)
 *             .addressType(&#34;Internet&#34;)
 *             .addressIpVersion(&#34;Ipv4&#34;)
 *             .tags(Map.ofEntries(
 *                 Map.entry(&#34;Created&#34;, &#34;tfTestAcc0&#34;),
 *                 Map.entry(&#34;For&#34;, &#34;Tftestacc 0&#34;)
 *             ))
 *             .vpcId(defaultNetworks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .zoneMappings(            
 *                 LoadBalancerZoneMappingArgs.builder()
 *                     .vswitchId(vswitchId1)
 *                     .zoneId(zoneId1)
 *                     .build(),
 *                 LoadBalancerZoneMappingArgs.builder()
 *                     .vswitchId(vswitchId2)
 *                     .zoneId(zoneId2)
 *                     .build())
 *             .build());
 * 
 *         var defaultListener = new Listener(&#34;defaultListener&#34;, ListenerArgs.builder()        
 *             .listenerProtocol(&#34;TCP&#34;)
 *             .listenerPort(&#34;80&#34;)
 *             .listenerDescription(var_.name())
 *             .loadBalancerId(defaultLoadBalancer.id())
 *             .serverGroupId(defaultServerGroup.id())
 *             .idleTimeout(&#34;900&#34;)
 *             .proxyProtocolEnabled(&#34;true&#34;)
 *             .secSensorEnabled(&#34;true&#34;)
 *             .cps(&#34;10000&#34;)
 *             .mss(&#34;0&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * NLB Listener can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:nlb/listener:Listener example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:nlb/listener:Listener")
public class Listener extends com.pulumi.resources.CustomResource {
    /**
     * Specifies whether to enable Application-Layer Protocol Negotiation (ALPN).
     * 
     */
    @Export(name="alpnEnabled", type=Boolean.class, parameters={})
    private Output<Boolean> alpnEnabled;

    /**
     * @return Specifies whether to enable Application-Layer Protocol Negotiation (ALPN).
     * 
     */
    public Output<Boolean> alpnEnabled() {
        return this.alpnEnabled;
    }
    /**
     * The ALPN policy.
     * 
     */
    @Export(name="alpnPolicy", type=String.class, parameters={})
    private Output</* @Nullable */ String> alpnPolicy;

    /**
     * @return The ALPN policy.
     * 
     */
    public Output<Optional<String>> alpnPolicy() {
        return Codegen.optional(this.alpnPolicy);
    }
    /**
     * The list of certificate authority (CA) certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one CA certificate is supported.
     * 
     */
    @Export(name="caCertificateIds", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> caCertificateIds;

    /**
     * @return The list of certificate authority (CA) certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one CA certificate is supported.
     * 
     */
    public Output<Optional<List<String>>> caCertificateIds() {
        return Codegen.optional(this.caCertificateIds);
    }
    /**
     * Specifies whether to enable mutual authentication.
     * 
     */
    @Export(name="caEnabled", type=Boolean.class, parameters={})
    private Output<Boolean> caEnabled;

    /**
     * @return Specifies whether to enable mutual authentication.
     * 
     */
    public Output<Boolean> caEnabled() {
        return this.caEnabled;
    }
    /**
     * The list of server certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one server certificate is supported.
     * 
     */
    @Export(name="certificateIds", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> certificateIds;

    /**
     * @return The list of server certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one server certificate is supported.
     * 
     */
    public Output<Optional<List<String>>> certificateIds() {
        return Codegen.optional(this.certificateIds);
    }
    /**
     * The maximum number of connections that can be created per second on the NLB instance. Valid values: 0 to 1000000. 0 specifies that the number of connections is unlimited.
     * 
     */
    @Export(name="cps", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> cps;

    /**
     * @return The maximum number of connections that can be created per second on the NLB instance. Valid values: 0 to 1000000. 0 specifies that the number of connections is unlimited.
     * 
     */
    public Output<Optional<Integer>> cps() {
        return Codegen.optional(this.cps);
    }
    /**
     * Full port listening end port. Valid values: `0` ~ `65535`. The value of the end port is less than the start port.
     * 
     */
    @Export(name="endPort", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> endPort;

    /**
     * @return Full port listening end port. Valid values: `0` ~ `65535`. The value of the end port is less than the start port.
     * 
     */
    public Output<Optional<Integer>> endPort() {
        return Codegen.optional(this.endPort);
    }
    /**
     * The timeout period of an idle connection. Unit: seconds. Valid values: `1` to `900`. Default value: `900`.
     * 
     */
    @Export(name="idleTimeout", type=Integer.class, parameters={})
    private Output<Integer> idleTimeout;

    /**
     * @return The timeout period of an idle connection. Unit: seconds. Valid values: `1` to `900`. Default value: `900`.
     * 
     */
    public Output<Integer> idleTimeout() {
        return this.idleTimeout;
    }
    /**
     * Custom listener name. The length is limited to 2 to 256 characters, supports Chinese and English letters, and can include numbers, commas (,), half-width periods (.), half-width semicolons (;), forward slashes (/), at(@), underscores (_), and dashes (-).
     * 
     */
    @Export(name="listenerDescription", type=String.class, parameters={})
    private Output</* @Nullable */ String> listenerDescription;

    /**
     * @return Custom listener name. The length is limited to 2 to 256 characters, supports Chinese and English letters, and can include numbers, commas (,), half-width periods (.), half-width semicolons (;), forward slashes (/), at(@), underscores (_), and dashes (-).
     * 
     */
    public Output<Optional<String>> listenerDescription() {
        return Codegen.optional(this.listenerDescription);
    }
    /**
     * Listening port. Valid values: 0 ~ 65535. `0`: indicates that full port listening is used. When set to `0`, you must configure `StartPort` and `EndPort`.
     * 
     */
    @Export(name="listenerPort", type=Integer.class, parameters={})
    private Output<Integer> listenerPort;

    /**
     * @return Listening port. Valid values: 0 ~ 65535. `0`: indicates that full port listening is used. When set to `0`, you must configure `StartPort` and `EndPort`.
     * 
     */
    public Output<Integer> listenerPort() {
        return this.listenerPort;
    }
    /**
     * The listening protocol. Valid values: `TCP`, `UDP`, or `TCPSSL`.
     * 
     */
    @Export(name="listenerProtocol", type=String.class, parameters={})
    private Output<String> listenerProtocol;

    /**
     * @return The listening protocol. Valid values: `TCP`, `UDP`, or `TCPSSL`.
     * 
     */
    public Output<String> listenerProtocol() {
        return this.listenerProtocol;
    }
    /**
     * The ID of the network-based server load balancer instance.
     * 
     */
    @Export(name="loadBalancerId", type=String.class, parameters={})
    private Output<String> loadBalancerId;

    /**
     * @return The ID of the network-based server load balancer instance.
     * 
     */
    public Output<String> loadBalancerId() {
        return this.loadBalancerId;
    }
    /**
     * The maximum size of a TCP segment. Unit: bytes. Valid values: 0 to 1500. 0 specifies that the maximum segment size remains unchanged. **Note:** This parameter is supported only by listeners that use SSL over TCP.
     * 
     */
    @Export(name="mss", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> mss;

    /**
     * @return The maximum size of a TCP segment. Unit: bytes. Valid values: 0 to 1500. 0 specifies that the maximum segment size remains unchanged. **Note:** This parameter is supported only by listeners that use SSL over TCP.
     * 
     */
    public Output<Optional<Integer>> mss() {
        return Codegen.optional(this.mss);
    }
    /**
     * Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers.
     * 
     */
    @Export(name="proxyProtocolEnabled", type=Boolean.class, parameters={})
    private Output<Boolean> proxyProtocolEnabled;

    /**
     * @return Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers.
     * 
     */
    public Output<Boolean> proxyProtocolEnabled() {
        return this.proxyProtocolEnabled;
    }
    /**
     * Specifies whether to enable fine-grained monitoring.
     * 
     */
    @Export(name="secSensorEnabled", type=Boolean.class, parameters={})
    private Output<Boolean> secSensorEnabled;

    /**
     * @return Specifies whether to enable fine-grained monitoring.
     * 
     */
    public Output<Boolean> secSensorEnabled() {
        return this.secSensorEnabled;
    }
    /**
     * The ID of the security policy. System security policies and custom security policies are supported. Valid values: `tls_cipher_policy_1_0` (default), `tls_cipher_policy_1_1,` `tls_cipher_policy_1_2`, `tls_cipher_policy_1_2_strict`, and `tls_cipher_policy_1_2_strict_with_1_3`.
     * 
     */
    @Export(name="securityPolicyId", type=String.class, parameters={})
    private Output<String> securityPolicyId;

    /**
     * @return The ID of the security policy. System security policies and custom security policies are supported. Valid values: `tls_cipher_policy_1_0` (default), `tls_cipher_policy_1_1,` `tls_cipher_policy_1_2`, `tls_cipher_policy_1_2_strict`, and `tls_cipher_policy_1_2_strict_with_1_3`.
     * 
     */
    public Output<String> securityPolicyId() {
        return this.securityPolicyId;
    }
    /**
     * The ID of the server group.
     * 
     */
    @Export(name="serverGroupId", type=String.class, parameters={})
    private Output<String> serverGroupId;

    /**
     * @return The ID of the server group.
     * 
     */
    public Output<String> serverGroupId() {
        return this.serverGroupId;
    }
    /**
     * Full Port listens to the starting port. Valid values: `0` ~ `65535`.
     * 
     */
    @Export(name="startPort", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> startPort;

    /**
     * @return Full Port listens to the starting port. Valid values: `0` ~ `65535`.
     * 
     */
    public Output<Optional<Integer>> startPort() {
        return Codegen.optional(this.startPort);
    }
    /**
     * The status of the resource. Valid values: `Running`, `Stopped`.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the resource. Valid values: `Running`, `Stopped`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Listener(String name) {
        this(name, ListenerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Listener(String name, ListenerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Listener(String name, ListenerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:nlb/listener:Listener", name, args == null ? ListenerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Listener(String name, Output<String> id, @Nullable ListenerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:nlb/listener:Listener", name, state, makeResourceOptions(options, id));
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
    public static Listener get(String name, Output<String> id, @Nullable ListenerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Listener(name, id, state, options);
    }
}