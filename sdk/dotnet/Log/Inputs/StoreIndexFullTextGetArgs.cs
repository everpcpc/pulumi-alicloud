// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Log.Inputs
{

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