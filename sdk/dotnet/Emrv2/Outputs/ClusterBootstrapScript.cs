// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Emrv2.Outputs
{

    [OutputType]
    public sealed class ClusterBootstrapScript
    {
        /// <summary>
        /// The bootstrap scripts execution fail strategy, ’FAILED_BLOCKED’ or ‘FAILED_CONTINUE’ .
        /// </summary>
        public readonly string ExecutionFailStrategy;
        /// <summary>
        /// The bootstrap scripts execution moment, ’BEFORE_INSTALL’ or ‘AFTER_STARTED’ .
        /// </summary>
        public readonly string ExecutionMoment;
        /// <summary>
        /// The bootstrap scripts execution target.
        /// </summary>
        public readonly Outputs.ClusterBootstrapScriptNodeSelector NodeSelector;
        /// <summary>
        /// The bootstrap scripts priority.
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// The bootstrap script args, e.g. "--a=b".
        /// </summary>
        public readonly string ScriptArgs;
        /// <summary>
        /// The bootstrap script name.
        /// </summary>
        public readonly string ScriptName;
        /// <summary>
        /// The bootstrap script path, e.g. "oss://bucket/path".
        /// </summary>
        public readonly string ScriptPath;

        [OutputConstructor]
        private ClusterBootstrapScript(
            string executionFailStrategy,

            string executionMoment,

            Outputs.ClusterBootstrapScriptNodeSelector nodeSelector,

            int? priority,

            string scriptArgs,

            string scriptName,

            string scriptPath)
        {
            ExecutionFailStrategy = executionFailStrategy;
            ExecutionMoment = executionMoment;
            NodeSelector = nodeSelector;
            Priority = priority;
            ScriptArgs = scriptArgs;
            ScriptName = scriptName;
            ScriptPath = scriptPath;
        }
    }
}
