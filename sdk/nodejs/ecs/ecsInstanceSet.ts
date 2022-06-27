// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a ECS Instance Set resource.
 *
 * For information about ECS Instance Set and how to use it, see [What is Instance Set](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/runinstances).
 *
 * > **NOTE:** Available in v1.173.0+.
 *
 * > **NOTE:** This resource is used to batch create a group of instance resources with the same configuration. However, this resource is not recommended. `alicloud.ecs.Instance` is preferred.
 *
 * > **NOTE:** In the instances managed by this resource, names are automatically generated based on `instanceName` and `uniqueSuffix`.
 *
 * > **NOTE:** Only `tags` support batch modification.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-testaccecsset";
 * const defaultZones = alicloud.getZones({
 *     availableDiskCategory: "cloud_efficiency",
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultInstanceTypes = defaultZones.then(defaultZones => alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones?[0]?.id,
 *     cpuCoreCount: 1,
 *     memorySize: 2,
 * }));
 * const defaultImages = alicloud.ecs.getImages({
 *     nameRegex: "^ubuntu_[0-9]+_[0-9]+_x64*",
 *     mostRecent: true,
 *     owners: "system",
 * });
 * const defaultNetworks = alicloud.vpc.getNetworks({
 *     nameRegex: "default-NODELETING",
 * });
 * const defaultSwitches = Promise.all([defaultNetworks, defaultZones]).then(([defaultNetworks, defaultZones]) => alicloud.vpc.getSwitches({
 *     vpcId: defaultNetworks.ids?[0],
 *     zoneId: defaultZones.zones?[0]?.id,
 * }));
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("defaultSecurityGroup", {vpcId: defaultNetworks.then(defaultNetworks => defaultNetworks.ids?[0])});
 * const beijingK = new alicloud.ecs.EcsInstanceSet("beijingK", {
 *     amount: 100,
 *     imageId: defaultImages.then(defaultImages => defaultImages.images?[0]?.id),
 *     instanceType: defaultInstanceTypes.then(defaultInstanceTypes => defaultInstanceTypes.instanceTypes?[0]?.id),
 *     instanceName: name,
 *     instanceChargeType: "PostPaid",
 *     systemDiskPerformanceLevel: "PL0",
 *     systemDiskCategory: "cloud_essd",
 *     systemDiskSize: 200,
 *     vswitchId: defaultSwitches.then(defaultSwitches => defaultSwitches.ids?[0]),
 *     securityGroupIds: [defaultSecurityGroup].map(__item => __item.id),
 *     zoneId: defaultZones.then(defaultZones => defaultZones.zones?[0]?.id),
 * });
 * ```
 */
export class EcsInstanceSet extends pulumi.CustomResource {
    /**
     * Get an existing EcsInstanceSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EcsInstanceSetState, opts?: pulumi.CustomResourceOptions): EcsInstanceSet {
        return new EcsInstanceSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/ecsInstanceSet:EcsInstanceSet';

    /**
     * Returns true if the given object is an instance of EcsInstanceSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EcsInstanceSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EcsInstanceSet.__pulumiType;
    }

    /**
     * The number of instances that you want to create. Valid values: `1` to `100`.
     */
    public readonly amount!: pulumi.Output<number | undefined>;
    /**
     * The automatic release time of the `PostPaid` instance.
     */
    public readonly autoReleaseTime!: pulumi.Output<string | undefined>;
    /**
     * Whether to enable auto-renewal for the instance. This parameter is valid only when the `instanceChargeType` is set to `PrePaid`.
     */
    public readonly autoRenew!: pulumi.Output<boolean | undefined>;
    /**
     * Auto renewal period of an instance, in the unit of month. It is valid when `instanceChargeType` is `PrePaid`.
     * - When `periodUnit` is `Month`, Valid values: `1`, `2`, `3`, `6`, `12`.
     * - When `periodUnit` is `Week`, Valid values: `1`, `2`, `3`.
     */
    public readonly autoRenewPeriod!: pulumi.Output<number | undefined>;
    /**
     * The list of data disks created with instance. See the following `Block dataDisks`.
     */
    public readonly dataDisks!: pulumi.Output<outputs.ecs.EcsInstanceSetDataDisk[] | undefined>;
    /**
     * The ID of the dedicated host on which to create the instance. If the `dedicatedHostId` is specified, the `spotStrategy` and `spotPriceLimit`  are ignored. This is because preemptible instances cannot be created on dedicated hosts.
     */
    public readonly dedicatedHostId!: pulumi.Output<string | undefined>;
    /**
     * Whether to enable release protection for the instance.
     */
    public readonly deletionProtection!: pulumi.Output<boolean>;
    /**
     * The ID of the deployment set to which to deploy the instance.
     */
    public readonly deploymentSetId!: pulumi.Output<string | undefined>;
    /**
     * The description of the instance, This description can have a string of 2 to 256 characters, It cannot begin with `http://` or `https://`.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The hostname of instance.
     */
    public readonly hostName!: pulumi.Output<string>;
    /**
     * The ID of the Elastic High Performance Computing (E-HPC) cluster to which to assign the instance.
     */
    public readonly hpcClusterId!: pulumi.Output<string | undefined>;
    /**
     * The Image to use for the instance.
     */
    public readonly imageId!: pulumi.Output<string>;
    /**
     * The billing method of the instance. Valid values: `PrePaid`, `PostPaid`.
     */
    public readonly instanceChargeType!: pulumi.Output<string | undefined>;
    /**
     * A list of ECS Instance ID.
     */
    public /*out*/ readonly instanceIds!: pulumi.Output<string[]>;
    /**
     * The name of the ECS. This instanceName can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen, and must not begin with `http://` or `https://`.
     */
    public readonly instanceName!: pulumi.Output<string | undefined>;
    /**
     * The type of instance to start.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The Internet charge type of the instance. Valid values are `PayByBandwidth`, `PayByTraffic`.
     */
    public readonly internetChargeType!: pulumi.Output<string>;
    /**
     * The Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bit per second). Value values: `1` to `100`.
     */
    public readonly internetMaxBandwidthOut!: pulumi.Output<number>;
    /**
     * The name of key pair that can login ECS instance successfully without password.
     */
    public readonly keyPairName!: pulumi.Output<string | undefined>;
    /**
     * The ID of the launch template.
     */
    public readonly launchTemplateId!: pulumi.Output<string | undefined>;
    /**
     * The name of the launch template. To use a launch template to create an instance, you must use the `launchTemplateId` or `launchTemplateName` parameter to specify the launch template.
     */
    public readonly launchTemplateName!: pulumi.Output<string | undefined>;
    /**
     * The version of the launch template.
     */
    public readonly launchTemplateVersion!: pulumi.Output<string | undefined>;
    /**
     * A list of NetworkInterface. See the following `Block networkInterfaces`.
     */
    public readonly networkInterfaces!: pulumi.Output<outputs.ecs.EcsInstanceSetNetworkInterface[] | undefined>;
    /**
     * The password to an instance is a string of 8 to 30 characters. It must contain uppercase/lowercase letters and numerals, but cannot contain special symbols.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Whether to use the password preset in the image.
     */
    public readonly passwordInherit!: pulumi.Output<boolean | undefined>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`.
     * - When `periodUnit` is `Month`, Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `12`, `24`, `36`, `48`, `60`.
     * - When `periodUnit` is `Week`, Valid values: `1`, `2`, `3`.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The duration unit that you will buy the resource. It is valid when `instanceChargeType` is 'PrePaid'. Valid value: `Week`, `Month`.
     */
    public readonly periodUnit!: pulumi.Output<string | undefined>;
    /**
     * The Instance RAM role name.
     */
    public readonly ramRoleName!: pulumi.Output<string | undefined>;
    /**
     * The ID of resource group which the instance belongs.
     */
    public readonly resourceGroupId!: pulumi.Output<string | undefined>;
    /**
     * The security enhancement strategy.
     * - `Active`: Enable security enhancement strategy, it only works on system images.
     * - `Deactive`: Disable security enhancement strategy, it works on all images.
     */
    public readonly securityEnhancementStrategy!: pulumi.Output<string | undefined>;
    /**
     * A list of security group ids to associate with.
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * The hourly price threshold of a instance, and it takes effect only when parameter 'spot_strategy' is 'SpotWithPriceLimit'. Three decimals is allowed at most.
     */
    public readonly spotPriceLimit!: pulumi.Output<number>;
    /**
     * The spot strategy of a Pay-As-You-Go instance, and it takes effect only when parameter `instanceChargeType` is 'PostPaid'.
     * - `NoSpot`: A regular Pay-As-You-Go instance.
     * - `SpotWithPriceLimit`: A price threshold for a spot instance.
     * - `SpotAsPriceGo`: A price that is based on the highest Pay-As-You-Go instance
     */
    public readonly spotStrategy!: pulumi.Output<string>;
    /**
     * The ID of the automatic snapshot policy applied to the system disk.
     */
    public readonly systemDiskAutoSnapshotPolicyId!: pulumi.Output<string | undefined>;
    /**
     * The category of the system disk. Valid values are `cloudEfficiency`, `cloudSsd`, `cloudEssd`, `cloud`.
     */
    public readonly systemDiskCategory!: pulumi.Output<string>;
    /**
     * The description of the system disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     */
    public readonly systemDiskDescription!: pulumi.Output<string | undefined>;
    /**
     * The name of the system disk.
     */
    public readonly systemDiskName!: pulumi.Output<string | undefined>;
    /**
     * The performance level of the ESSD used as the system disk. Valid values: `PL0`, `PL1`, `PL2`, `PL3`.
     */
    public readonly systemDiskPerformanceLevel!: pulumi.Output<string>;
    /**
     * The size of the system disk, measured in GiB. Value range:  values: `20` to `500`.
     */
    public readonly systemDiskSize!: pulumi.Output<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Whether to automatically append incremental suffixes to the hostname specified by the HostName parameter and to the instance name specified by the InstanceName parameter when you batch create instances. The incremental suffixes can range from `001` to `999`.
     */
    public readonly uniqueSuffix!: pulumi.Output<boolean | undefined>;
    /**
     * The virtual switch ID to launch in VPC.
     */
    public readonly vswitchId!: pulumi.Output<string | undefined>;
    /**
     * The ID of the zone in which to create the instance.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a EcsInstanceSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EcsInstanceSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EcsInstanceSetArgs | EcsInstanceSetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EcsInstanceSetState | undefined;
            resourceInputs["amount"] = state ? state.amount : undefined;
            resourceInputs["autoReleaseTime"] = state ? state.autoReleaseTime : undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["autoRenewPeriod"] = state ? state.autoRenewPeriod : undefined;
            resourceInputs["dataDisks"] = state ? state.dataDisks : undefined;
            resourceInputs["dedicatedHostId"] = state ? state.dedicatedHostId : undefined;
            resourceInputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            resourceInputs["deploymentSetId"] = state ? state.deploymentSetId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["hostName"] = state ? state.hostName : undefined;
            resourceInputs["hpcClusterId"] = state ? state.hpcClusterId : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            resourceInputs["instanceIds"] = state ? state.instanceIds : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["internetChargeType"] = state ? state.internetChargeType : undefined;
            resourceInputs["internetMaxBandwidthOut"] = state ? state.internetMaxBandwidthOut : undefined;
            resourceInputs["keyPairName"] = state ? state.keyPairName : undefined;
            resourceInputs["launchTemplateId"] = state ? state.launchTemplateId : undefined;
            resourceInputs["launchTemplateName"] = state ? state.launchTemplateName : undefined;
            resourceInputs["launchTemplateVersion"] = state ? state.launchTemplateVersion : undefined;
            resourceInputs["networkInterfaces"] = state ? state.networkInterfaces : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["passwordInherit"] = state ? state.passwordInherit : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["periodUnit"] = state ? state.periodUnit : undefined;
            resourceInputs["ramRoleName"] = state ? state.ramRoleName : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["securityEnhancementStrategy"] = state ? state.securityEnhancementStrategy : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["spotPriceLimit"] = state ? state.spotPriceLimit : undefined;
            resourceInputs["spotStrategy"] = state ? state.spotStrategy : undefined;
            resourceInputs["systemDiskAutoSnapshotPolicyId"] = state ? state.systemDiskAutoSnapshotPolicyId : undefined;
            resourceInputs["systemDiskCategory"] = state ? state.systemDiskCategory : undefined;
            resourceInputs["systemDiskDescription"] = state ? state.systemDiskDescription : undefined;
            resourceInputs["systemDiskName"] = state ? state.systemDiskName : undefined;
            resourceInputs["systemDiskPerformanceLevel"] = state ? state.systemDiskPerformanceLevel : undefined;
            resourceInputs["systemDiskSize"] = state ? state.systemDiskSize : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["uniqueSuffix"] = state ? state.uniqueSuffix : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as EcsInstanceSetArgs | undefined;
            if ((!args || args.imageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageId'");
            }
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            if ((!args || args.securityGroupIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupIds'");
            }
            resourceInputs["amount"] = args ? args.amount : undefined;
            resourceInputs["autoReleaseTime"] = args ? args.autoReleaseTime : undefined;
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["autoRenewPeriod"] = args ? args.autoRenewPeriod : undefined;
            resourceInputs["dataDisks"] = args ? args.dataDisks : undefined;
            resourceInputs["dedicatedHostId"] = args ? args.dedicatedHostId : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["deploymentSetId"] = args ? args.deploymentSetId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["hostName"] = args ? args.hostName : undefined;
            resourceInputs["hpcClusterId"] = args ? args.hpcClusterId : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["internetChargeType"] = args ? args.internetChargeType : undefined;
            resourceInputs["internetMaxBandwidthOut"] = args ? args.internetMaxBandwidthOut : undefined;
            resourceInputs["keyPairName"] = args ? args.keyPairName : undefined;
            resourceInputs["launchTemplateId"] = args ? args.launchTemplateId : undefined;
            resourceInputs["launchTemplateName"] = args ? args.launchTemplateName : undefined;
            resourceInputs["launchTemplateVersion"] = args ? args.launchTemplateVersion : undefined;
            resourceInputs["networkInterfaces"] = args ? args.networkInterfaces : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["passwordInherit"] = args ? args.passwordInherit : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["periodUnit"] = args ? args.periodUnit : undefined;
            resourceInputs["ramRoleName"] = args ? args.ramRoleName : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["securityEnhancementStrategy"] = args ? args.securityEnhancementStrategy : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["spotPriceLimit"] = args ? args.spotPriceLimit : undefined;
            resourceInputs["spotStrategy"] = args ? args.spotStrategy : undefined;
            resourceInputs["systemDiskAutoSnapshotPolicyId"] = args ? args.systemDiskAutoSnapshotPolicyId : undefined;
            resourceInputs["systemDiskCategory"] = args ? args.systemDiskCategory : undefined;
            resourceInputs["systemDiskDescription"] = args ? args.systemDiskDescription : undefined;
            resourceInputs["systemDiskName"] = args ? args.systemDiskName : undefined;
            resourceInputs["systemDiskPerformanceLevel"] = args ? args.systemDiskPerformanceLevel : undefined;
            resourceInputs["systemDiskSize"] = args ? args.systemDiskSize : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["uniqueSuffix"] = args ? args.uniqueSuffix : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["instanceIds"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EcsInstanceSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EcsInstanceSet resources.
 */
export interface EcsInstanceSetState {
    /**
     * The number of instances that you want to create. Valid values: `1` to `100`.
     */
    amount?: pulumi.Input<number>;
    /**
     * The automatic release time of the `PostPaid` instance.
     */
    autoReleaseTime?: pulumi.Input<string>;
    /**
     * Whether to enable auto-renewal for the instance. This parameter is valid only when the `instanceChargeType` is set to `PrePaid`.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * Auto renewal period of an instance, in the unit of month. It is valid when `instanceChargeType` is `PrePaid`.
     * - When `periodUnit` is `Month`, Valid values: `1`, `2`, `3`, `6`, `12`.
     * - When `periodUnit` is `Week`, Valid values: `1`, `2`, `3`.
     */
    autoRenewPeriod?: pulumi.Input<number>;
    /**
     * The list of data disks created with instance. See the following `Block dataDisks`.
     */
    dataDisks?: pulumi.Input<pulumi.Input<inputs.ecs.EcsInstanceSetDataDisk>[]>;
    /**
     * The ID of the dedicated host on which to create the instance. If the `dedicatedHostId` is specified, the `spotStrategy` and `spotPriceLimit`  are ignored. This is because preemptible instances cannot be created on dedicated hosts.
     */
    dedicatedHostId?: pulumi.Input<string>;
    /**
     * Whether to enable release protection for the instance.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * The ID of the deployment set to which to deploy the instance.
     */
    deploymentSetId?: pulumi.Input<string>;
    /**
     * The description of the instance, This description can have a string of 2 to 256 characters, It cannot begin with `http://` or `https://`.
     */
    description?: pulumi.Input<string>;
    /**
     * The hostname of instance.
     */
    hostName?: pulumi.Input<string>;
    /**
     * The ID of the Elastic High Performance Computing (E-HPC) cluster to which to assign the instance.
     */
    hpcClusterId?: pulumi.Input<string>;
    /**
     * The Image to use for the instance.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The billing method of the instance. Valid values: `PrePaid`, `PostPaid`.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * A list of ECS Instance ID.
     */
    instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the ECS. This instanceName can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen, and must not begin with `http://` or `https://`.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The type of instance to start.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The Internet charge type of the instance. Valid values are `PayByBandwidth`, `PayByTraffic`.
     */
    internetChargeType?: pulumi.Input<string>;
    /**
     * The Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bit per second). Value values: `1` to `100`.
     */
    internetMaxBandwidthOut?: pulumi.Input<number>;
    /**
     * The name of key pair that can login ECS instance successfully without password.
     */
    keyPairName?: pulumi.Input<string>;
    /**
     * The ID of the launch template.
     */
    launchTemplateId?: pulumi.Input<string>;
    /**
     * The name of the launch template. To use a launch template to create an instance, you must use the `launchTemplateId` or `launchTemplateName` parameter to specify the launch template.
     */
    launchTemplateName?: pulumi.Input<string>;
    /**
     * The version of the launch template.
     */
    launchTemplateVersion?: pulumi.Input<string>;
    /**
     * A list of NetworkInterface. See the following `Block networkInterfaces`.
     */
    networkInterfaces?: pulumi.Input<pulumi.Input<inputs.ecs.EcsInstanceSetNetworkInterface>[]>;
    /**
     * The password to an instance is a string of 8 to 30 characters. It must contain uppercase/lowercase letters and numerals, but cannot contain special symbols.
     */
    password?: pulumi.Input<string>;
    /**
     * Whether to use the password preset in the image.
     */
    passwordInherit?: pulumi.Input<boolean>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`.
     * - When `periodUnit` is `Month`, Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `12`, `24`, `36`, `48`, `60`.
     * - When `periodUnit` is `Week`, Valid values: `1`, `2`, `3`.
     */
    period?: pulumi.Input<number>;
    /**
     * The duration unit that you will buy the resource. It is valid when `instanceChargeType` is 'PrePaid'. Valid value: `Week`, `Month`.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The Instance RAM role name.
     */
    ramRoleName?: pulumi.Input<string>;
    /**
     * The ID of resource group which the instance belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The security enhancement strategy.
     * - `Active`: Enable security enhancement strategy, it only works on system images.
     * - `Deactive`: Disable security enhancement strategy, it works on all images.
     */
    securityEnhancementStrategy?: pulumi.Input<string>;
    /**
     * A list of security group ids to associate with.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The hourly price threshold of a instance, and it takes effect only when parameter 'spot_strategy' is 'SpotWithPriceLimit'. Three decimals is allowed at most.
     */
    spotPriceLimit?: pulumi.Input<number>;
    /**
     * The spot strategy of a Pay-As-You-Go instance, and it takes effect only when parameter `instanceChargeType` is 'PostPaid'.
     * - `NoSpot`: A regular Pay-As-You-Go instance.
     * - `SpotWithPriceLimit`: A price threshold for a spot instance.
     * - `SpotAsPriceGo`: A price that is based on the highest Pay-As-You-Go instance
     */
    spotStrategy?: pulumi.Input<string>;
    /**
     * The ID of the automatic snapshot policy applied to the system disk.
     */
    systemDiskAutoSnapshotPolicyId?: pulumi.Input<string>;
    /**
     * The category of the system disk. Valid values are `cloudEfficiency`, `cloudSsd`, `cloudEssd`, `cloud`.
     */
    systemDiskCategory?: pulumi.Input<string>;
    /**
     * The description of the system disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     */
    systemDiskDescription?: pulumi.Input<string>;
    /**
     * The name of the system disk.
     */
    systemDiskName?: pulumi.Input<string>;
    /**
     * The performance level of the ESSD used as the system disk. Valid values: `PL0`, `PL1`, `PL2`, `PL3`.
     */
    systemDiskPerformanceLevel?: pulumi.Input<string>;
    /**
     * The size of the system disk, measured in GiB. Value range:  values: `20` to `500`.
     */
    systemDiskSize?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether to automatically append incremental suffixes to the hostname specified by the HostName parameter and to the instance name specified by the InstanceName parameter when you batch create instances. The incremental suffixes can range from `001` to `999`.
     */
    uniqueSuffix?: pulumi.Input<boolean>;
    /**
     * The virtual switch ID to launch in VPC.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The ID of the zone in which to create the instance.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EcsInstanceSet resource.
 */
export interface EcsInstanceSetArgs {
    /**
     * The number of instances that you want to create. Valid values: `1` to `100`.
     */
    amount?: pulumi.Input<number>;
    /**
     * The automatic release time of the `PostPaid` instance.
     */
    autoReleaseTime?: pulumi.Input<string>;
    /**
     * Whether to enable auto-renewal for the instance. This parameter is valid only when the `instanceChargeType` is set to `PrePaid`.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * Auto renewal period of an instance, in the unit of month. It is valid when `instanceChargeType` is `PrePaid`.
     * - When `periodUnit` is `Month`, Valid values: `1`, `2`, `3`, `6`, `12`.
     * - When `periodUnit` is `Week`, Valid values: `1`, `2`, `3`.
     */
    autoRenewPeriod?: pulumi.Input<number>;
    /**
     * The list of data disks created with instance. See the following `Block dataDisks`.
     */
    dataDisks?: pulumi.Input<pulumi.Input<inputs.ecs.EcsInstanceSetDataDisk>[]>;
    /**
     * The ID of the dedicated host on which to create the instance. If the `dedicatedHostId` is specified, the `spotStrategy` and `spotPriceLimit`  are ignored. This is because preemptible instances cannot be created on dedicated hosts.
     */
    dedicatedHostId?: pulumi.Input<string>;
    /**
     * Whether to enable release protection for the instance.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * The ID of the deployment set to which to deploy the instance.
     */
    deploymentSetId?: pulumi.Input<string>;
    /**
     * The description of the instance, This description can have a string of 2 to 256 characters, It cannot begin with `http://` or `https://`.
     */
    description?: pulumi.Input<string>;
    /**
     * The hostname of instance.
     */
    hostName?: pulumi.Input<string>;
    /**
     * The ID of the Elastic High Performance Computing (E-HPC) cluster to which to assign the instance.
     */
    hpcClusterId?: pulumi.Input<string>;
    /**
     * The Image to use for the instance.
     */
    imageId: pulumi.Input<string>;
    /**
     * The billing method of the instance. Valid values: `PrePaid`, `PostPaid`.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The name of the ECS. This instanceName can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen, and must not begin with `http://` or `https://`.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The type of instance to start.
     */
    instanceType: pulumi.Input<string>;
    /**
     * The Internet charge type of the instance. Valid values are `PayByBandwidth`, `PayByTraffic`.
     */
    internetChargeType?: pulumi.Input<string>;
    /**
     * The Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bit per second). Value values: `1` to `100`.
     */
    internetMaxBandwidthOut?: pulumi.Input<number>;
    /**
     * The name of key pair that can login ECS instance successfully without password.
     */
    keyPairName?: pulumi.Input<string>;
    /**
     * The ID of the launch template.
     */
    launchTemplateId?: pulumi.Input<string>;
    /**
     * The name of the launch template. To use a launch template to create an instance, you must use the `launchTemplateId` or `launchTemplateName` parameter to specify the launch template.
     */
    launchTemplateName?: pulumi.Input<string>;
    /**
     * The version of the launch template.
     */
    launchTemplateVersion?: pulumi.Input<string>;
    /**
     * A list of NetworkInterface. See the following `Block networkInterfaces`.
     */
    networkInterfaces?: pulumi.Input<pulumi.Input<inputs.ecs.EcsInstanceSetNetworkInterface>[]>;
    /**
     * The password to an instance is a string of 8 to 30 characters. It must contain uppercase/lowercase letters and numerals, but cannot contain special symbols.
     */
    password?: pulumi.Input<string>;
    /**
     * Whether to use the password preset in the image.
     */
    passwordInherit?: pulumi.Input<boolean>;
    /**
     * The duration that you will buy the resource, in month. It is valid when `instanceChargeType` is `PrePaid`.
     * - When `periodUnit` is `Month`, Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `12`, `24`, `36`, `48`, `60`.
     * - When `periodUnit` is `Week`, Valid values: `1`, `2`, `3`.
     */
    period?: pulumi.Input<number>;
    /**
     * The duration unit that you will buy the resource. It is valid when `instanceChargeType` is 'PrePaid'. Valid value: `Week`, `Month`.
     */
    periodUnit?: pulumi.Input<string>;
    /**
     * The Instance RAM role name.
     */
    ramRoleName?: pulumi.Input<string>;
    /**
     * The ID of resource group which the instance belongs.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The security enhancement strategy.
     * - `Active`: Enable security enhancement strategy, it only works on system images.
     * - `Deactive`: Disable security enhancement strategy, it works on all images.
     */
    securityEnhancementStrategy?: pulumi.Input<string>;
    /**
     * A list of security group ids to associate with.
     */
    securityGroupIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The hourly price threshold of a instance, and it takes effect only when parameter 'spot_strategy' is 'SpotWithPriceLimit'. Three decimals is allowed at most.
     */
    spotPriceLimit?: pulumi.Input<number>;
    /**
     * The spot strategy of a Pay-As-You-Go instance, and it takes effect only when parameter `instanceChargeType` is 'PostPaid'.
     * - `NoSpot`: A regular Pay-As-You-Go instance.
     * - `SpotWithPriceLimit`: A price threshold for a spot instance.
     * - `SpotAsPriceGo`: A price that is based on the highest Pay-As-You-Go instance
     */
    spotStrategy?: pulumi.Input<string>;
    /**
     * The ID of the automatic snapshot policy applied to the system disk.
     */
    systemDiskAutoSnapshotPolicyId?: pulumi.Input<string>;
    /**
     * The category of the system disk. Valid values are `cloudEfficiency`, `cloudSsd`, `cloudEssd`, `cloud`.
     */
    systemDiskCategory?: pulumi.Input<string>;
    /**
     * The description of the system disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     */
    systemDiskDescription?: pulumi.Input<string>;
    /**
     * The name of the system disk.
     */
    systemDiskName?: pulumi.Input<string>;
    /**
     * The performance level of the ESSD used as the system disk. Valid values: `PL0`, `PL1`, `PL2`, `PL3`.
     */
    systemDiskPerformanceLevel?: pulumi.Input<string>;
    /**
     * The size of the system disk, measured in GiB. Value range:  values: `20` to `500`.
     */
    systemDiskSize?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether to automatically append incremental suffixes to the hostname specified by the HostName parameter and to the instance name specified by the InstanceName parameter when you batch create instances. The incremental suffixes can range from `001` to `999`.
     */
    uniqueSuffix?: pulumi.Input<boolean>;
    /**
     * The virtual switch ID to launch in VPC.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The ID of the zone in which to create the instance.
     */
    zoneId?: pulumi.Input<string>;
}
