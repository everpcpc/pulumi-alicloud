// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    /// <summary>
    /// Provides a VPC Dhcp Options Set resource.
    /// 
    /// For information about VPC Dhcp Options Set and how to use it, see [What is Dhcp Options Set](https://www.alibabacloud.com/help/doc-detail/174112.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.134.0+.
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
    ///     var example = new AliCloud.Vpc.DhcpOptionsSet("example", new()
    ///     {
    ///         DhcpOptionsSetDescription = "example_value",
    ///         DhcpOptionsSetName = "example_value",
    ///         DomainName = "example.com",
    ///         DomainNameServers = "100.100.2.136",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VPC Dhcp Options Set can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:vpc/dhcpOptionsSet:DhcpOptionsSet example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:vpc/dhcpOptionsSet:DhcpOptionsSet")]
    public partial class DhcpOptionsSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// AssociateVpcs. Number of VPCs that can be associated with each DHCP options set is 10. Field `associate_vpcs` has been deprecated from provider version 1.153.0. It will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
        /// </summary>
        [Output("associateVpcs")]
        public Output<ImmutableArray<Outputs.DhcpOptionsSetAssociateVpc>> AssociateVpcs { get; private set; } = null!;

        /// <summary>
        /// The description of the DHCP options set. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
        /// </summary>
        [Output("dhcpOptionsSetDescription")]
        public Output<string?> DhcpOptionsSetDescription { get; private set; } = null!;

        /// <summary>
        /// The name of the DHCP options set. The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
        /// </summary>
        [Output("dhcpOptionsSetName")]
        public Output<string?> DhcpOptionsSetName { get; private set; } = null!;

        /// <summary>
        /// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
        /// </summary>
        [Output("domainName")]
        public Output<string?> DomainName { get; private set; } = null!;

        /// <summary>
        /// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are `100.100.2.136` and `100.100.2.138`.
        /// </summary>
        [Output("domainNameServers")]
        public Output<string?> DomainNameServers { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to precheck this request only. Valid values: `true` or `false`.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// The ID of the account to which the DHCP options set belongs.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// The status of the DHCP options set. Valid values: `Available`, `InUse` or `Pending`. `Available`: The DHCP options set is available for use. `InUse`: The DHCP options set is in use. `Pending`: The DHCP options set is being configured.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a DhcpOptionsSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DhcpOptionsSet(string name, DhcpOptionsSetArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/dhcpOptionsSet:DhcpOptionsSet", name, args ?? new DhcpOptionsSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DhcpOptionsSet(string name, Input<string> id, DhcpOptionsSetState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/dhcpOptionsSet:DhcpOptionsSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DhcpOptionsSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DhcpOptionsSet Get(string name, Input<string> id, DhcpOptionsSetState? state = null, CustomResourceOptions? options = null)
        {
            return new DhcpOptionsSet(name, id, state, options);
        }
    }

    public sealed class DhcpOptionsSetArgs : global::Pulumi.ResourceArgs
    {
        [Input("associateVpcs")]
        private InputList<Inputs.DhcpOptionsSetAssociateVpcArgs>? _associateVpcs;

        /// <summary>
        /// AssociateVpcs. Number of VPCs that can be associated with each DHCP options set is 10. Field `associate_vpcs` has been deprecated from provider version 1.153.0. It will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
        /// </summary>
        [Obsolete(@"Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.")]
        public InputList<Inputs.DhcpOptionsSetAssociateVpcArgs> AssociateVpcs
        {
            get => _associateVpcs ?? (_associateVpcs = new InputList<Inputs.DhcpOptionsSetAssociateVpcArgs>());
            set => _associateVpcs = value;
        }

        /// <summary>
        /// The description of the DHCP options set. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
        /// </summary>
        [Input("dhcpOptionsSetDescription")]
        public Input<string>? DhcpOptionsSetDescription { get; set; }

        /// <summary>
        /// The name of the DHCP options set. The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
        /// </summary>
        [Input("dhcpOptionsSetName")]
        public Input<string>? DhcpOptionsSetName { get; set; }

        /// <summary>
        /// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are `100.100.2.136` and `100.100.2.138`.
        /// </summary>
        [Input("domainNameServers")]
        public Input<string>? DomainNameServers { get; set; }

        /// <summary>
        /// Specifies whether to precheck this request only. Valid values: `true` or `false`.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        public DhcpOptionsSetArgs()
        {
        }
        public static new DhcpOptionsSetArgs Empty => new DhcpOptionsSetArgs();
    }

    public sealed class DhcpOptionsSetState : global::Pulumi.ResourceArgs
    {
        [Input("associateVpcs")]
        private InputList<Inputs.DhcpOptionsSetAssociateVpcGetArgs>? _associateVpcs;

        /// <summary>
        /// AssociateVpcs. Number of VPCs that can be associated with each DHCP options set is 10. Field `associate_vpcs` has been deprecated from provider version 1.153.0. It will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.
        /// </summary>
        [Obsolete(@"Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.")]
        public InputList<Inputs.DhcpOptionsSetAssociateVpcGetArgs> AssociateVpcs
        {
            get => _associateVpcs ?? (_associateVpcs = new InputList<Inputs.DhcpOptionsSetAssociateVpcGetArgs>());
            set => _associateVpcs = value;
        }

        /// <summary>
        /// The description of the DHCP options set. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
        /// </summary>
        [Input("dhcpOptionsSetDescription")]
        public Input<string>? DhcpOptionsSetDescription { get; set; }

        /// <summary>
        /// The name of the DHCP options set. The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
        /// </summary>
        [Input("dhcpOptionsSetName")]
        public Input<string>? DhcpOptionsSetName { get; set; }

        /// <summary>
        /// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS instances in the VPC network.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP addresses must be separated with commas (,).Before you specify any DNS server IP address, all ECS instances in the associated VPC network use the IP addresses of the Alibaba Cloud DNS servers, which are `100.100.2.136` and `100.100.2.138`.
        /// </summary>
        [Input("domainNameServers")]
        public Input<string>? DomainNameServers { get; set; }

        /// <summary>
        /// Specifies whether to precheck this request only. Valid values: `true` or `false`.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// The ID of the account to which the DHCP options set belongs.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// The status of the DHCP options set. Valid values: `Available`, `InUse` or `Pending`. `Available`: The DHCP options set is available for use. `InUse`: The DHCP options set is in use. `Pending`: The DHCP options set is being configured.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public DhcpOptionsSetState()
        {
        }
        public static new DhcpOptionsSetState Empty => new DhcpOptionsSetState();
    }
}
