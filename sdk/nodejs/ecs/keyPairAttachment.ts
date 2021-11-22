// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class KeyPairAttachment extends pulumi.CustomResource {
    /**
     * Get an existing KeyPairAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyPairAttachmentState, opts?: pulumi.CustomResourceOptions): KeyPairAttachment {
        return new KeyPairAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ecs/keyPairAttachment:KeyPairAttachment';

    /**
     * Returns true if the given object is an instance of KeyPairAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KeyPairAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KeyPairAttachment.__pulumiType;
    }

    /**
     * Set it to true and it will reboot instances which attached with the key pair to make key pair affect immediately.
     */
    public readonly force!: pulumi.Output<boolean | undefined>;
    /**
     * The list of ECS instance's IDs.
     */
    public readonly instanceIds!: pulumi.Output<string[]>;
    /**
     * The name of key pair used to bind.
     *
     * @deprecated Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.
     */
    public readonly keyName!: pulumi.Output<string>;
    public readonly keyPairName!: pulumi.Output<string>;

    /**
     * Create a KeyPairAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyPairAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyPairAttachmentArgs | KeyPairAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeyPairAttachmentState | undefined;
            inputs["force"] = state ? state.force : undefined;
            inputs["instanceIds"] = state ? state.instanceIds : undefined;
            inputs["keyName"] = state ? state.keyName : undefined;
            inputs["keyPairName"] = state ? state.keyPairName : undefined;
        } else {
            const args = argsOrState as KeyPairAttachmentArgs | undefined;
            if ((!args || args.instanceIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceIds'");
            }
            inputs["force"] = args ? args.force : undefined;
            inputs["instanceIds"] = args ? args.instanceIds : undefined;
            inputs["keyName"] = args ? args.keyName : undefined;
            inputs["keyPairName"] = args ? args.keyPairName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(KeyPairAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KeyPairAttachment resources.
 */
export interface KeyPairAttachmentState {
    /**
     * Set it to true and it will reboot instances which attached with the key pair to make key pair affect immediately.
     */
    force?: pulumi.Input<boolean>;
    /**
     * The list of ECS instance's IDs.
     */
    instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of key pair used to bind.
     *
     * @deprecated Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.
     */
    keyName?: pulumi.Input<string>;
    keyPairName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KeyPairAttachment resource.
 */
export interface KeyPairAttachmentArgs {
    /**
     * Set it to true and it will reboot instances which attached with the key pair to make key pair affect immediately.
     */
    force?: pulumi.Input<boolean>;
    /**
     * The list of ECS instance's IDs.
     */
    instanceIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of key pair used to bind.
     *
     * @deprecated Field 'key_name' has been deprecated from provider version 1.121.0. New field 'key_pair_name' instead.
     */
    keyName?: pulumi.Input<string>;
    keyPairName?: pulumi.Input<string>;
}
