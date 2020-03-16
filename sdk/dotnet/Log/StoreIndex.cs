// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Log
{
    public partial class StoreIndex : Pulumi.CustomResource
    {
        /// <summary>
        /// List configurations of field search index. Valid item as follows:
        /// </summary>
        [Output("fieldSearches")]
        public Output<ImmutableArray<Outputs.StoreIndexFieldSearches>> FieldSearches { get; private set; } = null!;

        /// <summary>
        /// The configuration of full text index. Valid item as follows:
        /// </summary>
        [Output("fullText")]
        public Output<Outputs.StoreIndexFullText?> FullText { get; private set; } = null!;

        /// <summary>
        /// The log store name to the query index belongs.
        /// </summary>
        [Output("logstore")]
        public Output<string> Logstore { get; private set; } = null!;

        /// <summary>
        /// The project name to the log store belongs.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a StoreIndex resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StoreIndex(string name, StoreIndexArgs args, CustomResourceOptions? options = null)
            : base("alicloud:log/storeIndex:StoreIndex", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private StoreIndex(string name, Input<string> id, StoreIndexState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:log/storeIndex:StoreIndex", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StoreIndex resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StoreIndex Get(string name, Input<string> id, StoreIndexState? state = null, CustomResourceOptions? options = null)
        {
            return new StoreIndex(name, id, state, options);
        }
    }

    public sealed class StoreIndexArgs : Pulumi.ResourceArgs
    {
        [Input("fieldSearches")]
        private InputList<Inputs.StoreIndexFieldSearchesArgs>? _fieldSearches;

        /// <summary>
        /// List configurations of field search index. Valid item as follows:
        /// </summary>
        public InputList<Inputs.StoreIndexFieldSearchesArgs> FieldSearches
        {
            get => _fieldSearches ?? (_fieldSearches = new InputList<Inputs.StoreIndexFieldSearchesArgs>());
            set => _fieldSearches = value;
        }

        /// <summary>
        /// The configuration of full text index. Valid item as follows:
        /// </summary>
        [Input("fullText")]
        public Input<Inputs.StoreIndexFullTextArgs>? FullText { get; set; }

        /// <summary>
        /// The log store name to the query index belongs.
        /// </summary>
        [Input("logstore", required: true)]
        public Input<string> Logstore { get; set; } = null!;

        /// <summary>
        /// The project name to the log store belongs.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public StoreIndexArgs()
        {
        }
    }

    public sealed class StoreIndexState : Pulumi.ResourceArgs
    {
        [Input("fieldSearches")]
        private InputList<Inputs.StoreIndexFieldSearchesGetArgs>? _fieldSearches;

        /// <summary>
        /// List configurations of field search index. Valid item as follows:
        /// </summary>
        public InputList<Inputs.StoreIndexFieldSearchesGetArgs> FieldSearches
        {
            get => _fieldSearches ?? (_fieldSearches = new InputList<Inputs.StoreIndexFieldSearchesGetArgs>());
            set => _fieldSearches = value;
        }

        /// <summary>
        /// The configuration of full text index. Valid item as follows:
        /// </summary>
        [Input("fullText")]
        public Input<Inputs.StoreIndexFullTextGetArgs>? FullText { get; set; }

        /// <summary>
        /// The log store name to the query index belongs.
        /// </summary>
        [Input("logstore")]
        public Input<string>? Logstore { get; set; }

        /// <summary>
        /// The project name to the log store belongs.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public StoreIndexState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class StoreIndexFieldSearchesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The alias of one field.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Whether the case sensitive for the field. Default to false. It is valid when "type" is "text" or "json".
        /// </summary>
        [Input("caseSensitive")]
        public Input<bool>? CaseSensitive { get; set; }

        /// <summary>
        /// Whether to enable field analytics. Default to true.
        /// </summary>
        [Input("enableAnalytics")]
        public Input<bool>? EnableAnalytics { get; set; }

        /// <summary>
        /// Whether includes the chinese for the field. Default to false. It is valid when "type" is "text" or "json".
        /// </summary>
        [Input("includeChinese")]
        public Input<bool>? IncludeChinese { get; set; }

        [Input("jsonKeys")]
        private InputList<StoreIndexFieldSearchesJsonKeysArgs>? _jsonKeys;

        /// <summary>
        /// Use nested index when type is json
        /// </summary>
        public InputList<StoreIndexFieldSearchesJsonKeysArgs> JsonKeys
        {
            get => _jsonKeys ?? (_jsonKeys = new InputList<StoreIndexFieldSearchesJsonKeysArgs>());
            set => _jsonKeys = value;
        }

        /// <summary>
        /// When using the json_keys field, this field is required.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The string of several split words, like "\r", "#". It is valid when "type" is "text" or "json".
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// The type of one field. Valid values: ["long", "text", "double"]. Default to "long"
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public StoreIndexFieldSearchesArgs()
        {
        }
    }

    public sealed class StoreIndexFieldSearchesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The alias of one field.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Whether the case sensitive for the field. Default to false. It is valid when "type" is "text" or "json".
        /// </summary>
        [Input("caseSensitive")]
        public Input<bool>? CaseSensitive { get; set; }

        /// <summary>
        /// Whether to enable field analytics. Default to true.
        /// </summary>
        [Input("enableAnalytics")]
        public Input<bool>? EnableAnalytics { get; set; }

        /// <summary>
        /// Whether includes the chinese for the field. Default to false. It is valid when "type" is "text" or "json".
        /// </summary>
        [Input("includeChinese")]
        public Input<bool>? IncludeChinese { get; set; }

        [Input("jsonKeys")]
        private InputList<StoreIndexFieldSearchesJsonKeysGetArgs>? _jsonKeys;

        /// <summary>
        /// Use nested index when type is json
        /// </summary>
        public InputList<StoreIndexFieldSearchesJsonKeysGetArgs> JsonKeys
        {
            get => _jsonKeys ?? (_jsonKeys = new InputList<StoreIndexFieldSearchesJsonKeysGetArgs>());
            set => _jsonKeys = value;
        }

        /// <summary>
        /// When using the json_keys field, this field is required.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The string of several split words, like "\r", "#". It is valid when "type" is "text" or "json".
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// The type of one field. Valid values: ["long", "text", "double"]. Default to "long"
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public StoreIndexFieldSearchesGetArgs()
        {
        }
    }

    public sealed class StoreIndexFieldSearchesJsonKeysArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The alias of one field.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Whether to enable statistics. default to true.
        /// </summary>
        [Input("docValue")]
        public Input<bool>? DocValue { get; set; }

        /// <summary>
        /// When using the json_keys field, this field is required.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The type of one field. Valid values: ["long", "text", "double"]. Default to "long"
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public StoreIndexFieldSearchesJsonKeysArgs()
        {
        }
    }

    public sealed class StoreIndexFieldSearchesJsonKeysGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The alias of one field.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Whether to enable statistics. default to true.
        /// </summary>
        [Input("docValue")]
        public Input<bool>? DocValue { get; set; }

        /// <summary>
        /// When using the json_keys field, this field is required.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The type of one field. Valid values: ["long", "text", "double"]. Default to "long"
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public StoreIndexFieldSearchesJsonKeysGetArgs()
        {
        }
    }

    public sealed class StoreIndexFullTextArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the case sensitive for the field. Default to false. It is valid when "type" is "text" or "json".
        /// </summary>
        [Input("caseSensitive")]
        public Input<bool>? CaseSensitive { get; set; }

        /// <summary>
        /// Whether includes the chinese for the field. Default to false. It is valid when "type" is "text" or "json".
        /// </summary>
        [Input("includeChinese")]
        public Input<bool>? IncludeChinese { get; set; }

        /// <summary>
        /// The string of several split words, like "\r", "#". It is valid when "type" is "text" or "json".
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public StoreIndexFullTextArgs()
        {
        }
    }

    public sealed class StoreIndexFullTextGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the case sensitive for the field. Default to false. It is valid when "type" is "text" or "json".
        /// </summary>
        [Input("caseSensitive")]
        public Input<bool>? CaseSensitive { get; set; }

        /// <summary>
        /// Whether includes the chinese for the field. Default to false. It is valid when "type" is "text" or "json".
        /// </summary>
        [Input("includeChinese")]
        public Input<bool>? IncludeChinese { get; set; }

        /// <summary>
        /// The string of several split words, like "\r", "#". It is valid when "type" is "text" or "json".
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public StoreIndexFullTextGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class StoreIndexFieldSearches
    {
        /// <summary>
        /// The alias of one field.
        /// </summary>
        public readonly string? Alias;
        /// <summary>
        /// Whether the case sensitive for the field. Default to false. It is valid when "type" is "text" or "json".
        /// </summary>
        public readonly bool? CaseSensitive;
        /// <summary>
        /// Whether to enable field analytics. Default to true.
        /// </summary>
        public readonly bool? EnableAnalytics;
        /// <summary>
        /// Whether includes the chinese for the field. Default to false. It is valid when "type" is "text" or "json".
        /// </summary>
        public readonly bool? IncludeChinese;
        /// <summary>
        /// Use nested index when type is json
        /// </summary>
        public readonly ImmutableArray<StoreIndexFieldSearchesJsonKeys> JsonKeys;
        /// <summary>
        /// When using the json_keys field, this field is required.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The string of several split words, like "\r", "#". It is valid when "type" is "text" or "json".
        /// </summary>
        public readonly string? Token;
        /// <summary>
        /// The type of one field. Valid values: ["long", "text", "double"]. Default to "long"
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private StoreIndexFieldSearches(
            string? alias,
            bool? caseSensitive,
            bool? enableAnalytics,
            bool? includeChinese,
            ImmutableArray<StoreIndexFieldSearchesJsonKeys> jsonKeys,
            string name,
            string? token,
            string? type)
        {
            Alias = alias;
            CaseSensitive = caseSensitive;
            EnableAnalytics = enableAnalytics;
            IncludeChinese = includeChinese;
            JsonKeys = jsonKeys;
            Name = name;
            Token = token;
            Type = type;
        }
    }

    [OutputType]
    public sealed class StoreIndexFieldSearchesJsonKeys
    {
        /// <summary>
        /// The alias of one field.
        /// </summary>
        public readonly string? Alias;
        /// <summary>
        /// Whether to enable statistics. default to true.
        /// </summary>
        public readonly bool? DocValue;
        /// <summary>
        /// When using the json_keys field, this field is required.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of one field. Valid values: ["long", "text", "double"]. Default to "long"
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private StoreIndexFieldSearchesJsonKeys(
            string? alias,
            bool? docValue,
            string name,
            string? type)
        {
            Alias = alias;
            DocValue = docValue;
            Name = name;
            Type = type;
        }
    }

    [OutputType]
    public sealed class StoreIndexFullText
    {
        /// <summary>
        /// Whether the case sensitive for the field. Default to false. It is valid when "type" is "text" or "json".
        /// </summary>
        public readonly bool? CaseSensitive;
        /// <summary>
        /// Whether includes the chinese for the field. Default to false. It is valid when "type" is "text" or "json".
        /// </summary>
        public readonly bool? IncludeChinese;
        /// <summary>
        /// The string of several split words, like "\r", "#". It is valid when "type" is "text" or "json".
        /// </summary>
        public readonly string? Token;

        [OutputConstructor]
        private StoreIndexFullText(
            bool? caseSensitive,
            bool? includeChinese,
            string? token)
        {
            CaseSensitive = caseSensitive;
            IncludeChinese = includeChinese;
            Token = token;
        }
    }
    }
}
