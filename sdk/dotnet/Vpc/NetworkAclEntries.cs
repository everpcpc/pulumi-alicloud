// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    /// <summary>
    /// Provides a network acl entries resource to create ingress and egress entries.
    /// 
    /// &gt; **NOTE:** Available in 1.45.0+. Currently, the resource are only available in Hongkong(cn-hongkong), India(ap-south-1), and Indonesia(ap-southeast-1) regions.
    /// 
    /// &gt; **NOTE:** It doesn't support concurrency and the order of the ingress and egress entries determines the priority.
    /// 
    /// &gt; **NOTE:** Using this resource need to open a whitelist.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/network_acl_entries.html.markdown.
    /// </summary>
    public partial class NetworkAclEntries : Pulumi.CustomResource
    {
        /// <summary>
        /// List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block Egress.
        /// </summary>
        [Output("egresses")]
        public Output<ImmutableArray<Outputs.NetworkAclEntriesEgresses>> Egresses { get; private set; } = null!;

        /// <summary>
        /// List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block Ingress.
        /// </summary>
        [Output("ingresses")]
        public Output<ImmutableArray<Outputs.NetworkAclEntriesIngresses>> Ingresses { get; private set; } = null!;

        /// <summary>
        /// The id of the network acl, the field can't be changed.
        /// </summary>
        [Output("networkAclId")]
        public Output<string> NetworkAclId { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkAclEntries resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkAclEntries(string name, NetworkAclEntriesArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/networkAclEntries:NetworkAclEntries", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private NetworkAclEntries(string name, Input<string> id, NetworkAclEntriesState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/networkAclEntries:NetworkAclEntries", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkAclEntries resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkAclEntries Get(string name, Input<string> id, NetworkAclEntriesState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkAclEntries(name, id, state, options);
        }
    }

    public sealed class NetworkAclEntriesArgs : Pulumi.ResourceArgs
    {
        [Input("egresses")]
        private InputList<Inputs.NetworkAclEntriesEgressesArgs>? _egresses;

        /// <summary>
        /// List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block Egress.
        /// </summary>
        public InputList<Inputs.NetworkAclEntriesEgressesArgs> Egresses
        {
            get => _egresses ?? (_egresses = new InputList<Inputs.NetworkAclEntriesEgressesArgs>());
            set => _egresses = value;
        }

        [Input("ingresses")]
        private InputList<Inputs.NetworkAclEntriesIngressesArgs>? _ingresses;

        /// <summary>
        /// List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block Ingress.
        /// </summary>
        public InputList<Inputs.NetworkAclEntriesIngressesArgs> Ingresses
        {
            get => _ingresses ?? (_ingresses = new InputList<Inputs.NetworkAclEntriesIngressesArgs>());
            set => _ingresses = value;
        }

        /// <summary>
        /// The id of the network acl, the field can't be changed.
        /// </summary>
        [Input("networkAclId", required: true)]
        public Input<string> NetworkAclId { get; set; } = null!;

        public NetworkAclEntriesArgs()
        {
        }
    }

    public sealed class NetworkAclEntriesState : Pulumi.ResourceArgs
    {
        [Input("egresses")]
        private InputList<Inputs.NetworkAclEntriesEgressesGetArgs>? _egresses;

        /// <summary>
        /// List of the egress entries of the network acl. The order of the egress entries determines the priority. The details see Block Egress.
        /// </summary>
        public InputList<Inputs.NetworkAclEntriesEgressesGetArgs> Egresses
        {
            get => _egresses ?? (_egresses = new InputList<Inputs.NetworkAclEntriesEgressesGetArgs>());
            set => _egresses = value;
        }

        [Input("ingresses")]
        private InputList<Inputs.NetworkAclEntriesIngressesGetArgs>? _ingresses;

        /// <summary>
        /// List of the ingress entries of the network acl. The order of the ingress entries determines the priority. The details see Block Ingress.
        /// </summary>
        public InputList<Inputs.NetworkAclEntriesIngressesGetArgs> Ingresses
        {
            get => _ingresses ?? (_ingresses = new InputList<Inputs.NetworkAclEntriesIngressesGetArgs>());
            set => _ingresses = value;
        }

        /// <summary>
        /// The id of the network acl, the field can't be changed.
        /// </summary>
        [Input("networkAclId")]
        public Input<string>? NetworkAclId { get; set; }

        public NetworkAclEntriesState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class NetworkAclEntriesEgressesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the egress entry.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination ip of the egress entry.
        /// </summary>
        [Input("destinationCidrIp")]
        public Input<string>? DestinationCidrIp { get; set; }

        /// <summary>
        /// The entry type of the egress entry. It must be `custom` or `system`. Default value is `custom`.
        /// </summary>
        [Input("entryType")]
        public Input<string>? EntryType { get; set; }

        /// <summary>
        /// The name of the egress entry.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The policy of the egress entry. It must be `accept` or `drop`.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The port of the egress entry.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The protocol of the egress entry.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public NetworkAclEntriesEgressesArgs()
        {
        }
    }

    public sealed class NetworkAclEntriesEgressesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the egress entry.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination ip of the egress entry.
        /// </summary>
        [Input("destinationCidrIp")]
        public Input<string>? DestinationCidrIp { get; set; }

        /// <summary>
        /// The entry type of the egress entry. It must be `custom` or `system`. Default value is `custom`.
        /// </summary>
        [Input("entryType")]
        public Input<string>? EntryType { get; set; }

        /// <summary>
        /// The name of the egress entry.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The policy of the egress entry. It must be `accept` or `drop`.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The port of the egress entry.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The protocol of the egress entry.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public NetworkAclEntriesEgressesGetArgs()
        {
        }
    }

    public sealed class NetworkAclEntriesIngressesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the egress entry.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The entry type of the egress entry. It must be `custom` or `system`. Default value is `custom`.
        /// </summary>
        [Input("entryType")]
        public Input<string>? EntryType { get; set; }

        /// <summary>
        /// The name of the egress entry.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The policy of the egress entry. It must be `accept` or `drop`.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The port of the egress entry.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The protocol of the egress entry.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The source ip of the ingress entry.
        /// </summary>
        [Input("sourceCidrIp")]
        public Input<string>? SourceCidrIp { get; set; }

        public NetworkAclEntriesIngressesArgs()
        {
        }
    }

    public sealed class NetworkAclEntriesIngressesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the egress entry.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The entry type of the egress entry. It must be `custom` or `system`. Default value is `custom`.
        /// </summary>
        [Input("entryType")]
        public Input<string>? EntryType { get; set; }

        /// <summary>
        /// The name of the egress entry.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The policy of the egress entry. It must be `accept` or `drop`.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The port of the egress entry.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The protocol of the egress entry.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The source ip of the ingress entry.
        /// </summary>
        [Input("sourceCidrIp")]
        public Input<string>? SourceCidrIp { get; set; }

        public NetworkAclEntriesIngressesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class NetworkAclEntriesEgresses
    {
        /// <summary>
        /// The description of the egress entry.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The destination ip of the egress entry.
        /// </summary>
        public readonly string? DestinationCidrIp;
        /// <summary>
        /// The entry type of the egress entry. It must be `custom` or `system`. Default value is `custom`.
        /// </summary>
        public readonly string? EntryType;
        /// <summary>
        /// The name of the egress entry.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The policy of the egress entry. It must be `accept` or `drop`.
        /// </summary>
        public readonly string? Policy;
        /// <summary>
        /// The port of the egress entry.
        /// </summary>
        public readonly string? Port;
        /// <summary>
        /// The protocol of the egress entry.
        /// </summary>
        public readonly string? Protocol;

        [OutputConstructor]
        private NetworkAclEntriesEgresses(
            string? description,
            string? destinationCidrIp,
            string? entryType,
            string? name,
            string? policy,
            string? port,
            string? protocol)
        {
            Description = description;
            DestinationCidrIp = destinationCidrIp;
            EntryType = entryType;
            Name = name;
            Policy = policy;
            Port = port;
            Protocol = protocol;
        }
    }

    [OutputType]
    public sealed class NetworkAclEntriesIngresses
    {
        /// <summary>
        /// The description of the egress entry.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The entry type of the egress entry. It must be `custom` or `system`. Default value is `custom`.
        /// </summary>
        public readonly string? EntryType;
        /// <summary>
        /// The name of the egress entry.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The policy of the egress entry. It must be `accept` or `drop`.
        /// </summary>
        public readonly string? Policy;
        /// <summary>
        /// The port of the egress entry.
        /// </summary>
        public readonly string? Port;
        /// <summary>
        /// The protocol of the egress entry.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// The source ip of the ingress entry.
        /// </summary>
        public readonly string? SourceCidrIp;

        [OutputConstructor]
        private NetworkAclEntriesIngresses(
            string? description,
            string? entryType,
            string? name,
            string? policy,
            string? port,
            string? protocol,
            string? sourceCidrIp)
        {
            Description = description;
            EntryType = entryType;
            Name = name;
            Policy = policy;
            Port = port;
            Protocol = protocol;
            SourceCidrIp = sourceCidrIp;
        }
    }
    }
}
