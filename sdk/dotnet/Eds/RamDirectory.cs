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
    /// Provides a ECD Ram Directory resource.
    /// 
    /// For information about ECD Ram Directory and how to use it, see [What is Ram Directory](https://help.aliyun.com/document_detail/436216.html).
    /// 
    /// &gt; **NOTE:** Available in v1.174.0+.
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
    ///     var defaultZones = AliCloud.Eds.GetZones.Invoke();
    /// 
    ///     var defaultNetworks = AliCloud.Vpc.GetNetworks.Invoke(new()
    ///     {
    ///         NameRegex = "default-NODELETING",
    ///     });
    /// 
    ///     var defaultSwitches = AliCloud.Vpc.GetSwitches.Invoke(new()
    ///     {
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Ids[0]),
    ///     });
    /// 
    ///     var defaultRamDirectory = new AliCloud.Eds.RamDirectory("defaultRamDirectory", new()
    ///     {
    ///         DesktopAccessType = "INTERNET",
    ///         EnableAdminAccess = true,
    ///         EnableInternetAccess = true,
    ///         RamDirectoryName = @var.Name,
    ///         VswitchIds = new[]
    ///         {
    ///             defaultSwitches.Apply(getSwitchesResult =&gt; getSwitchesResult.Ids[0]),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECD Ram Directory can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:eds/ramDirectory:RamDirectory example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:eds/ramDirectory:RamDirectory")]
    public partial class RamDirectory : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
        /// </summary>
        [Output("desktopAccessType")]
        public Output<string> DesktopAccessType { get; private set; } = null!;

        /// <summary>
        /// Whether to enable public network access.
        /// </summary>
        [Output("enableAdminAccess")]
        public Output<bool> EnableAdminAccess { get; private set; } = null!;

        /// <summary>
        /// Whether to grant local administrator rights to users who use cloud desktops.
        /// </summary>
        [Output("enableInternetAccess")]
        public Output<bool> EnableInternetAccess { get; private set; } = null!;

        /// <summary>
        /// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
        /// </summary>
        [Output("ramDirectoryName")]
        public Output<string> RamDirectoryName { get; private set; } = null!;

        /// <summary>
        /// The status of directory.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// List of VSwitch IDs in the directory.
        /// </summary>
        [Output("vswitchIds")]
        public Output<ImmutableArray<string>> VswitchIds { get; private set; } = null!;


        /// <summary>
        /// Create a RamDirectory resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RamDirectory(string name, RamDirectoryArgs args, CustomResourceOptions? options = null)
            : base("alicloud:eds/ramDirectory:RamDirectory", name, args ?? new RamDirectoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RamDirectory(string name, Input<string> id, RamDirectoryState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:eds/ramDirectory:RamDirectory", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RamDirectory resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RamDirectory Get(string name, Input<string> id, RamDirectoryState? state = null, CustomResourceOptions? options = null)
        {
            return new RamDirectory(name, id, state, options);
        }
    }

    public sealed class RamDirectoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
        /// </summary>
        [Input("desktopAccessType")]
        public Input<string>? DesktopAccessType { get; set; }

        /// <summary>
        /// Whether to enable public network access.
        /// </summary>
        [Input("enableAdminAccess")]
        public Input<bool>? EnableAdminAccess { get; set; }

        /// <summary>
        /// Whether to grant local administrator rights to users who use cloud desktops.
        /// </summary>
        [Input("enableInternetAccess")]
        public Input<bool>? EnableInternetAccess { get; set; }

        /// <summary>
        /// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
        /// </summary>
        [Input("ramDirectoryName", required: true)]
        public Input<string> RamDirectoryName { get; set; } = null!;

        [Input("vswitchIds", required: true)]
        private InputList<string>? _vswitchIds;

        /// <summary>
        /// List of VSwitch IDs in the directory.
        /// </summary>
        public InputList<string> VswitchIds
        {
            get => _vswitchIds ?? (_vswitchIds = new InputList<string>());
            set => _vswitchIds = value;
        }

        public RamDirectoryArgs()
        {
        }
        public static new RamDirectoryArgs Empty => new RamDirectoryArgs();
    }

    public sealed class RamDirectoryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The desktop access type. Valid values: `VPC`, `INTERNET`, `ANY`.
        /// </summary>
        [Input("desktopAccessType")]
        public Input<string>? DesktopAccessType { get; set; }

        /// <summary>
        /// Whether to enable public network access.
        /// </summary>
        [Input("enableAdminAccess")]
        public Input<bool>? EnableAdminAccess { get; set; }

        /// <summary>
        /// Whether to grant local administrator rights to users who use cloud desktops.
        /// </summary>
        [Input("enableInternetAccess")]
        public Input<bool>? EnableInternetAccess { get; set; }

        /// <summary>
        /// The name of the directory. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
        /// </summary>
        [Input("ramDirectoryName")]
        public Input<string>? RamDirectoryName { get; set; }

        /// <summary>
        /// The status of directory.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("vswitchIds")]
        private InputList<string>? _vswitchIds;

        /// <summary>
        /// List of VSwitch IDs in the directory.
        /// </summary>
        public InputList<string> VswitchIds
        {
            get => _vswitchIds ?? (_vswitchIds = new InputList<string>());
            set => _vswitchIds = value;
        }

        public RamDirectoryState()
        {
        }
        public static new RamDirectoryState Empty => new RamDirectoryState();
    }
}
