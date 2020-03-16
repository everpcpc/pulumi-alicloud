// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpn
{
    public partial class Connection : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the customer gateway.
        /// </summary>
        [Output("customerGatewayId")]
        public Output<string> CustomerGatewayId { get; private set; } = null!;

        /// <summary>
        /// Whether to delete a successfully negotiated IPsec tunnel and initiate a negotiation again. Valid value:true,false.
        /// </summary>
        [Output("effectImmediately")]
        public Output<bool?> EffectImmediately { get; private set; } = null!;

        /// <summary>
        /// The configurations of phase-one negotiation.
        /// </summary>
        [Output("ikeConfigs")]
        public Output<ImmutableArray<Outputs.ConnectionIkeConfigs>> IkeConfigs { get; private set; } = null!;

        /// <summary>
        /// The configurations of phase-two negotiation.
        /// </summary>
        [Output("ipsecConfigs")]
        public Output<ImmutableArray<Outputs.ConnectionIpsecConfigs>> IpsecConfigs { get; private set; } = null!;

        /// <summary>
        /// The CIDR block of the VPC to be connected with the local data center. This parameter is used for phase-two negotiation.
        /// </summary>
        [Output("localSubnets")]
        public Output<ImmutableArray<string>> LocalSubnets { get; private set; } = null!;

        /// <summary>
        /// The name of the IPsec connection.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The CIDR block of the local data center. This parameter is used for phase-two negotiation.
        /// </summary>
        [Output("remoteSubnets")]
        public Output<ImmutableArray<string>> RemoteSubnets { get; private set; } = null!;

        /// <summary>
        /// The status of VPN connection.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPN gateway.
        /// </summary>
        [Output("vpnGatewayId")]
        public Output<string> VpnGatewayId { get; private set; } = null!;


        /// <summary>
        /// Create a Connection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Connection(string name, ConnectionArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpn/connection:Connection", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Connection(string name, Input<string> id, ConnectionState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpn/connection:Connection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Connection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Connection Get(string name, Input<string> id, ConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new Connection(name, id, state, options);
        }
    }

    public sealed class ConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the customer gateway.
        /// </summary>
        [Input("customerGatewayId", required: true)]
        public Input<string> CustomerGatewayId { get; set; } = null!;

        /// <summary>
        /// Whether to delete a successfully negotiated IPsec tunnel and initiate a negotiation again. Valid value:true,false.
        /// </summary>
        [Input("effectImmediately")]
        public Input<bool>? EffectImmediately { get; set; }

        [Input("ikeConfigs")]
        private InputList<Inputs.ConnectionIkeConfigsArgs>? _ikeConfigs;

        /// <summary>
        /// The configurations of phase-one negotiation.
        /// </summary>
        public InputList<Inputs.ConnectionIkeConfigsArgs> IkeConfigs
        {
            get => _ikeConfigs ?? (_ikeConfigs = new InputList<Inputs.ConnectionIkeConfigsArgs>());
            set => _ikeConfigs = value;
        }

        [Input("ipsecConfigs")]
        private InputList<Inputs.ConnectionIpsecConfigsArgs>? _ipsecConfigs;

        /// <summary>
        /// The configurations of phase-two negotiation.
        /// </summary>
        public InputList<Inputs.ConnectionIpsecConfigsArgs> IpsecConfigs
        {
            get => _ipsecConfigs ?? (_ipsecConfigs = new InputList<Inputs.ConnectionIpsecConfigsArgs>());
            set => _ipsecConfigs = value;
        }

        [Input("localSubnets", required: true)]
        private InputList<string>? _localSubnets;

        /// <summary>
        /// The CIDR block of the VPC to be connected with the local data center. This parameter is used for phase-two negotiation.
        /// </summary>
        public InputList<string> LocalSubnets
        {
            get => _localSubnets ?? (_localSubnets = new InputList<string>());
            set => _localSubnets = value;
        }

        /// <summary>
        /// The name of the IPsec connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("remoteSubnets", required: true)]
        private InputList<string>? _remoteSubnets;

        /// <summary>
        /// The CIDR block of the local data center. This parameter is used for phase-two negotiation.
        /// </summary>
        public InputList<string> RemoteSubnets
        {
            get => _remoteSubnets ?? (_remoteSubnets = new InputList<string>());
            set => _remoteSubnets = value;
        }

        /// <summary>
        /// The ID of the VPN gateway.
        /// </summary>
        [Input("vpnGatewayId", required: true)]
        public Input<string> VpnGatewayId { get; set; } = null!;

        public ConnectionArgs()
        {
        }
    }

    public sealed class ConnectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the customer gateway.
        /// </summary>
        [Input("customerGatewayId")]
        public Input<string>? CustomerGatewayId { get; set; }

        /// <summary>
        /// Whether to delete a successfully negotiated IPsec tunnel and initiate a negotiation again. Valid value:true,false.
        /// </summary>
        [Input("effectImmediately")]
        public Input<bool>? EffectImmediately { get; set; }

        [Input("ikeConfigs")]
        private InputList<Inputs.ConnectionIkeConfigsGetArgs>? _ikeConfigs;

        /// <summary>
        /// The configurations of phase-one negotiation.
        /// </summary>
        public InputList<Inputs.ConnectionIkeConfigsGetArgs> IkeConfigs
        {
            get => _ikeConfigs ?? (_ikeConfigs = new InputList<Inputs.ConnectionIkeConfigsGetArgs>());
            set => _ikeConfigs = value;
        }

        [Input("ipsecConfigs")]
        private InputList<Inputs.ConnectionIpsecConfigsGetArgs>? _ipsecConfigs;

        /// <summary>
        /// The configurations of phase-two negotiation.
        /// </summary>
        public InputList<Inputs.ConnectionIpsecConfigsGetArgs> IpsecConfigs
        {
            get => _ipsecConfigs ?? (_ipsecConfigs = new InputList<Inputs.ConnectionIpsecConfigsGetArgs>());
            set => _ipsecConfigs = value;
        }

        [Input("localSubnets")]
        private InputList<string>? _localSubnets;

        /// <summary>
        /// The CIDR block of the VPC to be connected with the local data center. This parameter is used for phase-two negotiation.
        /// </summary>
        public InputList<string> LocalSubnets
        {
            get => _localSubnets ?? (_localSubnets = new InputList<string>());
            set => _localSubnets = value;
        }

        /// <summary>
        /// The name of the IPsec connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("remoteSubnets")]
        private InputList<string>? _remoteSubnets;

        /// <summary>
        /// The CIDR block of the local data center. This parameter is used for phase-two negotiation.
        /// </summary>
        public InputList<string> RemoteSubnets
        {
            get => _remoteSubnets ?? (_remoteSubnets = new InputList<string>());
            set => _remoteSubnets = value;
        }

        /// <summary>
        /// The status of VPN connection.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the VPN gateway.
        /// </summary>
        [Input("vpnGatewayId")]
        public Input<string>? VpnGatewayId { get; set; }

        public ConnectionState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ConnectionIkeConfigsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication algorithm of phase-one negotiation. Valid value: md5 | sha1 | sha256 | sha384 | sha512 |. Default value: sha1
        /// </summary>
        [Input("ikeAuthAlg")]
        public Input<string>? IkeAuthAlg { get; set; }

        /// <summary>
        /// The encryption algorithm of phase-one negotiation. Valid value: aes | aes192 | aes256 | des | 3des. Default Valid value: aes
        /// </summary>
        [Input("ikeEncAlg")]
        public Input<string>? IkeEncAlg { get; set; }

        /// <summary>
        /// The SA lifecycle as the result of phase-one negotiation. The valid value of n is [0, 86400], the unit is second and the default value is 86400.
        /// </summary>
        [Input("ikeLifetime")]
        public Input<int>? IkeLifetime { get; set; }

        /// <summary>
        /// The identification of the VPN gateway.
        /// </summary>
        [Input("ikeLocalId")]
        public Input<string>? IkeLocalId { get; set; }

        /// <summary>
        /// The negotiation mode of IKE V1. Valid value: main (main mode) | aggressive (aggressive mode). Default value: main
        /// </summary>
        [Input("ikeMode")]
        public Input<string>? IkeMode { get; set; }

        /// <summary>
        /// The Diffie-Hellman key exchange algorithm used by phase-one negotiation. Valid value: group1 | group2 | group5 | group14 | group24. Default value: group2
        /// </summary>
        [Input("ikePfs")]
        public Input<string>? IkePfs { get; set; }

        /// <summary>
        /// The identification of the customer gateway.
        /// </summary>
        [Input("ikeRemoteId")]
        public Input<string>? IkeRemoteId { get; set; }

        /// <summary>
        /// The version of the IKE protocol. Valid value: ikev1 | ikev2. Default value: ikev1
        /// </summary>
        [Input("ikeVersion")]
        public Input<string>? IkeVersion { get; set; }

        /// <summary>
        /// Used for authentication between the IPsec VPN gateway and the customer gateway.
        /// </summary>
        [Input("psk")]
        public Input<string>? Psk { get; set; }

        public ConnectionIkeConfigsArgs()
        {
        }
    }

    public sealed class ConnectionIkeConfigsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication algorithm of phase-one negotiation. Valid value: md5 | sha1 | sha256 | sha384 | sha512 |. Default value: sha1
        /// </summary>
        [Input("ikeAuthAlg")]
        public Input<string>? IkeAuthAlg { get; set; }

        /// <summary>
        /// The encryption algorithm of phase-one negotiation. Valid value: aes | aes192 | aes256 | des | 3des. Default Valid value: aes
        /// </summary>
        [Input("ikeEncAlg")]
        public Input<string>? IkeEncAlg { get; set; }

        /// <summary>
        /// The SA lifecycle as the result of phase-one negotiation. The valid value of n is [0, 86400], the unit is second and the default value is 86400.
        /// </summary>
        [Input("ikeLifetime")]
        public Input<int>? IkeLifetime { get; set; }

        /// <summary>
        /// The identification of the VPN gateway.
        /// </summary>
        [Input("ikeLocalId")]
        public Input<string>? IkeLocalId { get; set; }

        /// <summary>
        /// The negotiation mode of IKE V1. Valid value: main (main mode) | aggressive (aggressive mode). Default value: main
        /// </summary>
        [Input("ikeMode")]
        public Input<string>? IkeMode { get; set; }

        /// <summary>
        /// The Diffie-Hellman key exchange algorithm used by phase-one negotiation. Valid value: group1 | group2 | group5 | group14 | group24. Default value: group2
        /// </summary>
        [Input("ikePfs")]
        public Input<string>? IkePfs { get; set; }

        /// <summary>
        /// The identification of the customer gateway.
        /// </summary>
        [Input("ikeRemoteId")]
        public Input<string>? IkeRemoteId { get; set; }

        /// <summary>
        /// The version of the IKE protocol. Valid value: ikev1 | ikev2. Default value: ikev1
        /// </summary>
        [Input("ikeVersion")]
        public Input<string>? IkeVersion { get; set; }

        /// <summary>
        /// Used for authentication between the IPsec VPN gateway and the customer gateway.
        /// </summary>
        [Input("psk")]
        public Input<string>? Psk { get; set; }

        public ConnectionIkeConfigsGetArgs()
        {
        }
    }

    public sealed class ConnectionIpsecConfigsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication algorithm of phase-two negotiation. Valid value: md5 | sha1 | sha256 | sha384 | sha512 |. Default value: sha1
        /// </summary>
        [Input("ipsecAuthAlg")]
        public Input<string>? IpsecAuthAlg { get; set; }

        /// <summary>
        /// The encryption algorithm of phase-two negotiation. Valid value: aes | aes192 | aes256 | des | 3des. Default value: aes
        /// </summary>
        [Input("ipsecEncAlg")]
        public Input<string>? IpsecEncAlg { get; set; }

        /// <summary>
        /// The SA lifecycle as the result of phase-two negotiation. The valid value is [0, 86400], the unit is second and the default value is 86400.
        /// </summary>
        [Input("ipsecLifetime")]
        public Input<int>? IpsecLifetime { get; set; }

        /// <summary>
        /// The Diffie-Hellman key exchange algorithm used by phase-two negotiation. Valid value: group1 | group2 | group5 | group14 | group24| disabled. Default value: group2
        /// </summary>
        [Input("ipsecPfs")]
        public Input<string>? IpsecPfs { get; set; }

        public ConnectionIpsecConfigsArgs()
        {
        }
    }

    public sealed class ConnectionIpsecConfigsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication algorithm of phase-two negotiation. Valid value: md5 | sha1 | sha256 | sha384 | sha512 |. Default value: sha1
        /// </summary>
        [Input("ipsecAuthAlg")]
        public Input<string>? IpsecAuthAlg { get; set; }

        /// <summary>
        /// The encryption algorithm of phase-two negotiation. Valid value: aes | aes192 | aes256 | des | 3des. Default value: aes
        /// </summary>
        [Input("ipsecEncAlg")]
        public Input<string>? IpsecEncAlg { get; set; }

        /// <summary>
        /// The SA lifecycle as the result of phase-two negotiation. The valid value is [0, 86400], the unit is second and the default value is 86400.
        /// </summary>
        [Input("ipsecLifetime")]
        public Input<int>? IpsecLifetime { get; set; }

        /// <summary>
        /// The Diffie-Hellman key exchange algorithm used by phase-two negotiation. Valid value: group1 | group2 | group5 | group14 | group24| disabled. Default value: group2
        /// </summary>
        [Input("ipsecPfs")]
        public Input<string>? IpsecPfs { get; set; }

        public ConnectionIpsecConfigsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ConnectionIkeConfigs
    {
        /// <summary>
        /// The authentication algorithm of phase-one negotiation. Valid value: md5 | sha1 | sha256 | sha384 | sha512 |. Default value: sha1
        /// </summary>
        public readonly string? IkeAuthAlg;
        /// <summary>
        /// The encryption algorithm of phase-one negotiation. Valid value: aes | aes192 | aes256 | des | 3des. Default Valid value: aes
        /// </summary>
        public readonly string? IkeEncAlg;
        /// <summary>
        /// The SA lifecycle as the result of phase-one negotiation. The valid value of n is [0, 86400], the unit is second and the default value is 86400.
        /// </summary>
        public readonly int? IkeLifetime;
        /// <summary>
        /// The identification of the VPN gateway.
        /// </summary>
        public readonly string? IkeLocalId;
        /// <summary>
        /// The negotiation mode of IKE V1. Valid value: main (main mode) | aggressive (aggressive mode). Default value: main
        /// </summary>
        public readonly string? IkeMode;
        /// <summary>
        /// The Diffie-Hellman key exchange algorithm used by phase-one negotiation. Valid value: group1 | group2 | group5 | group14 | group24. Default value: group2
        /// </summary>
        public readonly string? IkePfs;
        /// <summary>
        /// The identification of the customer gateway.
        /// </summary>
        public readonly string? IkeRemoteId;
        /// <summary>
        /// The version of the IKE protocol. Valid value: ikev1 | ikev2. Default value: ikev1
        /// </summary>
        public readonly string? IkeVersion;
        /// <summary>
        /// Used for authentication between the IPsec VPN gateway and the customer gateway.
        /// </summary>
        public readonly string? Psk;

        [OutputConstructor]
        private ConnectionIkeConfigs(
            string? ikeAuthAlg,
            string? ikeEncAlg,
            int? ikeLifetime,
            string? ikeLocalId,
            string? ikeMode,
            string? ikePfs,
            string? ikeRemoteId,
            string? ikeVersion,
            string? psk)
        {
            IkeAuthAlg = ikeAuthAlg;
            IkeEncAlg = ikeEncAlg;
            IkeLifetime = ikeLifetime;
            IkeLocalId = ikeLocalId;
            IkeMode = ikeMode;
            IkePfs = ikePfs;
            IkeRemoteId = ikeRemoteId;
            IkeVersion = ikeVersion;
            Psk = psk;
        }
    }

    [OutputType]
    public sealed class ConnectionIpsecConfigs
    {
        /// <summary>
        /// The authentication algorithm of phase-two negotiation. Valid value: md5 | sha1 | sha256 | sha384 | sha512 |. Default value: sha1
        /// </summary>
        public readonly string? IpsecAuthAlg;
        /// <summary>
        /// The encryption algorithm of phase-two negotiation. Valid value: aes | aes192 | aes256 | des | 3des. Default value: aes
        /// </summary>
        public readonly string? IpsecEncAlg;
        /// <summary>
        /// The SA lifecycle as the result of phase-two negotiation. The valid value is [0, 86400], the unit is second and the default value is 86400.
        /// </summary>
        public readonly int? IpsecLifetime;
        /// <summary>
        /// The Diffie-Hellman key exchange algorithm used by phase-two negotiation. Valid value: group1 | group2 | group5 | group14 | group24| disabled. Default value: group2
        /// </summary>
        public readonly string? IpsecPfs;

        [OutputConstructor]
        private ConnectionIpsecConfigs(
            string? ipsecAuthAlg,
            string? ipsecEncAlg,
            int? ipsecLifetime,
            string? ipsecPfs)
        {
            IpsecAuthAlg = ipsecAuthAlg;
            IpsecEncAlg = ipsecEncAlg;
            IpsecLifetime = ipsecLifetime;
            IpsecPfs = ipsecPfs;
        }
    }
    }
}
