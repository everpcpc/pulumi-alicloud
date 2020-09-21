// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source provides the VServer groups related to a server load balancer.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "slbservergroups";
 * const defaultZones = alicloud.getZones({
 *     availableDiskCategory: "cloud_efficiency",
 *     availableResourceCreation: "VSwitch",
 * });
 * const defaultNetwork = new alicloud.vpc.Network("defaultNetwork", {cidrBlock: "172.16.0.0/16"});
 * const defaultSwitch = new alicloud.vpc.Switch("defaultSwitch", {
 *     vpcId: defaultNetwork.id,
 *     cidrBlock: "172.16.0.0/16",
 *     availabilityZone: defaultZones.then(defaultZones => defaultZones.zones[0].id),
 * });
 * const defaultLoadBalancer = new alicloud.slb.LoadBalancer("defaultLoadBalancer", {vswitchId: defaultSwitch.id});
 * const defaultServerGroup = new alicloud.slb.ServerGroup("defaultServerGroup", {loadBalancerId: defaultLoadBalancer.id});
 * const sampleDs = defaultLoadBalancer.id.apply(id => alicloud.slb.getServerGroups({
 *     loadBalancerId: id,
 * }));
 * export const firstSlbServerGroupId = sampleDs.slbServerGroups[0].id;
 * ```
 */
export function getServerGroups(args: GetServerGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetServerGroupsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("alicloud:slb/getServerGroups:getServerGroups", {
        "ids": args.ids,
        "loadBalancerId": args.loadBalancerId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getServerGroups.
 */
export interface GetServerGroupsArgs {
    /**
     * A list of VServer group IDs to filter results.
     */
    readonly ids?: string[];
    /**
     * ID of the SLB.
     */
    readonly loadBalancerId: string;
    /**
     * A regex string to filter results by VServer group name.
     */
    readonly nameRegex?: string;
    readonly outputFile?: string;
}

/**
 * A collection of values returned by getServerGroups.
 */
export interface GetServerGroupsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of SLB VServer groups IDs.
     */
    readonly ids: string[];
    readonly loadBalancerId: string;
    readonly nameRegex?: string;
    /**
     * A list of SLB VServer groups names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    /**
     * A list of SLB VServer groups. Each element contains the following attributes:
     */
    readonly slbServerGroups: outputs.slb.GetServerGroupsSlbServerGroup[];
}
