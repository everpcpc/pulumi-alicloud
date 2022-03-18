// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * You can use DCDN to improve the overall performance of your website and accelerate content delivery to improve user experience. For information about Alicloud DCDN Domain and how to use it, see [What is Resource Alicloud DCDN Domain](https://www.alibabacloud.com/help/en/doc-detail/130628.htm).
 *
 * > **NOTE:** Available in v1.94.0+.
 *
 * > **NOTE:** You must activate the Dynamic Route for CDN (DCDN) service before you create an accelerated domain.
 *
 * > **NOTE:** Make sure that you have obtained an Internet content provider (ICP) filling for the accelerated domain.
 *
 * > **NOTE:** If the origin content is not saved on Alibaba Cloud, the content must be reviewed by Alibaba Cloud. The review will be completed by the next working day after you submit the application.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.dcdn.Domain("example", {
 *     domainName: "example.com",
 *     scope: "overseas",
 *     sources: [{
 *         content: "1.1.1.1",
 *         port: 80,
 *         priority: "20",
 *         type: "ipaddr",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * DCDN Domain can be imported using the id or DCDN Domain name, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:dcdn/domain:Domain example example.com
 * ```
 */
export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainState, opts?: pulumi.CustomResourceOptions): Domain {
        return new Domain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:dcdn/domain:Domain';

    /**
     * Returns true if the given object is an instance of Domain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Domain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Domain.__pulumiType;
    }

    /**
     * Indicates the name of the certificate if the HTTPS protocol is enabled.
     */
    public readonly certName!: pulumi.Output<string>;
    /**
     * The type of the certificate. Valid values:
     * `free`: a free certificate.
     * `cas`: a certificate purchased from Alibaba Cloud SSL Certificates Service.
     * `upload`: a user uploaded certificate.
     */
    public readonly certType!: pulumi.Output<string | undefined>;
    /**
     * The URL that is used to test the accessibility of the origin.
     */
    public readonly checkUrl!: pulumi.Output<string | undefined>;
    /**
     * The name of the accelerated domain.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * Specifies whether to check the certificate name for duplicates. If you set the value to 1, the system does not perform the check and overwrites the information of the existing certificate with the same name.
     */
    public readonly forceSet!: pulumi.Output<string | undefined>;
    /**
     * The ID of the resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * The acceleration region.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * The top-level domain name.
     */
    public readonly securityToken!: pulumi.Output<string | undefined>;
    /**
     * The origin information.
     */
    public readonly sources!: pulumi.Output<outputs.dcdn.DomainSource[]>;
    /**
     * The private key. Specify this parameter only if you enable the SSL certificate.
     */
    public readonly sslPri!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the SSL certificate is enabled. Valid values: `on` enabled, `off` disabled.
     */
    public readonly sslProtocol!: pulumi.Output<string | undefined>;
    /**
     * Indicates the public key of the certificate if the HTTPS protocol is enabled.
     */
    public readonly sslPub!: pulumi.Output<string | undefined>;
    /**
     * The status of DCDN Domain. Valid values: `online`, `offline`. Default to `online`.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * The top-level domain name.
     */
    public readonly topLevelDomain!: pulumi.Output<string | undefined>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainArgs | DomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainState | undefined;
            resourceInputs["certName"] = state ? state.certName : undefined;
            resourceInputs["certType"] = state ? state.certType : undefined;
            resourceInputs["checkUrl"] = state ? state.checkUrl : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["forceSet"] = state ? state.forceSet : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["securityToken"] = state ? state.securityToken : undefined;
            resourceInputs["sources"] = state ? state.sources : undefined;
            resourceInputs["sslPri"] = state ? state.sslPri : undefined;
            resourceInputs["sslProtocol"] = state ? state.sslProtocol : undefined;
            resourceInputs["sslPub"] = state ? state.sslPub : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["topLevelDomain"] = state ? state.topLevelDomain : undefined;
        } else {
            const args = argsOrState as DomainArgs | undefined;
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.sources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sources'");
            }
            resourceInputs["certName"] = args ? args.certName : undefined;
            resourceInputs["certType"] = args ? args.certType : undefined;
            resourceInputs["checkUrl"] = args ? args.checkUrl : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["forceSet"] = args ? args.forceSet : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["securityToken"] = args ? args.securityToken : undefined;
            resourceInputs["sources"] = args ? args.sources : undefined;
            resourceInputs["sslPri"] = args ? args.sslPri : undefined;
            resourceInputs["sslProtocol"] = args ? args.sslProtocol : undefined;
            resourceInputs["sslPub"] = args ? args.sslPub : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["topLevelDomain"] = args ? args.topLevelDomain : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Domain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Domain resources.
 */
export interface DomainState {
    /**
     * Indicates the name of the certificate if the HTTPS protocol is enabled.
     */
    certName?: pulumi.Input<string>;
    /**
     * The type of the certificate. Valid values:
     * `free`: a free certificate.
     * `cas`: a certificate purchased from Alibaba Cloud SSL Certificates Service.
     * `upload`: a user uploaded certificate.
     */
    certType?: pulumi.Input<string>;
    /**
     * The URL that is used to test the accessibility of the origin.
     */
    checkUrl?: pulumi.Input<string>;
    /**
     * The name of the accelerated domain.
     */
    domainName?: pulumi.Input<string>;
    /**
     * Specifies whether to check the certificate name for duplicates. If you set the value to 1, the system does not perform the check and overwrites the information of the existing certificate with the same name.
     */
    forceSet?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The acceleration region.
     */
    scope?: pulumi.Input<string>;
    /**
     * The top-level domain name.
     */
    securityToken?: pulumi.Input<string>;
    /**
     * The origin information.
     */
    sources?: pulumi.Input<pulumi.Input<inputs.dcdn.DomainSource>[]>;
    /**
     * The private key. Specify this parameter only if you enable the SSL certificate.
     */
    sslPri?: pulumi.Input<string>;
    /**
     * Indicates whether the SSL certificate is enabled. Valid values: `on` enabled, `off` disabled.
     */
    sslProtocol?: pulumi.Input<string>;
    /**
     * Indicates the public key of the certificate if the HTTPS protocol is enabled.
     */
    sslPub?: pulumi.Input<string>;
    /**
     * The status of DCDN Domain. Valid values: `online`, `offline`. Default to `online`.
     */
    status?: pulumi.Input<string>;
    /**
     * The top-level domain name.
     */
    topLevelDomain?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * Indicates the name of the certificate if the HTTPS protocol is enabled.
     */
    certName?: pulumi.Input<string>;
    /**
     * The type of the certificate. Valid values:
     * `free`: a free certificate.
     * `cas`: a certificate purchased from Alibaba Cloud SSL Certificates Service.
     * `upload`: a user uploaded certificate.
     */
    certType?: pulumi.Input<string>;
    /**
     * The URL that is used to test the accessibility of the origin.
     */
    checkUrl?: pulumi.Input<string>;
    /**
     * The name of the accelerated domain.
     */
    domainName: pulumi.Input<string>;
    /**
     * Specifies whether to check the certificate name for duplicates. If you set the value to 1, the system does not perform the check and overwrites the information of the existing certificate with the same name.
     */
    forceSet?: pulumi.Input<string>;
    /**
     * The ID of the resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * The acceleration region.
     */
    scope?: pulumi.Input<string>;
    /**
     * The top-level domain name.
     */
    securityToken?: pulumi.Input<string>;
    /**
     * The origin information.
     */
    sources: pulumi.Input<pulumi.Input<inputs.dcdn.DomainSource>[]>;
    /**
     * The private key. Specify this parameter only if you enable the SSL certificate.
     */
    sslPri?: pulumi.Input<string>;
    /**
     * Indicates whether the SSL certificate is enabled. Valid values: `on` enabled, `off` disabled.
     */
    sslProtocol?: pulumi.Input<string>;
    /**
     * Indicates the public key of the certificate if the HTTPS protocol is enabled.
     */
    sslPub?: pulumi.Input<string>;
    /**
     * The status of DCDN Domain. Valid values: `online`, `offline`. Default to `online`.
     */
    status?: pulumi.Input<string>;
    /**
     * The top-level domain name.
     */
    topLevelDomain?: pulumi.Input<string>;
}
