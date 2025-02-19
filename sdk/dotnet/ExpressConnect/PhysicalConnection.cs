// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ExpressConnect
{
    /// <summary>
    /// Provides a Express Connect Physical Connection resource.
    /// 
    /// For information about Express Connect Physical Connection and how to use it, see [What is Physical Connection](https://www.alibabacloud.com/help/doc-detail/44852.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.132.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var domestic = new AliCloud.ExpressConnect.PhysicalConnection("domestic", new()
    ///     {
    ///         AccessPointId = "ap-cn-hangzhou-yh-B",
    ///         Bandwidth = "100",
    ///         Description = "my domestic connection",
    ///         LineOperator = "CT",
    ///         PeerLocation = "example_value",
    ///         PhysicalConnectionName = "example_value",
    ///         PortType = "1000Base-LX",
    ///         Type = "VPC",
    ///     });
    /// 
    ///     var international = new AliCloud.ExpressConnect.PhysicalConnection("international", new()
    ///     {
    ///         AccessPointId = "ap-sg-singpore-A",
    ///         Bandwidth = "100",
    ///         Description = "my domestic connection",
    ///         LineOperator = "Other",
    ///         PeerLocation = "example_value",
    ///         PhysicalConnectionName = "example_value",
    ///         PortType = "1000Base-LX",
    ///         Type = "VPC",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Express Connect Physical Connection can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:expressconnect/physicalConnection:PhysicalConnection example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:expressconnect/physicalConnection:PhysicalConnection")]
    public partial class PhysicalConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Physical Leased Line Access Point ID.
        /// </summary>
        [Output("accessPointId")]
        public Output<string> AccessPointId { get; private set; } = null!;

        /// <summary>
        /// On the Bandwidth of the ECC Service and Physical Connection.
        /// </summary>
        [Output("bandwidth")]
        public Output<string> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// Operators for Physical Connection Circuit Provided Coding.
        /// </summary>
        [Output("circuitCode")]
        public Output<string?> CircuitCode { get; private set; } = null!;

        /// <summary>
        /// The Physical Connection to Which the Description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Provides Access to the Physical Line Operator. Valid values:
        /// * CT: China Telecom
        /// * CU: China Unicom
        /// * CM: china Mobile
        /// * CO: Other Chinese
        /// * Equinix: Equinix
        /// * Other: Other Overseas.
        /// </summary>
        [Output("lineOperator")]
        public Output<string> LineOperator { get; private set; } = null!;

        /// <summary>
        /// and an on-Premises Data Center Location.
        /// </summary>
        [Output("peerLocation")]
        public Output<string?> PeerLocation { get; private set; } = null!;

        /// <summary>
        /// on Behalf of the Resource Name of the Resources-Attribute Field.
        /// </summary>
        [Output("physicalConnectionName")]
        public Output<string?> PhysicalConnectionName { get; private set; } = null!;

        /// <summary>
        /// The Physical Leased Line Access Port Type. Valid value:
        /// * 100Base-T: Fast Electrical Ports
        /// * 1000Base-T: gigabit Electrical Ports
        /// * 1000Base-LX: Gigabit Singlemode Optical Ports (10Km)
        /// * 10GBase-T: Gigabit Electrical Port
        /// * 10GBase-LR: Gigabit Singlemode Optical Ports (10Km).
        /// * 40GBase-LR: 40 Gigabit Singlemode Optical Ports.
        /// * 100GBase-LR: One hundred thousand Gigabit Singlemode Optical Ports.
        /// 
        /// **NOTE:** From in v1.185.0+, The `40GBase-LR` and `100GBase-LR` is valid. and Set these values based on the water levels of background ports. For details about the water levels, contact the business manager.
        /// </summary>
        [Output("portType")]
        public Output<string?> PortType { get; private set; } = null!;

        /// <summary>
        /// Redundant Physical Connection to Which the ID.
        /// </summary>
        [Output("redundantPhysicalConnectionId")]
        public Output<string?> RedundantPhysicalConnectionId { get; private set; } = null!;

        /// <summary>
        /// Resources on Behalf of a State of the Resource Attribute Field. Valid values: `Canceled`, `Enabled`, `Terminated`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Physical Private Line of Type. Default Value: VPC.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PhysicalConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PhysicalConnection(string name, PhysicalConnectionArgs args, CustomResourceOptions? options = null)
            : base("alicloud:expressconnect/physicalConnection:PhysicalConnection", name, args ?? new PhysicalConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PhysicalConnection(string name, Input<string> id, PhysicalConnectionState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:expressconnect/physicalConnection:PhysicalConnection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PhysicalConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PhysicalConnection Get(string name, Input<string> id, PhysicalConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new PhysicalConnection(name, id, state, options);
        }
    }

    public sealed class PhysicalConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Physical Leased Line Access Point ID.
        /// </summary>
        [Input("accessPointId", required: true)]
        public Input<string> AccessPointId { get; set; } = null!;

        /// <summary>
        /// On the Bandwidth of the ECC Service and Physical Connection.
        /// </summary>
        [Input("bandwidth")]
        public Input<string>? Bandwidth { get; set; }

        /// <summary>
        /// Operators for Physical Connection Circuit Provided Coding.
        /// </summary>
        [Input("circuitCode")]
        public Input<string>? CircuitCode { get; set; }

        /// <summary>
        /// The Physical Connection to Which the Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Provides Access to the Physical Line Operator. Valid values:
        /// * CT: China Telecom
        /// * CU: China Unicom
        /// * CM: china Mobile
        /// * CO: Other Chinese
        /// * Equinix: Equinix
        /// * Other: Other Overseas.
        /// </summary>
        [Input("lineOperator", required: true)]
        public Input<string> LineOperator { get; set; } = null!;

        /// <summary>
        /// and an on-Premises Data Center Location.
        /// </summary>
        [Input("peerLocation")]
        public Input<string>? PeerLocation { get; set; }

        /// <summary>
        /// on Behalf of the Resource Name of the Resources-Attribute Field.
        /// </summary>
        [Input("physicalConnectionName")]
        public Input<string>? PhysicalConnectionName { get; set; }

        /// <summary>
        /// The Physical Leased Line Access Port Type. Valid value:
        /// * 100Base-T: Fast Electrical Ports
        /// * 1000Base-T: gigabit Electrical Ports
        /// * 1000Base-LX: Gigabit Singlemode Optical Ports (10Km)
        /// * 10GBase-T: Gigabit Electrical Port
        /// * 10GBase-LR: Gigabit Singlemode Optical Ports (10Km).
        /// * 40GBase-LR: 40 Gigabit Singlemode Optical Ports.
        /// * 100GBase-LR: One hundred thousand Gigabit Singlemode Optical Ports.
        /// 
        /// **NOTE:** From in v1.185.0+, The `40GBase-LR` and `100GBase-LR` is valid. and Set these values based on the water levels of background ports. For details about the water levels, contact the business manager.
        /// </summary>
        [Input("portType")]
        public Input<string>? PortType { get; set; }

        /// <summary>
        /// Redundant Physical Connection to Which the ID.
        /// </summary>
        [Input("redundantPhysicalConnectionId")]
        public Input<string>? RedundantPhysicalConnectionId { get; set; }

        /// <summary>
        /// Resources on Behalf of a State of the Resource Attribute Field. Valid values: `Canceled`, `Enabled`, `Terminated`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Physical Private Line of Type. Default Value: VPC.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public PhysicalConnectionArgs()
        {
        }
        public static new PhysicalConnectionArgs Empty => new PhysicalConnectionArgs();
    }

    public sealed class PhysicalConnectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Physical Leased Line Access Point ID.
        /// </summary>
        [Input("accessPointId")]
        public Input<string>? AccessPointId { get; set; }

        /// <summary>
        /// On the Bandwidth of the ECC Service and Physical Connection.
        /// </summary>
        [Input("bandwidth")]
        public Input<string>? Bandwidth { get; set; }

        /// <summary>
        /// Operators for Physical Connection Circuit Provided Coding.
        /// </summary>
        [Input("circuitCode")]
        public Input<string>? CircuitCode { get; set; }

        /// <summary>
        /// The Physical Connection to Which the Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Provides Access to the Physical Line Operator. Valid values:
        /// * CT: China Telecom
        /// * CU: China Unicom
        /// * CM: china Mobile
        /// * CO: Other Chinese
        /// * Equinix: Equinix
        /// * Other: Other Overseas.
        /// </summary>
        [Input("lineOperator")]
        public Input<string>? LineOperator { get; set; }

        /// <summary>
        /// and an on-Premises Data Center Location.
        /// </summary>
        [Input("peerLocation")]
        public Input<string>? PeerLocation { get; set; }

        /// <summary>
        /// on Behalf of the Resource Name of the Resources-Attribute Field.
        /// </summary>
        [Input("physicalConnectionName")]
        public Input<string>? PhysicalConnectionName { get; set; }

        /// <summary>
        /// The Physical Leased Line Access Port Type. Valid value:
        /// * 100Base-T: Fast Electrical Ports
        /// * 1000Base-T: gigabit Electrical Ports
        /// * 1000Base-LX: Gigabit Singlemode Optical Ports (10Km)
        /// * 10GBase-T: Gigabit Electrical Port
        /// * 10GBase-LR: Gigabit Singlemode Optical Ports (10Km).
        /// * 40GBase-LR: 40 Gigabit Singlemode Optical Ports.
        /// * 100GBase-LR: One hundred thousand Gigabit Singlemode Optical Ports.
        /// 
        /// **NOTE:** From in v1.185.0+, The `40GBase-LR` and `100GBase-LR` is valid. and Set these values based on the water levels of background ports. For details about the water levels, contact the business manager.
        /// </summary>
        [Input("portType")]
        public Input<string>? PortType { get; set; }

        /// <summary>
        /// Redundant Physical Connection to Which the ID.
        /// </summary>
        [Input("redundantPhysicalConnectionId")]
        public Input<string>? RedundantPhysicalConnectionId { get; set; }

        /// <summary>
        /// Resources on Behalf of a State of the Resource Attribute Field. Valid values: `Canceled`, `Enabled`, `Terminated`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Physical Private Line of Type. Default Value: VPC.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public PhysicalConnectionState()
        {
        }
        public static new PhysicalConnectionState Empty => new PhysicalConnectionState();
    }
}
