// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an Alicloud Function Compute custom domain resource.
 *  For the detailed information, please refer to the [developer guide](https://www.alibabacloud.com/help/doc-detail/90759.htm).
 *
 * > **NOTE:** Available in 1.98.0+
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "tf-testaccalicloudfcservice";
 * const defaultService = new alicloud.fc.Service("defaultService", {description: `${name}-description`});
 * const defaultBucket = new alicloud.oss.Bucket("defaultBucket", {bucket: name});
 * const defaultBucketObject = new alicloud.oss.BucketObject("defaultBucketObject", {
 *     bucket: defaultBucket.id,
 *     key: "fc/hello.zip",
 *     content: `		# -*- coding: utf-8 -*-
 * 	def handler(event, context):
 * 		print "hello world"
 * 		return 'hello world'
 * `,
 * });
 * const defaultFunction = new alicloud.fc.Function("defaultFunction", {
 *     service: defaultService.name,
 *     ossBucket: defaultBucket.id,
 *     ossKey: defaultBucketObject.key,
 *     memorySize: 512,
 *     runtime: "python2.7",
 *     handler: "hello.handler",
 * });
 * const defaultCustomDomain = new alicloud.fc.CustomDomain("defaultCustomDomain", {
 *     domainName: "terraform.functioncompute.com",
 *     protocol: "HTTP",
 *     routeConfigs: [{
 *         path: "/login/*",
 *         serviceName: defaultService.name,
 *         functionName: defaultFunction.name,
 *         qualifier: "v1",
 *         methods: [
 *             "GET",
 *             "POST",
 *         ],
 *     }],
 *     certConfig: {
 *         certName: "your certificate name",
 *         privateKey: "your private key",
 *         certificate: "your certificate data",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Function Compute custom domain can be imported using the id or the domain name, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:fc/customDomain:CustomDomain foo my-fc-custom-domain
 * ```
 */
export class CustomDomain extends pulumi.CustomResource {
    /**
     * Get an existing CustomDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomDomainState, opts?: pulumi.CustomResourceOptions): CustomDomain {
        return new CustomDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:fc/customDomain:CustomDomain';

    /**
     * Returns true if the given object is an instance of CustomDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomDomain.__pulumiType;
    }

    /**
     * The account id.
     */
    public /*out*/ readonly accountId!: pulumi.Output<string>;
    /**
     * The api version of Function Compute.
     */
    public /*out*/ readonly apiVersion!: pulumi.Output<string>;
    /**
     * The configuration of HTTPS certificate.
     */
    public readonly certConfig!: pulumi.Output<outputs.fc.CustomDomainCertConfig | undefined>;
    /**
     * The date this resource was created.
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * The custom domain name. For example, "example.com".
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The date this resource was last modified.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;
    /**
     * The protocol, `HTTP` or `HTTP,HTTPS`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The configuration of domain route, mapping the path and Function Compute function.
     */
    public readonly routeConfigs!: pulumi.Output<outputs.fc.CustomDomainRouteConfig[] | undefined>;

    /**
     * Create a CustomDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomDomainArgs | CustomDomainState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CustomDomainState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
            inputs["apiVersion"] = state ? state.apiVersion : undefined;
            inputs["certConfig"] = state ? state.certConfig : undefined;
            inputs["createdTime"] = state ? state.createdTime : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["lastModifiedTime"] = state ? state.lastModifiedTime : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["routeConfigs"] = state ? state.routeConfigs : undefined;
        } else {
            const args = argsOrState as CustomDomainArgs | undefined;
            if ((!args || args.domainName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.protocol === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'protocol'");
            }
            inputs["certConfig"] = args ? args.certConfig : undefined;
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["routeConfigs"] = args ? args.routeConfigs : undefined;
            inputs["accountId"] = undefined /*out*/;
            inputs["apiVersion"] = undefined /*out*/;
            inputs["createdTime"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CustomDomain.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomDomain resources.
 */
export interface CustomDomainState {
    /**
     * The account id.
     */
    readonly accountId?: pulumi.Input<string>;
    /**
     * The api version of Function Compute.
     */
    readonly apiVersion?: pulumi.Input<string>;
    /**
     * The configuration of HTTPS certificate.
     */
    readonly certConfig?: pulumi.Input<inputs.fc.CustomDomainCertConfig>;
    /**
     * The date this resource was created.
     */
    readonly createdTime?: pulumi.Input<string>;
    /**
     * The custom domain name. For example, "example.com".
     */
    readonly domainName?: pulumi.Input<string>;
    /**
     * The date this resource was last modified.
     */
    readonly lastModifiedTime?: pulumi.Input<string>;
    /**
     * The protocol, `HTTP` or `HTTP,HTTPS`.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * The configuration of domain route, mapping the path and Function Compute function.
     */
    readonly routeConfigs?: pulumi.Input<pulumi.Input<inputs.fc.CustomDomainRouteConfig>[]>;
}

/**
 * The set of arguments for constructing a CustomDomain resource.
 */
export interface CustomDomainArgs {
    /**
     * The configuration of HTTPS certificate.
     */
    readonly certConfig?: pulumi.Input<inputs.fc.CustomDomainCertConfig>;
    /**
     * The custom domain name. For example, "example.com".
     */
    readonly domainName: pulumi.Input<string>;
    /**
     * The protocol, `HTTP` or `HTTP,HTTPS`.
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * The configuration of domain route, mapping the path and Function Compute function.
     */
    readonly routeConfigs?: pulumi.Input<pulumi.Input<inputs.fc.CustomDomainRouteConfig>[]>;
}
