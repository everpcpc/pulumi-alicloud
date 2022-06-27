// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Hbr
{
    /// <summary>
    /// Provides a HBR Backup vault resource.
    /// 
    /// For information about HBR Backup vault and how to use it, see [What is Backup vault](https://www.alibabacloud.com/help/doc-detail/62362.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.129.0+.
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
    ///         var example = new AliCloud.Hbr.Vault("example", new AliCloud.Hbr.VaultArgs
    ///         {
    ///             VaultName = "example_value",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// HBR Vault can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:hbr/vault:Vault example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:hbr/vault:Vault")]
    public partial class Vault : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of Vault. Defaults to an empty string.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Source Encryption Type，It is valid only when vault_type is `STANDARD` or `OTS_BACKUP`. Valid values: `HBR_PRIVATE`,`KMS`. Defaults to `HBR_PRIVATE`.
        /// - `HBR_PRIVATE`: HBR is fully hosted, uses the backup service's own encryption method.
        /// - `KMS`: Use Alibaba Cloud Kms to encryption.
        /// </summary>
        [Output("encryptType")]
        public Output<string> EncryptType { get; private set; } = null!;

        /// <summary>
        /// The key id or alias name of Alibaba Cloud Kms. It is required and valid only when encrypt_type is `KMS`.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string?> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// The redundancy type of the vault. Valid values: `LRS`, and `ZRS`. Defaults to `LRS`.
        /// - `LRS`: means locally redundant storage, data will be stored on different storage devices in the same zone.
        /// - `ZRS`: means zone-redundant storage, the data will be replicated across three storage clusters in a single region. Each storage cluster is physically separated but within the same region.
        /// </summary>
        [Output("redundancyType")]
        public Output<string> RedundancyType { get; private set; } = null!;

        /// <summary>
        /// The status of Vault. Valid values: `INITIALIZING`, `CREATED`, `ERROR`, `UNKNOWN`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The name of Vault.
        /// </summary>
        [Output("vaultName")]
        public Output<string> VaultName { get; private set; } = null!;

        /// <summary>
        /// The storage class of Vault. Valid values: `STANDARD`.
        /// </summary>
        [Output("vaultStorageClass")]
        public Output<string> VaultStorageClass { get; private set; } = null!;

        /// <summary>
        /// The type of Vault. Valid values: `STANDARD`,`OTS_BACKUP`.
        /// </summary>
        [Output("vaultType")]
        public Output<string> VaultType { get; private set; } = null!;


        /// <summary>
        /// Create a Vault resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vault(string name, VaultArgs args, CustomResourceOptions? options = null)
            : base("alicloud:hbr/vault:Vault", name, args ?? new VaultArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Vault(string name, Input<string> id, VaultState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:hbr/vault:Vault", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Vault resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vault Get(string name, Input<string> id, VaultState? state = null, CustomResourceOptions? options = null)
        {
            return new Vault(name, id, state, options);
        }
    }

    public sealed class VaultArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of Vault. Defaults to an empty string.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Source Encryption Type，It is valid only when vault_type is `STANDARD` or `OTS_BACKUP`. Valid values: `HBR_PRIVATE`,`KMS`. Defaults to `HBR_PRIVATE`.
        /// - `HBR_PRIVATE`: HBR is fully hosted, uses the backup service's own encryption method.
        /// - `KMS`: Use Alibaba Cloud Kms to encryption.
        /// </summary>
        [Input("encryptType")]
        public Input<string>? EncryptType { get; set; }

        /// <summary>
        /// The key id or alias name of Alibaba Cloud Kms. It is required and valid only when encrypt_type is `KMS`.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// The redundancy type of the vault. Valid values: `LRS`, and `ZRS`. Defaults to `LRS`.
        /// - `LRS`: means locally redundant storage, data will be stored on different storage devices in the same zone.
        /// - `ZRS`: means zone-redundant storage, the data will be replicated across three storage clusters in a single region. Each storage cluster is physically separated but within the same region.
        /// </summary>
        [Input("redundancyType")]
        public Input<string>? RedundancyType { get; set; }

        /// <summary>
        /// The name of Vault.
        /// </summary>
        [Input("vaultName", required: true)]
        public Input<string> VaultName { get; set; } = null!;

        /// <summary>
        /// The storage class of Vault. Valid values: `STANDARD`.
        /// </summary>
        [Input("vaultStorageClass")]
        public Input<string>? VaultStorageClass { get; set; }

        /// <summary>
        /// The type of Vault. Valid values: `STANDARD`,`OTS_BACKUP`.
        /// </summary>
        [Input("vaultType")]
        public Input<string>? VaultType { get; set; }

        public VaultArgs()
        {
        }
    }

    public sealed class VaultState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of Vault. Defaults to an empty string.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Source Encryption Type，It is valid only when vault_type is `STANDARD` or `OTS_BACKUP`. Valid values: `HBR_PRIVATE`,`KMS`. Defaults to `HBR_PRIVATE`.
        /// - `HBR_PRIVATE`: HBR is fully hosted, uses the backup service's own encryption method.
        /// - `KMS`: Use Alibaba Cloud Kms to encryption.
        /// </summary>
        [Input("encryptType")]
        public Input<string>? EncryptType { get; set; }

        /// <summary>
        /// The key id or alias name of Alibaba Cloud Kms. It is required and valid only when encrypt_type is `KMS`.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// The redundancy type of the vault. Valid values: `LRS`, and `ZRS`. Defaults to `LRS`.
        /// - `LRS`: means locally redundant storage, data will be stored on different storage devices in the same zone.
        /// - `ZRS`: means zone-redundant storage, the data will be replicated across three storage clusters in a single region. Each storage cluster is physically separated but within the same region.
        /// </summary>
        [Input("redundancyType")]
        public Input<string>? RedundancyType { get; set; }

        /// <summary>
        /// The status of Vault. Valid values: `INITIALIZING`, `CREATED`, `ERROR`, `UNKNOWN`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The name of Vault.
        /// </summary>
        [Input("vaultName")]
        public Input<string>? VaultName { get; set; }

        /// <summary>
        /// The storage class of Vault. Valid values: `STANDARD`.
        /// </summary>
        [Input("vaultStorageClass")]
        public Input<string>? VaultStorageClass { get; set; }

        /// <summary>
        /// The type of Vault. Valid values: `STANDARD`,`OTS_BACKUP`.
        /// </summary>
        [Input("vaultType")]
        public Input<string>? VaultType { get; set; }

        public VaultState()
        {
        }
    }
}
