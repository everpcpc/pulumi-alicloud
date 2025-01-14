// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cen
{
    /// <summary>
    /// Provides a CEN child instance grant resource, which allow you to authorize a VPC or VBR to a CEN of a different account.
    /// 
    /// For more information about how to use it, see [Attach a network in a different account](https://www.alibabacloud.com/help/doc-detail/73645.htm).
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
    ///     // Create a new instance-grant and use it to grant one child instance of account1 to a new CEN of account 2.
    ///     var account1 = new AliCloud.Provider("account1", new()
    ///     {
    ///         AccessKey = "access123",
    ///         SecretKey = "secret123",
    ///     });
    /// 
    ///     var account2 = new AliCloud.Provider("account2", new()
    ///     {
    ///         AccessKey = "access456",
    ///         SecretKey = "secret456",
    ///     });
    /// 
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf-testAccCenInstanceGrantBasic";
    ///     var cen = new AliCloud.Cen.Instance("cen", new()
    ///     {
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Account2,
    ///     });
    /// 
    ///     var vpc = new AliCloud.Vpc.Network("vpc", new()
    ///     {
    ///         CidrBlock = "192.168.0.0/16",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Account1,
    ///     });
    /// 
    ///     var fooInstanceGrant = new AliCloud.Cen.InstanceGrant("fooInstanceGrant", new()
    ///     {
    ///         CenId = cen.Id,
    ///         ChildInstanceId = vpc.Id,
    ///         CenOwnerId = "uid2",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Account1,
    ///     });
    /// 
    ///     var fooInstanceAttachment = new AliCloud.Cen.InstanceAttachment("fooInstanceAttachment", new()
    ///     {
    ///         InstanceId = cen.Id,
    ///         ChildInstanceId = vpc.Id,
    ///         ChildInstanceType = "VPC",
    ///         ChildInstanceRegionId = "cn-qingdao",
    ///         ChildInstanceOwnerId = "uid1",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = alicloud.Account2,
    ///         DependsOn = new[]
    ///         {
    ///             fooInstanceGrant,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CEN instance can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cen/instanceGrant:InstanceGrant example cen-abc123456:vpc-abc123456:uid123456
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cen/instanceGrant:InstanceGrant")]
    public partial class InstanceGrant : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Output("cenId")]
        public Output<string> CenId { get; private set; } = null!;

        /// <summary>
        /// The owner UID of the  CEN which the child instance granted to.
        /// </summary>
        [Output("cenOwnerId")]
        public Output<string> CenOwnerId { get; private set; } = null!;

        /// <summary>
        /// The ID of the child instance to grant.
        /// </summary>
        [Output("childInstanceId")]
        public Output<string> ChildInstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceGrant resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceGrant(string name, InstanceGrantArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cen/instanceGrant:InstanceGrant", name, args ?? new InstanceGrantArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceGrant(string name, Input<string> id, InstanceGrantState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cen/instanceGrant:InstanceGrant", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceGrant resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceGrant Get(string name, Input<string> id, InstanceGrantState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceGrant(name, id, state, options);
        }
    }

    public sealed class InstanceGrantArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Input("cenId", required: true)]
        public Input<string> CenId { get; set; } = null!;

        /// <summary>
        /// The owner UID of the  CEN which the child instance granted to.
        /// </summary>
        [Input("cenOwnerId", required: true)]
        public Input<string> CenOwnerId { get; set; } = null!;

        /// <summary>
        /// The ID of the child instance to grant.
        /// </summary>
        [Input("childInstanceId", required: true)]
        public Input<string> ChildInstanceId { get; set; } = null!;

        public InstanceGrantArgs()
        {
        }
        public static new InstanceGrantArgs Empty => new InstanceGrantArgs();
    }

    public sealed class InstanceGrantState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the CEN.
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        /// <summary>
        /// The owner UID of the  CEN which the child instance granted to.
        /// </summary>
        [Input("cenOwnerId")]
        public Input<string>? CenOwnerId { get; set; }

        /// <summary>
        /// The ID of the child instance to grant.
        /// </summary>
        [Input("childInstanceId")]
        public Input<string>? ChildInstanceId { get; set; }

        public InstanceGrantState()
        {
        }
        public static new InstanceGrantState Empty => new InstanceGrantState();
    }
}
