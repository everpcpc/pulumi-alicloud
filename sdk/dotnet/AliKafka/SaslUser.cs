// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.AliKafka
{
    /// <summary>
    /// Provides an ALIKAFKA sasl user resource.
    /// 
    /// &gt; **NOTE:** Available in 1.66.0+
    /// 
    /// &gt; **NOTE:**  Only the following regions support create alikafka sasl user.
    /// [`cn-hangzhou`,`cn-beijing`,`cn-shenzhen`,`cn-shanghai`,`cn-qingdao`,`cn-hongkong`,`cn-huhehaote`,`cn-zhangjiakou`,`cn-chengdu`,`cn-heyuan`,`ap-southeast-1`,`ap-southeast-3`,`ap-southeast-5`,`ap-south-1`,`ap-northeast-1`,`eu-central-1`,`eu-west-1`,`us-west-1`,`us-east-1`]
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var username = config.Get("username") ?? "testusername";
    ///         var password = config.Get("password") ?? "testpassword";
    ///         var defaultZones = Output.Create(AliCloud.GetZones.InvokeAsync(new AliCloud.GetZonesArgs
    ///         {
    ///             AvailableResourceCreation = "VSwitch",
    ///         }));
    ///         var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new AliCloud.Vpc.NetworkArgs
    ///         {
    ///             CidrBlock = "172.16.0.0/12",
    ///         });
    ///         var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new AliCloud.Vpc.SwitchArgs
    ///         {
    ///             VpcId = defaultNetwork.Id,
    ///             CidrBlock = "172.16.0.0/24",
    ///             ZoneId = defaultZones.Apply(defaultZones =&gt; defaultZones.Zones[0].Id),
    ///         });
    ///         var defaultInstance = new AliCloud.AliKafka.Instance("defaultInstance", new AliCloud.AliKafka.InstanceArgs
    ///         {
    ///             TopicQuota = 50,
    ///             DiskType = 1,
    ///             DiskSize = 500,
    ///             DeployType = 5,
    ///             IoMax = 20,
    ///             VswitchId = defaultSwitch.Id,
    ///         });
    ///         var defaultSaslUser = new AliCloud.AliKafka.SaslUser("defaultSaslUser", new AliCloud.AliKafka.SaslUserArgs
    ///         {
    ///             InstanceId = defaultInstance.Id,
    ///             Username = username,
    ///             Password = password,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ALIKAFKA GROUP can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:alikafka/saslUser:SaslUser user alikafka_post-cn-123455abc:username
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:alikafka/saslUser:SaslUser")]
    public partial class SaslUser : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the ALIKAFKA Instance that owns the groups.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// An KMS encrypts password used to a db account. You have to specify one of `password` and `kms_encrypted_password` fields.
        /// </summary>
        [Output("kmsEncryptedPassword")]
        public Output<string?> KmsEncryptedPassword { get; private set; } = null!;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a user with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        /// </summary>
        [Output("kmsEncryptionContext")]
        public Output<ImmutableDictionary<string, object>?> KmsEncryptionContext { get; private set; } = null!;

        /// <summary>
        /// Operation password. It may consist of letters, digits, or underlines, with a length of 1 to 64 characters. You have to specify one of `password` and `kms_encrypted_password` fields.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Username for the sasl user. The length should between 1 to 64 characters. The characters can only contain 'a'-'z', 'A'-'Z', '0'-'9', '_' and '-'.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a SaslUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SaslUser(string name, SaslUserArgs args, CustomResourceOptions? options = null)
            : base("alicloud:alikafka/saslUser:SaslUser", name, args ?? new SaslUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SaslUser(string name, Input<string> id, SaslUserState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:alikafka/saslUser:SaslUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SaslUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SaslUser Get(string name, Input<string> id, SaslUserState? state = null, CustomResourceOptions? options = null)
        {
            return new SaslUser(name, id, state, options);
        }
    }

    public sealed class SaslUserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the ALIKAFKA Instance that owns the groups.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// An KMS encrypts password used to a db account. You have to specify one of `password` and `kms_encrypted_password` fields.
        /// </summary>
        [Input("kmsEncryptedPassword")]
        public Input<string>? KmsEncryptedPassword { get; set; }

        [Input("kmsEncryptionContext")]
        private InputMap<object>? _kmsEncryptionContext;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a user with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        /// </summary>
        public InputMap<object> KmsEncryptionContext
        {
            get => _kmsEncryptionContext ?? (_kmsEncryptionContext = new InputMap<object>());
            set => _kmsEncryptionContext = value;
        }

        /// <summary>
        /// Operation password. It may consist of letters, digits, or underlines, with a length of 1 to 64 characters. You have to specify one of `password` and `kms_encrypted_password` fields.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Username for the sasl user. The length should between 1 to 64 characters. The characters can only contain 'a'-'z', 'A'-'Z', '0'-'9', '_' and '-'.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SaslUserArgs()
        {
        }
    }

    public sealed class SaslUserState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the ALIKAFKA Instance that owns the groups.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// An KMS encrypts password used to a db account. You have to specify one of `password` and `kms_encrypted_password` fields.
        /// </summary>
        [Input("kmsEncryptedPassword")]
        public Input<string>? KmsEncryptedPassword { get; set; }

        [Input("kmsEncryptionContext")]
        private InputMap<object>? _kmsEncryptionContext;

        /// <summary>
        /// An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a user with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
        /// </summary>
        public InputMap<object> KmsEncryptionContext
        {
            get => _kmsEncryptionContext ?? (_kmsEncryptionContext = new InputMap<object>());
            set => _kmsEncryptionContext = value;
        }

        /// <summary>
        /// Operation password. It may consist of letters, digits, or underlines, with a length of 1 to 64 characters. You have to specify one of `password` and `kms_encrypted_password` fields.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Username for the sasl user. The length should between 1 to 64 characters. The characters can only contain 'a'-'z', 'A'-'Z', '0'-'9', '_' and '-'.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public SaslUserState()
        {
        }
    }
}
