// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ExpressConnect.Inputs
{

    public sealed class GetRouterInterfacesFilterArgs : global::Pulumi.InvokeArgs
    {
        [Input("key")]
        public string? Key { get; set; }

        [Input("values")]
        private List<string>? _values;
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetRouterInterfacesFilterArgs()
        {
        }
        public static new GetRouterInterfacesFilterArgs Empty => new GetRouterInterfacesFilterArgs();
    }
}
