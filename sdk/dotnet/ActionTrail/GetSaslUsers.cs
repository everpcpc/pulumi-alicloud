// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ActionTrail
{
    public static class GetSaslUsers
    {
        /// <summary>
        /// This data source provides a list of ALIKAFKA Sasl users in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.66.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var saslUsersDs = AliCloud.ActionTrail.GetSaslUsers.Invoke(new()
        ///     {
        ///         InstanceId = "xxx",
        ///         NameRegex = "username",
        ///         OutputFile = "saslUsers.txt",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstSaslUsername"] = saslUsersDs.Apply(getSaslUsersResult =&gt; getSaslUsersResult.Users[0]?.Username),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSaslUsersResult> InvokeAsync(GetSaslUsersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSaslUsersResult>("alicloud:actiontrail/getSaslUsers:getSaslUsers", args ?? new GetSaslUsersArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides a list of ALIKAFKA Sasl users in an Alibaba Cloud account according to the specified filters.
        /// 
        /// &gt; **NOTE:** Available in 1.66.0+
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AliCloud = Pulumi.AliCloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var saslUsersDs = AliCloud.ActionTrail.GetSaslUsers.Invoke(new()
        ///     {
        ///         InstanceId = "xxx",
        ///         NameRegex = "username",
        ///         OutputFile = "saslUsers.txt",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstSaslUsername"] = saslUsersDs.Apply(getSaslUsersResult =&gt; getSaslUsersResult.Users[0]?.Username),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSaslUsersResult> Invoke(GetSaslUsersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSaslUsersResult>("alicloud:actiontrail/getSaslUsers:getSaslUsers", args ?? new GetSaslUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSaslUsersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the ALIKAFKA Instance that owns the sasl users.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by the username.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetSaslUsersArgs()
        {
        }
        public static new GetSaslUsersArgs Empty => new GetSaslUsersArgs();
    }

    public sealed class GetSaslUsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the ALIKAFKA Instance that owns the sasl users.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// A regex string to filter results by the username.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetSaslUsersInvokeArgs()
        {
        }
        public static new GetSaslUsersInvokeArgs Empty => new GetSaslUsersInvokeArgs();
    }


    [OutputType]
    public sealed class GetSaslUsersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of sasl usernames.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of sasl users. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSaslUsersUserResult> Users;

        [OutputConstructor]
        private GetSaslUsersResult(
            string id,

            string instanceId,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            ImmutableArray<Outputs.GetSaslUsersUserResult> users)
        {
            Id = id;
            InstanceId = instanceId;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            Users = users;
        }
    }
}
