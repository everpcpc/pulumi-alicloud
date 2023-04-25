// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Nas
{
    /// <summary>
    /// Provides a NAS Access Group resource.
    /// 
    /// In NAS, the permission group acts as a whitelist that allows you to restrict file system access. You can allow specified IP addresses or CIDR blocks to access the file system, and assign different levels of access permission to different IP addresses or CIDR blocks by adding rules to the permission group.
    /// For information about NAS Access Group and how to use it, see [What is NAS Access Group](https://www.alibabacloud.com/help/en/doc-detail/27534)
    /// 
    /// &gt; **NOTE:** Available in v1.33.0+.
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
    ///     var foo = new AliCloud.Nas.AccessGroup("foo", new()
    ///     {
    ///         Description = "test_AccessG",
    ///         Type = "Classic",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Example after v1.92.0
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var foo = new AliCloud.Nas.AccessGroup("foo", new()
    ///     {
    ///         AccessGroupName = "CreateAccessGroup",
    ///         AccessGroupType = "Vpc",
    ///         Description = "test_AccessG",
    ///         FileSystemType = "extreme",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// NAS Access Group can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:nas/accessGroup:AccessGroup foo tf_testAccNasConfig:standard
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:nas/accessGroup:AccessGroup")]
    public partial class AccessGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A Name of one Access Group.
        /// </summary>
        [Output("accessGroupName")]
        public Output<string> AccessGroupName { get; private set; } = null!;

        /// <summary>
        /// A Type of one Access Group. Valid values: `Vpc` and `Classic`.
        /// </summary>
        [Output("accessGroupType")]
        public Output<string> AccessGroupType { get; private set; } = null!;

        /// <summary>
        /// The Access Group description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The type of file system. Valid values: `standard` and `extreme`. Default to `standard`. Note that the extreme only support Vpc Network.
        /// </summary>
        [Output("fileSystemType")]
        public Output<string?> FileSystemType { get; private set; } = null!;

        /// <summary>
        /// Replaced by `access_group_name` after version 1.92.0.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Replaced by `access_group_type` after version 1.92.0.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AccessGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessGroup(string name, AccessGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:nas/accessGroup:AccessGroup", name, args ?? new AccessGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessGroup(string name, Input<string> id, AccessGroupState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:nas/accessGroup:AccessGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccessGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessGroup Get(string name, Input<string> id, AccessGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new AccessGroup(name, id, state, options);
        }
    }

    public sealed class AccessGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A Name of one Access Group.
        /// </summary>
        [Input("accessGroupName")]
        public Input<string>? AccessGroupName { get; set; }

        /// <summary>
        /// A Type of one Access Group. Valid values: `Vpc` and `Classic`.
        /// </summary>
        [Input("accessGroupType")]
        public Input<string>? AccessGroupType { get; set; }

        /// <summary>
        /// The Access Group description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The type of file system. Valid values: `standard` and `extreme`. Default to `standard`. Note that the extreme only support Vpc Network.
        /// </summary>
        [Input("fileSystemType")]
        public Input<string>? FileSystemType { get; set; }

        /// <summary>
        /// Replaced by `access_group_name` after version 1.92.0.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Replaced by `access_group_type` after version 1.92.0.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AccessGroupArgs()
        {
        }
        public static new AccessGroupArgs Empty => new AccessGroupArgs();
    }

    public sealed class AccessGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A Name of one Access Group.
        /// </summary>
        [Input("accessGroupName")]
        public Input<string>? AccessGroupName { get; set; }

        /// <summary>
        /// A Type of one Access Group. Valid values: `Vpc` and `Classic`.
        /// </summary>
        [Input("accessGroupType")]
        public Input<string>? AccessGroupType { get; set; }

        /// <summary>
        /// The Access Group description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The type of file system. Valid values: `standard` and `extreme`. Default to `standard`. Note that the extreme only support Vpc Network.
        /// </summary>
        [Input("fileSystemType")]
        public Input<string>? FileSystemType { get; set; }

        /// <summary>
        /// Replaced by `access_group_name` after version 1.92.0.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Replaced by `access_group_type` after version 1.92.0.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AccessGroupState()
        {
        }
        public static new AccessGroupState Empty => new AccessGroupState();
    }
}
