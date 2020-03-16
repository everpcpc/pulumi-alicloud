// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sag
{
    public static partial class Invokes
    {
        /// <summary>
        /// This data source provides Sag Acls available to the user.
        /// 
        /// &gt; **NOTE:** Available in 1.60.0+
        /// 
        /// &gt; **NOTE:** Only the following regions support create Cloud Connect Network. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/d/sag_acls.html.markdown.
        /// </summary>
        public static Task<GetAclsResult> GetAcls(GetAclsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAclsResult>("alicloud:sag/getAcls:getAcls", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetAclsArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Sag Acl IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter Sag Acl instances by name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetAclsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetAclsResult
    {
        /// <summary>
        /// A list of Sag Acls. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAclsAclsResult> Acls;
        /// <summary>
        /// A list of Sag Acl IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of Sag Acls names. 
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAclsResult(
            ImmutableArray<Outputs.GetAclsAclsResult> acls,
            ImmutableArray<string> ids,
            string? nameRegex,
            ImmutableArray<string> names,
            string? outputFile,
            string id)
        {
            Acls = acls;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetAclsAclsResult
    {
        /// <summary>
        /// The ID of the ACL. For example "acl-xxx".
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the Acl.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetAclsAclsResult(
            string id,
            string name)
        {
            Id = id;
            Name = name;
        }
    }
    }
}
