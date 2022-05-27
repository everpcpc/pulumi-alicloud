// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Application Load Balancer (ALB) Listener resource.
 *
 * For information about Application Load Balancer (ALB) Listener and how to use it, see [What is Listener](https://www.alibabacloud.com/help/doc-detail/214348.htm).
 *
 * > **NOTE:** Available in v1.133.0+.
 *
 * ## Import
 *
 * Application Load Balancer (ALB) Listener can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:alb/listener:Listener example <id>
 * ```
 */
export class Listener extends pulumi.CustomResource {
    /**
     * Get an existing Listener resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerState, opts?: pulumi.CustomResourceOptions): Listener {
        return new Listener(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:alb/listener:Listener';

    /**
     * Returns true if the given object is an instance of Listener.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Listener {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Listener.__pulumiType;
    }

    /**
     * Indicates whether the access log has a custom header field. Valid values: true and false. Default value: false.
     */
    public readonly accessLogRecordCustomizedHeadersEnabled!: pulumi.Output<boolean>;
    /**
     * Xtrace Configuration Information. See the following `Block accessLogTracingConfig`.
     */
    public readonly accessLogTracingConfig!: pulumi.Output<outputs.alb.ListenerAccessLogTracingConfig | undefined>;
    /**
     * The configurations of the access control lists (ACLs). See the following `Block aclConfig`. **NOTE:** Field `aclConfig` has been deprecated from provider version 1.163.0, and it will be removed in the future version. Please use the new resource `alicloud.alb.ListenerAclAttachment`.,
     *
     * @deprecated Field 'acl_config' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_alb_listener_acl_attachment'.
     */
    public readonly aclConfig!: pulumi.Output<outputs.alb.ListenerAclConfig>;
    /**
     * The default certificate of the Listener. See the following `Block certificates`. **NOTE:** When `listenerProtocol` is `HTTPS`, The default certificate must be set one。
     */
    public readonly certificates!: pulumi.Output<outputs.alb.ListenerCertificates | undefined>;
    /**
     * The Default Rule Action List. See the following `Block defaultActions`.
     */
    public readonly defaultActions!: pulumi.Output<outputs.alb.ListenerDefaultAction[] | undefined>;
    /**
     * The dry run.
     */
    public readonly dryRun!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to Enable Gzip Compression, as a Specific File Type on a Compression. Valid values: `false`, `true`. Default Value: `true`. .
     */
    public readonly gzipEnabled!: pulumi.Output<boolean>;
    /**
     * Whether to Enable HTTP/2 Features. Valid Values: `True` Or `False`. Default Value: `True`.
     */
    public readonly http2Enabled!: pulumi.Output<boolean>;
    /**
     * Specify the Connection Idle Timeout Value: `1` to `60`. Unit: Seconds.
     */
    public readonly idleTimeout!: pulumi.Output<number>;
    /**
     * The description of the listener. The description must be 2 to 256 characters in length. The name can contain only the characters in the following string: `/^([^\x00-\xff]|[\w.,;/@-]){2,256}$/`.
     */
    public readonly listenerDescription!: pulumi.Output<string | undefined>;
    /**
     * The ALB Instance Front-End, and Those of the Ports Used. Value: `1` to `65535`.
     */
    public readonly listenerPort!: pulumi.Output<number>;
    /**
     * Snooping Protocols. Valid Values: `HTTP`, `HTTPS` Or `QUIC`.
     */
    public readonly listenerProtocol!: pulumi.Output<string>;
    /**
     * The ALB Instance Id.
     */
    public readonly loadBalancerId!: pulumi.Output<string>;
    /**
     * Configuration Associated with the QuIC Listening. See the following `Block quicConfig`.
     */
    public readonly quicConfig!: pulumi.Output<outputs.alb.ListenerQuicConfig>;
    /**
     * The Specified Request Timeout Time. Value: `1` to `180`. Unit: Seconds. Default Value: `60`. If the Timeout Time Within the Back-End Server Has Not Answered the ALB Will Give up Waiting, the Client Returns the HTTP 504 Error Code.
     */
    public readonly requestTimeout!: pulumi.Output<number>;
    /**
     * Security Policy.
     */
    public readonly securityPolicyId!: pulumi.Output<string>;
    /**
     * The state of the listener. Valid Values: `Running` Or `Stopped`. Valid values: `Running`: The listener is running. `Stopped`: The listener is stopped.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * The `xForwardFor` Related Attribute Configuration. See the following `Block xForwardedForConfig`. **NOTE:** The attribute is valid when the attribute `listenerProtocol` is `HTTPS`.
     */
    public readonly xForwardedForConfig!: pulumi.Output<outputs.alb.ListenerXForwardedForConfig>;

    /**
     * Create a Listener resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListenerArgs | ListenerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ListenerState | undefined;
            resourceInputs["accessLogRecordCustomizedHeadersEnabled"] = state ? state.accessLogRecordCustomizedHeadersEnabled : undefined;
            resourceInputs["accessLogTracingConfig"] = state ? state.accessLogTracingConfig : undefined;
            resourceInputs["aclConfig"] = state ? state.aclConfig : undefined;
            resourceInputs["certificates"] = state ? state.certificates : undefined;
            resourceInputs["defaultActions"] = state ? state.defaultActions : undefined;
            resourceInputs["dryRun"] = state ? state.dryRun : undefined;
            resourceInputs["gzipEnabled"] = state ? state.gzipEnabled : undefined;
            resourceInputs["http2Enabled"] = state ? state.http2Enabled : undefined;
            resourceInputs["idleTimeout"] = state ? state.idleTimeout : undefined;
            resourceInputs["listenerDescription"] = state ? state.listenerDescription : undefined;
            resourceInputs["listenerPort"] = state ? state.listenerPort : undefined;
            resourceInputs["listenerProtocol"] = state ? state.listenerProtocol : undefined;
            resourceInputs["loadBalancerId"] = state ? state.loadBalancerId : undefined;
            resourceInputs["quicConfig"] = state ? state.quicConfig : undefined;
            resourceInputs["requestTimeout"] = state ? state.requestTimeout : undefined;
            resourceInputs["securityPolicyId"] = state ? state.securityPolicyId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["xForwardedForConfig"] = state ? state.xForwardedForConfig : undefined;
        } else {
            const args = argsOrState as ListenerArgs | undefined;
            if ((!args || args.listenerPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerPort'");
            }
            if ((!args || args.listenerProtocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerProtocol'");
            }
            if ((!args || args.loadBalancerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerId'");
            }
            resourceInputs["accessLogRecordCustomizedHeadersEnabled"] = args ? args.accessLogRecordCustomizedHeadersEnabled : undefined;
            resourceInputs["accessLogTracingConfig"] = args ? args.accessLogTracingConfig : undefined;
            resourceInputs["aclConfig"] = args ? args.aclConfig : undefined;
            resourceInputs["certificates"] = args ? args.certificates : undefined;
            resourceInputs["defaultActions"] = args ? args.defaultActions : undefined;
            resourceInputs["dryRun"] = args ? args.dryRun : undefined;
            resourceInputs["gzipEnabled"] = args ? args.gzipEnabled : undefined;
            resourceInputs["http2Enabled"] = args ? args.http2Enabled : undefined;
            resourceInputs["idleTimeout"] = args ? args.idleTimeout : undefined;
            resourceInputs["listenerDescription"] = args ? args.listenerDescription : undefined;
            resourceInputs["listenerPort"] = args ? args.listenerPort : undefined;
            resourceInputs["listenerProtocol"] = args ? args.listenerProtocol : undefined;
            resourceInputs["loadBalancerId"] = args ? args.loadBalancerId : undefined;
            resourceInputs["quicConfig"] = args ? args.quicConfig : undefined;
            resourceInputs["requestTimeout"] = args ? args.requestTimeout : undefined;
            resourceInputs["securityPolicyId"] = args ? args.securityPolicyId : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["xForwardedForConfig"] = args ? args.xForwardedForConfig : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Listener.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Listener resources.
 */
export interface ListenerState {
    /**
     * Indicates whether the access log has a custom header field. Valid values: true and false. Default value: false.
     */
    accessLogRecordCustomizedHeadersEnabled?: pulumi.Input<boolean>;
    /**
     * Xtrace Configuration Information. See the following `Block accessLogTracingConfig`.
     */
    accessLogTracingConfig?: pulumi.Input<inputs.alb.ListenerAccessLogTracingConfig>;
    /**
     * The configurations of the access control lists (ACLs). See the following `Block aclConfig`. **NOTE:** Field `aclConfig` has been deprecated from provider version 1.163.0, and it will be removed in the future version. Please use the new resource `alicloud.alb.ListenerAclAttachment`.,
     *
     * @deprecated Field 'acl_config' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_alb_listener_acl_attachment'.
     */
    aclConfig?: pulumi.Input<inputs.alb.ListenerAclConfig>;
    /**
     * The default certificate of the Listener. See the following `Block certificates`. **NOTE:** When `listenerProtocol` is `HTTPS`, The default certificate must be set one。
     */
    certificates?: pulumi.Input<inputs.alb.ListenerCertificates>;
    /**
     * The Default Rule Action List. See the following `Block defaultActions`.
     */
    defaultActions?: pulumi.Input<pulumi.Input<inputs.alb.ListenerDefaultAction>[]>;
    /**
     * The dry run.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Whether to Enable Gzip Compression, as a Specific File Type on a Compression. Valid values: `false`, `true`. Default Value: `true`. .
     */
    gzipEnabled?: pulumi.Input<boolean>;
    /**
     * Whether to Enable HTTP/2 Features. Valid Values: `True` Or `False`. Default Value: `True`.
     */
    http2Enabled?: pulumi.Input<boolean>;
    /**
     * Specify the Connection Idle Timeout Value: `1` to `60`. Unit: Seconds.
     */
    idleTimeout?: pulumi.Input<number>;
    /**
     * The description of the listener. The description must be 2 to 256 characters in length. The name can contain only the characters in the following string: `/^([^\x00-\xff]|[\w.,;/@-]){2,256}$/`.
     */
    listenerDescription?: pulumi.Input<string>;
    /**
     * The ALB Instance Front-End, and Those of the Ports Used. Value: `1` to `65535`.
     */
    listenerPort?: pulumi.Input<number>;
    /**
     * Snooping Protocols. Valid Values: `HTTP`, `HTTPS` Or `QUIC`.
     */
    listenerProtocol?: pulumi.Input<string>;
    /**
     * The ALB Instance Id.
     */
    loadBalancerId?: pulumi.Input<string>;
    /**
     * Configuration Associated with the QuIC Listening. See the following `Block quicConfig`.
     */
    quicConfig?: pulumi.Input<inputs.alb.ListenerQuicConfig>;
    /**
     * The Specified Request Timeout Time. Value: `1` to `180`. Unit: Seconds. Default Value: `60`. If the Timeout Time Within the Back-End Server Has Not Answered the ALB Will Give up Waiting, the Client Returns the HTTP 504 Error Code.
     */
    requestTimeout?: pulumi.Input<number>;
    /**
     * Security Policy.
     */
    securityPolicyId?: pulumi.Input<string>;
    /**
     * The state of the listener. Valid Values: `Running` Or `Stopped`. Valid values: `Running`: The listener is running. `Stopped`: The listener is stopped.
     */
    status?: pulumi.Input<string>;
    /**
     * The `xForwardFor` Related Attribute Configuration. See the following `Block xForwardedForConfig`. **NOTE:** The attribute is valid when the attribute `listenerProtocol` is `HTTPS`.
     */
    xForwardedForConfig?: pulumi.Input<inputs.alb.ListenerXForwardedForConfig>;
}

/**
 * The set of arguments for constructing a Listener resource.
 */
export interface ListenerArgs {
    /**
     * Indicates whether the access log has a custom header field. Valid values: true and false. Default value: false.
     */
    accessLogRecordCustomizedHeadersEnabled?: pulumi.Input<boolean>;
    /**
     * Xtrace Configuration Information. See the following `Block accessLogTracingConfig`.
     */
    accessLogTracingConfig?: pulumi.Input<inputs.alb.ListenerAccessLogTracingConfig>;
    /**
     * The configurations of the access control lists (ACLs). See the following `Block aclConfig`. **NOTE:** Field `aclConfig` has been deprecated from provider version 1.163.0, and it will be removed in the future version. Please use the new resource `alicloud.alb.ListenerAclAttachment`.,
     *
     * @deprecated Field 'acl_config' has been deprecated from provider version 1.163.0 and it will be removed in the future version. Please use the new resource 'alicloud_alb_listener_acl_attachment'.
     */
    aclConfig?: pulumi.Input<inputs.alb.ListenerAclConfig>;
    /**
     * The default certificate of the Listener. See the following `Block certificates`. **NOTE:** When `listenerProtocol` is `HTTPS`, The default certificate must be set one。
     */
    certificates?: pulumi.Input<inputs.alb.ListenerCertificates>;
    /**
     * The Default Rule Action List. See the following `Block defaultActions`.
     */
    defaultActions?: pulumi.Input<pulumi.Input<inputs.alb.ListenerDefaultAction>[]>;
    /**
     * The dry run.
     */
    dryRun?: pulumi.Input<boolean>;
    /**
     * Whether to Enable Gzip Compression, as a Specific File Type on a Compression. Valid values: `false`, `true`. Default Value: `true`. .
     */
    gzipEnabled?: pulumi.Input<boolean>;
    /**
     * Whether to Enable HTTP/2 Features. Valid Values: `True` Or `False`. Default Value: `True`.
     */
    http2Enabled?: pulumi.Input<boolean>;
    /**
     * Specify the Connection Idle Timeout Value: `1` to `60`. Unit: Seconds.
     */
    idleTimeout?: pulumi.Input<number>;
    /**
     * The description of the listener. The description must be 2 to 256 characters in length. The name can contain only the characters in the following string: `/^([^\x00-\xff]|[\w.,;/@-]){2,256}$/`.
     */
    listenerDescription?: pulumi.Input<string>;
    /**
     * The ALB Instance Front-End, and Those of the Ports Used. Value: `1` to `65535`.
     */
    listenerPort: pulumi.Input<number>;
    /**
     * Snooping Protocols. Valid Values: `HTTP`, `HTTPS` Or `QUIC`.
     */
    listenerProtocol: pulumi.Input<string>;
    /**
     * The ALB Instance Id.
     */
    loadBalancerId: pulumi.Input<string>;
    /**
     * Configuration Associated with the QuIC Listening. See the following `Block quicConfig`.
     */
    quicConfig?: pulumi.Input<inputs.alb.ListenerQuicConfig>;
    /**
     * The Specified Request Timeout Time. Value: `1` to `180`. Unit: Seconds. Default Value: `60`. If the Timeout Time Within the Back-End Server Has Not Answered the ALB Will Give up Waiting, the Client Returns the HTTP 504 Error Code.
     */
    requestTimeout?: pulumi.Input<number>;
    /**
     * Security Policy.
     */
    securityPolicyId?: pulumi.Input<string>;
    /**
     * The state of the listener. Valid Values: `Running` Or `Stopped`. Valid values: `Running`: The listener is running. `Stopped`: The listener is stopped.
     */
    status?: pulumi.Input<string>;
    /**
     * The `xForwardFor` Related Attribute Configuration. See the following `Block xForwardedForConfig`. **NOTE:** The attribute is valid when the attribute `listenerProtocol` is `HTTPS`.
     */
    xForwardedForConfig?: pulumi.Input<inputs.alb.ListenerXForwardedForConfig>;
}
