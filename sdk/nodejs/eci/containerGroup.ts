// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides ECI Container Group resource.
 *
 * For information about ECI Container Group and how to use it, see [What is Container Group](https://www.alibabacloud.com/help/en/doc-detail/90341.htm).
 *
 * > **NOTE:** Available in v1.111.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.eci.ContainerGroup("example", {
 *     containerGroupName: "tf-testacc-eci-gruop",
 *     cpu: 8,
 *     memory: 16,
 *     restartPolicy: "OnFailure",
 *     securityGroupId: alicloud_security_group.group.id,
 *     vswitchId: data.alicloud_vpcs["default"].vpcs[0].vswitch_ids[0],
 *     tags: {
 *         TF: "create",
 *     },
 *     containers: [
 *         {
 *             image: "registry-vpc.cn-beijing.aliyuncs.com/eci_open/nginx:alpine",
 *             name: "nginx",
 *             workingDir: "/tmp/nginx",
 *             imagePullPolicy: "IfNotPresent",
 *             commands: [
 *                 "/bin/sh",
 *                 "-c",
 *                 "sleep 9999",
 *             ],
 *             volumeMounts: [{
 *                 mountPath: "/tmp/test",
 *                 readOnly: false,
 *                 name: "empty1",
 *             }],
 *             ports: [{
 *                 port: 80,
 *                 protocol: "TCP",
 *             }],
 *             environmentVars: [{
 *                 key: "test",
 *                 value: "nginx",
 *             }],
 *         },
 *         {
 *             image: "registry-vpc.cn-beijing.aliyuncs.com/eci_open/centos:7",
 *             name: "centos",
 *             commands: [
 *                 "/bin/sh",
 *                 "-c",
 *                 "sleep 9999",
 *             ],
 *         },
 *     ],
 *     initContainers: [{
 *         name: "init-busybox",
 *         image: "registry-vpc.cn-beijing.aliyuncs.com/eci_open/busybox:1.30",
 *         imagePullPolicy: "IfNotPresent",
 *         commands: ["echo"],
 *         args: ["hello initcontainer"],
 *     }],
 *     volumes: [
 *         {
 *             name: "empty1",
 *             type: "EmptyDirVolume",
 *         },
 *         {
 *             name: "empty2",
 *             type: "EmptyDirVolume",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * ECI Container Group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:eci/containerGroup:ContainerGroup example <container_group_id>
 * ```
 */
export class ContainerGroup extends pulumi.CustomResource {
    /**
     * Get an existing ContainerGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerGroupState, opts?: pulumi.CustomResourceOptions): ContainerGroup {
        return new ContainerGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:eci/containerGroup:ContainerGroup';

    /**
     * Returns true if the given object is an instance of ContainerGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerGroup.__pulumiType;
    }

    /**
     * Specifies whether to automatically create an EIP and bind the EIP to the elastic container instance.
     */
    public readonly autoCreateEip!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether to automatically match the image cache. Default value: false.
     */
    public readonly autoMatchImageCache!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the container group.
     */
    public readonly containerGroupName!: pulumi.Output<string>;
    /**
     * The list of containers.
     */
    public readonly containers!: pulumi.Output<outputs.eci.ContainerGroupContainer[]>;
    /**
     * The amount of CPU resources allocated to the container.
     */
    public readonly cpu!: pulumi.Output<number>;
    /**
     * The structure of dnsConfig.
     */
    public readonly dnsConfig!: pulumi.Output<outputs.eci.ContainerGroupDnsConfig | undefined>;
    /**
     * The security context of the container group.
     */
    public readonly eciSecurityContext!: pulumi.Output<outputs.eci.ContainerGroupEciSecurityContext | undefined>;
    /**
     * The bandwidth of the EIP. The default value is `5`.
     */
    public readonly eipBandwidth!: pulumi.Output<number | undefined>;
    /**
     * The ID of the elastic IP address (EIP).
     */
    public readonly eipInstanceId!: pulumi.Output<string | undefined>;
    /**
     * HostAliases.
     */
    public readonly hostAliases!: pulumi.Output<outputs.eci.ContainerGroupHostAlias[] | undefined>;
    /**
     * The image registry credential. The details see Block `imageRegistryCredential`.
     */
    public readonly imageRegistryCredentials!: pulumi.Output<outputs.eci.ContainerGroupImageRegistryCredential[] | undefined>;
    /**
     * The list of initContainers.
     */
    public readonly initContainers!: pulumi.Output<outputs.eci.ContainerGroupInitContainer[] | undefined>;
    /**
     * The address of the self-built mirror warehouse. When creating an image cache using an image in a self-built image repository with a self-signed certificate, you need to configure this parameter to skip certificate authentication to avoid image pull failure due to certificate authentication failure.
     */
    public readonly insecureRegistry!: pulumi.Output<string | undefined>;
    /**
     * The type of the ECS instance.
     */
    public readonly instanceType!: pulumi.Output<string | undefined>;
    /**
     * (Available in v1.170.0+) The Public IP of the container group.
     */
    public /*out*/ readonly internetIp!: pulumi.Output<string>;
    /**
     * (Available in v1.170.0+) The Private IP of the container group.
     */
    public /*out*/ readonly intranetIp!: pulumi.Output<string>;
    /**
     * The amount of memory resources allocated to the container.
     */
    public readonly memory!: pulumi.Output<number>;
    /**
     * The address of the self-built mirror warehouse. When creating an image cache from an image in a self-built image repository using the HTTP protocol, you need to configure this parameter so that the ECI uses the HTTP protocol to pull the image to avoid image pull failure due to different protocols.
     */
    public readonly plainHttpRegistry!: pulumi.Output<string | undefined>;
    /**
     * The RAM role that the container group assumes. ECI and ECS share the same RAM role.
     */
    public readonly ramRoleName!: pulumi.Output<string | undefined>;
    /**
     * The ID of the resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The restart policy of the container group. Valid values: `Always`, `Never`, `OnFailure`.
     */
    public readonly restartPolicy!: pulumi.Output<string>;
    /**
     * The ID of the security group to which the container group belongs. Container groups within the same security group can access each other.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * The status of container group.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The list of volumes.
     */
    public readonly volumes!: pulumi.Output<outputs.eci.ContainerGroupVolume[] | undefined>;
    /**
     * The ID of the VSwitch. Currently, container groups can only be deployed in VPC networks. The number of IP addresses in the VSwitch CIDR block determines the maximum number of container groups that can be created in the VSwitch. Before you can create an ECI instance, plan the CIDR block of the VSwitch.
     */
    public readonly vswitchId!: pulumi.Output<string>;
    /**
     * The ID of the zone where you want to deploy the container group. If no value is specified, the system assigns a zone to the container group. By default, no value is specified.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a ContainerGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerGroupArgs | ContainerGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContainerGroupState | undefined;
            resourceInputs["autoCreateEip"] = state ? state.autoCreateEip : undefined;
            resourceInputs["autoMatchImageCache"] = state ? state.autoMatchImageCache : undefined;
            resourceInputs["containerGroupName"] = state ? state.containerGroupName : undefined;
            resourceInputs["containers"] = state ? state.containers : undefined;
            resourceInputs["cpu"] = state ? state.cpu : undefined;
            resourceInputs["dnsConfig"] = state ? state.dnsConfig : undefined;
            resourceInputs["eciSecurityContext"] = state ? state.eciSecurityContext : undefined;
            resourceInputs["eipBandwidth"] = state ? state.eipBandwidth : undefined;
            resourceInputs["eipInstanceId"] = state ? state.eipInstanceId : undefined;
            resourceInputs["hostAliases"] = state ? state.hostAliases : undefined;
            resourceInputs["imageRegistryCredentials"] = state ? state.imageRegistryCredentials : undefined;
            resourceInputs["initContainers"] = state ? state.initContainers : undefined;
            resourceInputs["insecureRegistry"] = state ? state.insecureRegistry : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["internetIp"] = state ? state.internetIp : undefined;
            resourceInputs["intranetIp"] = state ? state.intranetIp : undefined;
            resourceInputs["memory"] = state ? state.memory : undefined;
            resourceInputs["plainHttpRegistry"] = state ? state.plainHttpRegistry : undefined;
            resourceInputs["ramRoleName"] = state ? state.ramRoleName : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["restartPolicy"] = state ? state.restartPolicy : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["volumes"] = state ? state.volumes : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ContainerGroupArgs | undefined;
            if ((!args || args.containerGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'containerGroupName'");
            }
            if ((!args || args.containers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'containers'");
            }
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            if ((!args || args.vswitchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchId'");
            }
            resourceInputs["autoCreateEip"] = args ? args.autoCreateEip : undefined;
            resourceInputs["autoMatchImageCache"] = args ? args.autoMatchImageCache : undefined;
            resourceInputs["containerGroupName"] = args ? args.containerGroupName : undefined;
            resourceInputs["containers"] = args ? args.containers : undefined;
            resourceInputs["cpu"] = args ? args.cpu : undefined;
            resourceInputs["dnsConfig"] = args ? args.dnsConfig : undefined;
            resourceInputs["eciSecurityContext"] = args ? args.eciSecurityContext : undefined;
            resourceInputs["eipBandwidth"] = args ? args.eipBandwidth : undefined;
            resourceInputs["eipInstanceId"] = args ? args.eipInstanceId : undefined;
            resourceInputs["hostAliases"] = args ? args.hostAliases : undefined;
            resourceInputs["imageRegistryCredentials"] = args ? args.imageRegistryCredentials : undefined;
            resourceInputs["initContainers"] = args ? args.initContainers : undefined;
            resourceInputs["insecureRegistry"] = args ? args.insecureRegistry : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["memory"] = args ? args.memory : undefined;
            resourceInputs["plainHttpRegistry"] = args ? args.plainHttpRegistry : undefined;
            resourceInputs["ramRoleName"] = args ? args.ramRoleName : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["restartPolicy"] = args ? args.restartPolicy : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["volumes"] = args ? args.volumes : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["internetIp"] = undefined /*out*/;
            resourceInputs["intranetIp"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ContainerGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContainerGroup resources.
 */
export interface ContainerGroupState {
    /**
     * Specifies whether to automatically create an EIP and bind the EIP to the elastic container instance.
     */
    autoCreateEip?: pulumi.Input<boolean>;
    /**
     * Specifies whether to automatically match the image cache. Default value: false.
     */
    autoMatchImageCache?: pulumi.Input<boolean>;
    /**
     * The name of the container group.
     */
    containerGroupName?: pulumi.Input<string>;
    /**
     * The list of containers.
     */
    containers?: pulumi.Input<pulumi.Input<inputs.eci.ContainerGroupContainer>[]>;
    /**
     * The amount of CPU resources allocated to the container.
     */
    cpu?: pulumi.Input<number>;
    /**
     * The structure of dnsConfig.
     */
    dnsConfig?: pulumi.Input<inputs.eci.ContainerGroupDnsConfig>;
    /**
     * The security context of the container group.
     */
    eciSecurityContext?: pulumi.Input<inputs.eci.ContainerGroupEciSecurityContext>;
    /**
     * The bandwidth of the EIP. The default value is `5`.
     */
    eipBandwidth?: pulumi.Input<number>;
    /**
     * The ID of the elastic IP address (EIP).
     */
    eipInstanceId?: pulumi.Input<string>;
    /**
     * HostAliases.
     */
    hostAliases?: pulumi.Input<pulumi.Input<inputs.eci.ContainerGroupHostAlias>[]>;
    /**
     * The image registry credential. The details see Block `imageRegistryCredential`.
     */
    imageRegistryCredentials?: pulumi.Input<pulumi.Input<inputs.eci.ContainerGroupImageRegistryCredential>[]>;
    /**
     * The list of initContainers.
     */
    initContainers?: pulumi.Input<pulumi.Input<inputs.eci.ContainerGroupInitContainer>[]>;
    /**
     * The address of the self-built mirror warehouse. When creating an image cache using an image in a self-built image repository with a self-signed certificate, you need to configure this parameter to skip certificate authentication to avoid image pull failure due to certificate authentication failure.
     */
    insecureRegistry?: pulumi.Input<string>;
    /**
     * The type of the ECS instance.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * (Available in v1.170.0+) The Public IP of the container group.
     */
    internetIp?: pulumi.Input<string>;
    /**
     * (Available in v1.170.0+) The Private IP of the container group.
     */
    intranetIp?: pulumi.Input<string>;
    /**
     * The amount of memory resources allocated to the container.
     */
    memory?: pulumi.Input<number>;
    /**
     * The address of the self-built mirror warehouse. When creating an image cache from an image in a self-built image repository using the HTTP protocol, you need to configure this parameter so that the ECI uses the HTTP protocol to pull the image to avoid image pull failure due to different protocols.
     */
    plainHttpRegistry?: pulumi.Input<string>;
    /**
     * The RAM role that the container group assumes. ECI and ECS share the same RAM role.
     */
    ramRoleName?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The restart policy of the container group. Valid values: `Always`, `Never`, `OnFailure`.
     */
    restartPolicy?: pulumi.Input<string>;
    /**
     * The ID of the security group to which the container group belongs. Container groups within the same security group can access each other.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * The status of container group.
     */
    status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The list of volumes.
     */
    volumes?: pulumi.Input<pulumi.Input<inputs.eci.ContainerGroupVolume>[]>;
    /**
     * The ID of the VSwitch. Currently, container groups can only be deployed in VPC networks. The number of IP addresses in the VSwitch CIDR block determines the maximum number of container groups that can be created in the VSwitch. Before you can create an ECI instance, plan the CIDR block of the VSwitch.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The ID of the zone where you want to deploy the container group. If no value is specified, the system assigns a zone to the container group. By default, no value is specified.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ContainerGroup resource.
 */
export interface ContainerGroupArgs {
    /**
     * Specifies whether to automatically create an EIP and bind the EIP to the elastic container instance.
     */
    autoCreateEip?: pulumi.Input<boolean>;
    /**
     * Specifies whether to automatically match the image cache. Default value: false.
     */
    autoMatchImageCache?: pulumi.Input<boolean>;
    /**
     * The name of the container group.
     */
    containerGroupName: pulumi.Input<string>;
    /**
     * The list of containers.
     */
    containers: pulumi.Input<pulumi.Input<inputs.eci.ContainerGroupContainer>[]>;
    /**
     * The amount of CPU resources allocated to the container.
     */
    cpu?: pulumi.Input<number>;
    /**
     * The structure of dnsConfig.
     */
    dnsConfig?: pulumi.Input<inputs.eci.ContainerGroupDnsConfig>;
    /**
     * The security context of the container group.
     */
    eciSecurityContext?: pulumi.Input<inputs.eci.ContainerGroupEciSecurityContext>;
    /**
     * The bandwidth of the EIP. The default value is `5`.
     */
    eipBandwidth?: pulumi.Input<number>;
    /**
     * The ID of the elastic IP address (EIP).
     */
    eipInstanceId?: pulumi.Input<string>;
    /**
     * HostAliases.
     */
    hostAliases?: pulumi.Input<pulumi.Input<inputs.eci.ContainerGroupHostAlias>[]>;
    /**
     * The image registry credential. The details see Block `imageRegistryCredential`.
     */
    imageRegistryCredentials?: pulumi.Input<pulumi.Input<inputs.eci.ContainerGroupImageRegistryCredential>[]>;
    /**
     * The list of initContainers.
     */
    initContainers?: pulumi.Input<pulumi.Input<inputs.eci.ContainerGroupInitContainer>[]>;
    /**
     * The address of the self-built mirror warehouse. When creating an image cache using an image in a self-built image repository with a self-signed certificate, you need to configure this parameter to skip certificate authentication to avoid image pull failure due to certificate authentication failure.
     */
    insecureRegistry?: pulumi.Input<string>;
    /**
     * The type of the ECS instance.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The amount of memory resources allocated to the container.
     */
    memory?: pulumi.Input<number>;
    /**
     * The address of the self-built mirror warehouse. When creating an image cache from an image in a self-built image repository using the HTTP protocol, you need to configure this parameter so that the ECI uses the HTTP protocol to pull the image to avoid image pull failure due to different protocols.
     */
    plainHttpRegistry?: pulumi.Input<string>;
    /**
     * The RAM role that the container group assumes. ECI and ECS share the same RAM role.
     */
    ramRoleName?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The restart policy of the container group. Valid values: `Always`, `Never`, `OnFailure`.
     */
    restartPolicy?: pulumi.Input<string>;
    /**
     * The ID of the security group to which the container group belongs. Container groups within the same security group can access each other.
     */
    securityGroupId: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     * - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The list of volumes.
     */
    volumes?: pulumi.Input<pulumi.Input<inputs.eci.ContainerGroupVolume>[]>;
    /**
     * The ID of the VSwitch. Currently, container groups can only be deployed in VPC networks. The number of IP addresses in the VSwitch CIDR block determines the maximum number of container groups that can be created in the VSwitch. Before you can create an ECI instance, plan the CIDR block of the VSwitch.
     */
    vswitchId: pulumi.Input<string>;
    /**
     * The ID of the zone where you want to deploy the container group. If no value is specified, the system assigns a zone to the container group. By default, no value is specified.
     */
    zoneId?: pulumi.Input<string>;
}
