// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Emrv2.Inputs
{

    public sealed class ClusterBootstrapScriptGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bootstrap scripts execution fail strategy, ’FAILED_BLOCKED’ or ‘FAILED_CONTINUE’ .
        /// </summary>
        [Input("executionFailStrategy", required: true)]
        public Input<string> ExecutionFailStrategy { get; set; } = null!;

        /// <summary>
        /// The bootstrap scripts execution moment, ’BEFORE_INSTALL’ or ‘AFTER_STARTED’ .
        /// </summary>
        [Input("executionMoment", required: true)]
        public Input<string> ExecutionMoment { get; set; } = null!;

        /// <summary>
        /// The bootstrap scripts execution target.
        /// </summary>
        [Input("nodeSelector", required: true)]
        public Input<Inputs.ClusterBootstrapScriptNodeSelectorGetArgs> NodeSelector { get; set; } = null!;

        /// <summary>
        /// The bootstrap scripts priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The bootstrap script args, e.g. "--a=b".
        /// </summary>
        [Input("scriptArgs", required: true)]
        public Input<string> ScriptArgs { get; set; } = null!;

        /// <summary>
        /// The bootstrap script name.
        /// </summary>
        [Input("scriptName", required: true)]
        public Input<string> ScriptName { get; set; } = null!;

        /// <summary>
        /// The bootstrap script path, e.g. "oss://bucket/path".
        /// </summary>
        [Input("scriptPath", required: true)]
        public Input<string> ScriptPath { get; set; } = null!;

        public ClusterBootstrapScriptGetArgs()
        {
        }
        public static new ClusterBootstrapScriptGetArgs Empty => new ClusterBootstrapScriptGetArgs();
    }
}
