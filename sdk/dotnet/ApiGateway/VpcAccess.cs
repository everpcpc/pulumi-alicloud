// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ApiGateway
{
    /// <summary>
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
    ///         var foo = new AliCloud.ApiGateway.VpcAccess("foo", new AliCloud.ApiGateway.VpcAccessArgs
    ///         {
    ///             InstanceId = "i-kai2ks92kzkw92ka",
    ///             Port = 8080,
    ///             VpcId = "vpc-awkcj192ka9zalz",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Api gateway app can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:apigateway/vpcAccess:VpcAccess example "APiGatewayVpc:vpc-aswcj19ajsz:i-ajdjfsdlf:8080"
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:apigateway/vpcAccess:VpcAccess")]
    public partial class VpcAccess : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the instance in VPC (ECS/Server Load Balance).
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The name of the vpc authorization.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the port corresponding to the instance.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// The vpc id of the vpc authorization.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcAccess resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcAccess(string name, VpcAccessArgs args, CustomResourceOptions? options = null)
            : base("alicloud:apigateway/vpcAccess:VpcAccess", name, args ?? new VpcAccessArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcAccess(string name, Input<string> id, VpcAccessState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:apigateway/vpcAccess:VpcAccess", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcAccess resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcAccess Get(string name, Input<string> id, VpcAccessState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcAccess(name, id, state, options);
        }
    }

    public sealed class VpcAccessArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the instance in VPC (ECS/Server Load Balance).
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The name of the vpc authorization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the port corresponding to the instance.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The vpc id of the vpc authorization.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public VpcAccessArgs()
        {
        }
    }

    public sealed class VpcAccessState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the instance in VPC (ECS/Server Load Balance).
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The name of the vpc authorization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the port corresponding to the instance.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The vpc id of the vpc authorization.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public VpcAccessState()
        {
        }
    }
}
