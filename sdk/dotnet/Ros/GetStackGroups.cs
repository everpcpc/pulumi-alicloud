// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ros
{
    public static class GetStackGroups
    {
        /// <summary>
        /// This data source provides the Ros Stack Groups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.107.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///         var example = Output.Create(AliCloud.Ros.GetStackGroups.InvokeAsync(new AliCloud.Ros.GetStackGroupsArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "example_value",
        ///             },
        ///             NameRegex = "the_resource_name",
        ///         }));
        ///         this.FirstRosStackGroupId = example.Apply(example =&gt; example.Groups[0].Id);
        ///     }
        /// 
        ///     [Output("firstRosStackGroupId")]
        ///     public Output&lt;string&gt; FirstRosStackGroupId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetStackGroupsResult> InvokeAsync(GetStackGroupsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStackGroupsResult>("alicloud:ros/getStackGroups:getStackGroups", args ?? new GetStackGroupsArgs(), options.WithVersion());
    }


    public sealed class GetStackGroupsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Default to `false`. Set it to `true` can output more details about resource attributes.
        /// </summary>
        [Input("enableDetails")]
        public bool? EnableDetails { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Stack Group IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Stack Group name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The status of Stack Group.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetStackGroupsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStackGroupsResult
    {
        public readonly bool? EnableDetails;
        public readonly ImmutableArray<Outputs.GetStackGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? Status;

        [OutputConstructor]
        private GetStackGroupsResult(
            bool? enableDetails,

            ImmutableArray<Outputs.GetStackGroupsGroupResult> groups,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? status)
        {
            EnableDetails = enableDetails;
            Groups = groups;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Status = status;
        }
    }
}