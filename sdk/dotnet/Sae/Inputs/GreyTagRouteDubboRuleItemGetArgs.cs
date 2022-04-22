// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae.Inputs
{

    public sealed class GreyTagRouteDubboRuleItemGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The comparison operator. Valid values: `&gt;`, `&lt;`, `&gt;=`, `&lt;=`, `==`, `!=`.
        /// </summary>
        [Input("cond")]
        public Input<string>? Cond { get; set; }

        /// <summary>
        /// The parameter value gets the expression.
        /// </summary>
        [Input("expr")]
        public Input<string>? Expr { get; set; }

        /// <summary>
        /// The parameter number.
        /// </summary>
        [Input("index")]
        public Input<int>? Index { get; set; }

        /// <summary>
        /// The operator. Valid values: `rawvalue`, `list`, `mod`, `deterministic_proportional_steaming_division`
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        /// <summary>
        /// The value of the parameter.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public GreyTagRouteDubboRuleItemGetArgs()
        {
        }
    }
}