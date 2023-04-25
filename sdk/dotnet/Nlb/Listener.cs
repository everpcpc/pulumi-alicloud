// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Nlb
{
    /// <summary>
    /// Provides a NLB Listener resource.
    /// 
    /// For information about NLB Listener and how to use it, see [What is Listener](https://www.alibabacloud.com/help/en/server-load-balancer/latest/createlistener-nl).
    /// 
    /// &gt; **NOTE:** Available in v1.191.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultNetworks = AliCloud.Vpc.GetNetworks.Invoke(new()
    ///     {
    ///         NameRegex = "default-NODELETING",
    ///     });
    /// 
    ///     var defaultResourceGroups = AliCloud.ResourceManager.GetResourceGroups.Invoke();
    /// 
    ///     var defaultServerGroup = new AliCloud.Nlb.ServerGroup("defaultServerGroup", new()
    ///     {
    ///         ResourceGroupId = defaultResourceGroups.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Ids[0]),
    ///         ServerGroupName = @var.Name,
    ///         ServerGroupType = "Instance",
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         Scheduler = "Wrr",
    ///         Protocol = "TCP",
    ///         HealthCheck = new AliCloud.Nlb.Inputs.ServerGroupHealthCheckArgs
    ///         {
    ///             HealthCheckUrl = "/test/index.html",
    ///             HealthCheckDomain = "tf-testAcc.com",
    ///             HealthCheckEnabled = true,
    ///             HealthCheckType = "TCP",
    ///             HealthCheckConnectPort = 0,
    ///             HealthyThreshold = 2,
    ///             UnhealthyThreshold = 2,
    ///             HealthCheckConnectTimeout = 5,
    ///             HealthCheckInterval = 10,
    ///             HttpCheckMethod = "GET",
    ///             HealthCheckHttpCodes = new[]
    ///             {
    ///                 "http_2xx",
    ///                 "http_3xx",
    ///                 "http_4xx",
    ///             },
    ///         },
    ///         ConnectionDrain = true,
    ///         ConnectionDrainTimeout = 60,
    ///         PreserveClientIpEnabled = true,
    ///         Tags = 
    ///         {
    ///             { "Created", "TF" },
    ///         },
    ///         AddressIpVersion = "Ipv4",
    ///     });
    /// 
    ///     var defaultZones = AliCloud.Nlb.GetZones.Invoke();
    /// 
    ///     var default1 = AliCloud.Vpc.GetSwitches.Invoke(new()
    ///     {
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var default2 = AliCloud.Vpc.GetSwitches.Invoke(new()
    ///     {
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[1]?.Id),
    ///     });
    /// 
    ///     var zoneId1 = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id);
    /// 
    ///     var vswitchId1 = default1.Apply(getSwitchesResult =&gt; getSwitchesResult.Ids[0]);
    /// 
    ///     var zoneId2 = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[1]?.Id);
    /// 
    ///     var vswitchId2 = default2.Apply(getSwitchesResult =&gt; getSwitchesResult.Ids[0]);
    /// 
    ///     var defaultLoadBalancer = new AliCloud.Nlb.LoadBalancer("defaultLoadBalancer", new()
    ///     {
    ///         LoadBalancerName = @var.Name,
    ///         ResourceGroupId = defaultResourceGroups.Apply(getResourceGroupsResult =&gt; getResourceGroupsResult.Ids[0]),
    ///         LoadBalancerType = "Network",
    ///         AddressType = "Internet",
    ///         AddressIpVersion = "Ipv4",
    ///         Tags = 
    ///         {
    ///             { "Created", "tfTestAcc0" },
    ///             { "For", "Tftestacc 0" },
    ///         },
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         ZoneMappings = new[]
    ///         {
    ///             new AliCloud.Nlb.Inputs.LoadBalancerZoneMappingArgs
    ///             {
    ///                 VswitchId = vswitchId1,
    ///                 ZoneId = zoneId1,
    ///             },
    ///             new AliCloud.Nlb.Inputs.LoadBalancerZoneMappingArgs
    ///             {
    ///                 VswitchId = vswitchId2,
    ///                 ZoneId = zoneId2,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var defaultListener = new AliCloud.Nlb.Listener("defaultListener", new()
    ///     {
    ///         ListenerProtocol = "TCP",
    ///         ListenerPort = 80,
    ///         ListenerDescription = @var.Name,
    ///         LoadBalancerId = defaultLoadBalancer.Id,
    ///         ServerGroupId = defaultServerGroup.Id,
    ///         IdleTimeout = 900,
    ///         ProxyProtocolEnabled = true,
    ///         SecSensorEnabled = true,
    ///         Cps = 10000,
    ///         Mss = 0,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// NLB Listener can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:nlb/listener:Listener example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:nlb/listener:Listener")]
    public partial class Listener : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether to enable Application-Layer Protocol Negotiation (ALPN).
        /// </summary>
        [Output("alpnEnabled")]
        public Output<bool> AlpnEnabled { get; private set; } = null!;

        /// <summary>
        /// The ALPN policy.
        /// </summary>
        [Output("alpnPolicy")]
        public Output<string?> AlpnPolicy { get; private set; } = null!;

        /// <summary>
        /// The list of certificate authority (CA) certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one CA certificate is supported.
        /// </summary>
        [Output("caCertificateIds")]
        public Output<ImmutableArray<string>> CaCertificateIds { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable mutual authentication.
        /// </summary>
        [Output("caEnabled")]
        public Output<bool> CaEnabled { get; private set; } = null!;

        /// <summary>
        /// The list of server certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one server certificate is supported.
        /// </summary>
        [Output("certificateIds")]
        public Output<ImmutableArray<string>> CertificateIds { get; private set; } = null!;

        /// <summary>
        /// The maximum number of connections that can be created per second on the NLB instance. Valid values: 0 to 1000000. 0 specifies that the number of connections is unlimited.
        /// </summary>
        [Output("cps")]
        public Output<int?> Cps { get; private set; } = null!;

        /// <summary>
        /// Full port listening end port. Valid values: `0` ~ `65535`. The value of the end port is less than the start port.
        /// </summary>
        [Output("endPort")]
        public Output<int?> EndPort { get; private set; } = null!;

        /// <summary>
        /// The timeout period of an idle connection. Unit: seconds. Valid values: `1` to `900`. Default value: `900`.
        /// </summary>
        [Output("idleTimeout")]
        public Output<int> IdleTimeout { get; private set; } = null!;

        /// <summary>
        /// Custom listener name. The length is limited to 2 to 256 characters, supports Chinese and English letters, and can include numbers, commas (,), half-width periods (.), half-width semicolons (;), forward slashes (/), at(@), underscores (_), and dashes (-).
        /// </summary>
        [Output("listenerDescription")]
        public Output<string?> ListenerDescription { get; private set; } = null!;

        /// <summary>
        /// Listening port. Valid values: 0 ~ 65535. `0`: indicates that full port listening is used. When set to `0`, you must configure `StartPort` and `EndPort`.
        /// </summary>
        [Output("listenerPort")]
        public Output<int> ListenerPort { get; private set; } = null!;

        /// <summary>
        /// The listening protocol. Valid values: `TCP`, `UDP`, or `TCPSSL`.
        /// </summary>
        [Output("listenerProtocol")]
        public Output<string> ListenerProtocol { get; private set; } = null!;

        /// <summary>
        /// The ID of the network-based server load balancer instance.
        /// </summary>
        [Output("loadBalancerId")]
        public Output<string> LoadBalancerId { get; private set; } = null!;

        /// <summary>
        /// The maximum size of a TCP segment. Unit: bytes. Valid values: 0 to 1500. 0 specifies that the maximum segment size remains unchanged. **Note:** This parameter is supported only by listeners that use SSL over TCP.
        /// </summary>
        [Output("mss")]
        public Output<int?> Mss { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers.
        /// </summary>
        [Output("proxyProtocolEnabled")]
        public Output<bool> ProxyProtocolEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable fine-grained monitoring.
        /// </summary>
        [Output("secSensorEnabled")]
        public Output<bool> SecSensorEnabled { get; private set; } = null!;

        /// <summary>
        /// The ID of the security policy. System security policies and custom security policies are supported. Valid values: `tls_cipher_policy_1_0` (default), `tls_cipher_policy_1_1,` `tls_cipher_policy_1_2`, `tls_cipher_policy_1_2_strict`, and `tls_cipher_policy_1_2_strict_with_1_3`.
        /// </summary>
        [Output("securityPolicyId")]
        public Output<string> SecurityPolicyId { get; private set; } = null!;

        /// <summary>
        /// The ID of the server group.
        /// </summary>
        [Output("serverGroupId")]
        public Output<string> ServerGroupId { get; private set; } = null!;

        /// <summary>
        /// Full Port listens to the starting port. Valid values: `0` ~ `65535`.
        /// </summary>
        [Output("startPort")]
        public Output<int?> StartPort { get; private set; } = null!;

        /// <summary>
        /// The status of the resource. Valid values: `Running`, `Stopped`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Listener resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Listener(string name, ListenerArgs args, CustomResourceOptions? options = null)
            : base("alicloud:nlb/listener:Listener", name, args ?? new ListenerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Listener(string name, Input<string> id, ListenerState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:nlb/listener:Listener", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Listener resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Listener Get(string name, Input<string> id, ListenerState? state = null, CustomResourceOptions? options = null)
        {
            return new Listener(name, id, state, options);
        }
    }

    public sealed class ListenerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to enable Application-Layer Protocol Negotiation (ALPN).
        /// </summary>
        [Input("alpnEnabled")]
        public Input<bool>? AlpnEnabled { get; set; }

        /// <summary>
        /// The ALPN policy.
        /// </summary>
        [Input("alpnPolicy")]
        public Input<string>? AlpnPolicy { get; set; }

        [Input("caCertificateIds")]
        private InputList<string>? _caCertificateIds;

        /// <summary>
        /// The list of certificate authority (CA) certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one CA certificate is supported.
        /// </summary>
        public InputList<string> CaCertificateIds
        {
            get => _caCertificateIds ?? (_caCertificateIds = new InputList<string>());
            set => _caCertificateIds = value;
        }

        /// <summary>
        /// Specifies whether to enable mutual authentication.
        /// </summary>
        [Input("caEnabled")]
        public Input<bool>? CaEnabled { get; set; }

        [Input("certificateIds")]
        private InputList<string>? _certificateIds;

        /// <summary>
        /// The list of server certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one server certificate is supported.
        /// </summary>
        public InputList<string> CertificateIds
        {
            get => _certificateIds ?? (_certificateIds = new InputList<string>());
            set => _certificateIds = value;
        }

        /// <summary>
        /// The maximum number of connections that can be created per second on the NLB instance. Valid values: 0 to 1000000. 0 specifies that the number of connections is unlimited.
        /// </summary>
        [Input("cps")]
        public Input<int>? Cps { get; set; }

        /// <summary>
        /// Full port listening end port. Valid values: `0` ~ `65535`. The value of the end port is less than the start port.
        /// </summary>
        [Input("endPort")]
        public Input<int>? EndPort { get; set; }

        /// <summary>
        /// The timeout period of an idle connection. Unit: seconds. Valid values: `1` to `900`. Default value: `900`.
        /// </summary>
        [Input("idleTimeout")]
        public Input<int>? IdleTimeout { get; set; }

        /// <summary>
        /// Custom listener name. The length is limited to 2 to 256 characters, supports Chinese and English letters, and can include numbers, commas (,), half-width periods (.), half-width semicolons (;), forward slashes (/), at(@), underscores (_), and dashes (-).
        /// </summary>
        [Input("listenerDescription")]
        public Input<string>? ListenerDescription { get; set; }

        /// <summary>
        /// Listening port. Valid values: 0 ~ 65535. `0`: indicates that full port listening is used. When set to `0`, you must configure `StartPort` and `EndPort`.
        /// </summary>
        [Input("listenerPort", required: true)]
        public Input<int> ListenerPort { get; set; } = null!;

        /// <summary>
        /// The listening protocol. Valid values: `TCP`, `UDP`, or `TCPSSL`.
        /// </summary>
        [Input("listenerProtocol", required: true)]
        public Input<string> ListenerProtocol { get; set; } = null!;

        /// <summary>
        /// The ID of the network-based server load balancer instance.
        /// </summary>
        [Input("loadBalancerId", required: true)]
        public Input<string> LoadBalancerId { get; set; } = null!;

        /// <summary>
        /// The maximum size of a TCP segment. Unit: bytes. Valid values: 0 to 1500. 0 specifies that the maximum segment size remains unchanged. **Note:** This parameter is supported only by listeners that use SSL over TCP.
        /// </summary>
        [Input("mss")]
        public Input<int>? Mss { get; set; }

        /// <summary>
        /// Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers.
        /// </summary>
        [Input("proxyProtocolEnabled")]
        public Input<bool>? ProxyProtocolEnabled { get; set; }

        /// <summary>
        /// Specifies whether to enable fine-grained monitoring.
        /// </summary>
        [Input("secSensorEnabled")]
        public Input<bool>? SecSensorEnabled { get; set; }

        /// <summary>
        /// The ID of the security policy. System security policies and custom security policies are supported. Valid values: `tls_cipher_policy_1_0` (default), `tls_cipher_policy_1_1,` `tls_cipher_policy_1_2`, `tls_cipher_policy_1_2_strict`, and `tls_cipher_policy_1_2_strict_with_1_3`.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        /// <summary>
        /// The ID of the server group.
        /// </summary>
        [Input("serverGroupId", required: true)]
        public Input<string> ServerGroupId { get; set; } = null!;

        /// <summary>
        /// Full Port listens to the starting port. Valid values: `0` ~ `65535`.
        /// </summary>
        [Input("startPort")]
        public Input<int>? StartPort { get; set; }

        /// <summary>
        /// The status of the resource. Valid values: `Running`, `Stopped`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ListenerArgs()
        {
        }
        public static new ListenerArgs Empty => new ListenerArgs();
    }

    public sealed class ListenerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to enable Application-Layer Protocol Negotiation (ALPN).
        /// </summary>
        [Input("alpnEnabled")]
        public Input<bool>? AlpnEnabled { get; set; }

        /// <summary>
        /// The ALPN policy.
        /// </summary>
        [Input("alpnPolicy")]
        public Input<string>? AlpnPolicy { get; set; }

        [Input("caCertificateIds")]
        private InputList<string>? _caCertificateIds;

        /// <summary>
        /// The list of certificate authority (CA) certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one CA certificate is supported.
        /// </summary>
        public InputList<string> CaCertificateIds
        {
            get => _caCertificateIds ?? (_caCertificateIds = new InputList<string>());
            set => _caCertificateIds = value;
        }

        /// <summary>
        /// Specifies whether to enable mutual authentication.
        /// </summary>
        [Input("caEnabled")]
        public Input<bool>? CaEnabled { get; set; }

        [Input("certificateIds")]
        private InputList<string>? _certificateIds;

        /// <summary>
        /// The list of server certificates. This parameter takes effect only for listeners that use SSL over TCP. **Note:** Only one server certificate is supported.
        /// </summary>
        public InputList<string> CertificateIds
        {
            get => _certificateIds ?? (_certificateIds = new InputList<string>());
            set => _certificateIds = value;
        }

        /// <summary>
        /// The maximum number of connections that can be created per second on the NLB instance. Valid values: 0 to 1000000. 0 specifies that the number of connections is unlimited.
        /// </summary>
        [Input("cps")]
        public Input<int>? Cps { get; set; }

        /// <summary>
        /// Full port listening end port. Valid values: `0` ~ `65535`. The value of the end port is less than the start port.
        /// </summary>
        [Input("endPort")]
        public Input<int>? EndPort { get; set; }

        /// <summary>
        /// The timeout period of an idle connection. Unit: seconds. Valid values: `1` to `900`. Default value: `900`.
        /// </summary>
        [Input("idleTimeout")]
        public Input<int>? IdleTimeout { get; set; }

        /// <summary>
        /// Custom listener name. The length is limited to 2 to 256 characters, supports Chinese and English letters, and can include numbers, commas (,), half-width periods (.), half-width semicolons (;), forward slashes (/), at(@), underscores (_), and dashes (-).
        /// </summary>
        [Input("listenerDescription")]
        public Input<string>? ListenerDescription { get; set; }

        /// <summary>
        /// Listening port. Valid values: 0 ~ 65535. `0`: indicates that full port listening is used. When set to `0`, you must configure `StartPort` and `EndPort`.
        /// </summary>
        [Input("listenerPort")]
        public Input<int>? ListenerPort { get; set; }

        /// <summary>
        /// The listening protocol. Valid values: `TCP`, `UDP`, or `TCPSSL`.
        /// </summary>
        [Input("listenerProtocol")]
        public Input<string>? ListenerProtocol { get; set; }

        /// <summary>
        /// The ID of the network-based server load balancer instance.
        /// </summary>
        [Input("loadBalancerId")]
        public Input<string>? LoadBalancerId { get; set; }

        /// <summary>
        /// The maximum size of a TCP segment. Unit: bytes. Valid values: 0 to 1500. 0 specifies that the maximum segment size remains unchanged. **Note:** This parameter is supported only by listeners that use SSL over TCP.
        /// </summary>
        [Input("mss")]
        public Input<int>? Mss { get; set; }

        /// <summary>
        /// Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers.
        /// </summary>
        [Input("proxyProtocolEnabled")]
        public Input<bool>? ProxyProtocolEnabled { get; set; }

        /// <summary>
        /// Specifies whether to enable fine-grained monitoring.
        /// </summary>
        [Input("secSensorEnabled")]
        public Input<bool>? SecSensorEnabled { get; set; }

        /// <summary>
        /// The ID of the security policy. System security policies and custom security policies are supported. Valid values: `tls_cipher_policy_1_0` (default), `tls_cipher_policy_1_1,` `tls_cipher_policy_1_2`, `tls_cipher_policy_1_2_strict`, and `tls_cipher_policy_1_2_strict_with_1_3`.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        /// <summary>
        /// The ID of the server group.
        /// </summary>
        [Input("serverGroupId")]
        public Input<string>? ServerGroupId { get; set; }

        /// <summary>
        /// Full Port listens to the starting port. Valid values: `0` ~ `65535`.
        /// </summary>
        [Input("startPort")]
        public Input<int>? StartPort { get; set; }

        /// <summary>
        /// The status of the resource. Valid values: `Running`, `Stopped`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ListenerState()
        {
        }
        public static new ListenerState Empty => new ListenerState();
    }
}
