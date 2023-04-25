// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dcdn
{
    /// <summary>
    /// Provides a Dcdn Kv resource.
    /// 
    /// For information about Dcdn Kv and how to use it, see [What is Kv](https://www.alibabacloud.com/help/en/dynamic-route-for-cdn/latest/putdcdnkv).
    /// 
    /// &gt; **NOTE:** Available in v1.198.0+.
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
    ///     var defaultKvAccount = AliCloud.Dcdn.GetKvAccount.Invoke(new()
    ///     {
    ///         Status = "online",
    ///     });
    /// 
    ///     var defaultKvNamespace = new AliCloud.Dcdn.KvNamespace("defaultKvNamespace", new()
    ///     {
    ///         Description = "wkmtest",
    ///         Namespace = @var.Name,
    ///     });
    /// 
    ///     var defaultKv = new AliCloud.Dcdn.Kv("defaultKv", new()
    ///     {
    ///         Value = "testvalue",
    ///         Key = @var.Name,
    ///         Namespace = defaultKvNamespace.Namespace,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Dcdn Kv can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:dcdn/kv:Kv example &lt;namespace&gt;:&lt;key&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dcdn/kv:Kv")]
    public partial class Kv : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the key to Put, the longest 512, cannot contain spaces.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The name specified when the customer calls PutDcdnKvNamespace
        /// </summary>
        [Output("namespace")]
        public Output<string> Namespace { get; private set; } = null!;

        /// <summary>
        /// The content of key, up to 2M(2*1000*1000)
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a Kv resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Kv(string name, KvArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dcdn/kv:Kv", name, args ?? new KvArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Kv(string name, Input<string> id, KvState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dcdn/kv:Kv", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Kv resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Kv Get(string name, Input<string> id, KvState? state = null, CustomResourceOptions? options = null)
        {
            return new Kv(name, id, state, options);
        }
    }

    public sealed class KvArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the key to Put, the longest 512, cannot contain spaces.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The name specified when the customer calls PutDcdnKvNamespace
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        /// <summary>
        /// The content of key, up to 2M(2*1000*1000)
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public KvArgs()
        {
        }
        public static new KvArgs Empty => new KvArgs();
    }

    public sealed class KvState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the key to Put, the longest 512, cannot contain spaces.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The name specified when the customer calls PutDcdnKvNamespace
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The content of key, up to 2M(2*1000*1000)
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public KvState()
        {
        }
        public static new KvState Empty => new KvState();
    }
}
