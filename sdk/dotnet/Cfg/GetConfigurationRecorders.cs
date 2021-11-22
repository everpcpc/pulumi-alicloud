// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.AliCloud.Cfg
{
    public static class GetConfigurationRecorders
    {
        /// <summary>
        /// This data source provides the Config Configuration Recorders of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:**  Available in 1.99.0+.
        /// 
        /// &gt; **NOTE:** The Cloud Config region only support `cn-shanghai` and `ap-northeast-1`.
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
        ///         var example = Output.Create(AliCloud.Cfg.GetConfigurationRecorders.InvokeAsync());
        ///         this.ListOfResourceTypes = data.Alicloud_config_configuration_recorders.This.Recorders[0].Resource_types;
        ///     }
        /// 
        ///     [Output("listOfResourceTypes")]
        ///     public Output&lt;string&gt; ListOfResourceTypes { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetConfigurationRecordersResult> InvokeAsync(GetConfigurationRecordersArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConfigurationRecordersResult>("alicloud:cfg/getConfigurationRecorders:getConfigurationRecorders", args ?? new GetConfigurationRecordersArgs(), options.WithVersion());

        /// <summary>
        /// This data source provides the Config Configuration Recorders of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:**  Available in 1.99.0+.
        /// 
        /// &gt; **NOTE:** The Cloud Config region only support `cn-shanghai` and `ap-northeast-1`.
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
        ///         var example = Output.Create(AliCloud.Cfg.GetConfigurationRecorders.InvokeAsync());
        ///         this.ListOfResourceTypes = data.Alicloud_config_configuration_recorders.This.Recorders[0].Resource_types;
        ///     }
        /// 
        ///     [Output("listOfResourceTypes")]
        ///     public Output&lt;string&gt; ListOfResourceTypes { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetConfigurationRecordersResult> Invoke(GetConfigurationRecordersInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConfigurationRecordersResult>("alicloud:cfg/getConfigurationRecorders:getConfigurationRecorders", args ?? new GetConfigurationRecordersInvokeArgs(), options.WithVersion());
    }


    public sealed class GetConfigurationRecordersArgs : Pulumi.InvokeArgs
    {
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetConfigurationRecordersArgs()
        {
        }
    }

    public sealed class GetConfigurationRecordersInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetConfigurationRecordersInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConfigurationRecordersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;
        /// <summary>
        /// A list of Config Configuration Recorders. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetConfigurationRecordersRecorderResult> Recorders;

        [OutputConstructor]
        private GetConfigurationRecordersResult(
            string id,

            string? outputFile,

            ImmutableArray<Outputs.GetConfigurationRecordersRecorderResult> recorders)
        {
            Id = id;
            OutputFile = outputFile;
            Recorders = recorders;
        }
    }
}
