// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Bastion Host Host Account Attachment resource to add list host accounts into one user group and one host group.
 *
 * > **NOTE:** Available in v1.135.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultHost = new alicloud.bastionhost.Host("defaultHost", {
 *     instanceId: "bastionhost-cn-tl3xxxxxxx",
 *     hostName: _var.name,
 *     activeAddressType: "Private",
 *     hostPrivateAddress: "172.16.0.10",
 *     osType: "Linux",
 *     source: "Local",
 * });
 * const defaultHostAccount: alicloud.bastionhost.HostAccount[] = [];
 * for (const range = {value: 0}; range.value < 3; range.value++) {
 *     defaultHostAccount.push(new alicloud.bastionhost.HostAccount(`defaultHostAccount-${range.value}`, {
 *         instanceId: defaultHost.instanceId,
 *         hostAccountName: `example_value-${range.value}`,
 *         hostId: defaultHost.hostId,
 *         protocolName: "SSH",
 *         password: "YourPassword12345",
 *     }));
 * }
 * const defaultUserGroup = new alicloud.bastionhost.UserGroup("defaultUserGroup", {
 *     instanceId: defaultHost.instanceId,
 *     userGroupName: "my-local-user",
 * });
 * const defaultHostGroup = new alicloud.bastionhost.HostGroup("defaultHostGroup", {
 *     hostGroupName: "example_value",
 *     instanceId: "bastionhost-cn-tl3xxxxxxx",
 * });
 * const defaultHostGroupAccountUserGroupAttachment = new alicloud.bastionhost.HostGroupAccountUserGroupAttachment("defaultHostGroupAccountUserGroupAttachment", {
 *     instanceId: defaultHost.instanceId,
 *     userGroupId: defaultUserGroup.userGroupId,
 *     hostGroupId: defaultHostGroup.hostGroupId,
 *     hostAccountNames: defaultHostAccount.map(__item => __item.hostAccountName),
 * });
 * ```
 *
 * ## Import
 *
 * Bastion Host Host Account can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:bastionhost/hostGroupAccountUserGroupAttachment:HostGroupAccountUserGroupAttachment example <instance_id>:<user_group_id>:<host_group_id>
 * ```
 */
export class HostGroupAccountUserGroupAttachment extends pulumi.CustomResource {
    /**
     * Get an existing HostGroupAccountUserGroupAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostGroupAccountUserGroupAttachmentState, opts?: pulumi.CustomResourceOptions): HostGroupAccountUserGroupAttachment {
        return new HostGroupAccountUserGroupAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:bastionhost/hostGroupAccountUserGroupAttachment:HostGroupAccountUserGroupAttachment';

    /**
     * Returns true if the given object is an instance of HostGroupAccountUserGroupAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HostGroupAccountUserGroupAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HostGroupAccountUserGroupAttachment.__pulumiType;
    }

    /**
     * A list names of the host account.
     */
    public readonly hostAccountNames!: pulumi.Output<string[]>;
    /**
     * The ID of the host group.
     */
    public readonly hostGroupId!: pulumi.Output<string>;
    /**
     * The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
     */
    public readonly userGroupId!: pulumi.Output<string>;

    /**
     * Create a HostGroupAccountUserGroupAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostGroupAccountUserGroupAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostGroupAccountUserGroupAttachmentArgs | HostGroupAccountUserGroupAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HostGroupAccountUserGroupAttachmentState | undefined;
            resourceInputs["hostAccountNames"] = state ? state.hostAccountNames : undefined;
            resourceInputs["hostGroupId"] = state ? state.hostGroupId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["userGroupId"] = state ? state.userGroupId : undefined;
        } else {
            const args = argsOrState as HostGroupAccountUserGroupAttachmentArgs | undefined;
            if ((!args || args.hostAccountNames === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostAccountNames'");
            }
            if ((!args || args.hostGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostGroupId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.userGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userGroupId'");
            }
            resourceInputs["hostAccountNames"] = args ? args.hostAccountNames : undefined;
            resourceInputs["hostGroupId"] = args ? args.hostGroupId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["userGroupId"] = args ? args.userGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HostGroupAccountUserGroupAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostGroupAccountUserGroupAttachment resources.
 */
export interface HostGroupAccountUserGroupAttachmentState {
    /**
     * A list names of the host account.
     */
    hostAccountNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the host group.
     */
    hostGroupId?: pulumi.Input<string>;
    /**
     * The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
     */
    userGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HostGroupAccountUserGroupAttachment resource.
 */
export interface HostGroupAccountUserGroupAttachmentArgs {
    /**
     * A list names of the host account.
     */
    hostAccountNames: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the host group.
     */
    hostGroupId: pulumi.Input<string>;
    /**
     * The ID of the Bastionhost instance where you want to authorize the user to manage the specified hosts and host accounts.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
     */
    userGroupId: pulumi.Input<string>;
}
