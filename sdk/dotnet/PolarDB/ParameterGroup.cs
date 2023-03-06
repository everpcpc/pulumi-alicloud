// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.PolarDB
{
    /// <summary>
    /// Provides a PolarDB Parameter Group resource.
    /// 
    /// For information about PolarDB Parameter Group and how to use it, see [What is Parameter Group](https://www.alibabacloud.com/help/en/polardb-for-mysql/latest/createparametergroup).
    /// 
    /// &gt; **NOTE:** Available in v1.183.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AliCloud.PolarDB.ParameterGroup("example", new()
    ///     {
    ///         DbType = "MySQL",
    ///         DbVersion = "8.0",
    ///         Description = "example_value",
    ///         Parameters = new[]
    ///         {
    ///             new AliCloud.PolarDB.Inputs.ParameterGroupParameterArgs
    ///             {
    ///                 ParamName = "wait_timeout",
    ///                 ParamValue = "86400",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// PolarDB Parameter Group can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:polardb/parameterGroup:ParameterGroup example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:polardb/parameterGroup:ParameterGroup")]
    public partial class ParameterGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The type of the database engine. Only `MySQL` is supported.
        /// </summary>
        [Output("dbType")]
        public Output<string> DbType { get; private set; } = null!;

        /// <summary>
        /// The version number of the database engine. Valid values: `5.6`, `5.7`, `8.0`.
        /// </summary>
        [Output("dbVersion")]
        public Output<string> DbVersion { get; private set; } = null!;

        /// <summary>
        /// The description of the parameter template. It must be 0 to 200 characters in length.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the parameter template. It must be 8 to 64 characters in length, and can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The parameter template. See the following `Block parameters`.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<Outputs.ParameterGroupParameter>> Parameters { get; private set; } = null!;


        /// <summary>
        /// Create a ParameterGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ParameterGroup(string name, ParameterGroupArgs args, CustomResourceOptions? options = null)
            : base("alicloud:polardb/parameterGroup:ParameterGroup", name, args ?? new ParameterGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ParameterGroup(string name, Input<string> id, ParameterGroupState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:polardb/parameterGroup:ParameterGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ParameterGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ParameterGroup Get(string name, Input<string> id, ParameterGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ParameterGroup(name, id, state, options);
        }
    }

    public sealed class ParameterGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the database engine. Only `MySQL` is supported.
        /// </summary>
        [Input("dbType", required: true)]
        public Input<string> DbType { get; set; } = null!;

        /// <summary>
        /// The version number of the database engine. Valid values: `5.6`, `5.7`, `8.0`.
        /// </summary>
        [Input("dbVersion", required: true)]
        public Input<string> DbVersion { get; set; } = null!;

        /// <summary>
        /// The description of the parameter template. It must be 0 to 200 characters in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the parameter template. It must be 8 to 64 characters in length, and can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters", required: true)]
        private InputList<Inputs.ParameterGroupParameterArgs>? _parameters;

        /// <summary>
        /// The parameter template. See the following `Block parameters`.
        /// </summary>
        public InputList<Inputs.ParameterGroupParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ParameterGroupParameterArgs>());
            set => _parameters = value;
        }

        public ParameterGroupArgs()
        {
        }
        public static new ParameterGroupArgs Empty => new ParameterGroupArgs();
    }

    public sealed class ParameterGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the database engine. Only `MySQL` is supported.
        /// </summary>
        [Input("dbType")]
        public Input<string>? DbType { get; set; }

        /// <summary>
        /// The version number of the database engine. Valid values: `5.6`, `5.7`, `8.0`.
        /// </summary>
        [Input("dbVersion")]
        public Input<string>? DbVersion { get; set; }

        /// <summary>
        /// The description of the parameter template. It must be 0 to 200 characters in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the parameter template. It must be 8 to 64 characters in length, and can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputList<Inputs.ParameterGroupParameterGetArgs>? _parameters;

        /// <summary>
        /// The parameter template. See the following `Block parameters`.
        /// </summary>
        public InputList<Inputs.ParameterGroupParameterGetArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ParameterGroupParameterGetArgs>());
            set => _parameters = value;
        }

        public ParameterGroupState()
        {
        }
        public static new ParameterGroupState Empty => new ParameterGroupState();
    }
}