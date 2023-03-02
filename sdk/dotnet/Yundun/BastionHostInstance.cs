// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Yundun
{
    [AliCloudResourceType("alicloud:yundun/bastionHostInstance:BastionHostInstance")]
    public partial class BastionHostInstance : global::Pulumi.CustomResource
    {
        [Output("adAuthServers")]
        public Output<ImmutableArray<Outputs.BastionHostInstanceAdAuthServer>> AdAuthServers { get; private set; } = null!;

        [Output("bandwidth")]
        public Output<string> Bandwidth { get; private set; } = null!;

        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("enablePublicAccess")]
        public Output<bool> EnablePublicAccess { get; private set; } = null!;

        [Output("ldapAuthServers")]
        public Output<ImmutableArray<Outputs.BastionHostInstanceLdapAuthServer>> LdapAuthServers { get; private set; } = null!;

        [Output("licenseCode")]
        public Output<string> LicenseCode { get; private set; } = null!;

        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        [Output("planCode")]
        public Output<string> PlanCode { get; private set; } = null!;

        [Output("publicWhiteLists")]
        public Output<ImmutableArray<string>> PublicWhiteLists { get; private set; } = null!;

        [Output("renewPeriod")]
        public Output<int?> RenewPeriod { get; private set; } = null!;

        [Output("renewalPeriodUnit")]
        public Output<string> RenewalPeriodUnit { get; private set; } = null!;

        [Output("renewalStatus")]
        public Output<string> RenewalStatus { get; private set; } = null!;

        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        [Output("storage")]
        public Output<string> Storage { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        [Output("vswitchId")]
        public Output<string> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a BastionHostInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BastionHostInstance(string name, BastionHostInstanceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:yundun/bastionHostInstance:BastionHostInstance", name, args ?? new BastionHostInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BastionHostInstance(string name, Input<string> id, BastionHostInstanceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:yundun/bastionHostInstance:BastionHostInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BastionHostInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BastionHostInstance Get(string name, Input<string> id, BastionHostInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new BastionHostInstance(name, id, state, options);
        }
    }

    public sealed class BastionHostInstanceArgs : global::Pulumi.ResourceArgs
    {
        [Input("adAuthServers")]
        private InputList<Inputs.BastionHostInstanceAdAuthServerArgs>? _adAuthServers;
        public InputList<Inputs.BastionHostInstanceAdAuthServerArgs> AdAuthServers
        {
            get => _adAuthServers ?? (_adAuthServers = new InputList<Inputs.BastionHostInstanceAdAuthServerArgs>());
            set => _adAuthServers = value;
        }

        [Input("bandwidth", required: true)]
        public Input<string> Bandwidth { get; set; } = null!;

        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("enablePublicAccess")]
        public Input<bool>? EnablePublicAccess { get; set; }

        [Input("ldapAuthServers")]
        private InputList<Inputs.BastionHostInstanceLdapAuthServerArgs>? _ldapAuthServers;
        public InputList<Inputs.BastionHostInstanceLdapAuthServerArgs> LdapAuthServers
        {
            get => _ldapAuthServers ?? (_ldapAuthServers = new InputList<Inputs.BastionHostInstanceLdapAuthServerArgs>());
            set => _ldapAuthServers = value;
        }

        [Input("licenseCode", required: true)]
        public Input<string> LicenseCode { get; set; } = null!;

        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("planCode", required: true)]
        public Input<string> PlanCode { get; set; } = null!;

        [Input("publicWhiteLists")]
        private InputList<string>? _publicWhiteLists;
        public InputList<string> PublicWhiteLists
        {
            get => _publicWhiteLists ?? (_publicWhiteLists = new InputList<string>());
            set => _publicWhiteLists = value;
        }

        [Input("renewPeriod")]
        public Input<int>? RenewPeriod { get; set; }

        [Input("renewalPeriodUnit")]
        public Input<string>? RenewalPeriodUnit { get; set; }

        [Input("renewalStatus")]
        public Input<string>? RenewalStatus { get; set; }

        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("securityGroupIds", required: true)]
        private InputList<string>? _securityGroupIds;
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("storage", required: true)]
        public Input<string> Storage { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vswitchId", required: true)]
        public Input<string> VswitchId { get; set; } = null!;

        public BastionHostInstanceArgs()
        {
        }
        public static new BastionHostInstanceArgs Empty => new BastionHostInstanceArgs();
    }

    public sealed class BastionHostInstanceState : global::Pulumi.ResourceArgs
    {
        [Input("adAuthServers")]
        private InputList<Inputs.BastionHostInstanceAdAuthServerGetArgs>? _adAuthServers;
        public InputList<Inputs.BastionHostInstanceAdAuthServerGetArgs> AdAuthServers
        {
            get => _adAuthServers ?? (_adAuthServers = new InputList<Inputs.BastionHostInstanceAdAuthServerGetArgs>());
            set => _adAuthServers = value;
        }

        [Input("bandwidth")]
        public Input<string>? Bandwidth { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enablePublicAccess")]
        public Input<bool>? EnablePublicAccess { get; set; }

        [Input("ldapAuthServers")]
        private InputList<Inputs.BastionHostInstanceLdapAuthServerGetArgs>? _ldapAuthServers;
        public InputList<Inputs.BastionHostInstanceLdapAuthServerGetArgs> LdapAuthServers
        {
            get => _ldapAuthServers ?? (_ldapAuthServers = new InputList<Inputs.BastionHostInstanceLdapAuthServerGetArgs>());
            set => _ldapAuthServers = value;
        }

        [Input("licenseCode")]
        public Input<string>? LicenseCode { get; set; }

        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("planCode")]
        public Input<string>? PlanCode { get; set; }

        [Input("publicWhiteLists")]
        private InputList<string>? _publicWhiteLists;
        public InputList<string> PublicWhiteLists
        {
            get => _publicWhiteLists ?? (_publicWhiteLists = new InputList<string>());
            set => _publicWhiteLists = value;
        }

        [Input("renewPeriod")]
        public Input<int>? RenewPeriod { get; set; }

        [Input("renewalPeriodUnit")]
        public Input<string>? RenewalPeriodUnit { get; set; }

        [Input("renewalStatus")]
        public Input<string>? RenewalStatus { get; set; }

        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("storage")]
        public Input<string>? Storage { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public BastionHostInstanceState()
        {
        }
        public static new BastionHostInstanceState Empty => new BastionHostInstanceState();
    }
}
