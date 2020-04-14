// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CEN route entry resource. Cloud Enterprise Network (CEN) supports publishing and withdrawing route entries of attached networks. You can publish a route entry of an attached VPC or VBR to a CEN instance, then other attached networks can learn the route if there is no route conflict. You can withdraw a published route entry when CEN does not need it any more.
 * 
 * For information about CEN route entries publishment and how to use it, see [Manage network routes](https://www.alibabacloud.com/help/doc-detail/86980.htm).
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 * 
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-testAccCenRouteEntryConfig";
 * 
 * const hz = new alicloud.Provider("hz", {
 *     region: "cn-hangzhou",
 * });
 * const defaultZones = alicloud.getZones({
 *     availableDiskCategory: "cloudEfficiency",
 *     availableResourceCreation: "VSwitch",
 * }, {provider: hz});
 * const defaultInstanceTypes = alicloud.ecs.getInstanceTypes({
 *     availabilityZone: defaultZones.zones[0].id,
 *     cpuCoreCount: 1,
 *     memorySize: 2,
 * }, {provider: hz});
 * const defaultImages = alicloud.ecs.getImages({
 *     mostRecent: true,
 *     nameRegex: "^ubuntu_18.*64",
 *     owners: "system",
 * }, {provider: hz});
 * const vpc = new alicloud.vpc.Network("vpc", {
 *     cidrBlock: "172.16.0.0/12",
 * }, {provider: hz});
 * const defaultSwitch = new alicloud.vpc.Switch("default", {
 *     availabilityZone: defaultZones.zones[0].id,
 *     cidrBlock: "172.16.0.0/21",
 *     vpcId: vpc.id,
 * }, {provider: hz});
 * const defaultSecurityGroup = new alicloud.ecs.SecurityGroup("default", {
 *     description: "foo",
 *     vpcId: vpc.id,
 * }, {provider: hz});
 * const defaultInstance = new alicloud.ecs.Instance("default", {
 *     imageId: defaultImages.images[0].id,
 *     instanceName: name,
 *     instanceType: defaultInstanceTypes.instanceTypes[0].id,
 *     internetChargeType: "PayByTraffic",
 *     internetMaxBandwidthOut: 5,
 *     securityGroups: [defaultSecurityGroup.id],
 *     systemDiskCategory: "cloudEfficiency",
 *     vswitchId: defaultSwitch.id,
 * }, {provider: hz});
 * const cen = new alicloud.cen.Instance("cen", {});
 * const attach = new alicloud.cen.InstanceAttachment("attach", {
 *     childInstanceId: vpc.id,
 *     childInstanceRegionId: "cn-hangzhou",
 *     instanceId: cen.id,
 * }, {dependsOn: [defaultSwitch]});
 * const route = new alicloud.vpc.RouteEntry("route", {
 *     destinationCidrblock: "11.0.0.0/16",
 *     nexthopId: defaultInstance.id,
 *     nexthopType: "Instance",
 *     routeTableId: vpc.routeTableId,
 * }, {provider: hz});
 * const foo = new alicloud.cen.RouteEntry("foo", {
 *     cidrBlock: route.destinationCidrblock,
 *     instanceId: cen.id,
 *     routeTableId: vpc.routeTableId,
 * }, {provider: hz,dependsOn: [attach]});
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-alicloud/blob/master/website/docs/r/cen_route_entry.html.markdown.
 */
export class RouteEntry extends pulumi.CustomResource {
    /**
     * Get an existing RouteEntry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteEntryState, opts?: pulumi.CustomResourceOptions): RouteEntry {
        return new RouteEntry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cen/routeEntry:RouteEntry';

    /**
     * Returns true if the given object is an instance of RouteEntry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteEntry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteEntry.__pulumiType;
    }

    /**
     * The destination CIDR block of the route entry to publish.
     */
    public readonly cidrBlock!: pulumi.Output<string>;
    /**
     * The ID of the CEN.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The route table of the attached VBR or VPC.
     */
    public readonly routeTableId!: pulumi.Output<string>;

    /**
     * Create a RouteEntry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteEntryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteEntryArgs | RouteEntryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RouteEntryState | undefined;
            inputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["routeTableId"] = state ? state.routeTableId : undefined;
        } else {
            const args = argsOrState as RouteEntryArgs | undefined;
            if (!args || args.cidrBlock === undefined) {
                throw new Error("Missing required property 'cidrBlock'");
            }
            if (!args || args.instanceId === undefined) {
                throw new Error("Missing required property 'instanceId'");
            }
            if (!args || args.routeTableId === undefined) {
                throw new Error("Missing required property 'routeTableId'");
            }
            inputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["routeTableId"] = args ? args.routeTableId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RouteEntry.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouteEntry resources.
 */
export interface RouteEntryState {
    /**
     * The destination CIDR block of the route entry to publish.
     */
    readonly cidrBlock?: pulumi.Input<string>;
    /**
     * The ID of the CEN.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * The route table of the attached VBR or VPC.
     */
    readonly routeTableId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouteEntry resource.
 */
export interface RouteEntryArgs {
    /**
     * The destination CIDR block of the route entry to publish.
     */
    readonly cidrBlock: pulumi.Input<string>;
    /**
     * The ID of the CEN.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * The route table of the attached VBR or VPC.
     */
    readonly routeTableId: pulumi.Input<string>;
}
