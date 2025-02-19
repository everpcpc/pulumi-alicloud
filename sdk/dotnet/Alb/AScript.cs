// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb
{
    /// <summary>
    /// Provides a Alb Ascript resource.
    /// 
    /// For information about Alb Ascript and how to use it, see [What is AScript](https://www.alibabacloud.com/help/en/server-load-balancer/latest/what-is-application-load-balancer).
    /// 
    /// &gt; **NOTE:** Available in v1.195.0+.
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
    ///     var @default = new AliCloud.Alb.AScript("default", new()
    ///     {
    ///         ScriptContent = "time()",
    ///         Position = "RequestHead",
    ///         AscriptName = "test",
    ///         Enabled = true,
    ///         ListenerId = @var.ListenerId,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Alb AScript can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:alb/aScript:AScript example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:alb/aScript:AScript")]
    public partial class AScript : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of AScript.
        /// </summary>
        [Output("ascriptName")]
        public Output<string> AscriptName { get; private set; } = null!;

        /// <summary>
        /// Whether scripts are enabled.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// Whether extension parameters are enabled.
        /// </summary>
        [Output("extAttributeEnabled")]
        public Output<bool> ExtAttributeEnabled { get; private set; } = null!;

        /// <summary>
        /// Extended attribute list. See the following `Block ExtAttributes`.
        /// </summary>
        [Output("extAttributes")]
        public Output<ImmutableArray<Outputs.AScriptExtAttribute>> ExtAttributes { get; private set; } = null!;

        /// <summary>
        /// Listener ID of script attribution
        /// </summary>
        [Output("listenerId")]
        public Output<string> ListenerId { get; private set; } = null!;

        /// <summary>
        /// The ID of load balancer instance.
        /// </summary>
        [Output("loadBalancerId")]
        public Output<string> LoadBalancerId { get; private set; } = null!;

        /// <summary>
        /// Execution location of AScript.
        /// </summary>
        [Output("position")]
        public Output<string> Position { get; private set; } = null!;

        /// <summary>
        /// The content of AScript.
        /// </summary>
        [Output("scriptContent")]
        public Output<string> ScriptContent { get; private set; } = null!;

        /// <summary>
        /// The status of AScript.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a AScript resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AScript(string name, AScriptArgs args, CustomResourceOptions? options = null)
            : base("alicloud:alb/aScript:AScript", name, args ?? new AScriptArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AScript(string name, Input<string> id, AScriptState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:alb/aScript:AScript", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AScript resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AScript Get(string name, Input<string> id, AScriptState? state = null, CustomResourceOptions? options = null)
        {
            return new AScript(name, id, state, options);
        }
    }

    public sealed class AScriptArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of AScript.
        /// </summary>
        [Input("ascriptName", required: true)]
        public Input<string> AscriptName { get; set; } = null!;

        /// <summary>
        /// Whether scripts are enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Whether extension parameters are enabled.
        /// </summary>
        [Input("extAttributeEnabled")]
        public Input<bool>? ExtAttributeEnabled { get; set; }

        [Input("extAttributes")]
        private InputList<Inputs.AScriptExtAttributeArgs>? _extAttributes;

        /// <summary>
        /// Extended attribute list. See the following `Block ExtAttributes`.
        /// </summary>
        public InputList<Inputs.AScriptExtAttributeArgs> ExtAttributes
        {
            get => _extAttributes ?? (_extAttributes = new InputList<Inputs.AScriptExtAttributeArgs>());
            set => _extAttributes = value;
        }

        /// <summary>
        /// Listener ID of script attribution
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        /// <summary>
        /// Execution location of AScript.
        /// </summary>
        [Input("position", required: true)]
        public Input<string> Position { get; set; } = null!;

        /// <summary>
        /// The content of AScript.
        /// </summary>
        [Input("scriptContent", required: true)]
        public Input<string> ScriptContent { get; set; } = null!;

        public AScriptArgs()
        {
        }
        public static new AScriptArgs Empty => new AScriptArgs();
    }

    public sealed class AScriptState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of AScript.
        /// </summary>
        [Input("ascriptName")]
        public Input<string>? AscriptName { get; set; }

        /// <summary>
        /// Whether scripts are enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Whether extension parameters are enabled.
        /// </summary>
        [Input("extAttributeEnabled")]
        public Input<bool>? ExtAttributeEnabled { get; set; }

        [Input("extAttributes")]
        private InputList<Inputs.AScriptExtAttributeGetArgs>? _extAttributes;

        /// <summary>
        /// Extended attribute list. See the following `Block ExtAttributes`.
        /// </summary>
        public InputList<Inputs.AScriptExtAttributeGetArgs> ExtAttributes
        {
            get => _extAttributes ?? (_extAttributes = new InputList<Inputs.AScriptExtAttributeGetArgs>());
            set => _extAttributes = value;
        }

        /// <summary>
        /// Listener ID of script attribution
        /// </summary>
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        /// <summary>
        /// The ID of load balancer instance.
        /// </summary>
        [Input("loadBalancerId")]
        public Input<string>? LoadBalancerId { get; set; }

        /// <summary>
        /// Execution location of AScript.
        /// </summary>
        [Input("position")]
        public Input<string>? Position { get; set; }

        /// <summary>
        /// The content of AScript.
        /// </summary>
        [Input("scriptContent")]
        public Input<string>? ScriptContent { get; set; }

        /// <summary>
        /// The status of AScript.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public AScriptState()
        {
        }
        public static new AScriptState Empty => new AScriptState();
    }
}
