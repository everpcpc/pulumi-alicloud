// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Eds.Inputs
{

    public sealed class EcdPolicyGroupAuthorizeAccessPolicyRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cidrip of authorize access rule.
        /// </summary>
        [Input("cidrIp")]
        public Input<string>? CidrIp { get; set; }

        /// <summary>
        /// The description of authorize access rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        public EcdPolicyGroupAuthorizeAccessPolicyRuleArgs()
        {
        }
    }
}
