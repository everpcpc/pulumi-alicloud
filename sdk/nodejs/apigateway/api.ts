// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const apiGroup = new alicloud.apigateway.Group("apiGroup", {description: "description of the api group"});
 * const apiGatewayApi = new alicloud.apigateway.Api("apiGatewayApi", {
 *     groupId: apiGroup.id,
 *     description: "your description",
 *     authType: "APP",
 *     forceNonceCheck: false,
 *     requestConfig: {
 *         protocol: "HTTP",
 *         method: "GET",
 *         path: "/test/path1",
 *         mode: "MAPPING",
 *     },
 *     serviceType: "HTTP",
 *     httpServiceConfig: {
 *         address: "http://apigateway-backend.alicloudapi.com:8080",
 *         method: "GET",
 *         path: "/web/cloudapi",
 *         timeout: 12,
 *         aoneName: "cloudapi-openapi",
 *     },
 *     requestParameters: [{
 *         name: "aaa",
 *         type: "STRING",
 *         required: "OPTIONAL",
 *         "in": "QUERY",
 *         inService: "QUERY",
 *         nameService: "testparams",
 *     }],
 *     stageNames: [
 *         "RELEASE",
 *         "TEST",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Api gateway api can be imported using the id.Format to `<API Group Id>:<API Id>` e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:apigateway/api:Api example "ab2351f2ce904edaa8d92a0510832b91:e4f728fca5a94148b023b99a3e5d0b62"
 * ```
 */
export class Api extends pulumi.CustomResource {
    /**
     * Get an existing Api resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiState, opts?: pulumi.CustomResourceOptions): Api {
        return new Api(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:apigateway/api:Api';

    /**
     * Returns true if the given object is an instance of Api.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Api {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Api.__pulumiType;
    }

    /**
     * The ID of the api of api gateway.
     */
    public /*out*/ readonly apiId!: pulumi.Output<string>;
    /**
     * The authorization Type including APP and ANONYMOUS. Defaults to null.
     */
    public readonly authType!: pulumi.Output<string>;
    /**
     * constant_parameters defines the constant parameters of the api.
     */
    public readonly constantParameters!: pulumi.Output<outputs.apigateway.ApiConstantParameter[] | undefined>;
    /**
     * The description of Constant parameter.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * fc_service_config defines the config when serviceType selected 'FunctionCompute'.
     */
    public readonly fcServiceConfig!: pulumi.Output<outputs.apigateway.ApiFcServiceConfig | undefined>;
    /**
     * Whether to prevent API replay attack. Default value: `false`.
     */
    public readonly forceNonceCheck!: pulumi.Output<boolean>;
    /**
     * The api gateway that the api belongs to. Defaults to null.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * http_service_config defines the config when serviceType selected 'HTTP'.
     */
    public readonly httpServiceConfig!: pulumi.Output<outputs.apigateway.ApiHttpServiceConfig | undefined>;
    /**
     * http_vpc_service_config defines the config when serviceType selected 'HTTP-VPC'.
     */
    public readonly httpVpcServiceConfig!: pulumi.Output<outputs.apigateway.ApiHttpVpcServiceConfig | undefined>;
    /**
     * http_service_config defines the config when serviceType selected 'MOCK'.
     */
    public readonly mockServiceConfig!: pulumi.Output<outputs.apigateway.ApiMockServiceConfig | undefined>;
    /**
     * System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Request_config defines how users can send requests to your API.
     */
    public readonly requestConfig!: pulumi.Output<outputs.apigateway.ApiRequestConfig>;
    /**
     * request_parameters defines the request parameters of the api.
     */
    public readonly requestParameters!: pulumi.Output<outputs.apigateway.ApiRequestParameter[] | undefined>;
    /**
     * The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
     */
    public readonly serviceType!: pulumi.Output<string>;
    /**
     * Stages that the api need to be deployed. Valid value: `RELEASE`,`PRE`,`TEST`.
     */
    public readonly stageNames!: pulumi.Output<string[] | undefined>;
    /**
     * system_parameters defines the system parameters of the api.
     */
    public readonly systemParameters!: pulumi.Output<outputs.apigateway.ApiSystemParameter[] | undefined>;

    /**
     * Create a Api resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiArgs | ApiState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiState | undefined;
            resourceInputs["apiId"] = state ? state.apiId : undefined;
            resourceInputs["authType"] = state ? state.authType : undefined;
            resourceInputs["constantParameters"] = state ? state.constantParameters : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["fcServiceConfig"] = state ? state.fcServiceConfig : undefined;
            resourceInputs["forceNonceCheck"] = state ? state.forceNonceCheck : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["httpServiceConfig"] = state ? state.httpServiceConfig : undefined;
            resourceInputs["httpVpcServiceConfig"] = state ? state.httpVpcServiceConfig : undefined;
            resourceInputs["mockServiceConfig"] = state ? state.mockServiceConfig : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["requestConfig"] = state ? state.requestConfig : undefined;
            resourceInputs["requestParameters"] = state ? state.requestParameters : undefined;
            resourceInputs["serviceType"] = state ? state.serviceType : undefined;
            resourceInputs["stageNames"] = state ? state.stageNames : undefined;
            resourceInputs["systemParameters"] = state ? state.systemParameters : undefined;
        } else {
            const args = argsOrState as ApiArgs | undefined;
            if ((!args || args.authType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authType'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.requestConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requestConfig'");
            }
            if ((!args || args.serviceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceType'");
            }
            resourceInputs["authType"] = args ? args.authType : undefined;
            resourceInputs["constantParameters"] = args ? args.constantParameters : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["fcServiceConfig"] = args ? args.fcServiceConfig : undefined;
            resourceInputs["forceNonceCheck"] = args ? args.forceNonceCheck : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["httpServiceConfig"] = args ? args.httpServiceConfig : undefined;
            resourceInputs["httpVpcServiceConfig"] = args ? args.httpVpcServiceConfig : undefined;
            resourceInputs["mockServiceConfig"] = args ? args.mockServiceConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["requestConfig"] = args ? args.requestConfig : undefined;
            resourceInputs["requestParameters"] = args ? args.requestParameters : undefined;
            resourceInputs["serviceType"] = args ? args.serviceType : undefined;
            resourceInputs["stageNames"] = args ? args.stageNames : undefined;
            resourceInputs["systemParameters"] = args ? args.systemParameters : undefined;
            resourceInputs["apiId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Api.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Api resources.
 */
export interface ApiState {
    /**
     * The ID of the api of api gateway.
     */
    apiId?: pulumi.Input<string>;
    /**
     * The authorization Type including APP and ANONYMOUS. Defaults to null.
     */
    authType?: pulumi.Input<string>;
    /**
     * constant_parameters defines the constant parameters of the api.
     */
    constantParameters?: pulumi.Input<pulumi.Input<inputs.apigateway.ApiConstantParameter>[]>;
    /**
     * The description of Constant parameter.
     */
    description?: pulumi.Input<string>;
    /**
     * fc_service_config defines the config when serviceType selected 'FunctionCompute'.
     */
    fcServiceConfig?: pulumi.Input<inputs.apigateway.ApiFcServiceConfig>;
    /**
     * Whether to prevent API replay attack. Default value: `false`.
     */
    forceNonceCheck?: pulumi.Input<boolean>;
    /**
     * The api gateway that the api belongs to. Defaults to null.
     */
    groupId?: pulumi.Input<string>;
    /**
     * http_service_config defines the config when serviceType selected 'HTTP'.
     */
    httpServiceConfig?: pulumi.Input<inputs.apigateway.ApiHttpServiceConfig>;
    /**
     * http_vpc_service_config defines the config when serviceType selected 'HTTP-VPC'.
     */
    httpVpcServiceConfig?: pulumi.Input<inputs.apigateway.ApiHttpVpcServiceConfig>;
    /**
     * http_service_config defines the config when serviceType selected 'MOCK'.
     */
    mockServiceConfig?: pulumi.Input<inputs.apigateway.ApiMockServiceConfig>;
    /**
     * System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html).
     */
    name?: pulumi.Input<string>;
    /**
     * Request_config defines how users can send requests to your API.
     */
    requestConfig?: pulumi.Input<inputs.apigateway.ApiRequestConfig>;
    /**
     * request_parameters defines the request parameters of the api.
     */
    requestParameters?: pulumi.Input<pulumi.Input<inputs.apigateway.ApiRequestParameter>[]>;
    /**
     * The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
     */
    serviceType?: pulumi.Input<string>;
    /**
     * Stages that the api need to be deployed. Valid value: `RELEASE`,`PRE`,`TEST`.
     */
    stageNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * system_parameters defines the system parameters of the api.
     */
    systemParameters?: pulumi.Input<pulumi.Input<inputs.apigateway.ApiSystemParameter>[]>;
}

/**
 * The set of arguments for constructing a Api resource.
 */
export interface ApiArgs {
    /**
     * The authorization Type including APP and ANONYMOUS. Defaults to null.
     */
    authType: pulumi.Input<string>;
    /**
     * constant_parameters defines the constant parameters of the api.
     */
    constantParameters?: pulumi.Input<pulumi.Input<inputs.apigateway.ApiConstantParameter>[]>;
    /**
     * The description of Constant parameter.
     */
    description: pulumi.Input<string>;
    /**
     * fc_service_config defines the config when serviceType selected 'FunctionCompute'.
     */
    fcServiceConfig?: pulumi.Input<inputs.apigateway.ApiFcServiceConfig>;
    /**
     * Whether to prevent API replay attack. Default value: `false`.
     */
    forceNonceCheck?: pulumi.Input<boolean>;
    /**
     * The api gateway that the api belongs to. Defaults to null.
     */
    groupId: pulumi.Input<string>;
    /**
     * http_service_config defines the config when serviceType selected 'HTTP'.
     */
    httpServiceConfig?: pulumi.Input<inputs.apigateway.ApiHttpServiceConfig>;
    /**
     * http_vpc_service_config defines the config when serviceType selected 'HTTP-VPC'.
     */
    httpVpcServiceConfig?: pulumi.Input<inputs.apigateway.ApiHttpVpcServiceConfig>;
    /**
     * http_service_config defines the config when serviceType selected 'MOCK'.
     */
    mockServiceConfig?: pulumi.Input<inputs.apigateway.ApiMockServiceConfig>;
    /**
     * System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html).
     */
    name?: pulumi.Input<string>;
    /**
     * Request_config defines how users can send requests to your API.
     */
    requestConfig: pulumi.Input<inputs.apigateway.ApiRequestConfig>;
    /**
     * request_parameters defines the request parameters of the api.
     */
    requestParameters?: pulumi.Input<pulumi.Input<inputs.apigateway.ApiRequestParameter>[]>;
    /**
     * The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
     */
    serviceType: pulumi.Input<string>;
    /**
     * Stages that the api need to be deployed. Valid value: `RELEASE`,`PRE`,`TEST`.
     */
    stageNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * system_parameters defines the system parameters of the api.
     */
    systemParameters?: pulumi.Input<pulumi.Input<inputs.apigateway.ApiSystemParameter>[]>;
}
