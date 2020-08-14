// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an RDS instance resource. A DB instance is an isolated database
 * environment in the cloud. A DB instance can contain multiple user-created
 * databases.
 *
 * ## Example Usage
 *
 * ### Create a RDS MySQL instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "dbInstanceconfig";
 * const creation = config.get("creation") || "Rds";
 *
 * const defaultZones = pulumi.output(alicloud.getZones({
 *     availableResourceCreation: creation,
 * }, { async: true }));
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     availabilityZone: defaultZones.zones[0].id,
 *     cidrBlock: "172.16.0.0/24",
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultInstance = new alicloud.rds.Instance("default", {
 *     engine: "MySQL",
 *     engineVersion: "5.6",
 *     instanceChargeType: "Postpaid",
 *     instanceName: name,
 *     instanceStorage: 30,
 *     instanceType: "rds.mysql.s2.large",
 *     monitoringPeriod: 60,
 *     vswitchId: defaultSwitch.id,
 * });
 * ```
 *
 * ### Create a RDS MySQL instance with specific parameters
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultNetwork = new alicloud.vpc.Network("default", {
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     availabilityZone: alicloud_zones_default.zones.0.id,
 *     cidrBlock: "172.16.0.0/24",
 *     vpcId: defaultNetwork.id,
 * });
 * const defaultInstance = new alicloud.rds.Instance("default", {
 *     dbInstanceClass: "rds.mysql.t1.small",
 *     dbInstanceStorage: "10",
 *     engine: "MySQL",
 *     engineVersion: "5.6",
 *     parameters: [
 *         {
 *             name: "innodbLargePrefix",
 *             value: "ON",
 *         },
 *         {
 *             name: "connectTimeout",
 *             value: "50",
 *         },
 *     ],
 * });
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:rds/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Whether to renewal a DB instance automatically or not. It is valid when instanceChargeType is `PrePaid`. Default to `false`.
     */
    public readonly autoRenew!: pulumi.Output<boolean | undefined>;
    /**
     * Auto-renewal period of an instance, in the unit of the month. It is valid when instanceChargeType is `PrePaid`. Valid value:[1~12], Default to 1.
     */
    public readonly autoRenewPeriod!: pulumi.Output<number | undefined>;
    /**
     * The upgrade method to use. Valid values:
     * - Auto: Instances are automatically upgraded to a higher minor version.
     * - Manual: Instances are forcibly upgraded to a higher minor version when the current version is unpublished.
     */
    public readonly autoUpgradeMinorVersion!: pulumi.Output<string>;
    /**
     * RDS database connection string.
     */
    public /*out*/ readonly connectionString!: pulumi.Output<string>;
    /**
     * The storage type of the instance. Valid values:
     * - local_ssd: specifies to use local SSDs. This value is recommended.
     * - cloud_ssd: specifies to use standard SSDs.
     * - cloud_essd: specifies to use enhanced SSDs (ESSDs).
     * - cloud_essd2: specifies to use enhanced SSDs (ESSDs).
     * - cloud_essd3: specifies to use enhanced SSDs (ESSDs).
     */
    public readonly dbInstanceStorageType!: pulumi.Output<string>;
    /**
     * Database type. Value options: MySQL, SQLServer, PostgreSQL, and PPAS.
     */
    public readonly engine!: pulumi.Output<string>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * Set it to true to make some parameter efficient when modifying them. Default to false.
     */
    public readonly forceRestart!: pulumi.Output<boolean | undefined>;
    /**
     * Valid values are `Prepaid`, `Postpaid`, Default to `Postpaid`. Currently, the resource only supports PostPaid to PrePaid.
     */
    public readonly instanceChargeType!: pulumi.Output<string | undefined>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    public readonly instanceName!: pulumi.Output<string | undefined>;
    /**
     * User-defined DB instance storage space. Value range:
     * - [5, 2000] for MySQL/PostgreSQL/PPAS HA dual node edition;
     * - [20,1000] for MySQL 5.7 basic single node edition;
     * - [10, 2000] for SQL Server 2008R2;
     * - [20,2000] for SQL Server 2012 basic single node edition
     * Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     * Note: There is extra 5 GB storage for SQL Server Instance and it is not in specified `instanceStorage`.
     */
    public readonly instanceStorage!: pulumi.Output<number>;
    /**
     * DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
     */
    public readonly maintainTime!: pulumi.Output<string>;
    /**
     * The monitoring frequency in seconds. Valid values are 5, 60, 300. Defaults to 300. 
     */
    public readonly monitoringPeriod!: pulumi.Output<number>;
    /**
     * Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
     */
    public readonly parameters!: pulumi.Output<outputs.rds.InstanceParameter[]>;
    /**
     * The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * RDS database connection port.
     */
    public /*out*/ readonly port!: pulumi.Output<string>;
    /**
     * The ID of resource group which the DB instance belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * It has been deprecated from 1.69.0 and use `securityGroupIds` instead.
     *
     * @deprecated Attribute `security_group_id` has been deprecated from 1.69.0 and use `security_group_ids` instead.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * , Available in 1.69.0+) The list IDs to join ECS Security Group. At most supports three security groups.
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * Valid values are `normal`, `safety`, Default to `normal`. support `safety` switch to high security access mode 
     */
    public readonly securityIpMode!: pulumi.Output<string | undefined>;
    /**
     * List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     */
    public readonly securityIps!: pulumi.Output<string[]>;
    /**
     * The sql collector keep time of the instance. Valid values are `30`, `180`, `365`, `1095`, `1825`, Default to `30`.
     */
    public readonly sqlCollectorConfigValue!: pulumi.Output<number | undefined>;
    /**
     * The sql collector status of the instance. Valid values are `Enabled`, `Disabled`, Default to `Disabled`.
     */
    public readonly sqlCollectorStatus!: pulumi.Output<string>;
    /**
     * Actions performed on SSL functions, Valid values: `Open`: turn on SSL encryption; `Close`: turn off SSL encryption; `Update`: update SSL certificate. See more [engine and engineVersion limitation](https://www.alibabacloud.com/help/zh/doc-detail/26254.htm).
     */
    public readonly sslAction!: pulumi.Output<string>;
    /**
     * Status of the SSL feature. `Yes`: SSL is turned on; `No`: SSL is turned off.
     */
    public /*out*/ readonly sslStatus!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The TDE(Transparent Data Encryption) status. See more [engine and engineVersion limitation](https://www.alibabacloud.com/help/zh/doc-detail/26256.htm).
     */
    public readonly tdeStatus!: pulumi.Output<string | undefined>;
    /**
     * The virtual switch ID to launch DB instances in one VPC. If there are multiple vswitches, separate them with commas.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;
    /**
     * The Zone to launch the DB instance. From version 1.8.1, it supports multiple zone.
     * If it is a multi-zone and `vswitchId` is specified, the vswitch must in the one of them.
     * The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `alicloud..getZones`.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceState | undefined;
            inputs["autoRenew"] = state ? state.autoRenew : undefined;
            inputs["autoRenewPeriod"] = state ? state.autoRenewPeriod : undefined;
            inputs["autoUpgradeMinorVersion"] = state ? state.autoUpgradeMinorVersion : undefined;
            inputs["connectionString"] = state ? state.connectionString : undefined;
            inputs["dbInstanceStorageType"] = state ? state.dbInstanceStorageType : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["forceRestart"] = state ? state.forceRestart : undefined;
            inputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            inputs["instanceName"] = state ? state.instanceName : undefined;
            inputs["instanceStorage"] = state ? state.instanceStorage : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["maintainTime"] = state ? state.maintainTime : undefined;
            inputs["monitoringPeriod"] = state ? state.monitoringPeriod : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["securityIpMode"] = state ? state.securityIpMode : undefined;
            inputs["securityIps"] = state ? state.securityIps : undefined;
            inputs["sqlCollectorConfigValue"] = state ? state.sqlCollectorConfigValue : undefined;
            inputs["sqlCollectorStatus"] = state ? state.sqlCollectorStatus : undefined;
            inputs["sslAction"] = state ? state.sslAction : undefined;
            inputs["sslStatus"] = state ? state.sslStatus : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tdeStatus"] = state ? state.tdeStatus : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if (!args || args.engine === undefined) {
                throw new Error("Missing required property 'engine'");
            }
            if (!args || args.engineVersion === undefined) {
                throw new Error("Missing required property 'engineVersion'");
            }
            if (!args || args.instanceStorage === undefined) {
                throw new Error("Missing required property 'instanceStorage'");
            }
            if (!args || args.instanceType === undefined) {
                throw new Error("Missing required property 'instanceType'");
            }
            inputs["autoRenew"] = args ? args.autoRenew : undefined;
            inputs["autoRenewPeriod"] = args ? args.autoRenewPeriod : undefined;
            inputs["autoUpgradeMinorVersion"] = args ? args.autoUpgradeMinorVersion : undefined;
            inputs["dbInstanceStorageType"] = args ? args.dbInstanceStorageType : undefined;
            inputs["engine"] = args ? args.engine : undefined;
            inputs["engineVersion"] = args ? args.engineVersion : undefined;
            inputs["forceRestart"] = args ? args.forceRestart : undefined;
            inputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            inputs["instanceName"] = args ? args.instanceName : undefined;
            inputs["instanceStorage"] = args ? args.instanceStorage : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["maintainTime"] = args ? args.maintainTime : undefined;
            inputs["monitoringPeriod"] = args ? args.monitoringPeriod : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            inputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["securityIpMode"] = args ? args.securityIpMode : undefined;
            inputs["securityIps"] = args ? args.securityIps : undefined;
            inputs["sqlCollectorConfigValue"] = args ? args.sqlCollectorConfigValue : undefined;
            inputs["sqlCollectorStatus"] = args ? args.sqlCollectorStatus : undefined;
            inputs["sslAction"] = args ? args.sslAction : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tdeStatus"] = args ? args.tdeStatus : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
            inputs["connectionString"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["sslStatus"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Whether to renewal a DB instance automatically or not. It is valid when instanceChargeType is `PrePaid`. Default to `false`.
     */
    readonly autoRenew?: pulumi.Input<boolean>;
    /**
     * Auto-renewal period of an instance, in the unit of the month. It is valid when instanceChargeType is `PrePaid`. Valid value:[1~12], Default to 1.
     */
    readonly autoRenewPeriod?: pulumi.Input<number>;
    /**
     * The upgrade method to use. Valid values:
     * - Auto: Instances are automatically upgraded to a higher minor version.
     * - Manual: Instances are forcibly upgraded to a higher minor version when the current version is unpublished.
     */
    readonly autoUpgradeMinorVersion?: pulumi.Input<string>;
    /**
     * RDS database connection string.
     */
    readonly connectionString?: pulumi.Input<string>;
    /**
     * The storage type of the instance. Valid values:
     * - local_ssd: specifies to use local SSDs. This value is recommended.
     * - cloud_ssd: specifies to use standard SSDs.
     * - cloud_essd: specifies to use enhanced SSDs (ESSDs).
     * - cloud_essd2: specifies to use enhanced SSDs (ESSDs).
     * - cloud_essd3: specifies to use enhanced SSDs (ESSDs).
     */
    readonly dbInstanceStorageType?: pulumi.Input<string>;
    /**
     * Database type. Value options: MySQL, SQLServer, PostgreSQL, and PPAS.
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * Set it to true to make some parameter efficient when modifying them. Default to false.
     */
    readonly forceRestart?: pulumi.Input<boolean>;
    /**
     * Valid values are `Prepaid`, `Postpaid`, Default to `Postpaid`. Currently, the resource only supports PostPaid to PrePaid.
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    readonly instanceName?: pulumi.Input<string>;
    /**
     * User-defined DB instance storage space. Value range:
     * - [5, 2000] for MySQL/PostgreSQL/PPAS HA dual node edition;
     * - [20,1000] for MySQL 5.7 basic single node edition;
     * - [10, 2000] for SQL Server 2008R2;
     * - [20,2000] for SQL Server 2012 basic single node edition
     * Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     * Note: There is extra 5 GB storage for SQL Server Instance and it is not in specified `instanceStorage`.
     */
    readonly instanceStorage?: pulumi.Input<number>;
    /**
     * DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     */
    readonly instanceType?: pulumi.Input<string>;
    /**
     * Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
     */
    readonly maintainTime?: pulumi.Input<string>;
    /**
     * The monitoring frequency in seconds. Valid values are 5, 60, 300. Defaults to 300. 
     */
    readonly monitoringPeriod?: pulumi.Input<number>;
    /**
     * Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
     */
    readonly parameters?: pulumi.Input<pulumi.Input<inputs.rds.InstanceParameter>[]>;
    /**
     * The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * RDS database connection port.
     */
    readonly port?: pulumi.Input<string>;
    /**
     * The ID of resource group which the DB instance belongs.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * It has been deprecated from 1.69.0 and use `securityGroupIds` instead.
     *
     * @deprecated Attribute `security_group_id` has been deprecated from 1.69.0 and use `security_group_ids` instead.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * , Available in 1.69.0+) The list IDs to join ECS Security Group. At most supports three security groups.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Valid values are `normal`, `safety`, Default to `normal`. support `safety` switch to high security access mode 
     */
    readonly securityIpMode?: pulumi.Input<string>;
    /**
     * List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     */
    readonly securityIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The sql collector keep time of the instance. Valid values are `30`, `180`, `365`, `1095`, `1825`, Default to `30`.
     */
    readonly sqlCollectorConfigValue?: pulumi.Input<number>;
    /**
     * The sql collector status of the instance. Valid values are `Enabled`, `Disabled`, Default to `Disabled`.
     */
    readonly sqlCollectorStatus?: pulumi.Input<string>;
    /**
     * Actions performed on SSL functions, Valid values: `Open`: turn on SSL encryption; `Close`: turn off SSL encryption; `Update`: update SSL certificate. See more [engine and engineVersion limitation](https://www.alibabacloud.com/help/zh/doc-detail/26254.htm).
     */
    readonly sslAction?: pulumi.Input<string>;
    /**
     * Status of the SSL feature. `Yes`: SSL is turned on; `No`: SSL is turned off.
     */
    readonly sslStatus?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The TDE(Transparent Data Encryption) status. See more [engine and engineVersion limitation](https://www.alibabacloud.com/help/zh/doc-detail/26256.htm).
     */
    readonly tdeStatus?: pulumi.Input<string>;
    /**
     * The virtual switch ID to launch DB instances in one VPC. If there are multiple vswitches, separate them with commas.
     */
    readonly vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the DB instance. From version 1.8.1, it supports multiple zone.
     * If it is a multi-zone and `vswitchId` is specified, the vswitch must in the one of them.
     * The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `alicloud..getZones`.
     */
    readonly zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Whether to renewal a DB instance automatically or not. It is valid when instanceChargeType is `PrePaid`. Default to `false`.
     */
    readonly autoRenew?: pulumi.Input<boolean>;
    /**
     * Auto-renewal period of an instance, in the unit of the month. It is valid when instanceChargeType is `PrePaid`. Valid value:[1~12], Default to 1.
     */
    readonly autoRenewPeriod?: pulumi.Input<number>;
    /**
     * The upgrade method to use. Valid values:
     * - Auto: Instances are automatically upgraded to a higher minor version.
     * - Manual: Instances are forcibly upgraded to a higher minor version when the current version is unpublished.
     */
    readonly autoUpgradeMinorVersion?: pulumi.Input<string>;
    /**
     * The storage type of the instance. Valid values:
     * - local_ssd: specifies to use local SSDs. This value is recommended.
     * - cloud_ssd: specifies to use standard SSDs.
     * - cloud_essd: specifies to use enhanced SSDs (ESSDs).
     * - cloud_essd2: specifies to use enhanced SSDs (ESSDs).
     * - cloud_essd3: specifies to use enhanced SSDs (ESSDs).
     */
    readonly dbInstanceStorageType?: pulumi.Input<string>;
    /**
     * Database type. Value options: MySQL, SQLServer, PostgreSQL, and PPAS.
     */
    readonly engine: pulumi.Input<string>;
    /**
     * Database version. Value options can refer to the latest docs [CreateDBInstance](https://www.alibabacloud.com/help/doc-detail/26228.htm) `EngineVersion`.
     */
    readonly engineVersion: pulumi.Input<string>;
    /**
     * Set it to true to make some parameter efficient when modifying them. Default to false.
     */
    readonly forceRestart?: pulumi.Input<boolean>;
    /**
     * Valid values are `Prepaid`, `Postpaid`, Default to `Postpaid`. Currently, the resource only supports PostPaid to PrePaid.
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    /**
     * The name of DB instance. It a string of 2 to 256 characters.
     */
    readonly instanceName?: pulumi.Input<string>;
    /**
     * User-defined DB instance storage space. Value range:
     * - [5, 2000] for MySQL/PostgreSQL/PPAS HA dual node edition;
     * - [20,1000] for MySQL 5.7 basic single node edition;
     * - [10, 2000] for SQL Server 2008R2;
     * - [20,2000] for SQL Server 2012 basic single node edition
     * Increase progressively at a rate of 5 GB. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     * Note: There is extra 5 GB storage for SQL Server Instance and it is not in specified `instanceStorage`.
     */
    readonly instanceStorage: pulumi.Input<number>;
    /**
     * DB Instance type. For details, see [Instance type table](https://www.alibabacloud.com/help/doc-detail/26312.htm).
     */
    readonly instanceType: pulumi.Input<string>;
    /**
     * Maintainable time period format of the instance: HH:MMZ-HH:MMZ (UTC time)
     */
    readonly maintainTime?: pulumi.Input<string>;
    /**
     * The monitoring frequency in seconds. Valid values are 5, 60, 300. Defaults to 300. 
     */
    readonly monitoringPeriod?: pulumi.Input<number>;
    /**
     * Set of parameters needs to be set after DB instance was launched. Available parameters can refer to the latest docs [View database parameter templates](https://www.alibabacloud.com/help/doc-detail/26284.htm) .
     */
    readonly parameters?: pulumi.Input<pulumi.Input<inputs.rds.InstanceParameter>[]>;
    /**
     * The duration that you will buy DB instance (in month). It is valid when instanceChargeType is `PrePaid`. Valid values: [1~9], 12, 24, 36. Default to 1.
     */
    readonly period?: pulumi.Input<number>;
    /**
     * The ID of resource group which the DB instance belongs.
     */
    readonly resourceGroupId?: pulumi.Input<string>;
    /**
     * It has been deprecated from 1.69.0 and use `securityGroupIds` instead.
     *
     * @deprecated Attribute `security_group_id` has been deprecated from 1.69.0 and use `security_group_ids` instead.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * , Available in 1.69.0+) The list IDs to join ECS Security Group. At most supports three security groups.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Valid values are `normal`, `safety`, Default to `normal`. support `safety` switch to high security access mode 
     */
    readonly securityIpMode?: pulumi.Input<string>;
    /**
     * List of IP addresses allowed to access all databases of an instance. The list contains up to 1,000 IP addresses, separated by commas. Supported formats include 0.0.0.0/0, 10.23.12.24 (IP), and 10.23.12.24/24 (Classless Inter-Domain Routing (CIDR) mode. /24 represents the length of the prefix in an IP address. The range of the prefix length is [1,32]).
     */
    readonly securityIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The sql collector keep time of the instance. Valid values are `30`, `180`, `365`, `1095`, `1825`, Default to `30`.
     */
    readonly sqlCollectorConfigValue?: pulumi.Input<number>;
    /**
     * The sql collector status of the instance. Valid values are `Enabled`, `Disabled`, Default to `Disabled`.
     */
    readonly sqlCollectorStatus?: pulumi.Input<string>;
    /**
     * Actions performed on SSL functions, Valid values: `Open`: turn on SSL encryption; `Close`: turn off SSL encryption; `Update`: update SSL certificate. See more [engine and engineVersion limitation](https://www.alibabacloud.com/help/zh/doc-detail/26254.htm).
     */
    readonly sslAction?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The TDE(Transparent Data Encryption) status. See more [engine and engineVersion limitation](https://www.alibabacloud.com/help/zh/doc-detail/26256.htm).
     */
    readonly tdeStatus?: pulumi.Input<string>;
    /**
     * The virtual switch ID to launch DB instances in one VPC. If there are multiple vswitches, separate them with commas.
     */
    readonly vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the DB instance. From version 1.8.1, it supports multiple zone.
     * If it is a multi-zone and `vswitchId` is specified, the vswitch must in the one of them.
     * The multiple zone ID can be retrieved by setting `multi` to "true" in the data source `alicloud..getZones`.
     */
    readonly zoneId?: pulumi.Input<string>;
}
