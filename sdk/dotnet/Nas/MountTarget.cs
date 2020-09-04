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
    /// Provides a NAS Mount Target resource.
    /// For information about NAS Mount Target and how to use it, see [Manage NAS Mount Targets](https://www.alibabacloud.com/help/en/doc-detail/27531.htm).
    /// 
    /// &gt; NOTE: Available in v1.34.0+.
    /// 
    /// &gt; NOTE: Currently this resource support create a mount point in a classic network only when current region is China mainland regions.
    /// 
    /// &gt; NOTE: You must grant NAS with specific RAM permissions when creating a classic mount targets,
    /// and it only can be achieved by creating a classic mount target mannually.
    /// See [Add a mount point](https://www.alibabacloud.com/help/doc-detail/60431.htm) and [Why do I need RAM permissions to create a mount point in a classic network](https://www.alibabacloud.com/help/faq-detail/42176.htm).
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
    ///         var exampleFileSystem = new AliCloud.Nas.FileSystem("exampleFileSystem", new AliCloud.Nas.FileSystemArgs
    ///         {
    ///             ProtocolType = "NFS",
    ///             StorageType = "Performance",
    ///             Description = "test file system",
    ///         });
    ///         var exampleAccessGroup = new AliCloud.Nas.AccessGroup("exampleAccessGroup", new AliCloud.Nas.AccessGroupArgs
    ///         {
    ///             AccessGroupName = "test_name",
    ///             AccessGroupType = "Classic",
    ///             Description = "test access group",
    ///         });
    ///         var exampleMountTarget = new AliCloud.Nas.MountTarget("exampleMountTarget", new AliCloud.Nas.MountTargetArgs
    ///         {
    ///             FileSystemId = exampleFileSystem.Id,
    ///             AccessGroupName = exampleAccessGroup.AccessGroupName,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class MountTarget : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the permission group that applies to the mount target.
        /// </summary>
        [Output("accessGroupName")]
        public Output<string> AccessGroupName { get; private set; } = null!;

        /// <summary>
        /// The ID of the file system.
        /// </summary>
        [Output("fileSystemId")]
        public Output<string> FileSystemId { get; private set; } = null!;

        /// <summary>
        /// The ID of security group.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string?> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// Whether the MountTarget is active. The status of the mount target. Valid values: `Active` and `Inactive`, Default value is `Active`. Before you mount a file system, make sure that the mount target is in the Active state.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of the VSwitch in the VPC where the mount target resides.
        /// </summary>
        [Output("vswitchId")]
        public Output<string?> VswitchId { get; private set; } = null!;


        /// <summary>
        /// Create a MountTarget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MountTarget(string name, MountTargetArgs args, CustomResourceOptions? options = null)
            : base("alicloud:nas/mountTarget:MountTarget", name, args ?? new MountTargetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MountTarget(string name, Input<string> id, MountTargetState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:nas/mountTarget:MountTarget", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MountTarget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MountTarget Get(string name, Input<string> id, MountTargetState? state = null, CustomResourceOptions? options = null)
        {
            return new MountTarget(name, id, state, options);
        }
    }

    public sealed class MountTargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the permission group that applies to the mount target.
        /// </summary>
        [Input("accessGroupName", required: true)]
        public Input<string> AccessGroupName { get; set; } = null!;

        /// <summary>
        /// The ID of the file system.
        /// </summary>
        [Input("fileSystemId", required: true)]
        public Input<string> FileSystemId { get; set; } = null!;

        /// <summary>
        /// The ID of security group.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// Whether the MountTarget is active. The status of the mount target. Valid values: `Active` and `Inactive`, Default value is `Active`. Before you mount a file system, make sure that the mount target is in the Active state.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the VSwitch in the VPC where the mount target resides.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public MountTargetArgs()
        {
        }
    }

    public sealed class MountTargetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the permission group that applies to the mount target.
        /// </summary>
        [Input("accessGroupName")]
        public Input<string>? AccessGroupName { get; set; }

        /// <summary>
        /// The ID of the file system.
        /// </summary>
        [Input("fileSystemId")]
        public Input<string>? FileSystemId { get; set; }

        /// <summary>
        /// The ID of security group.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// Whether the MountTarget is active. The status of the mount target. Valid values: `Active` and `Inactive`, Default value is `Active`. Before you mount a file system, make sure that the mount target is in the Active state.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the VSwitch in the VPC where the mount target resides.
        /// </summary>
        [Input("vswitchId")]
        public Input<string>? VswitchId { get; set; }

        public MountTargetState()
        {
        }
    }
}
