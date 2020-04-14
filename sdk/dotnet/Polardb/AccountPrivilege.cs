// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.PolarDB
{
    /// <summary>
    /// Provides a PolarDB account privilege resource and used to grant several database some access privilege. A database can be granted by multiple account.
    /// 
    /// &gt; **NOTE:** Available in v1.67.0+.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/polardb_account_privilege.html.markdown.
    /// </summary>
    public partial class AccountPrivilege : Pulumi.CustomResource
    {
        /// <summary>
        /// A specified account name.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"]. Default to "ReadOnly".
        /// </summary>
        [Output("accountPrivilege")]
        public Output<string?> Privilege { get; private set; } = null!;

        /// <summary>
        /// The Id of cluster in which account belongs.
        /// </summary>
        [Output("dbClusterId")]
        public Output<string> DbClusterId { get; private set; } = null!;

        /// <summary>
        /// List of specified database name.
        /// </summary>
        [Output("dbNames")]
        public Output<ImmutableArray<string>> DbNames { get; private set; } = null!;


        /// <summary>
        /// Create a AccountPrivilege resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountPrivilege(string name, AccountPrivilegeArgs args, CustomResourceOptions? options = null)
            : base("alicloud:polardb/accountPrivilege:AccountPrivilege", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AccountPrivilege(string name, Input<string> id, AccountPrivilegeState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:polardb/accountPrivilege:AccountPrivilege", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccountPrivilege resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountPrivilege Get(string name, Input<string> id, AccountPrivilegeState? state = null, CustomResourceOptions? options = null)
        {
            return new AccountPrivilege(name, id, state, options);
        }
    }

    public sealed class AccountPrivilegeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A specified account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"]. Default to "ReadOnly".
        /// </summary>
        [Input("accountPrivilege")]
        public Input<string>? Privilege { get; set; }

        /// <summary>
        /// The Id of cluster in which account belongs.
        /// </summary>
        [Input("dbClusterId", required: true)]
        public Input<string> DbClusterId { get; set; } = null!;

        [Input("dbNames", required: true)]
        private InputList<string>? _dbNames;

        /// <summary>
        /// List of specified database name.
        /// </summary>
        public InputList<string> DbNames
        {
            get => _dbNames ?? (_dbNames = new InputList<string>());
            set => _dbNames = value;
        }

        public AccountPrivilegeArgs()
        {
        }
    }

    public sealed class AccountPrivilegeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A specified account name.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// The privilege of one account access database. Valid values: ["ReadOnly", "ReadWrite"]. Default to "ReadOnly".
        /// </summary>
        [Input("accountPrivilege")]
        public Input<string>? Privilege { get; set; }

        /// <summary>
        /// The Id of cluster in which account belongs.
        /// </summary>
        [Input("dbClusterId")]
        public Input<string>? DbClusterId { get; set; }

        [Input("dbNames")]
        private InputList<string>? _dbNames;

        /// <summary>
        /// List of specified database name.
        /// </summary>
        public InputList<string> DbNames
        {
            get => _dbNames ?? (_dbNames = new InputList<string>());
            set => _dbNames = value;
        }

        public AccountPrivilegeState()
        {
        }
    }
}
