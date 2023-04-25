// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.FNF
{
    /// <summary>
    /// Provides a Serverless Workflow Flow resource.
    /// 
    /// For information about Serverless Workflow Flow and how to use it, see [What is Flow](https://www.alibabacloud.com/help/en/doc-detail/123079.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.105.0+.
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
    ///     var @default = new AliCloud.Ram.Role("default", new()
    ///     {
    ///         Document = @"  {
    ///     ""Statement"": [
    ///       {
    ///         ""Action"": ""sts:AssumeRole"",
    ///         ""Effect"": ""Allow"",
    ///         ""Principal"": {
    ///           ""Service"": [
    ///             ""fnf.aliyuncs.com""
    ///           ]
    ///         }
    ///       }
    ///     ],
    ///     ""Version"": ""1""
    ///   }
    /// ",
    ///     });
    /// 
    ///     var example = new AliCloud.FNF.Flow("example", new()
    ///     {
    ///         Definition = @"  version: v1beta1
    ///   type: flow
    ///   steps:
    ///     - type: pass
    ///       name: helloworld
    /// ",
    ///         RoleArn = @default.Arn,
    ///         Description = "Test for terraform fnf_flow.",
    ///         Type = "FDL",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Serverless Workflow Flow can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:fnf/flow:Flow example &lt;name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:fnf/flow:Flow")]
    public partial class Flow : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The definition of the flow. It must comply with the Flow Definition Language (FDL) syntax.
        /// </summary>
        [Output("definition")]
        public Output<string> Definition { get; private set; } = null!;

        /// <summary>
        /// The description of the flow.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The unique ID of the flow.
        /// </summary>
        [Output("flowId")]
        public Output<string> FlowId { get; private set; } = null!;

        /// <summary>
        /// The time when the flow was last modified.
        /// </summary>
        [Output("lastModifiedTime")]
        public Output<string> LastModifiedTime { get; private set; } = null!;

        /// <summary>
        /// The name of the flow. The name must be unique in an Alibaba Cloud account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ARN of the specified RAM role that Serverless Workflow uses to assume the role when Serverless Workflow executes a flow.
        /// </summary>
        [Output("roleArn")]
        public Output<string?> RoleArn { get; private set; } = null!;

        /// <summary>
        /// The type of the flow. Valid values are `FDL` or `DEFAULT`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Flow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Flow(string name, FlowArgs args, CustomResourceOptions? options = null)
            : base("alicloud:fnf/flow:Flow", name, args ?? new FlowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Flow(string name, Input<string> id, FlowState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:fnf/flow:Flow", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Flow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Flow Get(string name, Input<string> id, FlowState? state = null, CustomResourceOptions? options = null)
        {
            return new Flow(name, id, state, options);
        }
    }

    public sealed class FlowArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The definition of the flow. It must comply with the Flow Definition Language (FDL) syntax.
        /// </summary>
        [Input("definition", required: true)]
        public Input<string> Definition { get; set; } = null!;

        /// <summary>
        /// The description of the flow.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The name of the flow. The name must be unique in an Alibaba Cloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN of the specified RAM role that Serverless Workflow uses to assume the role when Serverless Workflow executes a flow.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// The type of the flow. Valid values are `FDL` or `DEFAULT`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FlowArgs()
        {
        }
        public static new FlowArgs Empty => new FlowArgs();
    }

    public sealed class FlowState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The definition of the flow. It must comply with the Flow Definition Language (FDL) syntax.
        /// </summary>
        [Input("definition")]
        public Input<string>? Definition { get; set; }

        /// <summary>
        /// The description of the flow.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The unique ID of the flow.
        /// </summary>
        [Input("flowId")]
        public Input<string>? FlowId { get; set; }

        /// <summary>
        /// The time when the flow was last modified.
        /// </summary>
        [Input("lastModifiedTime")]
        public Input<string>? LastModifiedTime { get; set; }

        /// <summary>
        /// The name of the flow. The name must be unique in an Alibaba Cloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN of the specified RAM role that Serverless Workflow uses to assume the role when Serverless Workflow executes a flow.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// The type of the flow. Valid values are `FDL` or `DEFAULT`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public FlowState()
        {
        }
        public static new FlowState Empty => new FlowState();
    }
}
