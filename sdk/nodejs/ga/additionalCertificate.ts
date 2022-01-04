// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Global Accelerator (GA) Additional Certificate resource.
 *
 * For information about Global Accelerator (GA) Additional Certificate and how to use it, see [What is Additional Certificate](https://www.alibabacloud.com/help/en/doc-detail/302356.html).
 *
 * > **NOTE:** Available in v1.150.0+.
 *
 * ## Import
 *
 * Global Accelerator (GA) Additional Certificate can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ga/additionalCertificate:AdditionalCertificate example <accelerator_id>:<listener_id>:<domain>
 * ```
 */
export class AdditionalCertificate extends pulumi.CustomResource {
    /**
     * Get an existing AdditionalCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AdditionalCertificateState, opts?: pulumi.CustomResourceOptions): AdditionalCertificate {
        return new AdditionalCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ga/additionalCertificate:AdditionalCertificate';

    /**
     * Returns true if the given object is an instance of AdditionalCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AdditionalCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AdditionalCertificate.__pulumiType;
    }

    /**
     * The ID of the GA instance.
     */
    public readonly acceleratorId!: pulumi.Output<string>;
    /**
     * The Certificate ID.
     */
    public readonly certificateId!: pulumi.Output<string>;
    /**
     * The domain name specified by the certificate. **NOTE:** You can associate each domain name with only one additional certificate.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * The ID of the listener. **NOTE:** Only HTTPS listeners support this parameter.
     */
    public readonly listenerId!: pulumi.Output<string>;

    /**
     * Create a AdditionalCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AdditionalCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AdditionalCertificateArgs | AdditionalCertificateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AdditionalCertificateState | undefined;
            inputs["acceleratorId"] = state ? state.acceleratorId : undefined;
            inputs["certificateId"] = state ? state.certificateId : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["listenerId"] = state ? state.listenerId : undefined;
        } else {
            const args = argsOrState as AdditionalCertificateArgs | undefined;
            if ((!args || args.acceleratorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acceleratorId'");
            }
            if ((!args || args.certificateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateId'");
            }
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.listenerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerId'");
            }
            inputs["acceleratorId"] = args ? args.acceleratorId : undefined;
            inputs["certificateId"] = args ? args.certificateId : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["listenerId"] = args ? args.listenerId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AdditionalCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AdditionalCertificate resources.
 */
export interface AdditionalCertificateState {
    /**
     * The ID of the GA instance.
     */
    acceleratorId?: pulumi.Input<string>;
    /**
     * The Certificate ID.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * The domain name specified by the certificate. **NOTE:** You can associate each domain name with only one additional certificate.
     */
    domain?: pulumi.Input<string>;
    /**
     * The ID of the listener. **NOTE:** Only HTTPS listeners support this parameter.
     */
    listenerId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AdditionalCertificate resource.
 */
export interface AdditionalCertificateArgs {
    /**
     * The ID of the GA instance.
     */
    acceleratorId: pulumi.Input<string>;
    /**
     * The Certificate ID.
     */
    certificateId: pulumi.Input<string>;
    /**
     * The domain name specified by the certificate. **NOTE:** You can associate each domain name with only one additional certificate.
     */
    domain: pulumi.Input<string>;
    /**
     * The ID of the listener. **NOTE:** Only HTTPS listeners support this parameter.
     */
    listenerId: pulumi.Input<string>;
}
