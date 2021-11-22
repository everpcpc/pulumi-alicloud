// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.ActionTrail
{
    public static class GetSaslAcls
    {
        /// <summary>
        /// This data source provides a list of ALIKAFKA Sasl acls in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.66.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var saslAclsDs = Output.Create(AliCloud.ActionTrail.GetSaslAcls.InvokeAsync(new AliCloud.ActionTrail.GetSaslAclsArgs
        ///         {
        ///             AclResourceName = "testTopic",
        ///             AclResourceType = "Topic",
        ///             InstanceId = "xxx",
        ///             OutputFile = "saslAcls.txt",
        ///             Username = "username",
        ///         }));
        ///         this.FirstSaslAclUsername = saslAclsDs.Apply(saslAclsDs =&gt; saslAclsDs.Acls?[0]?.Username);
        ///     }
        /// 
        ///     [Output("firstSaslAclUsername")]
        ///     public Output&lt;string&gt; FirstSaslAclUsername { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSaslAclsResult> InvokeAsync(GetSaslAclsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSaslAclsResult>("alicloud:actiontrail/getSaslAcls:getSaslAcls", args ?? new GetSaslAclsArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides a list of ALIKAFKA Sasl acls in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.66.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var saslAclsDs = Output.Create(AliCloud.ActionTrail.GetSaslAcls.InvokeAsync(new AliCloud.ActionTrail.GetSaslAclsArgs
        ///         {
        ///             AclResourceName = "testTopic",
        ///             AclResourceType = "Topic",
        ///             InstanceId = "xxx",
        ///             OutputFile = "saslAcls.txt",
        ///             Username = "username",
        ///         }));
        ///         this.FirstSaslAclUsername = saslAclsDs.Apply(saslAclsDs =&gt; saslAclsDs.Acls?[0]?.Username);
        ///     }
        /// 
        ///     [Output("firstSaslAclUsername")]
        ///     public Output&lt;string&gt; FirstSaslAclUsername { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSaslAclsResult> Invoke(GetSaslAclsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSaslAclsResult>("alicloud:actiontrail/getSaslAcls:getSaslAcls", args ?? new GetSaslAclsInvokeArgs(), options.WithVersion());
    }


    public sealed class GetSaslAclsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Get results for the specified resource name.
        /// </summary>
        [Input("aclResourceName", required: true)]
        public string AclResourceName { get; set; } = null!;

        /// <summary>
        /// Get results for the specified resource type.
        /// </summary>
        [Input("aclResourceType", required: true)]
        public string AclResourceType { get; set; } = null!;

        /// <summary>
        /// ID of the ALIKAFKA Instance that owns the sasl acls.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Get results for the specified username.
        /// </summary>
        [Input("username", required: true)]
        public string Username { get; set; } = null!;

        public GetSaslAclsArgs()
        {
        }
    }

    public sealed class GetSaslAclsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Get results for the specified resource name.
        /// </summary>
        [Input("aclResourceName", required: true)]
        public Input<string> AclResourceName { get; set; } = null!;

        /// <summary>
        /// Get results for the specified resource type.
        /// </summary>
        [Input("aclResourceType", required: true)]
        public Input<string> AclResourceType { get; set; } = null!;

        /// <summary>
        /// ID of the ALIKAFKA Instance that owns the sasl acls.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Get results for the specified username.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public GetSaslAclsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSaslAclsResult
    {
        /// <summary>
        /// The resource name of the sasl acl.
        /// </summary>
        public readonly string AclResourceName;
        /// <summary>
        /// The resource type of the sasl acl.
        /// </summary>
        public readonly string AclResourceType;
        /// <summary>
        /// A list of sasl acls. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSaslAclsAclResult> Acls;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? OutputFile;
        /// <summary>
        /// The username of the sasl acl.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetSaslAclsResult(
            string aclResourceName,

            string aclResourceType,

            ImmutableArray<Outputs.GetSaslAclsAclResult> acls,

            string id,

            string instanceId,

            string? outputFile,

            string username)
        {
            AclResourceName = aclResourceName;
            AclResourceType = aclResourceType;
            Acls = acls;
            Id = id;
            InstanceId = instanceId;
            OutputFile = outputFile;
            Username = username;
        }
    }
}
