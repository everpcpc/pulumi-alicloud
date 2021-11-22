// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an RDS readonly instance resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const creation = config.get("creation") || "Rds";
 * const name = config.get("name") || "dbInstancevpc";
 * const defaultZones = alicloud.getZones({
 *     availableResourceCreation: creation,
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {cidrBlock: "172.16.0.0/16"});
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?[0]?.id),
 *     vswitchName: name,
 * });
 * const defaultInstance = new alicloud.rds.Instance("defaultInstance", {
 *     engine: "MySQL",
 *     engineVersion: "5.6",
 *     instanceType: "rds.mysql.t1.small",
 *     instanceStorage: "20",
 *     instanceChargeType: "Postpaid",
 *     instanceName: name,
 *     vswitchId: defaultSwitch.id,
 *     securityIps: [
 *         "10.168.1.12",
 *         "100.69.7.112",
 *     ],
 * });
 * const defaultReadOnlyInstance = new alicloud.rds.ReadOnlyInstance("defaultReadOnlyInstance", {
 *     masterDbInstanceId: defaultInstance.id,
 *     zoneId: defaultInstance.zoneId,
 *     engineVersion: defaultInstance.engineVersion,
 *     instanceType: defaultInstance.instanceType,
 *     instanceStorage: "30",
 *     instanceName: `${name}ro`,
 *     vswitchId: defaultSwitch.id,
 * });
 * ```
 *
 * ## Import
 *
 * RDS readonly instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:rds/readOnlyInstance:ReadOnlyInstance example rm-abc12345678
 * ```
 */
