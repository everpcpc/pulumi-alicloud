// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Cloud Firewall Address Book resource.
 *
 * For information about Cloud Firewall Address Book and how to use it, see [What is Address Book](https://www.alibabacloud.com/help/en/cloud-firewall/latest/describeaddressbook).
 *
 * > **NOTE:** Available in v1.178.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.cloudfirewall.AddressBook("example", {
 *     autoAddTagEcs: 0,
 *     description: "example_value",
 *     ecsTags: [{
 *         tagKey: "created",
 *         tagValue: "tfTestAcc0",
 *     }],
 *     groupName: "example_value",
 *     groupType: "tag",
 *     tagRelation: "and",
 * });
 * ```
 *
 * ## Import
 *
 * Cloud Firewall Address Book can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cloudfirewall/addressBook:AddressBook example <id>
 * ```
 */
export class AddressBook extends pulumi.CustomResource {
    /**
     * Get an existing AddressBook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AddressBookState, opts?: pulumi.CustomResourceOptions): AddressBook {
        return new AddressBook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cloudfirewall/addressBook:AddressBook';

    /**
     * Returns true if the given object is an instance of AddressBook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AddressBook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AddressBook.__pulumiType;
    }

    /**
     * The list of addresses.
     */
    public readonly addressLists!: pulumi.Output<string[] | undefined>;
    /**
     * Whether you want to automatically add new matching tags of the ECS IP address to the Address Book. Valid values: `0`, `1`.
     */
    public readonly autoAddTagEcs!: pulumi.Output<number | undefined>;
    /**
     * The description of the Address Book.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * A list of ECS tags. See the following `Block ecsTags`.
     */
    public readonly ecsTags!: pulumi.Output<outputs.cloudfirewall.AddressBookEcsTag[] | undefined>;
    /**
     * The name of the Address Book.
     */
    public readonly groupName!: pulumi.Output<string>;
    /**
     * The type of the Address Book. Valid values:  `ip`, `tag`.
     */
    public readonly groupType!: pulumi.Output<string>;
    /**
     * The language of the content within the request and response. Valid values: `en`, `zh`.
     */
    public readonly lang!: pulumi.Output<string | undefined>;
    /**
     * The logical relation among the ECS tags that to be matched. Valid values:
     * - **and**: Only the public IP addresses of ECS instances that match all the specified tags can be added to the Address Book. This is the default value.
     * - **or**: The public IP addresses of ECS instances that match one of the specified tags can be added to the Address Book.
     */
    public readonly tagRelation!: pulumi.Output<string | undefined>;

    /**
     * Create a AddressBook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AddressBookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AddressBookArgs | AddressBookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AddressBookState | undefined;
            resourceInputs["addressLists"] = state ? state.addressLists : undefined;
            resourceInputs["autoAddTagEcs"] = state ? state.autoAddTagEcs : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ecsTags"] = state ? state.ecsTags : undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
            resourceInputs["groupType"] = state ? state.groupType : undefined;
            resourceInputs["lang"] = state ? state.lang : undefined;
            resourceInputs["tagRelation"] = state ? state.tagRelation : undefined;
        } else {
            const args = argsOrState as AddressBookArgs | undefined;
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.groupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupName'");
            }
            if ((!args || args.groupType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupType'");
            }
            resourceInputs["addressLists"] = args ? args.addressLists : undefined;
            resourceInputs["autoAddTagEcs"] = args ? args.autoAddTagEcs : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ecsTags"] = args ? args.ecsTags : undefined;
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["groupType"] = args ? args.groupType : undefined;
            resourceInputs["lang"] = args ? args.lang : undefined;
            resourceInputs["tagRelation"] = args ? args.tagRelation : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AddressBook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AddressBook resources.
 */
export interface AddressBookState {
    /**
     * The list of addresses.
     */
    addressLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether you want to automatically add new matching tags of the ECS IP address to the Address Book. Valid values: `0`, `1`.
     */
    autoAddTagEcs?: pulumi.Input<number>;
    /**
     * The description of the Address Book.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of ECS tags. See the following `Block ecsTags`.
     */
    ecsTags?: pulumi.Input<pulumi.Input<inputs.cloudfirewall.AddressBookEcsTag>[]>;
    /**
     * The name of the Address Book.
     */
    groupName?: pulumi.Input<string>;
    /**
     * The type of the Address Book. Valid values:  `ip`, `tag`.
     */
    groupType?: pulumi.Input<string>;
    /**
     * The language of the content within the request and response. Valid values: `en`, `zh`.
     */
    lang?: pulumi.Input<string>;
    /**
     * The logical relation among the ECS tags that to be matched. Valid values:
     * - **and**: Only the public IP addresses of ECS instances that match all the specified tags can be added to the Address Book. This is the default value.
     * - **or**: The public IP addresses of ECS instances that match one of the specified tags can be added to the Address Book.
     */
    tagRelation?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AddressBook resource.
 */
export interface AddressBookArgs {
    /**
     * The list of addresses.
     */
    addressLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether you want to automatically add new matching tags of the ECS IP address to the Address Book. Valid values: `0`, `1`.
     */
    autoAddTagEcs?: pulumi.Input<number>;
    /**
     * The description of the Address Book.
     */
    description: pulumi.Input<string>;
    /**
     * A list of ECS tags. See the following `Block ecsTags`.
     */
    ecsTags?: pulumi.Input<pulumi.Input<inputs.cloudfirewall.AddressBookEcsTag>[]>;
    /**
     * The name of the Address Book.
     */
    groupName: pulumi.Input<string>;
    /**
     * The type of the Address Book. Valid values:  `ip`, `tag`.
     */
    groupType: pulumi.Input<string>;
    /**
     * The language of the content within the request and response. Valid values: `en`, `zh`.
     */
    lang?: pulumi.Input<string>;
    /**
     * The logical relation among the ECS tags that to be matched. Valid values:
     * - **and**: Only the public IP addresses of ECS instances that match all the specified tags can be added to the Address Book. This is the default value.
     * - **or**: The public IP addresses of ECS instances that match one of the specified tags can be added to the Address Book.
     */
    tagRelation?: pulumi.Input<string>;
}