// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.Arms
{
    public static class GetAlertContactGroups
    {
        /// <summary>
        /// This data source provides the Arms Alert Contact Groups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.131.0+.
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
        ///         var nameRegex = Output.Create(AliCloud.Arms.GetAlertContactGroups.InvokeAsync(new AliCloud.Arms.GetAlertContactGroupsArgs
        ///         {
        ///             NameRegex = "^my-AlertContactGroup",
        ///         }));
        ///         this.ArmsAlertContactGroupId = nameRegex.Apply(nameRegex =&gt; nameRegex.Groups?[0]?.Id);
        ///     }
        /// 
        ///     [Output("armsAlertContactGroupId")]
        ///     public Output&lt;string&gt; ArmsAlertContactGroupId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAlertContactGroupsResult> InvokeAsync(GetAlertContactGroupsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAlertContactGroupsResult>("alicloud:arms/getAlertContactGroups:getAlertContactGroups", args ?? new GetAlertContactGroupsArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides the Arms Alert Contact Groups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.131.0+.
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
        ///         var nameRegex = Output.Create(AliCloud.Arms.GetAlertContactGroups.InvokeAsync(new AliCloud.Arms.GetAlertContactGroupsArgs
        ///         {
        ///             NameRegex = "^my-AlertContactGroup",
        ///         }));
        ///         this.ArmsAlertContactGroupId = nameRegex.Apply(nameRegex =&gt; nameRegex.Groups?[0]?.Id);
        ///     }
        /// 
        ///     [Output("armsAlertContactGroupId")]
        ///     public Output&lt;string&gt; ArmsAlertContactGroupId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAlertContactGroupsResult> Invoke(GetAlertContactGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAlertContactGroupsResult>("alicloud:arms/getAlertContactGroups:getAlertContactGroups", args ?? new GetAlertContactGroupsInvokeArgs(), options.WithVersion());
    }


    public sealed class GetAlertContactGroupsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("alertContactGroupName")]
        public string? AlertContactGroupName { get; set; }

        /// <summary>
        /// The contact id.
        /// </summary>
        [Input("contactId")]
        public string? ContactId { get; set; }

        /// <summary>
        /// The contact name.
        /// </summary>
        [Input("contactName")]
        public string? ContactName { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Alert Contact Group IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Alert Contact Group name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetAlertContactGroupsArgs()
        {
        }
    }

    public sealed class GetAlertContactGroupsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("alertContactGroupName")]
        public Input<string>? AlertContactGroupName { get; set; }

        /// <summary>
        /// The contact id.
        /// </summary>
        [Input("contactId")]
        public Input<string>? ContactId { get; set; }

        /// <summary>
        /// The contact name.
        /// </summary>
        [Input("contactName")]
        public Input<string>? ContactName { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Alert Contact Group IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Alert Contact Group name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetAlertContactGroupsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAlertContactGroupsResult
    {
        public readonly string? AlertContactGroupName;
        public readonly string? ContactId;
        public readonly string? ContactName;
        public readonly ImmutableArray<Outputs.GetAlertContactGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;

        [OutputConstructor]
        private GetAlertContactGroupsResult(
            string? alertContactGroupName,

            string? contactId,

            string? contactName,

            ImmutableArray<Outputs.GetAlertContactGroupsGroupResult> groups,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile)
        {
            AlertContactGroupName = alertContactGroupName;
            ContactId = contactId;
            ContactName = contactName;
            Groups = groups;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
        }
    }
}
