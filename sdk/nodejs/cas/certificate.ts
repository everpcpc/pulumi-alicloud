// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * @deprecated This resource has been deprecated in favour of ServiceCertificate
 */
export class Certificate extends pulumi.CustomResource {
    /**
     * Get an existing Certificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertificateState, opts?: pulumi.CustomResourceOptions): Certificate {
        pulumi.log.warn("Certificate is deprecated: This resource has been deprecated in favour of ServiceCertificate")
        return new Certificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cas/certificate:Certificate';

    /**
     * Returns true if the given object is an instance of Certificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Certificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Certificate.__pulumiType;
    }

    /**
     * Cert of the Certificate in which the Certificate will add.
     */
    public readonly cert!: pulumi.Output<string>;
    public readonly certificateName!: pulumi.Output<string>;
    /**
     * Key of the Certificate in which the Certificate will add.
     */
    public readonly key!: pulumi.Output<string>;
    public readonly lang!: pulumi.Output<string | undefined>;
    /**
     * Name of the Certificate. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.129.0 and it will be remove in the future version. Please use the new attribute 'certificate_name' instead.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated This resource has been deprecated in favour of ServiceCertificate */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated This resource has been deprecated in favour of ServiceCertificate */
    constructor(name: string, argsOrState?: CertificateArgs | CertificateState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Certificate is deprecated: This resource has been deprecated in favour of ServiceCertificate")
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CertificateState | undefined;
            inputs["cert"] = state ? state.cert : undefined;
            inputs["certificateName"] = state ? state.certificateName : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["lang"] = state ? state.lang : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as CertificateArgs | undefined;
            if ((!args || args.cert === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cert'");
            }
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            inputs["cert"] = args ? args.cert : undefined;
            inputs["certificateName"] = args ? args.certificateName : undefined;
            inputs["key"] = args ? args.key : undefined;
            inputs["lang"] = args ? args.lang : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Certificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Certificate resources.
 */
export interface CertificateState {
    /**
     * Cert of the Certificate in which the Certificate will add.
     */
    cert?: pulumi.Input<string>;
    certificateName?: pulumi.Input<string>;
    /**
     * Key of the Certificate in which the Certificate will add.
     */
    key?: pulumi.Input<string>;
    lang?: pulumi.Input<string>;
    /**
     * Name of the Certificate. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.129.0 and it will be remove in the future version. Please use the new attribute 'certificate_name' instead.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * Cert of the Certificate in which the Certificate will add.
     */
    cert: pulumi.Input<string>;
    certificateName?: pulumi.Input<string>;
    /**
     * Key of the Certificate in which the Certificate will add.
     */
    key: pulumi.Input<string>;
    lang?: pulumi.Input<string>;
    /**
     * Name of the Certificate. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     *
     * @deprecated Field 'name' has been deprecated from provider version 1.129.0 and it will be remove in the future version. Please use the new attribute 'certificate_name' instead.
     */
    name?: pulumi.Input<string>;
}
