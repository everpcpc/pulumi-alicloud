// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Log
{
    /// <summary>
    /// Log resource is a meta store service provided by log service, resource can be used to define meta store's table structure.
    /// 
    /// For information about SLS Resource and how to use it, see [Resource management](https://www.alibabacloud.com/help/en/doc-detail/207732.html)
    /// 
    /// &gt; **NOTE:** Available in 1.162.0+, log resource region should be set a main region: cn-heyuan
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
    ///     var example = new AliCloud.Log.Resource("example", new()
    ///     {
    ///         Description = "user tf test resource desc",
    ///         ExtInfo = "{}",
    ///         Schema = "{\"schema\":[{\"column\":\"col1\",\"desc\":\"col1 desc\",\"ext_info\":{},\"required\":true,\"type\":\"string\"},{\"column\":\"col2\",\"desc\":\"col2 desc\",\"ext_info\":\"optional\",\"required\":true,\"type\":\"string\"}]}",
    ///         Type = "userdefine",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Log resource can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:log/resource:Resource example user.tf.test_resource
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:log/resource:Resource")]
    public partial class Resource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The meta store's description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ext info of meta store.
        /// </summary>
        [Output("extInfo")]
        public Output<string?> ExtInfo { get; private set; } = null!;

        /// <summary>
        /// The meta store's name, can be used as table name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The meta store's schema info, which is json string format, used to define table's fields.
        /// </summary>
        [Output("schema")]
        public Output<string> Schema { get; private set; } = null!;

        /// <summary>
        /// The meta store's type, userdefine e.g.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Resource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Resource(string name, ResourceArgs args, CustomResourceOptions? options = null)
            : base("alicloud:log/resource:Resource", name, args ?? new ResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Resource(string name, Input<string> id, ResourceState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:log/resource:Resource", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Resource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Resource Get(string name, Input<string> id, ResourceState? state = null, CustomResourceOptions? options = null)
        {
            return new Resource(name, id, state, options);
        }
    }

    public sealed class ResourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The meta store's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ext info of meta store.
        /// </summary>
        [Input("extInfo")]
        public Input<string>? ExtInfo { get; set; }

        /// <summary>
        /// The meta store's name, can be used as table name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The meta store's schema info, which is json string format, used to define table's fields.
        /// </summary>
        [Input("schema", required: true)]
        public Input<string> Schema { get; set; } = null!;

        /// <summary>
        /// The meta store's type, userdefine e.g.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ResourceArgs()
        {
        }
        public static new ResourceArgs Empty => new ResourceArgs();
    }

    public sealed class ResourceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The meta store's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ext info of meta store.
        /// </summary>
        [Input("extInfo")]
        public Input<string>? ExtInfo { get; set; }

        /// <summary>
        /// The meta store's name, can be used as table name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The meta store's schema info, which is json string format, used to define table's fields.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// The meta store's type, userdefine e.g.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ResourceState()
        {
        }
        public static new ResourceState Empty => new ResourceState();
    }
}
