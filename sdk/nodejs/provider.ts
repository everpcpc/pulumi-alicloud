// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * The provider type for the alicloud package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'alicloud';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }


    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            inputs["accessKey"] = args ? args.accessKey : undefined;
            inputs["accountId"] = args ? args.accountId : undefined;
            inputs["assumeRole"] = pulumi.output(args ? args.assumeRole : undefined).apply(JSON.stringify);
            inputs["clientConnectTimeout"] = pulumi.output(args ? args.clientConnectTimeout : undefined).apply(JSON.stringify);
            inputs["clientReadTimeout"] = pulumi.output(args ? args.clientReadTimeout : undefined).apply(JSON.stringify);
            inputs["configurationSource"] = args ? args.configurationSource : undefined;
            inputs["credentialsUri"] = args ? args.credentialsUri : undefined;
            inputs["ecsRoleName"] = (args ? args.ecsRoleName : undefined) ?? utilities.getEnv("ALICLOUD_ECS_ROLE_NAME");
            inputs["endpoints"] = pulumi.output(args ? args.endpoints : undefined).apply(JSON.stringify);
            inputs["fc"] = args ? args.fc : undefined;
            inputs["logEndpoint"] = args ? args.logEndpoint : undefined;
            inputs["mnsEndpoint"] = args ? args.mnsEndpoint : undefined;
            inputs["otsInstanceName"] = args ? args.otsInstanceName : undefined;
            inputs["profile"] = (args ? args.profile : undefined) ?? utilities.getEnv("ALICLOUD_PROFILE");
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["region"] = (args ? args.region : undefined) ?? utilities.getEnv("ALICLOUD_REGION");
            inputs["secretKey"] = args ? args.secretKey : undefined;
            inputs["secureTransport"] = args ? args.secureTransport : undefined;
            inputs["securityToken"] = args ? args.securityToken : undefined;
            inputs["securityTransport"] = args ? args.securityTransport : undefined;
            inputs["sharedCredentialsFile"] = args ? args.sharedCredentialsFile : undefined;
            inputs["skipRegionValidation"] = pulumi.output(args ? args.skipRegionValidation : undefined).apply(JSON.stringify);
            inputs["sourceIp"] = args ? args.sourceIp : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Provider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * The access key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
     * console.
     */
    readonly accessKey?: pulumi.Input<string>;
    /**
     * The account ID for some service API operations. You can retrieve this from the 'Security Settings' section of the
     * Alibaba Cloud console.
     */
    readonly accountId?: pulumi.Input<string>;
    readonly assumeRole?: pulumi.Input<inputs.ProviderAssumeRole>;
    /**
     * The maximum timeout of the client connection server.
     */
    readonly clientConnectTimeout?: pulumi.Input<number>;
    /**
     * The maximum timeout of the client read request.
     */
    readonly clientReadTimeout?: pulumi.Input<number>;
    /**
     * Use this to mark a terraform configuration file source.
     */
    readonly configurationSource?: pulumi.Input<string>;
    /**
     * The URI of sidecar credentials service.
     */
    readonly credentialsUri?: pulumi.Input<string>;
    /**
     * The RAM Role Name attached on a ECS instance for API operations. You can retrieve this from the 'Access Control' section
     * of the Alibaba Cloud console.
     */
    readonly ecsRoleName?: pulumi.Input<string>;
    readonly endpoints?: pulumi.Input<pulumi.Input<inputs.ProviderEndpoint>[]>;
    /**
     * @deprecated Field 'fc' has been deprecated from provider version 1.28.0. New field 'fc' which in nested endpoints instead.
     */
    readonly fc?: pulumi.Input<string>;
    /**
     * @deprecated Field 'log_endpoint' has been deprecated from provider version 1.28.0. New field 'log' which in nested endpoints instead.
     */
    readonly logEndpoint?: pulumi.Input<string>;
    /**
     * @deprecated Field 'mns_endpoint' has been deprecated from provider version 1.28.0. New field 'mns' which in nested endpoints instead.
     */
    readonly mnsEndpoint?: pulumi.Input<string>;
    /**
     * @deprecated Field 'ots_instance_name' has been deprecated from provider version 1.10.0. New field 'instance_name' of resource 'alicloud_ots_table' instead.
     */
    readonly otsInstanceName?: pulumi.Input<string>;
    /**
     * The profile for API operations. If not set, the default profile created with `aliyun configure` will be used.
     */
    readonly profile?: pulumi.Input<string>;
    readonly protocol?: pulumi.Input<string>;
    /**
     * The region where Alibaba Cloud operations will take place. Examples are cn-beijing, cn-hangzhou, eu-central-1, etc.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The secret key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
     * console.
     */
    readonly secretKey?: pulumi.Input<string>;
    /**
     * The security transport for the assume role invoking.
     */
    readonly secureTransport?: pulumi.Input<string>;
    /**
     * security token. A security token is only required if you are using Security Token Service.
     */
    readonly securityToken?: pulumi.Input<string>;
    readonly securityTransport?: pulumi.Input<string>;
    /**
     * The path to the shared credentials file. If not set this defaults to ~/.aliyun/config.json
     */
    readonly sharedCredentialsFile?: pulumi.Input<string>;
    /**
     * Skip static validation of region ID. Used by users of alternative AlibabaCloud-like APIs or users w/ access to regions
     * that are not public (yet).
     */
    readonly skipRegionValidation?: pulumi.Input<boolean>;
    /**
     * The source ip for the assume role invoking.
     */
    readonly sourceIp?: pulumi.Input<string>;
}
