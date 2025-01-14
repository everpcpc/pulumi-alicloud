// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a MongoDB sharding instance resource supports replica set instances only. the MongoDB provides stable, reliable, and automatic scalable database services.
 * It offers a full range of database solutions, such as disaster recovery, backup, recovery, monitoring, and alarms.
 * You can see detail product introduction [here](https://www.alibabacloud.com/help/doc-detail/26558.htm)
 *
 * > **NOTE:**  Available in 1.40.0+
 *
 * > **NOTE:**  The following regions don't support create Classic network MongoDB sharding instance.
 * [`cn-zhangjiakou`,`cn-huhehaote`,`ap-southeast-2`,`ap-southeast-3`,`ap-southeast-5`,`ap-south-1`,`me-east-1`,`ap-northeast-1`,`eu-west-1`]
 *
 * > **NOTE:**  Create MongoDB Sharding instance or change instance type and storage would cost 10~20 minutes. Please make full preparation
 *
 * ## Example Usage
 * ## Module Support
 *
 * You can use to the existing mongodb-sharding module
 * to create a MongoDB sharding instance resource one-click.
 *
 * ## Import
 *
 * MongoDB can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:mongodb/shardingInstance:ShardingInstance example dds-bp1291daeda44195
 * ```
 */
export class ShardingInstance extends pulumi.CustomResource {
    /**
     * Get an existing ShardingInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ShardingInstanceState, opts?: pulumi.CustomResourceOptions): ShardingInstance {
        return new ShardingInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:mongodb/shardingInstance:ShardingInstance';

    /**
     * Returns true if the given object is an instance of ShardingInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ShardingInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ShardingInstance.__pulumiType;
    }

    /**
     * Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
     */
    public readonly accountPassword!: pulumi.Output<string | undefined>;
    /**
     * Auto renew for prepaid, true of false. Default is false.
     */
    public readonly autoRenew!: pulumi.Output<boolean | undefined>;
    /**
     * MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
     */
    public readonly backupPeriods!: pulumi.Output<string[]>;
    /**
     * MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
     */
    public readonly backupTime!: pulumi.Output<string>;
    /**
     * The node information list of config server. The details see Block `configServerList`. **NOTE:** Available in v1.140+.
     */
    public /*out*/ readonly configServerLists!: pulumi.Output<outputs.mongodb.ShardingInstanceConfigServerList[]>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/en/doc-detail/61884.htm) `EngineVersion`.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version v1.141.0.
     */
    public readonly instanceChargeType!: pulumi.Output<string>;
    /**
     * An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
     */
    public readonly kmsEncryptedPassword!: pulumi.Output<string | undefined>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    public readonly kmsEncryptionContext!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The mongo-node count can be purchased is in range of [2, 32].
     */
    public readonly mongoLists!: pulumi.Output<outputs.mongodb.ShardingInstanceMongoList[]>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
     */
    public readonly networkType!: pulumi.Output<string>;
    /**
     * The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
     * * UPGRADE: The specifications are upgraded.
     * * DOWNGRADE: The specifications are downgraded.
     * Note: This parameter is only applicable to instances when `instanceChargeType` is PrePaid.
     */
    public readonly orderType!: pulumi.Output<string | undefined>;
    /**
     * The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
     */
    public readonly period!: pulumi.Output<number>;
    /**
     * The type of the access protocol. Valid values: `mongodb` or `dynamodb`.
     */
    public readonly protocolType!: pulumi.Output<string>;
    /**
     * The ID of the Resource Group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * Instance log backup retention days. **NOTE:** Available in 1.42.0+.
     */
    public /*out*/ readonly retentionPeriod!: pulumi.Output<number>;
    /**
     * The Security Group ID of ECS.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `["127.0.0.1"]`.
     */
    public readonly securityIpLists!: pulumi.Output<string[]>;
    /**
     * the shard-node count can be purchased is in range of [2, 32].
     */
    public readonly shardLists!: pulumi.Output<outputs.mongodb.ShardingInstanceShardList[]>;
    /**
     * Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
     */
    public readonly storageEngine!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The TDE(Transparent Data Encryption) status. It can be updated from version 1.160.0+.
     */
    public readonly tdeStatus!: pulumi.Output<string | undefined>;
    /**
     * The ID of the VPC. > **NOTE:** This parameter is valid only when NetworkType is set to VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;
    /**
     * The Zone to launch the DB instance. MongoDB sharding instance does not support multiple-zone.
     * If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
     */
    public readonly zoneId!: pulumi.Output<string | undefined>;

    /**
     * Create a ShardingInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ShardingInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ShardingInstanceArgs | ShardingInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ShardingInstanceState | undefined;
            resourceInputs["accountPassword"] = state ? state.accountPassword : undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["backupPeriods"] = state ? state.backupPeriods : undefined;
            resourceInputs["backupTime"] = state ? state.backupTime : undefined;
            resourceInputs["configServerLists"] = state ? state.configServerLists : undefined;
            resourceInputs["engineVersion"] = state ? state.engineVersion : undefined;
            resourceInputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            resourceInputs["kmsEncryptedPassword"] = state ? state.kmsEncryptedPassword : undefined;
            resourceInputs["kmsEncryptionContext"] = state ? state.kmsEncryptionContext : undefined;
            resourceInputs["mongoLists"] = state ? state.mongoLists : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkType"] = state ? state.networkType : undefined;
            resourceInputs["orderType"] = state ? state.orderType : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["protocolType"] = state ? state.protocolType : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["retentionPeriod"] = state ? state.retentionPeriod : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["securityIpLists"] = state ? state.securityIpLists : undefined;
            resourceInputs["shardLists"] = state ? state.shardLists : undefined;
            resourceInputs["storageEngine"] = state ? state.storageEngine : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tdeStatus"] = state ? state.tdeStatus : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ShardingInstanceArgs | undefined;
            if ((!args || args.engineVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineVersion'");
            }
            if ((!args || args.mongoLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mongoLists'");
            }
            if ((!args || args.shardLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shardLists'");
            }
            resourceInputs["accountPassword"] = args?.accountPassword ? pulumi.secret(args.accountPassword) : undefined;
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["backupPeriods"] = args ? args.backupPeriods : undefined;
            resourceInputs["backupTime"] = args ? args.backupTime : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            resourceInputs["kmsEncryptedPassword"] = args ? args.kmsEncryptedPassword : undefined;
            resourceInputs["kmsEncryptionContext"] = args ? args.kmsEncryptionContext : undefined;
            resourceInputs["mongoLists"] = args ? args.mongoLists : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkType"] = args ? args.networkType : undefined;
            resourceInputs["orderType"] = args ? args.orderType : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["protocolType"] = args ? args.protocolType : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["securityIpLists"] = args ? args.securityIpLists : undefined;
            resourceInputs["shardLists"] = args ? args.shardLists : undefined;
            resourceInputs["storageEngine"] = args ? args.storageEngine : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tdeStatus"] = args ? args.tdeStatus : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["configServerLists"] = undefined /*out*/;
            resourceInputs["retentionPeriod"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accountPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ShardingInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ShardingInstance resources.
 */
export interface ShardingInstanceState {
    /**
     * Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
     */
    accountPassword?: pulumi.Input<string>;
    /**
     * Auto renew for prepaid, true of false. Default is false.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
     */
    backupPeriods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
     */
    backupTime?: pulumi.Input<string>;
    /**
     * The node information list of config server. The details see Block `configServerList`. **NOTE:** Available in v1.140+.
     */
    configServerLists?: pulumi.Input<pulumi.Input<inputs.mongodb.ShardingInstanceConfigServerList>[]>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/en/doc-detail/61884.htm) `EngineVersion`.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version v1.141.0.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
     */
    kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * The mongo-node count can be purchased is in range of [2, 32].
     */
    mongoLists?: pulumi.Input<pulumi.Input<inputs.mongodb.ShardingInstanceMongoList>[]>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
     */
    networkType?: pulumi.Input<string>;
    /**
     * The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
     * * UPGRADE: The specifications are upgraded.
     * * DOWNGRADE: The specifications are downgraded.
     * Note: This parameter is only applicable to instances when `instanceChargeType` is PrePaid.
     */
    orderType?: pulumi.Input<string>;
    /**
     * The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
     */
    period?: pulumi.Input<number>;
    /**
     * The type of the access protocol. Valid values: `mongodb` or `dynamodb`.
     */
    protocolType?: pulumi.Input<string>;
    /**
     * The ID of the Resource Group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Instance log backup retention days. **NOTE:** Available in 1.42.0+.
     */
    retentionPeriod?: pulumi.Input<number>;
    /**
     * The Security Group ID of ECS.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `["127.0.0.1"]`.
     */
    securityIpLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * the shard-node count can be purchased is in range of [2, 32].
     */
    shardLists?: pulumi.Input<pulumi.Input<inputs.mongodb.ShardingInstanceShardList>[]>;
    /**
     * Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
     */
    storageEngine?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The TDE(Transparent Data Encryption) status. It can be updated from version 1.160.0+.
     */
    tdeStatus?: pulumi.Input<string>;
    /**
     * The ID of the VPC. > **NOTE:** This parameter is valid only when NetworkType is set to VPC.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the DB instance. MongoDB sharding instance does not support multiple-zone.
     * If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ShardingInstance resource.
 */
export interface ShardingInstanceArgs {
    /**
     * Password of the root account. It is a string of 6 to 32 characters and is composed of letters, numbers, and underlines.
     */
    accountPassword?: pulumi.Input<string>;
    /**
     * Auto renew for prepaid, true of false. Default is false.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * MongoDB Instance backup period. It is required when `backupTime` was existed. Valid values: [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]. Default to [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]
     */
    backupPeriods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * MongoDB instance backup time. It is required when `backupPeriod` was existed. In the format of HH:mmZ- HH:mmZ. Time setting interval is one hour. If not set, the system will return a default, like "23:00Z-24:00Z".
     */
    backupTime?: pulumi.Input<string>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/en/doc-detail/61884.htm) `EngineVersion`.
     */
    engineVersion: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`,System default to `PostPaid`. **NOTE:** It can be modified from `PostPaid` to `PrePaid` after version v1.141.0.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * An KMS encrypts password used to a instance. If the `accountPassword` is filled in, this field will be ignored.
     */
    kmsEncryptedPassword?: pulumi.Input<string>;
    /**
     * An KMS encryption context used to decrypt `kmsEncryptedPassword` before creating or updating instance with `kmsEncryptedPassword`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kmsEncryptedPassword` is set.
     */
    kmsEncryptionContext?: pulumi.Input<{[key: string]: any}>;
    /**
     * The mongo-node count can be purchased is in range of [2, 32].
     */
    mongoLists: pulumi.Input<pulumi.Input<inputs.mongodb.ShardingInstanceMongoList>[]>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * The network type of the instance. Valid values:`Classic` or `VPC`. Default value: `Classic`.
     */
    networkType?: pulumi.Input<string>;
    /**
     * The type of configuration changes performed. Default value: DOWNGRADE. Valid values:
     * * UPGRADE: The specifications are upgraded.
     * * DOWNGRADE: The specifications are downgraded.
     * Note: This parameter is only applicable to instances when `instanceChargeType` is PrePaid.
     */
    orderType?: pulumi.Input<string>;
    /**
     * The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. System default to 1.
     */
    period?: pulumi.Input<number>;
    /**
     * The type of the access protocol. Valid values: `mongodb` or `dynamodb`.
     */
    protocolType?: pulumi.Input<string>;
    /**
     * The ID of the Resource Group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The Security Group ID of ECS.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]). System default to `["127.0.0.1"]`.
     */
    securityIpLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * the shard-node count can be purchased is in range of [2, 32].
     */
    shardLists: pulumi.Input<pulumi.Input<inputs.mongodb.ShardingInstanceShardList>[]>;
    /**
     * Storage engine: WiredTiger or RocksDB. System Default value: WiredTiger.
     */
    storageEngine?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The TDE(Transparent Data Encryption) status. It can be updated from version 1.160.0+.
     */
    tdeStatus?: pulumi.Input<string>;
    /**
     * The ID of the VPC. > **NOTE:** This parameter is valid only when NetworkType is set to VPC.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The virtual switch ID to launch DB instances in one VPC.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the DB instance. MongoDB sharding instance does not support multiple-zone.
     * If it is a multi-zone and `vswitchId` is specified, the vswitch must in one of them.
     */
    zoneId?: pulumi.Input<string>;
}
