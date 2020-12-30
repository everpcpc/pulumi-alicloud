// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.PrivateLink
{
    public static class GetVpcEndpointServiceUsers
    {
        /// <summary>
        /// This data source provides the Privatelink Vpc Endpoint Service Users of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.110.0+.
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
        ///         var example = Output.Create(AliCloud.PrivateLink.GetVpcEndpointServiceUsers.InvokeAsync(new AliCloud.PrivateLink.GetVpcEndpointServiceUsersArgs
        ///         {
        ///             ServiceId = "epsrv-gw81c6vxxxxxx",
        ///         }));
        ///         this.FirstPrivatelinkVpcEndpointServiceUserId = example.Apply(example =&gt; example.Users[0].Id);
        ///     }
        /// 
        ///     [Output("firstPrivatelinkVpcEndpointServiceUserId")]
        ///     public Output&lt;string&gt; FirstPrivatelinkVpcEndpointServiceUserId { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVpcEndpointServiceUsersResult> InvokeAsync(GetVpcEndpointServiceUsersArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcEndpointServiceUsersResult>("alicloud:privatelink/getVpcEndpointServiceUsers:getVpcEndpointServiceUsers", args ?? new GetVpcEndpointServiceUsersArgs(), options.WithVersion());
    }


    public sealed class GetVpcEndpointServiceUsersArgs : Pulumi.InvokeArgs
    {
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The Id of Vpc Endpoint Service.
        /// </summary>
        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        /// <summary>
        /// The Id of Ram User.
        /// </summary>
        [Input("userId")]
        public string? UserId { get; set; }

        public GetVpcEndpointServiceUsersArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVpcEndpointServiceUsersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        public readonly string ServiceId;
        public readonly string? UserId;
        public readonly ImmutableArray<Outputs.GetVpcEndpointServiceUsersUserResult> Users;

        [OutputConstructor]
        private GetVpcEndpointServiceUsersResult(
            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            string serviceId,

            string? userId,

            ImmutableArray<Outputs.GetVpcEndpointServiceUsersUserResult> users)
        {
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            ServiceId = serviceId;
            UserId = userId;
            Users = users;
        }
    }
}
