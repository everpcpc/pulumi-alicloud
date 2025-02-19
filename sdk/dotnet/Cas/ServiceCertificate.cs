// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Cas
{
    /// <summary>
    /// Provides a SSL Certificates Certificate resource.
    /// 
    /// For information about SSL Certificates Certificate and how to use it, see [What is Certificate](https://www.alibabacloud.com/help/product/28533.html).
    /// 
    /// &gt; **NOTE:** Available in v1.129.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AliCloud.Cas.ServiceCertificate("example", new()
    ///     {
    ///         CertificateName = "test",
    ///         Cert = File.ReadAllText($"{path.Module}/test.crt"),
    ///         Key = File.ReadAllText($"{path.Module}/test.key"),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SSL Certificates Certificate can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cas/serviceCertificate:ServiceCertificate example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cas/serviceCertificate:ServiceCertificate")]
    public partial class ServiceCertificate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cert of the Certificate in which the Certificate will add.
        /// </summary>
        [Output("cert")]
        public Output<string> Cert { get; private set; } = null!;

        /// <summary>
        /// Name of the Certificate. 
        /// This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-",
        /// and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time.
        /// Suffix .sh and .tel are not supported.
        /// **NOTE:** One of `certificate_name` and `name` must be specified.
        /// </summary>
        [Output("certificateName")]
        public Output<string> CertificateName { get; private set; } = null!;

        /// <summary>
        /// Key of the Certificate in which the Certificate will add.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The lang.
        /// </summary>
        [Output("lang")]
        public Output<string?> Lang { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.129.0 and using `certificate_name` instead.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceCertificate(string name, ServiceCertificateArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cas/serviceCertificate:ServiceCertificate", name, args ?? new ServiceCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceCertificate(string name, Input<string> id, ServiceCertificateState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cas/serviceCertificate:ServiceCertificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceCertificate Get(string name, Input<string> id, ServiceCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceCertificate(name, id, state, options);
        }
    }

    public sealed class ServiceCertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cert of the Certificate in which the Certificate will add.
        /// </summary>
        [Input("cert", required: true)]
        public Input<string> Cert { get; set; } = null!;

        /// <summary>
        /// Name of the Certificate. 
        /// This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-",
        /// and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time.
        /// Suffix .sh and .tel are not supported.
        /// **NOTE:** One of `certificate_name` and `name` must be specified.
        /// </summary>
        [Input("certificateName")]
        public Input<string>? CertificateName { get; set; }

        /// <summary>
        /// Key of the Certificate in which the Certificate will add.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The lang.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.129.0 and using `certificate_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ServiceCertificateArgs()
        {
        }
        public static new ServiceCertificateArgs Empty => new ServiceCertificateArgs();
    }

    public sealed class ServiceCertificateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cert of the Certificate in which the Certificate will add.
        /// </summary>
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        /// <summary>
        /// Name of the Certificate. 
        /// This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-",
        /// and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time.
        /// Suffix .sh and .tel are not supported.
        /// **NOTE:** One of `certificate_name` and `name` must be specified.
        /// </summary>
        [Input("certificateName")]
        public Input<string>? CertificateName { get; set; }

        /// <summary>
        /// Key of the Certificate in which the Certificate will add.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The lang.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.129.0 and using `certificate_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ServiceCertificateState()
        {
        }
        public static new ServiceCertificateState Empty => new ServiceCertificateState();
    }
}
