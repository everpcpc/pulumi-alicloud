// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae.Inputs
{

    public sealed class GreyTagRouteDubboRuleGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Conditional Patterns for Grayscale Rules. Valid values: `AND`, `OR`.
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        /// <summary>
        /// The service group.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("items")]
        private InputList<Inputs.GreyTagRouteDubboRuleItemGetArgs>? _items;

        /// <summary>
        /// A list of conditions items. The details see Block `dubbo_rules_items`.
        /// </summary>
        public InputList<Inputs.GreyTagRouteDubboRuleItemGetArgs> Items
        {
            get => _items ?? (_items = new InputList<Inputs.GreyTagRouteDubboRuleItemGetArgs>());
            set => _items = value;
        }

        /// <summary>
        /// The method name
        /// </summary>
        [Input("methodName")]
        public Input<string>? MethodName { get; set; }

        /// <summary>
        /// The service name.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// The service version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public GreyTagRouteDubboRuleGetArgs()
        {
        }
    }
}