export class ReadOnlyInstance extends pulumi.CustomResource {
    /**
     * Get an existing ReadOnlyInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReadOnlyInstanceState, opts?: pulumi.CustomResourceOptions): ReadOnlyInstance {
        return new ReadOnlyInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:rds/readOnlyInstance:ReadOnlyInstance';

    /**
     * Returns true if the given object is an instance of ReadOnlyInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReadOnlyInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReadOnlyInstance.__pulumiType;
    }

    /**
     * The method that is used to verify the identities of clients. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. In addition, this parameter is available only when the public key of the CA that issues client certificates is enabled. It is valid only when `sslEnabled  = 1`. Valid values:
     * - cert
     * - perfer
     * - verify-ca
     * - verify-full (supported only when the instance runs PostgreSQL 12 or later)
     */
    public readonly acl!: pulumi.Output<string>;
    /**
     * The type of the server certificate. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. If you set the SSLEnabled parameter to 1, the default value of this parameter is aliyun. It is valid only when `sslEnabled  = 1`. Value range:
     * - aliyun: a cloud certificate
     * - custom: a custom certificate
     */
    public readonly caType!: pulumi.Output<string>;
    /**
     * The public key of the CA that issues client certificates. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. If you set the ClientCAEbabled parameter to 1, you must also specify this parameter. It is valid only when `sslEnabled  = 1`.
     */
    public readonly clientCaCert!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to enable the public key of the CA that issues client certificates. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. It is valid only when `sslEnabled  = 1`. Valid values:
     * - 1: enables the public key
     * - 0: disables the public key
     */
    public readonly clientCaEnabled!: pulumi.Output<number | undefined>;
    /**
     * The CRL that contains revoked client certificates. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. If you set the ClientCrlEnabled parameter to 1, you must also specify this parameter. It is valid only when `sslEnabled  = 1`.
     */
    public readonly clientCertRevocationList!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether to enable a certificate revocation list (CRL) that contains revoked client certificates. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. In addition, this parameter is available only when the public key of the CA that issues client certificates is enabled. It is valid only when `sslEnabled  = 1`. Valid values:
     * - 1: enables the CRL
     * - 0: disables the CRL
     */
    public readonly clientCrlEnabled!: pulumi.Output<number | undefined>;
    /**
     * RDS database connection string.
     */
    public /*out*/ readonly connectionString!: pulumi.Output<string>;
    /**
     * Database type.
     */
    public /*out*/ readonly engine!: pulumi.Output<string>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * Set it to true to make some parameter efficient when modifying them. Default to false.
     */
    public readonly forceRestart!: pulumi.Output<boolean | undefined>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * User-defined DB instance storage space. Value range: [5, 2000] for MySQL/SQL Server HA dual node edition. Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     */
    public readonly instanceStorage!: pulumi.Output<number>;
    /**
     * DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * ID of the master instance.
     */
    public readonly masterDbInstanceId!: pulumi.Output<string>;
    /**
     * Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm).
     */
    public readonly parameters!: pulumi.Output<outputs.rds.ReadOnlyInstanceParameter[]>;
    /**
     * RDS database connection port.
     */
    public /*out*/ readonly port!: pulumi.Output<string>;
    /**
     * The method that is used to verify the replication permission. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. In addition, this parameter is available only when the public key of the CA that issues client certificates is enabled. It is valid only when `sslEnabled  = 1`. Valid values:
     * - cert
     * - perfer
     * - verify-ca
     * - verify-full (supported only when the instance runs PostgreSQL 12 or later)
     * > **NOTE:** Because of data backup and migration, change DB instance type and storage would cost 15~20 minutes. Please make full preparation before changing them.
     */
    public readonly replicationAcl!: pulumi.Output<string>;
    /**
     * The ID of resource group which the DB read-only instance belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The content of the server certificate. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. If you set the CAType parameter to custom, you must also specify this parameter. It is valid only when `sslEnabled  = 1`.
     */
    public readonly serverCert!: pulumi.Output<string>;
    /**
     * The private key of the server certificate. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. If you set the CAType parameter to custom, you must also specify this parameter. It is valid only when `sslEnabled  = 1`.
     */
    public readonly serverKey!: pulumi.Output<string>;
    /**
     * Specifies whether to enable or disable SSL encryption. Valid values:
     * - 1: enables SSL encryption
     * - 0: disables SSL encryption
     */
    public readonly sslEnabled!: pulumi.Output<number>;
    /**
     * The specific point in time when you want to perform the update. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. It is valid only when `upgradeDbInstanceKernelVersion = true`. The time must be in UTC.
     */
    public readonly switchTime!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The minor engine version to which you want to update the instance. If you do not specify this parameter, the instance is updated to the latest minor engine version. It is valid only when `upgradeDbInstanceKernelVersion = true`. You must specify the minor engine version in one of the following formats:
     * - PostgreSQL: rds_postgres_<Major engine version>00_<Minor engine version>. Example: rds_postgres_1200_20200830.
     * - MySQL: <RDS edition>_<Minor engine version>. Examples: rds_20200229, xcluster_20200229, and xcluster80_20200229. The following RDS editions are supported:
     * - rds: The instance runs RDS Basic or High-availability Edition.
     * - xcluster: The instance runs MySQL 5.7 on RDS Enterprise Edition.
     * - xcluster80: The instance runs MySQL 8.0 on RDS Enterprise Edition.
     * - SQLServer: <Minor engine version>. Example: 15.0.4073.23.
     */
    public readonly targetMinorVersion!: pulumi.Output<string>;
    /**
     * Whether to upgrade a minor version of the kernel. Valid values:
     * - true: upgrade
     * - false: not to upgrade
     */
    public readonly upgradeDbInstanceKernelVersion!: pulumi.Output<boolean | undefined>;
    /**
     * The method to update the minor engine version. Default value: Immediate. It is valid only when `upgradeDbInstanceKernelVersion = true`. Valid values:
     * - Immediate: The minor engine version is immediately updated.
     * - MaintainTime: The minor engine version is updated during the maintenance window. For more information about how to change the maintenance window, see ModifyDBInstanceMaintainTime.
     * - SpecifyTime: The minor engine version is updated at the point in time you specify.
     */
    public readonly upgradeTime!: pulumi.Output<string | undefined>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;
    /**
     * The Zone to launch the DB instance.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a ReadOnlyInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReadOnlyInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReadOnlyInstanceArgs | ReadOnlyInstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReadOnlyInstanceState | undefined;
            inputs["acl"] = state ? state.acl : undefined;
            inputs["caType"] = state ? state.caType : undefined;
            inputs["clientCaCert"] = state ? state.clientCaCert : undefined;
            inputs["clientCaEnabled"] = state ? state.clientCaEnabled : undefined;
            inputs["clientCertRevocationList"] = state ? state.clientCertRevocationList : undefined;
            inputs["clientCrlEnabled"] = state ? state.clientCrlEnabled : undefined;
            inputs["connectionString"] = state ? state.connectionString : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["forceRestart"] = state ? state.forceRestart : undefined;
            inputs["instanceName"] = state ? state.instanceName : undefined;
            inputs["instanceStorage"] = state ? state.instanceStorage : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["masterDbInstanceId"] = state ? state.masterDbInstanceId : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["replicationAcl"] = state ? state.replicationAcl : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["serverCert"] = state ? state.serverCert : undefined;
            inputs["serverKey"] = state ? state.serverKey : undefined;
            inputs["sslEnabled"] = state ? state.sslEnabled : undefined;
            inputs["switchTime"] = state ? state.switchTime : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["targetMinorVersion"] = state ? state.targetMinorVersion : undefined;
            inputs["upgradeDbInstanceKernelVersion"] = state ? state.upgradeDbInstanceKernelVersion : undefined;
            inputs["upgradeTime"] = state ? state.upgradeTime : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ReadOnlyInstanceArgs | undefined;
            if ((!args || args.engineVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineVersion'");
            }
            if ((!args || args.instanceStorage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceStorage'");
            }
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            if ((!args || args.masterDbInstanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'masterDbInstanceId'");
            }
            inputs["acl"] = args ? args.acl : undefined;
            inputs["caType"] = args ? args.caType : undefined;
            inputs["clientCaCert"] = args ? args.clientCaCert : undefined;
            inputs["clientCaEnabled"] = args ? args.clientCaEnabled : undefined;
            inputs["clientCertRevocationList"] = args ? args.clientCertRevocationList : undefined;
            inputs["clientCrlEnabled"] = args ? args.clientCrlEnabled : undefined;
            inputs["engineVersion"] = args ? args.engineVersion : undefined;
            inputs["forceRestart"] = args ? args.forceRestart : undefined;
            inputs["instanceName"] = args ? args.instanceName : undefined;
            inputs["instanceStorage"] = args ? args.instanceStorage : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["masterDbInstanceId"] = args ? args.masterDbInstanceId : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["replicationAcl"] = args ? args.replicationAcl : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["serverCert"] = args ? args.serverCert : undefined;
            inputs["serverKey"] = args ? args.serverKey : undefined;
            inputs["sslEnabled"] = args ? args.sslEnabled : undefined;
            inputs["switchTime"] = args ? args.switchTime : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetMinorVersion"] = args ? args.targetMinorVersion : undefined;
            inputs["upgradeDbInstanceKernelVersion"] = args ? args.upgradeDbInstanceKernelVersion : undefined;
            inputs["upgradeTime"] = args ? args.upgradeTime : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
            inputs["connectionString"] = undefined /*out*/;
            inputs["engine"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ReadOnlyInstance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReadOnlyInstance resources.
 */
export interface ReadOnlyInstanceState {
    /**
     * The method that is used to verify the identities of clients. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. In addition, this parameter is available only when the public key of the CA that issues client certificates is enabled. It is valid only when `sslEnabled  = 1`. Valid values:
     * - cert
     * - perfer
     * - verify-ca
     * - verify-full (supported only when the instance runs PostgreSQL 12 or later)
     */
    acl?: pulumi.Input<string>;
    /**
     * The type of the server certificate. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. If you set the SSLEnabled parameter to 1, the default value of this parameter is aliyun. It is valid only when `sslEnabled  = 1`. Value range:
     * - aliyun: a cloud certificate
     * - custom: a custom certificate
     */
    caType?: pulumi.Input<string>;
    /**
     * The public key of the CA that issues client certificates. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. If you set the ClientCAEbabled parameter to 1, you must also specify this parameter. It is valid only when `sslEnabled  = 1`.
     */
    clientCaCert?: pulumi.Input<string>;
    /**
     * Specifies whether to enable the public key of the CA that issues client certificates. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. It is valid only when `sslEnabled  = 1`. Valid values:
     * - 1: enables the public key
     * - 0: disables the public key
     */
    clientCaEnabled?: pulumi.Input<number>;
    /**
     * The CRL that contains revoked client certificates. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. If you set the ClientCrlEnabled parameter to 1, you must also specify this parameter. It is valid only when `sslEnabled  = 1`.
     */
    clientCertRevocationList?: pulumi.Input<string>;
    /**
     * Specifies whether to enable a certificate revocation list (CRL) that contains revoked client certificates. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. In addition, this parameter is available only when the public key of the CA that issues client certificates is enabled. It is valid only when `sslEnabled  = 1`. Valid values:
     * - 1: enables the CRL
     * - 0: disables the CRL
     */
    clientCrlEnabled?: pulumi.Input<number>;
    /**
     * RDS database connection string.
     */
    connectionString?: pulumi.Input<string>;
    /**
     * Database type.
     */
    engine?: pulumi.Input<string>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Set it to true to make some parameter efficient when modifying them. Default to false.
     */
    forceRestart?: pulumi.Input<boolean>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * User-defined DB instance storage space. Value range: [5, 2000] for MySQL/SQL Server HA dual node edition. Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     */
    instanceStorage?: pulumi.Input<number>;
    /**
     * DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     */
    instanceType?: pulumi.Input<string>;
    /**
     * ID of the master instance.
     */
    masterDbInstanceId?: pulumi.Input<string>;
    /**
     * Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm).
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.rds.ReadOnlyInstanceParameter>[]>;
    /**
     * RDS database connection port.
     */
    port?: pulumi.Input<string>;
    /**
     * The method that is used to verify the replication permission. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. In addition, this parameter is available only when the public key of the CA that issues client certificates is enabled. It is valid only when `sslEnabled  = 1`. Valid values:
     * - cert
     * - perfer
     * - verify-ca
     * - verify-full (supported only when the instance runs PostgreSQL 12 or later)
     * > **NOTE:** Because of data backup and migration, change DB instance type and storage would cost 15~20 minutes. Please make full preparation before changing them.
     */
    replicationAcl?: pulumi.Input<string>;
    /**
     * The ID of resource group which the DB read-only instance belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The content of the server certificate. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. If you set the CAType parameter to custom, you must also specify this parameter. It is valid only when `sslEnabled  = 1`.
     */
    serverCert?: pulumi.Input<string>;
    /**
     * The private key of the server certificate. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. If you set the CAType parameter to custom, you must also specify this parameter. It is valid only when `sslEnabled  = 1`.
     */
    serverKey?: pulumi.Input<string>;
    /**
     * Specifies whether to enable or disable SSL encryption. Valid values:
     * - 1: enables SSL encryption
     * - 0: disables SSL encryption
     */
    sslEnabled?: pulumi.Input<number>;
    /**
     * The specific point in time when you want to perform the update. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. It is valid only when `upgradeDbInstanceKernelVersion = true`. The time must be in UTC.
     */
    switchTime?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The minor engine version to which you want to update the instance. If you do not specify this parameter, the instance is updated to the latest minor engine version. It is valid only when `upgradeDbInstanceKernelVersion = true`. You must specify the minor engine version in one of the following formats:
     * - PostgreSQL: rds_postgres_<Major engine version>00_<Minor engine version>. Example: rds_postgres_1200_20200830.
     * - MySQL: <RDS edition>_<Minor engine version>. Examples: rds_20200229, xcluster_20200229, and xcluster80_20200229. The following RDS editions are supported:
     * - rds: The instance runs RDS Basic or High-availability Edition.
     * - xcluster: The instance runs MySQL 5.7 on RDS Enterprise Edition.
     * - xcluster80: The instance runs MySQL 8.0 on RDS Enterprise Edition.
     * - SQLServer: <Minor engine version>. Example: 15.0.4073.23.
     */
    targetMinorVersion?: pulumi.Input<string>;
    /**
     * Whether to upgrade a minor version of the kernel. Valid values:
     * - true: upgrade
     * - false: not to upgrade
     */
    upgradeDbInstanceKernelVersion?: pulumi.Input<boolean>;
    /**
     * The method to update the minor engine version. Default value: Immediate. It is valid only when `upgradeDbInstanceKernelVersion = true`. Valid values:
     * - Immediate: The minor engine version is immediately updated.
     * - MaintainTime: The minor engine version is updated during the maintenance window. For more information about how to change the maintenance window, see ModifyDBInstanceMaintainTime.
     * - SpecifyTime: The minor engine version is updated at the point in time you specify.
     */
    upgradeTime?: pulumi.Input<string>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the DB instance.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReadOnlyInstance resource.
 */
export interface ReadOnlyInstanceArgs {
    /**
     * The method that is used to verify the identities of clients. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. In addition, this parameter is available only when the public key of the CA that issues client certificates is enabled. It is valid only when `sslEnabled  = 1`. Valid values:
     * - cert
     * - perfer
     * - verify-ca
     * - verify-full (supported only when the instance runs PostgreSQL 12 or later)
     */
    acl?: pulumi.Input<string>;
    /**
     * The type of the server certificate. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. If you set the SSLEnabled parameter to 1, the default value of this parameter is aliyun. It is valid only when `sslEnabled  = 1`. Value range:
     * - aliyun: a cloud certificate
     * - custom: a custom certificate
     */
    caType?: pulumi.Input<string>;
    /**
     * The public key of the CA that issues client certificates. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. If you set the ClientCAEbabled parameter to 1, you must also specify this parameter. It is valid only when `sslEnabled  = 1`.
     */
    clientCaCert?: pulumi.Input<string>;
    /**
     * Specifies whether to enable the public key of the CA that issues client certificates. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. It is valid only when `sslEnabled  = 1`. Valid values:
     * - 1: enables the public key
     * - 0: disables the public key
     */
    clientCaEnabled?: pulumi.Input<number>;
    /**
     * The CRL that contains revoked client certificates. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. If you set the ClientCrlEnabled parameter to 1, you must also specify this parameter. It is valid only when `sslEnabled  = 1`.
     */
    clientCertRevocationList?: pulumi.Input<string>;
    /**
     * Specifies whether to enable a certificate revocation list (CRL) that contains revoked client certificates. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. In addition, this parameter is available only when the public key of the CA that issues client certificates is enabled. It is valid only when `sslEnabled  = 1`. Valid values:
     * - 1: enables the CRL
     * - 0: disables the CRL
     */
    clientCrlEnabled?: pulumi.Input<number>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
     */
    engineVersion: pulumi.Input<string>;
    /**
     * Set it to true to make some parameter efficient when modifying them. Default to false.
     */
    forceRestart?: pulumi.Input<boolean>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * User-defined DB instance storage space. Value range: [5, 2000] for MySQL/SQL Server HA dual node edition. Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     */
    instanceStorage: pulumi.Input<number>;
    /**
     * DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     */
    instanceType: pulumi.Input<string>;
    /**
     * ID of the master instance.
     */
    masterDbInstanceId: pulumi.Input<string>;
    /**
     * Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm).
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.rds.ReadOnlyInstanceParameter>[]>;
    /**
     * The method that is used to verify the replication permission. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. In addition, this parameter is available only when the public key of the CA that issues client certificates is enabled. It is valid only when `sslEnabled  = 1`. Valid values:
     * - cert
     * - perfer
     * - verify-ca
     * - verify-full (supported only when the instance runs PostgreSQL 12 or later)
     * > **NOTE:** Because of data backup and migration, change DB instance type and storage would cost 15~20 minutes. Please make full preparation before changing them.
     */
    replicationAcl?: pulumi.Input<string>;
    /**
     * The ID of resource group which the DB read-only instance belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The content of the server certificate. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. If you set the CAType parameter to custom, you must also specify this parameter. It is valid only when `sslEnabled  = 1`.
     */
    serverCert?: pulumi.Input<string>;
    /**
     * The private key of the server certificate. This parameter is supported only when the instance runs PostgreSQL with standard or enhanced SSDs. If you set the CAType parameter to custom, you must also specify this parameter. It is valid only when `sslEnabled  = 1`.
     */
    serverKey?: pulumi.Input<string>;
    /**
     * Specifies whether to enable or disable SSL encryption. Valid values:
     * - 1: enables SSL encryption
     * - 0: disables SSL encryption
     */
    sslEnabled?: pulumi.Input<number>;
    /**
     * The specific point in time when you want to perform the update. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. It is valid only when `upgradeDbInstanceKernelVersion = true`. The time must be in UTC.
     */
    switchTime?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The minor engine version to which you want to update the instance. If you do not specify this parameter, the instance is updated to the latest minor engine version. It is valid only when `upgradeDbInstanceKernelVersion = true`. You must specify the minor engine version in one of the following formats:
     * - PostgreSQL: rds_postgres_<Major engine version>00_<Minor engine version>. Example: rds_postgres_1200_20200830.
     * - MySQL: <RDS edition>_<Minor engine version>. Examples: rds_20200229, xcluster_20200229, and xcluster80_20200229. The following RDS editions are supported:
     * - rds: The instance runs RDS Basic or High-availability Edition.
     * - xcluster: The instance runs MySQL 5.7 on RDS Enterprise Edition.
     * - xcluster80: The instance runs MySQL 8.0 on RDS Enterprise Edition.
     * - SQLServer: <Minor engine version>. Example: 15.0.4073.23.
     */
    targetMinorVersion?: pulumi.Input<string>;
    /**
     * Whether to upgrade a minor version of the kernel. Valid values:
     * - true: upgrade
     * - false: not to upgrade
     */
    upgradeDbInstanceKernelVersion?: pulumi.Input<boolean>;
    /**
     * The method to update the minor engine version. Default value: Immediate. It is valid only when `upgradeDbInstanceKernelVersion = true`. Valid values:
     * - Immediate: The minor engine version is immediately updated.
     * - MaintainTime: The minor engine version is updated during the maintenance window. For more information about how to change the maintenance window, see ModifyDBInstanceMaintainTime.
     * - SpecifyTime: The minor engine version is updated at the point in time you specify.
     */
    upgradeTime?: pulumi.Input<string>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the DB instance.
     */
    zoneId?: pulumi.Input<string>;
}
