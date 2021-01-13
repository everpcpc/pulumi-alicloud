// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * HTTPS listeners of guaranteed-performance SLB support configuring multiple certificates, allowing you to forward requests with different domain names to different backend servers.
 * Please refer to the [documentation](https://www.alibabacloud.com/help/doc-detail/85956.htm?spm=a2c63.p38356.b99.40.1c881563Co8p6w) for details.
 *
 * > **NOTE:** Available in 1.60.0+
 *
 * > **NOTE:** The instance with shared loadBalancerSpec doesn't support domainExtension.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Create a new load balancer and domain extension
 * const instance = new alicloud.slb.LoadBalancer("instance", {
 *     internetChargeType: "PayByTraffic",
 *     internet: "true",
 * });
 * const foo = new alicloud.slb.ServerCertificate("foo", {
 *     serverCertificate: `-----BEGIN CERTIFICATE-----
 * MIIDdjCCAl4CCQCcm+erkcKN7DANBgkqhkiG9w0BAQsFADB9MQswCQYDVQQGEwJj
 * bjELMAkGA1UECAwCYmoxEDAOBgNVBAcMB2JlaWppbmcxDzANBgNVBAoMBmFsaXl1
 * bjELMAkGA1UECwwCc2MxFTATBgNVBAMMDHd3dy50ZXN0LmNvbTEaMBgGCSqGSIb3
 * DQEJARYLMTIzQDEyMy5jb20wHhcNMTkwNDI2MDM0ODAxWhcNMjQwNDI1MDM0ODAx
 * WjB9MQswCQYDVQQGEwJjbjELMAkGA1UECAwCYmoxEDAOBgNVBAcMB2JlaWppbmcx
 * DzANBgNVBAoMBmFsaXl1bjELMAkGA1UECwwCc2MxFTATBgNVBAMMDHd3dy50ZXN0
 * LmNvbTEaMBgGCSqGSIb3DQEJARYLMTIzQDEyMy5jb20wggEiMA0GCSqGSIb3DQEB
 * AQUAA4IBDwAwggEKAoIBAQDKMKF5qmN/uoMjdH3D8aPRcUOA0s8rZpYhG8zbkF1j
 * 8gHYoB/FDvM7G7dfVsyjbMwLOxKvAhWvHHSpEz/t7gB+QdwrAMiMJwGmtCnXrh2E
 * WiXgalMe1y4a/T5R7q+m4T1zFATf+kbnHWfkSGF4W7b6UBoaH+9StQ95CnqzNf/2
 * p/Of7+S0XzCxFXw8GIVzZk0xFe6lHJzaq06f3mvzrD+4rpO56tTUvrgTY/n61gsF
 * ZP7f0CJ2JQh6eNRFOEUSfxKu/Dy/+IsQxorCJY2Q59ZAf3rXrqDN104jw9PlwnLl
 * qfZz3RMODN6BWjxE8rvRtT0qMfuAfv1gjBdWZN0hUYBRAgMBAAEwDQYJKoZIhvcN
 * AQELBQADggEBAABzo82TxGp5poVkd5pLWj5ACgcBv8Cs6oH9D+4Jz9BmyuBUsQXh
 * 2aG0hQAe1mU61C9konsl/GTW8umJQ4M4lYEztXXwMf5PlBMGwebM0ZbSGg6jKtZg
 * WCgJ3eP/FMmyXGL5Jji5+e09eObhUDVle4tdi0On97zBoz85W02rgWFAqZJwiEAP
 * t+c7jX7uOSBq2/38iGStlrX5yB1at/gJXXiA5CL5OtlR3Okvb0/QH37efO1Nu39m
 * lFi0ODPAVyXjVypAiLguDxPn6AtDTdk9Iw9B19OD4NrzNRWgSSX5vuxo/VcRcgWk
 * 3gEe9Ca0ZKN20q9XgthAiFFjl1S9ZgdA6Zc=
 * -----END CERTIFICATE-----`,
 *     privateKey: `-----BEGIN RSA PRIVATE KEY-----
 * MIIEowIBAAKCAQEAyjCheapjf7qDI3R9w/Gj0XFDgNLPK2aWIRvM25BdY/IB2KAf
 * xQ7zOxu3X1bMo2zMCzsSrwIVrxx0qRM/7e4AfkHcKwDIjCcBprQp164dhFol4GpT
 * HtcuGv0+Ue6vpuE9cxQE3/pG5x1n5EhheFu2+lAaGh/vUrUPeQp6szX/9qfzn+/k
 * tF8wsRV8PBiFc2ZNMRXupRyc2qtOn95r86w/uK6TuerU1L64E2P5+tYLBWT+39Ai
 * diUIenjURThFEn8Srvw8v/iLEMaKwiWNkOfWQH96166gzddOI8PT5cJy5an2c90T
 * DgzegVo8RPK70bU9KjH7gH79YIwXVmTdIVGAUQIDAQABAoIBAE1J4a/8biR5S3/W
 * G+03BYQeY8tuyjqw8FqfoeOcf9agwAvqybouSNQjeCk9qOQfxq/UWQQFK/zQR9gJ
 * v7pX7GBXFK5rkj3g+0SaQhRsPmRFgY0Tl8qGPt2aSKRRNVv5ZeADmwlzRn86QmiF
 * Mp0rkfqFfDTYWEepZszCML0ouzuxsW/9tq7rvtSjsgATNt31B3vFa3D3JBi31jUh
 * 5nfR9A3bATze7mQw3byEDiVl5ASRDgYyur403P1fDnMy9DBHZ8NaPOsFF6OBpJal
 * BJsG5z00hll5PFN2jfmBQKlvAeU7wfwqdaSnGHOfqf2DeTTaFjIQ4gUhRn/m6pLo
 * 6kXttLECgYEA9sng0Qz/TcPFfM4tQ1gyvB1cKnnGIwg1FP8sfUjbbEgjaHhA224S
 * k3BxtX2Kq6fhTXuwusAFc6OVMAZ76FgrQ5K4Ci7+DTsrF28z4b8td+p+lO/DxgP9
 * lTgN+ddsiTOV4fUef9Z3yY0Zr0CnBUMbQYRaV2UIbCdiB0G4V/bt9TsCgYEA0bya
 * Oo9wGI0RJV0bYP7qwO74Ra1/i1viWbRlS7jU37Q+AZstrlKcQ5CTPzOjKFKMiUzl
 * 4miWacZ0/q2n+Mvd7NbXGXTLijahnyOYKaHJYyh4oBymfkgAifRstE0Ki9gdvArb
 * /I+emC0GvLSyfGN8UUeDJs4NmqdEXGqjo2JOV+MCgYALFv1MR5o9Y1u/hQBRs2fs
 * PiGDIx+9OUQxYloccyaxEfjNXAIGGkcpavchIbgWiJ++PJ2vdquIC8TLeK8evL+M
 * 9M3iX0Q5UfxYvD2HmnCvn9D6Xl/cyRcfGnq+TGjrLW9BzSMGuZt+aiHKV0xqFx7l
 * bc4leTvMqGRmURS4lzcQOwKBgQCDzA/i4sYfN25h21tcHXSpnsG3D2rJyQi5NCo/
 * ZjunA92/JqOTGuiFcLGHEszhhtY3ZXJET1LNz18vtzKJnpqrvOnYXlOVW/U+SqDQ
 * 8JDb1c/PVZGuY1KrXkR9HLiW3kz5IJ3S3PFdUVYdeTN8BQxXCyg4V12nJJtJs912
 * y0zN3wKBgGDS6YttCN6aI4EOABYE8fI1EYQ7vhfiYsaWGWSR1l6bQey7KR6M1ACz
 * ZzMASNyytVt12yXE4/Emv6/pYqigbDLfL1zQJSLJ3EHJYTh2RxjR+AaGDudYFG/T
 * liQ9YXhV5Iu2x1pNwrtFnssDdaaGpfA7l3xC00BL7Z+SAJyI4QKA
 * -----END RSA PRIVATE KEY-----`,
 * });
 * const https = new alicloud.slb.Listener("https", {
 *     loadBalancerId: instance.id,
 *     backendPort: 80,
 *     frontendPort: 443,
 *     protocol: "https",
 *     stickySession: "on",
 *     stickySessionType: "insert",
 *     cookie: "testslblistenercookie",
 *     cookieTimeout: 86400,
 *     healthCheck: "on",
 *     healthCheckUri: "/cons",
 *     healthCheckConnectPort: 20,
 *     healthyThreshold: 8,
 *     unhealthyThreshold: 8,
 *     healthCheckTimeout: 8,
 *     healthCheckInterval: 5,
 *     healthCheckHttpCode: "http_2xx,http_3xx",
 *     bandwidth: 10,
 *     sslCertificateId: foo.id,
 * });
 * const example1 = new alicloud.slb.DomainExtension("example1", {
 *     loadBalancerId: instance.id,
 *     frontendPort: https.frontendPort,
 *     domain: "www.test.com",
 *     serverCertificateId: foo.id,
 * });
 * ```
 *
 * ## Import
 *
 * Load balancer domain_extension can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:slb/domainExtension:DomainExtension example de-abc123456
 * ```
 */
export class DomainExtension extends pulumi.CustomResource {
    /**
     * Get an existing DomainExtension resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainExtensionState, opts?: pulumi.CustomResourceOptions): DomainExtension {
        return new DomainExtension(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:slb/domainExtension:DomainExtension';

    /**
     * Returns true if the given object is an instance of DomainExtension.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainExtension {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainExtension.__pulumiType;
    }

    /**
     * Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
     */
    public readonly deleteProtectionValidation!: pulumi.Output<boolean | undefined>;
    /**
     * The domain name,
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535.
     */
    public readonly frontendPort!: pulumi.Output<number>;
    /**
     * The ID of the SLB instance.
     */
    public readonly loadBalancerId!: pulumi.Output<string>;
    /**
     * The ID of the certificate used by the domain name.
     */
    public readonly serverCertificateId!: pulumi.Output<string>;

    /**
     * Create a DomainExtension resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainExtensionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainExtensionArgs | DomainExtensionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DomainExtensionState | undefined;
            inputs["deleteProtectionValidation"] = state ? state.deleteProtectionValidation : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["frontendPort"] = state ? state.frontendPort : undefined;
            inputs["loadBalancerId"] = state ? state.loadBalancerId : undefined;
            inputs["serverCertificateId"] = state ? state.serverCertificateId : undefined;
        } else {
            const args = argsOrState as DomainExtensionArgs | undefined;
            if ((!args || args.domain === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.frontendPort === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'frontendPort'");
            }
            if ((!args || args.loadBalancerId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'loadBalancerId'");
            }
            if ((!args || args.serverCertificateId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'serverCertificateId'");
            }
            inputs["deleteProtectionValidation"] = args ? args.deleteProtectionValidation : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["frontendPort"] = args ? args.frontendPort : undefined;
            inputs["loadBalancerId"] = args ? args.loadBalancerId : undefined;
            inputs["serverCertificateId"] = args ? args.serverCertificateId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DomainExtension.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainExtension resources.
 */
export interface DomainExtensionState {
    /**
     * Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
     */
    readonly deleteProtectionValidation?: pulumi.Input<boolean>;
    /**
     * The domain name,
     */
    readonly domain?: pulumi.Input<string>;
    /**
     * The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535.
     */
    readonly frontendPort?: pulumi.Input<number>;
    /**
     * The ID of the SLB instance.
     */
    readonly loadBalancerId?: pulumi.Input<string>;
    /**
     * The ID of the certificate used by the domain name.
     */
    readonly serverCertificateId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DomainExtension resource.
 */
export interface DomainExtensionArgs {
    /**
     * Checking DeleteProtection of SLB instance before deleting. If true, this resource will not be deleted when its SLB instance enabled DeleteProtection. Default to false.
     */
    readonly deleteProtectionValidation?: pulumi.Input<boolean>;
    /**
     * The domain name,
     */
    readonly domain: pulumi.Input<string>;
    /**
     * The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535.
     */
    readonly frontendPort: pulumi.Input<number>;
    /**
     * The ID of the SLB instance.
     */
    readonly loadBalancerId: pulumi.Input<string>;
    /**
     * The ID of the certificate used by the domain name.
     */
    readonly serverCertificateId: pulumi.Input<string>;
}
