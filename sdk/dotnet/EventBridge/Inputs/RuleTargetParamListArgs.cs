// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.EventBridge.Inputs
{

    public sealed class RuleTargetParamListArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The format of param.
        /// </summary>
        [Input("form", required: true)]
        public Input<string> Form { get; set; } = null!;

        /// <summary>
        /// The resource key of param. For more information, see [Event target parameters](https://help.aliyun.com/document_detail/185887.htm)
        /// </summary>
        [Input("resourceKey", required: true)]
        public Input<string> ResourceKey { get; set; } = null!;

        /// <summary>
        /// The template of param.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        /// <summary>
        /// The value of param.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public RuleTargetParamListArgs()
        {
        }
    }
}
