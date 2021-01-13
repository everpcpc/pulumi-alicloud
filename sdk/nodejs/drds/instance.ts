// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Distributed Relational Database Service (DRDS) is a lightweight (stateless), flexible, stable, and efficient middleware product independently developed by Alibaba Group to resolve scalability issues with single-host relational databases.
 * With its compatibility with MySQL protocols and syntaxes, DRDS enables database/table sharding, smooth scaling, configuration upgrade/downgrade,
 * transparent read/write splitting, and distributed transactions, providing O&M capabilities for distributed databases throughout their entire lifecycle.
 *
 * For information about DRDS and how to use it, see [What is DRDS](https://www.alibabacloud.com/help/doc-detail/29659.htm).
 *
 * > **NOTE:** At present, DRDS instance only can be supported in the regions: cn-shenzhen, cn-beijing, cn-hangzhou, cn-hongkong, cn-qingdao, ap-southeast-1.
 *
 * > **NOTE:** Currently, this resource only support `Domestic Site Account`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultInstance = new alicloud.drds.Instance("default", {
 *     description: "drds instance",
 *     instanceChargeType: "PostPaid",
 *     instanceSeries: "drds.sn1.4c8g",
 *     specification: "drds.sn1.4c8g.8C16G",
 *     vswitchId: "vsw-bp1jlu3swk8rq2yoi40ey",
 *     zoneId: "cn-hangzhou-e",
 * });
 * ```
 *
 * ## Import
 *
 * Distributed Relational Database Service (DRDS) can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:drds/instance:Instance example drds-abc123456
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
    public static readonly __pulumiType = 'alicloud:drds/instance:Instance';

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
     * Description of the DRDS instance, This description can have a string of 2 to 256 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
     */
    public readonly instanceChargeType!: pulumi.Output<string | undefined>;
    /**
     * User-defined DRDS instance node spec. Value range:
     * - `drds.sn1.4c8g` for DRDS instance Starter version;
     * - `drds.sn1.8c16g` for DRDS instance Standard edition;
     * - `drds.sn1.16c32g` for DRDS instance Enterprise Edition;
     * - `drds.sn1.32c64g` for DRDS instance Extreme Edition;
     */
    public readonly instanceSeries!: pulumi.Output<string>;
    /**
     * User-defined DRDS instance specification. Value range:
     * - `drds.sn1.4c8g` for DRDS instance Starter version;
     * - value range : `drds.sn1.4c8g.8c16g`, `drds.sn1.4c8g.16c32g`, `drds.sn1.4c8g.32c64g`, `drds.sn1.4c8g.64c128g`
     * - `drds.sn1.8c16g` for DRDS instance Standard edition;
     * - value range : `drds.sn1.8c16g.16c32g`, `drds.sn1.8c16g.32c64g`, `drds.sn1.8c16g.64c128g`
     * - `drds.sn1.16c32g` for DRDS instance Enterprise Edition;
     * - value range : `drds.sn1.16c32g.32c64g`, `drds.sn1.16c32g.64c128g`
     * - `drds.sn1.32c64g` for DRDS instance Extreme Edition;
     * - value range : `drds.sn1.32c64g.128c256g`
     */
    public readonly specification!: pulumi.Output<string>;
    /**
     * The VSwitch ID to launch in.
     */
    public readonly vswitchId!: pulumi.Output<string>;
    /**
     * The Zone to launch the DRDS instance.
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
            inputs["description"] = state ? state.description : undefined;
            inputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            inputs["instanceSeries"] = state ? state.instanceSeries : undefined;
            inputs["specification"] = state ? state.specification : undefined;
            inputs["vswitchId"] = state ? state.vswitchId : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.description === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.instanceSeries === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'instanceSeries'");
            }
            if ((!args || args.specification === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'specification'");
            }
            if ((!args || args.vswitchId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'vswitchId'");
            }
            if ((!args || args.zoneId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'zoneId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            inputs["instanceSeries"] = args ? args.instanceSeries : undefined;
            inputs["specification"] = args ? args.specification : undefined;
            inputs["vswitchId"] = args ? args.vswitchId : undefined;
            inputs["zoneId"] = args ? args.zoneId : undefined;
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
     * Description of the DRDS instance, This description can have a string of 2 to 256 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    /**
     * User-defined DRDS instance node spec. Value range:
     * - `drds.sn1.4c8g` for DRDS instance Starter version;
     * - `drds.sn1.8c16g` for DRDS instance Standard edition;
     * - `drds.sn1.16c32g` for DRDS instance Enterprise Edition;
     * - `drds.sn1.32c64g` for DRDS instance Extreme Edition;
     */
    readonly instanceSeries?: pulumi.Input<string>;
    /**
     * User-defined DRDS instance specification. Value range:
     * - `drds.sn1.4c8g` for DRDS instance Starter version;
     * - value range : `drds.sn1.4c8g.8c16g`, `drds.sn1.4c8g.16c32g`, `drds.sn1.4c8g.32c64g`, `drds.sn1.4c8g.64c128g`
     * - `drds.sn1.8c16g` for DRDS instance Standard edition;
     * - value range : `drds.sn1.8c16g.16c32g`, `drds.sn1.8c16g.32c64g`, `drds.sn1.8c16g.64c128g`
     * - `drds.sn1.16c32g` for DRDS instance Enterprise Edition;
     * - value range : `drds.sn1.16c32g.32c64g`, `drds.sn1.16c32g.64c128g`
     * - `drds.sn1.32c64g` for DRDS instance Extreme Edition;
     * - value range : `drds.sn1.32c64g.128c256g`
     */
    readonly specification?: pulumi.Input<string>;
    /**
     * The VSwitch ID to launch in.
     */
    readonly vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the DRDS instance.
     */
    readonly zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Description of the DRDS instance, This description can have a string of 2 to 256 characters.
     */
    readonly description: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
     */
    readonly instanceChargeType?: pulumi.Input<string>;
    /**
     * User-defined DRDS instance node spec. Value range:
     * - `drds.sn1.4c8g` for DRDS instance Starter version;
     * - `drds.sn1.8c16g` for DRDS instance Standard edition;
     * - `drds.sn1.16c32g` for DRDS instance Enterprise Edition;
     * - `drds.sn1.32c64g` for DRDS instance Extreme Edition;
     */
    readonly instanceSeries: pulumi.Input<string>;
    /**
     * User-defined DRDS instance specification. Value range:
     * - `drds.sn1.4c8g` for DRDS instance Starter version;
     * - value range : `drds.sn1.4c8g.8c16g`, `drds.sn1.4c8g.16c32g`, `drds.sn1.4c8g.32c64g`, `drds.sn1.4c8g.64c128g`
     * - `drds.sn1.8c16g` for DRDS instance Standard edition;
     * - value range : `drds.sn1.8c16g.16c32g`, `drds.sn1.8c16g.32c64g`, `drds.sn1.8c16g.64c128g`
     * - `drds.sn1.16c32g` for DRDS instance Enterprise Edition;
     * - value range : `drds.sn1.16c32g.32c64g`, `drds.sn1.16c32g.64c128g`
     * - `drds.sn1.32c64g` for DRDS instance Extreme Edition;
     * - value range : `drds.sn1.32c64g.128c256g`
     */
    readonly specification: pulumi.Input<string>;
    /**
     * The VSwitch ID to launch in.
     */
    readonly vswitchId: pulumi.Input<string>;
    /**
     * The Zone to launch the DRDS instance.
     */
    readonly zoneId: pulumi.Input<string>;
}
