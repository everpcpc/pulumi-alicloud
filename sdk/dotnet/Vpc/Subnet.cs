// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc
{
    public partial class Subnet : Pulumi.CustomResource
    {
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Subnet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subnet(string name, SubnetArgs args, CustomResourceOptions? options = null)
            : base("alicloud:vpc/subnet:Subnet", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Subnet(string name, Input<string> id, SubnetState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:vpc/subnet:Subnet", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Subnet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subnet Get(string name, Input<string> id, SubnetState? state = null, CustomResourceOptions? options = null)
        {
            return new Subnet(name, id, state, options);
        }
    }

    public sealed class SubnetArgs : Pulumi.ResourceArgs
    {
        [Input("availabilityZone", required: true)]
        public Input<string> AvailabilityZone { get; set; } = null!;

        [Input("cidrBlock", required: true)]
        public Input<string> CidrBlock { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public SubnetArgs()
        {
        }
    }

    public sealed class SubnetState : Pulumi.ResourceArgs
    {
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public SubnetState()
        {
        }
    }
}
