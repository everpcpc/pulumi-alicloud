// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Eds
{
    /// <summary>
    /// Provides a ECD Command resource.
    /// 
    /// For information about ECD Command and how to use it, see [What is Command](https://help.aliyun.com/document_detail/188382.html).
    /// 
    /// &gt; **NOTE:** Available in v1.146.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var defaultSimpleOfficeSite = new AliCloud.Eds.SimpleOfficeSite("defaultSimpleOfficeSite", new AliCloud.Eds.SimpleOfficeSiteArgs
    ///         {
    ///             CidrBlock = "172.16.0.0/12",
    ///             DesktopAccessType = "Internet",
    ///             OfficeSiteName = "your_office_site_name",
    ///         });
    ///         var defaultBundles = Output.Create(AliCloud.Eds.GetBundles.InvokeAsync(new AliCloud.Eds.GetBundlesArgs
    ///         {
    ///             BundleType = "SYSTEM",
    ///             NameRegex = "windows",
    ///         }));
    ///         var defaultEcdPolicyGroup = new AliCloud.Eds.EcdPolicyGroup("defaultEcdPolicyGroup", new AliCloud.Eds.EcdPolicyGroupArgs
    ///         {
    ///             PolicyGroupName = "your_policy_group_name",
    ///             Clipboard = "readwrite",
    ///             LocalDrive = "read",
    ///             AuthorizeAccessPolicyRules = 
    ///             {
    ///                 new AliCloud.Eds.Inputs.EcdPolicyGroupAuthorizeAccessPolicyRuleArgs
    ///                 {
    ///                     Description = "example_value",
    ///                     CidrIp = "1.2.3.4/24",
    ///                 },
    ///             },
    ///             AuthorizeSecurityPolicyRules = 
    ///             {
    ///                 new AliCloud.Eds.Inputs.EcdPolicyGroupAuthorizeSecurityPolicyRuleArgs
    ///                 {
    ///                     Type = "inflow",
    ///                     Policy = "accept",
    ///                     Description = "example_value",
    ///                     PortRange = "80/80",
    ///                     IpProtocol = "TCP",
    ///                     Priority = "1",
    ///                     CidrIp = "0.0.0.0/0",
    ///                 },
    ///             },
    ///         });
    ///         var defaultDesktop = new AliCloud.Eds.Desktop("defaultDesktop", new AliCloud.Eds.DesktopArgs
    ///         {
    ///             OfficeSiteId = defaultSimpleOfficeSite.Id,
    ///             PolicyGroupId = defaultEcdPolicyGroup.Id,
    ///             BundleId = defaultBundles.Apply(defaultBundles =&gt; defaultBundles.Bundles?[0]?.Id),
    ///             DesktopName = @var.Name,
    ///         });
    ///         var defaultCommand = new AliCloud.Eds.Command("defaultCommand", new AliCloud.Eds.CommandArgs
    ///         {
    ///             CommandContent = "ipconfig",
    ///             CommandType = "RunPowerShellScript",
    ///             DesktopId = defaultDesktop.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECD Command can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:eds/command:Command example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:eds/command:Command")]
    public partial class Command : Pulumi.CustomResource
    {
        /// <summary>
        /// The Contents of the Script to Base64 Encoded Transmission.
        /// </summary>
        [Output("commandContent")]
        public Output<string> CommandContent { get; private set; } = null!;

        /// <summary>
        /// The Script Type. Valid values: `RunBatScript`, `RunPowerShellScript`.
        /// </summary>
        [Output("commandType")]
        public Output<string> CommandType { get; private set; } = null!;

        /// <summary>
        /// That Returns the Data Encoding Method. Valid values: `Base64`, `PlainText`.
        /// </summary>
        [Output("contentEncoding")]
        public Output<string?> ContentEncoding { get; private set; } = null!;

        /// <summary>
        /// The desktop id of the Desktop.
        /// </summary>
        [Output("desktopId")]
        public Output<string> DesktopId { get; private set; } = null!;

        /// <summary>
        /// Script Is Executed in the Overall Implementation of the State. Valid values: `Pending`, `Failed`, `PartialFailed`, `Running`, `Stopped`, `Stopping`, `Finished`, `Success`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The timeout period for script execution the unit is seconds. Default to: `60`.
        /// </summary>
        [Output("timeout")]
        public Output<string?> Timeout { get; private set; } = null!;


        /// <summary>
        /// Create a Command resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Command(string name, CommandArgs args, CustomResourceOptions? options = null)
            : base("alicloud:eds/command:Command", name, args ?? new CommandArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Command(string name, Input<string> id, CommandState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:eds/command:Command", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Command resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Command Get(string name, Input<string> id, CommandState? state = null, CustomResourceOptions? options = null)
        {
            return new Command(name, id, state, options);
        }
    }

    public sealed class CommandArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Contents of the Script to Base64 Encoded Transmission.
        /// </summary>
        [Input("commandContent", required: true)]
        public Input<string> CommandContent { get; set; } = null!;

        /// <summary>
        /// The Script Type. Valid values: `RunBatScript`, `RunPowerShellScript`.
        /// </summary>
        [Input("commandType", required: true)]
        public Input<string> CommandType { get; set; } = null!;

        /// <summary>
        /// That Returns the Data Encoding Method. Valid values: `Base64`, `PlainText`.
        /// </summary>
        [Input("contentEncoding")]
        public Input<string>? ContentEncoding { get; set; }

        /// <summary>
        /// The desktop id of the Desktop.
        /// </summary>
        [Input("desktopId", required: true)]
        public Input<string> DesktopId { get; set; } = null!;

        /// <summary>
        /// The timeout period for script execution the unit is seconds. Default to: `60`.
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        public CommandArgs()
        {
        }
    }

    public sealed class CommandState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Contents of the Script to Base64 Encoded Transmission.
        /// </summary>
        [Input("commandContent")]
        public Input<string>? CommandContent { get; set; }

        /// <summary>
        /// The Script Type. Valid values: `RunBatScript`, `RunPowerShellScript`.
        /// </summary>
        [Input("commandType")]
        public Input<string>? CommandType { get; set; }

        /// <summary>
        /// That Returns the Data Encoding Method. Valid values: `Base64`, `PlainText`.
        /// </summary>
        [Input("contentEncoding")]
        public Input<string>? ContentEncoding { get; set; }

        /// <summary>
        /// The desktop id of the Desktop.
        /// </summary>
        [Input("desktopId")]
        public Input<string>? DesktopId { get; set; }

        /// <summary>
        /// Script Is Executed in the Overall Implementation of the State. Valid values: `Pending`, `Failed`, `PartialFailed`, `Running`, `Stopped`, `Stopping`, `Finished`, `Success`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The timeout period for script execution the unit is seconds. Default to: `60`.
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        public CommandState()
        {
        }
    }
}
