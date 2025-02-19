// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS.Inputs
{

    public sealed class ServerlessKubernetesAddonArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ACK add-on configurations.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// Disables the automatic installation of a component. Default is `false`.
        /// 
        /// The following example is the definition of addons block, The type of this field is list:
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Name of the ACK add-on. The name must match one of the names returned by [DescribeAddons](https://help.aliyun.com/document_detail/171524.html).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ServerlessKubernetesAddonArgs()
        {
        }
        public static new ServerlessKubernetesAddonArgs Empty => new ServerlessKubernetesAddonArgs();
    }
}
